// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a version of a CloudFormation Type.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudformation.NewCloudFormationType(ctx, "example", &cloudformation.CloudFormationTypeArgs{
//				SchemaHandlerPackage: pulumi.String(fmt.Sprintf("s3://%v/%v", aws_s3_object.Example.Bucket, aws_s3_object.Example.Key)),
//				Type:                 pulumi.String("RESOURCE"),
//				TypeName:             pulumi.String("ExampleCompany::ExampleService::ExampleResource"),
//				LoggingConfig: &cloudformation.CloudFormationTypeLoggingConfigArgs{
//					LogGroupName: pulumi.Any(aws_cloudwatch_log_group.Example.Name),
//					LogRoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
//				},
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
// terraform import {
//
//	to = aws_cloudformation_type.example
//
//	id = "arn:aws:cloudformation:us-east-1:123456789012:type/resource/ExampleCompany-ExampleService-ExampleType/1" } Using `pulumi import`, import `aws_cloudformation_type` using the type version Amazon Resource Name (ARN). For exampleconsole % pulumi import aws_cloudformation_type.example arn:aws:cloudformation:us-east-1:123456789012:type/resource/ExampleCompany-ExampleService-ExampleType/1
type CloudFormationType struct {
	pulumi.CustomResourceState

	// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type version. See also `typeArn`.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Identifier of the CloudFormation Type default version.
	DefaultVersionId pulumi.StringOutput `pulumi:"defaultVersionId"`
	// Deprecation status of the version.
	DeprecatedStatus pulumi.StringOutput `pulumi:"deprecatedStatus"`
	// Description of the version.
	Description pulumi.StringOutput `pulumi:"description"`
	// URL of the documentation for the CloudFormation Type.
	DocumentationUrl pulumi.StringOutput `pulumi:"documentationUrl"`
	// Amazon Resource Name (ARN) of the IAM Role for CloudFormation to assume when invoking the extension. If your extension calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the extension handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the extension handler, thereby supplying your extension with the appropriate credentials.
	ExecutionRoleArn pulumi.StringPtrOutput `pulumi:"executionRoleArn"`
	// Whether the CloudFormation Type version is the default version.
	IsDefaultVersion pulumi.BoolOutput `pulumi:"isDefaultVersion"`
	// Configuration block containing logging configuration.
	LoggingConfig CloudFormationTypeLoggingConfigPtrOutput `pulumi:"loggingConfig"`
	// Provisioning behavior of the CloudFormation Type.
	ProvisioningType pulumi.StringOutput `pulumi:"provisioningType"`
	// JSON document of the CloudFormation Type schema.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// URL to the S3 bucket containing the extension project package that contains the necessary files for the extension you want to register. Must begin with `s3://` or `https://`. For example, `s3://example-bucket/example-object`.
	SchemaHandlerPackage pulumi.StringOutput `pulumi:"schemaHandlerPackage"`
	// URL of the source code for the CloudFormation Type.
	SourceUrl pulumi.StringOutput `pulumi:"sourceUrl"`
	// CloudFormation Registry Type. For example, `RESOURCE` or `MODULE`.
	Type pulumi.StringOutput `pulumi:"type"`
	// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type. See also `arn`.
	TypeArn pulumi.StringOutput `pulumi:"typeArn"`
	// CloudFormation Type name. For example, `ExampleCompany::ExampleService::ExampleResource`.
	TypeName pulumi.StringOutput `pulumi:"typeName"`
	// (Optional) Identifier of the CloudFormation Type version.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
	// Scope of the CloudFormation Type.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewCloudFormationType registers a new resource with the given unique name, arguments, and options.
func NewCloudFormationType(ctx *pulumi.Context,
	name string, args *CloudFormationTypeArgs, opts ...pulumi.ResourceOption) (*CloudFormationType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SchemaHandlerPackage == nil {
		return nil, errors.New("invalid value for required argument 'SchemaHandlerPackage'")
	}
	if args.TypeName == nil {
		return nil, errors.New("invalid value for required argument 'TypeName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudFormationType
	err := ctx.RegisterResource("aws:cloudformation/cloudFormationType:CloudFormationType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudFormationType gets an existing CloudFormationType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudFormationType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudFormationTypeState, opts ...pulumi.ResourceOption) (*CloudFormationType, error) {
	var resource CloudFormationType
	err := ctx.ReadResource("aws:cloudformation/cloudFormationType:CloudFormationType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudFormationType resources.
type cloudFormationTypeState struct {
	// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type version. See also `typeArn`.
	Arn *string `pulumi:"arn"`
	// Identifier of the CloudFormation Type default version.
	DefaultVersionId *string `pulumi:"defaultVersionId"`
	// Deprecation status of the version.
	DeprecatedStatus *string `pulumi:"deprecatedStatus"`
	// Description of the version.
	Description *string `pulumi:"description"`
	// URL of the documentation for the CloudFormation Type.
	DocumentationUrl *string `pulumi:"documentationUrl"`
	// Amazon Resource Name (ARN) of the IAM Role for CloudFormation to assume when invoking the extension. If your extension calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the extension handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the extension handler, thereby supplying your extension with the appropriate credentials.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// Whether the CloudFormation Type version is the default version.
	IsDefaultVersion *bool `pulumi:"isDefaultVersion"`
	// Configuration block containing logging configuration.
	LoggingConfig *CloudFormationTypeLoggingConfig `pulumi:"loggingConfig"`
	// Provisioning behavior of the CloudFormation Type.
	ProvisioningType *string `pulumi:"provisioningType"`
	// JSON document of the CloudFormation Type schema.
	Schema *string `pulumi:"schema"`
	// URL to the S3 bucket containing the extension project package that contains the necessary files for the extension you want to register. Must begin with `s3://` or `https://`. For example, `s3://example-bucket/example-object`.
	SchemaHandlerPackage *string `pulumi:"schemaHandlerPackage"`
	// URL of the source code for the CloudFormation Type.
	SourceUrl *string `pulumi:"sourceUrl"`
	// CloudFormation Registry Type. For example, `RESOURCE` or `MODULE`.
	Type *string `pulumi:"type"`
	// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type. See also `arn`.
	TypeArn *string `pulumi:"typeArn"`
	// CloudFormation Type name. For example, `ExampleCompany::ExampleService::ExampleResource`.
	TypeName *string `pulumi:"typeName"`
	// (Optional) Identifier of the CloudFormation Type version.
	VersionId *string `pulumi:"versionId"`
	// Scope of the CloudFormation Type.
	Visibility *string `pulumi:"visibility"`
}

type CloudFormationTypeState struct {
	// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type version. See also `typeArn`.
	Arn pulumi.StringPtrInput
	// Identifier of the CloudFormation Type default version.
	DefaultVersionId pulumi.StringPtrInput
	// Deprecation status of the version.
	DeprecatedStatus pulumi.StringPtrInput
	// Description of the version.
	Description pulumi.StringPtrInput
	// URL of the documentation for the CloudFormation Type.
	DocumentationUrl pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role for CloudFormation to assume when invoking the extension. If your extension calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the extension handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the extension handler, thereby supplying your extension with the appropriate credentials.
	ExecutionRoleArn pulumi.StringPtrInput
	// Whether the CloudFormation Type version is the default version.
	IsDefaultVersion pulumi.BoolPtrInput
	// Configuration block containing logging configuration.
	LoggingConfig CloudFormationTypeLoggingConfigPtrInput
	// Provisioning behavior of the CloudFormation Type.
	ProvisioningType pulumi.StringPtrInput
	// JSON document of the CloudFormation Type schema.
	Schema pulumi.StringPtrInput
	// URL to the S3 bucket containing the extension project package that contains the necessary files for the extension you want to register. Must begin with `s3://` or `https://`. For example, `s3://example-bucket/example-object`.
	SchemaHandlerPackage pulumi.StringPtrInput
	// URL of the source code for the CloudFormation Type.
	SourceUrl pulumi.StringPtrInput
	// CloudFormation Registry Type. For example, `RESOURCE` or `MODULE`.
	Type pulumi.StringPtrInput
	// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type. See also `arn`.
	TypeArn pulumi.StringPtrInput
	// CloudFormation Type name. For example, `ExampleCompany::ExampleService::ExampleResource`.
	TypeName pulumi.StringPtrInput
	// (Optional) Identifier of the CloudFormation Type version.
	VersionId pulumi.StringPtrInput
	// Scope of the CloudFormation Type.
	Visibility pulumi.StringPtrInput
}

func (CloudFormationTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudFormationTypeState)(nil)).Elem()
}

type cloudFormationTypeArgs struct {
	// Amazon Resource Name (ARN) of the IAM Role for CloudFormation to assume when invoking the extension. If your extension calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the extension handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the extension handler, thereby supplying your extension with the appropriate credentials.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// Configuration block containing logging configuration.
	LoggingConfig *CloudFormationTypeLoggingConfig `pulumi:"loggingConfig"`
	// URL to the S3 bucket containing the extension project package that contains the necessary files for the extension you want to register. Must begin with `s3://` or `https://`. For example, `s3://example-bucket/example-object`.
	SchemaHandlerPackage string `pulumi:"schemaHandlerPackage"`
	// CloudFormation Registry Type. For example, `RESOURCE` or `MODULE`.
	Type *string `pulumi:"type"`
	// CloudFormation Type name. For example, `ExampleCompany::ExampleService::ExampleResource`.
	TypeName string `pulumi:"typeName"`
}

// The set of arguments for constructing a CloudFormationType resource.
type CloudFormationTypeArgs struct {
	// Amazon Resource Name (ARN) of the IAM Role for CloudFormation to assume when invoking the extension. If your extension calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the extension handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the extension handler, thereby supplying your extension with the appropriate credentials.
	ExecutionRoleArn pulumi.StringPtrInput
	// Configuration block containing logging configuration.
	LoggingConfig CloudFormationTypeLoggingConfigPtrInput
	// URL to the S3 bucket containing the extension project package that contains the necessary files for the extension you want to register. Must begin with `s3://` or `https://`. For example, `s3://example-bucket/example-object`.
	SchemaHandlerPackage pulumi.StringInput
	// CloudFormation Registry Type. For example, `RESOURCE` or `MODULE`.
	Type pulumi.StringPtrInput
	// CloudFormation Type name. For example, `ExampleCompany::ExampleService::ExampleResource`.
	TypeName pulumi.StringInput
}

func (CloudFormationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudFormationTypeArgs)(nil)).Elem()
}

type CloudFormationTypeInput interface {
	pulumi.Input

	ToCloudFormationTypeOutput() CloudFormationTypeOutput
	ToCloudFormationTypeOutputWithContext(ctx context.Context) CloudFormationTypeOutput
}

func (*CloudFormationType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudFormationType)(nil)).Elem()
}

