// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling
{
    /// <summary>
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
    ///         var bar = new Aws.AutoScaling.Group("bar", new Aws.AutoScaling.GroupArgs
    ///         {
    ///             AvailabilityZones = 
    ///             {
    ///                 "us-east-1a",
    ///             },
    ///             MaxSize = 5,
    ///             MinSize = 2,
    ///             HealthCheckGracePeriod = 300,
    ///             HealthCheckType = "ELB",
    ///             ForceDelete = true,
    ///             LaunchConfiguration = aws_launch_configuration.Foo.Name,
    ///         });
    ///         var bat = new Aws.AutoScaling.Policy("bat", new Aws.AutoScaling.PolicyArgs
    ///         {
    ///             ScalingAdjustment = 4,
    ///             AdjustmentType = "ChangeInCapacity",
    ///             Cooldown = 300,
    ///             AutoscalingGroupName = bar.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Create predictive scaling policy using customized metrics
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.AutoScaling.Policy("example", new Aws.AutoScaling.PolicyArgs
    ///         {
    ///             AutoscalingGroupName = "my-test-asg",
    ///             PolicyType = "PredictiveScaling",
    ///             PredictiveScalingConfiguration = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationArgs
    ///             {
    ///                 MetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs
    ///                 {
    ///                     CustomizedCapacityMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs
    ///                     {
    ///                         MetricDataQueries = 
    ///                         {
    ///                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs
    ///                             {
    ///                                 Expression = "SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
    ///                                 Id = "capacity_sum",
    ///                                 ReturnData = false,
    ///                             },
    ///                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs
    ///                             {
    ///                                 Expression = "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 300))",
    ///                                 Id = "load_sum",
    ///                                 ReturnData = false,
    ///                             },
    ///                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs
    ///                             {
    ///                                 Expression = "load_sum / capacity_sum",
    ///                                 Id = "weighted_average",
    ///                             },
    ///                         },
    ///                     },
    ///                     CustomizedLoadMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs
    ///                     {
    ///                         MetricDataQueries = 
    ///                         {
    ///                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryArgs
    ///                             {
    ///                                 Expression = "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 3600))",
    ///                                 Id = "load_sum",
    ///                             },
    ///                         },
    ///                     },
    ///                     CustomizedScalingMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs
    ///                     {
    ///                         MetricDataQueries = 
    ///                         {
    ///                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs
    ///                             {
    ///                                 Id = "scaling",
    ///                                 MetricStat = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatArgs
    ///                                 {
    ///                                     Metric = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricArgs
    ///                                     {
    ///                                         Dimensions = 
    ///                                         {
    ///                                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArgs
    ///                                             {
    ///                                                 Name = "AutoScalingGroupName",
    ///                                                 Value = "my-test-asg",
    ///                                             },
    ///                                         },
    ///                                         MetricName = "CPUUtilization",
    ///                                         Namespace = "AWS/EC2",
    ///                                     },
    ///                                     Stat = "Average",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     TargetValue = 10,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Create predictive scaling policy using customized scaling and predefined load metric
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.AutoScaling.Policy("example", new Aws.AutoScaling.PolicyArgs
    ///         {
    ///             AutoscalingGroupName = "my-test-asg",
    ///             PolicyType = "PredictiveScaling",
    ///             PredictiveScalingConfiguration = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationArgs
    ///             {
    ///                 MetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs
    ///                 {
    ///                     CustomizedScalingMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs
    ///                     {
    ///                         MetricDataQueries = 
    ///                         {
    ///                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs
    ///                             {
    ///                                 Id = "scaling",
    ///                                 MetricStat = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatArgs
    ///                                 {
    ///                                     Metric = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricArgs
    ///                                     {
    ///                                         Dimensions = 
    ///                                         {
    ///                                             new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArgs
    ///                                             {
    ///                                                 Name = "AutoScalingGroupName",
    ///                                                 Value = "my-test-asg",
    ///                                             },
    ///                                         },
    ///                                         MetricName = "CPUUtilization",
    ///                                         Namespace = "AWS/EC2",
    ///                                     },
    ///                                     Stat = "Average",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     PredefinedLoadMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs
    ///                     {
    ///                         PredefinedMetricType = "ASGTotalCPUUtilization",
    ///                         ResourceLabel = "testLabel",
    ///                     },
    ///                     TargetValue = 10,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// AutoScaling scaling policy can be imported using the role autoscaling_group_name and name separated by `/`.
    /// 
    /// ```sh
    ///  $ pulumi import aws:autoscaling/policy:Policy test-policy asg-name/policy-name
    /// ```
    /// </summary>
    [AwsResourceType("aws:autoscaling/policy:Policy")]
    public partial class Policy : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
        /// </summary>
        [Output("adjustmentType")]
        public Output<string?> AdjustmentType { get; private set; } = null!;

        /// <summary>
        /// The ARN assigned by AWS to the scaling policy.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the autoscaling group.
        /// </summary>
        [Output("autoscalingGroupName")]
        public Output<string> AutoscalingGroupName { get; private set; } = null!;

        /// <summary>
        /// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
        /// </summary>
        [Output("cooldown")]
        public Output<int?> Cooldown { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the scaling policy is enabled or disabled. Default: `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
        /// </summary>
        [Output("estimatedInstanceWarmup")]
        public Output<int?> EstimatedInstanceWarmup { get; private set; } = null!;

        /// <summary>
        /// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
        /// </summary>
        [Output("metricAggregationType")]
        public Output<string> MetricAggregationType { get; private set; } = null!;

        /// <summary>
        /// Minimum value to scale by when `adjustment_type` is set to `PercentChangeInCapacity`.
        /// </summary>
        [Output("minAdjustmentMagnitude")]
        public Output<int?> MinAdjustmentMagnitude { get; private set; } = null!;

        /// <summary>
        /// The name of the dimension.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
        /// </summary>
        [Output("policyType")]
        public Output<string?> PolicyType { get; private set; } = null!;

        /// <summary>
        /// The predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
        /// </summary>
        [Output("predictiveScalingConfiguration")]
        public Output<Outputs.PolicyPredictiveScalingConfiguration?> PredictiveScalingConfiguration { get; private set; } = null!;

        /// <summary>
        /// The number of members by which to
        /// scale, when the adjustment bounds are breached. A positive value scales
        /// up. A negative value scales down.
        /// </summary>
        [Output("scalingAdjustment")]
        public Output<int?> ScalingAdjustment { get; private set; } = null!;

        /// <summary>
        /// A set of adjustments that manage
        /// group scaling. These have the following structure:
        /// </summary>
        [Output("stepAdjustments")]
        public Output<ImmutableArray<Outputs.PolicyStepAdjustment>> StepAdjustments { get; private set; } = null!;

        /// <summary>
        /// A target tracking policy. These have the following structure:
        /// </summary>
        [Output("targetTrackingConfiguration")]
        public Output<Outputs.PolicyTargetTrackingConfiguration?> TargetTrackingConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:autoscaling/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:autoscaling/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
        /// </summary>
        [Input("adjustmentType")]
        public Input<string>? AdjustmentType { get; set; }

        /// <summary>
        /// The name of the autoscaling group.
        /// </summary>
        [Input("autoscalingGroupName", required: true)]
        public Input<string> AutoscalingGroupName { get; set; } = null!;

        /// <summary>
        /// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
        /// </summary>
        [Input("cooldown")]
        public Input<int>? Cooldown { get; set; }

        /// <summary>
        /// Indicates whether the scaling policy is enabled or disabled. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
        /// </summary>
        [Input("estimatedInstanceWarmup")]
        public Input<int>? EstimatedInstanceWarmup { get; set; }

        /// <summary>
        /// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
        /// </summary>
        [Input("metricAggregationType")]
        public Input<string>? MetricAggregationType { get; set; }

        /// <summary>
        /// Minimum value to scale by when `adjustment_type` is set to `PercentChangeInCapacity`.
        /// </summary>
        [Input("minAdjustmentMagnitude")]
        public Input<int>? MinAdjustmentMagnitude { get; set; }

        /// <summary>
        /// The name of the dimension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// The predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
        /// </summary>
        [Input("predictiveScalingConfiguration")]
        public Input<Inputs.PolicyPredictiveScalingConfigurationArgs>? PredictiveScalingConfiguration { get; set; }

        /// <summary>
        /// The number of members by which to
        /// scale, when the adjustment bounds are breached. A positive value scales
        /// up. A negative value scales down.
        /// </summary>
        [Input("scalingAdjustment")]
        public Input<int>? ScalingAdjustment { get; set; }

        [Input("stepAdjustments")]
        private InputList<Inputs.PolicyStepAdjustmentArgs>? _stepAdjustments;

        /// <summary>
        /// A set of adjustments that manage
        /// group scaling. These have the following structure:
        /// </summary>
        public InputList<Inputs.PolicyStepAdjustmentArgs> StepAdjustments
        {
            get => _stepAdjustments ?? (_stepAdjustments = new InputList<Inputs.PolicyStepAdjustmentArgs>());
            set => _stepAdjustments = value;
        }

        /// <summary>
        /// A target tracking policy. These have the following structure:
        /// </summary>
        [Input("targetTrackingConfiguration")]
        public Input<Inputs.PolicyTargetTrackingConfigurationArgs>? TargetTrackingConfiguration { get; set; }

        public PolicyArgs()
        {
        }
    }

    public sealed class PolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
        /// </summary>
        [Input("adjustmentType")]
        public Input<string>? AdjustmentType { get; set; }

        /// <summary>
        /// The ARN assigned by AWS to the scaling policy.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the autoscaling group.
        /// </summary>
        [Input("autoscalingGroupName")]
        public Input<string>? AutoscalingGroupName { get; set; }

        /// <summary>
        /// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
        /// </summary>
        [Input("cooldown")]
        public Input<int>? Cooldown { get; set; }

        /// <summary>
        /// Indicates whether the scaling policy is enabled or disabled. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
        /// </summary>
        [Input("estimatedInstanceWarmup")]
        public Input<int>? EstimatedInstanceWarmup { get; set; }

        /// <summary>
        /// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
        /// </summary>
        [Input("metricAggregationType")]
        public Input<string>? MetricAggregationType { get; set; }

        /// <summary>
        /// Minimum value to scale by when `adjustment_type` is set to `PercentChangeInCapacity`.
        /// </summary>
        [Input("minAdjustmentMagnitude")]
        public Input<int>? MinAdjustmentMagnitude { get; set; }

        /// <summary>
        /// The name of the dimension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// The predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
        /// </summary>
        [Input("predictiveScalingConfiguration")]
        public Input<Inputs.PolicyPredictiveScalingConfigurationGetArgs>? PredictiveScalingConfiguration { get; set; }

        /// <summary>
        /// The number of members by which to
        /// scale, when the adjustment bounds are breached. A positive value scales
        /// up. A negative value scales down.
        /// </summary>
        [Input("scalingAdjustment")]
        public Input<int>? ScalingAdjustment { get; set; }

        [Input("stepAdjustments")]
        private InputList<Inputs.PolicyStepAdjustmentGetArgs>? _stepAdjustments;

        /// <summary>
        /// A set of adjustments that manage
        /// group scaling. These have the following structure:
        /// </summary>
        public InputList<Inputs.PolicyStepAdjustmentGetArgs> StepAdjustments
        {
            get => _stepAdjustments ?? (_stepAdjustments = new InputList<Inputs.PolicyStepAdjustmentGetArgs>());
            set => _stepAdjustments = value;
        }

        /// <summary>
        /// A target tracking policy. These have the following structure:
        /// </summary>
        [Input("targetTrackingConfiguration")]
        public Input<Inputs.PolicyTargetTrackingConfigurationGetArgs>? TargetTrackingConfiguration { get; set; }

        public PolicyState()
        {
        }
    }
}
