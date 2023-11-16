// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an API Gateway Usage Plan Key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := apigateway.NewRestApi(ctx, "test", nil)
//			if err != nil {
//				return err
//			}
//			myusageplan, err := apigateway.NewUsagePlan(ctx, "myusageplan", &apigateway.UsagePlanArgs{
//				ApiStages: apigateway.UsagePlanApiStageArray{
//					&apigateway.UsagePlanApiStageArgs{
//						ApiId: test.ID(),
//						Stage: pulumi.Any(aws_api_gateway_stage.Foo.Stage_name),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			mykey, err := apigateway.NewApiKey(ctx, "mykey", nil)
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewUsagePlanKey(ctx, "main", &apigateway.UsagePlanKeyArgs{
//				KeyId:       mykey.ID(),
//				KeyType:     pulumi.String("API_KEY"),
//				UsagePlanId: myusageplan.ID(),
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
// Using `pulumi import`, import AWS API Gateway Usage Plan Key using the `USAGE-PLAN-ID/USAGE-PLAN-KEY-ID`. For example:
//
// ```sh
//
//	$ pulumi import aws:apigateway/usagePlanKey:UsagePlanKey key 12345abcde/zzz
//
// ```
type UsagePlanKey struct {
	pulumi.CustomResourceState

	// Identifier of the API key resource.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType pulumi.StringOutput `pulumi:"keyType"`
	// Name of a usage plan key.
	Name pulumi.StringOutput `pulumi:"name"`
	// Id of the usage plan resource representing to associate the key to.
	UsagePlanId pulumi.StringOutput `pulumi:"usagePlanId"`
	// Value of a usage plan key.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewUsagePlanKey registers a new resource with the given unique name, arguments, and options.
func NewUsagePlanKey(ctx *pulumi.Context,
	name string, args *UsagePlanKeyArgs, opts ...pulumi.ResourceOption) (*UsagePlanKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	if args.KeyType == nil {
		return nil, errors.New("invalid value for required argument 'KeyType'")
	}
	if args.UsagePlanId == nil {
		return nil, errors.New("invalid value for required argument 'UsagePlanId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UsagePlanKey
	err := ctx.RegisterResource("aws:apigateway/usagePlanKey:UsagePlanKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUsagePlanKey gets an existing UsagePlanKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUsagePlanKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UsagePlanKeyState, opts ...pulumi.ResourceOption) (*UsagePlanKey, error) {
	var resource UsagePlanKey
	err := ctx.ReadResource("aws:apigateway/usagePlanKey:UsagePlanKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UsagePlanKey resources.
type usagePlanKeyState struct {
	// Identifier of the API key resource.
	KeyId *string `pulumi:"keyId"`
	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType *string `pulumi:"keyType"`
	// Name of a usage plan key.
	Name *string `pulumi:"name"`
	// Id of the usage plan resource representing to associate the key to.
	UsagePlanId *string `pulumi:"usagePlanId"`
	// Value of a usage plan key.
	Value *string `pulumi:"value"`
}

type UsagePlanKeyState struct {
	// Identifier of the API key resource.
	KeyId pulumi.StringPtrInput
	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType pulumi.StringPtrInput
	// Name of a usage plan key.
	Name pulumi.StringPtrInput
	// Id of the usage plan resource representing to associate the key to.
	UsagePlanId pulumi.StringPtrInput
	// Value of a usage plan key.
	Value pulumi.StringPtrInput
}

func (UsagePlanKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanKeyState)(nil)).Elem()
}

type usagePlanKeyArgs struct {
	// Identifier of the API key resource.
	KeyId string `pulumi:"keyId"`
	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType string `pulumi:"keyType"`
	// Id of the usage plan resource representing to associate the key to.
	UsagePlanId string `pulumi:"usagePlanId"`
}

// The set of arguments for constructing a UsagePlanKey resource.
type UsagePlanKeyArgs struct {
	// Identifier of the API key resource.
	KeyId pulumi.StringInput
	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType pulumi.StringInput
	// Id of the usage plan resource representing to associate the key to.
	UsagePlanId pulumi.StringInput
}

func (UsagePlanKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanKeyArgs)(nil)).Elem()
}

type UsagePlanKeyInput interface {
	pulumi.Input

	ToUsagePlanKeyOutput() UsagePlanKeyOutput
	ToUsagePlanKeyOutputWithContext(ctx context.Context) UsagePlanKeyOutput
}

func (*UsagePlanKey) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlanKey)(nil)).Elem()
}

func (i *UsagePlanKey) ToUsagePlanKeyOutput() UsagePlanKeyOutput {
	return i.ToUsagePlanKeyOutputWithContext(context.Background())
}

func (i *UsagePlanKey) ToUsagePlanKeyOutputWithContext(ctx context.Context) UsagePlanKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanKeyOutput)
}

