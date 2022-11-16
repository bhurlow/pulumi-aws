// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.location;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.location.RouteCalculationArgs;
import com.pulumi.aws.location.inputs.RouteCalculationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Location Service Route Calculator.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.location.RouteCalculation;
 * import com.pulumi.aws.location.RouteCalculationArgs;
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
 *         var example = new RouteCalculation(&#34;example&#34;, RouteCalculationArgs.builder()        
 *             .calculatorName(&#34;example&#34;)
 *             .dataSource(&#34;Here&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_location_route_calculator` can be imported using the route calculator name, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:location/routeCalculation:RouteCalculation example example
 * ```
 * 
 */
@ResourceType(type="aws:location/routeCalculation:RouteCalculation")
public class RouteCalculation extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN when you specify a resource across AWS.
     * 
     */
    @Export(name="calculatorArn", type=String.class, parameters={})
    private Output<String> calculatorArn;

    /**
     * @return The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN when you specify a resource across AWS.
     * 
     */
    public Output<String> calculatorArn() {
        return this.calculatorArn;
    }
    /**
     * The name of the route calculator resource.
     * 
     */
    @Export(name="calculatorName", type=String.class, parameters={})
    private Output<String> calculatorName;

    /**
     * @return The name of the route calculator resource.
     * 
     */
    public Output<String> calculatorName() {
        return this.calculatorName;
    }
    /**
     * The timestamp for when the route calculator resource was created in ISO 8601 format.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The timestamp for when the route calculator resource was created in ISO 8601 format.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Specifies the data provider of traffic and road network data.
     * 
     */
    @Export(name="dataSource", type=String.class, parameters={})
    private Output<String> dataSource;

    /**
     * @return Specifies the data provider of traffic and road network data.
     * 
     */
    public Output<String> dataSource() {
        return this.dataSource;
    }
    /**
     * The optional description for the route calculator resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The optional description for the route calculator resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Key-value tags for the route calculator. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value tags for the route calculator. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * The timestamp for when the route calculator resource was last update in ISO 8601.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
    private Output<String> updateTime;

    /**
     * @return The timestamp for when the route calculator resource was last update in ISO 8601.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RouteCalculation(String name) {
        this(name, RouteCalculationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RouteCalculation(String name, RouteCalculationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RouteCalculation(String name, RouteCalculationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:location/routeCalculation:RouteCalculation", name, args == null ? RouteCalculationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RouteCalculation(String name, Output<String> id, @Nullable RouteCalculationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:location/routeCalculation:RouteCalculation", name, state, makeResourceOptions(options, id));
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
    public static RouteCalculation get(String name, Output<String> id, @Nullable RouteCalculationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RouteCalculation(name, id, state, options);
    }
}
