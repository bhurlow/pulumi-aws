// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class DomainDefaultUserSettingsTensorBoardAppSettings
    {
        /// <summary>
        /// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
        /// </summary>
        public readonly Outputs.DomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpec? DefaultResourceSpec;

        [OutputConstructor]
        private DomainDefaultUserSettingsTensorBoardAppSettings(Outputs.DomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpec? defaultResourceSpec)
        {
            DefaultResourceSpec = defaultResourceSpec;
        }
    }
}
