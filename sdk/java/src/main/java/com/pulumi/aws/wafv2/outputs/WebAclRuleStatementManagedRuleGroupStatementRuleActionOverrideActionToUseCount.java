// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.outputs;

import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandling;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount {
    /**
     * @return Defines custom handling for the web request. See Custom Request Handling below for details.
     * 
     */
    private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandling customRequestHandling;

    private WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount() {}
    /**
     * @return Defines custom handling for the web request. See Custom Request Handling below for details.
     * 
     */
    public Optional<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandling> customRequestHandling() {
        return Optional.ofNullable(this.customRequestHandling);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandling customRequestHandling;
        public Builder() {}
        public Builder(WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customRequestHandling = defaults.customRequestHandling;
        }

        @CustomType.Setter
        public Builder customRequestHandling(@Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandling customRequestHandling) {
            this.customRequestHandling = customRequestHandling;
            return this;
        }
        public WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount build() {
            final var o = new WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount();
            o.customRequestHandling = customRequestHandling;
            return o;
        }
    }
}
