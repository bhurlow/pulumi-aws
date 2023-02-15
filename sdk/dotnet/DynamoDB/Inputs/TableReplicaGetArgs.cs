// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Inputs
{

    public sealed class TableReplicaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the table
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Whether to enable Point In Time Recovery for the replica. Default is `false`.
        /// </summary>
        [Input("pointInTimeRecovery")]
        public Input<bool>? PointInTimeRecovery { get; set; }

        /// <summary>
        /// Whether to propagate the global table's tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
        /// </summary>
        [Input("propagateTags")]
        public Input<bool>? PropagateTags { get; set; }

        /// <summary>
        /// Region name of the replica.
        /// </summary>
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        /// <summary>
        /// ARN of the Table Stream. Only available when `stream_enabled = true`
        /// </summary>
        [Input("streamArn")]
        public Input<string>? StreamArn { get; set; }

        /// <summary>
        /// Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
        /// </summary>
        [Input("streamLabel")]
        public Input<string>? StreamLabel { get; set; }

        public TableReplicaGetArgs()
        {
        }
        public static new TableReplicaGetArgs Empty => new TableReplicaGetArgs();
    }
}
