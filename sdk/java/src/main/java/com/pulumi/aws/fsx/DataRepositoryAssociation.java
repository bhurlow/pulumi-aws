// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.fsx.DataRepositoryAssociationArgs;
import com.pulumi.aws.fsx.inputs.DataRepositoryAssociationState;
import com.pulumi.aws.fsx.outputs.DataRepositoryAssociationS3;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a FSx for Lustre Data Repository Association. See [Linking your file system to an S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html) for more information.
 * 
 * &gt; **NOTE:** Data Repository Associations are only compatible with AWS FSx for Lustre File Systems and `PERSISTENT_2` deployment type.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.BucketAclV2;
 * import com.pulumi.aws.s3.BucketAclV2Args;
 * import com.pulumi.aws.fsx.LustreFileSystem;
 * import com.pulumi.aws.fsx.LustreFileSystemArgs;
 * import com.pulumi.aws.fsx.DataRepositoryAssociation;
 * import com.pulumi.aws.fsx.DataRepositoryAssociationArgs;
 * import com.pulumi.aws.fsx.inputs.DataRepositoryAssociationS3Args;
 * import com.pulumi.aws.fsx.inputs.DataRepositoryAssociationS3AutoExportPolicyArgs;
 * import com.pulumi.aws.fsx.inputs.DataRepositoryAssociationS3AutoImportPolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleBucketV2 = new BucketV2(&#34;exampleBucketV2&#34;);
 * 
 *         var exampleBucketAclV2 = new BucketAclV2(&#34;exampleBucketAclV2&#34;, BucketAclV2Args.builder()        
 *             .bucket(exampleBucketV2.id())
 *             .acl(&#34;private&#34;)
 *             .build());
 * 
 *         var exampleLustreFileSystem = new LustreFileSystem(&#34;exampleLustreFileSystem&#34;, LustreFileSystemArgs.builder()        
 *             .storageCapacity(1200)
 *             .subnetIds(aws_subnet.example().id())
 *             .deploymentType(&#34;PERSISTENT_2&#34;)
 *             .perUnitStorageThroughput(125)
 *             .build());
 * 
 *         var exampleDataRepositoryAssociation = new DataRepositoryAssociation(&#34;exampleDataRepositoryAssociation&#34;, DataRepositoryAssociationArgs.builder()        
 *             .fileSystemId(exampleLustreFileSystem.id())
 *             .dataRepositoryPath(exampleBucketV2.id().applyValue(id -&gt; String.format(&#34;s3://%s&#34;, id)))
 *             .fileSystemPath(&#34;/my-bucket&#34;)
 *             .s3(DataRepositoryAssociationS3Args.builder()
 *                 .autoExportPolicy(DataRepositoryAssociationS3AutoExportPolicyArgs.builder()
 *                     .events(                    
 *                         &#34;NEW&#34;,
 *                         &#34;CHANGED&#34;,
 *                         &#34;DELETED&#34;)
 *                     .build())
 *                 .autoImportPolicy(DataRepositoryAssociationS3AutoImportPolicyArgs.builder()
 *                     .events(                    
 *                         &#34;NEW&#34;,
 *                         &#34;CHANGED&#34;,
 *                         &#34;DELETED&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_fsx_data_repository_association.example
 * 
 *  id = &#34;dra-0b1cfaeca11088b10&#34; } Using `pulumi import`, import FSx Data Repository Associations using the `id`. For exampleconsole % pulumi import aws_fsx_data_repository_association.example dra-0b1cfaeca11088b10
 * 
 */
@ResourceType(type="aws:fsx/dataRepositoryAssociation:DataRepositoryAssociation")
public class DataRepositoryAssociation extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name of the file system.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name of the file system.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    @Export(name="associationId", refs={String.class}, tree="[0]")
    private Output<String> associationId;

    public Output<String> associationId() {
        return this.associationId;
    }
    /**
     * Set to true to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created. Defaults to `false`.
     * 
     */
    @Export(name="batchImportMetaDataOnCreate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> batchImportMetaDataOnCreate;

    /**
     * @return Set to true to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> batchImportMetaDataOnCreate() {
        return Codegen.optional(this.batchImportMetaDataOnCreate);
    }
    /**
     * The path to the Amazon S3 data repository that will be linked to the file system. The path must be an S3 bucket s3://myBucket/myPrefix/. This path specifies where in the S3 data repository files will be imported from or exported to. The same S3 bucket cannot be linked more than once to the same file system.
     * 
     */
    @Export(name="dataRepositoryPath", refs={String.class}, tree="[0]")
    private Output<String> dataRepositoryPath;

    /**
     * @return The path to the Amazon S3 data repository that will be linked to the file system. The path must be an S3 bucket s3://myBucket/myPrefix/. This path specifies where in the S3 data repository files will be imported from or exported to. The same S3 bucket cannot be linked more than once to the same file system.
     * 
     */
    public Output<String> dataRepositoryPath() {
        return this.dataRepositoryPath;
    }
    /**
     * Set to true to delete files from the file system upon deleting this data repository association. Defaults to `false`.
     * 
     */
    @Export(name="deleteDataInFilesystem", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteDataInFilesystem;

    /**
     * @return Set to true to delete files from the file system upon deleting this data repository association. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> deleteDataInFilesystem() {
        return Codegen.optional(this.deleteDataInFilesystem);
    }
    /**
     * The ID of the Amazon FSx file system to on which to create a data repository association.
     * 
     */
    @Export(name="fileSystemId", refs={String.class}, tree="[0]")
    private Output<String> fileSystemId;

    /**
     * @return The ID of the Amazon FSx file system to on which to create a data repository association.
     * 
     */
    public Output<String> fileSystemId() {
        return this.fileSystemId;
    }
    /**
     * A path on the file system that points to a high-level directory (such as `/ns1/`) or subdirectory (such as `/ns1/subdir/`) that will be mapped 1-1 with `data_repository_path`. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path `/ns1/`, then you cannot link another data repository with file system path `/ns1/ns2`. This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
     * 
     */
    @Export(name="fileSystemPath", refs={String.class}, tree="[0]")
    private Output<String> fileSystemPath;

    /**
     * @return A path on the file system that points to a high-level directory (such as `/ns1/`) or subdirectory (such as `/ns1/subdir/`) that will be mapped 1-1 with `data_repository_path`. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path `/ns1/`, then you cannot link another data repository with file system path `/ns1/ns2`. This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
     * 
     */
    public Output<String> fileSystemPath() {
        return this.fileSystemPath;
    }
    /**
     * For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
     * 
     */
    @Export(name="importedFileChunkSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> importedFileChunkSize;

    /**
     * @return For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
     * 
     */
    public Output<Integer> importedFileChunkSize() {
        return this.importedFileChunkSize;
    }
    /**
     * See the `s3` configuration block. Max of 1.
     * The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
     * 
     */
    @Export(name="s3", refs={DataRepositoryAssociationS3.class}, tree="[0]")
    private Output<DataRepositoryAssociationS3> s3;

    /**
     * @return See the `s3` configuration block. Max of 1.
     * The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
     * 
     */
    public Output<DataRepositoryAssociationS3> s3() {
        return this.s3;
    }
    /**
     * A map of tags to assign to the data repository association. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the data repository association. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataRepositoryAssociation(String name) {
        this(name, DataRepositoryAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataRepositoryAssociation(String name, DataRepositoryAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataRepositoryAssociation(String name, DataRepositoryAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:fsx/dataRepositoryAssociation:DataRepositoryAssociation", name, args == null ? DataRepositoryAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataRepositoryAssociation(String name, Output<String> id, @Nullable DataRepositoryAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:fsx/dataRepositoryAssociation:DataRepositoryAssociation", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DataRepositoryAssociation get(String name, Output<String> id, @Nullable DataRepositoryAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataRepositoryAssociation(name, id, state, options);
    }
}
