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

// Associates a link to a device.
// A device can be associated to multiple links and a link can be associated to multiple devices.
// The device and link must be in the same global network and the same site.
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
//			_, err := networkmanager.NewLinkAssociation(ctx, "example", &networkmanager.LinkAssociationArgs{
//				GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
//				LinkId:          pulumi.Any(aws_networkmanager_link.Example.Id),
//				DeviceId:        pulumi.Any(aws_networkmanager_device.Example.Id),
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
//	to = aws_networkmanager_link_association.example
//
//	id = "global-network-0d47f6t230mz46dy4,link-444555aaabbb11223,device-07f6fd08867abc123" } Using `pulumi import`, import `aws_networkmanager_link_association` using the global network ID, link ID and device ID. For exampleconsole % pulumi import aws_networkmanager_link_association.example global-network-0d47f6t230mz46dy4,link-444555aaabbb11223,device-07f6fd08867abc123
type LinkAssociation struct {
	pulumi.CustomResourceState

	// The ID of the device.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId pulumi.StringOutput `pulumi:"globalNetworkId"`
	// The ID of the link.
	LinkId pulumi.StringOutput `pulumi:"linkId"`
}

// NewLinkAssociation registers a new resource with the given unique name, arguments, and options.
func NewLinkAssociation(ctx *pulumi.Context,
	name string, args *LinkAssociationArgs, opts ...pulumi.ResourceOption) (*LinkAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.GlobalNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'GlobalNetworkId'")
	}
	if args.LinkId == nil {
		return nil, errors.New("invalid value for required argument 'LinkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LinkAssociation
	err := ctx.RegisterResource("aws:networkmanager/linkAssociation:LinkAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkAssociation gets an existing LinkAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkAssociationState, opts ...pulumi.ResourceOption) (*LinkAssociation, error) {
	var resource LinkAssociation
	err := ctx.ReadResource("aws:networkmanager/linkAssociation:LinkAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkAssociation resources.
type linkAssociationState struct {
	// The ID of the device.
	DeviceId *string `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId *string `pulumi:"globalNetworkId"`
	// The ID of the link.
	LinkId *string `pulumi:"linkId"`
}

type LinkAssociationState struct {
	// The ID of the device.
	DeviceId pulumi.StringPtrInput
	// The ID of the global network.
	GlobalNetworkId pulumi.StringPtrInput
	// The ID of the link.
	LinkId pulumi.StringPtrInput
}

func (LinkAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAssociationState)(nil)).Elem()
}

type linkAssociationArgs struct {
	// The ID of the device.
	DeviceId string `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId string `pulumi:"globalNetworkId"`
	// The ID of the link.
	LinkId string `pulumi:"linkId"`
}

// The set of arguments for constructing a LinkAssociation resource.
type LinkAssociationArgs struct {
	// The ID of the device.
	DeviceId pulumi.StringInput
	// The ID of the global network.
	GlobalNetworkId pulumi.StringInput
	// The ID of the link.
	LinkId pulumi.StringInput
}

func (LinkAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAssociationArgs)(nil)).Elem()
}

type LinkAssociationInput interface {
	pulumi.Input

	ToLinkAssociationOutput() LinkAssociationOutput
	ToLinkAssociationOutputWithContext(ctx context.Context) LinkAssociationOutput
}

func (*LinkAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkAssociation)(nil)).Elem()
}

func (i *LinkAssociation) ToLinkAssociationOutput() LinkAssociationOutput {
	return i.ToLinkAssociationOutputWithContext(context.Background())
}

func (i *LinkAssociation) ToLinkAssociationOutputWithContext(ctx context.Context) LinkAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAssociationOutput)
}

// LinkAssociationArrayInput is an input type that accepts LinkAssociationArray and LinkAssociationArrayOutput values.
// You can construct a concrete instance of `LinkAssociationArrayInput` via:
//
//	LinkAssociationArray{ LinkAssociationArgs{...} }
type LinkAssociationArrayInput interface {
	pulumi.Input

	ToLinkAssociationArrayOutput() LinkAssociationArrayOutput
	ToLinkAssociationArrayOutputWithContext(context.Context) LinkAssociationArrayOutput
}

