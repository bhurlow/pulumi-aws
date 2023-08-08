// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage AWS Device Farm Uploads.
//
// > **NOTE:** AWS currently has limited regional support for Device Farm (e.g., `us-west-2`). See [AWS Device Farm endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/devicefarm.html) for information on supported regions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/devicefarm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleProject, err := devicefarm.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			_, err = devicefarm.NewUpload(ctx, "exampleUpload", &devicefarm.UploadArgs{
//				ProjectArn: exampleProject.Arn,
//				Type:       pulumi.String("APPIUM_JAVA_TESTNG_TEST_SPEC"),
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
//	to = aws_devicefarm_upload.example
//
//	id = "arn:aws:devicefarm:us-west-2:123456789012:upload:4fa784c7-ccb4-4dbf-ba4f-02198320daa1" } Using `pulumi import`, import DeviceFarm Uploads using their ARN. For exampleconsole % pulumi import aws_devicefarm_upload.example arn:aws:devicefarm:us-west-2:123456789012:upload:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
type Upload struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of this upload.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The upload's category.
	Category pulumi.StringOutput `pulumi:"category"`
	// The upload's content type (for example, application/octet-stream).
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The upload's metadata. For example, for Android, this contains information that is parsed from the manifest and is displayed in the AWS Device Farm console after the associated app is uploaded.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the project for the upload.
	ProjectArn pulumi.StringOutput `pulumi:"projectArn"`
	// The upload's upload type. See [AWS Docs](https://docs.aws.amazon.com/devicefarm/latest/APIReference/API_CreateUpload.html#API_CreateUpload_RequestSyntax) for valid list of values.
	Type pulumi.StringOutput `pulumi:"type"`
	// The presigned Amazon S3 URL that was used to store a file using a PUT request.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewUpload registers a new resource with the given unique name, arguments, and options.
func NewUpload(ctx *pulumi.Context,
	name string, args *UploadArgs, opts ...pulumi.ResourceOption) (*Upload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectArn == nil {
		return nil, errors.New("invalid value for required argument 'ProjectArn'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Upload
	err := ctx.RegisterResource("aws:devicefarm/upload:Upload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpload gets an existing Upload resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UploadState, opts ...pulumi.ResourceOption) (*Upload, error) {
	var resource Upload
	err := ctx.ReadResource("aws:devicefarm/upload:Upload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Upload resources.
type uploadState struct {
	// The Amazon Resource Name of this upload.
	Arn *string `pulumi:"arn"`
	// The upload's category.
	Category *string `pulumi:"category"`
	// The upload's content type (for example, application/octet-stream).
	ContentType *string `pulumi:"contentType"`
	// The upload's metadata. For example, for Android, this contains information that is parsed from the manifest and is displayed in the AWS Device Farm console after the associated app is uploaded.
	Metadata *string `pulumi:"metadata"`
	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name *string `pulumi:"name"`
	// The ARN of the project for the upload.
	ProjectArn *string `pulumi:"projectArn"`
	// The upload's upload type. See [AWS Docs](https://docs.aws.amazon.com/devicefarm/latest/APIReference/API_CreateUpload.html#API_CreateUpload_RequestSyntax) for valid list of values.
	Type *string `pulumi:"type"`
	// The presigned Amazon S3 URL that was used to store a file using a PUT request.
	Url *string `pulumi:"url"`
}

type UploadState struct {
	// The Amazon Resource Name of this upload.
	Arn pulumi.StringPtrInput
	// The upload's category.
	Category pulumi.StringPtrInput
	// The upload's content type (for example, application/octet-stream).
	ContentType pulumi.StringPtrInput
	// The upload's metadata. For example, for Android, this contains information that is parsed from the manifest and is displayed in the AWS Device Farm console after the associated app is uploaded.
	Metadata pulumi.StringPtrInput
	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name pulumi.StringPtrInput
	// The ARN of the project for the upload.
	ProjectArn pulumi.StringPtrInput
	// The upload's upload type. See [AWS Docs](https://docs.aws.amazon.com/devicefarm/latest/APIReference/API_CreateUpload.html#API_CreateUpload_RequestSyntax) for valid list of values.
	Type pulumi.StringPtrInput
	// The presigned Amazon S3 URL that was used to store a file using a PUT request.
	Url pulumi.StringPtrInput
}

func (UploadState) ElementType() reflect.Type {
	return reflect.TypeOf((*uploadState)(nil)).Elem()
}

type uploadArgs struct {
	// The upload's content type (for example, application/octet-stream).
	ContentType *string `pulumi:"contentType"`
	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name *string `pulumi:"name"`
	// The ARN of the project for the upload.
	ProjectArn string `pulumi:"projectArn"`
	// The upload's upload type. See [AWS Docs](https://docs.aws.amazon.com/devicefarm/latest/APIReference/API_CreateUpload.html#API_CreateUpload_RequestSyntax) for valid list of values.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Upload resource.
type UploadArgs struct {
	// The upload's content type (for example, application/octet-stream).
	ContentType pulumi.StringPtrInput
	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name pulumi.StringPtrInput
	// The ARN of the project for the upload.
	ProjectArn pulumi.StringInput
	// The upload's upload type. See [AWS Docs](https://docs.aws.amazon.com/devicefarm/latest/APIReference/API_CreateUpload.html#API_CreateUpload_RequestSyntax) for valid list of values.
	Type pulumi.StringInput
}

func (UploadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uploadArgs)(nil)).Elem()
}

type UploadInput interface {
	pulumi.Input

	ToUploadOutput() UploadOutput
	ToUploadOutputWithContext(ctx context.Context) UploadOutput
}

func (*Upload) ElementType() reflect.Type {
	return reflect.TypeOf((**Upload)(nil)).Elem()
}

func (i *Upload) ToUploadOutput() UploadOutput {
	return i.ToUploadOutputWithContext(context.Background())
}

func (i *Upload) ToUploadOutputWithContext(ctx context.Context) UploadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadOutput)
}

// UploadArrayInput is an input type that accepts UploadArray and UploadArrayOutput values.
// You can construct a concrete instance of `UploadArrayInput` via:
//
//	UploadArray{ UploadArgs{...} }
type UploadArrayInput interface {
	pulumi.Input

	ToUploadArrayOutput() UploadArrayOutput
	ToUploadArrayOutputWithContext(context.Context) UploadArrayOutput
}

type UploadArray []UploadInput

func (UploadArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Upload)(nil)).Elem()
}

