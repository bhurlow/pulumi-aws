// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Inputs
{

    public sealed class ClusterS3ImportArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket name where your backup is stored
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// Can be blank, but is the path to your backup
        /// </summary>
        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        /// <summary>
        /// Role applied to load the data.
        /// </summary>
        [Input("ingestionRole", required: true)]
        public Input<string> IngestionRole { get; set; } = null!;

        /// <summary>
        /// Source engine for the backup
        /// </summary>
        [Input("sourceEngine", required: true)]
        public Input<string> SourceEngine { get; set; } = null!;

        /// <summary>
        /// Version of the source engine used to make the backup
        /// 
        /// This will not recreate the resource if the S3 object changes in some way. It's only used to initialize the database. This only works currently with the aurora engine. See AWS for currently supported engines and options. See [Aurora S3 Migration Docs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Migrating.ExtMySQL.html#AuroraMySQL.Migrating.ExtMySQL.S3).
        /// </summary>
        [Input("sourceEngineVersion", required: true)]
        public Input<string> SourceEngineVersion { get; set; } = null!;

        public ClusterS3ImportArgs()
        {
        }
        public static new ClusterS3ImportArgs Empty => new ClusterS3ImportArgs();
    }
}
