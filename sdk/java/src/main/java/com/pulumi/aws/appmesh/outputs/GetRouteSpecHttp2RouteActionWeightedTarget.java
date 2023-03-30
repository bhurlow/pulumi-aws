// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRouteSpecHttp2RouteActionWeightedTarget {
    private Integer port;
    private String virtualNode;
    private Integer weight;

    private GetRouteSpecHttp2RouteActionWeightedTarget() {}
    public Integer port() {
        return this.port;
    }
    public String virtualNode() {
        return this.virtualNode;
    }
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouteSpecHttp2RouteActionWeightedTarget defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer port;
        private String virtualNode;
        private Integer weight;
        public Builder() {}
        public Builder(GetRouteSpecHttp2RouteActionWeightedTarget defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.port = defaults.port;
    	      this.virtualNode = defaults.virtualNode;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder virtualNode(String virtualNode) {
            this.virtualNode = Objects.requireNonNull(virtualNode);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public GetRouteSpecHttp2RouteActionWeightedTarget build() {
            final var o = new GetRouteSpecHttp2RouteActionWeightedTarget();
            o.port = port;
            o.virtualNode = virtualNode;
            o.weight = weight;
            return o;
        }
    }
}
