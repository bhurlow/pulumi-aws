# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcIpamPoolCidrAllocationArgs', 'VpcIpamPoolCidrAllocation']

@pulumi.input_type
class VpcIpamPoolCidrAllocationArgs:
    def __init__(__self__, *,
                 ipam_pool_id: pulumi.Input[str],
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disallowed_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 netmask_length: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a VpcIpamPoolCidrAllocation resource.
        :param pulumi.Input[str] ipam_pool_id: The ID of the pool to which you want to assign a CIDR.
        :param pulumi.Input[str] cidr: The CIDR you want to assign to the pool.
        :param pulumi.Input[str] description: The description for the allocation.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disallowed_cidrs: Exclude a particular CIDR range from being returned by the pool.
        :param pulumi.Input[int] netmask_length: The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
        """
        pulumi.set(__self__, "ipam_pool_id", ipam_pool_id)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disallowed_cidrs is not None:
            pulumi.set(__self__, "disallowed_cidrs", disallowed_cidrs)
        if netmask_length is not None:
            pulumi.set(__self__, "netmask_length", netmask_length)

    @property
    @pulumi.getter(name="ipamPoolId")
    def ipam_pool_id(self) -> pulumi.Input[str]:
        """
        The ID of the pool to which you want to assign a CIDR.
        """
        return pulumi.get(self, "ipam_pool_id")

    @ipam_pool_id.setter
    def ipam_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ipam_pool_id", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR you want to assign to the pool.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the allocation.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disallowedCidrs")
    def disallowed_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Exclude a particular CIDR range from being returned by the pool.
        """
        return pulumi.get(self, "disallowed_cidrs")

    @disallowed_cidrs.setter
    def disallowed_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disallowed_cidrs", value)

    @property
    @pulumi.getter(name="netmaskLength")
    def netmask_length(self) -> Optional[pulumi.Input[int]]:
        """
        The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
        """
        return pulumi.get(self, "netmask_length")

    @netmask_length.setter
    def netmask_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "netmask_length", value)


