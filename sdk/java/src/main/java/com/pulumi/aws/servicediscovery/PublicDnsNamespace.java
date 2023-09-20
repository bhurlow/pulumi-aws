// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.servicediscovery;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.servicediscovery.PublicDnsNamespaceArgs;
import com.pulumi.aws.servicediscovery.inputs.PublicDnsNamespaceState;
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
 * Provides a Service Discovery Public DNS Namespace resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.servicediscovery.PublicDnsNamespace;
 * import com.pulumi.aws.servicediscovery.PublicDnsNamespaceArgs;
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
 *         var example = new PublicDnsNamespace(&#34;example&#34;, PublicDnsNamespaceArgs.builder()        
 *             .description(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Service Discovery Public DNS Namespace using the namespace ID. For example:
 * 
 * ```sh
 *  $ pulumi import aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace example 0123456789
 * ```
 * 
 */
@ResourceType(type="aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace")
public class PublicDnsNamespace extends com.pulumi.resources.CustomResource {
    /**
     * The ARN that Amazon Route 53 assigns to the namespace when you create it.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN that Amazon Route 53 assigns to the namespace when you create it.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The description that you specify for the namespace when you create it.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description that you specify for the namespace when you create it.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
     * 
     */
    @Export(name="hostedZone", refs={String.class}, tree="[0]")
    private Output<String> hostedZone;

    /**
     * @return The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
     * 
     */
    public Output<String> hostedZone() {
        return this.hostedZone;
    }
    /**
     * The name of the namespace.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the namespace.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A map of tags to assign to the namespace. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the namespace. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PublicDnsNamespace(String name) {
        this(name, PublicDnsNamespaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PublicDnsNamespace(String name, @Nullable PublicDnsNamespaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PublicDnsNamespace(String name, @Nullable PublicDnsNamespaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, args == null ? PublicDnsNamespaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PublicDnsNamespace(String name, Output<String> id, @Nullable PublicDnsNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
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
    public static PublicDnsNamespace get(String name, Output<String> id, @Nullable PublicDnsNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PublicDnsNamespace(name, id, state, options);
    }
}
