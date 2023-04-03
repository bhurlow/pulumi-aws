// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Service Catalog Constraint.
//
// > **NOTE:** This resource does not associate a Service Catalog product and portfolio. However, the product and portfolio must be associated (see the `servicecatalog.ProductPortfolioAssociation` resource) prior to creating a constraint or you will receive an error.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/servicecatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"RoleArn": "arn:aws:iam::123456789012:role/LaunchRole",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = servicecatalog.NewConstraint(ctx, "example", &servicecatalog.ConstraintArgs{
//				Description: pulumi.String("Back off, man. I'm a scientist."),
//				PortfolioId: pulumi.Any(aws_servicecatalog_portfolio.Example.Id),
//				ProductId:   pulumi.Any(aws_servicecatalog_product.Example.Id),
//				Type:        pulumi.String("LAUNCH"),
//				Parameters:  pulumi.String(json0),
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
// `aws_servicecatalog_constraint` can be imported using the constraint ID, e.g.,
//
// ```sh
//
//	$ pulumi import aws:servicecatalog/constraint:Constraint example cons-nmdkb6cgxfcrs
//
// ```
type Constraint struct {
	pulumi.CustomResourceState

	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// Description of the constraint.
	Description pulumi.StringOutput `pulumi:"description"`
	// Owner of the constraint.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters pulumi.StringOutput `pulumi:"parameters"`
	// Portfolio identifier.
	PortfolioId pulumi.StringOutput `pulumi:"portfolioId"`
	// Product identifier.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	Status    pulumi.StringOutput `pulumi:"status"`
	// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `RESOURCE_UPDATE`, `STACKSET`, and `TEMPLATE`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConstraint registers a new resource with the given unique name, arguments, and options.
func NewConstraint(ctx *pulumi.Context,
	name string, args *ConstraintArgs, opts ...pulumi.ResourceOption) (*Constraint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Constraint
	err := ctx.RegisterResource("aws:servicecatalog/constraint:Constraint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConstraint gets an existing Constraint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConstraint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConstraintState, opts ...pulumi.ResourceOption) (*Constraint, error) {
	var resource Constraint
	err := ctx.ReadResource("aws:servicecatalog/constraint:Constraint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Constraint resources.
type constraintState struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// Description of the constraint.
	Description *string `pulumi:"description"`
	// Owner of the constraint.
	Owner *string `pulumi:"owner"`
	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters *string `pulumi:"parameters"`
	// Portfolio identifier.
	PortfolioId *string `pulumi:"portfolioId"`
	// Product identifier.
	ProductId *string `pulumi:"productId"`
	Status    *string `pulumi:"status"`
	// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `RESOURCE_UPDATE`, `STACKSET`, and `TEMPLATE`.
	Type *string `pulumi:"type"`
}

type ConstraintState struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrInput
	// Description of the constraint.
	Description pulumi.StringPtrInput
	// Owner of the constraint.
	Owner pulumi.StringPtrInput
	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters pulumi.StringPtrInput
	// Portfolio identifier.
	PortfolioId pulumi.StringPtrInput
	// Product identifier.
	ProductId pulumi.StringPtrInput
	Status    pulumi.StringPtrInput
	// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `RESOURCE_UPDATE`, `STACKSET`, and `TEMPLATE`.
	Type pulumi.StringPtrInput
}

func (ConstraintState) ElementType() reflect.Type {
	return reflect.TypeOf((*constraintState)(nil)).Elem()
}

type constraintArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// Description of the constraint.
	Description *string `pulumi:"description"`
	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters string `pulumi:"parameters"`
	// Portfolio identifier.
	PortfolioId string `pulumi:"portfolioId"`
	// Product identifier.
	ProductId string `pulumi:"productId"`
	// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `RESOURCE_UPDATE`, `STACKSET`, and `TEMPLATE`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Constraint resource.
type ConstraintArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrInput
	// Description of the constraint.
	Description pulumi.StringPtrInput
	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters pulumi.StringInput
	// Portfolio identifier.
	PortfolioId pulumi.StringInput
	// Product identifier.
	ProductId pulumi.StringInput
	// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `RESOURCE_UPDATE`, `STACKSET`, and `TEMPLATE`.
	Type pulumi.StringInput
}

func (ConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*constraintArgs)(nil)).Elem()
}

type ConstraintInput interface {
	pulumi.Input

	ToConstraintOutput() ConstraintOutput
	ToConstraintOutputWithContext(ctx context.Context) ConstraintOutput
}

func (*Constraint) ElementType() reflect.Type {
	return reflect.TypeOf((**Constraint)(nil)).Elem()
}

