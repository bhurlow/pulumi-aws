// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.autoscaling.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs extends com.pulumi.resources.ResourceArgs {

    public static final PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs Empty = new PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs();

    /**
     * Describes a scaling metric for a predictive scaling policy. Valid values are `ASGAverageCPUUtilization`, `ASGAverageNetworkIn`, `ASGAverageNetworkOut`, or `ALBRequestCountPerTarget`.
     * 
     */
    @Import(name="predefinedMetricType", required=true)
    private Output<String> predefinedMetricType;

    /**
     * @return Describes a scaling metric for a predictive scaling policy. Valid values are `ASGAverageCPUUtilization`, `ASGAverageNetworkIn`, `ASGAverageNetworkOut`, or `ALBRequestCountPerTarget`.
     * 
     */
    public Output<String> predefinedMetricType() {
        return this.predefinedMetricType;
    }

    /**
     * Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group. You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). Refer to [PredefinedMetricSpecification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredefinedMetricSpecification.html) for more information.
     * 
     */
    @Import(name="resourceLabel")
    private @Nullable Output<String> resourceLabel;

    /**
     * @return Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group. You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). Refer to [PredefinedMetricSpecification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredefinedMetricSpecification.html) for more information.
     * 
     */
    public Optional<Output<String>> resourceLabel() {
        return Optional.ofNullable(this.resourceLabel);
    }

    private PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs() {}

    private PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs(PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs $) {
        this.predefinedMetricType = $.predefinedMetricType;
        this.resourceLabel = $.resourceLabel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs $;

        public Builder() {
            $ = new PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs();
        }

        public Builder(PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs defaults) {
            $ = new PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param predefinedMetricType Describes a scaling metric for a predictive scaling policy. Valid values are `ASGAverageCPUUtilization`, `ASGAverageNetworkIn`, `ASGAverageNetworkOut`, or `ALBRequestCountPerTarget`.
         * 
         * @return builder
         * 
         */
        public Builder predefinedMetricType(Output<String> predefinedMetricType) {
            $.predefinedMetricType = predefinedMetricType;
            return this;
        }

        /**
         * @param predefinedMetricType Describes a scaling metric for a predictive scaling policy. Valid values are `ASGAverageCPUUtilization`, `ASGAverageNetworkIn`, `ASGAverageNetworkOut`, or `ALBRequestCountPerTarget`.
         * 
         * @return builder
         * 
         */
        public Builder predefinedMetricType(String predefinedMetricType) {
            return predefinedMetricType(Output.of(predefinedMetricType));
        }

        /**
         * @param resourceLabel Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group. You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). Refer to [PredefinedMetricSpecification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredefinedMetricSpecification.html) for more information.
         * 
         * @return builder
         * 
         */
        public Builder resourceLabel(@Nullable Output<String> resourceLabel) {
            $.resourceLabel = resourceLabel;
            return this;
        }

        /**
         * @param resourceLabel Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group. You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). Refer to [PredefinedMetricSpecification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredefinedMetricSpecification.html) for more information.
         * 
         * @return builder
         * 
         */
        public Builder resourceLabel(String resourceLabel) {
            return resourceLabel(Output.of(resourceLabel));
        }

        public PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedScalingMetricSpecificationArgs build() {
            $.predefinedMetricType = Objects.requireNonNull($.predefinedMetricType, "expected parameter 'predefinedMetricType' to be non-null");
            return $;
        }
    }

}
