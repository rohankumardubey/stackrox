// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

const (
	baseTable  = "secrets"
	countStmt  = "SELECT COUNT(*) FROM secrets"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM secrets WHERE id = $1)"

	getStmt     = "SELECT serialized FROM secrets WHERE id = $1"
	deleteStmt  = "DELETE FROM secrets WHERE id = $1"
	walkStmt    = "SELECT serialized FROM secrets"
	getIDsStmt  = "SELECT id FROM secrets"
	getManyStmt = "SELECT serialized FROM secrets WHERE id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM secrets WHERE id = ANY($1::text[])"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.Secret)(nil)), baseTable)
	log    = logging.LoggerForModule()
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.Secret, bool, error)
	Upsert(ctx context.Context, obj *storage.Secret) error
	UpsertMany(ctx context.Context, objs []*storage.Secret) error
	Delete(ctx context.Context, id string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.Secret, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.Secret) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableSecrets(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists secrets (
    id varchar,
    name varchar,
    clusterid varchar,
    clustername varchar,
    namespace varchar,
    type varchar,
    labels jsonb,
    annotations jsonb,
    createdat timestamp,
    relationship_id varchar,
    serialized bytea,
    PRIMARY KEY(id)
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableSecretsFiles(ctx, db)
	createTableSecretsContainerRelationships(ctx, db)
	createTableSecretsDeploymentRelationships(ctx, db)
}

func createTableSecretsFiles(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists secrets_Files (
    secretid varchar,
    idx integer,
    name varchar,
    type integer,
    cert_subject_commonname varchar,
    cert_subject_country varchar,
    cert_subject_organization varchar,
    cert_subject_organizationunit varchar,
    cert_subject_locality varchar,
    cert_subject_province varchar,
    cert_subject_streetaddress varchar,
    cert_subject_postalcode varchar,
    cert_subject_names text[],
    cert_issuer_commonname varchar,
    cert_issuer_country varchar,
    cert_issuer_organization varchar,
    cert_issuer_organizationunit varchar,
    cert_issuer_locality varchar,
    cert_issuer_province varchar,
    cert_issuer_streetaddress varchar,
    cert_issuer_postalcode varchar,
    cert_issuer_names text[],
    cert_sans text[],
    cert_startdate timestamp,
    cert_enddate timestamp,
    cert_algorithm varchar,
    PRIMARY KEY(secretid, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (secretid) REFERENCES secrets(id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists secretsFiles_idx on secrets_Files using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableSecretsFilesRegistries(ctx, db)
}

func createTableSecretsFilesRegistries(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists secrets_Files_Registries (
    secretid varchar,
    secretdatafileidx integer,
    idx integer,
    name varchar,
    username varchar,
    PRIMARY KEY(secretid, secretdatafileidx, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (secretid, secretdatafileidx) REFERENCES secrets_Files(secretid, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists secretsFilesRegistries_idx on secrets_Files_Registries using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func createTableSecretsContainerRelationships(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists secrets_ContainerRelationships (
    secretid varchar,
    idx integer,
    id varchar,
    path varchar,
    PRIMARY KEY(secretid, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (secretid) REFERENCES secrets(id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists secretsContainerRelationships_idx on secrets_ContainerRelationships using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func createTableSecretsDeploymentRelationships(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists secrets_DeploymentRelationships (
    secretid varchar,
    idx integer,
    id varchar,
    name varchar,
    PRIMARY KEY(secretid, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (secretid) REFERENCES secrets(id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists secretsDeploymentRelationships_idx on secrets_DeploymentRelationships using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoSecrets(ctx context.Context, tx pgx.Tx, obj *storage.Secret) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetName(),
		obj.GetClusterId(),
		obj.GetClusterName(),
		obj.GetNamespace(),
		obj.GetType(),
		obj.GetLabels(),
		obj.GetAnnotations(),
		pgutils.NilOrTime(obj.GetCreatedAt()),
		obj.GetRelationship().GetId(),
		serialized,
	}

	finalStr := "INSERT INTO secrets (id, name, clusterid, clustername, namespace, type, labels, annotations, createdat, relationship_id, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT(id) DO UPDATE SET id = EXCLUDED.id, name = EXCLUDED.name, clusterid = EXCLUDED.clusterid, clustername = EXCLUDED.clustername, namespace = EXCLUDED.namespace, type = EXCLUDED.type, labels = EXCLUDED.labels, annotations = EXCLUDED.annotations, createdat = EXCLUDED.createdat, relationship_id = EXCLUDED.relationship_id, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetFiles() {
		if err := insertIntoSecretsFiles(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_Files where secretid = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetFiles()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetRelationship().GetContainerRelationships() {
		if err := insertIntoSecretsContainerRelationships(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_ContainerRelationships where secretid = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetRelationship().GetContainerRelationships()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetRelationship().GetDeploymentRelationships() {
		if err := insertIntoSecretsDeploymentRelationships(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_DeploymentRelationships where secretid = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetRelationship().GetDeploymentRelationships()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSecretsFiles(ctx context.Context, tx pgx.Tx, obj *storage.SecretDataFile, secretid string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		secretid,
		idx,
		obj.GetName(),
		obj.GetType(),
		obj.GetCert().GetSubject().GetCommonName(),
		obj.GetCert().GetSubject().GetCountry(),
		obj.GetCert().GetSubject().GetOrganization(),
		obj.GetCert().GetSubject().GetOrganizationUnit(),
		obj.GetCert().GetSubject().GetLocality(),
		obj.GetCert().GetSubject().GetProvince(),
		obj.GetCert().GetSubject().GetStreetAddress(),
		obj.GetCert().GetSubject().GetPostalCode(),
		obj.GetCert().GetSubject().GetNames(),
		obj.GetCert().GetIssuer().GetCommonName(),
		obj.GetCert().GetIssuer().GetCountry(),
		obj.GetCert().GetIssuer().GetOrganization(),
		obj.GetCert().GetIssuer().GetOrganizationUnit(),
		obj.GetCert().GetIssuer().GetLocality(),
		obj.GetCert().GetIssuer().GetProvince(),
		obj.GetCert().GetIssuer().GetStreetAddress(),
		obj.GetCert().GetIssuer().GetPostalCode(),
		obj.GetCert().GetIssuer().GetNames(),
		obj.GetCert().GetSans(),
		pgutils.NilOrTime(obj.GetCert().GetStartDate()),
		pgutils.NilOrTime(obj.GetCert().GetEndDate()),
		obj.GetCert().GetAlgorithm(),
	}

	finalStr := "INSERT INTO secrets_Files (secretid, idx, name, type, cert_subject_commonname, cert_subject_country, cert_subject_organization, cert_subject_organizationunit, cert_subject_locality, cert_subject_province, cert_subject_streetaddress, cert_subject_postalcode, cert_subject_names, cert_issuer_commonname, cert_issuer_country, cert_issuer_organization, cert_issuer_organizationunit, cert_issuer_locality, cert_issuer_province, cert_issuer_streetaddress, cert_issuer_postalcode, cert_issuer_names, cert_sans, cert_startdate, cert_enddate, cert_algorithm) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26) ON CONFLICT(secretid, idx) DO UPDATE SET secretid = EXCLUDED.secretid, idx = EXCLUDED.idx, name = EXCLUDED.name, type = EXCLUDED.type, cert_subject_commonname = EXCLUDED.cert_subject_commonname, cert_subject_country = EXCLUDED.cert_subject_country, cert_subject_organization = EXCLUDED.cert_subject_organization, cert_subject_organizationunit = EXCLUDED.cert_subject_organizationunit, cert_subject_locality = EXCLUDED.cert_subject_locality, cert_subject_province = EXCLUDED.cert_subject_province, cert_subject_streetaddress = EXCLUDED.cert_subject_streetaddress, cert_subject_postalcode = EXCLUDED.cert_subject_postalcode, cert_subject_names = EXCLUDED.cert_subject_names, cert_issuer_commonname = EXCLUDED.cert_issuer_commonname, cert_issuer_country = EXCLUDED.cert_issuer_country, cert_issuer_organization = EXCLUDED.cert_issuer_organization, cert_issuer_organizationunit = EXCLUDED.cert_issuer_organizationunit, cert_issuer_locality = EXCLUDED.cert_issuer_locality, cert_issuer_province = EXCLUDED.cert_issuer_province, cert_issuer_streetaddress = EXCLUDED.cert_issuer_streetaddress, cert_issuer_postalcode = EXCLUDED.cert_issuer_postalcode, cert_issuer_names = EXCLUDED.cert_issuer_names, cert_sans = EXCLUDED.cert_sans, cert_startdate = EXCLUDED.cert_startdate, cert_enddate = EXCLUDED.cert_enddate, cert_algorithm = EXCLUDED.cert_algorithm"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetImagePullSecret().GetRegistries() {
		if err := insertIntoSecretsFilesRegistries(ctx, tx, child, secretid, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_Files_Registries where secretid = $1 AND secretdatafileidx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, secretid, idx, len(obj.GetImagePullSecret().GetRegistries()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSecretsFilesRegistries(ctx context.Context, tx pgx.Tx, obj *storage.ImagePullSecret_Registry, secretid string, secretdatafileidx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		secretid,
		secretdatafileidx,
		idx,
		obj.GetName(),
		obj.GetUsername(),
	}

	finalStr := "INSERT INTO secrets_Files_Registries (secretid, secretdatafileidx, idx, name, username) VALUES($1, $2, $3, $4, $5) ON CONFLICT(secretid, secretdatafileidx, idx) DO UPDATE SET secretid = EXCLUDED.secretid, secretdatafileidx = EXCLUDED.secretdatafileidx, idx = EXCLUDED.idx, name = EXCLUDED.name, username = EXCLUDED.username"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoSecretsContainerRelationships(ctx context.Context, tx pgx.Tx, obj *storage.SecretContainerRelationship, secretid string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		secretid,
		idx,
		obj.GetId(),
		obj.GetPath(),
	}

	finalStr := "INSERT INTO secrets_ContainerRelationships (secretid, idx, id, path) VALUES($1, $2, $3, $4) ON CONFLICT(secretid, idx) DO UPDATE SET secretid = EXCLUDED.secretid, idx = EXCLUDED.idx, id = EXCLUDED.id, path = EXCLUDED.path"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoSecretsDeploymentRelationships(ctx context.Context, tx pgx.Tx, obj *storage.SecretDeploymentRelationship, secretid string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		secretid,
		idx,
		obj.GetId(),
		obj.GetName(),
	}

	finalStr := "INSERT INTO secrets_DeploymentRelationships (secretid, idx, id, name) VALUES($1, $2, $3, $4) ON CONFLICT(secretid, idx) DO UPDATE SET secretid = EXCLUDED.secretid, idx = EXCLUDED.idx, id = EXCLUDED.id, name = EXCLUDED.name"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) copyFromSecrets(ctx context.Context, tx pgx.Tx, objs ...*storage.Secret) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"id",

		"name",

		"clusterid",

		"clustername",

		"namespace",

		"type",

		"labels",

		"annotations",

		"createdat",

		"relationship_id",

		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{

			obj.GetId(),

			obj.GetName(),

			obj.GetClusterId(),

			obj.GetClusterName(),

			obj.GetNamespace(),

			obj.GetType(),

			obj.GetLabels(),

			obj.GetAnnotations(),

			pgutils.NilOrTime(obj.GetCreatedAt()),

			obj.GetRelationship().GetId(),

			serialized,
		})

		// Add the id to be deleted.
		deletes = append(deletes, obj.GetId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.Exec(ctx, deleteManyStmt, deletes)
			if err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = nil

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"secrets"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for _, obj := range objs {

		if err = s.copyFromSecretsFiles(ctx, tx, obj.GetId(), obj.GetFiles()...); err != nil {
			return err
		}
		if err = s.copyFromSecretsContainerRelationships(ctx, tx, obj.GetId(), obj.GetRelationship().GetContainerRelationships()...); err != nil {
			return err
		}
		if err = s.copyFromSecretsDeploymentRelationships(ctx, tx, obj.GetId(), obj.GetRelationship().GetDeploymentRelationships()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromSecretsFiles(ctx context.Context, tx pgx.Tx, secretid string, objs ...*storage.SecretDataFile) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"secretid",

		"idx",

		"name",

		"type",

		"cert_subject_commonname",

		"cert_subject_country",

		"cert_subject_organization",

		"cert_subject_organizationunit",

		"cert_subject_locality",

		"cert_subject_province",

		"cert_subject_streetaddress",

		"cert_subject_postalcode",

		"cert_subject_names",

		"cert_issuer_commonname",

		"cert_issuer_country",

		"cert_issuer_organization",

		"cert_issuer_organizationunit",

		"cert_issuer_locality",

		"cert_issuer_province",

		"cert_issuer_streetaddress",

		"cert_issuer_postalcode",

		"cert_issuer_names",

		"cert_sans",

		"cert_startdate",

		"cert_enddate",

		"cert_algorithm",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			secretid,

			idx,

			obj.GetName(),

			obj.GetType(),

			obj.GetCert().GetSubject().GetCommonName(),

			obj.GetCert().GetSubject().GetCountry(),

			obj.GetCert().GetSubject().GetOrganization(),

			obj.GetCert().GetSubject().GetOrganizationUnit(),

			obj.GetCert().GetSubject().GetLocality(),

			obj.GetCert().GetSubject().GetProvince(),

			obj.GetCert().GetSubject().GetStreetAddress(),

			obj.GetCert().GetSubject().GetPostalCode(),

			obj.GetCert().GetSubject().GetNames(),

			obj.GetCert().GetIssuer().GetCommonName(),

			obj.GetCert().GetIssuer().GetCountry(),

			obj.GetCert().GetIssuer().GetOrganization(),

			obj.GetCert().GetIssuer().GetOrganizationUnit(),

			obj.GetCert().GetIssuer().GetLocality(),

			obj.GetCert().GetIssuer().GetProvince(),

			obj.GetCert().GetIssuer().GetStreetAddress(),

			obj.GetCert().GetIssuer().GetPostalCode(),

			obj.GetCert().GetIssuer().GetNames(),

			obj.GetCert().GetSans(),

			pgutils.NilOrTime(obj.GetCert().GetStartDate()),

			pgutils.NilOrTime(obj.GetCert().GetEndDate()),

			obj.GetCert().GetAlgorithm(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"secrets_files"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {

		if err = s.copyFromSecretsFilesRegistries(ctx, tx, secretid, idx, obj.GetImagePullSecret().GetRegistries()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromSecretsFilesRegistries(ctx context.Context, tx pgx.Tx, secretid string, secretdatafileidx int, objs ...*storage.ImagePullSecret_Registry) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"secretid",

		"secretdatafileidx",

		"idx",

		"name",

		"username",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			secretid,

			secretdatafileidx,

			idx,

			obj.GetName(),

			obj.GetUsername(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"secrets_files_registries"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) copyFromSecretsContainerRelationships(ctx context.Context, tx pgx.Tx, secretid string, objs ...*storage.SecretContainerRelationship) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"secretid",

		"idx",

		"id",

		"path",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			secretid,

			idx,

			obj.GetId(),

			obj.GetPath(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"secrets_containerrelationships"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) copyFromSecretsDeploymentRelationships(ctx context.Context, tx pgx.Tx, secretid string, objs ...*storage.SecretDeploymentRelationship) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"secretid",

		"idx",

		"id",

		"name",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			secretid,

			idx,

			obj.GetId(),

			obj.GetName(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"secrets_deploymentrelationships"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	createTableSecrets(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.Secret) error {
	conn, release := s.acquireConn(ctx, ops.Get, "Secret")
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromSecrets(ctx, tx, objs...); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.Secret) error {
	conn, release := s.acquireConn(ctx, ops.Get, "Secret")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoSecrets(ctx, tx, obj); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.Secret) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "Secret")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.Secret) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "Secret")

	if len(objs) < batchAfter {
		return s.upsert(ctx, objs...)
	} else {
		return s.copyFrom(ctx, objs...)
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "Secret")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "Secret")

	row := s.db.QueryRow(ctx, existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string) (*storage.Secret, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "Secret")

	conn, release := s.acquireConn(ctx, ops.Get, "Secret")
	defer release()

	row := conn.QueryRow(ctx, getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.Secret
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(ctx context.Context, id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "Secret")

	conn, release := s.acquireConn(ctx, ops.Remove, "Secret")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.SecretIDs")

	rows, err := s.db.Query(ctx, getIDsStmt)
	if err != nil {
		return nil, pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.Secret, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "Secret")

	conn, release := s.acquireConn(ctx, ops.GetMany, "Secret")
	defer release()

	rows, err := conn.Query(ctx, getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	resultsByID := make(map[string]*storage.Secret)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.Secret{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.Secret, 0, len(resultsByID))
	for i, id := range ids {
		if result, ok := resultsByID[id]; !ok {
			missingIndices = append(missingIndices, i)
		} else {
			elems = append(elems, result)
		}
	}
	return elems, missingIndices, nil
}

// Delete removes the specified IDs from the store
func (s *storeImpl) DeleteMany(ctx context.Context, ids []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "Secret")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "Secret")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.Secret) error) error {
	rows, err := s.db.Query(ctx, walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.Secret
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableSecrets(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS secrets CASCADE")
	dropTableSecretsFiles(ctx, db)
	dropTableSecretsContainerRelationships(ctx, db)
	dropTableSecretsDeploymentRelationships(ctx, db)

}

func dropTableSecretsFiles(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS secrets_Files CASCADE")
	dropTableSecretsFilesRegistries(ctx, db)

}

func dropTableSecretsFilesRegistries(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS secrets_Files_Registries CASCADE")

}

func dropTableSecretsContainerRelationships(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS secrets_ContainerRelationships CASCADE")

}

func dropTableSecretsDeploymentRelationships(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS secrets_DeploymentRelationships CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableSecrets(ctx, db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}
