// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.inputs;

import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecBackendArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecBackendDefaultsArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecListenerArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecLoggingArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualNodeSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualNodeSpecArgs Empty = new VirtualNodeSpecArgs();

    /**
     * Defaults for backends.
     * 
     */
    @Import(name="backendDefaults")
    private @Nullable Output<VirtualNodeSpecBackendDefaultsArgs> backendDefaults;

    /**
     * @return Defaults for backends.
     * 
     */
    public Optional<Output<VirtualNodeSpecBackendDefaultsArgs>> backendDefaults() {
        return Optional.ofNullable(this.backendDefaults);
    }

    /**
     * Backends to which the virtual node is expected to send outbound traffic.
     * 
     */
    @Import(name="backends")
    private @Nullable Output<List<VirtualNodeSpecBackendArgs>> backends;

    /**
     * @return Backends to which the virtual node is expected to send outbound traffic.
     * 
     */
    public Optional<Output<List<VirtualNodeSpecBackendArgs>>> backends() {
        return Optional.ofNullable(this.backends);
    }

    /**
     * Listeners from which the virtual node is expected to receive inbound traffic.
     * 
     */
    @Import(name="listeners")
    private @Nullable Output<List<VirtualNodeSpecListenerArgs>> listeners;

    /**
     * @return Listeners from which the virtual node is expected to receive inbound traffic.
     * 
     */
    public Optional<Output<List<VirtualNodeSpecListenerArgs>>> listeners() {
        return Optional.ofNullable(this.listeners);
    }

    /**
     * Inbound and outbound access logging information for the virtual node.
     * 
     */
    @Import(name="logging")
    private @Nullable Output<VirtualNodeSpecLoggingArgs> logging;

    /**
     * @return Inbound and outbound access logging information for the virtual node.
     * 
     */
    public Optional<Output<VirtualNodeSpecLoggingArgs>> logging() {
        return Optional.ofNullable(this.logging);
    }

    /**
     * Service discovery information for the virtual node.
     * 
     */
    @Import(name="serviceDiscovery")
    private @Nullable Output<VirtualNodeSpecServiceDiscoveryArgs> serviceDiscovery;

    /**
     * @return Service discovery information for the virtual node.
     * 
     */
    public Optional<Output<VirtualNodeSpecServiceDiscoveryArgs>> serviceDiscovery() {
        return Optional.ofNullable(this.serviceDiscovery);
    }

    private VirtualNodeSpecArgs() {}

    private VirtualNodeSpecArgs(VirtualNodeSpecArgs $) {
        this.backendDefaults = $.backendDefaults;
        this.backends = $.backends;
        this.listeners = $.listeners;
        this.logging = $.logging;
        this.serviceDiscovery = $.serviceDiscovery;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualNodeSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualNodeSpecArgs $;

        public Builder() {
            $ = new VirtualNodeSpecArgs();
        }

        public Builder(VirtualNodeSpecArgs defaults) {
            $ = new VirtualNodeSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backendDefaults Defaults for backends.
         * 
         * @return builder
         * 
         */
        public Builder backendDefaults(@Nullable Output<VirtualNodeSpecBackendDefaultsArgs> backendDefaults) {
            $.backendDefaults = backendDefaults;
            return this;
        }

        /**
         * @param backendDefaults Defaults for backends.
         * 
         * @return builder
         * 
         */
        public Builder backendDefaults(VirtualNodeSpecBackendDefaultsArgs backendDefaults) {
            return backendDefaults(Output.of(backendDefaults));
        }

        /**
         * @param backends Backends to which the virtual node is expected to send outbound traffic.
         * 
         * @return builder
         * 
         */
        public Builder backends(@Nullable Output<List<VirtualNodeSpecBackendArgs>> backends) {
            $.backends = backends;
            return this;
        }

        /**
         * @param backends Backends to which the virtual node is expected to send outbound traffic.
         * 
         * @return builder
         * 
         */
        public Builder backends(List<VirtualNodeSpecBackendArgs> backends) {
            return backends(Output.of(backends));
        }

        /**
         * @param backends Backends to which the virtual node is expected to send outbound traffic.
         * 
         * @return builder
         * 
         */
        public Builder backends(VirtualNodeSpecBackendArgs... backends) {
            return backends(List.of(backends));
        }

        /**
         * @param listeners Listeners from which the virtual node is expected to receive inbound traffic.
         * 
         * @return builder
         * 
         */
        public Builder listeners(@Nullable Output<List<VirtualNodeSpecListenerArgs>> listeners) {
            $.listeners = listeners;
            return this;
        }

        /**
         * @param listeners Listeners from which the virtual node is expected to receive inbound traffic.
         * 
         * @return builder
         * 
         */
        public Builder listeners(List<VirtualNodeSpecListenerArgs> listeners) {
            return listeners(Output.of(listeners));
        }

        /**
         * @param listeners Listeners from which the virtual node is expected to receive inbound traffic.
         * 
         * @return builder
         * 
         */
        public Builder listeners(VirtualNodeSpecListenerArgs... listeners) {
            return listeners(List.of(listeners));
        }

        /**
         * @param logging Inbound and outbound access logging information for the virtual node.
         * 
         * @return builder
         * 
         */
        public Builder logging(@Nullable Output<VirtualNodeSpecLoggingArgs> logging) {
            $.logging = logging;
            return this;
        }

        /**
         * @param logging Inbound and outbound access logging information for the virtual node.
         * 
         * @return builder
         * 
         */
        public Builder logging(VirtualNodeSpecLoggingArgs logging) {
            return logging(Output.of(logging));
        }

        /**
         * @param serviceDiscovery Service discovery information for the virtual node.
         * 
         * @return builder
         * 
         */
        public Builder serviceDiscovery(@Nullable Output<VirtualNodeSpecServiceDiscoveryArgs> serviceDiscovery) {
            $.serviceDiscovery = serviceDiscovery;
            return this;
        }

        /**
         * @param serviceDiscovery Service discovery information for the virtual node.
         * 
         * @return builder
         * 
         */
        public Builder serviceDiscovery(VirtualNodeSpecServiceDiscoveryArgs serviceDiscovery) {
            return serviceDiscovery(Output.of(serviceDiscovery));
        }

        public VirtualNodeSpecArgs build() {
            return $;
        }
    }

}
