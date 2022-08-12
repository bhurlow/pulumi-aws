// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appautoscaling.outputs;

import com.pulumi.aws.appautoscaling.outputs.PolicyStepScalingPolicyConfigurationStepAdjustment;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PolicyStepScalingPolicyConfiguration {
    /**
     * @return Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     * 
     */
    private final @Nullable String adjustmentType;
    /**
     * @return The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     * 
     */
    private final @Nullable Integer cooldown;
    /**
     * @return The aggregation type for the policy&#39;s metrics. Valid values are &#34;Minimum&#34;, &#34;Maximum&#34;, and &#34;Average&#34;. Without a value, AWS will treat the aggregation type as &#34;Average&#34;.
     * 
     */
    private final @Nullable String metricAggregationType;
    /**
     * @return The minimum number to adjust your scalable dimension as a result of a scaling activity. If the adjustment type is PercentChangeInCapacity, the scaling policy changes the scalable dimension of the scalable target by this amount.
     * 
     */
    private final @Nullable Integer minAdjustmentMagnitude;
    /**
     * @return A set of adjustments that manage scaling. These have the following structure:
     * 
     */
    private final @Nullable List<PolicyStepScalingPolicyConfigurationStepAdjustment> stepAdjustments;

    @CustomType.Constructor
    private PolicyStepScalingPolicyConfiguration(
        @CustomType.Parameter("adjustmentType") @Nullable String adjustmentType,
        @CustomType.Parameter("cooldown") @Nullable Integer cooldown,
        @CustomType.Parameter("metricAggregationType") @Nullable String metricAggregationType,
        @CustomType.Parameter("minAdjustmentMagnitude") @Nullable Integer minAdjustmentMagnitude,
        @CustomType.Parameter("stepAdjustments") @Nullable List<PolicyStepScalingPolicyConfigurationStepAdjustment> stepAdjustments) {
        this.adjustmentType = adjustmentType;
        this.cooldown = cooldown;
        this.metricAggregationType = metricAggregationType;
        this.minAdjustmentMagnitude = minAdjustmentMagnitude;
        this.stepAdjustments = stepAdjustments;
    }

    /**
     * @return Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     * 
     */
    public Optional<String> adjustmentType() {
        return Optional.ofNullable(this.adjustmentType);
    }
    /**
     * @return The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     * 
     */
    public Optional<Integer> cooldown() {
        return Optional.ofNullable(this.cooldown);
    }
    /**
     * @return The aggregation type for the policy&#39;s metrics. Valid values are &#34;Minimum&#34;, &#34;Maximum&#34;, and &#34;Average&#34;. Without a value, AWS will treat the aggregation type as &#34;Average&#34;.
     * 
     */
    public Optional<String> metricAggregationType() {
        return Optional.ofNullable(this.metricAggregationType);
    }
    /**
     * @return The minimum number to adjust your scalable dimension as a result of a scaling activity. If the adjustment type is PercentChangeInCapacity, the scaling policy changes the scalable dimension of the scalable target by this amount.
     * 
     */
    public Optional<Integer> minAdjustmentMagnitude() {
        return Optional.ofNullable(this.minAdjustmentMagnitude);
    }
    /**
     * @return A set of adjustments that manage scaling. These have the following structure:
     * 
     */
    public List<PolicyStepScalingPolicyConfigurationStepAdjustment> stepAdjustments() {
        return this.stepAdjustments == null ? List.of() : this.stepAdjustments;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicyStepScalingPolicyConfiguration defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String adjustmentType;
        private @Nullable Integer cooldown;
        private @Nullable String metricAggregationType;
        private @Nullable Integer minAdjustmentMagnitude;
        private @Nullable List<PolicyStepScalingPolicyConfigurationStepAdjustment> stepAdjustments;

        public Builder() {
    	      // Empty
        }

        public Builder(PolicyStepScalingPolicyConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adjustmentType = defaults.adjustmentType;
    	      this.cooldown = defaults.cooldown;
    	      this.metricAggregationType = defaults.metricAggregationType;
    	      this.minAdjustmentMagnitude = defaults.minAdjustmentMagnitude;
    	      this.stepAdjustments = defaults.stepAdjustments;
        }

        public Builder adjustmentType(@Nullable String adjustmentType) {
            this.adjustmentType = adjustmentType;
            return this;
        }
        public Builder cooldown(@Nullable Integer cooldown) {
            this.cooldown = cooldown;
            return this;
        }
        public Builder metricAggregationType(@Nullable String metricAggregationType) {
            this.metricAggregationType = metricAggregationType;
            return this;
        }
        public Builder minAdjustmentMagnitude(@Nullable Integer minAdjustmentMagnitude) {
            this.minAdjustmentMagnitude = minAdjustmentMagnitude;
            return this;
        }
        public Builder stepAdjustments(@Nullable List<PolicyStepScalingPolicyConfigurationStepAdjustment> stepAdjustments) {
            this.stepAdjustments = stepAdjustments;
            return this;
        }
        public Builder stepAdjustments(PolicyStepScalingPolicyConfigurationStepAdjustment... stepAdjustments) {
            return stepAdjustments(List.of(stepAdjustments));
        }        public PolicyStepScalingPolicyConfiguration build() {
            return new PolicyStepScalingPolicyConfiguration(adjustmentType, cooldown, metricAggregationType, minAdjustmentMagnitude, stepAdjustments);
        }
    }
}