// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.inputs;

import com.pulumi.aws.ecs.inputs.ServiceServiceConnectConfigurationServiceTlsIssuerCertAuthorityArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceServiceConnectConfigurationServiceTlsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceServiceConnectConfigurationServiceTlsArgs Empty = new ServiceServiceConnectConfigurationServiceTlsArgs();

    /**
     * The details of the certificate authority which will issue the certificate.
     * 
     */
    @Import(name="issuerCertAuthority", required=true)
    private Output<ServiceServiceConnectConfigurationServiceTlsIssuerCertAuthorityArgs> issuerCertAuthority;

    /**
     * @return The details of the certificate authority which will issue the certificate.
     * 
     */
    public Output<ServiceServiceConnectConfigurationServiceTlsIssuerCertAuthorityArgs> issuerCertAuthority() {
        return this.issuerCertAuthority;
    }

    /**
     * The KMS key used to encrypt the private key in Secrets Manager.
     * 
     */
    @Import(name="kmsKey")
    private @Nullable Output<String> kmsKey;

    /**
     * @return The KMS key used to encrypt the private key in Secrets Manager.
     * 
     */
    public Optional<Output<String>> kmsKey() {
        return Optional.ofNullable(this.kmsKey);
    }

    /**
     * The ARN of the IAM Role that&#39;s associated with the Service Connect TLS.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return The ARN of the IAM Role that&#39;s associated with the Service Connect TLS.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    private ServiceServiceConnectConfigurationServiceTlsArgs() {}

    private ServiceServiceConnectConfigurationServiceTlsArgs(ServiceServiceConnectConfigurationServiceTlsArgs $) {
        this.issuerCertAuthority = $.issuerCertAuthority;
        this.kmsKey = $.kmsKey;
        this.roleArn = $.roleArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceServiceConnectConfigurationServiceTlsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceServiceConnectConfigurationServiceTlsArgs $;

        public Builder() {
            $ = new ServiceServiceConnectConfigurationServiceTlsArgs();
        }

        public Builder(ServiceServiceConnectConfigurationServiceTlsArgs defaults) {
            $ = new ServiceServiceConnectConfigurationServiceTlsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param issuerCertAuthority The details of the certificate authority which will issue the certificate.
         * 
         * @return builder
         * 
         */
        public Builder issuerCertAuthority(Output<ServiceServiceConnectConfigurationServiceTlsIssuerCertAuthorityArgs> issuerCertAuthority) {
            $.issuerCertAuthority = issuerCertAuthority;
            return this;
        }

        /**
         * @param issuerCertAuthority The details of the certificate authority which will issue the certificate.
         * 
         * @return builder
         * 
         */
        public Builder issuerCertAuthority(ServiceServiceConnectConfigurationServiceTlsIssuerCertAuthorityArgs issuerCertAuthority) {
            return issuerCertAuthority(Output.of(issuerCertAuthority));
        }

        /**
         * @param kmsKey The KMS key used to encrypt the private key in Secrets Manager.
         * 
         * @return builder
         * 
         */
        public Builder kmsKey(@Nullable Output<String> kmsKey) {
            $.kmsKey = kmsKey;
            return this;
        }

        /**
         * @param kmsKey The KMS key used to encrypt the private key in Secrets Manager.
         * 
         * @return builder
         * 
         */
        public Builder kmsKey(String kmsKey) {
            return kmsKey(Output.of(kmsKey));
        }

        /**
         * @param roleArn The ARN of the IAM Role that&#39;s associated with the Service Connect TLS.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The ARN of the IAM Role that&#39;s associated with the Service Connect TLS.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        public ServiceServiceConnectConfigurationServiceTlsArgs build() {
            if ($.issuerCertAuthority == null) {
                throw new MissingRequiredPropertyException("ServiceServiceConnectConfigurationServiceTlsArgs", "issuerCertAuthority");
            }
            return $;
        }
    }

}