func (i *Constraint) ToConstraintOutput() ConstraintOutput {
	return i.ToConstraintOutputWithContext(context.Background())
}

func (i *Constraint) ToConstraintOutputWithContext(ctx context.Context) ConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConstraintOutput)
}

// ConstraintArrayInput is an input type that accepts ConstraintArray and ConstraintArrayOutput values.
// You can construct a concrete instance of `ConstraintArrayInput` via:
//
//	ConstraintArray{ ConstraintArgs{...} }
type ConstraintArrayInput interface {
	pulumi.Input

	ToConstraintArrayOutput() ConstraintArrayOutput
	ToConstraintArrayOutputWithContext(context.Context) ConstraintArrayOutput
}

type ConstraintArray []ConstraintInput

func (ConstraintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Constraint)(nil)).Elem()
}

func (i ConstraintArray) ToConstraintArrayOutput() ConstraintArrayOutput {
	return i.ToConstraintArrayOutputWithContext(context.Background())
}

func (i ConstraintArray) ToConstraintArrayOutputWithContext(ctx context.Context) ConstraintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConstraintArrayOutput)
}

// ConstraintMapInput is an input type that accepts ConstraintMap and ConstraintMapOutput values.
// You can construct a concrete instance of `ConstraintMapInput` via:
//
//	ConstraintMap{ "key": ConstraintArgs{...} }
type ConstraintMapInput interface {
	pulumi.Input

	ToConstraintMapOutput() ConstraintMapOutput
	ToConstraintMapOutputWithContext(context.Context) ConstraintMapOutput
}

type ConstraintMap map[string]ConstraintInput

func (ConstraintMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Constraint)(nil)).Elem()
}

func (i ConstraintMap) ToConstraintMapOutput() ConstraintMapOutput {
	return i.ToConstraintMapOutputWithContext(context.Background())
}

func (i ConstraintMap) ToConstraintMapOutputWithContext(ctx context.Context) ConstraintMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConstraintMapOutput)
}

type ConstraintOutput struct{ *pulumi.OutputState }

func (ConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Constraint)(nil)).Elem()
}

func (o ConstraintOutput) ToConstraintOutput() ConstraintOutput {
	return o
}

func (o ConstraintOutput) ToConstraintOutputWithContext(ctx context.Context) ConstraintOutput {
	return o
}

// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
func (o ConstraintOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

// Description of the constraint.
func (o ConstraintOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Owner of the constraint.
func (o ConstraintOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
func (o ConstraintOutput) Parameters() pulumi.StringOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringOutput { return v.Parameters }).(pulumi.StringOutput)
}

// Portfolio identifier.
func (o ConstraintOutput) PortfolioId() pulumi.StringOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringOutput { return v.PortfolioId }).(pulumi.StringOutput)
}

// Product identifier.
func (o ConstraintOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o ConstraintOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `RESOURCE_UPDATE`, `STACKSET`, and `TEMPLATE`.
func (o ConstraintOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Constraint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ConstraintArrayOutput struct{ *pulumi.OutputState }

func (ConstraintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Constraint)(nil)).Elem()
}

func (o ConstraintArrayOutput) ToConstraintArrayOutput() ConstraintArrayOutput {
	return o
}

func (o ConstraintArrayOutput) ToConstraintArrayOutputWithContext(ctx context.Context) ConstraintArrayOutput {
	return o
}

func (o ConstraintArrayOutput) Index(i pulumi.IntInput) ConstraintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Constraint {
		return vs[0].([]*Constraint)[vs[1].(int)]
	}).(ConstraintOutput)
}

type ConstraintMapOutput struct{ *pulumi.OutputState }

func (ConstraintMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Constraint)(nil)).Elem()
}

func (o ConstraintMapOutput) ToConstraintMapOutput() ConstraintMapOutput {
	return o
}

func (o ConstraintMapOutput) ToConstraintMapOutputWithContext(ctx context.Context) ConstraintMapOutput {
	return o
}

func (o ConstraintMapOutput) MapIndex(k pulumi.StringInput) ConstraintOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Constraint {
		return vs[0].(map[string]*Constraint)[vs[1].(string)]
	}).(ConstraintOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConstraintInput)(nil)).Elem(), &Constraint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConstraintArrayInput)(nil)).Elem(), ConstraintArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConstraintMapInput)(nil)).Elem(), ConstraintMap{})
	pulumi.RegisterOutputType(ConstraintOutput{})
	pulumi.RegisterOutputType(ConstraintArrayOutput{})
	pulumi.RegisterOutputType(ConstraintMapOutput{})
}
