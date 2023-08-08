// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appsync;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.appsync.DomainNameArgs;
import com.pulumi.aws.appsync.inputs.DomainNameState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an AppSync Domain Name.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.appsync.DomainName;
 * import com.pulumi.aws.appsync.DomainNameArgs;
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
 *         var example = new DomainName(&#34;example&#34;, DomainNameArgs.builder()        
 *             .domainName(&#34;api.example.com&#34;)
 *             .certificateArn(aws_acm_certificate.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_appsync_domain_name.example
 * 
 *  id = &#34;example.com&#34; } Using `pulumi import`, import `aws_appsync_domain_name` using the AppSync domain name. For exampleconsole % pulumi import aws_appsync_domain_name.example example.com
 * 
 */
@ResourceType(type="aws:appsync/domainName:DomainName")
public class DomainName extends com.pulumi.resources.CustomResource {
    /**
     * Domain name that AppSync provides.
     * 
     */
    @Export(name="appsyncDomainName", refs={String.class}, tree="[0]")
    private Output<String> appsyncDomainName;

    /**
     * @return Domain name that AppSync provides.
     * 
     */
    public Output<String> appsyncDomainName() {
        return this.appsyncDomainName;
    }
    /**
     * ARN of the certificate. This can be an Certificate Manager (ACM) certificate or an Identity and Access Management (IAM) server certificate. The certifiacte must reside in us-east-1.
     * 
     */
    @Export(name="certificateArn", refs={String.class}, tree="[0]")
    private Output<String> certificateArn;

    /**
     * @return ARN of the certificate. This can be an Certificate Manager (ACM) certificate or an Identity and Access Management (IAM) server certificate. The certifiacte must reside in us-east-1.
     * 
     */
    public Output<String> certificateArn() {
        return this.certificateArn;
    }
    /**
     * A description of the Domain Name.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the Domain Name.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Domain name.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return Domain name.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * ID of your Amazon Route 53 hosted zone.
     * 
     */
    @Export(name="hostedZoneId", refs={String.class}, tree="[0]")
    private Output<String> hostedZoneId;

    /**
     * @return ID of your Amazon Route 53 hosted zone.
     * 
     */
    public Output<String> hostedZoneId() {
        return this.hostedZoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DomainName(String name) {
        this(name, DomainNameArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DomainName(String name, DomainNameArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DomainName(String name, DomainNameArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appsync/domainName:DomainName", name, args == null ? DomainNameArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DomainName(String name, Output<String> id, @Nullable DomainNameState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appsync/domainName:DomainName", name, state, makeResourceOptions(options, id));
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
    public static DomainName get(String name, Output<String> id, @Nullable DomainNameState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DomainName(name, id, state, options);
    }
}
