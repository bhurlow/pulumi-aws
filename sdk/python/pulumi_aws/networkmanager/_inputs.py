# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DeviceAwsLocationArgs',
    'DeviceLocationArgs',
    'LinkBandwidthArgs',
    'SiteLocationArgs',
    'GetCoreNetworkPolicyDocumentAttachmentPolicyArgs',
    'GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs',
    'GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs',
    'GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs',
    'GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs',
    'GetCoreNetworkPolicyDocumentSegmentArgs',
    'GetCoreNetworkPolicyDocumentSegmentActionArgs',
]

@pulumi.input_type
class DeviceAwsLocationArgs:
    def __init__(__self__, *,
                 subnet_arn: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] subnet_arn: The Amazon Resource Name (ARN) of the subnet that the device is located in.
        :param pulumi.Input[str] zone: The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
        """
        if subnet_arn is not None:
            pulumi.set(__self__, "subnet_arn", subnet_arn)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="subnetArn")
    def subnet_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the subnet that the device is located in.
        """
        return pulumi.get(self, "subnet_arn")

    @subnet_arn.setter
    def subnet_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_arn", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class DeviceLocationArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 latitude: Optional[pulumi.Input[str]] = None,
                 longitude: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] address: The physical address.
        :param pulumi.Input[str] latitude: The latitude.
        :param pulumi.Input[str] longitude: The longitude.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The physical address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def latitude(self) -> Optional[pulumi.Input[str]]:
        """
        The latitude.
        """
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> Optional[pulumi.Input[str]]:
        """
        The longitude.
        """
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "longitude", value)


