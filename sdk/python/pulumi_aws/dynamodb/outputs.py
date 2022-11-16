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
    'GlobalTableReplica',
    'TableAttribute',
    'TableGlobalSecondaryIndex',
    'TableLocalSecondaryIndex',
    'TablePointInTimeRecovery',
    'TableReplica',
    'TableServerSideEncryption',
    'TableTtl',
    'GetTableAttributeResult',
    'GetTableGlobalSecondaryIndexResult',
    'GetTableLocalSecondaryIndexResult',
    'GetTablePointInTimeRecoveryResult',
    'GetTableReplicaResult',
    'GetTableServerSideEncryptionResult',
    'GetTableTtlResult',
]

@pulumi.output_type
class GlobalTableReplica(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "regionName":
            suggest = "region_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GlobalTableReplica. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GlobalTableReplica.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GlobalTableReplica.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 region_name: str):
        """
        :param str region_name: AWS region name of replica DynamoDB TableE.g., `us-east-1`
        """
        pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> str:
        """
        AWS region name of replica DynamoDB TableE.g., `us-east-1`
        """
        return pulumi.get(self, "region_name")


@pulumi.output_type
class TableAttribute(dict):
    def __init__(__self__, *,
                 name: str,
                 type: str):
        """
        :param str name: Name of the index
        :param str type: Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the index
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class TableGlobalSecondaryIndex(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "hashKey":
            suggest = "hash_key"
        elif key == "projectionType":
            suggest = "projection_type"
        elif key == "nonKeyAttributes":
            suggest = "non_key_attributes"
        elif key == "rangeKey":
            suggest = "range_key"
        elif key == "readCapacity":
            suggest = "read_capacity"
        elif key == "writeCapacity":
            suggest = "write_capacity"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableGlobalSecondaryIndex. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableGlobalSecondaryIndex.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableGlobalSecondaryIndex.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 hash_key: str,
                 name: str,
                 projection_type: str,
                 non_key_attributes: Optional[Sequence[str]] = None,
                 range_key: Optional[str] = None,
                 read_capacity: Optional[int] = None,
                 write_capacity: Optional[int] = None):
        """
        :param str hash_key: Name of the hash key in the index; must be defined as an attribute in the resource.
        :param str name: Name of the index
        :param str projection_type: One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        :param Sequence[str] non_key_attributes: Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        :param str range_key: Name of the range key.
        :param int read_capacity: Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        :param int write_capacity: Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
        """
        pulumi.set(__self__, "hash_key", hash_key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "projection_type", projection_type)
        if non_key_attributes is not None:
            pulumi.set(__self__, "non_key_attributes", non_key_attributes)
        if range_key is not None:
            pulumi.set(__self__, "range_key", range_key)
        if read_capacity is not None:
            pulumi.set(__self__, "read_capacity", read_capacity)
        if write_capacity is not None:
            pulumi.set(__self__, "write_capacity", write_capacity)

    @property
    @pulumi.getter(name="hashKey")
    def hash_key(self) -> str:
        """
        Name of the hash key in the index; must be defined as an attribute in the resource.
        """
        return pulumi.get(self, "hash_key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the index
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectionType")
    def projection_type(self) -> str:
        """
        One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        """
        return pulumi.get(self, "projection_type")

    @property
    @pulumi.getter(name="nonKeyAttributes")
    def non_key_attributes(self) -> Optional[Sequence[str]]:
        """
        Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        """
        return pulumi.get(self, "non_key_attributes")

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> Optional[str]:
        """
        Name of the range key.
        """
        return pulumi.get(self, "range_key")

    @property
    @pulumi.getter(name="readCapacity")
    def read_capacity(self) -> Optional[int]:
        """
        Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        """
        return pulumi.get(self, "read_capacity")

    @property
    @pulumi.getter(name="writeCapacity")
    def write_capacity(self) -> Optional[int]:
        """
        Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
        """
        return pulumi.get(self, "write_capacity")


@pulumi.output_type
class TableLocalSecondaryIndex(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "projectionType":
            suggest = "projection_type"
        elif key == "rangeKey":
            suggest = "range_key"
        elif key == "nonKeyAttributes":
            suggest = "non_key_attributes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableLocalSecondaryIndex. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableLocalSecondaryIndex.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableLocalSecondaryIndex.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 projection_type: str,
                 range_key: str,
                 non_key_attributes: Optional[Sequence[str]] = None):
        """
        :param str name: Name of the index
        :param str projection_type: One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        :param str range_key: Name of the range key.
        :param Sequence[str] non_key_attributes: Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "projection_type", projection_type)
        pulumi.set(__self__, "range_key", range_key)
        if non_key_attributes is not None:
            pulumi.set(__self__, "non_key_attributes", non_key_attributes)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the index
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectionType")
    def projection_type(self) -> str:
        """
        One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        """
        return pulumi.get(self, "projection_type")

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> str:
        """
        Name of the range key.
        """
        return pulumi.get(self, "range_key")

    @property
    @pulumi.getter(name="nonKeyAttributes")
    def non_key_attributes(self) -> Optional[Sequence[str]]:
        """
        Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        """
        return pulumi.get(self, "non_key_attributes")


