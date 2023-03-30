// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd {
    /**
     * @return Name of the secret for a virtual gateway&#39;s Transport Layer Security (TLS) Secret Discovery Service validation context trust.
     * 
     */
    private String secretName;

    private GetVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd() {}
    /**
     * @return Name of the secret for a virtual gateway&#39;s Transport Layer Security (TLS) Secret Discovery Service validation context trust.
     * 
     */
    public String secretName() {
        return this.secretName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String secretName;
        public Builder() {}
        public Builder(GetVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.secretName = defaults.secretName;
        }

        @CustomType.Setter
        public Builder secretName(String secretName) {
            this.secretName = Objects.requireNonNull(secretName);
            return this;
        }
        public GetVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd build() {
            final var o = new GetVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd();
            o.secretName = secretName;
            return o;
        }
    }
}
