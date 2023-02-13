// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class NetworkInsightsAnalysisExplanationArgs : global::Pulumi.ResourceArgs
    {
        [Input("aclRules")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationAclRuleArgs>? _aclRules;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationAclRuleArgs> AclRules
        {
            get => _aclRules ?? (_aclRules = new InputList<Inputs.NetworkInsightsAnalysisExplanationAclRuleArgs>());
            set => _aclRules = value;
        }

        [Input("acls")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationAclArgs>? _acls;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationAclArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.NetworkInsightsAnalysisExplanationAclArgs>());
            set => _acls = value;
        }

        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("addresses")]
        private InputList<string>? _addresses;
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        [Input("attachedTos")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationAttachedToArgs>? _attachedTos;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationAttachedToArgs> AttachedTos
        {
            get => _attachedTos ?? (_attachedTos = new InputList<Inputs.NetworkInsightsAnalysisExplanationAttachedToArgs>());
            set => _attachedTos = value;
        }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("cidrs")]
        private InputList<string>? _cidrs;
        public InputList<string> Cidrs
        {
            get => _cidrs ?? (_cidrs = new InputList<string>());
            set => _cidrs = value;
        }

        [Input("classicLoadBalancerListeners")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationClassicLoadBalancerListenerArgs>? _classicLoadBalancerListeners;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationClassicLoadBalancerListenerArgs> ClassicLoadBalancerListeners
        {
            get => _classicLoadBalancerListeners ?? (_classicLoadBalancerListeners = new InputList<Inputs.NetworkInsightsAnalysisExplanationClassicLoadBalancerListenerArgs>());
            set => _classicLoadBalancerListeners = value;
        }

        [Input("components")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationComponentArgs>? _components;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationComponentArgs> Components
        {
            get => _components ?? (_components = new InputList<Inputs.NetworkInsightsAnalysisExplanationComponentArgs>());
            set => _components = value;
        }

        [Input("customerGateways")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationCustomerGatewayArgs>? _customerGateways;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationCustomerGatewayArgs> CustomerGateways
        {
            get => _customerGateways ?? (_customerGateways = new InputList<Inputs.NetworkInsightsAnalysisExplanationCustomerGatewayArgs>());
            set => _customerGateways = value;
        }

        [Input("destinationVpcs")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationDestinationVpcArgs>? _destinationVpcs;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationDestinationVpcArgs> DestinationVpcs
        {
            get => _destinationVpcs ?? (_destinationVpcs = new InputList<Inputs.NetworkInsightsAnalysisExplanationDestinationVpcArgs>());
            set => _destinationVpcs = value;
        }

        [Input("destinations")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationDestinationArgs>? _destinations;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.NetworkInsightsAnalysisExplanationDestinationArgs>());
            set => _destinations = value;
        }

        [Input("direction")]
        public Input<string>? Direction { get; set; }

        [Input("elasticLoadBalancerListeners")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationElasticLoadBalancerListenerArgs>? _elasticLoadBalancerListeners;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationElasticLoadBalancerListenerArgs> ElasticLoadBalancerListeners
        {
            get => _elasticLoadBalancerListeners ?? (_elasticLoadBalancerListeners = new InputList<Inputs.NetworkInsightsAnalysisExplanationElasticLoadBalancerListenerArgs>());
            set => _elasticLoadBalancerListeners = value;
        }

        [Input("explanationCode")]
        public Input<string>? ExplanationCode { get; set; }

        [Input("ingressRouteTables")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationIngressRouteTableArgs>? _ingressRouteTables;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationIngressRouteTableArgs> IngressRouteTables
        {
            get => _ingressRouteTables ?? (_ingressRouteTables = new InputList<Inputs.NetworkInsightsAnalysisExplanationIngressRouteTableArgs>());
            set => _ingressRouteTables = value;
        }

        [Input("internetGateways")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationInternetGatewayArgs>? _internetGateways;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationInternetGatewayArgs> InternetGateways
        {
            get => _internetGateways ?? (_internetGateways = new InputList<Inputs.NetworkInsightsAnalysisExplanationInternetGatewayArgs>());
            set => _internetGateways = value;
        }

        [Input("loadBalancerArn")]
        public Input<string>? LoadBalancerArn { get; set; }

        [Input("loadBalancerListenerPort")]
        public Input<int>? LoadBalancerListenerPort { get; set; }

        [Input("loadBalancerTargetGroup")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroupArgs>? _loadBalancerTargetGroup;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroupArgs> LoadBalancerTargetGroup
        {
            get => _loadBalancerTargetGroup ?? (_loadBalancerTargetGroup = new InputList<Inputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroupArgs>());
            set => _loadBalancerTargetGroup = value;
        }

        [Input("loadBalancerTargetGroups")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroupArgs>? _loadBalancerTargetGroups;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroupArgs> LoadBalancerTargetGroups
        {
            get => _loadBalancerTargetGroups ?? (_loadBalancerTargetGroups = new InputList<Inputs.NetworkInsightsAnalysisExplanationLoadBalancerTargetGroupArgs>());
            set => _loadBalancerTargetGroups = value;
        }

        [Input("loadBalancerTargetPort")]
        public Input<int>? LoadBalancerTargetPort { get; set; }

        [Input("missingComponent")]
        public Input<string>? MissingComponent { get; set; }

        [Input("natGateways")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationNatGatewayArgs>? _natGateways;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationNatGatewayArgs> NatGateways
        {
            get => _natGateways ?? (_natGateways = new InputList<Inputs.NetworkInsightsAnalysisExplanationNatGatewayArgs>());
            set => _natGateways = value;
        }

        [Input("networkInterfaces")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationNetworkInterfaceArgs>? _networkInterfaces;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.NetworkInsightsAnalysisExplanationNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("packetField")]
        public Input<string>? PacketField { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationPortRangeArgs>? _portRanges;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationPortRangeArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.NetworkInsightsAnalysisExplanationPortRangeArgs>());
            set => _portRanges = value;
        }

        [Input("prefixLists")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationPrefixListArgs>? _prefixLists;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationPrefixListArgs> PrefixLists
        {
            get => _prefixLists ?? (_prefixLists = new InputList<Inputs.NetworkInsightsAnalysisExplanationPrefixListArgs>());
            set => _prefixLists = value;
        }

        [Input("protocols")]
        private InputList<string>? _protocols;
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        [Input("routeTableRoutes")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationRouteTableRouteArgs>? _routeTableRoutes;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationRouteTableRouteArgs> RouteTableRoutes
        {
            get => _routeTableRoutes ?? (_routeTableRoutes = new InputList<Inputs.NetworkInsightsAnalysisExplanationRouteTableRouteArgs>());
            set => _routeTableRoutes = value;
        }

        [Input("routeTables")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationRouteTableArgs>? _routeTables;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationRouteTableArgs> RouteTables
        {
            get => _routeTables ?? (_routeTables = new InputList<Inputs.NetworkInsightsAnalysisExplanationRouteTableArgs>());
            set => _routeTables = value;
        }

        [Input("securityGroup")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupArgs>? _securityGroup;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupArgs> SecurityGroup
        {
            get => _securityGroup ?? (_securityGroup = new InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupArgs>());
            set => _securityGroup = value;
        }

        [Input("securityGroupRules")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupRuleArgs>? _securityGroupRules;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupRuleArgs> SecurityGroupRules
        {
            get => _securityGroupRules ?? (_securityGroupRules = new InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupRuleArgs>());
            set => _securityGroupRules = value;
        }

        [Input("securityGroups")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupArgs>? _securityGroups;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupArgs> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<Inputs.NetworkInsightsAnalysisExplanationSecurityGroupArgs>());
            set => _securityGroups = value;
        }

        [Input("sourceVpcs")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationSourceVpcArgs>? _sourceVpcs;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationSourceVpcArgs> SourceVpcs
        {
            get => _sourceVpcs ?? (_sourceVpcs = new InputList<Inputs.NetworkInsightsAnalysisExplanationSourceVpcArgs>());
            set => _sourceVpcs = value;
        }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("subnetRouteTables")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationSubnetRouteTableArgs>? _subnetRouteTables;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationSubnetRouteTableArgs> SubnetRouteTables
        {
            get => _subnetRouteTables ?? (_subnetRouteTables = new InputList<Inputs.NetworkInsightsAnalysisExplanationSubnetRouteTableArgs>());
            set => _subnetRouteTables = value;
        }

        [Input("subnets")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationSubnetArgs>? _subnets;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationSubnetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.NetworkInsightsAnalysisExplanationSubnetArgs>());
            set => _subnets = value;
        }

        [Input("transitGatewayAttachments")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayAttachmentArgs>? _transitGatewayAttachments;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayAttachmentArgs> TransitGatewayAttachments
        {
            get => _transitGatewayAttachments ?? (_transitGatewayAttachments = new InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayAttachmentArgs>());
            set => _transitGatewayAttachments = value;
        }

        [Input("transitGatewayRouteTableRoutes")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableRouteArgs>? _transitGatewayRouteTableRoutes;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableRouteArgs> TransitGatewayRouteTableRoutes
        {
            get => _transitGatewayRouteTableRoutes ?? (_transitGatewayRouteTableRoutes = new InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableRouteArgs>());
            set => _transitGatewayRouteTableRoutes = value;
        }

        [Input("transitGatewayRouteTables")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableArgs>? _transitGatewayRouteTables;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableArgs> TransitGatewayRouteTables
        {
            get => _transitGatewayRouteTables ?? (_transitGatewayRouteTables = new InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayRouteTableArgs>());
            set => _transitGatewayRouteTables = value;
        }

        [Input("transitGateways")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayArgs>? _transitGateways;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayArgs> TransitGateways
        {
            get => _transitGateways ?? (_transitGateways = new InputList<Inputs.NetworkInsightsAnalysisExplanationTransitGatewayArgs>());
            set => _transitGateways = value;
        }

        [Input("vpcEndpoints")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationVpcEndpointArgs>? _vpcEndpoints;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationVpcEndpointArgs> VpcEndpoints
        {
            get => _vpcEndpoints ?? (_vpcEndpoints = new InputList<Inputs.NetworkInsightsAnalysisExplanationVpcEndpointArgs>());
            set => _vpcEndpoints = value;
        }

        [Input("vpcPeeringConnections")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationVpcPeeringConnectionArgs>? _vpcPeeringConnections;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationVpcPeeringConnectionArgs> VpcPeeringConnections
        {
            get => _vpcPeeringConnections ?? (_vpcPeeringConnections = new InputList<Inputs.NetworkInsightsAnalysisExplanationVpcPeeringConnectionArgs>());
            set => _vpcPeeringConnections = value;
        }

        [Input("vpcs")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationVpcArgs>? _vpcs;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationVpcArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.NetworkInsightsAnalysisExplanationVpcArgs>());
            set => _vpcs = value;
        }

        [Input("vpnConnections")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationVpnConnectionArgs>? _vpnConnections;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationVpnConnectionArgs> VpnConnections
        {
            get => _vpnConnections ?? (_vpnConnections = new InputList<Inputs.NetworkInsightsAnalysisExplanationVpnConnectionArgs>());
            set => _vpnConnections = value;
        }

        [Input("vpnGateways")]
        private InputList<Inputs.NetworkInsightsAnalysisExplanationVpnGatewayArgs>? _vpnGateways;
        public InputList<Inputs.NetworkInsightsAnalysisExplanationVpnGatewayArgs> VpnGateways
        {
            get => _vpnGateways ?? (_vpnGateways = new InputList<Inputs.NetworkInsightsAnalysisExplanationVpnGatewayArgs>());
            set => _vpnGateways = value;
        }

        public NetworkInsightsAnalysisExplanationArgs()
        {
        }
        public static new NetworkInsightsAnalysisExplanationArgs Empty => new NetworkInsightsAnalysisExplanationArgs();
    }
}
