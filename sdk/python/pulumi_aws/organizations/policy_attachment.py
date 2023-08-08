# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PolicyAttachmentArgs', 'PolicyAttachment']

@pulumi.input_type
class PolicyAttachmentArgs:
    def __init__(__self__, *,
                 policy_id: pulumi.Input[str],
                 target_id: pulumi.Input[str],
                 skip_destroy: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a PolicyAttachment resource.
        :param pulumi.Input[str] policy_id: The unique identifier (ID) of the policy that you want to attach to the target.
        :param pulumi.Input[str] target_id: The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
        :param pulumi.Input[bool] skip_destroy: If set to `true`, destroy will **not** detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
        """
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "target_id", target_id)
        if skip_destroy is not None:
            pulumi.set(__self__, "skip_destroy", skip_destroy)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        The unique identifier (ID) of the policy that you want to attach to the target.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Input[str]:
        """
        The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter(name="skipDestroy")
    def skip_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, destroy will **not** detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
        """
        return pulumi.get(self, "skip_destroy")

    @skip_destroy.setter
    def skip_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_destroy", value)


@pulumi.input_type
class _PolicyAttachmentState:
    def __init__(__self__, *,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 skip_destroy: Optional[pulumi.Input[bool]] = None,
                 target_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyAttachment resources.
        :param pulumi.Input[str] policy_id: The unique identifier (ID) of the policy that you want to attach to the target.
        :param pulumi.Input[bool] skip_destroy: If set to `true`, destroy will **not** detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
        :param pulumi.Input[str] target_id: The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
        """
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if skip_destroy is not None:
            pulumi.set(__self__, "skip_destroy", skip_destroy)
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier (ID) of the policy that you want to attach to the target.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="skipDestroy")
    def skip_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, destroy will **not** detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
        """
        return pulumi.get(self, "skip_destroy")

    @skip_destroy.setter
    def skip_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_destroy", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_id", value)


class PolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 skip_destroy: Optional[pulumi.Input[bool]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to attach an AWS Organizations policy to an organization account, root, or unit.

        ## Example Usage
        ### Organization Account

        ```python
        import pulumi
        import pulumi_aws as aws

        account = aws.organizations.PolicyAttachment("account",
            policy_id=aws_organizations_policy["example"]["id"],
            target_id="123456789012")
        ```
        ### Organization Root

        ```python
        import pulumi
        import pulumi_aws as aws

        root = aws.organizations.PolicyAttachment("root",
            policy_id=aws_organizations_policy["example"]["id"],
            target_id=aws_organizations_organization["example"]["roots"][0]["id"])
        ```
        ### Organization Unit

        ```python
        import pulumi
        import pulumi_aws as aws

        unit = aws.organizations.PolicyAttachment("unit",
            policy_id=aws_organizations_policy["example"]["id"],
            target_id=aws_organizations_organizational_unit["example"]["id"])
        ```

        ## Import

        With an account targetterraform import {

         to = aws_organizations_policy_attachment.account

         id = "123456789012:p-12345678" } Using `pulumi import`, import `aws_organizations_policy_attachment` using the target ID and policy ID. For exampleWith an account targetconsole % pulumi import aws_organizations_policy_attachment.account 123456789012:p-12345678

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_id: The unique identifier (ID) of the policy that you want to attach to the target.
        :param pulumi.Input[bool] skip_destroy: If set to `true`, destroy will **not** detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
        :param pulumi.Input[str] target_id: The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to attach an AWS Organizations policy to an organization account, root, or unit.

        ## Example Usage
        ### Organization Account

        ```python
        import pulumi
        import pulumi_aws as aws

        account = aws.organizations.PolicyAttachment("account",
            policy_id=aws_organizations_policy["example"]["id"],
            target_id="123456789012")
        ```
        ### Organization Root

        ```python
        import pulumi
        import pulumi_aws as aws

        root = aws.organizations.PolicyAttachment("root",
            policy_id=aws_organizations_policy["example"]["id"],
            target_id=aws_organizations_organization["example"]["roots"][0]["id"])
        ```
        ### Organization Unit

        ```python
        import pulumi
        import pulumi_aws as aws

        unit = aws.organizations.PolicyAttachment("unit",
            policy_id=aws_organizations_policy["example"]["id"],
            target_id=aws_organizations_organizational_unit["example"]["id"])
        ```

        ## Import

        With an account targetterraform import {

         to = aws_organizations_policy_attachment.account

         id = "123456789012:p-12345678" } Using `pulumi import`, import `aws_organizations_policy_attachment` using the target ID and policy ID. For exampleWith an account targetconsole % pulumi import aws_organizations_policy_attachment.account 123456789012:p-12345678

        :param str resource_name: The name of the resource.
        :param PolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 skip_destroy: Optional[pulumi.Input[bool]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyAttachmentArgs.__new__(PolicyAttachmentArgs)

            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            __props__.__dict__["skip_destroy"] = skip_destroy
            if target_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_id'")
            __props__.__dict__["target_id"] = target_id
        super(PolicyAttachment, __self__).__init__(
            'aws:organizations/policyAttachment:PolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            skip_destroy: Optional[pulumi.Input[bool]] = None,
            target_id: Optional[pulumi.Input[str]] = None) -> 'PolicyAttachment':
        """
        Get an existing PolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_id: The unique identifier (ID) of the policy that you want to attach to the target.
        :param pulumi.Input[bool] skip_destroy: If set to `true`, destroy will **not** detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
        :param pulumi.Input[str] target_id: The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyAttachmentState.__new__(_PolicyAttachmentState)

        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["skip_destroy"] = skip_destroy
        __props__.__dict__["target_id"] = target_id
        return PolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The unique identifier (ID) of the policy that you want to attach to the target.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="skipDestroy")
    def skip_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, destroy will **not** detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
        """
        return pulumi.get(self, "skip_destroy")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Output[str]:
        """
        The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
        """
        return pulumi.get(self, "target_id")

