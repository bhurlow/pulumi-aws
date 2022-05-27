// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Keyspaces.Outputs
{

    [OutputType]
    public sealed class TableSchemaDefinitionClusteringKey
    {
        /// <summary>
        /// The name of the clustering key column.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The order modifier. Valid values: `ASC`, `DESC`.
        /// </summary>
        public readonly string OrderBy;

        [OutputConstructor]
        private TableSchemaDefinitionClusteringKey(
            string name,

            string orderBy)
        {
            Name = name;
            OrderBy = orderBy;
        }
    }
}
