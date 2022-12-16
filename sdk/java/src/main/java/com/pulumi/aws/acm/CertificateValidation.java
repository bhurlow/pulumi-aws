// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.acm;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.acm.CertificateValidationArgs;
import com.pulumi.aws.acm.inputs.CertificateValidationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource represents a successful validation of an ACM certificate in concert
 * with other resources.
 * 
 * Most commonly, this resource is used together with `aws.route53.Record` and
 * `aws.acm.Certificate` to request a DNS validated certificate,
 * deploy the required validation records and wait for validation to complete.
 * 
 * &gt; **WARNING:** This resource implements a part of the validation workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.
 * 
 * ## Example Usage
 * 
 * {{% //examples %}}
 * 
 */
@ResourceType(type="aws:acm/certificateValidation:CertificateValidation")
public class CertificateValidation extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the certificate that is being validated.
     * 
     */
    @Export(name="certificateArn", refs={String.class}, tree="[0]")
    private Output<String> certificateArn;

    /**
     * @return ARN of the certificate that is being validated.
     * 
     */
    public Output<String> certificateArn() {
        return this.certificateArn;
    }
    /**
     * List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
     * 
     */
    @Export(name="validationRecordFqdns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> validationRecordFqdns;

    /**
     * @return List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
     * 
     */
    public Output<Optional<List<String>>> validationRecordFqdns() {
        return Codegen.optional(this.validationRecordFqdns);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CertificateValidation(String name) {
        this(name, CertificateValidationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CertificateValidation(String name, CertificateValidationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CertificateValidation(String name, CertificateValidationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:acm/certificateValidation:CertificateValidation", name, args == null ? CertificateValidationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CertificateValidation(String name, Output<String> id, @Nullable CertificateValidationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:acm/certificateValidation:CertificateValidation", name, state, makeResourceOptions(options, id));
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
    public static CertificateValidation get(String name, Output<String> id, @Nullable CertificateValidationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CertificateValidation(name, id, state, options);
    }
}
