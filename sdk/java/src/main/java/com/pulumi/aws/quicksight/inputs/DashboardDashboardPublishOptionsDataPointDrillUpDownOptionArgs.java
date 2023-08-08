// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs Empty = new DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs();

    /**
     * Availability status. Possibles values: ENABLED, DISABLED.
     * 
     */
    @Import(name="availabilityStatus")
    private @Nullable Output<String> availabilityStatus;

    /**
     * @return Availability status. Possibles values: ENABLED, DISABLED.
     * 
     */
    public Optional<Output<String>> availabilityStatus() {
        return Optional.ofNullable(this.availabilityStatus);
    }

    private DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs() {}

    private DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs(DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs $) {
        this.availabilityStatus = $.availabilityStatus;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs $;

        public Builder() {
            $ = new DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs();
        }

        public Builder(DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs defaults) {
            $ = new DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param availabilityStatus Availability status. Possibles values: ENABLED, DISABLED.
         * 
         * @return builder
         * 
         */
        public Builder availabilityStatus(@Nullable Output<String> availabilityStatus) {
            $.availabilityStatus = availabilityStatus;
            return this;
        }

        /**
         * @param availabilityStatus Availability status. Possibles values: ENABLED, DISABLED.
         * 
         * @return builder
         * 
         */
        public Builder availabilityStatus(String availabilityStatus) {
            return availabilityStatus(Output.of(availabilityStatus));
        }

        public DashboardDashboardPublishOptionsDataPointDrillUpDownOptionArgs build() {
            return $;
        }
    }

}
