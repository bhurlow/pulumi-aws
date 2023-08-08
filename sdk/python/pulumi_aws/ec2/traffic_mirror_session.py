# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TrafficMirrorSessionArgs', 'TrafficMirrorSession']

@pulumi.input_type
class TrafficMirrorSessionArgs:
    def __init__(__self__, *,
                 network_interface_id: pulumi.Input[str],
                 session_number: pulumi.Input[int],
                 traffic_mirror_filter_id: pulumi.Input[str],
                 traffic_mirror_target_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 packet_length: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a TrafficMirrorSession resource.
        :param pulumi.Input[str] network_interface_id: ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        :param pulumi.Input[int] session_number: The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[int] packet_length: The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[int] virtual_network_id: The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        pulumi.set(__self__, "session_number", session_number)
        pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        pulumi.set(__self__, "traffic_mirror_target_id", traffic_mirror_target_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if packet_length is not None:
            pulumi.set(__self__, "packet_length", packet_length)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if virtual_network_id is not None:
            pulumi.set(__self__, "virtual_network_id", virtual_network_id)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Input[str]:
        """
        ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="sessionNumber")
    def session_number(self) -> pulumi.Input[int]:
        """
        The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        """
        return pulumi.get(self, "session_number")

    @session_number.setter
    def session_number(self, value: pulumi.Input[int]):
        pulumi.set(self, "session_number", value)

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Input[str]:
        """
        ID of the traffic mirror filter to be used
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @traffic_mirror_filter_id.setter
    def traffic_mirror_filter_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_mirror_filter_id", value)

    @property
    @pulumi.getter(name="trafficMirrorTargetId")
    def traffic_mirror_target_id(self) -> pulumi.Input[str]:
        """
        ID of the traffic mirror target to be used
        """
        return pulumi.get(self, "traffic_mirror_target_id")

    @traffic_mirror_target_id.setter
    def traffic_mirror_target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_mirror_target_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the traffic mirror session.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="packetLength")
    def packet_length(self) -> Optional[pulumi.Input[int]]:
        """
        The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        """
        return pulumi.get(self, "packet_length")

    @packet_length.setter
    def packet_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_length", value)

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
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> Optional[pulumi.Input[int]]:
        """
        The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        return pulumi.get(self, "virtual_network_id")

    @virtual_network_id.setter
    def virtual_network_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "virtual_network_id", value)


@pulumi.input_type
class _TrafficMirrorSessionState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 packet_length: Optional[pulumi.Input[int]] = None,
                 session_number: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_target_id: Optional[pulumi.Input[str]] = None,
                 virtual_network_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering TrafficMirrorSession resources.
        :param pulumi.Input[str] arn: The ARN of the traffic mirror session.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        :param pulumi.Input[str] owner_id: The AWS account ID of the session owner.
        :param pulumi.Input[int] packet_length: The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        :param pulumi.Input[int] session_number: The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[int] virtual_network_id: The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if packet_length is not None:
            pulumi.set(__self__, "packet_length", packet_length)
        if session_number is not None:
            pulumi.set(__self__, "session_number", session_number)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if traffic_mirror_filter_id is not None:
            pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        if traffic_mirror_target_id is not None:
            pulumi.set(__self__, "traffic_mirror_target_id", traffic_mirror_target_id)
        if virtual_network_id is not None:
            pulumi.set(__self__, "virtual_network_id", virtual_network_id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the traffic mirror session.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the traffic mirror session.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS account ID of the session owner.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter(name="packetLength")
    def packet_length(self) -> Optional[pulumi.Input[int]]:
        """
        The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        """
        return pulumi.get(self, "packet_length")

    @packet_length.setter
    def packet_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_length", value)

    @property
    @pulumi.getter(name="sessionNumber")
    def session_number(self) -> Optional[pulumi.Input[int]]:
        """
        The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        """
        return pulumi.get(self, "session_number")

    @session_number.setter
    def session_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "session_number", value)

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
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the traffic mirror filter to be used
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @traffic_mirror_filter_id.setter
    def traffic_mirror_filter_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_mirror_filter_id", value)

    @property
    @pulumi.getter(name="trafficMirrorTargetId")
    def traffic_mirror_target_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the traffic mirror target to be used
        """
        return pulumi.get(self, "traffic_mirror_target_id")

    @traffic_mirror_target_id.setter
    def traffic_mirror_target_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_mirror_target_id", value)

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> Optional[pulumi.Input[int]]:
        """
        The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        return pulumi.get(self, "virtual_network_id")

    @virtual_network_id.setter
    def virtual_network_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "virtual_network_id", value)


class TrafficMirrorSession(pulumi.CustomResource):
    @overload
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
                 __props__=None):
        """
        Provides an Traffic mirror session.\\
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
            session_number=1,
            traffic_mirror_filter_id=filter.id,
            traffic_mirror_target_id=target.id)
        ```

        ## Import

        terraform import {

         to = aws_ec2_traffic_mirror_session.session

         id = "tms-0d8aa3ca35897b82e" } Using `pulumi import`, import traffic mirror sessions using the `id`. For exampleconsole % pulumi import aws_ec2_traffic_mirror_session.session tms-0d8aa3ca35897b82e

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        :param pulumi.Input[int] packet_length: The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        :param pulumi.Input[int] session_number: The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[int] virtual_network_id: The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrafficMirrorSessionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Traffic mirror session.\\
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
            session_number=1,
            traffic_mirror_filter_id=filter.id,
            traffic_mirror_target_id=target.id)
        ```

        ## Import

        terraform import {

         to = aws_ec2_traffic_mirror_session.session

         id = "tms-0d8aa3ca35897b82e" } Using `pulumi import`, import traffic mirror sessions using the `id`. For exampleconsole % pulumi import aws_ec2_traffic_mirror_session.session tms-0d8aa3ca35897b82e

        :param str resource_name: The name of the resource.
        :param TrafficMirrorSessionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficMirrorSessionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
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
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrafficMirrorSessionArgs.__new__(TrafficMirrorSessionArgs)

            __props__.__dict__["description"] = description
            if network_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__.__dict__["network_interface_id"] = network_interface_id
            __props__.__dict__["packet_length"] = packet_length
            if session_number is None and not opts.urn:
                raise TypeError("Missing required property 'session_number'")
            __props__.__dict__["session_number"] = session_number
            __props__.__dict__["tags"] = tags
            if traffic_mirror_filter_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_filter_id'")
            __props__.__dict__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
            if traffic_mirror_target_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_target_id'")
            __props__.__dict__["traffic_mirror_target_id"] = traffic_mirror_target_id
            __props__.__dict__["virtual_network_id"] = virtual_network_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["owner_id"] = None
            __props__.__dict__["tags_all"] = None
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
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
        :param pulumi.Input[int] session_number: The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[int] virtual_network_id: The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrafficMirrorSessionState.__new__(_TrafficMirrorSessionState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["network_interface_id"] = network_interface_id
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["packet_length"] = packet_length
        __props__.__dict__["session_number"] = session_number
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
        __props__.__dict__["traffic_mirror_target_id"] = traffic_mirror_target_id
        __props__.__dict__["virtual_network_id"] = virtual_network_id
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
        The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
        """
        return pulumi.get(self, "session_number")

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
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

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
        The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        return pulumi.get(self, "virtual_network_id")

