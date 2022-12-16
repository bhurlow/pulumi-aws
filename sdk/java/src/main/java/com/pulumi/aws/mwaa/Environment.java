// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.mwaa;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.mwaa.EnvironmentArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentState;
import com.pulumi.aws.mwaa.outputs.EnvironmentLastUpdated;
import com.pulumi.aws.mwaa.outputs.EnvironmentLoggingConfiguration;
import com.pulumi.aws.mwaa.outputs.EnvironmentNetworkConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a MWAA Environment resource.
 * 
 * ## Example Usage
 * 
 * A MWAA Environment requires an IAM role (`aws.iam.Role`), two subnets in the private zone (`aws.ec2.Subnet`) and a versioned S3 bucket (`aws.s3.BucketV2`).
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.mwaa.Environment;
 * import com.pulumi.aws.mwaa.EnvironmentArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
 *         var example = new Environment(&#34;example&#34;, EnvironmentArgs.builder()        
 *             .dagS3Path(&#34;dags/&#34;)
 *             .executionRoleArn(aws_iam_role.example().arn())
 *             .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
 *                 .securityGroupIds(aws_security_group.example().id())
 *                 .subnetIds(aws_subnet.private().stream().map(element -&gt; element.id()).collect(toList()))
 *                 .build())
 *             .sourceBucketArn(aws_s3_bucket.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example with Airflow configuration options
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.mwaa.Environment;
 * import com.pulumi.aws.mwaa.EnvironmentArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
 *         var example = new Environment(&#34;example&#34;, EnvironmentArgs.builder()        
 *             .airflowConfigurationOptions(Map.ofEntries(
 *                 Map.entry(&#34;core.default_task_retries&#34;, 16),
 *                 Map.entry(&#34;core.parallelism&#34;, 1)
 *             ))
 *             .dagS3Path(&#34;dags/&#34;)
 *             .executionRoleArn(aws_iam_role.example().arn())
 *             .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
 *                 .securityGroupIds(aws_security_group.example().id())
 *                 .subnetIds(aws_subnet.private().stream().map(element -&gt; element.id()).collect(toList()))
 *                 .build())
 *             .sourceBucketArn(aws_s3_bucket.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example with logging configurations
 * 
 * Note that Airflow task logs are enabled by default with the `INFO` log level.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.mwaa.Environment;
 * import com.pulumi.aws.mwaa.EnvironmentArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationDagProcessingLogsArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationSchedulerLogsArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationTaskLogsArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationWebserverLogsArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationWorkerLogsArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
 *         var example = new Environment(&#34;example&#34;, EnvironmentArgs.builder()        
 *             .dagS3Path(&#34;dags/&#34;)
 *             .executionRoleArn(aws_iam_role.example().arn())
 *             .loggingConfiguration(EnvironmentLoggingConfigurationArgs.builder()
 *                 .dagProcessingLogs(EnvironmentLoggingConfigurationDagProcessingLogsArgs.builder()
 *                     .enabled(true)
 *                     .logLevel(&#34;DEBUG&#34;)
 *                     .build())
 *                 .schedulerLogs(EnvironmentLoggingConfigurationSchedulerLogsArgs.builder()
 *                     .enabled(true)
 *                     .logLevel(&#34;INFO&#34;)
 *                     .build())
 *                 .taskLogs(EnvironmentLoggingConfigurationTaskLogsArgs.builder()
 *                     .enabled(true)
 *                     .logLevel(&#34;WARNING&#34;)
 *                     .build())
 *                 .webserverLogs(EnvironmentLoggingConfigurationWebserverLogsArgs.builder()
 *                     .enabled(true)
 *                     .logLevel(&#34;ERROR&#34;)
 *                     .build())
 *                 .workerLogs(EnvironmentLoggingConfigurationWorkerLogsArgs.builder()
 *                     .enabled(true)
 *                     .logLevel(&#34;CRITICAL&#34;)
 *                     .build())
 *                 .build())
 *             .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
 *                 .securityGroupIds(aws_security_group.example().id())
 *                 .subnetIds(aws_subnet.private().stream().map(element -&gt; element.id()).collect(toList()))
 *                 .build())
 *             .sourceBucketArn(aws_s3_bucket.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example with tags
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.mwaa.Environment;
 * import com.pulumi.aws.mwaa.EnvironmentArgs;
 * import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
 *         var example = new Environment(&#34;example&#34;, EnvironmentArgs.builder()        
 *             .dagS3Path(&#34;dags/&#34;)
 *             .executionRoleArn(aws_iam_role.example().arn())
 *             .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
 *                 .securityGroupIds(aws_security_group.example().id())
 *                 .subnetIds(aws_subnet.private().stream().map(element -&gt; element.id()).collect(toList()))
 *                 .build())
 *             .sourceBucketArn(aws_s3_bucket.example().arn())
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Name&#34;, &#34;example&#34;),
 *                 Map.entry(&#34;Environment&#34;, &#34;production&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * MWAA Environment can be imported using `Name` e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:mwaa/environment:Environment example MyAirflowEnvironment
 * ```
 * 
 */
@ResourceType(type="aws:mwaa/environment:Environment")
public class Environment extends com.pulumi.resources.CustomResource {
    /**
     * The `airflow_configuration_options` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
     * 
     */
    @Export(name="airflowConfigurationOptions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> airflowConfigurationOptions;

    /**
     * @return The `airflow_configuration_options` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
     * 
     */
    public Output<Optional<Map<String,String>>> airflowConfigurationOptions() {
        return Codegen.optional(this.airflowConfigurationOptions);
    }
    /**
     * Airflow version of your environment, will be set by default to the latest version that MWAA supports.
     * 
     */
    @Export(name="airflowVersion", refs={String.class}, tree="[0]")
    private Output<String> airflowVersion;

    /**
     * @return Airflow version of your environment, will be set by default to the latest version that MWAA supports.
     * 
     */
    public Output<String> airflowVersion() {
        return this.airflowVersion;
    }
    /**
     * The ARN of the MWAA Environment
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the MWAA Environment
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The Created At date of the MWAA Environment
     * * `logging_configuration[0].&lt;LOG_CONFIGURATION_TYPE&gt;[0].cloud_watch_log_group_arn` - Provides the ARN for the CloudWatch group where the logs will be published
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The Created At date of the MWAA Environment
     * * `logging_configuration[0].&lt;LOG_CONFIGURATION_TYPE&gt;[0].cloud_watch_log_group_arn` - Provides the ARN for the CloudWatch group where the logs will be published
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
     * 
     */
    @Export(name="dagS3Path", refs={String.class}, tree="[0]")
    private Output<String> dagS3Path;

    /**
     * @return The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
     * 
     */
    public Output<String> dagS3Path() {
        return this.dagS3Path;
    }
    /**
     * Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
     * 
     */
    @Export(name="environmentClass", refs={String.class}, tree="[0]")
    private Output<String> environmentClass;

    /**
     * @return Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
     * 
     */
    public Output<String> environmentClass() {
        return this.environmentClass;
    }
    /**
     * The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
     * 
     */
    @Export(name="executionRoleArn", refs={String.class}, tree="[0]")
    private Output<String> executionRoleArn;

    /**
     * @return The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
     * 
     */
    public Output<String> executionRoleArn() {
        return this.executionRoleArn;
    }
    /**
     * The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
     * 
     */
    @Export(name="kmsKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKey;

    /**
     * @return The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
     * 
     */
    public Output<Optional<String>> kmsKey() {
        return Codegen.optional(this.kmsKey);
    }
    @Export(name="lastUpdateds", refs={List.class,EnvironmentLastUpdated.class}, tree="[0,1]")
    private Output<List<EnvironmentLastUpdated>> lastUpdateds;

    public Output<List<EnvironmentLastUpdated>> lastUpdateds() {
        return this.lastUpdateds;
    }
    /**
     * The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
     * 
     */
    @Export(name="loggingConfiguration", refs={EnvironmentLoggingConfiguration.class}, tree="[0]")
    private Output<EnvironmentLoggingConfiguration> loggingConfiguration;

    /**
     * @return The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
     * 
     */
    public Output<EnvironmentLoggingConfiguration> loggingConfiguration() {
        return this.loggingConfiguration;
    }
    /**
     * The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
     * 
     */
    @Export(name="maxWorkers", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxWorkers;

    /**
     * @return The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
     * 
     */
    public Output<Integer> maxWorkers() {
        return this.maxWorkers;
    }
    /**
     * The minimum number of workers that you want to run in your environment. Will be `1` by default.
     * 
     */
    @Export(name="minWorkers", refs={Integer.class}, tree="[0]")
    private Output<Integer> minWorkers;

    /**
     * @return The minimum number of workers that you want to run in your environment. Will be `1` by default.
     * 
     */
    public Output<Integer> minWorkers() {
        return this.minWorkers;
    }
    /**
     * The name of the Apache Airflow Environment
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Apache Airflow Environment
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
     * 
     */
    @Export(name="networkConfiguration", refs={EnvironmentNetworkConfiguration.class}, tree="[0]")
    private Output<EnvironmentNetworkConfiguration> networkConfiguration;

    /**
     * @return Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
     * 
     */
    public Output<EnvironmentNetworkConfiguration> networkConfiguration() {
        return this.networkConfiguration;
    }
    /**
     * The plugins.zip file version you want to use.
     * 
     */
    @Export(name="pluginsS3ObjectVersion", refs={String.class}, tree="[0]")
    private Output<String> pluginsS3ObjectVersion;

    /**
     * @return The plugins.zip file version you want to use.
     * 
     */
    public Output<String> pluginsS3ObjectVersion() {
        return this.pluginsS3ObjectVersion;
    }
    /**
     * The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then plugins_s3_object_version is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
     * 
     */
    @Export(name="pluginsS3Path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pluginsS3Path;

    /**
     * @return The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then plugins_s3_object_version is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
     * 
     */
    public Output<Optional<String>> pluginsS3Path() {
        return Codegen.optional(this.pluginsS3Path);
    }
    /**
     * The requirements.txt file version you want to use.
     * 
     */
    @Export(name="requirementsS3ObjectVersion", refs={String.class}, tree="[0]")
    private Output<String> requirementsS3ObjectVersion;

    /**
     * @return The requirements.txt file version you want to use.
     * 
     */
    public Output<String> requirementsS3ObjectVersion() {
        return this.requirementsS3ObjectVersion;
    }
    /**
     * The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirements_s3_object_version is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
     * 
     */
    @Export(name="requirementsS3Path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requirementsS3Path;

    /**
     * @return The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirements_s3_object_version is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
     * 
     */
    public Output<Optional<String>> requirementsS3Path() {
        return Codegen.optional(this.requirementsS3Path);
    }
    /**
     * The number of schedulers that you want to run in your environment. v2.0.2 and above accepts `2` - `5`, default `2`. v1.10.12 accepts `1`.
     * 
     */
    @Export(name="schedulers", refs={Integer.class}, tree="[0]")
    private Output<Integer> schedulers;

    /**
     * @return The number of schedulers that you want to run in your environment. v2.0.2 and above accepts `2` - `5`, default `2`. v1.10.12 accepts `1`.
     * 
     */
    public Output<Integer> schedulers() {
        return this.schedulers;
    }
    /**
     * The Service Role ARN of the Amazon MWAA Environment
     * 
     */
    @Export(name="serviceRoleArn", refs={String.class}, tree="[0]")
    private Output<String> serviceRoleArn;

    /**
     * @return The Service Role ARN of the Amazon MWAA Environment
     * 
     */
    public Output<String> serviceRoleArn() {
        return this.serviceRoleArn;
    }
    /**
     * The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
     * 
     */
    @Export(name="sourceBucketArn", refs={String.class}, tree="[0]")
    private Output<String> sourceBucketArn;

    /**
     * @return The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
     * 
     */
    public Output<String> sourceBucketArn() {
        return this.sourceBucketArn;
    }
    /**
     * The status of the Amazon MWAA Environment
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Amazon MWAA Environment
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A map of resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
     * 
     */
    @Export(name="webserverAccessMode", refs={String.class}, tree="[0]")
    private Output<String> webserverAccessMode;

    /**
     * @return Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
     * 
     */
    public Output<String> webserverAccessMode() {
        return this.webserverAccessMode;
    }
    /**
     * The webserver URL of the MWAA Environment
     * 
     */
    @Export(name="webserverUrl", refs={String.class}, tree="[0]")
    private Output<String> webserverUrl;

    /**
     * @return The webserver URL of the MWAA Environment
     * 
     */
    public Output<String> webserverUrl() {
        return this.webserverUrl;
    }
    /**
     * Specifies the start date for the weekly maintenance window.
     * 
     */
    @Export(name="weeklyMaintenanceWindowStart", refs={String.class}, tree="[0]")
    private Output<String> weeklyMaintenanceWindowStart;

    /**
     * @return Specifies the start date for the weekly maintenance window.
     * 
     */
    public Output<String> weeklyMaintenanceWindowStart() {
        return this.weeklyMaintenanceWindowStart;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Environment(String name) {
        this(name, EnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Environment(String name, EnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Environment(String name, EnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:mwaa/environment:Environment", name, args == null ? EnvironmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Environment(String name, Output<String> id, @Nullable EnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:mwaa/environment:Environment", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "airflowConfigurationOptions"
            ))
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
    public static Environment get(String name, Output<String> id, @Nullable EnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Environment(name, id, state, options);
    }
}
