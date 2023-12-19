// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms
{
    /// <summary>
    /// Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.
    /// 
    /// &gt; **Note:** All arguments including the password will be stored in the raw state as plain-text. &gt; **Note:** The `s3_settings` argument is deprecated, may not be maintained, and will be removed in a future version. Use the `aws.dms.S3Endpoint` resource instead.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new endpoint
    ///     var test = new Aws.Dms.Endpoint("test", new()
    ///     {
    ///         CertificateArn = "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
    ///         DatabaseName = "test",
    ///         EndpointId = "test-dms-endpoint-tf",
    ///         EndpointType = "source",
    ///         EngineName = "aurora",
    ///         ExtraConnectionAttributes = "",
    ///         KmsKeyArn = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    ///         Password = "test",
    ///         Port = 3306,
    ///         ServerName = "test",
    ///         SslMode = "none",
    ///         Tags = 
    ///         {
    ///             { "Name", "test" },
    ///         },
    ///         Username = "test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import endpoints using the `endpoint_id`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:dms/endpoint:Endpoint test test-dms-endpoint-tf
    /// ```
    /// </summary>
    [AwsResourceType("aws:dms/endpoint:Endpoint")]
    public partial class Endpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN for the certificate.
        /// </summary>
        [Output("certificateArn")]
        public Output<string> CertificateArn { get; private set; } = null!;

        /// <summary>
        /// Name of the endpoint database.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Configuration block for OpenSearch settings. See below.
        /// </summary>
        [Output("elasticsearchSettings")]
        public Output<Outputs.EndpointElasticsearchSettings?> ElasticsearchSettings { get; private set; } = null!;

        /// <summary>
        /// ARN for the endpoint.
        /// </summary>
        [Output("endpointArn")]
        public Output<string> EndpointArn { get; private set; } = null!;

        /// <summary>
        /// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// Type of endpoint. Valid values are `source`, `target`.
        /// </summary>
        [Output("endpointType")]
        public Output<string> EndpointType { get; private set; } = null!;

        /// <summary>
        /// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
        /// </summary>
        [Output("engineName")]
        public Output<string> EngineName { get; private set; } = null!;

        /// <summary>
        /// Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
        /// </summary>
        [Output("extraConnectionAttributes")]
        public Output<string> ExtraConnectionAttributes { get; private set; } = null!;

        /// <summary>
        /// Configuration block for Kafka settings. See below.
        /// </summary>
        [Output("kafkaSettings")]
        public Output<Outputs.EndpointKafkaSettings?> KafkaSettings { get; private set; } = null!;

        /// <summary>
        /// Configuration block for Kinesis settings. See below.
        /// </summary>
        [Output("kinesisSettings")]
        public Output<Outputs.EndpointKinesisSettings?> KinesisSettings { get; private set; } = null!;

        /// <summary>
        /// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// Configuration block for MongoDB settings. See below.
        /// </summary>
        [Output("mongodbSettings")]
        public Output<Outputs.EndpointMongodbSettings?> MongodbSettings { get; private set; } = null!;

        /// <summary>
        /// Password to be used to login to the endpoint database.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("pauseReplicationTasks")]
        public Output<bool?> PauseReplicationTasks { get; private set; } = null!;

        /// <summary>
        /// Port used by the endpoint database.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// Configuration block for Postgres settings. See below.
        /// </summary>
        [Output("postgresSettings")]
        public Output<Outputs.EndpointPostgresSettings?> PostgresSettings { get; private set; } = null!;

        [Output("redisSettings")]
        public Output<Outputs.EndpointRedisSettings?> RedisSettings { get; private set; } = null!;

        /// <summary>
        /// Configuration block for Redshift settings. See below.
        /// </summary>
        [Output("redshiftSettings")]
        public Output<Outputs.EndpointRedshiftSettings> RedshiftSettings { get; private set; } = null!;

        /// <summary>
        /// (**Deprecated**, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
        /// </summary>
        [Output("s3Settings")]
        public Output<Outputs.EndpointS3Settings?> S3Settings { get; private set; } = null!;

        /// <summary>
        /// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
        /// </summary>
        [Output("secretsManagerAccessRoleArn")]
        public Output<string?> SecretsManagerAccessRoleArn { get; private set; } = null!;

        /// <summary>
        /// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
        /// </summary>
        [Output("secretsManagerArn")]
        public Output<string?> SecretsManagerArn { get; private set; } = null!;

        /// <summary>
        /// Host name of the server.
        /// </summary>
        [Output("serverName")]
        public Output<string?> ServerName { get; private set; } = null!;

        /// <summary>
        /// ARN used by the service access IAM role for dynamodb endpoints.
        /// </summary>
        [Output("serviceAccessRole")]
        public Output<string?> ServiceAccessRole { get; private set; } = null!;

        /// <summary>
        /// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
        /// </summary>
        [Output("sslMode")]
        public Output<string> SslMode { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// User name to be used to login to the endpoint database.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("aws:dms/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("aws:dms/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN for the certificate.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// Name of the endpoint database.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Configuration block for OpenSearch settings. See below.
        /// </summary>
        [Input("elasticsearchSettings")]
        public Input<Inputs.EndpointElasticsearchSettingsArgs>? ElasticsearchSettings { get; set; }

        /// <summary>
        /// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
        /// </summary>
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        /// <summary>
        /// Type of endpoint. Valid values are `source`, `target`.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
        /// </summary>
        [Input("engineName", required: true)]
        public Input<string> EngineName { get; set; } = null!;

        /// <summary>
        /// Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
        /// </summary>
        [Input("extraConnectionAttributes")]
        public Input<string>? ExtraConnectionAttributes { get; set; }

        /// <summary>
        /// Configuration block for Kafka settings. See below.
        /// </summary>
        [Input("kafkaSettings")]
        public Input<Inputs.EndpointKafkaSettingsArgs>? KafkaSettings { get; set; }

        /// <summary>
        /// Configuration block for Kinesis settings. See below.
        /// </summary>
        [Input("kinesisSettings")]
        public Input<Inputs.EndpointKinesisSettingsArgs>? KinesisSettings { get; set; }

        /// <summary>
        /// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Configuration block for MongoDB settings. See below.
        /// </summary>
        [Input("mongodbSettings")]
        public Input<Inputs.EndpointMongodbSettingsArgs>? MongodbSettings { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password to be used to login to the endpoint database.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("pauseReplicationTasks")]
        public Input<bool>? PauseReplicationTasks { get; set; }

        /// <summary>
        /// Port used by the endpoint database.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Configuration block for Postgres settings. See below.
        /// </summary>
        [Input("postgresSettings")]
        public Input<Inputs.EndpointPostgresSettingsArgs>? PostgresSettings { get; set; }

        [Input("redisSettings")]
        public Input<Inputs.EndpointRedisSettingsArgs>? RedisSettings { get; set; }

        /// <summary>
        /// Configuration block for Redshift settings. See below.
        /// </summary>
        [Input("redshiftSettings")]
        public Input<Inputs.EndpointRedshiftSettingsArgs>? RedshiftSettings { get; set; }

        /// <summary>
        /// (**Deprecated**, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
        /// </summary>
        [Input("s3Settings")]
        public Input<Inputs.EndpointS3SettingsArgs>? S3Settings { get; set; }

        /// <summary>
        /// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
        /// </summary>
        [Input("secretsManagerAccessRoleArn")]
        public Input<string>? SecretsManagerAccessRoleArn { get; set; }

        /// <summary>
        /// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
        /// </summary>
        [Input("secretsManagerArn")]
        public Input<string>? SecretsManagerArn { get; set; }

        /// <summary>
        /// Host name of the server.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// ARN used by the service access IAM role for dynamodb endpoints.
        /// </summary>
        [Input("serviceAccessRole")]
        public Input<string>? ServiceAccessRole { get; set; }

        /// <summary>
        /// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// User name to be used to login to the endpoint database.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public EndpointArgs()
        {
        }
        public static new EndpointArgs Empty => new EndpointArgs();
    }

    public sealed class EndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN for the certificate.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// Name of the endpoint database.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Configuration block for OpenSearch settings. See below.
        /// </summary>
        [Input("elasticsearchSettings")]
        public Input<Inputs.EndpointElasticsearchSettingsGetArgs>? ElasticsearchSettings { get; set; }

        /// <summary>
        /// ARN for the endpoint.
        /// </summary>
        [Input("endpointArn")]
        public Input<string>? EndpointArn { get; set; }

        /// <summary>
        /// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Type of endpoint. Valid values are `source`, `target`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
        /// </summary>
        [Input("engineName")]
        public Input<string>? EngineName { get; set; }

        /// <summary>
        /// Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
        /// </summary>
        [Input("extraConnectionAttributes")]
        public Input<string>? ExtraConnectionAttributes { get; set; }

        /// <summary>
        /// Configuration block for Kafka settings. See below.
        /// </summary>
        [Input("kafkaSettings")]
        public Input<Inputs.EndpointKafkaSettingsGetArgs>? KafkaSettings { get; set; }

        /// <summary>
        /// Configuration block for Kinesis settings. See below.
        /// </summary>
        [Input("kinesisSettings")]
        public Input<Inputs.EndpointKinesisSettingsGetArgs>? KinesisSettings { get; set; }

        /// <summary>
        /// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Configuration block for MongoDB settings. See below.
        /// </summary>
        [Input("mongodbSettings")]
        public Input<Inputs.EndpointMongodbSettingsGetArgs>? MongodbSettings { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password to be used to login to the endpoint database.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("pauseReplicationTasks")]
        public Input<bool>? PauseReplicationTasks { get; set; }

        /// <summary>
        /// Port used by the endpoint database.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Configuration block for Postgres settings. See below.
        /// </summary>
        [Input("postgresSettings")]
        public Input<Inputs.EndpointPostgresSettingsGetArgs>? PostgresSettings { get; set; }

        [Input("redisSettings")]
        public Input<Inputs.EndpointRedisSettingsGetArgs>? RedisSettings { get; set; }

        /// <summary>
        /// Configuration block for Redshift settings. See below.
        /// </summary>
        [Input("redshiftSettings")]
        public Input<Inputs.EndpointRedshiftSettingsGetArgs>? RedshiftSettings { get; set; }

        /// <summary>
        /// (**Deprecated**, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
        /// </summary>
        [Input("s3Settings")]
        public Input<Inputs.EndpointS3SettingsGetArgs>? S3Settings { get; set; }

        /// <summary>
        /// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
        /// </summary>
        [Input("secretsManagerAccessRoleArn")]
        public Input<string>? SecretsManagerAccessRoleArn { get; set; }

        /// <summary>
        /// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
        /// </summary>
        [Input("secretsManagerArn")]
        public Input<string>? SecretsManagerArn { get; set; }

        /// <summary>
        /// Host name of the server.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// ARN used by the service access IAM role for dynamodb endpoints.
        /// </summary>
        [Input("serviceAccessRole")]
        public Input<string>? ServiceAccessRole { get; set; }

        /// <summary>
        /// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// User name to be used to login to the endpoint database.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public EndpointState()
        {
        }
        public static new EndpointState Empty => new EndpointState();
    }
}
