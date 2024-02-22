// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.aws.sagemaker.outputs.SpaceSpaceSettingsCodeEditorAppSettings;
import com.pulumi.aws.sagemaker.outputs.SpaceSpaceSettingsCustomFileSystem;
import com.pulumi.aws.sagemaker.outputs.SpaceSpaceSettingsJupyterLabAppSettings;
import com.pulumi.aws.sagemaker.outputs.SpaceSpaceSettingsJupyterServerAppSettings;
import com.pulumi.aws.sagemaker.outputs.SpaceSpaceSettingsKernelGatewayAppSettings;
import com.pulumi.aws.sagemaker.outputs.SpaceSpaceSettingsSpaceStorageSettings;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SpaceSpaceSettings {
    /**
     * @return The type of app created within the space.
     * 
     */
    private @Nullable String appType;
    /**
     * @return The Code Editor application settings. See Code Editor App Settings below.
     * 
     */
    private @Nullable SpaceSpaceSettingsCodeEditorAppSettings codeEditorAppSettings;
    /**
     * @return A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
     * 
     */
    private @Nullable List<SpaceSpaceSettingsCustomFileSystem> customFileSystems;
    /**
     * @return The settings for the JupyterLab application. See Jupyter Lab App Settings below.
     * 
     */
    private @Nullable SpaceSpaceSettingsJupyterLabAppSettings jupyterLabAppSettings;
    /**
     * @return The Jupyter server&#39;s app settings. See Jupyter Server App Settings below.
     * 
     */
    private @Nullable SpaceSpaceSettingsJupyterServerAppSettings jupyterServerAppSettings;
    /**
     * @return The kernel gateway app settings. See Kernel Gateway App Settings below.
     * 
     */
    private @Nullable SpaceSpaceSettingsKernelGatewayAppSettings kernelGatewayAppSettings;
    private @Nullable SpaceSpaceSettingsSpaceStorageSettings spaceStorageSettings;

    private SpaceSpaceSettings() {}
    /**
     * @return The type of app created within the space.
     * 
     */
    public Optional<String> appType() {
        return Optional.ofNullable(this.appType);
    }
    /**
     * @return The Code Editor application settings. See Code Editor App Settings below.
     * 
     */
    public Optional<SpaceSpaceSettingsCodeEditorAppSettings> codeEditorAppSettings() {
        return Optional.ofNullable(this.codeEditorAppSettings);
    }
    /**
     * @return A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
     * 
     */
    public List<SpaceSpaceSettingsCustomFileSystem> customFileSystems() {
        return this.customFileSystems == null ? List.of() : this.customFileSystems;
    }
    /**
     * @return The settings for the JupyterLab application. See Jupyter Lab App Settings below.
     * 
     */
    public Optional<SpaceSpaceSettingsJupyterLabAppSettings> jupyterLabAppSettings() {
        return Optional.ofNullable(this.jupyterLabAppSettings);
    }
    /**
     * @return The Jupyter server&#39;s app settings. See Jupyter Server App Settings below.
     * 
     */
    public Optional<SpaceSpaceSettingsJupyterServerAppSettings> jupyterServerAppSettings() {
        return Optional.ofNullable(this.jupyterServerAppSettings);
    }
    /**
     * @return The kernel gateway app settings. See Kernel Gateway App Settings below.
     * 
     */
    public Optional<SpaceSpaceSettingsKernelGatewayAppSettings> kernelGatewayAppSettings() {
        return Optional.ofNullable(this.kernelGatewayAppSettings);
    }
    public Optional<SpaceSpaceSettingsSpaceStorageSettings> spaceStorageSettings() {
        return Optional.ofNullable(this.spaceStorageSettings);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SpaceSpaceSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String appType;
        private @Nullable SpaceSpaceSettingsCodeEditorAppSettings codeEditorAppSettings;
        private @Nullable List<SpaceSpaceSettingsCustomFileSystem> customFileSystems;
        private @Nullable SpaceSpaceSettingsJupyterLabAppSettings jupyterLabAppSettings;
        private @Nullable SpaceSpaceSettingsJupyterServerAppSettings jupyterServerAppSettings;
        private @Nullable SpaceSpaceSettingsKernelGatewayAppSettings kernelGatewayAppSettings;
        private @Nullable SpaceSpaceSettingsSpaceStorageSettings spaceStorageSettings;
        public Builder() {}
        public Builder(SpaceSpaceSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appType = defaults.appType;
    	      this.codeEditorAppSettings = defaults.codeEditorAppSettings;
    	      this.customFileSystems = defaults.customFileSystems;
    	      this.jupyterLabAppSettings = defaults.jupyterLabAppSettings;
    	      this.jupyterServerAppSettings = defaults.jupyterServerAppSettings;
    	      this.kernelGatewayAppSettings = defaults.kernelGatewayAppSettings;
    	      this.spaceStorageSettings = defaults.spaceStorageSettings;
        }

        @CustomType.Setter
        public Builder appType(@Nullable String appType) {

            this.appType = appType;
            return this;
        }
        @CustomType.Setter
        public Builder codeEditorAppSettings(@Nullable SpaceSpaceSettingsCodeEditorAppSettings codeEditorAppSettings) {

            this.codeEditorAppSettings = codeEditorAppSettings;
            return this;
        }
        @CustomType.Setter
        public Builder customFileSystems(@Nullable List<SpaceSpaceSettingsCustomFileSystem> customFileSystems) {

            this.customFileSystems = customFileSystems;
            return this;
        }
        public Builder customFileSystems(SpaceSpaceSettingsCustomFileSystem... customFileSystems) {
            return customFileSystems(List.of(customFileSystems));
        }
        @CustomType.Setter
        public Builder jupyterLabAppSettings(@Nullable SpaceSpaceSettingsJupyterLabAppSettings jupyterLabAppSettings) {

            this.jupyterLabAppSettings = jupyterLabAppSettings;
            return this;
        }
        @CustomType.Setter
        public Builder jupyterServerAppSettings(@Nullable SpaceSpaceSettingsJupyterServerAppSettings jupyterServerAppSettings) {

            this.jupyterServerAppSettings = jupyterServerAppSettings;
            return this;
        }
        @CustomType.Setter
        public Builder kernelGatewayAppSettings(@Nullable SpaceSpaceSettingsKernelGatewayAppSettings kernelGatewayAppSettings) {

            this.kernelGatewayAppSettings = kernelGatewayAppSettings;
            return this;
        }
        @CustomType.Setter
        public Builder spaceStorageSettings(@Nullable SpaceSpaceSettingsSpaceStorageSettings spaceStorageSettings) {

            this.spaceStorageSettings = spaceStorageSettings;
            return this;
        }
        public SpaceSpaceSettings build() {
            final var _resultValue = new SpaceSpaceSettings();
            _resultValue.appType = appType;
            _resultValue.codeEditorAppSettings = codeEditorAppSettings;
            _resultValue.customFileSystems = customFileSystems;
            _resultValue.jupyterLabAppSettings = jupyterLabAppSettings;
            _resultValue.jupyterServerAppSettings = jupyterServerAppSettings;
            _resultValue.kernelGatewayAppSettings = kernelGatewayAppSettings;
            _resultValue.spaceStorageSettings = spaceStorageSettings;
            return _resultValue;
        }
    }
}
