// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chime

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enable Termination settings to control outbound calling from your SIP infrastructure.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chime"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultVoiceConnector, err := chime.NewVoiceConnector(ctx, "defaultVoiceConnector", &chime.VoiceConnectorArgs{
//				RequireEncryption: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = chime.NewVoiceConnectorTermination(ctx, "defaultVoiceConnectorTermination", &chime.VoiceConnectorTerminationArgs{
//				Disabled: pulumi.Bool(false),
//				CpsLimit: pulumi.Int(1),
//				CidrAllowLists: pulumi.StringArray{
//					pulumi.String("50.35.78.96/31"),
//				},
//				CallingRegions: pulumi.StringArray{
//					pulumi.String("US"),
//					pulumi.String("CA"),
//				},
//				VoiceConnectorId: defaultVoiceConnector.ID(),
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
//	to = aws_chime_voice_connector_termination.default
//
//	id = "abcdef1ghij2klmno3pqr4" } Using `pulumi import`, import Chime Voice Connector Termination using the `voice_connector_id`. For exampleconsole % pulumi import aws_chime_voice_connector_termination.default abcdef1ghij2klmno3pqr4
type VoiceConnectorTermination struct {
	pulumi.CustomResourceState

	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	CallingRegions pulumi.StringArrayOutput `pulumi:"callingRegions"`
	// The IP addresses allowed to make calls, in CIDR format.
	CidrAllowLists pulumi.StringArrayOutput `pulumi:"cidrAllowLists"`
	// The limit on calls per second. Max value based on account service quota. Default value of `1`.
	CpsLimit pulumi.IntPtrOutput `pulumi:"cpsLimit"`
	// The default caller ID phone number.
	DefaultPhoneNumber pulumi.StringPtrOutput `pulumi:"defaultPhoneNumber"`
	// When termination settings are disabled, outbound calls can not be made.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId pulumi.StringOutput `pulumi:"voiceConnectorId"`
}

// NewVoiceConnectorTermination registers a new resource with the given unique name, arguments, and options.
func NewVoiceConnectorTermination(ctx *pulumi.Context,
	name string, args *VoiceConnectorTerminationArgs, opts ...pulumi.ResourceOption) (*VoiceConnectorTermination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CallingRegions == nil {
		return nil, errors.New("invalid value for required argument 'CallingRegions'")
	}
	if args.CidrAllowLists == nil {
		return nil, errors.New("invalid value for required argument 'CidrAllowLists'")
	}
	if args.VoiceConnectorId == nil {
		return nil, errors.New("invalid value for required argument 'VoiceConnectorId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VoiceConnectorTermination
	err := ctx.RegisterResource("aws:chime/voiceConnectorTermination:VoiceConnectorTermination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVoiceConnectorTermination gets an existing VoiceConnectorTermination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVoiceConnectorTermination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VoiceConnectorTerminationState, opts ...pulumi.ResourceOption) (*VoiceConnectorTermination, error) {
	var resource VoiceConnectorTermination
	err := ctx.ReadResource("aws:chime/voiceConnectorTermination:VoiceConnectorTermination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VoiceConnectorTermination resources.
type voiceConnectorTerminationState struct {
	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	CallingRegions []string `pulumi:"callingRegions"`
	// The IP addresses allowed to make calls, in CIDR format.
	CidrAllowLists []string `pulumi:"cidrAllowLists"`
	// The limit on calls per second. Max value based on account service quota. Default value of `1`.
	CpsLimit *int `pulumi:"cpsLimit"`
	// The default caller ID phone number.
	DefaultPhoneNumber *string `pulumi:"defaultPhoneNumber"`
	// When termination settings are disabled, outbound calls can not be made.
	Disabled *bool `pulumi:"disabled"`
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string `pulumi:"voiceConnectorId"`
}

type VoiceConnectorTerminationState struct {
	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	CallingRegions pulumi.StringArrayInput
	// The IP addresses allowed to make calls, in CIDR format.
	CidrAllowLists pulumi.StringArrayInput
	// The limit on calls per second. Max value based on account service quota. Default value of `1`.
	CpsLimit pulumi.IntPtrInput
	// The default caller ID phone number.
	DefaultPhoneNumber pulumi.StringPtrInput
	// When termination settings are disabled, outbound calls can not be made.
	Disabled pulumi.BoolPtrInput
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId pulumi.StringPtrInput
}

func (VoiceConnectorTerminationState) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceConnectorTerminationState)(nil)).Elem()
}

type voiceConnectorTerminationArgs struct {
	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	CallingRegions []string `pulumi:"callingRegions"`
	// The IP addresses allowed to make calls, in CIDR format.
	CidrAllowLists []string `pulumi:"cidrAllowLists"`
	// The limit on calls per second. Max value based on account service quota. Default value of `1`.
	CpsLimit *int `pulumi:"cpsLimit"`
	// The default caller ID phone number.
	DefaultPhoneNumber *string `pulumi:"defaultPhoneNumber"`
	// When termination settings are disabled, outbound calls can not be made.
	Disabled *bool `pulumi:"disabled"`
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId string `pulumi:"voiceConnectorId"`
}

// The set of arguments for constructing a VoiceConnectorTermination resource.
type VoiceConnectorTerminationArgs struct {
	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	CallingRegions pulumi.StringArrayInput
	// The IP addresses allowed to make calls, in CIDR format.
	CidrAllowLists pulumi.StringArrayInput
	// The limit on calls per second. Max value based on account service quota. Default value of `1`.
	CpsLimit pulumi.IntPtrInput
	// The default caller ID phone number.
	DefaultPhoneNumber pulumi.StringPtrInput
	// When termination settings are disabled, outbound calls can not be made.
	Disabled pulumi.BoolPtrInput
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId pulumi.StringInput
}

func (VoiceConnectorTerminationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceConnectorTerminationArgs)(nil)).Elem()
}

