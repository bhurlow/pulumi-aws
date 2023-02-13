// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class NetworkInsightsAnalysisExplanation
    {
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationAclRule> AclRules;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationAcl> Acls;
        public readonly string? Address;
        public readonly ImmutableArray<string> Addresses;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationAttachedTo> AttachedTos;
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly ImmutableArray<string> Cidrs;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationClassicLoadBalancerListener> ClassicLoadBalancerListeners;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationComponent> Components;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationCustomerGateway> CustomerGateways;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationDestinationVpc> DestinationVpcs;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationDestination> Destinations;
        public readonly string? Direction;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationElasticLoadBalancerListener> ElasticLoadBalancerListeners;
        public readonly string? ExplanationCode;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationIngressRouteTable> IngressRouteTables;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationInternetGateway> InternetGateways;
        public readonly string? LoadBalancerArn;
        public readonly int? LoadBalancerListenerPort;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup> LoadBalancerTargetGroup;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup> LoadBalancerTargetGroups;
        public readonly int? LoadBalancerTargetPort;
        public readonly string? MissingComponent;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationNatGateway> NatGateways;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationNetworkInterface> NetworkInterfaces;
        public readonly string? PacketField;
        public readonly int? Port;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationPortRange> PortRanges;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationPrefixList> PrefixLists;
        public readonly ImmutableArray<string> Protocols;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationRouteTableRoute> RouteTableRoutes;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationRouteTable> RouteTables;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSecurityGroup> SecurityGroup;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSecurityGroupRule> SecurityGroupRules;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSecurityGroup> SecurityGroups;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSourceVpc> SourceVpcs;
        public readonly string? State;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSubnetRouteTable> SubnetRouteTables;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSubnet> Subnets;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGatewayAttachment> TransitGatewayAttachments;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableRoute> TransitGatewayRouteTableRoutes;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTable> TransitGatewayRouteTables;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGateway> TransitGateways;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpcEndpoint> VpcEndpoints;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpcPeeringConnection> VpcPeeringConnections;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpc> Vpcs;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpnConnection> VpnConnections;
        public readonly ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpnGateway> VpnGateways;

        [OutputConstructor]
        private NetworkInsightsAnalysisExplanation(
            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationAclRule> aclRules,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationAcl> acls,

            string? address,

            ImmutableArray<string> addresses,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationAttachedTo> attachedTos,

            ImmutableArray<string> availabilityZones,

            ImmutableArray<string> cidrs,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationClassicLoadBalancerListener> classicLoadBalancerListeners,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationComponent> components,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationCustomerGateway> customerGateways,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationDestinationVpc> destinationVpcs,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationDestination> destinations,

            string? direction,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationElasticLoadBalancerListener> elasticLoadBalancerListeners,

            string? explanationCode,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationIngressRouteTable> ingressRouteTables,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationInternetGateway> internetGateways,

            string? loadBalancerArn,

            int? loadBalancerListenerPort,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup> loadBalancerTargetGroup,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup> loadBalancerTargetGroups,

            int? loadBalancerTargetPort,

            string? missingComponent,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationNatGateway> natGateways,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationNetworkInterface> networkInterfaces,

            string? packetField,

            int? port,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationPortRange> portRanges,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationPrefixList> prefixLists,

            ImmutableArray<string> protocols,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationRouteTableRoute> routeTableRoutes,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationRouteTable> routeTables,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSecurityGroup> securityGroup,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSecurityGroupRule> securityGroupRules,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSecurityGroup> securityGroups,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSourceVpc> sourceVpcs,

            string? state,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSubnetRouteTable> subnetRouteTables,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationSubnet> subnets,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGatewayAttachment> transitGatewayAttachments,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableRoute> transitGatewayRouteTableRoutes,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTable> transitGatewayRouteTables,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationTransitGateway> transitGateways,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpcEndpoint> vpcEndpoints,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpcPeeringConnection> vpcPeeringConnections,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpc> vpcs,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpnConnection> vpnConnections,

            ImmutableArray<Outputs.NetworkInsightsAnalysisExplanationVpnGateway> vpnGateways)
        {
            AclRules = aclRules;
            Acls = acls;
            Address = address;
            Addresses = addresses;
            AttachedTos = attachedTos;
            AvailabilityZones = availabilityZones;
            Cidrs = cidrs;
            ClassicLoadBalancerListeners = classicLoadBalancerListeners;
            Components = components;
            CustomerGateways = customerGateways;
            DestinationVpcs = destinationVpcs;
            Destinations = destinations;
            Direction = direction;
            ElasticLoadBalancerListeners = elasticLoadBalancerListeners;
            ExplanationCode = explanationCode;
            IngressRouteTables = ingressRouteTables;
            InternetGateways = internetGateways;
            LoadBalancerArn = loadBalancerArn;
            LoadBalancerListenerPort = loadBalancerListenerPort;
            LoadBalancerTargetGroup = loadBalancerTargetGroup;
            LoadBalancerTargetGroups = loadBalancerTargetGroups;
            LoadBalancerTargetPort = loadBalancerTargetPort;
            MissingComponent = missingComponent;
            NatGateways = natGateways;
            NetworkInterfaces = networkInterfaces;
            PacketField = packetField;
            Port = port;
            PortRanges = portRanges;
            PrefixLists = prefixLists;
            Protocols = protocols;
            RouteTableRoutes = routeTableRoutes;
            RouteTables = routeTables;
            SecurityGroup = securityGroup;
            SecurityGroupRules = securityGroupRules;
            SecurityGroups = securityGroups;
            SourceVpcs = sourceVpcs;
            State = state;
            SubnetRouteTables = subnetRouteTables;
            Subnets = subnets;
            TransitGatewayAttachments = transitGatewayAttachments;
            TransitGatewayRouteTableRoutes = transitGatewayRouteTableRoutes;
            TransitGatewayRouteTables = transitGatewayRouteTables;
            TransitGateways = transitGateways;
            VpcEndpoints = vpcEndpoints;
            VpcPeeringConnections = vpcPeeringConnections;
            Vpcs = vpcs;
            VpnConnections = vpnConnections;
            VpnGateways = vpnGateways;
        }
    }
}
