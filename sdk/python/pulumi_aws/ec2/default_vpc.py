# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DefaultVpcArgs', 'DefaultVpc']

@pulumi.input_type
class DefaultVpcArgs:
    def __init__(__self__, *,
                 enable_classiclink: Optional[pulumi.Input[bool]] = None,
                 enable_classiclink_dns_support: Optional[pulumi.Input[bool]] = None,
                 enable_dns_hostnames: Optional[pulumi.Input[bool]] = None,
                 enable_dns_support: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DefaultVpc resource.
        :param pulumi.Input[bool] enable_classiclink: A boolean flag to enable/disable ClassicLink
               for the VPC. Only valid in regions and accounts that support EC2 Classic.
               See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        :param pulumi.Input[bool] enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        :param pulumi.Input[bool] enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        if enable_classiclink is not None:
            pulumi.set(__self__, "enable_classiclink", enable_classiclink)
        if enable_classiclink_dns_support is not None:
            pulumi.set(__self__, "enable_classiclink_dns_support", enable_classiclink_dns_support)
        if enable_dns_hostnames is not None:
            pulumi.set(__self__, "enable_dns_hostnames", enable_dns_hostnames)
        if enable_dns_support is not None:
            pulumi.set(__self__, "enable_dns_support", enable_dns_support)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="enableClassiclink")
    def enable_classiclink(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable ClassicLink
        for the VPC. Only valid in regions and accounts that support EC2 Classic.
        See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        """
        return pulumi.get(self, "enable_classiclink")

    @enable_classiclink.setter
    def enable_classiclink(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_classiclink", value)

    @property
    @pulumi.getter(name="enableClassiclinkDnsSupport")
    def enable_classiclink_dns_support(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_classiclink_dns_support")

    @enable_classiclink_dns_support.setter
    def enable_classiclink_dns_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_classiclink_dns_support", value)

    @property
    @pulumi.getter(name="enableDnsHostnames")
    def enable_dns_hostnames(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        """
        return pulumi.get(self, "enable_dns_hostnames")

    @enable_dns_hostnames.setter
    def enable_dns_hostnames(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_dns_hostnames", value)

    @property
    @pulumi.getter(name="enableDnsSupport")
    def enable_dns_support(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        """
        return pulumi.get(self, "enable_dns_support")

    @enable_dns_support.setter
    def enable_dns_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_dns_support", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DefaultVpcState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 assign_generated_ipv6_cidr_block: Optional[pulumi.Input[bool]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 default_network_acl_id: Optional[pulumi.Input[str]] = None,
                 default_route_table_id: Optional[pulumi.Input[str]] = None,
                 default_security_group_id: Optional[pulumi.Input[str]] = None,
                 dhcp_options_id: Optional[pulumi.Input[str]] = None,
                 enable_classiclink: Optional[pulumi.Input[bool]] = None,
                 enable_classiclink_dns_support: Optional[pulumi.Input[bool]] = None,
                 enable_dns_hostnames: Optional[pulumi.Input[bool]] = None,
                 enable_dns_support: Optional[pulumi.Input[bool]] = None,
                 instance_tenancy: Optional[pulumi.Input[str]] = None,
                 ipv6_association_id: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 main_route_table_id: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DefaultVpc resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of VPC
        :param pulumi.Input[bool] assign_generated_ipv6_cidr_block: Whether or not an Amazon-provided IPv6 CIDR
               block with a /56 prefix length for the VPC was assigned
        :param pulumi.Input[str] cidr_block: The CIDR block of the VPC
        :param pulumi.Input[str] default_network_acl_id: The ID of the network ACL created by default on VPC creation
        :param pulumi.Input[str] default_route_table_id: The ID of the route table created by default on VPC creation
        :param pulumi.Input[str] default_security_group_id: The ID of the security group created by default on VPC creation
        :param pulumi.Input[bool] enable_classiclink: A boolean flag to enable/disable ClassicLink
               for the VPC. Only valid in regions and accounts that support EC2 Classic.
               See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        :param pulumi.Input[bool] enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        :param pulumi.Input[bool] enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param pulumi.Input[str] instance_tenancy: Tenancy of instances spin up within VPC.
        :param pulumi.Input[str] ipv6_association_id: The association ID for the IPv6 CIDR block of the VPC
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block of the VPC
        :param pulumi.Input[str] main_route_table_id: The ID of the main route table associated with
               this VPC. Note that you can change a VPC's main route table by using an
               `ec2.MainRouteTableAssociation`
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the VPC.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if assign_generated_ipv6_cidr_block is not None:
            pulumi.set(__self__, "assign_generated_ipv6_cidr_block", assign_generated_ipv6_cidr_block)
        if cidr_block is not None:
            pulumi.set(__self__, "cidr_block", cidr_block)
        if default_network_acl_id is not None:
            pulumi.set(__self__, "default_network_acl_id", default_network_acl_id)
        if default_route_table_id is not None:
            pulumi.set(__self__, "default_route_table_id", default_route_table_id)
        if default_security_group_id is not None:
            pulumi.set(__self__, "default_security_group_id", default_security_group_id)
        if dhcp_options_id is not None:
            pulumi.set(__self__, "dhcp_options_id", dhcp_options_id)
        if enable_classiclink is not None:
            pulumi.set(__self__, "enable_classiclink", enable_classiclink)
        if enable_classiclink_dns_support is not None:
            pulumi.set(__self__, "enable_classiclink_dns_support", enable_classiclink_dns_support)
        if enable_dns_hostnames is not None:
            pulumi.set(__self__, "enable_dns_hostnames", enable_dns_hostnames)
        if enable_dns_support is not None:
            pulumi.set(__self__, "enable_dns_support", enable_dns_support)
        if instance_tenancy is not None:
            pulumi.set(__self__, "instance_tenancy", instance_tenancy)
        if ipv6_association_id is not None:
            pulumi.set(__self__, "ipv6_association_id", ipv6_association_id)
        if ipv6_cidr_block is not None:
            pulumi.set(__self__, "ipv6_cidr_block", ipv6_cidr_block)
        if main_route_table_id is not None:
            pulumi.set(__self__, "main_route_table_id", main_route_table_id)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of VPC
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="assignGeneratedIpv6CidrBlock")
    def assign_generated_ipv6_cidr_block(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not an Amazon-provided IPv6 CIDR
        block with a /56 prefix length for the VPC was assigned
        """
        return pulumi.get(self, "assign_generated_ipv6_cidr_block")

    @assign_generated_ipv6_cidr_block.setter
    def assign_generated_ipv6_cidr_block(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "assign_generated_ipv6_cidr_block", value)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block of the VPC
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter(name="defaultNetworkAclId")
    def default_network_acl_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the network ACL created by default on VPC creation
        """
        return pulumi.get(self, "default_network_acl_id")

    @default_network_acl_id.setter
    def default_network_acl_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_network_acl_id", value)

    @property
    @pulumi.getter(name="defaultRouteTableId")
    def default_route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the route table created by default on VPC creation
        """
        return pulumi.get(self, "default_route_table_id")

    @default_route_table_id.setter
    def default_route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_route_table_id", value)

    @property
    @pulumi.getter(name="defaultSecurityGroupId")
    def default_security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the security group created by default on VPC creation
        """
        return pulumi.get(self, "default_security_group_id")

    @default_security_group_id.setter
    def default_security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_security_group_id", value)

    @property
    @pulumi.getter(name="dhcpOptionsId")
    def dhcp_options_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dhcp_options_id")

    @dhcp_options_id.setter
    def dhcp_options_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_options_id", value)

    @property
    @pulumi.getter(name="enableClassiclink")
    def enable_classiclink(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable ClassicLink
        for the VPC. Only valid in regions and accounts that support EC2 Classic.
        See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        """
        return pulumi.get(self, "enable_classiclink")

    @enable_classiclink.setter
    def enable_classiclink(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_classiclink", value)

    @property
    @pulumi.getter(name="enableClassiclinkDnsSupport")
    def enable_classiclink_dns_support(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_classiclink_dns_support")

    @enable_classiclink_dns_support.setter
    def enable_classiclink_dns_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_classiclink_dns_support", value)

    @property
    @pulumi.getter(name="enableDnsHostnames")
    def enable_dns_hostnames(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        """
        return pulumi.get(self, "enable_dns_hostnames")

    @enable_dns_hostnames.setter
    def enable_dns_hostnames(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_dns_hostnames", value)

    @property
    @pulumi.getter(name="enableDnsSupport")
    def enable_dns_support(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        """
        return pulumi.get(self, "enable_dns_support")

    @enable_dns_support.setter
    def enable_dns_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_dns_support", value)

    @property
    @pulumi.getter(name="instanceTenancy")
    def instance_tenancy(self) -> Optional[pulumi.Input[str]]:
        """
        Tenancy of instances spin up within VPC.
        """
        return pulumi.get(self, "instance_tenancy")

    @instance_tenancy.setter
    def instance_tenancy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_tenancy", value)

    @property
    @pulumi.getter(name="ipv6AssociationId")
    def ipv6_association_id(self) -> Optional[pulumi.Input[str]]:
        """
        The association ID for the IPv6 CIDR block of the VPC
        """
        return pulumi.get(self, "ipv6_association_id")

    @ipv6_association_id.setter
    def ipv6_association_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_association_id", value)

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 CIDR block of the VPC
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @ipv6_cidr_block.setter
    def ipv6_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_cidr_block", value)

    @property
    @pulumi.getter(name="mainRouteTableId")
    def main_route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the main route table associated with
        this VPC. Note that you can change a VPC's main route table by using an
        `ec2.MainRouteTableAssociation`
        """
        return pulumi.get(self, "main_route_table_id")

    @main_route_table_id.setter
    def main_route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "main_route_table_id", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the AWS account that owns the VPC.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class DefaultVpc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_classiclink: Optional[pulumi.Input[bool]] = None,
                 enable_classiclink_dns_support: Optional[pulumi.Input[bool]] = None,
                 enable_dns_hostnames: Optional[pulumi.Input[bool]] = None,
                 enable_dns_support: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
        in the current region.

        For AWS accounts created after 2013-12-04, each region comes with a Default VPC.
        **This is an advanced resource**, and has special caveats to be aware of when
        using it. Please read this document in its entirety before using this resource.

        The `ec2.DefaultVpc` behaves differently from normal resources, in that
        this provider does not _create_ this resource, but instead "adopts" it
        into management.

        ## Example Usage

        Basic usage with tags:

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.ec2.DefaultVpc("default", tags={
            "Name": "Default VPC",
        })
        ```

        ## Import

        Default VPCs can be imported using the `vpc id`, e.g.

        ```sh
         $ pulumi import aws:ec2/defaultVpc:DefaultVpc default vpc-a01106c2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_classiclink: A boolean flag to enable/disable ClassicLink
               for the VPC. Only valid in regions and accounts that support EC2 Classic.
               See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        :param pulumi.Input[bool] enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        :param pulumi.Input[bool] enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DefaultVpcArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
        in the current region.

        For AWS accounts created after 2013-12-04, each region comes with a Default VPC.
        **This is an advanced resource**, and has special caveats to be aware of when
        using it. Please read this document in its entirety before using this resource.

        The `ec2.DefaultVpc` behaves differently from normal resources, in that
        this provider does not _create_ this resource, but instead "adopts" it
        into management.

        ## Example Usage

        Basic usage with tags:

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.ec2.DefaultVpc("default", tags={
            "Name": "Default VPC",
        })
        ```

        ## Import

        Default VPCs can be imported using the `vpc id`, e.g.

        ```sh
         $ pulumi import aws:ec2/defaultVpc:DefaultVpc default vpc-a01106c2
        ```

        :param str resource_name: The name of the resource.
        :param DefaultVpcArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultVpcArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_classiclink: Optional[pulumi.Input[bool]] = None,
                 enable_classiclink_dns_support: Optional[pulumi.Input[bool]] = None,
                 enable_dns_hostnames: Optional[pulumi.Input[bool]] = None,
                 enable_dns_support: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultVpcArgs.__new__(DefaultVpcArgs)

            __props__.__dict__["enable_classiclink"] = enable_classiclink
            __props__.__dict__["enable_classiclink_dns_support"] = enable_classiclink_dns_support
            __props__.__dict__["enable_dns_hostnames"] = enable_dns_hostnames
            __props__.__dict__["enable_dns_support"] = enable_dns_support
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["assign_generated_ipv6_cidr_block"] = None
            __props__.__dict__["cidr_block"] = None
            __props__.__dict__["default_network_acl_id"] = None
            __props__.__dict__["default_route_table_id"] = None
            __props__.__dict__["default_security_group_id"] = None
            __props__.__dict__["dhcp_options_id"] = None
            __props__.__dict__["instance_tenancy"] = None
            __props__.__dict__["ipv6_association_id"] = None
            __props__.__dict__["ipv6_cidr_block"] = None
            __props__.__dict__["main_route_table_id"] = None
            __props__.__dict__["owner_id"] = None
            __props__.__dict__["tags_all"] = None
        super(DefaultVpc, __self__).__init__(
            'aws:ec2/defaultVpc:DefaultVpc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            assign_generated_ipv6_cidr_block: Optional[pulumi.Input[bool]] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            default_network_acl_id: Optional[pulumi.Input[str]] = None,
            default_route_table_id: Optional[pulumi.Input[str]] = None,
            default_security_group_id: Optional[pulumi.Input[str]] = None,
            dhcp_options_id: Optional[pulumi.Input[str]] = None,
            enable_classiclink: Optional[pulumi.Input[bool]] = None,
            enable_classiclink_dns_support: Optional[pulumi.Input[bool]] = None,
            enable_dns_hostnames: Optional[pulumi.Input[bool]] = None,
            enable_dns_support: Optional[pulumi.Input[bool]] = None,
            instance_tenancy: Optional[pulumi.Input[str]] = None,
            ipv6_association_id: Optional[pulumi.Input[str]] = None,
            ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
            main_route_table_id: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'DefaultVpc':
        """
        Get an existing DefaultVpc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of VPC
        :param pulumi.Input[bool] assign_generated_ipv6_cidr_block: Whether or not an Amazon-provided IPv6 CIDR
               block with a /56 prefix length for the VPC was assigned
        :param pulumi.Input[str] cidr_block: The CIDR block of the VPC
        :param pulumi.Input[str] default_network_acl_id: The ID of the network ACL created by default on VPC creation
        :param pulumi.Input[str] default_route_table_id: The ID of the route table created by default on VPC creation
        :param pulumi.Input[str] default_security_group_id: The ID of the security group created by default on VPC creation
        :param pulumi.Input[bool] enable_classiclink: A boolean flag to enable/disable ClassicLink
               for the VPC. Only valid in regions and accounts that support EC2 Classic.
               See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        :param pulumi.Input[bool] enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        :param pulumi.Input[bool] enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param pulumi.Input[str] instance_tenancy: Tenancy of instances spin up within VPC.
        :param pulumi.Input[str] ipv6_association_id: The association ID for the IPv6 CIDR block of the VPC
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block of the VPC
        :param pulumi.Input[str] main_route_table_id: The ID of the main route table associated with
               this VPC. Note that you can change a VPC's main route table by using an
               `ec2.MainRouteTableAssociation`
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the VPC.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultVpcState.__new__(_DefaultVpcState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["assign_generated_ipv6_cidr_block"] = assign_generated_ipv6_cidr_block
        __props__.__dict__["cidr_block"] = cidr_block
        __props__.__dict__["default_network_acl_id"] = default_network_acl_id
        __props__.__dict__["default_route_table_id"] = default_route_table_id
        __props__.__dict__["default_security_group_id"] = default_security_group_id
        __props__.__dict__["dhcp_options_id"] = dhcp_options_id
        __props__.__dict__["enable_classiclink"] = enable_classiclink
        __props__.__dict__["enable_classiclink_dns_support"] = enable_classiclink_dns_support
        __props__.__dict__["enable_dns_hostnames"] = enable_dns_hostnames
        __props__.__dict__["enable_dns_support"] = enable_dns_support
        __props__.__dict__["instance_tenancy"] = instance_tenancy
        __props__.__dict__["ipv6_association_id"] = ipv6_association_id
        __props__.__dict__["ipv6_cidr_block"] = ipv6_cidr_block
        __props__.__dict__["main_route_table_id"] = main_route_table_id
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return DefaultVpc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of VPC
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assignGeneratedIpv6CidrBlock")
    def assign_generated_ipv6_cidr_block(self) -> pulumi.Output[bool]:
        """
        Whether or not an Amazon-provided IPv6 CIDR
        block with a /56 prefix length for the VPC was assigned
        """
        return pulumi.get(self, "assign_generated_ipv6_cidr_block")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[str]:
        """
        The CIDR block of the VPC
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter(name="defaultNetworkAclId")
    def default_network_acl_id(self) -> pulumi.Output[str]:
        """
        The ID of the network ACL created by default on VPC creation
        """
        return pulumi.get(self, "default_network_acl_id")

    @property
    @pulumi.getter(name="defaultRouteTableId")
    def default_route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the route table created by default on VPC creation
        """
        return pulumi.get(self, "default_route_table_id")

    @property
    @pulumi.getter(name="defaultSecurityGroupId")
    def default_security_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the security group created by default on VPC creation
        """
        return pulumi.get(self, "default_security_group_id")

    @property
    @pulumi.getter(name="dhcpOptionsId")
    def dhcp_options_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dhcp_options_id")

    @property
    @pulumi.getter(name="enableClassiclink")
    def enable_classiclink(self) -> pulumi.Output[bool]:
        """
        A boolean flag to enable/disable ClassicLink
        for the VPC. Only valid in regions and accounts that support EC2 Classic.
        See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        """
        return pulumi.get(self, "enable_classiclink")

    @property
    @pulumi.getter(name="enableClassiclinkDnsSupport")
    def enable_classiclink_dns_support(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "enable_classiclink_dns_support")

    @property
    @pulumi.getter(name="enableDnsHostnames")
    def enable_dns_hostnames(self) -> pulumi.Output[bool]:
        """
        A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        """
        return pulumi.get(self, "enable_dns_hostnames")

    @property
    @pulumi.getter(name="enableDnsSupport")
    def enable_dns_support(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        """
        return pulumi.get(self, "enable_dns_support")

    @property
    @pulumi.getter(name="instanceTenancy")
    def instance_tenancy(self) -> pulumi.Output[str]:
        """
        Tenancy of instances spin up within VPC.
        """
        return pulumi.get(self, "instance_tenancy")

    @property
    @pulumi.getter(name="ipv6AssociationId")
    def ipv6_association_id(self) -> pulumi.Output[str]:
        """
        The association ID for the IPv6 CIDR block of the VPC
        """
        return pulumi.get(self, "ipv6_association_id")

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> pulumi.Output[str]:
        """
        The IPv6 CIDR block of the VPC
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @property
    @pulumi.getter(name="mainRouteTableId")
    def main_route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the main route table associated with
        this VPC. Note that you can change a VPC's main route table by using an
        `ec2.MainRouteTableAssociation`
        """
        return pulumi.get(self, "main_route_table_id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The ID of the AWS account that owns the VPC.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags_all")

