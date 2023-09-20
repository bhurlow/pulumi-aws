// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transfer;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.transfer.CertificateArgs;
import com.pulumi.aws.transfer.inputs.CertificateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a AWS Transfer AS2 Certificate resource.
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.transfer.Certificate;
 * import com.pulumi.aws.transfer.CertificateArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Certificate(&#34;example&#34;, CertificateArgs.builder()        
 *             .certificate(Files.readString(Paths.get(String.format(&#34;%s/example.com/example.crt&#34;, path.module()))))
 *             .certificateChain(Files.readString(Paths.get(String.format(&#34;%s/example.com/ca.crt&#34;, path.module()))))
 *             .privateKey(Files.readString(Paths.get(String.format(&#34;%s/example.com/example.key&#34;, path.module()))))
 *             .description(&#34;example&#34;)
 *             .usage(&#34;SIGNING&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Transfer AS2 Certificate using the `certificate_id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:transfer/certificate:Certificate example c-4221a88afd5f4362a
 * ```
 * 
 */
@ResourceType(type="aws:transfer/certificate:Certificate")
public class Certificate extends com.pulumi.resources.CustomResource {
    /**
     * An date when the certificate becomes active
     * 
     */
    @Export(name="activeDate", refs={String.class}, tree="[0]")
    private Output<String> activeDate;

    /**
     * @return An date when the certificate becomes active
     * 
     */
    public Output<String> activeDate() {
        return this.activeDate;
    }
    /**
     * The ARN of the certificate
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the certificate
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The valid certificate file required for the transfer.
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return The valid certificate file required for the transfer.
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * The optional list of certificate that make up the chain for the certificate that is being imported.
     * 
     */
    @Export(name="certificateChain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certificateChain;

    /**
     * @return The optional list of certificate that make up the chain for the certificate that is being imported.
     * 
     */
    public Output<Optional<String>> certificateChain() {
        return Codegen.optional(this.certificateChain);
    }
    /**
     * The unique identifier for the AS2 certificate
     * 
     */
    @Export(name="certificateId", refs={String.class}, tree="[0]")
    private Output<String> certificateId;

    /**
     * @return The unique identifier for the AS2 certificate
     * 
     */
    public Output<String> certificateId() {
        return this.certificateId;
    }
    /**
     * A short description that helps identify the certificate.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A short description that helps identify the certificate.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * An date when the certificate becomes inactive
     * 
     */
    @Export(name="inactiveDate", refs={String.class}, tree="[0]")
    private Output<String> inactiveDate;

    /**
     * @return An date when the certificate becomes inactive
     * 
     */
    public Output<String> inactiveDate() {
        return this.inactiveDate;
    }
    /**
     * The private key associated with the certificate being imported.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privateKey;

    /**
     * @return The private key associated with the certificate being imported.
     * 
     */
    public Output<Optional<String>> privateKey() {
        return Codegen.optional(this.privateKey);
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * Specifies if a certificate is being used for signing or encryption. The valid values are SIGNING and ENCRYPTION.
     * 
     */
    @Export(name="usage", refs={String.class}, tree="[0]")
    private Output<String> usage;

    /**
     * @return Specifies if a certificate is being used for signing or encryption. The valid values are SIGNING and ENCRYPTION.
     * 
     */
    public Output<String> usage() {
        return this.usage;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Certificate(String name) {
        this(name, CertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Certificate(String name, CertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Certificate(String name, CertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:transfer/certificate:Certificate", name, args == null ? CertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Certificate(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:transfer/certificate:Certificate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "certificate",
                "certificateChain",
                "privateKey",
                "tagsAll"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Certificate get(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Certificate(name, id, state, options);
    }
}
