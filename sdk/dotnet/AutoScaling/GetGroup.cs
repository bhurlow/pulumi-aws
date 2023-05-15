// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling
{
    public static class GetGroup
    {
        /// <summary>
        /// Use this data source to get information on an existing autoscaling group.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aws.AutoScaling.GetGroup.Invoke(new()
        ///     {
        ///         Name = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("aws:autoscaling/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an existing autoscaling group.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aws.AutoScaling.GetGroup.Invoke(new()
        ///     {
        ///         Name = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("aws:autoscaling/getGroup:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the exact name of the desired autoscaling group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the exact name of the desired autoscaling group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// ARN of the Auto Scaling group.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// One or more Availability Zones for the group.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly int DefaultCooldown;
        /// <summary>
        /// Desired size of the group.
        /// </summary>
        public readonly int DesiredCapacity;
        /// <summary>
        /// The unit of measurement for the value returned for `desired_capacity`.
        /// </summary>
        public readonly string DesiredCapacityType;
        /// <summary>
        /// List of metrics enabled for collection.
        /// </summary>
        public readonly ImmutableArray<string> EnabledMetrics;
        /// <summary>
        /// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
        /// </summary>
        public readonly int HealthCheckGracePeriod;
        /// <summary>
        /// Service to use for the health checks. The valid values are EC2 and ELB.
        /// </summary>
        public readonly string HealthCheckType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the associated launch configuration.
        /// </summary>
        public readonly string LaunchConfiguration;
        /// <summary>
        /// List of launch templates along with the overrides.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupLaunchTemplateResult> LaunchTemplates;
        /// <summary>
        /// One or more load balancers associated with the group.
        /// </summary>
        public readonly ImmutableArray<string> LoadBalancers;
        /// <summary>
        /// Maximum amount of time, in seconds, that an instance can be in service.
        /// </summary>
        public readonly int MaxInstanceLifetime;
        /// <summary>
        /// Maximum size of the group.
        /// </summary>
        public readonly int MaxSize;
        /// <summary>
        /// Minimum number of instances to maintain in the warm pool.
        /// </summary>
        public readonly int MinSize;
        /// <summary>
        /// List of mixed instances policy objects for the group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMixedInstancesPolicyResult> MixedInstancesPolicies;
        /// <summary>
        /// Name of the Auto Scaling Group.
        /// </summary>
        public readonly string Name;
        public readonly bool NewInstancesProtectedFromScaleIn;
        /// <summary>
        /// Name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.
        /// </summary>
        public readonly string PlacementGroup;
        /// <summary>
        /// Predicted capacity of the group.
        /// </summary>
        public readonly int PredictedCapacity;
        /// <summary>
        /// ARN of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
        /// </summary>
        public readonly string ServiceLinkedRoleArn;
        /// <summary>
        /// Current state of the group when DeleteAutoScalingGroup is in progress.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// List of processes suspended processes for the Auto Scaling Group.
        /// </summary>
        public readonly ImmutableArray<string> SuspendedProcesses;
        /// <summary>
        /// List of tags for the group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupTagResult> Tags;
        /// <summary>
        /// ARNs of the target groups for your load balancer.
        /// </summary>
        public readonly ImmutableArray<string> TargetGroupArns;
        /// <summary>
        /// The termination policies for the group.
        /// </summary>
        public readonly ImmutableArray<string> TerminationPolicies;
        /// <summary>
        /// VPC ID for the group.
        /// </summary>
        public readonly string VpcZoneIdentifier;
        /// <summary>
        /// Current size of the warm pool.
        /// </summary>
        public readonly int WarmPoolSize;
        /// <summary>
        /// List of warm pool configuration objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupWarmPoolResult> WarmPools;

        [OutputConstructor]
        private GetGroupResult(
            string arn,

            ImmutableArray<string> availabilityZones,

            int defaultCooldown,

            int desiredCapacity,

            string desiredCapacityType,

            ImmutableArray<string> enabledMetrics,

            int healthCheckGracePeriod,

            string healthCheckType,

            string id,

            string launchConfiguration,

            ImmutableArray<Outputs.GetGroupLaunchTemplateResult> launchTemplates,

            ImmutableArray<string> loadBalancers,

            int maxInstanceLifetime,

            int maxSize,

            int minSize,

            ImmutableArray<Outputs.GetGroupMixedInstancesPolicyResult> mixedInstancesPolicies,

            string name,

            bool newInstancesProtectedFromScaleIn,

            string placementGroup,

            int predictedCapacity,

            string serviceLinkedRoleArn,

            string status,

            ImmutableArray<string> suspendedProcesses,

            ImmutableArray<Outputs.GetGroupTagResult> tags,

            ImmutableArray<string> targetGroupArns,

            ImmutableArray<string> terminationPolicies,

            string vpcZoneIdentifier,

            int warmPoolSize,

            ImmutableArray<Outputs.GetGroupWarmPoolResult> warmPools)
        {
            Arn = arn;
            AvailabilityZones = availabilityZones;
            DefaultCooldown = defaultCooldown;
            DesiredCapacity = desiredCapacity;
            DesiredCapacityType = desiredCapacityType;
            EnabledMetrics = enabledMetrics;
            HealthCheckGracePeriod = healthCheckGracePeriod;
            HealthCheckType = healthCheckType;
            Id = id;
            LaunchConfiguration = launchConfiguration;
            LaunchTemplates = launchTemplates;
            LoadBalancers = loadBalancers;
            MaxInstanceLifetime = maxInstanceLifetime;
            MaxSize = maxSize;
            MinSize = minSize;
            MixedInstancesPolicies = mixedInstancesPolicies;
            Name = name;
            NewInstancesProtectedFromScaleIn = newInstancesProtectedFromScaleIn;
            PlacementGroup = placementGroup;
            PredictedCapacity = predictedCapacity;
            ServiceLinkedRoleArn = serviceLinkedRoleArn;
            Status = status;
            SuspendedProcesses = suspendedProcesses;
            Tags = tags;
            TargetGroupArns = targetGroupArns;
            TerminationPolicies = terminationPolicies;
            VpcZoneIdentifier = vpcZoneIdentifier;
            WarmPoolSize = warmPoolSize;
            WarmPools = warmPools;
        }
    }
}
