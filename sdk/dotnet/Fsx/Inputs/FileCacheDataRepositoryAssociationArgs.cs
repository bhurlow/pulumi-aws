// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fsx.Inputs
{

    public sealed class FileCacheDataRepositoryAssociationArgs : global::Pulumi.ResourceArgs
    {
        [Input("associationId")]
        public Input<string>? AssociationId { get; set; }

        /// <summary>
        /// The path to the S3 or NFS data repository that links to the cache.
        /// </summary>
        [Input("dataRepositoryPath", required: true)]
        public Input<string> DataRepositoryPath { get; set; } = null!;

        [Input("dataRepositorySubdirectories")]
        private InputList<string>? _dataRepositorySubdirectories;

        /// <summary>
        /// A list of NFS Exports that will be linked with this data repository association. The Export paths are in the format /exportpath1. To use this parameter, you must configure DataRepositoryPath as the domain name of the NFS file system. The NFS file system domain name in effect is the root of the subdirectories. Note that DataRepositorySubdirectories is not supported for S3 data repositories. Max of 500.
        /// </summary>
        public InputList<string> DataRepositorySubdirectories
        {
            get => _dataRepositorySubdirectories ?? (_dataRepositorySubdirectories = new InputList<string>());
            set => _dataRepositorySubdirectories = value;
        }

        /// <summary>
        /// The system-generated, unique ID of the cache.
        /// </summary>
        [Input("fileCacheId")]
        public Input<string>? FileCacheId { get; set; }

        /// <summary>
        /// A path on the cache that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath. The leading forward slash in the name is required. Two data repository associations cannot have overlapping cache paths. For example, if a data repository is associated with cache path /ns1/, then you cannot link another data repository with cache path /ns1/ns2. This path specifies where in your cache files will be exported from. This cache directory can be linked to only one data repository, and no data repository other can be linked to the directory. Note: The cache path can only be set to root (/) on an NFS DRA when DataRepositorySubdirectories is specified. If you specify root (/) as the cache path, you can create only one DRA on the cache. The cache path cannot be set to root (/) for an S3 DRA.
        /// </summary>
        [Input("fileCachePath", required: true)]
        public Input<string> FileCachePath { get; set; } = null!;

        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        [Input("fileSystemPath")]
        public Input<string>? FileSystemPath { get; set; }

        [Input("importedFileChunkSize")]
        public Input<int>? ImportedFileChunkSize { get; set; }

        [Input("nfs")]
        private InputList<Inputs.FileCacheDataRepositoryAssociationNfArgs>? _nfs;

        /// <summary>
        /// - (Optional) See the `nfs` configuration block.
        /// </summary>
        public InputList<Inputs.FileCacheDataRepositoryAssociationNfArgs> Nfs
        {
            get => _nfs ?? (_nfs = new InputList<Inputs.FileCacheDataRepositoryAssociationNfArgs>());
            set => _nfs = value;
        }

        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the file cache. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FileCacheDataRepositoryAssociationArgs()
        {
        }
        public static new FileCacheDataRepositoryAssociationArgs Empty => new FileCacheDataRepositoryAssociationArgs();
    }
}
