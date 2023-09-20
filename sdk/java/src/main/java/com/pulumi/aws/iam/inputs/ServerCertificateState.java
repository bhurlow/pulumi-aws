// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.iam.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerCertificateState extends com.pulumi.resources.ResourceArgs {

    public static final ServerCertificateState Empty = new ServerCertificateState();

    /**
     * The Amazon Resource Name (ARN) specifying the server certificate.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) specifying the server certificate.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The contents of the public key certificate in
     * PEM-encoded format.
     * 
     */
    @Import(name="certificateBody")
    private @Nullable Output<String> certificateBody;

    /**
     * @return The contents of the public key certificate in
     * PEM-encoded format.
     * 
     */
    public Optional<Output<String>> certificateBody() {
        return Optional.ofNullable(this.certificateBody);
    }

    /**
     * The contents of the certificate chain.
     * This is typically a concatenation of the PEM-encoded public key certificates
     * of the chain.
     * 
     */
    @Import(name="certificateChain")
    private @Nullable Output<String> certificateChain;

    /**
     * @return The contents of the certificate chain.
     * This is typically a concatenation of the PEM-encoded public key certificates
     * of the chain.
     * 
     */
    public Optional<Output<String>> certificateChain() {
        return Optional.ofNullable(this.certificateChain);
    }

    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
     * 
     */
    @Import(name="expiration")
    private @Nullable Output<String> expiration;

    /**
     * @return Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
     * 
     */
    public Optional<Output<String>> expiration() {
        return Optional.ofNullable(this.expiration);
    }

    /**
     * The name of the Server Certificate. Do not include the
     * path in this value. If omitted, the provider will assign a random, unique name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Server Certificate. Do not include the
     * path in this value. If omitted, the provider will assign a random, unique name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * The IAM path for the server certificate.  If it is not
     * included, it defaults to a slash (/). If this certificate is for use with
     * AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
     * See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The IAM path for the server certificate.  If it is not
     * included, it defaults to a slash (/). If this certificate is for use with
     * AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
     * See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The contents of the private key in PEM-encoded format.
     * 
     */
    @Import(name="privateKey")
    private @Nullable Output<String> privateKey;

    /**
     * @return The contents of the private key in PEM-encoded format.
     * 
     */
    public Optional<Output<String>> privateKey() {
        return Optional.ofNullable(this.privateKey);
    }

    /**
     * Map of resource tags for the server certificate. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * &gt; **NOTE:** AWS performs behind-the-scenes modifications to some certificate files if they do not adhere to a specific format. These modifications will result in this provider forever believing that it needs to update the resources since the local and AWS file contents will not match after theses modifications occur. In order to prevent this from happening you must ensure that all your PEM-encoded files use UNIX line-breaks and that `certificate_body` contains only one certificate. All other certificates should go in `certificate_chain`. It is common for some Certificate Authorities to issue certificate files that have DOS line-breaks and that are actually multiple certificates concatenated together in order to form a full certificate chain.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of resource tags for the server certificate. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * &gt; **NOTE:** AWS performs behind-the-scenes modifications to some certificate files if they do not adhere to a specific format. These modifications will result in this provider forever believing that it needs to update the resources since the local and AWS file contents will not match after theses modifications occur. In order to prevent this from happening you must ensure that all your PEM-encoded files use UNIX line-breaks and that `certificate_body` contains only one certificate. All other certificates should go in `certificate_chain`. It is common for some Certificate Authorities to issue certificate files that have DOS line-breaks and that are actually multiple certificates concatenated together in order to form a full certificate chain.
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

    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
     * 
     */
    @Import(name="uploadDate")
    private @Nullable Output<String> uploadDate;

    /**
     * @return Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
     * 
     */
    public Optional<Output<String>> uploadDate() {
        return Optional.ofNullable(this.uploadDate);
    }

    private ServerCertificateState() {}

    private ServerCertificateState(ServerCertificateState $) {
        this.arn = $.arn;
        this.certificateBody = $.certificateBody;
        this.certificateChain = $.certificateChain;
        this.expiration = $.expiration;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.path = $.path;
        this.privateKey = $.privateKey;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.uploadDate = $.uploadDate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerCertificateState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerCertificateState $;

        public Builder() {
            $ = new ServerCertificateState();
        }

        public Builder(ServerCertificateState defaults) {
            $ = new ServerCertificateState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) specifying the server certificate.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) specifying the server certificate.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param certificateBody The contents of the public key certificate in
         * PEM-encoded format.
         * 
         * @return builder
         * 
         */
        public Builder certificateBody(@Nullable Output<String> certificateBody) {
            $.certificateBody = certificateBody;
            return this;
        }

        /**
         * @param certificateBody The contents of the public key certificate in
         * PEM-encoded format.
         * 
         * @return builder
         * 
         */
        public Builder certificateBody(String certificateBody) {
            return certificateBody(Output.of(certificateBody));
        }

        /**
         * @param certificateChain The contents of the certificate chain.
         * This is typically a concatenation of the PEM-encoded public key certificates
         * of the chain.
         * 
         * @return builder
         * 
         */
        public Builder certificateChain(@Nullable Output<String> certificateChain) {
            $.certificateChain = certificateChain;
            return this;
        }

        /**
         * @param certificateChain The contents of the certificate chain.
         * This is typically a concatenation of the PEM-encoded public key certificates
         * of the chain.
         * 
         * @return builder
         * 
         */
        public Builder certificateChain(String certificateChain) {
            return certificateChain(Output.of(certificateChain));
        }

        /**
         * @param expiration Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
         * 
         * @return builder
         * 
         */
        public Builder expiration(@Nullable Output<String> expiration) {
            $.expiration = expiration;
            return this;
        }

        /**
         * @param expiration Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
         * 
         * @return builder
         * 
         */
        public Builder expiration(String expiration) {
            return expiration(Output.of(expiration));
        }

        /**
         * @param name The name of the Server Certificate. Do not include the
         * path in this value. If omitted, the provider will assign a random, unique name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Server Certificate. Do not include the
         * path in this value. If omitted, the provider will assign a random, unique name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified
         * prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified
         * prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param path The IAM path for the server certificate.  If it is not
         * included, it defaults to a slash (/). If this certificate is for use with
         * AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
         * See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The IAM path for the server certificate.  If it is not
         * included, it defaults to a slash (/). If this certificate is for use with
         * AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
         * See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param privateKey The contents of the private key in PEM-encoded format.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(@Nullable Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey The contents of the private key in PEM-encoded format.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param tags Map of resource tags for the server certificate. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * &gt; **NOTE:** AWS performs behind-the-scenes modifications to some certificate files if they do not adhere to a specific format. These modifications will result in this provider forever believing that it needs to update the resources since the local and AWS file contents will not match after theses modifications occur. In order to prevent this from happening you must ensure that all your PEM-encoded files use UNIX line-breaks and that `certificate_body` contains only one certificate. All other certificates should go in `certificate_chain`. It is common for some Certificate Authorities to issue certificate files that have DOS line-breaks and that are actually multiple certificates concatenated together in order to form a full certificate chain.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of resource tags for the server certificate. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * &gt; **NOTE:** AWS performs behind-the-scenes modifications to some certificate files if they do not adhere to a specific format. These modifications will result in this provider forever believing that it needs to update the resources since the local and AWS file contents will not match after theses modifications occur. In order to prevent this from happening you must ensure that all your PEM-encoded files use UNIX line-breaks and that `certificate_body` contains only one certificate. All other certificates should go in `certificate_chain`. It is common for some Certificate Authorities to issue certificate files that have DOS line-breaks and that are actually multiple certificates concatenated together in order to form a full certificate chain.
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

        /**
         * @param uploadDate Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
         * 
         * @return builder
         * 
         */
        public Builder uploadDate(@Nullable Output<String> uploadDate) {
            $.uploadDate = uploadDate;
            return this;
        }

        /**
         * @param uploadDate Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
         * 
         * @return builder
         * 
         */
        public Builder uploadDate(String uploadDate) {
            return uploadDate(Output.of(uploadDate));
        }

        public ServerCertificateState build() {
            return $;
        }
    }

}
