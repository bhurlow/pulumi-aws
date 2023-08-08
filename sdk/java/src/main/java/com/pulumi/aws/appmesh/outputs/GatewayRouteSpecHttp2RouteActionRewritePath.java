// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GatewayRouteSpecHttp2RouteActionRewritePath {
    /**
     * @return The exact path to match on.
     * 
     */
    private String exact;

    private GatewayRouteSpecHttp2RouteActionRewritePath() {}
    /**
     * @return The exact path to match on.
     * 
     */
    public String exact() {
        return this.exact;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GatewayRouteSpecHttp2RouteActionRewritePath defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String exact;
        public Builder() {}
        public Builder(GatewayRouteSpecHttp2RouteActionRewritePath defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.exact = defaults.exact;
        }

        @CustomType.Setter
        public Builder exact(String exact) {
            this.exact = Objects.requireNonNull(exact);
            return this;
        }
        public GatewayRouteSpecHttp2RouteActionRewritePath build() {
            final var o = new GatewayRouteSpecHttp2RouteActionRewritePath();
            o.exact = exact;
            return o;
        }
    }
}