func (i *CloudFormationType) ToCloudFormationTypeOutput() CloudFormationTypeOutput {
	return i.ToCloudFormationTypeOutputWithContext(context.Background())
}

func (i *CloudFormationType) ToCloudFormationTypeOutputWithContext(ctx context.Context) CloudFormationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationTypeOutput)
}

// CloudFormationTypeArrayInput is an input type that accepts CloudFormationTypeArray and CloudFormationTypeArrayOutput values.
// You can construct a concrete instance of `CloudFormationTypeArrayInput` via:
//
//	CloudFormationTypeArray{ CloudFormationTypeArgs{...} }
type CloudFormationTypeArrayInput interface {
	pulumi.Input

	ToCloudFormationTypeArrayOutput() CloudFormationTypeArrayOutput
	ToCloudFormationTypeArrayOutputWithContext(context.Context) CloudFormationTypeArrayOutput
}

type CloudFormationTypeArray []CloudFormationTypeInput

func (CloudFormationTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudFormationType)(nil)).Elem()
}

func (i CloudFormationTypeArray) ToCloudFormationTypeArrayOutput() CloudFormationTypeArrayOutput {
	return i.ToCloudFormationTypeArrayOutputWithContext(context.Background())
}

func (i CloudFormationTypeArray) ToCloudFormationTypeArrayOutputWithContext(ctx context.Context) CloudFormationTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationTypeArrayOutput)
}

