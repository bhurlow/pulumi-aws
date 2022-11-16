// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    /// <summary>
    /// Manages an EC2 Transit Gateway Connect Peer.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleConnect = new Aws.Ec2TransitGateway.Connect("exampleConnect", new()
    ///     {
    ///         TransportAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
    ///         TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
    ///     });
    /// 
    ///     var exampleConnectPeer = new Aws.Ec2TransitGateway.ConnectPeer("exampleConnectPeer", new()
    ///     {
    ///         PeerAddress = "10.1.2.3",
    ///         InsideCidrBlocks = new[]
    ///         {
    ///             "169.254.100.0/29",
    ///         },
    ///         TransitGatewayAttachmentId = exampleConnect.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_ec2_transit_gateway_connect_peer` can be imported by using the EC2 Transit Gateway Connect Peer identifier, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2transitgateway/connectPeer:ConnectPeer example tgw-connect-peer-12345678
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2transitgateway/connectPeer:ConnectPeer")]
    public partial class ConnectPeer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// EC2 Transit Gateway Connect Peer ARN
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        /// </summary>
        [Output("bgpAsn")]
        public Output<string> BgpAsn { get; private set; } = null!;

        /// <summary>
        /// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        /// </summary>
        [Output("insideCidrBlocks")]
        public Output<ImmutableArray<string>> InsideCidrBlocks { get; private set; } = null!;

        /// <summary>
        /// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        /// </summary>
        [Output("peerAddress")]
        public Output<string> PeerAddress { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        /// </summary>
        [Output("transitGatewayAddress")]
        public Output<string> TransitGatewayAddress { get; private set; } = null!;

        /// <summary>
        /// The Transit Gateway Connect
        /// </summary>
        [Output("transitGatewayAttachmentId")]
        public Output<string> TransitGatewayAttachmentId { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectPeer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectPeer(string name, ConnectPeerArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/connectPeer:ConnectPeer", name, args ?? new ConnectPeerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectPeer(string name, Input<string> id, ConnectPeerState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/connectPeer:ConnectPeer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectPeer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectPeer Get(string name, Input<string> id, ConnectPeerState? state = null, CustomResourceOptions? options = null)
        {
            return new ConnectPeer(name, id, state, options);
        }
    }

    public sealed class ConnectPeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        /// </summary>
        [Input("bgpAsn")]
        public Input<string>? BgpAsn { get; set; }

        [Input("insideCidrBlocks", required: true)]
        private InputList<string>? _insideCidrBlocks;

        /// <summary>
        /// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        /// </summary>
        public InputList<string> InsideCidrBlocks
        {
            get => _insideCidrBlocks ?? (_insideCidrBlocks = new InputList<string>());
            set => _insideCidrBlocks = value;
        }

        /// <summary>
        /// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        /// </summary>
        [Input("peerAddress", required: true)]
        public Input<string> PeerAddress { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        /// </summary>
        [Input("transitGatewayAddress")]
        public Input<string>? TransitGatewayAddress { get; set; }

        /// <summary>
        /// The Transit Gateway Connect
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public Input<string> TransitGatewayAttachmentId { get; set; } = null!;

        public ConnectPeerArgs()
        {
        }
        public static new ConnectPeerArgs Empty => new ConnectPeerArgs();
    }

    public sealed class ConnectPeerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// EC2 Transit Gateway Connect Peer ARN
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
        /// </summary>
        [Input("bgpAsn")]
        public Input<string>? BgpAsn { get; set; }

        [Input("insideCidrBlocks")]
        private InputList<string>? _insideCidrBlocks;

        /// <summary>
        /// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
        /// </summary>
        public InputList<string> InsideCidrBlocks
        {
            get => _insideCidrBlocks ?? (_insideCidrBlocks = new InputList<string>());
            set => _insideCidrBlocks = value;
        }

        /// <summary>
        /// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
        /// </summary>
        [Input("peerAddress")]
        public Input<string>? PeerAddress { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
        /// </summary>
        [Input("transitGatewayAddress")]
        public Input<string>? TransitGatewayAddress { get; set; }

        /// <summary>
        /// The Transit Gateway Connect
        /// </summary>
        [Input("transitGatewayAttachmentId")]
        public Input<string>? TransitGatewayAttachmentId { get; set; }

        public ConnectPeerState()
        {
        }
        public static new ConnectPeerState Empty => new ConnectPeerState();
    }
}
