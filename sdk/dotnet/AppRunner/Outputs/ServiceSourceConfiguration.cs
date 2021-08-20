// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppRunner.Outputs
{

    [OutputType]
    public sealed class ServiceSourceConfiguration
    {
        /// <summary>
        /// Describes resources needed to authenticate access to some source repositories. See Authentication Configuration below for more details.
        /// </summary>
        public readonly Outputs.ServiceSourceConfigurationAuthenticationConfiguration? AuthenticationConfiguration;
        /// <summary>
        /// Whether continuous integration from the source repository is enabled for the App Runner service. If set to `true`, each repository change (source code commit or new image version) starts a deployment. Defaults to `true`.
        /// </summary>
        public readonly bool? AutoDeploymentsEnabled;
        /// <summary>
        /// Description of a source code repository. See Code Repository below for more details.
        /// </summary>
        public readonly Outputs.ServiceSourceConfigurationCodeRepository? CodeRepository;
        /// <summary>
        /// Description of a source image repository. See Image Repository below for more details.
        /// </summary>
        public readonly Outputs.ServiceSourceConfigurationImageRepository? ImageRepository;

        [OutputConstructor]
        private ServiceSourceConfiguration(
            Outputs.ServiceSourceConfigurationAuthenticationConfiguration? authenticationConfiguration,

            bool? autoDeploymentsEnabled,

            Outputs.ServiceSourceConfigurationCodeRepository? codeRepository,

            Outputs.ServiceSourceConfigurationImageRepository? imageRepository)
        {
            AuthenticationConfiguration = authenticationConfiguration;
            AutoDeploymentsEnabled = autoDeploymentsEnabled;
            CodeRepository = codeRepository;
            ImageRepository = imageRepository;
        }
    }
}