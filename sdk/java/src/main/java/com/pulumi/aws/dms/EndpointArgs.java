// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dms;

import com.pulumi.aws.dms.inputs.EndpointElasticsearchSettingsArgs;
import com.pulumi.aws.dms.inputs.EndpointKafkaSettingsArgs;
import com.pulumi.aws.dms.inputs.EndpointKinesisSettingsArgs;
import com.pulumi.aws.dms.inputs.EndpointMongodbSettingsArgs;
import com.pulumi.aws.dms.inputs.EndpointPostgresSettingsArgs;
import com.pulumi.aws.dms.inputs.EndpointRedisSettingsArgs;
import com.pulumi.aws.dms.inputs.EndpointRedshiftSettingsArgs;
import com.pulumi.aws.dms.inputs.EndpointS3SettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointArgs Empty = new EndpointArgs();

    /**
     * ARN for the certificate.
     * 
     */
    @Import(name="certificateArn")
    private @Nullable Output<String> certificateArn;

    /**
     * @return ARN for the certificate.
     * 
     */
    public Optional<Output<String>> certificateArn() {
        return Optional.ofNullable(this.certificateArn);
    }

    /**
     * Name of the endpoint database.
     * 
     */
    @Import(name="databaseName")
    private @Nullable Output<String> databaseName;

    /**
     * @return Name of the endpoint database.
     * 
     */
    public Optional<Output<String>> databaseName() {
        return Optional.ofNullable(this.databaseName);
    }

    /**
     * Configuration block for OpenSearch settings. See below.
     * 
     */
    @Import(name="elasticsearchSettings")
    private @Nullable Output<EndpointElasticsearchSettingsArgs> elasticsearchSettings;

    /**
     * @return Configuration block for OpenSearch settings. See below.
     * 
     */
    public Optional<Output<EndpointElasticsearchSettingsArgs>> elasticsearchSettings() {
        return Optional.ofNullable(this.elasticsearchSettings);
    }

    /**
     * Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
     * 
     */
    @Import(name="endpointId", required=true)
    private Output<String> endpointId;

    /**
     * @return Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
     * 
     */
    public Output<String> endpointId() {
        return this.endpointId;
    }

    /**
     * Type of endpoint. Valid values are `source`, `target`.
     * 
     */
    @Import(name="endpointType", required=true)
    private Output<String> endpointType;

    /**
     * @return Type of endpoint. Valid values are `source`, `target`.
     * 
     */
    public Output<String> endpointType() {
        return this.endpointType;
    }

    /**
     * Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
     * 
     */
    @Import(name="engineName", required=true)
    private Output<String> engineName;

    /**
     * @return Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
     * 
     */
    public Output<String> engineName() {
        return this.engineName;
    }

    /**
     * Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
     * 
     */
    @Import(name="extraConnectionAttributes")
    private @Nullable Output<String> extraConnectionAttributes;

    /**
     * @return Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
     * 
     */
    public Optional<Output<String>> extraConnectionAttributes() {
        return Optional.ofNullable(this.extraConnectionAttributes);
    }

    /**
     * Configuration block for Kafka settings. See below.
     * 
     */
    @Import(name="kafkaSettings")
    private @Nullable Output<EndpointKafkaSettingsArgs> kafkaSettings;

    /**
     * @return Configuration block for Kafka settings. See below.
     * 
     */
    public Optional<Output<EndpointKafkaSettingsArgs>> kafkaSettings() {
        return Optional.ofNullable(this.kafkaSettings);
    }

    /**
     * Configuration block for Kinesis settings. See below.
     * 
     */
    @Import(name="kinesisSettings")
    private @Nullable Output<EndpointKinesisSettingsArgs> kinesisSettings;

    /**
     * @return Configuration block for Kinesis settings. See below.
     * 
     */
    public Optional<Output<EndpointKinesisSettingsArgs>> kinesisSettings() {
        return Optional.ofNullable(this.kinesisSettings);
    }

    /**
     * ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="kmsKeyArn")
    private @Nullable Output<String> kmsKeyArn;

    /**
     * @return ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> kmsKeyArn() {
        return Optional.ofNullable(this.kmsKeyArn);
    }

    /**
     * Configuration block for MongoDB settings. See below.
     * 
     */
    @Import(name="mongodbSettings")
    private @Nullable Output<EndpointMongodbSettingsArgs> mongodbSettings;

    /**
     * @return Configuration block for MongoDB settings. See below.
     * 
     */
    public Optional<Output<EndpointMongodbSettingsArgs>> mongodbSettings() {
        return Optional.ofNullable(this.mongodbSettings);
    }

    /**
     * Password to be used to login to the endpoint database.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password to be used to login to the endpoint database.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="pauseReplicationTasks")
    private @Nullable Output<Boolean> pauseReplicationTasks;

    public Optional<Output<Boolean>> pauseReplicationTasks() {
        return Optional.ofNullable(this.pauseReplicationTasks);
    }

    /**
     * Port used by the endpoint database.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port used by the endpoint database.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Configuration block for Postgres settings. See below.
     * 
     */
    @Import(name="postgresSettings")
    private @Nullable Output<EndpointPostgresSettingsArgs> postgresSettings;

    /**
     * @return Configuration block for Postgres settings. See below.
     * 
     */
    public Optional<Output<EndpointPostgresSettingsArgs>> postgresSettings() {
        return Optional.ofNullable(this.postgresSettings);
    }

    @Import(name="redisSettings")
    private @Nullable Output<EndpointRedisSettingsArgs> redisSettings;

    public Optional<Output<EndpointRedisSettingsArgs>> redisSettings() {
        return Optional.ofNullable(this.redisSettings);
    }

    /**
     * Configuration block for Redshift settings. See below.
     * 
     */
    @Import(name="redshiftSettings")
    private @Nullable Output<EndpointRedshiftSettingsArgs> redshiftSettings;

    /**
     * @return Configuration block for Redshift settings. See below.
     * 
     */
    public Optional<Output<EndpointRedshiftSettingsArgs>> redshiftSettings() {
        return Optional.ofNullable(this.redshiftSettings);
    }

    /**
     * (**Deprecated**, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
     * 
     */
    @Import(name="s3Settings")
    private @Nullable Output<EndpointS3SettingsArgs> s3Settings;

    /**
     * @return (**Deprecated**, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
     * 
     */
    public Optional<Output<EndpointS3SettingsArgs>> s3Settings() {
        return Optional.ofNullable(this.s3Settings);
    }

    /**
     * ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
     * 
     */
    @Import(name="secretsManagerAccessRoleArn")
    private @Nullable Output<String> secretsManagerAccessRoleArn;

    /**
     * @return ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
     * 
     */
    public Optional<Output<String>> secretsManagerAccessRoleArn() {
        return Optional.ofNullable(this.secretsManagerAccessRoleArn);
    }

    /**
     * Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
     * 
     */
    @Import(name="secretsManagerArn")
    private @Nullable Output<String> secretsManagerArn;

    /**
     * @return Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
     * 
     */
    public Optional<Output<String>> secretsManagerArn() {
        return Optional.ofNullable(this.secretsManagerArn);
    }

    /**
     * Host name of the server.
     * 
     */
    @Import(name="serverName")
    private @Nullable Output<String> serverName;

    /**
     * @return Host name of the server.
     * 
     */
    public Optional<Output<String>> serverName() {
        return Optional.ofNullable(this.serverName);
    }

    /**
     * ARN used by the service access IAM role for dynamodb endpoints.
     * 
     */
    @Import(name="serviceAccessRole")
    private @Nullable Output<String> serviceAccessRole;

    /**
     * @return ARN used by the service access IAM role for dynamodb endpoints.
     * 
     */
    public Optional<Output<String>> serviceAccessRole() {
        return Optional.ofNullable(this.serviceAccessRole);
    }

    /**
     * SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
     * 
     */
    @Import(name="sslMode")
    private @Nullable Output<String> sslMode;

    /**
     * @return SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
     * 
     */
    public Optional<Output<String>> sslMode() {
        return Optional.ofNullable(this.sslMode);
    }

    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * User name to be used to login to the endpoint database.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return User name to be used to login to the endpoint database.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private EndpointArgs() {}

    private EndpointArgs(EndpointArgs $) {
        this.certificateArn = $.certificateArn;
        this.databaseName = $.databaseName;
        this.elasticsearchSettings = $.elasticsearchSettings;
        this.endpointId = $.endpointId;
        this.endpointType = $.endpointType;
        this.engineName = $.engineName;
        this.extraConnectionAttributes = $.extraConnectionAttributes;
        this.kafkaSettings = $.kafkaSettings;
        this.kinesisSettings = $.kinesisSettings;
        this.kmsKeyArn = $.kmsKeyArn;
        this.mongodbSettings = $.mongodbSettings;
        this.password = $.password;
        this.pauseReplicationTasks = $.pauseReplicationTasks;
        this.port = $.port;
        this.postgresSettings = $.postgresSettings;
        this.redisSettings = $.redisSettings;
        this.redshiftSettings = $.redshiftSettings;
        this.s3Settings = $.s3Settings;
        this.secretsManagerAccessRoleArn = $.secretsManagerAccessRoleArn;
        this.secretsManagerArn = $.secretsManagerArn;
        this.serverName = $.serverName;
        this.serviceAccessRole = $.serviceAccessRole;
        this.sslMode = $.sslMode;
        this.tags = $.tags;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointArgs $;

        public Builder() {
            $ = new EndpointArgs();
        }

        public Builder(EndpointArgs defaults) {
            $ = new EndpointArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificateArn ARN for the certificate.
         * 
         * @return builder
         * 
         */
        public Builder certificateArn(@Nullable Output<String> certificateArn) {
            $.certificateArn = certificateArn;
            return this;
        }

        /**
         * @param certificateArn ARN for the certificate.
         * 
         * @return builder
         * 
         */
        public Builder certificateArn(String certificateArn) {
            return certificateArn(Output.of(certificateArn));
        }

        /**
         * @param databaseName Name of the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder databaseName(@Nullable Output<String> databaseName) {
            $.databaseName = databaseName;
            return this;
        }

        /**
         * @param databaseName Name of the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder databaseName(String databaseName) {
            return databaseName(Output.of(databaseName));
        }

        /**
         * @param elasticsearchSettings Configuration block for OpenSearch settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder elasticsearchSettings(@Nullable Output<EndpointElasticsearchSettingsArgs> elasticsearchSettings) {
            $.elasticsearchSettings = elasticsearchSettings;
            return this;
        }

        /**
         * @param elasticsearchSettings Configuration block for OpenSearch settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder elasticsearchSettings(EndpointElasticsearchSettingsArgs elasticsearchSettings) {
            return elasticsearchSettings(Output.of(elasticsearchSettings));
        }

        /**
         * @param endpointId Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
         * 
         * @return builder
         * 
         */
        public Builder endpointId(Output<String> endpointId) {
            $.endpointId = endpointId;
            return this;
        }

        /**
         * @param endpointId Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
         * 
         * @return builder
         * 
         */
        public Builder endpointId(String endpointId) {
            return endpointId(Output.of(endpointId));
        }

        /**
         * @param endpointType Type of endpoint. Valid values are `source`, `target`.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(Output<String> endpointType) {
            $.endpointType = endpointType;
            return this;
        }

        /**
         * @param endpointType Type of endpoint. Valid values are `source`, `target`.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(String endpointType) {
            return endpointType(Output.of(endpointType));
        }

        /**
         * @param engineName Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
         * 
         * @return builder
         * 
         */
        public Builder engineName(Output<String> engineName) {
            $.engineName = engineName;
            return this;
        }

        /**
         * @param engineName Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
         * 
         * @return builder
         * 
         */
        public Builder engineName(String engineName) {
            return engineName(Output.of(engineName));
        }

        /**
         * @param extraConnectionAttributes Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
         * 
         * @return builder
         * 
         */
        public Builder extraConnectionAttributes(@Nullable Output<String> extraConnectionAttributes) {
            $.extraConnectionAttributes = extraConnectionAttributes;
            return this;
        }

        /**
         * @param extraConnectionAttributes Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
         * 
         * @return builder
         * 
         */
        public Builder extraConnectionAttributes(String extraConnectionAttributes) {
            return extraConnectionAttributes(Output.of(extraConnectionAttributes));
        }

        /**
         * @param kafkaSettings Configuration block for Kafka settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder kafkaSettings(@Nullable Output<EndpointKafkaSettingsArgs> kafkaSettings) {
            $.kafkaSettings = kafkaSettings;
            return this;
        }

        /**
         * @param kafkaSettings Configuration block for Kafka settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder kafkaSettings(EndpointKafkaSettingsArgs kafkaSettings) {
            return kafkaSettings(Output.of(kafkaSettings));
        }

        /**
         * @param kinesisSettings Configuration block for Kinesis settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder kinesisSettings(@Nullable Output<EndpointKinesisSettingsArgs> kinesisSettings) {
            $.kinesisSettings = kinesisSettings;
            return this;
        }

        /**
         * @param kinesisSettings Configuration block for Kinesis settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder kinesisSettings(EndpointKinesisSettingsArgs kinesisSettings) {
            return kinesisSettings(Output.of(kinesisSettings));
        }

        /**
         * @param kmsKeyArn ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(@Nullable Output<String> kmsKeyArn) {
            $.kmsKeyArn = kmsKeyArn;
            return this;
        }

        /**
         * @param kmsKeyArn ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(String kmsKeyArn) {
            return kmsKeyArn(Output.of(kmsKeyArn));
        }

        /**
         * @param mongodbSettings Configuration block for MongoDB settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder mongodbSettings(@Nullable Output<EndpointMongodbSettingsArgs> mongodbSettings) {
            $.mongodbSettings = mongodbSettings;
            return this;
        }

        /**
         * @param mongodbSettings Configuration block for MongoDB settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder mongodbSettings(EndpointMongodbSettingsArgs mongodbSettings) {
            return mongodbSettings(Output.of(mongodbSettings));
        }

        /**
         * @param password Password to be used to login to the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password to be used to login to the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder pauseReplicationTasks(@Nullable Output<Boolean> pauseReplicationTasks) {
            $.pauseReplicationTasks = pauseReplicationTasks;
            return this;
        }

        public Builder pauseReplicationTasks(Boolean pauseReplicationTasks) {
            return pauseReplicationTasks(Output.of(pauseReplicationTasks));
        }

        /**
         * @param port Port used by the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port used by the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param postgresSettings Configuration block for Postgres settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder postgresSettings(@Nullable Output<EndpointPostgresSettingsArgs> postgresSettings) {
            $.postgresSettings = postgresSettings;
            return this;
        }

        /**
         * @param postgresSettings Configuration block for Postgres settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder postgresSettings(EndpointPostgresSettingsArgs postgresSettings) {
            return postgresSettings(Output.of(postgresSettings));
        }

        public Builder redisSettings(@Nullable Output<EndpointRedisSettingsArgs> redisSettings) {
            $.redisSettings = redisSettings;
            return this;
        }

        public Builder redisSettings(EndpointRedisSettingsArgs redisSettings) {
            return redisSettings(Output.of(redisSettings));
        }

        /**
         * @param redshiftSettings Configuration block for Redshift settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder redshiftSettings(@Nullable Output<EndpointRedshiftSettingsArgs> redshiftSettings) {
            $.redshiftSettings = redshiftSettings;
            return this;
        }

        /**
         * @param redshiftSettings Configuration block for Redshift settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder redshiftSettings(EndpointRedshiftSettingsArgs redshiftSettings) {
            return redshiftSettings(Output.of(redshiftSettings));
        }

        /**
         * @param s3Settings (**Deprecated**, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder s3Settings(@Nullable Output<EndpointS3SettingsArgs> s3Settings) {
            $.s3Settings = s3Settings;
            return this;
        }

        /**
         * @param s3Settings (**Deprecated**, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
         * 
         * @return builder
         * 
         */
        public Builder s3Settings(EndpointS3SettingsArgs s3Settings) {
            return s3Settings(Output.of(s3Settings));
        }

        /**
         * @param secretsManagerAccessRoleArn ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
         * 
         * @return builder
         * 
         */
        public Builder secretsManagerAccessRoleArn(@Nullable Output<String> secretsManagerAccessRoleArn) {
            $.secretsManagerAccessRoleArn = secretsManagerAccessRoleArn;
            return this;
        }

        /**
         * @param secretsManagerAccessRoleArn ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
         * 
         * @return builder
         * 
         */
        public Builder secretsManagerAccessRoleArn(String secretsManagerAccessRoleArn) {
            return secretsManagerAccessRoleArn(Output.of(secretsManagerAccessRoleArn));
        }

        /**
         * @param secretsManagerArn Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
         * 
         * @return builder
         * 
         */
        public Builder secretsManagerArn(@Nullable Output<String> secretsManagerArn) {
            $.secretsManagerArn = secretsManagerArn;
            return this;
        }

        /**
         * @param secretsManagerArn Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
         * 
         * @return builder
         * 
         */
        public Builder secretsManagerArn(String secretsManagerArn) {
            return secretsManagerArn(Output.of(secretsManagerArn));
        }

        /**
         * @param serverName Host name of the server.
         * 
         * @return builder
         * 
         */
        public Builder serverName(@Nullable Output<String> serverName) {
            $.serverName = serverName;
            return this;
        }

        /**
         * @param serverName Host name of the server.
         * 
         * @return builder
         * 
         */
        public Builder serverName(String serverName) {
            return serverName(Output.of(serverName));
        }

        /**
         * @param serviceAccessRole ARN used by the service access IAM role for dynamodb endpoints.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccessRole(@Nullable Output<String> serviceAccessRole) {
            $.serviceAccessRole = serviceAccessRole;
            return this;
        }

        /**
         * @param serviceAccessRole ARN used by the service access IAM role for dynamodb endpoints.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccessRole(String serviceAccessRole) {
            return serviceAccessRole(Output.of(serviceAccessRole));
        }

        /**
         * @param sslMode SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
         * 
         * @return builder
         * 
         */
        public Builder sslMode(@Nullable Output<String> sslMode) {
            $.sslMode = sslMode;
            return this;
        }

        /**
         * @param sslMode SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
         * 
         * @return builder
         * 
         */
        public Builder sslMode(String sslMode) {
            return sslMode(Output.of(sslMode));
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param username User name to be used to login to the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username User name to be used to login to the endpoint database.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public EndpointArgs build() {
            $.endpointId = Objects.requireNonNull($.endpointId, "expected parameter 'endpointId' to be non-null");
            $.endpointType = Objects.requireNonNull($.endpointType, "expected parameter 'endpointType' to be non-null");
            $.engineName = Objects.requireNonNull($.engineName, "expected parameter 'engineName' to be non-null");
            return $;
        }
    }

}
