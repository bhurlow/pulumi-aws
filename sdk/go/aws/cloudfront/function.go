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

// Provides a CloudFront Function resource. With CloudFront Functions in Amazon CloudFront, you can write lightweight functions in JavaScript for high-scale, latency-sensitive CDN customizations.
//
// See [CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-functions.html)
//
// > **NOTE:** You cannot delete a function if it’s associated with a cache behavior. First, update your distributions to remove the function association from all cache behaviors, then delete the function.
//
// ## Example Usage
// ### Basic Example
//
// ```go
// package main
//
// import (
//
//	"fmt"
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
//			_, err := cloudfront.NewFunction(ctx, "test", &cloudfront.FunctionArgs{
//				Runtime: pulumi.String("cloudfront-js-2.0"),
//				Comment: pulumi.String("my function"),
//				Publish: pulumi.Bool(true),
//				Code:    readFileOrPanic(fmt.Sprintf("%v/function.js", path.Module)),
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
// Using `pulumi import`, import CloudFront Functions using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:cloudfront/function:Function test my_test_function
//
// ```
type Function struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) identifying your CloudFront Function.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Source code of the function
	Code pulumi.StringOutput `pulumi:"code"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// ETag hash of any `LIVE` stage of the function.
	LiveStageEtag pulumi.StringOutput `pulumi:"liveStageEtag"`
	// Unique name for your CloudFront Function.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
	Publish pulumi.BoolPtrOutput `pulumi:"publish"`
	// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
	//
	// The following arguments are optional:
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Code == nil {
		return nil, errors.New("invalid value for required argument 'Code'")
	}
	if args.Runtime == nil {
		return nil, errors.New("invalid value for required argument 'Runtime'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Function
	err := ctx.RegisterResource("aws:cloudfront/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("aws:cloudfront/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// Amazon Resource Name (ARN) identifying your CloudFront Function.
	Arn *string `pulumi:"arn"`
	// Source code of the function
	Code *string `pulumi:"code"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
	Etag *string `pulumi:"etag"`
	// ETag hash of any `LIVE` stage of the function.
	LiveStageEtag *string `pulumi:"liveStageEtag"`
	// Unique name for your CloudFront Function.
	Name *string `pulumi:"name"`
	// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
	Publish *bool `pulumi:"publish"`
	// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
	//
	// The following arguments are optional:
	Runtime *string `pulumi:"runtime"`
	// Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
	Status *string `pulumi:"status"`
}

type FunctionState struct {
	// Amazon Resource Name (ARN) identifying your CloudFront Function.
	Arn pulumi.StringPtrInput
	// Source code of the function
	Code pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
	Etag pulumi.StringPtrInput
	// ETag hash of any `LIVE` stage of the function.
	LiveStageEtag pulumi.StringPtrInput
	// Unique name for your CloudFront Function.
	Name pulumi.StringPtrInput
	// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
	Publish pulumi.BoolPtrInput
	// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
	//
	// The following arguments are optional:
	Runtime pulumi.StringPtrInput
	// Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
	Status pulumi.StringPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// Source code of the function
	Code string `pulumi:"code"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Unique name for your CloudFront Function.
	Name *string `pulumi:"name"`
	// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
	Publish *bool `pulumi:"publish"`
	// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
	//
	// The following arguments are optional:
	Runtime string `pulumi:"runtime"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// Source code of the function
	Code pulumi.StringInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Unique name for your CloudFront Function.
	Name pulumi.StringPtrInput
	// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
	Publish pulumi.BoolPtrInput
	// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
	//
	// The following arguments are optional:
	Runtime pulumi.StringInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

// FunctionArrayInput is an input type that accepts FunctionArray and FunctionArrayOutput values.
// You can construct a concrete instance of `FunctionArrayInput` via:
//
//	FunctionArray{ FunctionArgs{...} }
type FunctionArrayInput interface {
	pulumi.Input

	ToFunctionArrayOutput() FunctionArrayOutput
	ToFunctionArrayOutputWithContext(context.Context) FunctionArrayOutput
}

type FunctionArray []FunctionInput

func (FunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (i FunctionArray) ToFunctionArrayOutput() FunctionArrayOutput {
	return i.ToFunctionArrayOutputWithContext(context.Background())
}

func (i FunctionArray) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionArrayOutput)
}

// FunctionMapInput is an input type that accepts FunctionMap and FunctionMapOutput values.
// You can construct a concrete instance of `FunctionMapInput` via:
//
//	FunctionMap{ "key": FunctionArgs{...} }
type FunctionMapInput interface {
	pulumi.Input

	ToFunctionMapOutput() FunctionMapOutput
	ToFunctionMapOutputWithContext(context.Context) FunctionMapOutput
}

type FunctionMap map[string]FunctionInput

func (FunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (i FunctionMap) ToFunctionMapOutput() FunctionMapOutput {
	return i.ToFunctionMapOutputWithContext(context.Background())
}

func (i FunctionMap) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionMapOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

// Amazon Resource Name (ARN) identifying your CloudFront Function.
func (o FunctionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Source code of the function
func (o FunctionOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Code }).(pulumi.StringOutput)
}

// Comment.
func (o FunctionOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
func (o FunctionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// ETag hash of any `LIVE` stage of the function.
func (o FunctionOutput) LiveStageEtag() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.LiveStageEtag }).(pulumi.StringOutput)
}

// Unique name for your CloudFront Function.
func (o FunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
func (o FunctionOutput) Publish() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.BoolPtrOutput { return v.Publish }).(pulumi.BoolPtrOutput)
}

// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
//
// The following arguments are optional:
func (o FunctionOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Runtime }).(pulumi.StringOutput)
}

// Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
func (o FunctionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type FunctionArrayOutput struct{ *pulumi.OutputState }

func (FunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (o FunctionArrayOutput) ToFunctionArrayOutput() FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) Index(i pulumi.IntInput) FunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Function {
		return vs[0].([]*Function)[vs[1].(int)]
	}).(FunctionOutput)
}

type FunctionMapOutput struct{ *pulumi.OutputState }

func (FunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (o FunctionMapOutput) ToFunctionMapOutput() FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) MapIndex(k pulumi.StringInput) FunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Function {
		return vs[0].(map[string]*Function)[vs[1].(string)]
	}).(FunctionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionInput)(nil)).Elem(), &Function{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionArrayInput)(nil)).Elem(), FunctionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionMapInput)(nil)).Elem(), FunctionMap{})
	pulumi.RegisterOutputType(FunctionOutput{})
	pulumi.RegisterOutputType(FunctionArrayOutput{})
	pulumi.RegisterOutputType(FunctionMapOutput{})
}
