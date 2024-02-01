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
    'AwsLogSourceSourceArgs',
    'DataLakeConfigurationArgs',
    'DataLakeConfigurationLifecycleConfigurationArgs',
    'DataLakeConfigurationLifecycleConfigurationExpirationArgs',
    'DataLakeConfigurationLifecycleConfigurationTransitionArgs',
    'DataLakeConfigurationReplicationConfigurationArgs',
    'DataLakeTimeoutsArgs',
]

@pulumi.input_type
class AwsLogSourceSourceArgs:
    def __init__(__self__, *,
                 regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 source_name: pulumi.Input[str],
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: Specify the Regions where you want to enable Security Lake.
        :param pulumi.Input[str] source_name: The name for a AWS source. This must be a Regionally unique value. Valid values: `ROUTE53`, `VPC_FLOW`, `SH_FINDINGS`, `CLOUD_TRAIL_MGMT`, `LAMBDA_EXECUTION`, `S3_DATA`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] accounts: Specify the AWS account information where you want to enable Security Lake.
        :param pulumi.Input[str] source_version: The version for a AWS source. This must be a Regionally unique value.
        """
        pulumi.set(__self__, "regions", regions)
        pulumi.set(__self__, "source_name", source_name)
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if source_version is not None:
            pulumi.set(__self__, "source_version", source_version)

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specify the Regions where you want to enable Security Lake.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="sourceName")
    def source_name(self) -> pulumi.Input[str]:
        """
        The name for a AWS source. This must be a Regionally unique value. Valid values: `ROUTE53`, `VPC_FLOW`, `SH_FINDINGS`, `CLOUD_TRAIL_MGMT`, `LAMBDA_EXECUTION`, `S3_DATA`.
        """
        return pulumi.get(self, "source_name")

    @source_name.setter
    def source_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_name", value)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specify the AWS account information where you want to enable Security Lake.
        """
        return pulumi.get(self, "accounts")

    @accounts.setter
    def accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "accounts", value)

    @property
    @pulumi.getter(name="sourceVersion")
    def source_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version for a AWS source. This must be a Regionally unique value.
        """
        return pulumi.get(self, "source_version")

    @source_version.setter
    def source_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_version", value)


@pulumi.input_type
class DataLakeConfigurationArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 encryption_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]] = None,
                 lifecycle_configuration: Optional[pulumi.Input['DataLakeConfigurationLifecycleConfigurationArgs']] = None,
                 replication_configuration: Optional[pulumi.Input['DataLakeConfigurationReplicationConfigurationArgs']] = None):
        """
        :param pulumi.Input[str] region: The AWS Regions where Security Lake is automatically enabled.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]] encryption_configurations: Provides encryption details of Amazon Security Lake object.
        :param pulumi.Input['DataLakeConfigurationLifecycleConfigurationArgs'] lifecycle_configuration: Provides lifecycle details of Amazon Security Lake object.
        :param pulumi.Input['DataLakeConfigurationReplicationConfigurationArgs'] replication_configuration: Provides replication details of Amazon Security Lake object.
        """
        pulumi.set(__self__, "region", region)
        if encryption_configurations is not None:
            pulumi.set(__self__, "encryption_configurations", encryption_configurations)
        if lifecycle_configuration is not None:
            pulumi.set(__self__, "lifecycle_configuration", lifecycle_configuration)
        if replication_configuration is not None:
            pulumi.set(__self__, "replication_configuration", replication_configuration)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The AWS Regions where Security Lake is automatically enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="encryptionConfigurations")
    def encryption_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]]:
        """
        Provides encryption details of Amazon Security Lake object.
        """
        return pulumi.get(self, "encryption_configurations")

    @encryption_configurations.setter
    def encryption_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "encryption_configurations", value)

    @property
    @pulumi.getter(name="lifecycleConfiguration")
    def lifecycle_configuration(self) -> Optional[pulumi.Input['DataLakeConfigurationLifecycleConfigurationArgs']]:
        """
        Provides lifecycle details of Amazon Security Lake object.
        """
        return pulumi.get(self, "lifecycle_configuration")

    @lifecycle_configuration.setter
    def lifecycle_configuration(self, value: Optional[pulumi.Input['DataLakeConfigurationLifecycleConfigurationArgs']]):
        pulumi.set(self, "lifecycle_configuration", value)

    @property
    @pulumi.getter(name="replicationConfiguration")
    def replication_configuration(self) -> Optional[pulumi.Input['DataLakeConfigurationReplicationConfigurationArgs']]:
        """
        Provides replication details of Amazon Security Lake object.
        """
        return pulumi.get(self, "replication_configuration")

    @replication_configuration.setter
    def replication_configuration(self, value: Optional[pulumi.Input['DataLakeConfigurationReplicationConfigurationArgs']]):
        pulumi.set(self, "replication_configuration", value)


@pulumi.input_type
class DataLakeConfigurationLifecycleConfigurationArgs:
    def __init__(__self__, *,
                 expiration: Optional[pulumi.Input['DataLakeConfigurationLifecycleConfigurationExpirationArgs']] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeConfigurationLifecycleConfigurationTransitionArgs']]]] = None):
        """
        :param pulumi.Input['DataLakeConfigurationLifecycleConfigurationExpirationArgs'] expiration: Provides data expiration details of Amazon Security Lake object.
        :param pulumi.Input[Sequence[pulumi.Input['DataLakeConfigurationLifecycleConfigurationTransitionArgs']]] transitions: Provides data storage transition details of Amazon Security Lake object.
        """
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if transitions is not None:
            pulumi.set(__self__, "transitions", transitions)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input['DataLakeConfigurationLifecycleConfigurationExpirationArgs']]:
        """
        Provides data expiration details of Amazon Security Lake object.
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input['DataLakeConfigurationLifecycleConfigurationExpirationArgs']]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def transitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeConfigurationLifecycleConfigurationTransitionArgs']]]]:
        """
        Provides data storage transition details of Amazon Security Lake object.
        """
        return pulumi.get(self, "transitions")

    @transitions.setter
    def transitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeConfigurationLifecycleConfigurationTransitionArgs']]]]):
        pulumi.set(self, "transitions", value)


