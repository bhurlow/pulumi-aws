// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.inputs;

import com.pulumi.aws.sagemaker.inputs.SpaceSpaceSettingsCodeEditorAppSettingsArgs;
import com.pulumi.aws.sagemaker.inputs.SpaceSpaceSettingsCustomFileSystemArgs;
import com.pulumi.aws.sagemaker.inputs.SpaceSpaceSettingsJupyterLabAppSettingsArgs;
import com.pulumi.aws.sagemaker.inputs.SpaceSpaceSettingsJupyterServerAppSettingsArgs;
import com.pulumi.aws.sagemaker.inputs.SpaceSpaceSettingsKernelGatewayAppSettingsArgs;
import com.pulumi.aws.sagemaker.inputs.SpaceSpaceSettingsSpaceStorageSettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SpaceSpaceSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final SpaceSpaceSettingsArgs Empty = new SpaceSpaceSettingsArgs();

    /**
     * The type of app created within the space.
     * 
     */
    @Import(name="appType")
    private @Nullable Output<String> appType;

    /**
     * @return The type of app created within the space.
     * 
     */
    public Optional<Output<String>> appType() {
        return Optional.ofNullable(this.appType);
    }

    /**
     * The Code Editor application settings. See Code Editor App Settings below.
     * 
     */
    @Import(name="codeEditorAppSettings")
    private @Nullable Output<SpaceSpaceSettingsCodeEditorAppSettingsArgs> codeEditorAppSettings;

    /**
     * @return The Code Editor application settings. See Code Editor App Settings below.
     * 
     */
    public Optional<Output<SpaceSpaceSettingsCodeEditorAppSettingsArgs>> codeEditorAppSettings() {
        return Optional.ofNullable(this.codeEditorAppSettings);
    }

    /**
     * A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
     * 
     */
    @Import(name="customFileSystems")
    private @Nullable Output<List<SpaceSpaceSettingsCustomFileSystemArgs>> customFileSystems;

    /**
     * @return A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
     * 
     */
    public Optional<Output<List<SpaceSpaceSettingsCustomFileSystemArgs>>> customFileSystems() {
        return Optional.ofNullable(this.customFileSystems);
    }

    /**
     * The settings for the JupyterLab application. See Jupyter Lab App Settings below.
     * 
     */
    @Import(name="jupyterLabAppSettings")
    private @Nullable Output<SpaceSpaceSettingsJupyterLabAppSettingsArgs> jupyterLabAppSettings;

    /**
     * @return The settings for the JupyterLab application. See Jupyter Lab App Settings below.
     * 
     */
    public Optional<Output<SpaceSpaceSettingsJupyterLabAppSettingsArgs>> jupyterLabAppSettings() {
        return Optional.ofNullable(this.jupyterLabAppSettings);
    }

    /**
     * The Jupyter server&#39;s app settings. See Jupyter Server App Settings below.
     * 
     */
    @Import(name="jupyterServerAppSettings")
    private @Nullable Output<SpaceSpaceSettingsJupyterServerAppSettingsArgs> jupyterServerAppSettings;

    /**
     * @return The Jupyter server&#39;s app settings. See Jupyter Server App Settings below.
     * 
     */
    public Optional<Output<SpaceSpaceSettingsJupyterServerAppSettingsArgs>> jupyterServerAppSettings() {
        return Optional.ofNullable(this.jupyterServerAppSettings);
    }

    /**
     * The kernel gateway app settings. See Kernel Gateway App Settings below.
     * 
     */
    @Import(name="kernelGatewayAppSettings")
    private @Nullable Output<SpaceSpaceSettingsKernelGatewayAppSettingsArgs> kernelGatewayAppSettings;

    /**
     * @return The kernel gateway app settings. See Kernel Gateway App Settings below.
     * 
     */
    public Optional<Output<SpaceSpaceSettingsKernelGatewayAppSettingsArgs>> kernelGatewayAppSettings() {
        return Optional.ofNullable(this.kernelGatewayAppSettings);
    }

    @Import(name="spaceStorageSettings")
    private @Nullable Output<SpaceSpaceSettingsSpaceStorageSettingsArgs> spaceStorageSettings;

    public Optional<Output<SpaceSpaceSettingsSpaceStorageSettingsArgs>> spaceStorageSettings() {
        return Optional.ofNullable(this.spaceStorageSettings);
    }

    private SpaceSpaceSettingsArgs() {}

    private SpaceSpaceSettingsArgs(SpaceSpaceSettingsArgs $) {
        this.appType = $.appType;
        this.codeEditorAppSettings = $.codeEditorAppSettings;
        this.customFileSystems = $.customFileSystems;
        this.jupyterLabAppSettings = $.jupyterLabAppSettings;
        this.jupyterServerAppSettings = $.jupyterServerAppSettings;
        this.kernelGatewayAppSettings = $.kernelGatewayAppSettings;
        this.spaceStorageSettings = $.spaceStorageSettings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SpaceSpaceSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SpaceSpaceSettingsArgs $;

        public Builder() {
            $ = new SpaceSpaceSettingsArgs();
        }

        public Builder(SpaceSpaceSettingsArgs defaults) {
            $ = new SpaceSpaceSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param appType The type of app created within the space.
         * 
         * @return builder
         * 
         */
        public Builder appType(@Nullable Output<String> appType) {
            $.appType = appType;
            return this;
        }

        /**
         * @param appType The type of app created within the space.
         * 
         * @return builder
         * 
         */
        public Builder appType(String appType) {
            return appType(Output.of(appType));
        }

        /**
         * @param codeEditorAppSettings The Code Editor application settings. See Code Editor App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder codeEditorAppSettings(@Nullable Output<SpaceSpaceSettingsCodeEditorAppSettingsArgs> codeEditorAppSettings) {
            $.codeEditorAppSettings = codeEditorAppSettings;
            return this;
        }

        /**
         * @param codeEditorAppSettings The Code Editor application settings. See Code Editor App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder codeEditorAppSettings(SpaceSpaceSettingsCodeEditorAppSettingsArgs codeEditorAppSettings) {
            return codeEditorAppSettings(Output.of(codeEditorAppSettings));
        }

        /**
         * @param customFileSystems A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
         * 
         * @return builder
         * 
         */
        public Builder customFileSystems(@Nullable Output<List<SpaceSpaceSettingsCustomFileSystemArgs>> customFileSystems) {
            $.customFileSystems = customFileSystems;
            return this;
        }

        /**
         * @param customFileSystems A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
         * 
         * @return builder
         * 
         */
        public Builder customFileSystems(List<SpaceSpaceSettingsCustomFileSystemArgs> customFileSystems) {
            return customFileSystems(Output.of(customFileSystems));
        }

        /**
         * @param customFileSystems A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
         * 
         * @return builder
         * 
         */
        public Builder customFileSystems(SpaceSpaceSettingsCustomFileSystemArgs... customFileSystems) {
            return customFileSystems(List.of(customFileSystems));
        }

        /**
         * @param jupyterLabAppSettings The settings for the JupyterLab application. See Jupyter Lab App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder jupyterLabAppSettings(@Nullable Output<SpaceSpaceSettingsJupyterLabAppSettingsArgs> jupyterLabAppSettings) {
            $.jupyterLabAppSettings = jupyterLabAppSettings;
            return this;
        }

        /**
         * @param jupyterLabAppSettings The settings for the JupyterLab application. See Jupyter Lab App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder jupyterLabAppSettings(SpaceSpaceSettingsJupyterLabAppSettingsArgs jupyterLabAppSettings) {
            return jupyterLabAppSettings(Output.of(jupyterLabAppSettings));
        }

        /**
         * @param jupyterServerAppSettings The Jupyter server&#39;s app settings. See Jupyter Server App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder jupyterServerAppSettings(@Nullable Output<SpaceSpaceSettingsJupyterServerAppSettingsArgs> jupyterServerAppSettings) {
            $.jupyterServerAppSettings = jupyterServerAppSettings;
            return this;
        }

        /**
         * @param jupyterServerAppSettings The Jupyter server&#39;s app settings. See Jupyter Server App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder jupyterServerAppSettings(SpaceSpaceSettingsJupyterServerAppSettingsArgs jupyterServerAppSettings) {
            return jupyterServerAppSettings(Output.of(jupyterServerAppSettings));
        }

        /**
         * @param kernelGatewayAppSettings The kernel gateway app settings. See Kernel Gateway App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder kernelGatewayAppSettings(@Nullable Output<SpaceSpaceSettingsKernelGatewayAppSettingsArgs> kernelGatewayAppSettings) {
            $.kernelGatewayAppSettings = kernelGatewayAppSettings;
            return this;
        }

        /**
         * @param kernelGatewayAppSettings The kernel gateway app settings. See Kernel Gateway App Settings below.
         * 
         * @return builder
         * 
         */
        public Builder kernelGatewayAppSettings(SpaceSpaceSettingsKernelGatewayAppSettingsArgs kernelGatewayAppSettings) {
            return kernelGatewayAppSettings(Output.of(kernelGatewayAppSettings));
        }

        public Builder spaceStorageSettings(@Nullable Output<SpaceSpaceSettingsSpaceStorageSettingsArgs> spaceStorageSettings) {
            $.spaceStorageSettings = spaceStorageSettings;
            return this;
        }

        public Builder spaceStorageSettings(SpaceSpaceSettingsSpaceStorageSettingsArgs spaceStorageSettings) {
            return spaceStorageSettings(Output.of(spaceStorageSettings));
        }

        public SpaceSpaceSettingsArgs build() {
            return $;
        }
    }

}
