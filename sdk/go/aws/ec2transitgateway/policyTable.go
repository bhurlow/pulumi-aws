// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EC2 Transit Gateway Policy Table.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2transitgateway.NewPolicyTable(ctx, "example", &ec2transitgateway.PolicyTableArgs{
//				TransitGatewayId: pulumi.Any(aws_ec2_transit_gateway.Example.Id),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("Example Policy Table"),
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
//	to = aws_ec2_transit_gateway_policy_table.example
//
//	id = "tgw-rtb-12345678" } Using `pulumi import`, import `aws_ec2_transit_gateway_policy_table` using the EC2 Transit Gateway Policy Table identifier. For exampleconsole % pulumi import aws_ec2_transit_gateway_policy_table.example tgw-rtb-12345678
type PolicyTable struct {
	pulumi.CustomResourceState

	// EC2 Transit Gateway Policy Table Amazon Resource Name (ARN).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The state of the EC2 Transit Gateway Policy Table.
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// EC2 Transit Gateway identifier.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
}

// NewPolicyTable registers a new resource with the given unique name, arguments, and options.
func NewPolicyTable(ctx *pulumi.Context,
	name string, args *PolicyTableArgs, opts ...pulumi.ResourceOption) (*PolicyTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyTable
	err := ctx.RegisterResource("aws:ec2transitgateway/policyTable:PolicyTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTable gets an existing PolicyTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTableState, opts ...pulumi.ResourceOption) (*PolicyTable, error) {
	var resource PolicyTable
	err := ctx.ReadResource("aws:ec2transitgateway/policyTable:PolicyTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTable resources.
type policyTableState struct {
	// EC2 Transit Gateway Policy Table Amazon Resource Name (ARN).
	Arn *string `pulumi:"arn"`
	// The state of the EC2 Transit Gateway Policy Table.
	State *string `pulumi:"state"`
	// Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// EC2 Transit Gateway identifier.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
}

type PolicyTableState struct {
	// EC2 Transit Gateway Policy Table Amazon Resource Name (ARN).
	Arn pulumi.StringPtrInput
	// The state of the EC2 Transit Gateway Policy Table.
	State pulumi.StringPtrInput
	// Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// EC2 Transit Gateway identifier.
	TransitGatewayId pulumi.StringPtrInput
}

func (PolicyTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTableState)(nil)).Elem()
}

type policyTableArgs struct {
	// Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// EC2 Transit Gateway identifier.
	TransitGatewayId string `pulumi:"transitGatewayId"`
}

// The set of arguments for constructing a PolicyTable resource.
type PolicyTableArgs struct {
	// Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// EC2 Transit Gateway identifier.
	TransitGatewayId pulumi.StringInput
}

func (PolicyTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTableArgs)(nil)).Elem()
}

type PolicyTableInput interface {
	pulumi.Input

	ToPolicyTableOutput() PolicyTableOutput
	ToPolicyTableOutputWithContext(ctx context.Context) PolicyTableOutput
}

func (*PolicyTable) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTable)(nil)).Elem()
}

func (i *PolicyTable) ToPolicyTableOutput() PolicyTableOutput {
	return i.ToPolicyTableOutputWithContext(context.Background())
}

func (i *PolicyTable) ToPolicyTableOutputWithContext(ctx context.Context) PolicyTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTableOutput)
}

// PolicyTableArrayInput is an input type that accepts PolicyTableArray and PolicyTableArrayOutput values.
// You can construct a concrete instance of `PolicyTableArrayInput` via:
//
//	PolicyTableArray{ PolicyTableArgs{...} }
type PolicyTableArrayInput interface {
	pulumi.Input

	ToPolicyTableArrayOutput() PolicyTableArrayOutput
	ToPolicyTableArrayOutputWithContext(context.Context) PolicyTableArrayOutput
}

type PolicyTableArray []PolicyTableInput

func (PolicyTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyTable)(nil)).Elem()
}

func (i PolicyTableArray) ToPolicyTableArrayOutput() PolicyTableArrayOutput {
	return i.ToPolicyTableArrayOutputWithContext(context.Background())
}

func (i PolicyTableArray) ToPolicyTableArrayOutputWithContext(ctx context.Context) PolicyTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTableArrayOutput)
}

// PolicyTableMapInput is an input type that accepts PolicyTableMap and PolicyTableMapOutput values.
// You can construct a concrete instance of `PolicyTableMapInput` via:
//
//	PolicyTableMap{ "key": PolicyTableArgs{...} }
type PolicyTableMapInput interface {
	pulumi.Input

	ToPolicyTableMapOutput() PolicyTableMapOutput
	ToPolicyTableMapOutputWithContext(context.Context) PolicyTableMapOutput
}

type PolicyTableMap map[string]PolicyTableInput

func (PolicyTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyTable)(nil)).Elem()
}

func (i PolicyTableMap) ToPolicyTableMapOutput() PolicyTableMapOutput {
	return i.ToPolicyTableMapOutputWithContext(context.Background())
}

func (i PolicyTableMap) ToPolicyTableMapOutputWithContext(ctx context.Context) PolicyTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTableMapOutput)
}

type PolicyTableOutput struct{ *pulumi.OutputState }

func (PolicyTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTable)(nil)).Elem()
}

func (o PolicyTableOutput) ToPolicyTableOutput() PolicyTableOutput {
	return o
}

func (o PolicyTableOutput) ToPolicyTableOutputWithContext(ctx context.Context) PolicyTableOutput {
	return o
}

// EC2 Transit Gateway Policy Table Amazon Resource Name (ARN).
func (o PolicyTableOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTable) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The state of the EC2 Transit Gateway Policy Table.
func (o PolicyTableOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTable) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Key-value tags for the EC2 Transit Gateway Policy Table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o PolicyTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PolicyTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o PolicyTableOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PolicyTable) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// EC2 Transit Gateway identifier.
func (o PolicyTableOutput) TransitGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyTable) pulumi.StringOutput { return v.TransitGatewayId }).(pulumi.StringOutput)
}

type PolicyTableArrayOutput struct{ *pulumi.OutputState }

func (PolicyTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyTable)(nil)).Elem()
}

func (o PolicyTableArrayOutput) ToPolicyTableArrayOutput() PolicyTableArrayOutput {
	return o
}

func (o PolicyTableArrayOutput) ToPolicyTableArrayOutputWithContext(ctx context.Context) PolicyTableArrayOutput {
	return o
}

func (o PolicyTableArrayOutput) Index(i pulumi.IntInput) PolicyTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyTable {
		return vs[0].([]*PolicyTable)[vs[1].(int)]
	}).(PolicyTableOutput)
}

type PolicyTableMapOutput struct{ *pulumi.OutputState }

func (PolicyTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyTable)(nil)).Elem()
}

func (o PolicyTableMapOutput) ToPolicyTableMapOutput() PolicyTableMapOutput {
	return o
}

func (o PolicyTableMapOutput) ToPolicyTableMapOutputWithContext(ctx context.Context) PolicyTableMapOutput {
	return o
}

func (o PolicyTableMapOutput) MapIndex(k pulumi.StringInput) PolicyTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyTable {
		return vs[0].(map[string]*PolicyTable)[vs[1].(string)]
	}).(PolicyTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTableInput)(nil)).Elem(), &PolicyTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTableArrayInput)(nil)).Elem(), PolicyTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTableMapInput)(nil)).Elem(), PolicyTableMap{})
	pulumi.RegisterOutputType(PolicyTableOutput{})
	pulumi.RegisterOutputType(PolicyTableArrayOutput{})
	pulumi.RegisterOutputType(PolicyTableMapOutput{})
}
