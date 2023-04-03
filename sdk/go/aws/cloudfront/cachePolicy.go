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
// The following example below creates a CloudFront cache policy.
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
//			_, err := cloudfront.NewCachePolicy(ctx, "example", &cloudfront.CachePolicyArgs{
//				Comment:    pulumi.String("test comment"),
//				DefaultTtl: pulumi.Int(50),
//				MaxTtl:     pulumi.Int(100),
//				MinTtl:     pulumi.Int(1),
//				ParametersInCacheKeyAndForwardedToOrigin: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs{
//					CookiesConfig: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigArgs{
//						CookieBehavior: pulumi.String("whitelist"),
//						Cookies: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesArgs{
//							Items: pulumi.StringArray{
//								pulumi.String("example"),
//							},
//						},
//					},
//					HeadersConfig: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigArgs{
//						HeaderBehavior: pulumi.String("whitelist"),
//						Headers: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersArgs{
//							Items: pulumi.StringArray{
//								pulumi.String("example"),
//							},
//						},
//					},
//					QueryStringsConfig: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigArgs{
//						QueryStringBehavior: pulumi.String("whitelist"),
//						QueryStrings: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsArgs{
//							Items: pulumi.StringArray{
//								pulumi.String("example"),
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
// Cloudfront Cache Policies can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:cloudfront/cachePolicy:CachePolicy policy 658327ea-f89d-4fab-a63d-7e88639e58f6
//
// ```
type CachePolicy struct {
	pulumi.CustomResourceState

	// A comment to describe the cache policy.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// The current version of the cache policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl pulumi.IntPtrOutput `pulumi:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTtl pulumi.IntPtrOutput `pulumi:"minTtl"`
	// A unique name to identify the cache policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigin CachePolicyParametersInCacheKeyAndForwardedToOriginOutput `pulumi:"parametersInCacheKeyAndForwardedToOrigin"`
}

