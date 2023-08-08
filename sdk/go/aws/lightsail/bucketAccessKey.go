// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a lightsail bucket access key. This is a set of credentials that allow API requests to be made to the lightsail bucket.
//
// ## Import
//
// terraform import {
//
//	to = aws_lightsail_bucket_access_key.test
//
//	id = "example-bucket,AKIA47VOQ2KPR7LLRZ6D" } Using `pulumi import`, import `aws_lightsail_bucket_access_key` using the `id` attribute. For exampleconsole % pulumi import aws_lightsail_bucket_access_key.test example-bucket,AKIA47VOQ2KPR7LLRZ6D
type BucketAccessKey struct {
	pulumi.CustomResourceState

	// The ID of the access key.
	AccessKeyId pulumi.StringOutput `pulumi:"accessKeyId"`
	// The name of the bucket that the new access key will belong to, and grant access to.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The timestamp when the access key was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The secret access key used to sign requests. This attribute is not available for imported resources. Note that this will be written to the state file.
	SecretAccessKey pulumi.StringOutput `pulumi:"secretAccessKey"`
	// The status of the access key.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBucketAccessKey registers a new resource with the given unique name, arguments, and options.
func NewBucketAccessKey(ctx *pulumi.Context,
	name string, args *BucketAccessKeyArgs, opts ...pulumi.ResourceOption) (*BucketAccessKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketAccessKey
	err := ctx.RegisterResource("aws:lightsail/bucketAccessKey:BucketAccessKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketAccessKey gets an existing BucketAccessKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketAccessKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketAccessKeyState, opts ...pulumi.ResourceOption) (*BucketAccessKey, error) {
	var resource BucketAccessKey
	err := ctx.ReadResource("aws:lightsail/bucketAccessKey:BucketAccessKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketAccessKey resources.
type bucketAccessKeyState struct {
	// The ID of the access key.
	AccessKeyId *string `pulumi:"accessKeyId"`
	// The name of the bucket that the new access key will belong to, and grant access to.
	BucketName *string `pulumi:"bucketName"`
	// The timestamp when the access key was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The secret access key used to sign requests. This attribute is not available for imported resources. Note that this will be written to the state file.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// The status of the access key.
	Status *string `pulumi:"status"`
}

type BucketAccessKeyState struct {
	// The ID of the access key.
	AccessKeyId pulumi.StringPtrInput
	// The name of the bucket that the new access key will belong to, and grant access to.
	BucketName pulumi.StringPtrInput
	// The timestamp when the access key was created.
	CreatedAt pulumi.StringPtrInput
	// The secret access key used to sign requests. This attribute is not available for imported resources. Note that this will be written to the state file.
	SecretAccessKey pulumi.StringPtrInput
	// The status of the access key.
	Status pulumi.StringPtrInput
}

func (BucketAccessKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessKeyState)(nil)).Elem()
}

type bucketAccessKeyArgs struct {
	// The name of the bucket that the new access key will belong to, and grant access to.
	BucketName string `pulumi:"bucketName"`
}

// The set of arguments for constructing a BucketAccessKey resource.
type BucketAccessKeyArgs struct {
	// The name of the bucket that the new access key will belong to, and grant access to.
	BucketName pulumi.StringInput
}

func (BucketAccessKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessKeyArgs)(nil)).Elem()
}

type BucketAccessKeyInput interface {
	pulumi.Input

	ToBucketAccessKeyOutput() BucketAccessKeyOutput
	ToBucketAccessKeyOutputWithContext(ctx context.Context) BucketAccessKeyOutput
}

func (*BucketAccessKey) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccessKey)(nil)).Elem()
}

func (i *BucketAccessKey) ToBucketAccessKeyOutput() BucketAccessKeyOutput {
	return i.ToBucketAccessKeyOutputWithContext(context.Background())
}

func (i *BucketAccessKey) ToBucketAccessKeyOutputWithContext(ctx context.Context) BucketAccessKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccessKeyOutput)
}

// BucketAccessKeyArrayInput is an input type that accepts BucketAccessKeyArray and BucketAccessKeyArrayOutput values.
// You can construct a concrete instance of `BucketAccessKeyArrayInput` via:
//
//	BucketAccessKeyArray{ BucketAccessKeyArgs{...} }
type BucketAccessKeyArrayInput interface {
	pulumi.Input

	ToBucketAccessKeyArrayOutput() BucketAccessKeyArrayOutput
	ToBucketAccessKeyArrayOutputWithContext(context.Context) BucketAccessKeyArrayOutput
}

type BucketAccessKeyArray []BucketAccessKeyInput

