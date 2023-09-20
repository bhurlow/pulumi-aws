// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkflowState extends com.pulumi.resources.ResourceArgs {

    public static final WorkflowState Empty = new WorkflowState();

    /**
     * Amazon Resource Name (ARN) of Glue Workflow
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of Glue Workflow
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
     * 
     */
    @Import(name="defaultRunProperties")
    private @Nullable Output<Map<String,String>> defaultRunProperties;

    /**
     * @return A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
     * 
     */
    public Optional<Output<Map<String,String>>> defaultRunProperties() {
        return Optional.ofNullable(this.defaultRunProperties);
    }

    /**
     * Description of the workflow.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the workflow.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
     * 
     */
    @Import(name="maxConcurrentRuns")
    private @Nullable Output<Integer> maxConcurrentRuns;

    /**
     * @return Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
     * 
     */
    public Optional<Output<Integer>> maxConcurrentRuns() {
        return Optional.ofNullable(this.maxConcurrentRuns);
    }

    /**
     * The name you assign to this workflow.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name you assign to this workflow.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private WorkflowState() {}

    private WorkflowState(WorkflowState $) {
        this.arn = $.arn;
        this.defaultRunProperties = $.defaultRunProperties;
        this.description = $.description;
        this.maxConcurrentRuns = $.maxConcurrentRuns;
        this.name = $.name;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkflowState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkflowState $;

        public Builder() {
            $ = new WorkflowState();
        }

        public Builder(WorkflowState defaults) {
            $ = new WorkflowState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of Glue Workflow
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of Glue Workflow
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param defaultRunProperties A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
         * 
         * @return builder
         * 
         */
        public Builder defaultRunProperties(@Nullable Output<Map<String,String>> defaultRunProperties) {
            $.defaultRunProperties = defaultRunProperties;
            return this;
        }

        /**
         * @param defaultRunProperties A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
         * 
         * @return builder
         * 
         */
        public Builder defaultRunProperties(Map<String,String> defaultRunProperties) {
            return defaultRunProperties(Output.of(defaultRunProperties));
        }

        /**
         * @param description Description of the workflow.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the workflow.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param maxConcurrentRuns Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
         * 
         * @return builder
         * 
         */
        public Builder maxConcurrentRuns(@Nullable Output<Integer> maxConcurrentRuns) {
            $.maxConcurrentRuns = maxConcurrentRuns;
            return this;
        }

        /**
         * @param maxConcurrentRuns Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
         * 
         * @return builder
         * 
         */
        public Builder maxConcurrentRuns(Integer maxConcurrentRuns) {
            return maxConcurrentRuns(Output.of(maxConcurrentRuns));
        }

        /**
         * @param name The name you assign to this workflow.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name you assign to this workflow.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public WorkflowState build() {
            return $;
        }
    }

}
