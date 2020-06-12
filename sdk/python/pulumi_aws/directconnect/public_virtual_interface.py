# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PublicVirtualInterface(pulumi.CustomResource):
    address_family: pulumi.Output[str]
    """
    The address family for the BGP peer. `ipv4 ` or `ipv6`.
    """
    amazon_address: pulumi.Output[str]
    """
    The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
    """
    amazon_side_asn: pulumi.Output[str]
    arn: pulumi.Output[str]
    """
    The ARN of the virtual interface.
    """
    aws_device: pulumi.Output[str]
    """
    The Direct Connect endpoint on which the virtual interface terminates.
    """
    bgp_asn: pulumi.Output[float]
    """
    The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
    """
    bgp_auth_key: pulumi.Output[str]
    """
    The authentication key for BGP configuration.
    """
    connection_id: pulumi.Output[str]
    """
    The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
    """
    customer_address: pulumi.Output[str]
    """
    The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
    """
    name: pulumi.Output[str]
    """
    The name for the virtual interface.
    """
    route_filter_prefixes: pulumi.Output[list]
    """
    A list of routes to be advertised to the AWS network in this region.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    vlan: pulumi.Output[float]
    """
    The VLAN ID.
    """
    def __init__(__self__, resource_name, opts=None, address_family=None, amazon_address=None, bgp_asn=None, bgp_auth_key=None, connection_id=None, customer_address=None, name=None, route_filter_prefixes=None, tags=None, vlan=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Direct Connect public virtual interface resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.directconnect.PublicVirtualInterface("foo",
            address_family="ipv4",
            amazon_address="175.45.176.2/30",
            bgp_asn=65352,
            connection_id="dxcon-zzzzzzzz",
            customer_address="175.45.176.1/30",
            route_filter_prefixes=[
                "210.52.109.0/24",
                "175.45.176.0/22",
            ],
            vlan=4094)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The address family for the BGP peer. `ipv4 ` or `ipv6`.
        :param pulumi.Input[str] amazon_address: The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        :param pulumi.Input[float] bgp_asn: The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        :param pulumi.Input[str] bgp_auth_key: The authentication key for BGP configuration.
        :param pulumi.Input[str] connection_id: The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        :param pulumi.Input[str] customer_address: The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        :param pulumi.Input[str] name: The name for the virtual interface.
        :param pulumi.Input[list] route_filter_prefixes: A list of routes to be advertised to the AWS network in this region.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] vlan: The VLAN ID.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if address_family is None:
                raise TypeError("Missing required property 'address_family'")
            __props__['address_family'] = address_family
            __props__['amazon_address'] = amazon_address
            if bgp_asn is None:
                raise TypeError("Missing required property 'bgp_asn'")
            __props__['bgp_asn'] = bgp_asn
            __props__['bgp_auth_key'] = bgp_auth_key
            if connection_id is None:
                raise TypeError("Missing required property 'connection_id'")
            __props__['connection_id'] = connection_id
            __props__['customer_address'] = customer_address
            __props__['name'] = name
            if route_filter_prefixes is None:
                raise TypeError("Missing required property 'route_filter_prefixes'")
            __props__['route_filter_prefixes'] = route_filter_prefixes
            __props__['tags'] = tags
            if vlan is None:
                raise TypeError("Missing required property 'vlan'")
            __props__['vlan'] = vlan
            __props__['amazon_side_asn'] = None
            __props__['arn'] = None
            __props__['aws_device'] = None
        super(PublicVirtualInterface, __self__).__init__(
            'aws:directconnect/publicVirtualInterface:PublicVirtualInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address_family=None, amazon_address=None, amazon_side_asn=None, arn=None, aws_device=None, bgp_asn=None, bgp_auth_key=None, connection_id=None, customer_address=None, name=None, route_filter_prefixes=None, tags=None, vlan=None):
        """
        Get an existing PublicVirtualInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The address family for the BGP peer. `ipv4 ` or `ipv6`.
        :param pulumi.Input[str] amazon_address: The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        :param pulumi.Input[str] arn: The ARN of the virtual interface.
        :param pulumi.Input[str] aws_device: The Direct Connect endpoint on which the virtual interface terminates.
        :param pulumi.Input[float] bgp_asn: The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        :param pulumi.Input[str] bgp_auth_key: The authentication key for BGP configuration.
        :param pulumi.Input[str] connection_id: The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        :param pulumi.Input[str] customer_address: The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        :param pulumi.Input[str] name: The name for the virtual interface.
        :param pulumi.Input[list] route_filter_prefixes: A list of routes to be advertised to the AWS network in this region.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] vlan: The VLAN ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_family"] = address_family
        __props__["amazon_address"] = amazon_address
        __props__["amazon_side_asn"] = amazon_side_asn
        __props__["arn"] = arn
        __props__["aws_device"] = aws_device
        __props__["bgp_asn"] = bgp_asn
        __props__["bgp_auth_key"] = bgp_auth_key
        __props__["connection_id"] = connection_id
        __props__["customer_address"] = customer_address
        __props__["name"] = name
        __props__["route_filter_prefixes"] = route_filter_prefixes
        __props__["tags"] = tags
        __props__["vlan"] = vlan
        return PublicVirtualInterface(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

