// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// The following example below creates a CloudFront origin request policy.
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
//			_, err := cloudfront.NewOriginRequestPolicy(ctx, "example", &cloudfront.OriginRequestPolicyArgs{
//				Comment: pulumi.String("example comment"),
//				CookiesConfig: &cloudfront.OriginRequestPolicyCookiesConfigArgs{
//					CookieBehavior: pulumi.String("whitelist"),
//					Cookies: &cloudfront.OriginRequestPolicyCookiesConfigCookiesArgs{
//						Items: pulumi.StringArray{
//							pulumi.String("example"),
//						},
//					},
//				},
//				HeadersConfig: &cloudfront.OriginRequestPolicyHeadersConfigArgs{
//					HeaderBehavior: pulumi.String("whitelist"),
//					Headers: &cloudfront.OriginRequestPolicyHeadersConfigHeadersArgs{
//						Items: pulumi.StringArray{
//							pulumi.String("example"),
//						},
//					},
//				},
//				QueryStringsConfig: &cloudfront.OriginRequestPolicyQueryStringsConfigArgs{
//					QueryStringBehavior: pulumi.String("whitelist"),
//					QueryStrings: &cloudfront.OriginRequestPolicyQueryStringsConfigQueryStringsArgs{
//						Items: pulumi.StringArray{
//							pulumi.String("example"),
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
// Cloudfront Origin Request Policies can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import aws:cloudfront/originRequestPolicy:OriginRequestPolicy policy ccca32ef-dce3-4df3-80df-1bd3000bc4d3
//
// ```
type OriginRequestPolicy struct {
	pulumi.CustomResourceState

	// Comment to describe the origin request policy.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig OriginRequestPolicyCookiesConfigOutput `pulumi:"cookiesConfig"`
	// The current version of the origin request policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig OriginRequestPolicyHeadersConfigOutput `pulumi:"headersConfig"`
	// Unique name to identify the origin request policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig OriginRequestPolicyQueryStringsConfigOutput `pulumi:"queryStringsConfig"`
}

// NewOriginRequestPolicy registers a new resource with the given unique name, arguments, and options.
func NewOriginRequestPolicy(ctx *pulumi.Context,
	name string, args *OriginRequestPolicyArgs, opts ...pulumi.ResourceOption) (*OriginRequestPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CookiesConfig == nil {
		return nil, errors.New("invalid value for required argument 'CookiesConfig'")
	}
	if args.HeadersConfig == nil {
		return nil, errors.New("invalid value for required argument 'HeadersConfig'")
	}
	if args.QueryStringsConfig == nil {
		return nil, errors.New("invalid value for required argument 'QueryStringsConfig'")
	}
	var resource OriginRequestPolicy
	err := ctx.RegisterResource("aws:cloudfront/originRequestPolicy:OriginRequestPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOriginRequestPolicy gets an existing OriginRequestPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginRequestPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginRequestPolicyState, opts ...pulumi.ResourceOption) (*OriginRequestPolicy, error) {
	var resource OriginRequestPolicy
	err := ctx.ReadResource("aws:cloudfront/originRequestPolicy:OriginRequestPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OriginRequestPolicy resources.
type originRequestPolicyState struct {
	// Comment to describe the origin request policy.
	Comment *string `pulumi:"comment"`
	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig *OriginRequestPolicyCookiesConfig `pulumi:"cookiesConfig"`
	// The current version of the origin request policy.
	Etag *string `pulumi:"etag"`
	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig *OriginRequestPolicyHeadersConfig `pulumi:"headersConfig"`
	// Unique name to identify the origin request policy.
	Name *string `pulumi:"name"`
	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig *OriginRequestPolicyQueryStringsConfig `pulumi:"queryStringsConfig"`
}

type OriginRequestPolicyState struct {
	// Comment to describe the origin request policy.
	Comment pulumi.StringPtrInput
	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig OriginRequestPolicyCookiesConfigPtrInput
	// The current version of the origin request policy.
	Etag pulumi.StringPtrInput
	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig OriginRequestPolicyHeadersConfigPtrInput
	// Unique name to identify the origin request policy.
	Name pulumi.StringPtrInput
	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig OriginRequestPolicyQueryStringsConfigPtrInput
}

func (OriginRequestPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*originRequestPolicyState)(nil)).Elem()
}

type originRequestPolicyArgs struct {
	// Comment to describe the origin request policy.
	Comment *string `pulumi:"comment"`
	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig OriginRequestPolicyCookiesConfig `pulumi:"cookiesConfig"`
	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig OriginRequestPolicyHeadersConfig `pulumi:"headersConfig"`
	// Unique name to identify the origin request policy.
	Name *string `pulumi:"name"`
	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig OriginRequestPolicyQueryStringsConfig `pulumi:"queryStringsConfig"`
}

// The set of arguments for constructing a OriginRequestPolicy resource.
type OriginRequestPolicyArgs struct {
	// Comment to describe the origin request policy.
	Comment pulumi.StringPtrInput
	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig OriginRequestPolicyCookiesConfigInput
	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig OriginRequestPolicyHeadersConfigInput
	// Unique name to identify the origin request policy.
	Name pulumi.StringPtrInput
	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig OriginRequestPolicyQueryStringsConfigInput
}

func (OriginRequestPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originRequestPolicyArgs)(nil)).Elem()
}

