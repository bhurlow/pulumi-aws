// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.inputs;

import com.pulumi.aws.wafv2.inputs.RuleGroupRuleStatementArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;


public final class RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs Empty = new RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs();

    /**
     * The statements to combine.
     * 
     */
    @Import(name="statements", required=true)
    private Output<List<RuleGroupRuleStatementArgs>> statements;

    /**
     * @return The statements to combine.
     * 
     */
    public Output<List<RuleGroupRuleStatementArgs>> statements() {
        return this.statements;
    }

    private RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs() {}

    private RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs(RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs $) {
        this.statements = $.statements;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs $;

        public Builder() {
            $ = new RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs();
        }

        public Builder(RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs defaults) {
            $ = new RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param statements The statements to combine.
         * 
         * @return builder
         * 
         */
        public Builder statements(Output<List<RuleGroupRuleStatementArgs>> statements) {
            $.statements = statements;
            return this;
        }

        /**
         * @param statements The statements to combine.
         * 
         * @return builder
         * 
         */
        public Builder statements(List<RuleGroupRuleStatementArgs> statements) {
            return statements(Output.of(statements));
        }

        /**
         * @param statements The statements to combine.
         * 
         * @return builder
         * 
         */
        public Builder statements(RuleGroupRuleStatementArgs... statements) {
            return statements(List.of(statements));
        }

        public RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs build() {
            $.statements = Objects.requireNonNull($.statements, "expected parameter 'statements' to be non-null");
            return $;
        }
    }

}