// NewCachePolicy registers a new resource with the given unique name, arguments, and options.
func NewCachePolicy(ctx *pulumi.Context,
	name string, args *CachePolicyArgs, opts ...pulumi.ResourceOption) (*CachePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParametersInCacheKeyAndForwardedToOrigin == nil {
		return nil, errors.New("invalid value for required argument 'ParametersInCacheKeyAndForwardedToOrigin'")
	}
	var resource CachePolicy
	err := ctx.RegisterResource("aws:cloudfront/cachePolicy:CachePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCachePolicy gets an existing CachePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCachePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CachePolicyState, opts ...pulumi.ResourceOption) (*CachePolicy, error) {
	var resource CachePolicy
	err := ctx.ReadResource("aws:cloudfront/cachePolicy:CachePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CachePolicy resources.
type cachePolicyState struct {
	// A comment to describe the cache policy.
	Comment *string `pulumi:"comment"`
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The current version of the cache policy.
	Etag *string `pulumi:"etag"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl *int `pulumi:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTtl *int `pulumi:"minTtl"`
	// A unique name to identify the cache policy.
	Name *string `pulumi:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigin *CachePolicyParametersInCacheKeyAndForwardedToOrigin `pulumi:"parametersInCacheKeyAndForwardedToOrigin"`
}

type CachePolicyState struct {
	// A comment to describe the cache policy.
	Comment pulumi.StringPtrInput
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTtl pulumi.IntPtrInput
	// The current version of the cache policy.
	Etag pulumi.StringPtrInput
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl pulumi.IntPtrInput
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTtl pulumi.IntPtrInput
	// A unique name to identify the cache policy.
	Name pulumi.StringPtrInput
	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigin CachePolicyParametersInCacheKeyAndForwardedToOriginPtrInput
}

func (CachePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*cachePolicyState)(nil)).Elem()
}

type cachePolicyArgs struct {
	// A comment to describe the cache policy.
	Comment *string `pulumi:"comment"`
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl *int `pulumi:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTtl *int `pulumi:"minTtl"`
	// A unique name to identify the cache policy.
	Name *string `pulumi:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigin CachePolicyParametersInCacheKeyAndForwardedToOrigin `pulumi:"parametersInCacheKeyAndForwardedToOrigin"`
}

// The set of arguments for constructing a CachePolicy resource.
type CachePolicyArgs struct {
	// A comment to describe the cache policy.
	Comment pulumi.StringPtrInput
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTtl pulumi.IntPtrInput
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl pulumi.IntPtrInput
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTtl pulumi.IntPtrInput
	// A unique name to identify the cache policy.
	Name pulumi.StringPtrInput
	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigin CachePolicyParametersInCacheKeyAndForwardedToOriginInput
}

func (CachePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cachePolicyArgs)(nil)).Elem()
}

type CachePolicyInput interface {
	pulumi.Input

	ToCachePolicyOutput() CachePolicyOutput
	ToCachePolicyOutputWithContext(ctx context.Context) CachePolicyOutput
}

func (*CachePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CachePolicy)(nil)).Elem()
}

func (i *CachePolicy) ToCachePolicyOutput() CachePolicyOutput {
	return i.ToCachePolicyOutputWithContext(context.Background())
}

func (i *CachePolicy) ToCachePolicyOutputWithContext(ctx context.Context) CachePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachePolicyOutput)
}

// CachePolicyArrayInput is an input type that accepts CachePolicyArray and CachePolicyArrayOutput values.
// You can construct a concrete instance of `CachePolicyArrayInput` via:
//
//	CachePolicyArray{ CachePolicyArgs{...} }
type CachePolicyArrayInput interface {
	pulumi.Input

	ToCachePolicyArrayOutput() CachePolicyArrayOutput
	ToCachePolicyArrayOutputWithContext(context.Context) CachePolicyArrayOutput
}

type CachePolicyArray []CachePolicyInput

func (CachePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CachePolicy)(nil)).Elem()
}

func (i CachePolicyArray) ToCachePolicyArrayOutput() CachePolicyArrayOutput {
	return i.ToCachePolicyArrayOutputWithContext(context.Background())
}

func (i CachePolicyArray) ToCachePolicyArrayOutputWithContext(ctx context.Context) CachePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachePolicyArrayOutput)
}

// CachePolicyMapInput is an input type that accepts CachePolicyMap and CachePolicyMapOutput values.
// You can construct a concrete instance of `CachePolicyMapInput` via:
//
//	CachePolicyMap{ "key": CachePolicyArgs{...} }
type CachePolicyMapInput interface {
	pulumi.Input

	ToCachePolicyMapOutput() CachePolicyMapOutput
	ToCachePolicyMapOutputWithContext(context.Context) CachePolicyMapOutput
}

type CachePolicyMap map[string]CachePolicyInput

func (CachePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CachePolicy)(nil)).Elem()
}

func (i CachePolicyMap) ToCachePolicyMapOutput() CachePolicyMapOutput {
	return i.ToCachePolicyMapOutputWithContext(context.Background())
}

func (i CachePolicyMap) ToCachePolicyMapOutputWithContext(ctx context.Context) CachePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachePolicyMapOutput)
}

type CachePolicyOutput struct{ *pulumi.OutputState }

func (CachePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CachePolicy)(nil)).Elem()
}

func (o CachePolicyOutput) ToCachePolicyOutput() CachePolicyOutput {
	return o
}

func (o CachePolicyOutput) ToCachePolicyOutputWithContext(ctx context.Context) CachePolicyOutput {
	return o
}

// A comment to describe the cache policy.
func (o CachePolicyOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
func (o CachePolicyOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.IntPtrOutput { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

// The current version of the cache policy.
func (o CachePolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
func (o CachePolicyOutput) MaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.IntPtrOutput { return v.MaxTtl }).(pulumi.IntPtrOutput)
}

// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
func (o CachePolicyOutput) MinTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.IntPtrOutput { return v.MinTtl }).(pulumi.IntPtrOutput)
}

// A unique name to identify the cache policy.
func (o CachePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
func (o CachePolicyOutput) ParametersInCacheKeyAndForwardedToOrigin() CachePolicyParametersInCacheKeyAndForwardedToOriginOutput {
	return o.ApplyT(func(v *CachePolicy) CachePolicyParametersInCacheKeyAndForwardedToOriginOutput {
		return v.ParametersInCacheKeyAndForwardedToOrigin
	}).(CachePolicyParametersInCacheKeyAndForwardedToOriginOutput)
}

type CachePolicyArrayOutput struct{ *pulumi.OutputState }

func (CachePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CachePolicy)(nil)).Elem()
}

func (o CachePolicyArrayOutput) ToCachePolicyArrayOutput() CachePolicyArrayOutput {
	return o
}

func (o CachePolicyArrayOutput) ToCachePolicyArrayOutputWithContext(ctx context.Context) CachePolicyArrayOutput {
	return o
}

func (o CachePolicyArrayOutput) Index(i pulumi.IntInput) CachePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CachePolicy {
		return vs[0].([]*CachePolicy)[vs[1].(int)]
	}).(CachePolicyOutput)
}

type CachePolicyMapOutput struct{ *pulumi.OutputState }

func (CachePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CachePolicy)(nil)).Elem()
}

func (o CachePolicyMapOutput) ToCachePolicyMapOutput() CachePolicyMapOutput {
	return o
}

func (o CachePolicyMapOutput) ToCachePolicyMapOutputWithContext(ctx context.Context) CachePolicyMapOutput {
	return o
}

func (o CachePolicyMapOutput) MapIndex(k pulumi.StringInput) CachePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CachePolicy {
		return vs[0].(map[string]*CachePolicy)[vs[1].(string)]
	}).(CachePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CachePolicyInput)(nil)).Elem(), &CachePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*CachePolicyArrayInput)(nil)).Elem(), CachePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CachePolicyMapInput)(nil)).Elem(), CachePolicyMap{})
	pulumi.RegisterOutputType(CachePolicyOutput{})
	pulumi.RegisterOutputType(CachePolicyArrayOutput{})
	pulumi.RegisterOutputType(CachePolicyMapOutput{})
}
