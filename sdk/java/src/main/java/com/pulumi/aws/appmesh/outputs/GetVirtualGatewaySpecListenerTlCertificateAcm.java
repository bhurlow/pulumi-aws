// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetVirtualGatewaySpecListenerTlCertificateAcm {
    /**
     * @return ARN for the certificate.
     * 
     */
    private String certificateArn;

    private GetVirtualGatewaySpecListenerTlCertificateAcm() {}
    /**
     * @return ARN for the certificate.
     * 
     */
    public String certificateArn() {
        return this.certificateArn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualGatewaySpecListenerTlCertificateAcm defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String certificateArn;
        public Builder() {}
        public Builder(GetVirtualGatewaySpecListenerTlCertificateAcm defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certificateArn = defaults.certificateArn;
        }

        @CustomType.Setter
        public Builder certificateArn(String certificateArn) {
            this.certificateArn = Objects.requireNonNull(certificateArn);
            return this;
        }
        public GetVirtualGatewaySpecListenerTlCertificateAcm build() {
            final var o = new GetVirtualGatewaySpecListenerTlCertificateAcm();
            o.certificateArn = certificateArn;
            return o;
        }
    }
}
