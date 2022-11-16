// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fis.inputs;

import com.pulumi.aws.fis.inputs.ExperimentTemplateActionArgs;
import com.pulumi.aws.fis.inputs.ExperimentTemplateStopConditionArgs;
import com.pulumi.aws.fis.inputs.ExperimentTemplateTargetArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExperimentTemplateState extends com.pulumi.resources.ResourceArgs {

    public static final ExperimentTemplateState Empty = new ExperimentTemplateState();

    /**
     * Action to be performed during an experiment. See below.
     * 
     */
    @Import(name="actions")
    private @Nullable Output<List<ExperimentTemplateActionArgs>> actions;

    /**
     * @return Action to be performed during an experiment. See below.
     * 
     */
    public Optional<Output<List<ExperimentTemplateActionArgs>>> actions() {
        return Optional.ofNullable(this.actions);
    }

    /**
     * Description of the action.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the action.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * When an ongoing experiment should be stopped. See below.
     * 
     */
    @Import(name="stopConditions")
    private @Nullable Output<List<ExperimentTemplateStopConditionArgs>> stopConditions;

    /**
     * @return When an ongoing experiment should be stopped. See below.
     * 
     */
    public Optional<Output<List<ExperimentTemplateStopConditionArgs>>> stopConditions() {
        return Optional.ofNullable(this.stopConditions);
    }

    /**
     * Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Action&#39;s target, if applicable. See below.
     * 
     */
    @Import(name="targets")
    private @Nullable Output<List<ExperimentTemplateTargetArgs>> targets;

    /**
     * @return Action&#39;s target, if applicable. See below.
     * 
     */
    public Optional<Output<List<ExperimentTemplateTargetArgs>>> targets() {
        return Optional.ofNullable(this.targets);
    }

    private ExperimentTemplateState() {}

    private ExperimentTemplateState(ExperimentTemplateState $) {
        this.actions = $.actions;
        this.description = $.description;
        this.roleArn = $.roleArn;
        this.stopConditions = $.stopConditions;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.targets = $.targets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExperimentTemplateState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExperimentTemplateState $;

        public Builder() {
            $ = new ExperimentTemplateState();
        }

        public Builder(ExperimentTemplateState defaults) {
            $ = new ExperimentTemplateState(Objects.requireNonNull(defaults));
        }

        /**
         * @param actions Action to be performed during an experiment. See below.
         * 
         * @return builder
         * 
         */
        public Builder actions(@Nullable Output<List<ExperimentTemplateActionArgs>> actions) {
            $.actions = actions;
            return this;
        }

        /**
         * @param actions Action to be performed during an experiment. See below.
         * 
         * @return builder
         * 
         */
        public Builder actions(List<ExperimentTemplateActionArgs> actions) {
            return actions(Output.of(actions));
        }

        /**
         * @param actions Action to be performed during an experiment. See below.
         * 
         * @return builder
         * 
         */
        public Builder actions(ExperimentTemplateActionArgs... actions) {
            return actions(List.of(actions));
        }

        /**
         * @param description Description of the action.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the action.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param roleArn ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param stopConditions When an ongoing experiment should be stopped. See below.
         * 
         * @return builder
         * 
         */
        public Builder stopConditions(@Nullable Output<List<ExperimentTemplateStopConditionArgs>> stopConditions) {
            $.stopConditions = stopConditions;
            return this;
        }

        /**
         * @param stopConditions When an ongoing experiment should be stopped. See below.
         * 
         * @return builder
         * 
         */
        public Builder stopConditions(List<ExperimentTemplateStopConditionArgs> stopConditions) {
            return stopConditions(Output.of(stopConditions));
        }

        /**
         * @param stopConditions When an ongoing experiment should be stopped. See below.
         * 
         * @return builder
         * 
         */
        public Builder stopConditions(ExperimentTemplateStopConditionArgs... stopConditions) {
            return stopConditions(List.of(stopConditions));
        }

        /**
         * @param tags Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param targets Action&#39;s target, if applicable. See below.
         * 
         * @return builder
         * 
         */
        public Builder targets(@Nullable Output<List<ExperimentTemplateTargetArgs>> targets) {
            $.targets = targets;
            return this;
        }

        /**
         * @param targets Action&#39;s target, if applicable. See below.
         * 
         * @return builder
         * 
         */
        public Builder targets(List<ExperimentTemplateTargetArgs> targets) {
            return targets(Output.of(targets));
        }

        /**
         * @param targets Action&#39;s target, if applicable. See below.
         * 
         * @return builder
         * 
         */
        public Builder targets(ExperimentTemplateTargetArgs... targets) {
            return targets(List.of(targets));
        }

        public ExperimentTemplateState build() {
            return $;
        }
    }

}
