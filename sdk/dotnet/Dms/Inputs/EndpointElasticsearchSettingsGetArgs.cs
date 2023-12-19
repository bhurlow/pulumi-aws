// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms.Inputs
{

    public sealed class EndpointElasticsearchSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint for the OpenSearch cluster.
        /// </summary>
        [Input("endpointUri", required: true)]
        public Input<string> EndpointUri { get; set; } = null!;

        /// <summary>
        /// Maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster. Default is `300`.
        /// </summary>
        [Input("errorRetryDuration")]
        public Input<int>? ErrorRetryDuration { get; set; }

        /// <summary>
        /// Maximum percentage of records that can fail to be written before a full load operation stops. Default is `10`.
        /// </summary>
        [Input("fullLoadErrorPercentage")]
        public Input<int>? FullLoadErrorPercentage { get; set; }

        /// <summary>
        /// ARN of the IAM Role with permissions to write to the OpenSearch cluster.
        /// </summary>
        [Input("serviceAccessRoleArn", required: true)]
        public Input<string> ServiceAccessRoleArn { get; set; } = null!;

        /// <summary>
        /// Enable to migrate documentation using the documentation type `_doc`. OpenSearch and an Elasticsearch clusters only support the _doc documentation type in versions 7.x and later. The default value is `false`.
        /// </summary>
        [Input("useNewMappingType")]
        public Input<bool>? UseNewMappingType { get; set; }

        public EndpointElasticsearchSettingsGetArgs()
        {
        }
        public static new EndpointElasticsearchSettingsGetArgs Empty => new EndpointElasticsearchSettingsGetArgs();
    }
}
