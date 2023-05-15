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
    'GetPolicyResult',
    'AwaitableGetPolicyResult',
    'get_policy',
    'get_policy_output',
]

@pulumi.output_type
class GetPolicyResult:
    """
    A collection of values returned by getPolicy.
    """
    def __init__(__self__, arn=None, aws_managed=None, content=None, description=None, id=None, name=None, policy_id=None, type=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if aws_managed and not isinstance(aws_managed, bool):
            raise TypeError("Expected argument 'aws_managed' to be a bool")
        pulumi.set(__self__, "aws_managed", aws_managed)
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        pulumi.set(__self__, "content", content)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        pulumi.set(__self__, "policy_id", policy_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name of the policy.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsManaged")
    def aws_managed(self) -> bool:
        """
        Indicates if a policy is an AWS managed policy.
        """
        return pulumi.get(self, "aws_managed")

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The text content of the policy.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The friendly name of the policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> str:
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of policy values can be `SERVICE_CONTROL_POLICY | TAG_POLICY | BACKUP_POLICY | AISERVICES_OPT_OUT_POLICY`
        """
        return pulumi.get(self, "type")


class AwaitableGetPolicyResult(GetPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyResult(
            arn=self.arn,
            aws_managed=self.aws_managed,
            content=self.content,
            description=self.description,
            id=self.id,
            name=self.name,
            policy_id=self.policy_id,
            type=self.type)


def get_policy(policy_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyResult:
    """
    Data source for managing an AWS Organizations Policy.

    ## Example Usage


    :param str policy_id: The unique identifier (ID) of the policy that you want more details on. Policy id starts with a "p-" followed by 8-28 lowercase or uppercase letters, digits, and underscores.
    """
    __args__ = dict()
    __args__['policyId'] = policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:organizations/getPolicy:getPolicy', __args__, opts=opts, typ=GetPolicyResult).value

    return AwaitableGetPolicyResult(
        arn=__ret__.arn,
        aws_managed=__ret__.aws_managed,
        content=__ret__.content,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        policy_id=__ret__.policy_id,
        type=__ret__.type)


@_utilities.lift_output_func(get_policy)
def get_policy_output(policy_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPolicyResult]:
    """
    Data source for managing an AWS Organizations Policy.

    ## Example Usage


    :param str policy_id: The unique identifier (ID) of the policy that you want more details on. Policy id starts with a "p-" followed by 8-28 lowercase or uppercase letters, digits, and underscores.
    """
    ...
