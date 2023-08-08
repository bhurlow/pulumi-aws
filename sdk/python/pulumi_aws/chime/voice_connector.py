# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VoiceConnectorArgs', 'VoiceConnector']

@pulumi.input_type
class VoiceConnectorArgs:
    def __init__(__self__, *,
                 require_encryption: pulumi.Input[bool],
                 aws_region: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VoiceConnector resource.
        :param pulumi.Input[bool] require_encryption: When enabled, requires encryption for the Amazon Chime Voice Connector.
               
               The following arguments are optional:
        :param pulumi.Input[str] aws_region: The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
        :param pulumi.Input[str] name: The name of the Amazon Chime Voice Connector.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "require_encryption", require_encryption)
        if aws_region is not None:
            pulumi.set(__self__, "aws_region", aws_region)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="requireEncryption")
    def require_encryption(self) -> pulumi.Input[bool]:
        """
        When enabled, requires encryption for the Amazon Chime Voice Connector.

        The following arguments are optional:
        """
        return pulumi.get(self, "require_encryption")

    @require_encryption.setter
    def require_encryption(self, value: pulumi.Input[bool]):
        pulumi.set(self, "require_encryption", value)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
        """
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VoiceConnectorState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 aws_region: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_host_name: Optional[pulumi.Input[str]] = None,
                 require_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering VoiceConnector resources.
        :param pulumi.Input[str] arn: ARN (Amazon Resource Name) of the Amazon Chime Voice Connector.
        :param pulumi.Input[str] aws_region: The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
        :param pulumi.Input[str] name: The name of the Amazon Chime Voice Connector.
        :param pulumi.Input[str] outbound_host_name: The outbound host name for the Amazon Chime Voice Connector.
        :param pulumi.Input[bool] require_encryption: When enabled, requires encryption for the Amazon Chime Voice Connector.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if aws_region is not None:
            pulumi.set(__self__, "aws_region", aws_region)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if outbound_host_name is not None:
            pulumi.set(__self__, "outbound_host_name", outbound_host_name)
        if require_encryption is not None:
            pulumi.set(__self__, "require_encryption", require_encryption)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN (Amazon Resource Name) of the Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
        """
        return pulumi.get(self, "aws_region")

    @aws_region.setter
    def aws_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_region", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outboundHostName")
    def outbound_host_name(self) -> Optional[pulumi.Input[str]]:
        """
        The outbound host name for the Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "outbound_host_name")

    @outbound_host_name.setter
    def outbound_host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outbound_host_name", value)

    @property
    @pulumi.getter(name="requireEncryption")
    def require_encryption(self) -> Optional[pulumi.Input[bool]]:
        """
        When enabled, requires encryption for the Amazon Chime Voice Connector.

        The following arguments are optional:
        """
        return pulumi.get(self, "require_encryption")

    @require_encryption.setter
    def require_encryption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_encryption", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class VoiceConnector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_region: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 require_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Enables you to connect your phone system to the telephone network at a substantial cost savings by using SIP trunking.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.chime.VoiceConnector("test",
            aws_region="us-east-1",
            require_encryption=True)
        ```

        ## Import

        terraform import {

         to = aws_chime_voice_connector.test

         id = "example" } Using `pulumi import`, import Configuration Recorder using the name. For exampleconsole % pulumi import aws_chime_voice_connector.test example

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_region: The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
        :param pulumi.Input[str] name: The name of the Amazon Chime Voice Connector.
        :param pulumi.Input[bool] require_encryption: When enabled, requires encryption for the Amazon Chime Voice Connector.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VoiceConnectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enables you to connect your phone system to the telephone network at a substantial cost savings by using SIP trunking.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.chime.VoiceConnector("test",
            aws_region="us-east-1",
            require_encryption=True)
        ```

        ## Import

        terraform import {

         to = aws_chime_voice_connector.test

         id = "example" } Using `pulumi import`, import Configuration Recorder using the name. For exampleconsole % pulumi import aws_chime_voice_connector.test example

        :param str resource_name: The name of the resource.
        :param VoiceConnectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VoiceConnectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_region: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 require_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VoiceConnectorArgs.__new__(VoiceConnectorArgs)

            __props__.__dict__["aws_region"] = aws_region
            __props__.__dict__["name"] = name
            if require_encryption is None and not opts.urn:
                raise TypeError("Missing required property 'require_encryption'")
            __props__.__dict__["require_encryption"] = require_encryption
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["outbound_host_name"] = None
            __props__.__dict__["tags_all"] = None
        super(VoiceConnector, __self__).__init__(
            'aws:chime/voiceConnector:VoiceConnector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_region: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            outbound_host_name: Optional[pulumi.Input[str]] = None,
            require_encryption: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'VoiceConnector':
        """
        Get an existing VoiceConnector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN (Amazon Resource Name) of the Amazon Chime Voice Connector.
        :param pulumi.Input[str] aws_region: The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
        :param pulumi.Input[str] name: The name of the Amazon Chime Voice Connector.
        :param pulumi.Input[str] outbound_host_name: The outbound host name for the Amazon Chime Voice Connector.
        :param pulumi.Input[bool] require_encryption: When enabled, requires encryption for the Amazon Chime Voice Connector.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VoiceConnectorState.__new__(_VoiceConnectorState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["aws_region"] = aws_region
        __props__.__dict__["name"] = name
        __props__.__dict__["outbound_host_name"] = outbound_host_name
        __props__.__dict__["require_encryption"] = require_encryption
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return VoiceConnector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN (Amazon Resource Name) of the Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsRegion")
    def aws_region(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
        """
        return pulumi.get(self, "aws_region")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundHostName")
    def outbound_host_name(self) -> pulumi.Output[str]:
        """
        The outbound host name for the Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "outbound_host_name")

    @property
    @pulumi.getter(name="requireEncryption")
    def require_encryption(self) -> pulumi.Output[bool]:
        """
        When enabled, requires encryption for the Amazon Chime Voice Connector.

        The following arguments are optional:
        """
        return pulumi.get(self, "require_encryption")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

