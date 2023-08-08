// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// The following example below creates a CloudFront public key.
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfront.NewPublicKey(ctx, "example", &cloudfront.PublicKeyArgs{
//				Comment:    pulumi.String("test public key"),
//				EncodedKey: readFileOrPanic("public_key.pem"),
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
//	to = aws_cloudfront_public_key.example
//
//	id = "K3D5EWEUDCCXON" } Using `pulumi import`, import CloudFront Public Key using the `id`. For exampleconsole % pulumi import aws_cloudfront_public_key.example K3D5EWEUDCCXON
type PublicKey struct {
	pulumi.CustomResourceState

	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference pulumi.StringOutput `pulumi:"callerReference"`
	// An optional comment about the public key.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey pulumi.StringOutput `pulumi:"encodedKey"`
	// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name for the public key. By default generated by this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name for the public key. Conflicts with `name`.
	//
	// **NOTE:** When setting `encodedKey` value, there needs a newline at the end of string. Otherwise, multiple runs of pulumi will want to recreate the `cloudfront.PublicKey` resource.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
}

// NewPublicKey registers a new resource with the given unique name, arguments, and options.
func NewPublicKey(ctx *pulumi.Context,
	name string, args *PublicKeyArgs, opts ...pulumi.ResourceOption) (*PublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EncodedKey == nil {
		return nil, errors.New("invalid value for required argument 'EncodedKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicKey
	err := ctx.RegisterResource("aws:cloudfront/publicKey:PublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicKey gets an existing PublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicKeyState, opts ...pulumi.ResourceOption) (*PublicKey, error) {
	var resource PublicKey
	err := ctx.ReadResource("aws:cloudfront/publicKey:PublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicKey resources.
type publicKeyState struct {
	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference *string `pulumi:"callerReference"`
	// An optional comment about the public key.
	Comment *string `pulumi:"comment"`
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey *string `pulumi:"encodedKey"`
	// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
	Etag *string `pulumi:"etag"`
	// The name for the public key. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name for the public key. Conflicts with `name`.
	//
	// **NOTE:** When setting `encodedKey` value, there needs a newline at the end of string. Otherwise, multiple runs of pulumi will want to recreate the `cloudfront.PublicKey` resource.
	NamePrefix *string `pulumi:"namePrefix"`
}

type PublicKeyState struct {
	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference pulumi.StringPtrInput
	// An optional comment about the public key.
	Comment pulumi.StringPtrInput
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey pulumi.StringPtrInput
	// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringPtrInput
	// The name for the public key. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name for the public key. Conflicts with `name`.
	//
	// **NOTE:** When setting `encodedKey` value, there needs a newline at the end of string. Otherwise, multiple runs of pulumi will want to recreate the `cloudfront.PublicKey` resource.
	NamePrefix pulumi.StringPtrInput
}

func (PublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicKeyState)(nil)).Elem()
}

type publicKeyArgs struct {
	// An optional comment about the public key.
	Comment *string `pulumi:"comment"`
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey string `pulumi:"encodedKey"`
	// The name for the public key. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name for the public key. Conflicts with `name`.
	//
	// **NOTE:** When setting `encodedKey` value, there needs a newline at the end of string. Otherwise, multiple runs of pulumi will want to recreate the `cloudfront.PublicKey` resource.
	NamePrefix *string `pulumi:"namePrefix"`
}

// The set of arguments for constructing a PublicKey resource.
type PublicKeyArgs struct {
	// An optional comment about the public key.
	Comment pulumi.StringPtrInput
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey pulumi.StringInput
	// The name for the public key. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name for the public key. Conflicts with `name`.
	//
	// **NOTE:** When setting `encodedKey` value, there needs a newline at the end of string. Otherwise, multiple runs of pulumi will want to recreate the `cloudfront.PublicKey` resource.
	NamePrefix pulumi.StringPtrInput
}

func (PublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicKeyArgs)(nil)).Elem()
}

type PublicKeyInput interface {
	pulumi.Input

	ToPublicKeyOutput() PublicKeyOutput
	ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput
}

func (*PublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicKey)(nil)).Elem()
}

func (i *PublicKey) ToPublicKeyOutput() PublicKeyOutput {
	return i.ToPublicKeyOutputWithContext(context.Background())
}

func (i *PublicKey) ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicKeyOutput)
}

