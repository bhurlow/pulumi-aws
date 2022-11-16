// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apigatewayv2.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DomainNameDomainNameConfiguration {
    /**
     * @return ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the `aws.acm.Certificate` resource to configure an ACM certificate.
     * 
     */
    private String certificateArn;
    /**
     * @return Endpoint type. Valid values: `REGIONAL`.
     * 
     */
    private String endpointType;
    /**
     * @return Amazon Route 53 Hosted Zone ID of the endpoint.
     * 
     */
    private @Nullable String hostedZoneId;
    /**
     * @return ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
     * 
     */
    private @Nullable String ownershipVerificationCertificateArn;
    /**
     * @return Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
     * 
     */
    private String securityPolicy;
    /**
     * @return Target domain name.
     * 
     */
    private @Nullable String targetDomainName;

    private DomainNameDomainNameConfiguration() {}
    /**
     * @return ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the `aws.acm.Certificate` resource to configure an ACM certificate.
     * 
     */
    public String certificateArn() {
        return this.certificateArn;
    }
    /**
     * @return Endpoint type. Valid values: `REGIONAL`.
     * 
     */
    public String endpointType() {
        return this.endpointType;
    }
    /**
     * @return Amazon Route 53 Hosted Zone ID of the endpoint.
     * 
     */
    public Optional<String> hostedZoneId() {
        return Optional.ofNullable(this.hostedZoneId);
    }
    /**
     * @return ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
     * 
     */
    public Optional<String> ownershipVerificationCertificateArn() {
        return Optional.ofNullable(this.ownershipVerificationCertificateArn);
    }
    /**
     * @return Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
     * 
     */
    public String securityPolicy() {
        return this.securityPolicy;
    }
    /**
     * @return Target domain name.
     * 
     */
    public Optional<String> targetDomainName() {
        return Optional.ofNullable(this.targetDomainName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DomainNameDomainNameConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String certificateArn;
        private String endpointType;
        private @Nullable String hostedZoneId;
        private @Nullable String ownershipVerificationCertificateArn;
        private String securityPolicy;
        private @Nullable String targetDomainName;
        public Builder() {}
        public Builder(DomainNameDomainNameConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certificateArn = defaults.certificateArn;
    	      this.endpointType = defaults.endpointType;
    	      this.hostedZoneId = defaults.hostedZoneId;
    	      this.ownershipVerificationCertificateArn = defaults.ownershipVerificationCertificateArn;
    	      this.securityPolicy = defaults.securityPolicy;
    	      this.targetDomainName = defaults.targetDomainName;
        }

        @CustomType.Setter
        public Builder certificateArn(String certificateArn) {
            this.certificateArn = Objects.requireNonNull(certificateArn);
            return this;
        }
        @CustomType.Setter
        public Builder endpointType(String endpointType) {
            this.endpointType = Objects.requireNonNull(endpointType);
            return this;
        }
        @CustomType.Setter
        public Builder hostedZoneId(@Nullable String hostedZoneId) {
            this.hostedZoneId = hostedZoneId;
            return this;
        }
        @CustomType.Setter
        public Builder ownershipVerificationCertificateArn(@Nullable String ownershipVerificationCertificateArn) {
            this.ownershipVerificationCertificateArn = ownershipVerificationCertificateArn;
            return this;
        }
        @CustomType.Setter
        public Builder securityPolicy(String securityPolicy) {
            this.securityPolicy = Objects.requireNonNull(securityPolicy);
            return this;
        }
        @CustomType.Setter
        public Builder targetDomainName(@Nullable String targetDomainName) {
            this.targetDomainName = targetDomainName;
            return this;
        }
        public DomainNameDomainNameConfiguration build() {
            final var o = new DomainNameDomainNameConfiguration();
            o.certificateArn = certificateArn;
            o.endpointType = endpointType;
            o.hostedZoneId = hostedZoneId;
            o.ownershipVerificationCertificateArn = ownershipVerificationCertificateArn;
            o.securityPolicy = securityPolicy;
            o.targetDomainName = targetDomainName;
            return o;
        }
    }
}
