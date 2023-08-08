// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride {
    /**
     * @return The Elastic Inference accelerator device name to override for the task. This parameter must match a deviceName specified in the task definition.
     * 
     */
    private @Nullable String deviceName;
    /**
     * @return The Elastic Inference accelerator type to use.
     * 
     */
    private @Nullable String deviceType;

    private PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride() {}
    /**
     * @return The Elastic Inference accelerator device name to override for the task. This parameter must match a deviceName specified in the task definition.
     * 
     */
    public Optional<String> deviceName() {
        return Optional.ofNullable(this.deviceName);
    }
    /**
     * @return The Elastic Inference accelerator type to use.
     * 
     */
    public Optional<String> deviceType() {
        return Optional.ofNullable(this.deviceType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String deviceName;
        private @Nullable String deviceType;
        public Builder() {}
        public Builder(PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.deviceName = defaults.deviceName;
    	      this.deviceType = defaults.deviceType;
        }

        @CustomType.Setter
        public Builder deviceName(@Nullable String deviceName) {
            this.deviceName = deviceName;
            return this;
        }
        @CustomType.Setter
        public Builder deviceType(@Nullable String deviceType) {
            this.deviceType = deviceType;
            return this;
        }
        public PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride build() {
            final var o = new PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride();
            o.deviceName = deviceName;
            o.deviceType = deviceType;
            return o;
        }
    }
}
