// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudfront.outputs;

import com.pulumi.aws.cloudfront.outputs.GetOriginRequestPolicyQueryStringsConfigQueryString;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOriginRequestPolicyQueryStringsConfig {
    private String queryStringBehavior;
    private List<GetOriginRequestPolicyQueryStringsConfigQueryString> queryStrings;

    private GetOriginRequestPolicyQueryStringsConfig() {}
    public String queryStringBehavior() {
        return this.queryStringBehavior;
    }
    public List<GetOriginRequestPolicyQueryStringsConfigQueryString> queryStrings() {
        return this.queryStrings;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOriginRequestPolicyQueryStringsConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String queryStringBehavior;
        private List<GetOriginRequestPolicyQueryStringsConfigQueryString> queryStrings;
        public Builder() {}
        public Builder(GetOriginRequestPolicyQueryStringsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.queryStringBehavior = defaults.queryStringBehavior;
    	      this.queryStrings = defaults.queryStrings;
        }

        @CustomType.Setter
        public Builder queryStringBehavior(String queryStringBehavior) {
            if (queryStringBehavior == null) {
              throw new MissingRequiredPropertyException("GetOriginRequestPolicyQueryStringsConfig", "queryStringBehavior");
            }
            this.queryStringBehavior = queryStringBehavior;
            return this;
        }
        @CustomType.Setter
        public Builder queryStrings(List<GetOriginRequestPolicyQueryStringsConfigQueryString> queryStrings) {
            if (queryStrings == null) {
              throw new MissingRequiredPropertyException("GetOriginRequestPolicyQueryStringsConfig", "queryStrings");
            }
            this.queryStrings = queryStrings;
            return this;
        }
        public Builder queryStrings(GetOriginRequestPolicyQueryStringsConfigQueryString... queryStrings) {
            return queryStrings(List.of(queryStrings));
        }
        public GetOriginRequestPolicyQueryStringsConfig build() {
            final var _resultValue = new GetOriginRequestPolicyQueryStringsConfig();
            _resultValue.queryStringBehavior = queryStringBehavior;
            _resultValue.queryStrings = queryStrings;
            return _resultValue;
        }
    }
}
