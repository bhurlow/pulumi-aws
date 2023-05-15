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
    /// Provides an ECS task set - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).
    /// 
    /// See [ECS Task Set section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-external.html).
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
    ///     var example = new Aws.Ecs.TaskSet("example", new()
    ///     {
    ///         Service = aws_ecs_service.Example.Id,
    ///         Cluster = aws_ecs_cluster.Example.Id,
    ///         TaskDefinition = aws_ecs_task_definition.Example.Arn,
    ///         LoadBalancers = new[]
    ///         {
    ///             new Aws.Ecs.Inputs.TaskSetLoadBalancerArgs
    ///             {
    ///                 TargetGroupArn = aws_lb_target_group.Example.Arn,
    ///                 ContainerName = "mongo",
    ///                 ContainerPort = 8080,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS Task Sets can be imported via the `task_set_id`, `service`, and `cluster` separated by commas (`,`) e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecs/taskSet:TaskSet example ecs-svc/7177320696926227436,arn:aws:ecs:us-west-2:123456789101:service/example/example-1234567890,arn:aws:ecs:us-west-2:123456789101:cluster/example
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecs/taskSet:TaskSet")]
    public partial class TaskSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that identifies the task set.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
        /// </summary>
        [Output("capacityProviderStrategies")]
        public Output<ImmutableArray<Outputs.TaskSetCapacityProviderStrategy>> CapacityProviderStrategies { get; private set; } = null!;

        /// <summary>
        /// The short name or ARN of the cluster that hosts the service to create the task set in.
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// The external ID associated with the task set.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
        /// </summary>
        [Output("launchType")]
        public Output<string> LaunchType { get; private set; } = null!;

        /// <summary>
        /// Details on load balancers that are used with a task set. Detailed below.
        /// </summary>
        [Output("loadBalancers")]
        public Output<ImmutableArray<Outputs.TaskSetLoadBalancer>> LoadBalancers { get; private set; } = null!;

        /// <summary>
        /// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
        /// </summary>
        [Output("networkConfiguration")]
        public Output<Outputs.TaskSetNetworkConfiguration?> NetworkConfiguration { get; private set; } = null!;

        /// <summary>
        /// The platform version on which to run your service. Only applicable for `launch_type` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
        /// </summary>
        [Output("platformVersion")]
        public Output<string> PlatformVersion { get; private set; } = null!;

        /// <summary>
        /// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
        /// </summary>
        [Output("scale")]
        public Output<Outputs.TaskSetScale> Scale { get; private set; } = null!;

        /// <summary>
        /// The short name or ARN of the ECS service.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// The service discovery registries for the service. The maximum number of `service_registries` blocks is `1`. Detailed below.
        /// </summary>
        [Output("serviceRegistries")]
        public Output<Outputs.TaskSetServiceRegistries?> ServiceRegistries { get; private set; } = null!;

        /// <summary>
        /// The stability status. This indicates whether the task set has reached a steady state.
        /// </summary>
        [Output("stabilityStatus")]
        public Output<string> StabilityStatus { get; private set; } = null!;

        /// <summary>
        /// The status of the task set.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
        /// </summary>
        [Output("taskDefinition")]
        public Output<string> TaskDefinition { get; private set; } = null!;

        /// <summary>
        /// The ID of the task set.
        /// </summary>
        [Output("taskSetId")]
        public Output<string> TaskSetId { get; private set; } = null!;

        /// <summary>
        /// Whether the provider should wait until the task set has reached `STEADY_STATE`.
        /// </summary>
        [Output("waitUntilStable")]
        public Output<bool?> WaitUntilStable { get; private set; } = null!;

        /// <summary>
        /// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
        /// </summary>
        [Output("waitUntilStableTimeout")]
        public Output<string?> WaitUntilStableTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a TaskSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TaskSet(string name, TaskSetArgs args, CustomResourceOptions? options = null)
            : base("aws:ecs/taskSet:TaskSet", name, args ?? new TaskSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TaskSet(string name, Input<string> id, TaskSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecs/taskSet:TaskSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TaskSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TaskSet Get(string name, Input<string> id, TaskSetState? state = null, CustomResourceOptions? options = null)
        {
            return new TaskSet(name, id, state, options);
        }
    }

    public sealed class TaskSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("capacityProviderStrategies")]
        private InputList<Inputs.TaskSetCapacityProviderStrategyArgs>? _capacityProviderStrategies;

        /// <summary>
        /// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
        /// </summary>
        public InputList<Inputs.TaskSetCapacityProviderStrategyArgs> CapacityProviderStrategies
        {
            get => _capacityProviderStrategies ?? (_capacityProviderStrategies = new InputList<Inputs.TaskSetCapacityProviderStrategyArgs>());
            set => _capacityProviderStrategies = value;
        }

        /// <summary>
        /// The short name or ARN of the cluster that hosts the service to create the task set in.
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        /// <summary>
        /// The external ID associated with the task set.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
        /// </summary>
        [Input("launchType")]
        public Input<string>? LaunchType { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.TaskSetLoadBalancerArgs>? _loadBalancers;

        /// <summary>
        /// Details on load balancers that are used with a task set. Detailed below.
        /// </summary>
        public InputList<Inputs.TaskSetLoadBalancerArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.TaskSetLoadBalancerArgs>());
            set => _loadBalancers = value;
        }

        /// <summary>
        /// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
        /// </summary>
        [Input("networkConfiguration")]
        public Input<Inputs.TaskSetNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        /// <summary>
        /// The platform version on which to run your service. Only applicable for `launch_type` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
        /// </summary>
        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        /// <summary>
        /// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
        /// </summary>
        [Input("scale")]
        public Input<Inputs.TaskSetScaleArgs>? Scale { get; set; }

        /// <summary>
        /// The short name or ARN of the ECS service.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// The service discovery registries for the service. The maximum number of `service_registries` blocks is `1`. Detailed below.
        /// </summary>
        [Input("serviceRegistries")]
        public Input<Inputs.TaskSetServiceRegistriesArgs>? ServiceRegistries { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
        /// </summary>
        [Input("taskDefinition", required: true)]
        public Input<string> TaskDefinition { get; set; } = null!;

        /// <summary>
        /// Whether the provider should wait until the task set has reached `STEADY_STATE`.
        /// </summary>
        [Input("waitUntilStable")]
        public Input<bool>? WaitUntilStable { get; set; }

        /// <summary>
        /// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
        /// </summary>
        [Input("waitUntilStableTimeout")]
        public Input<string>? WaitUntilStableTimeout { get; set; }

        public TaskSetArgs()
        {
        }
        public static new TaskSetArgs Empty => new TaskSetArgs();
    }

    public sealed class TaskSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that identifies the task set.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("capacityProviderStrategies")]
        private InputList<Inputs.TaskSetCapacityProviderStrategyGetArgs>? _capacityProviderStrategies;

        /// <summary>
        /// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
        /// </summary>
        public InputList<Inputs.TaskSetCapacityProviderStrategyGetArgs> CapacityProviderStrategies
        {
            get => _capacityProviderStrategies ?? (_capacityProviderStrategies = new InputList<Inputs.TaskSetCapacityProviderStrategyGetArgs>());
            set => _capacityProviderStrategies = value;
        }

        /// <summary>
        /// The short name or ARN of the cluster that hosts the service to create the task set in.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// The external ID associated with the task set.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
        /// </summary>
        [Input("launchType")]
        public Input<string>? LaunchType { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.TaskSetLoadBalancerGetArgs>? _loadBalancers;

        /// <summary>
        /// Details on load balancers that are used with a task set. Detailed below.
        /// </summary>
        public InputList<Inputs.TaskSetLoadBalancerGetArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.TaskSetLoadBalancerGetArgs>());
            set => _loadBalancers = value;
        }

        /// <summary>
        /// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
        /// </summary>
        [Input("networkConfiguration")]
        public Input<Inputs.TaskSetNetworkConfigurationGetArgs>? NetworkConfiguration { get; set; }

        /// <summary>
        /// The platform version on which to run your service. Only applicable for `launch_type` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
        /// </summary>
        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        /// <summary>
        /// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
        /// </summary>
        [Input("scale")]
        public Input<Inputs.TaskSetScaleGetArgs>? Scale { get; set; }

        /// <summary>
        /// The short name or ARN of the ECS service.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// The service discovery registries for the service. The maximum number of `service_registries` blocks is `1`. Detailed below.
        /// </summary>
        [Input("serviceRegistries")]
        public Input<Inputs.TaskSetServiceRegistriesGetArgs>? ServiceRegistries { get; set; }

        /// <summary>
        /// The stability status. This indicates whether the task set has reached a steady state.
        /// </summary>
        [Input("stabilityStatus")]
        public Input<string>? StabilityStatus { get; set; }

        /// <summary>
        /// The status of the task set.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
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

        /// <summary>
        /// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
        /// </summary>
        [Input("taskDefinition")]
        public Input<string>? TaskDefinition { get; set; }

        /// <summary>
        /// The ID of the task set.
        /// </summary>
        [Input("taskSetId")]
        public Input<string>? TaskSetId { get; set; }

        /// <summary>
        /// Whether the provider should wait until the task set has reached `STEADY_STATE`.
        /// </summary>
        [Input("waitUntilStable")]
        public Input<bool>? WaitUntilStable { get; set; }

        /// <summary>
        /// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
        /// </summary>
        [Input("waitUntilStableTimeout")]
        public Input<string>? WaitUntilStableTimeout { get; set; }

        public TaskSetState()
        {
        }
        public static new TaskSetState Empty => new TaskSetState();
    }
}
