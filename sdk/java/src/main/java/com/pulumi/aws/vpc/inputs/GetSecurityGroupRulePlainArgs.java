// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpc.inputs;

import com.pulumi.aws.vpc.inputs.GetSecurityGroupRuleFilter;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecurityGroupRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecurityGroupRulePlainArgs Empty = new GetSecurityGroupRulePlainArgs();

    /**
     * Configuration block(s) for filtering. Detailed below.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetSecurityGroupRuleFilter> filters;

    /**
     * @return Configuration block(s) for filtering. Detailed below.
     * 
     */
    public Optional<List<GetSecurityGroupRuleFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * ID of the security group rule to select.
     * 
     */
    @Import(name="securityGroupRuleId")
    private @Nullable String securityGroupRuleId;

    /**
     * @return ID of the security group rule to select.
     * 
     */
    public Optional<String> securityGroupRuleId() {
        return Optional.ofNullable(this.securityGroupRuleId);
    }

    private GetSecurityGroupRulePlainArgs() {}

    private GetSecurityGroupRulePlainArgs(GetSecurityGroupRulePlainArgs $) {
        this.filters = $.filters;
        this.securityGroupRuleId = $.securityGroupRuleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecurityGroupRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecurityGroupRulePlainArgs $;

        public Builder() {
            $ = new GetSecurityGroupRulePlainArgs();
        }

        public Builder(GetSecurityGroupRulePlainArgs defaults) {
            $ = new GetSecurityGroupRulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Configuration block(s) for filtering. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetSecurityGroupRuleFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Configuration block(s) for filtering. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetSecurityGroupRuleFilter... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param securityGroupRuleId ID of the security group rule to select.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupRuleId(@Nullable String securityGroupRuleId) {
            $.securityGroupRuleId = securityGroupRuleId;
            return this;
        }

        public GetSecurityGroupRulePlainArgs build() {
            return $;
        }
    }

}
