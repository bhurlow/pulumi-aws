// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.FinSpace.Inputs
{

    public sealed class KxClusterSavedownStorageConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of temporary storage in gigabytes. Must be between 10 and 16000.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Type of writeable storage space for temporarily storing your savedown data. The valid values are:
        /// * SDS01 - This type represents 3000 IOPS and io2 ebs volume type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The name of the kdb volume that you want to use as writeable save-down storage for clusters.
        /// </summary>
        [Input("volumeName")]
        public Input<string>? VolumeName { get; set; }

        public KxClusterSavedownStorageConfigurationArgs()
        {
        }
        public static new KxClusterSavedownStorageConfigurationArgs Empty => new KxClusterSavedownStorageConfigurationArgs();
    }
}
