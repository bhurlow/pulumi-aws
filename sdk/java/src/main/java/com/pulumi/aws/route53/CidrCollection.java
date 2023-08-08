// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.route53.CidrCollectionArgs;
import com.pulumi.aws.route53.inputs.CidrCollectionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Route53 CIDR collection resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53.CidrCollection;
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
 *         var example = new CidrCollection(&#34;example&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_route53_cidr_collection.example
 * 
 *  id = &#34;9ac32814-3e67-0932-6048-8d779cc6f511&#34; } Using `pulumi import`, import CIDR collections using their ID. For exampleconsole % pulumi import aws_route53_cidr_collection.example 9ac32814-3e67-0932-6048-8d779cc6f511
 * 
 */
@ResourceType(type="aws:route53/cidrCollection:CidrCollection")
public class CidrCollection extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the CIDR collection.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the CIDR collection.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Unique name for the CIDR collection.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name for the CIDR collection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The lastest version of the CIDR collection.
     * 
     */
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output<Integer> version;

    /**
     * @return The lastest version of the CIDR collection.
     * 
     */
    public Output<Integer> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CidrCollection(String name) {
        this(name, CidrCollectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CidrCollection(String name, @Nullable CidrCollectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CidrCollection(String name, @Nullable CidrCollectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/cidrCollection:CidrCollection", name, args == null ? CidrCollectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CidrCollection(String name, Output<String> id, @Nullable CidrCollectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/cidrCollection:CidrCollection", name, state, makeResourceOptions(options, id));
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
    public static CidrCollection get(String name, Output<String> id, @Nullable CidrCollectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CidrCollection(name, id, state, options);
    }
}
