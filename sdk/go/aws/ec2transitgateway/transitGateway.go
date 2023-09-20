// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages an EC2 Transit Gateway.
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
//			_, err := ec2transitgateway.NewTransitGateway(ctx, "example", &ec2transitgateway.TransitGatewayArgs{
//				Description: pulumi.String("example"),
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
// Using `pulumi import`, import `aws_ec2_transit_gateway` using the EC2 Transit Gateway identifier. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2transitgateway/transitGateway:TransitGateway example tgw-12345678
//
// ```
type TransitGateway struct {
	pulumi.CustomResourceState

	// Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
	//
	// > **NOTE:** Modifying `amazonSideAsn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazonSideAsn`.
	AmazonSideAsn pulumi.IntPtrOutput `pulumi:"amazonSideAsn"`
	// EC2 Transit Gateway Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Identifier of the default association route table
	AssociationDefaultRouteTableId pulumi.StringOutput `pulumi:"associationDefaultRouteTableId"`
	// Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
	AutoAcceptSharedAttachments pulumi.StringPtrOutput `pulumi:"autoAcceptSharedAttachments"`
	// Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTableAssociation pulumi.StringPtrOutput `pulumi:"defaultRouteTableAssociation"`
	// Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTablePropagation pulumi.StringPtrOutput `pulumi:"defaultRouteTablePropagation"`
	// Description of the EC2 Transit Gateway.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport pulumi.StringPtrOutput `pulumi:"dnsSupport"`
	// Whether Multicast support is enabled. Required to use `ec2TransitGatewayMulticastDomain`. Valid values: `disable`, `enable`. Default value: `disable`.
	MulticastSupport pulumi.StringPtrOutput `pulumi:"multicastSupport"`
	// Identifier of the AWS account that owns the EC2 Transit Gateway
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Identifier of the default propagation route table
	PropagationDefaultRouteTableId pulumi.StringOutput `pulumi:"propagationDefaultRouteTableId"`
	// Key-value tags for the EC2 Transit Gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
	TransitGatewayCidrBlocks pulumi.StringArrayOutput `pulumi:"transitGatewayCidrBlocks"`
	// Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	VpnEcmpSupport pulumi.StringPtrOutput `pulumi:"vpnEcmpSupport"`
}

