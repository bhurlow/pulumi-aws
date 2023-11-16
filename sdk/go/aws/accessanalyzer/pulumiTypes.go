// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accessanalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ArchiveRuleFilter struct {
	// Contains comparator.
	Contains []string `pulumi:"contains"`
	// Filter criteria.
	Criteria string `pulumi:"criteria"`
	// Equals comparator.
	Eqs []string `pulumi:"eqs"`
	// Boolean comparator.
	Exists *string `pulumi:"exists"`
	// Not Equals comparator.
	Neqs []string `pulumi:"neqs"`
}

// ArchiveRuleFilterInput is an input type that accepts ArchiveRuleFilterArgs and ArchiveRuleFilterOutput values.
// You can construct a concrete instance of `ArchiveRuleFilterInput` via:
//
//	ArchiveRuleFilterArgs{...}
type ArchiveRuleFilterInput interface {
	pulumi.Input

	ToArchiveRuleFilterOutput() ArchiveRuleFilterOutput
	ToArchiveRuleFilterOutputWithContext(context.Context) ArchiveRuleFilterOutput
}

type ArchiveRuleFilterArgs struct {
	// Contains comparator.
	Contains pulumi.StringArrayInput `pulumi:"contains"`
	// Filter criteria.
	Criteria pulumi.StringInput `pulumi:"criteria"`
	// Equals comparator.
	Eqs pulumi.StringArrayInput `pulumi:"eqs"`
	// Boolean comparator.
	Exists pulumi.StringPtrInput `pulumi:"exists"`
	// Not Equals comparator.
	Neqs pulumi.StringArrayInput `pulumi:"neqs"`
}

func (ArchiveRuleFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArchiveRuleFilter)(nil)).Elem()
}

func (i ArchiveRuleFilterArgs) ToArchiveRuleFilterOutput() ArchiveRuleFilterOutput {
	return i.ToArchiveRuleFilterOutputWithContext(context.Background())
}

func (i ArchiveRuleFilterArgs) ToArchiveRuleFilterOutputWithContext(ctx context.Context) ArchiveRuleFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArchiveRuleFilterOutput)
}

// ArchiveRuleFilterArrayInput is an input type that accepts ArchiveRuleFilterArray and ArchiveRuleFilterArrayOutput values.
// You can construct a concrete instance of `ArchiveRuleFilterArrayInput` via:
//
//	ArchiveRuleFilterArray{ ArchiveRuleFilterArgs{...} }
type ArchiveRuleFilterArrayInput interface {
	pulumi.Input

	ToArchiveRuleFilterArrayOutput() ArchiveRuleFilterArrayOutput
	ToArchiveRuleFilterArrayOutputWithContext(context.Context) ArchiveRuleFilterArrayOutput
}

type ArchiveRuleFilterArray []ArchiveRuleFilterInput

func (ArchiveRuleFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArchiveRuleFilter)(nil)).Elem()
}

func (i ArchiveRuleFilterArray) ToArchiveRuleFilterArrayOutput() ArchiveRuleFilterArrayOutput {
	return i.ToArchiveRuleFilterArrayOutputWithContext(context.Background())
}

func (i ArchiveRuleFilterArray) ToArchiveRuleFilterArrayOutputWithContext(ctx context.Context) ArchiveRuleFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArchiveRuleFilterArrayOutput)
}

type ArchiveRuleFilterOutput struct{ *pulumi.OutputState }

func (ArchiveRuleFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArchiveRuleFilter)(nil)).Elem()
}

func (o ArchiveRuleFilterOutput) ToArchiveRuleFilterOutput() ArchiveRuleFilterOutput {
	return o
}

func (o ArchiveRuleFilterOutput) ToArchiveRuleFilterOutputWithContext(ctx context.Context) ArchiveRuleFilterOutput {
	return o
}

// Contains comparator.
func (o ArchiveRuleFilterOutput) Contains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ArchiveRuleFilter) []string { return v.Contains }).(pulumi.StringArrayOutput)
}

// Filter criteria.
func (o ArchiveRuleFilterOutput) Criteria() pulumi.StringOutput {
	return o.ApplyT(func(v ArchiveRuleFilter) string { return v.Criteria }).(pulumi.StringOutput)
}

// Equals comparator.
func (o ArchiveRuleFilterOutput) Eqs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ArchiveRuleFilter) []string { return v.Eqs }).(pulumi.StringArrayOutput)
}

// Boolean comparator.
func (o ArchiveRuleFilterOutput) Exists() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArchiveRuleFilter) *string { return v.Exists }).(pulumi.StringPtrOutput)
}

// Not Equals comparator.
func (o ArchiveRuleFilterOutput) Neqs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ArchiveRuleFilter) []string { return v.Neqs }).(pulumi.StringArrayOutput)
}

type ArchiveRuleFilterArrayOutput struct{ *pulumi.OutputState }

func (ArchiveRuleFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArchiveRuleFilter)(nil)).Elem()
}

func (o ArchiveRuleFilterArrayOutput) ToArchiveRuleFilterArrayOutput() ArchiveRuleFilterArrayOutput {
	return o
}

func (o ArchiveRuleFilterArrayOutput) ToArchiveRuleFilterArrayOutputWithContext(ctx context.Context) ArchiveRuleFilterArrayOutput {
	return o
}

func (o ArchiveRuleFilterArrayOutput) Index(i pulumi.IntInput) ArchiveRuleFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArchiveRuleFilter {
		return vs[0].([]ArchiveRuleFilter)[vs[1].(int)]
	}).(ArchiveRuleFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArchiveRuleFilterInput)(nil)).Elem(), ArchiveRuleFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArchiveRuleFilterArrayInput)(nil)).Elem(), ArchiveRuleFilterArray{})
	pulumi.RegisterOutputType(ArchiveRuleFilterOutput{})
	pulumi.RegisterOutputType(ArchiveRuleFilterArrayOutput{})
}
