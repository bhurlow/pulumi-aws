// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudfront.outputs;

import com.pulumi.aws.cloudfront.outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig {
    /**
     * @return Whether URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for `query_string_behavior` are `none`, `whitelist`, `allExcept`, and `all`.
     * 
     */
    private String queryStringBehavior;
    /**
     * @return Configuration parameter that contains a list of query string names. See Items for more information.
     * 
     */
    private @Nullable CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings queryStrings;

    private CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig() {}
    /**
     * @return Whether URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for `query_string_behavior` are `none`, `whitelist`, `allExcept`, and `all`.
     * 
     */
    public String queryStringBehavior() {
        return this.queryStringBehavior;
    }
    /**
     * @return Configuration parameter that contains a list of query string names. See Items for more information.
     * 
     */
    public Optional<CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings> queryStrings() {
        return Optional.ofNullable(this.queryStrings);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String queryStringBehavior;
        private @Nullable CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings queryStrings;
        public Builder() {}
        public Builder(CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.queryStringBehavior = defaults.queryStringBehavior;
    	      this.queryStrings = defaults.queryStrings;
        }

        @CustomType.Setter
        public Builder queryStringBehavior(String queryStringBehavior) {
            this.queryStringBehavior = Objects.requireNonNull(queryStringBehavior);
            return this;
        }
        @CustomType.Setter
        public Builder queryStrings(@Nullable CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings queryStrings) {
            this.queryStrings = queryStrings;
            return this;
        }
        public CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig build() {
            final var o = new CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig();
            o.queryStringBehavior = queryStringBehavior;
            o.queryStrings = queryStrings;
            return o;
        }
    }
}
