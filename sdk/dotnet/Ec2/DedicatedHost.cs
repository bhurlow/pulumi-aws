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
    /// Provides an EC2 Host resource. This allows Dedicated Hosts to be allocated, modified, and released.
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
    ///     // Create a new host with instance type of c5.18xlarge with Auto Placement
    ///     // and Host Recovery enabled.
    ///     var test = new Aws.Ec2.DedicatedHost("test", new()
    ///     {
    ///         AutoPlacement = "on",
    ///         AvailabilityZone = "us-west-2a",
    ///         HostRecovery = "on",
    ///         InstanceType = "c5.18xlarge",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_ec2_host.example
    /// 
    ///  id = "h-0385a99d0e4b20cbb" } Using `pulumi import`, import hosts using the host `id`. For exampleconsole % pulumi import aws_ec2_host.example h-0385a99d0e4b20cbb
    /// </summary>
    [AwsResourceType("aws:ec2/dedicatedHost:DedicatedHost")]
    public partial class DedicatedHost : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Dedicated Host.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        /// </summary>
        [Output("assetId")]
        public Output<string?> AssetId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        /// </summary>
        [Output("autoPlacement")]
        public Output<string?> AutoPlacement { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone in which to allocate the Dedicated Host.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        /// </summary>
        [Output("hostRecovery")]
        public Output<string?> HostRecovery { get; private set; } = null!;

        /// <summary>
        /// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        /// </summary>
        [Output("instanceFamily")]
        public Output<string?> InstanceFamily { get; private set; } = null!;

        /// <summary>
        /// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        /// </summary>
        [Output("instanceType")]
        public Output<string?> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        /// </summary>
        [Output("outpostArn")]
        public Output<string?> OutpostArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the Dedicated Host.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedHost(string name, DedicatedHostArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/dedicatedHost:DedicatedHost", name, args ?? new DedicatedHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedHost(string name, Input<string> id, DedicatedHostState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/dedicatedHost:DedicatedHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DedicatedHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedHost Get(string name, Input<string> id, DedicatedHostState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedHost(name, id, state, options);
        }
    }

    public sealed class DedicatedHostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        /// </summary>
        [Input("assetId")]
        public Input<string>? AssetId { get; set; }

        /// <summary>
        /// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        /// </summary>
        [Input("autoPlacement")]
        public Input<string>? AutoPlacement { get; set; }

        /// <summary>
        /// The Availability Zone in which to allocate the Dedicated Host.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        /// </summary>
        [Input("hostRecovery")]
        public Input<string>? HostRecovery { get; set; }

        /// <summary>
        /// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        /// </summary>
        [Input("instanceFamily")]
        public Input<string>? InstanceFamily { get; set; }

        /// <summary>
        /// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DedicatedHostArgs()
        {
        }
        public static new DedicatedHostArgs Empty => new DedicatedHostArgs();
    }

    public sealed class DedicatedHostState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Dedicated Host.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
        /// </summary>
        [Input("assetId")]
        public Input<string>? AssetId { get; set; }

        /// <summary>
        /// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
        /// </summary>
        [Input("autoPlacement")]
        public Input<string>? AutoPlacement { get; set; }

        /// <summary>
        /// The Availability Zone in which to allocate the Dedicated Host.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
        /// </summary>
        [Input("hostRecovery")]
        public Input<string>? HostRecovery { get; set; }

        /// <summary>
        /// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
        /// </summary>
        [Input("instanceFamily")]
        public Input<string>? InstanceFamily { get; set; }

        /// <summary>
        /// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the Dedicated Host.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public DedicatedHostState()
        {
        }
        public static new DedicatedHostState Empty => new DedicatedHostState();
    }
}
