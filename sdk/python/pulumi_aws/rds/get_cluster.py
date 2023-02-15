# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
    'get_cluster_output',
]

@pulumi.output_type
class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, arn=None, availability_zones=None, backtrack_window=None, backup_retention_period=None, cluster_identifier=None, cluster_members=None, cluster_resource_id=None, database_name=None, db_cluster_parameter_group_name=None, db_subnet_group_name=None, enabled_cloudwatch_logs_exports=None, endpoint=None, engine=None, engine_mode=None, engine_version=None, final_snapshot_identifier=None, hosted_zone_id=None, iam_database_authentication_enabled=None, iam_roles=None, id=None, kms_key_id=None, master_username=None, network_type=None, port=None, preferred_backup_window=None, preferred_maintenance_window=None, reader_endpoint=None, replication_source_identifier=None, storage_encrypted=None, tags=None, vpc_security_group_ids=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError("Expected argument 'availability_zones' to be a list")
        pulumi.set(__self__, "availability_zones", availability_zones)
        if backtrack_window and not isinstance(backtrack_window, int):
            raise TypeError("Expected argument 'backtrack_window' to be a int")
        pulumi.set(__self__, "backtrack_window", backtrack_window)
        if backup_retention_period and not isinstance(backup_retention_period, int):
            raise TypeError("Expected argument 'backup_retention_period' to be a int")
        pulumi.set(__self__, "backup_retention_period", backup_retention_period)
        if cluster_identifier and not isinstance(cluster_identifier, str):
            raise TypeError("Expected argument 'cluster_identifier' to be a str")
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        if cluster_members and not isinstance(cluster_members, list):
            raise TypeError("Expected argument 'cluster_members' to be a list")
        pulumi.set(__self__, "cluster_members", cluster_members)
        if cluster_resource_id and not isinstance(cluster_resource_id, str):
            raise TypeError("Expected argument 'cluster_resource_id' to be a str")
        pulumi.set(__self__, "cluster_resource_id", cluster_resource_id)
        if database_name and not isinstance(database_name, str):
            raise TypeError("Expected argument 'database_name' to be a str")
        pulumi.set(__self__, "database_name", database_name)
        if db_cluster_parameter_group_name and not isinstance(db_cluster_parameter_group_name, str):
            raise TypeError("Expected argument 'db_cluster_parameter_group_name' to be a str")
        pulumi.set(__self__, "db_cluster_parameter_group_name", db_cluster_parameter_group_name)
        if db_subnet_group_name and not isinstance(db_subnet_group_name, str):
            raise TypeError("Expected argument 'db_subnet_group_name' to be a str")
        pulumi.set(__self__, "db_subnet_group_name", db_subnet_group_name)
        if enabled_cloudwatch_logs_exports and not isinstance(enabled_cloudwatch_logs_exports, list):
            raise TypeError("Expected argument 'enabled_cloudwatch_logs_exports' to be a list")
        pulumi.set(__self__, "enabled_cloudwatch_logs_exports", enabled_cloudwatch_logs_exports)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if engine_mode and not isinstance(engine_mode, str):
            raise TypeError("Expected argument 'engine_mode' to be a str")
        pulumi.set(__self__, "engine_mode", engine_mode)
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if final_snapshot_identifier and not isinstance(final_snapshot_identifier, str):
            raise TypeError("Expected argument 'final_snapshot_identifier' to be a str")
        pulumi.set(__self__, "final_snapshot_identifier", final_snapshot_identifier)
        if hosted_zone_id and not isinstance(hosted_zone_id, str):
            raise TypeError("Expected argument 'hosted_zone_id' to be a str")
        pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)
        if iam_database_authentication_enabled and not isinstance(iam_database_authentication_enabled, bool):
            raise TypeError("Expected argument 'iam_database_authentication_enabled' to be a bool")
        pulumi.set(__self__, "iam_database_authentication_enabled", iam_database_authentication_enabled)
        if iam_roles and not isinstance(iam_roles, list):
            raise TypeError("Expected argument 'iam_roles' to be a list")
        pulumi.set(__self__, "iam_roles", iam_roles)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if master_username and not isinstance(master_username, str):
            raise TypeError("Expected argument 'master_username' to be a str")
        pulumi.set(__self__, "master_username", master_username)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if preferred_backup_window and not isinstance(preferred_backup_window, str):
            raise TypeError("Expected argument 'preferred_backup_window' to be a str")
        pulumi.set(__self__, "preferred_backup_window", preferred_backup_window)
        if preferred_maintenance_window and not isinstance(preferred_maintenance_window, str):
            raise TypeError("Expected argument 'preferred_maintenance_window' to be a str")
        pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if reader_endpoint and not isinstance(reader_endpoint, str):
            raise TypeError("Expected argument 'reader_endpoint' to be a str")
        pulumi.set(__self__, "reader_endpoint", reader_endpoint)
        if replication_source_identifier and not isinstance(replication_source_identifier, str):
            raise TypeError("Expected argument 'replication_source_identifier' to be a str")
        pulumi.set(__self__, "replication_source_identifier", replication_source_identifier)
        if storage_encrypted and not isinstance(storage_encrypted, bool):
            raise TypeError("Expected argument 'storage_encrypted' to be a bool")
        pulumi.set(__self__, "storage_encrypted", storage_encrypted)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_security_group_ids and not isinstance(vpc_security_group_ids, list):
            raise TypeError("Expected argument 'vpc_security_group_ids' to be a list")
        pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Sequence[str]:
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="backtrackWindow")
    def backtrack_window(self) -> int:
        return pulumi.get(self, "backtrack_window")

    @property
    @pulumi.getter(name="backupRetentionPeriod")
    def backup_retention_period(self) -> int:
        return pulumi.get(self, "backup_retention_period")

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> str:
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter(name="clusterMembers")
    def cluster_members(self) -> Sequence[str]:
        return pulumi.get(self, "cluster_members")

    @property
    @pulumi.getter(name="clusterResourceId")
    def cluster_resource_id(self) -> str:
        return pulumi.get(self, "cluster_resource_id")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="dbClusterParameterGroupName")
    def db_cluster_parameter_group_name(self) -> str:
        return pulumi.get(self, "db_cluster_parameter_group_name")

    @property
    @pulumi.getter(name="dbSubnetGroupName")
    def db_subnet_group_name(self) -> str:
        return pulumi.get(self, "db_subnet_group_name")

    @property
    @pulumi.getter(name="enabledCloudwatchLogsExports")
    def enabled_cloudwatch_logs_exports(self) -> Sequence[str]:
        return pulumi.get(self, "enabled_cloudwatch_logs_exports")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def engine(self) -> str:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineMode")
    def engine_mode(self) -> str:
        return pulumi.get(self, "engine_mode")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="finalSnapshotIdentifier")
    def final_snapshot_identifier(self) -> str:
        return pulumi.get(self, "final_snapshot_identifier")

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> str:
        return pulumi.get(self, "hosted_zone_id")

    @property
    @pulumi.getter(name="iamDatabaseAuthenticationEnabled")
    def iam_database_authentication_enabled(self) -> bool:
        return pulumi.get(self, "iam_database_authentication_enabled")

    @property
    @pulumi.getter(name="iamRoles")
    def iam_roles(self) -> Sequence[str]:
        return pulumi.get(self, "iam_roles")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> str:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="masterUsername")
    def master_username(self) -> str:
        return pulumi.get(self, "master_username")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="preferredBackupWindow")
    def preferred_backup_window(self) -> str:
        return pulumi.get(self, "preferred_backup_window")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> str:
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter(name="readerEndpoint")
    def reader_endpoint(self) -> str:
        return pulumi.get(self, "reader_endpoint")

    @property
    @pulumi.getter(name="replicationSourceIdentifier")
    def replication_source_identifier(self) -> str:
        return pulumi.get(self, "replication_source_identifier")

    @property
    @pulumi.getter(name="storageEncrypted")
    def storage_encrypted(self) -> bool:
        return pulumi.get(self, "storage_encrypted")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Sequence[str]:
        return pulumi.get(self, "vpc_security_group_ids")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            arn=self.arn,
            availability_zones=self.availability_zones,
            backtrack_window=self.backtrack_window,
            backup_retention_period=self.backup_retention_period,
            cluster_identifier=self.cluster_identifier,
            cluster_members=self.cluster_members,
            cluster_resource_id=self.cluster_resource_id,
            database_name=self.database_name,
            db_cluster_parameter_group_name=self.db_cluster_parameter_group_name,
            db_subnet_group_name=self.db_subnet_group_name,
            enabled_cloudwatch_logs_exports=self.enabled_cloudwatch_logs_exports,
            endpoint=self.endpoint,
            engine=self.engine,
            engine_mode=self.engine_mode,
            engine_version=self.engine_version,
            final_snapshot_identifier=self.final_snapshot_identifier,
            hosted_zone_id=self.hosted_zone_id,
            iam_database_authentication_enabled=self.iam_database_authentication_enabled,
            iam_roles=self.iam_roles,
            id=self.id,
            kms_key_id=self.kms_key_id,
            master_username=self.master_username,
            network_type=self.network_type,
            port=self.port,
            preferred_backup_window=self.preferred_backup_window,
            preferred_maintenance_window=self.preferred_maintenance_window,
            reader_endpoint=self.reader_endpoint,
            replication_source_identifier=self.replication_source_identifier,
            storage_encrypted=self.storage_encrypted,
            tags=self.tags,
            vpc_security_group_ids=self.vpc_security_group_ids)


