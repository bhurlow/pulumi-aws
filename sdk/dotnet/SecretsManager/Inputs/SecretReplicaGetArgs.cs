// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecretsManager.Inputs
{

    public sealed class SecretReplicaGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account's default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Date that you last accessed the secret in the Region.
        /// </summary>
        [Input("lastAccessedDate")]
        public Input<string>? LastAccessedDate { get; set; }

        /// <summary>
        /// Region for replicating the secret.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Status can be `InProgress`, `Failed`, or `InSync`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Message such as `Replication succeeded` or `Secret with this name already exists in this region`.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        public SecretReplicaGetArgs()
        {
        }
    }
}
