# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DedicatedHostArgs', 'DedicatedHost']

@pulumi.input_type
class DedicatedHostArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[str],
                 asset_id: Optional[pulumi.Input[str]] = None,
                 auto_placement: Optional[pulumi.Input[str]] = None,
                 host_recovery: Optional[pulumi.Input[str]] = None,
                 instance_family: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 outpost_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DedicatedHost resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone in which to allocate the Dedicated Host.
        :param pulumi.Input[str] asset_id: The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        :param pulumi.Input[str] auto_placement: Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        :param pulumi.Input[str] host_recovery: Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        :param pulumi.Input[str] instance_family: Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] instance_type: Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] outpost_arn: The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        if asset_id is not None:
            pulumi.set(__self__, "asset_id", asset_id)
        if auto_placement is not None:
            pulumi.set(__self__, "auto_placement", auto_placement)
        if host_recovery is not None:
            pulumi.set(__self__, "host_recovery", host_recovery)
        if instance_family is not None:
            pulumi.set(__self__, "instance_family", instance_family)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if outpost_arn is not None:
            pulumi.set(__self__, "outpost_arn", outpost_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        """
        The Availability Zone in which to allocate the Dedicated Host.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="assetId")
    def asset_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        """
        return pulumi.get(self, "asset_id")

    @asset_id.setter
    def asset_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asset_id", value)

    @property
    @pulumi.getter(name="autoPlacement")
    def auto_placement(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        """
        return pulumi.get(self, "auto_placement")

    @auto_placement.setter
    def auto_placement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_placement", value)

    @property
    @pulumi.getter(name="hostRecovery")
    def host_recovery(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        """
        return pulumi.get(self, "host_recovery")

    @host_recovery.setter
    def host_recovery(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_recovery", value)

    @property
    @pulumi.getter(name="instanceFamily")
    def instance_family(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        """
        return pulumi.get(self, "instance_family")

    @instance_family.setter
    def instance_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_family", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        """
        return pulumi.get(self, "outpost_arn")

    @outpost_arn.setter
    def outpost_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outpost_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DedicatedHostState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 asset_id: Optional[pulumi.Input[str]] = None,
                 auto_placement: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 host_recovery: Optional[pulumi.Input[str]] = None,
                 instance_family: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 outpost_arn: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DedicatedHost resources.
        :param pulumi.Input[str] arn: The ARN of the Dedicated Host.
        :param pulumi.Input[str] asset_id: The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        :param pulumi.Input[str] auto_placement: Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        :param pulumi.Input[str] availability_zone: The Availability Zone in which to allocate the Dedicated Host.
        :param pulumi.Input[str] host_recovery: Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        :param pulumi.Input[str] instance_family: Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] instance_type: Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] outpost_arn: The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the Dedicated Host.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if asset_id is not None:
            pulumi.set(__self__, "asset_id", asset_id)
        if auto_placement is not None:
            pulumi.set(__self__, "auto_placement", auto_placement)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if host_recovery is not None:
            pulumi.set(__self__, "host_recovery", host_recovery)
        if instance_family is not None:
            pulumi.set(__self__, "instance_family", instance_family)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if outpost_arn is not None:
            pulumi.set(__self__, "outpost_arn", outpost_arn)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
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
        The ARN of the Dedicated Host.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="assetId")
    def asset_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        """
        return pulumi.get(self, "asset_id")

    @asset_id.setter
    def asset_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asset_id", value)

    @property
    @pulumi.getter(name="autoPlacement")
    def auto_placement(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        """
        return pulumi.get(self, "auto_placement")

    @auto_placement.setter
    def auto_placement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_placement", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The Availability Zone in which to allocate the Dedicated Host.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="hostRecovery")
    def host_recovery(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        """
        return pulumi.get(self, "host_recovery")

    @host_recovery.setter
    def host_recovery(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_recovery", value)

    @property
    @pulumi.getter(name="instanceFamily")
    def instance_family(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        """
        return pulumi.get(self, "instance_family")

    @instance_family.setter
    def instance_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_family", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        """
        return pulumi.get(self, "outpost_arn")

    @outpost_arn.setter
    def outpost_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outpost_arn", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the AWS account that owns the Dedicated Host.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class DedicatedHost(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_id: Optional[pulumi.Input[str]] = None,
                 auto_placement: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 host_recovery: Optional[pulumi.Input[str]] = None,
                 instance_family: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 outpost_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an EC2 Host resource. This allows Dedicated Hosts to be allocated, modified, and released.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new host with instance type of c5.18xlarge with Auto Placement
        # and Host Recovery enabled.
        test = aws.ec2.DedicatedHost("test",
            auto_placement="on",
            availability_zone="us-west-2a",
            host_recovery="on",
            instance_type="c5.18xlarge")
        ```

        ## Import

        Using `pulumi import`, import hosts using the host `id`. For example:

        ```sh
         $ pulumi import aws:ec2/dedicatedHost:DedicatedHost example h-0385a99d0e4b20cbb
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asset_id: The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        :param pulumi.Input[str] auto_placement: Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        :param pulumi.Input[str] availability_zone: The Availability Zone in which to allocate the Dedicated Host.
        :param pulumi.Input[str] host_recovery: Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        :param pulumi.Input[str] instance_family: Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] instance_type: Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] outpost_arn: The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedHostArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an EC2 Host resource. This allows Dedicated Hosts to be allocated, modified, and released.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new host with instance type of c5.18xlarge with Auto Placement
        # and Host Recovery enabled.
        test = aws.ec2.DedicatedHost("test",
            auto_placement="on",
            availability_zone="us-west-2a",
            host_recovery="on",
            instance_type="c5.18xlarge")
        ```

        ## Import

        Using `pulumi import`, import hosts using the host `id`. For example:

        ```sh
         $ pulumi import aws:ec2/dedicatedHost:DedicatedHost example h-0385a99d0e4b20cbb
        ```

        :param str resource_name: The name of the resource.
        :param DedicatedHostArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedHostArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_id: Optional[pulumi.Input[str]] = None,
                 auto_placement: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 host_recovery: Optional[pulumi.Input[str]] = None,
                 instance_family: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 outpost_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedHostArgs.__new__(DedicatedHostArgs)

            __props__.__dict__["asset_id"] = asset_id
            __props__.__dict__["auto_placement"] = auto_placement
            if availability_zone is None and not opts.urn:
                raise TypeError("Missing required property 'availability_zone'")
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["host_recovery"] = host_recovery
            __props__.__dict__["instance_family"] = instance_family
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["outpost_arn"] = outpost_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["owner_id"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(DedicatedHost, __self__).__init__(
            'aws:ec2/dedicatedHost:DedicatedHost',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            asset_id: Optional[pulumi.Input[str]] = None,
            auto_placement: Optional[pulumi.Input[str]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            host_recovery: Optional[pulumi.Input[str]] = None,
            instance_family: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            outpost_arn: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'DedicatedHost':
        """
        Get an existing DedicatedHost resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Dedicated Host.
        :param pulumi.Input[str] asset_id: The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        :param pulumi.Input[str] auto_placement: Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        :param pulumi.Input[str] availability_zone: The Availability Zone in which to allocate the Dedicated Host.
        :param pulumi.Input[str] host_recovery: Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        :param pulumi.Input[str] instance_family: Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] instance_type: Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        :param pulumi.Input[str] outpost_arn: The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the Dedicated Host.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DedicatedHostState.__new__(_DedicatedHostState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["asset_id"] = asset_id
        __props__.__dict__["auto_placement"] = auto_placement
        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["host_recovery"] = host_recovery
        __props__.__dict__["instance_family"] = instance_family
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["outpost_arn"] = outpost_arn
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return DedicatedHost(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Dedicated Host.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assetId")
    def asset_id(self) -> pulumi.Output[str]:
        """
        The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        """
        return pulumi.get(self, "asset_id")

    @property
    @pulumi.getter(name="autoPlacement")
    def auto_placement(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        """
        return pulumi.get(self, "auto_placement")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The Availability Zone in which to allocate the Dedicated Host.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="hostRecovery")
    def host_recovery(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        """
        return pulumi.get(self, "host_recovery")

    @property
    @pulumi.getter(name="instanceFamily")
    def instance_family(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        """
        return pulumi.get(self, "instance_family")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        """
        return pulumi.get(self, "outpost_arn")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The ID of the AWS account that owns the Dedicated Host.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

