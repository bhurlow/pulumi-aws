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
    public sealed class SpaceSpaceSettingsJupyterLabAppSettingsCodeRepository
    {
        /// <summary>
        /// The URL of the Git repository.
        /// </summary>
        public readonly string RepositoryUrl;

        [OutputConstructor]
        private SpaceSpaceSettingsJupyterLabAppSettingsCodeRepository(string repositoryUrl)
        {
            RepositoryUrl = repositoryUrl;
        }
    }
}
