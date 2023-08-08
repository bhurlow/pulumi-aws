// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.route53.CidrLocationArgs;
import com.pulumi.aws.route53.inputs.CidrLocationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Route53 CIDR location resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53.CidrCollection;
 * import com.pulumi.aws.route53.CidrLocation;
 * import com.pulumi.aws.route53.CidrLocationArgs;
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
 *         var exampleCidrCollection = new CidrCollection(&#34;exampleCidrCollection&#34;);
 * 
 *         var exampleCidrLocation = new CidrLocation(&#34;exampleCidrLocation&#34;, CidrLocationArgs.builder()        
 *             .cidrCollectionId(exampleCidrCollection.id())
 *             .cidrBlocks(            
 *                 &#34;200.5.3.0/24&#34;,
 *                 &#34;200.6.3.0/24&#34;)
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
 *  to = aws_route53_cidr_location.example
 * 
 *  id = &#34;9ac32814-3e67-0932-6048-8d779cc6f511,office&#34; } Using `pulumi import`, import CIDR locations using their the CIDR collection ID and location name. For exampleconsole % pulumi import aws_route53_cidr_location.example 9ac32814-3e67-0932-6048-8d779cc6f511,office
 * 
 */
@ResourceType(type="aws:route53/cidrLocation:CidrLocation")
public class CidrLocation extends com.pulumi.resources.CustomResource {
    /**
     * CIDR blocks for the location.
     * 
     */
    @Export(name="cidrBlocks", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> cidrBlocks;

    /**
     * @return CIDR blocks for the location.
     * 
     */
    public Output<List<String>> cidrBlocks() {
        return this.cidrBlocks;
    }
    /**
     * The ID of the CIDR collection to update.
     * 
     */
    @Export(name="cidrCollectionId", refs={String.class}, tree="[0]")
    private Output<String> cidrCollectionId;

    /**
     * @return The ID of the CIDR collection to update.
     * 
     */
    public Output<String> cidrCollectionId() {
        return this.cidrCollectionId;
    }
    /**
     * Name for the CIDR location.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name for the CIDR location.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CidrLocation(String name) {
        this(name, CidrLocationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CidrLocation(String name, CidrLocationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CidrLocation(String name, CidrLocationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/cidrLocation:CidrLocation", name, args == null ? CidrLocationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CidrLocation(String name, Output<String> id, @Nullable CidrLocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/cidrLocation:CidrLocation", name, state, makeResourceOptions(options, id));
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
    public static CidrLocation get(String name, Output<String> id, @Nullable CidrLocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CidrLocation(name, id, state, options);
    }
}
