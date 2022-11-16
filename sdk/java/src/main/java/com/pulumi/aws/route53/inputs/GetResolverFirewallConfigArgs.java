// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetResolverFirewallConfigArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetResolverFirewallConfigArgs Empty = new GetResolverFirewallConfigArgs();

    /**
     * The ID of the VPC from Amazon VPC that the configuration is for.
     * 
     */
    @Import(name="resourceId", required=true)
    private Output<String> resourceId;

    /**
     * @return The ID of the VPC from Amazon VPC that the configuration is for.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }

    private GetResolverFirewallConfigArgs() {}

    private GetResolverFirewallConfigArgs(GetResolverFirewallConfigArgs $) {
        this.resourceId = $.resourceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResolverFirewallConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResolverFirewallConfigArgs $;

        public Builder() {
            $ = new GetResolverFirewallConfigArgs();
        }

        public Builder(GetResolverFirewallConfigArgs defaults) {
            $ = new GetResolverFirewallConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param resourceId The ID of the VPC from Amazon VPC that the configuration is for.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of the VPC from Amazon VPC that the configuration is for.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        public GetResolverFirewallConfigArgs build() {
            $.resourceId = Objects.requireNonNull($.resourceId, "expected parameter 'resourceId' to be non-null");
            return $;
        }
    }

}
