// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.inputs;

import com.pulumi.aws.wafv2.inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class WebAclLoggingConfigurationLoggingFilterFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final WebAclLoggingConfigurationLoggingFilterFilterArgs Empty = new WebAclLoggingConfigurationLoggingFilterFilterArgs();

    /**
     * How to handle logs that satisfy the filter&#39;s conditions and requirement. Valid values: `KEEP` or `DROP`.
     * 
     */
    @Import(name="behavior", required=true)
    private Output<String> behavior;

    /**
     * @return How to handle logs that satisfy the filter&#39;s conditions and requirement. Valid values: `KEEP` or `DROP`.
     * 
     */
    public Output<String> behavior() {
        return this.behavior;
    }

    /**
     * Match condition(s) for the filter. See Condition below for more details.
     * 
     */
    @Import(name="conditions", required=true)
    private Output<List<WebAclLoggingConfigurationLoggingFilterFilterConditionArgs>> conditions;

    /**
     * @return Match condition(s) for the filter. See Condition below for more details.
     * 
     */
    public Output<List<WebAclLoggingConfigurationLoggingFilterFilterConditionArgs>> conditions() {
        return this.conditions;
    }

    /**
     * Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition. Valid values: `MEETS_ALL` or `MEETS_ANY`.
     * 
     */
    @Import(name="requirement", required=true)
    private Output<String> requirement;

    /**
     * @return Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition. Valid values: `MEETS_ALL` or `MEETS_ANY`.
     * 
     */
    public Output<String> requirement() {
        return this.requirement;
    }

    private WebAclLoggingConfigurationLoggingFilterFilterArgs() {}

    private WebAclLoggingConfigurationLoggingFilterFilterArgs(WebAclLoggingConfigurationLoggingFilterFilterArgs $) {
        this.behavior = $.behavior;
        this.conditions = $.conditions;
        this.requirement = $.requirement;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebAclLoggingConfigurationLoggingFilterFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebAclLoggingConfigurationLoggingFilterFilterArgs $;

        public Builder() {
            $ = new WebAclLoggingConfigurationLoggingFilterFilterArgs();
        }

        public Builder(WebAclLoggingConfigurationLoggingFilterFilterArgs defaults) {
            $ = new WebAclLoggingConfigurationLoggingFilterFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param behavior How to handle logs that satisfy the filter&#39;s conditions and requirement. Valid values: `KEEP` or `DROP`.
         * 
         * @return builder
         * 
         */
        public Builder behavior(Output<String> behavior) {
            $.behavior = behavior;
            return this;
        }

        /**
         * @param behavior How to handle logs that satisfy the filter&#39;s conditions and requirement. Valid values: `KEEP` or `DROP`.
         * 
         * @return builder
         * 
         */
        public Builder behavior(String behavior) {
            return behavior(Output.of(behavior));
        }

        /**
         * @param conditions Match condition(s) for the filter. See Condition below for more details.
         * 
         * @return builder
         * 
         */
        public Builder conditions(Output<List<WebAclLoggingConfigurationLoggingFilterFilterConditionArgs>> conditions) {
            $.conditions = conditions;
            return this;
        }

        /**
         * @param conditions Match condition(s) for the filter. See Condition below for more details.
         * 
         * @return builder
         * 
         */
        public Builder conditions(List<WebAclLoggingConfigurationLoggingFilterFilterConditionArgs> conditions) {
            return conditions(Output.of(conditions));
        }

        /**
         * @param conditions Match condition(s) for the filter. See Condition below for more details.
         * 
         * @return builder
         * 
         */
        public Builder conditions(WebAclLoggingConfigurationLoggingFilterFilterConditionArgs... conditions) {
            return conditions(List.of(conditions));
        }

        /**
         * @param requirement Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition. Valid values: `MEETS_ALL` or `MEETS_ANY`.
         * 
         * @return builder
         * 
         */
        public Builder requirement(Output<String> requirement) {
            $.requirement = requirement;
            return this;
        }

        /**
         * @param requirement Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition. Valid values: `MEETS_ALL` or `MEETS_ANY`.
         * 
         * @return builder
         * 
         */
        public Builder requirement(String requirement) {
            return requirement(Output.of(requirement));
        }

        public WebAclLoggingConfigurationLoggingFilterFilterArgs build() {
            $.behavior = Objects.requireNonNull($.behavior, "expected parameter 'behavior' to be non-null");
            $.conditions = Objects.requireNonNull($.conditions, "expected parameter 'conditions' to be non-null");
            $.requirement = Objects.requireNonNull($.requirement, "expected parameter 'requirement' to be non-null");
            return $;
        }
    }

}