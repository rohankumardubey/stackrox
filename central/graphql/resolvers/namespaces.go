package resolvers

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/stackrox/rox/central/graphql/resolvers/loaders"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/central/namespace"
	"github.com/stackrox/rox/central/policy/matcher"
	riskDS "github.com/stackrox/rox/central/risk/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	pkgMetrics "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/scoped"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/utils"
)

func init() {
	schema := getBuilder()
	utils.Must(
		schema.AddQuery("namespaces(query: String, pagination: Pagination): [Namespace!]!"),
		schema.AddQuery("namespace(id: ID!): Namespace"),
		schema.AddQuery("namespaceByClusterIDAndName(clusterID: ID!, name: String!): Namespace"),
		schema.AddQuery("namespaceCount(query: String): Int!"),
		schema.AddExtraResolver("Namespace", "complianceResults(query: String): [ControlResult!]!"),
		schema.AddExtraResolver("Namespace", `subjects(query: String, pagination: Pagination): [Subject!]!`),
		schema.AddExtraResolver("Namespace", `subjectCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `serviceAccountCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `serviceAccounts(query: String, pagination: Pagination): [ServiceAccount!]!`),
		schema.AddExtraResolver("Namespace", `k8sRoleCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `k8sRoles(query: String, pagination: Pagination): [K8SRole!]!`),
		schema.AddExtraResolver("Namespace", `policyCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `policyStatus(query: String): PolicyStatus!`),
		schema.AddExtraResolver("Namespace", `policyStatusOnly(query: String): String!`),
		schema.AddExtraResolver("Namespace", `policies(query: String, pagination: Pagination): [Policy!]!`),
		schema.AddExtraResolver("Namespace", `failingPolicyCounter(query: String): PolicyCounter`),
		schema.AddExtraResolver("Namespace", `images(query: String, pagination: Pagination): [Image!]!`),
		schema.AddExtraResolver("Namespace", `imageCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `components(query: String, pagination: Pagination): [EmbeddedImageScanComponent!]!`),
		schema.AddExtraResolver("Namespace", `componentCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `vulns(query: String, scopeQuery: String, pagination: Pagination): [EmbeddedVulnerability!]!`),
		schema.AddExtraResolver("Namespace", `vulnCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `vulnCounter(query: String): VulnerabilityCounter!`),
		schema.AddExtraResolver("Namespace", `secrets(query: String, pagination: Pagination): [Secret!]!`),
		schema.AddExtraResolver("Namespace", `deployments(query: String, pagination: Pagination): [Deployment!]!`),
		schema.AddExtraResolver("Namespace", "cluster: Cluster!"),
		schema.AddExtraResolver("Namespace", `secretCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `deploymentCount(query: String): Int!`),
		schema.AddExtraResolver("Namespace", `risk: Risk`),
		schema.AddExtraResolver("Namespace", "latestViolation(query: String): Time"),
		schema.AddExtraResolver("Namespace", `unusedVarSink(query: String): Int`),
		schema.AddExtraResolver("Namespace", "plottedVulns(query: String): PlottedVulnerabilities!"),
	)
}

func (resolver *namespaceResolver) getNamespaceIDRawQuery() string {
	return search.NewQueryBuilder().
		AddExactMatches(search.NamespaceID, resolver.data.GetMetadata().GetId()).
		Query()
}

func (resolver *namespaceResolver) getClusterNamespaceRawQuery() string {
	return search.NewQueryBuilder().
		AddExactMatches(search.ClusterID, resolver.data.GetMetadata().GetClusterId()).
		AddExactMatches(search.Namespace, resolver.data.Metadata.GetName()).
		Query()
}

func (resolver *namespaceResolver) getClusterNamespaceQuery() *v1.Query {
	return search.NewQueryBuilder().
		AddExactMatches(search.ClusterID, resolver.data.GetMetadata().GetClusterId()).
		AddExactMatches(search.Namespace, resolver.data.Metadata.GetName()).
		ProtoQuery()
}

func (resolver *namespaceResolver) getNamespaceConjunctionQuery(args RawQuery) (*v1.Query, error) {
	q1 := resolver.getClusterNamespaceQuery()
	if args.String() == "" {
		return q1, nil
	}

	q2, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	return search.AddAsConjunction(q2, q1)
}

