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
    public sealed class GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehoseResult
    {
        /// <summary>
        /// Name of the Kinesis Firehose Delivery Stream to send findings to.
        /// </summary>
        public readonly string DeliveryStream;

        [OutputConstructor]
        private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehoseResult(string deliveryStream)
        {
            DeliveryStream = deliveryStream;
        }
    }
}
