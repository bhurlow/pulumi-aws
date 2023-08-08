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
    public sealed class EndpointConfigurationProductionVariantServerlessConfig
    {
        /// <summary>
        /// The maximum number of concurrent invocations your serverless endpoint can process. Valid values are between `1` and `200`.
        /// </summary>
        public readonly int MaxConcurrency;
        /// <summary>
        /// The memory size of your serverless endpoint. Valid values are in 1 GB increments: `1024` MB, `2048` MB, `3072` MB, `4096` MB, `5120` MB, or `6144` MB.
        /// </summary>
        public readonly int MemorySizeInMb;
        /// <summary>
        /// The amount of provisioned concurrency to allocate for the serverless endpoint. Should be less than or equal to `max_concurrency`. Valid values are between `1` and `200`.
        /// </summary>
        public readonly int? ProvisionedConcurrency;

        [OutputConstructor]
        private EndpointConfigurationProductionVariantServerlessConfig(
            int maxConcurrency,

            int memorySizeInMb,

            int? provisionedConcurrency)
        {
            MaxConcurrency = maxConcurrency;
            MemorySizeInMb = memorySizeInMb;
            ProvisionedConcurrency = provisionedConcurrency;
        }
    }
}