// NewTransitGateway registers a new resource with the given unique name, arguments, and options.
func NewTransitGateway(ctx *pulumi.Context,
	name string, args *TransitGatewayArgs, opts ...pulumi.ResourceOption) (*TransitGateway, error) {
	if args == nil {
		args = &TransitGatewayArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitGateway
	err := ctx.RegisterResource("aws:ec2transitgateway/transitGateway:TransitGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGateway gets an existing TransitGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayState, opts ...pulumi.ResourceOption) (*TransitGateway, error) {
	var resource TransitGateway
	err := ctx.ReadResource("aws:ec2transitgateway/transitGateway:TransitGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGateway resources.
type transitGatewayState struct {
	// Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
	//
	// > **NOTE:** Modifying `amazonSideAsn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazonSideAsn`.
	AmazonSideAsn *int `pulumi:"amazonSideAsn"`
	// EC2 Transit Gateway Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// Identifier of the default association route table
	AssociationDefaultRouteTableId *string `pulumi:"associationDefaultRouteTableId"`
	// Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
	AutoAcceptSharedAttachments *string `pulumi:"autoAcceptSharedAttachments"`
	// Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTableAssociation *string `pulumi:"defaultRouteTableAssociation"`
	// Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTablePropagation *string `pulumi:"defaultRouteTablePropagation"`
	// Description of the EC2 Transit Gateway.
	Description *string `pulumi:"description"`
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport *string `pulumi:"dnsSupport"`
	// Whether Multicast support is enabled. Required to use `ec2TransitGatewayMulticastDomain`. Valid values: `disable`, `enable`. Default value: `disable`.
	MulticastSupport *string `pulumi:"multicastSupport"`
	// Identifier of the AWS account that owns the EC2 Transit Gateway
	OwnerId *string `pulumi:"ownerId"`
	// Identifier of the default propagation route table
	PropagationDefaultRouteTableId *string `pulumi:"propagationDefaultRouteTableId"`
	// Key-value tags for the EC2 Transit Gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
	TransitGatewayCidrBlocks []string `pulumi:"transitGatewayCidrBlocks"`
	// Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	VpnEcmpSupport *string `pulumi:"vpnEcmpSupport"`
}

type TransitGatewayState struct {
	// Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
	//
	// > **NOTE:** Modifying `amazonSideAsn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazonSideAsn`.
	AmazonSideAsn pulumi.IntPtrInput
	// EC2 Transit Gateway Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// Identifier of the default association route table
	AssociationDefaultRouteTableId pulumi.StringPtrInput
	// Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
	AutoAcceptSharedAttachments pulumi.StringPtrInput
	// Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTableAssociation pulumi.StringPtrInput
	// Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTablePropagation pulumi.StringPtrInput
	// Description of the EC2 Transit Gateway.
	Description pulumi.StringPtrInput
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport pulumi.StringPtrInput
	// Whether Multicast support is enabled. Required to use `ec2TransitGatewayMulticastDomain`. Valid values: `disable`, `enable`. Default value: `disable`.
	MulticastSupport pulumi.StringPtrInput
	// Identifier of the AWS account that owns the EC2 Transit Gateway
	OwnerId pulumi.StringPtrInput
	// Identifier of the default propagation route table
	PropagationDefaultRouteTableId pulumi.StringPtrInput
	// Key-value tags for the EC2 Transit Gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
	TransitGatewayCidrBlocks pulumi.StringArrayInput
	// Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	VpnEcmpSupport pulumi.StringPtrInput
}

func (TransitGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayState)(nil)).Elem()
}

type transitGatewayArgs struct {
	// Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
	//
	// > **NOTE:** Modifying `amazonSideAsn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazonSideAsn`.
	AmazonSideAsn *int `pulumi:"amazonSideAsn"`
	// Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
	AutoAcceptSharedAttachments *string `pulumi:"autoAcceptSharedAttachments"`
	// Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTableAssociation *string `pulumi:"defaultRouteTableAssociation"`
	// Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTablePropagation *string `pulumi:"defaultRouteTablePropagation"`
	// Description of the EC2 Transit Gateway.
	Description *string `pulumi:"description"`
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport *string `pulumi:"dnsSupport"`
	// Whether Multicast support is enabled. Required to use `ec2TransitGatewayMulticastDomain`. Valid values: `disable`, `enable`. Default value: `disable`.
	MulticastSupport *string `pulumi:"multicastSupport"`
	// Key-value tags for the EC2 Transit Gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
	TransitGatewayCidrBlocks []string `pulumi:"transitGatewayCidrBlocks"`
	// Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	VpnEcmpSupport *string `pulumi:"vpnEcmpSupport"`
}

// The set of arguments for constructing a TransitGateway resource.
type TransitGatewayArgs struct {
	// Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
	//
	// > **NOTE:** Modifying `amazonSideAsn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazonSideAsn`.
	AmazonSideAsn pulumi.IntPtrInput
	// Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
	AutoAcceptSharedAttachments pulumi.StringPtrInput
	// Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTableAssociation pulumi.StringPtrInput
	// Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
	DefaultRouteTablePropagation pulumi.StringPtrInput
	// Description of the EC2 Transit Gateway.
	Description pulumi.StringPtrInput
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport pulumi.StringPtrInput
	// Whether Multicast support is enabled. Required to use `ec2TransitGatewayMulticastDomain`. Valid values: `disable`, `enable`. Default value: `disable`.
	MulticastSupport pulumi.StringPtrInput
	// Key-value tags for the EC2 Transit Gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
	TransitGatewayCidrBlocks pulumi.StringArrayInput
	// Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	VpnEcmpSupport pulumi.StringPtrInput
}

func (TransitGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayArgs)(nil)).Elem()
}

type TransitGatewayInput interface {
	pulumi.Input

	ToTransitGatewayOutput() TransitGatewayOutput
	ToTransitGatewayOutputWithContext(ctx context.Context) TransitGatewayOutput
}

