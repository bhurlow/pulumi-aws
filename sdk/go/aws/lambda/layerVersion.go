// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Lambda Layer Version resource. Lambda Layers allow you to reuse shared bits of code across multiple lambda functions.
//
// For information about Lambda Layers and how to use them, see [AWS Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
//
// ## Example Usage
// ## Specifying the Deployment Package
//
// AWS Lambda Layers expect source code to be provided as a deployment package whose structure varies depending on which `compatibleRuntimes` this layer specifies.
// See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) for the valid values of `compatibleRuntimes`.
//
// Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
// indirectly via Amazon S3 (using the `s3Bucket`, `s3Key` and `s3ObjectVersion` arguments). When providing the deployment
// package via S3 it may be useful to use the `s3.BucketObject` resource to upload it.
//
// For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading
// large files efficiently.
//
// ## Import
//
// Lambda Layers can be imported using `arn`.
//
// ```sh
//  $ pulumi import aws:lambda/layerVersion:LayerVersion \
// ```
//
//  aws_lambda_layer_version.test_layer \
//
//  arn:aws:lambda:_REGION_:_ACCOUNT_ID_:layer:_LAYER_NAME_:_LAYER_VERSION_
type LayerVersion struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Lambda Layer with version.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.ArchiveOutput `pulumi:"code"`
	// The compatible architectures for the specific Lambda Layer Version. Valid values are `arm64` and `x8664`. If not supplied, a value of `null` is assumed and will default to the value used for the function architecture.
	CompatibleArchitectures pulumi.StringArrayOutput `pulumi:"compatibleArchitectures"`
	// A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes pulumi.StringArrayOutput `pulumi:"compatibleRuntimes"`
	// The date this resource was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Description of what your Lambda Layer does.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the Lambda Layer without version.
	LayerArn pulumi.StringOutput `pulumi:"layerArn"`
	// A unique name for your Lambda Layer
	LayerName pulumi.StringOutput `pulumi:"layerName"`
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo pulumi.StringPtrOutput `pulumi:"licenseInfo"`
	// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket pulumi.StringPtrOutput `pulumi:"s3Bucket"`
	// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key pulumi.StringPtrOutput `pulumi:"s3Key"`
	// The object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion pulumi.StringPtrOutput `pulumi:"s3ObjectVersion"`
	// The Amazon Resource Name (ARN) of a signing job.
	SigningJobArn pulumi.StringOutput `pulumi:"signingJobArn"`
	// The Amazon Resource Name (ARN) for a signing profile version.
	SigningProfileVersionArn pulumi.StringOutput `pulumi:"signingProfileVersionArn"`
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
	SourceCodeHash pulumi.StringOutput `pulumi:"sourceCodeHash"`
	// The size in bytes of the function .zip file.
	SourceCodeSize pulumi.IntOutput `pulumi:"sourceCodeSize"`
	// This Lamba Layer version.
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
	// The Amazon Resource Name (ARN) of the Lambda Layer with version.
	Arn *string `pulumi:"arn"`
	// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.Archive `pulumi:"code"`
	// The compatible architectures for the specific Lambda Layer Version. Valid values are `arm64` and `x8664`. If not supplied, a value of `null` is assumed and will default to the value used for the function architecture.
	CompatibleArchitectures []string `pulumi:"compatibleArchitectures"`
	// A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// The date this resource was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Description of what your Lambda Layer does.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the Lambda Layer without version.
	LayerArn *string `pulumi:"layerArn"`
	// A unique name for your Lambda Layer
	LayerName *string `pulumi:"layerName"`
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo *string `pulumi:"licenseInfo"`
	// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket *string `pulumi:"s3Bucket"`
	// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key *string `pulumi:"s3Key"`
	// The object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion *string `pulumi:"s3ObjectVersion"`
	// The Amazon Resource Name (ARN) of a signing job.
	SigningJobArn *string `pulumi:"signingJobArn"`
	// The Amazon Resource Name (ARN) for a signing profile version.
	SigningProfileVersionArn *string `pulumi:"signingProfileVersionArn"`
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
	SourceCodeHash *string `pulumi:"sourceCodeHash"`
	// The size in bytes of the function .zip file.
	SourceCodeSize *int `pulumi:"sourceCodeSize"`
	// This Lamba Layer version.
	Version *string `pulumi:"version"`
}

