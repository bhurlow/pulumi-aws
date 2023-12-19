// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms.Outputs
{

    [OutputType]
    public sealed class EndpointElasticsearchSettings
    {
        /// <summary>
        /// Endpoint for the OpenSearch cluster.
        /// </summary>
        public readonly string EndpointUri;
        /// <summary>
        /// Maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster. Default is `300`.
        /// </summary>
        public readonly int? ErrorRetryDuration;
        /// <summary>
        /// Maximum percentage of records that can fail to be written before a full load operation stops. Default is `10`.
        /// </summary>
        public readonly int? FullLoadErrorPercentage;
        /// <summary>
        /// ARN of the IAM Role with permissions to write to the OpenSearch cluster.
        /// </summary>
        public readonly string ServiceAccessRoleArn;
        /// <summary>
        /// Enable to migrate documentation using the documentation type `_doc`. OpenSearch and an Elasticsearch clusters only support the _doc documentation type in versions 7.x and later. The default value is `false`.
        /// </summary>
        public readonly bool? UseNewMappingType;

        [OutputConstructor]
        private EndpointElasticsearchSettings(
            string endpointUri,

            int? errorRetryDuration,

            int? fullLoadErrorPercentage,

            string serviceAccessRoleArn,

            bool? useNewMappingType)
        {
            EndpointUri = endpointUri;
            ErrorRetryDuration = errorRetryDuration;
            FullLoadErrorPercentage = fullLoadErrorPercentage;
            ServiceAccessRoleArn = serviceAccessRoleArn;
            UseNewMappingType = useNewMappingType;
        }
    }
}