@pulumi.input_type
class LinkBandwidthArgs:
    def __init__(__self__, *,
                 download_speed: Optional[pulumi.Input[int]] = None,
                 upload_speed: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] download_speed: Download speed in Mbps.
        :param pulumi.Input[int] upload_speed: Upload speed in Mbps.
        """
        if download_speed is not None:
            pulumi.set(__self__, "download_speed", download_speed)
        if upload_speed is not None:
            pulumi.set(__self__, "upload_speed", upload_speed)

    @property
    @pulumi.getter(name="downloadSpeed")
    def download_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Download speed in Mbps.
        """
        return pulumi.get(self, "download_speed")

    @download_speed.setter
    def download_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "download_speed", value)

    @property
    @pulumi.getter(name="uploadSpeed")
    def upload_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Upload speed in Mbps.
        """
        return pulumi.get(self, "upload_speed")

    @upload_speed.setter
    def upload_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "upload_speed", value)


@pulumi.input_type
class SiteLocationArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 latitude: Optional[pulumi.Input[str]] = None,
                 longitude: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] address: Address of the location.
        :param pulumi.Input[str] latitude: Latitude of the location.
        :param pulumi.Input[str] longitude: Longitude of the location.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the location.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def latitude(self) -> Optional[pulumi.Input[str]]:
        """
        Latitude of the location.
        """
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> Optional[pulumi.Input[str]]:
        """
        Longitude of the location.
        """
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "longitude", value)


@pulumi.input_type
class GetCoreNetworkPolicyDocumentAttachmentPolicyArgs:
    def __init__(__self__, *,
                 action: 'GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs',
                 conditions: Sequence['GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs'],
                 rule_number: int,
                 condition_logic: Optional[str] = None,
                 description: Optional[str] = None):
        """
        :param 'GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs' action: The action to take for the chosen segment. Valid values `create-route` or `share`.
        :param Sequence['GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs'] conditions: A block argument. Detailed Below.
        :param int rule_number: An integer from `1` to `65535` indicating the rule's order number. Rules are processed in order from the lowest numbered rule to the highest. Rules stop processing when a rule is matched. It's important to make sure that you number your rules in the exact order that you want them processed.
        :param str condition_logic: Valid values include `and` or `or`. This is a mandatory parameter only if you have more than one condition. The `condition_logic` apply to all of the conditions for a rule, which also means nested conditions of `and` or `or` are not supported. Use `or` if you want to associate the attachment with the segment by either the segment name or attachment tag value, or by the chosen conditions. Use `and` if you want to associate the attachment with the segment by either the segment name or attachment tag value and by the chosen conditions. Detailed Below.
        :param str description: A user-defined string describing the segment action.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "conditions", conditions)
        pulumi.set(__self__, "rule_number", rule_number)
        if condition_logic is not None:
            pulumi.set(__self__, "condition_logic", condition_logic)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def action(self) -> 'GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs':
        """
        The action to take for the chosen segment. Valid values `create-route` or `share`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: 'GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs'):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def conditions(self) -> Sequence['GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs']:
        """
        A block argument. Detailed Below.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Sequence['GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs']):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="ruleNumber")
    def rule_number(self) -> int:
        """
        An integer from `1` to `65535` indicating the rule's order number. Rules are processed in order from the lowest numbered rule to the highest. Rules stop processing when a rule is matched. It's important to make sure that you number your rules in the exact order that you want them processed.
        """
        return pulumi.get(self, "rule_number")

    @rule_number.setter
    def rule_number(self, value: int):
        pulumi.set(self, "rule_number", value)

    @property
    @pulumi.getter(name="conditionLogic")
    def condition_logic(self) -> Optional[str]:
        """
        Valid values include `and` or `or`. This is a mandatory parameter only if you have more than one condition. The `condition_logic` apply to all of the conditions for a rule, which also means nested conditions of `and` or `or` are not supported. Use `or` if you want to associate the attachment with the segment by either the segment name or attachment tag value, or by the chosen conditions. Use `and` if you want to associate the attachment with the segment by either the segment name or attachment tag value and by the chosen conditions. Detailed Below.
        """
        return pulumi.get(self, "condition_logic")

    @condition_logic.setter
    def condition_logic(self, value: Optional[str]):
        pulumi.set(self, "condition_logic", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A user-defined string describing the segment action.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[str]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs:
    def __init__(__self__, *,
                 association_method: str,
                 require_acceptance: Optional[bool] = None,
                 segment: Optional[str] = None,
                 tag_value_of_key: Optional[str] = None):
        """
        :param str association_method: Defines how a segment is mapped. Values can be `constant` or `tag`. `constant` statically defines the segment to associate the attachment to. `tag` uses the value of a tag to dynamically try to map to a segment.reference_policies_elements_condition_operators.html) to evaluate.
        :param bool require_acceptance: Determines if this mapping should override the segment value for `require_attachment_acceptance`. You can only set this to `true`, indicating that this setting applies only to segments that have `require_attachment_acceptance` set to `false`. If the segment already has the default `require_attachment_acceptance`, you can set this to inherit segment’s acceptance value.
        :param str segment: The name of the segment.
        :param str tag_value_of_key: Maps the attachment to the value of a known key. This is used with the `association_method` is `tag`. For example a `tag` of `stage = “test”`, will map to a segment named `test`. The value must exactly match the name of a segment. This allows you to have many segments, but use only a single rule without having to define multiple nearly identical conditions. This prevents creating many similar conditions that all use the same keys to map to segments.
        """
        pulumi.set(__self__, "association_method", association_method)
        if require_acceptance is not None:
            pulumi.set(__self__, "require_acceptance", require_acceptance)
        if segment is not None:
            pulumi.set(__self__, "segment", segment)
        if tag_value_of_key is not None:
            pulumi.set(__self__, "tag_value_of_key", tag_value_of_key)

    @property
    @pulumi.getter(name="associationMethod")
    def association_method(self) -> str:
        """
        Defines how a segment is mapped. Values can be `constant` or `tag`. `constant` statically defines the segment to associate the attachment to. `tag` uses the value of a tag to dynamically try to map to a segment.reference_policies_elements_condition_operators.html) to evaluate.
        """
        return pulumi.get(self, "association_method")

    @association_method.setter
    def association_method(self, value: str):
        pulumi.set(self, "association_method", value)

    @property
    @pulumi.getter(name="requireAcceptance")
    def require_acceptance(self) -> Optional[bool]:
        """
        Determines if this mapping should override the segment value for `require_attachment_acceptance`. You can only set this to `true`, indicating that this setting applies only to segments that have `require_attachment_acceptance` set to `false`. If the segment already has the default `require_attachment_acceptance`, you can set this to inherit segment’s acceptance value.
        """
        return pulumi.get(self, "require_acceptance")

    @require_acceptance.setter
    def require_acceptance(self, value: Optional[bool]):
        pulumi.set(self, "require_acceptance", value)

    @property
    @pulumi.getter
    def segment(self) -> Optional[str]:
        """
        The name of the segment.
        """
        return pulumi.get(self, "segment")

    @segment.setter
    def segment(self, value: Optional[str]):
        pulumi.set(self, "segment", value)

    @property
    @pulumi.getter(name="tagValueOfKey")
    def tag_value_of_key(self) -> Optional[str]:
        """
        Maps the attachment to the value of a known key. This is used with the `association_method` is `tag`. For example a `tag` of `stage = “test”`, will map to a segment named `test`. The value must exactly match the name of a segment. This allows you to have many segments, but use only a single rule without having to define multiple nearly identical conditions. This prevents creating many similar conditions that all use the same keys to map to segments.
        """
        return pulumi.get(self, "tag_value_of_key")

    @tag_value_of_key.setter
    def tag_value_of_key(self, value: Optional[str]):
        pulumi.set(self, "tag_value_of_key", value)


@pulumi.input_type
class GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs:
    def __init__(__self__, *,
                 type: str,
                 key: Optional[str] = None,
                 operator: Optional[str] = None,
                 value: Optional[str] = None):
        """
        :param str type: Valid values include: `account-id`, `any`, `tag-value`, `tag-exists`, `resource-id`, `region`, `attachment-type`.
        :param str key: string value
        :param str operator: Valid values include: `equals`, `not-equals`, `contains`, `begins-with`.
        :param str value: string value
        """
        pulumi.set(__self__, "type", type)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Valid values include: `account-id`, `any`, `tag-value`, `tag-exists`, `resource-id`, `region`, `attachment-type`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: str):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        string value
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[str]:
        """
        Valid values include: `equals`, `not-equals`, `contains`, `begins-with`.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[str]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        string value
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs:
    def __init__(__self__, *,
                 asn_ranges: Sequence[str],
                 edge_locations: Sequence['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs'],
                 inside_cidr_blocks: Optional[Sequence[str]] = None,
                 vpn_ecmp_support: Optional[bool] = None):
        """
        :param Sequence[str] asn_ranges: List of strings containing Autonomous System Numbers (ASNs) to assign to Core Network Edges. By default, the core network automatically assigns an ASN for each Core Network Edge but you can optionally define the ASN in the edge-locations for each Region. The ASN uses an array of integer ranges only from `64512` to `65534` and `4200000000` to `4294967294` expressed as a string like `"64512-65534"`. No other ASN ranges can be used.
        :param Sequence['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs'] edge_locations: A list of strings of AWS Region names. Allows you to define a more restrictive set of Regions for a segment. The edge location must be a subset of the locations that are defined for `edge_locations` in the `core_network_configuration`.
        :param Sequence[str] inside_cidr_blocks: The local CIDR blocks for this Core Network Edge for AWS Transit Gateway Connect attachments. By default, this CIDR block will be one or more optional IPv4 and IPv6 CIDR prefixes auto-assigned from `inside_cidr_blocks`.
        :param bool vpn_ecmp_support: Indicates whether the core network forwards traffic over multiple equal-cost routes using VPN. The value can be either `true` or `false`. The default is `true`.
        """
        pulumi.set(__self__, "asn_ranges", asn_ranges)
        pulumi.set(__self__, "edge_locations", edge_locations)
        if inside_cidr_blocks is not None:
            pulumi.set(__self__, "inside_cidr_blocks", inside_cidr_blocks)
        if vpn_ecmp_support is not None:
            pulumi.set(__self__, "vpn_ecmp_support", vpn_ecmp_support)

    @property
    @pulumi.getter(name="asnRanges")
    def asn_ranges(self) -> Sequence[str]:
        """
        List of strings containing Autonomous System Numbers (ASNs) to assign to Core Network Edges. By default, the core network automatically assigns an ASN for each Core Network Edge but you can optionally define the ASN in the edge-locations for each Region. The ASN uses an array of integer ranges only from `64512` to `65534` and `4200000000` to `4294967294` expressed as a string like `"64512-65534"`. No other ASN ranges can be used.
        """
        return pulumi.get(self, "asn_ranges")

    @asn_ranges.setter
    def asn_ranges(self, value: Sequence[str]):
        pulumi.set(self, "asn_ranges", value)

    @property
    @pulumi.getter(name="edgeLocations")
    def edge_locations(self) -> Sequence['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs']:
        """
        A list of strings of AWS Region names. Allows you to define a more restrictive set of Regions for a segment. The edge location must be a subset of the locations that are defined for `edge_locations` in the `core_network_configuration`.
        """
        return pulumi.get(self, "edge_locations")

    @edge_locations.setter
    def edge_locations(self, value: Sequence['GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs']):
        pulumi.set(self, "edge_locations", value)

    @property
    @pulumi.getter(name="insideCidrBlocks")
    def inside_cidr_blocks(self) -> Optional[Sequence[str]]:
        """
        The local CIDR blocks for this Core Network Edge for AWS Transit Gateway Connect attachments. By default, this CIDR block will be one or more optional IPv4 and IPv6 CIDR prefixes auto-assigned from `inside_cidr_blocks`.
        """
        return pulumi.get(self, "inside_cidr_blocks")

    @inside_cidr_blocks.setter
    def inside_cidr_blocks(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "inside_cidr_blocks", value)

    @property
    @pulumi.getter(name="vpnEcmpSupport")
    def vpn_ecmp_support(self) -> Optional[bool]:
        """
        Indicates whether the core network forwards traffic over multiple equal-cost routes using VPN. The value can be either `true` or `false`. The default is `true`.
        """
        return pulumi.get(self, "vpn_ecmp_support")

    @vpn_ecmp_support.setter
    def vpn_ecmp_support(self, value: Optional[bool]):
        pulumi.set(self, "vpn_ecmp_support", value)


@pulumi.input_type
class GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs:
    def __init__(__self__, *,
                 location: str,
                 asn: Optional[int] = None,
                 inside_cidr_blocks: Optional[Sequence[str]] = None):
        """
        :param int asn: The ASN of the Core Network Edge in an AWS Region. By default, the ASN will be a single integer automatically assigned from `asn_ranges`
        :param Sequence[str] inside_cidr_blocks: The local CIDR blocks for this Core Network Edge for AWS Transit Gateway Connect attachments. By default, this CIDR block will be one or more optional IPv4 and IPv6 CIDR prefixes auto-assigned from `inside_cidr_blocks`.
        """
        pulumi.set(__self__, "location", location)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if inside_cidr_blocks is not None:
            pulumi.set(__self__, "inside_cidr_blocks", inside_cidr_blocks)

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: str):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def asn(self) -> Optional[int]:
        """
        The ASN of the Core Network Edge in an AWS Region. By default, the ASN will be a single integer automatically assigned from `asn_ranges`
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: Optional[int]):
        pulumi.set(self, "asn", value)

    @property
    @pulumi.getter(name="insideCidrBlocks")
    def inside_cidr_blocks(self) -> Optional[Sequence[str]]:
        """
        The local CIDR blocks for this Core Network Edge for AWS Transit Gateway Connect attachments. By default, this CIDR block will be one or more optional IPv4 and IPv6 CIDR prefixes auto-assigned from `inside_cidr_blocks`.
        """
        return pulumi.get(self, "inside_cidr_blocks")

    @inside_cidr_blocks.setter
    def inside_cidr_blocks(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "inside_cidr_blocks", value)


