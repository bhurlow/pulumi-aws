// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.quicksight.DataSourceArgs;
import com.pulumi.aws.quicksight.inputs.DataSourceState;
import com.pulumi.aws.quicksight.outputs.DataSourceCredentials;
import com.pulumi.aws.quicksight.outputs.DataSourceParameters;
import com.pulumi.aws.quicksight.outputs.DataSourcePermission;
import com.pulumi.aws.quicksight.outputs.DataSourceSslProperties;
import com.pulumi.aws.quicksight.outputs.DataSourceVpcConnectionProperties;
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
 * Resource for managing QuickSight Data Source
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.DataSource;
 * import com.pulumi.aws.quicksight.DataSourceArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSourceParametersArgs;
 * import com.pulumi.aws.quicksight.inputs.DataSourceParametersS3Args;
 * import com.pulumi.aws.quicksight.inputs.DataSourceParametersS3ManifestFileLocationArgs;
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
 *         var default_ = new DataSource(&#34;default&#34;, DataSourceArgs.builder()        
 *             .dataSourceId(&#34;example-id&#34;)
 *             .parameters(DataSourceParametersArgs.builder()
 *                 .s3(DataSourceParametersS3Args.builder()
 *                     .manifestFileLocation(DataSourceParametersS3ManifestFileLocationArgs.builder()
 *                         .bucket(&#34;my-bucket&#34;)
 *                         .key(&#34;path/to/manifest.json&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .type(&#34;S3&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * A QuickSight data source can be imported using the AWS account ID, and data source ID name separated by a slash (`/`) e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:quicksight/dataSource:DataSource example 123456789123/my-data-source-id
 * ```
 * 
 */
@ResourceType(type="aws:quicksight/dataSource:DataSource")
public class DataSource extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the data source
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the data source
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     * 
     */
    @Export(name="awsAccountId", refs={String.class}, tree="[0]")
    private Output<String> awsAccountId;

    /**
     * @return The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     * 
     */
    public Output<String> awsAccountId() {
        return this.awsAccountId;
    }
    /**
     * The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
     * 
     */
    @Export(name="credentials", refs={DataSourceCredentials.class}, tree="[0]")
    private Output</* @Nullable */ DataSourceCredentials> credentials;

    /**
     * @return The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
     * 
     */
    public Output<Optional<DataSourceCredentials>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * An identifier for the data source.
     * 
     */
    @Export(name="dataSourceId", refs={String.class}, tree="[0]")
    private Output<String> dataSourceId;

    /**
     * @return An identifier for the data source.
     * 
     */
    public Output<String> dataSourceId() {
        return this.dataSourceId;
    }
    /**
     * A name for the data source, maximum of 128 characters.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name for the data source, maximum of 128 characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parameters used to connect to this data source (exactly one).
     * 
     */
    @Export(name="parameters", refs={DataSourceParameters.class}, tree="[0]")
    private Output<DataSourceParameters> parameters;

    /**
     * @return The parameters used to connect to this data source (exactly one).
     * 
     */
    public Output<DataSourceParameters> parameters() {
        return this.parameters;
    }
    /**
     * A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
     * 
     */
    @Export(name="permissions", refs={List.class,DataSourcePermission.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataSourcePermission>> permissions;

    /**
     * @return A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
     * 
     */
    public Output<Optional<List<DataSourcePermission>>> permissions() {
        return Codegen.optional(this.permissions);
    }
    /**
     * Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
     * 
     */
    @Export(name="sslProperties", refs={DataSourceSslProperties.class}, tree="[0]")
    private Output</* @Nullable */ DataSourceSslProperties> sslProperties;

    /**
     * @return Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
     * 
     */
    public Output<Optional<DataSourceSslProperties>> sslProperties() {
        return Codegen.optional(this.sslProperties);
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
     * The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
     * 
     */
    @Export(name="vpcConnectionProperties", refs={DataSourceVpcConnectionProperties.class}, tree="[0]")
    private Output</* @Nullable */ DataSourceVpcConnectionProperties> vpcConnectionProperties;

    /**
     * @return Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
     * 
     */
    public Output<Optional<DataSourceVpcConnectionProperties>> vpcConnectionProperties() {
        return Codegen.optional(this.vpcConnectionProperties);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataSource(String name) {
        this(name, DataSourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataSource(String name, DataSourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataSource(String name, DataSourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/dataSource:DataSource", name, args == null ? DataSourceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataSource(String name, Output<String> id, @Nullable DataSourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/dataSource:DataSource", name, state, makeResourceOptions(options, id));
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
    public static DataSource get(String name, Output<String> id, @Nullable DataSourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataSource(name, id, state, options);
    }
}
