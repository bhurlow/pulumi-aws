// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.TimestreamWrite.Outputs
{

    [OutputType]
    public sealed class TableSchema
    {
        /// <summary>
        /// A non-empty list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed. See Composite Partition Key below for more details.
        /// </summary>
        public readonly Outputs.TableSchemaCompositePartitionKey? CompositePartitionKey;

        [OutputConstructor]
        private TableSchema(Outputs.TableSchemaCompositePartitionKey? compositePartitionKey)
        {
            CompositePartitionKey = compositePartitionKey;
        }
    }
}
