// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.aws.appmesh.outputs.GetVirtualGatewaySpecLoggingAccessLog;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetVirtualGatewaySpecLogging {
    /**
     * @return Access log configuration for a virtual gateway.
     * 
     */
    private List<GetVirtualGatewaySpecLoggingAccessLog> accessLogs;

    private GetVirtualGatewaySpecLogging() {}
    /**
     * @return Access log configuration for a virtual gateway.
     * 
     */
    public List<GetVirtualGatewaySpecLoggingAccessLog> accessLogs() {
        return this.accessLogs;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualGatewaySpecLogging defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetVirtualGatewaySpecLoggingAccessLog> accessLogs;
        public Builder() {}
        public Builder(GetVirtualGatewaySpecLogging defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessLogs = defaults.accessLogs;
        }

        @CustomType.Setter
        public Builder accessLogs(List<GetVirtualGatewaySpecLoggingAccessLog> accessLogs) {
            this.accessLogs = Objects.requireNonNull(accessLogs);
            return this;
        }
        public Builder accessLogs(GetVirtualGatewaySpecLoggingAccessLog... accessLogs) {
            return accessLogs(List.of(accessLogs));
        }
        public GetVirtualGatewaySpecLogging build() {
            final var o = new GetVirtualGatewaySpecLogging();
            o.accessLogs = accessLogs;
            return o;
        }
    }
}