// PublicKeyArrayInput is an input type that accepts PublicKeyArray and PublicKeyArrayOutput values.
// You can construct a concrete instance of `PublicKeyArrayInput` via:
//
//	PublicKeyArray{ PublicKeyArgs{...} }
type PublicKeyArrayInput interface {
	pulumi.Input

	ToPublicKeyArrayOutput() PublicKeyArrayOutput
	ToPublicKeyArrayOutputWithContext(context.Context) PublicKeyArrayOutput
}

type PublicKeyArray []PublicKeyInput

func (PublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicKey)(nil)).Elem()
}

func (i PublicKeyArray) ToPublicKeyArrayOutput() PublicKeyArrayOutput {
	return i.ToPublicKeyArrayOutputWithContext(context.Background())
}

func (i PublicKeyArray) ToPublicKeyArrayOutputWithContext(ctx context.Context) PublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicKeyArrayOutput)
}

// PublicKeyMapInput is an input type that accepts PublicKeyMap and PublicKeyMapOutput values.
// You can construct a concrete instance of `PublicKeyMapInput` via:
//
//	PublicKeyMap{ "key": PublicKeyArgs{...} }
type PublicKeyMapInput interface {
	pulumi.Input

	ToPublicKeyMapOutput() PublicKeyMapOutput
	ToPublicKeyMapOutputWithContext(context.Context) PublicKeyMapOutput
}

type PublicKeyMap map[string]PublicKeyInput

func (PublicKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicKey)(nil)).Elem()
}

func (i PublicKeyMap) ToPublicKeyMapOutput() PublicKeyMapOutput {
	return i.ToPublicKeyMapOutputWithContext(context.Background())
}

func (i PublicKeyMap) ToPublicKeyMapOutputWithContext(ctx context.Context) PublicKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicKeyMapOutput)
}

type PublicKeyOutput struct{ *pulumi.OutputState }

func (PublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicKey)(nil)).Elem()
}

func (o PublicKeyOutput) ToPublicKeyOutput() PublicKeyOutput {
	return o
}

func (o PublicKeyOutput) ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput {
	return o
}

// Internal value used by CloudFront to allow future updates to the public key configuration.
func (o PublicKeyOutput) CallerReference() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicKey) pulumi.StringOutput { return v.CallerReference }).(pulumi.StringOutput)
}

// An optional comment about the public key.
func (o PublicKeyOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicKey) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
func (o PublicKeyOutput) EncodedKey() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicKey) pulumi.StringOutput { return v.EncodedKey }).(pulumi.StringOutput)
}

// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
func (o PublicKeyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicKey) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name for the public key. By default generated by this provider.
func (o PublicKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name for the public key. Conflicts with `name`.
//
// **NOTE:** When setting `encodedKey` value, there needs a newline at the end of string. Otherwise, multiple runs of pulumi will want to recreate the `cloudfront.PublicKey` resource.
func (o PublicKeyOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicKey) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

type PublicKeyArrayOutput struct{ *pulumi.OutputState }

func (PublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicKey)(nil)).Elem()
}

func (o PublicKeyArrayOutput) ToPublicKeyArrayOutput() PublicKeyArrayOutput {
	return o
}

func (o PublicKeyArrayOutput) ToPublicKeyArrayOutputWithContext(ctx context.Context) PublicKeyArrayOutput {
	return o
}

func (o PublicKeyArrayOutput) Index(i pulumi.IntInput) PublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicKey {
		return vs[0].([]*PublicKey)[vs[1].(int)]
	}).(PublicKeyOutput)
}

type PublicKeyMapOutput struct{ *pulumi.OutputState }

func (PublicKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicKey)(nil)).Elem()
}

func (o PublicKeyMapOutput) ToPublicKeyMapOutput() PublicKeyMapOutput {
	return o
}

func (o PublicKeyMapOutput) ToPublicKeyMapOutputWithContext(ctx context.Context) PublicKeyMapOutput {
	return o
}

func (o PublicKeyMapOutput) MapIndex(k pulumi.StringInput) PublicKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicKey {
		return vs[0].(map[string]*PublicKey)[vs[1].(string)]
	}).(PublicKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicKeyInput)(nil)).Elem(), &PublicKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicKeyArrayInput)(nil)).Elem(), PublicKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicKeyMapInput)(nil)).Elem(), PublicKeyMap{})
	pulumi.RegisterOutputType(PublicKeyOutput{})
	pulumi.RegisterOutputType(PublicKeyArrayOutput{})
	pulumi.RegisterOutputType(PublicKeyMapOutput{})
}
