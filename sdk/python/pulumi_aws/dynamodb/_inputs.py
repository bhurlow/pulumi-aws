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
    'GlobalTableReplicaArgs',
    'TableAttributeArgs',
    'TableGlobalSecondaryIndexArgs',
    'TableLocalSecondaryIndexArgs',
    'TablePointInTimeRecoveryArgs',
    'TableReplicaArgs',
    'TableServerSideEncryptionArgs',
    'TableTtlArgs',
    'GetTableServerSideEncryptionArgs',
]

@pulumi.input_type
class GlobalTableReplicaArgs:
    def __init__(__self__, *,
                 region_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] region_name: AWS region name of replica DynamoDB TableE.g., `us-east-1`
        """
        pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Input[str]:
        """
        AWS region name of replica DynamoDB TableE.g., `us-east-1`
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_name", value)


@pulumi.input_type
class TableAttributeArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Name of the attribute
        :param pulumi.Input[str] type: Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the attribute
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class TableGlobalSecondaryIndexArgs:
    def __init__(__self__, *,
                 hash_key: pulumi.Input[str],
                 name: pulumi.Input[str],
                 projection_type: pulumi.Input[str],
                 non_key_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 range_key: Optional[pulumi.Input[str]] = None,
                 read_capacity: Optional[pulumi.Input[int]] = None,
                 write_capacity: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] hash_key: Name of the hash key in the index; must be defined as an attribute in the resource.
        :param pulumi.Input[str] name: Name of the index.
        :param pulumi.Input[str] projection_type: One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] non_key_attributes: Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        :param pulumi.Input[str] range_key: Name of the range key; must be defined
        :param pulumi.Input[int] read_capacity: Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        :param pulumi.Input[int] write_capacity: Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
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
    def hash_key(self) -> pulumi.Input[str]:
        """
        Name of the hash key in the index; must be defined as an attribute in the resource.
        """
        return pulumi.get(self, "hash_key")

    @hash_key.setter
    def hash_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "hash_key", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the index.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectionType")
    def projection_type(self) -> pulumi.Input[str]:
        """
        One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        """
        return pulumi.get(self, "projection_type")

    @projection_type.setter
    def projection_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "projection_type", value)

    @property
    @pulumi.getter(name="nonKeyAttributes")
    def non_key_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        """
        return pulumi.get(self, "non_key_attributes")

    @non_key_attributes.setter
    def non_key_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "non_key_attributes", value)

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the range key; must be defined
        """
        return pulumi.get(self, "range_key")

    @range_key.setter
    def range_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range_key", value)

    @property
    @pulumi.getter(name="readCapacity")
    def read_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        """
        return pulumi.get(self, "read_capacity")

    @read_capacity.setter
    def read_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "read_capacity", value)

    @property
    @pulumi.getter(name="writeCapacity")
    def write_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
        """
        return pulumi.get(self, "write_capacity")

    @write_capacity.setter
    def write_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "write_capacity", value)


