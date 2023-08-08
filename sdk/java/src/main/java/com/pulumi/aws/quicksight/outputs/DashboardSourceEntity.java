// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight.outputs;

import com.pulumi.aws.quicksight.outputs.DashboardSourceEntitySourceTemplate;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardSourceEntity {
    /**
     * @return The source template. See source_template.
     * 
     */
    private @Nullable DashboardSourceEntitySourceTemplate sourceTemplate;

    private DashboardSourceEntity() {}
    /**
     * @return The source template. See source_template.
     * 
     */
    public Optional<DashboardSourceEntitySourceTemplate> sourceTemplate() {
        return Optional.ofNullable(this.sourceTemplate);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardSourceEntity defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable DashboardSourceEntitySourceTemplate sourceTemplate;
        public Builder() {}
        public Builder(DashboardSourceEntity defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.sourceTemplate = defaults.sourceTemplate;
        }

        @CustomType.Setter
        public Builder sourceTemplate(@Nullable DashboardSourceEntitySourceTemplate sourceTemplate) {
            this.sourceTemplate = sourceTemplate;
            return this;
        }
        public DashboardSourceEntity build() {
            final var o = new DashboardSourceEntity();
            o.sourceTemplate = sourceTemplate;
            return o;
        }
    }
}
