// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.quicksight.DataSetArgs;
import com.pulumi.aws.quicksight.inputs.DataSetState;
import com.pulumi.aws.quicksight.outputs.DataSetColumnGroup;
import com.pulumi.aws.quicksight.outputs.DataSetColumnLevelPermissionRule;
import com.pulumi.aws.quicksight.outputs.DataSetDataSetUsageConfiguration;
import com.pulumi.aws.quicksight.outputs.DataSetFieldFolder;
import com.pulumi.aws.quicksight.outputs.DataSetLogicalTableMap;
import com.pulumi.aws.quicksight.outputs.DataSetOutputColumn;
import com.pulumi.aws.quicksight.outputs.DataSetPermission;
import com.pulumi.aws.quicksight.outputs.DataSetPhysicalTableMap;
import com.pulumi.aws.quicksight.outputs.DataSetRefreshProperties;
import com.pulumi.aws.quicksight.outputs.DataSetRowLevelPermissionDataSet;
import com.pulumi.aws.quicksight.outputs.DataSetRowLevelPermissionTagConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing a QuickSight Data Set.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.DataSet;
 * import com.pulumi.aws.quicksight.DataSetArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
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
 *         var example = new DataSet(&#34;example&#34;, DataSetArgs.builder()        
 *             .dataSetId(&#34;example-id&#34;)
 *             .importMode(&#34;SPICE&#34;)
 *             .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
 *                 .physicalTableMapId(&#34;example-id&#34;)
 *                 .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
 *                     .dataSourceArn(aws_quicksight_data_source.example().arn())
 *                     .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
 *                         .name(&#34;Column1&#34;)
 *                         .type(&#34;STRING&#34;)
 *                         .build())
 *                     .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
 *                         .format(&#34;JSON&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Column Level Permission Rules
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.DataSet;
 * import com.pulumi.aws.quicksight.DataSetArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetColumnLevelPermissionRuleArgs;
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
 *         var example = new DataSet(&#34;example&#34;, DataSetArgs.builder()        
 *             .dataSetId(&#34;example-id&#34;)
 *             .importMode(&#34;SPICE&#34;)
 *             .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
 *                 .physicalTableMapId(&#34;example-id&#34;)
 *                 .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
 *                     .dataSourceArn(aws_quicksight_data_source.example().arn())
 *                     .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
 *                         .name(&#34;Column1&#34;)
 *                         .type(&#34;STRING&#34;)
 *                         .build())
 *                     .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
 *                         .format(&#34;JSON&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .columnLevelPermissionRules(DataSetColumnLevelPermissionRuleArgs.builder()
 *                 .columnNames(&#34;Column1&#34;)
 *                 .principals(aws_quicksight_user.example().arn())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Field Folders
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.DataSet;
 * import com.pulumi.aws.quicksight.DataSetArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetFieldFolderArgs;
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
 *         var example = new DataSet(&#34;example&#34;, DataSetArgs.builder()        
 *             .dataSetId(&#34;example-id&#34;)
 *             .importMode(&#34;SPICE&#34;)
 *             .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
 *                 .physicalTableMapId(&#34;example-id&#34;)
 *                 .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
 *                     .dataSourceArn(aws_quicksight_data_source.example().arn())
 *                     .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
 *                         .name(&#34;Column1&#34;)
 *                         .type(&#34;STRING&#34;)
 *                         .build())
 *                     .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
 *                         .format(&#34;JSON&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .fieldFolders(DataSetFieldFolderArgs.builder()
 *                 .fieldFoldersId(&#34;example-id&#34;)
 *                 .columns(&#34;Column1&#34;)
 *                 .description(&#34;example description&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Permissions
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.DataSet;
 * import com.pulumi.aws.quicksight.DataSetArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPermissionArgs;
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
 *         var example = new DataSet(&#34;example&#34;, DataSetArgs.builder()        
 *             .dataSetId(&#34;example-id&#34;)
 *             .importMode(&#34;SPICE&#34;)
 *             .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
 *                 .physicalTableMapId(&#34;example-id&#34;)
 *                 .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
 *                     .dataSourceArn(aws_quicksight_data_source.example().arn())
 *                     .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
 *                         .name(&#34;Column1&#34;)
 *                         .type(&#34;STRING&#34;)
 *                         .build())
 *                     .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
 *                         .format(&#34;JSON&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .permissions(DataSetPermissionArgs.builder()
 *                 .actions(                
 *                     &#34;quicksight:DescribeDataSet&#34;,
 *                     &#34;quicksight:DescribeDataSetPermissions&#34;,
 *                     &#34;quicksight:PassDataSet&#34;,
 *                     &#34;quicksight:DescribeIngestion&#34;,
 *                     &#34;quicksight:ListIngestions&#34;)
 *                 .principal(aws_quicksight_user.example().arn())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Row Level Permission Tag Configuration
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.DataSet;
 * import com.pulumi.aws.quicksight.DataSetArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSetRowLevelPermissionTagConfigurationArgs;
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
 *         var example = new DataSet(&#34;example&#34;, DataSetArgs.builder()        
 *             .dataSetId(&#34;example-id&#34;)
 *             .importMode(&#34;SPICE&#34;)
 *             .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
 *                 .physicalTableMapId(&#34;example-id&#34;)
 *                 .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
 *                     .dataSourceArn(aws_quicksight_data_source.example().arn())
 *                     .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
 *                         .name(&#34;Column1&#34;)
 *                         .type(&#34;STRING&#34;)
 *                         .build())
 *                     .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
 *                         .format(&#34;JSON&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .rowLevelPermissionTagConfiguration(DataSetRowLevelPermissionTagConfigurationArgs.builder()
 *                 .status(&#34;ENABLED&#34;)
 *                 .tagRules(DataSetRowLevelPermissionTagConfigurationTagRuleArgs.builder()
 *                     .columnName(&#34;Column1&#34;)
 *                     .tagKey(&#34;tagkey&#34;)
 *                     .matchAllValue(&#34;*&#34;)
 *                     .tagMultiValueDelimiter(&#34;,&#34;)
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
 * A QuickSight Data Set can be imported using the AWS account ID and data set ID separated by a comma (`,`) e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:quicksight/dataSet:DataSet example 123456789012,example-id
 * ```
 * 
 */
@ResourceType(type="aws:quicksight/dataSet:DataSet")
public class DataSet extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the dataset that contains permissions for RLS.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the dataset that contains permissions for RLS.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * AWS account ID.
     * 
     */
    @Export(name="awsAccountId", refs={String.class}, tree="[0]")
    private Output<String> awsAccountId;

    /**
     * @return AWS account ID.
     * 
     */
    public Output<String> awsAccountId() {
        return this.awsAccountId;
    }
    /**
     * Groupings of columns that work together in certain Amazon QuickSight features. Currently, only geospatial hierarchy is supported. See column_groups.
     * 
     */
    @Export(name="columnGroups", refs={List.class,DataSetColumnGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataSetColumnGroup>> columnGroups;

    /**
     * @return Groupings of columns that work together in certain Amazon QuickSight features. Currently, only geospatial hierarchy is supported. See column_groups.
     * 
     */
    public Output<Optional<List<DataSetColumnGroup>>> columnGroups() {
        return Codegen.optional(this.columnGroups);
    }
    /**
     * A set of 1 or more definitions of a [ColumnLevelPermissionRule](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnLevelPermissionRule.html). See column_level_permission_rules.
     * 
     */
    @Export(name="columnLevelPermissionRules", refs={List.class,DataSetColumnLevelPermissionRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataSetColumnLevelPermissionRule>> columnLevelPermissionRules;

    /**
     * @return A set of 1 or more definitions of a [ColumnLevelPermissionRule](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnLevelPermissionRule.html). See column_level_permission_rules.
     * 
     */
    public Output<Optional<List<DataSetColumnLevelPermissionRule>>> columnLevelPermissionRules() {
        return Codegen.optional(this.columnLevelPermissionRules);
    }
    /**
     * Identifier for the data set.
     * 
     */
    @Export(name="dataSetId", refs={String.class}, tree="[0]")
    private Output<String> dataSetId;

    /**
     * @return Identifier for the data set.
     * 
     */
    public Output<String> dataSetId() {
        return this.dataSetId;
    }
    /**
     * The usage configuration to apply to child datasets that reference this dataset as a source. See data_set_usage_configuration.
     * 
     */
    @Export(name="dataSetUsageConfiguration", refs={DataSetDataSetUsageConfiguration.class}, tree="[0]")
    private Output<DataSetDataSetUsageConfiguration> dataSetUsageConfiguration;

    /**
     * @return The usage configuration to apply to child datasets that reference this dataset as a source. See data_set_usage_configuration.
     * 
     */
    public Output<DataSetDataSetUsageConfiguration> dataSetUsageConfiguration() {
        return this.dataSetUsageConfiguration;
    }
    /**
     * The folder that contains fields and nested subfolders for your dataset. See field_folders.
     * 
     */
    @Export(name="fieldFolders", refs={List.class,DataSetFieldFolder.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataSetFieldFolder>> fieldFolders;

    /**
     * @return The folder that contains fields and nested subfolders for your dataset. See field_folders.
     * 
     */
    public Output<Optional<List<DataSetFieldFolder>>> fieldFolders() {
        return Codegen.optional(this.fieldFolders);
    }
    /**
     * Indicates whether you want to import the data into SPICE. Valid values are `SPICE` and `DIRECT_QUERY`.
     * 
     */
    @Export(name="importMode", refs={String.class}, tree="[0]")
    private Output<String> importMode;

    /**
     * @return Indicates whether you want to import the data into SPICE. Valid values are `SPICE` and `DIRECT_QUERY`.
     * 
     */
    public Output<String> importMode() {
        return this.importMode;
    }
    /**
     * Configures the combination and transformation of the data from the physical tables. Maximum of 1 entry. See logical_table_map.
     * 
     */
    @Export(name="logicalTableMaps", refs={List.class,DataSetLogicalTableMap.class}, tree="[0,1]")
    private Output<List<DataSetLogicalTableMap>> logicalTableMaps;

    /**
     * @return Configures the combination and transformation of the data from the physical tables. Maximum of 1 entry. See logical_table_map.
     * 
     */
    public Output<List<DataSetLogicalTableMap>> logicalTableMaps() {
        return this.logicalTableMaps;
    }
    /**
     * Display name for the dataset.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name for the dataset.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="outputColumns", refs={List.class,DataSetOutputColumn.class}, tree="[0,1]")
    private Output<List<DataSetOutputColumn>> outputColumns;

    public Output<List<DataSetOutputColumn>> outputColumns() {
        return this.outputColumns;
    }
    /**
     * A set of resource permissions on the data source. Maximum of 64 items. See permissions.
     * 
     */
    @Export(name="permissions", refs={List.class,DataSetPermission.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataSetPermission>> permissions;

    /**
     * @return A set of resource permissions on the data source. Maximum of 64 items. See permissions.
     * 
     */
    public Output<Optional<List<DataSetPermission>>> permissions() {
        return Codegen.optional(this.permissions);
    }
    /**
     * Declares the physical tables that are available in the underlying data sources. See physical_table_map.
     * 
     */
    @Export(name="physicalTableMaps", refs={List.class,DataSetPhysicalTableMap.class}, tree="[0,1]")
    private Output<List<DataSetPhysicalTableMap>> physicalTableMaps;

    /**
     * @return Declares the physical tables that are available in the underlying data sources. See physical_table_map.
     * 
     */
    public Output<List<DataSetPhysicalTableMap>> physicalTableMaps() {
        return this.physicalTableMaps;
    }
    /**
     * The refresh properties for the data set. **NOTE**: Only valid when `import_mode` is set to `SPICE`. See refresh_properties.
     * 
     */
    @Export(name="refreshProperties", refs={DataSetRefreshProperties.class}, tree="[0]")
    private Output</* @Nullable */ DataSetRefreshProperties> refreshProperties;

    /**
     * @return The refresh properties for the data set. **NOTE**: Only valid when `import_mode` is set to `SPICE`. See refresh_properties.
     * 
     */
    public Output<Optional<DataSetRefreshProperties>> refreshProperties() {
        return Codegen.optional(this.refreshProperties);
    }
    /**
     * The row-level security configuration for the data that you want to create. See row_level_permission_data_set.
     * 
     */
    @Export(name="rowLevelPermissionDataSet", refs={DataSetRowLevelPermissionDataSet.class}, tree="[0]")
    private Output</* @Nullable */ DataSetRowLevelPermissionDataSet> rowLevelPermissionDataSet;

    /**
     * @return The row-level security configuration for the data that you want to create. See row_level_permission_data_set.
     * 
     */
    public Output<Optional<DataSetRowLevelPermissionDataSet>> rowLevelPermissionDataSet() {
        return Codegen.optional(this.rowLevelPermissionDataSet);
    }
    /**
     * The configuration of tags on a dataset to set row-level security. Row-level security tags are currently supported for anonymous embedding only. See row_level_permission_tag_configuration.
     * 
     */
    @Export(name="rowLevelPermissionTagConfiguration", refs={DataSetRowLevelPermissionTagConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ DataSetRowLevelPermissionTagConfiguration> rowLevelPermissionTagConfiguration;

    /**
     * @return The configuration of tags on a dataset to set row-level security. Row-level security tags are currently supported for anonymous embedding only. See row_level_permission_tag_configuration.
     * 
     */
    public Output<Optional<DataSetRowLevelPermissionTagConfiguration>> rowLevelPermissionTagConfiguration() {
        return Codegen.optional(this.rowLevelPermissionTagConfiguration);
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public DataSet(String name) {
        this(name, DataSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataSet(String name, DataSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataSet(String name, DataSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/dataSet:DataSet", name, args == null ? DataSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataSet(String name, Output<String> id, @Nullable DataSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/dataSet:DataSet", name, state, makeResourceOptions(options, id));
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
    public static DataSet get(String name, Output<String> id, @Nullable DataSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataSet(name, id, state, options);
    }
}
