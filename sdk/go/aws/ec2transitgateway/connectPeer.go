// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EC2 Transit Gateway Connect Peer.
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
//			exampleConnect, err := ec2transitgateway.NewConnect(ctx, "exampleConnect", &ec2transitgateway.ConnectArgs{
//				TransportAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
//				TransitGatewayId:      pulumi.Any(aws_ec2_transit_gateway.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2transitgateway.NewConnectPeer(ctx, "exampleConnectPeer", &ec2transitgateway.ConnectPeerArgs{
//				PeerAddress: pulumi.String("10.1.2.3"),
//				InsideCidrBlocks: pulumi.StringArray{
//					pulumi.String("169.254.100.0/29"),
//				},
//				TransitGatewayAttachmentId: exampleConnect.ID(),
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
// `aws_ec2_transit_gateway_connect_peer` can be imported by using the EC2 Transit Gateway Connect Peer identifier, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2transitgateway/connectPeer:ConnectPeer example tgw-connect-peer-12345678
//
// ```
type ConnectPeer struct {
	pulumi.CustomResourceState

	// EC2 Transit Gateway Connect Peer ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	BgpAsn pulumi.StringOutput `pulumi:"bgpAsn"`
	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	InsideCidrBlocks pulumi.StringArrayOutput `pulumi:"insideCidrBlocks"`
	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
	PeerAddress pulumi.StringOutput `pulumi:"peerAddress"`
	// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	TransitGatewayAddress pulumi.StringOutput `pulumi:"transitGatewayAddress"`
	// The Transit Gateway Connect
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`
}

// NewConnectPeer registers a new resource with the given unique name, arguments, and options.
func NewConnectPeer(ctx *pulumi.Context,
	name string, args *ConnectPeerArgs, opts ...pulumi.ResourceOption) (*ConnectPeer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InsideCidrBlocks == nil {
		return nil, errors.New("invalid value for required argument 'InsideCidrBlocks'")
	}
	if args.PeerAddress == nil {
		return nil, errors.New("invalid value for required argument 'PeerAddress'")
	}
	if args.TransitGatewayAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayAttachmentId'")
	}
	var resource ConnectPeer
	err := ctx.RegisterResource("aws:ec2transitgateway/connectPeer:ConnectPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectPeer gets an existing ConnectPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectPeerState, opts ...pulumi.ResourceOption) (*ConnectPeer, error) {
	var resource ConnectPeer
	err := ctx.ReadResource("aws:ec2transitgateway/connectPeer:ConnectPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectPeer resources.
type connectPeerState struct {
	// EC2 Transit Gateway Connect Peer ARN
	Arn *string `pulumi:"arn"`
	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	BgpAsn *string `pulumi:"bgpAsn"`
	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	InsideCidrBlocks []string `pulumi:"insideCidrBlocks"`
	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
	PeerAddress *string `pulumi:"peerAddress"`
	// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	TransitGatewayAddress *string `pulumi:"transitGatewayAddress"`
	// The Transit Gateway Connect
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
}

type ConnectPeerState struct {
	// EC2 Transit Gateway Connect Peer ARN
	Arn pulumi.StringPtrInput
	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	BgpAsn pulumi.StringPtrInput
	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	InsideCidrBlocks pulumi.StringArrayInput
	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
	PeerAddress pulumi.StringPtrInput
	// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	TransitGatewayAddress pulumi.StringPtrInput
	// The Transit Gateway Connect
	TransitGatewayAttachmentId pulumi.StringPtrInput
}

func (ConnectPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectPeerState)(nil)).Elem()
}

type connectPeerArgs struct {
	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	BgpAsn *string `pulumi:"bgpAsn"`
	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	InsideCidrBlocks []string `pulumi:"insideCidrBlocks"`
	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
	PeerAddress string `pulumi:"peerAddress"`
	// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	TransitGatewayAddress *string `pulumi:"transitGatewayAddress"`
	// The Transit Gateway Connect
	TransitGatewayAttachmentId string `pulumi:"transitGatewayAttachmentId"`
}

// The set of arguments for constructing a ConnectPeer resource.
type ConnectPeerArgs struct {
	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	BgpAsn pulumi.StringPtrInput
	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	InsideCidrBlocks pulumi.StringArrayInput
	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
	PeerAddress pulumi.StringInput
	// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	TransitGatewayAddress pulumi.StringPtrInput
	// The Transit Gateway Connect
	TransitGatewayAttachmentId pulumi.StringInput
}

func (ConnectPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectPeerArgs)(nil)).Elem()
}

type ConnectPeerInput interface {
	pulumi.Input

	ToConnectPeerOutput() ConnectPeerOutput
	ToConnectPeerOutputWithContext(ctx context.Context) ConnectPeerOutput
}

func (*ConnectPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectPeer)(nil)).Elem()
}

func (i *ConnectPeer) ToConnectPeerOutput() ConnectPeerOutput {
	return i.ToConnectPeerOutputWithContext(context.Background())
}

func (i *ConnectPeer) ToConnectPeerOutputWithContext(ctx context.Context) ConnectPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectPeerOutput)
}

// ConnectPeerArrayInput is an input type that accepts ConnectPeerArray and ConnectPeerArrayOutput values.
// You can construct a concrete instance of `ConnectPeerArrayInput` via:
//
//	ConnectPeerArray{ ConnectPeerArgs{...} }
type ConnectPeerArrayInput interface {
	pulumi.Input

	ToConnectPeerArrayOutput() ConnectPeerArrayOutput
	ToConnectPeerArrayOutputWithContext(context.Context) ConnectPeerArrayOutput
}

type ConnectPeerArray []ConnectPeerInput

func (ConnectPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectPeer)(nil)).Elem()
}

func (i ConnectPeerArray) ToConnectPeerArrayOutput() ConnectPeerArrayOutput {
	return i.ToConnectPeerArrayOutputWithContext(context.Background())
}

func (i ConnectPeerArray) ToConnectPeerArrayOutputWithContext(ctx context.Context) ConnectPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectPeerArrayOutput)
}

// ConnectPeerMapInput is an input type that accepts ConnectPeerMap and ConnectPeerMapOutput values.
// You can construct a concrete instance of `ConnectPeerMapInput` via:
//
//	ConnectPeerMap{ "key": ConnectPeerArgs{...} }
type ConnectPeerMapInput interface {
	pulumi.Input

	ToConnectPeerMapOutput() ConnectPeerMapOutput
	ToConnectPeerMapOutputWithContext(context.Context) ConnectPeerMapOutput
}

type ConnectPeerMap map[string]ConnectPeerInput

func (ConnectPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectPeer)(nil)).Elem()
}

func (i ConnectPeerMap) ToConnectPeerMapOutput() ConnectPeerMapOutput {
	return i.ToConnectPeerMapOutputWithContext(context.Background())
}

func (i ConnectPeerMap) ToConnectPeerMapOutputWithContext(ctx context.Context) ConnectPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectPeerMapOutput)
}

type ConnectPeerOutput struct{ *pulumi.OutputState }

func (ConnectPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectPeer)(nil)).Elem()
}

func (o ConnectPeerOutput) ToConnectPeerOutput() ConnectPeerOutput {
	return o
}

func (o ConnectPeerOutput) ToConnectPeerOutputWithContext(ctx context.Context) ConnectPeerOutput {
	return o
}

// EC2 Transit Gateway Connect Peer ARN
func (o ConnectPeerOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
func (o ConnectPeerOutput) BgpAsn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringOutput { return v.BgpAsn }).(pulumi.StringOutput)
}

// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
func (o ConnectPeerOutput) InsideCidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringArrayOutput { return v.InsideCidrBlocks }).(pulumi.StringArrayOutput)
}

// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transitGatewayAddress`
func (o ConnectPeerOutput) PeerAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringOutput { return v.PeerAddress }).(pulumi.StringOutput)
}

// Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ConnectPeerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ConnectPeerOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peerAddress`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
func (o ConnectPeerOutput) TransitGatewayAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringOutput { return v.TransitGatewayAddress }).(pulumi.StringOutput)
}

// The Transit Gateway Connect
func (o ConnectPeerOutput) TransitGatewayAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectPeer) pulumi.StringOutput { return v.TransitGatewayAttachmentId }).(pulumi.StringOutput)
}

type ConnectPeerArrayOutput struct{ *pulumi.OutputState }

func (ConnectPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectPeer)(nil)).Elem()
}

func (o ConnectPeerArrayOutput) ToConnectPeerArrayOutput() ConnectPeerArrayOutput {
	return o
}

func (o ConnectPeerArrayOutput) ToConnectPeerArrayOutputWithContext(ctx context.Context) ConnectPeerArrayOutput {
	return o
}

func (o ConnectPeerArrayOutput) Index(i pulumi.IntInput) ConnectPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConnectPeer {
		return vs[0].([]*ConnectPeer)[vs[1].(int)]
	}).(ConnectPeerOutput)
}

type ConnectPeerMapOutput struct{ *pulumi.OutputState }

func (ConnectPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectPeer)(nil)).Elem()
}

func (o ConnectPeerMapOutput) ToConnectPeerMapOutput() ConnectPeerMapOutput {
	return o
}

func (o ConnectPeerMapOutput) ToConnectPeerMapOutputWithContext(ctx context.Context) ConnectPeerMapOutput {
	return o
}

func (o ConnectPeerMapOutput) MapIndex(k pulumi.StringInput) ConnectPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConnectPeer {
		return vs[0].(map[string]*ConnectPeer)[vs[1].(string)]
	}).(ConnectPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectPeerInput)(nil)).Elem(), &ConnectPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectPeerArrayInput)(nil)).Elem(), ConnectPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectPeerMapInput)(nil)).Elem(), ConnectPeerMap{})
	pulumi.RegisterOutputType(ConnectPeerOutput{})
	pulumi.RegisterOutputType(ConnectPeerArrayOutput{})
	pulumi.RegisterOutputType(ConnectPeerMapOutput{})
}
