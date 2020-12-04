// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./ami";
export * from "./amiCopy";
export * from "./amiFromInstance";
export * from "./amiLaunchPermission";
export * from "./availabilityZoneGroup";
export * from "./capacityReservation";
export * from "./customerGateway";
export * from "./defaultNetworkAcl";
export * from "./defaultRouteTable";
export * from "./defaultSecurityGroup";
export * from "./defaultSubnet";
export * from "./defaultVpc";
export * from "./defaultVpcDhcpOptions";
export * from "./egressOnlyInternetGateway";
export * from "./eip";
export * from "./eipAssociation";
export * from "./fleet";
export * from "./flowLog";
export * from "./getCoipPool";
export * from "./getCoipPools";
export * from "./getCustomerGateway";
export * from "./getInstance";
export * from "./getInstanceType";
export * from "./getInstanceTypeOffering";
export * from "./getInstanceTypeOfferings";
export * from "./getInstances";
export * from "./getInternetGateway";
export * from "./getLaunchConfiguration";
export * from "./getLaunchTemplate";
export * from "./getLocalGateway";
export * from "./getLocalGatewayRouteTable";
export * from "./getLocalGatewayRouteTables";
export * from "./getLocalGatewayVirtualInterface";
export * from "./getLocalGatewayVirtualInterfaceGroup";
export * from "./getLocalGatewayVirtualInterfaceGroups";
export * from "./getLocalGateways";
export * from "./getNatGateway";
export * from "./getNetworkAcls";
export * from "./getNetworkInterface";
export * from "./getNetworkInterfaces";
export * from "./getRoute";
export * from "./getRouteTable";
export * from "./getRouteTables";
export * from "./getSecurityGroup";
export * from "./getSecurityGroups";
export * from "./getSpotPrice";
export * from "./getSubnet";
export * from "./getSubnetIds";
export * from "./getVpc";
export * from "./getVpcDhcpOptions";
export * from "./getVpcEndpoint";
export * from "./getVpcEndpointService";
export * from "./getVpcPeeringConnection";
export * from "./getVpcPeeringConnections";
export * from "./getVpcs";
export * from "./getVpnGateway";
export * from "./instance";
export * from "./instancePlatform";
export * from "./instanceType";
export * from "./internetGateway";
export * from "./keyPair";
export * from "./launchConfiguration";
export * from "./launchTemplate";
export * from "./localGatewayRoute";
export * from "./localGatewayRouteTableVpcAssociation";
export * from "./mainRouteTableAssociation";
export * from "./natGateway";
export * from "./networkAcl";
export * from "./networkAclRule";
export * from "./networkInterface";
export * from "./networkInterfaceAttachment";
export * from "./networkInterfaceSecurityGroupAttachment";
export * from "./peeringConnectionOptions";
export * from "./placementGroup";
export * from "./placementStrategy";
export * from "./protocolType";
export * from "./proxyProtocolPolicy";
export * from "./route";
export * from "./routeTable";
export * from "./routeTableAssociation";
export * from "./securityGroup";
export * from "./securityGroupRule";
export * from "./snapshotCreateVolumePermission";
export * from "./spotDatafeedSubscription";
export * from "./spotFleetRequest";
export * from "./spotInstanceRequest";
export * from "./subnet";
export * from "./tag";
export * from "./tenancy";
export * from "./trafficMirrorFilter";
export * from "./trafficMirrorFilterRule";
export * from "./trafficMirrorSession";
export * from "./trafficMirrorTarget";
export * from "./transitGatewayPeeringAttachmentAccepter";
export * from "./volumeAttachment";
export * from "./vpc";
export * from "./vpcDhcpOptions";
export * from "./vpcDhcpOptionsAssociation";
export * from "./vpcEndpoint";
export * from "./vpcEndpointConnectionNotification";
export * from "./vpcEndpointRouteTableAssociation";
export * from "./vpcEndpointService";
export * from "./vpcEndpointServiceAllowedPrinciple";
export * from "./vpcEndpointSubnetAssociation";
export * from "./vpcIpv4CidrBlockAssociation";
export * from "./vpcPeeringConnection";
export * from "./vpcPeeringConnectionAccepter";
export * from "./vpnConnection";
export * from "./vpnConnectionRoute";
export * from "./vpnGateway";
export * from "./vpnGatewayAttachment";
export * from "./vpnGatewayRoutePropagation";

