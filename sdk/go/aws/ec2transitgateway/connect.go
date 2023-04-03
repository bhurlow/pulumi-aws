// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EC2 Transit Gateway Connect.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2transitgateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ec2transitgateway.NewVpcAttachment(ctx, "example", &ec2transitgateway.VpcAttachmentArgs{
//				SubnetIds: pulumi.StringArray{
//					aws_subnet.Example.Id,
//				},
//				TransitGatewayId: pulumi.Any(aws_ec2_transit_gateway.Example.Id),
//				VpcId:            pulumi.Any(aws_vpc.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2transitgateway.NewConnect(ctx, "attachment", &ec2transitgateway.ConnectArgs{
//				TransportAttachmentId: example.ID(),
//				TransitGatewayId:      pulumi.Any(aws_ec2_transit_gateway.Example.Id),
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
// `aws_ec2_transit_gateway_connect` can be imported by using the EC2 Transit Gateway Connect identifier, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2transitgateway/connect:Connect example tgw-attach-12345678
//
// ```
type Connect struct {
	pulumi.CustomResourceState

	// The tunnel protocol. Valida values: `gre`. Default is `gre`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation pulumi.BoolPtrOutput `pulumi:"transitGatewayDefaultRouteTableAssociation"`
	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation pulumi.BoolPtrOutput `pulumi:"transitGatewayDefaultRouteTablePropagation"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
	// The underlaying VPC attachment
	TransportAttachmentId pulumi.StringOutput `pulumi:"transportAttachmentId"`
}

// NewConnect registers a new resource with the given unique name, arguments, and options.
func NewConnect(ctx *pulumi.Context,
	name string, args *ConnectArgs, opts ...pulumi.ResourceOption) (*Connect, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayId'")
	}
	if args.TransportAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'TransportAttachmentId'")
	}
	var resource Connect
	err := ctx.RegisterResource("aws:ec2transitgateway/connect:Connect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnect gets an existing Connect resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnect(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectState, opts ...pulumi.ResourceOption) (*Connect, error) {
	var resource Connect
	err := ctx.ReadResource("aws:ec2transitgateway/connect:Connect", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connect resources.
type connectState struct {
	// The tunnel protocol. Valida values: `gre`. Default is `gre`.
	Protocol *string `pulumi:"protocol"`
	// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation *bool `pulumi:"transitGatewayDefaultRouteTableAssociation"`
	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation *bool `pulumi:"transitGatewayDefaultRouteTablePropagation"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// The underlaying VPC attachment
	TransportAttachmentId *string `pulumi:"transportAttachmentId"`
}

type ConnectState struct {
	// The tunnel protocol. Valida values: `gre`. Default is `gre`.
	Protocol pulumi.StringPtrInput
	// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation pulumi.BoolPtrInput
	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation pulumi.BoolPtrInput
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrInput
	// The underlaying VPC attachment
	TransportAttachmentId pulumi.StringPtrInput
}

func (ConnectState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectState)(nil)).Elem()
}

