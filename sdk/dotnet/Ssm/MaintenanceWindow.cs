// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm
{
    /// <summary>
    /// Provides an SSM Maintenance Window resource
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
    ///     var production = new Aws.Ssm.MaintenanceWindow("production", new()
    ///     {
    ///         Cutoff = 1,
    ///         Duration = 3,
    ///         Schedule = "cron(0 16 ? * TUE *)",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SSM
    /// 
    /// Maintenance Windows can be imported using the `maintenance window id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ssm/maintenanceWindow:MaintenanceWindow imported-window mw-0123456789
    /// ```
    /// </summary>
    [AwsResourceType("aws:ssm/maintenanceWindow:MaintenanceWindow")]
    public partial class MaintenanceWindow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
        /// </summary>
        [Output("allowUnassociatedTargets")]
        public Output<bool?> AllowUnassociatedTargets { get; private set; } = null!;

        /// <summary>
        /// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
        /// </summary>
        [Output("cutoff")]
        public Output<int> Cutoff { get; private set; } = null!;

        /// <summary>
        /// A description for the maintenance window.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The duration of the Maintenance Window in hours.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// Whether the maintenance window is enabled. Default: `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
        /// </summary>
        [Output("endDate")]
        public Output<string?> EndDate { get; private set; } = null!;

        /// <summary>
        /// The name of the maintenance window.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
        /// </summary>
        [Output("scheduleOffset")]
        public Output<int?> ScheduleOffset { get; private set; } = null!;

        /// <summary>
        /// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
        /// </summary>
        [Output("scheduleTimezone")]
        public Output<string?> ScheduleTimezone { get; private set; } = null!;

        /// <summary>
        /// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
        /// </summary>
        [Output("startDate")]
        public Output<string?> StartDate { get; private set; } = null!;

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
        /// Create a MaintenanceWindow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MaintenanceWindow(string name, MaintenanceWindowArgs args, CustomResourceOptions? options = null)
            : base("aws:ssm/maintenanceWindow:MaintenanceWindow", name, args ?? new MaintenanceWindowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MaintenanceWindow(string name, Input<string> id, MaintenanceWindowState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssm/maintenanceWindow:MaintenanceWindow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MaintenanceWindow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MaintenanceWindow Get(string name, Input<string> id, MaintenanceWindowState? state = null, CustomResourceOptions? options = null)
        {
            return new MaintenanceWindow(name, id, state, options);
        }
    }

    public sealed class MaintenanceWindowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
        /// </summary>
        [Input("allowUnassociatedTargets")]
        public Input<bool>? AllowUnassociatedTargets { get; set; }

        /// <summary>
        /// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
        /// </summary>
        [Input("cutoff", required: true)]
        public Input<int> Cutoff { get; set; } = null!;

        /// <summary>
        /// A description for the maintenance window.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The duration of the Maintenance Window in hours.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// Whether the maintenance window is enabled. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// The name of the maintenance window.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        /// <summary>
        /// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
        /// </summary>
        [Input("scheduleOffset")]
        public Input<int>? ScheduleOffset { get; set; }

        /// <summary>
        /// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
        /// </summary>
        [Input("scheduleTimezone")]
        public Input<string>? ScheduleTimezone { get; set; }

        /// <summary>
        /// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

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

        public MaintenanceWindowArgs()
        {
        }
        public static new MaintenanceWindowArgs Empty => new MaintenanceWindowArgs();
    }

    public sealed class MaintenanceWindowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
        /// </summary>
        [Input("allowUnassociatedTargets")]
        public Input<bool>? AllowUnassociatedTargets { get; set; }

        /// <summary>
        /// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
        /// </summary>
        [Input("cutoff")]
        public Input<int>? Cutoff { get; set; }

        /// <summary>
        /// A description for the maintenance window.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The duration of the Maintenance Window in hours.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Whether the maintenance window is enabled. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// The name of the maintenance window.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
        /// </summary>
        [Input("scheduleOffset")]
        public Input<int>? ScheduleOffset { get; set; }

        /// <summary>
        /// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
        /// </summary>
        [Input("scheduleTimezone")]
        public Input<string>? ScheduleTimezone { get; set; }

        /// <summary>
        /// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

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

        public MaintenanceWindowState()
        {
        }
        public static new MaintenanceWindowState Empty => new MaintenanceWindowState();
    }
}
