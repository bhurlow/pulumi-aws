// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class SpotInstanceRequestCpuOptions
    {
        /// <summary>
        /// Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only. Valid values are `enabled` and `disabled`.
        /// </summary>
        public readonly string? AmdSevSnp;
        /// <summary>
        /// Sets the number of CPU cores for an instance. This option is only supported on creation of instance type that support CPU Options [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
        /// </summary>
        public readonly int? CoreCount;
        /// <summary>
        /// If set to 1, hyperthreading is disabled on the launched instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
        /// </summary>
        public readonly int? ThreadsPerCore;

        [OutputConstructor]
        private SpotInstanceRequestCpuOptions(
            string? amdSevSnp,

            int? coreCount,

            int? threadsPerCore)
        {
            AmdSevSnp = amdSevSnp;
            CoreCount = coreCount;
            ThreadsPerCore = threadsPerCore;
        }
    }
}
