// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a directory's multi-factor authentication (MFA) using a Remote Authentication Dial In User Service (RADIUS) server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := directoryservice.NewRadiusSettings(ctx, "example", &directoryservice.RadiusSettingsArgs{
//				DirectoryId:            pulumi.Any(aws_directory_service_directory.Example.Id),
//				AuthenticationProtocol: pulumi.String("PAP"),
//				DisplayLabel:           pulumi.String("example"),
//				RadiusPort:             pulumi.Int(1812),
//				RadiusRetries:          pulumi.Int(4),
//				RadiusServers: pulumi.StringArray{
//					pulumi.String("10.0.1.5"),
//				},
//				RadiusTimeout: pulumi.Int(1),
//				SharedSecret:  pulumi.String("12345678"),
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
//	to = aws_directory_service_radius_settings.example
//
//	id = "d-926724cf57" } Using `pulumi import`, import RADIUS settings using the directory ID. For exampleconsole % pulumi import aws_directory_service_radius_settings.example d-926724cf57
type RadiusSettings struct {
	pulumi.CustomResourceState

	// The protocol specified for your RADIUS endpoints. Valid values: `PAP`, `CHAP`, `MS-CHAPv1`, `MS-CHAPv2`.
	AuthenticationProtocol pulumi.StringOutput `pulumi:"authenticationProtocol"`
	// The identifier of the directory for which you want to manager RADIUS settings.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// Display label.
	DisplayLabel pulumi.StringOutput `pulumi:"displayLabel"`
	// The port that your RADIUS server is using for communications. Your self-managed network must allow inbound traffic over this port from the AWS Directory Service servers.
	RadiusPort pulumi.IntOutput `pulumi:"radiusPort"`
	// The maximum number of times that communication with the RADIUS server is attempted. Minimum value of `0`. Maximum value of `10`.
	RadiusRetries pulumi.IntOutput `pulumi:"radiusRetries"`
	// An array of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
	RadiusServers pulumi.StringArrayOutput `pulumi:"radiusServers"`
	// The amount of time, in seconds, to wait for the RADIUS server to respond. Minimum value of `1`. Maximum value of `50`.
	RadiusTimeout pulumi.IntOutput `pulumi:"radiusTimeout"`
	// Required for enabling RADIUS on the directory.
	SharedSecret pulumi.StringOutput `pulumi:"sharedSecret"`
	// Not currently used.
	UseSameUsername pulumi.BoolPtrOutput `pulumi:"useSameUsername"`
}

