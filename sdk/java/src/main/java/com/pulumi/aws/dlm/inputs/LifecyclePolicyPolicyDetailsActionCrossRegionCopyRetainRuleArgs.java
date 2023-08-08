// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dlm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs Empty = new LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs();

    /**
     * How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
     * 
     */
    @Import(name="interval", required=true)
    private Output<Integer> interval;

    /**
     * @return How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
     * 
     */
    public Output<Integer> interval() {
        return this.interval;
    }

    /**
     * The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
     * 
     */
    @Import(name="intervalUnit", required=true)
    private Output<String> intervalUnit;

    /**
     * @return The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
     * 
     */
    public Output<String> intervalUnit() {
        return this.intervalUnit;
    }

    private LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs() {}

    private LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs(LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs $) {
        this.interval = $.interval;
        this.intervalUnit = $.intervalUnit;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs $;

        public Builder() {
            $ = new LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs();
        }

        public Builder(LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs defaults) {
            $ = new LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param interval How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
         * 
         * @return builder
         * 
         */
        public Builder interval(Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param intervalUnit The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
         * 
         * @return builder
         * 
         */
        public Builder intervalUnit(Output<String> intervalUnit) {
            $.intervalUnit = intervalUnit;
            return this;
        }

        /**
         * @param intervalUnit The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
         * 
         * @return builder
         * 
         */
        public Builder intervalUnit(String intervalUnit) {
            return intervalUnit(Output.of(intervalUnit));
        }

        public LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs build() {
            $.interval = Objects.requireNonNull($.interval, "expected parameter 'interval' to be non-null");
            $.intervalUnit = Objects.requireNonNull($.intervalUnit, "expected parameter 'intervalUnit' to be non-null");
            return $;
        }
    }

}
