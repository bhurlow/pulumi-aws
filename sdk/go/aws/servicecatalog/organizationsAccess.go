// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Service Catalog AWS Organizations Access, a portfolio sharing feature through AWS Organizations. This allows Service Catalog to receive updates on your organization in order to sync your shares with the current structure. This resource will prompt AWS to set `organizations:EnableAWSServiceAccess` on your behalf so that your shares can be in sync with any changes in your AWS Organizations structure.
//
// > **NOTE:** This resource can only be used by the management account in the organization. In other words, a delegated administrator is not authorized to use the resource.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/servicecatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicecatalog.NewOrganizationsAccess(ctx, "example", &servicecatalog.OrganizationsAccessArgs{
//				Enabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type OrganizationsAccess struct {
	pulumi.CustomResourceState

	// Whether to enable AWS Organizations access.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
}

// NewOrganizationsAccess registers a new resource with the given unique name, arguments, and options.
func NewOrganizationsAccess(ctx *pulumi.Context,
	name string, args *OrganizationsAccessArgs, opts ...pulumi.ResourceOption) (*OrganizationsAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	var resource OrganizationsAccess
	err := ctx.RegisterResource("aws:servicecatalog/organizationsAccess:OrganizationsAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationsAccess gets an existing OrganizationsAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationsAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationsAccessState, opts ...pulumi.ResourceOption) (*OrganizationsAccess, error) {
	var resource OrganizationsAccess
	err := ctx.ReadResource("aws:servicecatalog/organizationsAccess:OrganizationsAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationsAccess resources.
type organizationsAccessState struct {
	// Whether to enable AWS Organizations access.
	Enabled *bool `pulumi:"enabled"`
}

type OrganizationsAccessState struct {
	// Whether to enable AWS Organizations access.
	Enabled pulumi.BoolPtrInput
}

func (OrganizationsAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationsAccessState)(nil)).Elem()
}

type organizationsAccessArgs struct {
	// Whether to enable AWS Organizations access.
	Enabled bool `pulumi:"enabled"`
}

// The set of arguments for constructing a OrganizationsAccess resource.
type OrganizationsAccessArgs struct {
	// Whether to enable AWS Organizations access.
	Enabled pulumi.BoolInput
}

func (OrganizationsAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationsAccessArgs)(nil)).Elem()
}

type OrganizationsAccessInput interface {
	pulumi.Input

	ToOrganizationsAccessOutput() OrganizationsAccessOutput
	ToOrganizationsAccessOutputWithContext(ctx context.Context) OrganizationsAccessOutput
}

func (*OrganizationsAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationsAccess)(nil)).Elem()
}

func (i *OrganizationsAccess) ToOrganizationsAccessOutput() OrganizationsAccessOutput {
	return i.ToOrganizationsAccessOutputWithContext(context.Background())
}

func (i *OrganizationsAccess) ToOrganizationsAccessOutputWithContext(ctx context.Context) OrganizationsAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationsAccessOutput)
}

// OrganizationsAccessArrayInput is an input type that accepts OrganizationsAccessArray and OrganizationsAccessArrayOutput values.
// You can construct a concrete instance of `OrganizationsAccessArrayInput` via:
//
//	OrganizationsAccessArray{ OrganizationsAccessArgs{...} }
type OrganizationsAccessArrayInput interface {
	pulumi.Input

	ToOrganizationsAccessArrayOutput() OrganizationsAccessArrayOutput
	ToOrganizationsAccessArrayOutputWithContext(context.Context) OrganizationsAccessArrayOutput
}

type OrganizationsAccessArray []OrganizationsAccessInput

func (OrganizationsAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationsAccess)(nil)).Elem()
}

func (i OrganizationsAccessArray) ToOrganizationsAccessArrayOutput() OrganizationsAccessArrayOutput {
	return i.ToOrganizationsAccessArrayOutputWithContext(context.Background())
}

func (i OrganizationsAccessArray) ToOrganizationsAccessArrayOutputWithContext(ctx context.Context) OrganizationsAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationsAccessArrayOutput)
}

// OrganizationsAccessMapInput is an input type that accepts OrganizationsAccessMap and OrganizationsAccessMapOutput values.
// You can construct a concrete instance of `OrganizationsAccessMapInput` via:
//
//	OrganizationsAccessMap{ "key": OrganizationsAccessArgs{...} }
type OrganizationsAccessMapInput interface {
	pulumi.Input

	ToOrganizationsAccessMapOutput() OrganizationsAccessMapOutput
	ToOrganizationsAccessMapOutputWithContext(context.Context) OrganizationsAccessMapOutput
}

type OrganizationsAccessMap map[string]OrganizationsAccessInput

func (OrganizationsAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationsAccess)(nil)).Elem()
}

func (i OrganizationsAccessMap) ToOrganizationsAccessMapOutput() OrganizationsAccessMapOutput {
	return i.ToOrganizationsAccessMapOutputWithContext(context.Background())
}

func (i OrganizationsAccessMap) ToOrganizationsAccessMapOutputWithContext(ctx context.Context) OrganizationsAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationsAccessMapOutput)
}

type OrganizationsAccessOutput struct{ *pulumi.OutputState }

func (OrganizationsAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationsAccess)(nil)).Elem()
}

func (o OrganizationsAccessOutput) ToOrganizationsAccessOutput() OrganizationsAccessOutput {
	return o
}

func (o OrganizationsAccessOutput) ToOrganizationsAccessOutputWithContext(ctx context.Context) OrganizationsAccessOutput {
	return o
}

// Whether to enable AWS Organizations access.
func (o OrganizationsAccessOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrganizationsAccess) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

type OrganizationsAccessArrayOutput struct{ *pulumi.OutputState }

func (OrganizationsAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationsAccess)(nil)).Elem()
}

func (o OrganizationsAccessArrayOutput) ToOrganizationsAccessArrayOutput() OrganizationsAccessArrayOutput {
	return o
}

func (o OrganizationsAccessArrayOutput) ToOrganizationsAccessArrayOutputWithContext(ctx context.Context) OrganizationsAccessArrayOutput {
	return o
}

func (o OrganizationsAccessArrayOutput) Index(i pulumi.IntInput) OrganizationsAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationsAccess {
		return vs[0].([]*OrganizationsAccess)[vs[1].(int)]
	}).(OrganizationsAccessOutput)
}

type OrganizationsAccessMapOutput struct{ *pulumi.OutputState }

func (OrganizationsAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationsAccess)(nil)).Elem()
}

func (o OrganizationsAccessMapOutput) ToOrganizationsAccessMapOutput() OrganizationsAccessMapOutput {
	return o
}

func (o OrganizationsAccessMapOutput) ToOrganizationsAccessMapOutputWithContext(ctx context.Context) OrganizationsAccessMapOutput {
	return o
}

func (o OrganizationsAccessMapOutput) MapIndex(k pulumi.StringInput) OrganizationsAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationsAccess {
		return vs[0].(map[string]*OrganizationsAccess)[vs[1].(string)]
	}).(OrganizationsAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationsAccessInput)(nil)).Elem(), &OrganizationsAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationsAccessArrayInput)(nil)).Elem(), OrganizationsAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationsAccessMapInput)(nil)).Elem(), OrganizationsAccessMap{})
	pulumi.RegisterOutputType(OrganizationsAccessOutput{})
	pulumi.RegisterOutputType(OrganizationsAccessArrayOutput{})
	pulumi.RegisterOutputType(OrganizationsAccessMapOutput{})
}