// Namespace returns a GraphQL resolver for the given namespace.
func (resolver *Resolver) Namespace(ctx context.Context, args struct{ graphql.ID }) (*namespaceResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "Namespace")
	if err := readNamespaces(ctx); err != nil {
		return nil, err
	}
	return resolver.wrapNamespace(namespace.ResolveByID(ctx, string(args.ID), resolver.NamespaceDataStore, resolver.DeploymentDataStore, resolver.SecretsDataStore, resolver.NetworkPoliciesStore))
}

// Namespaces returns GraphQL resolvers for all namespaces based on an optional query.
func (resolver *Resolver) Namespaces(ctx context.Context, args PaginatedQuery) ([]*namespaceResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "Namespaces")
	if err := readNamespaces(ctx); err != nil {
		return nil, err
	}
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	return resolver.wrapNamespaces(namespace.ResolveByQuery(ctx, query, resolver.NamespaceDataStore, resolver.DeploymentDataStore, resolver.SecretsDataStore, resolver.NetworkPoliciesStore))
}

type clusterIDAndNameQuery struct {
	ClusterID graphql.ID
	Name      string
}

// NamespaceByClusterIDAndName returns a GraphQL resolver for the (unique) namespace specified by this query.
func (resolver *Resolver) NamespaceByClusterIDAndName(ctx context.Context, args clusterIDAndNameQuery) (*namespaceResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "NamespaceByClusterIDAndName")
	if err := readNamespaces(ctx); err != nil {
		return nil, err
	}

	return resolver.wrapNamespace(namespace.ResolveByClusterIDAndName(ctx, string(args.ClusterID), args.Name, resolver.NamespaceDataStore, resolver.DeploymentDataStore, resolver.SecretsDataStore, resolver.NetworkPoliciesStore))
}

// NamespaceCount returns count of all clusters across infrastructure
func (resolver *Resolver) NamespaceCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "NamespaceCount")
	if err := readNamespaces(ctx); err != nil {
		return 0, err
	}
	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}
	results, err := resolver.NamespaceDataStore.Search(ctx, q)
	if err != nil {
		return 0, err
	}
	return int32(len(results)), nil
}

func (resolver *namespaceResolver) ComplianceResults(ctx context.Context, args RawQuery) ([]*controlResultResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "ComplianceResults")
	if err := readCompliance(ctx); err != nil {
		return nil, err
	}

	runResults, err := resolver.root.ComplianceAggregator.GetResultsWithEvidence(ctx, args.String())
	if err != nil {
		return nil, err
	}
	output := newBulkControlResults()
	nsID := resolver.data.GetMetadata().GetId()
	output.addDeploymentData(resolver.root, runResults, func(d *storage.Deployment, _ *v1.ComplianceControl) bool {
		return d.GetNamespaceId() == nsID
	})

	return *output, nil
}

// SubjectCount returns the count of Subjects which have any permission on this namespace or the cluster it belongs to
func (resolver *namespaceResolver) SubjectCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "SubjectCount")
	if err := readK8sSubjects(ctx); err != nil {
		return 0, err
	}
	if err := readK8sRoleBindings(ctx); err != nil {
		return 0, err
	}

	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}

	subjects, err := resolver.getSubjects(ctx, q)
	if err != nil {
		return 0, err
	}
	return int32(len(subjects)), nil
}

// Subjects returns the Subjects which have any permission in namespace or cluster wide
func (resolver *namespaceResolver) Subjects(ctx context.Context, args PaginatedQuery) ([]*subjectResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Subjects")
	if err := readK8sSubjects(ctx); err != nil {
		return nil, err
	}
	if err := readK8sRoleBindings(ctx); err != nil {
		return nil, err
	}
	var resolvers []*subjectResolver
	baseQuery, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	pagination := baseQuery.GetPagination()
	baseQuery.Pagination = nil

	subjects, err := resolver.getSubjects(ctx, baseQuery)
	if err != nil {
		return nil, err
	}
	for _, subject := range subjects {
		resolvers = append(resolvers, &subjectResolver{ctx, resolver.root, subject})
	}

	paginatedResolvers, err := paginationWrapper{
		pv: pagination,
	}.paginate(resolvers, nil)
	return paginatedResolvers.([]*subjectResolver), err
}

