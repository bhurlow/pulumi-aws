// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LightSail
{
    /// <summary>
    /// Provides a Lightsail Disk resource.
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
    ///     var available = Aws.GetAvailabilityZones.Invoke(new()
    ///     {
    ///         State = "available",
    ///         Filters = new[]
    ///         {
    ///             new Aws.Inputs.GetAvailabilityZonesFilterInputArgs
    ///             {
    ///                 Name = "opt-in-status",
    ///                 Values = new[]
    ///                 {
    ///                     "opt-in-not-required",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var test = new Aws.LightSail.Disk("test", new()
    ///     {
    ///         SizeInGb = 8,
    ///         AvailabilityZone = available.Apply(getAvailabilityZonesResult =&gt; getAvailabilityZonesResult.Names[0]),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_lightsail_disk` using the name attribute. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:lightsail/disk:Disk test test
    /// ```
    /// </summary>
    [AwsResourceType("aws:lightsail/disk:Disk")]
    public partial class Disk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Lightsail load balancer.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone in which to create your disk.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the load balancer was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The name of the Lightsail load balancer.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The instance port the load balancer will connect.
        /// </summary>
        [Output("sizeInGb")]
        public Output<int> SizeInGb { get; private set; } = null!;

        /// <summary>
        /// The support code for the disk. Include this code in your email to support when you have questions about a disk in Lightsail. This code enables our support team to look up your Lightsail information more easily.
        /// </summary>
        [Output("supportCode")]
        public Output<string> SupportCode { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs args, CustomResourceOptions? options = null)
            : base("aws:lightsail/disk:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
            : base("aws:lightsail/disk:Disk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, state, options);
        }
    }

    public sealed class DiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Availability Zone in which to create your disk.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// The name of the Lightsail load balancer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The instance port the load balancer will connect.
        /// </summary>
        [Input("sizeInGb", required: true)]
        public Input<int> SizeInGb { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DiskArgs()
        {
        }
        public static new DiskArgs Empty => new DiskArgs();
    }

    public sealed class DiskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Lightsail load balancer.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Availability Zone in which to create your disk.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The timestamp when the load balancer was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The name of the Lightsail load balancer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The instance port the load balancer will connect.
        /// </summary>
        [Input("sizeInGb")]
        public Input<int>? SizeInGb { get; set; }

        /// <summary>
        /// The support code for the disk. Include this code in your email to support when you have questions about a disk in Lightsail. This code enables our support team to look up your Lightsail information more easily.
        /// </summary>
        [Input("supportCode")]
        public Input<string>? SupportCode { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public DiskState()
        {
        }
        public static new DiskState Empty => new DiskState();
    }
}
