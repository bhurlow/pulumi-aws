// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.FinSpace.Inputs
{

    public sealed class KxClusterDatabaseCacheConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of disk cache.
        /// </summary>
        [Input("cacheType", required: true)]
        public Input<string> CacheType { get; set; } = null!;

        [Input("dbPaths")]
        private InputList<string>? _dbPaths;

        /// <summary>
        /// Paths within the database to cache.
        /// </summary>
        public InputList<string> DbPaths
        {
            get => _dbPaths ?? (_dbPaths = new InputList<string>());
            set => _dbPaths = value;
        }

        public KxClusterDatabaseCacheConfigurationArgs()
        {
        }
        public static new KxClusterDatabaseCacheConfigurationArgs Empty => new KxClusterDatabaseCacheConfigurationArgs();
    }
}
