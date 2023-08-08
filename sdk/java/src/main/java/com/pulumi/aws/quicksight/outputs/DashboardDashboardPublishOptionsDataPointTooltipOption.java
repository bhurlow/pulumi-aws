// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardDashboardPublishOptionsDataPointTooltipOption {
    /**
     * @return Availability status. Possibles values: ENABLED, DISABLED.
     * 
     */
    private @Nullable String availabilityStatus;

    private DashboardDashboardPublishOptionsDataPointTooltipOption() {}
    /**
     * @return Availability status. Possibles values: ENABLED, DISABLED.
     * 
     */
    public Optional<String> availabilityStatus() {
        return Optional.ofNullable(this.availabilityStatus);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardDashboardPublishOptionsDataPointTooltipOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String availabilityStatus;
        public Builder() {}
        public Builder(DashboardDashboardPublishOptionsDataPointTooltipOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availabilityStatus = defaults.availabilityStatus;
        }

        @CustomType.Setter
        public Builder availabilityStatus(@Nullable String availabilityStatus) {
            this.availabilityStatus = availabilityStatus;
            return this;
        }
        public DashboardDashboardPublishOptionsDataPointTooltipOption build() {
            final var o = new DashboardDashboardPublishOptionsDataPointTooltipOption();
            o.availabilityStatus = availabilityStatus;
            return o;
        }
    }
}