type LayerVersionState struct {
	// The Amazon Resource Name (ARN) of the Lambda Layer with version.
	Arn pulumi.StringPtrInput
	// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.ArchiveInput
	// The compatible architectures for the specific Lambda Layer Version. Valid values are `arm64` and `x8664`. If not supplied, a value of `null` is assumed and will default to the value used for the function architecture.
	CompatibleArchitectures pulumi.StringArrayInput
	// A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes pulumi.StringArrayInput
	// The date this resource was created.
	CreatedDate pulumi.StringPtrInput
	// Description of what your Lambda Layer does.
	Description pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Lambda Layer without version.
	LayerArn pulumi.StringPtrInput
	// A unique name for your Lambda Layer
	LayerName pulumi.StringPtrInput
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo pulumi.StringPtrInput
	// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket pulumi.StringPtrInput
	// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key pulumi.StringPtrInput
	// The object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of a signing job.
	SigningJobArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for a signing profile version.
	SigningProfileVersionArn pulumi.StringPtrInput
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
	SourceCodeHash pulumi.StringPtrInput
	// The size in bytes of the function .zip file.
	SourceCodeSize pulumi.IntPtrInput
	// This Lamba Layer version.
	Version pulumi.StringPtrInput
}

func (LayerVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionState)(nil)).Elem()
}

type layerVersionArgs struct {
	// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.Archive `pulumi:"code"`
	// The compatible architectures for the specific Lambda Layer Version. Valid values are `arm64` and `x8664`. If not supplied, a value of `null` is assumed and will default to the value used for the function architecture.
	CompatibleArchitectures []string `pulumi:"compatibleArchitectures"`
	// A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// Description of what your Lambda Layer does.
	Description *string `pulumi:"description"`
	// A unique name for your Lambda Layer
	LayerName string `pulumi:"layerName"`
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo *string `pulumi:"licenseInfo"`
	// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket *string `pulumi:"s3Bucket"`
	// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key *string `pulumi:"s3Key"`
	// The object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion *string `pulumi:"s3ObjectVersion"`
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
	SourceCodeHash *string `pulumi:"sourceCodeHash"`
}

// The set of arguments for constructing a LayerVersion resource.
type LayerVersionArgs struct {
	// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
	Code pulumi.ArchiveInput
	// The compatible architectures for the specific Lambda Layer Version. Valid values are `arm64` and `x8664`. If not supplied, a value of `null` is assumed and will default to the value used for the function architecture.
	CompatibleArchitectures pulumi.StringArrayInput
	// A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes pulumi.StringArrayInput
	// Description of what your Lambda Layer does.
	Description pulumi.StringPtrInput
	// A unique name for your Lambda Layer
	LayerName pulumi.StringInput
	// License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
	LicenseInfo pulumi.StringPtrInput
	// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket pulumi.StringPtrInput
	// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
	S3Key pulumi.StringPtrInput
	// The object version containing the function's deployment package. Conflicts with `filename`.
	S3ObjectVersion pulumi.StringPtrInput
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
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
	return reflect.TypeOf((*LayerVersion)(nil))
}

func (i *LayerVersion) ToLayerVersionOutput() LayerVersionOutput {
	return i.ToLayerVersionOutputWithContext(context.Background())
}

func (i *LayerVersion) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionOutput)
}

func (i *LayerVersion) ToLayerVersionPtrOutput() LayerVersionPtrOutput {
	return i.ToLayerVersionPtrOutputWithContext(context.Background())
}

func (i *LayerVersion) ToLayerVersionPtrOutputWithContext(ctx context.Context) LayerVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionPtrOutput)
}

