// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.redshift.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HsmClientCertificateState extends com.pulumi.resources.ResourceArgs {

    public static final HsmClientCertificateState Empty = new HsmClientCertificateState();

    /**
     * Amazon Resource Name (ARN) of the Hsm Client Certificate.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the Hsm Client Certificate.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The identifier of the HSM client certificate.
     * 
     */
    @Import(name="hsmClientCertificateIdentifier")
    private @Nullable Output<String> hsmClientCertificateIdentifier;

    /**
     * @return The identifier of the HSM client certificate.
     * 
     */
    public Optional<Output<String>> hsmClientCertificateIdentifier() {
        return Optional.ofNullable(this.hsmClientCertificateIdentifier);
    }

    /**
     * The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
     * 
     */
    @Import(name="hsmClientCertificatePublicKey")
    private @Nullable Output<String> hsmClientCertificatePublicKey;

    /**
     * @return The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
     * 
     */
    public Optional<Output<String>> hsmClientCertificatePublicKey() {
        return Optional.ofNullable(this.hsmClientCertificatePublicKey);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private HsmClientCertificateState() {}

    private HsmClientCertificateState(HsmClientCertificateState $) {
        this.arn = $.arn;
        this.hsmClientCertificateIdentifier = $.hsmClientCertificateIdentifier;
        this.hsmClientCertificatePublicKey = $.hsmClientCertificatePublicKey;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HsmClientCertificateState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HsmClientCertificateState $;

        public Builder() {
            $ = new HsmClientCertificateState();
        }

        public Builder(HsmClientCertificateState defaults) {
            $ = new HsmClientCertificateState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the Hsm Client Certificate.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the Hsm Client Certificate.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param hsmClientCertificateIdentifier The identifier of the HSM client certificate.
         * 
         * @return builder
         * 
         */
        public Builder hsmClientCertificateIdentifier(@Nullable Output<String> hsmClientCertificateIdentifier) {
            $.hsmClientCertificateIdentifier = hsmClientCertificateIdentifier;
            return this;
        }

        /**
         * @param hsmClientCertificateIdentifier The identifier of the HSM client certificate.
         * 
         * @return builder
         * 
         */
        public Builder hsmClientCertificateIdentifier(String hsmClientCertificateIdentifier) {
            return hsmClientCertificateIdentifier(Output.of(hsmClientCertificateIdentifier));
        }

        /**
         * @param hsmClientCertificatePublicKey The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
         * 
         * @return builder
         * 
         */
        public Builder hsmClientCertificatePublicKey(@Nullable Output<String> hsmClientCertificatePublicKey) {
            $.hsmClientCertificatePublicKey = hsmClientCertificatePublicKey;
            return this;
        }

        /**
         * @param hsmClientCertificatePublicKey The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
         * 
         * @return builder
         * 
         */
        public Builder hsmClientCertificatePublicKey(String hsmClientCertificatePublicKey) {
            return hsmClientCertificatePublicKey(Output.of(hsmClientCertificatePublicKey));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public HsmClientCertificateState build() {
            return $;
        }
    }

}
