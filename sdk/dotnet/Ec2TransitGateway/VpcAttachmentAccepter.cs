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
    /// Manages the accepter's side of an EC2 Transit Gateway VPC Attachment.
    /// 
    /// When a cross-account (requester's AWS account differs from the accepter's AWS account) EC2 Transit Gateway VPC Attachment
    /// is created, an EC2 Transit Gateway VPC Attachment resource is automatically created in the accepter's account.
    /// The requester can use the `aws.ec2transitgateway.VpcAttachment` resource to manage its side of the connection
    /// and the accepter can use the `aws.ec2transitgateway.VpcAttachmentAccepter` resource to "adopt" its side of the
    /// connection into management.
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
    ///     var example = new Aws.Ec2TransitGateway.VpcAttachmentAccepter("example", new()
    ///     {
    ///         TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
    ///         Tags = 
    ///         {
    ///             { "Name", "Example cross-account attachment" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_ec2_transit_gateway_vpc_attachment_accepter` using the EC2 Transit Gateway Attachment identifier. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2transitgateway/vpcAttachmentAccepter:VpcAttachmentAccepter example tgw-attach-12345678
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2transitgateway/vpcAttachmentAccepter:VpcAttachmentAccepter")]
    public partial class VpcAttachmentAccepter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether Appliance Mode support is enabled. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("applianceModeSupport")]
        public Output<string> ApplianceModeSupport { get; private set; } = null!;

        /// <summary>
        /// Whether DNS support is enabled. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("dnsSupport")]
        public Output<string> DnsSupport { get; private set; } = null!;

        /// <summary>
        /// Whether IPv6 support is enabled. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("ipv6Support")]
        public Output<string> Ipv6Support { get; private set; } = null!;

        /// <summary>
        /// Identifiers of EC2 Subnets.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway VPC Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The ID of the EC2 Transit Gateway Attachment to manage.
        /// </summary>
        [Output("transitGatewayAttachmentId")]
        public Output<string> TransitGatewayAttachmentId { get; private set; } = null!;

        /// <summary>
        /// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: `true`.
        /// </summary>
        [Output("transitGatewayDefaultRouteTableAssociation")]
        public Output<bool?> TransitGatewayDefaultRouteTableAssociation { get; private set; } = null!;

        /// <summary>
        /// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: `true`.
        /// </summary>
        [Output("transitGatewayDefaultRouteTablePropagation")]
        public Output<bool?> TransitGatewayDefaultRouteTablePropagation { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Output("transitGatewayId")]
        public Output<string> TransitGatewayId { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Identifier of the AWS account that owns the EC2 VPC.
        /// </summary>
        [Output("vpcOwnerId")]
        public Output<string> VpcOwnerId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcAttachmentAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcAttachmentAccepter(string name, VpcAttachmentAccepterArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/vpcAttachmentAccepter:VpcAttachmentAccepter", name, args ?? new VpcAttachmentAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcAttachmentAccepter(string name, Input<string> id, VpcAttachmentAccepterState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/vpcAttachmentAccepter:VpcAttachmentAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcAttachmentAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcAttachmentAccepter Get(string name, Input<string> id, VpcAttachmentAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcAttachmentAccepter(name, id, state, options);
        }
    }

    public sealed class VpcAttachmentAccepterArgs : global::Pulumi.ResourceArgs
    {
        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway VPC Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the EC2 Transit Gateway Attachment to manage.
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public Input<string> TransitGatewayAttachmentId { get; set; } = null!;

        /// <summary>
        /// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTableAssociation")]
        public Input<bool>? TransitGatewayDefaultRouteTableAssociation { get; set; }

        /// <summary>
        /// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTablePropagation")]
        public Input<bool>? TransitGatewayDefaultRouteTablePropagation { get; set; }

        public VpcAttachmentAccepterArgs()
        {
        }
        public static new VpcAttachmentAccepterArgs Empty => new VpcAttachmentAccepterArgs();
    }

    public sealed class VpcAttachmentAccepterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether Appliance Mode support is enabled. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("applianceModeSupport")]
        public Input<string>? ApplianceModeSupport { get; set; }

        /// <summary>
        /// Whether DNS support is enabled. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("dnsSupport")]
        public Input<string>? DnsSupport { get; set; }

        /// <summary>
        /// Whether IPv6 support is enabled. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("ipv6Support")]
        public Input<string>? Ipv6Support { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Identifiers of EC2 Subnets.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway VPC Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// The ID of the EC2 Transit Gateway Attachment to manage.
        /// </summary>
        [Input("transitGatewayAttachmentId")]
        public Input<string>? TransitGatewayAttachmentId { get; set; }

        /// <summary>
        /// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTableAssociation")]
        public Input<bool>? TransitGatewayDefaultRouteTableAssociation { get; set; }

        /// <summary>
        /// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: `true`.
        /// </summary>
        [Input("transitGatewayDefaultRouteTablePropagation")]
        public Input<bool>? TransitGatewayDefaultRouteTablePropagation { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// Identifier of EC2 VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Identifier of the AWS account that owns the EC2 VPC.
        /// </summary>
        [Input("vpcOwnerId")]
        public Input<string>? VpcOwnerId { get; set; }

        public VpcAttachmentAccepterState()
        {
        }
        public static new VpcAttachmentAccepterState Empty => new VpcAttachmentAccepterState();
    }
}
