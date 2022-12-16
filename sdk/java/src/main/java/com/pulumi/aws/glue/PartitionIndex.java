// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.glue.PartitionIndexArgs;
import com.pulumi.aws.glue.inputs.PartitionIndexState;
import com.pulumi.aws.glue.outputs.PartitionIndexPartitionIndex;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.glue.CatalogDatabase;
 * import com.pulumi.aws.glue.CatalogDatabaseArgs;
 * import com.pulumi.aws.glue.CatalogTable;
 * import com.pulumi.aws.glue.CatalogTableArgs;
 * import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorArgs;
 * import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorSerDeInfoArgs;
 * import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorSkewedInfoArgs;
 * import com.pulumi.aws.glue.inputs.CatalogTablePartitionKeyArgs;
 * import com.pulumi.aws.glue.PartitionIndex;
 * import com.pulumi.aws.glue.PartitionIndexArgs;
 * import com.pulumi.aws.glue.inputs.PartitionIndexPartitionIndexArgs;
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
 *         var exampleCatalogDatabase = new CatalogDatabase(&#34;exampleCatalogDatabase&#34;, CatalogDatabaseArgs.builder()        
 *             .name(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleCatalogTable = new CatalogTable(&#34;exampleCatalogTable&#34;, CatalogTableArgs.builder()        
 *             .name(&#34;example&#34;)
 *             .databaseName(exampleCatalogDatabase.name())
 *             .owner(&#34;my_owner&#34;)
 *             .retention(1)
 *             .tableType(&#34;VIRTUAL_VIEW&#34;)
 *             .viewExpandedText(&#34;view_expanded_text_1&#34;)
 *             .viewOriginalText(&#34;view_original_text_1&#34;)
 *             .storageDescriptor(CatalogTableStorageDescriptorArgs.builder()
 *                 .bucketColumns(&#34;bucket_column_1&#34;)
 *                 .compressed(false)
 *                 .inputFormat(&#34;SequenceFileInputFormat&#34;)
 *                 .location(&#34;my_location&#34;)
 *                 .numberOfBuckets(1)
 *                 .outputFormat(&#34;SequenceFileInputFormat&#34;)
 *                 .storedAsSubDirectories(false)
 *                 .parameters(Map.of(&#34;param1&#34;, &#34;param1_val&#34;))
 *                 .columns(                
 *                     CatalogTableStorageDescriptorColumnArgs.builder()
 *                         .name(&#34;my_column_1&#34;)
 *                         .type(&#34;int&#34;)
 *                         .comment(&#34;my_column1_comment&#34;)
 *                         .build(),
 *                     CatalogTableStorageDescriptorColumnArgs.builder()
 *                         .name(&#34;my_column_2&#34;)
 *                         .type(&#34;string&#34;)
 *                         .comment(&#34;my_column2_comment&#34;)
 *                         .build())
 *                 .serDeInfo(CatalogTableStorageDescriptorSerDeInfoArgs.builder()
 *                     .name(&#34;ser_de_name&#34;)
 *                     .parameters(Map.of(&#34;param1&#34;, &#34;param_val_1&#34;))
 *                     .serializationLibrary(&#34;org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe&#34;)
 *                     .build())
 *                 .sortColumns(CatalogTableStorageDescriptorSortColumnArgs.builder()
 *                     .column(&#34;my_column_1&#34;)
 *                     .sortOrder(1)
 *                     .build())
 *                 .skewedInfo(CatalogTableStorageDescriptorSkewedInfoArgs.builder()
 *                     .skewedColumnNames(&#34;my_column_1&#34;)
 *                     .skewedColumnValueLocationMaps(Map.of(&#34;my_column_1&#34;, &#34;my_column_1_val_loc_map&#34;))
 *                     .skewedColumnValues(&#34;skewed_val_1&#34;)
 *                     .build())
 *                 .build())
 *             .partitionKeys(            
 *                 CatalogTablePartitionKeyArgs.builder()
 *                     .name(&#34;my_column_1&#34;)
 *                     .type(&#34;int&#34;)
 *                     .comment(&#34;my_column_1_comment&#34;)
 *                     .build(),
 *                 CatalogTablePartitionKeyArgs.builder()
 *                     .name(&#34;my_column_2&#34;)
 *                     .type(&#34;string&#34;)
 *                     .comment(&#34;my_column_2_comment&#34;)
 *                     .build())
 *             .parameters(Map.of(&#34;param1&#34;, &#34;param1_val&#34;))
 *             .build());
 * 
 *         var examplePartitionIndex = new PartitionIndex(&#34;examplePartitionIndex&#34;, PartitionIndexArgs.builder()        
 *             .databaseName(exampleCatalogDatabase.name())
 *             .tableName(exampleCatalogTable.name())
 *             .partitionIndex(PartitionIndexPartitionIndexArgs.builder()
 *                 .indexName(&#34;example&#34;)
 *                 .keys(                
 *                     &#34;my_column_1&#34;,
 *                     &#34;my_column_2&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Glue Partition Indexes can be imported with their catalog ID (usually AWS account ID), database name, table name, and index name, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:glue/partitionIndex:PartitionIndex example 123456789012:MyDatabase:MyTable:index-name
 * ```
 * 
 */
@ResourceType(type="aws:glue/partitionIndex:PartitionIndex")
public class PartitionIndex extends com.pulumi.resources.CustomResource {
    /**
     * The catalog ID where the table resides.
     * 
     */
    @Export(name="catalogId", refs={String.class}, tree="[0]")
    private Output<String> catalogId;

    /**
     * @return The catalog ID where the table resides.
     * 
     */
    public Output<String> catalogId() {
        return this.catalogId;
    }
    /**
     * Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
     * 
     */
    @Export(name="databaseName", refs={String.class}, tree="[0]")
    private Output<String> databaseName;

    /**
     * @return Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }
    /**
     * Configuration block for a partition index. See `partition_index` below.
     * 
     */
    @Export(name="partitionIndex", refs={PartitionIndexPartitionIndex.class}, tree="[0]")
    private Output<PartitionIndexPartitionIndex> partitionIndex;

    /**
     * @return Configuration block for a partition index. See `partition_index` below.
     * 
     */
    public Output<PartitionIndexPartitionIndex> partitionIndex() {
        return this.partitionIndex;
    }
    /**
     * Name of the table. For Hive compatibility, this must be entirely lowercase.
     * 
     */
    @Export(name="tableName", refs={String.class}, tree="[0]")
    private Output<String> tableName;

    /**
     * @return Name of the table. For Hive compatibility, this must be entirely lowercase.
     * 
     */
    public Output<String> tableName() {
        return this.tableName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PartitionIndex(String name) {
        this(name, PartitionIndexArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PartitionIndex(String name, PartitionIndexArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PartitionIndex(String name, PartitionIndexArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:glue/partitionIndex:PartitionIndex", name, args == null ? PartitionIndexArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PartitionIndex(String name, Output<String> id, @Nullable PartitionIndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:glue/partitionIndex:PartitionIndex", name, state, makeResourceOptions(options, id));
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
    public static PartitionIndex get(String name, Output<String> id, @Nullable PartitionIndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PartitionIndex(name, id, state, options);
    }
}
