// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Outputs
{

    [OutputType]
    public sealed class CrawlerCatalogTarget
    {
        /// <summary>
        /// The name of the connection to use to connect to the Delta table target.
        /// </summary>
        public readonly string? ConnectionName;
        /// <summary>
        /// The name of the Glue database to be synchronized.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// A valid Amazon SQS ARN.
        /// </summary>
        public readonly string? DlqEventQueueArn;
        /// <summary>
        /// A valid Amazon SQS ARN.
        /// </summary>
        public readonly string? EventQueueArn;
        /// <summary>
        /// A list of catalog tables to be synchronized.
        /// </summary>
        public readonly ImmutableArray<string> Tables;

        [OutputConstructor]
        private CrawlerCatalogTarget(
            string? connectionName,

            string databaseName,

            string? dlqEventQueueArn,

            string? eventQueueArn,

            ImmutableArray<string> tables)
        {
            ConnectionName = connectionName;
            DatabaseName = databaseName;
            DlqEventQueueArn = dlqEventQueueArn;
            EventQueueArn = eventQueueArn;
            Tables = tables;
        }
    }
}
