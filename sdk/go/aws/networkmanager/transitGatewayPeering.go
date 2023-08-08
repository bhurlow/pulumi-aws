// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a peering connection between an AWS Cloud WAN core network and an AWS Transit Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkmanager.NewTransitGatewayPeering(ctx, "example", &networkmanager.TransitGatewayPeeringArgs{
//				CoreNetworkId:     pulumi.Any(awscc_networkmanager_core_network.Example.Id),
//				TransitGatewayArn: pulumi.Any(aws_ec2_transit_gateway.Example.Arn),
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
//	to = aws_networkmanager_transit_gateway_peering.example
//
//	id = "peering-444555aaabbb11223" } Using `pulumi import`, import `aws_networkmanager_transit_gateway_peering` using the peering ID. For exampleconsole % pulumi import aws_networkmanager_transit_gateway_peering.example peering-444555aaabbb11223
type TransitGatewayPeering struct {
	pulumi.CustomResourceState

	// Peering Amazon Resource Name (ARN).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the core network.
	CoreNetworkArn pulumi.StringOutput `pulumi:"coreNetworkArn"`
	// The ID of a core network.
	CoreNetworkId pulumi.StringOutput `pulumi:"coreNetworkId"`
	// The edge location for the peer.
	EdgeLocation pulumi.StringOutput `pulumi:"edgeLocation"`
	// The ID of the account owner.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The type of peering. This will be `TRANSIT_GATEWAY`.
	PeeringType pulumi.StringOutput `pulumi:"peeringType"`
	// The resource ARN of the peer.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// Key-value tags for the peering. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ARN of the transit gateway for the peering request.
	TransitGatewayArn pulumi.StringOutput `pulumi:"transitGatewayArn"`
	// The ID of the transit gateway peering attachment.
	TransitGatewayPeeringAttachmentId pulumi.StringOutput `pulumi:"transitGatewayPeeringAttachmentId"`
}

// NewTransitGatewayPeering registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayPeering(ctx *pulumi.Context,
	name string, args *TransitGatewayPeeringArgs, opts ...pulumi.ResourceOption) (*TransitGatewayPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CoreNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'CoreNetworkId'")
	}
	if args.TransitGatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitGatewayPeering
	err := ctx.RegisterResource("aws:networkmanager/transitGatewayPeering:TransitGatewayPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayPeering gets an existing TransitGatewayPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayPeeringState, opts ...pulumi.ResourceOption) (*TransitGatewayPeering, error) {
	var resource TransitGatewayPeering
	err := ctx.ReadResource("aws:networkmanager/transitGatewayPeering:TransitGatewayPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayPeering resources.
type transitGatewayPeeringState struct {
	// Peering Amazon Resource Name (ARN).
	Arn *string `pulumi:"arn"`
	// The ARN of the core network.
	CoreNetworkArn *string `pulumi:"coreNetworkArn"`
	// The ID of a core network.
	CoreNetworkId *string `pulumi:"coreNetworkId"`
	// The edge location for the peer.
	EdgeLocation *string `pulumi:"edgeLocation"`
	// The ID of the account owner.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
	// The type of peering. This will be `TRANSIT_GATEWAY`.
	PeeringType *string `pulumi:"peeringType"`
	// The resource ARN of the peer.
	ResourceArn *string `pulumi:"resourceArn"`
	// Key-value tags for the peering. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ARN of the transit gateway for the peering request.
	TransitGatewayArn *string `pulumi:"transitGatewayArn"`
	// The ID of the transit gateway peering attachment.
	TransitGatewayPeeringAttachmentId *string `pulumi:"transitGatewayPeeringAttachmentId"`
}

type TransitGatewayPeeringState struct {
	// Peering Amazon Resource Name (ARN).
	Arn pulumi.StringPtrInput
	// The ARN of the core network.
	CoreNetworkArn pulumi.StringPtrInput
	// The ID of a core network.
	CoreNetworkId pulumi.StringPtrInput
	// The edge location for the peer.
	EdgeLocation pulumi.StringPtrInput
	// The ID of the account owner.
	OwnerAccountId pulumi.StringPtrInput
	// The type of peering. This will be `TRANSIT_GATEWAY`.
	PeeringType pulumi.StringPtrInput
	// The resource ARN of the peer.
	ResourceArn pulumi.StringPtrInput
	// Key-value tags for the peering. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The ARN of the transit gateway for the peering request.
	TransitGatewayArn pulumi.StringPtrInput
	// The ID of the transit gateway peering attachment.
	TransitGatewayPeeringAttachmentId pulumi.StringPtrInput
}

func (TransitGatewayPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayPeeringState)(nil)).Elem()
}

