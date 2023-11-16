// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2clientvpn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides additional routes for AWS Client VPN endpoints. For more information on usage, please see the
// [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2clientvpn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleEndpoint, err := ec2clientvpn.NewEndpoint(ctx, "exampleEndpoint", &ec2clientvpn.EndpointArgs{
//				Description:          pulumi.String("Example Client VPN endpoint"),
//				ServerCertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
//				ClientCidrBlock:      pulumi.String("10.0.0.0/16"),
//				AuthenticationOptions: ec2clientvpn.EndpointAuthenticationOptionArray{
//					&ec2clientvpn.EndpointAuthenticationOptionArgs{
//						Type:                    pulumi.String("certificate-authentication"),
//						RootCertificateChainArn: pulumi.Any(aws_acm_certificate.Example.Arn),
//					},
//				},
//				ConnectionLogOptions: &ec2clientvpn.EndpointConnectionLogOptionsArgs{
//					Enabled: pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleNetworkAssociation, err := ec2clientvpn.NewNetworkAssociation(ctx, "exampleNetworkAssociation", &ec2clientvpn.NetworkAssociationArgs{
//				ClientVpnEndpointId: exampleEndpoint.ID(),
//				SubnetId:            pulumi.Any(aws_subnet.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2clientvpn.NewRoute(ctx, "exampleRoute", &ec2clientvpn.RouteArgs{
//				ClientVpnEndpointId:  exampleEndpoint.ID(),
//				DestinationCidrBlock: pulumi.String("0.0.0.0/0"),
//				TargetVpcSubnetId:    exampleNetworkAssociation.SubnetId,
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
// Using `pulumi import`, import AWS Client VPN routes using the endpoint ID, target subnet ID, and destination CIDR block. All values are separated by a `,`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2clientvpn/route:Route example cvpn-endpoint-1234567890abcdef,subnet-9876543210fedcba,10.1.0.0/24
//
// ```
type Route struct {
	pulumi.CustomResourceState

	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId pulumi.StringOutput `pulumi:"clientVpnEndpointId"`
	// A brief description of the route.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IPv4 address range, in CIDR notation, of the route destination.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// Indicates how the Client VPN route was added. Will be `add-route` for routes created by this resource.
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
	TargetVpcSubnetId pulumi.StringOutput `pulumi:"targetVpcSubnetId"`
	// The type of the route.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientVpnEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'ClientVpnEndpointId'")
	}
	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.TargetVpcSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetVpcSubnetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("aws:ec2clientvpn/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws:ec2clientvpn/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string `pulumi:"clientVpnEndpointId"`
	// A brief description of the route.
	Description *string `pulumi:"description"`
	// The IPv4 address range, in CIDR notation, of the route destination.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// Indicates how the Client VPN route was added. Will be `add-route` for routes created by this resource.
	Origin *string `pulumi:"origin"`
	// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
	TargetVpcSubnetId *string `pulumi:"targetVpcSubnetId"`
	// The type of the route.
	Type *string `pulumi:"type"`
}

type RouteState struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId pulumi.StringPtrInput
	// A brief description of the route.
	Description pulumi.StringPtrInput
	// The IPv4 address range, in CIDR notation, of the route destination.
	DestinationCidrBlock pulumi.StringPtrInput
	// Indicates how the Client VPN route was added. Will be `add-route` for routes created by this resource.
	Origin pulumi.StringPtrInput
	// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
	TargetVpcSubnetId pulumi.StringPtrInput
	// The type of the route.
	Type pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId string `pulumi:"clientVpnEndpointId"`
	// A brief description of the route.
	Description *string `pulumi:"description"`
	// The IPv4 address range, in CIDR notation, of the route destination.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
	TargetVpcSubnetId string `pulumi:"targetVpcSubnetId"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId pulumi.StringInput
	// A brief description of the route.
	Description pulumi.StringPtrInput
	// The IPv4 address range, in CIDR notation, of the route destination.
	DestinationCidrBlock pulumi.StringInput
	// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
	TargetVpcSubnetId pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//	RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//	RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

// The ID of the Client VPN endpoint.
func (o RouteOutput) ClientVpnEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ClientVpnEndpointId }).(pulumi.StringOutput)
}

// A brief description of the route.
func (o RouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The IPv4 address range, in CIDR notation, of the route destination.
func (o RouteOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// Indicates how the Client VPN route was added. Will be `add-route` for routes created by this resource.
func (o RouteOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
func (o RouteOutput) TargetVpcSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.TargetVpcSubnetId }).(pulumi.StringOutput)
}

// The type of the route.
func (o RouteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Route {
		return vs[0].([]*Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Route {
		return vs[0].(map[string]*Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteArrayInput)(nil)).Elem(), RouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapInput)(nil)).Elem(), RouteMap{})
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