// NewRadiusSettings registers a new resource with the given unique name, arguments, and options.
func NewRadiusSettings(ctx *pulumi.Context,
	name string, args *RadiusSettingsArgs, opts ...pulumi.ResourceOption) (*RadiusSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationProtocol == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationProtocol'")
	}
	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.DisplayLabel == nil {
		return nil, errors.New("invalid value for required argument 'DisplayLabel'")
	}
	if args.RadiusPort == nil {
		return nil, errors.New("invalid value for required argument 'RadiusPort'")
	}
	if args.RadiusRetries == nil {
		return nil, errors.New("invalid value for required argument 'RadiusRetries'")
	}
	if args.RadiusServers == nil {
		return nil, errors.New("invalid value for required argument 'RadiusServers'")
	}
	if args.RadiusTimeout == nil {
		return nil, errors.New("invalid value for required argument 'RadiusTimeout'")
	}
	if args.SharedSecret == nil {
		return nil, errors.New("invalid value for required argument 'SharedSecret'")
	}
	if args.SharedSecret != nil {
		args.SharedSecret = pulumi.ToSecret(args.SharedSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"sharedSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RadiusSettings
	err := ctx.RegisterResource("aws:directoryservice/radiusSettings:RadiusSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRadiusSettings gets an existing RadiusSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRadiusSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RadiusSettingsState, opts ...pulumi.ResourceOption) (*RadiusSettings, error) {
	var resource RadiusSettings
	err := ctx.ReadResource("aws:directoryservice/radiusSettings:RadiusSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RadiusSettings resources.
type radiusSettingsState struct {
	// The protocol specified for your RADIUS endpoints. Valid values: `PAP`, `CHAP`, `MS-CHAPv1`, `MS-CHAPv2`.
	AuthenticationProtocol *string `pulumi:"authenticationProtocol"`
	// The identifier of the directory for which you want to manager RADIUS settings.
	DirectoryId *string `pulumi:"directoryId"`
	// Display label.
	DisplayLabel *string `pulumi:"displayLabel"`
	// The port that your RADIUS server is using for communications. Your self-managed network must allow inbound traffic over this port from the AWS Directory Service servers.
	RadiusPort *int `pulumi:"radiusPort"`
	// The maximum number of times that communication with the RADIUS server is attempted. Minimum value of `0`. Maximum value of `10`.
	RadiusRetries *int `pulumi:"radiusRetries"`
	// An array of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
	RadiusServers []string `pulumi:"radiusServers"`
	// The amount of time, in seconds, to wait for the RADIUS server to respond. Minimum value of `1`. Maximum value of `50`.
	RadiusTimeout *int `pulumi:"radiusTimeout"`
	// Required for enabling RADIUS on the directory.
	SharedSecret *string `pulumi:"sharedSecret"`
	// Not currently used.
	UseSameUsername *bool `pulumi:"useSameUsername"`
}

type RadiusSettingsState struct {
	// The protocol specified for your RADIUS endpoints. Valid values: `PAP`, `CHAP`, `MS-CHAPv1`, `MS-CHAPv2`.
	AuthenticationProtocol pulumi.StringPtrInput
	// The identifier of the directory for which you want to manager RADIUS settings.
	DirectoryId pulumi.StringPtrInput
	// Display label.
	DisplayLabel pulumi.StringPtrInput
	// The port that your RADIUS server is using for communications. Your self-managed network must allow inbound traffic over this port from the AWS Directory Service servers.
	RadiusPort pulumi.IntPtrInput
	// The maximum number of times that communication with the RADIUS server is attempted. Minimum value of `0`. Maximum value of `10`.
	RadiusRetries pulumi.IntPtrInput
	// An array of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
	RadiusServers pulumi.StringArrayInput
	// The amount of time, in seconds, to wait for the RADIUS server to respond. Minimum value of `1`. Maximum value of `50`.
	RadiusTimeout pulumi.IntPtrInput
	// Required for enabling RADIUS on the directory.
	SharedSecret pulumi.StringPtrInput
	// Not currently used.
	UseSameUsername pulumi.BoolPtrInput
}

func (RadiusSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*radiusSettingsState)(nil)).Elem()
}

type radiusSettingsArgs struct {
	// The protocol specified for your RADIUS endpoints. Valid values: `PAP`, `CHAP`, `MS-CHAPv1`, `MS-CHAPv2`.
	AuthenticationProtocol string `pulumi:"authenticationProtocol"`
	// The identifier of the directory for which you want to manager RADIUS settings.
	DirectoryId string `pulumi:"directoryId"`
	// Display label.
	DisplayLabel string `pulumi:"displayLabel"`
	// The port that your RADIUS server is using for communications. Your self-managed network must allow inbound traffic over this port from the AWS Directory Service servers.
	RadiusPort int `pulumi:"radiusPort"`
	// The maximum number of times that communication with the RADIUS server is attempted. Minimum value of `0`. Maximum value of `10`.
	RadiusRetries int `pulumi:"radiusRetries"`
	// An array of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
	RadiusServers []string `pulumi:"radiusServers"`
	// The amount of time, in seconds, to wait for the RADIUS server to respond. Minimum value of `1`. Maximum value of `50`.
	RadiusTimeout int `pulumi:"radiusTimeout"`
	// Required for enabling RADIUS on the directory.
	SharedSecret string `pulumi:"sharedSecret"`
	// Not currently used.
	UseSameUsername *bool `pulumi:"useSameUsername"`
}

// The set of arguments for constructing a RadiusSettings resource.
type RadiusSettingsArgs struct {
	// The protocol specified for your RADIUS endpoints. Valid values: `PAP`, `CHAP`, `MS-CHAPv1`, `MS-CHAPv2`.
	AuthenticationProtocol pulumi.StringInput
	// The identifier of the directory for which you want to manager RADIUS settings.
	DirectoryId pulumi.StringInput
	// Display label.
	DisplayLabel pulumi.StringInput
	// The port that your RADIUS server is using for communications. Your self-managed network must allow inbound traffic over this port from the AWS Directory Service servers.
	RadiusPort pulumi.IntInput
	// The maximum number of times that communication with the RADIUS server is attempted. Minimum value of `0`. Maximum value of `10`.
	RadiusRetries pulumi.IntInput
	// An array of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
	RadiusServers pulumi.StringArrayInput
	// The amount of time, in seconds, to wait for the RADIUS server to respond. Minimum value of `1`. Maximum value of `50`.
	RadiusTimeout pulumi.IntInput
	// Required for enabling RADIUS on the directory.
	SharedSecret pulumi.StringInput
	// Not currently used.
	UseSameUsername pulumi.BoolPtrInput
}

func (RadiusSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*radiusSettingsArgs)(nil)).Elem()
}

