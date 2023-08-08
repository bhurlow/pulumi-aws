// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a resource to manage EC2 Fleets.
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
    ///     var example = new Aws.Ec2.Fleet("example", new()
    ///     {
    ///         LaunchTemplateConfigs = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.FleetLaunchTemplateConfigArgs
    ///             {
    ///                 LaunchTemplateSpecification = new Aws.Ec2.Inputs.FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs
    ///                 {
    ///                     LaunchTemplateId = aws_launch_template.Example.Id,
    ///                     Version = aws_launch_template.Example.Latest_version,
    ///                 },
    ///             },
    ///         },
    ///         TargetCapacitySpecification = new Aws.Ec2.Inputs.FleetTargetCapacitySpecificationArgs
    ///         {
    ///             DefaultTargetCapacityType = "spot",
    ///             TotalTargetCapacity = 5,
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
    ///  to = aws_ec2_fleet.example
    /// 
    ///  id = "fleet-b9b55d27-c5fc-41ac-a6f3-48fcc91f080c" } Using `pulumi import`, import `aws_ec2_fleet` using the Fleet identifier. For exampleconsole % pulumi import aws_ec2_fleet.example fleet-b9b55d27-c5fc-41ac-a6f3-48fcc91f080c
    /// </summary>
    [AwsResourceType("aws:ec2/fleet:Fleet")]
    public partial class Fleet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the fleet
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Reserved.
        /// </summary>
        [Output("context")]
        public Output<string?> Context { get; private set; } = null!;

        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
        /// </summary>
        [Output("excessCapacityTerminationPolicy")]
        public Output<string?> ExcessCapacityTerminationPolicy { get; private set; } = null!;

        /// <summary>
        /// Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
        /// </summary>
        [Output("fleetInstanceSets")]
        public Output<ImmutableArray<Outputs.FleetFleetInstanceSet>> FleetInstanceSets { get; private set; } = null!;

        /// <summary>
        /// The state of the EC2 Fleet.
        /// </summary>
        [Output("fleetState")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The number of units fulfilled by this request compared to the set target capacity.
        /// </summary>
        [Output("fulfilledCapacity")]
        public Output<double> FulfilledCapacity { get; private set; } = null!;

        /// <summary>
        /// The number of units fulfilled by this request compared to the set target On-Demand capacity.
        /// </summary>
        [Output("fulfilledOnDemandCapacity")]
        public Output<double> FulfilledOnDemandCapacity { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        [Output("launchTemplateConfigs")]
        public Output<ImmutableArray<Outputs.FleetLaunchTemplateConfig>> LaunchTemplateConfigs { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Output("onDemandOptions")]
        public Output<Outputs.FleetOnDemandOptions?> OnDemandOptions { get; private set; } = null!;

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
        /// </summary>
        [Output("replaceUnhealthyInstances")]
        public Output<bool?> ReplaceUnhealthyInstances { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Output("spotOptions")]
        public Output<Outputs.FleetSpotOptions?> SpotOptions { get; private set; } = null!;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Output("targetCapacitySpecification")]
        public Output<Outputs.FleetTargetCapacitySpecification> TargetCapacitySpecification { get; private set; } = null!;

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Output("terminateInstances")]
        public Output<bool?> TerminateInstances { get; private set; } = null!;

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Output("terminateInstancesWithExpiration")]
        public Output<bool?> TerminateInstancesWithExpiration { get; private set; } = null!;

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        /// </summary>
        [Output("validFrom")]
        public Output<string?> ValidFrom { get; private set; } = null!;

        /// <summary>
        /// The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
        /// </summary>
        [Output("validUntil")]
        public Output<string?> ValidUntil { get; private set; } = null!;


        /// <summary>
        /// Create a Fleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fleet(string name, FleetArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/fleet:Fleet", name, args ?? new FleetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fleet(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/fleet:Fleet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fleet Get(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
        {
            return new Fleet(name, id, state, options);
        }
    }

    public sealed class FleetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reserved.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
        /// </summary>
        [Input("excessCapacityTerminationPolicy")]
        public Input<string>? ExcessCapacityTerminationPolicy { get; set; }

        [Input("fleetInstanceSets")]
        private InputList<Inputs.FleetFleetInstanceSetArgs>? _fleetInstanceSets;

        /// <summary>
        /// Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
        /// </summary>
        public InputList<Inputs.FleetFleetInstanceSetArgs> FleetInstanceSets
        {
            get => _fleetInstanceSets ?? (_fleetInstanceSets = new InputList<Inputs.FleetFleetInstanceSetArgs>());
            set => _fleetInstanceSets = value;
        }

        /// <summary>
        /// The state of the EC2 Fleet.
        /// </summary>
        [Input("fleetState")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The number of units fulfilled by this request compared to the set target capacity.
        /// </summary>
        [Input("fulfilledCapacity")]
        public Input<double>? FulfilledCapacity { get; set; }

        /// <summary>
        /// The number of units fulfilled by this request compared to the set target On-Demand capacity.
        /// </summary>
        [Input("fulfilledOnDemandCapacity")]
        public Input<double>? FulfilledOnDemandCapacity { get; set; }

        [Input("launchTemplateConfigs", required: true)]
        private InputList<Inputs.FleetLaunchTemplateConfigArgs>? _launchTemplateConfigs;

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        public InputList<Inputs.FleetLaunchTemplateConfigArgs> LaunchTemplateConfigs
        {
            get => _launchTemplateConfigs ?? (_launchTemplateConfigs = new InputList<Inputs.FleetLaunchTemplateConfigArgs>());
            set => _launchTemplateConfigs = value;
        }

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Input("onDemandOptions")]
        public Input<Inputs.FleetOnDemandOptionsArgs>? OnDemandOptions { get; set; }

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
        /// </summary>
        [Input("replaceUnhealthyInstances")]
        public Input<bool>? ReplaceUnhealthyInstances { get; set; }

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Input("spotOptions")]
        public Input<Inputs.FleetSpotOptionsArgs>? SpotOptions { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Input("targetCapacitySpecification", required: true)]
        public Input<Inputs.FleetTargetCapacitySpecificationArgs> TargetCapacitySpecification { get; set; } = null!;

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Input("terminateInstances")]
        public Input<bool>? TerminateInstances { get; set; }

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Input("terminateInstancesWithExpiration")]
        public Input<bool>? TerminateInstancesWithExpiration { get; set; }

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        /// </summary>
        [Input("validFrom")]
        public Input<string>? ValidFrom { get; set; }

        /// <summary>
        /// The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public FleetArgs()
        {
        }
        public static new FleetArgs Empty => new FleetArgs();
    }

    public sealed class FleetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the fleet
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Reserved.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
        /// </summary>
        [Input("excessCapacityTerminationPolicy")]
        public Input<string>? ExcessCapacityTerminationPolicy { get; set; }

        [Input("fleetInstanceSets")]
        private InputList<Inputs.FleetFleetInstanceSetGetArgs>? _fleetInstanceSets;

        /// <summary>
        /// Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
        /// </summary>
        public InputList<Inputs.FleetFleetInstanceSetGetArgs> FleetInstanceSets
        {
            get => _fleetInstanceSets ?? (_fleetInstanceSets = new InputList<Inputs.FleetFleetInstanceSetGetArgs>());
            set => _fleetInstanceSets = value;
        }

        /// <summary>
        /// The state of the EC2 Fleet.
        /// </summary>
        [Input("fleetState")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The number of units fulfilled by this request compared to the set target capacity.
        /// </summary>
        [Input("fulfilledCapacity")]
        public Input<double>? FulfilledCapacity { get; set; }

        /// <summary>
        /// The number of units fulfilled by this request compared to the set target On-Demand capacity.
        /// </summary>
        [Input("fulfilledOnDemandCapacity")]
        public Input<double>? FulfilledOnDemandCapacity { get; set; }

        [Input("launchTemplateConfigs")]
        private InputList<Inputs.FleetLaunchTemplateConfigGetArgs>? _launchTemplateConfigs;

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        public InputList<Inputs.FleetLaunchTemplateConfigGetArgs> LaunchTemplateConfigs
        {
            get => _launchTemplateConfigs ?? (_launchTemplateConfigs = new InputList<Inputs.FleetLaunchTemplateConfigGetArgs>());
            set => _launchTemplateConfigs = value;
        }

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Input("onDemandOptions")]
        public Input<Inputs.FleetOnDemandOptionsGetArgs>? OnDemandOptions { get; set; }

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
        /// </summary>
        [Input("replaceUnhealthyInstances")]
        public Input<bool>? ReplaceUnhealthyInstances { get; set; }

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Input("spotOptions")]
        public Input<Inputs.FleetSpotOptionsGetArgs>? SpotOptions { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Input("targetCapacitySpecification")]
        public Input<Inputs.FleetTargetCapacitySpecificationGetArgs>? TargetCapacitySpecification { get; set; }

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Input("terminateInstances")]
        public Input<bool>? TerminateInstances { get; set; }

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Input("terminateInstancesWithExpiration")]
        public Input<bool>? TerminateInstancesWithExpiration { get; set; }

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        /// </summary>
        [Input("validFrom")]
        public Input<string>? ValidFrom { get; set; }

        /// <summary>
        /// The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public FleetState()
        {
        }
        public static new FleetState Empty => new FleetState();
    }
}