type LayerVersionPtrInput interface {
	pulumi.Input

	ToLayerVersionPtrOutput() LayerVersionPtrOutput
	ToLayerVersionPtrOutputWithContext(ctx context.Context) LayerVersionPtrOutput
}

type layerVersionPtrType LayerVersionArgs

func (*layerVersionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersion)(nil))
}

func (i *layerVersionPtrType) ToLayerVersionPtrOutput() LayerVersionPtrOutput {
	return i.ToLayerVersionPtrOutputWithContext(context.Background())
}

func (i *layerVersionPtrType) ToLayerVersionPtrOutputWithContext(ctx context.Context) LayerVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionPtrOutput)
}

// LayerVersionArrayInput is an input type that accepts LayerVersionArray and LayerVersionArrayOutput values.
// You can construct a concrete instance of `LayerVersionArrayInput` via:
//
//          LayerVersionArray{ LayerVersionArgs{...} }
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
//          LayerVersionMap{ "key": LayerVersionArgs{...} }
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
	return reflect.TypeOf((*LayerVersion)(nil))
}

func (o LayerVersionOutput) ToLayerVersionOutput() LayerVersionOutput {
	return o
}

func (o LayerVersionOutput) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return o
}

func (o LayerVersionOutput) ToLayerVersionPtrOutput() LayerVersionPtrOutput {
	return o.ToLayerVersionPtrOutputWithContext(context.Background())
}

func (o LayerVersionOutput) ToLayerVersionPtrOutputWithContext(ctx context.Context) LayerVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LayerVersion) *LayerVersion {
		return &v
	}).(LayerVersionPtrOutput)
}

type LayerVersionPtrOutput struct{ *pulumi.OutputState }

func (LayerVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersion)(nil))
}

func (o LayerVersionPtrOutput) ToLayerVersionPtrOutput() LayerVersionPtrOutput {
	return o
}

func (o LayerVersionPtrOutput) ToLayerVersionPtrOutputWithContext(ctx context.Context) LayerVersionPtrOutput {
	return o
}

func (o LayerVersionPtrOutput) Elem() LayerVersionOutput {
	return o.ApplyT(func(v *LayerVersion) LayerVersion {
		if v != nil {
			return *v
		}
		var ret LayerVersion
		return ret
	}).(LayerVersionOutput)
}

type LayerVersionArrayOutput struct{ *pulumi.OutputState }

func (LayerVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LayerVersion)(nil))
}

func (o LayerVersionArrayOutput) ToLayerVersionArrayOutput() LayerVersionArrayOutput {
	return o
}

func (o LayerVersionArrayOutput) ToLayerVersionArrayOutputWithContext(ctx context.Context) LayerVersionArrayOutput {
	return o
}

func (o LayerVersionArrayOutput) Index(i pulumi.IntInput) LayerVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LayerVersion {
		return vs[0].([]LayerVersion)[vs[1].(int)]
	}).(LayerVersionOutput)
}

type LayerVersionMapOutput struct{ *pulumi.OutputState }

func (LayerVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LayerVersion)(nil))
}

func (o LayerVersionMapOutput) ToLayerVersionMapOutput() LayerVersionMapOutput {
	return o
}

func (o LayerVersionMapOutput) ToLayerVersionMapOutputWithContext(ctx context.Context) LayerVersionMapOutput {
	return o
}

func (o LayerVersionMapOutput) MapIndex(k pulumi.StringInput) LayerVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LayerVersion {
		return vs[0].(map[string]LayerVersion)[vs[1].(string)]
	}).(LayerVersionOutput)
}

func init() {
	pulumi.RegisterOutputType(LayerVersionOutput{})
	pulumi.RegisterOutputType(LayerVersionPtrOutput{})
	pulumi.RegisterOutputType(LayerVersionArrayOutput{})
	pulumi.RegisterOutputType(LayerVersionMapOutput{})
}
