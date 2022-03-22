// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"encoding/json"
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

var (
	log = logging.LoggerForModule()
)

const (
	baseTable  = "networkbaseline"
	countStmt  = "SELECT COUNT(*) FROM networkbaseline"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM networkbaseline WHERE DeploymentId = $1)"

	getStmt           = "SELECT serialized FROM networkbaseline WHERE DeploymentId = $1"
	deleteStmt        = "DELETE FROM networkbaseline WHERE DeploymentId = $1"
	walkStmt          = "SELECT serialized FROM networkbaseline"
	getWithRollupStmt = "select row_to_json((select record from (select table0.DeploymentId as DeploymentId, table0.ClusterId as ClusterId, table0.Namespace as Namespace, table0.ObservationPeriodEnd as ObservationPeriodEnd, table0.Locked as Locked, table0.DeploymentName as DeploymentName, to_json(join0)->'array' as join0, to_json(join1)->'array' as join1 from networkbaseline table0 left join lateral (select array(select json_build_object('idx', table1.idx, 'Entity_Info_Type', table1.Entity_Info_Type, 'Entity_Info_Id', table1.Entity_Info_Id, 'Entity_Info_Deployment_Name', table1.Entity_Info_Deployment_Name, 'Entity_Info_Deployment_Namespace', table1.Entity_Info_Deployment_Namespace, 'Entity_Info_Deployment_Cluster', table1.Entity_Info_Deployment_Cluster, 'Entity_Info_ExternalSource_Name', table1.Entity_Info_ExternalSource_Name, 'Entity_Info_ExternalSource_Cidr', table1.Entity_Info_ExternalSource_Cidr, 'Entity_Info_ExternalSource_Default', table1.Entity_Info_ExternalSource_Default, 'Entity_Scope_ClusterId', table1.Entity_Scope_ClusterId, 'join0', to_json(join0)->'array', 'join1', to_json(join1)->'array') from networkbaseline_Peers table1 left join lateral (select array(select json_build_object('idx', table2.idx, 'Port', table2.Port, 'L4Protocol', table2.L4Protocol) from networkbaseline_Peers_ListenPorts table2 where (table0.DeploymentId = table2.networkbaseline_DeploymentId and table1.idx = table2.networkbaseline_Peers_idx))) join0 on true left join lateral (select array(select json_build_object('idx', table3.idx, 'Ingress', table3.Ingress, 'Port', table3.Port, 'Protocol', table3.Protocol) from networkbaseline_Peers_Properties table3 where (table0.DeploymentId = table3.networkbaseline_DeploymentId and table1.idx = table3.networkbaseline_Peers_idx))) join1 on true where (table0.DeploymentId = table1.networkbaseline_DeploymentId))) join0 on true left join lateral (select array(select json_build_object('idx', table4.idx, 'Entity_Info_Type', table4.Entity_Info_Type, 'Entity_Info_Id', table4.Entity_Info_Id, 'Entity_Info_Deployment_Name', table4.Entity_Info_Deployment_Name, 'Entity_Info_Deployment_Namespace', table4.Entity_Info_Deployment_Namespace, 'Entity_Info_Deployment_Cluster', table4.Entity_Info_Deployment_Cluster, 'Entity_Info_ExternalSource_Name', table4.Entity_Info_ExternalSource_Name, 'Entity_Info_ExternalSource_Cidr', table4.Entity_Info_ExternalSource_Cidr, 'Entity_Info_ExternalSource_Default', table4.Entity_Info_ExternalSource_Default, 'Entity_Scope_ClusterId', table4.Entity_Scope_ClusterId, 'join0', to_json(join0)->'array', 'join1', to_json(join1)->'array') from networkbaseline_ForbiddenPeers table4 left join lateral (select array(select json_build_object('idx', table5.idx, 'Port', table5.Port, 'L4Protocol', table5.L4Protocol) from networkbaseline_ForbiddenPeers_ListenPorts table5 where (table0.DeploymentId = table5.networkbaseline_DeploymentId and table4.idx = table5.networkbaseline_ForbiddenPeers_idx))) join0 on true left join lateral (select array(select json_build_object('idx', table6.idx, 'Ingress', table6.Ingress, 'Port', table6.Port, 'Protocol', table6.Protocol) from networkbaseline_ForbiddenPeers_Properties table6 where (table0.DeploymentId = table6.networkbaseline_DeploymentId and table4.idx = table6.networkbaseline_ForbiddenPeers_idx))) join1 on true where (table0.DeploymentId = table4.networkbaseline_DeploymentId))) join1 on true where (table0.DeploymentId = $1)) record ))"
	getIDsStmt        = "SELECT DeploymentId FROM networkbaseline"
	getManyStmt       = "SELECT serialized FROM networkbaseline WHERE DeploymentId = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM networkbaseline WHERE DeploymentId = ANY($1::text[])"
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.NetworkBaseline)(nil)), baseTable)
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, deploymentId string) (bool, error)
	Get(ctx context.Context, deploymentId string) (*storage.NetworkBaseline, bool, error)
	Upsert(ctx context.Context, obj *storage.NetworkBaseline) error
	UpsertMany(ctx context.Context, objs []*storage.NetworkBaseline) error
	Delete(ctx context.Context, deploymentId string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.NetworkBaseline, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.NetworkBaseline) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableNetworkbaseline(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists networkbaseline (
    DeploymentId varchar,
    ClusterId varchar,
    Namespace varchar,
    ObservationPeriodEnd timestamp,
    Locked bool,
    DeploymentName varchar,
    serialized bytea,
    PRIMARY KEY(DeploymentId)
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

	createTableNetworkbaselinePeers(ctx, db)
	createTableNetworkbaselineForbiddenPeers(ctx, db)
}

func createTableNetworkbaselinePeers(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists networkbaseline_Peers (
    networkbaseline_DeploymentId varchar,
    idx integer,
    Entity_Info_Type integer,
    Entity_Info_Id varchar,
    Entity_Info_Deployment_Name varchar,
    Entity_Info_Deployment_Namespace varchar,
    Entity_Info_Deployment_Cluster varchar,
    Entity_Info_ExternalSource_Name varchar,
    Entity_Info_ExternalSource_Cidr varchar,
    Entity_Info_ExternalSource_Default bool,
    Entity_Scope_ClusterId varchar,
    PRIMARY KEY(networkbaseline_DeploymentId, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (networkbaseline_DeploymentId) REFERENCES networkbaseline(DeploymentId) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists networkbaselinePeers_idx on networkbaseline_Peers using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableNetworkbaselinePeersListenPorts(ctx, db)
	createTableNetworkbaselinePeersProperties(ctx, db)
}

func createTableNetworkbaselinePeersListenPorts(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists networkbaseline_Peers_ListenPorts (
    networkbaseline_DeploymentId varchar,
    networkbaseline_Peers_idx integer,
    idx integer,
    Port integer,
    L4Protocol integer,
    PRIMARY KEY(networkbaseline_DeploymentId, networkbaseline_Peers_idx, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (networkbaseline_DeploymentId, networkbaseline_Peers_idx) REFERENCES networkbaseline_Peers(networkbaseline_DeploymentId, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists networkbaselinePeersListenPorts_idx on networkbaseline_Peers_ListenPorts using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func createTableNetworkbaselinePeersProperties(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists networkbaseline_Peers_Properties (
    networkbaseline_DeploymentId varchar,
    networkbaseline_Peers_idx integer,
    idx integer,
    Ingress bool,
    Port integer,
    Protocol integer,
    PRIMARY KEY(networkbaseline_DeploymentId, networkbaseline_Peers_idx, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (networkbaseline_DeploymentId, networkbaseline_Peers_idx) REFERENCES networkbaseline_Peers(networkbaseline_DeploymentId, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists networkbaselinePeersProperties_idx on networkbaseline_Peers_Properties using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func createTableNetworkbaselineForbiddenPeers(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists networkbaseline_ForbiddenPeers (
    networkbaseline_DeploymentId varchar,
    idx integer,
    Entity_Info_Type integer,
    Entity_Info_Id varchar,
    Entity_Info_Deployment_Name varchar,
    Entity_Info_Deployment_Namespace varchar,
    Entity_Info_Deployment_Cluster varchar,
    Entity_Info_ExternalSource_Name varchar,
    Entity_Info_ExternalSource_Cidr varchar,
    Entity_Info_ExternalSource_Default bool,
    Entity_Scope_ClusterId varchar,
    PRIMARY KEY(networkbaseline_DeploymentId, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (networkbaseline_DeploymentId) REFERENCES networkbaseline(DeploymentId) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists networkbaselineForbiddenPeers_idx on networkbaseline_ForbiddenPeers using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableNetworkbaselineForbiddenPeersListenPorts(ctx, db)
	createTableNetworkbaselineForbiddenPeersProperties(ctx, db)
}

func createTableNetworkbaselineForbiddenPeersListenPorts(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists networkbaseline_ForbiddenPeers_ListenPorts (
    networkbaseline_DeploymentId varchar,
    networkbaseline_ForbiddenPeers_idx integer,
    idx integer,
    Port integer,
    L4Protocol integer,
    PRIMARY KEY(networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx) REFERENCES networkbaseline_ForbiddenPeers(networkbaseline_DeploymentId, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists networkbaselineForbiddenPeersListenPorts_idx on networkbaseline_ForbiddenPeers_ListenPorts using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func createTableNetworkbaselineForbiddenPeersProperties(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists networkbaseline_ForbiddenPeers_Properties (
    networkbaseline_DeploymentId varchar,
    networkbaseline_ForbiddenPeers_idx integer,
    idx integer,
    Ingress bool,
    Port integer,
    Protocol integer,
    PRIMARY KEY(networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx) REFERENCES networkbaseline_ForbiddenPeers(networkbaseline_DeploymentId, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists networkbaselineForbiddenPeersProperties_idx on networkbaseline_ForbiddenPeers_Properties using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoNetworkbaseline(ctx context.Context, tx pgx.Tx, obj *storage.NetworkBaseline) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetDeploymentId(),
		obj.GetClusterId(),
		obj.GetNamespace(),
		pgutils.NilOrStringTimestamp(obj.GetObservationPeriodEnd()),
		obj.GetLocked(),
		obj.GetDeploymentName(),
		serialized,
	}

	finalStr := "INSERT INTO networkbaseline (DeploymentId, ClusterId, Namespace, ObservationPeriodEnd, Locked, DeploymentName, serialized) VALUES($1, $2, $3, $4, $5, $6, $7) ON CONFLICT(DeploymentId) DO UPDATE SET DeploymentId = EXCLUDED.DeploymentId, ClusterId = EXCLUDED.ClusterId, Namespace = EXCLUDED.Namespace, ObservationPeriodEnd = EXCLUDED.ObservationPeriodEnd, Locked = EXCLUDED.Locked, DeploymentName = EXCLUDED.DeploymentName, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetPeers() {
		if err := insertIntoNetworkbaselinePeers(ctx, tx, child, obj.GetDeploymentId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from networkbaseline_Peers where networkbaseline_DeploymentId = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetDeploymentId(), len(obj.GetPeers()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetForbiddenPeers() {
		if err := insertIntoNetworkbaselineForbiddenPeers(ctx, tx, child, obj.GetDeploymentId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from networkbaseline_ForbiddenPeers where networkbaseline_DeploymentId = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetDeploymentId(), len(obj.GetForbiddenPeers()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoNetworkbaselinePeers(ctx context.Context, tx pgx.Tx, obj *storage.NetworkBaselinePeer, networkbaseline_DeploymentId string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		networkbaseline_DeploymentId,
		idx,
		obj.GetEntity().GetInfo().GetType(),
		obj.GetEntity().GetInfo().GetId(),
		obj.GetEntity().GetInfo().GetDeployment().GetName(),
		obj.GetEntity().GetInfo().GetDeployment().GetNamespace(),
		obj.GetEntity().GetInfo().GetDeployment().GetCluster(),
		obj.GetEntity().GetInfo().GetExternalSource().GetName(),
		obj.GetEntity().GetInfo().GetExternalSource().GetCidr(),
		obj.GetEntity().GetInfo().GetExternalSource().GetDefault(),
		obj.GetEntity().GetScope().GetClusterId(),
	}

	finalStr := "INSERT INTO networkbaseline_Peers (networkbaseline_DeploymentId, idx, Entity_Info_Type, Entity_Info_Id, Entity_Info_Deployment_Name, Entity_Info_Deployment_Namespace, Entity_Info_Deployment_Cluster, Entity_Info_ExternalSource_Name, Entity_Info_ExternalSource_Cidr, Entity_Info_ExternalSource_Default, Entity_Scope_ClusterId) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT(networkbaseline_DeploymentId, idx) DO UPDATE SET networkbaseline_DeploymentId = EXCLUDED.networkbaseline_DeploymentId, idx = EXCLUDED.idx, Entity_Info_Type = EXCLUDED.Entity_Info_Type, Entity_Info_Id = EXCLUDED.Entity_Info_Id, Entity_Info_Deployment_Name = EXCLUDED.Entity_Info_Deployment_Name, Entity_Info_Deployment_Namespace = EXCLUDED.Entity_Info_Deployment_Namespace, Entity_Info_Deployment_Cluster = EXCLUDED.Entity_Info_Deployment_Cluster, Entity_Info_ExternalSource_Name = EXCLUDED.Entity_Info_ExternalSource_Name, Entity_Info_ExternalSource_Cidr = EXCLUDED.Entity_Info_ExternalSource_Cidr, Entity_Info_ExternalSource_Default = EXCLUDED.Entity_Info_ExternalSource_Default, Entity_Scope_ClusterId = EXCLUDED.Entity_Scope_ClusterId"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetEntity().GetInfo().GetDeployment().GetListenPorts() {
		if err := insertIntoNetworkbaselinePeersListenPorts(ctx, tx, child, networkbaseline_DeploymentId, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from networkbaseline_Peers_ListenPorts where networkbaseline_DeploymentId = $1 AND networkbaseline_Peers_idx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, networkbaseline_DeploymentId, idx, len(obj.GetEntity().GetInfo().GetDeployment().GetListenPorts()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetProperties() {
		if err := insertIntoNetworkbaselinePeersProperties(ctx, tx, child, networkbaseline_DeploymentId, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from networkbaseline_Peers_Properties where networkbaseline_DeploymentId = $1 AND networkbaseline_Peers_idx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, networkbaseline_DeploymentId, idx, len(obj.GetProperties()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoNetworkbaselinePeersListenPorts(ctx context.Context, tx pgx.Tx, obj *storage.NetworkEntityInfo_Deployment_ListenPort, networkbaseline_DeploymentId string, networkbaseline_Peers_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		networkbaseline_DeploymentId,
		networkbaseline_Peers_idx,
		idx,
		obj.GetPort(),
		obj.GetL4Protocol(),
	}

	finalStr := "INSERT INTO networkbaseline_Peers_ListenPorts (networkbaseline_DeploymentId, networkbaseline_Peers_idx, idx, Port, L4Protocol) VALUES($1, $2, $3, $4, $5) ON CONFLICT(networkbaseline_DeploymentId, networkbaseline_Peers_idx, idx) DO UPDATE SET networkbaseline_DeploymentId = EXCLUDED.networkbaseline_DeploymentId, networkbaseline_Peers_idx = EXCLUDED.networkbaseline_Peers_idx, idx = EXCLUDED.idx, Port = EXCLUDED.Port, L4Protocol = EXCLUDED.L4Protocol"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoNetworkbaselinePeersProperties(ctx context.Context, tx pgx.Tx, obj *storage.NetworkBaselineConnectionProperties, networkbaseline_DeploymentId string, networkbaseline_Peers_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		networkbaseline_DeploymentId,
		networkbaseline_Peers_idx,
		idx,
		obj.GetIngress(),
		obj.GetPort(),
		obj.GetProtocol(),
	}

	finalStr := "INSERT INTO networkbaseline_Peers_Properties (networkbaseline_DeploymentId, networkbaseline_Peers_idx, idx, Ingress, Port, Protocol) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT(networkbaseline_DeploymentId, networkbaseline_Peers_idx, idx) DO UPDATE SET networkbaseline_DeploymentId = EXCLUDED.networkbaseline_DeploymentId, networkbaseline_Peers_idx = EXCLUDED.networkbaseline_Peers_idx, idx = EXCLUDED.idx, Ingress = EXCLUDED.Ingress, Port = EXCLUDED.Port, Protocol = EXCLUDED.Protocol"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoNetworkbaselineForbiddenPeers(ctx context.Context, tx pgx.Tx, obj *storage.NetworkBaselinePeer, networkbaseline_DeploymentId string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		networkbaseline_DeploymentId,
		idx,
		obj.GetEntity().GetInfo().GetType(),
		obj.GetEntity().GetInfo().GetId(),
		obj.GetEntity().GetInfo().GetDeployment().GetName(),
		obj.GetEntity().GetInfo().GetDeployment().GetNamespace(),
		obj.GetEntity().GetInfo().GetDeployment().GetCluster(),
		obj.GetEntity().GetInfo().GetExternalSource().GetName(),
		obj.GetEntity().GetInfo().GetExternalSource().GetCidr(),
		obj.GetEntity().GetInfo().GetExternalSource().GetDefault(),
		obj.GetEntity().GetScope().GetClusterId(),
	}

	finalStr := "INSERT INTO networkbaseline_ForbiddenPeers (networkbaseline_DeploymentId, idx, Entity_Info_Type, Entity_Info_Id, Entity_Info_Deployment_Name, Entity_Info_Deployment_Namespace, Entity_Info_Deployment_Cluster, Entity_Info_ExternalSource_Name, Entity_Info_ExternalSource_Cidr, Entity_Info_ExternalSource_Default, Entity_Scope_ClusterId) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT(networkbaseline_DeploymentId, idx) DO UPDATE SET networkbaseline_DeploymentId = EXCLUDED.networkbaseline_DeploymentId, idx = EXCLUDED.idx, Entity_Info_Type = EXCLUDED.Entity_Info_Type, Entity_Info_Id = EXCLUDED.Entity_Info_Id, Entity_Info_Deployment_Name = EXCLUDED.Entity_Info_Deployment_Name, Entity_Info_Deployment_Namespace = EXCLUDED.Entity_Info_Deployment_Namespace, Entity_Info_Deployment_Cluster = EXCLUDED.Entity_Info_Deployment_Cluster, Entity_Info_ExternalSource_Name = EXCLUDED.Entity_Info_ExternalSource_Name, Entity_Info_ExternalSource_Cidr = EXCLUDED.Entity_Info_ExternalSource_Cidr, Entity_Info_ExternalSource_Default = EXCLUDED.Entity_Info_ExternalSource_Default, Entity_Scope_ClusterId = EXCLUDED.Entity_Scope_ClusterId"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetEntity().GetInfo().GetDeployment().GetListenPorts() {
		if err := insertIntoNetworkbaselineForbiddenPeersListenPorts(ctx, tx, child, networkbaseline_DeploymentId, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from networkbaseline_ForbiddenPeers_ListenPorts where networkbaseline_DeploymentId = $1 AND networkbaseline_ForbiddenPeers_idx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, networkbaseline_DeploymentId, idx, len(obj.GetEntity().GetInfo().GetDeployment().GetListenPorts()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetProperties() {
		if err := insertIntoNetworkbaselineForbiddenPeersProperties(ctx, tx, child, networkbaseline_DeploymentId, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from networkbaseline_ForbiddenPeers_Properties where networkbaseline_DeploymentId = $1 AND networkbaseline_ForbiddenPeers_idx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, networkbaseline_DeploymentId, idx, len(obj.GetProperties()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoNetworkbaselineForbiddenPeersListenPorts(ctx context.Context, tx pgx.Tx, obj *storage.NetworkEntityInfo_Deployment_ListenPort, networkbaseline_DeploymentId string, networkbaseline_ForbiddenPeers_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		networkbaseline_DeploymentId,
		networkbaseline_ForbiddenPeers_idx,
		idx,
		obj.GetPort(),
		obj.GetL4Protocol(),
	}

	finalStr := "INSERT INTO networkbaseline_ForbiddenPeers_ListenPorts (networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx, idx, Port, L4Protocol) VALUES($1, $2, $3, $4, $5) ON CONFLICT(networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx, idx) DO UPDATE SET networkbaseline_DeploymentId = EXCLUDED.networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx = EXCLUDED.networkbaseline_ForbiddenPeers_idx, idx = EXCLUDED.idx, Port = EXCLUDED.Port, L4Protocol = EXCLUDED.L4Protocol"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoNetworkbaselineForbiddenPeersProperties(ctx context.Context, tx pgx.Tx, obj *storage.NetworkBaselineConnectionProperties, networkbaseline_DeploymentId string, networkbaseline_ForbiddenPeers_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		networkbaseline_DeploymentId,
		networkbaseline_ForbiddenPeers_idx,
		idx,
		obj.GetIngress(),
		obj.GetPort(),
		obj.GetProtocol(),
	}

	finalStr := "INSERT INTO networkbaseline_ForbiddenPeers_Properties (networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx, idx, Ingress, Port, Protocol) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT(networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx, idx) DO UPDATE SET networkbaseline_DeploymentId = EXCLUDED.networkbaseline_DeploymentId, networkbaseline_ForbiddenPeers_idx = EXCLUDED.networkbaseline_ForbiddenPeers_idx, idx = EXCLUDED.idx, Ingress = EXCLUDED.Ingress, Port = EXCLUDED.Port, Protocol = EXCLUDED.Protocol"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	createTableNetworkbaseline(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.NetworkBaseline) error {
	conn, release := s.acquireConn(ctx, ops.Get, "NetworkBaseline")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoNetworkbaseline(ctx, tx, obj); err != nil {
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

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.NetworkBaseline) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "NetworkBaseline")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.NetworkBaseline) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "NetworkBaseline")

	return s.upsert(ctx, objs...)
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "NetworkBaseline")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, deploymentId string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "NetworkBaseline")

	row := s.db.QueryRow(ctx, existsStmt, deploymentId)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

func (s *storeImpl) GetWithRollup(ctx context.Context, deploymentId string) (map[string]interface{}, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "NetworkBaseline")

	row := s.db.QueryRow(ctx, getWithRollupStmt, deploymentId)
	var serializedRow []byte
	if err := row.Scan(&serializedRow); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var out map[string]interface{}
	if err := json.Unmarshal(serializedRow, &out); err != nil {
		return nil, false, err
	}
	return out, true, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, deploymentId string) (*storage.NetworkBaseline, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "NetworkBaseline")

	conn, release := s.acquireConn(ctx, ops.Get, "NetworkBaseline")
	defer release()

	row := conn.QueryRow(ctx, getStmt, deploymentId)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.NetworkBaseline
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
func (s *storeImpl) Delete(ctx context.Context, deploymentId string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "NetworkBaseline")

	conn, release := s.acquireConn(ctx, ops.Remove, "NetworkBaseline")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, deploymentId); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.NetworkBaselineIDs")

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
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.NetworkBaseline, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "NetworkBaseline")

	conn, release := s.acquireConn(ctx, ops.GetMany, "NetworkBaseline")
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
	resultsByID := make(map[string]*storage.NetworkBaseline)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.NetworkBaseline{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetDeploymentId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.NetworkBaseline, 0, len(resultsByID))
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "NetworkBaseline")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "NetworkBaseline")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.NetworkBaseline) error) error {
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
		var msg storage.NetworkBaseline
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

func dropTableNetworkbaseline(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS networkbaseline CASCADE")
	dropTableNetworkbaselinePeers(ctx, db)
	dropTableNetworkbaselineForbiddenPeers(ctx, db)

}

func dropTableNetworkbaselinePeers(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS networkbaseline_Peers CASCADE")
	dropTableNetworkbaselinePeersListenPorts(ctx, db)
	dropTableNetworkbaselinePeersProperties(ctx, db)

}

func dropTableNetworkbaselinePeersListenPorts(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS networkbaseline_Peers_ListenPorts CASCADE")

}

func dropTableNetworkbaselinePeersProperties(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS networkbaseline_Peers_Properties CASCADE")

}

func dropTableNetworkbaselineForbiddenPeers(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS networkbaseline_ForbiddenPeers CASCADE")
	dropTableNetworkbaselineForbiddenPeersListenPorts(ctx, db)
	dropTableNetworkbaselineForbiddenPeersProperties(ctx, db)

}

func dropTableNetworkbaselineForbiddenPeersListenPorts(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS networkbaseline_ForbiddenPeers_ListenPorts CASCADE")

}

func dropTableNetworkbaselineForbiddenPeersProperties(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS networkbaseline_ForbiddenPeers_Properties CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableNetworkbaseline(ctx, db)
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