// CloudFormationTypeMapInput is an input type that accepts CloudFormationTypeMap and CloudFormationTypeMapOutput values.
// You can construct a concrete instance of `CloudFormationTypeMapInput` via:
//
//	CloudFormationTypeMap{ "key": CloudFormationTypeArgs{...} }
type CloudFormationTypeMapInput interface {
	pulumi.Input

	ToCloudFormationTypeMapOutput() CloudFormationTypeMapOutput
	ToCloudFormationTypeMapOutputWithContext(context.Context) CloudFormationTypeMapOutput
}

type CloudFormationTypeMap map[string]CloudFormationTypeInput

func (CloudFormationTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudFormationType)(nil)).Elem()
}

func (i CloudFormationTypeMap) ToCloudFormationTypeMapOutput() CloudFormationTypeMapOutput {
	return i.ToCloudFormationTypeMapOutputWithContext(context.Background())
}

func (i CloudFormationTypeMap) ToCloudFormationTypeMapOutputWithContext(ctx context.Context) CloudFormationTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationTypeMapOutput)
}

type CloudFormationTypeOutput struct{ *pulumi.OutputState }

func (CloudFormationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudFormationType)(nil)).Elem()
}

func (o CloudFormationTypeOutput) ToCloudFormationTypeOutput() CloudFormationTypeOutput {
	return o
}

func (o CloudFormationTypeOutput) ToCloudFormationTypeOutputWithContext(ctx context.Context) CloudFormationTypeOutput {
	return o
}

// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type version. See also `typeArn`.
func (o CloudFormationTypeOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Identifier of the CloudFormation Type default version.
func (o CloudFormationTypeOutput) DefaultVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.DefaultVersionId }).(pulumi.StringOutput)
}

