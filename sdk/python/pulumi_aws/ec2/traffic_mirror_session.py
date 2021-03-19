# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['TrafficMirrorSession']


class TrafficMirrorSession(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 packet_length: Optional[pulumi.Input[int]] = None,
                 session_number: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_target_id: Optional[pulumi.Input[str]] = None,
                 virtual_network_id: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Traffic mirror session.\
        Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring

        ## Example Usage

        To create a basic traffic mirror session

        ```python
        import pulumi
        import pulumi_aws as aws

        filter = aws.ec2.TrafficMirrorFilter("filter",
            description="traffic mirror filter - example",
            network_services=["amazon-dns"])
        target = aws.ec2.TrafficMirrorTarget("target", network_load_balancer_arn=aws_lb["lb"]["arn"])
        session = aws.ec2.TrafficMirrorSession("session",
            description="traffic mirror session - example",
            network_interface_id=aws_instance["test"]["primary_network_interface_id"],
            traffic_mirror_filter_id=filter.id,
            traffic_mirror_target_id=target.id)
        ```

        ## Import

        Traffic mirror sessions can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/trafficMirrorSession:TrafficMirrorSession session tms-0d8aa3ca35897b82e
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        :param pulumi.Input[int] packet_length: The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        :param pulumi.Input[int] session_number: - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[int] virtual_network_id: - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
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

            __props__['description'] = description
            if network_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__['network_interface_id'] = network_interface_id
            __props__['packet_length'] = packet_length
            if session_number is None and not opts.urn:
                raise TypeError("Missing required property 'session_number'")
            __props__['session_number'] = session_number
            __props__['tags'] = tags
            if traffic_mirror_filter_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_filter_id'")
            __props__['traffic_mirror_filter_id'] = traffic_mirror_filter_id
            if traffic_mirror_target_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_target_id'")
            __props__['traffic_mirror_target_id'] = traffic_mirror_target_id
            __props__['virtual_network_id'] = virtual_network_id
            __props__['arn'] = None
            __props__['owner_id'] = None
        super(TrafficMirrorSession, __self__).__init__(
            'aws:ec2/trafficMirrorSession:TrafficMirrorSession',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            packet_length: Optional[pulumi.Input[int]] = None,
            session_number: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
            traffic_mirror_target_id: Optional[pulumi.Input[str]] = None,
            virtual_network_id: Optional[pulumi.Input[int]] = None) -> 'TrafficMirrorSession':
        """
        Get an existing TrafficMirrorSession resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the traffic mirror session.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        :param pulumi.Input[str] owner_id: The AWS account ID of the session owner.
        :param pulumi.Input[int] packet_length: The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        :param pulumi.Input[int] session_number: - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[int] virtual_network_id: - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["network_interface_id"] = network_interface_id
        __props__["owner_id"] = owner_id
        __props__["packet_length"] = packet_length
        __props__["session_number"] = session_number
        __props__["tags"] = tags
        __props__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
        __props__["traffic_mirror_target_id"] = traffic_mirror_target_id
        __props__["virtual_network_id"] = virtual_network_id
        return TrafficMirrorSession(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the traffic mirror session.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the traffic mirror session.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID of the session owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="packetLength")
    def packet_length(self) -> pulumi.Output[Optional[int]]:
        """
        The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        """
        return pulumi.get(self, "packet_length")

    @property
    @pulumi.getter(name="sessionNumber")
    def session_number(self) -> pulumi.Output[int]:
        """
        - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        """
        return pulumi.get(self, "session_number")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Output[str]:
        """
        ID of the traffic mirror filter to be used
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @property
    @pulumi.getter(name="trafficMirrorTargetId")
    def traffic_mirror_target_id(self) -> pulumi.Output[str]:
        """
        ID of the traffic mirror target to be used
        """
        return pulumi.get(self, "traffic_mirror_target_id")

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> pulumi.Output[int]:
        """
        - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        return pulumi.get(self, "virtual_network_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