def get_cluster(cluster_identifier: Optional[str] = None,
                tags: Optional[Mapping[str, str]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Provides information about an RDS cluster.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    cluster_name = aws.rds.get_cluster(cluster_identifier="clusterName")
    ```


    :param str cluster_identifier: Cluster identifier of the RDS cluster.
    :param Mapping[str, str] tags: A map of tags assigned to the resource.
    """
    __args__ = dict()
    __args__['clusterIdentifier'] = cluster_identifier
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:rds/getCluster:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        arn=__ret__.arn,
        availability_zones=__ret__.availability_zones,
        backtrack_window=__ret__.backtrack_window,
        backup_retention_period=__ret__.backup_retention_period,
        cluster_identifier=__ret__.cluster_identifier,
        cluster_members=__ret__.cluster_members,
        cluster_resource_id=__ret__.cluster_resource_id,
        database_name=__ret__.database_name,
        db_cluster_parameter_group_name=__ret__.db_cluster_parameter_group_name,
        db_subnet_group_name=__ret__.db_subnet_group_name,
        enabled_cloudwatch_logs_exports=__ret__.enabled_cloudwatch_logs_exports,
        endpoint=__ret__.endpoint,
        engine=__ret__.engine,
        engine_mode=__ret__.engine_mode,
        engine_version=__ret__.engine_version,
        final_snapshot_identifier=__ret__.final_snapshot_identifier,
        hosted_zone_id=__ret__.hosted_zone_id,
        iam_database_authentication_enabled=__ret__.iam_database_authentication_enabled,
        iam_roles=__ret__.iam_roles,
        id=__ret__.id,
        kms_key_id=__ret__.kms_key_id,
        master_username=__ret__.master_username,
        network_type=__ret__.network_type,
        port=__ret__.port,
        preferred_backup_window=__ret__.preferred_backup_window,
        preferred_maintenance_window=__ret__.preferred_maintenance_window,
        reader_endpoint=__ret__.reader_endpoint,
        replication_source_identifier=__ret__.replication_source_identifier,
        storage_encrypted=__ret__.storage_encrypted,
        tags=__ret__.tags,
        vpc_security_group_ids=__ret__.vpc_security_group_ids)


@_utilities.lift_output_func(get_cluster)
def get_cluster_output(cluster_identifier: Optional[pulumi.Input[str]] = None,
                       tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterResult]:
    """
    Provides information about an RDS cluster.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    cluster_name = aws.rds.get_cluster(cluster_identifier="clusterName")
    ```


    :param str cluster_identifier: Cluster identifier of the RDS cluster.
    :param Mapping[str, str] tags: A map of tags assigned to the resource.
    """
    ...