@pulumi.output_type
class TablePointInTimeRecovery(dict):
    def __init__(__self__, *,
                 enabled: bool):
        """
        :param bool enabled: Whether TTL is enabled.
        """
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether TTL is enabled.
        """
        return pulumi.get(self, "enabled")


@pulumi.output_type
class TableReplica(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "regionName":
            suggest = "region_name"
        elif key == "kmsKeyArn":
            suggest = "kms_key_arn"
        elif key == "pointInTimeRecovery":
            suggest = "point_in_time_recovery"
        elif key == "propagateTags":
            suggest = "propagate_tags"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableReplica. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableReplica.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableReplica.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 region_name: str,
                 kms_key_arn: Optional[str] = None,
                 point_in_time_recovery: Optional[bool] = None,
                 propagate_tags: Optional[bool] = None):
        """
        :param str region_name: Region name of the replica.
        :param str kms_key_arn: ARN of the CMK that should be used for the AWS KMS encryption. This attribute should only be specified if the key is different from the default DynamoDB CMK, `alias/aws/dynamodb`.
        :param bool point_in_time_recovery: Whether to enable Point In Time Recovery for the replica. Default is `false`.
        :param bool propagate_tags: Whether to propagate the global table's tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
        """
        pulumi.set(__self__, "region_name", region_name)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if point_in_time_recovery is not None:
            pulumi.set(__self__, "point_in_time_recovery", point_in_time_recovery)
        if propagate_tags is not None:
            pulumi.set(__self__, "propagate_tags", propagate_tags)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> str:
        """
        Region name of the replica.
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[str]:
        """
        ARN of the CMK that should be used for the AWS KMS encryption. This attribute should only be specified if the key is different from the default DynamoDB CMK, `alias/aws/dynamodb`.
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="pointInTimeRecovery")
    def point_in_time_recovery(self) -> Optional[bool]:
        """
        Whether to enable Point In Time Recovery for the replica. Default is `false`.
        """
        return pulumi.get(self, "point_in_time_recovery")

    @property
    @pulumi.getter(name="propagateTags")
    def propagate_tags(self) -> Optional[bool]:
        """
        Whether to propagate the global table's tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
        """
        return pulumi.get(self, "propagate_tags")


@pulumi.output_type
class TableServerSideEncryption(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "kmsKeyArn":
            suggest = "kms_key_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableServerSideEncryption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableServerSideEncryption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableServerSideEncryption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: bool,
                 kms_key_arn: Optional[str] = None):
        """
        :param bool enabled: Whether TTL is enabled.
        :param str kms_key_arn: ARN of the CMK that should be used for the AWS KMS encryption. This attribute should only be specified if the key is different from the default DynamoDB CMK, `alias/aws/dynamodb`.
        """
        pulumi.set(__self__, "enabled", enabled)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether TTL is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[str]:
        """
        ARN of the CMK that should be used for the AWS KMS encryption. This attribute should only be specified if the key is different from the default DynamoDB CMK, `alias/aws/dynamodb`.
        """
        return pulumi.get(self, "kms_key_arn")


@pulumi.output_type
class TableTtl(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "attributeName":
            suggest = "attribute_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TableTtl. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TableTtl.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TableTtl.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 attribute_name: str,
                 enabled: Optional[bool] = None):
        """
        :param str attribute_name: Name of the table attribute to store the TTL timestamp in.
        :param bool enabled: Whether TTL is enabled.
        """
        pulumi.set(__self__, "attribute_name", attribute_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> str:
        """
        Name of the table attribute to store the TTL timestamp in.
        """
        return pulumi.get(self, "attribute_name")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether TTL is enabled.
        """
        return pulumi.get(self, "enabled")


