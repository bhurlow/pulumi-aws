// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.bedrock.inputs;

import com.pulumi.aws.bedrock.inputs.ProvisionedModelThroughputTimeoutsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProvisionedModelThroughputState extends com.pulumi.resources.ResourceArgs {

    public static final ProvisionedModelThroughputState Empty = new ProvisionedModelThroughputState();

    /**
     * Commitment duration requested for the Provisioned Throughput. For custom models, you can purchase on-demand Provisioned Throughput by omitting this argument. Valid values: `OneMonth`, `SixMonths`.
     * 
     */
    @Import(name="commitmentDuration")
    private @Nullable Output<String> commitmentDuration;

    /**
     * @return Commitment duration requested for the Provisioned Throughput. For custom models, you can purchase on-demand Provisioned Throughput by omitting this argument. Valid values: `OneMonth`, `SixMonths`.
     * 
     */
    public Optional<Output<String>> commitmentDuration() {
        return Optional.ofNullable(this.commitmentDuration);
    }

    /**
     * ARN of the model to associate with this Provisioned Throughput.
     * 
     */
    @Import(name="modelArn")
    private @Nullable Output<String> modelArn;

    /**
     * @return ARN of the model to associate with this Provisioned Throughput.
     * 
     */
    public Optional<Output<String>> modelArn() {
        return Optional.ofNullable(this.modelArn);
    }

    /**
     * Number of model units to allocate. A model unit delivers a specific throughput level for the specified model.
     * 
     */
    @Import(name="modelUnits")
    private @Nullable Output<Integer> modelUnits;

    /**
     * @return Number of model units to allocate. A model unit delivers a specific throughput level for the specified model.
     * 
     */
    public Optional<Output<Integer>> modelUnits() {
        return Optional.ofNullable(this.modelUnits);
    }

    /**
     * The ARN of the Provisioned Throughput.
     * 
     */
    @Import(name="provisionedModelArn")
    private @Nullable Output<String> provisionedModelArn;

    /**
     * @return The ARN of the Provisioned Throughput.
     * 
     */
    public Optional<Output<String>> provisionedModelArn() {
        return Optional.ofNullable(this.provisionedModelArn);
    }

    /**
     * Unique name for this Provisioned Throughput.
     * 
     */
    @Import(name="provisionedModelName")
    private @Nullable Output<String> provisionedModelName;

    /**
     * @return Unique name for this Provisioned Throughput.
     * 
     */
    public Optional<Output<String>> provisionedModelName() {
        return Optional.ofNullable(this.provisionedModelName);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    @Import(name="timeouts")
    private @Nullable Output<ProvisionedModelThroughputTimeoutsArgs> timeouts;

    public Optional<Output<ProvisionedModelThroughputTimeoutsArgs>> timeouts() {
        return Optional.ofNullable(this.timeouts);
    }

    private ProvisionedModelThroughputState() {}

    private ProvisionedModelThroughputState(ProvisionedModelThroughputState $) {
        this.commitmentDuration = $.commitmentDuration;
        this.modelArn = $.modelArn;
        this.modelUnits = $.modelUnits;
        this.provisionedModelArn = $.provisionedModelArn;
        this.provisionedModelName = $.provisionedModelName;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.timeouts = $.timeouts;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProvisionedModelThroughputState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProvisionedModelThroughputState $;

        public Builder() {
            $ = new ProvisionedModelThroughputState();
        }

        public Builder(ProvisionedModelThroughputState defaults) {
            $ = new ProvisionedModelThroughputState(Objects.requireNonNull(defaults));
        }

        /**
         * @param commitmentDuration Commitment duration requested for the Provisioned Throughput. For custom models, you can purchase on-demand Provisioned Throughput by omitting this argument. Valid values: `OneMonth`, `SixMonths`.
         * 
         * @return builder
         * 
         */
        public Builder commitmentDuration(@Nullable Output<String> commitmentDuration) {
            $.commitmentDuration = commitmentDuration;
            return this;
        }

        /**
         * @param commitmentDuration Commitment duration requested for the Provisioned Throughput. For custom models, you can purchase on-demand Provisioned Throughput by omitting this argument. Valid values: `OneMonth`, `SixMonths`.
         * 
         * @return builder
         * 
         */
        public Builder commitmentDuration(String commitmentDuration) {
            return commitmentDuration(Output.of(commitmentDuration));
        }

        /**
         * @param modelArn ARN of the model to associate with this Provisioned Throughput.
         * 
         * @return builder
         * 
         */
        public Builder modelArn(@Nullable Output<String> modelArn) {
            $.modelArn = modelArn;
            return this;
        }

        /**
         * @param modelArn ARN of the model to associate with this Provisioned Throughput.
         * 
         * @return builder
         * 
         */
        public Builder modelArn(String modelArn) {
            return modelArn(Output.of(modelArn));
        }

        /**
         * @param modelUnits Number of model units to allocate. A model unit delivers a specific throughput level for the specified model.
         * 
         * @return builder
         * 
         */
        public Builder modelUnits(@Nullable Output<Integer> modelUnits) {
            $.modelUnits = modelUnits;
            return this;
        }

        /**
         * @param modelUnits Number of model units to allocate. A model unit delivers a specific throughput level for the specified model.
         * 
         * @return builder
         * 
         */
        public Builder modelUnits(Integer modelUnits) {
            return modelUnits(Output.of(modelUnits));
        }

        /**
         * @param provisionedModelArn The ARN of the Provisioned Throughput.
         * 
         * @return builder
         * 
         */
        public Builder provisionedModelArn(@Nullable Output<String> provisionedModelArn) {
            $.provisionedModelArn = provisionedModelArn;
            return this;
        }

        /**
         * @param provisionedModelArn The ARN of the Provisioned Throughput.
         * 
         * @return builder
         * 
         */
        public Builder provisionedModelArn(String provisionedModelArn) {
            return provisionedModelArn(Output.of(provisionedModelArn));
        }

        /**
         * @param provisionedModelName Unique name for this Provisioned Throughput.
         * 
         * @return builder
         * 
         */
        public Builder provisionedModelName(@Nullable Output<String> provisionedModelName) {
            $.provisionedModelName = provisionedModelName;
            return this;
        }

        /**
         * @param provisionedModelName Unique name for this Provisioned Throughput.
         * 
         * @return builder
         * 
         */
        public Builder provisionedModelName(String provisionedModelName) {
            return provisionedModelName(Output.of(provisionedModelName));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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

        public Builder timeouts(@Nullable Output<ProvisionedModelThroughputTimeoutsArgs> timeouts) {
            $.timeouts = timeouts;
            return this;
        }

        public Builder timeouts(ProvisionedModelThroughputTimeoutsArgs timeouts) {
            return timeouts(Output.of(timeouts));
        }

        public ProvisionedModelThroughputState build() {
            return $;
        }
    }

}
