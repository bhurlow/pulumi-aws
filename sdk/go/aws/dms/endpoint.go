// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dms.NewEndpoint(ctx, "test", &dms.EndpointArgs{
//				CertificateArn:            pulumi.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"),
//				DatabaseName:              pulumi.String("test"),
//				EndpointId:                pulumi.String("test-dms-endpoint-tf"),
//				EndpointType:              pulumi.String("source"),
//				EngineName:                pulumi.String("aurora"),
//				ExtraConnectionAttributes: pulumi.String(""),
//				KmsKeyArn:                 pulumi.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"),
//				Password:                  pulumi.String("test"),
//				Port:                      pulumi.Int(3306),
//				ServerName:                pulumi.String("test"),
//				SslMode:                   pulumi.String("none"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("test"),
//				},
//				Username: pulumi.String("test"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Endpoints can be imported using the `endpoint_id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:dms/endpoint:Endpoint test test-dms-endpoint-tf
//
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// ARN for the certificate.
	CertificateArn pulumi.StringOutput `pulumi:"certificateArn"`
	// Name of the endpoint database.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// Configuration block for OpenSearch settings. See below.
	ElasticsearchSettings EndpointElasticsearchSettingsPtrOutput `pulumi:"elasticsearchSettings"`
	// ARN for the endpoint.
	EndpointArn pulumi.StringOutput `pulumi:"endpointArn"`
	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// Type of endpoint. Valid values are `source`, `target`.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `db2`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
	EngineName pulumi.StringOutput `pulumi:"engineName"`
	// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html).
	ExtraConnectionAttributes pulumi.StringOutput `pulumi:"extraConnectionAttributes"`
	// Configuration block for Kafka settings. See below.
	KafkaSettings EndpointKafkaSettingsPtrOutput `pulumi:"kafkaSettings"`
	// Configuration block for Kinesis settings. See below.
	KinesisSettings EndpointKinesisSettingsPtrOutput `pulumi:"kinesisSettings"`
	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`
	// Configuration block for MongoDB settings. See below.
	MongodbSettings EndpointMongodbSettingsPtrOutput `pulumi:"mongodbSettings"`
	// Password to be used to login to the endpoint database.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port          pulumi.IntPtrOutput            `pulumi:"port"`
	RedisSettings EndpointRedisSettingsPtrOutput `pulumi:"redisSettings"`
	// Configuration block for Redshift settings. See below.
	RedshiftSettings EndpointRedshiftSettingsOutput `pulumi:"redshiftSettings"`
	// Configuration block for S3 settings. See below.
	S3Settings EndpointS3SettingsPtrOutput `pulumi:"s3Settings"`
	// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
	SecretsManagerAccessRoleArn pulumi.StringPtrOutput `pulumi:"secretsManagerAccessRoleArn"`
	// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only for `engineName` as `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift` or `sqlserver`.
	SecretsManagerArn pulumi.StringPtrOutput `pulumi:"secretsManagerArn"`
	// Fully qualified domain name of the endpoint.
	ServerName pulumi.StringPtrOutput `pulumi:"serverName"`
	// ARN used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole pulumi.StringPtrOutput `pulumi:"serviceAccessRole"`
	// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
	SslMode pulumi.StringOutput `pulumi:"sslMode"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// User name to be used to login to the endpoint database.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.EndpointType == nil {
		return nil, errors.New("invalid value for required argument 'EndpointType'")
	}
	if args.EngineName == nil {
		return nil, errors.New("invalid value for required argument 'EngineName'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	var resource Endpoint
	err := ctx.RegisterResource("aws:dms/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("aws:dms/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// ARN for the certificate.
	CertificateArn *string `pulumi:"certificateArn"`
	// Name of the endpoint database.
	DatabaseName *string `pulumi:"databaseName"`
	// Configuration block for OpenSearch settings. See below.
	ElasticsearchSettings *EndpointElasticsearchSettings `pulumi:"elasticsearchSettings"`
	// ARN for the endpoint.
	EndpointArn *string `pulumi:"endpointArn"`
	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId *string `pulumi:"endpointId"`
	// Type of endpoint. Valid values are `source`, `target`.
	EndpointType *string `pulumi:"endpointType"`
	// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `db2`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
	EngineName *string `pulumi:"engineName"`
	// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html).
	ExtraConnectionAttributes *string `pulumi:"extraConnectionAttributes"`
	// Configuration block for Kafka settings. See below.
	KafkaSettings *EndpointKafkaSettings `pulumi:"kafkaSettings"`
	// Configuration block for Kinesis settings. See below.
	KinesisSettings *EndpointKinesisSettings `pulumi:"kinesisSettings"`
	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Configuration block for MongoDB settings. See below.
	MongodbSettings *EndpointMongodbSettings `pulumi:"mongodbSettings"`
	// Password to be used to login to the endpoint database.
	Password *string `pulumi:"password"`
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port          *int                   `pulumi:"port"`
	RedisSettings *EndpointRedisSettings `pulumi:"redisSettings"`
	// Configuration block for Redshift settings. See below.
	RedshiftSettings *EndpointRedshiftSettings `pulumi:"redshiftSettings"`
	// Configuration block for S3 settings. See below.
	S3Settings *EndpointS3Settings `pulumi:"s3Settings"`
	// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
	SecretsManagerAccessRoleArn *string `pulumi:"secretsManagerAccessRoleArn"`
	// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only for `engineName` as `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift` or `sqlserver`.
	SecretsManagerArn *string `pulumi:"secretsManagerArn"`
	// Fully qualified domain name of the endpoint.
	ServerName *string `pulumi:"serverName"`
	// ARN used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole *string `pulumi:"serviceAccessRole"`
	// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
	SslMode *string `pulumi:"sslMode"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// User name to be used to login to the endpoint database.
	Username *string `pulumi:"username"`
}

type EndpointState struct {
	// ARN for the certificate.
	CertificateArn pulumi.StringPtrInput
	// Name of the endpoint database.
	DatabaseName pulumi.StringPtrInput
	// Configuration block for OpenSearch settings. See below.
	ElasticsearchSettings EndpointElasticsearchSettingsPtrInput
	// ARN for the endpoint.
	EndpointArn pulumi.StringPtrInput
	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId pulumi.StringPtrInput
	// Type of endpoint. Valid values are `source`, `target`.
	EndpointType pulumi.StringPtrInput
	// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `db2`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
	EngineName pulumi.StringPtrInput
	// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html).
	ExtraConnectionAttributes pulumi.StringPtrInput
	// Configuration block for Kafka settings. See below.
	KafkaSettings EndpointKafkaSettingsPtrInput
	// Configuration block for Kinesis settings. See below.
	KinesisSettings EndpointKinesisSettingsPtrInput
	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringPtrInput
	// Configuration block for MongoDB settings. See below.
	MongodbSettings EndpointMongodbSettingsPtrInput
	// Password to be used to login to the endpoint database.
	Password pulumi.StringPtrInput
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port          pulumi.IntPtrInput
	RedisSettings EndpointRedisSettingsPtrInput
	// Configuration block for Redshift settings. See below.
	RedshiftSettings EndpointRedshiftSettingsPtrInput
	// Configuration block for S3 settings. See below.
	S3Settings EndpointS3SettingsPtrInput
	// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
	SecretsManagerAccessRoleArn pulumi.StringPtrInput
	// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only for `engineName` as `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift` or `sqlserver`.
	SecretsManagerArn pulumi.StringPtrInput
	// Fully qualified domain name of the endpoint.
	ServerName pulumi.StringPtrInput
	// ARN used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole pulumi.StringPtrInput
	// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
	SslMode pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// User name to be used to login to the endpoint database.
	Username pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// ARN for the certificate.
	CertificateArn *string `pulumi:"certificateArn"`
	// Name of the endpoint database.
	DatabaseName *string `pulumi:"databaseName"`
	// Configuration block for OpenSearch settings. See below.
	ElasticsearchSettings *EndpointElasticsearchSettings `pulumi:"elasticsearchSettings"`
	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId string `pulumi:"endpointId"`
	// Type of endpoint. Valid values are `source`, `target`.
	EndpointType string `pulumi:"endpointType"`
	// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `db2`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
	EngineName string `pulumi:"engineName"`
	// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html).
	ExtraConnectionAttributes *string `pulumi:"extraConnectionAttributes"`
	// Configuration block for Kafka settings. See below.
	KafkaSettings *EndpointKafkaSettings `pulumi:"kafkaSettings"`
	// Configuration block for Kinesis settings. See below.
	KinesisSettings *EndpointKinesisSettings `pulumi:"kinesisSettings"`
	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Configuration block for MongoDB settings. See below.
	MongodbSettings *EndpointMongodbSettings `pulumi:"mongodbSettings"`
	// Password to be used to login to the endpoint database.
	Password *string `pulumi:"password"`
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port          *int                   `pulumi:"port"`
	RedisSettings *EndpointRedisSettings `pulumi:"redisSettings"`
	// Configuration block for Redshift settings. See below.
	RedshiftSettings *EndpointRedshiftSettings `pulumi:"redshiftSettings"`
	// Configuration block for S3 settings. See below.
	S3Settings *EndpointS3Settings `pulumi:"s3Settings"`
	// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
	SecretsManagerAccessRoleArn *string `pulumi:"secretsManagerAccessRoleArn"`
	// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only for `engineName` as `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift` or `sqlserver`.
	SecretsManagerArn *string `pulumi:"secretsManagerArn"`
	// Fully qualified domain name of the endpoint.
	ServerName *string `pulumi:"serverName"`
	// ARN used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole *string `pulumi:"serviceAccessRole"`
	// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
	SslMode *string `pulumi:"sslMode"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// User name to be used to login to the endpoint database.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// ARN for the certificate.
	CertificateArn pulumi.StringPtrInput
	// Name of the endpoint database.
	DatabaseName pulumi.StringPtrInput
	// Configuration block for OpenSearch settings. See below.
	ElasticsearchSettings EndpointElasticsearchSettingsPtrInput
	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId pulumi.StringInput
	// Type of endpoint. Valid values are `source`, `target`.
	EndpointType pulumi.StringInput
	// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `db2`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
	EngineName pulumi.StringInput
	// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html).
	ExtraConnectionAttributes pulumi.StringPtrInput
	// Configuration block for Kafka settings. See below.
	KafkaSettings EndpointKafkaSettingsPtrInput
	// Configuration block for Kinesis settings. See below.
	KinesisSettings EndpointKinesisSettingsPtrInput
	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringPtrInput
	// Configuration block for MongoDB settings. See below.
	MongodbSettings EndpointMongodbSettingsPtrInput
	// Password to be used to login to the endpoint database.
	Password pulumi.StringPtrInput
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port          pulumi.IntPtrInput
	RedisSettings EndpointRedisSettingsPtrInput
	// Configuration block for Redshift settings. See below.
	RedshiftSettings EndpointRedshiftSettingsPtrInput
	// Configuration block for S3 settings. See below.
	S3Settings EndpointS3SettingsPtrInput
	// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
	SecretsManagerAccessRoleArn pulumi.StringPtrInput
	// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only for `engineName` as `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift` or `sqlserver`.
	SecretsManagerArn pulumi.StringPtrInput
	// Fully qualified domain name of the endpoint.
	ServerName pulumi.StringPtrInput
	// ARN used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole pulumi.StringPtrInput
	// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
	SslMode pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// User name to be used to login to the endpoint database.
	Username pulumi.StringPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//	EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//	EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

// ARN for the certificate.
func (o EndpointOutput) CertificateArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.CertificateArn }).(pulumi.StringOutput)
}