@pulumi.output_type
class GetTableAttributeResult(dict):
    def __init__(__self__, *,
                 name: str,
                 type: str):
        """
        :param str name: Name of the DynamoDB table.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the DynamoDB table.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class GetTableGlobalSecondaryIndexResult(dict):
    def __init__(__self__, *,
                 hash_key: str,
                 name: str,
                 non_key_attributes: Sequence[str],
                 projection_type: str,
                 range_key: str,
                 read_capacity: int,
                 write_capacity: int):
        """
        :param str name: Name of the DynamoDB table.
        """
        pulumi.set(__self__, "hash_key", hash_key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "non_key_attributes", non_key_attributes)
        pulumi.set(__self__, "projection_type", projection_type)
        pulumi.set(__self__, "range_key", range_key)
        pulumi.set(__self__, "read_capacity", read_capacity)
        pulumi.set(__self__, "write_capacity", write_capacity)

    @property
    @pulumi.getter(name="hashKey")
    def hash_key(self) -> str:
        return pulumi.get(self, "hash_key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the DynamoDB table.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nonKeyAttributes")
    def non_key_attributes(self) -> Sequence[str]:
        return pulumi.get(self, "non_key_attributes")

    @property
    @pulumi.getter(name="projectionType")
    def projection_type(self) -> str:
        return pulumi.get(self, "projection_type")

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> str:
        return pulumi.get(self, "range_key")

    @property
    @pulumi.getter(name="readCapacity")
    def read_capacity(self) -> int:
        return pulumi.get(self, "read_capacity")

    @property
    @pulumi.getter(name="writeCapacity")
    def write_capacity(self) -> int:
        return pulumi.get(self, "write_capacity")


@pulumi.output_type
class GetTableLocalSecondaryIndexResult(dict):
    def __init__(__self__, *,
                 name: str,
                 non_key_attributes: Sequence[str],
                 projection_type: str,
                 range_key: str):
        """
        :param str name: Name of the DynamoDB table.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "non_key_attributes", non_key_attributes)
        pulumi.set(__self__, "projection_type", projection_type)
        pulumi.set(__self__, "range_key", range_key)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the DynamoDB table.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nonKeyAttributes")
    def non_key_attributes(self) -> Sequence[str]:
        return pulumi.get(self, "non_key_attributes")

    @property
    @pulumi.getter(name="projectionType")
    def projection_type(self) -> str:
        return pulumi.get(self, "projection_type")

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> str:
        return pulumi.get(self, "range_key")


@pulumi.output_type
class GetTablePointInTimeRecoveryResult(dict):
    def __init__(__self__, *,
                 enabled: bool):
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")


@pulumi.output_type
class GetTableReplicaResult(dict):
    def __init__(__self__, *,
                 kms_key_arn: str,
                 region_name: str):
        pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> str:
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> str:
        return pulumi.get(self, "region_name")


@pulumi.output_type
class GetTableServerSideEncryptionResult(dict):
    def __init__(__self__, *,
                 enabled: bool,
                 kms_key_arn: str):
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "kms_key_arn", kms_key_arn)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> str:
        return pulumi.get(self, "kms_key_arn")


@pulumi.output_type
class GetTableTtlResult(dict):
    def __init__(__self__, *,
                 attribute_name: str,
                 enabled: bool):
        pulumi.set(__self__, "attribute_name", attribute_name)
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> str:
        return pulumi.get(self, "attribute_name")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")


