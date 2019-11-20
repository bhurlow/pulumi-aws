// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a resource to create a VPC routing table.
    /// 
    /// &gt; **NOTE on Route Tables and Routes:** This provider currently
    /// provides both a standalone Route resource and a Route Table resource with routes
    /// defined in-line. At this time you cannot use a Route Table with in-line routes
    /// in conjunction with any Route resources. Doing so will cause
    /// a conflict of rule settings and will overwrite rules.
    /// 
    /// &gt; **NOTE on `gateway_id` and `nat_gateway_id`:** The AWS API is very forgiving with these two
    /// attributes and the `aws.ec2.RouteTable` resource can be created with a NAT ID specified as a Gateway ID attribute.
    /// This _will_ lead to a permanent diff between your configuration and statefile, as the API returns the correct
    /// parameters in the returned route table. If you're experiencing constant diffs in your `aws.ec2.RouteTable` resources,
    /// the first thing to check is whether or not you're specifying a NAT ID instead of a Gateway ID, or vice-versa.
    /// 
    /// &gt; **NOTE on `propagating_vgws` and the `aws.ec2.VpnGatewayRoutePropagation` resource:**
    /// If the `propagating_vgws` argument is present, it's not supported to _also_
    /// define route propagations using `aws.ec2.VpnGatewayRoutePropagation`, since
    /// this resource will delete any propagating gateways not explicitly listed in
    /// `propagating_vgws`. Omit this argument when defining route propagation using
    /// the separate resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route_table.html.markdown.
    /// </summary>
    public partial class RouteTable : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the AWS account that owns the route table.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// A list of virtual gateways for propagation.
        /// </summary>
        [Output("propagatingVgws")]
        public Output<ImmutableArray<string>> PropagatingVgws { get; private set; } = null!;

        /// <summary>
        /// A list of route objects. Their keys are documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<Outputs.RouteTableRoutes>> Routes { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteTable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteTable(string name, RouteTableArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/routeTable:RouteTable", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private RouteTable(string name, Input<string> id, RouteTableState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/routeTable:RouteTable", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RouteTable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteTable Get(string name, Input<string> id, RouteTableState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteTable(name, id, state, options);
        }
    }

    public sealed class RouteTableArgs : Pulumi.ResourceArgs
    {
        [Input("propagatingVgws")]
        private InputList<string>? _propagatingVgws;

        /// <summary>
        /// A list of virtual gateways for propagation.
        /// </summary>
        public InputList<string> PropagatingVgws
        {
            get => _propagatingVgws ?? (_propagatingVgws = new InputList<string>());
            set => _propagatingVgws = value;
        }

        [Input("routes")]
        private InputList<Inputs.RouteTableRoutesArgs>? _routes;

        /// <summary>
        /// A list of route objects. Their keys are documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.RouteTableRoutesArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.RouteTableRoutesArgs>());
            set => _routes = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public RouteTableArgs()
        {
        }
    }

    public sealed class RouteTableState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the AWS account that owns the route table.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("propagatingVgws")]
        private InputList<string>? _propagatingVgws;

        /// <summary>
        /// A list of virtual gateways for propagation.
        /// </summary>
        public InputList<string> PropagatingVgws
        {
            get => _propagatingVgws ?? (_propagatingVgws = new InputList<string>());
            set => _propagatingVgws = value;
        }

        [Input("routes")]
        private InputList<Inputs.RouteTableRoutesGetArgs>? _routes;

        /// <summary>
        /// A list of route objects. Their keys are documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.RouteTableRoutesGetArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.RouteTableRoutesGetArgs>());
            set => _routes = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public RouteTableState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RouteTableRoutesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC Egress Only Internet Gateway.
        /// </summary>
        [Input("egressOnlyGatewayId")]
        public Input<string>? EgressOnlyGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC internet gateway or a virtual private gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The Ipv6 CIDR block of the route.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC NAT gateway.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 network interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// Identifier of an EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC peering connection.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public Input<string>? VpcPeeringConnectionId { get; set; }

        public RouteTableRoutesArgs()
        {
        }
    }

    public sealed class RouteTableRoutesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC Egress Only Internet Gateway.
        /// </summary>
        [Input("egressOnlyGatewayId")]
        public Input<string>? EgressOnlyGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC internet gateway or a virtual private gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The Ipv6 CIDR block of the route.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC NAT gateway.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 network interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// Identifier of an EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC peering connection.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public Input<string>? VpcPeeringConnectionId { get; set; }

        public RouteTableRoutesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RouteTableRoutes
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// Identifier of a VPC Egress Only Internet Gateway.
        /// </summary>
        public readonly string? EgressOnlyGatewayId;
        /// <summary>
        /// Identifier of a VPC internet gateway or a virtual private gateway.
        /// </summary>
        public readonly string? GatewayId;
        /// <summary>
        /// Identifier of an EC2 instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The Ipv6 CIDR block of the route.
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// Identifier of a VPC NAT gateway.
        /// </summary>
        public readonly string? NatGatewayId;
        /// <summary>
        /// Identifier of an EC2 network interface.
        /// </summary>
        public readonly string? NetworkInterfaceId;
        /// <summary>
        /// Identifier of an EC2 Transit Gateway.
        /// </summary>
        public readonly string? TransitGatewayId;
        /// <summary>
        /// Identifier of a VPC peering connection.
        /// </summary>
        public readonly string? VpcPeeringConnectionId;

        [OutputConstructor]
        private RouteTableRoutes(
            string? cidrBlock,
            string? egressOnlyGatewayId,
            string? gatewayId,
            string? instanceId,
            string? ipv6CidrBlock,
            string? natGatewayId,
            string? networkInterfaceId,
            string? transitGatewayId,
            string? vpcPeeringConnectionId)
        {
            CidrBlock = cidrBlock;
            EgressOnlyGatewayId = egressOnlyGatewayId;
            GatewayId = gatewayId;
            InstanceId = instanceId;
            Ipv6CidrBlock = ipv6CidrBlock;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            TransitGatewayId = transitGatewayId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
        }
    }
    }
}
