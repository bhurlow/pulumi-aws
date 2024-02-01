# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceConnectEndpointTimeouts',
    'GetAttachmentFilterResult',
    'GetAttachmentsFilterResult',
    'GetConnectFilterResult',
    'GetConnectPeerFilterResult',
    'GetDirectConnectGatewayAttachmentFilterResult',
    'GetMulticastDomainAssociationResult',
    'GetMulticastDomainFilterResult',
    'GetMulticastDomainMemberResult',
    'GetMulticastDomainSourceResult',
    'GetPeeringAttachmentFilterResult',
    'GetRouteTableAssociationsFilterResult',
    'GetRouteTableFilterResult',
    'GetRouteTablePropagationsFilterResult',
    'GetRouteTableRoutesFilterResult',
    'GetRouteTableRoutesRouteResult',
    'GetTransitGatewayFilterResult',
    'GetVpcAttachmentFilterResult',
    'GetVpcAttachmentsFilterResult',
    'GetVpnAttachmentFilterResult',
]

@pulumi.output_type
class InstanceConnectEndpointTimeouts(dict):
    def __init__(__self__, *,
                 create: Optional[str] = None,
                 delete: Optional[str] = None):
        """
        :param str create: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        :param str delete: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
        """
        if create is not None:
            pulumi.set(__self__, "create", create)
        if delete is not None:
            pulumi.set(__self__, "delete", delete)

    @property
    @pulumi.getter
    def create(self) -> Optional[str]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        """
        return pulumi.get(self, "create")

    @property
    @pulumi.getter
    def delete(self) -> Optional[str]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
        """
        return pulumi.get(self, "delete")