type transitGatewayPeeringArgs struct {
	// The ID of a core network.
	CoreNetworkId string `pulumi:"coreNetworkId"`
	// Key-value tags for the peering. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The ARN of the transit gateway for the peering request.
	TransitGatewayArn string `pulumi:"transitGatewayArn"`
}

// The set of arguments for constructing a TransitGatewayPeering resource.
type TransitGatewayPeeringArgs struct {
	// The ID of a core network.
	CoreNetworkId pulumi.StringInput
	// Key-value tags for the peering. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The ARN of the transit gateway for the peering request.
	TransitGatewayArn pulumi.StringInput
}

func (TransitGatewayPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayPeeringArgs)(nil)).Elem()
}

type TransitGatewayPeeringInput interface {
	pulumi.Input

	ToTransitGatewayPeeringOutput() TransitGatewayPeeringOutput
	ToTransitGatewayPeeringOutputWithContext(ctx context.Context) TransitGatewayPeeringOutput
}

func (*TransitGatewayPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGatewayPeering)(nil)).Elem()
}

func (i *TransitGatewayPeering) ToTransitGatewayPeeringOutput() TransitGatewayPeeringOutput {
	return i.ToTransitGatewayPeeringOutputWithContext(context.Background())
}

func (i *TransitGatewayPeering) ToTransitGatewayPeeringOutputWithContext(ctx context.Context) TransitGatewayPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayPeeringOutput)
}

// TransitGatewayPeeringArrayInput is an input type that accepts TransitGatewayPeeringArray and TransitGatewayPeeringArrayOutput values.
// You can construct a concrete instance of `TransitGatewayPeeringArrayInput` via:
//
//	TransitGatewayPeeringArray{ TransitGatewayPeeringArgs{...} }
type TransitGatewayPeeringArrayInput interface {
	pulumi.Input

	ToTransitGatewayPeeringArrayOutput() TransitGatewayPeeringArrayOutput
	ToTransitGatewayPeeringArrayOutputWithContext(context.Context) TransitGatewayPeeringArrayOutput
}

type TransitGatewayPeeringArray []TransitGatewayPeeringInput

func (TransitGatewayPeeringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitGatewayPeering)(nil)).Elem()
}

func (i TransitGatewayPeeringArray) ToTransitGatewayPeeringArrayOutput() TransitGatewayPeeringArrayOutput {
	return i.ToTransitGatewayPeeringArrayOutputWithContext(context.Background())
}

func (i TransitGatewayPeeringArray) ToTransitGatewayPeeringArrayOutputWithContext(ctx context.Context) TransitGatewayPeeringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayPeeringArrayOutput)
}

// TransitGatewayPeeringMapInput is an input type that accepts TransitGatewayPeeringMap and TransitGatewayPeeringMapOutput values.
// You can construct a concrete instance of `TransitGatewayPeeringMapInput` via:
//
//	TransitGatewayPeeringMap{ "key": TransitGatewayPeeringArgs{...} }
type TransitGatewayPeeringMapInput interface {
	pulumi.Input

	ToTransitGatewayPeeringMapOutput() TransitGatewayPeeringMapOutput
	ToTransitGatewayPeeringMapOutputWithContext(context.Context) TransitGatewayPeeringMapOutput
}

type TransitGatewayPeeringMap map[string]TransitGatewayPeeringInput

func (TransitGatewayPeeringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitGatewayPeering)(nil)).Elem()
}

func (i TransitGatewayPeeringMap) ToTransitGatewayPeeringMapOutput() TransitGatewayPeeringMapOutput {
	return i.ToTransitGatewayPeeringMapOutputWithContext(context.Background())
}

