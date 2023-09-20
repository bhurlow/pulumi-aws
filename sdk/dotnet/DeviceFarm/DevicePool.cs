// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DeviceFarm
{
    /// <summary>
    /// Provides a resource to manage AWS Device Farm Device Pools.
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
    ///     var example = new Aws.DeviceFarm.DevicePool("example", new()
    ///     {
    ///         ProjectArn = aws_devicefarm_project.Example.Arn,
    ///         Rules = new[]
    ///         {
    ///             new Aws.DeviceFarm.Inputs.DevicePoolRuleArgs
    ///             {
    ///                 Attribute = "OS_VERSION",
    ///                 Operator = "EQUALS",
    ///                 Value = "\"AVAILABLE\"",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import DeviceFarm Device Pools using their ARN. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:devicefarm/devicePool:DevicePool example arn:aws:devicefarm:us-west-2:123456789012:devicepool:4fa784c7-ccb4-4dbf-ba4f-02198320daa1/4fa784c7-ccb4-4dbf-ba4f-02198320daa1
    /// ```
    /// </summary>
    [AwsResourceType("aws:devicefarm/devicePool:DevicePool")]
    public partial class DevicePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name of this Device Pool
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The device pool's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The number of devices that Device Farm can add to your device pool.
        /// </summary>
        [Output("maxDevices")]
        public Output<int?> MaxDevices { get; private set; } = null!;

        /// <summary>
        /// The name of the Device Pool
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the project for the device pool.
        /// </summary>
        [Output("projectArn")]
        public Output<string> ProjectArn { get; private set; } = null!;

        /// <summary>
        /// The device pool's rules. See Rule.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.DevicePoolRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DevicePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DevicePool(string name, DevicePoolArgs args, CustomResourceOptions? options = null)
            : base("aws:devicefarm/devicePool:DevicePool", name, args ?? new DevicePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DevicePool(string name, Input<string> id, DevicePoolState? state = null, CustomResourceOptions? options = null)
            : base("aws:devicefarm/devicePool:DevicePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DevicePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DevicePool Get(string name, Input<string> id, DevicePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new DevicePool(name, id, state, options);
        }
    }

    public sealed class DevicePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device pool's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The number of devices that Device Farm can add to your device pool.
        /// </summary>
        [Input("maxDevices")]
        public Input<int>? MaxDevices { get; set; }

        /// <summary>
        /// The name of the Device Pool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the project for the device pool.
        /// </summary>
        [Input("projectArn", required: true)]
        public Input<string> ProjectArn { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.DevicePoolRuleArgs>? _rules;

        /// <summary>
        /// The device pool's rules. See Rule.
        /// </summary>
        public InputList<Inputs.DevicePoolRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.DevicePoolRuleArgs>());
            set => _rules = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DevicePoolArgs()
        {
        }
        public static new DevicePoolArgs Empty => new DevicePoolArgs();
    }

    public sealed class DevicePoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name of this Device Pool
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The device pool's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The number of devices that Device Farm can add to your device pool.
        /// </summary>
        [Input("maxDevices")]
        public Input<int>? MaxDevices { get; set; }

        /// <summary>
        /// The name of the Device Pool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the project for the device pool.
        /// </summary>
        [Input("projectArn")]
        public Input<string>? ProjectArn { get; set; }

        [Input("rules")]
        private InputList<Inputs.DevicePoolRuleGetArgs>? _rules;

        /// <summary>
        /// The device pool's rules. See Rule.
        /// </summary>
        public InputList<Inputs.DevicePoolRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.DevicePoolRuleGetArgs>());
            set => _rules = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        [Input("type")]
        public Input<string>? Type { get; set; }

        public DevicePoolState()
        {
        }
        public static new DevicePoolState Empty => new DevicePoolState();
    }
}
