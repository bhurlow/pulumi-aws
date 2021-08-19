// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM SAML provider.
//
// ## Import
//
// IAM SAML Providers can be imported using the `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:iam/samlProvider:SamlProvider default arn:aws:iam::123456789012:saml-provider/SAMLADFS
// ```
type SamlProvider struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS for this provider.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the provider to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// An XML document generated by an identity provider that supports SAML 2.0.
	SamlMetadataDocument pulumi.StringOutput `pulumi:"samlMetadataDocument"`
	// Map of resource tags for the IAM SAML provider. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The expiration date and time for the SAML provider in RFC1123 format, e.g. `Mon, 02 Jan 2006 15:04:05 MST`.
	ValidUntil pulumi.StringOutput `pulumi:"validUntil"`
}

// NewSamlProvider registers a new resource with the given unique name, arguments, and options.
func NewSamlProvider(ctx *pulumi.Context,
	name string, args *SamlProviderArgs, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SamlMetadataDocument == nil {
		return nil, errors.New("invalid value for required argument 'SamlMetadataDocument'")
	}
	var resource SamlProvider
	err := ctx.RegisterResource("aws:iam/samlProvider:SamlProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlProvider gets an existing SamlProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlProviderState, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	var resource SamlProvider
	err := ctx.ReadResource("aws:iam/samlProvider:SamlProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlProvider resources.
type samlProviderState struct {
	// The ARN assigned by AWS for this provider.
	Arn *string `pulumi:"arn"`
	// The name of the provider to create.
	Name *string `pulumi:"name"`
	// An XML document generated by an identity provider that supports SAML 2.0.
	SamlMetadataDocument *string `pulumi:"samlMetadataDocument"`
	// Map of resource tags for the IAM SAML provider. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The expiration date and time for the SAML provider in RFC1123 format, e.g. `Mon, 02 Jan 2006 15:04:05 MST`.
	ValidUntil *string `pulumi:"validUntil"`
}

type SamlProviderState struct {
	// The ARN assigned by AWS for this provider.
	Arn pulumi.StringPtrInput
	// The name of the provider to create.
	Name pulumi.StringPtrInput
	// An XML document generated by an identity provider that supports SAML 2.0.
	SamlMetadataDocument pulumi.StringPtrInput
	// Map of resource tags for the IAM SAML provider. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The expiration date and time for the SAML provider in RFC1123 format, e.g. `Mon, 02 Jan 2006 15:04:05 MST`.
	ValidUntil pulumi.StringPtrInput
}

func (SamlProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderState)(nil)).Elem()
}

type samlProviderArgs struct {
	// The name of the provider to create.
	Name *string `pulumi:"name"`
	// An XML document generated by an identity provider that supports SAML 2.0.
	SamlMetadataDocument string `pulumi:"samlMetadataDocument"`
	// Map of resource tags for the IAM SAML provider. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SamlProvider resource.
type SamlProviderArgs struct {
	// The name of the provider to create.
	Name pulumi.StringPtrInput
	// An XML document generated by an identity provider that supports SAML 2.0.
	SamlMetadataDocument pulumi.StringInput
	// Map of resource tags for the IAM SAML provider. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (SamlProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderArgs)(nil)).Elem()
}

type SamlProviderInput interface {
	pulumi.Input

	ToSamlProviderOutput() SamlProviderOutput
	ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput
}

func (*SamlProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*SamlProvider)(nil))
}

func (i *SamlProvider) ToSamlProviderOutput() SamlProviderOutput {
	return i.ToSamlProviderOutputWithContext(context.Background())
}

func (i *SamlProvider) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderOutput)
}

func (i *SamlProvider) ToSamlProviderPtrOutput() SamlProviderPtrOutput {
	return i.ToSamlProviderPtrOutputWithContext(context.Background())
}

func (i *SamlProvider) ToSamlProviderPtrOutputWithContext(ctx context.Context) SamlProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderPtrOutput)
}

type SamlProviderPtrInput interface {
	pulumi.Input

	ToSamlProviderPtrOutput() SamlProviderPtrOutput
	ToSamlProviderPtrOutputWithContext(ctx context.Context) SamlProviderPtrOutput
}

