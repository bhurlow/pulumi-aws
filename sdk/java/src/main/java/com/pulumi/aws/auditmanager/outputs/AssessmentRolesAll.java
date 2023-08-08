// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.auditmanager.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AssessmentRolesAll {
    /**
     * @return Amazon Resource Name (ARN) of the IAM role.
     * 
     */
    private String roleArn;
    /**
     * @return Type of customer persona. For assessment creation, type must always be `PROCESS_OWNER`.
     * 
     */
    private String roleType;

    private AssessmentRolesAll() {}
    /**
     * @return Amazon Resource Name (ARN) of the IAM role.
     * 
     */
    public String roleArn() {
        return this.roleArn;
    }
    /**
     * @return Type of customer persona. For assessment creation, type must always be `PROCESS_OWNER`.
     * 
     */
    public String roleType() {
        return this.roleType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AssessmentRolesAll defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String roleArn;
        private String roleType;
        public Builder() {}
        public Builder(AssessmentRolesAll defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.roleArn = defaults.roleArn;
    	      this.roleType = defaults.roleType;
        }

        @CustomType.Setter
        public Builder roleArn(String roleArn) {
            this.roleArn = Objects.requireNonNull(roleArn);
            return this;
        }
        @CustomType.Setter
        public Builder roleType(String roleType) {
            this.roleType = Objects.requireNonNull(roleType);
            return this;
        }
        public AssessmentRolesAll build() {
            final var o = new AssessmentRolesAll();
            o.roleArn = roleArn;
            o.roleType = roleType;
            return o;
        }
    }
}
