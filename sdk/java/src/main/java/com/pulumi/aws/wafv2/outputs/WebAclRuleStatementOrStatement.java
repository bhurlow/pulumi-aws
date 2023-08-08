// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.outputs;

import com.pulumi.aws.wafv2.outputs.WebAclRuleStatement;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;

@CustomType
public final class WebAclRuleStatementOrStatement {
    /**
     * @return The statements to combine.
     * 
     */
    private List<WebAclRuleStatement> statements;

    private WebAclRuleStatementOrStatement() {}
    /**
     * @return The statements to combine.
     * 
     */
    public List<WebAclRuleStatement> statements() {
        return this.statements;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WebAclRuleStatementOrStatement defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<WebAclRuleStatement> statements;
        public Builder() {}
        public Builder(WebAclRuleStatementOrStatement defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.statements = defaults.statements;
        }

        @CustomType.Setter
        public Builder statements(List<WebAclRuleStatement> statements) {
            this.statements = Objects.requireNonNull(statements);
            return this;
        }
        public Builder statements(WebAclRuleStatement... statements) {
            return statements(List.of(statements));
        }
        public WebAclRuleStatementOrStatement build() {
            final var o = new WebAclRuleStatementOrStatement();
            o.statements = statements;
            return o;
        }
    }
}
