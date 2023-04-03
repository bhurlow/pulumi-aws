// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudFront Field-level Encryption Config resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudfront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfront.NewFieldLevelEncryptionConfig(ctx, "test", &cloudfront.FieldLevelEncryptionConfigArgs{
//				Comment: pulumi.String("test comment"),
//				ContentTypeProfileConfig: &cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigArgs{
//					ForwardWhenContentTypeIsUnknown: pulumi.Bool(true),
//					ContentTypeProfiles: &cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesArgs{
//						Items: cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItemArray{
//							&cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItemArgs{
//								ContentType: pulumi.String("application/x-www-form-urlencoded"),
//								Format:      pulumi.String("URLEncoded"),
//							},
//						},
//					},
//				},
//				QueryArgProfileConfig: &cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigArgs{
//					ForwardWhenQueryArgProfileIsUnknown: pulumi.Bool(true),
//					QueryArgProfiles: &cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesArgs{
//						Items: cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItemArray{
//							&cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItemArgs{
//								ProfileId: pulumi.Any(aws_cloudfront_field_level_encryption_profile.Test.Id),
//								QueryArg:  pulumi.String("Arg1"),
//							},
//						},
//					},
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
// Cloudfront Field Level Encryption Config can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import aws:cloudfront/fieldLevelEncryptionConfig:FieldLevelEncryptionConfig config E74FTE3AEXAMPLE
//
// ```
type FieldLevelEncryptionConfig struct {
	pulumi.CustomResourceState

	// Internal value used by CloudFront to allow future updates to the Field Level Encryption Config.
	CallerReference pulumi.StringOutput `pulumi:"callerReference"`
	// An optional comment about the Field Level Encryption Config.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig FieldLevelEncryptionConfigContentTypeProfileConfigOutput `pulumi:"contentTypeProfileConfig"`
	// The current version of the Field Level Encryption Config. For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig FieldLevelEncryptionConfigQueryArgProfileConfigOutput `pulumi:"queryArgProfileConfig"`
}

// NewFieldLevelEncryptionConfig registers a new resource with the given unique name, arguments, and options.
func NewFieldLevelEncryptionConfig(ctx *pulumi.Context,
	name string, args *FieldLevelEncryptionConfigArgs, opts ...pulumi.ResourceOption) (*FieldLevelEncryptionConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentTypeProfileConfig == nil {
		return nil, errors.New("invalid value for required argument 'ContentTypeProfileConfig'")
	}
	if args.QueryArgProfileConfig == nil {
		return nil, errors.New("invalid value for required argument 'QueryArgProfileConfig'")
	}
	var resource FieldLevelEncryptionConfig
	err := ctx.RegisterResource("aws:cloudfront/fieldLevelEncryptionConfig:FieldLevelEncryptionConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFieldLevelEncryptionConfig gets an existing FieldLevelEncryptionConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFieldLevelEncryptionConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FieldLevelEncryptionConfigState, opts ...pulumi.ResourceOption) (*FieldLevelEncryptionConfig, error) {
	var resource FieldLevelEncryptionConfig
	err := ctx.ReadResource("aws:cloudfront/fieldLevelEncryptionConfig:FieldLevelEncryptionConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FieldLevelEncryptionConfig resources.
type fieldLevelEncryptionConfigState struct {
	// Internal value used by CloudFront to allow future updates to the Field Level Encryption Config.
	CallerReference *string `pulumi:"callerReference"`
	// An optional comment about the Field Level Encryption Config.
	Comment *string `pulumi:"comment"`
	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig *FieldLevelEncryptionConfigContentTypeProfileConfig `pulumi:"contentTypeProfileConfig"`
	// The current version of the Field Level Encryption Config. For example: `E2QWRUHAPOMQZL`.
	Etag *string `pulumi:"etag"`
	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig *FieldLevelEncryptionConfigQueryArgProfileConfig `pulumi:"queryArgProfileConfig"`
}

type FieldLevelEncryptionConfigState struct {
	// Internal value used by CloudFront to allow future updates to the Field Level Encryption Config.
	CallerReference pulumi.StringPtrInput
	// An optional comment about the Field Level Encryption Config.
	Comment pulumi.StringPtrInput
	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig FieldLevelEncryptionConfigContentTypeProfileConfigPtrInput
	// The current version of the Field Level Encryption Config. For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringPtrInput
	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig FieldLevelEncryptionConfigQueryArgProfileConfigPtrInput
}

func (FieldLevelEncryptionConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*fieldLevelEncryptionConfigState)(nil)).Elem()
}

type fieldLevelEncryptionConfigArgs struct {
	// An optional comment about the Field Level Encryption Config.
	Comment *string `pulumi:"comment"`
	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig FieldLevelEncryptionConfigContentTypeProfileConfig `pulumi:"contentTypeProfileConfig"`
	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig FieldLevelEncryptionConfigQueryArgProfileConfig `pulumi:"queryArgProfileConfig"`
}

// The set of arguments for constructing a FieldLevelEncryptionConfig resource.
type FieldLevelEncryptionConfigArgs struct {
	// An optional comment about the Field Level Encryption Config.
	Comment pulumi.StringPtrInput
	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig FieldLevelEncryptionConfigContentTypeProfileConfigInput
	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig FieldLevelEncryptionConfigQueryArgProfileConfigInput
}

func (FieldLevelEncryptionConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fieldLevelEncryptionConfigArgs)(nil)).Elem()
}

