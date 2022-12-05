// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceServiceConnectConfigurationServiceClientAliasArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceServiceConnectConfigurationServiceClientAliasArgs Empty = new ServiceServiceConnectConfigurationServiceClientAliasArgs();

    /**
     * The name that you use in the applications of client tasks to connect to this service.
     * 
     */
    @Import(name="dnsName")
    private @Nullable Output<String> dnsName;

    /**
     * @return The name that you use in the applications of client tasks to connect to this service.
     * 
     */
    public Optional<Output<String>> dnsName() {
        return Optional.ofNullable(this.dnsName);
    }

    /**
     * The listening port number for the Service Connect proxy. This port is available inside of all of the tasks within the same namespace.
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return The listening port number for the Service Connect proxy. This port is available inside of all of the tasks within the same namespace.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    private ServiceServiceConnectConfigurationServiceClientAliasArgs() {}

    private ServiceServiceConnectConfigurationServiceClientAliasArgs(ServiceServiceConnectConfigurationServiceClientAliasArgs $) {
        this.dnsName = $.dnsName;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceServiceConnectConfigurationServiceClientAliasArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceServiceConnectConfigurationServiceClientAliasArgs $;

        public Builder() {
            $ = new ServiceServiceConnectConfigurationServiceClientAliasArgs();
        }

        public Builder(ServiceServiceConnectConfigurationServiceClientAliasArgs defaults) {
            $ = new ServiceServiceConnectConfigurationServiceClientAliasArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dnsName The name that you use in the applications of client tasks to connect to this service.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(@Nullable Output<String> dnsName) {
            $.dnsName = dnsName;
            return this;
        }

        /**
         * @param dnsName The name that you use in the applications of client tasks to connect to this service.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(String dnsName) {
            return dnsName(Output.of(dnsName));
        }

        /**
         * @param port The listening port number for the Service Connect proxy. This port is available inside of all of the tasks within the same namespace.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The listening port number for the Service Connect proxy. This port is available inside of all of the tasks within the same namespace.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public ServiceServiceConnectConfigurationServiceClientAliasArgs build() {
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            return $;
        }
    }

}