// ServiceAccountCount returns the count of ServiceAccounts which have any permission on this cluster namespace
func (resolver *namespaceResolver) ServiceAccountCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "ServiceAccountCount")
	if err := readServiceAccounts(ctx); err != nil {
		return 0, err
	}

	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}

	q, err = search.AddAsConjunction(resolver.getClusterNamespaceQuery(), q)
	if err != nil {
		return 0, err
	}

	results, err := resolver.root.ServiceAccountsDataStore.Search(ctx, q)
	if err != nil {
		return 0, err
	}
	return int32(len(results)), nil
}

// ServiceAccounts returns the ServiceAccounts which have any permission on this cluster namespace
func (resolver *namespaceResolver) ServiceAccounts(ctx context.Context, args PaginatedQuery) ([]*serviceAccountResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "ServiceAccounts")
	if err := readServiceAccounts(ctx); err != nil {
		return nil, err
	}

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.ServiceAccounts(ctx, PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

// K8sRoleCount returns count of K8s roles in this cluster namespace
func (resolver *namespaceResolver) K8sRoleCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "K8sRoleCount")
	if err := readK8sRoles(ctx); err != nil {
		return 0, err
	}

	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}

	q, err = search.AddAsConjunction(resolver.getClusterNamespaceQuery(), q)
	if err != nil {
		return 0, err
	}

	results, err := resolver.root.K8sRoleStore.Search(ctx, q)
	if err != nil {
		return 0, err
	}
	return int32(len(results)), nil
}

// K8sRoles returns count of K8s roles in this cluster namespace
func (resolver *namespaceResolver) K8sRoles(ctx context.Context, args PaginatedQuery) ([]*k8SRoleResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "K8sRoles")
	if err := readK8sRoles(ctx); err != nil {
		return nil, err
	}

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.K8sRoles(ctx, PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

func (resolver *namespaceResolver) Images(ctx context.Context, args PaginatedQuery) ([]*imageResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Images")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.Images(ctx, PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

func (resolver *namespaceResolver) ImageCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "ImageCount")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.ImageCount(ctx, RawQuery{Query: &query})
}

func (resolver *namespaceResolver) getApplicablePolicies(ctx context.Context, q *v1.Query) ([]*storage.Policy, error) {
	policyLoader, err := loaders.GetPolicyLoader(ctx)
	if err != nil {
		return nil, err
	}

	policies, err := policyLoader.FromQuery(ctx, q)
	if err != nil {
		return nil, err
	}

	applicable, _ := matcher.NewNamespaceMatcher(resolver.data.Metadata).FilterApplicablePolicies(policies)
	return applicable, nil
}

// PolicyCount returns count of policies applicable to this namespace
func (resolver *namespaceResolver) PolicyCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "PolicyCount")

	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}

	policies, err := resolver.getApplicablePolicies(ctx, query)
	if err != nil {
		return 0, err
	}

	return int32(len(policies)), nil
}

// Policies returns all the policies applicable to this namespace
func (resolver *namespaceResolver) Policies(ctx context.Context, args PaginatedQuery) ([]*policyResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Policies")

	if err := readPolicies(ctx); err != nil {
		return nil, err
	}

	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	// remove pagination from query since we want to paginate the final result
	pagination := q.GetPagination()
	q.Pagination = &v1.QueryPagination{
		SortOptions: pagination.GetSortOptions(),
	}

	policyResolvers, err := resolver.root.wrapPolicies(resolver.getApplicablePolicies(ctx, q))
	if err != nil {
		return nil, err
	}
	for _, policyResolver := range policyResolvers {
		policyResolver.ctx = scoped.Context(ctx, scoped.Scope{
			Level: v1.SearchCategory_NAMESPACES,
			ID:    resolver.data.GetMetadata().GetId(),
		})
	}

	resolvers, err := paginationWrapper{
		pv: pagination,
	}.paginate(policyResolvers, nil)
	return resolvers.([]*policyResolver), err
}

