// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.FinSpace.Outputs
{

    [OutputType]
    public sealed class KxDataviewSegmentConfiguration
    {
        /// <summary>
        /// The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
        /// </summary>
        public readonly ImmutableArray<string> DbPaths;
        /// <summary>
        /// The name of the volume that you want to attach to a dataview. This volume must be in the same availability zone as the dataview that you are attaching to.
        /// </summary>
        public readonly string VolumeName;

        [OutputConstructor]
        private KxDataviewSegmentConfiguration(
            ImmutableArray<string> dbPaths,

            string volumeName)
        {
            DbPaths = dbPaths;
            VolumeName = volumeName;
        }
    }
}