type RadiusSettingsInput interface {
	pulumi.Input

	ToRadiusSettingsOutput() RadiusSettingsOutput
	ToRadiusSettingsOutputWithContext(ctx context.Context) RadiusSettingsOutput
}

func (*RadiusSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**RadiusSettings)(nil)).Elem()
}

func (i *RadiusSettings) ToRadiusSettingsOutput() RadiusSettingsOutput {
	return i.ToRadiusSettingsOutputWithContext(context.Background())
}

func (i *RadiusSettings) ToRadiusSettingsOutputWithContext(ctx context.Context) RadiusSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusSettingsOutput)
}

// RadiusSettingsArrayInput is an input type that accepts RadiusSettingsArray and RadiusSettingsArrayOutput values.
// You can construct a concrete instance of `RadiusSettingsArrayInput` via:
//
//	RadiusSettingsArray{ RadiusSettingsArgs{...} }
type RadiusSettingsArrayInput interface {
	pulumi.Input

	ToRadiusSettingsArrayOutput() RadiusSettingsArrayOutput
	ToRadiusSettingsArrayOutputWithContext(context.Context) RadiusSettingsArrayOutput
}

type RadiusSettingsArray []RadiusSettingsInput

func (RadiusSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RadiusSettings)(nil)).Elem()
}

func (i RadiusSettingsArray) ToRadiusSettingsArrayOutput() RadiusSettingsArrayOutput {
	return i.ToRadiusSettingsArrayOutputWithContext(context.Background())
}

func (i RadiusSettingsArray) ToRadiusSettingsArrayOutputWithContext(ctx context.Context) RadiusSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusSettingsArrayOutput)
}

// RadiusSettingsMapInput is an input type that accepts RadiusSettingsMap and RadiusSettingsMapOutput values.
// You can construct a concrete instance of `RadiusSettingsMapInput` via:
//
//	RadiusSettingsMap{ "key": RadiusSettingsArgs{...} }
type RadiusSettingsMapInput interface {
	pulumi.Input

	ToRadiusSettingsMapOutput() RadiusSettingsMapOutput
	ToRadiusSettingsMapOutputWithContext(context.Context) RadiusSettingsMapOutput
}

type RadiusSettingsMap map[string]RadiusSettingsInput

func (RadiusSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RadiusSettings)(nil)).Elem()
}

func (i RadiusSettingsMap) ToRadiusSettingsMapOutput() RadiusSettingsMapOutput {
	return i.ToRadiusSettingsMapOutputWithContext(context.Background())
}

func (i RadiusSettingsMap) ToRadiusSettingsMapOutputWithContext(ctx context.Context) RadiusSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusSettingsMapOutput)
}

type RadiusSettingsOutput struct{ *pulumi.OutputState }

func (RadiusSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RadiusSettings)(nil)).Elem()
}

