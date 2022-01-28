# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DevicePoolRule',
    'TestGridProjectVpcConfig',
]

@pulumi.output_type
class DevicePoolRule(dict):
    def __init__(__self__, *,
                 attribute: Optional[str] = None,
                 operator: Optional[str] = None,
                 value: Optional[str] = None):
        """
        :param str attribute: The rule's stringified attribute. Valid values are: `APPIUM_VERSION`, `ARN`, `AVAILABILITY`, `FLEET_TYPE`, `FORM_FACTOR`, `INSTANCE_ARN`, `INSTANCE_LABELS`, `MANUFACTURER`, `MODEL`, `OS_VERSION`, `PLATFORM`, `REMOTE_ACCESS_ENABLED`, `REMOTE_DEBUG_ENABLED`.
        :param str operator: Specifies how Device Farm compares the rule's attribute to the value. For the operators that are supported by each attribute. Valid values are: `EQUALS`, `NOT_IN`, `IN`, `GREATER_THAN`, `GREATER_THAN_OR_EQUALS`, `LESS_THAN`, `LESS_THAN_OR_EQUALS`, `CONTAINS`.
        :param str value: The rule's value.
        """
        if attribute is not None:
            pulumi.set(__self__, "attribute", attribute)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def attribute(self) -> Optional[str]:
        """
        The rule's stringified attribute. Valid values are: `APPIUM_VERSION`, `ARN`, `AVAILABILITY`, `FLEET_TYPE`, `FORM_FACTOR`, `INSTANCE_ARN`, `INSTANCE_LABELS`, `MANUFACTURER`, `MODEL`, `OS_VERSION`, `PLATFORM`, `REMOTE_ACCESS_ENABLED`, `REMOTE_DEBUG_ENABLED`.
        """
        return pulumi.get(self, "attribute")

    @property
    @pulumi.getter
    def operator(self) -> Optional[str]:
        """
        Specifies how Device Farm compares the rule's attribute to the value. For the operators that are supported by each attribute. Valid values are: `EQUALS`, `NOT_IN`, `IN`, `GREATER_THAN`, `GREATER_THAN_OR_EQUALS`, `LESS_THAN`, `LESS_THAN_OR_EQUALS`, `CONTAINS`.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The rule's value.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class TestGridProjectVpcConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroupIds":
            suggest = "security_group_ids"
        elif key == "subnetIds":
            suggest = "subnet_ids"
        elif key == "vpcId":
            suggest = "vpc_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TestGridProjectVpcConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TestGridProjectVpcConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TestGridProjectVpcConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_group_ids: Sequence[str],
                 subnet_ids: Sequence[str],
                 vpc_id: str):
        """
        :param Sequence[str] security_group_ids: A list of VPC security group IDs in your Amazon VPC.
        :param Sequence[str] subnet_ids: A list of VPC subnet IDs in your Amazon VPC.
        :param str vpc_id: The ID of the Amazon VPC.
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[str]:
        """
        A list of VPC security group IDs in your Amazon VPC.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Sequence[str]:
        """
        A list of VPC subnet IDs in your Amazon VPC.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the Amazon VPC.
        """
        return pulumi.get(self, "vpc_id")