func (i UploadArray) ToUploadArrayOutput() UploadArrayOutput {
	return i.ToUploadArrayOutputWithContext(context.Background())
}

func (i UploadArray) ToUploadArrayOutputWithContext(ctx context.Context) UploadArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadArrayOutput)
}

// UploadMapInput is an input type that accepts UploadMap and UploadMapOutput values.
// You can construct a concrete instance of `UploadMapInput` via:
//
//	UploadMap{ "key": UploadArgs{...} }
type UploadMapInput interface {
	pulumi.Input

	ToUploadMapOutput() UploadMapOutput
	ToUploadMapOutputWithContext(context.Context) UploadMapOutput
}

type UploadMap map[string]UploadInput

func (UploadMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Upload)(nil)).Elem()
}

func (i UploadMap) ToUploadMapOutput() UploadMapOutput {
	return i.ToUploadMapOutputWithContext(context.Background())
}

func (i UploadMap) ToUploadMapOutputWithContext(ctx context.Context) UploadMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadMapOutput)
}

type UploadOutput struct{ *pulumi.OutputState }

func (UploadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Upload)(nil)).Elem()
}

func (o UploadOutput) ToUploadOutput() UploadOutput {
	return o
}

func (o UploadOutput) ToUploadOutputWithContext(ctx context.Context) UploadOutput {
	return o
}

// The Amazon Resource Name of this upload.
func (o UploadOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The upload's category.
func (o UploadOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// The upload's content type (for example, application/octet-stream).
func (o UploadOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The upload's metadata. For example, for Android, this contains information that is parsed from the manifest and is displayed in the AWS Device Farm console after the associated app is uploaded.
func (o UploadOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringOutput { return v.Metadata }).(pulumi.StringOutput)
}

// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
func (o UploadOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the project for the upload.
func (o UploadOutput) ProjectArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringOutput { return v.ProjectArn }).(pulumi.StringOutput)
}

// The upload's upload type. See [AWS Docs](https://docs.aws.amazon.com/devicefarm/latest/APIReference/API_CreateUpload.html#API_CreateUpload_RequestSyntax) for valid list of values.
func (o UploadOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The presigned Amazon S3 URL that was used to store a file using a PUT request.
func (o UploadOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Upload) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type UploadArrayOutput struct{ *pulumi.OutputState }

func (UploadArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Upload)(nil)).Elem()
}

func (o UploadArrayOutput) ToUploadArrayOutput() UploadArrayOutput {
	return o
}

func (o UploadArrayOutput) ToUploadArrayOutputWithContext(ctx context.Context) UploadArrayOutput {
	return o
}

func (o UploadArrayOutput) Index(i pulumi.IntInput) UploadOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Upload {
		return vs[0].([]*Upload)[vs[1].(int)]
	}).(UploadOutput)
}

type UploadMapOutput struct{ *pulumi.OutputState }

func (UploadMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Upload)(nil)).Elem()
}

func (o UploadMapOutput) ToUploadMapOutput() UploadMapOutput {
	return o
}

func (o UploadMapOutput) ToUploadMapOutputWithContext(ctx context.Context) UploadMapOutput {
	return o
}

func (o UploadMapOutput) MapIndex(k pulumi.StringInput) UploadOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Upload {
		return vs[0].(map[string]*Upload)[vs[1].(string)]
	}).(UploadOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UploadInput)(nil)).Elem(), &Upload{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadArrayInput)(nil)).Elem(), UploadArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadMapInput)(nil)).Elem(), UploadMap{})
	pulumi.RegisterOutputType(UploadOutput{})
	pulumi.RegisterOutputType(UploadArrayOutput{})
	pulumi.RegisterOutputType(UploadMapOutput{})
}