@pulumi.input_type
class TableLocalSecondaryIndexArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 projection_type: pulumi.Input[str],
                 range_key: pulumi.Input[str],
                 non_key_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] name: Name of the index
        :param pulumi.Input[str] projection_type: One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        :param pulumi.Input[str] range_key: Name of the range key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] non_key_attributes: Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "projection_type", projection_type)
        pulumi.set(__self__, "range_key", range_key)
        if non_key_attributes is not None:
            pulumi.set(__self__, "non_key_attributes", non_key_attributes)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the index
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectionType")
    def projection_type(self) -> pulumi.Input[str]:
        """
        One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        """
        return pulumi.get(self, "projection_type")

    @projection_type.setter
    def projection_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "projection_type", value)

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> pulumi.Input[str]:
        """
        Name of the range key.
        """
        return pulumi.get(self, "range_key")

    @range_key.setter
    def range_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "range_key", value)

    @property
    @pulumi.getter(name="nonKeyAttributes")
    def non_key_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        """
        return pulumi.get(self, "non_key_attributes")

    @non_key_attributes.setter
    def non_key_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "non_key_attributes", value)


@pulumi.input_type
class TablePointInTimeRecoveryArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool]):
        """
        :param pulumi.Input[bool] enabled: Whether to enable point-in-time recovery. It can take 10 minutes to enable for new tables. If the `point_in_time_recovery` block is not provided, this defaults to `false`.
        """
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Whether to enable point-in-time recovery. It can take 10 minutes to enable for new tables. If the `point_in_time_recovery` block is not provided, this defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class TableReplicaArgs:
    def __init__(__self__, *,
                 region_name: pulumi.Input[str],
                 arn: Optional[pulumi.Input[str]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 point_in_time_recovery: Optional[pulumi.Input[bool]] = None,
                 propagate_tags: Optional[pulumi.Input[bool]] = None,
                 stream_arn: Optional[pulumi.Input[str]] = None,
                 stream_label: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] region_name: Region name of the replica.
        :param pulumi.Input[str] arn: ARN of the table
        :param pulumi.Input[str] kms_key_arn: ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
        :param pulumi.Input[bool] point_in_time_recovery: Whether to enable Point In Time Recovery for the replica. Default is `false`.
        :param pulumi.Input[bool] propagate_tags: Whether to propagate the global table's tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
        :param pulumi.Input[str] stream_arn: ARN of the Table Stream. Only available when `stream_enabled = true`
        :param pulumi.Input[str] stream_label: Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
        """
        pulumi.set(__self__, "region_name", region_name)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if point_in_time_recovery is not None:
            pulumi.set(__self__, "point_in_time_recovery", point_in_time_recovery)
        if propagate_tags is not None:
            pulumi.set(__self__, "propagate_tags", propagate_tags)
        if stream_arn is not None:
            pulumi.set(__self__, "stream_arn", stream_arn)
        if stream_label is not None:
            pulumi.set(__self__, "stream_label", stream_label)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Input[str]:
        """
        Region name of the replica.
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the table
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter(name="pointInTimeRecovery")
    def point_in_time_recovery(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable Point In Time Recovery for the replica. Default is `false`.
        """
        return pulumi.get(self, "point_in_time_recovery")

    @point_in_time_recovery.setter
    def point_in_time_recovery(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "point_in_time_recovery", value)

    @property
    @pulumi.getter(name="propagateTags")
    def propagate_tags(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to propagate the global table's tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
        """
        return pulumi.get(self, "propagate_tags")

    @propagate_tags.setter
    def propagate_tags(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "propagate_tags", value)

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the Table Stream. Only available when `stream_enabled = true`
        """
        return pulumi.get(self, "stream_arn")

    @stream_arn.setter
    def stream_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_arn", value)

    @property
    @pulumi.getter(name="streamLabel")
    def stream_label(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
        """
        return pulumi.get(self, "stream_label")

    @stream_label.setter
    def stream_label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_label", value)


@pulumi.input_type
class TableServerSideEncryptionArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 kms_key_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK). If `enabled` is `false` then server-side encryption is set to AWS-_owned_ key (shown as `DEFAULT` in the AWS console). Potentially confusingly, if `enabled` is `true` and no `kms_key_arn` is specified then server-side encryption is set to the _default_ KMS-_managed_ key (shown as `KMS` in the AWS console). The [AWS KMS documentation](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html) explains the difference between AWS-_owned_ and KMS-_managed_ keys.
        :param pulumi.Input[str] kms_key_arn: ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
        """
        pulumi.set(__self__, "enabled", enabled)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK). If `enabled` is `false` then server-side encryption is set to AWS-_owned_ key (shown as `DEFAULT` in the AWS console). Potentially confusingly, if `enabled` is `true` and no `kms_key_arn` is specified then server-side encryption is set to the _default_ KMS-_managed_ key (shown as `KMS` in the AWS console). The [AWS KMS documentation](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html) explains the difference between AWS-_owned_ and KMS-_managed_ keys.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)


@pulumi.input_type
class TableTtlArgs:
    def __init__(__self__, *,
                 attribute_name: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] attribute_name: Name of the table attribute to store the TTL timestamp in.
        :param pulumi.Input[bool] enabled: Whether TTL is enabled.
        """
        pulumi.set(__self__, "attribute_name", attribute_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> pulumi.Input[str]:
        """
        Name of the table attribute to store the TTL timestamp in.
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether TTL is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class GetTableServerSideEncryptionArgs:
    def __init__(__self__, *,
                 enabled: bool,
                 kms_key_arn: str):
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "kms_key_arn", kms_key_arn)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: bool):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> str:
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: str):
        pulumi.set(self, "kms_key_arn", value)


