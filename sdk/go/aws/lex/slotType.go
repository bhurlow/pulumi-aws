// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon Lex Slot Type resource. For more information see
// [Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := lex.NewSlotType(ctx, "flowerTypes", &lex.SlotTypeArgs{
//				CreateVersion: pulumi.Bool(true),
//				Description:   pulumi.String("Types of flowers to order"),
//				EnumerationValues: lex.SlotTypeEnumerationValueArray{
//					&lex.SlotTypeEnumerationValueArgs{
//						Synonyms: pulumi.StringArray{
//							pulumi.String("Lirium"),
//							pulumi.String("Martagon"),
//						},
//						Value: pulumi.String("lilies"),
//					},
//					&lex.SlotTypeEnumerationValueArgs{
//						Synonyms: pulumi.StringArray{
//							pulumi.String("Eduardoregelia"),
//							pulumi.String("Podonix"),
//						},
//						Value: pulumi.String("tulips"),
//					},
//				},
//				Name:                   pulumi.String("FlowerTypes"),
//				ValueSelectionStrategy: pulumi.String("ORIGINAL_VALUE"),
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
//	to = aws_lex_slot_type.flower_types
//
//	id = "FlowerTypes" } Using `pulumi import`, import slot types using their name. For exampleconsole % pulumi import aws_lex_slot_type.flower_types FlowerTypes
type SlotType struct {
	pulumi.CustomResourceState

	// Checksum identifying the version of the slot type that was created. The checksum is
	// not included as an argument because the resource will add it automatically when updating the slot type.
	Checksum pulumi.StringOutput `pulumi:"checksum"`
	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrOutput `pulumi:"createVersion"`
	// The date when the slot type version was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	EnumerationValues SlotTypeEnumerationValueArrayOutput `pulumi:"enumerationValues"`
	// The date when the `$LATEST` version of this slot type was updated.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
	// value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
	ValueSelectionStrategy pulumi.StringPtrOutput `pulumi:"valueSelectionStrategy"`
	// The version of the slot type.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSlotType registers a new resource with the given unique name, arguments, and options.
func NewSlotType(ctx *pulumi.Context,
	name string, args *SlotTypeArgs, opts ...pulumi.ResourceOption) (*SlotType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnumerationValues == nil {
		return nil, errors.New("invalid value for required argument 'EnumerationValues'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SlotType
	err := ctx.RegisterResource("aws:lex/slotType:SlotType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlotType gets an existing SlotType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlotType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SlotTypeState, opts ...pulumi.ResourceOption) (*SlotType, error) {
	var resource SlotType
	err := ctx.ReadResource("aws:lex/slotType:SlotType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SlotType resources.
type slotTypeState struct {
	// Checksum identifying the version of the slot type that was created. The checksum is
	// not included as an argument because the resource will add it automatically when updating the slot type.
	Checksum *string `pulumi:"checksum"`
	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to `false`.
	CreateVersion *bool `pulumi:"createVersion"`
	// The date when the slot type version was created.
	CreatedDate *string `pulumi:"createdDate"`
	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description *string `pulumi:"description"`
	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	EnumerationValues []SlotTypeEnumerationValue `pulumi:"enumerationValues"`
	// The date when the `$LATEST` version of this slot type was updated.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
	Name *string `pulumi:"name"`
	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
	// value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
	ValueSelectionStrategy *string `pulumi:"valueSelectionStrategy"`
	// The version of the slot type.
	Version *string `pulumi:"version"`
}

type SlotTypeState struct {
	// Checksum identifying the version of the slot type that was created. The checksum is
	// not included as an argument because the resource will add it automatically when updating the slot type.
	Checksum pulumi.StringPtrInput
	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrInput
	// The date when the slot type version was created.
	CreatedDate pulumi.StringPtrInput
	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description pulumi.StringPtrInput
	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	EnumerationValues SlotTypeEnumerationValueArrayInput
	// The date when the `$LATEST` version of this slot type was updated.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
	Name pulumi.StringPtrInput
	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
	// value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
	ValueSelectionStrategy pulumi.StringPtrInput
	// The version of the slot type.
	Version pulumi.StringPtrInput
}

func (SlotTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*slotTypeState)(nil)).Elem()
}

type slotTypeArgs struct {
	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to `false`.
	CreateVersion *bool `pulumi:"createVersion"`
	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description *string `pulumi:"description"`
	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	EnumerationValues []SlotTypeEnumerationValue `pulumi:"enumerationValues"`
	// The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
	Name *string `pulumi:"name"`
	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
	// value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
	ValueSelectionStrategy *string `pulumi:"valueSelectionStrategy"`
}

// The set of arguments for constructing a SlotType resource.
type SlotTypeArgs struct {
	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrInput
	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description pulumi.StringPtrInput
	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	EnumerationValues SlotTypeEnumerationValueArrayInput
	// The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
	Name pulumi.StringPtrInput
	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
	// value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
	ValueSelectionStrategy pulumi.StringPtrInput
}

func (SlotTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*slotTypeArgs)(nil)).Elem()
}

type SlotTypeInput interface {
	pulumi.Input

	ToSlotTypeOutput() SlotTypeOutput
	ToSlotTypeOutputWithContext(ctx context.Context) SlotTypeOutput
}

func (*SlotType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlotType)(nil)).Elem()
}

func (i *SlotType) ToSlotTypeOutput() SlotTypeOutput {
	return i.ToSlotTypeOutputWithContext(context.Background())
}

func (i *SlotType) ToSlotTypeOutputWithContext(ctx context.Context) SlotTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotTypeOutput)
}

