// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer.Outputs
{

    [OutputType]
    public sealed class WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocation
    {
        /// <summary>
        /// Specifies the details for the EFS file being copied.
        /// </summary>
        public readonly Outputs.WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocation? EfsFileLocation;
        /// <summary>
        /// Specifies the details for the S3 file being copied.
        /// </summary>
        public readonly Outputs.WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocation? S3FileLocation;

        [OutputConstructor]
        private WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocation(
            Outputs.WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocation? efsFileLocation,

            Outputs.WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocation? s3FileLocation)
        {
            EfsFileLocation = efsFileLocation;
            S3FileLocation = s3FileLocation;
        }
    }
}
