// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.redshift.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ScheduledActionTargetActionResumeClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ScheduledActionTargetActionResumeClusterArgs Empty = new ScheduledActionTargetActionResumeClusterArgs();

    /**
     * The identifier of the cluster to be resumed.
     * 
     */
    @Import(name="clusterIdentifier", required=true)
    private Output<String> clusterIdentifier;

    /**
     * @return The identifier of the cluster to be resumed.
     * 
     */
    public Output<String> clusterIdentifier() {
        return this.clusterIdentifier;
    }

    private ScheduledActionTargetActionResumeClusterArgs() {}

    private ScheduledActionTargetActionResumeClusterArgs(ScheduledActionTargetActionResumeClusterArgs $) {
        this.clusterIdentifier = $.clusterIdentifier;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ScheduledActionTargetActionResumeClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ScheduledActionTargetActionResumeClusterArgs $;

        public Builder() {
            $ = new ScheduledActionTargetActionResumeClusterArgs();
        }

        public Builder(ScheduledActionTargetActionResumeClusterArgs defaults) {
            $ = new ScheduledActionTargetActionResumeClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterIdentifier The identifier of the cluster to be resumed.
         * 
         * @return builder
         * 
         */
        public Builder clusterIdentifier(Output<String> clusterIdentifier) {
            $.clusterIdentifier = clusterIdentifier;
            return this;
        }

        /**
         * @param clusterIdentifier The identifier of the cluster to be resumed.
         * 
         * @return builder
         * 
         */
        public Builder clusterIdentifier(String clusterIdentifier) {
            return clusterIdentifier(Output.of(clusterIdentifier));
        }

        public ScheduledActionTargetActionResumeClusterArgs build() {
            $.clusterIdentifier = Objects.requireNonNull($.clusterIdentifier, "expected parameter 'clusterIdentifier' to be non-null");
            return $;
        }
    }

}