@pulumi.input_type
class GetCoreNetworkPolicyDocumentSegmentArgs:
    def __init__(__self__, *,
                 name: str,
                 allow_filters: Optional[Sequence[str]] = None,
                 deny_filters: Optional[Sequence[str]] = None,
                 description: Optional[str] = None,
                 edge_locations: Optional[Sequence[str]] = None,
                 isolate_attachments: Optional[bool] = None,
                 require_attachment_acceptance: Optional[bool] = None):
        """
        :param str name: A unique name for a segment. The name is a string used in other parts of the policy document, as well as in the console for metrics and other reference points. Valid characters are a–z, and 0–9.
        :param Sequence[str] allow_filters: List of strings of segment names that explicitly allows only routes from the segments that are listed in the array. Use the `allow_filter` setting if a segment has a well-defined group of other segments that connectivity should be restricted to. It is applied after routes have been shared in `segment_actions`. If a segment is listed in `allow_filter`, attachments between the two segments will have routes if they are also shared in the segment-actions area. For example, you might have a segment named "video-producer" that should only ever share routes with a "video-distributor" segment, no matter how many other share statements are created.
        :param Sequence[str] deny_filters: An array of segments that disallows routes from the segments listed in the array. It is applied only after routes have been shared in `segment_actions`. If a segment is listed in the `deny_filter`, attachments between the two segments will never have routes shared across them. For example, you might have a "financial" payment segment that should never share routes with a "development" segment, regardless of how many other share statements are created. Adding the payments segment to the deny-filter parameter prevents any shared routes from being created with other segments.
        :param str description: A user-defined string describing the segment action.
        :param Sequence[str] edge_locations: A list of strings of AWS Region names. Allows you to define a more restrictive set of Regions for a segment. The edge location must be a subset of the locations that are defined for `edge_locations` in the `core_network_configuration`.
        :param bool isolate_attachments: This Boolean setting determines whether attachments on the same segment can communicate with each other. If set to `true`, the only routes available will be either shared routes through the share actions, which are attachments in other segments, or static routes. The default value is `false`. For example, you might have a segment dedicated to "development" that should never allow VPCs to talk to each other, even if they’re on the same segment. In this example, you would keep the default parameter of `false`.
        :param bool require_attachment_acceptance: This Boolean setting determines whether attachment requests are automatically approved or require acceptance. The default is `true`, indicating that attachment requests require acceptance. For example, you might use this setting to allow a "sandbox" segment to allow any attachment request so that a core network or attachment administrator does not need to review and approve attachment requests. In this example, `require_attachment_acceptance` is set to `false`.
        """
        pulumi.set(__self__, "name", name)
        if allow_filters is not None:
            pulumi.set(__self__, "allow_filters", allow_filters)
        if deny_filters is not None:
            pulumi.set(__self__, "deny_filters", deny_filters)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if edge_locations is not None:
            pulumi.set(__self__, "edge_locations", edge_locations)
        if isolate_attachments is not None:
            pulumi.set(__self__, "isolate_attachments", isolate_attachments)
        if require_attachment_acceptance is not None:
            pulumi.set(__self__, "require_attachment_acceptance", require_attachment_acceptance)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A unique name for a segment. The name is a string used in other parts of the policy document, as well as in the console for metrics and other reference points. Valid characters are a–z, and 0–9.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="allowFilters")
    def allow_filters(self) -> Optional[Sequence[str]]:
        """
        List of strings of segment names that explicitly allows only routes from the segments that are listed in the array. Use the `allow_filter` setting if a segment has a well-defined group of other segments that connectivity should be restricted to. It is applied after routes have been shared in `segment_actions`. If a segment is listed in `allow_filter`, attachments between the two segments will have routes if they are also shared in the segment-actions area. For example, you might have a segment named "video-producer" that should only ever share routes with a "video-distributor" segment, no matter how many other share statements are created.
        """
        return pulumi.get(self, "allow_filters")

    @allow_filters.setter
    def allow_filters(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "allow_filters", value)

    @property
    @pulumi.getter(name="denyFilters")
    def deny_filters(self) -> Optional[Sequence[str]]:
        """
        An array of segments that disallows routes from the segments listed in the array. It is applied only after routes have been shared in `segment_actions`. If a segment is listed in the `deny_filter`, attachments between the two segments will never have routes shared across them. For example, you might have a "financial" payment segment that should never share routes with a "development" segment, regardless of how many other share statements are created. Adding the payments segment to the deny-filter parameter prevents any shared routes from being created with other segments.
        """
        return pulumi.get(self, "deny_filters")

    @deny_filters.setter
    def deny_filters(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "deny_filters", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A user-defined string describing the segment action.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="edgeLocations")
    def edge_locations(self) -> Optional[Sequence[str]]:
        """
        A list of strings of AWS Region names. Allows you to define a more restrictive set of Regions for a segment. The edge location must be a subset of the locations that are defined for `edge_locations` in the `core_network_configuration`.
        """
        return pulumi.get(self, "edge_locations")

    @edge_locations.setter
    def edge_locations(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "edge_locations", value)

    @property
    @pulumi.getter(name="isolateAttachments")
    def isolate_attachments(self) -> Optional[bool]:
        """
        This Boolean setting determines whether attachments on the same segment can communicate with each other. If set to `true`, the only routes available will be either shared routes through the share actions, which are attachments in other segments, or static routes. The default value is `false`. For example, you might have a segment dedicated to "development" that should never allow VPCs to talk to each other, even if they’re on the same segment. In this example, you would keep the default parameter of `false`.
        """
        return pulumi.get(self, "isolate_attachments")

    @isolate_attachments.setter
    def isolate_attachments(self, value: Optional[bool]):
        pulumi.set(self, "isolate_attachments", value)

    @property
    @pulumi.getter(name="requireAttachmentAcceptance")
    def require_attachment_acceptance(self) -> Optional[bool]:
        """
        This Boolean setting determines whether attachment requests are automatically approved or require acceptance. The default is `true`, indicating that attachment requests require acceptance. For example, you might use this setting to allow a "sandbox" segment to allow any attachment request so that a core network or attachment administrator does not need to review and approve attachment requests. In this example, `require_attachment_acceptance` is set to `false`.
        """
        return pulumi.get(self, "require_attachment_acceptance")

    @require_attachment_acceptance.setter
    def require_attachment_acceptance(self, value: Optional[bool]):
        pulumi.set(self, "require_attachment_acceptance", value)


