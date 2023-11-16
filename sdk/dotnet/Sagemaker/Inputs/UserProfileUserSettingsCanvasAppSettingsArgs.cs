// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class UserProfileUserSettingsCanvasAppSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The model deployment settings for the SageMaker Canvas application. See Direct Deploy Settings below.
        /// </summary>
        [Input("directDeploySettings")]
        public Input<Inputs.UserProfileUserSettingsCanvasAppSettingsDirectDeploySettingsArgs>? DirectDeploySettings { get; set; }

        [Input("identityProviderOauthSettings")]
        private InputList<Inputs.UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSettingArgs>? _identityProviderOauthSettings;

        /// <summary>
        /// The settings for connecting to an external data source with OAuth. See Identity Provider OAuth Settings below.
        /// </summary>
        public InputList<Inputs.UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSettingArgs> IdentityProviderOauthSettings
        {
            get => _identityProviderOauthSettings ?? (_identityProviderOauthSettings = new InputList<Inputs.UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSettingArgs>());
            set => _identityProviderOauthSettings = value;
        }

        /// <summary>
        /// The settings for document querying. See Kendra Settings below.
        /// </summary>
        [Input("kendraSettings")]
        public Input<Inputs.UserProfileUserSettingsCanvasAppSettingsKendraSettingsArgs>? KendraSettings { get; set; }

        /// <summary>
        /// The model registry settings for the SageMaker Canvas application. See Model Register Settings below.
        /// </summary>
        [Input("modelRegisterSettings")]
        public Input<Inputs.UserProfileUserSettingsCanvasAppSettingsModelRegisterSettingsArgs>? ModelRegisterSettings { get; set; }

        /// <summary>
        /// Time series forecast settings for the Canvas app. see Time Series Forecasting Settings below.
        /// </summary>
        [Input("timeSeriesForecastingSettings")]
        public Input<Inputs.UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsArgs>? TimeSeriesForecastingSettings { get; set; }

        /// <summary>
        /// The workspace settings for the SageMaker Canvas application. See Workspace Settings below.
        /// </summary>
        [Input("workspaceSettings")]
        public Input<Inputs.UserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsArgs>? WorkspaceSettings { get; set; }

        public UserProfileUserSettingsCanvasAppSettingsArgs()
        {
        }
        public static new UserProfileUserSettingsCanvasAppSettingsArgs Empty => new UserProfileUserSettingsCanvasAppSettingsArgs();
    }
}
