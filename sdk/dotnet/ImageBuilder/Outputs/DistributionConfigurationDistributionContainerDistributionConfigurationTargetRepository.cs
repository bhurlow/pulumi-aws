// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ImageBuilder.Outputs
{

    [OutputType]
    public sealed class DistributionConfigurationDistributionContainerDistributionConfigurationTargetRepository
    {
        /// <summary>
        /// The name of the container repository where the output container image is stored. This name is prefixed by the repository location.
        /// </summary>
        public readonly string RepositoryName;
        /// <summary>
        /// The service in which this image is registered. Valid values: `ECR`.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private DistributionConfigurationDistributionContainerDistributionConfigurationTargetRepository(
            string repositoryName,

            string service)
        {
            RepositoryName = repositoryName;
            Service = service;
        }
    }
}
