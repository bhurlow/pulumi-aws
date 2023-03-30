// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRouteSpecHttp2RouteTimeoutIdle {
    private String unit;
    private Integer value;

    private GetRouteSpecHttp2RouteTimeoutIdle() {}
    public String unit() {
        return this.unit;
    }
    public Integer value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouteSpecHttp2RouteTimeoutIdle defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String unit;
        private Integer value;
        public Builder() {}
        public Builder(GetRouteSpecHttp2RouteTimeoutIdle defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.unit = defaults.unit;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder unit(String unit) {
            this.unit = Objects.requireNonNull(unit);
            return this;
        }
        @CustomType.Setter
        public Builder value(Integer value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public GetRouteSpecHttp2RouteTimeoutIdle build() {
            final var o = new GetRouteSpecHttp2RouteTimeoutIdle();
            o.unit = unit;
            o.value = value;
            return o;
        }
    }
}