// FailingPolicyCounter returns a policy counter for all the failed policies.
func (resolver *namespaceResolver) FailingPolicyCounter(ctx context.Context, args RawQuery) (*PolicyCounterResolver, error) {
	if err := readAlerts(ctx); err != nil {
		return nil, err
	}

	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	q, err = search.AddAsConjunction(q, resolver.getClusterNamespaceQuery())
	if err != nil {
		return nil, err
	}

	alerts, err := resolver.root.ViolationsDataStore.SearchListAlerts(ctx, q)
	if err != nil {
		return nil, nil
	}
	return mapListAlertsToPolicySeverityCount(alerts), nil
}

// PolicyStatus returns true if there is no policy violation for this namespace
func (resolver *namespaceResolver) PolicyStatus(ctx context.Context, args RawQuery) (*policyStatusResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "PolicyStatus")

	if err := readAlerts(ctx); err != nil {
		return nil, err
	}

	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}

	alerts, err := resolver.getActiveDeployAlerts(ctx, query)
	if err != nil {
		return nil, err
	}

	scopedCtx := scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_NAMESPACES,
		ID:    resolver.data.GetMetadata().GetId(),
	})

	if len(alerts) == 0 {
		return &policyStatusResolver{scopedCtx, resolver.root, "pass", nil}, nil
	}

	policyIDs := set.NewStringSet()
	for _, alert := range alerts {
		policyIDs.Add(alert.GetPolicy().GetId())
	}

	return &policyStatusResolver{scopedCtx, resolver.root, "fail", policyIDs.AsSlice()}, nil
}

// PolicyStatusOnly returns 'fail' if there are policy violations for this namespace
func (resolver *namespaceResolver) PolicyStatusOnly(ctx context.Context, args RawQuery) (string, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "PolicyStatusOnly")
	if err := readAlerts(ctx); err != nil {
		return "", err
	}

	q, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return "", err
	}

	results, err := resolver.root.ViolationsDataStore.Search(ctx,
		search.ConjunctionQuery(q,
			search.NewQueryBuilder().AddExactMatches(search.ClusterID, resolver.data.GetMetadata().GetClusterId()).
				AddExactMatches(search.Namespace, resolver.data.GetMetadata().GetName()).
				AddStrings(search.ViolationState, storage.ViolationState_ACTIVE.String()).ProtoQuery()))
	if err != nil {
		return "", err
	}

	if len(results) > 0 {
		return "fail", nil
	}
	return "pass", nil
}

func (resolver *namespaceResolver) getActiveDeployAlerts(ctx context.Context, q *v1.Query) ([]*storage.ListAlert, error) {
	if err := readAlerts(ctx); err != nil {
		return nil, err
	}

	namespace := resolver.data

	return resolver.root.ViolationsDataStore.SearchListAlerts(ctx,
		search.ConjunctionQuery(q,
			search.NewQueryBuilder().AddExactMatches(search.ClusterID, namespace.GetMetadata().GetClusterId()).
				AddExactMatches(search.Namespace, namespace.GetMetadata().GetName()).
				AddStrings(search.ViolationState, storage.ViolationState_ACTIVE.String()).
				AddStrings(search.LifecycleStage, storage.LifecycleStage_DEPLOY.String()).ProtoQuery()))
}

func (resolver *namespaceResolver) Components(ctx context.Context, args PaginatedQuery) ([]ComponentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Components")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getNamespaceIDRawQuery())

	return resolver.root.Components(scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_NAMESPACES,
		ID:    resolver.data.GetMetadata().GetId(),
	}), PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

func (resolver *namespaceResolver) ComponentCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "ComponentCount")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getNamespaceIDRawQuery())

	return resolver.root.ComponentCount(scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_NAMESPACES,
		ID:    resolver.data.GetMetadata().GetId(),
	}), RawQuery{Query: &query})
}

func (resolver *namespaceResolver) Vulns(ctx context.Context, args PaginatedQuery) ([]VulnerabilityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Vulns")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getNamespaceIDRawQuery())

	return resolver.root.Vulnerabilities(scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_NAMESPACES,
		ID:    resolver.data.GetMetadata().GetId(),
	}), PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

func (resolver *namespaceResolver) VulnCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "VulnCount")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getNamespaceIDRawQuery())

	return resolver.root.VulnerabilityCount(scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_NAMESPACES,
		ID:    resolver.data.GetMetadata().GetId(),
	}), RawQuery{Query: &query})
}