type FieldLevelEncryptionConfigInput interface {
	pulumi.Input

	ToFieldLevelEncryptionConfigOutput() FieldLevelEncryptionConfigOutput
	ToFieldLevelEncryptionConfigOutputWithContext(ctx context.Context) FieldLevelEncryptionConfigOutput
}

func (*FieldLevelEncryptionConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**FieldLevelEncryptionConfig)(nil)).Elem()
}

func (i *FieldLevelEncryptionConfig) ToFieldLevelEncryptionConfigOutput() FieldLevelEncryptionConfigOutput {
	return i.ToFieldLevelEncryptionConfigOutputWithContext(context.Background())
}

func (i *FieldLevelEncryptionConfig) ToFieldLevelEncryptionConfigOutputWithContext(ctx context.Context) FieldLevelEncryptionConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldLevelEncryptionConfigOutput)
}

// FieldLevelEncryptionConfigArrayInput is an input type that accepts FieldLevelEncryptionConfigArray and FieldLevelEncryptionConfigArrayOutput values.
// You can construct a concrete instance of `FieldLevelEncryptionConfigArrayInput` via:
//
//	FieldLevelEncryptionConfigArray{ FieldLevelEncryptionConfigArgs{...} }
type FieldLevelEncryptionConfigArrayInput interface {
	pulumi.Input

	ToFieldLevelEncryptionConfigArrayOutput() FieldLevelEncryptionConfigArrayOutput
	ToFieldLevelEncryptionConfigArrayOutputWithContext(context.Context) FieldLevelEncryptionConfigArrayOutput
}

type FieldLevelEncryptionConfigArray []FieldLevelEncryptionConfigInput

func (FieldLevelEncryptionConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FieldLevelEncryptionConfig)(nil)).Elem()
}

func (i FieldLevelEncryptionConfigArray) ToFieldLevelEncryptionConfigArrayOutput() FieldLevelEncryptionConfigArrayOutput {
	return i.ToFieldLevelEncryptionConfigArrayOutputWithContext(context.Background())
}

func (i FieldLevelEncryptionConfigArray) ToFieldLevelEncryptionConfigArrayOutputWithContext(ctx context.Context) FieldLevelEncryptionConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldLevelEncryptionConfigArrayOutput)
}

// FieldLevelEncryptionConfigMapInput is an input type that accepts FieldLevelEncryptionConfigMap and FieldLevelEncryptionConfigMapOutput values.
// You can construct a concrete instance of `FieldLevelEncryptionConfigMapInput` via:
//
//	FieldLevelEncryptionConfigMap{ "key": FieldLevelEncryptionConfigArgs{...} }
type FieldLevelEncryptionConfigMapInput interface {
	pulumi.Input

	ToFieldLevelEncryptionConfigMapOutput() FieldLevelEncryptionConfigMapOutput
	ToFieldLevelEncryptionConfigMapOutputWithContext(context.Context) FieldLevelEncryptionConfigMapOutput
}

type FieldLevelEncryptionConfigMap map[string]FieldLevelEncryptionConfigInput

func (FieldLevelEncryptionConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FieldLevelEncryptionConfig)(nil)).Elem()
}

func (i FieldLevelEncryptionConfigMap) ToFieldLevelEncryptionConfigMapOutput() FieldLevelEncryptionConfigMapOutput {
	return i.ToFieldLevelEncryptionConfigMapOutputWithContext(context.Background())
}

func (i FieldLevelEncryptionConfigMap) ToFieldLevelEncryptionConfigMapOutputWithContext(ctx context.Context) FieldLevelEncryptionConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldLevelEncryptionConfigMapOutput)
}

