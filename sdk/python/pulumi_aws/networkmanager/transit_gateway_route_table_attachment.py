# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransitGatewayRouteTableAttachmentArgs', 'TransitGatewayRouteTableAttachment']

@pulumi.input_type
class TransitGatewayRouteTableAttachmentArgs:
    def __init__(__self__, *,
                 peering_id: pulumi.Input[str],
                 transit_gateway_route_table_arn: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a TransitGatewayRouteTableAttachment resource.
        :param pulumi.Input[str] peering_id: The ID of the peer for the attachment.
        :param pulumi.Input[str] transit_gateway_route_table_arn: The ARN of the transit gateway route table for the attachment.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "peering_id", peering_id)
        pulumi.set(__self__, "transit_gateway_route_table_arn", transit_gateway_route_table_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="peeringId")
    def peering_id(self) -> pulumi.Input[str]:
        """
        The ID of the peer for the attachment.
        """
        return pulumi.get(self, "peering_id")

    @peering_id.setter
    def peering_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peering_id", value)

    @property
    @pulumi.getter(name="transitGatewayRouteTableArn")
    def transit_gateway_route_table_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the transit gateway route table for the attachment.
        """
        return pulumi.get(self, "transit_gateway_route_table_arn")

    @transit_gateway_route_table_arn.setter
    def transit_gateway_route_table_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_gateway_route_table_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _TransitGatewayRouteTableAttachmentState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 attachment_policy_rule_number: Optional[pulumi.Input[int]] = None,
                 attachment_type: Optional[pulumi.Input[str]] = None,
                 core_network_arn: Optional[pulumi.Input[str]] = None,
                 core_network_id: Optional[pulumi.Input[str]] = None,
                 edge_location: Optional[pulumi.Input[str]] = None,
                 owner_account_id: Optional[pulumi.Input[str]] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 resource_arn: Optional[pulumi.Input[str]] = None,
                 segment_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_route_table_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransitGatewayRouteTableAttachment resources.
        :param pulumi.Input[str] arn: Attachment Amazon Resource Name (ARN).
        :param pulumi.Input[int] attachment_policy_rule_number: The policy rule number associated with the attachment.
        :param pulumi.Input[str] attachment_type: The type of attachment.
        :param pulumi.Input[str] core_network_arn: The ARN of the core network.
        :param pulumi.Input[str] core_network_id: The ID of the core network.
        :param pulumi.Input[str] edge_location: The edge location for the peer.
        :param pulumi.Input[str] owner_account_id: The ID of the attachment account owner.
        :param pulumi.Input[str] peering_id: The ID of the peer for the attachment.
        :param pulumi.Input[str] resource_arn: The attachment resource ARN.
        :param pulumi.Input[str] segment_name: The name of the segment attachment.
        :param pulumi.Input[str] state: The state of the attachment.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] transit_gateway_route_table_arn: The ARN of the transit gateway route table for the attachment.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if attachment_policy_rule_number is not None:
            pulumi.set(__self__, "attachment_policy_rule_number", attachment_policy_rule_number)
        if attachment_type is not None:
            pulumi.set(__self__, "attachment_type", attachment_type)
        if core_network_arn is not None:
            pulumi.set(__self__, "core_network_arn", core_network_arn)
        if core_network_id is not None:
            pulumi.set(__self__, "core_network_id", core_network_id)
        if edge_location is not None:
            pulumi.set(__self__, "edge_location", edge_location)
        if owner_account_id is not None:
            pulumi.set(__self__, "owner_account_id", owner_account_id)
        if peering_id is not None:
            pulumi.set(__self__, "peering_id", peering_id)
        if resource_arn is not None:
            pulumi.set(__self__, "resource_arn", resource_arn)
        if segment_name is not None:
            pulumi.set(__self__, "segment_name", segment_name)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if transit_gateway_route_table_arn is not None:
            pulumi.set(__self__, "transit_gateway_route_table_arn", transit_gateway_route_table_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Attachment Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="attachmentPolicyRuleNumber")
    def attachment_policy_rule_number(self) -> Optional[pulumi.Input[int]]:
        """
        The policy rule number associated with the attachment.
        """
        return pulumi.get(self, "attachment_policy_rule_number")

    @attachment_policy_rule_number.setter
    def attachment_policy_rule_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "attachment_policy_rule_number", value)

    @property
    @pulumi.getter(name="attachmentType")
    def attachment_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of attachment.
        """
        return pulumi.get(self, "attachment_type")

    @attachment_type.setter
    def attachment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attachment_type", value)

    @property
    @pulumi.getter(name="coreNetworkArn")
    def core_network_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the core network.
        """
        return pulumi.get(self, "core_network_arn")

    @core_network_arn.setter
    def core_network_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "core_network_arn", value)

    @property
    @pulumi.getter(name="coreNetworkId")
    def core_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the core network.
        """
        return pulumi.get(self, "core_network_id")

    @core_network_id.setter
    def core_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "core_network_id", value)

    @property
    @pulumi.getter(name="edgeLocation")
    def edge_location(self) -> Optional[pulumi.Input[str]]:
        """
        The edge location for the peer.
        """
        return pulumi.get(self, "edge_location")

    @edge_location.setter
    def edge_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_location", value)

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the attachment account owner.
        """
        return pulumi.get(self, "owner_account_id")

    @owner_account_id.setter
    def owner_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_account_id", value)

    @property
    @pulumi.getter(name="peeringId")
    def peering_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the peer for the attachment.
        """
        return pulumi.get(self, "peering_id")

    @peering_id.setter
    def peering_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_id", value)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The attachment resource ARN.
        """
        return pulumi.get(self, "resource_arn")

    @resource_arn.setter
    def resource_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_arn", value)

    @property
    @pulumi.getter(name="segmentName")
    def segment_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the segment attachment.
        """
        return pulumi.get(self, "segment_name")

    @segment_name.setter
    def segment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "segment_name", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the attachment.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    @pulumi.getter(name="transitGatewayRouteTableArn")
    def transit_gateway_route_table_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the transit gateway route table for the attachment.
        """
        return pulumi.get(self, "transit_gateway_route_table_arn")

    @transit_gateway_route_table_arn.setter
    def transit_gateway_route_table_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_gateway_route_table_arn", value)


class TransitGatewayRouteTableAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_route_table_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a transit gateway route table attachment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.networkmanager.TransitGatewayRouteTableAttachment("example",
            peering_id=aws_networkmanager_transit_gateway_peering["example"]["id"],
            transit_gateway_route_table_arn=aws_ec2_transit_gateway_route_table["example"]["arn"])
        ```

        ## Import

        `aws_networkmanager_transit_gateway_route_table_attachment` can be imported using the attachment ID, e.g.

        ```sh
         $ pulumi import aws:networkmanager/transitGatewayRouteTableAttachment:TransitGatewayRouteTableAttachment example attachment-0f8fa60d2238d1bd8
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peering_id: The ID of the peer for the attachment.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] transit_gateway_route_table_arn: The ARN of the transit gateway route table for the attachment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransitGatewayRouteTableAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a transit gateway route table attachment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.networkmanager.TransitGatewayRouteTableAttachment("example",
            peering_id=aws_networkmanager_transit_gateway_peering["example"]["id"],
            transit_gateway_route_table_arn=aws_ec2_transit_gateway_route_table["example"]["arn"])
        ```

        ## Import

        `aws_networkmanager_transit_gateway_route_table_attachment` can be imported using the attachment ID, e.g.

        ```sh
         $ pulumi import aws:networkmanager/transitGatewayRouteTableAttachment:TransitGatewayRouteTableAttachment example attachment-0f8fa60d2238d1bd8
        ```

        :param str resource_name: The name of the resource.
        :param TransitGatewayRouteTableAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitGatewayRouteTableAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peering_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_route_table_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransitGatewayRouteTableAttachmentArgs.__new__(TransitGatewayRouteTableAttachmentArgs)

            if peering_id is None and not opts.urn:
                raise TypeError("Missing required property 'peering_id'")
            __props__.__dict__["peering_id"] = peering_id
            __props__.__dict__["tags"] = tags
            if transit_gateway_route_table_arn is None and not opts.urn:
                raise TypeError("Missing required property 'transit_gateway_route_table_arn'")
            __props__.__dict__["transit_gateway_route_table_arn"] = transit_gateway_route_table_arn
            __props__.__dict__["arn"] = None
            __props__.__dict__["attachment_policy_rule_number"] = None
            __props__.__dict__["attachment_type"] = None
            __props__.__dict__["core_network_arn"] = None
            __props__.__dict__["core_network_id"] = None
            __props__.__dict__["edge_location"] = None
            __props__.__dict__["owner_account_id"] = None
            __props__.__dict__["resource_arn"] = None
            __props__.__dict__["segment_name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["tags_all"] = None
        super(TransitGatewayRouteTableAttachment, __self__).__init__(
            'aws:networkmanager/transitGatewayRouteTableAttachment:TransitGatewayRouteTableAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            attachment_policy_rule_number: Optional[pulumi.Input[int]] = None,
            attachment_type: Optional[pulumi.Input[str]] = None,
            core_network_arn: Optional[pulumi.Input[str]] = None,
            core_network_id: Optional[pulumi.Input[str]] = None,
            edge_location: Optional[pulumi.Input[str]] = None,
            owner_account_id: Optional[pulumi.Input[str]] = None,
            peering_id: Optional[pulumi.Input[str]] = None,
            resource_arn: Optional[pulumi.Input[str]] = None,
            segment_name: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            transit_gateway_route_table_arn: Optional[pulumi.Input[str]] = None) -> 'TransitGatewayRouteTableAttachment':
        """
        Get an existing TransitGatewayRouteTableAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Attachment Amazon Resource Name (ARN).
        :param pulumi.Input[int] attachment_policy_rule_number: The policy rule number associated with the attachment.
        :param pulumi.Input[str] attachment_type: The type of attachment.
        :param pulumi.Input[str] core_network_arn: The ARN of the core network.
        :param pulumi.Input[str] core_network_id: The ID of the core network.
        :param pulumi.Input[str] edge_location: The edge location for the peer.
        :param pulumi.Input[str] owner_account_id: The ID of the attachment account owner.
        :param pulumi.Input[str] peering_id: The ID of the peer for the attachment.
        :param pulumi.Input[str] resource_arn: The attachment resource ARN.
        :param pulumi.Input[str] segment_name: The name of the segment attachment.
        :param pulumi.Input[str] state: The state of the attachment.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] transit_gateway_route_table_arn: The ARN of the transit gateway route table for the attachment.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransitGatewayRouteTableAttachmentState.__new__(_TransitGatewayRouteTableAttachmentState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["attachment_policy_rule_number"] = attachment_policy_rule_number
        __props__.__dict__["attachment_type"] = attachment_type
        __props__.__dict__["core_network_arn"] = core_network_arn
        __props__.__dict__["core_network_id"] = core_network_id
        __props__.__dict__["edge_location"] = edge_location
        __props__.__dict__["owner_account_id"] = owner_account_id
        __props__.__dict__["peering_id"] = peering_id
        __props__.__dict__["resource_arn"] = resource_arn
        __props__.__dict__["segment_name"] = segment_name
        __props__.__dict__["state"] = state
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["transit_gateway_route_table_arn"] = transit_gateway_route_table_arn
        return TransitGatewayRouteTableAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Attachment Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="attachmentPolicyRuleNumber")
    def attachment_policy_rule_number(self) -> pulumi.Output[int]:
        """
        The policy rule number associated with the attachment.
        """
        return pulumi.get(self, "attachment_policy_rule_number")

    @property
    @pulumi.getter(name="attachmentType")
    def attachment_type(self) -> pulumi.Output[str]:
        """
        The type of attachment.
        """
        return pulumi.get(self, "attachment_type")

    @property
    @pulumi.getter(name="coreNetworkArn")
    def core_network_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the core network.
        """
        return pulumi.get(self, "core_network_arn")

    @property
    @pulumi.getter(name="coreNetworkId")
    def core_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the core network.
        """
        return pulumi.get(self, "core_network_id")

    @property
    @pulumi.getter(name="edgeLocation")
    def edge_location(self) -> pulumi.Output[str]:
        """
        The edge location for the peer.
        """
        return pulumi.get(self, "edge_location")

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the attachment account owner.
        """
        return pulumi.get(self, "owner_account_id")

    @property
    @pulumi.getter(name="peeringId")
    def peering_id(self) -> pulumi.Output[str]:
        """
        The ID of the peer for the attachment.
        """
        return pulumi.get(self, "peering_id")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Output[str]:
        """
        The attachment resource ARN.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="segmentName")
    def segment_name(self) -> pulumi.Output[str]:
        """
        The name of the segment attachment.
        """
        return pulumi.get(self, "segment_name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the attachment.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    @pulumi.getter(name="transitGatewayRouteTableArn")
    def transit_gateway_route_table_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the transit gateway route table for the attachment.
        """
        return pulumi.get(self, "transit_gateway_route_table_arn")

