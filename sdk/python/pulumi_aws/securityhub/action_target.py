# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ActionTargetArgs', 'ActionTarget']

@pulumi.input_type
class ActionTargetArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 identifier: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ActionTarget resource.
        :param pulumi.Input[str] description: The name of the custom action target.
        :param pulumi.Input[str] identifier: The ID for the custom action target.
        :param pulumi.Input[str] name: The description for the custom action target.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "identifier", identifier)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The name of the custom action target.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Input[str]:
        """
        The ID for the custom action target.
        """
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the custom action target.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ActionTargetState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ActionTarget resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Security Hub custom action target.
        :param pulumi.Input[str] description: The name of the custom action target.
        :param pulumi.Input[str] identifier: The ID for the custom action target.
        :param pulumi.Input[str] name: The description for the custom action target.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if identifier is not None:
            pulumi.set(__self__, "identifier", identifier)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Security Hub custom action target.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the custom action target.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The ID for the custom action target.
        """
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the custom action target.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ActionTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates Security Hub custom action.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_account = aws.securityhub.Account("exampleAccount")
        example_action_target = aws.securityhub.ActionTarget("exampleActionTarget",
            identifier="SendToChat",
            description="This is custom action sends selected findings to chat",
            opts=pulumi.ResourceOptions(depends_on=[example_account]))
        ```

        ## Import

        terraform import {

         to = aws_securityhub_action_target.example

         id = "arn:aws:securityhub:eu-west-1:312940875350:action/custom/a" } Using `pulumi import`, import Security Hub custom action using the action target ARN. For exampleconsole % pulumi import aws_securityhub_action_target.example arn:aws:securityhub:eu-west-1:312940875350:action/custom/a

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The name of the custom action target.
        :param pulumi.Input[str] identifier: The ID for the custom action target.
        :param pulumi.Input[str] name: The description for the custom action target.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates Security Hub custom action.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_account = aws.securityhub.Account("exampleAccount")
        example_action_target = aws.securityhub.ActionTarget("exampleActionTarget",
            identifier="SendToChat",
            description="This is custom action sends selected findings to chat",
            opts=pulumi.ResourceOptions(depends_on=[example_account]))
        ```

        ## Import

        terraform import {

         to = aws_securityhub_action_target.example

         id = "arn:aws:securityhub:eu-west-1:312940875350:action/custom/a" } Using `pulumi import`, import Security Hub custom action using the action target ARN. For exampleconsole % pulumi import aws_securityhub_action_target.example arn:aws:securityhub:eu-west-1:312940875350:action/custom/a

        :param str resource_name: The name of the resource.
        :param ActionTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActionTargetArgs.__new__(ActionTargetArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if identifier is None and not opts.urn:
                raise TypeError("Missing required property 'identifier'")
            __props__.__dict__["identifier"] = identifier
            __props__.__dict__["name"] = name
            __props__.__dict__["arn"] = None
        super(ActionTarget, __self__).__init__(
            'aws:securityhub/actionTarget:ActionTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ActionTarget':
        """
        Get an existing ActionTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Security Hub custom action target.
        :param pulumi.Input[str] description: The name of the custom action target.
        :param pulumi.Input[str] identifier: The ID for the custom action target.
        :param pulumi.Input[str] name: The description for the custom action target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionTargetState.__new__(_ActionTargetState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["identifier"] = identifier
        __props__.__dict__["name"] = name
        return ActionTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Security Hub custom action target.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The name of the custom action target.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        """
        The ID for the custom action target.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The description for the custom action target.
        """
        return pulumi.get(self, "name")

