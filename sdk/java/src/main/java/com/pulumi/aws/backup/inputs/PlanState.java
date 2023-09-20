// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.backup.inputs;

import com.pulumi.aws.backup.inputs.PlanAdvancedBackupSettingArgs;
import com.pulumi.aws.backup.inputs.PlanRuleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PlanState extends com.pulumi.resources.ResourceArgs {

    public static final PlanState Empty = new PlanState();

    /**
     * An object that specifies backup options for each resource type.
     * 
     */
    @Import(name="advancedBackupSettings")
    private @Nullable Output<List<PlanAdvancedBackupSettingArgs>> advancedBackupSettings;

    /**
     * @return An object that specifies backup options for each resource type.
     * 
     */
    public Optional<Output<List<PlanAdvancedBackupSettingArgs>>> advancedBackupSettings() {
        return Optional.ofNullable(this.advancedBackupSettings);
    }

    /**
     * The ARN of the backup plan.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the backup plan.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The display name of a backup plan.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The display name of a backup plan.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A rule object that specifies a scheduled task that is used to back up a selection of resources.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<PlanRuleArgs>> rules;

    /**
     * @return A rule object that specifies a scheduled task that is used to back up a selection of resources.
     * 
     */
    public Optional<Output<List<PlanRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * Metadata that you can assign to help organize the plans you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Metadata that you can assign to help organize the plans you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    /**
     * Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private PlanState() {}

    private PlanState(PlanState $) {
        this.advancedBackupSettings = $.advancedBackupSettings;
        this.arn = $.arn;
        this.name = $.name;
        this.rules = $.rules;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PlanState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PlanState $;

        public Builder() {
            $ = new PlanState();
        }

        public Builder(PlanState defaults) {
            $ = new PlanState(Objects.requireNonNull(defaults));
        }

        /**
         * @param advancedBackupSettings An object that specifies backup options for each resource type.
         * 
         * @return builder
         * 
         */
        public Builder advancedBackupSettings(@Nullable Output<List<PlanAdvancedBackupSettingArgs>> advancedBackupSettings) {
            $.advancedBackupSettings = advancedBackupSettings;
            return this;
        }

        /**
         * @param advancedBackupSettings An object that specifies backup options for each resource type.
         * 
         * @return builder
         * 
         */
        public Builder advancedBackupSettings(List<PlanAdvancedBackupSettingArgs> advancedBackupSettings) {
            return advancedBackupSettings(Output.of(advancedBackupSettings));
        }

        /**
         * @param advancedBackupSettings An object that specifies backup options for each resource type.
         * 
         * @return builder
         * 
         */
        public Builder advancedBackupSettings(PlanAdvancedBackupSettingArgs... advancedBackupSettings) {
            return advancedBackupSettings(List.of(advancedBackupSettings));
        }

        /**
         * @param arn The ARN of the backup plan.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the backup plan.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param name The display name of a backup plan.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The display name of a backup plan.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param rules A rule object that specifies a scheduled task that is used to back up a selection of resources.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<PlanRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules A rule object that specifies a scheduled task that is used to back up a selection of resources.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<PlanRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules A rule object that specifies a scheduled task that is used to back up a selection of resources.
         * 
         * @return builder
         * 
         */
        public Builder rules(PlanRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param tags Metadata that you can assign to help organize the plans you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Metadata that you can assign to help organize the plans you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /**
         * @param version Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public PlanState build() {
            return $;
        }
    }

}
