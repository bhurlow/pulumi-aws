// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator listener.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/globalaccelerator"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleAccelerator, err := globalaccelerator.NewAccelerator(ctx, "exampleAccelerator", &globalaccelerator.AcceleratorArgs{
//				IpAddressType: pulumi.String("IPV4"),
//				Enabled:       pulumi.Bool(true),
//				Attributes: &globalaccelerator.AcceleratorAttributesArgs{
//					FlowLogsEnabled:  pulumi.Bool(true),
//					FlowLogsS3Bucket: pulumi.String("example-bucket"),
//					FlowLogsS3Prefix: pulumi.String("flow-logs/"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = globalaccelerator.NewListener(ctx, "exampleListener", &globalaccelerator.ListenerArgs{
//				AcceleratorArn: exampleAccelerator.ID(),
//				ClientAffinity: pulumi.String("SOURCE_IP"),
//				Protocol:       pulumi.String("TCP"),
//				PortRanges: globalaccelerator.ListenerPortRangeArray{
//					&globalaccelerator.ListenerPortRangeArgs{
//						FromPort: pulumi.Int(80),
//						ToPort:   pulumi.Int(80),
//					},
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
//	to = aws_globalaccelerator_listener.example
//
//	id = "arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx" } Using `pulumi import`, import Global Accelerator listeners using the `id`. For exampleconsole % pulumi import aws_globalaccelerator_listener.example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
type Listener struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn pulumi.StringOutput `pulumi:"acceleratorArn"`
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity pulumi.StringPtrOutput `pulumi:"clientAffinity"`
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges ListenerPortRangeArrayOutput `pulumi:"portRanges"`
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorArn == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorArn'")
	}
	if args.PortRanges == nil {
		return nil, errors.New("invalid value for required argument 'PortRanges'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Listener
	err := ctx.RegisterResource("aws:globalaccelerator/listener:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("aws:globalaccelerator/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn *string `pulumi:"acceleratorArn"`
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity *string `pulumi:"clientAffinity"`
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges []ListenerPortRange `pulumi:"portRanges"`
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol *string `pulumi:"protocol"`
}

type ListenerState struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn pulumi.StringPtrInput
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity pulumi.StringPtrInput
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges ListenerPortRangeArrayInput
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn string `pulumi:"acceleratorArn"`
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity *string `pulumi:"clientAffinity"`
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges []ListenerPortRange `pulumi:"portRanges"`
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol string `pulumi:"protocol"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn pulumi.StringInput
	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity pulumi.StringPtrInput
	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges ListenerPortRangeArrayInput
	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol pulumi.StringInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

// ListenerArrayInput is an input type that accepts ListenerArray and ListenerArrayOutput values.
// You can construct a concrete instance of `ListenerArrayInput` via:
//
//	ListenerArray{ ListenerArgs{...} }
type ListenerArrayInput interface {
	pulumi.Input

	ToListenerArrayOutput() ListenerArrayOutput
	ToListenerArrayOutputWithContext(context.Context) ListenerArrayOutput
}

type ListenerArray []ListenerInput

func (ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (i ListenerArray) ToListenerArrayOutput() ListenerArrayOutput {
	return i.ToListenerArrayOutputWithContext(context.Background())
}

func (i ListenerArray) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerArrayOutput)
}

// ListenerMapInput is an input type that accepts ListenerMap and ListenerMapOutput values.
// You can construct a concrete instance of `ListenerMapInput` via:
//
//	ListenerMap{ "key": ListenerArgs{...} }
type ListenerMapInput interface {
	pulumi.Input

	ToListenerMapOutput() ListenerMapOutput
	ToListenerMapOutputWithContext(context.Context) ListenerMapOutput
}

type ListenerMap map[string]ListenerInput

func (ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (i ListenerMap) ToListenerMapOutput() ListenerMapOutput {
	return i.ToListenerMapOutputWithContext(context.Background())
}

func (i ListenerMap) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerMapOutput)
}

type ListenerOutput struct{ *pulumi.OutputState }

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

// The Amazon Resource Name (ARN) of your accelerator.
func (o ListenerOutput) AcceleratorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.AcceleratorArn }).(pulumi.StringOutput)
}

// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
func (o ListenerOutput) ClientAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.ClientAffinity }).(pulumi.StringPtrOutput)
}

// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
func (o ListenerOutput) PortRanges() ListenerPortRangeArrayOutput {
	return o.ApplyT(func(v *Listener) ListenerPortRangeArrayOutput { return v.PortRanges }).(ListenerPortRangeArrayOutput)
}

// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
func (o ListenerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

type ListenerArrayOutput struct{ *pulumi.OutputState }

func (ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (o ListenerArrayOutput) ToListenerArrayOutput() ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) Index(i pulumi.IntInput) ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].([]*Listener)[vs[1].(int)]
	}).(ListenerOutput)
}

type ListenerMapOutput struct{ *pulumi.OutputState }

func (ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (o ListenerMapOutput) ToListenerMapOutput() ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) MapIndex(k pulumi.StringInput) ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].(map[string]*Listener)[vs[1].(string)]
	}).(ListenerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerArrayInput)(nil)).Elem(), ListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerMapInput)(nil)).Elem(), ListenerMap{})
	pulumi.RegisterOutputType(ListenerOutput{})
	pulumi.RegisterOutputType(ListenerArrayOutput{})
	pulumi.RegisterOutputType(ListenerMapOutput{})
}
