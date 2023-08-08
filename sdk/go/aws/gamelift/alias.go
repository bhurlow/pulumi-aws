// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a GameLift Alias resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/gamelift"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gamelift.NewAlias(ctx, "example", &gamelift.AliasArgs{
//				Description: pulumi.String("Example Description"),
//				RoutingStrategy: &gamelift.AliasRoutingStrategyArgs{
//					Message: pulumi.String("Example Message"),
//					Type:    pulumi.String("TERMINAL"),
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
// terraform import {
//
//	to = aws_gamelift_alias.example
//
//	id = "<alias-id>" } Using `pulumi import`, import GameLift Aliases using the ID. For exampleconsole % pulumi import aws_gamelift_alias.example <alias-id>
type Alias struct {
	pulumi.CustomResourceState

	// Alias ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the alias.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the alias.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyOutput `pulumi:"routingStrategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoutingStrategy == nil {
		return nil, errors.New("invalid value for required argument 'RoutingStrategy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alias
	err := ctx.RegisterResource("aws:gamelift/alias:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws:gamelift/alias:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
	// Alias ARN.
	Arn *string `pulumi:"arn"`
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy *AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type AliasState struct {
	// Alias ARN.
	Arn pulumi.StringPtrInput
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

// AliasArrayInput is an input type that accepts AliasArray and AliasArrayOutput values.
// You can construct a concrete instance of `AliasArrayInput` via:
//
//	AliasArray{ AliasArgs{...} }
type AliasArrayInput interface {
	pulumi.Input

	ToAliasArrayOutput() AliasArrayOutput
	ToAliasArrayOutputWithContext(context.Context) AliasArrayOutput
}

type AliasArray []AliasInput

func (AliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alias)(nil)).Elem()
}

func (i AliasArray) ToAliasArrayOutput() AliasArrayOutput {
	return i.ToAliasArrayOutputWithContext(context.Background())
}

func (i AliasArray) ToAliasArrayOutputWithContext(ctx context.Context) AliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasArrayOutput)
}

// AliasMapInput is an input type that accepts AliasMap and AliasMapOutput values.
// You can construct a concrete instance of `AliasMapInput` via:
//
//	AliasMap{ "key": AliasArgs{...} }
type AliasMapInput interface {
	pulumi.Input

	ToAliasMapOutput() AliasMapOutput
	ToAliasMapOutputWithContext(context.Context) AliasMapOutput
}

type AliasMap map[string]AliasInput

func (AliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alias)(nil)).Elem()
}

func (i AliasMap) ToAliasMapOutput() AliasMapOutput {
	return i.ToAliasMapOutputWithContext(context.Background())
}

func (i AliasMap) ToAliasMapOutputWithContext(ctx context.Context) AliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasMapOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

// Alias ARN.
func (o AliasOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the alias.
func (o AliasOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the alias.
func (o AliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the fleet and/or routing type to use for the alias.
func (o AliasOutput) RoutingStrategy() AliasRoutingStrategyOutput {
	return o.ApplyT(func(v *Alias) AliasRoutingStrategyOutput { return v.RoutingStrategy }).(AliasRoutingStrategyOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o AliasOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o AliasOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type AliasArrayOutput struct{ *pulumi.OutputState }

func (AliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alias)(nil)).Elem()
}

func (o AliasArrayOutput) ToAliasArrayOutput() AliasArrayOutput {
	return o
}

func (o AliasArrayOutput) ToAliasArrayOutputWithContext(ctx context.Context) AliasArrayOutput {
	return o
}

func (o AliasArrayOutput) Index(i pulumi.IntInput) AliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alias {
		return vs[0].([]*Alias)[vs[1].(int)]
	}).(AliasOutput)
}

type AliasMapOutput struct{ *pulumi.OutputState }

func (AliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alias)(nil)).Elem()
}

func (o AliasMapOutput) ToAliasMapOutput() AliasMapOutput {
	return o
}

func (o AliasMapOutput) ToAliasMapOutputWithContext(ctx context.Context) AliasMapOutput {
	return o
}

func (o AliasMapOutput) MapIndex(k pulumi.StringInput) AliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alias {
		return vs[0].(map[string]*Alias)[vs[1].(string)]
	}).(AliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasArrayInput)(nil)).Elem(), AliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasMapInput)(nil)).Elem(), AliasMap{})
	pulumi.RegisterOutputType(AliasOutput{})
	pulumi.RegisterOutputType(AliasArrayOutput{})
	pulumi.RegisterOutputType(AliasMapOutput{})
}
