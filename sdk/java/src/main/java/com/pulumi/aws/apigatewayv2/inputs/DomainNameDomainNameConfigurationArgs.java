// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apigatewayv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainNameDomainNameConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainNameDomainNameConfigurationArgs Empty = new DomainNameDomainNameConfigurationArgs();

    /**
     * ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the `aws.acm.Certificate` resource to configure an ACM certificate.
     * 
     */
    @Import(name="certificateArn", required=true)
    private Output<String> certificateArn;

    /**
     * @return ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the `aws.acm.Certificate` resource to configure an ACM certificate.
     * 
     */
    public Output<String> certificateArn() {
        return this.certificateArn;
    }

    /**
     * Endpoint type. Valid values: `REGIONAL`.
     * 
     */
    @Import(name="endpointType", required=true)
    private Output<String> endpointType;

    /**
     * @return Endpoint type. Valid values: `REGIONAL`.
     * 
     */
    public Output<String> endpointType() {
        return this.endpointType;
    }

    /**
     * Amazon Route 53 Hosted Zone ID of the endpoint.
     * 
     */
    @Import(name="hostedZoneId")
    private @Nullable Output<String> hostedZoneId;

    /**
     * @return Amazon Route 53 Hosted Zone ID of the endpoint.
     * 
     */
    public Optional<Output<String>> hostedZoneId() {
        return Optional.ofNullable(this.hostedZoneId);
    }

    /**
     * ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
     * 
     */
    @Import(name="ownershipVerificationCertificateArn")
    private @Nullable Output<String> ownershipVerificationCertificateArn;

    /**
     * @return ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
     * 
     */
    public Optional<Output<String>> ownershipVerificationCertificateArn() {
        return Optional.ofNullable(this.ownershipVerificationCertificateArn);
    }

    /**
     * Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
     * 
     */
    @Import(name="securityPolicy", required=true)
    private Output<String> securityPolicy;

    /**
     * @return Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
     * 
     */
    public Output<String> securityPolicy() {
        return this.securityPolicy;
    }

    /**
     * Target domain name.
     * 
     */
    @Import(name="targetDomainName")
    private @Nullable Output<String> targetDomainName;

    /**
     * @return Target domain name.
     * 
     */
    public Optional<Output<String>> targetDomainName() {
        return Optional.ofNullable(this.targetDomainName);
    }

    private DomainNameDomainNameConfigurationArgs() {}

    private DomainNameDomainNameConfigurationArgs(DomainNameDomainNameConfigurationArgs $) {
        this.certificateArn = $.certificateArn;
        this.endpointType = $.endpointType;
        this.hostedZoneId = $.hostedZoneId;
        this.ownershipVerificationCertificateArn = $.ownershipVerificationCertificateArn;
        this.securityPolicy = $.securityPolicy;
        this.targetDomainName = $.targetDomainName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainNameDomainNameConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainNameDomainNameConfigurationArgs $;

        public Builder() {
            $ = new DomainNameDomainNameConfigurationArgs();
        }

        public Builder(DomainNameDomainNameConfigurationArgs defaults) {
            $ = new DomainNameDomainNameConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificateArn ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the `aws.acm.Certificate` resource to configure an ACM certificate.
         * 
         * @return builder
         * 
         */
        public Builder certificateArn(Output<String> certificateArn) {
            $.certificateArn = certificateArn;
            return this;
        }

        /**
         * @param certificateArn ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the `aws.acm.Certificate` resource to configure an ACM certificate.
         * 
         * @return builder
         * 
         */
        public Builder certificateArn(String certificateArn) {
            return certificateArn(Output.of(certificateArn));
        }

        /**
         * @param endpointType Endpoint type. Valid values: `REGIONAL`.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(Output<String> endpointType) {
            $.endpointType = endpointType;
            return this;
        }

        /**
         * @param endpointType Endpoint type. Valid values: `REGIONAL`.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(String endpointType) {
            return endpointType(Output.of(endpointType));
        }

        /**
         * @param hostedZoneId Amazon Route 53 Hosted Zone ID of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder hostedZoneId(@Nullable Output<String> hostedZoneId) {
            $.hostedZoneId = hostedZoneId;
            return this;
        }

        /**
         * @param hostedZoneId Amazon Route 53 Hosted Zone ID of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder hostedZoneId(String hostedZoneId) {
            return hostedZoneId(Output.of(hostedZoneId));
        }

        /**
         * @param ownershipVerificationCertificateArn ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
         * 
         * @return builder
         * 
         */
        public Builder ownershipVerificationCertificateArn(@Nullable Output<String> ownershipVerificationCertificateArn) {
            $.ownershipVerificationCertificateArn = ownershipVerificationCertificateArn;
            return this;
        }

        /**
         * @param ownershipVerificationCertificateArn ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
         * 
         * @return builder
         * 
         */
        public Builder ownershipVerificationCertificateArn(String ownershipVerificationCertificateArn) {
            return ownershipVerificationCertificateArn(Output.of(ownershipVerificationCertificateArn));
        }

        /**
         * @param securityPolicy Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicy(Output<String> securityPolicy) {
            $.securityPolicy = securityPolicy;
            return this;
        }

        /**
         * @param securityPolicy Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicy(String securityPolicy) {
            return securityPolicy(Output.of(securityPolicy));
        }

        /**
         * @param targetDomainName Target domain name.
         * 
         * @return builder
         * 
         */
        public Builder targetDomainName(@Nullable Output<String> targetDomainName) {
            $.targetDomainName = targetDomainName;
            return this;
        }

        /**
         * @param targetDomainName Target domain name.
         * 
         * @return builder
         * 
         */
        public Builder targetDomainName(String targetDomainName) {
            return targetDomainName(Output.of(targetDomainName));
        }

        public DomainNameDomainNameConfigurationArgs build() {
            $.certificateArn = Objects.requireNonNull($.certificateArn, "expected parameter 'certificateArn' to be non-null");
            $.endpointType = Objects.requireNonNull($.endpointType, "expected parameter 'endpointType' to be non-null");
            $.securityPolicy = Objects.requireNonNull($.securityPolicy, "expected parameter 'securityPolicy' to be non-null");
            return $;
        }
    }

}