type samlProviderPtrType SamlProviderArgs

func (*samlProviderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil))
}

func (i *samlProviderPtrType) ToSamlProviderPtrOutput() SamlProviderPtrOutput {
	return i.ToSamlProviderPtrOutputWithContext(context.Background())
}

func (i *samlProviderPtrType) ToSamlProviderPtrOutputWithContext(ctx context.Context) SamlProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderPtrOutput)
}

// SamlProviderArrayInput is an input type that accepts SamlProviderArray and SamlProviderArrayOutput values.
// You can construct a concrete instance of `SamlProviderArrayInput` via:
//
//          SamlProviderArray{ SamlProviderArgs{...} }
type SamlProviderArrayInput interface {
	pulumi.Input

	ToSamlProviderArrayOutput() SamlProviderArrayOutput
	ToSamlProviderArrayOutputWithContext(context.Context) SamlProviderArrayOutput
}

type SamlProviderArray []SamlProviderInput

func (SamlProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderArray) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return i.ToSamlProviderArrayOutputWithContext(context.Background())
}

func (i SamlProviderArray) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderArrayOutput)
}

// SamlProviderMapInput is an input type that accepts SamlProviderMap and SamlProviderMapOutput values.
// You can construct a concrete instance of `SamlProviderMapInput` via:
//
//          SamlProviderMap{ "key": SamlProviderArgs{...} }
type SamlProviderMapInput interface {
	pulumi.Input

	ToSamlProviderMapOutput() SamlProviderMapOutput
	ToSamlProviderMapOutputWithContext(context.Context) SamlProviderMapOutput
}

type SamlProviderMap map[string]SamlProviderInput

func (SamlProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderMap) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return i.ToSamlProviderMapOutputWithContext(context.Background())
}

func (i SamlProviderMap) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderMapOutput)
}

type SamlProviderOutput struct{ *pulumi.OutputState }

func (SamlProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SamlProvider)(nil))
}

func (o SamlProviderOutput) ToSamlProviderOutput() SamlProviderOutput {
	return o
}

func (o SamlProviderOutput) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return o
}

func (o SamlProviderOutput) ToSamlProviderPtrOutput() SamlProviderPtrOutput {
	return o.ToSamlProviderPtrOutputWithContext(context.Background())
}

func (o SamlProviderOutput) ToSamlProviderPtrOutputWithContext(ctx context.Context) SamlProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SamlProvider) *SamlProvider {
		return &v
	}).(SamlProviderPtrOutput)
}

type SamlProviderPtrOutput struct{ *pulumi.OutputState }

func (SamlProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil))
}

func (o SamlProviderPtrOutput) ToSamlProviderPtrOutput() SamlProviderPtrOutput {
	return o
}

func (o SamlProviderPtrOutput) ToSamlProviderPtrOutputWithContext(ctx context.Context) SamlProviderPtrOutput {
	return o
}

func (o SamlProviderPtrOutput) Elem() SamlProviderOutput {
	return o.ApplyT(func(v *SamlProvider) SamlProvider {
		if v != nil {
			return *v
		}
		var ret SamlProvider
		return ret
	}).(SamlProviderOutput)
}

type SamlProviderArrayOutput struct{ *pulumi.OutputState }

func (SamlProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SamlProvider)(nil))
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) Index(i pulumi.IntInput) SamlProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SamlProvider {
		return vs[0].([]SamlProvider)[vs[1].(int)]
	}).(SamlProviderOutput)
}

type SamlProviderMapOutput struct{ *pulumi.OutputState }

func (SamlProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SamlProvider)(nil))
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) MapIndex(k pulumi.StringInput) SamlProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SamlProvider {
		return vs[0].(map[string]SamlProvider)[vs[1].(string)]
	}).(SamlProviderOutput)
}

func init() {
	pulumi.RegisterOutputType(SamlProviderOutput{})
	pulumi.RegisterOutputType(SamlProviderPtrOutput{})
	pulumi.RegisterOutputType(SamlProviderArrayOutput{})
	pulumi.RegisterOutputType(SamlProviderMapOutput{})
}
