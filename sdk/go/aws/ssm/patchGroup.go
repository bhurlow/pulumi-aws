// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Patch Group resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			production, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
//				ApprovedPatches: pulumi.StringArray{
//					pulumi.String("KB123456"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ssm.NewPatchGroup(ctx, "patchgroup", &ssm.PatchGroupArgs{
//				BaselineId: production.ID(),
//				PatchGroup: pulumi.String("patch-group-name"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PatchGroup struct {
	pulumi.CustomResourceState

	// The ID of the patch baseline to register the patch group with.
	BaselineId pulumi.StringOutput `pulumi:"baselineId"`
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup pulumi.StringOutput `pulumi:"patchGroup"`
}

// NewPatchGroup registers a new resource with the given unique name, arguments, and options.
func NewPatchGroup(ctx *pulumi.Context,
	name string, args *PatchGroupArgs, opts ...pulumi.ResourceOption) (*PatchGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaselineId == nil {
		return nil, errors.New("invalid value for required argument 'BaselineId'")
	}
	if args.PatchGroup == nil {
		return nil, errors.New("invalid value for required argument 'PatchGroup'")
	}
	var resource PatchGroup
	err := ctx.RegisterResource("aws:ssm/patchGroup:PatchGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchGroup gets an existing PatchGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchGroupState, opts ...pulumi.ResourceOption) (*PatchGroup, error) {
	var resource PatchGroup
	err := ctx.ReadResource("aws:ssm/patchGroup:PatchGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchGroup resources.
type patchGroupState struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId *string `pulumi:"baselineId"`
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup *string `pulumi:"patchGroup"`
}

type PatchGroupState struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId pulumi.StringPtrInput
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup pulumi.StringPtrInput
}

func (PatchGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchGroupState)(nil)).Elem()
}

type patchGroupArgs struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId string `pulumi:"baselineId"`
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup string `pulumi:"patchGroup"`
}

// The set of arguments for constructing a PatchGroup resource.
type PatchGroupArgs struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId pulumi.StringInput
	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup pulumi.StringInput
}

func (PatchGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchGroupArgs)(nil)).Elem()
}

type PatchGroupInput interface {
	pulumi.Input

	ToPatchGroupOutput() PatchGroupOutput
	ToPatchGroupOutputWithContext(ctx context.Context) PatchGroupOutput
}

func (*PatchGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchGroup)(nil)).Elem()
}

func (i *PatchGroup) ToPatchGroupOutput() PatchGroupOutput {
	return i.ToPatchGroupOutputWithContext(context.Background())
}

func (i *PatchGroup) ToPatchGroupOutputWithContext(ctx context.Context) PatchGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchGroupOutput)
}

// PatchGroupArrayInput is an input type that accepts PatchGroupArray and PatchGroupArrayOutput values.
// You can construct a concrete instance of `PatchGroupArrayInput` via:
//
//	PatchGroupArray{ PatchGroupArgs{...} }
type PatchGroupArrayInput interface {
	pulumi.Input

	ToPatchGroupArrayOutput() PatchGroupArrayOutput
	ToPatchGroupArrayOutputWithContext(context.Context) PatchGroupArrayOutput
}

type PatchGroupArray []PatchGroupInput

func (PatchGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchGroup)(nil)).Elem()
}

func (i PatchGroupArray) ToPatchGroupArrayOutput() PatchGroupArrayOutput {
	return i.ToPatchGroupArrayOutputWithContext(context.Background())
}

func (i PatchGroupArray) ToPatchGroupArrayOutputWithContext(ctx context.Context) PatchGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchGroupArrayOutput)
}

// PatchGroupMapInput is an input type that accepts PatchGroupMap and PatchGroupMapOutput values.
// You can construct a concrete instance of `PatchGroupMapInput` via:
//
//	PatchGroupMap{ "key": PatchGroupArgs{...} }
type PatchGroupMapInput interface {
	pulumi.Input

	ToPatchGroupMapOutput() PatchGroupMapOutput
	ToPatchGroupMapOutputWithContext(context.Context) PatchGroupMapOutput
}

type PatchGroupMap map[string]PatchGroupInput

func (PatchGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchGroup)(nil)).Elem()
}

func (i PatchGroupMap) ToPatchGroupMapOutput() PatchGroupMapOutput {
	return i.ToPatchGroupMapOutputWithContext(context.Background())
}

func (i PatchGroupMap) ToPatchGroupMapOutputWithContext(ctx context.Context) PatchGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchGroupMapOutput)
}

type PatchGroupOutput struct{ *pulumi.OutputState }

func (PatchGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchGroup)(nil)).Elem()
}

func (o PatchGroupOutput) ToPatchGroupOutput() PatchGroupOutput {
	return o
}

func (o PatchGroupOutput) ToPatchGroupOutputWithContext(ctx context.Context) PatchGroupOutput {
	return o
}

// The ID of the patch baseline to register the patch group with.
func (o PatchGroupOutput) BaselineId() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchGroup) pulumi.StringOutput { return v.BaselineId }).(pulumi.StringOutput)
}

// The name of the patch group that should be registered with the patch baseline.
func (o PatchGroupOutput) PatchGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchGroup) pulumi.StringOutput { return v.PatchGroup }).(pulumi.StringOutput)
}

type PatchGroupArrayOutput struct{ *pulumi.OutputState }

func (PatchGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchGroup)(nil)).Elem()
}

func (o PatchGroupArrayOutput) ToPatchGroupArrayOutput() PatchGroupArrayOutput {
	return o
}

func (o PatchGroupArrayOutput) ToPatchGroupArrayOutputWithContext(ctx context.Context) PatchGroupArrayOutput {
	return o
}

func (o PatchGroupArrayOutput) Index(i pulumi.IntInput) PatchGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PatchGroup {
		return vs[0].([]*PatchGroup)[vs[1].(int)]
	}).(PatchGroupOutput)
}

type PatchGroupMapOutput struct{ *pulumi.OutputState }

func (PatchGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchGroup)(nil)).Elem()
}

func (o PatchGroupMapOutput) ToPatchGroupMapOutput() PatchGroupMapOutput {
	return o
}

func (o PatchGroupMapOutput) ToPatchGroupMapOutputWithContext(ctx context.Context) PatchGroupMapOutput {
	return o
}

func (o PatchGroupMapOutput) MapIndex(k pulumi.StringInput) PatchGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PatchGroup {
		return vs[0].(map[string]*PatchGroup)[vs[1].(string)]
	}).(PatchGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PatchGroupInput)(nil)).Elem(), &PatchGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchGroupArrayInput)(nil)).Elem(), PatchGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchGroupMapInput)(nil)).Elem(), PatchGroupMap{})
	pulumi.RegisterOutputType(PatchGroupOutput{})
	pulumi.RegisterOutputType(PatchGroupArrayOutput{})
	pulumi.RegisterOutputType(PatchGroupMapOutput{})
}
