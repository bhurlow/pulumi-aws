// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class EndpointConfigurationShadowProductionVariantRoutingConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointConfigurationShadowProductionVariantRoutingConfigArgs Empty = new EndpointConfigurationShadowProductionVariantRoutingConfigArgs();

    /**
     * Sets how the endpoint routes incoming traffic. Valid values are `LEAST_OUTSTANDING_REQUESTS` and `RANDOM`. `LEAST_OUTSTANDING_REQUESTS` routes requests to the specific instances that have more capacity to process them. `RANDOM` routes each request to a randomly chosen instance.
     * 
     */
    @Import(name="routingStrategy", required=true)
    private Output<String> routingStrategy;

    /**
     * @return Sets how the endpoint routes incoming traffic. Valid values are `LEAST_OUTSTANDING_REQUESTS` and `RANDOM`. `LEAST_OUTSTANDING_REQUESTS` routes requests to the specific instances that have more capacity to process them. `RANDOM` routes each request to a randomly chosen instance.
     * 
     */
    public Output<String> routingStrategy() {
        return this.routingStrategy;
    }

    private EndpointConfigurationShadowProductionVariantRoutingConfigArgs() {}

    private EndpointConfigurationShadowProductionVariantRoutingConfigArgs(EndpointConfigurationShadowProductionVariantRoutingConfigArgs $) {
        this.routingStrategy = $.routingStrategy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointConfigurationShadowProductionVariantRoutingConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointConfigurationShadowProductionVariantRoutingConfigArgs $;

        public Builder() {
            $ = new EndpointConfigurationShadowProductionVariantRoutingConfigArgs();
        }

        public Builder(EndpointConfigurationShadowProductionVariantRoutingConfigArgs defaults) {
            $ = new EndpointConfigurationShadowProductionVariantRoutingConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param routingStrategy Sets how the endpoint routes incoming traffic. Valid values are `LEAST_OUTSTANDING_REQUESTS` and `RANDOM`. `LEAST_OUTSTANDING_REQUESTS` routes requests to the specific instances that have more capacity to process them. `RANDOM` routes each request to a randomly chosen instance.
         * 
         * @return builder
         * 
         */
        public Builder routingStrategy(Output<String> routingStrategy) {
            $.routingStrategy = routingStrategy;
            return this;
        }

        /**
         * @param routingStrategy Sets how the endpoint routes incoming traffic. Valid values are `LEAST_OUTSTANDING_REQUESTS` and `RANDOM`. `LEAST_OUTSTANDING_REQUESTS` routes requests to the specific instances that have more capacity to process them. `RANDOM` routes each request to a randomly chosen instance.
         * 
         * @return builder
         * 
         */
        public Builder routingStrategy(String routingStrategy) {
            return routingStrategy(Output.of(routingStrategy));
        }

        public EndpointConfigurationShadowProductionVariantRoutingConfigArgs build() {
            if ($.routingStrategy == null) {
                throw new MissingRequiredPropertyException("EndpointConfigurationShadowProductionVariantRoutingConfigArgs", "routingStrategy");
            }
            return $;
        }
    }

}
