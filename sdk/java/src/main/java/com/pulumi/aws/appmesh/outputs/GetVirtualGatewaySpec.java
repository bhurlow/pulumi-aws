// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.aws.appmesh.outputs.GetVirtualGatewaySpecBackendDefault;
import com.pulumi.aws.appmesh.outputs.GetVirtualGatewaySpecListener;
import com.pulumi.aws.appmesh.outputs.GetVirtualGatewaySpecLogging;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetVirtualGatewaySpec {
    /**
     * @return Defaults for backends.
     * 
     */
    private List<GetVirtualGatewaySpecBackendDefault> backendDefaults;
    /**
     * @return Listeners that the mesh endpoint is expected to receive inbound traffic from. You can specify one listener.
     * 
     */
    private List<GetVirtualGatewaySpecListener> listeners;
    /**
     * @return Inbound and outbound access logging information for the virtual gateway.
     * 
     */
    private List<GetVirtualGatewaySpecLogging> loggings;

    private GetVirtualGatewaySpec() {}
    /**
     * @return Defaults for backends.
     * 
     */
    public List<GetVirtualGatewaySpecBackendDefault> backendDefaults() {
        return this.backendDefaults;
    }
    /**
     * @return Listeners that the mesh endpoint is expected to receive inbound traffic from. You can specify one listener.
     * 
     */
    public List<GetVirtualGatewaySpecListener> listeners() {
        return this.listeners;
    }
    /**
     * @return Inbound and outbound access logging information for the virtual gateway.
     * 
     */
    public List<GetVirtualGatewaySpecLogging> loggings() {
        return this.loggings;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualGatewaySpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetVirtualGatewaySpecBackendDefault> backendDefaults;
        private List<GetVirtualGatewaySpecListener> listeners;
        private List<GetVirtualGatewaySpecLogging> loggings;
        public Builder() {}
        public Builder(GetVirtualGatewaySpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backendDefaults = defaults.backendDefaults;
    	      this.listeners = defaults.listeners;
    	      this.loggings = defaults.loggings;
        }

        @CustomType.Setter
        public Builder backendDefaults(List<GetVirtualGatewaySpecBackendDefault> backendDefaults) {
            this.backendDefaults = Objects.requireNonNull(backendDefaults);
            return this;
        }
        public Builder backendDefaults(GetVirtualGatewaySpecBackendDefault... backendDefaults) {
            return backendDefaults(List.of(backendDefaults));
        }
        @CustomType.Setter
        public Builder listeners(List<GetVirtualGatewaySpecListener> listeners) {
            this.listeners = Objects.requireNonNull(listeners);
            return this;
        }
        public Builder listeners(GetVirtualGatewaySpecListener... listeners) {
            return listeners(List.of(listeners));
        }
        @CustomType.Setter
        public Builder loggings(List<GetVirtualGatewaySpecLogging> loggings) {
            this.loggings = Objects.requireNonNull(loggings);
            return this;
        }
        public Builder loggings(GetVirtualGatewaySpecLogging... loggings) {
            return loggings(List.of(loggings));
        }
        public GetVirtualGatewaySpec build() {
            final var o = new GetVirtualGatewaySpec();
            o.backendDefaults = backendDefaults;
            o.listeners = listeners;
            o.loggings = loggings;
            return o;
        }
    }
}
