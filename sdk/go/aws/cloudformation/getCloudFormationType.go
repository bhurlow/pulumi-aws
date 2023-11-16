// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a CloudFormation Type.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudformation.LookupCloudFormationType(ctx, &cloudformation.LookupCloudFormationTypeArgs{
//				Type:     pulumi.StringRef("RESOURCE"),
//				TypeName: pulumi.StringRef("AWS::Athena::WorkGroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCloudFormationType(ctx *pulumi.Context, args *LookupCloudFormationTypeArgs, opts ...pulumi.InvokeOption) (*LookupCloudFormationTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCloudFormationTypeResult
	err := ctx.Invoke("aws:cloudformation/getCloudFormationType:getCloudFormationType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudFormationType.
type LookupCloudFormationTypeArgs struct {
	// ARN of the CloudFormation Type. For example, `arn:aws:cloudformation:us-west-2::type/resource/AWS-EC2-VPC`.
	Arn *string `pulumi:"arn"`
	// CloudFormation Registry Type. For example, `RESOURCE`.
	Type *string `pulumi:"type"`
	// CloudFormation Type name. For example, `AWS::EC2::VPC`.
	TypeName *string `pulumi:"typeName"`
	// Identifier of the CloudFormation Type version.
	VersionId *string `pulumi:"versionId"`
}

// A collection of values returned by getCloudFormationType.
type LookupCloudFormationTypeResult struct {
	Arn string `pulumi:"arn"`
	// Identifier of the CloudFormation Type default version.
	DefaultVersionId string `pulumi:"defaultVersionId"`
	// Deprecation status of the CloudFormation Type.
	DeprecatedStatus string `pulumi:"deprecatedStatus"`
	// Description of the CloudFormation Type.
	Description string `pulumi:"description"`
	// URL of the documentation for the CloudFormation Type.
	DocumentationUrl string `pulumi:"documentationUrl"`
	// ARN of the IAM Role used to register the CloudFormation Type.
	ExecutionRoleArn string `pulumi:"executionRoleArn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether the CloudFormation Type version is the default version.
	IsDefaultVersion bool `pulumi:"isDefaultVersion"`
	// List of objects containing logging configuration.
	LoggingConfigs []GetCloudFormationTypeLoggingConfig `pulumi:"loggingConfigs"`
	// Provisioning behavior of the CloudFormation Type.
	ProvisioningType string `pulumi:"provisioningType"`
	// JSON document of the CloudFormation Type schema.
	Schema string `pulumi:"schema"`
	// URL of the source code for the CloudFormation Type.
	SourceUrl string  `pulumi:"sourceUrl"`
	Type      string  `pulumi:"type"`
	TypeArn   string  `pulumi:"typeArn"`
	TypeName  string  `pulumi:"typeName"`
	VersionId *string `pulumi:"versionId"`
	// Scope of the CloudFormation Type.
	Visibility string `pulumi:"visibility"`
}

func LookupCloudFormationTypeOutput(ctx *pulumi.Context, args LookupCloudFormationTypeOutputArgs, opts ...pulumi.InvokeOption) LookupCloudFormationTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudFormationTypeResult, error) {
			args := v.(LookupCloudFormationTypeArgs)
			r, err := LookupCloudFormationType(ctx, &args, opts...)
			var s LookupCloudFormationTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudFormationTypeResultOutput)
}

// A collection of arguments for invoking getCloudFormationType.
type LookupCloudFormationTypeOutputArgs struct {
	// ARN of the CloudFormation Type. For example, `arn:aws:cloudformation:us-west-2::type/resource/AWS-EC2-VPC`.
	Arn pulumi.StringPtrInput `pulumi:"arn"`
	// CloudFormation Registry Type. For example, `RESOURCE`.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// CloudFormation Type name. For example, `AWS::EC2::VPC`.
	TypeName pulumi.StringPtrInput `pulumi:"typeName"`
	// Identifier of the CloudFormation Type version.
	VersionId pulumi.StringPtrInput `pulumi:"versionId"`
}

func (LookupCloudFormationTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudFormationTypeArgs)(nil)).Elem()
}

// A collection of values returned by getCloudFormationType.
type LookupCloudFormationTypeResultOutput struct{ *pulumi.OutputState }

func (LookupCloudFormationTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudFormationTypeResult)(nil)).Elem()
}

func (o LookupCloudFormationTypeResultOutput) ToLookupCloudFormationTypeResultOutput() LookupCloudFormationTypeResultOutput {
	return o
}

func (o LookupCloudFormationTypeResultOutput) ToLookupCloudFormationTypeResultOutputWithContext(ctx context.Context) LookupCloudFormationTypeResultOutput {
	return o
}

func (o LookupCloudFormationTypeResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Identifier of the CloudFormation Type default version.
func (o LookupCloudFormationTypeResultOutput) DefaultVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.DefaultVersionId }).(pulumi.StringOutput)
}

// Deprecation status of the CloudFormation Type.
func (o LookupCloudFormationTypeResultOutput) DeprecatedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.DeprecatedStatus }).(pulumi.StringOutput)
}

// Description of the CloudFormation Type.
func (o LookupCloudFormationTypeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.Description }).(pulumi.StringOutput)
}

// URL of the documentation for the CloudFormation Type.
func (o LookupCloudFormationTypeResultOutput) DocumentationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.DocumentationUrl }).(pulumi.StringOutput)
}

// ARN of the IAM Role used to register the CloudFormation Type.
func (o LookupCloudFormationTypeResultOutput) ExecutionRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.ExecutionRoleArn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudFormationTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether the CloudFormation Type version is the default version.
func (o LookupCloudFormationTypeResultOutput) IsDefaultVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) bool { return v.IsDefaultVersion }).(pulumi.BoolOutput)
}

// List of objects containing logging configuration.
func (o LookupCloudFormationTypeResultOutput) LoggingConfigs() GetCloudFormationTypeLoggingConfigArrayOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) []GetCloudFormationTypeLoggingConfig { return v.LoggingConfigs }).(GetCloudFormationTypeLoggingConfigArrayOutput)
}

// Provisioning behavior of the CloudFormation Type.
func (o LookupCloudFormationTypeResultOutput) ProvisioningType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.ProvisioningType }).(pulumi.StringOutput)
}

// JSON document of the CloudFormation Type schema.
func (o LookupCloudFormationTypeResultOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.Schema }).(pulumi.StringOutput)
}

// URL of the source code for the CloudFormation Type.
func (o LookupCloudFormationTypeResultOutput) SourceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.SourceUrl }).(pulumi.StringOutput)
}

func (o LookupCloudFormationTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCloudFormationTypeResultOutput) TypeArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.TypeArn }).(pulumi.StringOutput)
}

func (o LookupCloudFormationTypeResultOutput) TypeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.TypeName }).(pulumi.StringOutput)
}

func (o LookupCloudFormationTypeResultOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) *string { return v.VersionId }).(pulumi.StringPtrOutput)
}

// Scope of the CloudFormation Type.
func (o LookupCloudFormationTypeResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudFormationTypeResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudFormationTypeResultOutput{})
}
