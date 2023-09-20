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

__all__ = ['CapacityProviderArgs', 'CapacityProvider']

@pulumi.input_type
class CapacityProviderArgs:
    def __init__(__self__, *,
                 auto_scaling_group_provider: pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs'],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CapacityProvider resource.
        :param pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs'] auto_scaling_group_provider: Configuration block for the provider for the ECS auto scaling group. Detailed below.
        :param pulumi.Input[str] name: Name of the capacity provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "auto_scaling_group_provider", auto_scaling_group_provider)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="autoScalingGroupProvider")
    def auto_scaling_group_provider(self) -> pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs']:
        """
        Configuration block for the provider for the ECS auto scaling group. Detailed below.
        """
        return pulumi.get(self, "auto_scaling_group_provider")

    @auto_scaling_group_provider.setter
    def auto_scaling_group_provider(self, value: pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs']):
        pulumi.set(self, "auto_scaling_group_provider", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the capacity provider.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _CapacityProviderState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 auto_scaling_group_provider: Optional[pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering CapacityProvider resources.
        :param pulumi.Input[str] arn: ARN that identifies the capacity provider.
        :param pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs'] auto_scaling_group_provider: Configuration block for the provider for the ECS auto scaling group. Detailed below.
        :param pulumi.Input[str] name: Name of the capacity provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if auto_scaling_group_provider is not None:
            pulumi.set(__self__, "auto_scaling_group_provider", auto_scaling_group_provider)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN that identifies the capacity provider.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="autoScalingGroupProvider")
    def auto_scaling_group_provider(self) -> Optional[pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs']]:
        """
        Configuration block for the provider for the ECS auto scaling group. Detailed below.
        """
        return pulumi.get(self, "auto_scaling_group_provider")

    @auto_scaling_group_provider.setter
    def auto_scaling_group_provider(self, value: Optional[pulumi.Input['CapacityProviderAutoScalingGroupProviderArgs']]):
        pulumi.set(self, "auto_scaling_group_provider", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the capacity provider.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class CapacityProvider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_provider: Optional[pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).

        > **NOTE:** Associating an ECS Capacity Provider to an Auto Scaling Group will automatically add the `AmazonECSManaged` tag to the Auto Scaling Group. This tag should be included in the `autoscaling.Group` resource configuration to prevent the provider from removing it in subsequent executions as well as ensuring the `AmazonECSManaged` tag is propagated to all EC2 Instances in the Auto Scaling Group if `min_size` is above 0 on creation. Any EC2 Instances in the Auto Scaling Group without this tag must be manually be updated, otherwise they may cause unexpected scaling behavior and metrics.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configuration, including potentially other tags ...
        test_group = aws.autoscaling.Group("testGroup", tags=[aws.autoscaling.GroupTagArgs(
            key="AmazonECSManaged",
            value="true",
            propagate_at_launch=True,
        )])
        test_capacity_provider = aws.ecs.CapacityProvider("testCapacityProvider", auto_scaling_group_provider=aws.ecs.CapacityProviderAutoScalingGroupProviderArgs(
            auto_scaling_group_arn=test_group.arn,
            managed_termination_protection="ENABLED",
            managed_scaling=aws.ecs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs(
                maximum_scaling_step_size=1000,
                minimum_scaling_step_size=1,
                status="ENABLED",
                target_capacity=10,
            ),
        ))
        ```

        ## Import

        Using `pulumi import`, import ECS Capacity Providers using the `name`. For example:

        ```sh
         $ pulumi import aws:ecs/capacityProvider:CapacityProvider example example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']] auto_scaling_group_provider: Configuration block for the provider for the ECS auto scaling group. Detailed below.
        :param pulumi.Input[str] name: Name of the capacity provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CapacityProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).

        > **NOTE:** Associating an ECS Capacity Provider to an Auto Scaling Group will automatically add the `AmazonECSManaged` tag to the Auto Scaling Group. This tag should be included in the `autoscaling.Group` resource configuration to prevent the provider from removing it in subsequent executions as well as ensuring the `AmazonECSManaged` tag is propagated to all EC2 Instances in the Auto Scaling Group if `min_size` is above 0 on creation. Any EC2 Instances in the Auto Scaling Group without this tag must be manually be updated, otherwise they may cause unexpected scaling behavior and metrics.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configuration, including potentially other tags ...
        test_group = aws.autoscaling.Group("testGroup", tags=[aws.autoscaling.GroupTagArgs(
            key="AmazonECSManaged",
            value="true",
            propagate_at_launch=True,
        )])
        test_capacity_provider = aws.ecs.CapacityProvider("testCapacityProvider", auto_scaling_group_provider=aws.ecs.CapacityProviderAutoScalingGroupProviderArgs(
            auto_scaling_group_arn=test_group.arn,
            managed_termination_protection="ENABLED",
            managed_scaling=aws.ecs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs(
                maximum_scaling_step_size=1000,
                minimum_scaling_step_size=1,
                status="ENABLED",
                target_capacity=10,
            ),
        ))
        ```

        ## Import

        Using `pulumi import`, import ECS Capacity Providers using the `name`. For example:

        ```sh
         $ pulumi import aws:ecs/capacityProvider:CapacityProvider example example
        ```

        :param str resource_name: The name of the resource.
        :param CapacityProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CapacityProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_provider: Optional[pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CapacityProviderArgs.__new__(CapacityProviderArgs)

            if auto_scaling_group_provider is None and not opts.urn:
                raise TypeError("Missing required property 'auto_scaling_group_provider'")
            __props__.__dict__["auto_scaling_group_provider"] = auto_scaling_group_provider
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(CapacityProvider, __self__).__init__(
            'aws:ecs/capacityProvider:CapacityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            auto_scaling_group_provider: Optional[pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'CapacityProvider':
        """
        Get an existing CapacityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN that identifies the capacity provider.
        :param pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']] auto_scaling_group_provider: Configuration block for the provider for the ECS auto scaling group. Detailed below.
        :param pulumi.Input[str] name: Name of the capacity provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CapacityProviderState.__new__(_CapacityProviderState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["auto_scaling_group_provider"] = auto_scaling_group_provider
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return CapacityProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN that identifies the capacity provider.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoScalingGroupProvider")
    def auto_scaling_group_provider(self) -> pulumi.Output['outputs.CapacityProviderAutoScalingGroupProvider']:
        """
        Configuration block for the provider for the ECS auto scaling group. Detailed below.
        """
        return pulumi.get(self, "auto_scaling_group_provider")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the capacity provider.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

