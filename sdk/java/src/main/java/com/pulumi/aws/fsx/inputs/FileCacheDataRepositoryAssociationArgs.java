// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx.inputs;

import com.pulumi.aws.fsx.inputs.FileCacheDataRepositoryAssociationNfArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FileCacheDataRepositoryAssociationArgs extends com.pulumi.resources.ResourceArgs {

    public static final FileCacheDataRepositoryAssociationArgs Empty = new FileCacheDataRepositoryAssociationArgs();

    @Import(name="associationId")
    private @Nullable Output<String> associationId;

    public Optional<Output<String>> associationId() {
        return Optional.ofNullable(this.associationId);
    }

    /**
     * The path to the S3 or NFS data repository that links to the cache.
     * 
     */
    @Import(name="dataRepositoryPath", required=true)
    private Output<String> dataRepositoryPath;

    /**
     * @return The path to the S3 or NFS data repository that links to the cache.
     * 
     */
    public Output<String> dataRepositoryPath() {
        return this.dataRepositoryPath;
    }

    /**
     * A list of NFS Exports that will be linked with this data repository association. The Export paths are in the format /exportpath1. To use this parameter, you must configure DataRepositoryPath as the domain name of the NFS file system. The NFS file system domain name in effect is the root of the subdirectories. Note that DataRepositorySubdirectories is not supported for S3 data repositories. Max of 500.
     * 
     */
    @Import(name="dataRepositorySubdirectories")
    private @Nullable Output<List<String>> dataRepositorySubdirectories;

    /**
     * @return A list of NFS Exports that will be linked with this data repository association. The Export paths are in the format /exportpath1. To use this parameter, you must configure DataRepositoryPath as the domain name of the NFS file system. The NFS file system domain name in effect is the root of the subdirectories. Note that DataRepositorySubdirectories is not supported for S3 data repositories. Max of 500.
     * 
     */
    public Optional<Output<List<String>>> dataRepositorySubdirectories() {
        return Optional.ofNullable(this.dataRepositorySubdirectories);
    }

    /**
     * The system-generated, unique ID of the cache.
     * 
     */
    @Import(name="fileCacheId")
    private @Nullable Output<String> fileCacheId;

    /**
     * @return The system-generated, unique ID of the cache.
     * 
     */
    public Optional<Output<String>> fileCacheId() {
        return Optional.ofNullable(this.fileCacheId);
    }

    /**
     * A path on the cache that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath. The leading forward slash in the name is required. Two data repository associations cannot have overlapping cache paths. For example, if a data repository is associated with cache path /ns1/, then you cannot link another data repository with cache path /ns1/ns2. This path specifies where in your cache files will be exported from. This cache directory can be linked to only one data repository, and no data repository other can be linked to the directory. Note: The cache path can only be set to root (/) on an NFS DRA when DataRepositorySubdirectories is specified. If you specify root (/) as the cache path, you can create only one DRA on the cache. The cache path cannot be set to root (/) for an S3 DRA.
     * 
     */
    @Import(name="fileCachePath", required=true)
    private Output<String> fileCachePath;

    /**
     * @return A path on the cache that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath. The leading forward slash in the name is required. Two data repository associations cannot have overlapping cache paths. For example, if a data repository is associated with cache path /ns1/, then you cannot link another data repository with cache path /ns1/ns2. This path specifies where in your cache files will be exported from. This cache directory can be linked to only one data repository, and no data repository other can be linked to the directory. Note: The cache path can only be set to root (/) on an NFS DRA when DataRepositorySubdirectories is specified. If you specify root (/) as the cache path, you can create only one DRA on the cache. The cache path cannot be set to root (/) for an S3 DRA.
     * 
     */
    public Output<String> fileCachePath() {
        return this.fileCachePath;
    }

    @Import(name="fileSystemId")
    private @Nullable Output<String> fileSystemId;

    public Optional<Output<String>> fileSystemId() {
        return Optional.ofNullable(this.fileSystemId);
    }

    @Import(name="fileSystemPath")
    private @Nullable Output<String> fileSystemPath;

    public Optional<Output<String>> fileSystemPath() {
        return Optional.ofNullable(this.fileSystemPath);
    }

    @Import(name="importedFileChunkSize")
    private @Nullable Output<Integer> importedFileChunkSize;

    public Optional<Output<Integer>> importedFileChunkSize() {
        return Optional.ofNullable(this.importedFileChunkSize);
    }

    /**
     * - (Optional) See the `nfs` configuration block.
     * 
     */
    @Import(name="nfs")
    private @Nullable Output<List<FileCacheDataRepositoryAssociationNfArgs>> nfs;

    /**
     * @return - (Optional) See the `nfs` configuration block.
     * 
     */
    public Optional<Output<List<FileCacheDataRepositoryAssociationNfArgs>>> nfs() {
        return Optional.ofNullable(this.nfs);
    }

    @Import(name="resourceArn")
    private @Nullable Output<String> resourceArn;

    public Optional<Output<String>> resourceArn() {
        return Optional.ofNullable(this.resourceArn);
    }

    /**
     * A map of tags to assign to the file cache. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the file cache. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private FileCacheDataRepositoryAssociationArgs() {}

    private FileCacheDataRepositoryAssociationArgs(FileCacheDataRepositoryAssociationArgs $) {
        this.associationId = $.associationId;
        this.dataRepositoryPath = $.dataRepositoryPath;
        this.dataRepositorySubdirectories = $.dataRepositorySubdirectories;
        this.fileCacheId = $.fileCacheId;
        this.fileCachePath = $.fileCachePath;
        this.fileSystemId = $.fileSystemId;
        this.fileSystemPath = $.fileSystemPath;
        this.importedFileChunkSize = $.importedFileChunkSize;
        this.nfs = $.nfs;
        this.resourceArn = $.resourceArn;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FileCacheDataRepositoryAssociationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FileCacheDataRepositoryAssociationArgs $;

        public Builder() {
            $ = new FileCacheDataRepositoryAssociationArgs();
        }

        public Builder(FileCacheDataRepositoryAssociationArgs defaults) {
            $ = new FileCacheDataRepositoryAssociationArgs(Objects.requireNonNull(defaults));
        }

        public Builder associationId(@Nullable Output<String> associationId) {
            $.associationId = associationId;
            return this;
        }

        public Builder associationId(String associationId) {
            return associationId(Output.of(associationId));
        }

        /**
         * @param dataRepositoryPath The path to the S3 or NFS data repository that links to the cache.
         * 
         * @return builder
         * 
         */
        public Builder dataRepositoryPath(Output<String> dataRepositoryPath) {
            $.dataRepositoryPath = dataRepositoryPath;
            return this;
        }

        /**
         * @param dataRepositoryPath The path to the S3 or NFS data repository that links to the cache.
         * 
         * @return builder
         * 
         */
        public Builder dataRepositoryPath(String dataRepositoryPath) {
            return dataRepositoryPath(Output.of(dataRepositoryPath));
        }

        /**
         * @param dataRepositorySubdirectories A list of NFS Exports that will be linked with this data repository association. The Export paths are in the format /exportpath1. To use this parameter, you must configure DataRepositoryPath as the domain name of the NFS file system. The NFS file system domain name in effect is the root of the subdirectories. Note that DataRepositorySubdirectories is not supported for S3 data repositories. Max of 500.
         * 
         * @return builder
         * 
         */
        public Builder dataRepositorySubdirectories(@Nullable Output<List<String>> dataRepositorySubdirectories) {
            $.dataRepositorySubdirectories = dataRepositorySubdirectories;
            return this;
        }

        /**
         * @param dataRepositorySubdirectories A list of NFS Exports that will be linked with this data repository association. The Export paths are in the format /exportpath1. To use this parameter, you must configure DataRepositoryPath as the domain name of the NFS file system. The NFS file system domain name in effect is the root of the subdirectories. Note that DataRepositorySubdirectories is not supported for S3 data repositories. Max of 500.
         * 
         * @return builder
         * 
         */
        public Builder dataRepositorySubdirectories(List<String> dataRepositorySubdirectories) {
            return dataRepositorySubdirectories(Output.of(dataRepositorySubdirectories));
        }

        /**
         * @param dataRepositorySubdirectories A list of NFS Exports that will be linked with this data repository association. The Export paths are in the format /exportpath1. To use this parameter, you must configure DataRepositoryPath as the domain name of the NFS file system. The NFS file system domain name in effect is the root of the subdirectories. Note that DataRepositorySubdirectories is not supported for S3 data repositories. Max of 500.
         * 
         * @return builder
         * 
         */
        public Builder dataRepositorySubdirectories(String... dataRepositorySubdirectories) {
            return dataRepositorySubdirectories(List.of(dataRepositorySubdirectories));
        }

        /**
         * @param fileCacheId The system-generated, unique ID of the cache.
         * 
         * @return builder
         * 
         */
        public Builder fileCacheId(@Nullable Output<String> fileCacheId) {
            $.fileCacheId = fileCacheId;
            return this;
        }

        /**
         * @param fileCacheId The system-generated, unique ID of the cache.
         * 
         * @return builder
         * 
         */
        public Builder fileCacheId(String fileCacheId) {
            return fileCacheId(Output.of(fileCacheId));
        }

        /**
         * @param fileCachePath A path on the cache that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath. The leading forward slash in the name is required. Two data repository associations cannot have overlapping cache paths. For example, if a data repository is associated with cache path /ns1/, then you cannot link another data repository with cache path /ns1/ns2. This path specifies where in your cache files will be exported from. This cache directory can be linked to only one data repository, and no data repository other can be linked to the directory. Note: The cache path can only be set to root (/) on an NFS DRA when DataRepositorySubdirectories is specified. If you specify root (/) as the cache path, you can create only one DRA on the cache. The cache path cannot be set to root (/) for an S3 DRA.
         * 
         * @return builder
         * 
         */
        public Builder fileCachePath(Output<String> fileCachePath) {
            $.fileCachePath = fileCachePath;
            return this;
        }

        /**
         * @param fileCachePath A path on the cache that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath. The leading forward slash in the name is required. Two data repository associations cannot have overlapping cache paths. For example, if a data repository is associated with cache path /ns1/, then you cannot link another data repository with cache path /ns1/ns2. This path specifies where in your cache files will be exported from. This cache directory can be linked to only one data repository, and no data repository other can be linked to the directory. Note: The cache path can only be set to root (/) on an NFS DRA when DataRepositorySubdirectories is specified. If you specify root (/) as the cache path, you can create only one DRA on the cache. The cache path cannot be set to root (/) for an S3 DRA.
         * 
         * @return builder
         * 
         */
        public Builder fileCachePath(String fileCachePath) {
            return fileCachePath(Output.of(fileCachePath));
        }

        public Builder fileSystemId(@Nullable Output<String> fileSystemId) {
            $.fileSystemId = fileSystemId;
            return this;
        }

        public Builder fileSystemId(String fileSystemId) {
            return fileSystemId(Output.of(fileSystemId));
        }

        public Builder fileSystemPath(@Nullable Output<String> fileSystemPath) {
            $.fileSystemPath = fileSystemPath;
            return this;
        }

        public Builder fileSystemPath(String fileSystemPath) {
            return fileSystemPath(Output.of(fileSystemPath));
        }

        public Builder importedFileChunkSize(@Nullable Output<Integer> importedFileChunkSize) {
            $.importedFileChunkSize = importedFileChunkSize;
            return this;
        }

        public Builder importedFileChunkSize(Integer importedFileChunkSize) {
            return importedFileChunkSize(Output.of(importedFileChunkSize));
        }

        /**
         * @param nfs - (Optional) See the `nfs` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder nfs(@Nullable Output<List<FileCacheDataRepositoryAssociationNfArgs>> nfs) {
            $.nfs = nfs;
            return this;
        }

        /**
         * @param nfs - (Optional) See the `nfs` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder nfs(List<FileCacheDataRepositoryAssociationNfArgs> nfs) {
            return nfs(Output.of(nfs));
        }

        /**
         * @param nfs - (Optional) See the `nfs` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder nfs(FileCacheDataRepositoryAssociationNfArgs... nfs) {
            return nfs(List.of(nfs));
        }

        public Builder resourceArn(@Nullable Output<String> resourceArn) {
            $.resourceArn = resourceArn;
            return this;
        }

        public Builder resourceArn(String resourceArn) {
            return resourceArn(Output.of(resourceArn));
        }

        /**
         * @param tags A map of tags to assign to the file cache. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the file cache. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public FileCacheDataRepositoryAssociationArgs build() {
            $.dataRepositoryPath = Objects.requireNonNull($.dataRepositoryPath, "expected parameter 'dataRepositoryPath' to be non-null");
            $.fileCachePath = Objects.requireNonNull($.fileCachePath, "expected parameter 'fileCachePath' to be non-null");
            return $;
        }
    }

}