func (BucketAccessKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketAccessKey)(nil)).Elem()
}

func (i BucketAccessKeyArray) ToBucketAccessKeyArrayOutput() BucketAccessKeyArrayOutput {
	return i.ToBucketAccessKeyArrayOutputWithContext(context.Background())
}

func (i BucketAccessKeyArray) ToBucketAccessKeyArrayOutputWithContext(ctx context.Context) BucketAccessKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccessKeyArrayOutput)
}

// BucketAccessKeyMapInput is an input type that accepts BucketAccessKeyMap and BucketAccessKeyMapOutput values.
// You can construct a concrete instance of `BucketAccessKeyMapInput` via:
//
//	BucketAccessKeyMap{ "key": BucketAccessKeyArgs{...} }
type BucketAccessKeyMapInput interface {
	pulumi.Input

	ToBucketAccessKeyMapOutput() BucketAccessKeyMapOutput
	ToBucketAccessKeyMapOutputWithContext(context.Context) BucketAccessKeyMapOutput
}

type BucketAccessKeyMap map[string]BucketAccessKeyInput

func (BucketAccessKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketAccessKey)(nil)).Elem()
}

func (i BucketAccessKeyMap) ToBucketAccessKeyMapOutput() BucketAccessKeyMapOutput {
	return i.ToBucketAccessKeyMapOutputWithContext(context.Background())
}

func (i BucketAccessKeyMap) ToBucketAccessKeyMapOutputWithContext(ctx context.Context) BucketAccessKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccessKeyMapOutput)
}

type BucketAccessKeyOutput struct{ *pulumi.OutputState }

func (BucketAccessKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccessKey)(nil)).Elem()
}

func (o BucketAccessKeyOutput) ToBucketAccessKeyOutput() BucketAccessKeyOutput {
	return o
}

func (o BucketAccessKeyOutput) ToBucketAccessKeyOutputWithContext(ctx context.Context) BucketAccessKeyOutput {
	return o
}

// The ID of the access key.
func (o BucketAccessKeyOutput) AccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessKey) pulumi.StringOutput { return v.AccessKeyId }).(pulumi.StringOutput)
}

// The name of the bucket that the new access key will belong to, and grant access to.
func (o BucketAccessKeyOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessKey) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// The timestamp when the access key was created.
func (o BucketAccessKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The secret access key used to sign requests. This attribute is not available for imported resources. Note that this will be written to the state file.
func (o BucketAccessKeyOutput) SecretAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessKey) pulumi.StringOutput { return v.SecretAccessKey }).(pulumi.StringOutput)
}

// The status of the access key.
func (o BucketAccessKeyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessKey) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BucketAccessKeyArrayOutput struct{ *pulumi.OutputState }

func (BucketAccessKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketAccessKey)(nil)).Elem()
}

func (o BucketAccessKeyArrayOutput) ToBucketAccessKeyArrayOutput() BucketAccessKeyArrayOutput {
	return o
}

func (o BucketAccessKeyArrayOutput) ToBucketAccessKeyArrayOutputWithContext(ctx context.Context) BucketAccessKeyArrayOutput {
	return o
}

func (o BucketAccessKeyArrayOutput) Index(i pulumi.IntInput) BucketAccessKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketAccessKey {
		return vs[0].([]*BucketAccessKey)[vs[1].(int)]
	}).(BucketAccessKeyOutput)
}

type BucketAccessKeyMapOutput struct{ *pulumi.OutputState }

func (BucketAccessKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketAccessKey)(nil)).Elem()
}

func (o BucketAccessKeyMapOutput) ToBucketAccessKeyMapOutput() BucketAccessKeyMapOutput {
	return o
}

func (o BucketAccessKeyMapOutput) ToBucketAccessKeyMapOutputWithContext(ctx context.Context) BucketAccessKeyMapOutput {
	return o
}

func (o BucketAccessKeyMapOutput) MapIndex(k pulumi.StringInput) BucketAccessKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketAccessKey {
		return vs[0].(map[string]*BucketAccessKey)[vs[1].(string)]
	}).(BucketAccessKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccessKeyInput)(nil)).Elem(), &BucketAccessKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccessKeyArrayInput)(nil)).Elem(), BucketAccessKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccessKeyMapInput)(nil)).Elem(), BucketAccessKeyMap{})
	pulumi.RegisterOutputType(BucketAccessKeyOutput{})
	pulumi.RegisterOutputType(BucketAccessKeyArrayOutput{})
	pulumi.RegisterOutputType(BucketAccessKeyMapOutput{})
}
