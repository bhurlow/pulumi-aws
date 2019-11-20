// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms
{
    /// <summary>
    /// Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dms_replication_task.html.markdown.
    /// </summary>
    public partial class ReplicationTask : Pulumi.CustomResource
    {
        /// <summary>
        /// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        /// </summary>
        [Output("cdcStartTime")]
        public Output<string?> CdcStartTime { get; private set; } = null!;

        /// <summary>
        /// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        /// </summary>
        [Output("migrationType")]
        public Output<string> MigrationType { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the replication instance.
        /// </summary>
        [Output("replicationInstanceArn")]
        public Output<string> ReplicationInstanceArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the replication task.
        /// </summary>
        [Output("replicationTaskArn")]
        public Output<string> ReplicationTaskArn { get; private set; } = null!;

        /// <summary>
        /// The replication task identifier.
        /// </summary>
        [Output("replicationTaskId")]
        public Output<string> ReplicationTaskId { get; private set; } = null!;

        /// <summary>
        /// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        /// </summary>
        [Output("replicationTaskSettings")]
        public Output<string?> ReplicationTaskSettings { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        /// </summary>
        [Output("sourceEndpointArn")]
        public Output<string> SourceEndpointArn { get; private set; } = null!;

        /// <summary>
        /// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        /// </summary>
        [Output("tableMappings")]
        public Output<string> TableMappings { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        /// </summary>
        [Output("targetEndpointArn")]
        public Output<string> TargetEndpointArn { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationTask(string name, ReplicationTaskArgs args, CustomResourceOptions? options = null)
            : base("aws:dms/replicationTask:ReplicationTask", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ReplicationTask(string name, Input<string> id, ReplicationTaskState? state = null, CustomResourceOptions? options = null)
            : base("aws:dms/replicationTask:ReplicationTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicationTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationTask Get(string name, Input<string> id, ReplicationTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationTask(name, id, state, options);
        }
    }

    public sealed class ReplicationTaskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        /// </summary>
        [Input("cdcStartTime")]
        public Input<string>? CdcStartTime { get; set; }

        /// <summary>
        /// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        /// </summary>
        [Input("migrationType", required: true)]
        public Input<string> MigrationType { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the replication instance.
        /// </summary>
        [Input("replicationInstanceArn", required: true)]
        public Input<string> ReplicationInstanceArn { get; set; } = null!;

        /// <summary>
        /// The replication task identifier.
        /// </summary>
        [Input("replicationTaskId", required: true)]
        public Input<string> ReplicationTaskId { get; set; } = null!;

        /// <summary>
        /// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        /// </summary>
        [Input("replicationTaskSettings")]
        public Input<string>? ReplicationTaskSettings { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        /// </summary>
        [Input("sourceEndpointArn", required: true)]
        public Input<string> SourceEndpointArn { get; set; } = null!;

        /// <summary>
        /// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        /// </summary>
        [Input("tableMappings", required: true)]
        public Input<string> TableMappings { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        /// </summary>
        [Input("targetEndpointArn", required: true)]
        public Input<string> TargetEndpointArn { get; set; } = null!;

        public ReplicationTaskArgs()
        {
        }
    }

    public sealed class ReplicationTaskState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        /// </summary>
        [Input("cdcStartTime")]
        public Input<string>? CdcStartTime { get; set; }

        /// <summary>
        /// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        /// </summary>
        [Input("migrationType")]
        public Input<string>? MigrationType { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the replication instance.
        /// </summary>
        [Input("replicationInstanceArn")]
        public Input<string>? ReplicationInstanceArn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the replication task.
        /// </summary>
        [Input("replicationTaskArn")]
        public Input<string>? ReplicationTaskArn { get; set; }

        /// <summary>
        /// The replication task identifier.
        /// </summary>
        [Input("replicationTaskId")]
        public Input<string>? ReplicationTaskId { get; set; }

        /// <summary>
        /// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        /// </summary>
        [Input("replicationTaskSettings")]
        public Input<string>? ReplicationTaskSettings { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        /// </summary>
        [Input("sourceEndpointArn")]
        public Input<string>? SourceEndpointArn { get; set; }

        /// <summary>
        /// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        /// </summary>
        [Input("tableMappings")]
        public Input<string>? TableMappings { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        /// </summary>
        [Input("targetEndpointArn")]
        public Input<string>? TargetEndpointArn { get; set; }

        public ReplicationTaskState()
        {
        }
    }
}
