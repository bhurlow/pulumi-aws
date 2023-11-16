// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ssmcontacts.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PlanStageTargetContactTargetInfo {
    /**
     * @return The Amazon Resource Name (ARN) of the contact.
     * 
     */
    private @Nullable String contactId;
    /**
     * @return A Boolean value determining if the contact&#39;s acknowledgement stops the progress of stages in the plan.
     * 
     */
    private Boolean isEssential;

    private PlanStageTargetContactTargetInfo() {}
    /**
     * @return The Amazon Resource Name (ARN) of the contact.
     * 
     */
    public Optional<String> contactId() {
        return Optional.ofNullable(this.contactId);
    }
    /**
     * @return A Boolean value determining if the contact&#39;s acknowledgement stops the progress of stages in the plan.
     * 
     */
    public Boolean isEssential() {
        return this.isEssential;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PlanStageTargetContactTargetInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String contactId;
        private Boolean isEssential;
        public Builder() {}
        public Builder(PlanStageTargetContactTargetInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contactId = defaults.contactId;
    	      this.isEssential = defaults.isEssential;
        }

        @CustomType.Setter
        public Builder contactId(@Nullable String contactId) {
            this.contactId = contactId;
            return this;
        }
        @CustomType.Setter
        public Builder isEssential(Boolean isEssential) {
            this.isEssential = Objects.requireNonNull(isEssential);
            return this;
        }
        public PlanStageTargetContactTargetInfo build() {
            final var o = new PlanStageTargetContactTargetInfo();
            o.contactId = contactId;
            o.isEssential = isEssential;
            return o;
        }
    }
}
