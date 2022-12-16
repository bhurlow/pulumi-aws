// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.networkmanager;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.networkmanager.CoreNetworkArgs;
import com.pulumi.aws.networkmanager.inputs.CoreNetworkState;
import com.pulumi.aws.networkmanager.outputs.CoreNetworkEdge;
import com.pulumi.aws.networkmanager.outputs.CoreNetworkSegment;
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
 * Provides a core network resource.
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.networkmanager.CoreNetwork;
 * import com.pulumi.aws.networkmanager.CoreNetworkArgs;
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
 *         var example = new CoreNetwork(&#34;example&#34;, CoreNetworkArgs.builder()        
 *             .globalNetworkId(aws_networkmanager_global_network.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With description
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.networkmanager.CoreNetwork;
 * import com.pulumi.aws.networkmanager.CoreNetworkArgs;
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
 *         var example = new CoreNetwork(&#34;example&#34;, CoreNetworkArgs.builder()        
 *             .globalNetworkId(aws_networkmanager_global_network.example().id())
 *             .description(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With policy document
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.networkmanager.CoreNetwork;
 * import com.pulumi.aws.networkmanager.CoreNetworkArgs;
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
 *         var example = new CoreNetwork(&#34;example&#34;, CoreNetworkArgs.builder()        
 *             .globalNetworkId(aws_networkmanager_global_network.example().id())
 *             .policyDocument(data.aws_networkmanager_core_network_policy_document().example().json())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With tags
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.networkmanager.CoreNetwork;
 * import com.pulumi.aws.networkmanager.CoreNetworkArgs;
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
 *         var example = new CoreNetwork(&#34;example&#34;, CoreNetworkArgs.builder()        
 *             .globalNetworkId(aws_networkmanager_global_network.example().id())
 *             .tags(Map.of(&#34;hello&#34;, &#34;world&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_networkmanager_core_network` can be imported using the core network ID, e.g.
 * 
 * ```sh
 *  $ pulumi import aws:networkmanager/coreNetwork:CoreNetwork example core-network-0d47f6t230mz46dy4
 * ```
 * 
 */
@ResourceType(type="aws:networkmanager/coreNetwork:CoreNetwork")
public class CoreNetwork extends com.pulumi.resources.CustomResource {
    /**
     * Core Network Amazon Resource Name (ARN).
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Core Network Amazon Resource Name (ARN).
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Timestamp when a core network was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Timestamp when a core network was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Description of the Core Network.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the Core Network.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * One or more blocks detailing the edges within a core network. Detailed below.
     * 
     */
    @Export(name="edges", refs={List.class,CoreNetworkEdge.class}, tree="[0,1]")
    private Output<List<CoreNetworkEdge>> edges;

    /**
     * @return One or more blocks detailing the edges within a core network. Detailed below.
     * 
     */
    public Output<List<CoreNetworkEdge>> edges() {
        return this.edges;
    }
    /**
     * The ID of the global network that a core network will be a part of.
     * 
     */
    @Export(name="globalNetworkId", refs={String.class}, tree="[0]")
    private Output<String> globalNetworkId;

    /**
     * @return The ID of the global network that a core network will be a part of.
     * 
     */
    public Output<String> globalNetworkId() {
        return this.globalNetworkId;
    }
    /**
     * Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
     * 
     */
    @Export(name="policyDocument", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> policyDocument;

    /**
     * @return Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
     * 
     */
    public Output<Optional<String>> policyDocument() {
        return Codegen.optional(this.policyDocument);
    }
    /**
     * One or more blocks detailing the segments within a core network. Detailed below.
     * 
     */
    @Export(name="segments", refs={List.class,CoreNetworkSegment.class}, tree="[0,1]")
    private Output<List<CoreNetworkSegment>> segments;

    /**
     * @return One or more blocks detailing the segments within a core network. Detailed below.
     * 
     */
    public Output<List<CoreNetworkSegment>> segments() {
        return this.segments;
    }
    /**
     * Current state of a core network.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Current state of a core network.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Key-value tags for the Core Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value tags for the Core Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
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
    public CoreNetwork(String name) {
        this(name, CoreNetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CoreNetwork(String name, CoreNetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CoreNetwork(String name, CoreNetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:networkmanager/coreNetwork:CoreNetwork", name, args == null ? CoreNetworkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CoreNetwork(String name, Output<String> id, @Nullable CoreNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:networkmanager/coreNetwork:CoreNetwork", name, state, makeResourceOptions(options, id));
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
    public static CoreNetwork get(String name, Output<String> id, @Nullable CoreNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CoreNetwork(name, id, state, options);
    }
}
