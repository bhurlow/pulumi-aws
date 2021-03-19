# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VpcPeeringConnection']


class VpcPeeringConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepter: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterArgs']]] = None,
                 auto_accept: Optional[pulumi.Input[bool]] = None,
                 peer_owner_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_vpc_id: Optional[pulumi.Input[str]] = None,
                 requester: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionRequesterArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage a VPC peering connection.

        > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
        both a standalone VPC Peering Connection Options and a VPC Peering Connection
        resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
        connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
        Doing so will cause a conflict of options and will overwrite the options.
        Using a VPC Peering Connection Options resource decouples management of the connection options from
        management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.

        > **Note:** For cross-account (requester's AWS account differs from the accepter's AWS account) or inter-region
        VPC Peering Connections use the `ec2.VpcPeeringConnection` resource to manage the requester's side of the
        connection and use the `ec2.VpcPeeringConnectionAccepter` resource to manage the accepter's side of the connection.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ec2.VpcPeeringConnection("foo",
            peer_owner_id=var["peer_owner_id"],
            peer_vpc_id=aws_vpc["bar"]["id"],
            vpc_id=aws_vpc["foo"]["id"])
        ```

        Basic usage with connection options:

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ec2.VpcPeeringConnection("foo",
            peer_owner_id=var["peer_owner_id"],
            peer_vpc_id=aws_vpc["bar"]["id"],
            vpc_id=aws_vpc["foo"]["id"],
            accepter=aws.ec2.VpcPeeringConnectionAccepterArgs(
                allow_remote_vpc_dns_resolution=True,
            ),
            requester=aws.ec2.VpcPeeringConnectionRequesterArgs(
                allow_remote_vpc_dns_resolution=True,
            ))
        ```

        Basic usage with tags:

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16")
        bar = aws.ec2.Vpc("bar", cidr_block="10.2.0.0/16")
        foo_vpc_peering_connection = aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection",
            peer_owner_id=var["peer_owner_id"],
            peer_vpc_id=bar.id,
            vpc_id=foo_vpc.id,
            auto_accept=True,
            tags={
                "Name": "VPC Peering between foo and bar",
            })
        ```

        Basic usage with region:

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16",
        opts=pulumi.ResourceOptions(provider=aws["us-west-2"]))
        bar = aws.ec2.Vpc("bar", cidr_block="10.2.0.0/16",
        opts=pulumi.ResourceOptions(provider=aws["us-east-1"]))
        foo_vpc_peering_connection = aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection",
            peer_owner_id=var["peer_owner_id"],
            peer_vpc_id=bar.id,
            vpc_id=foo_vpc.id,
            peer_region="us-east-1")
        ```
        ## Notes

        If both VPCs are not in the same AWS account do not enable the `auto_accept` attribute.
        The accepter can manage its side of the connection using the `ec2.VpcPeeringConnectionAccepter` resource
        or accept the connection manually using the AWS Management Console, AWS CLI, through SDKs, etc.

        ## Import

        VPC Peering resources can be imported using the `vpc peering id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcPeeringConnection:VpcPeeringConnection test_connection pcx-111aaa111
        ```

         [1]/docs/providers/aws/index.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterArgs']] accepter: An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
               the peering connection (a maximum of one).
        :param pulumi.Input[bool] auto_accept: Accept the peering (both VPCs need to be in the same AWS account).
        :param pulumi.Input[str] peer_owner_id: The AWS account ID of the owner of the peer VPC.
               Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
        :param pulumi.Input[str] peer_region: The region of the accepter VPC of the VPC Peering Connection. `auto_accept` must be `false`,
               and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
        :param pulumi.Input[str] peer_vpc_id: The ID of the VPC with which you are creating the VPC Peering Connection.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionRequesterArgs']] requester: A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
               the peering connection (a maximum of one).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the requester VPC.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['accepter'] = accepter
            __props__['auto_accept'] = auto_accept
            __props__['peer_owner_id'] = peer_owner_id
            __props__['peer_region'] = peer_region
            if peer_vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'peer_vpc_id'")
            __props__['peer_vpc_id'] = peer_vpc_id
            __props__['requester'] = requester
            __props__['tags'] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['accept_status'] = None
        super(VpcPeeringConnection, __self__).__init__(
            'aws:ec2/vpcPeeringConnection:VpcPeeringConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accept_status: Optional[pulumi.Input[str]] = None,
            accepter: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterArgs']]] = None,
            auto_accept: Optional[pulumi.Input[bool]] = None,
            peer_owner_id: Optional[pulumi.Input[str]] = None,
            peer_region: Optional[pulumi.Input[str]] = None,
            peer_vpc_id: Optional[pulumi.Input[str]] = None,
            requester: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionRequesterArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'VpcPeeringConnection':
        """
        Get an existing VpcPeeringConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_status: The status of the VPC Peering Connection request.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterArgs']] accepter: An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
               the peering connection (a maximum of one).
        :param pulumi.Input[bool] auto_accept: Accept the peering (both VPCs need to be in the same AWS account).
        :param pulumi.Input[str] peer_owner_id: The AWS account ID of the owner of the peer VPC.
               Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
        :param pulumi.Input[str] peer_region: The region of the accepter VPC of the VPC Peering Connection. `auto_accept` must be `false`,
               and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
        :param pulumi.Input[str] peer_vpc_id: The ID of the VPC with which you are creating the VPC Peering Connection.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionRequesterArgs']] requester: A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
               the peering connection (a maximum of one).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the requester VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accept_status"] = accept_status
        __props__["accepter"] = accepter
        __props__["auto_accept"] = auto_accept
        __props__["peer_owner_id"] = peer_owner_id
        __props__["peer_region"] = peer_region
        __props__["peer_vpc_id"] = peer_vpc_id
        __props__["requester"] = requester
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return VpcPeeringConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptStatus")
    def accept_status(self) -> pulumi.Output[str]:
        """
        The status of the VPC Peering Connection request.
        """
        return pulumi.get(self, "accept_status")

    @property
    @pulumi.getter
    def accepter(self) -> pulumi.Output['outputs.VpcPeeringConnectionAccepter']:
        """
        An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
        the peering connection (a maximum of one).
        """
        return pulumi.get(self, "accepter")

    @property
    @pulumi.getter(name="autoAccept")
    def auto_accept(self) -> pulumi.Output[Optional[bool]]:
        """
        Accept the peering (both VPCs need to be in the same AWS account).
        """
        return pulumi.get(self, "auto_accept")

    @property
    @pulumi.getter(name="peerOwnerId")
    def peer_owner_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID of the owner of the peer VPC.
        Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
        """
        return pulumi.get(self, "peer_owner_id")

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> pulumi.Output[str]:
        """
        The region of the accepter VPC of the VPC Peering Connection. `auto_accept` must be `false`,
        and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
        """
        return pulumi.get(self, "peer_region")

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC with which you are creating the VPC Peering Connection.
        """
        return pulumi.get(self, "peer_vpc_id")

    @property
    @pulumi.getter
    def requester(self) -> pulumi.Output['outputs.VpcPeeringConnectionRequester']:
        """
        A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
        the peering connection (a maximum of one).
        """
        return pulumi.get(self, "requester")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the requester VPC.
        """
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

