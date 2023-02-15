// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.evidently.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class LaunchScheduledSplitsConfigStepSegmentOverride {
    /**
     * @return Specifies a number indicating the order to use to evaluate segment overrides, if there are more than one. Segment overrides with lower numbers are evaluated first.
     * 
     */
    private Integer evaluationOrder;
    /**
     * @return The name or ARN of the segment to use.
     * 
     */
    private String segment;
    /**
     * @return The traffic allocation percentages among the feature variations to assign to this segment. This is a set of key-value pairs. The keys are variation names. The values represent the amount of traffic to allocate to that variation for this segment. This is expressed in thousandths of a percent, so a weight of 50000 represents 50% of traffic.
     * 
     */
    private Map<String,Integer> weights;

    private LaunchScheduledSplitsConfigStepSegmentOverride() {}
    /**
     * @return Specifies a number indicating the order to use to evaluate segment overrides, if there are more than one. Segment overrides with lower numbers are evaluated first.
     * 
     */
    public Integer evaluationOrder() {
        return this.evaluationOrder;
    }
    /**
     * @return The name or ARN of the segment to use.
     * 
     */
    public String segment() {
        return this.segment;
    }
    /**
     * @return The traffic allocation percentages among the feature variations to assign to this segment. This is a set of key-value pairs. The keys are variation names. The values represent the amount of traffic to allocate to that variation for this segment. This is expressed in thousandths of a percent, so a weight of 50000 represents 50% of traffic.
     * 
     */
    public Map<String,Integer> weights() {
        return this.weights;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LaunchScheduledSplitsConfigStepSegmentOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer evaluationOrder;
        private String segment;
        private Map<String,Integer> weights;
        public Builder() {}
        public Builder(LaunchScheduledSplitsConfigStepSegmentOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.evaluationOrder = defaults.evaluationOrder;
    	      this.segment = defaults.segment;
    	      this.weights = defaults.weights;
        }

        @CustomType.Setter
        public Builder evaluationOrder(Integer evaluationOrder) {
            this.evaluationOrder = Objects.requireNonNull(evaluationOrder);
            return this;
        }
        @CustomType.Setter
        public Builder segment(String segment) {
            this.segment = Objects.requireNonNull(segment);
            return this;
        }
        @CustomType.Setter
        public Builder weights(Map<String,Integer> weights) {
            this.weights = Objects.requireNonNull(weights);
            return this;
        }
        public LaunchScheduledSplitsConfigStepSegmentOverride build() {
            final var o = new LaunchScheduledSplitsConfigStepSegmentOverride();
            o.evaluationOrder = evaluationOrder;
            o.segment = segment;
            o.weights = weights;
            return o;
        }
    }
}