@pulumi.input_type
class GetCoreNetworkPolicyDocumentSegmentActionArgs:
    def __init__(__self__, *,
                 action: str,
                 segment: str,
                 description: Optional[str] = None,
                 destination_cidr_blocks: Optional[Sequence[str]] = None,
                 destinations: Optional[Sequence[str]] = None,
                 mode: Optional[str] = None,
                 share_with_excepts: Optional[Sequence[str]] = None,
                 share_withs: Optional[Sequence[str]] = None):
        """
        :param str action: The action to take for the chosen segment. Valid values `create-route` or `share`.
        :param str segment: The name of the segment.
        :param str description: A user-defined string describing the segment action.
        :param Sequence[str] destination_cidr_blocks: List of strings containing CIDRs. You can define the IPv4 and IPv6 CIDR notation for each AWS Region. For example, `10.1.0.0/16` or `2001:db8::/56`. This is an array of CIDR notation strings.
        :param Sequence[str] destinations: A list of strings. Valid values include `["blackhole"]` or a list of attachment ids.
        :param str mode: A string. This mode places the attachment and return routes in each of the `share_with` segments. Valid values include: `attachment-route`.
        :param Sequence[str] share_with_excepts: A set subtraction of segments to not share with.
        :param Sequence[str] share_withs: A list of strings to share with. Must be a substring is all segments. Valid values include: `["*"]` or `["<segment-names>"]`.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "segment", segment)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_cidr_blocks is not None:
            pulumi.set(__self__, "destination_cidr_blocks", destination_cidr_blocks)
        if destinations is not None:
            pulumi.set(__self__, "destinations", destinations)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if share_with_excepts is not None:
            pulumi.set(__self__, "share_with_excepts", share_with_excepts)
        if share_withs is not None:
            pulumi.set(__self__, "share_withs", share_withs)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        The action to take for the chosen segment. Valid values `create-route` or `share`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: str):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def segment(self) -> str:
        """
        The name of the segment.
        """
        return pulumi.get(self, "segment")

    @segment.setter
    def segment(self, value: str):
        pulumi.set(self, "segment", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A user-defined string describing the segment action.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationCidrBlocks")
    def destination_cidr_blocks(self) -> Optional[Sequence[str]]:
        """
        List of strings containing CIDRs. You can define the IPv4 and IPv6 CIDR notation for each AWS Region. For example, `10.1.0.0/16` or `2001:db8::/56`. This is an array of CIDR notation strings.
        """
        return pulumi.get(self, "destination_cidr_blocks")

    @destination_cidr_blocks.setter
    def destination_cidr_blocks(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "destination_cidr_blocks", value)

    @property
    @pulumi.getter
    def destinations(self) -> Optional[Sequence[str]]:
        """
        A list of strings. Valid values include `["blackhole"]` or a list of attachment ids.
        """
        return pulumi.get(self, "destinations")

    @destinations.setter
    def destinations(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "destinations", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[str]:
        """
        A string. This mode places the attachment and return routes in each of the `share_with` segments. Valid values include: `attachment-route`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="shareWithExcepts")
    def share_with_excepts(self) -> Optional[Sequence[str]]:
        """
        A set subtraction of segments to not share with.
        """
        return pulumi.get(self, "share_with_excepts")

    @share_with_excepts.setter
    def share_with_excepts(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "share_with_excepts", value)

    @property
    @pulumi.getter(name="shareWiths")
    def share_withs(self) -> Optional[Sequence[str]]:
        """
        A list of strings to share with. Must be a substring is all segments. Valid values include: `["*"]` or `["<segment-names>"]`.
        """
        return pulumi.get(self, "share_withs")

    @share_withs.setter
    def share_withs(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "share_withs", value)