@pulumi.output_type
class GetAttachmentFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the field to filter by, as defined by the [underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html).
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the field to filter by, as defined by the [underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetAttachmentsFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter check available value on [official documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html)
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter check available value on [official documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetConnectFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter.
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetConnectPeerFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter.
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetDirectConnectGatewayAttachmentFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter field. Valid values can be found in the [EC2 DescribeTransitGatewayAttachments API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html).
        :param Sequence[str] values: Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter field. Valid values can be found in the [EC2 DescribeTransitGatewayAttachments API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetMulticastDomainAssociationResult(dict):
    def __init__(__self__, *,
                 subnet_id: str,
                 transit_gateway_attachment_id: str):
        """
        :param str subnet_id: The ID of the subnet associated with the transit gateway multicast domain.
        :param str transit_gateway_attachment_id: The ID of the transit gateway attachment.
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "transit_gateway_attachment_id", transit_gateway_attachment_id)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of the subnet associated with the transit gateway multicast domain.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="transitGatewayAttachmentId")
    def transit_gateway_attachment_id(self) -> str:
        """
        The ID of the transit gateway attachment.
        """
        return pulumi.get(self, "transit_gateway_attachment_id")


@pulumi.output_type
class GetMulticastDomainFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayMulticastDomains.html).
        :param Sequence[str] values: Set of values that are accepted for the given field. A multicast domain will be selected if any one of the given values matches.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayMulticastDomains.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given field. A multicast domain will be selected if any one of the given values matches.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetMulticastDomainMemberResult(dict):
    def __init__(__self__, *,
                 group_ip_address: str,
                 network_interface_id: str):
        """
        :param str group_ip_address: The IP address assigned to the transit gateway multicast group.
        :param str network_interface_id: The group members' network interface ID.
        """
        pulumi.set(__self__, "group_ip_address", group_ip_address)
        pulumi.set(__self__, "network_interface_id", network_interface_id)

    @property
    @pulumi.getter(name="groupIpAddress")
    def group_ip_address(self) -> str:
        """
        The IP address assigned to the transit gateway multicast group.
        """
        return pulumi.get(self, "group_ip_address")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        """
        The group members' network interface ID.
        """
        return pulumi.get(self, "network_interface_id")


@pulumi.output_type
class GetMulticastDomainSourceResult(dict):
    def __init__(__self__, *,
                 group_ip_address: str,
                 network_interface_id: str):
        """
        :param str group_ip_address: The IP address assigned to the transit gateway multicast group.
        :param str network_interface_id: The group members' network interface ID.
        """
        pulumi.set(__self__, "group_ip_address", group_ip_address)
        pulumi.set(__self__, "network_interface_id", network_interface_id)

    @property
    @pulumi.getter(name="groupIpAddress")
    def group_ip_address(self) -> str:
        """
        The IP address assigned to the transit gateway multicast group.
        """
        return pulumi.get(self, "group_ip_address")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        """
        The group members' network interface ID.
        """
        return pulumi.get(self, "network_interface_id")


@pulumi.output_type
class GetPeeringAttachmentFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the field to filter by, as defined by
               [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayPeeringAttachments.html).
        :param Sequence[str] values: Set of values that are accepted for the given field.
               An EC2 Transit Gateway Peering Attachment be selected if any one of the given values matches.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the field to filter by, as defined by
        [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayPeeringAttachments.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given field.
        An EC2 Transit Gateway Peering Attachment be selected if any one of the given values matches.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetRouteTableAssociationsFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the field to filter by, as defined by
               [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetTransitGatewayRouteTableAssociations.html).
        :param Sequence[str] values: Set of values that are accepted for the given field.
               A Transit Gateway Route Table will be selected if any one of the given values matches.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the field to filter by, as defined by
        [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetTransitGatewayRouteTableAssociations.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given field.
        A Transit Gateway Route Table will be selected if any one of the given values matches.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetRouteTableFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter.
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetRouteTablePropagationsFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the field to filter by, as defined by
               [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetTransitGatewayRouteTablePropagations.html).
        :param Sequence[str] values: Set of values that are accepted for the given field.
               A Transit Gateway Route Table will be selected if any one of the given values matches.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the field to filter by, as defined by
        [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetTransitGatewayRouteTablePropagations.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given field.
        A Transit Gateway Route Table will be selected if any one of the given values matches.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetRouteTableRoutesFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the field to filter by, as defined by
               [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayRoutes.html).
        :param Sequence[str] values: Set of values that are accepted for the given field.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the field to filter by, as defined by
        [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayRoutes.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given field.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetRouteTableRoutesRouteResult(dict):
    def __init__(__self__, *,
                 destination_cidr_block: str,
                 prefix_list_id: str,
                 state: str,
                 transit_gateway_route_table_announcement_id: str,
                 type: str):
        """
        :param str destination_cidr_block: The CIDR used for route destination matches.
        :param str prefix_list_id: The ID of the prefix list used for destination matches.
        :param str state: The current state of the route, can be `active`, `deleted`, `pending`, `blackhole`, `deleting`.
        :param str transit_gateway_route_table_announcement_id: The id of the transit gateway route table announcement, most of the time it is an empty string.
        :param str type: The type of the route, can be `propagated` or `static`.
        """
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "prefix_list_id", prefix_list_id)
        pulumi.set(__self__, "state", state)
        pulumi.set(__self__, "transit_gateway_route_table_announcement_id", transit_gateway_route_table_announcement_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> str:
        """
        The CIDR used for route destination matches.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="prefixListId")
    def prefix_list_id(self) -> str:
        """
        The ID of the prefix list used for destination matches.
        """
        return pulumi.get(self, "prefix_list_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the route, can be `active`, `deleted`, `pending`, `blackhole`, `deleting`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="transitGatewayRouteTableAnnouncementId")
    def transit_gateway_route_table_announcement_id(self) -> str:
        """
        The id of the transit gateway route table announcement, most of the time it is an empty string.
        """
        return pulumi.get(self, "transit_gateway_route_table_announcement_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the route, can be `propagated` or `static`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetTransitGatewayFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the field to filter by, as defined by the [underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGateways.html).
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the field to filter by, as defined by the [underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGateways.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetVpcAttachmentFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter.
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetVpcAttachmentsFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter check available value on [official documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayVpcAttachments.html)
        :param Sequence[str] values: List of one or more values for the filter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter check available value on [official documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayVpcAttachments.html)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        List of one or more values for the filter.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetVpnAttachmentFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Name of the filter field. Valid values can be found in the [EC2 DescribeTransitGatewayAttachments API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html).
        :param Sequence[str] values: Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the filter field. Valid values can be found in the [EC2 DescribeTransitGatewayAttachments API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        return pulumi.get(self, "values")


