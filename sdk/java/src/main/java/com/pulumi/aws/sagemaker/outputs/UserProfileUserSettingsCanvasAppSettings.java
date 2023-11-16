// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.aws.sagemaker.outputs.UserProfileUserSettingsCanvasAppSettingsDirectDeploySettings;
import com.pulumi.aws.sagemaker.outputs.UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSetting;
import com.pulumi.aws.sagemaker.outputs.UserProfileUserSettingsCanvasAppSettingsKendraSettings;
import com.pulumi.aws.sagemaker.outputs.UserProfileUserSettingsCanvasAppSettingsModelRegisterSettings;
import com.pulumi.aws.sagemaker.outputs.UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings;
import com.pulumi.aws.sagemaker.outputs.UserProfileUserSettingsCanvasAppSettingsWorkspaceSettings;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UserProfileUserSettingsCanvasAppSettings {
    /**
     * @return The model deployment settings for the SageMaker Canvas application. See Direct Deploy Settings below.
     * 
     */
    private @Nullable UserProfileUserSettingsCanvasAppSettingsDirectDeploySettings directDeploySettings;
    /**
     * @return The settings for connecting to an external data source with OAuth. See Identity Provider OAuth Settings below.
     * 
     */
    private @Nullable List<UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSetting> identityProviderOauthSettings;
    /**
     * @return The settings for document querying. See Kendra Settings below.
     * 
     */
    private @Nullable UserProfileUserSettingsCanvasAppSettingsKendraSettings kendraSettings;
    /**
     * @return The model registry settings for the SageMaker Canvas application. See Model Register Settings below.
     * 
     */
    private @Nullable UserProfileUserSettingsCanvasAppSettingsModelRegisterSettings modelRegisterSettings;
    /**
     * @return Time series forecast settings for the Canvas app. see Time Series Forecasting Settings below.
     * 
     */
    private @Nullable UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings timeSeriesForecastingSettings;
    /**
     * @return The workspace settings for the SageMaker Canvas application. See Workspace Settings below.
     * 
     */
    private @Nullable UserProfileUserSettingsCanvasAppSettingsWorkspaceSettings workspaceSettings;

    private UserProfileUserSettingsCanvasAppSettings() {}
    /**
     * @return The model deployment settings for the SageMaker Canvas application. See Direct Deploy Settings below.
     * 
     */
    public Optional<UserProfileUserSettingsCanvasAppSettingsDirectDeploySettings> directDeploySettings() {
        return Optional.ofNullable(this.directDeploySettings);
    }
    /**
     * @return The settings for connecting to an external data source with OAuth. See Identity Provider OAuth Settings below.
     * 
     */
    public List<UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSetting> identityProviderOauthSettings() {
        return this.identityProviderOauthSettings == null ? List.of() : this.identityProviderOauthSettings;
    }
    /**
     * @return The settings for document querying. See Kendra Settings below.
     * 
     */
    public Optional<UserProfileUserSettingsCanvasAppSettingsKendraSettings> kendraSettings() {
        return Optional.ofNullable(this.kendraSettings);
    }
    /**
     * @return The model registry settings for the SageMaker Canvas application. See Model Register Settings below.
     * 
     */
    public Optional<UserProfileUserSettingsCanvasAppSettingsModelRegisterSettings> modelRegisterSettings() {
        return Optional.ofNullable(this.modelRegisterSettings);
    }
    /**
     * @return Time series forecast settings for the Canvas app. see Time Series Forecasting Settings below.
     * 
     */
    public Optional<UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings> timeSeriesForecastingSettings() {
        return Optional.ofNullable(this.timeSeriesForecastingSettings);
    }
    /**
     * @return The workspace settings for the SageMaker Canvas application. See Workspace Settings below.
     * 
     */
    public Optional<UserProfileUserSettingsCanvasAppSettingsWorkspaceSettings> workspaceSettings() {
        return Optional.ofNullable(this.workspaceSettings);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UserProfileUserSettingsCanvasAppSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable UserProfileUserSettingsCanvasAppSettingsDirectDeploySettings directDeploySettings;
        private @Nullable List<UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSetting> identityProviderOauthSettings;
        private @Nullable UserProfileUserSettingsCanvasAppSettingsKendraSettings kendraSettings;
        private @Nullable UserProfileUserSettingsCanvasAppSettingsModelRegisterSettings modelRegisterSettings;
        private @Nullable UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings timeSeriesForecastingSettings;
        private @Nullable UserProfileUserSettingsCanvasAppSettingsWorkspaceSettings workspaceSettings;
        public Builder() {}
        public Builder(UserProfileUserSettingsCanvasAppSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.directDeploySettings = defaults.directDeploySettings;
    	      this.identityProviderOauthSettings = defaults.identityProviderOauthSettings;
    	      this.kendraSettings = defaults.kendraSettings;
    	      this.modelRegisterSettings = defaults.modelRegisterSettings;
    	      this.timeSeriesForecastingSettings = defaults.timeSeriesForecastingSettings;
    	      this.workspaceSettings = defaults.workspaceSettings;
        }

        @CustomType.Setter
        public Builder directDeploySettings(@Nullable UserProfileUserSettingsCanvasAppSettingsDirectDeploySettings directDeploySettings) {
            this.directDeploySettings = directDeploySettings;
            return this;
        }
        @CustomType.Setter
        public Builder identityProviderOauthSettings(@Nullable List<UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSetting> identityProviderOauthSettings) {
            this.identityProviderOauthSettings = identityProviderOauthSettings;
            return this;
        }
        public Builder identityProviderOauthSettings(UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSetting... identityProviderOauthSettings) {
            return identityProviderOauthSettings(List.of(identityProviderOauthSettings));
        }
        @CustomType.Setter
        public Builder kendraSettings(@Nullable UserProfileUserSettingsCanvasAppSettingsKendraSettings kendraSettings) {
            this.kendraSettings = kendraSettings;
            return this;
        }
        @CustomType.Setter
        public Builder modelRegisterSettings(@Nullable UserProfileUserSettingsCanvasAppSettingsModelRegisterSettings modelRegisterSettings) {
            this.modelRegisterSettings = modelRegisterSettings;
            return this;
        }
        @CustomType.Setter
        public Builder timeSeriesForecastingSettings(@Nullable UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings timeSeriesForecastingSettings) {
            this.timeSeriesForecastingSettings = timeSeriesForecastingSettings;
            return this;
        }
        @CustomType.Setter
        public Builder workspaceSettings(@Nullable UserProfileUserSettingsCanvasAppSettingsWorkspaceSettings workspaceSettings) {
            this.workspaceSettings = workspaceSettings;
            return this;
        }
        public UserProfileUserSettingsCanvasAppSettings build() {
            final var o = new UserProfileUserSettingsCanvasAppSettings();
            o.directDeploySettings = directDeploySettings;
            o.identityProviderOauthSettings = identityProviderOauthSettings;
            o.kendraSettings = kendraSettings;
            o.modelRegisterSettings = modelRegisterSettings;
            o.timeSeriesForecastingSettings = timeSeriesForecastingSettings;
            o.workspaceSettings = workspaceSettings;
            return o;
        }
    }
}
