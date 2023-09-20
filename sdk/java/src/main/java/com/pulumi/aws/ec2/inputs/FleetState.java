// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.inputs;

import com.pulumi.aws.ec2.inputs.FleetFleetInstanceSetArgs;
import com.pulumi.aws.ec2.inputs.FleetLaunchTemplateConfigArgs;
import com.pulumi.aws.ec2.inputs.FleetOnDemandOptionsArgs;
import com.pulumi.aws.ec2.inputs.FleetSpotOptionsArgs;
import com.pulumi.aws.ec2.inputs.FleetTargetCapacitySpecificationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FleetState extends com.pulumi.resources.ResourceArgs {

    public static final FleetState Empty = new FleetState();

    /**
     * The ARN of the fleet
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the fleet
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Reserved.
     * 
     */
    @Import(name="context")
    private @Nullable Output<String> context;

    /**
     * @return Reserved.
     * 
     */
    public Optional<Output<String>> context() {
        return Optional.ofNullable(this.context);
    }

    /**
     * Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
     * 
     */
    @Import(name="excessCapacityTerminationPolicy")
    private @Nullable Output<String> excessCapacityTerminationPolicy;

    /**
     * @return Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
     * 
     */
    public Optional<Output<String>> excessCapacityTerminationPolicy() {
        return Optional.ofNullable(this.excessCapacityTerminationPolicy);
    }

    /**
     * Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
     * 
     */
    @Import(name="fleetInstanceSets")
    private @Nullable Output<List<FleetFleetInstanceSetArgs>> fleetInstanceSets;

    /**
     * @return Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
     * 
     */
    public Optional<Output<List<FleetFleetInstanceSetArgs>>> fleetInstanceSets() {
        return Optional.ofNullable(this.fleetInstanceSets);
    }

    /**
     * The state of the EC2 Fleet.
     * 
     */
    @Import(name="fleetState")
    private @Nullable Output<String> fleetState;

    /**
     * @return The state of the EC2 Fleet.
     * 
     */
    public Optional<Output<String>> fleetState() {
        return Optional.ofNullable(this.fleetState);
    }

    /**
     * The number of units fulfilled by this request compared to the set target capacity.
     * 
     */
    @Import(name="fulfilledCapacity")
    private @Nullable Output<Double> fulfilledCapacity;

    /**
     * @return The number of units fulfilled by this request compared to the set target capacity.
     * 
     */
    public Optional<Output<Double>> fulfilledCapacity() {
        return Optional.ofNullable(this.fulfilledCapacity);
    }

    /**
     * The number of units fulfilled by this request compared to the set target On-Demand capacity.
     * 
     */
    @Import(name="fulfilledOnDemandCapacity")
    private @Nullable Output<Double> fulfilledOnDemandCapacity;

    /**
     * @return The number of units fulfilled by this request compared to the set target On-Demand capacity.
     * 
     */
    public Optional<Output<Double>> fulfilledOnDemandCapacity() {
        return Optional.ofNullable(this.fulfilledOnDemandCapacity);
    }

    /**
     * Nested argument containing EC2 Launch Template configurations. Defined below.
     * 
     */
    @Import(name="launchTemplateConfigs")
    private @Nullable Output<List<FleetLaunchTemplateConfigArgs>> launchTemplateConfigs;

    /**
     * @return Nested argument containing EC2 Launch Template configurations. Defined below.
     * 
     */
    public Optional<Output<List<FleetLaunchTemplateConfigArgs>>> launchTemplateConfigs() {
        return Optional.ofNullable(this.launchTemplateConfigs);
    }

    /**
     * Nested argument containing On-Demand configurations. Defined below.
     * 
     */
    @Import(name="onDemandOptions")
    private @Nullable Output<FleetOnDemandOptionsArgs> onDemandOptions;

    /**
     * @return Nested argument containing On-Demand configurations. Defined below.
     * 
     */
    public Optional<Output<FleetOnDemandOptionsArgs>> onDemandOptions() {
        return Optional.ofNullable(this.onDemandOptions);
    }

    /**
     * Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
     * 
     */
    @Import(name="replaceUnhealthyInstances")
    private @Nullable Output<Boolean> replaceUnhealthyInstances;

    /**
     * @return Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
     * 
     */
    public Optional<Output<Boolean>> replaceUnhealthyInstances() {
        return Optional.ofNullable(this.replaceUnhealthyInstances);
    }

    /**
     * Nested argument containing Spot configurations. Defined below.
     * 
     */
    @Import(name="spotOptions")
    private @Nullable Output<FleetSpotOptionsArgs> spotOptions;

    /**
     * @return Nested argument containing Spot configurations. Defined below.
     * 
     */
    public Optional<Output<FleetSpotOptionsArgs>> spotOptions() {
        return Optional.ofNullable(this.spotOptions);
    }

    /**
     * Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Nested argument containing target capacity configurations. Defined below.
     * 
     */
    @Import(name="targetCapacitySpecification")
    private @Nullable Output<FleetTargetCapacitySpecificationArgs> targetCapacitySpecification;

    /**
     * @return Nested argument containing target capacity configurations. Defined below.
     * 
     */
    public Optional<Output<FleetTargetCapacitySpecificationArgs>> targetCapacitySpecification() {
        return Optional.ofNullable(this.targetCapacitySpecification);
    }

    /**
     * Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
     * 
     */
    @Import(name="terminateInstances")
    private @Nullable Output<Boolean> terminateInstances;

    /**
     * @return Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> terminateInstances() {
        return Optional.ofNullable(this.terminateInstances);
    }

    /**
     * Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
     * 
     */
    @Import(name="terminateInstancesWithExpiration")
    private @Nullable Output<Boolean> terminateInstancesWithExpiration;

    /**
     * @return Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> terminateInstancesWithExpiration() {
        return Optional.ofNullable(this.terminateInstancesWithExpiration);
    }

    /**
     * The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
     * 
     */
    @Import(name="validFrom")
    private @Nullable Output<String> validFrom;

    /**
     * @return The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
     * 
     */
    public Optional<Output<String>> validFrom() {
        return Optional.ofNullable(this.validFrom);
    }

    /**
     * The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
     * 
     */
    @Import(name="validUntil")
    private @Nullable Output<String> validUntil;

    /**
     * @return The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
     * 
     */
    public Optional<Output<String>> validUntil() {
        return Optional.ofNullable(this.validUntil);
    }

    private FleetState() {}

    private FleetState(FleetState $) {
        this.arn = $.arn;
        this.context = $.context;
        this.excessCapacityTerminationPolicy = $.excessCapacityTerminationPolicy;
        this.fleetInstanceSets = $.fleetInstanceSets;
        this.fleetState = $.fleetState;
        this.fulfilledCapacity = $.fulfilledCapacity;
        this.fulfilledOnDemandCapacity = $.fulfilledOnDemandCapacity;
        this.launchTemplateConfigs = $.launchTemplateConfigs;
        this.onDemandOptions = $.onDemandOptions;
        this.replaceUnhealthyInstances = $.replaceUnhealthyInstances;
        this.spotOptions = $.spotOptions;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.targetCapacitySpecification = $.targetCapacitySpecification;
        this.terminateInstances = $.terminateInstances;
        this.terminateInstancesWithExpiration = $.terminateInstancesWithExpiration;
        this.type = $.type;
        this.validFrom = $.validFrom;
        this.validUntil = $.validUntil;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FleetState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FleetState $;

        public Builder() {
            $ = new FleetState();
        }

        public Builder(FleetState defaults) {
            $ = new FleetState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the fleet
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the fleet
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param context Reserved.
         * 
         * @return builder
         * 
         */
        public Builder context(@Nullable Output<String> context) {
            $.context = context;
            return this;
        }

        /**
         * @param context Reserved.
         * 
         * @return builder
         * 
         */
        public Builder context(String context) {
            return context(Output.of(context));
        }

        /**
         * @param excessCapacityTerminationPolicy Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
         * 
         * @return builder
         * 
         */
        public Builder excessCapacityTerminationPolicy(@Nullable Output<String> excessCapacityTerminationPolicy) {
            $.excessCapacityTerminationPolicy = excessCapacityTerminationPolicy;
            return this;
        }

        /**
         * @param excessCapacityTerminationPolicy Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`. Supported only for fleets of type `maintain`.
         * 
         * @return builder
         * 
         */
        public Builder excessCapacityTerminationPolicy(String excessCapacityTerminationPolicy) {
            return excessCapacityTerminationPolicy(Output.of(excessCapacityTerminationPolicy));
        }

        /**
         * @param fleetInstanceSets Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
         * 
         * @return builder
         * 
         */
        public Builder fleetInstanceSets(@Nullable Output<List<FleetFleetInstanceSetArgs>> fleetInstanceSets) {
            $.fleetInstanceSets = fleetInstanceSets;
            return this;
        }

        /**
         * @param fleetInstanceSets Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
         * 
         * @return builder
         * 
         */
        public Builder fleetInstanceSets(List<FleetFleetInstanceSetArgs> fleetInstanceSets) {
            return fleetInstanceSets(Output.of(fleetInstanceSets));
        }

        /**
         * @param fleetInstanceSets Information about the instances that were launched by the fleet. Available only when `type` is set to `instant`.
         * 
         * @return builder
         * 
         */
        public Builder fleetInstanceSets(FleetFleetInstanceSetArgs... fleetInstanceSets) {
            return fleetInstanceSets(List.of(fleetInstanceSets));
        }

        /**
         * @param fleetState The state of the EC2 Fleet.
         * 
         * @return builder
         * 
         */
        public Builder fleetState(@Nullable Output<String> fleetState) {
            $.fleetState = fleetState;
            return this;
        }

        /**
         * @param fleetState The state of the EC2 Fleet.
         * 
         * @return builder
         * 
         */
        public Builder fleetState(String fleetState) {
            return fleetState(Output.of(fleetState));
        }

        /**
         * @param fulfilledCapacity The number of units fulfilled by this request compared to the set target capacity.
         * 
         * @return builder
         * 
         */
        public Builder fulfilledCapacity(@Nullable Output<Double> fulfilledCapacity) {
            $.fulfilledCapacity = fulfilledCapacity;
            return this;
        }

        /**
         * @param fulfilledCapacity The number of units fulfilled by this request compared to the set target capacity.
         * 
         * @return builder
         * 
         */
        public Builder fulfilledCapacity(Double fulfilledCapacity) {
            return fulfilledCapacity(Output.of(fulfilledCapacity));
        }

        /**
         * @param fulfilledOnDemandCapacity The number of units fulfilled by this request compared to the set target On-Demand capacity.
         * 
         * @return builder
         * 
         */
        public Builder fulfilledOnDemandCapacity(@Nullable Output<Double> fulfilledOnDemandCapacity) {
            $.fulfilledOnDemandCapacity = fulfilledOnDemandCapacity;
            return this;
        }

        /**
         * @param fulfilledOnDemandCapacity The number of units fulfilled by this request compared to the set target On-Demand capacity.
         * 
         * @return builder
         * 
         */
        public Builder fulfilledOnDemandCapacity(Double fulfilledOnDemandCapacity) {
            return fulfilledOnDemandCapacity(Output.of(fulfilledOnDemandCapacity));
        }

        /**
         * @param launchTemplateConfigs Nested argument containing EC2 Launch Template configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateConfigs(@Nullable Output<List<FleetLaunchTemplateConfigArgs>> launchTemplateConfigs) {
            $.launchTemplateConfigs = launchTemplateConfigs;
            return this;
        }

        /**
         * @param launchTemplateConfigs Nested argument containing EC2 Launch Template configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateConfigs(List<FleetLaunchTemplateConfigArgs> launchTemplateConfigs) {
            return launchTemplateConfigs(Output.of(launchTemplateConfigs));
        }

        /**
         * @param launchTemplateConfigs Nested argument containing EC2 Launch Template configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateConfigs(FleetLaunchTemplateConfigArgs... launchTemplateConfigs) {
            return launchTemplateConfigs(List.of(launchTemplateConfigs));
        }

        /**
         * @param onDemandOptions Nested argument containing On-Demand configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder onDemandOptions(@Nullable Output<FleetOnDemandOptionsArgs> onDemandOptions) {
            $.onDemandOptions = onDemandOptions;
            return this;
        }

        /**
         * @param onDemandOptions Nested argument containing On-Demand configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder onDemandOptions(FleetOnDemandOptionsArgs onDemandOptions) {
            return onDemandOptions(Output.of(onDemandOptions));
        }

        /**
         * @param replaceUnhealthyInstances Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
         * 
         * @return builder
         * 
         */
        public Builder replaceUnhealthyInstances(@Nullable Output<Boolean> replaceUnhealthyInstances) {
            $.replaceUnhealthyInstances = replaceUnhealthyInstances;
            return this;
        }

        /**
         * @param replaceUnhealthyInstances Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`. Supported only for fleets of type `maintain`.
         * 
         * @return builder
         * 
         */
        public Builder replaceUnhealthyInstances(Boolean replaceUnhealthyInstances) {
            return replaceUnhealthyInstances(Output.of(replaceUnhealthyInstances));
        }

        /**
         * @param spotOptions Nested argument containing Spot configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder spotOptions(@Nullable Output<FleetSpotOptionsArgs> spotOptions) {
            $.spotOptions = spotOptions;
            return this;
        }

        /**
         * @param spotOptions Nested argument containing Spot configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder spotOptions(FleetSpotOptionsArgs spotOptions) {
            return spotOptions(Output.of(spotOptions));
        }

        /**
         * @param tags Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param targetCapacitySpecification Nested argument containing target capacity configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder targetCapacitySpecification(@Nullable Output<FleetTargetCapacitySpecificationArgs> targetCapacitySpecification) {
            $.targetCapacitySpecification = targetCapacitySpecification;
            return this;
        }

        /**
         * @param targetCapacitySpecification Nested argument containing target capacity configurations. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder targetCapacitySpecification(FleetTargetCapacitySpecificationArgs targetCapacitySpecification) {
            return targetCapacitySpecification(Output.of(targetCapacitySpecification));
        }

        /**
         * @param terminateInstances Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder terminateInstances(@Nullable Output<Boolean> terminateInstances) {
            $.terminateInstances = terminateInstances;
            return this;
        }

        /**
         * @param terminateInstances Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder terminateInstances(Boolean terminateInstances) {
            return terminateInstances(Output.of(terminateInstances));
        }

        /**
         * @param terminateInstancesWithExpiration Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder terminateInstancesWithExpiration(@Nullable Output<Boolean> terminateInstancesWithExpiration) {
            $.terminateInstancesWithExpiration = terminateInstancesWithExpiration;
            return this;
        }

        /**
         * @param terminateInstancesWithExpiration Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder terminateInstancesWithExpiration(Boolean terminateInstancesWithExpiration) {
            return terminateInstancesWithExpiration(Output.of(terminateInstancesWithExpiration));
        }

        /**
         * @param type The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`, `instant`. Defaults to `maintain`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param validFrom The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
         * 
         * @return builder
         * 
         */
        public Builder validFrom(@Nullable Output<String> validFrom) {
            $.validFrom = validFrom;
            return this;
        }

        /**
         * @param validFrom The start date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
         * 
         * @return builder
         * 
         */
        public Builder validFrom(String validFrom) {
            return validFrom(Output.of(validFrom));
        }

        /**
         * @param validUntil The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
         * 
         * @return builder
         * 
         */
        public Builder validUntil(@Nullable Output<String> validUntil) {
            $.validUntil = validUntil;
            return this;
        }

        /**
         * @param validUntil The end date and time of the request, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
         * 
         * @return builder
         * 
         */
        public Builder validUntil(String validUntil) {
            return validUntil(Output.of(validUntil));
        }

        public FleetState build() {
            return $;
        }
    }

}
