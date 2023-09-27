// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fsx.Outputs
{

    [OutputType]
    public sealed class GetOntapFileSystemDiskIopsConfigurationResult
    {
        /// <summary>
        /// The total number of SSD IOPS provisioned for the file system.
        /// </summary>
        public readonly int Iops;
        /// <summary>
        /// Specifies whether the file system is using the `AUTOMATIC` setting of SSD IOPS of 3 IOPS per GB of storage capacity, or if it using a `USER_PROVISIONED` value.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private GetOntapFileSystemDiskIopsConfigurationResult(
            int iops,

            string mode)
        {
            Iops = iops;
            Mode = mode;
        }
    }
}
