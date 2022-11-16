// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.location;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.location.TrackerArgs;
import com.pulumi.aws.location.inputs.TrackerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Location Service Tracker.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.location.Tracker;
 * import com.pulumi.aws.location.TrackerArgs;
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
 *         var example = new Tracker(&#34;example&#34;, TrackerArgs.builder()        
 *             .trackerName(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_location_tracker` resources can be imported using the tracker name, e.g.
 * 
 * ```sh
 *  $ pulumi import aws:location/tracker:Tracker example example
 * ```
 * 
 */
@ResourceType(type="aws:location/tracker:Tracker")
public class Tracker extends com.pulumi.resources.CustomResource {
    /**
     * The timestamp for when the tracker resource was created in ISO 8601 format.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The timestamp for when the tracker resource was created in ISO 8601 format.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The optional description for the tracker resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The optional description for the tracker resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
     * 
     */
    @Export(name="kmsKeyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * The position filtering method of the tracker resource. Valid values: `TimeBased`, `DistanceBased`, `AccuracyBased`. Default: `TimeBased`.
     * 
     */
    @Export(name="positionFiltering", type=String.class, parameters={})
    private Output</* @Nullable */ String> positionFiltering;

    /**
     * @return The position filtering method of the tracker resource. Valid values: `TimeBased`, `DistanceBased`, `AccuracyBased`. Default: `TimeBased`.
     * 
     */
    public Output<Optional<String>> positionFiltering() {
        return Codegen.optional(this.positionFiltering);
    }
    /**
     * Key-value tags for the tracker. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value tags for the tracker. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The Amazon Resource Name (ARN) for the tracker resource. Used when you need to specify a resource across all AWS.
     * 
     */
    @Export(name="trackerArn", type=String.class, parameters={})
    private Output<String> trackerArn;

    /**
     * @return The Amazon Resource Name (ARN) for the tracker resource. Used when you need to specify a resource across all AWS.
     * 
     */
    public Output<String> trackerArn() {
        return this.trackerArn;
    }
    /**
     * The name of the tracker resource.
     * 
     */
    @Export(name="trackerName", type=String.class, parameters={})
    private Output<String> trackerName;

    /**
     * @return The name of the tracker resource.
     * 
     */
    public Output<String> trackerName() {
        return this.trackerName;
    }
    /**
     * The timestamp for when the tracker resource was last updated in ISO 8601 format.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
    private Output<String> updateTime;

    /**
     * @return The timestamp for when the tracker resource was last updated in ISO 8601 format.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Tracker(String name) {
        this(name, TrackerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Tracker(String name, TrackerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Tracker(String name, TrackerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:location/tracker:Tracker", name, args == null ? TrackerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Tracker(String name, Output<String> id, @Nullable TrackerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:location/tracker:Tracker", name, state, makeResourceOptions(options, id));
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
    public static Tracker get(String name, Output<String> id, @Nullable TrackerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Tracker(name, id, state, options);
    }
}