// Deprecation status of the version.
func (o CloudFormationTypeOutput) DeprecatedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.DeprecatedStatus }).(pulumi.StringOutput)
}

// Description of the version.
func (o CloudFormationTypeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// URL of the documentation for the CloudFormation Type.
func (o CloudFormationTypeOutput) DocumentationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.DocumentationUrl }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the IAM Role for CloudFormation to assume when invoking the extension. If your extension calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. When CloudFormation needs to invoke the extension handler, CloudFormation assumes this execution role to create a temporary session token, which it then passes to the extension handler, thereby supplying your extension with the appropriate credentials.
func (o CloudFormationTypeOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringPtrOutput { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

// Whether the CloudFormation Type version is the default version.
func (o CloudFormationTypeOutput) IsDefaultVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.BoolOutput { return v.IsDefaultVersion }).(pulumi.BoolOutput)
}

// Configuration block containing logging configuration.
func (o CloudFormationTypeOutput) LoggingConfig() CloudFormationTypeLoggingConfigPtrOutput {
	return o.ApplyT(func(v *CloudFormationType) CloudFormationTypeLoggingConfigPtrOutput { return v.LoggingConfig }).(CloudFormationTypeLoggingConfigPtrOutput)
}

// Provisioning behavior of the CloudFormation Type.
func (o CloudFormationTypeOutput) ProvisioningType() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.ProvisioningType }).(pulumi.StringOutput)
}

// JSON document of the CloudFormation Type schema.
func (o CloudFormationTypeOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// URL to the S3 bucket containing the extension project package that contains the necessary files for the extension you want to register. Must begin with `s3://` or `https://`. For example, `s3://example-bucket/example-object`.
func (o CloudFormationTypeOutput) SchemaHandlerPackage() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.SchemaHandlerPackage }).(pulumi.StringOutput)
}

// URL of the source code for the CloudFormation Type.
func (o CloudFormationTypeOutput) SourceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.SourceUrl }).(pulumi.StringOutput)
}

// CloudFormation Registry Type. For example, `RESOURCE` or `MODULE`.
func (o CloudFormationTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// (Optional) Amazon Resource Name (ARN) of the CloudFormation Type. See also `arn`.
func (o CloudFormationTypeOutput) TypeArn() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.TypeArn }).(pulumi.StringOutput)
}

// CloudFormation Type name. For example, `ExampleCompany::ExampleService::ExampleResource`.
func (o CloudFormationTypeOutput) TypeName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.TypeName }).(pulumi.StringOutput)
}

// (Optional) Identifier of the CloudFormation Type version.
func (o CloudFormationTypeOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

// Scope of the CloudFormation Type.
func (o CloudFormationTypeOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationType) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type CloudFormationTypeArrayOutput struct{ *pulumi.OutputState }

func (CloudFormationTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudFormationType)(nil)).Elem()
}

func (o CloudFormationTypeArrayOutput) ToCloudFormationTypeArrayOutput() CloudFormationTypeArrayOutput {
	return o
}

func (o CloudFormationTypeArrayOutput) ToCloudFormationTypeArrayOutputWithContext(ctx context.Context) CloudFormationTypeArrayOutput {
	return o
}

func (o CloudFormationTypeArrayOutput) Index(i pulumi.IntInput) CloudFormationTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudFormationType {
		return vs[0].([]*CloudFormationType)[vs[1].(int)]
	}).(CloudFormationTypeOutput)
}

type CloudFormationTypeMapOutput struct{ *pulumi.OutputState }

func (CloudFormationTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudFormationType)(nil)).Elem()
}

func (o CloudFormationTypeMapOutput) ToCloudFormationTypeMapOutput() CloudFormationTypeMapOutput {
	return o
}

func (o CloudFormationTypeMapOutput) ToCloudFormationTypeMapOutputWithContext(ctx context.Context) CloudFormationTypeMapOutput {
	return o
}

func (o CloudFormationTypeMapOutput) MapIndex(k pulumi.StringInput) CloudFormationTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudFormationType {
		return vs[0].(map[string]*CloudFormationType)[vs[1].(string)]
	}).(CloudFormationTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudFormationTypeInput)(nil)).Elem(), &CloudFormationType{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudFormationTypeArrayInput)(nil)).Elem(), CloudFormationTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudFormationTypeMapInput)(nil)).Elem(), CloudFormationTypeMap{})
	pulumi.RegisterOutputType(CloudFormationTypeOutput{})
	pulumi.RegisterOutputType(CloudFormationTypeArrayOutput{})
	pulumi.RegisterOutputType(CloudFormationTypeMapOutput{})
}
