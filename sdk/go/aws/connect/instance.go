// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon Connect instance resource. For more information see
// [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)
//
// !> **WARN:** Amazon Connect enforces a limit of [100 combined instance creation and deletions every 30 days](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#feature-limits).  For example, if you create 80 instances and delete 20 of them, you must wait 30 days to create or delete another instance.  Use care when creating or deleting instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewInstance(ctx, "test", &connect.InstanceArgs{
//				IdentityManagementType: pulumi.String("CONNECT_MANAGED"),
//				InboundCallsEnabled:    pulumi.Bool(true),
//				InstanceAlias:          pulumi.String("friendly-name-connect"),
//				OutboundCallsEnabled:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With Existing Active Directory
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewInstance(ctx, "test", &connect.InstanceArgs{
//				DirectoryId:            pulumi.Any(aws_directory_service_directory.Test.Id),
//				IdentityManagementType: pulumi.String("EXISTING_DIRECTORY"),
//				InboundCallsEnabled:    pulumi.Bool(true),
//				InstanceAlias:          pulumi.String("friendly-name-connect"),
//				OutboundCallsEnabled:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With SAML
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewInstance(ctx, "test", &connect.InstanceArgs{
//				IdentityManagementType: pulumi.String("SAML"),
//				InboundCallsEnabled:    pulumi.Bool(true),
//				InstanceAlias:          pulumi.String("friendly-name-connect"),
//				OutboundCallsEnabled:   pulumi.Bool(true),
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
// Connect instances can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:connect/instance:Instance example f1288a1f-6193-445a-b47e-af739b2
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the instance.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether auto resolve best voices is enabled. Defaults to `true`.
	AutoResolveBestVoicesEnabled pulumi.BoolPtrOutput `pulumi:"autoResolveBestVoicesEnabled"`
	// Specifies whether contact flow logs are enabled. Defaults to `false`.
	ContactFlowLogsEnabled pulumi.BoolPtrOutput `pulumi:"contactFlowLogsEnabled"`
	// Specifies whether contact lens is enabled. Defaults to `true`.
	ContactLensEnabled pulumi.BoolPtrOutput `pulumi:"contactLensEnabled"`
	// When the instance was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The identifier for the directory if identityManagementType is `EXISTING_DIRECTORY`.
	DirectoryId pulumi.StringPtrOutput `pulumi:"directoryId"`
	// Specifies whether early media for outbound calls is enabled . Defaults to `true` if outbound calls is enabled.
	EarlyMediaEnabled pulumi.BoolPtrOutput `pulumi:"earlyMediaEnabled"`
	// Specifies the identity management type attached to the instance. Allowed Values are: `SAML`, `CONNECT_MANAGED`, `EXISTING_DIRECTORY`.
	IdentityManagementType pulumi.StringOutput `pulumi:"identityManagementType"`
	// Specifies whether inbound calls are enabled.
	InboundCallsEnabled pulumi.BoolOutput `pulumi:"inboundCallsEnabled"`
	// Specifies the name of the instance. Required if `directoryId` not specified.
	InstanceAlias pulumi.StringPtrOutput `pulumi:"instanceAlias"`
	// Specifies whether multi-party calls/conference is enabled. Defaults to `false`.
	MultiPartyConferenceEnabled pulumi.BoolPtrOutput `pulumi:"multiPartyConferenceEnabled"`
	// Specifies whether outbound calls are enabled.
	// <!-- * `useCustomTtsVoices` - (Optional) Whether use custom tts voices is enabled. Defaults to `false` -->
	OutboundCallsEnabled pulumi.BoolOutput `pulumi:"outboundCallsEnabled"`
	// The service role of the instance.
	ServiceRole pulumi.StringOutput `pulumi:"serviceRole"`
	// The state of the instance.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityManagementType == nil {
		return nil, errors.New("invalid value for required argument 'IdentityManagementType'")
	}
	if args.InboundCallsEnabled == nil {
		return nil, errors.New("invalid value for required argument 'InboundCallsEnabled'")
	}
	if args.OutboundCallsEnabled == nil {
		return nil, errors.New("invalid value for required argument 'OutboundCallsEnabled'")
	}
	var resource Instance
	err := ctx.RegisterResource("aws:connect/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("aws:connect/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Amazon Resource Name (ARN) of the instance.
	Arn *string `pulumi:"arn"`
	// Specifies whether auto resolve best voices is enabled. Defaults to `true`.
	AutoResolveBestVoicesEnabled *bool `pulumi:"autoResolveBestVoicesEnabled"`
	// Specifies whether contact flow logs are enabled. Defaults to `false`.
	ContactFlowLogsEnabled *bool `pulumi:"contactFlowLogsEnabled"`
	// Specifies whether contact lens is enabled. Defaults to `true`.
	ContactLensEnabled *bool `pulumi:"contactLensEnabled"`
	// When the instance was created.
	CreatedTime *string `pulumi:"createdTime"`
	// The identifier for the directory if identityManagementType is `EXISTING_DIRECTORY`.
	DirectoryId *string `pulumi:"directoryId"`
	// Specifies whether early media for outbound calls is enabled . Defaults to `true` if outbound calls is enabled.
	EarlyMediaEnabled *bool `pulumi:"earlyMediaEnabled"`
	// Specifies the identity management type attached to the instance. Allowed Values are: `SAML`, `CONNECT_MANAGED`, `EXISTING_DIRECTORY`.
	IdentityManagementType *string `pulumi:"identityManagementType"`
	// Specifies whether inbound calls are enabled.
	InboundCallsEnabled *bool `pulumi:"inboundCallsEnabled"`
	// Specifies the name of the instance. Required if `directoryId` not specified.
	InstanceAlias *string `pulumi:"instanceAlias"`
	// Specifies whether multi-party calls/conference is enabled. Defaults to `false`.
	MultiPartyConferenceEnabled *bool `pulumi:"multiPartyConferenceEnabled"`
	// Specifies whether outbound calls are enabled.
	// <!-- * `useCustomTtsVoices` - (Optional) Whether use custom tts voices is enabled. Defaults to `false` -->
	OutboundCallsEnabled *bool `pulumi:"outboundCallsEnabled"`
	// The service role of the instance.
	ServiceRole *string `pulumi:"serviceRole"`
	// The state of the instance.
	Status *string `pulumi:"status"`
}

type InstanceState struct {
	// Amazon Resource Name (ARN) of the instance.
	Arn pulumi.StringPtrInput
	// Specifies whether auto resolve best voices is enabled. Defaults to `true`.
	AutoResolveBestVoicesEnabled pulumi.BoolPtrInput
	// Specifies whether contact flow logs are enabled. Defaults to `false`.
	ContactFlowLogsEnabled pulumi.BoolPtrInput
	// Specifies whether contact lens is enabled. Defaults to `true`.
	ContactLensEnabled pulumi.BoolPtrInput
	// When the instance was created.
	CreatedTime pulumi.StringPtrInput
	// The identifier for the directory if identityManagementType is `EXISTING_DIRECTORY`.
	DirectoryId pulumi.StringPtrInput
	// Specifies whether early media for outbound calls is enabled . Defaults to `true` if outbound calls is enabled.
	EarlyMediaEnabled pulumi.BoolPtrInput
	// Specifies the identity management type attached to the instance. Allowed Values are: `SAML`, `CONNECT_MANAGED`, `EXISTING_DIRECTORY`.
	IdentityManagementType pulumi.StringPtrInput
	// Specifies whether inbound calls are enabled.
	InboundCallsEnabled pulumi.BoolPtrInput
	// Specifies the name of the instance. Required if `directoryId` not specified.
	InstanceAlias pulumi.StringPtrInput
	// Specifies whether multi-party calls/conference is enabled. Defaults to `false`.
	MultiPartyConferenceEnabled pulumi.BoolPtrInput
	// Specifies whether outbound calls are enabled.
	// <!-- * `useCustomTtsVoices` - (Optional) Whether use custom tts voices is enabled. Defaults to `false` -->
	OutboundCallsEnabled pulumi.BoolPtrInput
	// The service role of the instance.
	ServiceRole pulumi.StringPtrInput
	// The state of the instance.
	Status pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Specifies whether auto resolve best voices is enabled. Defaults to `true`.
	AutoResolveBestVoicesEnabled *bool `pulumi:"autoResolveBestVoicesEnabled"`
	// Specifies whether contact flow logs are enabled. Defaults to `false`.
	ContactFlowLogsEnabled *bool `pulumi:"contactFlowLogsEnabled"`
	// Specifies whether contact lens is enabled. Defaults to `true`.
	ContactLensEnabled *bool `pulumi:"contactLensEnabled"`
	// The identifier for the directory if identityManagementType is `EXISTING_DIRECTORY`.
	DirectoryId *string `pulumi:"directoryId"`
	// Specifies whether early media for outbound calls is enabled . Defaults to `true` if outbound calls is enabled.
	EarlyMediaEnabled *bool `pulumi:"earlyMediaEnabled"`
	// Specifies the identity management type attached to the instance. Allowed Values are: `SAML`, `CONNECT_MANAGED`, `EXISTING_DIRECTORY`.
	IdentityManagementType string `pulumi:"identityManagementType"`
	// Specifies whether inbound calls are enabled.
	InboundCallsEnabled bool `pulumi:"inboundCallsEnabled"`
	// Specifies the name of the instance. Required if `directoryId` not specified.
	InstanceAlias *string `pulumi:"instanceAlias"`
	// Specifies whether multi-party calls/conference is enabled. Defaults to `false`.
	MultiPartyConferenceEnabled *bool `pulumi:"multiPartyConferenceEnabled"`
	// Specifies whether outbound calls are enabled.
	// <!-- * `useCustomTtsVoices` - (Optional) Whether use custom tts voices is enabled. Defaults to `false` -->
	OutboundCallsEnabled bool `pulumi:"outboundCallsEnabled"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Specifies whether auto resolve best voices is enabled. Defaults to `true`.
	AutoResolveBestVoicesEnabled pulumi.BoolPtrInput
	// Specifies whether contact flow logs are enabled. Defaults to `false`.
	ContactFlowLogsEnabled pulumi.BoolPtrInput
	// Specifies whether contact lens is enabled. Defaults to `true`.
	ContactLensEnabled pulumi.BoolPtrInput
	// The identifier for the directory if identityManagementType is `EXISTING_DIRECTORY`.
	DirectoryId pulumi.StringPtrInput
	// Specifies whether early media for outbound calls is enabled . Defaults to `true` if outbound calls is enabled.
	EarlyMediaEnabled pulumi.BoolPtrInput
	// Specifies the identity management type attached to the instance. Allowed Values are: `SAML`, `CONNECT_MANAGED`, `EXISTING_DIRECTORY`.
	IdentityManagementType pulumi.StringInput
	// Specifies whether inbound calls are enabled.
	InboundCallsEnabled pulumi.BoolInput
	// Specifies the name of the instance. Required if `directoryId` not specified.
	InstanceAlias pulumi.StringPtrInput
	// Specifies whether multi-party calls/conference is enabled. Defaults to `false`.
	MultiPartyConferenceEnabled pulumi.BoolPtrInput
	// Specifies whether outbound calls are enabled.
	// <!-- * `useCustomTtsVoices` - (Optional) Whether use custom tts voices is enabled. Defaults to `false` -->
	OutboundCallsEnabled pulumi.BoolInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Amazon Resource Name (ARN) of the instance.
func (o InstanceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies whether auto resolve best voices is enabled. Defaults to `true`.
func (o InstanceOutput) AutoResolveBestVoicesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.AutoResolveBestVoicesEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether contact flow logs are enabled. Defaults to `false`.
func (o InstanceOutput) ContactFlowLogsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.ContactFlowLogsEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether contact lens is enabled. Defaults to `true`.
func (o InstanceOutput) ContactLensEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.ContactLensEnabled }).(pulumi.BoolPtrOutput)
}

// When the instance was created.
func (o InstanceOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The identifier for the directory if identityManagementType is `EXISTING_DIRECTORY`.
func (o InstanceOutput) DirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.DirectoryId }).(pulumi.StringPtrOutput)
}

// Specifies whether early media for outbound calls is enabled . Defaults to `true` if outbound calls is enabled.
func (o InstanceOutput) EarlyMediaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.EarlyMediaEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies the identity management type attached to the instance. Allowed Values are: `SAML`, `CONNECT_MANAGED`, `EXISTING_DIRECTORY`.
func (o InstanceOutput) IdentityManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.IdentityManagementType }).(pulumi.StringOutput)
}

// Specifies whether inbound calls are enabled.
func (o InstanceOutput) InboundCallsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.InboundCallsEnabled }).(pulumi.BoolOutput)
}

// Specifies the name of the instance. Required if `directoryId` not specified.
func (o InstanceOutput) InstanceAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceAlias }).(pulumi.StringPtrOutput)
}

// Specifies whether multi-party calls/conference is enabled. Defaults to `false`.
func (o InstanceOutput) MultiPartyConferenceEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.MultiPartyConferenceEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether outbound calls are enabled.
// <!-- * `useCustomTtsVoices` - (Optional) Whether use custom tts voices is enabled. Defaults to `false` -->
func (o InstanceOutput) OutboundCallsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.OutboundCallsEnabled }).(pulumi.BoolOutput)
}

// The service role of the instance.
func (o InstanceOutput) ServiceRole() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ServiceRole }).(pulumi.StringOutput)
}

// The state of the instance.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