func (*TransitGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGateway)(nil)).Elem()
}

func (i *TransitGateway) ToTransitGatewayOutput() TransitGatewayOutput {
	return i.ToTransitGatewayOutputWithContext(context.Background())
}

func (i *TransitGateway) ToTransitGatewayOutputWithContext(ctx context.Context) TransitGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayOutput)
}

func (i *TransitGateway) ToOutput(ctx context.Context) pulumix.Output[*TransitGateway] {
	return pulumix.Output[*TransitGateway]{
		OutputState: i.ToTransitGatewayOutputWithContext(ctx).OutputState,
	}
}

// TransitGatewayArrayInput is an input type that accepts TransitGatewayArray and TransitGatewayArrayOutput values.
// You can construct a concrete instance of `TransitGatewayArrayInput` via:
//
//	TransitGatewayArray{ TransitGatewayArgs{...} }
type TransitGatewayArrayInput interface {
	pulumi.Input

	ToTransitGatewayArrayOutput() TransitGatewayArrayOutput
	ToTransitGatewayArrayOutputWithContext(context.Context) TransitGatewayArrayOutput
}

type TransitGatewayArray []TransitGatewayInput

func (TransitGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitGateway)(nil)).Elem()
}

func (i TransitGatewayArray) ToTransitGatewayArrayOutput() TransitGatewayArrayOutput {
	return i.ToTransitGatewayArrayOutputWithContext(context.Background())
}

func (i TransitGatewayArray) ToTransitGatewayArrayOutputWithContext(ctx context.Context) TransitGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayArrayOutput)
}

func (i TransitGatewayArray) ToOutput(ctx context.Context) pulumix.Output[[]*TransitGateway] {
	return pulumix.Output[[]*TransitGateway]{
		OutputState: i.ToTransitGatewayArrayOutputWithContext(ctx).OutputState,
	}
}

// TransitGatewayMapInput is an input type that accepts TransitGatewayMap and TransitGatewayMapOutput values.
// You can construct a concrete instance of `TransitGatewayMapInput` via:
//
//	TransitGatewayMap{ "key": TransitGatewayArgs{...} }
type TransitGatewayMapInput interface {
	pulumi.Input

	ToTransitGatewayMapOutput() TransitGatewayMapOutput
	ToTransitGatewayMapOutputWithContext(context.Context) TransitGatewayMapOutput
}

type TransitGatewayMap map[string]TransitGatewayInput

func (TransitGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitGateway)(nil)).Elem()
}

func (i TransitGatewayMap) ToTransitGatewayMapOutput() TransitGatewayMapOutput {
	return i.ToTransitGatewayMapOutputWithContext(context.Background())
}

func (i TransitGatewayMap) ToTransitGatewayMapOutputWithContext(ctx context.Context) TransitGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayMapOutput)
}

func (i TransitGatewayMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TransitGateway] {
	return pulumix.Output[map[string]*TransitGateway]{
		OutputState: i.ToTransitGatewayMapOutputWithContext(ctx).OutputState,
	}
}

type TransitGatewayOutput struct{ *pulumi.OutputState }

func (TransitGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGateway)(nil)).Elem()
}

func (o TransitGatewayOutput) ToTransitGatewayOutput() TransitGatewayOutput {
	return o
}

func (o TransitGatewayOutput) ToTransitGatewayOutputWithContext(ctx context.Context) TransitGatewayOutput {
	return o
}

func (o TransitGatewayOutput) ToOutput(ctx context.Context) pulumix.Output[*TransitGateway] {
	return pulumix.Output[*TransitGateway]{
		OutputState: o.OutputState,
	}
}

// Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
//
// > **NOTE:** Modifying `amazonSideAsn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazonSideAsn`.
func (o TransitGatewayOutput) AmazonSideAsn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.IntPtrOutput { return v.AmazonSideAsn }).(pulumi.IntPtrOutput)
}