func (o RadiusSettingsOutput) ToRadiusSettingsOutput() RadiusSettingsOutput {
	return o
}

func (o RadiusSettingsOutput) ToRadiusSettingsOutputWithContext(ctx context.Context) RadiusSettingsOutput {
	return o
}

// The protocol specified for your RADIUS endpoints. Valid values: `PAP`, `CHAP`, `MS-CHAPv1`, `MS-CHAPv2`.
func (o RadiusSettingsOutput) AuthenticationProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.StringOutput { return v.AuthenticationProtocol }).(pulumi.StringOutput)
}

// The identifier of the directory for which you want to manager RADIUS settings.
func (o RadiusSettingsOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// Display label.
func (o RadiusSettingsOutput) DisplayLabel() pulumi.StringOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.StringOutput { return v.DisplayLabel }).(pulumi.StringOutput)
}

// The port that your RADIUS server is using for communications. Your self-managed network must allow inbound traffic over this port from the AWS Directory Service servers.
func (o RadiusSettingsOutput) RadiusPort() pulumi.IntOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.IntOutput { return v.RadiusPort }).(pulumi.IntOutput)
}

// The maximum number of times that communication with the RADIUS server is attempted. Minimum value of `0`. Maximum value of `10`.
func (o RadiusSettingsOutput) RadiusRetries() pulumi.IntOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.IntOutput { return v.RadiusRetries }).(pulumi.IntOutput)
}

// An array of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
func (o RadiusSettingsOutput) RadiusServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.StringArrayOutput { return v.RadiusServers }).(pulumi.StringArrayOutput)
}

// The amount of time, in seconds, to wait for the RADIUS server to respond. Minimum value of `1`. Maximum value of `50`.
func (o RadiusSettingsOutput) RadiusTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.IntOutput { return v.RadiusTimeout }).(pulumi.IntOutput)
}

// Required for enabling RADIUS on the directory.
func (o RadiusSettingsOutput) SharedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.StringOutput { return v.SharedSecret }).(pulumi.StringOutput)
}

// Not currently used.
func (o RadiusSettingsOutput) UseSameUsername() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RadiusSettings) pulumi.BoolPtrOutput { return v.UseSameUsername }).(pulumi.BoolPtrOutput)
}

type RadiusSettingsArrayOutput struct{ *pulumi.OutputState }

func (RadiusSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RadiusSettings)(nil)).Elem()
}

func (o RadiusSettingsArrayOutput) ToRadiusSettingsArrayOutput() RadiusSettingsArrayOutput {
	return o
}

func (o RadiusSettingsArrayOutput) ToRadiusSettingsArrayOutputWithContext(ctx context.Context) RadiusSettingsArrayOutput {
	return o
}

func (o RadiusSettingsArrayOutput) Index(i pulumi.IntInput) RadiusSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RadiusSettings {
		return vs[0].([]*RadiusSettings)[vs[1].(int)]
	}).(RadiusSettingsOutput)
}

type RadiusSettingsMapOutput struct{ *pulumi.OutputState }

func (RadiusSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RadiusSettings)(nil)).Elem()
}

func (o RadiusSettingsMapOutput) ToRadiusSettingsMapOutput() RadiusSettingsMapOutput {
	return o
}

func (o RadiusSettingsMapOutput) ToRadiusSettingsMapOutputWithContext(ctx context.Context) RadiusSettingsMapOutput {
	return o
}

func (o RadiusSettingsMapOutput) MapIndex(k pulumi.StringInput) RadiusSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RadiusSettings {
		return vs[0].(map[string]*RadiusSettings)[vs[1].(string)]
	}).(RadiusSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusSettingsInput)(nil)).Elem(), &RadiusSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusSettingsArrayInput)(nil)).Elem(), RadiusSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusSettingsMapInput)(nil)).Elem(), RadiusSettingsMap{})
	pulumi.RegisterOutputType(RadiusSettingsOutput{})
	pulumi.RegisterOutputType(RadiusSettingsArrayOutput{})
	pulumi.RegisterOutputType(RadiusSettingsMapOutput{})
}
