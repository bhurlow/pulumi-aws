# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['PlacementGroupArgs', 'PlacementGroup']

@pulumi.input_type
class PlacementGroupArgs:
    def __init__(__self__, *,
                 strategy: pulumi.Input[Union[str, 'PlacementStrategy']],
                 name: Optional[pulumi.Input[str]] = None,
                 partition_count: Optional[pulumi.Input[int]] = None,
                 spread_level: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a PlacementGroup resource.
        :param pulumi.Input[Union[str, 'PlacementStrategy']] strategy: The placement strategy. Can be `cluster`, `partition` or `spread`.
        :param pulumi.Input[str] name: The name of the placement group.
        :param pulumi.Input[int] partition_count: The number of partitions to create in the
               placement group.  Can only be specified when the `strategy` is set to
               `partition`.  Valid values are 1 - 7 (default is `2`).
        :param pulumi.Input[str] spread_level: Determines how placement groups spread instances. Can only be used
               when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "strategy", strategy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partition_count is not None:
            pulumi.set(__self__, "partition_count", partition_count)
        if spread_level is not None:
            pulumi.set(__self__, "spread_level", spread_level)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def strategy(self) -> pulumi.Input[Union[str, 'PlacementStrategy']]:
        """
        The placement strategy. Can be `cluster`, `partition` or `spread`.
        """
        return pulumi.get(self, "strategy")

    @strategy.setter
    def strategy(self, value: pulumi.Input[Union[str, 'PlacementStrategy']]):
        pulumi.set(self, "strategy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the placement group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of partitions to create in the
        placement group.  Can only be specified when the `strategy` is set to
        `partition`.  Valid values are 1 - 7 (default is `2`).
        """
        return pulumi.get(self, "partition_count")

    @partition_count.setter
    def partition_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "partition_count", value)

    @property
    @pulumi.getter(name="spreadLevel")
    def spread_level(self) -> Optional[pulumi.Input[str]]:
        """
        Determines how placement groups spread instances. Can only be used
        when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
        """
        return pulumi.get(self, "spread_level")

    @spread_level.setter
    def spread_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spread_level", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _PlacementGroupState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_count: Optional[pulumi.Input[int]] = None,
                 placement_group_id: Optional[pulumi.Input[str]] = None,
                 spread_level: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[Union[str, 'PlacementStrategy']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering PlacementGroup resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the placement group.
        :param pulumi.Input[str] name: The name of the placement group.
        :param pulumi.Input[int] partition_count: The number of partitions to create in the
               placement group.  Can only be specified when the `strategy` is set to
               `partition`.  Valid values are 1 - 7 (default is `2`).
        :param pulumi.Input[str] placement_group_id: The ID of the placement group.
        :param pulumi.Input[str] spread_level: Determines how placement groups spread instances. Can only be used
               when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
        :param pulumi.Input[Union[str, 'PlacementStrategy']] strategy: The placement strategy. Can be `cluster`, `partition` or `spread`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partition_count is not None:
            pulumi.set(__self__, "partition_count", partition_count)
        if placement_group_id is not None:
            pulumi.set(__self__, "placement_group_id", placement_group_id)
        if spread_level is not None:
            pulumi.set(__self__, "spread_level", spread_level)
        if strategy is not None:
            pulumi.set(__self__, "strategy", strategy)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the placement group.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the placement group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of partitions to create in the
        placement group.  Can only be specified when the `strategy` is set to
        `partition`.  Valid values are 1 - 7 (default is `2`).
        """
        return pulumi.get(self, "partition_count")

    @partition_count.setter
    def partition_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "partition_count", value)

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the placement group.
        """
        return pulumi.get(self, "placement_group_id")

    @placement_group_id.setter
    def placement_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "placement_group_id", value)

    @property
    @pulumi.getter(name="spreadLevel")
    def spread_level(self) -> Optional[pulumi.Input[str]]:
        """
        Determines how placement groups spread instances. Can only be used
        when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
        """
        return pulumi.get(self, "spread_level")

    @spread_level.setter
    def spread_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spread_level", value)

    @property
    @pulumi.getter
    def strategy(self) -> Optional[pulumi.Input[Union[str, 'PlacementStrategy']]]:
        """
        The placement strategy. Can be `cluster`, `partition` or `spread`.
        """
        return pulumi.get(self, "strategy")

    @strategy.setter
    def strategy(self, value: Optional[pulumi.Input[Union[str, 'PlacementStrategy']]]):
        pulumi.set(self, "strategy", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class PlacementGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_count: Optional[pulumi.Input[int]] = None,
                 spread_level: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[Union[str, 'PlacementStrategy']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an EC2 placement group. Read more about placement groups
        in [AWS Docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        web = aws.ec2.PlacementGroup("web", strategy="cluster")
        ```

        ## Import

        terraform import {

         to = aws_placement_group.prod_pg

         id = "production-placement-group" } Using `pulumi import`, import placement groups using the `name`. For exampleconsole % pulumi import aws_placement_group.prod_pg production-placement-group

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the placement group.
        :param pulumi.Input[int] partition_count: The number of partitions to create in the
               placement group.  Can only be specified when the `strategy` is set to
               `partition`.  Valid values are 1 - 7 (default is `2`).
        :param pulumi.Input[str] spread_level: Determines how placement groups spread instances. Can only be used
               when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
        :param pulumi.Input[Union[str, 'PlacementStrategy']] strategy: The placement strategy. Can be `cluster`, `partition` or `spread`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PlacementGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an EC2 placement group. Read more about placement groups
        in [AWS Docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        web = aws.ec2.PlacementGroup("web", strategy="cluster")
        ```

        ## Import

        terraform import {

         to = aws_placement_group.prod_pg

         id = "production-placement-group" } Using `pulumi import`, import placement groups using the `name`. For exampleconsole % pulumi import aws_placement_group.prod_pg production-placement-group

        :param str resource_name: The name of the resource.
        :param PlacementGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PlacementGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_count: Optional[pulumi.Input[int]] = None,
                 spread_level: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[Union[str, 'PlacementStrategy']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PlacementGroupArgs.__new__(PlacementGroupArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["partition_count"] = partition_count
            __props__.__dict__["spread_level"] = spread_level
            if strategy is None and not opts.urn:
                raise TypeError("Missing required property 'strategy'")
            __props__.__dict__["strategy"] = strategy
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["placement_group_id"] = None
            __props__.__dict__["tags_all"] = None
        super(PlacementGroup, __self__).__init__(
            'aws:ec2/placementGroup:PlacementGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition_count: Optional[pulumi.Input[int]] = None,
            placement_group_id: Optional[pulumi.Input[str]] = None,
            spread_level: Optional[pulumi.Input[str]] = None,
            strategy: Optional[pulumi.Input[Union[str, 'PlacementStrategy']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'PlacementGroup':
        """
        Get an existing PlacementGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the placement group.
        :param pulumi.Input[str] name: The name of the placement group.
        :param pulumi.Input[int] partition_count: The number of partitions to create in the
               placement group.  Can only be specified when the `strategy` is set to
               `partition`.  Valid values are 1 - 7 (default is `2`).
        :param pulumi.Input[str] placement_group_id: The ID of the placement group.
        :param pulumi.Input[str] spread_level: Determines how placement groups spread instances. Can only be used
               when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
        :param pulumi.Input[Union[str, 'PlacementStrategy']] strategy: The placement strategy. Can be `cluster`, `partition` or `spread`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PlacementGroupState.__new__(_PlacementGroupState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["name"] = name
        __props__.__dict__["partition_count"] = partition_count
        __props__.__dict__["placement_group_id"] = placement_group_id
        __props__.__dict__["spread_level"] = spread_level
        __props__.__dict__["strategy"] = strategy
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return PlacementGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the placement group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the placement group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> pulumi.Output[int]:
        """
        The number of partitions to create in the
        placement group.  Can only be specified when the `strategy` is set to
        `partition`.  Valid values are 1 - 7 (default is `2`).
        """
        return pulumi.get(self, "partition_count")

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the placement group.
        """
        return pulumi.get(self, "placement_group_id")

    @property
    @pulumi.getter(name="spreadLevel")
    def spread_level(self) -> pulumi.Output[str]:
        """
        Determines how placement groups spread instances. Can only be used
        when the `strategy` is set to `spread`. Can be `host` or `rack`. `host` can only be used for Outpost placement groups. Defaults to `rack`.
        """
        return pulumi.get(self, "spread_level")

    @property
    @pulumi.getter
    def strategy(self) -> pulumi.Output[str]:
        """
        The placement strategy. Can be `cluster`, `partition` or `spread`.
        """
        return pulumi.get(self, "strategy")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

