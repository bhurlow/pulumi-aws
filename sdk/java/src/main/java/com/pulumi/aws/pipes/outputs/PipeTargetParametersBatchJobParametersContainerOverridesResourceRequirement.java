// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirement {
    /**
     * @return The type of resource to assign to a container. The supported resources include GPU, MEMORY, and VCPU.
     * 
     */
    private String type;
    /**
     * @return The value of the key-value pair. For environment variables, this is the value of the environment variable.
     * 
     */
    private String value;

    private PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirement() {}
    /**
     * @return The type of resource to assign to a container. The supported resources include GPU, MEMORY, and VCPU.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return The value of the key-value pair. For environment variables, this is the value of the environment variable.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirement defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String type;
        private String value;
        public Builder() {}
        public Builder(PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirement defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.type = defaults.type;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirement build() {
            final var o = new PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirement();
            o.type = type;
            o.value = value;
            return o;
        }
    }
}