// UsagePlanKeyArrayInput is an input type that accepts UsagePlanKeyArray and UsagePlanKeyArrayOutput values.
// You can construct a concrete instance of `UsagePlanKeyArrayInput` via:
//
//	UsagePlanKeyArray{ UsagePlanKeyArgs{...} }
type UsagePlanKeyArrayInput interface {
	pulumi.Input

	ToUsagePlanKeyArrayOutput() UsagePlanKeyArrayOutput
	ToUsagePlanKeyArrayOutputWithContext(context.Context) UsagePlanKeyArrayOutput
}

type UsagePlanKeyArray []UsagePlanKeyInput

func (UsagePlanKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UsagePlanKey)(nil)).Elem()
}

func (i UsagePlanKeyArray) ToUsagePlanKeyArrayOutput() UsagePlanKeyArrayOutput {
	return i.ToUsagePlanKeyArrayOutputWithContext(context.Background())
}

func (i UsagePlanKeyArray) ToUsagePlanKeyArrayOutputWithContext(ctx context.Context) UsagePlanKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanKeyArrayOutput)
}

// UsagePlanKeyMapInput is an input type that accepts UsagePlanKeyMap and UsagePlanKeyMapOutput values.
// You can construct a concrete instance of `UsagePlanKeyMapInput` via:
//
//	UsagePlanKeyMap{ "key": UsagePlanKeyArgs{...} }
type UsagePlanKeyMapInput interface {
	pulumi.Input

	ToUsagePlanKeyMapOutput() UsagePlanKeyMapOutput
	ToUsagePlanKeyMapOutputWithContext(context.Context) UsagePlanKeyMapOutput
}

type UsagePlanKeyMap map[string]UsagePlanKeyInput

func (UsagePlanKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UsagePlanKey)(nil)).Elem()
}

func (i UsagePlanKeyMap) ToUsagePlanKeyMapOutput() UsagePlanKeyMapOutput {
	return i.ToUsagePlanKeyMapOutputWithContext(context.Background())
}

func (i UsagePlanKeyMap) ToUsagePlanKeyMapOutputWithContext(ctx context.Context) UsagePlanKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanKeyMapOutput)
}

type UsagePlanKeyOutput struct{ *pulumi.OutputState }

func (UsagePlanKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlanKey)(nil)).Elem()
}

func (o UsagePlanKeyOutput) ToUsagePlanKeyOutput() UsagePlanKeyOutput {
	return o
}

func (o UsagePlanKeyOutput) ToUsagePlanKeyOutputWithContext(ctx context.Context) UsagePlanKeyOutput {
	return o
}

// Identifier of the API key resource.
func (o UsagePlanKeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanKey) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// Type of the API key resource. Currently, the valid key type is API_KEY.
func (o UsagePlanKeyOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanKey) pulumi.StringOutput { return v.KeyType }).(pulumi.StringOutput)
}

// Name of a usage plan key.
func (o UsagePlanKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Id of the usage plan resource representing to associate the key to.
func (o UsagePlanKeyOutput) UsagePlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanKey) pulumi.StringOutput { return v.UsagePlanId }).(pulumi.StringOutput)
}

// Value of a usage plan key.
func (o UsagePlanKeyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanKey) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type UsagePlanKeyArrayOutput struct{ *pulumi.OutputState }

func (UsagePlanKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UsagePlanKey)(nil)).Elem()
}

func (o UsagePlanKeyArrayOutput) ToUsagePlanKeyArrayOutput() UsagePlanKeyArrayOutput {
	return o
}

func (o UsagePlanKeyArrayOutput) ToUsagePlanKeyArrayOutputWithContext(ctx context.Context) UsagePlanKeyArrayOutput {
	return o
}

func (o UsagePlanKeyArrayOutput) Index(i pulumi.IntInput) UsagePlanKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UsagePlanKey {
		return vs[0].([]*UsagePlanKey)[vs[1].(int)]
	}).(UsagePlanKeyOutput)
}

type UsagePlanKeyMapOutput struct{ *pulumi.OutputState }

func (UsagePlanKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UsagePlanKey)(nil)).Elem()
}

func (o UsagePlanKeyMapOutput) ToUsagePlanKeyMapOutput() UsagePlanKeyMapOutput {
	return o
}

func (o UsagePlanKeyMapOutput) ToUsagePlanKeyMapOutputWithContext(ctx context.Context) UsagePlanKeyMapOutput {
	return o
}

func (o UsagePlanKeyMapOutput) MapIndex(k pulumi.StringInput) UsagePlanKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UsagePlanKey {
		return vs[0].(map[string]*UsagePlanKey)[vs[1].(string)]
	}).(UsagePlanKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UsagePlanKeyInput)(nil)).Elem(), &UsagePlanKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsagePlanKeyArrayInput)(nil)).Elem(), UsagePlanKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsagePlanKeyMapInput)(nil)).Elem(), UsagePlanKeyMap{})
	pulumi.RegisterOutputType(UsagePlanKeyOutput{})
	pulumi.RegisterOutputType(UsagePlanKeyArrayOutput{})
	pulumi.RegisterOutputType(UsagePlanKeyMapOutput{})
}
