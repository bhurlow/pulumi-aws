// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.servicecatalog;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.servicecatalog.PortfolioArgs;
import com.pulumi.aws.servicecatalog.inputs.PortfolioState;
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
 * Provides a resource to create a Service Catalog Portfolio.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.servicecatalog.Portfolio;
 * import com.pulumi.aws.servicecatalog.PortfolioArgs;
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
 *         var portfolio = new Portfolio(&#34;portfolio&#34;, PortfolioArgs.builder()        
 *             .description(&#34;List of my organizations apps&#34;)
 *             .providerName(&#34;Brett&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Service Catalog Portfolios using the Service Catalog Portfolio `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:servicecatalog/portfolio:Portfolio testfolio port-12344321
 * ```
 * 
 */
@ResourceType(type="aws:servicecatalog/portfolio:Portfolio")
public class Portfolio extends com.pulumi.resources.CustomResource {
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    public Output<String> arn() {
        return this.arn;
    }
    @Export(name="createdTime", refs={String.class}, tree="[0]")
    private Output<String> createdTime;

    public Output<String> createdTime() {
        return this.createdTime;
    }
    /**
     * Description of the portfolio
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description of the portfolio
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The name of the portfolio.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the portfolio.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Name of the person or organization who owns the portfolio.
     * 
     */
    @Export(name="providerName", refs={String.class}, tree="[0]")
    private Output<String> providerName;

    /**
     * @return Name of the person or organization who owns the portfolio.
     * 
     */
    public Output<String> providerName() {
        return this.providerName;
    }
    /**
     * Tags to apply to the connection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Tags to apply to the connection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public Portfolio(String name) {
        this(name, PortfolioArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Portfolio(String name, PortfolioArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Portfolio(String name, PortfolioArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:servicecatalog/portfolio:Portfolio", name, args == null ? PortfolioArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Portfolio(String name, Output<String> id, @Nullable PortfolioState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:servicecatalog/portfolio:Portfolio", name, state, makeResourceOptions(options, id));
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
    public static Portfolio get(String name, Output<String> id, @Nullable PortfolioState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Portfolio(name, id, state, options);
    }
}