func (i TransitGatewayPeeringMap) ToTransitGatewayPeeringMapOutputWithContext(ctx context.Context) TransitGatewayPeeringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayPeeringMapOutput)
}

type TransitGatewayPeeringOutput struct{ *pulumi.OutputState }

func (TransitGatewayPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGatewayPeering)(nil)).Elem()
}

func (o TransitGatewayPeeringOutput) ToTransitGatewayPeeringOutput() TransitGatewayPeeringOutput {
	return o
}

func (o TransitGatewayPeeringOutput) ToTransitGatewayPeeringOutputWithContext(ctx context.Context) TransitGatewayPeeringOutput {
	return o
}

// Peering Amazon Resource Name (ARN).
func (o TransitGatewayPeeringOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ARN of the core network.
func (o TransitGatewayPeeringOutput) CoreNetworkArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.CoreNetworkArn }).(pulumi.StringOutput)
}

// The ID of a core network.
func (o TransitGatewayPeeringOutput) CoreNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.CoreNetworkId }).(pulumi.StringOutput)
}

// The edge location for the peer.
func (o TransitGatewayPeeringOutput) EdgeLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.EdgeLocation }).(pulumi.StringOutput)
}

// The ID of the account owner.
func (o TransitGatewayPeeringOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// The type of peering. This will be `TRANSIT_GATEWAY`.
func (o TransitGatewayPeeringOutput) PeeringType() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.PeeringType }).(pulumi.StringOutput)
}

// The resource ARN of the peer.
func (o TransitGatewayPeeringOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

// Key-value tags for the peering. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TransitGatewayPeeringOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o TransitGatewayPeeringOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The ARN of the transit gateway for the peering request.
func (o TransitGatewayPeeringOutput) TransitGatewayArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.TransitGatewayArn }).(pulumi.StringOutput)
}

// The ID of the transit gateway peering attachment.
func (o TransitGatewayPeeringOutput) TransitGatewayPeeringAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayPeering) pulumi.StringOutput { return v.TransitGatewayPeeringAttachmentId }).(pulumi.StringOutput)
}

type TransitGatewayPeeringArrayOutput struct{ *pulumi.OutputState }

func (TransitGatewayPeeringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitGatewayPeering)(nil)).Elem()
}

func (o TransitGatewayPeeringArrayOutput) ToTransitGatewayPeeringArrayOutput() TransitGatewayPeeringArrayOutput {
	return o
}

func (o TransitGatewayPeeringArrayOutput) ToTransitGatewayPeeringArrayOutputWithContext(ctx context.Context) TransitGatewayPeeringArrayOutput {
	return o
}

func (o TransitGatewayPeeringArrayOutput) Index(i pulumi.IntInput) TransitGatewayPeeringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitGatewayPeering {
		return vs[0].([]*TransitGatewayPeering)[vs[1].(int)]
	}).(TransitGatewayPeeringOutput)
}

type TransitGatewayPeeringMapOutput struct{ *pulumi.OutputState }

func (TransitGatewayPeeringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitGatewayPeering)(nil)).Elem()
}

func (o TransitGatewayPeeringMapOutput) ToTransitGatewayPeeringMapOutput() TransitGatewayPeeringMapOutput {
	return o
}

func (o TransitGatewayPeeringMapOutput) ToTransitGatewayPeeringMapOutputWithContext(ctx context.Context) TransitGatewayPeeringMapOutput {
	return o
}

func (o TransitGatewayPeeringMapOutput) MapIndex(k pulumi.StringInput) TransitGatewayPeeringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitGatewayPeering {
		return vs[0].(map[string]*TransitGatewayPeering)[vs[1].(string)]
	}).(TransitGatewayPeeringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayPeeringInput)(nil)).Elem(), &TransitGatewayPeering{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayPeeringArrayInput)(nil)).Elem(), TransitGatewayPeeringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayPeeringMapInput)(nil)).Elem(), TransitGatewayPeeringMap{})
	pulumi.RegisterOutputType(TransitGatewayPeeringOutput{})
	pulumi.RegisterOutputType(TransitGatewayPeeringArrayOutput{})
	pulumi.RegisterOutputType(TransitGatewayPeeringMapOutput{})
}
