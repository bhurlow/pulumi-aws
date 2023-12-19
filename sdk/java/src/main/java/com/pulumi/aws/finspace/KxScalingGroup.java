// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.finspace;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.finspace.KxScalingGroupArgs;
import com.pulumi.aws.finspace.inputs.KxScalingGroupState;
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
 * Resource for managing an AWS FinSpace Kx Scaling Group.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.finspace.KxScalingGroup;
 * import com.pulumi.aws.finspace.KxScalingGroupArgs;
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
 *         var example = new KxScalingGroup(&#34;example&#34;, KxScalingGroupArgs.builder()        
 *             .environmentId(aws_finspace_kx_environment.example().id())
 *             .availabilityZoneId(&#34;use1-az2&#34;)
 *             .hostType(&#34;kx.sg.4xlarge&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import an AWS FinSpace Kx Scaling Group using the `id` (environment ID and scaling group name, comma-delimited). For example:
 * 
 * ```sh
 *  $ pulumi import aws:finspace/kxScalingGroup:KxScalingGroup example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-scalinggroup
 * ```
 * 
 */
@ResourceType(type="aws:finspace/kxScalingGroup:KxScalingGroup")
public class KxScalingGroup extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) identifier of the KX Scaling Group.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) identifier of the KX Scaling Group.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The availability zone identifiers for the requested regions.
     * 
     */
    @Export(name="availabilityZoneId", refs={String.class}, tree="[0]")
    private Output<String> availabilityZoneId;

    /**
     * @return The availability zone identifiers for the requested regions.
     * 
     */
    public Output<String> availabilityZoneId() {
        return this.availabilityZoneId;
    }
    /**
     * The list of Managed kdb clusters that are currently active in the given scaling group.
     * 
     */
    @Export(name="clusters", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> clusters;

    /**
     * @return The list of Managed kdb clusters that are currently active in the given scaling group.
     * 
     */
    public Output<List<String>> clusters() {
        return this.clusters;
    }
    /**
     * The timestamp at which the scaling group was created in FinSpace. The value is determined as epoch time in milliseconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
     * 
     */
    @Export(name="createdTimestamp", refs={String.class}, tree="[0]")
    private Output<String> createdTimestamp;

    /**
     * @return The timestamp at which the scaling group was created in FinSpace. The value is determined as epoch time in milliseconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
     * 
     */
    public Output<String> createdTimestamp() {
        return this.createdTimestamp;
    }
    /**
     * A unique identifier for the kdb environment, where you want to create the scaling group.
     * 
     */
    @Export(name="environmentId", refs={String.class}, tree="[0]")
    private Output<String> environmentId;

    /**
     * @return A unique identifier for the kdb environment, where you want to create the scaling group.
     * 
     */
    public Output<String> environmentId() {
        return this.environmentId;
    }
    /**
     * The memory and CPU capabilities of the scaling group host on which FinSpace Managed kdb clusters will be placed.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="hostType", refs={String.class}, tree="[0]")
    private Output<String> hostType;

    /**
     * @return The memory and CPU capabilities of the scaling group host on which FinSpace Managed kdb clusters will be placed.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> hostType() {
        return this.hostType;
    }
    /**
     * Last timestamp at which the scaling group was updated in FinSpace. Value determined as epoch time in seconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000.
     * 
     */
    @Export(name="lastModifiedTimestamp", refs={String.class}, tree="[0]")
    private Output<String> lastModifiedTimestamp;

    /**
     * @return Last timestamp at which the scaling group was updated in FinSpace. Value determined as epoch time in seconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000.
     * 
     */
    public Output<String> lastModifiedTimestamp() {
        return this.lastModifiedTimestamp;
    }
    /**
     * Unique name for the scaling group that you want to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name for the scaling group that you want to create.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The status of scaling group.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of scaling group.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The error message when a failed state occurs.
     * 
     */
    @Export(name="statusReason", refs={String.class}, tree="[0]")
    private Output<String> statusReason;

    /**
     * @return The error message when a failed state occurs.
     * 
     */
    public Output<String> statusReason() {
        return this.statusReason;
    }
    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. You can add up to 50 tags to a scaling group.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. You can add up to 50 tags to a scaling group.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KxScalingGroup(String name) {
        this(name, KxScalingGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KxScalingGroup(String name, KxScalingGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KxScalingGroup(String name, KxScalingGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:finspace/kxScalingGroup:KxScalingGroup", name, args == null ? KxScalingGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KxScalingGroup(String name, Output<String> id, @Nullable KxScalingGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:finspace/kxScalingGroup:KxScalingGroup", name, state, makeResourceOptions(options, id));
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
    public static KxScalingGroup get(String name, Output<String> id, @Nullable KxScalingGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KxScalingGroup(name, id, state, options);
    }
}