// EC2 Transit Gateway Amazon Resource Name (ARN)
func (o TransitGatewayOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Identifier of the default association route table
func (o TransitGatewayOutput) AssociationDefaultRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringOutput { return v.AssociationDefaultRouteTableId }).(pulumi.StringOutput)
}

// Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
func (o TransitGatewayOutput) AutoAcceptSharedAttachments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringPtrOutput { return v.AutoAcceptSharedAttachments }).(pulumi.StringPtrOutput)
}

// Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
func (o TransitGatewayOutput) DefaultRouteTableAssociation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringPtrOutput { return v.DefaultRouteTableAssociation }).(pulumi.StringPtrOutput)
}

// Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
func (o TransitGatewayOutput) DefaultRouteTablePropagation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringPtrOutput { return v.DefaultRouteTablePropagation }).(pulumi.StringPtrOutput)
}

// Description of the EC2 Transit Gateway.
func (o TransitGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
func (o TransitGatewayOutput) DnsSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringPtrOutput { return v.DnsSupport }).(pulumi.StringPtrOutput)
}

// Whether Multicast support is enabled. Required to use `ec2TransitGatewayMulticastDomain`. Valid values: `disable`, `enable`. Default value: `disable`.
func (o TransitGatewayOutput) MulticastSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringPtrOutput { return v.MulticastSupport }).(pulumi.StringPtrOutput)
}

// Identifier of the AWS account that owns the EC2 Transit Gateway
func (o TransitGatewayOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Identifier of the default propagation route table
func (o TransitGatewayOutput) PropagationDefaultRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringOutput { return v.PropagationDefaultRouteTableId }).(pulumi.StringOutput)
}

// Key-value tags for the EC2 Transit Gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TransitGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o TransitGatewayOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
func (o TransitGatewayOutput) TransitGatewayCidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringArrayOutput { return v.TransitGatewayCidrBlocks }).(pulumi.StringArrayOutput)
}

// Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
func (o TransitGatewayOutput) VpnEcmpSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitGateway) pulumi.StringPtrOutput { return v.VpnEcmpSupport }).(pulumi.StringPtrOutput)
}

type TransitGatewayArrayOutput struct{ *pulumi.OutputState }

func (TransitGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitGateway)(nil)).Elem()
}

func (o TransitGatewayArrayOutput) ToTransitGatewayArrayOutput() TransitGatewayArrayOutput {
	return o
}

func (o TransitGatewayArrayOutput) ToTransitGatewayArrayOutputWithContext(ctx context.Context) TransitGatewayArrayOutput {
	return o
}

func (o TransitGatewayArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TransitGateway] {
	return pulumix.Output[[]*TransitGateway]{
		OutputState: o.OutputState,
	}
}

func (o TransitGatewayArrayOutput) Index(i pulumi.IntInput) TransitGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitGateway {
		return vs[0].([]*TransitGateway)[vs[1].(int)]
	}).(TransitGatewayOutput)
}

type TransitGatewayMapOutput struct{ *pulumi.OutputState }

func (TransitGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitGateway)(nil)).Elem()
}

func (o TransitGatewayMapOutput) ToTransitGatewayMapOutput() TransitGatewayMapOutput {
	return o
}

func (o TransitGatewayMapOutput) ToTransitGatewayMapOutputWithContext(ctx context.Context) TransitGatewayMapOutput {
	return o
}

func (o TransitGatewayMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TransitGateway] {
	return pulumix.Output[map[string]*TransitGateway]{
		OutputState: o.OutputState,
	}
}

func (o TransitGatewayMapOutput) MapIndex(k pulumi.StringInput) TransitGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitGateway {
		return vs[0].(map[string]*TransitGateway)[vs[1].(string)]
	}).(TransitGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayInput)(nil)).Elem(), &TransitGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayArrayInput)(nil)).Elem(), TransitGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayMapInput)(nil)).Elem(), TransitGatewayMap{})
	pulumi.RegisterOutputType(TransitGatewayOutput{})
	pulumi.RegisterOutputType(TransitGatewayArrayOutput{})
	pulumi.RegisterOutputType(TransitGatewayMapOutput{})
}
