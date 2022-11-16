// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.waf.outputs;

import com.pulumi.aws.waf.outputs.RuleGroupActivatedRuleAction;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RuleGroupActivatedRule {
    /**
     * @return Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
     * 
     */
    private RuleGroupActivatedRuleAction action;
    /**
     * @return Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
     * 
     */
    private Integer priority;
    /**
     * @return The ID of a rule
     * 
     */
    private String ruleId;
    /**
     * @return The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
     * 
     */
    private @Nullable String type;

    private RuleGroupActivatedRule() {}
    /**
     * @return Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
     * 
     */
    public RuleGroupActivatedRuleAction action() {
        return this.action;
    }
    /**
     * @return Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
     * 
     */
    public Integer priority() {
        return this.priority;
    }
    /**
     * @return The ID of a rule
     * 
     */
    public String ruleId() {
        return this.ruleId;
    }
    /**
     * @return The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleGroupActivatedRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private RuleGroupActivatedRuleAction action;
        private Integer priority;
        private String ruleId;
        private @Nullable String type;
        public Builder() {}
        public Builder(RuleGroupActivatedRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.action = defaults.action;
    	      this.priority = defaults.priority;
    	      this.ruleId = defaults.ruleId;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder action(RuleGroupActivatedRuleAction action) {
            this.action = Objects.requireNonNull(action);
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        @CustomType.Setter
        public Builder ruleId(String ruleId) {
            this.ruleId = Objects.requireNonNull(ruleId);
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public RuleGroupActivatedRule build() {
            final var o = new RuleGroupActivatedRule();
            o.action = action;
            o.priority = priority;
            o.ruleId = ruleId;
            o.type = type;
            return o;
        }
    }
}
