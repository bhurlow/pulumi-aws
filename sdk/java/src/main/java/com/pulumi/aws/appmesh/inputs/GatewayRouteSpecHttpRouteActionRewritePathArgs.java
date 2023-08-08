// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GatewayRouteSpecHttpRouteActionRewritePathArgs extends com.pulumi.resources.ResourceArgs {

    public static final GatewayRouteSpecHttpRouteActionRewritePathArgs Empty = new GatewayRouteSpecHttpRouteActionRewritePathArgs();

    /**
     * The exact path to match on.
     * 
     */
    @Import(name="exact", required=true)
    private Output<String> exact;

    /**
     * @return The exact path to match on.
     * 
     */
    public Output<String> exact() {
        return this.exact;
    }

    private GatewayRouteSpecHttpRouteActionRewritePathArgs() {}

    private GatewayRouteSpecHttpRouteActionRewritePathArgs(GatewayRouteSpecHttpRouteActionRewritePathArgs $) {
        this.exact = $.exact;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GatewayRouteSpecHttpRouteActionRewritePathArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GatewayRouteSpecHttpRouteActionRewritePathArgs $;

        public Builder() {
            $ = new GatewayRouteSpecHttpRouteActionRewritePathArgs();
        }

        public Builder(GatewayRouteSpecHttpRouteActionRewritePathArgs defaults) {
            $ = new GatewayRouteSpecHttpRouteActionRewritePathArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exact The exact path to match on.
         * 
         * @return builder
         * 
         */
        public Builder exact(Output<String> exact) {
            $.exact = exact;
            return this;
        }

        /**
         * @param exact The exact path to match on.
         * 
         * @return builder
         * 
         */
        public Builder exact(String exact) {
            return exact(Output.of(exact));
        }

        public GatewayRouteSpecHttpRouteActionRewritePathArgs build() {
            $.exact = Objects.requireNonNull($.exact, "expected parameter 'exact' to be non-null");
            return $;
        }
    }

}