type VoiceConnectorTerminationInput interface {
	pulumi.Input

	ToVoiceConnectorTerminationOutput() VoiceConnectorTerminationOutput
	ToVoiceConnectorTerminationOutputWithContext(ctx context.Context) VoiceConnectorTerminationOutput
}

func (*VoiceConnectorTermination) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceConnectorTermination)(nil)).Elem()
}

func (i *VoiceConnectorTermination) ToVoiceConnectorTerminationOutput() VoiceConnectorTerminationOutput {
	return i.ToVoiceConnectorTerminationOutputWithContext(context.Background())
}

func (i *VoiceConnectorTermination) ToVoiceConnectorTerminationOutputWithContext(ctx context.Context) VoiceConnectorTerminationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationOutput)
}

// VoiceConnectorTerminationArrayInput is an input type that accepts VoiceConnectorTerminationArray and VoiceConnectorTerminationArrayOutput values.
// You can construct a concrete instance of `VoiceConnectorTerminationArrayInput` via:
//
//	VoiceConnectorTerminationArray{ VoiceConnectorTerminationArgs{...} }
type VoiceConnectorTerminationArrayInput interface {
	pulumi.Input

	ToVoiceConnectorTerminationArrayOutput() VoiceConnectorTerminationArrayOutput
	ToVoiceConnectorTerminationArrayOutputWithContext(context.Context) VoiceConnectorTerminationArrayOutput
}

type VoiceConnectorTerminationArray []VoiceConnectorTerminationInput

func (VoiceConnectorTerminationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VoiceConnectorTermination)(nil)).Elem()
}

func (i VoiceConnectorTerminationArray) ToVoiceConnectorTerminationArrayOutput() VoiceConnectorTerminationArrayOutput {
	return i.ToVoiceConnectorTerminationArrayOutputWithContext(context.Background())
}

func (i VoiceConnectorTerminationArray) ToVoiceConnectorTerminationArrayOutputWithContext(ctx context.Context) VoiceConnectorTerminationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationArrayOutput)
}

// VoiceConnectorTerminationMapInput is an input type that accepts VoiceConnectorTerminationMap and VoiceConnectorTerminationMapOutput values.
// You can construct a concrete instance of `VoiceConnectorTerminationMapInput` via:
//
//	VoiceConnectorTerminationMap{ "key": VoiceConnectorTerminationArgs{...} }
type VoiceConnectorTerminationMapInput interface {
	pulumi.Input

	ToVoiceConnectorTerminationMapOutput() VoiceConnectorTerminationMapOutput
	ToVoiceConnectorTerminationMapOutputWithContext(context.Context) VoiceConnectorTerminationMapOutput
}

type VoiceConnectorTerminationMap map[string]VoiceConnectorTerminationInput

func (VoiceConnectorTerminationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VoiceConnectorTermination)(nil)).Elem()
}

func (i VoiceConnectorTerminationMap) ToVoiceConnectorTerminationMapOutput() VoiceConnectorTerminationMapOutput {
	return i.ToVoiceConnectorTerminationMapOutputWithContext(context.Background())
}

