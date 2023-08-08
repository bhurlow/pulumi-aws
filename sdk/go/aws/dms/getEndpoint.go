// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for managing an AWS DMS (Database Migration) Endpoint.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dms.LookupEndpoint(ctx, &dms.LookupEndpointArgs{
//				EndpointId: "test_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointResult
	err := ctx.Invoke("aws:dms/getEndpoint:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpoint.
type LookupEndpointArgs struct {
	ElasticsearchSettings []GetEndpointElasticsearchSetting `pulumi:"elasticsearchSettings"`
	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId                string                      `pulumi:"endpointId"`
	ExtraConnectionAttributes *string                     `pulumi:"extraConnectionAttributes"`
	KafkaSettings             []GetEndpointKafkaSetting   `pulumi:"kafkaSettings"`
	MongodbSettings           []GetEndpointMongodbSetting `pulumi:"mongodbSettings"`
	Tags                      map[string]string           `pulumi:"tags"`
}

// A collection of values returned by getEndpoint.
type LookupEndpointResult struct {
	CertificateArn            string                            `pulumi:"certificateArn"`
	DatabaseName              string                            `pulumi:"databaseName"`
	ElasticsearchSettings     []GetEndpointElasticsearchSetting `pulumi:"elasticsearchSettings"`
	EndpointArn               string                            `pulumi:"endpointArn"`
	EndpointId                string                            `pulumi:"endpointId"`
	EndpointType              string                            `pulumi:"endpointType"`
	EngineName                string                            `pulumi:"engineName"`
	ExtraConnectionAttributes *string                           `pulumi:"extraConnectionAttributes"`
	// The provider-assigned unique ID for this managed resource.
	Id                          string                       `pulumi:"id"`
	KafkaSettings               []GetEndpointKafkaSetting    `pulumi:"kafkaSettings"`
	KinesisSettings             []GetEndpointKinesisSetting  `pulumi:"kinesisSettings"`
	KmsKeyArn                   string                       `pulumi:"kmsKeyArn"`
	MongodbSettings             []GetEndpointMongodbSetting  `pulumi:"mongodbSettings"`
	Password                    string                       `pulumi:"password"`
	Port                        int                          `pulumi:"port"`
	RedisSettings               []GetEndpointRedisSetting    `pulumi:"redisSettings"`
	RedshiftSettings            []GetEndpointRedshiftSetting `pulumi:"redshiftSettings"`
	S3Settings                  []GetEndpointS3Setting       `pulumi:"s3Settings"`
	SecretsManagerAccessRoleArn string                       `pulumi:"secretsManagerAccessRoleArn"`
	SecretsManagerArn           string                       `pulumi:"secretsManagerArn"`
	ServerName                  string                       `pulumi:"serverName"`
	ServiceAccessRole           string                       `pulumi:"serviceAccessRole"`
	SslMode                     string                       `pulumi:"sslMode"`
	Tags                        map[string]string            `pulumi:"tags"`
	Username                    string                       `pulumi:"username"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

// A collection of arguments for invoking getEndpoint.
type LookupEndpointOutputArgs struct {
	ElasticsearchSettings GetEndpointElasticsearchSettingArrayInput `pulumi:"elasticsearchSettings"`
	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId                pulumi.StringInput                  `pulumi:"endpointId"`
	ExtraConnectionAttributes pulumi.StringPtrInput               `pulumi:"extraConnectionAttributes"`
	KafkaSettings             GetEndpointKafkaSettingArrayInput   `pulumi:"kafkaSettings"`
	MongodbSettings           GetEndpointMongodbSettingArrayInput `pulumi:"mongodbSettings"`
	Tags                      pulumi.StringMapInput               `pulumi:"tags"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getEndpoint.
type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) CertificateArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.CertificateArn }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) ElasticsearchSettings() GetEndpointElasticsearchSettingArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GetEndpointElasticsearchSetting { return v.ElasticsearchSettings }).(GetEndpointElasticsearchSettingArrayOutput)
}

func (o LookupEndpointResultOutput) EndpointArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.EndpointArn }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.EndpointId }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) EngineName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.EngineName }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) ExtraConnectionAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.ExtraConnectionAttributes }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) KafkaSettings() GetEndpointKafkaSettingArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GetEndpointKafkaSetting { return v.KafkaSettings }).(GetEndpointKafkaSettingArrayOutput)
}

func (o LookupEndpointResultOutput) KinesisSettings() GetEndpointKinesisSettingArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GetEndpointKinesisSetting { return v.KinesisSettings }).(GetEndpointKinesisSettingArrayOutput)
}

func (o LookupEndpointResultOutput) KmsKeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.KmsKeyArn }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) MongodbSettings() GetEndpointMongodbSettingArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GetEndpointMongodbSetting { return v.MongodbSettings }).(GetEndpointMongodbSettingArrayOutput)
}

func (o LookupEndpointResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupEndpointResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupEndpointResultOutput) RedisSettings() GetEndpointRedisSettingArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GetEndpointRedisSetting { return v.RedisSettings }).(GetEndpointRedisSettingArrayOutput)
}

func (o LookupEndpointResultOutput) RedshiftSettings() GetEndpointRedshiftSettingArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GetEndpointRedshiftSetting { return v.RedshiftSettings }).(GetEndpointRedshiftSettingArrayOutput)
}

func (o LookupEndpointResultOutput) S3Settings() GetEndpointS3SettingArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GetEndpointS3Setting { return v.S3Settings }).(GetEndpointS3SettingArrayOutput)
}

func (o LookupEndpointResultOutput) SecretsManagerAccessRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.SecretsManagerAccessRoleArn }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) SecretsManagerArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.SecretsManagerArn }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) ServiceAccessRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.ServiceAccessRole }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) SslMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.SslMode }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEndpointResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}
