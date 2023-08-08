// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.NetworkManager
{
    /// <summary>
    /// Creates a peering connection between an AWS Cloud WAN core network and an AWS Transit Gateway.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.NetworkManager.TransitGatewayPeering("example", new()
    ///     {
    ///         CoreNetworkId = awscc_networkmanager_core_network.Example.Id,
    ///         TransitGatewayArn = aws_ec2_transit_gateway.Example.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_networkmanager_transit_gateway_peering.example
    /// 
    ///  id = "peering-444555aaabbb11223" } Using `pulumi import`, import `aws_networkmanager_transit_gateway_peering` using the peering ID. For exampleconsole % pulumi import aws_networkmanager_transit_gateway_peering.example peering-444555aaabbb11223
    /// </summary>
    [AwsResourceType("aws:networkmanager/transitGatewayPeering:TransitGatewayPeering")]
    public partial class TransitGatewayPeering : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Peering Amazon Resource Name (ARN).
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the core network.
        /// </summary>
        [Output("coreNetworkArn")]
        public Output<string> CoreNetworkArn { get; private set; } = null!;

        /// <summary>
        /// The ID of a core network.
        /// </summary>
        [Output("coreNetworkId")]
        public Output<string> CoreNetworkId { get; private set; } = null!;

        /// <summary>
        /// The edge location for the peer.
        /// </summary>
        [Output("edgeLocation")]
        public Output<string> EdgeLocation { get; private set; } = null!;

        /// <summary>
        /// The ID of the account owner.
        /// </summary>
        [Output("ownerAccountId")]
        public Output<string> OwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// The type of peering. This will be `TRANSIT_GATEWAY`.
        /// </summary>
        [Output("peeringType")]
        public Output<string> PeeringType { get; private set; } = null!;

        /// <summary>
        /// The resource ARN of the peer.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the peering. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The ARN of the transit gateway for the peering request.
        /// </summary>
        [Output("transitGatewayArn")]
        public Output<string> TransitGatewayArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the transit gateway peering attachment.
        /// </summary>
        [Output("transitGatewayPeeringAttachmentId")]
        public Output<string> TransitGatewayPeeringAttachmentId { get; private set; } = null!;


        /// <summary>
        /// Create a TransitGatewayPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitGatewayPeering(string name, TransitGatewayPeeringArgs args, CustomResourceOptions? options = null)
            : base("aws:networkmanager/transitGatewayPeering:TransitGatewayPeering", name, args ?? new TransitGatewayPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitGatewayPeering(string name, Input<string> id, TransitGatewayPeeringState? state = null, CustomResourceOptions? options = null)
            : base("aws:networkmanager/transitGatewayPeering:TransitGatewayPeering", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransitGatewayPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitGatewayPeering Get(string name, Input<string> id, TransitGatewayPeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new TransitGatewayPeering(name, id, state, options);
        }
    }

    public sealed class TransitGatewayPeeringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a core network.
        /// </summary>
        [Input("coreNetworkId", required: true)]
        public Input<string> CoreNetworkId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the peering. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ARN of the transit gateway for the peering request.
        /// </summary>
        [Input("transitGatewayArn", required: true)]
        public Input<string> TransitGatewayArn { get; set; } = null!;

        public TransitGatewayPeeringArgs()
        {
        }
        public static new TransitGatewayPeeringArgs Empty => new TransitGatewayPeeringArgs();
    }

    public sealed class TransitGatewayPeeringState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Peering Amazon Resource Name (ARN).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN of the core network.
        /// </summary>
        [Input("coreNetworkArn")]
        public Input<string>? CoreNetworkArn { get; set; }

        /// <summary>
        /// The ID of a core network.
        /// </summary>
        [Input("coreNetworkId")]
        public Input<string>? CoreNetworkId { get; set; }

        /// <summary>
        /// The edge location for the peer.
        /// </summary>
        [Input("edgeLocation")]
        public Input<string>? EdgeLocation { get; set; }

        /// <summary>
        /// The ID of the account owner.
        /// </summary>
        [Input("ownerAccountId")]
        public Input<string>? OwnerAccountId { get; set; }

        /// <summary>
        /// The type of peering. This will be `TRANSIT_GATEWAY`.
        /// </summary>
        [Input("peeringType")]
        public Input<string>? PeeringType { get; set; }

        /// <summary>
        /// The resource ARN of the peer.
        /// </summary>
        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the peering. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// The ARN of the transit gateway for the peering request.
        /// </summary>
        [Input("transitGatewayArn")]
        public Input<string>? TransitGatewayArn { get; set; }

        /// <summary>
        /// The ID of the transit gateway peering attachment.
        /// </summary>
        [Input("transitGatewayPeeringAttachmentId")]
        public Input<string>? TransitGatewayPeeringAttachmentId { get; set; }

        public TransitGatewayPeeringState()
        {
        }
        public static new TransitGatewayPeeringState Empty => new TransitGatewayPeeringState();
    }
}
