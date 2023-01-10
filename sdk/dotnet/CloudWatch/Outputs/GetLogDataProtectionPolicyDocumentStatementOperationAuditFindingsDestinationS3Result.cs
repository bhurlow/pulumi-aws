// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Outputs
{

    [OutputType]
    public sealed class GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3Result
    {
        /// <summary>
        /// Name of the S3 Bucket to send findings to.
        /// </summary>
        public readonly string Bucket;

        [OutputConstructor]
        private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3Result(string bucket)
        {
            Bucket = bucket;
        }
    }
}