type FieldLevelEncryptionConfigOutput struct{ *pulumi.OutputState }

func (FieldLevelEncryptionConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FieldLevelEncryptionConfig)(nil)).Elem()
}

func (o FieldLevelEncryptionConfigOutput) ToFieldLevelEncryptionConfigOutput() FieldLevelEncryptionConfigOutput {
	return o
}

func (o FieldLevelEncryptionConfigOutput) ToFieldLevelEncryptionConfigOutputWithContext(ctx context.Context) FieldLevelEncryptionConfigOutput {
	return o
}

// Internal value used by CloudFront to allow future updates to the Field Level Encryption Config.
func (o FieldLevelEncryptionConfigOutput) CallerReference() pulumi.StringOutput {
	return o.ApplyT(func(v *FieldLevelEncryptionConfig) pulumi.StringOutput { return v.CallerReference }).(pulumi.StringOutput)
}

// An optional comment about the Field Level Encryption Config.
func (o FieldLevelEncryptionConfigOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FieldLevelEncryptionConfig) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
func (o FieldLevelEncryptionConfigOutput) ContentTypeProfileConfig() FieldLevelEncryptionConfigContentTypeProfileConfigOutput {
	return o.ApplyT(func(v *FieldLevelEncryptionConfig) FieldLevelEncryptionConfigContentTypeProfileConfigOutput {
		return v.ContentTypeProfileConfig
	}).(FieldLevelEncryptionConfigContentTypeProfileConfigOutput)
}

// The current version of the Field Level Encryption Config. For example: `E2QWRUHAPOMQZL`.
func (o FieldLevelEncryptionConfigOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FieldLevelEncryptionConfig) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
func (o FieldLevelEncryptionConfigOutput) QueryArgProfileConfig() FieldLevelEncryptionConfigQueryArgProfileConfigOutput {
	return o.ApplyT(func(v *FieldLevelEncryptionConfig) FieldLevelEncryptionConfigQueryArgProfileConfigOutput {
		return v.QueryArgProfileConfig
	}).(FieldLevelEncryptionConfigQueryArgProfileConfigOutput)
}

type FieldLevelEncryptionConfigArrayOutput struct{ *pulumi.OutputState }

func (FieldLevelEncryptionConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FieldLevelEncryptionConfig)(nil)).Elem()
}

func (o FieldLevelEncryptionConfigArrayOutput) ToFieldLevelEncryptionConfigArrayOutput() FieldLevelEncryptionConfigArrayOutput {
	return o
}

func (o FieldLevelEncryptionConfigArrayOutput) ToFieldLevelEncryptionConfigArrayOutputWithContext(ctx context.Context) FieldLevelEncryptionConfigArrayOutput {
	return o
}

func (o FieldLevelEncryptionConfigArrayOutput) Index(i pulumi.IntInput) FieldLevelEncryptionConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FieldLevelEncryptionConfig {
		return vs[0].([]*FieldLevelEncryptionConfig)[vs[1].(int)]
	}).(FieldLevelEncryptionConfigOutput)
}

type FieldLevelEncryptionConfigMapOutput struct{ *pulumi.OutputState }

func (FieldLevelEncryptionConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FieldLevelEncryptionConfig)(nil)).Elem()
}

func (o FieldLevelEncryptionConfigMapOutput) ToFieldLevelEncryptionConfigMapOutput() FieldLevelEncryptionConfigMapOutput {
	return o
}

func (o FieldLevelEncryptionConfigMapOutput) ToFieldLevelEncryptionConfigMapOutputWithContext(ctx context.Context) FieldLevelEncryptionConfigMapOutput {
	return o
}

func (o FieldLevelEncryptionConfigMapOutput) MapIndex(k pulumi.StringInput) FieldLevelEncryptionConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FieldLevelEncryptionConfig {
		return vs[0].(map[string]*FieldLevelEncryptionConfig)[vs[1].(string)]
	}).(FieldLevelEncryptionConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FieldLevelEncryptionConfigInput)(nil)).Elem(), &FieldLevelEncryptionConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*FieldLevelEncryptionConfigArrayInput)(nil)).Elem(), FieldLevelEncryptionConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FieldLevelEncryptionConfigMapInput)(nil)).Elem(), FieldLevelEncryptionConfigMap{})
	pulumi.RegisterOutputType(FieldLevelEncryptionConfigOutput{})
	pulumi.RegisterOutputType(FieldLevelEncryptionConfigArrayOutput{})
	pulumi.RegisterOutputType(FieldLevelEncryptionConfigMapOutput{})
}
