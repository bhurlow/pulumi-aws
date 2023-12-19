// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.FinSpace.Inputs
{

    public sealed class KxClusterDatabaseGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cacheConfigurations")]
        private InputList<Inputs.KxClusterDatabaseCacheConfigurationGetArgs>? _cacheConfigurations;

        /// <summary>
        /// Configuration details for the disk cache to increase performance reading from a KX database mounted to the cluster. See cache_configurations.
        /// </summary>
        public InputList<Inputs.KxClusterDatabaseCacheConfigurationGetArgs> CacheConfigurations
        {
            get => _cacheConfigurations ?? (_cacheConfigurations = new InputList<Inputs.KxClusterDatabaseCacheConfigurationGetArgs>());
            set => _cacheConfigurations = value;
        }

        /// <summary>
        /// A unique identifier of the changeset that is associated with the cluster.
        /// </summary>
        [Input("changesetId")]
        public Input<string>? ChangesetId { get; set; }

        /// <summary>
        /// Name of the KX database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the dataview to be used for caching historical data on disk. You cannot update to a different dataview name once a cluster is created. Use `lifecycle` `ignore_changes` for database to prevent any undesirable behaviors.
        /// </summary>
        [Input("dataviewName")]
        public Input<string>? DataviewName { get; set; }

        public KxClusterDatabaseGetArgs()
        {
        }
        public static new KxClusterDatabaseGetArgs Empty => new KxClusterDatabaseGetArgs();
    }
}
