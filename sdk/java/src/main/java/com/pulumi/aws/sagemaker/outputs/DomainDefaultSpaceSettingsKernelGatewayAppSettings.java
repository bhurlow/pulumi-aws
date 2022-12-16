// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.aws.sagemaker.outputs.DomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImage;
import com.pulumi.aws.sagemaker.outputs.DomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DomainDefaultSpaceSettingsKernelGatewayAppSettings {
    /**
     * @return A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
     * 
     */
    private @Nullable List<DomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImage> customImages;
    /**
     * @return The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
     * 
     */
    private @Nullable DomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec defaultResourceSpec;
    /**
     * @return The Amazon Resource Name (ARN) of the Lifecycle Configurations.
     * 
     */
    private @Nullable List<String> lifecycleConfigArns;

    private DomainDefaultSpaceSettingsKernelGatewayAppSettings() {}
    /**
     * @return A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
     * 
     */
    public List<DomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImage> customImages() {
        return this.customImages == null ? List.of() : this.customImages;
    }
    /**
     * @return The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
     * 
     */
    public Optional<DomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec> defaultResourceSpec() {
        return Optional.ofNullable(this.defaultResourceSpec);
    }
    /**
     * @return The Amazon Resource Name (ARN) of the Lifecycle Configurations.
     * 
     */
    public List<String> lifecycleConfigArns() {
        return this.lifecycleConfigArns == null ? List.of() : this.lifecycleConfigArns;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DomainDefaultSpaceSettingsKernelGatewayAppSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<DomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImage> customImages;
        private @Nullable DomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec defaultResourceSpec;
        private @Nullable List<String> lifecycleConfigArns;
        public Builder() {}
        public Builder(DomainDefaultSpaceSettingsKernelGatewayAppSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customImages = defaults.customImages;
    	      this.defaultResourceSpec = defaults.defaultResourceSpec;
    	      this.lifecycleConfigArns = defaults.lifecycleConfigArns;
        }

        @CustomType.Setter
        public Builder customImages(@Nullable List<DomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImage> customImages) {
            this.customImages = customImages;
            return this;
        }
        public Builder customImages(DomainDefaultSpaceSettingsKernelGatewayAppSettingsCustomImage... customImages) {
            return customImages(List.of(customImages));
        }
        @CustomType.Setter
        public Builder defaultResourceSpec(@Nullable DomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec defaultResourceSpec) {
            this.defaultResourceSpec = defaultResourceSpec;
            return this;
        }
        @CustomType.Setter
        public Builder lifecycleConfigArns(@Nullable List<String> lifecycleConfigArns) {
            this.lifecycleConfigArns = lifecycleConfigArns;
            return this;
        }
        public Builder lifecycleConfigArns(String... lifecycleConfigArns) {
            return lifecycleConfigArns(List.of(lifecycleConfigArns));
        }
        public DomainDefaultSpaceSettingsKernelGatewayAppSettings build() {
            final var o = new DomainDefaultSpaceSettingsKernelGatewayAppSettings();
            o.customImages = customImages;
            o.defaultResourceSpec = defaultResourceSpec;
            o.lifecycleConfigArns = lifecycleConfigArns;
            return o;
        }
    }
}
