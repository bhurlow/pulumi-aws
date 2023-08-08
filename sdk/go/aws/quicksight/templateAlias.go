// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS QuickSight Template Alias.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := quicksight.NewTemplateAlias(ctx, "example", &quicksight.TemplateAliasArgs{
//				AliasName:             pulumi.String("example-alias"),
//				TemplateId:            pulumi.Any(aws_quicksight_template.Test.Template_id),
//				TemplateVersionNumber: pulumi.Any(aws_quicksight_template.Test.Version_number),
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
//	to = aws_quicksight_template_alias.example
//
//	id = "123456789012,example-id,example-alias" } Using `pulumi import`, import QuickSight Template Alias using the AWS account ID, template ID, and alias name separated by a comma (`,`). For exampleconsole % pulumi import aws_quicksight_template_alias.example 123456789012,example-id,example-alias
type TemplateAlias struct {
	pulumi.CustomResourceState

	// Display name of the template alias.
	AliasName pulumi.StringOutput `pulumi:"aliasName"`
	// Amazon Resource Name (ARN) of the template alias.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// ID of the template.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// Version number of the template.
	//
	// The following arguments are optional:
	TemplateVersionNumber pulumi.IntOutput `pulumi:"templateVersionNumber"`
}

// NewTemplateAlias registers a new resource with the given unique name, arguments, and options.
func NewTemplateAlias(ctx *pulumi.Context,
	name string, args *TemplateAliasArgs, opts ...pulumi.ResourceOption) (*TemplateAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AliasName == nil {
		return nil, errors.New("invalid value for required argument 'AliasName'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	if args.TemplateVersionNumber == nil {
		return nil, errors.New("invalid value for required argument 'TemplateVersionNumber'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TemplateAlias
	err := ctx.RegisterResource("aws:quicksight/templateAlias:TemplateAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateAlias gets an existing TemplateAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateAliasState, opts ...pulumi.ResourceOption) (*TemplateAlias, error) {
	var resource TemplateAlias
	err := ctx.ReadResource("aws:quicksight/templateAlias:TemplateAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateAlias resources.
type templateAliasState struct {
	// Display name of the template alias.
	AliasName *string `pulumi:"aliasName"`
	// Amazon Resource Name (ARN) of the template alias.
	Arn *string `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// ID of the template.
	TemplateId *string `pulumi:"templateId"`
	// Version number of the template.
	//
	// The following arguments are optional:
	TemplateVersionNumber *int `pulumi:"templateVersionNumber"`
}

type TemplateAliasState struct {
	// Display name of the template alias.
	AliasName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the template alias.
	Arn pulumi.StringPtrInput
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// ID of the template.
	TemplateId pulumi.StringPtrInput
	// Version number of the template.
	//
	// The following arguments are optional:
	TemplateVersionNumber pulumi.IntPtrInput
}

func (TemplateAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateAliasState)(nil)).Elem()
}

type templateAliasArgs struct {
	// Display name of the template alias.
	AliasName string `pulumi:"aliasName"`
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// ID of the template.
	TemplateId string `pulumi:"templateId"`
	// Version number of the template.
	//
	// The following arguments are optional:
	TemplateVersionNumber int `pulumi:"templateVersionNumber"`
}

// The set of arguments for constructing a TemplateAlias resource.
type TemplateAliasArgs struct {
	// Display name of the template alias.
	AliasName pulumi.StringInput
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// ID of the template.
	TemplateId pulumi.StringInput
	// Version number of the template.
	//
	// The following arguments are optional:
	TemplateVersionNumber pulumi.IntInput
}

func (TemplateAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateAliasArgs)(nil)).Elem()
}

type TemplateAliasInput interface {
	pulumi.Input

	ToTemplateAliasOutput() TemplateAliasOutput
	ToTemplateAliasOutputWithContext(ctx context.Context) TemplateAliasOutput
}

func (*TemplateAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateAlias)(nil)).Elem()
}

func (i *TemplateAlias) ToTemplateAliasOutput() TemplateAliasOutput {
	return i.ToTemplateAliasOutputWithContext(context.Background())
}

func (i *TemplateAlias) ToTemplateAliasOutputWithContext(ctx context.Context) TemplateAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateAliasOutput)
}

// TemplateAliasArrayInput is an input type that accepts TemplateAliasArray and TemplateAliasArrayOutput values.
// You can construct a concrete instance of `TemplateAliasArrayInput` via:
//
//	TemplateAliasArray{ TemplateAliasArgs{...} }
type TemplateAliasArrayInput interface {
	pulumi.Input

	ToTemplateAliasArrayOutput() TemplateAliasArrayOutput
	ToTemplateAliasArrayOutputWithContext(context.Context) TemplateAliasArrayOutput
}

type TemplateAliasArray []TemplateAliasInput

func (TemplateAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateAlias)(nil)).Elem()
}

func (i TemplateAliasArray) ToTemplateAliasArrayOutput() TemplateAliasArrayOutput {
	return i.ToTemplateAliasArrayOutputWithContext(context.Background())
}

func (i TemplateAliasArray) ToTemplateAliasArrayOutputWithContext(ctx context.Context) TemplateAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateAliasArrayOutput)
}

// TemplateAliasMapInput is an input type that accepts TemplateAliasMap and TemplateAliasMapOutput values.
// You can construct a concrete instance of `TemplateAliasMapInput` via:
//
//	TemplateAliasMap{ "key": TemplateAliasArgs{...} }
type TemplateAliasMapInput interface {
	pulumi.Input

	ToTemplateAliasMapOutput() TemplateAliasMapOutput
	ToTemplateAliasMapOutputWithContext(context.Context) TemplateAliasMapOutput
}

type TemplateAliasMap map[string]TemplateAliasInput

func (TemplateAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateAlias)(nil)).Elem()
}

func (i TemplateAliasMap) ToTemplateAliasMapOutput() TemplateAliasMapOutput {
	return i.ToTemplateAliasMapOutputWithContext(context.Background())
}

func (i TemplateAliasMap) ToTemplateAliasMapOutputWithContext(ctx context.Context) TemplateAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateAliasMapOutput)
}

type TemplateAliasOutput struct{ *pulumi.OutputState }

func (TemplateAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateAlias)(nil)).Elem()
}

func (o TemplateAliasOutput) ToTemplateAliasOutput() TemplateAliasOutput {
	return o
}

func (o TemplateAliasOutput) ToTemplateAliasOutputWithContext(ctx context.Context) TemplateAliasOutput {
	return o
}

// Display name of the template alias.
func (o TemplateAliasOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateAlias) pulumi.StringOutput { return v.AliasName }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the template alias.
func (o TemplateAliasOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateAlias) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// AWS account ID.
func (o TemplateAliasOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateAlias) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

// ID of the template.
func (o TemplateAliasOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateAlias) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// Version number of the template.
//
// The following arguments are optional:
func (o TemplateAliasOutput) TemplateVersionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *TemplateAlias) pulumi.IntOutput { return v.TemplateVersionNumber }).(pulumi.IntOutput)
}

type TemplateAliasArrayOutput struct{ *pulumi.OutputState }

func (TemplateAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateAlias)(nil)).Elem()
}

func (o TemplateAliasArrayOutput) ToTemplateAliasArrayOutput() TemplateAliasArrayOutput {
	return o
}

func (o TemplateAliasArrayOutput) ToTemplateAliasArrayOutputWithContext(ctx context.Context) TemplateAliasArrayOutput {
	return o
}

func (o TemplateAliasArrayOutput) Index(i pulumi.IntInput) TemplateAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TemplateAlias {
		return vs[0].([]*TemplateAlias)[vs[1].(int)]
	}).(TemplateAliasOutput)
}

type TemplateAliasMapOutput struct{ *pulumi.OutputState }

func (TemplateAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateAlias)(nil)).Elem()
}

func (o TemplateAliasMapOutput) ToTemplateAliasMapOutput() TemplateAliasMapOutput {
	return o
}

func (o TemplateAliasMapOutput) ToTemplateAliasMapOutputWithContext(ctx context.Context) TemplateAliasMapOutput {
	return o
}

func (o TemplateAliasMapOutput) MapIndex(k pulumi.StringInput) TemplateAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TemplateAlias {
		return vs[0].(map[string]*TemplateAlias)[vs[1].(string)]
	}).(TemplateAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateAliasInput)(nil)).Elem(), &TemplateAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateAliasArrayInput)(nil)).Elem(), TemplateAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateAliasMapInput)(nil)).Elem(), TemplateAliasMap{})
	pulumi.RegisterOutputType(TemplateAliasOutput{})
	pulumi.RegisterOutputType(TemplateAliasArrayOutput{})
	pulumi.RegisterOutputType(TemplateAliasMapOutput{})
}