type OriginRequestPolicyInput interface {
	pulumi.Input

	ToOriginRequestPolicyOutput() OriginRequestPolicyOutput
	ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput
}

func (*OriginRequestPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginRequestPolicy)(nil)).Elem()
}

func (i *OriginRequestPolicy) ToOriginRequestPolicyOutput() OriginRequestPolicyOutput {
	return i.ToOriginRequestPolicyOutputWithContext(context.Background())
}

func (i *OriginRequestPolicy) ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginRequestPolicyOutput)
}

// OriginRequestPolicyArrayInput is an input type that accepts OriginRequestPolicyArray and OriginRequestPolicyArrayOutput values.
// You can construct a concrete instance of `OriginRequestPolicyArrayInput` via:
//
//	OriginRequestPolicyArray{ OriginRequestPolicyArgs{...} }
type OriginRequestPolicyArrayInput interface {
	pulumi.Input

	ToOriginRequestPolicyArrayOutput() OriginRequestPolicyArrayOutput
	ToOriginRequestPolicyArrayOutputWithContext(context.Context) OriginRequestPolicyArrayOutput
}

type OriginRequestPolicyArray []OriginRequestPolicyInput

func (OriginRequestPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OriginRequestPolicy)(nil)).Elem()
}

func (i OriginRequestPolicyArray) ToOriginRequestPolicyArrayOutput() OriginRequestPolicyArrayOutput {
	return i.ToOriginRequestPolicyArrayOutputWithContext(context.Background())
}

func (i OriginRequestPolicyArray) ToOriginRequestPolicyArrayOutputWithContext(ctx context.Context) OriginRequestPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginRequestPolicyArrayOutput)
}

// OriginRequestPolicyMapInput is an input type that accepts OriginRequestPolicyMap and OriginRequestPolicyMapOutput values.
// You can construct a concrete instance of `OriginRequestPolicyMapInput` via:
//
//	OriginRequestPolicyMap{ "key": OriginRequestPolicyArgs{...} }
type OriginRequestPolicyMapInput interface {
	pulumi.Input

	ToOriginRequestPolicyMapOutput() OriginRequestPolicyMapOutput
	ToOriginRequestPolicyMapOutputWithContext(context.Context) OriginRequestPolicyMapOutput
}

type OriginRequestPolicyMap map[string]OriginRequestPolicyInput

func (OriginRequestPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OriginRequestPolicy)(nil)).Elem()
}

