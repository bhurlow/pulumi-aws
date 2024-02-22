// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceServiceConnectConfigurationServiceTimeoutArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceServiceConnectConfigurationServiceTimeoutArgs Empty = new ServiceServiceConnectConfigurationServiceTimeoutArgs();

    /**
     * The amount of time in seconds a connection will stay active while idle. A value of 0 can be set to disable idleTimeout.
     * 
     */
    @Import(name="idleTimeoutSeconds")
    private @Nullable Output<Integer> idleTimeoutSeconds;

    /**
     * @return The amount of time in seconds a connection will stay active while idle. A value of 0 can be set to disable idleTimeout.
     * 
     */
    public Optional<Output<Integer>> idleTimeoutSeconds() {
        return Optional.ofNullable(this.idleTimeoutSeconds);
    }

    /**
     * The amount of time in seconds for the upstream to respond with a complete response per request. A value of 0 can be set to disable perRequestTimeout. Can only be set when appProtocol isn&#39;t TCP.
     * 
     */
    @Import(name="perRequestTimeoutSeconds")
    private @Nullable Output<Integer> perRequestTimeoutSeconds;

    /**
     * @return The amount of time in seconds for the upstream to respond with a complete response per request. A value of 0 can be set to disable perRequestTimeout. Can only be set when appProtocol isn&#39;t TCP.
     * 
     */
    public Optional<Output<Integer>> perRequestTimeoutSeconds() {
        return Optional.ofNullable(this.perRequestTimeoutSeconds);
    }

    private ServiceServiceConnectConfigurationServiceTimeoutArgs() {}

    private ServiceServiceConnectConfigurationServiceTimeoutArgs(ServiceServiceConnectConfigurationServiceTimeoutArgs $) {
        this.idleTimeoutSeconds = $.idleTimeoutSeconds;
        this.perRequestTimeoutSeconds = $.perRequestTimeoutSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceServiceConnectConfigurationServiceTimeoutArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceServiceConnectConfigurationServiceTimeoutArgs $;

        public Builder() {
            $ = new ServiceServiceConnectConfigurationServiceTimeoutArgs();
        }

        public Builder(ServiceServiceConnectConfigurationServiceTimeoutArgs defaults) {
            $ = new ServiceServiceConnectConfigurationServiceTimeoutArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param idleTimeoutSeconds The amount of time in seconds a connection will stay active while idle. A value of 0 can be set to disable idleTimeout.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeoutSeconds(@Nullable Output<Integer> idleTimeoutSeconds) {
            $.idleTimeoutSeconds = idleTimeoutSeconds;
            return this;
        }

        /**
         * @param idleTimeoutSeconds The amount of time in seconds a connection will stay active while idle. A value of 0 can be set to disable idleTimeout.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeoutSeconds(Integer idleTimeoutSeconds) {
            return idleTimeoutSeconds(Output.of(idleTimeoutSeconds));
        }

        /**
         * @param perRequestTimeoutSeconds The amount of time in seconds for the upstream to respond with a complete response per request. A value of 0 can be set to disable perRequestTimeout. Can only be set when appProtocol isn&#39;t TCP.
         * 
         * @return builder
         * 
         */
        public Builder perRequestTimeoutSeconds(@Nullable Output<Integer> perRequestTimeoutSeconds) {
            $.perRequestTimeoutSeconds = perRequestTimeoutSeconds;
            return this;
        }

        /**
         * @param perRequestTimeoutSeconds The amount of time in seconds for the upstream to respond with a complete response per request. A value of 0 can be set to disable perRequestTimeout. Can only be set when appProtocol isn&#39;t TCP.
         * 
         * @return builder
         * 
         */
        public Builder perRequestTimeoutSeconds(Integer perRequestTimeoutSeconds) {
            return perRequestTimeoutSeconds(Output.of(perRequestTimeoutSeconds));
        }

        public ServiceServiceConnectConfigurationServiceTimeoutArgs build() {
            return $;
        }
    }

}