func (resolver *namespaceResolver) VulnCounter(ctx context.Context, args RawQuery) (*VulnerabilityCounterResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "VulnCounter")

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getNamespaceIDRawQuery())

	return resolver.root.VulnCounter(scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_NAMESPACES,
		ID:    resolver.data.GetMetadata().GetId(),
	}), RawQuery{Query: &query})
}

func (resolver *namespaceResolver) Secrets(ctx context.Context, args PaginatedQuery) ([]*secretResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Secrets")
	if err := readSecrets(ctx); err != nil {
		return nil, err
	}
	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.Secrets(ctx, PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

func (resolver *namespaceResolver) Deployments(ctx context.Context, args PaginatedQuery) ([]*deploymentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Deployments")
	if err := readDeployments(ctx); err != nil {
		return nil, err
	}

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.Deployments(ctx, PaginatedQuery{Query: &query, Pagination: args.Pagination})
}

func (resolver *namespaceResolver) Cluster(ctx context.Context) (*clusterResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Cluster")
	if err := readClusters(ctx); err != nil {
		return nil, err
	}
	return resolver.root.wrapCluster(resolver.root.ClusterDataStore.GetCluster(ctx, resolver.data.GetMetadata().GetClusterId()))
}

func (resolver *namespaceResolver) SecretCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "SecretCount")
	if err := readSecrets(ctx); err != nil {
		return 0, err
	}

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.SecretCount(ctx, RawQuery{Query: &query})
}

func (resolver *namespaceResolver) DeploymentCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "DeploymentCount")
	if err := readDeployments(ctx); err != nil {
		return 0, err
	}

	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())

	return resolver.root.DeploymentCount(ctx, RawQuery{Query: &query})
}

func (resolver *namespaceResolver) Risk(ctx context.Context) (*riskResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Risk")
	if err := readRisks(ctx); err != nil {
		return nil, err
	}
	return resolver.root.wrapRisk(resolver.getNamespaceRisk(ctx))
}

func (resolver *namespaceResolver) getNamespaceRisk(ctx context.Context) (*storage.Risk, bool, error) {
	ns := resolver.data

	riskQuery := search.NewQueryBuilder().
		AddExactMatches(search.Namespace, ns.GetMetadata().GetName()).
		AddExactMatches(search.ClusterID, ns.GetMetadata().GetClusterId()).
		AddExactMatches(search.RiskSubjectType, storage.RiskSubjectType_DEPLOYMENT.String()).
		ProtoQuery()

	risks, err := resolver.root.RiskDataStore.SearchRawRisks(ctx, riskQuery)
	if err != nil {
		return nil, false, err
	}

	risks = filterDeploymentRisksOnScope(ctx, risks...)
	scrubRiskFactors(risks...)
	aggregateRiskScore := getAggregateRiskScore(risks...)
	if aggregateRiskScore == float32(0.0) {
		return nil, false, nil
	}

	risk := &storage.Risk{
		Score: aggregateRiskScore,
		Subject: &storage.RiskSubject{
			Id:        ns.GetMetadata().GetId(),
			Namespace: ns.GetMetadata().GetName(),
			ClusterId: ns.GetMetadata().GetClusterId(),
			Type:      storage.RiskSubjectType_NAMESPACE,
		},
	}

	id, err := riskDS.GetID(risk.GetSubject().GetId(), risk.GetSubject().GetType())
	if err != nil {
		return nil, false, err
	}
	risk.Id = id

	return risk, true, nil
}

func (resolver *namespaceResolver) LatestViolation(ctx context.Context, args RawQuery) (*graphql.Time, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Namespaces, "Latest Violation")

	q, err := resolver.getNamespaceConjunctionQuery(args)
	if err != nil {
		return nil, nil
	}

	return getLatestViolationTime(ctx, resolver.root, q)
}

func (resolver *namespaceResolver) PlottedVulns(ctx context.Context, args PaginatedQuery) (*PlottedVulnerabilitiesResolver, error) {
	query := search.AddRawQueriesAsConjunction(args.String(), resolver.getClusterNamespaceRawQuery())
	return newPlottedVulnerabilitiesResolver(ctx, resolver.root, RawQuery{Query: &query})
}

func (resolver *namespaceResolver) UnusedVarSink(ctx context.Context, args RawQuery) *int32 {
	return nil
}
