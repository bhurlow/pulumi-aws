// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Lambda Layer Version resource. Lambda Layers allow you to reuse shared bits of code across multiple lambda functions.
//
// For information about Lambda Layers and how to use them, see [AWS Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
//
// > **NOTE:** Setting `skipDestroy` to `true` means that the AWS Provider will _not_ destroy any layer version, even when running destroy. Layer versions are thus intentional dangling resources that are _not_ managed by the provider and may incur extra expense in your AWS account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := lambda.NewLayerVersion(ctx, "lambdaLayer", &lambda.LayerVersionArgs{
//				CompatibleRuntimes: pulumi.StringArray{
//					pulumi.String("nodejs16.x"),
//				},
//				Code:      pulumi.NewFileArchive("lambda_layer_payload.zip"),
//				LayerName: pulumi.String("lambda_layer_name"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Specifying the Deployment Package
//
// AWS Lambda Layers expect source code to be provided as a deployment package whose structure varies depending on which `compatibleRuntimes` this layer specifies.
// See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) for the valid values of `compatibleRuntimes`.
//
// Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
// indirectly via Amazon S3 (using the `s3Bucket`, `s3Key` and `s3ObjectVersion` arguments). When providing the deployment
// package via S3 it may be useful to use the `s3.BucketObjectv2` resource to upload it.
//
// For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading large files efficiently.
//
// ## Import
//
// terraform import {
//
//	to = aws_lambda_layer_version.test_layer
//
//	id = "arn:aws:lambda:_REGION_:_ACCOUNT_ID_:layer:_LAYER_NAME_:_LAYER_VERSION_" } Using `pulumi import`, import Lambda Layers using `arn`. For exampleconsole % pulumi import \
//
//	aws_lambda_layer_version.test_layer \
//
//	arn:aws:lambda:_REGION_:_ACCOUNT_ID_:layer:_LAYER_NAME_:_LAYER_VERSION_
type LayerVersion struct {
	pulumi.CustomResourceState

	// ARN of the Lambda Layer with version.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.ArchiveOutput `pulumi:"code"`
	// List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
	CompatibleArchitectures pulumi.StringArrayOutput `pulumi:"compatibleArchitectures"`
	// List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes pulumi.StringArrayOutput `pulumi:"compatibleRuntimes"`
	// Date this resource was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Description of what your Lambda Layer does.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ARN of the Lambda Layer without version.
	LayerArn pulumi.StringOutput `pulumi:"layerArn"`
	// Unique name for your Lambda Layer
	//
	// The following arguments are optional:
	LayerName pulumi.StringOutput `pulumi:"layerName"`
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo pulumi.StringPtrOutput `pulumi:"licenseInfo"`
	// S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket pulumi.StringPtrOutput `pulumi:"s3Bucket"`
	// S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key pulumi.StringPtrOutput `pulumi:"s3Key"`
	// Object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion pulumi.StringPtrOutput `pulumi:"s3ObjectVersion"`
	// ARN of a signing job.
	SigningJobArn pulumi.StringOutput `pulumi:"signingJobArn"`
	// ARN for a signing profile version.
	SigningProfileVersionArn pulumi.StringOutput `pulumi:"signingProfileVersionArn"`
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy pulumi.BoolPtrOutput `pulumi:"skipDestroy"`
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
	SourceCodeHash pulumi.StringOutput `pulumi:"sourceCodeHash"`
	// Size in bytes of the function .zip file.
	SourceCodeSize pulumi.IntOutput `pulumi:"sourceCodeSize"`
	// Lambda Layer version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewLayerVersion registers a new resource with the given unique name, arguments, and options.
func NewLayerVersion(ctx *pulumi.Context,
	name string, args *LayerVersionArgs, opts ...pulumi.ResourceOption) (*LayerVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LayerName == nil {
		return nil, errors.New("invalid value for required argument 'LayerName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LayerVersion
	err := ctx.RegisterResource("aws:lambda/layerVersion:LayerVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLayerVersion gets an existing LayerVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLayerVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LayerVersionState, opts ...pulumi.ResourceOption) (*LayerVersion, error) {
	var resource LayerVersion
	err := ctx.ReadResource("aws:lambda/layerVersion:LayerVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LayerVersion resources.
type layerVersionState struct {
	// ARN of the Lambda Layer with version.
	Arn *string `pulumi:"arn"`
	// Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.Archive `pulumi:"code"`
	// List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
	CompatibleArchitectures []string `pulumi:"compatibleArchitectures"`
	// List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// Date this resource was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Description of what your Lambda Layer does.
	Description *string `pulumi:"description"`
	// ARN of the Lambda Layer without version.
	LayerArn *string `pulumi:"layerArn"`
	// Unique name for your Lambda Layer
	//
	// The following arguments are optional:
	LayerName *string `pulumi:"layerName"`
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo *string `pulumi:"licenseInfo"`
	// S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket *string `pulumi:"s3Bucket"`
	// S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key *string `pulumi:"s3Key"`
	// Object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion *string `pulumi:"s3ObjectVersion"`
	// ARN of a signing job.
	SigningJobArn *string `pulumi:"signingJobArn"`
	// ARN for a signing profile version.
	SigningProfileVersionArn *string `pulumi:"signingProfileVersionArn"`
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
	SourceCodeHash *string `pulumi:"sourceCodeHash"`
	// Size in bytes of the function .zip file.
	SourceCodeSize *int `pulumi:"sourceCodeSize"`
	// Lambda Layer version.
	Version *string `pulumi:"version"`
}

type LayerVersionState struct {
	// ARN of the Lambda Layer with version.
	Arn pulumi.StringPtrInput
	// Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.ArchiveInput
	// List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
	CompatibleArchitectures pulumi.StringArrayInput
	// List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes pulumi.StringArrayInput
	// Date this resource was created.
	CreatedDate pulumi.StringPtrInput
	// Description of what your Lambda Layer does.
	Description pulumi.StringPtrInput
	// ARN of the Lambda Layer without version.
	LayerArn pulumi.StringPtrInput
	// Unique name for your Lambda Layer
	//
	// The following arguments are optional:
	LayerName pulumi.StringPtrInput
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo pulumi.StringPtrInput
	// S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket pulumi.StringPtrInput
	// S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key pulumi.StringPtrInput
	// Object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion pulumi.StringPtrInput
	// ARN of a signing job.
	SigningJobArn pulumi.StringPtrInput
	// ARN for a signing profile version.
	SigningProfileVersionArn pulumi.StringPtrInput
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy pulumi.BoolPtrInput
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
	SourceCodeHash pulumi.StringPtrInput
	// Size in bytes of the function .zip file.
	SourceCodeSize pulumi.IntPtrInput
	// Lambda Layer version.
	Version pulumi.StringPtrInput
}

func (LayerVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionState)(nil)).Elem()
}

type layerVersionArgs struct {
	// Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.Archive `pulumi:"code"`
	// List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
	CompatibleArchitectures []string `pulumi:"compatibleArchitectures"`
	// List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// Description of what your Lambda Layer does.
	Description *string `pulumi:"description"`
	// Unique name for your Lambda Layer
	//
	// The following arguments are optional:
	LayerName string `pulumi:"layerName"`
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo *string `pulumi:"licenseInfo"`
	// S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket *string `pulumi:"s3Bucket"`
	// S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key *string `pulumi:"s3Key"`
	// Object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion *string `pulumi:"s3ObjectVersion"`
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
	SourceCodeHash *string `pulumi:"sourceCodeHash"`
}

// The set of arguments for constructing a LayerVersion resource.
type LayerVersionArgs struct {
	// Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.ArchiveInput
	// List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
	CompatibleArchitectures pulumi.StringArrayInput
	// List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes pulumi.StringArrayInput
	// Description of what your Lambda Layer does.
	Description pulumi.StringPtrInput
	// Unique name for your Lambda Layer
	//
	// The following arguments are optional:
	LayerName pulumi.StringInput
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo pulumi.StringPtrInput
	// S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket pulumi.StringPtrInput
	// S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key pulumi.StringPtrInput
	// Object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion pulumi.StringPtrInput
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy pulumi.BoolPtrInput
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
	SourceCodeHash pulumi.StringPtrInput
}

func (LayerVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionArgs)(nil)).Elem()
}

type LayerVersionInput interface {
	pulumi.Input

	ToLayerVersionOutput() LayerVersionOutput
	ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput
}

func (*LayerVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersion)(nil)).Elem()
}

