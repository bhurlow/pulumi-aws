// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a resource to create a VPC NAT Gateway.
    /// 
    /// ## Example Usage
    /// ### Public NAT
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Ec2.NatGateway("example", new()
    ///     {
    ///         AllocationId = aws_eip.Example.Id,
    ///         SubnetId = aws_subnet.Example.Id,
    ///         Tags = 
    ///         {
    ///             { "Name", "gw NAT" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_internet_gateway.Example,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Private NAT
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Ec2.NatGateway("example", new()
    ///     {
    ///         ConnectivityType = "private",
    ///         SubnetId = aws_subnet.Example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NAT Gateways can be imported using the `id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/natGateway:NatGateway private_gw nat-05dba92075d71c408
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/natGateway:NatGateway")]
    public partial class NatGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        /// </summary>
        [Output("allocationId")]
        public Output<string?> AllocationId { get; private set; } = null!;

        /// <summary>
        /// The association ID of the Elastic IP address that's associated with the NAT gateway. Only available when `connectivity_type` is `public`.
        /// </summary>
        [Output("associationId")]
        public Output<string> AssociationId { get; private set; } = null!;

        /// <summary>
        /// Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        /// </summary>
        [Output("connectivityType")]
        public Output<string?> ConnectivityType { get; private set; } = null!;

        /// <summary>
        /// The ID of the network interface associated with the NAT gateway.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// The Elastic IP address associated with the NAT gateway.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// The Subnet ID of the subnet in which to place the gateway.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a NatGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NatGateway(string name, NatGatewayArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/natGateway:NatGateway", name, args ?? new NatGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NatGateway(string name, Input<string> id, NatGatewayState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/natGateway:NatGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NatGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NatGateway Get(string name, Input<string> id, NatGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new NatGateway(name, id, state, options);
        }
    }

    public sealed class NatGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        /// <summary>
        /// Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        /// </summary>
        [Input("connectivityType")]
        public Input<string>? ConnectivityType { get; set; }

        /// <summary>
        /// The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// The Subnet ID of the subnet in which to place the gateway.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public NatGatewayArgs()
        {
        }
        public static new NatGatewayArgs Empty => new NatGatewayArgs();
    }

    public sealed class NatGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        /// <summary>
        /// The association ID of the Elastic IP address that's associated with the NAT gateway. Only available when `connectivity_type` is `public`.
        /// </summary>
        [Input("associationId")]
        public Input<string>? AssociationId { get; set; }

        /// <summary>
        /// Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        /// </summary>
        [Input("connectivityType")]
        public Input<string>? ConnectivityType { get; set; }

        /// <summary>
        /// The ID of the network interface associated with the NAT gateway.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// The Elastic IP address associated with the NAT gateway.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// The Subnet ID of the subnet in which to place the gateway.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public NatGatewayState()
        {
        }
        public static new NatGatewayState Empty => new NatGatewayState();
    }
}
