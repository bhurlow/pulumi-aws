// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Pinpoint APNs VoIP Sandbox Channel resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/pinpoint"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			app, err := pinpoint.NewApp(ctx, "app", nil)
//			if err != nil {
//				return err
//			}
//			_, err = pinpoint.NewApnsVoipSandboxChannel(ctx, "apnsVoipSandbox", &pinpoint.ApnsVoipSandboxChannelArgs{
//				ApplicationId: app.ApplicationId,
//				Certificate:   readFileOrPanic("./certificate.pem"),
//				PrivateKey:    readFileOrPanic("./private_key.key"),
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
// Pinpoint APNs VoIP Sandbox Channel can be imported using the `application-id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:pinpoint/apnsVoipSandboxChannel:ApnsVoipSandboxChannel apns_voip_sandbox application-id
//
// ```
type ApnsVoipSandboxChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrOutput `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrOutput `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrOutput `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrOutput `pulumi:"tokenKeyId"`
}

// NewApnsVoipSandboxChannel registers a new resource with the given unique name, arguments, and options.
func NewApnsVoipSandboxChannel(ctx *pulumi.Context,
	name string, args *ApnsVoipSandboxChannelArgs, opts ...pulumi.ResourceOption) (*ApnsVoipSandboxChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.BundleId != nil {
		args.BundleId = pulumi.ToSecret(args.BundleId).(pulumi.StringPtrInput)
	}
	if args.Certificate != nil {
		args.Certificate = pulumi.ToSecret(args.Certificate).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	if args.TeamId != nil {
		args.TeamId = pulumi.ToSecret(args.TeamId).(pulumi.StringPtrInput)
	}
	if args.TokenKey != nil {
		args.TokenKey = pulumi.ToSecret(args.TokenKey).(pulumi.StringPtrInput)
	}
	if args.TokenKeyId != nil {
		args.TokenKeyId = pulumi.ToSecret(args.TokenKeyId).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"bundleId",
		"certificate",
		"privateKey",
		"teamId",
		"tokenKey",
		"tokenKeyId",
	})
	opts = append(opts, secrets)
	var resource ApnsVoipSandboxChannel
	err := ctx.RegisterResource("aws:pinpoint/apnsVoipSandboxChannel:ApnsVoipSandboxChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApnsVoipSandboxChannel gets an existing ApnsVoipSandboxChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApnsVoipSandboxChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApnsVoipSandboxChannelState, opts ...pulumi.ResourceOption) (*ApnsVoipSandboxChannel, error) {
	var resource ApnsVoipSandboxChannel
	err := ctx.ReadResource("aws:pinpoint/apnsVoipSandboxChannel:ApnsVoipSandboxChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApnsVoipSandboxChannel resources.
type apnsVoipSandboxChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey *string `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId *string `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey *string `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId *string `pulumi:"tokenKeyId"`
}

type ApnsVoipSandboxChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrInput
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrInput
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrInput
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrInput
}

func (ApnsVoipSandboxChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsVoipSandboxChannelState)(nil)).Elem()
}

type apnsVoipSandboxChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey *string `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId *string `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey *string `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId *string `pulumi:"tokenKeyId"`
}

// The set of arguments for constructing a ApnsVoipSandboxChannel resource.
type ApnsVoipSandboxChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrInput
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrInput
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrInput
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrInput
}

func (ApnsVoipSandboxChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsVoipSandboxChannelArgs)(nil)).Elem()
}

type ApnsVoipSandboxChannelInput interface {
	pulumi.Input

	ToApnsVoipSandboxChannelOutput() ApnsVoipSandboxChannelOutput
	ToApnsVoipSandboxChannelOutputWithContext(ctx context.Context) ApnsVoipSandboxChannelOutput
}

func (*ApnsVoipSandboxChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsVoipSandboxChannel)(nil)).Elem()
}

func (i *ApnsVoipSandboxChannel) ToApnsVoipSandboxChannelOutput() ApnsVoipSandboxChannelOutput {
	return i.ToApnsVoipSandboxChannelOutputWithContext(context.Background())
}

func (i *ApnsVoipSandboxChannel) ToApnsVoipSandboxChannelOutputWithContext(ctx context.Context) ApnsVoipSandboxChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipSandboxChannelOutput)
}

// ApnsVoipSandboxChannelArrayInput is an input type that accepts ApnsVoipSandboxChannelArray and ApnsVoipSandboxChannelArrayOutput values.
// You can construct a concrete instance of `ApnsVoipSandboxChannelArrayInput` via:
//
//	ApnsVoipSandboxChannelArray{ ApnsVoipSandboxChannelArgs{...} }
type ApnsVoipSandboxChannelArrayInput interface {
	pulumi.Input

	ToApnsVoipSandboxChannelArrayOutput() ApnsVoipSandboxChannelArrayOutput
	ToApnsVoipSandboxChannelArrayOutputWithContext(context.Context) ApnsVoipSandboxChannelArrayOutput
}

type ApnsVoipSandboxChannelArray []ApnsVoipSandboxChannelInput

func (ApnsVoipSandboxChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApnsVoipSandboxChannel)(nil)).Elem()
}

func (i ApnsVoipSandboxChannelArray) ToApnsVoipSandboxChannelArrayOutput() ApnsVoipSandboxChannelArrayOutput {
	return i.ToApnsVoipSandboxChannelArrayOutputWithContext(context.Background())
}

func (i ApnsVoipSandboxChannelArray) ToApnsVoipSandboxChannelArrayOutputWithContext(ctx context.Context) ApnsVoipSandboxChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipSandboxChannelArrayOutput)
}

// ApnsVoipSandboxChannelMapInput is an input type that accepts ApnsVoipSandboxChannelMap and ApnsVoipSandboxChannelMapOutput values.
// You can construct a concrete instance of `ApnsVoipSandboxChannelMapInput` via:
//
//	ApnsVoipSandboxChannelMap{ "key": ApnsVoipSandboxChannelArgs{...} }
type ApnsVoipSandboxChannelMapInput interface {
	pulumi.Input

	ToApnsVoipSandboxChannelMapOutput() ApnsVoipSandboxChannelMapOutput
	ToApnsVoipSandboxChannelMapOutputWithContext(context.Context) ApnsVoipSandboxChannelMapOutput
}

type ApnsVoipSandboxChannelMap map[string]ApnsVoipSandboxChannelInput

func (ApnsVoipSandboxChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApnsVoipSandboxChannel)(nil)).Elem()
}

func (i ApnsVoipSandboxChannelMap) ToApnsVoipSandboxChannelMapOutput() ApnsVoipSandboxChannelMapOutput {
	return i.ToApnsVoipSandboxChannelMapOutputWithContext(context.Background())
}

func (i ApnsVoipSandboxChannelMap) ToApnsVoipSandboxChannelMapOutputWithContext(ctx context.Context) ApnsVoipSandboxChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipSandboxChannelMapOutput)
}

type ApnsVoipSandboxChannelOutput struct{ *pulumi.OutputState }

func (ApnsVoipSandboxChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsVoipSandboxChannel)(nil)).Elem()
}

func (o ApnsVoipSandboxChannelOutput) ToApnsVoipSandboxChannelOutput() ApnsVoipSandboxChannelOutput {
	return o
}

func (o ApnsVoipSandboxChannelOutput) ToApnsVoipSandboxChannelOutputWithContext(ctx context.Context) ApnsVoipSandboxChannelOutput {
	return o
}

// The application ID.
func (o ApnsVoipSandboxChannelOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
func (o ApnsVoipSandboxChannelOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringPtrOutput { return v.BundleId }).(pulumi.StringPtrOutput)
}

// The pem encoded TLS Certificate from Apple.
func (o ApnsVoipSandboxChannelOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

// The default authentication method used for APNs.
// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
func (o ApnsVoipSandboxChannelOutput) DefaultAuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringPtrOutput { return v.DefaultAuthenticationMethod }).(pulumi.StringPtrOutput)
}

// Whether the channel is enabled or disabled. Defaults to `true`.
func (o ApnsVoipSandboxChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The Certificate Private Key file (ie. `.key` file).
func (o ApnsVoipSandboxChannelOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
func (o ApnsVoipSandboxChannelOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

// The `.p8` file that you download from your Apple developer account when you create an authentication key.
func (o ApnsVoipSandboxChannelOutput) TokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringPtrOutput { return v.TokenKey }).(pulumi.StringPtrOutput)
}

// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
func (o ApnsVoipSandboxChannelOutput) TokenKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsVoipSandboxChannel) pulumi.StringPtrOutput { return v.TokenKeyId }).(pulumi.StringPtrOutput)
}

type ApnsVoipSandboxChannelArrayOutput struct{ *pulumi.OutputState }

func (ApnsVoipSandboxChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApnsVoipSandboxChannel)(nil)).Elem()
}

func (o ApnsVoipSandboxChannelArrayOutput) ToApnsVoipSandboxChannelArrayOutput() ApnsVoipSandboxChannelArrayOutput {
	return o
}

func (o ApnsVoipSandboxChannelArrayOutput) ToApnsVoipSandboxChannelArrayOutputWithContext(ctx context.Context) ApnsVoipSandboxChannelArrayOutput {
	return o
}

func (o ApnsVoipSandboxChannelArrayOutput) Index(i pulumi.IntInput) ApnsVoipSandboxChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApnsVoipSandboxChannel {
		return vs[0].([]*ApnsVoipSandboxChannel)[vs[1].(int)]
	}).(ApnsVoipSandboxChannelOutput)
}

type ApnsVoipSandboxChannelMapOutput struct{ *pulumi.OutputState }

func (ApnsVoipSandboxChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApnsVoipSandboxChannel)(nil)).Elem()
}

func (o ApnsVoipSandboxChannelMapOutput) ToApnsVoipSandboxChannelMapOutput() ApnsVoipSandboxChannelMapOutput {
	return o
}

func (o ApnsVoipSandboxChannelMapOutput) ToApnsVoipSandboxChannelMapOutputWithContext(ctx context.Context) ApnsVoipSandboxChannelMapOutput {
	return o
}

func (o ApnsVoipSandboxChannelMapOutput) MapIndex(k pulumi.StringInput) ApnsVoipSandboxChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApnsVoipSandboxChannel {
		return vs[0].(map[string]*ApnsVoipSandboxChannel)[vs[1].(string)]
	}).(ApnsVoipSandboxChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApnsVoipSandboxChannelInput)(nil)).Elem(), &ApnsVoipSandboxChannel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApnsVoipSandboxChannelArrayInput)(nil)).Elem(), ApnsVoipSandboxChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApnsVoipSandboxChannelMapInput)(nil)).Elem(), ApnsVoipSandboxChannelMap{})
	pulumi.RegisterOutputType(ApnsVoipSandboxChannelOutput{})
	pulumi.RegisterOutputType(ApnsVoipSandboxChannelArrayOutput{})
	pulumi.RegisterOutputType(ApnsVoipSandboxChannelMapOutput{})
}