@pulumi.input_type
class _VpcIpamPoolCidrAllocationState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disallowed_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ipam_pool_allocation_id: Optional[pulumi.Input[str]] = None,
                 ipam_pool_id: Optional[pulumi.Input[str]] = None,
                 netmask_length: Optional[pulumi.Input[int]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_owner: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcIpamPoolCidrAllocation resources.
        :param pulumi.Input[str] cidr: The CIDR you want to assign to the pool.
        :param pulumi.Input[str] description: The description for the allocation.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disallowed_cidrs: Exclude a particular CIDR range from being returned by the pool.
        :param pulumi.Input[str] ipam_pool_id: The ID of the pool to which you want to assign a CIDR.
        :param pulumi.Input[int] netmask_length: The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_owner: The owner of the resource.
        :param pulumi.Input[str] resource_type: The type of the resource.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disallowed_cidrs is not None:
            pulumi.set(__self__, "disallowed_cidrs", disallowed_cidrs)
        if ipam_pool_allocation_id is not None:
            pulumi.set(__self__, "ipam_pool_allocation_id", ipam_pool_allocation_id)
        if ipam_pool_id is not None:
            pulumi.set(__self__, "ipam_pool_id", ipam_pool_id)
        if netmask_length is not None:
            pulumi.set(__self__, "netmask_length", netmask_length)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_owner is not None:
            pulumi.set(__self__, "resource_owner", resource_owner)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR you want to assign to the pool.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the allocation.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disallowedCidrs")
    def disallowed_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Exclude a particular CIDR range from being returned by the pool.
        """
        return pulumi.get(self, "disallowed_cidrs")

    @disallowed_cidrs.setter
    def disallowed_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disallowed_cidrs", value)

    @property
    @pulumi.getter(name="ipamPoolAllocationId")
    def ipam_pool_allocation_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipam_pool_allocation_id")

    @ipam_pool_allocation_id.setter
    def ipam_pool_allocation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipam_pool_allocation_id", value)

    @property
    @pulumi.getter(name="ipamPoolId")
    def ipam_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the pool to which you want to assign a CIDR.
        """
        return pulumi.get(self, "ipam_pool_id")

    @ipam_pool_id.setter
    def ipam_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipam_pool_id", value)

    @property
    @pulumi.getter(name="netmaskLength")
    def netmask_length(self) -> Optional[pulumi.Input[int]]:
        """
        The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
        """
        return pulumi.get(self, "netmask_length")

    @netmask_length.setter
    def netmask_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "netmask_length", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceOwner")
    def resource_owner(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the resource.
        """
        return pulumi.get(self, "resource_owner")

    @resource_owner.setter
    def resource_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_owner", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


class VpcIpamPoolCidrAllocation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disallowed_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ipam_pool_id: Optional[pulumi.Input[str]] = None,
                 netmask_length: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Allocates (reserves) a CIDR from an IPAM address pool, preventing usage by IPAM. Only works for private IPv4.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        current = aws.get_region()
        example_vpc_ipam = aws.ec2.VpcIpam("exampleVpcIpam", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
            region_name=current.name,
        )])
        example_vpc_ipam_pool = aws.ec2.VpcIpamPool("exampleVpcIpamPool",
            address_family="ipv4",
            ipam_scope_id=example_vpc_ipam.private_default_scope_id,
            locale=current.name)
        example_vpc_ipam_pool_cidr = aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr",
            ipam_pool_id=example_vpc_ipam_pool.id,
            cidr="172.2.0.0/16")
        example_vpc_ipam_pool_cidr_allocation = aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation",
            ipam_pool_id=example_vpc_ipam_pool.id,
            cidr="172.2.0.0/24",
            opts=pulumi.ResourceOptions(depends_on=[example_vpc_ipam_pool_cidr]))
        ```

        With the `disallowed_cidrs` attribute:

        ```python
        import pulumi
        import pulumi_aws as aws

        current = aws.get_region()
        example_vpc_ipam = aws.ec2.VpcIpam("exampleVpcIpam", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
            region_name=current.name,
        )])
        example_vpc_ipam_pool = aws.ec2.VpcIpamPool("exampleVpcIpamPool",
            address_family="ipv4",
            ipam_scope_id=example_vpc_ipam.private_default_scope_id,
            locale=current.name)
        example_vpc_ipam_pool_cidr = aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr",
            ipam_pool_id=example_vpc_ipam_pool.id,
            cidr="172.2.0.0/16")
        example_vpc_ipam_pool_cidr_allocation = aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation",
            ipam_pool_id=example_vpc_ipam_pool.id,
            netmask_length=28,
            disallowed_cidrs=["172.2.0.0/28"],
            opts=pulumi.ResourceOptions(depends_on=[example_vpc_ipam_pool_cidr]))
        ```

        ## Import

        terraform import {

         to = aws_vpc_ipam_pool_cidr_allocation.example

         id = "ipam-pool-alloc-0dc6d196509c049ba8b549ff99f639736_ipam-pool-07cfb559e0921fcbe" } Using `pulumi import`, import IPAM allocations using the allocation `id` and `pool id`, separated by `_`. For exampleconsole % pulumi import aws_vpc_ipam_pool_cidr_allocation.example ipam-pool-alloc-0dc6d196509c049ba8b549ff99f639736_ipam-pool-07cfb559e0921fcbe

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The CIDR you want to assign to the pool.
        :param pulumi.Input[str] description: The description for the allocation.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disallowed_cidrs: Exclude a particular CIDR range from being returned by the pool.
        :param pulumi.Input[str] ipam_pool_id: The ID of the pool to which you want to assign a CIDR.
        :param pulumi.Input[int] netmask_length: The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcIpamPoolCidrAllocationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allocates (reserves) a CIDR from an IPAM address pool, preventing usage by IPAM. Only works for private IPv4.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        current = aws.get_region()
        example_vpc_ipam = aws.ec2.VpcIpam("exampleVpcIpam", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
            region_name=current.name,
        )])
        example_vpc_ipam_pool = aws.ec2.VpcIpamPool("exampleVpcIpamPool",
            address_family="ipv4",
            ipam_scope_id=example_vpc_ipam.private_default_scope_id,
            locale=current.name)
        example_vpc_ipam_pool_cidr = aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr",
            ipam_pool_id=example_vpc_ipam_pool.id,
            cidr="172.2.0.0/16")
        example_vpc_ipam_pool_cidr_allocation = aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation",
            ipam_pool_id=example_vpc_ipam_pool.id,
            cidr="172.2.0.0/24",
            opts=pulumi.ResourceOptions(depends_on=[example_vpc_ipam_pool_cidr]))
        ```

        With the `disallowed_cidrs` attribute:

        ```python
        import pulumi
        import pulumi_aws as aws

        current = aws.get_region()
        example_vpc_ipam = aws.ec2.VpcIpam("exampleVpcIpam", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
            region_name=current.name,
        )])
        example_vpc_ipam_pool = aws.ec2.VpcIpamPool("exampleVpcIpamPool",
            address_family="ipv4",
            ipam_scope_id=example_vpc_ipam.private_default_scope_id,
            locale=current.name)
        example_vpc_ipam_pool_cidr = aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr",
            ipam_pool_id=example_vpc_ipam_pool.id,
            cidr="172.2.0.0/16")
        example_vpc_ipam_pool_cidr_allocation = aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation",
            ipam_pool_id=example_vpc_ipam_pool.id,
            netmask_length=28,
            disallowed_cidrs=["172.2.0.0/28"],
            opts=pulumi.ResourceOptions(depends_on=[example_vpc_ipam_pool_cidr]))
        ```

        ## Import

        terraform import {

         to = aws_vpc_ipam_pool_cidr_allocation.example

         id = "ipam-pool-alloc-0dc6d196509c049ba8b549ff99f639736_ipam-pool-07cfb559e0921fcbe" } Using `pulumi import`, import IPAM allocations using the allocation `id` and `pool id`, separated by `_`. For exampleconsole % pulumi import aws_vpc_ipam_pool_cidr_allocation.example ipam-pool-alloc-0dc6d196509c049ba8b549ff99f639736_ipam-pool-07cfb559e0921fcbe

        :param str resource_name: The name of the resource.
        :param VpcIpamPoolCidrAllocationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcIpamPoolCidrAllocationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disallowed_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ipam_pool_id: Optional[pulumi.Input[str]] = None,
                 netmask_length: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcIpamPoolCidrAllocationArgs.__new__(VpcIpamPoolCidrAllocationArgs)

            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["description"] = description
            __props__.__dict__["disallowed_cidrs"] = disallowed_cidrs
            if ipam_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'ipam_pool_id'")
            __props__.__dict__["ipam_pool_id"] = ipam_pool_id
            __props__.__dict__["netmask_length"] = netmask_length
            __props__.__dict__["ipam_pool_allocation_id"] = None
            __props__.__dict__["resource_id"] = None
            __props__.__dict__["resource_owner"] = None
            __props__.__dict__["resource_type"] = None
        super(VpcIpamPoolCidrAllocation, __self__).__init__(
            'aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disallowed_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ipam_pool_allocation_id: Optional[pulumi.Input[str]] = None,
            ipam_pool_id: Optional[pulumi.Input[str]] = None,
            netmask_length: Optional[pulumi.Input[int]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_owner: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None) -> 'VpcIpamPoolCidrAllocation':
        """
        Get an existing VpcIpamPoolCidrAllocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The CIDR you want to assign to the pool.
        :param pulumi.Input[str] description: The description for the allocation.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disallowed_cidrs: Exclude a particular CIDR range from being returned by the pool.
        :param pulumi.Input[str] ipam_pool_id: The ID of the pool to which you want to assign a CIDR.
        :param pulumi.Input[int] netmask_length: The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_owner: The owner of the resource.
        :param pulumi.Input[str] resource_type: The type of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcIpamPoolCidrAllocationState.__new__(_VpcIpamPoolCidrAllocationState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["description"] = description
        __props__.__dict__["disallowed_cidrs"] = disallowed_cidrs
        __props__.__dict__["ipam_pool_allocation_id"] = ipam_pool_allocation_id
        __props__.__dict__["ipam_pool_id"] = ipam_pool_id
        __props__.__dict__["netmask_length"] = netmask_length
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_owner"] = resource_owner
        __props__.__dict__["resource_type"] = resource_type
        return VpcIpamPoolCidrAllocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        The CIDR you want to assign to the pool.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the allocation.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disallowedCidrs")
    def disallowed_cidrs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Exclude a particular CIDR range from being returned by the pool.
        """
        return pulumi.get(self, "disallowed_cidrs")

    @property
    @pulumi.getter(name="ipamPoolAllocationId")
    def ipam_pool_allocation_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ipam_pool_allocation_id")

    @property
    @pulumi.getter(name="ipamPoolId")
    def ipam_pool_id(self) -> pulumi.Output[str]:
        """
        The ID of the pool to which you want to assign a CIDR.
        """
        return pulumi.get(self, "ipam_pool_id")

    @property
    @pulumi.getter(name="netmaskLength")
    def netmask_length(self) -> pulumi.Output[Optional[int]]:
        """
        The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
        """
        return pulumi.get(self, "netmask_length")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceOwner")
    def resource_owner(self) -> pulumi.Output[str]:
        """
        The owner of the resource.
        """
        return pulumi.get(self, "resource_owner")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "resource_type")