// Export enums:
export * from "../types/enums/ec2";

// Import resources to register:
import { Ami } from "./ami";
import { AmiCopy } from "./amiCopy";
import { AmiFromInstance } from "./amiFromInstance";
import { AmiLaunchPermission } from "./amiLaunchPermission";
import { AvailabilityZoneGroup } from "./availabilityZoneGroup";
import { CapacityReservation } from "./capacityReservation";
import { CustomerGateway } from "./customerGateway";
import { DefaultNetworkAcl } from "./defaultNetworkAcl";
import { DefaultRouteTable } from "./defaultRouteTable";
import { DefaultSecurityGroup } from "./defaultSecurityGroup";
import { DefaultSubnet } from "./defaultSubnet";
import { DefaultVpc } from "./defaultVpc";
import { DefaultVpcDhcpOptions } from "./defaultVpcDhcpOptions";
import { EgressOnlyInternetGateway } from "./egressOnlyInternetGateway";
import { Eip } from "./eip";
import { EipAssociation } from "./eipAssociation";
import { Fleet } from "./fleet";
import { FlowLog } from "./flowLog";
import { Instance } from "./instance";
import { InternetGateway } from "./internetGateway";
import { KeyPair } from "./keyPair";
import { LaunchConfiguration } from "./launchConfiguration";
import { LaunchTemplate } from "./launchTemplate";
import { LocalGatewayRoute } from "./localGatewayRoute";
import { LocalGatewayRouteTableVpcAssociation } from "./localGatewayRouteTableVpcAssociation";
import { MainRouteTableAssociation } from "./mainRouteTableAssociation";
import { NatGateway } from "./natGateway";
import { NetworkAcl } from "./networkAcl";
import { NetworkAclRule } from "./networkAclRule";
import { NetworkInterface } from "./networkInterface";
import { NetworkInterfaceAttachment } from "./networkInterfaceAttachment";
import { NetworkInterfaceSecurityGroupAttachment } from "./networkInterfaceSecurityGroupAttachment";
import { PeeringConnectionOptions } from "./peeringConnectionOptions";
import { PlacementGroup } from "./placementGroup";
import { ProxyProtocolPolicy } from "./proxyProtocolPolicy";
import { Route } from "./route";
import { RouteTable } from "./routeTable";
import { RouteTableAssociation } from "./routeTableAssociation";
import { SecurityGroup } from "./securityGroup";
import { SecurityGroupRule } from "./securityGroupRule";
import { SnapshotCreateVolumePermission } from "./snapshotCreateVolumePermission";
import { SpotDatafeedSubscription } from "./spotDatafeedSubscription";
import { SpotFleetRequest } from "./spotFleetRequest";
import { SpotInstanceRequest } from "./spotInstanceRequest";
import { Subnet } from "./subnet";
import { Tag } from "./tag";
import { TrafficMirrorFilter } from "./trafficMirrorFilter";
import { TrafficMirrorFilterRule } from "./trafficMirrorFilterRule";
import { TrafficMirrorSession } from "./trafficMirrorSession";
import { TrafficMirrorTarget } from "./trafficMirrorTarget";
import { TransitGatewayPeeringAttachmentAccepter } from "./transitGatewayPeeringAttachmentAccepter";
import { VolumeAttachment } from "./volumeAttachment";
import { Vpc } from "./vpc";
import { VpcDhcpOptions } from "./vpcDhcpOptions";
import { VpcDhcpOptionsAssociation } from "./vpcDhcpOptionsAssociation";
import { VpcEndpoint } from "./vpcEndpoint";
import { VpcEndpointConnectionNotification } from "./vpcEndpointConnectionNotification";
import { VpcEndpointRouteTableAssociation } from "./vpcEndpointRouteTableAssociation";
import { VpcEndpointService } from "./vpcEndpointService";
import { VpcEndpointServiceAllowedPrinciple } from "./vpcEndpointServiceAllowedPrinciple";
import { VpcEndpointSubnetAssociation } from "./vpcEndpointSubnetAssociation";
import { VpcIpv4CidrBlockAssociation } from "./vpcIpv4CidrBlockAssociation";
import { VpcPeeringConnection } from "./vpcPeeringConnection";
import { VpcPeeringConnectionAccepter } from "./vpcPeeringConnectionAccepter";
import { VpnConnection } from "./vpnConnection";
import { VpnConnectionRoute } from "./vpnConnectionRoute";
import { VpnGateway } from "./vpnGateway";
import { VpnGatewayAttachment } from "./vpnGatewayAttachment";
import { VpnGatewayRoutePropagation } from "./vpnGatewayRoutePropagation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:ec2/ami:Ami":
                return new Ami(name, <any>undefined, { urn })
            case "aws:ec2/amiCopy:AmiCopy":
                return new AmiCopy(name, <any>undefined, { urn })
            case "aws:ec2/amiFromInstance:AmiFromInstance":
                return new AmiFromInstance(name, <any>undefined, { urn })
            case "aws:ec2/amiLaunchPermission:AmiLaunchPermission":
                return new AmiLaunchPermission(name, <any>undefined, { urn })
            case "aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup":
                return new AvailabilityZoneGroup(name, <any>undefined, { urn })
            case "aws:ec2/capacityReservation:CapacityReservation":
                return new CapacityReservation(name, <any>undefined, { urn })
            case "aws:ec2/customerGateway:CustomerGateway":
                return new CustomerGateway(name, <any>undefined, { urn })
            case "aws:ec2/defaultNetworkAcl:DefaultNetworkAcl":
                return new DefaultNetworkAcl(name, <any>undefined, { urn })
            case "aws:ec2/defaultRouteTable:DefaultRouteTable":
                return new DefaultRouteTable(name, <any>undefined, { urn })
            case "aws:ec2/defaultSecurityGroup:DefaultSecurityGroup":
                return new DefaultSecurityGroup(name, <any>undefined, { urn })
            case "aws:ec2/defaultSubnet:DefaultSubnet":
                return new DefaultSubnet(name, <any>undefined, { urn })
            case "aws:ec2/defaultVpc:DefaultVpc":
                return new DefaultVpc(name, <any>undefined, { urn })
            case "aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions":
                return new DefaultVpcDhcpOptions(name, <any>undefined, { urn })
            case "aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway":
                return new EgressOnlyInternetGateway(name, <any>undefined, { urn })
            case "aws:ec2/eip:Eip":
                return new Eip(name, <any>undefined, { urn })
            case "aws:ec2/eipAssociation:EipAssociation":
                return new EipAssociation(name, <any>undefined, { urn })
            case "aws:ec2/fleet:Fleet":
                return new Fleet(name, <any>undefined, { urn })
            case "aws:ec2/flowLog:FlowLog":
                return new FlowLog(name, <any>undefined, { urn })
            case "aws:ec2/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "aws:ec2/internetGateway:InternetGateway":
                return new InternetGateway(name, <any>undefined, { urn })
            case "aws:ec2/keyPair:KeyPair":
                return new KeyPair(name, <any>undefined, { urn })
            case "aws:ec2/launchConfiguration:LaunchConfiguration":
                return new LaunchConfiguration(name, <any>undefined, { urn })
            case "aws:ec2/launchTemplate:LaunchTemplate":
                return new LaunchTemplate(name, <any>undefined, { urn })
            case "aws:ec2/localGatewayRoute:LocalGatewayRoute":
                return new LocalGatewayRoute(name, <any>undefined, { urn })
            case "aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation":
                return new LocalGatewayRouteTableVpcAssociation(name, <any>undefined, { urn })
            case "aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation":
                return new MainRouteTableAssociation(name, <any>undefined, { urn })
            case "aws:ec2/natGateway:NatGateway":
                return new NatGateway(name, <any>undefined, { urn })
            case "aws:ec2/networkAcl:NetworkAcl":
                return new NetworkAcl(name, <any>undefined, { urn })
            case "aws:ec2/networkAclRule:NetworkAclRule":
                return new NetworkAclRule(name, <any>undefined, { urn })
            case "aws:ec2/networkInterface:NetworkInterface":
                return new NetworkInterface(name, <any>undefined, { urn })
            case "aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment":
                return new NetworkInterfaceAttachment(name, <any>undefined, { urn })
            case "aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment":
                return new NetworkInterfaceSecurityGroupAttachment(name, <any>undefined, { urn })
            case "aws:ec2/peeringConnectionOptions:PeeringConnectionOptions":
                return new PeeringConnectionOptions(name, <any>undefined, { urn })
            case "aws:ec2/placementGroup:PlacementGroup":
                return new PlacementGroup(name, <any>undefined, { urn })
            case "aws:ec2/proxyProtocolPolicy:ProxyProtocolPolicy":
                return new ProxyProtocolPolicy(name, <any>undefined, { urn })
            case "aws:ec2/route:Route":
                return new Route(name, <any>undefined, { urn })
            case "aws:ec2/routeTable:RouteTable":
                return new RouteTable(name, <any>undefined, { urn })
            case "aws:ec2/routeTableAssociation:RouteTableAssociation":
                return new RouteTableAssociation(name, <any>undefined, { urn })
            case "aws:ec2/securityGroup:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "aws:ec2/securityGroupRule:SecurityGroupRule":
                return new SecurityGroupRule(name, <any>undefined, { urn })
            case "aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission":
                return new SnapshotCreateVolumePermission(name, <any>undefined, { urn })
            case "aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription":
                return new SpotDatafeedSubscription(name, <any>undefined, { urn })
            case "aws:ec2/spotFleetRequest:SpotFleetRequest":
                return new SpotFleetRequest(name, <any>undefined, { urn })
            case "aws:ec2/spotInstanceRequest:SpotInstanceRequest":
                return new SpotInstanceRequest(name, <any>undefined, { urn })
            case "aws:ec2/subnet:Subnet":
                return new Subnet(name, <any>undefined, { urn })
            case "aws:ec2/tag:Tag":
                return new Tag(name, <any>undefined, { urn })
            case "aws:ec2/trafficMirrorFilter:TrafficMirrorFilter":
                return new TrafficMirrorFilter(name, <any>undefined, { urn })
            case "aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule":
                return new TrafficMirrorFilterRule(name, <any>undefined, { urn })
            case "aws:ec2/trafficMirrorSession:TrafficMirrorSession":
                return new TrafficMirrorSession(name, <any>undefined, { urn })
            case "aws:ec2/trafficMirrorTarget:TrafficMirrorTarget":
                return new TrafficMirrorTarget(name, <any>undefined, { urn })
            case "aws:ec2/transitGatewayPeeringAttachmentAccepter:TransitGatewayPeeringAttachmentAccepter":
                return new TransitGatewayPeeringAttachmentAccepter(name, <any>undefined, { urn })
            case "aws:ec2/volumeAttachment:VolumeAttachment":
                return new VolumeAttachment(name, <any>undefined, { urn })
            case "aws:ec2/vpc:Vpc":
                return new Vpc(name, <any>undefined, { urn })
            case "aws:ec2/vpcDhcpOptions:VpcDhcpOptions":
                return new VpcDhcpOptions(name, <any>undefined, { urn })
            case "aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation":
                return new VpcDhcpOptionsAssociation(name, <any>undefined, { urn })
            case "aws:ec2/vpcEndpoint:VpcEndpoint":
                return new VpcEndpoint(name, <any>undefined, { urn })
            case "aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification":
                return new VpcEndpointConnectionNotification(name, <any>undefined, { urn })
            case "aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation":
                return new VpcEndpointRouteTableAssociation(name, <any>undefined, { urn })
            case "aws:ec2/vpcEndpointService:VpcEndpointService":
                return new VpcEndpointService(name, <any>undefined, { urn })
            case "aws:ec2/vpcEndpointServiceAllowedPrinciple:VpcEndpointServiceAllowedPrinciple":
                return new VpcEndpointServiceAllowedPrinciple(name, <any>undefined, { urn })
            case "aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation":
                return new VpcEndpointSubnetAssociation(name, <any>undefined, { urn })
            case "aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation":
                return new VpcIpv4CidrBlockAssociation(name, <any>undefined, { urn })
            case "aws:ec2/vpcPeeringConnection:VpcPeeringConnection":
                return new VpcPeeringConnection(name, <any>undefined, { urn })
            case "aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter":
                return new VpcPeeringConnectionAccepter(name, <any>undefined, { urn })
            case "aws:ec2/vpnConnection:VpnConnection":
                return new VpnConnection(name, <any>undefined, { urn })
            case "aws:ec2/vpnConnectionRoute:VpnConnectionRoute":
                return new VpnConnectionRoute(name, <any>undefined, { urn })
            case "aws:ec2/vpnGateway:VpnGateway":
                return new VpnGateway(name, <any>undefined, { urn })
            case "aws:ec2/vpnGatewayAttachment:VpnGatewayAttachment":
                return new VpnGatewayAttachment(name, <any>undefined, { urn })
            case "aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation":
                return new VpnGatewayRoutePropagation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "ec2/ami", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/amiCopy", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/amiFromInstance", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/amiLaunchPermission", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/availabilityZoneGroup", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/capacityReservation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/customerGateway", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/defaultNetworkAcl", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/defaultRouteTable", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/defaultSecurityGroup", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/defaultSubnet", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/defaultVpc", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/defaultVpcDhcpOptions", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/egressOnlyInternetGateway", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/eip", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/eipAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/fleet", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/flowLog", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/instance", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/internetGateway", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/keyPair", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/launchConfiguration", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/launchTemplate", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/localGatewayRoute", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/localGatewayRouteTableVpcAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/mainRouteTableAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/natGateway", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/networkAcl", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/networkAclRule", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/networkInterface", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/networkInterfaceAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/networkInterfaceSecurityGroupAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/peeringConnectionOptions", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/placementGroup", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/proxyProtocolPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/route", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/routeTable", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/routeTableAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/securityGroup", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/securityGroupRule", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/snapshotCreateVolumePermission", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/spotDatafeedSubscription", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/spotFleetRequest", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/spotInstanceRequest", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/subnet", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/tag", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/trafficMirrorFilter", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/trafficMirrorFilterRule", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/trafficMirrorSession", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/trafficMirrorTarget", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/transitGatewayPeeringAttachmentAccepter", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/volumeAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpc", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcDhcpOptions", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcDhcpOptionsAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcEndpoint", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcEndpointConnectionNotification", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcEndpointRouteTableAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcEndpointService", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcEndpointServiceAllowedPrinciple", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcEndpointSubnetAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcIpv4CidrBlockAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcPeeringConnection", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpcPeeringConnectionAccepter", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpnConnection", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpnConnectionRoute", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpnGateway", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpnGatewayAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "ec2/vpnGatewayRoutePropagation", _module)
