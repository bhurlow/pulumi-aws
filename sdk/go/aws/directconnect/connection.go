// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Connection of Direct Connect.
//
// ## Example Usage
// ### Create a connection
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := directconnect.NewConnection(ctx, "hoge", &directconnect.ConnectionArgs{
//				Bandwidth: pulumi.String("1Gbps"),
//				Location:  pulumi.String("EqDC2"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Request a MACsec-capable connection
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := directconnect.NewConnection(ctx, "example", &directconnect.ConnectionArgs{
//				Bandwidth:     pulumi.String("10Gbps"),
//				Location:      pulumi.String("EqDA2"),
//				RequestMacsec: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Configure encryption mode for MACsec-capable connections
//
// > **NOTE:** You can only specify the `encryptionMode` argument once the connection is in an `Available` state.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := directconnect.NewConnection(ctx, "example", &directconnect.ConnectionArgs{
//				Bandwidth:      pulumi.String("10Gbps"),
//				EncryptionMode: pulumi.String("must_encrypt"),
//				Location:       pulumi.String("EqDC2"),
//				RequestMacsec:  pulumi.Bool(true),
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
// Using `pulumi import`, import Direct Connect connections using the connection `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:directconnect/connection:Connection test_connection dxcon-ffre0ec3
//
// ```
type Connection struct {
	pulumi.CustomResourceState

	// The ARN of the connection.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice pulumi.StringOutput `pulumi:"awsDevice"`
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth pulumi.StringOutput `pulumi:"bandwidth"`
	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `noEncrypt`, `shouldEncrypt`, and `mustEncrypt`.
	EncryptionMode pulumi.StringOutput `pulumi:"encryptionMode"`
	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringOutput `pulumi:"hasLogicalRedundancy"`
	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable pulumi.BoolOutput `pulumi:"jumboFrameCapable"`
	// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringOutput `pulumi:"location"`
	// Boolean value indicating whether the connection supports MAC Security (MACsec).
	MacsecCapable pulumi.BoolOutput `pulumi:"macsecCapable"`
	// The name of the connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the AWS account that owns the connection.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName pulumi.StringOutput `pulumi:"partnerName"`
	// The MAC Security (MACsec) port link status of the connection.
	PortEncryptionStatus pulumi.StringOutput `pulumi:"portEncryptionStatus"`
	// The name of the service provider associated with the connection.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
	//
	// > **NOTE:** Changing the value of `requestMacsec` will cause the resource to be destroyed and re-created.
	RequestMacsec pulumi.BoolPtrOutput `pulumi:"requestMacsec"`
	// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
	SkipDestroy pulumi.BoolPtrOutput `pulumi:"skipDestroy"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The VLAN ID.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("aws:directconnect/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("aws:directconnect/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// The ARN of the connection.
	Arn *string `pulumi:"arn"`
	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `pulumi:"awsDevice"`
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth *string `pulumi:"bandwidth"`
	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `noEncrypt`, `shouldEncrypt`, and `mustEncrypt`.
	EncryptionMode *string `pulumi:"encryptionMode"`
	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy *string `pulumi:"hasLogicalRedundancy"`
	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable *bool `pulumi:"jumboFrameCapable"`
	// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location *string `pulumi:"location"`
	// Boolean value indicating whether the connection supports MAC Security (MACsec).
	MacsecCapable *bool `pulumi:"macsecCapable"`
	// The name of the connection.
	Name *string `pulumi:"name"`
	// The ID of the AWS account that owns the connection.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName *string `pulumi:"partnerName"`
	// The MAC Security (MACsec) port link status of the connection.
	PortEncryptionStatus *string `pulumi:"portEncryptionStatus"`
	// The name of the service provider associated with the connection.
	ProviderName *string `pulumi:"providerName"`
	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
	//
	// > **NOTE:** Changing the value of `requestMacsec` will cause the resource to be destroyed and re-created.
	RequestMacsec *bool `pulumi:"requestMacsec"`
	// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

type ConnectionState struct {
	// The ARN of the connection.
	Arn pulumi.StringPtrInput
	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice pulumi.StringPtrInput
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth pulumi.StringPtrInput
	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `noEncrypt`, `shouldEncrypt`, and `mustEncrypt`.
	EncryptionMode pulumi.StringPtrInput
	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringPtrInput
	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable pulumi.BoolPtrInput
	// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringPtrInput
	// Boolean value indicating whether the connection supports MAC Security (MACsec).
	MacsecCapable pulumi.BoolPtrInput
	// The name of the connection.
	Name pulumi.StringPtrInput
	// The ID of the AWS account that owns the connection.
	OwnerAccountId pulumi.StringPtrInput
	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName pulumi.StringPtrInput
	// The MAC Security (MACsec) port link status of the connection.
	PortEncryptionStatus pulumi.StringPtrInput
	// The name of the service provider associated with the connection.
	ProviderName pulumi.StringPtrInput
	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
	//
	// > **NOTE:** Changing the value of `requestMacsec` will cause the resource to be destroyed and re-created.
	RequestMacsec pulumi.BoolPtrInput
	// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
	SkipDestroy pulumi.BoolPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The VLAN ID.
	VlanId pulumi.IntPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth string `pulumi:"bandwidth"`
	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `noEncrypt`, `shouldEncrypt`, and `mustEncrypt`.
	EncryptionMode *string `pulumi:"encryptionMode"`
	// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location string `pulumi:"location"`
	// The name of the connection.
	Name *string `pulumi:"name"`
	// The name of the service provider associated with the connection.
	ProviderName *string `pulumi:"providerName"`
	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
	//
	// > **NOTE:** Changing the value of `requestMacsec` will cause the resource to be destroyed and re-created.
	RequestMacsec *bool `pulumi:"requestMacsec"`
	// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth pulumi.StringInput
	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `noEncrypt`, `shouldEncrypt`, and `mustEncrypt`.
	EncryptionMode pulumi.StringPtrInput
	// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringInput
	// The name of the connection.
	Name pulumi.StringPtrInput
	// The name of the service provider associated with the connection.
	ProviderName pulumi.StringPtrInput
	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
	//
	// > **NOTE:** Changing the value of `requestMacsec` will cause the resource to be destroyed and re-created.
	RequestMacsec pulumi.BoolPtrInput
	// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
	SkipDestroy pulumi.BoolPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i *Connection) ToOutput(ctx context.Context) pulumix.Output[*Connection] {
	return pulumix.Output[*Connection]{
		OutputState: i.ToConnectionOutputWithContext(ctx).OutputState,
	}
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//	ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

func (i ConnectionArray) ToOutput(ctx context.Context) pulumix.Output[[]*Connection] {
	return pulumix.Output[[]*Connection]{
		OutputState: i.ToConnectionArrayOutputWithContext(ctx).OutputState,
	}
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//	ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

func (i ConnectionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Connection] {
	return pulumix.Output[map[string]*Connection]{
		OutputState: i.ToConnectionMapOutputWithContext(ctx).OutputState,
	}
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*Connection] {
	return pulumix.Output[*Connection]{
		OutputState: o.OutputState,
	}
}

// The ARN of the connection.
func (o ConnectionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Direct Connect endpoint on which the physical connection terminates.
func (o ConnectionOutput) AwsDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.AwsDevice }).(pulumi.StringOutput)
}

// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
func (o ConnectionOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Bandwidth }).(pulumi.StringOutput)
}

// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `noEncrypt`, `shouldEncrypt`, and `mustEncrypt`.
func (o ConnectionOutput) EncryptionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.EncryptionMode }).(pulumi.StringOutput)
}

// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
func (o ConnectionOutput) HasLogicalRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.HasLogicalRedundancy }).(pulumi.StringOutput)
}

// Boolean value representing if jumbo frames have been enabled for this connection.
func (o ConnectionOutput) JumboFrameCapable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolOutput { return v.JumboFrameCapable }).(pulumi.BoolOutput)
}

// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
func (o ConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Boolean value indicating whether the connection supports MAC Security (MACsec).
func (o ConnectionOutput) MacsecCapable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolOutput { return v.MacsecCapable }).(pulumi.BoolOutput)
}

// The name of the connection.
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the AWS account that owns the connection.
func (o ConnectionOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// The name of the AWS Direct Connect service provider associated with the connection.
func (o ConnectionOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.PartnerName }).(pulumi.StringOutput)
}

// The MAC Security (MACsec) port link status of the connection.
func (o ConnectionOutput) PortEncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.PortEncryptionStatus }).(pulumi.StringOutput)
}

// The name of the service provider associated with the connection.
func (o ConnectionOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ProviderName }).(pulumi.StringOutput)
}

// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
//
// > **NOTE:** Changing the value of `requestMacsec` will cause the resource to be destroyed and re-created.
func (o ConnectionOutput) RequestMacsec() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolPtrOutput { return v.RequestMacsec }).(pulumi.BoolPtrOutput)
}

// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
func (o ConnectionOutput) SkipDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolPtrOutput { return v.SkipDestroy }).(pulumi.BoolPtrOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ConnectionOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The VLAN ID.
func (o ConnectionOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *Connection) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Connection] {
	return pulumix.Output[[]*Connection]{
		OutputState: o.OutputState,
	}
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].([]*Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Connection] {
	return pulumix.Output[map[string]*Connection]{
		OutputState: o.OutputState,
	}
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].(map[string]*Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