func (i *LayerVersion) ToLayerVersionOutput() LayerVersionOutput {
	return i.ToLayerVersionOutputWithContext(context.Background())
}

func (i *LayerVersion) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionOutput)
}

// LayerVersionArrayInput is an input type that accepts LayerVersionArray and LayerVersionArrayOutput values.
// You can construct a concrete instance of `LayerVersionArrayInput` via:
//
//	LayerVersionArray{ LayerVersionArgs{...} }
type LayerVersionArrayInput interface {
	pulumi.Input

	ToLayerVersionArrayOutput() LayerVersionArrayOutput
	ToLayerVersionArrayOutputWithContext(context.Context) LayerVersionArrayOutput
}

type LayerVersionArray []LayerVersionInput

func (LayerVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LayerVersion)(nil)).Elem()
}

func (i LayerVersionArray) ToLayerVersionArrayOutput() LayerVersionArrayOutput {
	return i.ToLayerVersionArrayOutputWithContext(context.Background())
}

func (i LayerVersionArray) ToLayerVersionArrayOutputWithContext(ctx context.Context) LayerVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionArrayOutput)
}

// LayerVersionMapInput is an input type that accepts LayerVersionMap and LayerVersionMapOutput values.
// You can construct a concrete instance of `LayerVersionMapInput` via:
//
//	LayerVersionMap{ "key": LayerVersionArgs{...} }
type LayerVersionMapInput interface {
	pulumi.Input

	ToLayerVersionMapOutput() LayerVersionMapOutput
	ToLayerVersionMapOutputWithContext(context.Context) LayerVersionMapOutput
}

