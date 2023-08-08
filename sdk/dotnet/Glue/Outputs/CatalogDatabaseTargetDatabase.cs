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
    public sealed class CatalogDatabaseTargetDatabase
    {
        /// <summary>
        /// ID of the Data Catalog in which the database resides.
        /// </summary>
        public readonly string CatalogId;
        /// <summary>
        /// Name of the catalog database.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// Region of the target database.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private CatalogDatabaseTargetDatabase(
            string catalogId,

            string databaseName,

            string? region)
        {
            CatalogId = catalogId;
            DatabaseName = databaseName;
            Region = region;
        }
    }
}
