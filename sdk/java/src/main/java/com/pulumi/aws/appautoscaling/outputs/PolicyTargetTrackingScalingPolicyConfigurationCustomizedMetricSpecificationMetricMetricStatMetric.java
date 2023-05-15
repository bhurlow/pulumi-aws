// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appautoscaling.outputs;

import com.pulumi.aws.appautoscaling.outputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimension;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetric {
    /**
     * @return Configuration block(s) with the dimensions of the metric if the metric was published with dimensions. Detailed below.
     * 
     */
    private @Nullable List<PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimension> dimensions;
    /**
     * @return Name of the metric.
     * 
     */
    private String metricName;
    /**
     * @return Namespace of the metric.
     * 
     */
    private String namespace;

    private PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetric() {}
    /**
     * @return Configuration block(s) with the dimensions of the metric if the metric was published with dimensions. Detailed below.
     * 
     */
    public List<PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimension> dimensions() {
        return this.dimensions == null ? List.of() : this.dimensions;
    }
    /**
     * @return Name of the metric.
     * 
     */
    public String metricName() {
        return this.metricName;
    }
    /**
     * @return Namespace of the metric.
     * 
     */
    public String namespace() {
        return this.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetric defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimension> dimensions;
        private String metricName;
        private String namespace;
        public Builder() {}
        public Builder(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetric defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dimensions = defaults.dimensions;
    	      this.metricName = defaults.metricName;
    	      this.namespace = defaults.namespace;
        }

        @CustomType.Setter
        public Builder dimensions(@Nullable List<PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimension> dimensions) {
            this.dimensions = dimensions;
            return this;
        }
        public Builder dimensions(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimension... dimensions) {
            return dimensions(List.of(dimensions));
        }
        @CustomType.Setter
        public Builder metricName(String metricName) {
            this.metricName = Objects.requireNonNull(metricName);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }
        public PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetric build() {
            final var o = new PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetric();
            o.dimensions = dimensions;
            o.metricName = metricName;
            o.namespace = namespace;
            return o;
        }
    }
}