func (i VoiceConnectorTerminationMap) ToVoiceConnectorTerminationMapOutputWithContext(ctx context.Context) VoiceConnectorTerminationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationMapOutput)
}

type VoiceConnectorTerminationOutput struct{ *pulumi.OutputState }

func (VoiceConnectorTerminationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceConnectorTermination)(nil)).Elem()
}

func (o VoiceConnectorTerminationOutput) ToVoiceConnectorTerminationOutput() VoiceConnectorTerminationOutput {
	return o
}

func (o VoiceConnectorTerminationOutput) ToVoiceConnectorTerminationOutputWithContext(ctx context.Context) VoiceConnectorTerminationOutput {
	return o
}

// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
func (o VoiceConnectorTerminationOutput) CallingRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VoiceConnectorTermination) pulumi.StringArrayOutput { return v.CallingRegions }).(pulumi.StringArrayOutput)
}

// The IP addresses allowed to make calls, in CIDR format.
func (o VoiceConnectorTerminationOutput) CidrAllowLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VoiceConnectorTermination) pulumi.StringArrayOutput { return v.CidrAllowLists }).(pulumi.StringArrayOutput)
}

// The limit on calls per second. Max value based on account service quota. Default value of `1`.
func (o VoiceConnectorTerminationOutput) CpsLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VoiceConnectorTermination) pulumi.IntPtrOutput { return v.CpsLimit }).(pulumi.IntPtrOutput)
}

// The default caller ID phone number.
func (o VoiceConnectorTerminationOutput) DefaultPhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VoiceConnectorTermination) pulumi.StringPtrOutput { return v.DefaultPhoneNumber }).(pulumi.StringPtrOutput)
}

// When termination settings are disabled, outbound calls can not be made.
func (o VoiceConnectorTerminationOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VoiceConnectorTermination) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The Amazon Chime Voice Connector ID.
func (o VoiceConnectorTerminationOutput) VoiceConnectorId() pulumi.StringOutput {
	return o.ApplyT(func(v *VoiceConnectorTermination) pulumi.StringOutput { return v.VoiceConnectorId }).(pulumi.StringOutput)
}

type VoiceConnectorTerminationArrayOutput struct{ *pulumi.OutputState }

func (VoiceConnectorTerminationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VoiceConnectorTermination)(nil)).Elem()
}

func (o VoiceConnectorTerminationArrayOutput) ToVoiceConnectorTerminationArrayOutput() VoiceConnectorTerminationArrayOutput {
	return o
}

func (o VoiceConnectorTerminationArrayOutput) ToVoiceConnectorTerminationArrayOutputWithContext(ctx context.Context) VoiceConnectorTerminationArrayOutput {
	return o
}

func (o VoiceConnectorTerminationArrayOutput) Index(i pulumi.IntInput) VoiceConnectorTerminationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VoiceConnectorTermination {
		return vs[0].([]*VoiceConnectorTermination)[vs[1].(int)]
	}).(VoiceConnectorTerminationOutput)
}

type VoiceConnectorTerminationMapOutput struct{ *pulumi.OutputState }

func (VoiceConnectorTerminationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VoiceConnectorTermination)(nil)).Elem()
}

func (o VoiceConnectorTerminationMapOutput) ToVoiceConnectorTerminationMapOutput() VoiceConnectorTerminationMapOutput {
	return o
}

func (o VoiceConnectorTerminationMapOutput) ToVoiceConnectorTerminationMapOutputWithContext(ctx context.Context) VoiceConnectorTerminationMapOutput {
	return o
}

func (o VoiceConnectorTerminationMapOutput) MapIndex(k pulumi.StringInput) VoiceConnectorTerminationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VoiceConnectorTermination {
		return vs[0].(map[string]*VoiceConnectorTermination)[vs[1].(string)]
	}).(VoiceConnectorTerminationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorTerminationInput)(nil)).Elem(), &VoiceConnectorTermination{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorTerminationArrayInput)(nil)).Elem(), VoiceConnectorTerminationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorTerminationMapInput)(nil)).Elem(), VoiceConnectorTerminationMap{})
	pulumi.RegisterOutputType(VoiceConnectorTerminationOutput{})
	pulumi.RegisterOutputType(VoiceConnectorTerminationArrayOutput{})
	pulumi.RegisterOutputType(VoiceConnectorTerminationMapOutput{})
}
