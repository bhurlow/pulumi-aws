// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Keyspaces.Inputs
{

    public sealed class TableSchemaDefinitionGetArgs : Pulumi.ResourceArgs
    {
        [Input("clusteringKeys")]
        private InputList<Inputs.TableSchemaDefinitionClusteringKeyGetArgs>? _clusteringKeys;

        /// <summary>
        /// The columns that are part of the clustering key of the table.
        /// </summary>
        public InputList<Inputs.TableSchemaDefinitionClusteringKeyGetArgs> ClusteringKeys
        {
            get => _clusteringKeys ?? (_clusteringKeys = new InputList<Inputs.TableSchemaDefinitionClusteringKeyGetArgs>());
            set => _clusteringKeys = value;
        }

        [Input("columns", required: true)]
        private InputList<Inputs.TableSchemaDefinitionColumnGetArgs>? _columns;

        /// <summary>
        /// The regular columns of the table.
        /// </summary>
        public InputList<Inputs.TableSchemaDefinitionColumnGetArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.TableSchemaDefinitionColumnGetArgs>());
            set => _columns = value;
        }

        [Input("partitionKeys", required: true)]
        private InputList<Inputs.TableSchemaDefinitionPartitionKeyGetArgs>? _partitionKeys;

        /// <summary>
        /// The columns that are part of the partition key of the table .
        /// </summary>
        public InputList<Inputs.TableSchemaDefinitionPartitionKeyGetArgs> PartitionKeys
        {
            get => _partitionKeys ?? (_partitionKeys = new InputList<Inputs.TableSchemaDefinitionPartitionKeyGetArgs>());
            set => _partitionKeys = value;
        }

        [Input("staticColumns")]
        private InputList<Inputs.TableSchemaDefinitionStaticColumnGetArgs>? _staticColumns;

        /// <summary>
        /// The columns that have been defined as `STATIC`. Static columns store values that are shared by all rows in the same partition.
        /// </summary>
        public InputList<Inputs.TableSchemaDefinitionStaticColumnGetArgs> StaticColumns
        {
            get => _staticColumns ?? (_staticColumns = new InputList<Inputs.TableSchemaDefinitionStaticColumnGetArgs>());
            set => _staticColumns = value;
        }

        public TableSchemaDefinitionGetArgs()
        {
        }
    }
}