type LayerVersionMap map[string]LayerVersionInput

func (LayerVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LayerVersion)(nil)).Elem()
}

func (i LayerVersionMap) ToLayerVersionMapOutput() LayerVersionMapOutput {
	return i.ToLayerVersionMapOutputWithContext(context.Background())
}

func (i LayerVersionMap) ToLayerVersionMapOutputWithContext(ctx context.Context) LayerVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionMapOutput)
}

type LayerVersionOutput struct{ *pulumi.OutputState }

func (LayerVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersion)(nil)).Elem()
}

func (o LayerVersionOutput) ToLayerVersionOutput() LayerVersionOutput {
	return o
}

func (o LayerVersionOutput) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return o
}

// ARN of the Lambda Layer with version.
func (o LayerVersionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
func (o LayerVersionOutput) Code() pulumi.ArchiveOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.ArchiveOutput { return v.Code }).(pulumi.ArchiveOutput)
}

// List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
func (o LayerVersionOutput) CompatibleArchitectures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringArrayOutput { return v.CompatibleArchitectures }).(pulumi.StringArrayOutput)
}

// List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
func (o LayerVersionOutput) CompatibleRuntimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringArrayOutput { return v.CompatibleRuntimes }).(pulumi.StringArrayOutput)
}

// Date this resource was created.
func (o LayerVersionOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// Description of what your Lambda Layer does.
func (o LayerVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ARN of the Lambda Layer without version.
func (o LayerVersionOutput) LayerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.LayerArn }).(pulumi.StringOutput)
}

// Unique name for your Lambda Layer
//
// The following arguments are optional:
func (o LayerVersionOutput) LayerName() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.LayerName }).(pulumi.StringOutput)
}

// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
func (o LayerVersionOutput) LicenseInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.LicenseInfo }).(pulumi.StringPtrOutput)
}

// S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
func (o LayerVersionOutput) S3Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.S3Bucket }).(pulumi.StringPtrOutput)
}

// S3 key of an object containing the function's deployment package. Conflicts with `filename`.
func (o LayerVersionOutput) S3Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.S3Key }).(pulumi.StringPtrOutput)
}

// Object version containing the function's deployment package. Conflicts with `filename`.
func (o LayerVersionOutput) S3ObjectVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.S3ObjectVersion }).(pulumi.StringPtrOutput)
}

// ARN of a signing job.
func (o LayerVersionOutput) SigningJobArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.SigningJobArn }).(pulumi.StringOutput)
}

// ARN for a signing profile version.
func (o LayerVersionOutput) SigningProfileVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.SigningProfileVersionArn }).(pulumi.StringOutput)
}

// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
func (o LayerVersionOutput) SkipDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.BoolPtrOutput { return v.SkipDestroy }).(pulumi.BoolPtrOutput)
}

// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
func (o LayerVersionOutput) SourceCodeHash() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.SourceCodeHash }).(pulumi.StringOutput)
}

// Size in bytes of the function .zip file.
func (o LayerVersionOutput) SourceCodeSize() pulumi.IntOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.IntOutput { return v.SourceCodeSize }).(pulumi.IntOutput)
}

// Lambda Layer version.
func (o LayerVersionOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type LayerVersionArrayOutput struct{ *pulumi.OutputState }

func (LayerVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LayerVersion)(nil)).Elem()
}

func (o LayerVersionArrayOutput) ToLayerVersionArrayOutput() LayerVersionArrayOutput {
	return o
}

func (o LayerVersionArrayOutput) ToLayerVersionArrayOutputWithContext(ctx context.Context) LayerVersionArrayOutput {
	return o
}

func (o LayerVersionArrayOutput) Index(i pulumi.IntInput) LayerVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LayerVersion {
		return vs[0].([]*LayerVersion)[vs[1].(int)]
	}).(LayerVersionOutput)
}

type LayerVersionMapOutput struct{ *pulumi.OutputState }

func (LayerVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LayerVersion)(nil)).Elem()
}

func (o LayerVersionMapOutput) ToLayerVersionMapOutput() LayerVersionMapOutput {
	return o
}

func (o LayerVersionMapOutput) ToLayerVersionMapOutputWithContext(ctx context.Context) LayerVersionMapOutput {
	return o
}

func (o LayerVersionMapOutput) MapIndex(k pulumi.StringInput) LayerVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LayerVersion {
		return vs[0].(map[string]*LayerVersion)[vs[1].(string)]
	}).(LayerVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LayerVersionInput)(nil)).Elem(), &LayerVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*LayerVersionArrayInput)(nil)).Elem(), LayerVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LayerVersionMapInput)(nil)).Elem(), LayerVersionMap{})
	pulumi.RegisterOutputType(LayerVersionOutput{})
	pulumi.RegisterOutputType(LayerVersionArrayOutput{})
	pulumi.RegisterOutputType(LayerVersionMapOutput{})
}
