// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.servicequotas.inputs;

import com.pulumi.aws.servicequotas.inputs.ServiceQuotaUsageMetricMetricDimensionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceQuotaUsageMetricArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceQuotaUsageMetricArgs Empty = new ServiceQuotaUsageMetricArgs();

    /**
     * The metric dimensions.
     * 
     */
    @Import(name="metricDimensions")
    private @Nullable Output<List<ServiceQuotaUsageMetricMetricDimensionArgs>> metricDimensions;

    /**
     * @return The metric dimensions.
     * 
     */
    public Optional<Output<List<ServiceQuotaUsageMetricMetricDimensionArgs>>> metricDimensions() {
        return Optional.ofNullable(this.metricDimensions);
    }

    /**
     * The name of the metric.
     * 
     */
    @Import(name="metricName")
    private @Nullable Output<String> metricName;

    /**
     * @return The name of the metric.
     * 
     */
    public Optional<Output<String>> metricName() {
        return Optional.ofNullable(this.metricName);
    }

    /**
     * The namespace of the metric.
     * 
     */
    @Import(name="metricNamespace")
    private @Nullable Output<String> metricNamespace;

    /**
     * @return The namespace of the metric.
     * 
     */
    public Optional<Output<String>> metricNamespace() {
        return Optional.ofNullable(this.metricNamespace);
    }

    /**
     * The metric statistic that AWS recommend you use when determining quota usage.
     * 
     */
    @Import(name="metricStatisticRecommendation")
    private @Nullable Output<String> metricStatisticRecommendation;

    /**
     * @return The metric statistic that AWS recommend you use when determining quota usage.
     * 
     */
    public Optional<Output<String>> metricStatisticRecommendation() {
        return Optional.ofNullable(this.metricStatisticRecommendation);
    }

    private ServiceQuotaUsageMetricArgs() {}

    private ServiceQuotaUsageMetricArgs(ServiceQuotaUsageMetricArgs $) {
        this.metricDimensions = $.metricDimensions;
        this.metricName = $.metricName;
        this.metricNamespace = $.metricNamespace;
        this.metricStatisticRecommendation = $.metricStatisticRecommendation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceQuotaUsageMetricArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceQuotaUsageMetricArgs $;

        public Builder() {
            $ = new ServiceQuotaUsageMetricArgs();
        }

        public Builder(ServiceQuotaUsageMetricArgs defaults) {
            $ = new ServiceQuotaUsageMetricArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param metricDimensions The metric dimensions.
         * 
         * @return builder
         * 
         */
        public Builder metricDimensions(@Nullable Output<List<ServiceQuotaUsageMetricMetricDimensionArgs>> metricDimensions) {
            $.metricDimensions = metricDimensions;
            return this;
        }

        /**
         * @param metricDimensions The metric dimensions.
         * 
         * @return builder
         * 
         */
        public Builder metricDimensions(List<ServiceQuotaUsageMetricMetricDimensionArgs> metricDimensions) {
            return metricDimensions(Output.of(metricDimensions));
        }

        /**
         * @param metricDimensions The metric dimensions.
         * 
         * @return builder
         * 
         */
        public Builder metricDimensions(ServiceQuotaUsageMetricMetricDimensionArgs... metricDimensions) {
            return metricDimensions(List.of(metricDimensions));
        }

        /**
         * @param metricName The name of the metric.
         * 
         * @return builder
         * 
         */
        public Builder metricName(@Nullable Output<String> metricName) {
            $.metricName = metricName;
            return this;
        }

        /**
         * @param metricName The name of the metric.
         * 
         * @return builder
         * 
         */
        public Builder metricName(String metricName) {
            return metricName(Output.of(metricName));
        }

        /**
         * @param metricNamespace The namespace of the metric.
         * 
         * @return builder
         * 
         */
        public Builder metricNamespace(@Nullable Output<String> metricNamespace) {
            $.metricNamespace = metricNamespace;
            return this;
        }

        /**
         * @param metricNamespace The namespace of the metric.
         * 
         * @return builder
         * 
         */
        public Builder metricNamespace(String metricNamespace) {
            return metricNamespace(Output.of(metricNamespace));
        }

        /**
         * @param metricStatisticRecommendation The metric statistic that AWS recommend you use when determining quota usage.
         * 
         * @return builder
         * 
         */
        public Builder metricStatisticRecommendation(@Nullable Output<String> metricStatisticRecommendation) {
            $.metricStatisticRecommendation = metricStatisticRecommendation;
            return this;
        }

        /**
         * @param metricStatisticRecommendation The metric statistic that AWS recommend you use when determining quota usage.
         * 
         * @return builder
         * 
         */
        public Builder metricStatisticRecommendation(String metricStatisticRecommendation) {
            return metricStatisticRecommendation(Output.of(metricStatisticRecommendation));
        }

        public ServiceQuotaUsageMetricArgs build() {
            return $;
        }
    }

}