func (i OriginRequestPolicyMap) ToOriginRequestPolicyMapOutput() OriginRequestPolicyMapOutput {
	return i.ToOriginRequestPolicyMapOutputWithContext(context.Background())
}

func (i OriginRequestPolicyMap) ToOriginRequestPolicyMapOutputWithContext(ctx context.Context) OriginRequestPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginRequestPolicyMapOutput)
}

type OriginRequestPolicyOutput struct{ *pulumi.OutputState }

func (OriginRequestPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginRequestPolicy)(nil)).Elem()
}

func (o OriginRequestPolicyOutput) ToOriginRequestPolicyOutput() OriginRequestPolicyOutput {
	return o
}

func (o OriginRequestPolicyOutput) ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput {
	return o
}

// Comment to describe the origin request policy.
func (o OriginRequestPolicyOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
func (o OriginRequestPolicyOutput) CookiesConfig() OriginRequestPolicyCookiesConfigOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) OriginRequestPolicyCookiesConfigOutput { return v.CookiesConfig }).(OriginRequestPolicyCookiesConfigOutput)
}

// The current version of the origin request policy.
func (o OriginRequestPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
func (o OriginRequestPolicyOutput) HeadersConfig() OriginRequestPolicyHeadersConfigOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) OriginRequestPolicyHeadersConfigOutput { return v.HeadersConfig }).(OriginRequestPolicyHeadersConfigOutput)
}

// Unique name to identify the origin request policy.
func (o OriginRequestPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
func (o OriginRequestPolicyOutput) QueryStringsConfig() OriginRequestPolicyQueryStringsConfigOutput {
	return o.ApplyT(func(v *OriginRequestPolicy) OriginRequestPolicyQueryStringsConfigOutput { return v.QueryStringsConfig }).(OriginRequestPolicyQueryStringsConfigOutput)
}

type OriginRequestPolicyArrayOutput struct{ *pulumi.OutputState }

func (OriginRequestPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OriginRequestPolicy)(nil)).Elem()
}

func (o OriginRequestPolicyArrayOutput) ToOriginRequestPolicyArrayOutput() OriginRequestPolicyArrayOutput {
	return o
}

func (o OriginRequestPolicyArrayOutput) ToOriginRequestPolicyArrayOutputWithContext(ctx context.Context) OriginRequestPolicyArrayOutput {
	return o
}

func (o OriginRequestPolicyArrayOutput) Index(i pulumi.IntInput) OriginRequestPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OriginRequestPolicy {
		return vs[0].([]*OriginRequestPolicy)[vs[1].(int)]
	}).(OriginRequestPolicyOutput)
}

type OriginRequestPolicyMapOutput struct{ *pulumi.OutputState }

func (OriginRequestPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OriginRequestPolicy)(nil)).Elem()
}

func (o OriginRequestPolicyMapOutput) ToOriginRequestPolicyMapOutput() OriginRequestPolicyMapOutput {
	return o
}

func (o OriginRequestPolicyMapOutput) ToOriginRequestPolicyMapOutputWithContext(ctx context.Context) OriginRequestPolicyMapOutput {
	return o
}

func (o OriginRequestPolicyMapOutput) MapIndex(k pulumi.StringInput) OriginRequestPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OriginRequestPolicy {
		return vs[0].(map[string]*OriginRequestPolicy)[vs[1].(string)]
	}).(OriginRequestPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OriginRequestPolicyInput)(nil)).Elem(), &OriginRequestPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*OriginRequestPolicyArrayInput)(nil)).Elem(), OriginRequestPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OriginRequestPolicyMapInput)(nil)).Elem(), OriginRequestPolicyMap{})
	pulumi.RegisterOutputType(OriginRequestPolicyOutput{})
	pulumi.RegisterOutputType(OriginRequestPolicyArrayOutput{})
	pulumi.RegisterOutputType(OriginRequestPolicyMapOutput{})
}
