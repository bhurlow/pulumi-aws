// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sfn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AliasRoutingConfiguration {
    /**
     * @return A version of the state machine.
     * 
     */
    private String stateMachineVersionArn;
    /**
     * @return Percentage of traffic routed to the state machine version.
     * 
     * The following arguments are optional:
     * 
     */
    private Integer weight;

    private AliasRoutingConfiguration() {}
    /**
     * @return A version of the state machine.
     * 
     */
    public String stateMachineVersionArn() {
        return this.stateMachineVersionArn;
    }
    /**
     * @return Percentage of traffic routed to the state machine version.
     * 
     * The following arguments are optional:
     * 
     */
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AliasRoutingConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String stateMachineVersionArn;
        private Integer weight;
        public Builder() {}
        public Builder(AliasRoutingConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.stateMachineVersionArn = defaults.stateMachineVersionArn;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder stateMachineVersionArn(String stateMachineVersionArn) {
            this.stateMachineVersionArn = Objects.requireNonNull(stateMachineVersionArn);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public AliasRoutingConfiguration build() {
            final var o = new AliasRoutingConfiguration();
            o.stateMachineVersionArn = stateMachineVersionArn;
            o.weight = weight;
            return o;
        }
    }
}
