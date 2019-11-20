// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm
{
    /// <summary>
    /// Provides an SSM Maintenance Window Target resource
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_maintenance_window_target.html.markdown.
    /// </summary>
    public partial class MaintenanceWindowTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the maintenance window target.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the maintenance window target.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        /// </summary>
        [Output("ownerInformation")]
        public Output<string?> OwnerInformation { get; private set; } = null!;

        /// <summary>
        /// The type of target being registered with the Maintenance Window. Possible values `INSTANCE`.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// The targets (either instances or tags). Instances are specified using Key=InstanceIds,Values=InstanceId1,InstanceId2. Tags are specified using Key=tag name,Values=tag value.
        /// </summary>
        [Output("targets")]
        public Output<ImmutableArray<Outputs.MaintenanceWindowTargetTargets>> Targets { get; private set; } = null!;

        /// <summary>
        /// The Id of the maintenance window to register the target with.
        /// </summary>
        [Output("windowId")]
        public Output<string> WindowId { get; private set; } = null!;


        /// <summary>
        /// Create a MaintenanceWindowTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MaintenanceWindowTarget(string name, MaintenanceWindowTargetArgs args, CustomResourceOptions? options = null)
            : base("aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private MaintenanceWindowTarget(string name, Input<string> id, MaintenanceWindowTargetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MaintenanceWindowTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MaintenanceWindowTarget Get(string name, Input<string> id, MaintenanceWindowTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new MaintenanceWindowTarget(name, id, state, options);
        }
    }

    public sealed class MaintenanceWindowTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the maintenance window target.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the maintenance window target.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        /// </summary>
        [Input("ownerInformation")]
        public Input<string>? OwnerInformation { get; set; }

        /// <summary>
        /// The type of target being registered with the Maintenance Window. Possible values `INSTANCE`.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("targets", required: true)]
        private InputList<Inputs.MaintenanceWindowTargetTargetsArgs>? _targets;

        /// <summary>
        /// The targets (either instances or tags). Instances are specified using Key=InstanceIds,Values=InstanceId1,InstanceId2. Tags are specified using Key=tag name,Values=tag value.
        /// </summary>
        public InputList<Inputs.MaintenanceWindowTargetTargetsArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MaintenanceWindowTargetTargetsArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// The Id of the maintenance window to register the target with.
        /// </summary>
        [Input("windowId", required: true)]
        public Input<string> WindowId { get; set; } = null!;

        public MaintenanceWindowTargetArgs()
        {
        }
    }

    public sealed class MaintenanceWindowTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the maintenance window target.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the maintenance window target.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
        /// </summary>
        [Input("ownerInformation")]
        public Input<string>? OwnerInformation { get; set; }

        /// <summary>
        /// The type of target being registered with the Maintenance Window. Possible values `INSTANCE`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        [Input("targets")]
        private InputList<Inputs.MaintenanceWindowTargetTargetsGetArgs>? _targets;

        /// <summary>
        /// The targets (either instances or tags). Instances are specified using Key=InstanceIds,Values=InstanceId1,InstanceId2. Tags are specified using Key=tag name,Values=tag value.
        /// </summary>
        public InputList<Inputs.MaintenanceWindowTargetTargetsGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MaintenanceWindowTargetTargetsGetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// The Id of the maintenance window to register the target with.
        /// </summary>
        [Input("windowId")]
        public Input<string>? WindowId { get; set; }

        public MaintenanceWindowTargetState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class MaintenanceWindowTargetTargetsArgs : Pulumi.ResourceArgs
    {
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public MaintenanceWindowTargetTargetsArgs()
        {
        }
    }

    public sealed class MaintenanceWindowTargetTargetsGetArgs : Pulumi.ResourceArgs
    {
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public MaintenanceWindowTargetTargetsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class MaintenanceWindowTargetTargets
    {
        public readonly string Key;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private MaintenanceWindowTargetTargets(
            string key,
            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }
    }
}