type connectArgs struct {
	// The tunnel protocol. Valida values: `gre`. Default is `gre`.
	Protocol *string `pulumi:"protocol"`
	// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation *bool `pulumi:"transitGatewayDefaultRouteTableAssociation"`
	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation *bool `pulumi:"transitGatewayDefaultRouteTablePropagation"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId string `pulumi:"transitGatewayId"`
	// The underlaying VPC attachment
	TransportAttachmentId string `pulumi:"transportAttachmentId"`
}

// The set of arguments for constructing a Connect resource.
type ConnectArgs struct {
	// The tunnel protocol. Valida values: `gre`. Default is `gre`.
	Protocol pulumi.StringPtrInput
	// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation pulumi.BoolPtrInput
	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation pulumi.BoolPtrInput
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringInput
	// The underlaying VPC attachment
	TransportAttachmentId pulumi.StringInput
}

func (ConnectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectArgs)(nil)).Elem()
}

type ConnectInput interface {
	pulumi.Input

	ToConnectOutput() ConnectOutput
	ToConnectOutputWithContext(ctx context.Context) ConnectOutput
}

func (*Connect) ElementType() reflect.Type {
	return reflect.TypeOf((**Connect)(nil)).Elem()
}

func (i *Connect) ToConnectOutput() ConnectOutput {
	return i.ToConnectOutputWithContext(context.Background())
}

func (i *Connect) ToConnectOutputWithContext(ctx context.Context) ConnectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectOutput)
}

// ConnectArrayInput is an input type that accepts ConnectArray and ConnectArrayOutput values.
// You can construct a concrete instance of `ConnectArrayInput` via:
//
//	ConnectArray{ ConnectArgs{...} }
type ConnectArrayInput interface {
	pulumi.Input

	ToConnectArrayOutput() ConnectArrayOutput
	ToConnectArrayOutputWithContext(context.Context) ConnectArrayOutput
}

type ConnectArray []ConnectInput

func (ConnectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connect)(nil)).Elem()
}

func (i ConnectArray) ToConnectArrayOutput() ConnectArrayOutput {
	return i.ToConnectArrayOutputWithContext(context.Background())
}

func (i ConnectArray) ToConnectArrayOutputWithContext(ctx context.Context) ConnectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectArrayOutput)
}

// ConnectMapInput is an input type that accepts ConnectMap and ConnectMapOutput values.
// You can construct a concrete instance of `ConnectMapInput` via:
//
//	ConnectMap{ "key": ConnectArgs{...} }
type ConnectMapInput interface {
	pulumi.Input

	ToConnectMapOutput() ConnectMapOutput
	ToConnectMapOutputWithContext(context.Context) ConnectMapOutput
}

type ConnectMap map[string]ConnectInput

func (ConnectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connect)(nil)).Elem()
}

func (i ConnectMap) ToConnectMapOutput() ConnectMapOutput {
	return i.ToConnectMapOutputWithContext(context.Background())
}

func (i ConnectMap) ToConnectMapOutputWithContext(ctx context.Context) ConnectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectMapOutput)
}

type ConnectOutput struct{ *pulumi.OutputState }

func (ConnectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connect)(nil)).Elem()
}

func (o ConnectOutput) ToConnectOutput() ConnectOutput {
	return o
}

func (o ConnectOutput) ToConnectOutputWithContext(ctx context.Context) ConnectOutput {
	return o
}

// The tunnel protocol. Valida values: `gre`. Default is `gre`.
func (o ConnectOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connect) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ConnectOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connect) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ConnectOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connect) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
func (o ConnectOutput) TransitGatewayDefaultRouteTableAssociation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connect) pulumi.BoolPtrOutput { return v.TransitGatewayDefaultRouteTableAssociation }).(pulumi.BoolPtrOutput)
}

// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
func (o ConnectOutput) TransitGatewayDefaultRouteTablePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connect) pulumi.BoolPtrOutput { return v.TransitGatewayDefaultRouteTablePropagation }).(pulumi.BoolPtrOutput)
}

// Identifier of EC2 Transit Gateway.
func (o ConnectOutput) TransitGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connect) pulumi.StringOutput { return v.TransitGatewayId }).(pulumi.StringOutput)
}

// The underlaying VPC attachment
func (o ConnectOutput) TransportAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connect) pulumi.StringOutput { return v.TransportAttachmentId }).(pulumi.StringOutput)
}

type ConnectArrayOutput struct{ *pulumi.OutputState }

func (ConnectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connect)(nil)).Elem()
}

func (o ConnectArrayOutput) ToConnectArrayOutput() ConnectArrayOutput {
	return o
}

func (o ConnectArrayOutput) ToConnectArrayOutputWithContext(ctx context.Context) ConnectArrayOutput {
	return o
}

func (o ConnectArrayOutput) Index(i pulumi.IntInput) ConnectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connect {
		return vs[0].([]*Connect)[vs[1].(int)]
	}).(ConnectOutput)
}

type ConnectMapOutput struct{ *pulumi.OutputState }

func (ConnectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connect)(nil)).Elem()
}

func (o ConnectMapOutput) ToConnectMapOutput() ConnectMapOutput {
	return o
}

func (o ConnectMapOutput) ToConnectMapOutputWithContext(ctx context.Context) ConnectMapOutput {
	return o
}

func (o ConnectMapOutput) MapIndex(k pulumi.StringInput) ConnectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connect {
		return vs[0].(map[string]*Connect)[vs[1].(string)]
	}).(ConnectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectInput)(nil)).Elem(), &Connect{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectArrayInput)(nil)).Elem(), ConnectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectMapInput)(nil)).Elem(), ConnectMap{})
	pulumi.RegisterOutputType(ConnectOutput{})
	pulumi.RegisterOutputType(ConnectArrayOutput{})
	pulumi.RegisterOutputType(ConnectMapOutput{})
}
