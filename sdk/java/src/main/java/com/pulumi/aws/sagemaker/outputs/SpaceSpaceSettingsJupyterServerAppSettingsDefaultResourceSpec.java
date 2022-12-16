// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec {
    /**
     * @return The instance type.
     * 
     */
    private @Nullable String instanceType;
    /**
     * @return The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
     * 
     */
    private @Nullable String lifecycleConfigArn;
    /**
     * @return The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
     * 
     */
    private @Nullable String sagemakerImageArn;
    /**
     * @return The ARN of the image version created on the instance.
     * 
     */
    private @Nullable String sagemakerImageVersionArn;

    private SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec() {}
    /**
     * @return The instance type.
     * 
     */
    public Optional<String> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }
    /**
     * @return The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
     * 
     */
    public Optional<String> lifecycleConfigArn() {
        return Optional.ofNullable(this.lifecycleConfigArn);
    }
    /**
     * @return The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
     * 
     */
    public Optional<String> sagemakerImageArn() {
        return Optional.ofNullable(this.sagemakerImageArn);
    }
    /**
     * @return The ARN of the image version created on the instance.
     * 
     */
    public Optional<String> sagemakerImageVersionArn() {
        return Optional.ofNullable(this.sagemakerImageVersionArn);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String instanceType;
        private @Nullable String lifecycleConfigArn;
        private @Nullable String sagemakerImageArn;
        private @Nullable String sagemakerImageVersionArn;
        public Builder() {}
        public Builder(SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.instanceType = defaults.instanceType;
    	      this.lifecycleConfigArn = defaults.lifecycleConfigArn;
    	      this.sagemakerImageArn = defaults.sagemakerImageArn;
    	      this.sagemakerImageVersionArn = defaults.sagemakerImageVersionArn;
        }

        @CustomType.Setter
        public Builder instanceType(@Nullable String instanceType) {
            this.instanceType = instanceType;
            return this;
        }
        @CustomType.Setter
        public Builder lifecycleConfigArn(@Nullable String lifecycleConfigArn) {
            this.lifecycleConfigArn = lifecycleConfigArn;
            return this;
        }
        @CustomType.Setter
        public Builder sagemakerImageArn(@Nullable String sagemakerImageArn) {
            this.sagemakerImageArn = sagemakerImageArn;
            return this;
        }
        @CustomType.Setter
        public Builder sagemakerImageVersionArn(@Nullable String sagemakerImageVersionArn) {
            this.sagemakerImageVersionArn = sagemakerImageVersionArn;
            return this;
        }
        public SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec build() {
            final var o = new SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec();
            o.instanceType = instanceType;
            o.lifecycleConfigArn = lifecycleConfigArn;
            o.sagemakerImageArn = sagemakerImageArn;
            o.sagemakerImageVersionArn = sagemakerImageVersionArn;
            return o;
        }
    }
}
