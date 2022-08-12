// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.autoscaling.inputs;

import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatMetricArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs extends com.pulumi.resources.ResourceArgs {

    public static final PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs Empty = new PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs();

    /**
     * A structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
     * 
     */
    @Import(name="metric", required=true)
    private Output<PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatMetricArgs> metric;

    /**
     * @return A structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
     * 
     */
    public Output<PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatMetricArgs> metric() {
        return this.metric;
    }

    /**
     * The statistic of the metrics to return.
     * 
     */
    @Import(name="stat", required=true)
    private Output<String> stat;

    /**
     * @return The statistic of the metrics to return.
     * 
     */
    public Output<String> stat() {
        return this.stat;
    }

    /**
     * The unit of the metrics to return.
     * 
     */
    @Import(name="unit")
    private @Nullable Output<String> unit;

    /**
     * @return The unit of the metrics to return.
     * 
     */
    public Optional<Output<String>> unit() {
        return Optional.ofNullable(this.unit);
    }

    private PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs() {}

    private PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs $) {
        this.metric = $.metric;
        this.stat = $.stat;
        this.unit = $.unit;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs $;

        public Builder() {
            $ = new PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs();
        }

        public Builder(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs defaults) {
            $ = new PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param metric A structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
         * 
         * @return builder
         * 
         */
        public Builder metric(Output<PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatMetricArgs> metric) {
            $.metric = metric;
            return this;
        }

        /**
         * @param metric A structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
         * 
         * @return builder
         * 
         */
        public Builder metric(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatMetricArgs metric) {
            return metric(Output.of(metric));
        }

        /**
         * @param stat The statistic of the metrics to return.
         * 
         * @return builder
         * 
         */
        public Builder stat(Output<String> stat) {
            $.stat = stat;
            return this;
        }

        /**
         * @param stat The statistic of the metrics to return.
         * 
         * @return builder
         * 
         */
        public Builder stat(String stat) {
            return stat(Output.of(stat));
        }

        /**
         * @param unit The unit of the metrics to return.
         * 
         * @return builder
         * 
         */
        public Builder unit(@Nullable Output<String> unit) {
            $.unit = unit;
            return this;
        }

        /**
         * @param unit The unit of the metrics to return.
         * 
         * @return builder
         * 
         */
        public Builder unit(String unit) {
            return unit(Output.of(unit));
        }

        public PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatArgs build() {
            $.metric = Objects.requireNonNull($.metric, "expected parameter 'metric' to be non-null");
            $.stat = Objects.requireNonNull($.stat, "expected parameter 'stat' to be non-null");
            return $;
        }
    }

}