type LinkAssociationArray []LinkAssociationInput

func (LinkAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkAssociation)(nil)).Elem()
}

func (i LinkAssociationArray) ToLinkAssociationArrayOutput() LinkAssociationArrayOutput {
	return i.ToLinkAssociationArrayOutputWithContext(context.Background())
}

func (i LinkAssociationArray) ToLinkAssociationArrayOutputWithContext(ctx context.Context) LinkAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAssociationArrayOutput)
}

// LinkAssociationMapInput is an input type that accepts LinkAssociationMap and LinkAssociationMapOutput values.
// You can construct a concrete instance of `LinkAssociationMapInput` via:
//
//	LinkAssociationMap{ "key": LinkAssociationArgs{...} }
type LinkAssociationMapInput interface {
	pulumi.Input

	ToLinkAssociationMapOutput() LinkAssociationMapOutput
	ToLinkAssociationMapOutputWithContext(context.Context) LinkAssociationMapOutput
}

type LinkAssociationMap map[string]LinkAssociationInput

func (LinkAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkAssociation)(nil)).Elem()
}

func (i LinkAssociationMap) ToLinkAssociationMapOutput() LinkAssociationMapOutput {
	return i.ToLinkAssociationMapOutputWithContext(context.Background())
}

func (i LinkAssociationMap) ToLinkAssociationMapOutputWithContext(ctx context.Context) LinkAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAssociationMapOutput)
}

type LinkAssociationOutput struct{ *pulumi.OutputState }

func (LinkAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkAssociation)(nil)).Elem()
}

func (o LinkAssociationOutput) ToLinkAssociationOutput() LinkAssociationOutput {
	return o
}

func (o LinkAssociationOutput) ToLinkAssociationOutputWithContext(ctx context.Context) LinkAssociationOutput {
	return o
}

// The ID of the device.
func (o LinkAssociationOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkAssociation) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// The ID of the global network.
func (o LinkAssociationOutput) GlobalNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkAssociation) pulumi.StringOutput { return v.GlobalNetworkId }).(pulumi.StringOutput)
}

// The ID of the link.
func (o LinkAssociationOutput) LinkId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkAssociation) pulumi.StringOutput { return v.LinkId }).(pulumi.StringOutput)
}

type LinkAssociationArrayOutput struct{ *pulumi.OutputState }

func (LinkAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkAssociation)(nil)).Elem()
}

func (o LinkAssociationArrayOutput) ToLinkAssociationArrayOutput() LinkAssociationArrayOutput {
	return o
}

func (o LinkAssociationArrayOutput) ToLinkAssociationArrayOutputWithContext(ctx context.Context) LinkAssociationArrayOutput {
	return o
}

func (o LinkAssociationArrayOutput) Index(i pulumi.IntInput) LinkAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LinkAssociation {
		return vs[0].([]*LinkAssociation)[vs[1].(int)]
	}).(LinkAssociationOutput)
}

type LinkAssociationMapOutput struct{ *pulumi.OutputState }

func (LinkAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkAssociation)(nil)).Elem()
}

func (o LinkAssociationMapOutput) ToLinkAssociationMapOutput() LinkAssociationMapOutput {
	return o
}

func (o LinkAssociationMapOutput) ToLinkAssociationMapOutputWithContext(ctx context.Context) LinkAssociationMapOutput {
	return o
}

func (o LinkAssociationMapOutput) MapIndex(k pulumi.StringInput) LinkAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LinkAssociation {
		return vs[0].(map[string]*LinkAssociation)[vs[1].(string)]
	}).(LinkAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkAssociationInput)(nil)).Elem(), &LinkAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkAssociationArrayInput)(nil)).Elem(), LinkAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkAssociationMapInput)(nil)).Elem(), LinkAssociationMap{})
	pulumi.RegisterOutputType(LinkAssociationOutput{})
	pulumi.RegisterOutputType(LinkAssociationArrayOutput{})
	pulumi.RegisterOutputType(LinkAssociationMapOutput{})
}
