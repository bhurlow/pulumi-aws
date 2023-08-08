// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dlm.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LifecyclePolicyPolicyDetailsScheduleCreateRule {
    /**
     * @return The schedule, as a Cron expression. The schedule interval must be between 1 hour and 1 year. Conflicts with `interval`, `interval_unit`, and `times`.
     * 
     */
    private @Nullable String cronExpression;
    /**
     * @return How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
     * 
     */
    private @Nullable Integer interval;
    /**
     * @return The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
     * 
     */
    private @Nullable String intervalUnit;
    /**
     * @return Specifies the destination for snapshots created by the policy. To create snapshots in the same Region as the source resource, specify `CLOUD`. To create snapshots on the same Outpost as the source resource, specify `OUTPOST_LOCAL`. If you omit this parameter, `CLOUD` is used by default. If the policy targets resources in an AWS Region, then you must create snapshots in the same Region as the source resource. If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost as the source resource, or in the Region of that Outpost. Valid values are `CLOUD` and `OUTPOST_LOCAL`.
     * 
     */
    private @Nullable String location;
    /**
     * @return A list of times in 24 hour clock format that sets when the lifecycle policy should be evaluated. Max of 1. Conflicts with `cron_expression`. Must be set if `interval` is set.
     * 
     */
    private @Nullable String times;

    private LifecyclePolicyPolicyDetailsScheduleCreateRule() {}
    /**
     * @return The schedule, as a Cron expression. The schedule interval must be between 1 hour and 1 year. Conflicts with `interval`, `interval_unit`, and `times`.
     * 
     */
    public Optional<String> cronExpression() {
        return Optional.ofNullable(this.cronExpression);
    }
    /**
     * @return How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
     * 
     */
    public Optional<Integer> interval() {
        return Optional.ofNullable(this.interval);
    }
    /**
     * @return The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
     * 
     */
    public Optional<String> intervalUnit() {
        return Optional.ofNullable(this.intervalUnit);
    }
    /**
     * @return Specifies the destination for snapshots created by the policy. To create snapshots in the same Region as the source resource, specify `CLOUD`. To create snapshots on the same Outpost as the source resource, specify `OUTPOST_LOCAL`. If you omit this parameter, `CLOUD` is used by default. If the policy targets resources in an AWS Region, then you must create snapshots in the same Region as the source resource. If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost as the source resource, or in the Region of that Outpost. Valid values are `CLOUD` and `OUTPOST_LOCAL`.
     * 
     */
    public Optional<String> location() {
        return Optional.ofNullable(this.location);
    }
    /**
     * @return A list of times in 24 hour clock format that sets when the lifecycle policy should be evaluated. Max of 1. Conflicts with `cron_expression`. Must be set if `interval` is set.
     * 
     */
    public Optional<String> times() {
        return Optional.ofNullable(this.times);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LifecyclePolicyPolicyDetailsScheduleCreateRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cronExpression;
        private @Nullable Integer interval;
        private @Nullable String intervalUnit;
        private @Nullable String location;
        private @Nullable String times;
        public Builder() {}
        public Builder(LifecyclePolicyPolicyDetailsScheduleCreateRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cronExpression = defaults.cronExpression;
    	      this.interval = defaults.interval;
    	      this.intervalUnit = defaults.intervalUnit;
    	      this.location = defaults.location;
    	      this.times = defaults.times;
        }

        @CustomType.Setter
        public Builder cronExpression(@Nullable String cronExpression) {
            this.cronExpression = cronExpression;
            return this;
        }
        @CustomType.Setter
        public Builder interval(@Nullable Integer interval) {
            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder intervalUnit(@Nullable String intervalUnit) {
            this.intervalUnit = intervalUnit;
            return this;
        }
        @CustomType.Setter
        public Builder location(@Nullable String location) {
            this.location = location;
            return this;
        }
        @CustomType.Setter
        public Builder times(@Nullable String times) {
            this.times = times;
            return this;
        }
        public LifecyclePolicyPolicyDetailsScheduleCreateRule build() {
            final var o = new LifecyclePolicyPolicyDetailsScheduleCreateRule();
            o.cronExpression = cronExpression;
            o.interval = interval;
            o.intervalUnit = intervalUnit;
            o.location = location;
            o.times = times;
            return o;
        }
    }
}
