// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS Config Aggregate Authorization
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cfg.NewAggregateAuthorization(ctx, "example", &cfg.AggregateAuthorizationArgs{
//				AccountId: pulumi.String("123456789012"),
//				Region:    pulumi.String("eu-west-2"),
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
// Using `pulumi import`, import Config aggregate authorizations using `account_id:region`. For example:
//
// ```sh
//
//	$ pulumi import aws:cfg/aggregateAuthorization:AggregateAuthorization example 123456789012:us-east-1
//
// ```
type AggregateAuthorization struct {
	pulumi.CustomResourceState

	// Account ID
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The ARN of the authorization
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Region
	Region pulumi.StringOutput `pulumi:"region"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewAggregateAuthorization registers a new resource with the given unique name, arguments, and options.
func NewAggregateAuthorization(ctx *pulumi.Context,
	name string, args *AggregateAuthorizationArgs, opts ...pulumi.ResourceOption) (*AggregateAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AggregateAuthorization
	err := ctx.RegisterResource("aws:cfg/aggregateAuthorization:AggregateAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAggregateAuthorization gets an existing AggregateAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAggregateAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AggregateAuthorizationState, opts ...pulumi.ResourceOption) (*AggregateAuthorization, error) {
	var resource AggregateAuthorization
	err := ctx.ReadResource("aws:cfg/aggregateAuthorization:AggregateAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AggregateAuthorization resources.
type aggregateAuthorizationState struct {
	// Account ID
	AccountId *string `pulumi:"accountId"`
	// The ARN of the authorization
	Arn *string `pulumi:"arn"`
	// Region
	Region *string `pulumi:"region"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type AggregateAuthorizationState struct {
	// Account ID
	AccountId pulumi.StringPtrInput
	// The ARN of the authorization
	Arn pulumi.StringPtrInput
	// Region
	Region pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (AggregateAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateAuthorizationState)(nil)).Elem()
}

type aggregateAuthorizationArgs struct {
	// Account ID
	AccountId string `pulumi:"accountId"`
	// Region
	Region string `pulumi:"region"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AggregateAuthorization resource.
type AggregateAuthorizationArgs struct {
	// Account ID
	AccountId pulumi.StringInput
	// Region
	Region pulumi.StringInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (AggregateAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateAuthorizationArgs)(nil)).Elem()
}

type AggregateAuthorizationInput interface {
	pulumi.Input

	ToAggregateAuthorizationOutput() AggregateAuthorizationOutput
	ToAggregateAuthorizationOutputWithContext(ctx context.Context) AggregateAuthorizationOutput
}

func (*AggregateAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateAuthorization)(nil)).Elem()
}

func (i *AggregateAuthorization) ToAggregateAuthorizationOutput() AggregateAuthorizationOutput {
	return i.ToAggregateAuthorizationOutputWithContext(context.Background())
}

func (i *AggregateAuthorization) ToAggregateAuthorizationOutputWithContext(ctx context.Context) AggregateAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateAuthorizationOutput)
}

// AggregateAuthorizationArrayInput is an input type that accepts AggregateAuthorizationArray and AggregateAuthorizationArrayOutput values.
// You can construct a concrete instance of `AggregateAuthorizationArrayInput` via:
//
//	AggregateAuthorizationArray{ AggregateAuthorizationArgs{...} }
type AggregateAuthorizationArrayInput interface {
	pulumi.Input

	ToAggregateAuthorizationArrayOutput() AggregateAuthorizationArrayOutput
	ToAggregateAuthorizationArrayOutputWithContext(context.Context) AggregateAuthorizationArrayOutput
}

type AggregateAuthorizationArray []AggregateAuthorizationInput

func (AggregateAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateAuthorization)(nil)).Elem()
}

func (i AggregateAuthorizationArray) ToAggregateAuthorizationArrayOutput() AggregateAuthorizationArrayOutput {
	return i.ToAggregateAuthorizationArrayOutputWithContext(context.Background())
}

func (i AggregateAuthorizationArray) ToAggregateAuthorizationArrayOutputWithContext(ctx context.Context) AggregateAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateAuthorizationArrayOutput)
}

// AggregateAuthorizationMapInput is an input type that accepts AggregateAuthorizationMap and AggregateAuthorizationMapOutput values.
// You can construct a concrete instance of `AggregateAuthorizationMapInput` via:
//
//	AggregateAuthorizationMap{ "key": AggregateAuthorizationArgs{...} }
type AggregateAuthorizationMapInput interface {
	pulumi.Input

	ToAggregateAuthorizationMapOutput() AggregateAuthorizationMapOutput
	ToAggregateAuthorizationMapOutputWithContext(context.Context) AggregateAuthorizationMapOutput
}

type AggregateAuthorizationMap map[string]AggregateAuthorizationInput

func (AggregateAuthorizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateAuthorization)(nil)).Elem()
}

func (i AggregateAuthorizationMap) ToAggregateAuthorizationMapOutput() AggregateAuthorizationMapOutput {
	return i.ToAggregateAuthorizationMapOutputWithContext(context.Background())
}

func (i AggregateAuthorizationMap) ToAggregateAuthorizationMapOutputWithContext(ctx context.Context) AggregateAuthorizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateAuthorizationMapOutput)
}

type AggregateAuthorizationOutput struct{ *pulumi.OutputState }

func (AggregateAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateAuthorization)(nil)).Elem()
}

func (o AggregateAuthorizationOutput) ToAggregateAuthorizationOutput() AggregateAuthorizationOutput {
	return o
}

func (o AggregateAuthorizationOutput) ToAggregateAuthorizationOutputWithContext(ctx context.Context) AggregateAuthorizationOutput {
	return o
}

// Account ID
func (o AggregateAuthorizationOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateAuthorization) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The ARN of the authorization
func (o AggregateAuthorizationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateAuthorization) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Region
func (o AggregateAuthorizationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateAuthorization) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o AggregateAuthorizationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AggregateAuthorization) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o AggregateAuthorizationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AggregateAuthorization) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type AggregateAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (AggregateAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateAuthorization)(nil)).Elem()
}

func (o AggregateAuthorizationArrayOutput) ToAggregateAuthorizationArrayOutput() AggregateAuthorizationArrayOutput {
	return o
}

func (o AggregateAuthorizationArrayOutput) ToAggregateAuthorizationArrayOutputWithContext(ctx context.Context) AggregateAuthorizationArrayOutput {
	return o
}

func (o AggregateAuthorizationArrayOutput) Index(i pulumi.IntInput) AggregateAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AggregateAuthorization {
		return vs[0].([]*AggregateAuthorization)[vs[1].(int)]
	}).(AggregateAuthorizationOutput)
}

type AggregateAuthorizationMapOutput struct{ *pulumi.OutputState }

func (AggregateAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateAuthorization)(nil)).Elem()
}

func (o AggregateAuthorizationMapOutput) ToAggregateAuthorizationMapOutput() AggregateAuthorizationMapOutput {
	return o
}

func (o AggregateAuthorizationMapOutput) ToAggregateAuthorizationMapOutputWithContext(ctx context.Context) AggregateAuthorizationMapOutput {
	return o
}

func (o AggregateAuthorizationMapOutput) MapIndex(k pulumi.StringInput) AggregateAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AggregateAuthorization {
		return vs[0].(map[string]*AggregateAuthorization)[vs[1].(string)]
	}).(AggregateAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateAuthorizationInput)(nil)).Elem(), &AggregateAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateAuthorizationArrayInput)(nil)).Elem(), AggregateAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateAuthorizationMapInput)(nil)).Elem(), AggregateAuthorizationMap{})
	pulumi.RegisterOutputType(AggregateAuthorizationOutput{})
	pulumi.RegisterOutputType(AggregateAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(AggregateAuthorizationMapOutput{})
}