@pulumi.input_type
class DataLakeConfigurationLifecycleConfigurationExpirationArgs:
    def __init__(__self__, *,
                 days: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] days: Number of days before data transition to a different S3 Storage Class in the Amazon Security Lake object.
        """
        if days is not None:
            pulumi.set(__self__, "days", days)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[int]]:
        """
        Number of days before data transition to a different S3 Storage Class in the Amazon Security Lake object.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "days", value)


@pulumi.input_type
class DataLakeConfigurationLifecycleConfigurationTransitionArgs:
    def __init__(__self__, *,
                 days: Optional[pulumi.Input[int]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] days: Number of days before data transition to a different S3 Storage Class in the Amazon Security Lake object.
        :param pulumi.Input[str] storage_class: The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
        """
        if days is not None:
            pulumi.set(__self__, "days", days)
        if storage_class is not None:
            pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[int]]:
        """
        Number of days before data transition to a different S3 Storage Class in the Amazon Security Lake object.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> Optional[pulumi.Input[str]]:
        """
        The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_class", value)


@pulumi.input_type
class DataLakeConfigurationReplicationConfigurationArgs:
    def __init__(__self__, *,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. Amazon S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different AWS Regions or within the same Region as the source bucket.
        :param pulumi.Input[str] role_arn: Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
        """
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. Amazon S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different AWS Regions or within the same Region as the source bucket.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


@pulumi.input_type
class DataLakeTimeoutsArgs:
    def __init__(__self__, *,
                 create: Optional[pulumi.Input[str]] = None,
                 delete: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] create: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        :param pulumi.Input[str] delete: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
        :param pulumi.Input[str] update: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        """
        if create is not None:
            pulumi.set(__self__, "create", create)
        if delete is not None:
            pulumi.set(__self__, "delete", delete)
        if update is not None:
            pulumi.set(__self__, "update", update)

    @property
    @pulumi.getter
    def create(self) -> Optional[pulumi.Input[str]]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        """
        return pulumi.get(self, "create")

    @create.setter
    def create(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create", value)

    @property
    @pulumi.getter
    def delete(self) -> Optional[pulumi.Input[str]]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
        """
        return pulumi.get(self, "delete")

    @delete.setter
    def delete(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete", value)

    @property
    @pulumi.getter
    def update(self) -> Optional[pulumi.Input[str]]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        """
        return pulumi.get(self, "update")

    @update.setter
    def update(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update", value)


