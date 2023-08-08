// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PipeTargetParametersHttpParameters {
    /**
     * @return Key-value mapping of the headers that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
     * 
     */
    private @Nullable Map<String,String> headerParameters;
    /**
     * @return The path parameter values to be used to populate API Gateway REST API or EventBridge ApiDestination path wildcards (&#34;*&#34;).
     * 
     */
    private @Nullable String pathParameterValues;
    /**
     * @return Key-value mapping of the query strings that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
     * 
     */
    private @Nullable Map<String,String> queryStringParameters;

    private PipeTargetParametersHttpParameters() {}
    /**
     * @return Key-value mapping of the headers that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
     * 
     */
    public Map<String,String> headerParameters() {
        return this.headerParameters == null ? Map.of() : this.headerParameters;
    }
    /**
     * @return The path parameter values to be used to populate API Gateway REST API or EventBridge ApiDestination path wildcards (&#34;*&#34;).
     * 
     */
    public Optional<String> pathParameterValues() {
        return Optional.ofNullable(this.pathParameterValues);
    }
    /**
     * @return Key-value mapping of the query strings that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
     * 
     */
    public Map<String,String> queryStringParameters() {
        return this.queryStringParameters == null ? Map.of() : this.queryStringParameters;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PipeTargetParametersHttpParameters defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> headerParameters;
        private @Nullable String pathParameterValues;
        private @Nullable Map<String,String> queryStringParameters;
        public Builder() {}
        public Builder(PipeTargetParametersHttpParameters defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.headerParameters = defaults.headerParameters;
    	      this.pathParameterValues = defaults.pathParameterValues;
    	      this.queryStringParameters = defaults.queryStringParameters;
        }

        @CustomType.Setter
        public Builder headerParameters(@Nullable Map<String,String> headerParameters) {
            this.headerParameters = headerParameters;
            return this;
        }
        @CustomType.Setter
        public Builder pathParameterValues(@Nullable String pathParameterValues) {
            this.pathParameterValues = pathParameterValues;
            return this;
        }
        @CustomType.Setter
        public Builder queryStringParameters(@Nullable Map<String,String> queryStringParameters) {
            this.queryStringParameters = queryStringParameters;
            return this;
        }
        public PipeTargetParametersHttpParameters build() {
            final var o = new PipeTargetParametersHttpParameters();
            o.headerParameters = headerParameters;
            o.pathParameterValues = pathParameterValues;
            o.queryStringParameters = queryStringParameters;
            return o;
        }
    }
}