// SlotTypeArrayInput is an input type that accepts SlotTypeArray and SlotTypeArrayOutput values.
// You can construct a concrete instance of `SlotTypeArrayInput` via:
//
//	SlotTypeArray{ SlotTypeArgs{...} }
type SlotTypeArrayInput interface {
	pulumi.Input

	ToSlotTypeArrayOutput() SlotTypeArrayOutput
	ToSlotTypeArrayOutputWithContext(context.Context) SlotTypeArrayOutput
}

type SlotTypeArray []SlotTypeInput

func (SlotTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SlotType)(nil)).Elem()
}

func (i SlotTypeArray) ToSlotTypeArrayOutput() SlotTypeArrayOutput {
	return i.ToSlotTypeArrayOutputWithContext(context.Background())
}

func (i SlotTypeArray) ToSlotTypeArrayOutputWithContext(ctx context.Context) SlotTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotTypeArrayOutput)
}

// SlotTypeMapInput is an input type that accepts SlotTypeMap and SlotTypeMapOutput values.
// You can construct a concrete instance of `SlotTypeMapInput` via:
//
//	SlotTypeMap{ "key": SlotTypeArgs{...} }
type SlotTypeMapInput interface {
	pulumi.Input

	ToSlotTypeMapOutput() SlotTypeMapOutput
	ToSlotTypeMapOutputWithContext(context.Context) SlotTypeMapOutput
}

type SlotTypeMap map[string]SlotTypeInput

func (SlotTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SlotType)(nil)).Elem()
}

func (i SlotTypeMap) ToSlotTypeMapOutput() SlotTypeMapOutput {
	return i.ToSlotTypeMapOutputWithContext(context.Background())
}

func (i SlotTypeMap) ToSlotTypeMapOutputWithContext(ctx context.Context) SlotTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotTypeMapOutput)
}

type SlotTypeOutput struct{ *pulumi.OutputState }

func (SlotTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlotType)(nil)).Elem()
}

func (o SlotTypeOutput) ToSlotTypeOutput() SlotTypeOutput {
	return o
}

func (o SlotTypeOutput) ToSlotTypeOutputWithContext(ctx context.Context) SlotTypeOutput {
	return o
}

// Checksum identifying the version of the slot type that was created. The checksum is
// not included as an argument because the resource will add it automatically when updating the slot type.
func (o SlotTypeOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v *SlotType) pulumi.StringOutput { return v.Checksum }).(pulumi.StringOutput)
}

// Determines if a new slot type version is created when the initial resource is created and on each
// update. Defaults to `false`.
func (o SlotTypeOutput) CreateVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SlotType) pulumi.BoolPtrOutput { return v.CreateVersion }).(pulumi.BoolPtrOutput)
}

// The date when the slot type version was created.
func (o SlotTypeOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SlotType) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// A description of the slot type. Must be less than or equal to 200 characters in length.
func (o SlotTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlotType) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of EnumerationValue objects that defines the values that
// the slot type can take. Each value can have a list of synonyms, which are additional values that help
// train the machine learning model about the values that it resolves for a slot. Attributes are
// documented under enumeration_value.
func (o SlotTypeOutput) EnumerationValues() SlotTypeEnumerationValueArrayOutput {
	return o.ApplyT(func(v *SlotType) SlotTypeEnumerationValueArrayOutput { return v.EnumerationValues }).(SlotTypeEnumerationValueArrayOutput)
}

// The date when the `$LATEST` version of this slot type was updated.
func (o SlotTypeOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SlotType) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

// The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
func (o SlotTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SlotType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Determines the slot resolution strategy that Amazon Lex
// uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
// value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
// if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
func (o SlotTypeOutput) ValueSelectionStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlotType) pulumi.StringPtrOutput { return v.ValueSelectionStrategy }).(pulumi.StringPtrOutput)
}

// The version of the slot type.
func (o SlotTypeOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SlotType) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type SlotTypeArrayOutput struct{ *pulumi.OutputState }

func (SlotTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SlotType)(nil)).Elem()
}

func (o SlotTypeArrayOutput) ToSlotTypeArrayOutput() SlotTypeArrayOutput {
	return o
}

func (o SlotTypeArrayOutput) ToSlotTypeArrayOutputWithContext(ctx context.Context) SlotTypeArrayOutput {
	return o
}

func (o SlotTypeArrayOutput) Index(i pulumi.IntInput) SlotTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SlotType {
		return vs[0].([]*SlotType)[vs[1].(int)]
	}).(SlotTypeOutput)
}

type SlotTypeMapOutput struct{ *pulumi.OutputState }

func (SlotTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SlotType)(nil)).Elem()
}

func (o SlotTypeMapOutput) ToSlotTypeMapOutput() SlotTypeMapOutput {
	return o
}

func (o SlotTypeMapOutput) ToSlotTypeMapOutputWithContext(ctx context.Context) SlotTypeMapOutput {
	return o
}

func (o SlotTypeMapOutput) MapIndex(k pulumi.StringInput) SlotTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SlotType {
		return vs[0].(map[string]*SlotType)[vs[1].(string)]
	}).(SlotTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SlotTypeInput)(nil)).Elem(), &SlotType{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlotTypeArrayInput)(nil)).Elem(), SlotTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlotTypeMapInput)(nil)).Elem(), SlotTypeMap{})
	pulumi.RegisterOutputType(SlotTypeOutput{})
	pulumi.RegisterOutputType(SlotTypeArrayOutput{})
	pulumi.RegisterOutputType(SlotTypeMapOutput{})
}
