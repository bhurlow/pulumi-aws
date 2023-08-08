// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.auditmanager.inputs;

import com.pulumi.aws.auditmanager.inputs.ControlControlMappingSourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControlState extends com.pulumi.resources.ResourceArgs {

    public static final ControlState Empty = new ControlState();

    /**
     * Recommended actions to carry out if the control isn&#39;t fulfilled.
     * 
     */
    @Import(name="actionPlanInstructions")
    private @Nullable Output<String> actionPlanInstructions;

    /**
     * @return Recommended actions to carry out if the control isn&#39;t fulfilled.
     * 
     */
    public Optional<Output<String>> actionPlanInstructions() {
        return Optional.ofNullable(this.actionPlanInstructions);
    }

    /**
     * Title of the action plan for remediating the control.
     * 
     */
    @Import(name="actionPlanTitle")
    private @Nullable Output<String> actionPlanTitle;

    /**
     * @return Title of the action plan for remediating the control.
     * 
     */
    public Optional<Output<String>> actionPlanTitle() {
        return Optional.ofNullable(this.actionPlanTitle);
    }

    /**
     * Amazon Resource Name (ARN) of the control.
     * * `control_mapping_sources.*.source_id` - Unique identifier for the source.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the control.
     * * `control_mapping_sources.*.source_id` - Unique identifier for the source.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Data mapping sources. See `control_mapping_sources` below.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="controlMappingSources")
    private @Nullable Output<List<ControlControlMappingSourceArgs>> controlMappingSources;

    /**
     * @return Data mapping sources. See `control_mapping_sources` below.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<List<ControlControlMappingSourceArgs>>> controlMappingSources() {
        return Optional.ofNullable(this.controlMappingSources);
    }

    /**
     * Description of the control.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the control.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of the control.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the control.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Steps to follow to determine if the control is satisfied.
     * 
     */
    @Import(name="testingInformation")
    private @Nullable Output<String> testingInformation;

    /**
     * @return Steps to follow to determine if the control is satisfied.
     * 
     */
    public Optional<Output<String>> testingInformation() {
        return Optional.ofNullable(this.testingInformation);
    }

    /**
     * Type of control, such as a custom control or a standard control.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of control, such as a custom control or a standard control.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ControlState() {}

    private ControlState(ControlState $) {
        this.actionPlanInstructions = $.actionPlanInstructions;
        this.actionPlanTitle = $.actionPlanTitle;
        this.arn = $.arn;
        this.controlMappingSources = $.controlMappingSources;
        this.description = $.description;
        this.name = $.name;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.testingInformation = $.testingInformation;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControlState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControlState $;

        public Builder() {
            $ = new ControlState();
        }

        public Builder(ControlState defaults) {
            $ = new ControlState(Objects.requireNonNull(defaults));
        }

        /**
         * @param actionPlanInstructions Recommended actions to carry out if the control isn&#39;t fulfilled.
         * 
         * @return builder
         * 
         */
        public Builder actionPlanInstructions(@Nullable Output<String> actionPlanInstructions) {
            $.actionPlanInstructions = actionPlanInstructions;
            return this;
        }

        /**
         * @param actionPlanInstructions Recommended actions to carry out if the control isn&#39;t fulfilled.
         * 
         * @return builder
         * 
         */
        public Builder actionPlanInstructions(String actionPlanInstructions) {
            return actionPlanInstructions(Output.of(actionPlanInstructions));
        }

        /**
         * @param actionPlanTitle Title of the action plan for remediating the control.
         * 
         * @return builder
         * 
         */
        public Builder actionPlanTitle(@Nullable Output<String> actionPlanTitle) {
            $.actionPlanTitle = actionPlanTitle;
            return this;
        }

        /**
         * @param actionPlanTitle Title of the action plan for remediating the control.
         * 
         * @return builder
         * 
         */
        public Builder actionPlanTitle(String actionPlanTitle) {
            return actionPlanTitle(Output.of(actionPlanTitle));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the control.
         * * `control_mapping_sources.*.source_id` - Unique identifier for the source.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the control.
         * * `control_mapping_sources.*.source_id` - Unique identifier for the source.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param controlMappingSources Data mapping sources. See `control_mapping_sources` below.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder controlMappingSources(@Nullable Output<List<ControlControlMappingSourceArgs>> controlMappingSources) {
            $.controlMappingSources = controlMappingSources;
            return this;
        }

        /**
         * @param controlMappingSources Data mapping sources. See `control_mapping_sources` below.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder controlMappingSources(List<ControlControlMappingSourceArgs> controlMappingSources) {
            return controlMappingSources(Output.of(controlMappingSources));
        }

        /**
         * @param controlMappingSources Data mapping sources. See `control_mapping_sources` below.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder controlMappingSources(ControlControlMappingSourceArgs... controlMappingSources) {
            return controlMappingSources(List.of(controlMappingSources));
        }

        /**
         * @param description Description of the control.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the control.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name of the control.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the control.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the control. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param testingInformation Steps to follow to determine if the control is satisfied.
         * 
         * @return builder
         * 
         */
        public Builder testingInformation(@Nullable Output<String> testingInformation) {
            $.testingInformation = testingInformation;
            return this;
        }

        /**
         * @param testingInformation Steps to follow to determine if the control is satisfied.
         * 
         * @return builder
         * 
         */
        public Builder testingInformation(String testingInformation) {
            return testingInformation(Output.of(testingInformation));
        }

        /**
         * @param type Type of control, such as a custom control or a standard control.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of control, such as a custom control or a standard control.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ControlState build() {
            return $;
        }
    }

}
