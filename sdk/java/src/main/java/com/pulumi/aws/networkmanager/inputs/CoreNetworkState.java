// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.networkmanager.inputs;

import com.pulumi.aws.networkmanager.inputs.CoreNetworkEdgeArgs;
import com.pulumi.aws.networkmanager.inputs.CoreNetworkSegmentArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CoreNetworkState extends com.pulumi.resources.ResourceArgs {

    public static final CoreNetworkState Empty = new CoreNetworkState();

    /**
     * Core Network Amazon Resource Name (ARN).
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Core Network Amazon Resource Name (ARN).
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The base policy created by setting the `create_base_policy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `base_policy_region` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
     * 
     * @deprecated
     * Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
     * 
     */
    @Deprecated /* Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider. */
    @Import(name="basePolicyRegion")
    private @Nullable Output<String> basePolicyRegion;

    /**
     * @return The base policy created by setting the `create_base_policy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `base_policy_region` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
     * 
     * @deprecated
     * Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
     * 
     */
    @Deprecated /* Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider. */
    public Optional<Output<String>> basePolicyRegion() {
        return Optional.ofNullable(this.basePolicyRegion);
    }

    /**
     * A list of regions to add to the base policy. The base policy created by setting the `create_base_policy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `base_policy_regions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
     * 
     */
    @Import(name="basePolicyRegions")
    private @Nullable Output<List<String>> basePolicyRegions;

    /**
     * @return A list of regions to add to the base policy. The base policy created by setting the `create_base_policy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `base_policy_regions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
     * 
     */
    public Optional<Output<List<String>>> basePolicyRegions() {
        return Optional.ofNullable(this.basePolicyRegions);
    }

    /**
     * Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `aws.networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `aws.networkmanager.CoreNetworkPolicyAttachment` resource.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
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
     *     }
     * }
     * ```
     * 
     */
    @Import(name="createBasePolicy")
    private @Nullable Output<Boolean> createBasePolicy;

    /**
     * @return Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `aws.networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `aws.networkmanager.CoreNetworkPolicyAttachment` resource.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
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
     *     }
     * }
     * ```
     * 
     */
    public Optional<Output<Boolean>> createBasePolicy() {
        return Optional.ofNullable(this.createBasePolicy);
    }

    /**
     * Timestamp when a core network was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Timestamp when a core network was created.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Description of the Core Network.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the Core Network.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * One or more blocks detailing the edges within a core network. Detailed below.
     * 
     */
    @Import(name="edges")
    private @Nullable Output<List<CoreNetworkEdgeArgs>> edges;

    /**
     * @return One or more blocks detailing the edges within a core network. Detailed below.
     * 
     */
    public Optional<Output<List<CoreNetworkEdgeArgs>>> edges() {
        return Optional.ofNullable(this.edges);
    }

    /**
     * The ID of the global network that a core network will be a part of.
     * 
     */
    @Import(name="globalNetworkId")
    private @Nullable Output<String> globalNetworkId;

    /**
     * @return The ID of the global network that a core network will be a part of.
     * 
     */
    public Optional<Output<String>> globalNetworkId() {
        return Optional.ofNullable(this.globalNetworkId);
    }

    /**
     * One or more blocks detailing the segments within a core network. Detailed below.
     * 
     */
    @Import(name="segments")
    private @Nullable Output<List<CoreNetworkSegmentArgs>> segments;

    /**
     * @return One or more blocks detailing the segments within a core network. Detailed below.
     * 
     */
    public Optional<Output<List<CoreNetworkSegmentArgs>>> segments() {
        return Optional.ofNullable(this.segments);
    }

    /**
     * Current state of a core network.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return Current state of a core network.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * Key-value tags for the Core Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value tags for the Core Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    private CoreNetworkState() {}

    private CoreNetworkState(CoreNetworkState $) {
        this.arn = $.arn;
        this.basePolicyRegion = $.basePolicyRegion;
        this.basePolicyRegions = $.basePolicyRegions;
        this.createBasePolicy = $.createBasePolicy;
        this.createdAt = $.createdAt;
        this.description = $.description;
        this.edges = $.edges;
        this.globalNetworkId = $.globalNetworkId;
        this.segments = $.segments;
        this.state = $.state;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CoreNetworkState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CoreNetworkState $;

        public Builder() {
            $ = new CoreNetworkState();
        }

        public Builder(CoreNetworkState defaults) {
            $ = new CoreNetworkState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Core Network Amazon Resource Name (ARN).
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Core Network Amazon Resource Name (ARN).
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param basePolicyRegion The base policy created by setting the `create_base_policy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `base_policy_region` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
         * 
         * @return builder
         * 
         * @deprecated
         * Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
         * 
         */
        @Deprecated /* Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider. */
        public Builder basePolicyRegion(@Nullable Output<String> basePolicyRegion) {
            $.basePolicyRegion = basePolicyRegion;
            return this;
        }

        /**
         * @param basePolicyRegion The base policy created by setting the `create_base_policy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `base_policy_region` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
         * 
         * @return builder
         * 
         * @deprecated
         * Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
         * 
         */
        @Deprecated /* Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider. */
        public Builder basePolicyRegion(String basePolicyRegion) {
            return basePolicyRegion(Output.of(basePolicyRegion));
        }

        /**
         * @param basePolicyRegions A list of regions to add to the base policy. The base policy created by setting the `create_base_policy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `base_policy_regions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
         * 
         * @return builder
         * 
         */
        public Builder basePolicyRegions(@Nullable Output<List<String>> basePolicyRegions) {
            $.basePolicyRegions = basePolicyRegions;
            return this;
        }

        /**
         * @param basePolicyRegions A list of regions to add to the base policy. The base policy created by setting the `create_base_policy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `base_policy_regions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
         * 
         * @return builder
         * 
         */
        public Builder basePolicyRegions(List<String> basePolicyRegions) {
            return basePolicyRegions(Output.of(basePolicyRegions));
        }

        /**
         * @param basePolicyRegions A list of regions to add to the base policy. The base policy created by setting the `create_base_policy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `base_policy_regions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
         * 
         * @return builder
         * 
         */
        public Builder basePolicyRegions(String... basePolicyRegions) {
            return basePolicyRegions(List.of(basePolicyRegions));
        }

        /**
         * @param createBasePolicy Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `aws.networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `aws.networkmanager.CoreNetworkPolicyAttachment` resource.
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
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
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder createBasePolicy(@Nullable Output<Boolean> createBasePolicy) {
            $.createBasePolicy = createBasePolicy;
            return this;
        }

        /**
         * @param createBasePolicy Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `aws.networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `aws.networkmanager.CoreNetworkPolicyAttachment` resource.
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
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
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder createBasePolicy(Boolean createBasePolicy) {
            return createBasePolicy(Output.of(createBasePolicy));
        }

        /**
         * @param createdAt Timestamp when a core network was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Timestamp when a core network was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param description Description of the Core Network.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the Core Network.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param edges One or more blocks detailing the edges within a core network. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder edges(@Nullable Output<List<CoreNetworkEdgeArgs>> edges) {
            $.edges = edges;
            return this;
        }

        /**
         * @param edges One or more blocks detailing the edges within a core network. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder edges(List<CoreNetworkEdgeArgs> edges) {
            return edges(Output.of(edges));
        }

        /**
         * @param edges One or more blocks detailing the edges within a core network. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder edges(CoreNetworkEdgeArgs... edges) {
            return edges(List.of(edges));
        }

        /**
         * @param globalNetworkId The ID of the global network that a core network will be a part of.
         * 
         * @return builder
         * 
         */
        public Builder globalNetworkId(@Nullable Output<String> globalNetworkId) {
            $.globalNetworkId = globalNetworkId;
            return this;
        }

        /**
         * @param globalNetworkId The ID of the global network that a core network will be a part of.
         * 
         * @return builder
         * 
         */
        public Builder globalNetworkId(String globalNetworkId) {
            return globalNetworkId(Output.of(globalNetworkId));
        }

        /**
         * @param segments One or more blocks detailing the segments within a core network. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder segments(@Nullable Output<List<CoreNetworkSegmentArgs>> segments) {
            $.segments = segments;
            return this;
        }

        /**
         * @param segments One or more blocks detailing the segments within a core network. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder segments(List<CoreNetworkSegmentArgs> segments) {
            return segments(Output.of(segments));
        }

        /**
         * @param segments One or more blocks detailing the segments within a core network. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder segments(CoreNetworkSegmentArgs... segments) {
            return segments(List.of(segments));
        }

        /**
         * @param state Current state of a core network.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state Current state of a core network.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param tags Key-value tags for the Core Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value tags for the Core Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        public CoreNetworkState build() {
            return $;
        }
    }

}
