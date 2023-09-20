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
    /// Manages an EC2 Transit Gateway Connect.
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
    ///     var example = new Aws.Ec2TransitGateway.VpcAttachment("example", new()
    ///     {
    ///         SubnetIds = new[]
    ///         {
    ///             aws_subnet.Example.Id,
    ///         },
    ///         TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
    ///         VpcId = aws_vpc.Example.Id,
    ///     });
    /// 
    ///     var attachment = new Aws.Ec2TransitGateway.Connect("attachment", new()
    ///     {
    ///         TransportAttachmentId = example.Id,
    ///         TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_ec2_transit_gateway_connect` using the EC2 Transit Gateway Connect identifier. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2transitgateway/connect:Connect example tgw-attach-12345678
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2transitgateway/connect:Connect")]
    public partial class Connect : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The tunnel protocol. Valid values: `gre`. Default is `gre`.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        /// </summary>
        [Output("transitGatewayDefaultRouteTableAssociation")]
        public Output<bool?> TransitGatewayDefaultRouteTableAssociation { get; private set; } = null!;

        /// <summary>
        /// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        /// </summary>
        [Output("transitGatewayDefaultRouteTablePropagation")]
        public Output<bool?> TransitGatewayDefaultRouteTablePropagation { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Output("transitGatewayId")]
        public Output<string> TransitGatewayId { get; private set; } = null!;

        /// <summary>
        /// The underlaying VPC attachment
        /// </summary>
        [Output("transportAttachmentId")]
        public Output<string> TransportAttachmentId { get; private set; } = null!;


        /// <summary>
        /// Create a Connect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connect(string name, ConnectArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/connect:Connect", name, args ?? new ConnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connect(string name, Input<string> id, ConnectState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/connect:Connect", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Connect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connect Get(string name, Input<string> id, ConnectState? state = null, CustomResourceOptions? options = null)
        {
            return new Connect(name, id, state, options);
        }
    }

    public sealed class ConnectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tunnel protocol. Valid values: `gre`. Default is `gre`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTableAssociation")]
        public Input<bool>? TransitGatewayDefaultRouteTableAssociation { get; set; }

        /// <summary>
        /// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTablePropagation")]
        public Input<bool>? TransitGatewayDefaultRouteTablePropagation { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId", required: true)]
        public Input<string> TransitGatewayId { get; set; } = null!;

        /// <summary>
        /// The underlaying VPC attachment
        /// </summary>
        [Input("transportAttachmentId", required: true)]
        public Input<string> TransportAttachmentId { get; set; } = null!;

        public ConnectArgs()
        {
        }
        public static new ConnectArgs Empty => new ConnectArgs();
    }

    public sealed class ConnectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tunnel protocol. Valid values: `gre`. Default is `gre`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTableAssociation")]
        public Input<bool>? TransitGatewayDefaultRouteTableAssociation { get; set; }

        /// <summary>
        /// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTablePropagation")]
        public Input<bool>? TransitGatewayDefaultRouteTablePropagation { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// The underlaying VPC attachment
        /// </summary>
        [Input("transportAttachmentId")]
        public Input<string>? TransportAttachmentId { get; set; }

        public ConnectState()
        {
        }
        public static new ConnectState Empty => new ConnectState();
    }
}
