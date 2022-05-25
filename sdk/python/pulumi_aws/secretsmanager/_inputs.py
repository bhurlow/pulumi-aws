# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'SecretReplicaArgs',
    'SecretRotationRotationRulesArgs',
    'SecretRotationRulesArgs',
    'GetSecretsFilterArgs',
]

@pulumi.input_type
class SecretReplicaArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 last_accessed_date: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_message: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] region: Region for replicating the secret.
        :param pulumi.Input[str] kms_key_id: ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account's default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
        :param pulumi.Input[str] last_accessed_date: Date that you last accessed the secret in the Region.
        :param pulumi.Input[str] status: Status can be `InProgress`, `Failed`, or `InSync`.
        :param pulumi.Input[str] status_message: Message such as `Replication succeeded` or `Secret with this name already exists in this region`.
        """
        pulumi.set(__self__, "region", region)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if last_accessed_date is not None:
            pulumi.set(__self__, "last_accessed_date", last_accessed_date)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if status_message is not None:
            pulumi.set(__self__, "status_message", status_message)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Region for replicating the secret.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account's default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="lastAccessedDate")
    def last_accessed_date(self) -> Optional[pulumi.Input[str]]:
        """
        Date that you last accessed the secret in the Region.
        """
        return pulumi.get(self, "last_accessed_date")

    @last_accessed_date.setter
    def last_accessed_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_accessed_date", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status can be `InProgress`, `Failed`, or `InSync`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> Optional[pulumi.Input[str]]:
        """
        Message such as `Replication succeeded` or `Secret with this name already exists in this region`.
        """
        return pulumi.get(self, "status_message")

    @status_message.setter
    def status_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status_message", value)


@pulumi.input_type
class SecretRotationRotationRulesArgs:
    def __init__(__self__, *,
                 automatically_after_days: pulumi.Input[int]):
        """
        :param pulumi.Input[int] automatically_after_days: Specifies the number of days between automatic scheduled rotations of the secret.
        """
        pulumi.set(__self__, "automatically_after_days", automatically_after_days)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days between automatic scheduled rotations of the secret.
        """
        return pulumi.get(self, "automatically_after_days")

    @automatically_after_days.setter
    def automatically_after_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "automatically_after_days", value)


@pulumi.input_type
class SecretRotationRulesArgs:
    def __init__(__self__, *,
                 automatically_after_days: pulumi.Input[int]):
        """
        :param pulumi.Input[int] automatically_after_days: Specifies the number of days between automatic scheduled rotations of the secret.
        """
        pulumi.set(__self__, "automatically_after_days", automatically_after_days)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days between automatic scheduled rotations of the secret.
        """
        return pulumi.get(self, "automatically_after_days")

    @automatically_after_days.setter
    def automatically_after_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "automatically_after_days", value)


@pulumi.input_type
class GetSecretsFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: The name of the filter field. Valid values can be found in the [Secrets Manager ListSecrets API Reference](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html).
        :param Sequence[str] values: Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter field. Valid values can be found in the [Secrets Manager ListSecrets API Reference](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


