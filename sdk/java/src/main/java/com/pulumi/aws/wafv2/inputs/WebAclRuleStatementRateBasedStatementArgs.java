// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.inputs;

import com.pulumi.aws.wafv2.inputs.WebAclRuleStatementRateBasedStatementForwardedIpConfigArgs;
import com.pulumi.aws.wafv2.inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WebAclRuleStatementRateBasedStatementArgs extends com.pulumi.resources.ResourceArgs {

    public static final WebAclRuleStatementRateBasedStatementArgs Empty = new WebAclRuleStatementRateBasedStatementArgs();

    /**
     * Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `FORWARDED_IP` or `IP`. Default: `IP`.
     * 
     */
    @Import(name="aggregateKeyType")
    private @Nullable Output<String> aggregateKeyType;

    /**
     * @return Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `FORWARDED_IP` or `IP`. Default: `IP`.
     * 
     */
    public Optional<Output<String>> aggregateKeyType() {
        return Optional.ofNullable(this.aggregateKeyType);
    }

    /**
     * Configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that&#39;s reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See `forwarded_ip_config` below for details.
     * 
     */
    @Import(name="forwardedIpConfig")
    private @Nullable Output<WebAclRuleStatementRateBasedStatementForwardedIpConfigArgs> forwardedIpConfig;

    /**
     * @return Configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that&#39;s reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See `forwarded_ip_config` below for details.
     * 
     */
    public Optional<Output<WebAclRuleStatementRateBasedStatementForwardedIpConfigArgs>> forwardedIpConfig() {
        return Optional.ofNullable(this.forwardedIpConfig);
    }

    /**
     * Limit on requests per 5-minute period for a single originating IP address.
     * 
     */
    @Import(name="limit", required=true)
    private Output<Integer> limit;

    /**
     * @return Limit on requests per 5-minute period for a single originating IP address.
     * 
     */
    public Output<Integer> limit() {
        return this.limit;
    }

    /**
     * Optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
     * 
     */
    @Import(name="scopeDownStatement")
    private @Nullable Output<WebAclRuleStatementRateBasedStatementScopeDownStatementArgs> scopeDownStatement;

    /**
     * @return Optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
     * 
     */
    public Optional<Output<WebAclRuleStatementRateBasedStatementScopeDownStatementArgs>> scopeDownStatement() {
        return Optional.ofNullable(this.scopeDownStatement);
    }

    private WebAclRuleStatementRateBasedStatementArgs() {}

    private WebAclRuleStatementRateBasedStatementArgs(WebAclRuleStatementRateBasedStatementArgs $) {
        this.aggregateKeyType = $.aggregateKeyType;
        this.forwardedIpConfig = $.forwardedIpConfig;
        this.limit = $.limit;
        this.scopeDownStatement = $.scopeDownStatement;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebAclRuleStatementRateBasedStatementArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebAclRuleStatementRateBasedStatementArgs $;

        public Builder() {
            $ = new WebAclRuleStatementRateBasedStatementArgs();
        }

        public Builder(WebAclRuleStatementRateBasedStatementArgs defaults) {
            $ = new WebAclRuleStatementRateBasedStatementArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aggregateKeyType Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `FORWARDED_IP` or `IP`. Default: `IP`.
         * 
         * @return builder
         * 
         */
        public Builder aggregateKeyType(@Nullable Output<String> aggregateKeyType) {
            $.aggregateKeyType = aggregateKeyType;
            return this;
        }

        /**
         * @param aggregateKeyType Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `FORWARDED_IP` or `IP`. Default: `IP`.
         * 
         * @return builder
         * 
         */
        public Builder aggregateKeyType(String aggregateKeyType) {
            return aggregateKeyType(Output.of(aggregateKeyType));
        }

        /**
         * @param forwardedIpConfig Configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that&#39;s reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See `forwarded_ip_config` below for details.
         * 
         * @return builder
         * 
         */
        public Builder forwardedIpConfig(@Nullable Output<WebAclRuleStatementRateBasedStatementForwardedIpConfigArgs> forwardedIpConfig) {
            $.forwardedIpConfig = forwardedIpConfig;
            return this;
        }

        /**
         * @param forwardedIpConfig Configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that&#39;s reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See `forwarded_ip_config` below for details.
         * 
         * @return builder
         * 
         */
        public Builder forwardedIpConfig(WebAclRuleStatementRateBasedStatementForwardedIpConfigArgs forwardedIpConfig) {
            return forwardedIpConfig(Output.of(forwardedIpConfig));
        }

        /**
         * @param limit Limit on requests per 5-minute period for a single originating IP address.
         * 
         * @return builder
         * 
         */
        public Builder limit(Output<Integer> limit) {
            $.limit = limit;
            return this;
        }

        /**
         * @param limit Limit on requests per 5-minute period for a single originating IP address.
         * 
         * @return builder
         * 
         */
        public Builder limit(Integer limit) {
            return limit(Output.of(limit));
        }

        /**
         * @param scopeDownStatement Optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
         * 
         * @return builder
         * 
         */
        public Builder scopeDownStatement(@Nullable Output<WebAclRuleStatementRateBasedStatementScopeDownStatementArgs> scopeDownStatement) {
            $.scopeDownStatement = scopeDownStatement;
            return this;
        }

        /**
         * @param scopeDownStatement Optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
         * 
         * @return builder
         * 
         */
        public Builder scopeDownStatement(WebAclRuleStatementRateBasedStatementScopeDownStatementArgs scopeDownStatement) {
            return scopeDownStatement(Output.of(scopeDownStatement));
        }

        public WebAclRuleStatementRateBasedStatementArgs build() {
            $.limit = Objects.requireNonNull($.limit, "expected parameter 'limit' to be non-null");
            return $;
        }
    }

}
