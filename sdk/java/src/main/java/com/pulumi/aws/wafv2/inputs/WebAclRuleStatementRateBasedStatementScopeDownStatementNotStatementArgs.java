// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.inputs;

import com.pulumi.aws.wafv2.inputs.WebAclRuleStatementArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;


public final class WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs extends com.pulumi.resources.ResourceArgs {

    public static final WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs Empty = new WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs();

    /**
     * The statements to combine.
     * 
     */
    @Import(name="statements", required=true)
    private Output<List<WebAclRuleStatementArgs>> statements;

    /**
     * @return The statements to combine.
     * 
     */
    public Output<List<WebAclRuleStatementArgs>> statements() {
        return this.statements;
    }

    private WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs() {}

    private WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs(WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs $) {
        this.statements = $.statements;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs $;

        public Builder() {
            $ = new WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs();
        }

        public Builder(WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs defaults) {
            $ = new WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param statements The statements to combine.
         * 
         * @return builder
         * 
         */
        public Builder statements(Output<List<WebAclRuleStatementArgs>> statements) {
            $.statements = statements;
            return this;
        }

        /**
         * @param statements The statements to combine.
         * 
         * @return builder
         * 
         */
        public Builder statements(List<WebAclRuleStatementArgs> statements) {
            return statements(Output.of(statements));
        }

        /**
         * @param statements The statements to combine.
         * 
         * @return builder
         * 
         */
        public Builder statements(WebAclRuleStatementArgs... statements) {
            return statements(List.of(statements));
        }

        public WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementArgs build() {
            $.statements = Objects.requireNonNull($.statements, "expected parameter 'statements' to be non-null");
            return $;
        }
    }

}
