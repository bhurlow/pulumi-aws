// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.acmpca;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.acmpca.CertificateAuthorityCertificateArgs;
import com.pulumi.aws.acmpca.inputs.CertificateAuthorityCertificateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Associates a certificate with an AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority). An ACM PCA Certificate Authority is unable to issue certificates until it has a certificate associated with it. A root level ACM PCA Certificate Authority is able to self-sign its own root certificate.
 * 
 * ## Example Usage
 * 
 */
@ResourceType(type="aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate")
public class CertificateAuthorityCertificate extends com.pulumi.resources.CustomResource {
    /**
     * PEM-encoded certificate for the Certificate Authority.
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return PEM-encoded certificate for the Certificate Authority.
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * ARN of the Certificate Authority.
     * 
     */
    @Export(name="certificateAuthorityArn", refs={String.class}, tree="[0]")
    private Output<String> certificateAuthorityArn;

    /**
     * @return ARN of the Certificate Authority.
     * 
     */
    public Output<String> certificateAuthorityArn() {
        return this.certificateAuthorityArn;
    }
    /**
     * PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
     * 
     */
    @Export(name="certificateChain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certificateChain;

    /**
     * @return PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
     * 
     */
    public Output<Optional<String>> certificateChain() {
        return Codegen.optional(this.certificateChain);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CertificateAuthorityCertificate(String name) {
        this(name, CertificateAuthorityCertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CertificateAuthorityCertificate(String name, CertificateAuthorityCertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CertificateAuthorityCertificate(String name, CertificateAuthorityCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate", name, args == null ? CertificateAuthorityCertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CertificateAuthorityCertificate(String name, Output<String> id, @Nullable CertificateAuthorityCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static CertificateAuthorityCertificate get(String name, Output<String> id, @Nullable CertificateAuthorityCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CertificateAuthorityCertificate(name, id, state, options);
    }
}
