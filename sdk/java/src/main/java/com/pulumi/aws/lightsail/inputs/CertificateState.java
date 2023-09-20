// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lightsail.inputs;

import com.pulumi.aws.lightsail.inputs.CertificateDomainValidationOptionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertificateState extends com.pulumi.resources.ResourceArgs {

    public static final CertificateState Empty = new CertificateState();

    /**
     * The ARN of the lightsail certificate.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the lightsail certificate.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The timestamp when the instance was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The timestamp when the instance was created.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * A domain name for which the certificate should be issued.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return A domain name for which the certificate should be issued.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g., if SANs are defined.
     * 
     */
    @Import(name="domainValidationOptions")
    private @Nullable Output<List<CertificateDomainValidationOptionArgs>> domainValidationOptions;

    /**
     * @return Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g., if SANs are defined.
     * 
     */
    public Optional<Output<List<CertificateDomainValidationOptionArgs>>> domainValidationOptions() {
        return Optional.ofNullable(this.domainValidationOptions);
    }

    /**
     * The name of the Lightsail load balancer.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Lightsail load balancer.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
     * 
     */
    @Import(name="subjectAlternativeNames")
    private @Nullable Output<List<String>> subjectAlternativeNames;

    /**
     * @return Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
     * 
     */
    public Optional<Output<List<String>>> subjectAlternativeNames() {
        return Optional.ofNullable(this.subjectAlternativeNames);
    }

    /**
     * A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    private CertificateState() {}

    private CertificateState(CertificateState $) {
        this.arn = $.arn;
        this.createdAt = $.createdAt;
        this.domainName = $.domainName;
        this.domainValidationOptions = $.domainValidationOptions;
        this.name = $.name;
        this.subjectAlternativeNames = $.subjectAlternativeNames;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateState $;

        public Builder() {
            $ = new CertificateState();
        }

        public Builder(CertificateState defaults) {
            $ = new CertificateState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the lightsail certificate.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the lightsail certificate.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param createdAt The timestamp when the instance was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The timestamp when the instance was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param domainName A domain name for which the certificate should be issued.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName A domain name for which the certificate should be issued.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param domainValidationOptions Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g., if SANs are defined.
         * 
         * @return builder
         * 
         */
        public Builder domainValidationOptions(@Nullable Output<List<CertificateDomainValidationOptionArgs>> domainValidationOptions) {
            $.domainValidationOptions = domainValidationOptions;
            return this;
        }

        /**
         * @param domainValidationOptions Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g., if SANs are defined.
         * 
         * @return builder
         * 
         */
        public Builder domainValidationOptions(List<CertificateDomainValidationOptionArgs> domainValidationOptions) {
            return domainValidationOptions(Output.of(domainValidationOptions));
        }

        /**
         * @param domainValidationOptions Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g., if SANs are defined.
         * 
         * @return builder
         * 
         */
        public Builder domainValidationOptions(CertificateDomainValidationOptionArgs... domainValidationOptions) {
            return domainValidationOptions(List.of(domainValidationOptions));
        }

        /**
         * @param name The name of the Lightsail load balancer.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Lightsail load balancer.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param subjectAlternativeNames Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
         * 
         * @return builder
         * 
         */
        public Builder subjectAlternativeNames(@Nullable Output<List<String>> subjectAlternativeNames) {
            $.subjectAlternativeNames = subjectAlternativeNames;
            return this;
        }

        /**
         * @param subjectAlternativeNames Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
         * 
         * @return builder
         * 
         */
        public Builder subjectAlternativeNames(List<String> subjectAlternativeNames) {
            return subjectAlternativeNames(Output.of(subjectAlternativeNames));
        }

        /**
         * @param subjectAlternativeNames Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
         * 
         * @return builder
         * 
         */
        public Builder subjectAlternativeNames(String... subjectAlternativeNames) {
            return subjectAlternativeNames(List.of(subjectAlternativeNames));
        }

        /**
         * @param tags A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        public CertificateState build() {
            return $;
        }
    }

}
