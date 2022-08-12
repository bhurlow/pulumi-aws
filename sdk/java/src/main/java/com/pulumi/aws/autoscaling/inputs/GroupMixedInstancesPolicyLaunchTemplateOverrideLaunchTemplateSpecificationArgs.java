// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.autoscaling.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs Empty = new GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs();

    /**
     * The ID of the launch template. Conflicts with `launch_template_name`.
     * 
     */
    @Import(name="launchTemplateId")
    private @Nullable Output<String> launchTemplateId;

    /**
     * @return The ID of the launch template. Conflicts with `launch_template_name`.
     * 
     */
    public Optional<Output<String>> launchTemplateId() {
        return Optional.ofNullable(this.launchTemplateId);
    }

    /**
     * The name of the launch template. Conflicts with `launch_template_id`.
     * 
     */
    @Import(name="launchTemplateName")
    private @Nullable Output<String> launchTemplateName;

    /**
     * @return The name of the launch template. Conflicts with `launch_template_id`.
     * 
     */
    public Optional<Output<String>> launchTemplateName() {
        return Optional.ofNullable(this.launchTemplateName);
    }

    /**
     * Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs() {}

    private GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs(GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs $) {
        this.launchTemplateId = $.launchTemplateId;
        this.launchTemplateName = $.launchTemplateName;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs $;

        public Builder() {
            $ = new GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs();
        }

        public Builder(GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs defaults) {
            $ = new GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param launchTemplateId The ID of the launch template. Conflicts with `launch_template_name`.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateId(@Nullable Output<String> launchTemplateId) {
            $.launchTemplateId = launchTemplateId;
            return this;
        }

        /**
         * @param launchTemplateId The ID of the launch template. Conflicts with `launch_template_name`.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateId(String launchTemplateId) {
            return launchTemplateId(Output.of(launchTemplateId));
        }

        /**
         * @param launchTemplateName The name of the launch template. Conflicts with `launch_template_id`.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateName(@Nullable Output<String> launchTemplateName) {
            $.launchTemplateName = launchTemplateName;
            return this;
        }

        /**
         * @param launchTemplateName The name of the launch template. Conflicts with `launch_template_id`.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateName(String launchTemplateName) {
            return launchTemplateName(Output.of(launchTemplateName));
        }

        /**
         * @param version Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs build() {
            return $;
        }
    }

}