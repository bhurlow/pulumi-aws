# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConnectPeerArgs', 'ConnectPeer']

@pulumi.input_type
class ConnectPeerArgs:
    def __init__(__self__, *,
                 inside_cidr_blocks: pulumi.Input[Sequence[pulumi.Input[str]]],
                 peer_address: pulumi.Input[str],
                 transit_gateway_attachment_id: pulumi.Input[str],
                 bgp_asn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_address: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConnectPeer resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inside_cidr_blocks: The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        :param pulumi.Input[str] peer_address: The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        :param pulumi.Input[str] transit_gateway_attachment_id: The Transit Gateway Connect
        :param pulumi.Input[str] bgp_asn: The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] transit_gateway_address: The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        """
        pulumi.set(__self__, "inside_cidr_blocks", inside_cidr_blocks)
        pulumi.set(__self__, "peer_address", peer_address)
        pulumi.set(__self__, "transit_gateway_attachment_id", transit_gateway_attachment_id)
        if bgp_asn is not None:
            pulumi.set(__self__, "bgp_asn", bgp_asn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if transit_gateway_address is not None:
            pulumi.set(__self__, "transit_gateway_address", transit_gateway_address)

    @property
    @pulumi.getter(name="insideCidrBlocks")
    def inside_cidr_blocks(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        """
        return pulumi.get(self, "inside_cidr_blocks")

    @inside_cidr_blocks.setter
    def inside_cidr_blocks(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "inside_cidr_blocks", value)

    @property
    @pulumi.getter(name="peerAddress")
    def peer_address(self) -> pulumi.Input[str]:
        """
        The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        """
        return pulumi.get(self, "peer_address")

    @peer_address.setter
    def peer_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_address", value)

    @property
    @pulumi.getter(name="transitGatewayAttachmentId")
    def transit_gateway_attachment_id(self) -> pulumi.Input[str]:
        """
        The Transit Gateway Connect
        """
        return pulumi.get(self, "transit_gateway_attachment_id")

    @transit_gateway_attachment_id.setter
    def transit_gateway_attachment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_gateway_attachment_id", value)

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> Optional[pulumi.Input[str]]:
        """
        The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        """
        return pulumi.get(self, "bgp_asn")

    @bgp_asn.setter
    def bgp_asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_asn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="transitGatewayAddress")
    def transit_gateway_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        """
        return pulumi.get(self, "transit_gateway_address")

    @transit_gateway_address.setter
    def transit_gateway_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_gateway_address", value)


@pulumi.input_type
class _ConnectPeerState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 bgp_asn: Optional[pulumi.Input[str]] = None,
                 inside_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 peer_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_address: Optional[pulumi.Input[str]] = None,
                 transit_gateway_attachment_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConnectPeer resources.
        :param pulumi.Input[str] arn: EC2 Transit Gateway Connect Peer ARN
        :param pulumi.Input[str] bgp_asn: The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inside_cidr_blocks: The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        :param pulumi.Input[str] peer_address: The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] transit_gateway_address: The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        :param pulumi.Input[str] transit_gateway_attachment_id: The Transit Gateway Connect
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if bgp_asn is not None:
            pulumi.set(__self__, "bgp_asn", bgp_asn)
        if inside_cidr_blocks is not None:
            pulumi.set(__self__, "inside_cidr_blocks", inside_cidr_blocks)
        if peer_address is not None:
            pulumi.set(__self__, "peer_address", peer_address)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if transit_gateway_address is not None:
            pulumi.set(__self__, "transit_gateway_address", transit_gateway_address)
        if transit_gateway_attachment_id is not None:
            pulumi.set(__self__, "transit_gateway_attachment_id", transit_gateway_attachment_id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        EC2 Transit Gateway Connect Peer ARN
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> Optional[pulumi.Input[str]]:
        """
        The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        """
        return pulumi.get(self, "bgp_asn")

    @bgp_asn.setter
    def bgp_asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_asn", value)

    @property
    @pulumi.getter(name="insideCidrBlocks")
    def inside_cidr_blocks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        """
        return pulumi.get(self, "inside_cidr_blocks")

    @inside_cidr_blocks.setter
    def inside_cidr_blocks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "inside_cidr_blocks", value)

    @property
    @pulumi.getter(name="peerAddress")
    def peer_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        """
        return pulumi.get(self, "peer_address")

    @peer_address.setter
    def peer_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_address", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    @property
    @pulumi.getter(name="transitGatewayAddress")
    def transit_gateway_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        """
        return pulumi.get(self, "transit_gateway_address")

    @transit_gateway_address.setter
    def transit_gateway_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_gateway_address", value)

    @property
    @pulumi.getter(name="transitGatewayAttachmentId")
    def transit_gateway_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Transit Gateway Connect
        """
        return pulumi.get(self, "transit_gateway_attachment_id")

    @transit_gateway_attachment_id.setter
    def transit_gateway_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_gateway_attachment_id", value)


class ConnectPeer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bgp_asn: Optional[pulumi.Input[str]] = None,
                 inside_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 peer_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_address: Optional[pulumi.Input[str]] = None,
                 transit_gateway_attachment_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an EC2 Transit Gateway Connect Peer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_connect = aws.ec2transitgateway.Connect("exampleConnect",
            transport_attachment_id=aws_ec2_transit_gateway_vpc_attachment["example"]["id"],
            transit_gateway_id=aws_ec2_transit_gateway["example"]["id"])
        example_connect_peer = aws.ec2transitgateway.ConnectPeer("exampleConnectPeer",
            peer_address="10.1.2.3",
            inside_cidr_blocks=["169.254.100.0/29"],
            transit_gateway_attachment_id=example_connect.id)
        ```

        ## Import

        `aws_ec2_transit_gateway_connect_peer` can be imported by using the EC2 Transit Gateway Connect Peer identifier, e.g.,

        ```sh
         $ pulumi import aws:ec2transitgateway/connectPeer:ConnectPeer example tgw-connect-peer-12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bgp_asn: The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inside_cidr_blocks: The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        :param pulumi.Input[str] peer_address: The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] transit_gateway_address: The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        :param pulumi.Input[str] transit_gateway_attachment_id: The Transit Gateway Connect
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectPeerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an EC2 Transit Gateway Connect Peer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_connect = aws.ec2transitgateway.Connect("exampleConnect",
            transport_attachment_id=aws_ec2_transit_gateway_vpc_attachment["example"]["id"],
            transit_gateway_id=aws_ec2_transit_gateway["example"]["id"])
        example_connect_peer = aws.ec2transitgateway.ConnectPeer("exampleConnectPeer",
            peer_address="10.1.2.3",
            inside_cidr_blocks=["169.254.100.0/29"],
            transit_gateway_attachment_id=example_connect.id)
        ```

        ## Import

        `aws_ec2_transit_gateway_connect_peer` can be imported by using the EC2 Transit Gateway Connect Peer identifier, e.g.,

        ```sh
         $ pulumi import aws:ec2transitgateway/connectPeer:ConnectPeer example tgw-connect-peer-12345678
        ```

        :param str resource_name: The name of the resource.
        :param ConnectPeerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectPeerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bgp_asn: Optional[pulumi.Input[str]] = None,
                 inside_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 peer_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_address: Optional[pulumi.Input[str]] = None,
                 transit_gateway_attachment_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConnectPeerArgs.__new__(ConnectPeerArgs)

            __props__.__dict__["bgp_asn"] = bgp_asn
            if inside_cidr_blocks is None and not opts.urn:
                raise TypeError("Missing required property 'inside_cidr_blocks'")
            __props__.__dict__["inside_cidr_blocks"] = inside_cidr_blocks
            if peer_address is None and not opts.urn:
                raise TypeError("Missing required property 'peer_address'")
            __props__.__dict__["peer_address"] = peer_address
            __props__.__dict__["tags"] = tags
            __props__.__dict__["transit_gateway_address"] = transit_gateway_address
            if transit_gateway_attachment_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_gateway_attachment_id'")
            __props__.__dict__["transit_gateway_attachment_id"] = transit_gateway_attachment_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(ConnectPeer, __self__).__init__(
            'aws:ec2transitgateway/connectPeer:ConnectPeer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            bgp_asn: Optional[pulumi.Input[str]] = None,
            inside_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            peer_address: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            transit_gateway_address: Optional[pulumi.Input[str]] = None,
            transit_gateway_attachment_id: Optional[pulumi.Input[str]] = None) -> 'ConnectPeer':
        """
        Get an existing ConnectPeer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: EC2 Transit Gateway Connect Peer ARN
        :param pulumi.Input[str] bgp_asn: The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inside_cidr_blocks: The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        :param pulumi.Input[str] peer_address: The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] transit_gateway_address: The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        :param pulumi.Input[str] transit_gateway_attachment_id: The Transit Gateway Connect
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConnectPeerState.__new__(_ConnectPeerState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["bgp_asn"] = bgp_asn
        __props__.__dict__["inside_cidr_blocks"] = inside_cidr_blocks
        __props__.__dict__["peer_address"] = peer_address
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["transit_gateway_address"] = transit_gateway_address
        __props__.__dict__["transit_gateway_attachment_id"] = transit_gateway_attachment_id
        return ConnectPeer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        EC2 Transit Gateway Connect Peer ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> pulumi.Output[str]:
        """
        The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        """
        return pulumi.get(self, "bgp_asn")

    @property
    @pulumi.getter(name="insideCidrBlocks")
    def inside_cidr_blocks(self) -> pulumi.Output[Sequence[str]]:
        """
        The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        """
        return pulumi.get(self, "inside_cidr_blocks")

    @property
    @pulumi.getter(name="peerAddress")
    def peer_address(self) -> pulumi.Output[str]:
        """
        The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        """
        return pulumi.get(self, "peer_address")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="transitGatewayAddress")
    def transit_gateway_address(self) -> pulumi.Output[str]:
        """
        The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        """
        return pulumi.get(self, "transit_gateway_address")

    @property
    @pulumi.getter(name="transitGatewayAttachmentId")
    def transit_gateway_attachment_id(self) -> pulumi.Output[str]:
        """
        The Transit Gateway Connect
        """
        return pulumi.get(self, "transit_gateway_attachment_id")

