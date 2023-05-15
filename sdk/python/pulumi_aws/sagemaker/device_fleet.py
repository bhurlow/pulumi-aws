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

__all__ = ['DeviceFleetArgs', 'DeviceFleet']

@pulumi.input_type
class DeviceFleetArgs:
    def __init__(__self__, *,
                 device_fleet_name: pulumi.Input[str],
                 output_config: pulumi.Input['DeviceFleetOutputConfigArgs'],
                 role_arn: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 enable_iot_role_alias: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DeviceFleet resource.
        :param pulumi.Input[str] device_fleet_name: The name of the Device Fleet (must be unique).
        :param pulumi.Input['DeviceFleetOutputConfigArgs'] output_config: Specifies details about the repository. see Output Config details below.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
        :param pulumi.Input[str] description: A description of the fleet.
        :param pulumi.Input[bool] enable_iot_role_alias: Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "device_fleet_name", device_fleet_name)
        pulumi.set(__self__, "output_config", output_config)
        pulumi.set(__self__, "role_arn", role_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_iot_role_alias is not None:
            pulumi.set(__self__, "enable_iot_role_alias", enable_iot_role_alias)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="deviceFleetName")
    def device_fleet_name(self) -> pulumi.Input[str]:
        """
        The name of the Device Fleet (must be unique).
        """
        return pulumi.get(self, "device_fleet_name")

    @device_fleet_name.setter
    def device_fleet_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_fleet_name", value)

    @property
    @pulumi.getter(name="outputConfig")
    def output_config(self) -> pulumi.Input['DeviceFleetOutputConfigArgs']:
        """
        Specifies details about the repository. see Output Config details below.
        """
        return pulumi.get(self, "output_config")

    @output_config.setter
    def output_config(self, value: pulumi.Input['DeviceFleetOutputConfigArgs']):
        pulumi.set(self, "output_config", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the fleet.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableIotRoleAlias")
    def enable_iot_role_alias(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
        """
        return pulumi.get(self, "enable_iot_role_alias")

    @enable_iot_role_alias.setter
    def enable_iot_role_alias(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_iot_role_alias", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DeviceFleetState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_fleet_name: Optional[pulumi.Input[str]] = None,
                 enable_iot_role_alias: Optional[pulumi.Input[bool]] = None,
                 iot_role_alias: Optional[pulumi.Input[str]] = None,
                 output_config: Optional[pulumi.Input['DeviceFleetOutputConfigArgs']] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DeviceFleet resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
        :param pulumi.Input[str] description: A description of the fleet.
        :param pulumi.Input[str] device_fleet_name: The name of the Device Fleet (must be unique).
        :param pulumi.Input[bool] enable_iot_role_alias: Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
        :param pulumi.Input['DeviceFleetOutputConfigArgs'] output_config: Specifies details about the repository. see Output Config details below.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_fleet_name is not None:
            pulumi.set(__self__, "device_fleet_name", device_fleet_name)
        if enable_iot_role_alias is not None:
            pulumi.set(__self__, "enable_iot_role_alias", enable_iot_role_alias)
        if iot_role_alias is not None:
            pulumi.set(__self__, "iot_role_alias", iot_role_alias)
        if output_config is not None:
            pulumi.set(__self__, "output_config", output_config)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the fleet.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceFleetName")
    def device_fleet_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Device Fleet (must be unique).
        """
        return pulumi.get(self, "device_fleet_name")

    @device_fleet_name.setter
    def device_fleet_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_fleet_name", value)

    @property
    @pulumi.getter(name="enableIotRoleAlias")
    def enable_iot_role_alias(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
        """
        return pulumi.get(self, "enable_iot_role_alias")

    @enable_iot_role_alias.setter
    def enable_iot_role_alias(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_iot_role_alias", value)

    @property
    @pulumi.getter(name="iotRoleAlias")
    def iot_role_alias(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iot_role_alias")

    @iot_role_alias.setter
    def iot_role_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iot_role_alias", value)

    @property
    @pulumi.getter(name="outputConfig")
    def output_config(self) -> Optional[pulumi.Input['DeviceFleetOutputConfigArgs']]:
        """
        Specifies details about the repository. see Output Config details below.
        """
        return pulumi.get(self, "output_config")

    @output_config.setter
    def output_config(self, value: Optional[pulumi.Input['DeviceFleetOutputConfigArgs']]):
        pulumi.set(self, "output_config", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class DeviceFleet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_fleet_name: Optional[pulumi.Input[str]] = None,
                 enable_iot_role_alias: Optional[pulumi.Input[bool]] = None,
                 output_config: Optional[pulumi.Input[pulumi.InputType['DeviceFleetOutputConfigArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a SageMaker Device Fleet resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.DeviceFleet("example",
            device_fleet_name="example",
            role_arn=aws_iam_role["test"]["arn"],
            output_config=aws.sagemaker.DeviceFleetOutputConfigArgs(
                s3_output_location=f"s3://{aws_s3_bucket['example']['bucket']}/prefix/",
            ))
        ```

        ## Import

        SageMaker Device Fleets can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:sagemaker/deviceFleet:DeviceFleet example my-fleet
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the fleet.
        :param pulumi.Input[str] device_fleet_name: The name of the Device Fleet (must be unique).
        :param pulumi.Input[bool] enable_iot_role_alias: Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
        :param pulumi.Input[pulumi.InputType['DeviceFleetOutputConfigArgs']] output_config: Specifies details about the repository. see Output Config details below.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceFleetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a SageMaker Device Fleet resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.DeviceFleet("example",
            device_fleet_name="example",
            role_arn=aws_iam_role["test"]["arn"],
            output_config=aws.sagemaker.DeviceFleetOutputConfigArgs(
                s3_output_location=f"s3://{aws_s3_bucket['example']['bucket']}/prefix/",
            ))
        ```

        ## Import

        SageMaker Device Fleets can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:sagemaker/deviceFleet:DeviceFleet example my-fleet
        ```

        :param str resource_name: The name of the resource.
        :param DeviceFleetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceFleetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_fleet_name: Optional[pulumi.Input[str]] = None,
                 enable_iot_role_alias: Optional[pulumi.Input[bool]] = None,
                 output_config: Optional[pulumi.Input[pulumi.InputType['DeviceFleetOutputConfigArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceFleetArgs.__new__(DeviceFleetArgs)

            __props__.__dict__["description"] = description
            if device_fleet_name is None and not opts.urn:
                raise TypeError("Missing required property 'device_fleet_name'")
            __props__.__dict__["device_fleet_name"] = device_fleet_name
            __props__.__dict__["enable_iot_role_alias"] = enable_iot_role_alias
            if output_config is None and not opts.urn:
                raise TypeError("Missing required property 'output_config'")
            __props__.__dict__["output_config"] = output_config
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["iot_role_alias"] = None
            __props__.__dict__["tags_all"] = None
        super(DeviceFleet, __self__).__init__(
            'aws:sagemaker/deviceFleet:DeviceFleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            device_fleet_name: Optional[pulumi.Input[str]] = None,
            enable_iot_role_alias: Optional[pulumi.Input[bool]] = None,
            iot_role_alias: Optional[pulumi.Input[str]] = None,
            output_config: Optional[pulumi.Input[pulumi.InputType['DeviceFleetOutputConfigArgs']]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'DeviceFleet':
        """
        Get an existing DeviceFleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
        :param pulumi.Input[str] description: A description of the fleet.
        :param pulumi.Input[str] device_fleet_name: The name of the Device Fleet (must be unique).
        :param pulumi.Input[bool] enable_iot_role_alias: Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
        :param pulumi.Input[pulumi.InputType['DeviceFleetOutputConfigArgs']] output_config: Specifies details about the repository. see Output Config details below.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceFleetState.__new__(_DeviceFleetState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["device_fleet_name"] = device_fleet_name
        __props__.__dict__["enable_iot_role_alias"] = enable_iot_role_alias
        __props__.__dict__["iot_role_alias"] = iot_role_alias
        __props__.__dict__["output_config"] = output_config
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return DeviceFleet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the fleet.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceFleetName")
    def device_fleet_name(self) -> pulumi.Output[str]:
        """
        The name of the Device Fleet (must be unique).
        """
        return pulumi.get(self, "device_fleet_name")

    @property
    @pulumi.getter(name="enableIotRoleAlias")
    def enable_iot_role_alias(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
        """
        return pulumi.get(self, "enable_iot_role_alias")

    @property
    @pulumi.getter(name="iotRoleAlias")
    def iot_role_alias(self) -> pulumi.Output[str]:
        return pulumi.get(self, "iot_role_alias")

    @property
    @pulumi.getter(name="outputConfig")
    def output_config(self) -> pulumi.Output['outputs.DeviceFleetOutputConfig']:
        """
        Specifies details about the repository. see Output Config details below.
        """
        return pulumi.get(self, "output_config")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

