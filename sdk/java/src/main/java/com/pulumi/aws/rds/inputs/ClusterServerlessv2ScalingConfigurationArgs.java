// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.util.Objects;


public final class ClusterServerlessv2ScalingConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterServerlessv2ScalingConfigurationArgs Empty = new ClusterServerlessv2ScalingConfigurationArgs();

    /**
     * Maximum capacity for an Aurora DB cluster in `serverless` DB engine mode. The maximum capacity must be greater than or equal to the minimum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `16`.
     * 
     */
    @Import(name="maxCapacity", required=true)
    private Output<Double> maxCapacity;

    /**
     * @return Maximum capacity for an Aurora DB cluster in `serverless` DB engine mode. The maximum capacity must be greater than or equal to the minimum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `16`.
     * 
     */
    public Output<Double> maxCapacity() {
        return this.maxCapacity;
    }

    /**
     * Minimum capacity for an Aurora DB cluster in `serverless` DB engine mode. The minimum capacity must be lesser than or equal to the maximum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `1`.
     * 
     */
    @Import(name="minCapacity", required=true)
    private Output<Double> minCapacity;

    /**
     * @return Minimum capacity for an Aurora DB cluster in `serverless` DB engine mode. The minimum capacity must be lesser than or equal to the maximum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `1`.
     * 
     */
    public Output<Double> minCapacity() {
        return this.minCapacity;
    }

    private ClusterServerlessv2ScalingConfigurationArgs() {}

    private ClusterServerlessv2ScalingConfigurationArgs(ClusterServerlessv2ScalingConfigurationArgs $) {
        this.maxCapacity = $.maxCapacity;
        this.minCapacity = $.minCapacity;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterServerlessv2ScalingConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterServerlessv2ScalingConfigurationArgs $;

        public Builder() {
            $ = new ClusterServerlessv2ScalingConfigurationArgs();
        }

        public Builder(ClusterServerlessv2ScalingConfigurationArgs defaults) {
            $ = new ClusterServerlessv2ScalingConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param maxCapacity Maximum capacity for an Aurora DB cluster in `serverless` DB engine mode. The maximum capacity must be greater than or equal to the minimum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `16`.
         * 
         * @return builder
         * 
         */
        public Builder maxCapacity(Output<Double> maxCapacity) {
            $.maxCapacity = maxCapacity;
            return this;
        }

        /**
         * @param maxCapacity Maximum capacity for an Aurora DB cluster in `serverless` DB engine mode. The maximum capacity must be greater than or equal to the minimum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `16`.
         * 
         * @return builder
         * 
         */
        public Builder maxCapacity(Double maxCapacity) {
            return maxCapacity(Output.of(maxCapacity));
        }

        /**
         * @param minCapacity Minimum capacity for an Aurora DB cluster in `serverless` DB engine mode. The minimum capacity must be lesser than or equal to the maximum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder minCapacity(Output<Double> minCapacity) {
            $.minCapacity = minCapacity;
            return this;
        }

        /**
         * @param minCapacity Minimum capacity for an Aurora DB cluster in `serverless` DB engine mode. The minimum capacity must be lesser than or equal to the maximum capacity. Valid Aurora MySQL capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`. Valid Aurora PostgreSQL capacity values are (`2`, `4`, `8`, `16`, `32`, `64`, `192`, and `384`). Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder minCapacity(Double minCapacity) {
            return minCapacity(Output.of(minCapacity));
        }

        public ClusterServerlessv2ScalingConfigurationArgs build() {
            $.maxCapacity = Objects.requireNonNull($.maxCapacity, "expected parameter 'maxCapacity' to be non-null");
            $.minCapacity = Objects.requireNonNull($.minCapacity, "expected parameter 'minCapacity' to be non-null");
            return $;
        }
    }

}
