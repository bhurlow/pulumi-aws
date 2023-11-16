// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class DomainDefaultUserSettingsCanvasAppSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The model deployment settings for the SageMaker Canvas application. See Direct Deploy Settings below.
        /// </summary>
        [Input("directDeploySettings")]
        public Input<Inputs.DomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettingsGetArgs>? DirectDeploySettings { get; set; }

        [Input("identityProviderOauthSettings")]
        private InputList<Inputs.DomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSettingGetArgs>? _identityProviderOauthSettings;

        /// <summary>
        /// The settings for connecting to an external data source with OAuth. See Identity Provider OAuth Settings below.
        /// </summary>
        public InputList<Inputs.DomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSettingGetArgs> IdentityProviderOauthSettings
        {
            get => _identityProviderOauthSettings ?? (_identityProviderOauthSettings = new InputList<Inputs.DomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSettingGetArgs>());
            set => _identityProviderOauthSettings = value;
        }

        /// <summary>
        /// The settings for document querying. See Kendra Settings below.
        /// </summary>
        [Input("kendraSettings")]
        public Input<Inputs.DomainDefaultUserSettingsCanvasAppSettingsKendraSettingsGetArgs>? KendraSettings { get; set; }

        /// <summary>
        /// The model registry settings for the SageMaker Canvas application. See Model Register Settings below.
        /// </summary>
        [Input("modelRegisterSettings")]
        public Input<Inputs.DomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettingsGetArgs>? ModelRegisterSettings { get; set; }

        /// <summary>
        /// Time series forecast settings for the Canvas app. See Time Series Forecasting Settings below.
        /// </summary>
        [Input("timeSeriesForecastingSettings")]
        public Input<Inputs.DomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsGetArgs>? TimeSeriesForecastingSettings { get; set; }

        /// <summary>
        /// The workspace settings for the SageMaker Canvas application. See Workspace Settings below.
        /// </summary>
        [Input("workspaceSettings")]
        public Input<Inputs.DomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettingsGetArgs>? WorkspaceSettings { get; set; }

        public DomainDefaultUserSettingsCanvasAppSettingsGetArgs()
        {
        }
        public static new DomainDefaultUserSettingsCanvasAppSettingsGetArgs Empty => new DomainDefaultUserSettingsCanvasAppSettingsGetArgs();
    }
}
