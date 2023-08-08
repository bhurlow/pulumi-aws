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
    /// Provides a resource to create an EventBridge Global Endpoint.
    /// 
    /// &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
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
    ///     var @this = new Aws.CloudWatch.EventEndpoint("this", new()
    ///     {
    ///         RoleArn = aws_iam_role.Replication.Arn,
    ///         EventBuses = new[]
    ///         {
    ///             new Aws.CloudWatch.Inputs.EventEndpointEventBusArgs
    ///             {
    ///                 EventBusArn = aws_cloudwatch_event_bus.Primary.Arn,
    ///             },
    ///             new Aws.CloudWatch.Inputs.EventEndpointEventBusArgs
    ///             {
    ///                 EventBusArn = aws_cloudwatch_event_bus.Secondary.Arn,
    ///             },
    ///         },
    ///         ReplicationConfig = new Aws.CloudWatch.Inputs.EventEndpointReplicationConfigArgs
    ///         {
    ///             State = "DISABLED",
    ///         },
    ///         RoutingConfig = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigArgs
    ///         {
    ///             FailoverConfig = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigFailoverConfigArgs
    ///             {
    ///                 Primary = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigFailoverConfigPrimaryArgs
    ///                 {
    ///                     HealthCheck = aws_route53_health_check.Primary.Arn,
    ///                 },
    ///                 Secondary = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigFailoverConfigSecondaryArgs
    ///                 {
    ///                     Route = "us-east-2",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_cloudwatch_event_endpoint.imported_endpoint
    /// 
    ///  id = "example-endpoint" } Using `pulumi import`, import EventBridge Global Endpoints using the `name`. For exampleconsole % pulumi import aws_cloudwatch_event_endpoint.imported_endpoint example-endpoint
    /// </summary>
    [AwsResourceType("aws:cloudwatch/eventEndpoint:EventEndpoint")]
    public partial class EventEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the endpoint that was created.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A description of the global endpoint.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The URL of the endpoint that was created.
        /// </summary>
        [Output("endpointUrl")]
        public Output<string> EndpointUrl { get; private set; } = null!;

        /// <summary>
        /// The event buses to use. The names of the event buses must be identical in each Region. Exactly two event buses are required. Documented below.
        /// </summary>
        [Output("eventBuses")]
        public Output<ImmutableArray<Outputs.EventEndpointEventBus>> EventBuses { get; private set; } = null!;

        /// <summary>
        /// The name of the global endpoint.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Parameters used for replication. Documented below.
        /// </summary>
        [Output("replicationConfig")]
        public Output<Outputs.EventEndpointReplicationConfig?> ReplicationConfig { get; private set; } = null!;

        /// <summary>
        /// The ARN of the IAM role used for replication between event buses.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Parameters used for routing, including the health check and secondary Region. Documented below.
        /// </summary>
        [Output("routingConfig")]
        public Output<Outputs.EventEndpointRoutingConfig> RoutingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a EventEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventEndpoint(string name, EventEndpointArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventEndpoint:EventEndpoint", name, args ?? new EventEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventEndpoint(string name, Input<string> id, EventEndpointState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventEndpoint:EventEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventEndpoint Get(string name, Input<string> id, EventEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new EventEndpoint(name, id, state, options);
        }
    }

    public sealed class EventEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the global endpoint.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("eventBuses", required: true)]
        private InputList<Inputs.EventEndpointEventBusArgs>? _eventBuses;

        /// <summary>
        /// The event buses to use. The names of the event buses must be identical in each Region. Exactly two event buses are required. Documented below.
        /// </summary>
        public InputList<Inputs.EventEndpointEventBusArgs> EventBuses
        {
            get => _eventBuses ?? (_eventBuses = new InputList<Inputs.EventEndpointEventBusArgs>());
            set => _eventBuses = value;
        }

        /// <summary>
        /// The name of the global endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parameters used for replication. Documented below.
        /// </summary>
        [Input("replicationConfig")]
        public Input<Inputs.EventEndpointReplicationConfigArgs>? ReplicationConfig { get; set; }

        /// <summary>
        /// The ARN of the IAM role used for replication between event buses.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Parameters used for routing, including the health check and secondary Region. Documented below.
        /// </summary>
        [Input("routingConfig", required: true)]
        public Input<Inputs.EventEndpointRoutingConfigArgs> RoutingConfig { get; set; } = null!;

        public EventEndpointArgs()
        {
        }
        public static new EventEndpointArgs Empty => new EventEndpointArgs();
    }

    public sealed class EventEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the endpoint that was created.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A description of the global endpoint.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The URL of the endpoint that was created.
        /// </summary>
        [Input("endpointUrl")]
        public Input<string>? EndpointUrl { get; set; }

        [Input("eventBuses")]
        private InputList<Inputs.EventEndpointEventBusGetArgs>? _eventBuses;

        /// <summary>
        /// The event buses to use. The names of the event buses must be identical in each Region. Exactly two event buses are required. Documented below.
        /// </summary>
        public InputList<Inputs.EventEndpointEventBusGetArgs> EventBuses
        {
            get => _eventBuses ?? (_eventBuses = new InputList<Inputs.EventEndpointEventBusGetArgs>());
            set => _eventBuses = value;
        }

        /// <summary>
        /// The name of the global endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parameters used for replication. Documented below.
        /// </summary>
        [Input("replicationConfig")]
        public Input<Inputs.EventEndpointReplicationConfigGetArgs>? ReplicationConfig { get; set; }

        /// <summary>
        /// The ARN of the IAM role used for replication between event buses.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Parameters used for routing, including the health check and secondary Region. Documented below.
        /// </summary>
        [Input("routingConfig")]
        public Input<Inputs.EventEndpointRoutingConfigGetArgs>? RoutingConfig { get; set; }

        public EventEndpointState()
        {
        }
        public static new EventEndpointState Empty => new EventEndpointState();
    }
}
