// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs
{
    /// <summary>
    /// Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).
    /// 
    /// &gt; **NOTE:** Associating an ECS Capacity Provider to an Auto Scaling Group will automatically add the `AmazonECSManaged` tag to the Auto Scaling Group. This tag should be included in the `aws.autoscaling.Group` resource configuration to prevent the provider from removing it in subsequent executions as well as ensuring the `AmazonECSManaged` tag is propagated to all EC2 Instances in the Auto Scaling Group if `min_size` is above 0 on creation. Any EC2 Instances in the Auto Scaling Group without this tag must be manually be updated, otherwise they may cause unexpected scaling behavior and metrics.
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
    ///     // ... other configuration, including potentially other tags ...
    ///     var testGroup = new Aws.AutoScaling.Group("testGroup", new()
    ///     {
    ///         Tags = new[]
    ///         {
    ///             new Aws.AutoScaling.Inputs.GroupTagArgs
    ///             {
    ///                 Key = "AmazonECSManaged",
    ///                 Value = "true",
    ///                 PropagateAtLaunch = true,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var testCapacityProvider = new Aws.Ecs.CapacityProvider("testCapacityProvider", new()
    ///     {
    ///         AutoScalingGroupProvider = new Aws.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
    ///         {
    ///             AutoScalingGroupArn = testGroup.Arn,
    ///             ManagedTerminationProtection = "ENABLED",
    ///             ManagedScaling = new Aws.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs
    ///             {
    ///                 MaximumScalingStepSize = 1000,
    ///                 MinimumScalingStepSize = 1,
    ///                 Status = "ENABLED",
    ///                 TargetCapacity = 10,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import ECS Capacity Providers using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecs/capacityProvider:CapacityProvider example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecs/capacityProvider:CapacityProvider")]
    public partial class CapacityProvider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN that identifies the capacity provider.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the provider for the ECS auto scaling group. Detailed below.
        /// </summary>
        [Output("autoScalingGroupProvider")]
        public Output<Outputs.CapacityProviderAutoScalingGroupProvider> AutoScalingGroupProvider { get; private set; } = null!;

        /// <summary>
        /// Name of the capacity provider.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a CapacityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CapacityProvider(string name, CapacityProviderArgs args, CustomResourceOptions? options = null)
            : base("aws:ecs/capacityProvider:CapacityProvider", name, args ?? new CapacityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CapacityProvider(string name, Input<string> id, CapacityProviderState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecs/capacityProvider:CapacityProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CapacityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CapacityProvider Get(string name, Input<string> id, CapacityProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new CapacityProvider(name, id, state, options);
        }
    }

    public sealed class CapacityProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block for the provider for the ECS auto scaling group. Detailed below.
        /// </summary>
        [Input("autoScalingGroupProvider", required: true)]
        public Input<Inputs.CapacityProviderAutoScalingGroupProviderArgs> AutoScalingGroupProvider { get; set; } = null!;

        /// <summary>
        /// Name of the capacity provider.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CapacityProviderArgs()
        {
        }
        public static new CapacityProviderArgs Empty => new CapacityProviderArgs();
    }

    public sealed class CapacityProviderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN that identifies the capacity provider.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Configuration block for the provider for the ECS auto scaling group. Detailed below.
        /// </summary>
        [Input("autoScalingGroupProvider")]
        public Input<Inputs.CapacityProviderAutoScalingGroupProviderGetArgs>? AutoScalingGroupProvider { get; set; }

        /// <summary>
        /// Name of the capacity provider.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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

        public CapacityProviderState()
        {
        }
        public static new CapacityProviderState Empty => new CapacityProviderState();
    }
}
