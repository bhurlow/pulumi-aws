// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.outputs;

import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow;
import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock;
import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha;
import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse {
    /**
     * @return Specifies that AWS WAF should allow requests by default. See Allow below for details.
     * 
     */
    private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow allow;
    /**
     * @return Specifies that AWS WAF should block requests by default. See Block below for details.
     * 
     */
    private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock block;
    /**
     * @return Instructs AWS WAF to run a Captcha check against the web request. See Captcha below for details.
     * 
     */
    private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha captcha;
    /**
     * @return Instructs AWS WAF to count the web request and allow it. See Count below for details.
     * 
     */
    private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount count;

    private WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse() {}
    /**
     * @return Specifies that AWS WAF should allow requests by default. See Allow below for details.
     * 
     */
    public Optional<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow> allow() {
        return Optional.ofNullable(this.allow);
    }
    /**
     * @return Specifies that AWS WAF should block requests by default. See Block below for details.
     * 
     */
    public Optional<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock> block() {
        return Optional.ofNullable(this.block);
    }
    /**
     * @return Instructs AWS WAF to run a Captcha check against the web request. See Captcha below for details.
     * 
     */
    public Optional<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha> captcha() {
        return Optional.ofNullable(this.captcha);
    }
    /**
     * @return Instructs AWS WAF to count the web request and allow it. See Count below for details.
     * 
     */
    public Optional<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount> count() {
        return Optional.ofNullable(this.count);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow allow;
        private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock block;
        private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha captcha;
        private @Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount count;
        public Builder() {}
        public Builder(WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allow = defaults.allow;
    	      this.block = defaults.block;
    	      this.captcha = defaults.captcha;
    	      this.count = defaults.count;
        }

        @CustomType.Setter
        public Builder allow(@Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow allow) {
            this.allow = allow;
            return this;
        }
        @CustomType.Setter
        public Builder block(@Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock block) {
            this.block = block;
            return this;
        }
        @CustomType.Setter
        public Builder captcha(@Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha captcha) {
            this.captcha = captcha;
            return this;
        }
        @CustomType.Setter
        public Builder count(@Nullable WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount count) {
            this.count = count;
            return this;
        }
        public WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse build() {
            final var o = new WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse();
            o.allow = allow;
            o.block = block;
            o.captcha = captcha;
            o.count = count;
            return o;
        }
    }
}
