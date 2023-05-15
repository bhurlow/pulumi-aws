// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.autoscaling.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification {
    /**
     * @return ID of the launch template.
     * 
     */
    private String launchTemplateId;
    /**
     * @return Name of the launch template.
     * 
     */
    private String launchTemplateName;
    /**
     * @return Template version.
     * 
     */
    private String version;

    private GetGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification() {}
    /**
     * @return ID of the launch template.
     * 
     */
    public String launchTemplateId() {
        return this.launchTemplateId;
    }
    /**
     * @return Name of the launch template.
     * 
     */
    public String launchTemplateName() {
        return this.launchTemplateName;
    }
    /**
     * @return Template version.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String launchTemplateId;
        private String launchTemplateName;
        private String version;
        public Builder() {}
        public Builder(GetGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.launchTemplateId = defaults.launchTemplateId;
    	      this.launchTemplateName = defaults.launchTemplateName;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder launchTemplateId(String launchTemplateId) {
            this.launchTemplateId = Objects.requireNonNull(launchTemplateId);
            return this;
        }
        @CustomType.Setter
        public Builder launchTemplateName(String launchTemplateName) {
            this.launchTemplateName = Objects.requireNonNull(launchTemplateName);
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification build() {
            final var o = new GetGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification();
            o.launchTemplateId = launchTemplateId;
            o.launchTemplateName = launchTemplateName;
            o.version = version;
            return o;
        }
    }
}
