// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.waf.inputs;

import com.pulumi.aws.waf.inputs.RuleGroupActivatedRuleActionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleGroupActivatedRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleGroupActivatedRuleArgs Empty = new RuleGroupActivatedRuleArgs();

    /**
     * Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
     * 
     */
    @Import(name="action", required=true)
    private Output<RuleGroupActivatedRuleActionArgs> action;

    /**
     * @return Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
     * 
     */
    public Output<RuleGroupActivatedRuleActionArgs> action() {
        return this.action;
    }

    /**
     * Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
     * 
     */
    @Import(name="priority", required=true)
    private Output<Integer> priority;

    /**
     * @return Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }

    /**
     * The ID of a rule
     * 
     */
    @Import(name="ruleId", required=true)
    private Output<String> ruleId;

    /**
     * @return The ID of a rule
     * 
     */
    public Output<String> ruleId() {
        return this.ruleId;
    }

    /**
     * The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private RuleGroupActivatedRuleArgs() {}

    private RuleGroupActivatedRuleArgs(RuleGroupActivatedRuleArgs $) {
        this.action = $.action;
        this.priority = $.priority;
        this.ruleId = $.ruleId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleGroupActivatedRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleGroupActivatedRuleArgs $;

        public Builder() {
            $ = new RuleGroupActivatedRuleArgs();
        }

        public Builder(RuleGroupActivatedRuleArgs defaults) {
            $ = new RuleGroupActivatedRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param action Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
         * 
         * @return builder
         * 
         */
        public Builder action(Output<RuleGroupActivatedRuleActionArgs> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
         * 
         * @return builder
         * 
         */
        public Builder action(RuleGroupActivatedRuleActionArgs action) {
            return action(Output.of(action));
        }

        /**
         * @param priority Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
         * 
         * @return builder
         * 
         */
        public Builder priority(Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param ruleId The ID of a rule
         * 
         * @return builder
         * 
         */
        public Builder ruleId(Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId The ID of a rule
         * 
         * @return builder
         * 
         */
        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        /**
         * @param type The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public RuleGroupActivatedRuleArgs build() {
            $.action = Objects.requireNonNull($.action, "expected parameter 'action' to be non-null");
            $.priority = Objects.requireNonNull($.priority, "expected parameter 'priority' to be non-null");
            $.ruleId = Objects.requireNonNull($.ruleId, "expected parameter 'ruleId' to be non-null");
            return $;
        }
    }

}