// Name of the endpoint database.
func (o EndpointOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// Configuration block for OpenSearch settings. See below.
func (o EndpointOutput) ElasticsearchSettings() EndpointElasticsearchSettingsPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointElasticsearchSettingsPtrOutput { return v.ElasticsearchSettings }).(EndpointElasticsearchSettingsPtrOutput)
}

// ARN for the endpoint.
func (o EndpointOutput) EndpointArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointArn }).(pulumi.StringOutput)
}

// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
func (o EndpointOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// Type of endpoint. Valid values are `source`, `target`.
func (o EndpointOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointType }).(pulumi.StringOutput)
}

// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `db2`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
func (o EndpointOutput) EngineName() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EngineName }).(pulumi.StringOutput)
}

// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html).
func (o EndpointOutput) ExtraConnectionAttributes() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ExtraConnectionAttributes }).(pulumi.StringOutput)
}

// Configuration block for Kafka settings. See below.
func (o EndpointOutput) KafkaSettings() EndpointKafkaSettingsPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointKafkaSettingsPtrOutput { return v.KafkaSettings }).(EndpointKafkaSettingsPtrOutput)
}

// Configuration block for Kinesis settings. See below.
func (o EndpointOutput) KinesisSettings() EndpointKinesisSettingsPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointKinesisSettingsPtrOutput { return v.KinesisSettings }).(EndpointKinesisSettingsPtrOutput)
}

// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
func (o EndpointOutput) KmsKeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.KmsKeyArn }).(pulumi.StringOutput)
}

// Configuration block for MongoDB settings. See below.
func (o EndpointOutput) MongodbSettings() EndpointMongodbSettingsPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointMongodbSettingsPtrOutput { return v.MongodbSettings }).(EndpointMongodbSettingsPtrOutput)
}

// Password to be used to login to the endpoint database.
func (o EndpointOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Transmission Control Protocol (TCP) port for the endpoint.
func (o EndpointOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o EndpointOutput) RedisSettings() EndpointRedisSettingsPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointRedisSettingsPtrOutput { return v.RedisSettings }).(EndpointRedisSettingsPtrOutput)
}

// Configuration block for Redshift settings. See below.
func (o EndpointOutput) RedshiftSettings() EndpointRedshiftSettingsOutput {
	return o.ApplyT(func(v *Endpoint) EndpointRedshiftSettingsOutput { return v.RedshiftSettings }).(EndpointRedshiftSettingsOutput)
}

// Configuration block for S3 settings. See below.
func (o EndpointOutput) S3Settings() EndpointS3SettingsPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointS3SettingsPtrOutput { return v.S3Settings }).(EndpointS3SettingsPtrOutput)
}

// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
func (o EndpointOutput) SecretsManagerAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.SecretsManagerAccessRoleArn }).(pulumi.StringPtrOutput)
}

// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only for `engineName` as `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift` or `sqlserver`.
func (o EndpointOutput) SecretsManagerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.SecretsManagerArn }).(pulumi.StringPtrOutput)
}

// Fully qualified domain name of the endpoint.
func (o EndpointOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.ServerName }).(pulumi.StringPtrOutput)
}

// ARN used by the service access IAM role for dynamodb endpoints.
func (o EndpointOutput) ServiceAccessRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.ServiceAccessRole }).(pulumi.StringPtrOutput)
}

// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
func (o EndpointOutput) SslMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.SslMode }).(pulumi.StringOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o EndpointOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// User name to be used to login to the endpoint database.
func (o EndpointOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
