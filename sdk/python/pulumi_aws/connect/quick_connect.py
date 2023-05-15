# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['QuickConnectArgs', 'QuickConnect']

@pulumi.input_type
class QuickConnectArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 quick_connect_config: pulumi.Input['QuickConnectQuickConnectConfigArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a QuickConnect resource.
        :param pulumi.Input[str] instance_id: Specifies the identifier of the hosting Amazon Connect Instance.
        :param pulumi.Input['QuickConnectQuickConnectConfigArgs'] quick_connect_config: A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
        :param pulumi.Input[str] description: Specifies the description of the Quick Connect.
        :param pulumi.Input[str] name: Specifies the name of the Quick Connect.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "quick_connect_config", quick_connect_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the identifier of the hosting Amazon Connect Instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="quickConnectConfig")
    def quick_connect_config(self) -> pulumi.Input['QuickConnectQuickConnectConfigArgs']:
        """
        A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
        """
        return pulumi.get(self, "quick_connect_config")

    @quick_connect_config.setter
    def quick_connect_config(self, value: pulumi.Input['QuickConnectQuickConnectConfigArgs']):
        pulumi.set(self, "quick_connect_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the Quick Connect.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Quick Connect.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _QuickConnectState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 quick_connect_config: Optional[pulumi.Input['QuickConnectQuickConnectConfigArgs']] = None,
                 quick_connect_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering QuickConnect resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Quick Connect.
        :param pulumi.Input[str] description: Specifies the description of the Quick Connect.
        :param pulumi.Input[str] instance_id: Specifies the identifier of the hosting Amazon Connect Instance.
        :param pulumi.Input[str] name: Specifies the name of the Quick Connect.
        :param pulumi.Input['QuickConnectQuickConnectConfigArgs'] quick_connect_config: A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
        :param pulumi.Input[str] quick_connect_id: The identifier for the Quick Connect.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if quick_connect_config is not None:
            pulumi.set(__self__, "quick_connect_config", quick_connect_config)
        if quick_connect_id is not None:
            pulumi.set(__self__, "quick_connect_id", quick_connect_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the Quick Connect.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the Quick Connect.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the identifier of the hosting Amazon Connect Instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Quick Connect.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="quickConnectConfig")
    def quick_connect_config(self) -> Optional[pulumi.Input['QuickConnectQuickConnectConfigArgs']]:
        """
        A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
        """
        return pulumi.get(self, "quick_connect_config")

    @quick_connect_config.setter
    def quick_connect_config(self, value: Optional[pulumi.Input['QuickConnectQuickConnectConfigArgs']]):
        pulumi.set(self, "quick_connect_config", value)

    @property
    @pulumi.getter(name="quickConnectId")
    def quick_connect_id(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the Quick Connect.
        """
        return pulumi.get(self, "quick_connect_id")

    @quick_connect_id.setter
    def quick_connect_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quick_connect_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class QuickConnect(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 quick_connect_config: Optional[pulumi.Input[pulumi.InputType['QuickConnectQuickConnectConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an Amazon Connect Quick Connect resource. For more information see
        [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.connect.QuickConnect("test",
            description="quick connect phone number",
            instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
            quick_connect_config=aws.connect.QuickConnectQuickConnectConfigArgs(
                phone_configs=[aws.connect.QuickConnectQuickConnectConfigPhoneConfigArgs(
                    phone_number="+12345678912",
                )],
                quick_connect_type="PHONE_NUMBER",
            ),
            tags={
                "Name": "Example Quick Connect",
            })
        ```

        ## Import

        Amazon Connect Quick Connects can be imported using the `instance_id` and `quick_connect_id` separated by a colon (`:`), e.g.,

        ```sh
         $ pulumi import aws:connect/quickConnect:QuickConnect example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description of the Quick Connect.
        :param pulumi.Input[str] instance_id: Specifies the identifier of the hosting Amazon Connect Instance.
        :param pulumi.Input[str] name: Specifies the name of the Quick Connect.
        :param pulumi.Input[pulumi.InputType['QuickConnectQuickConnectConfigArgs']] quick_connect_config: A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QuickConnectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Amazon Connect Quick Connect resource. For more information see
        [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.connect.QuickConnect("test",
            description="quick connect phone number",
            instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
            quick_connect_config=aws.connect.QuickConnectQuickConnectConfigArgs(
                phone_configs=[aws.connect.QuickConnectQuickConnectConfigPhoneConfigArgs(
                    phone_number="+12345678912",
                )],
                quick_connect_type="PHONE_NUMBER",
            ),
            tags={
                "Name": "Example Quick Connect",
            })
        ```

        ## Import

        Amazon Connect Quick Connects can be imported using the `instance_id` and `quick_connect_id` separated by a colon (`:`), e.g.,

        ```sh
         $ pulumi import aws:connect/quickConnect:QuickConnect example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
        ```

        :param str resource_name: The name of the resource.
        :param QuickConnectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QuickConnectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 quick_connect_config: Optional[pulumi.Input[pulumi.InputType['QuickConnectQuickConnectConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QuickConnectArgs.__new__(QuickConnectArgs)

            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            if quick_connect_config is None and not opts.urn:
                raise TypeError("Missing required property 'quick_connect_config'")
            __props__.__dict__["quick_connect_config"] = quick_connect_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["quick_connect_id"] = None
            __props__.__dict__["tags_all"] = None
        super(QuickConnect, __self__).__init__(
            'aws:connect/quickConnect:QuickConnect',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            quick_connect_config: Optional[pulumi.Input[pulumi.InputType['QuickConnectQuickConnectConfigArgs']]] = None,
            quick_connect_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'QuickConnect':
        """
        Get an existing QuickConnect resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Quick Connect.
        :param pulumi.Input[str] description: Specifies the description of the Quick Connect.
        :param pulumi.Input[str] instance_id: Specifies the identifier of the hosting Amazon Connect Instance.
        :param pulumi.Input[str] name: Specifies the name of the Quick Connect.
        :param pulumi.Input[pulumi.InputType['QuickConnectQuickConnectConfigArgs']] quick_connect_config: A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
        :param pulumi.Input[str] quick_connect_id: The identifier for the Quick Connect.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QuickConnectState.__new__(_QuickConnectState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["quick_connect_config"] = quick_connect_config
        __props__.__dict__["quick_connect_id"] = quick_connect_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return QuickConnect(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Quick Connect.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the description of the Quick Connect.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the identifier of the hosting Amazon Connect Instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Quick Connect.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="quickConnectConfig")
    def quick_connect_config(self) -> pulumi.Output['outputs.QuickConnectQuickConnectConfig']:
        """
        A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
        """
        return pulumi.get(self, "quick_connect_config")

    @property
    @pulumi.getter(name="quickConnectId")
    def quick_connect_id(self) -> pulumi.Output[str]:
        """
        The identifier for the Quick Connect.
        """
        return pulumi.get(self, "quick_connect_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

