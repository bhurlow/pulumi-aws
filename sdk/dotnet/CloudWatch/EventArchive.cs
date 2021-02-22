// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides an EventBridge event archive resource.
    /// 
    /// &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var orderEventBus = new Aws.CloudWatch.EventBus("orderEventBus", new Aws.CloudWatch.EventBusArgs
    ///         {
    ///         });
    ///         var orderEventArchive = new Aws.CloudWatch.EventArchive("orderEventArchive", new Aws.CloudWatch.EventArchiveArgs
    ///         {
    ///             EventSourceArn = orderEventBus.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example all optional arguments
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var orderEventBus = new Aws.CloudWatch.EventBus("orderEventBus", new Aws.CloudWatch.EventBusArgs
    ///         {
    ///         });
    ///         var orderEventArchive = new Aws.CloudWatch.EventArchive("orderEventArchive", new Aws.CloudWatch.EventArchiveArgs
    ///         {
    ///             Description = "Archived events from order service",
    ///             EventSourceArn = orderEventBus.Arn,
    ///             RetentionDays = 7,
    ///             EventPattern = @"{
    ///   ""source"": [""company.team.order""]
    /// }
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Event Archive can be imported using their name, for example bash
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudwatch/eventArchive:EventArchive imported_event_archive order-archive
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudwatch/eventArchive:EventArchive")]
    public partial class EventArchive : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the event archive.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the new event archive.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `event_source_arn`.
        /// </summary>
        [Output("eventPattern")]
        public Output<string?> EventPattern { get; private set; } = null!;

        /// <summary>
        /// Event bus source ARN from where these events should be archived.
        /// </summary>
        [Output("eventSourceArn")]
        public Output<string> EventSourceArn { get; private set; } = null!;

        /// <summary>
        /// The name of the new event archive. The archive name cannot exceed 48 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;


        /// <summary>
        /// Create a EventArchive resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventArchive(string name, EventArchiveArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventArchive:EventArchive", name, args ?? new EventArchiveArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventArchive(string name, Input<string> id, EventArchiveState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventArchive:EventArchive", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventArchive resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventArchive Get(string name, Input<string> id, EventArchiveState? state = null, CustomResourceOptions? options = null)
        {
            return new EventArchive(name, id, state, options);
        }
    }

    public sealed class EventArchiveArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the new event archive.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `event_source_arn`.
        /// </summary>
        [Input("eventPattern")]
        public Input<string>? EventPattern { get; set; }

        /// <summary>
        /// Event bus source ARN from where these events should be archived.
        /// </summary>
        [Input("eventSourceArn", required: true)]
        public Input<string> EventSourceArn { get; set; } = null!;

        /// <summary>
        /// The name of the new event archive. The archive name cannot exceed 48 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public EventArchiveArgs()
        {
        }
    }

    public sealed class EventArchiveState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the event archive.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description of the new event archive.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `event_source_arn`.
        /// </summary>
        [Input("eventPattern")]
        public Input<string>? EventPattern { get; set; }

        /// <summary>
        /// Event bus source ARN from where these events should be archived.
        /// </summary>
        [Input("eventSourceArn")]
        public Input<string>? EventSourceArn { get; set; }

        /// <summary>
        /// The name of the new event archive. The archive name cannot exceed 48 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public EventArchiveState()
        {
        }
    }
}
