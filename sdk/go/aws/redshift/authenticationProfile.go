// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Redshift authentication profile
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"AllowDBUserOverride": "1",
//				"Client_ID":           "ExampleClientID",
//				"App_ID":              "example",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = redshift.NewAuthenticationProfile(ctx, "example", &redshift.AuthenticationProfileArgs{
//				AuthenticationProfileName:    pulumi.String("example"),
//				AuthenticationProfileContent: pulumi.String(json0),
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
//	to = aws_redshift_authentication_profile.test
//
//	id = "example" } Using `pulumi import`, import Redshift Authentication by `authentication_profile_name`. For exampleconsole % pulumi import aws_redshift_authentication_profile.test example
type AuthenticationProfile struct {
	pulumi.CustomResourceState

	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent pulumi.StringOutput `pulumi:"authenticationProfileContent"`
	// The name of the authentication profile.
	AuthenticationProfileName pulumi.StringOutput `pulumi:"authenticationProfileName"`
}

// NewAuthenticationProfile registers a new resource with the given unique name, arguments, and options.
func NewAuthenticationProfile(ctx *pulumi.Context,
	name string, args *AuthenticationProfileArgs, opts ...pulumi.ResourceOption) (*AuthenticationProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationProfileContent == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationProfileContent'")
	}
	if args.AuthenticationProfileName == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationProfileName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthenticationProfile
	err := ctx.RegisterResource("aws:redshift/authenticationProfile:AuthenticationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthenticationProfile gets an existing AuthenticationProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthenticationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthenticationProfileState, opts ...pulumi.ResourceOption) (*AuthenticationProfile, error) {
	var resource AuthenticationProfile
	err := ctx.ReadResource("aws:redshift/authenticationProfile:AuthenticationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthenticationProfile resources.
type authenticationProfileState struct {
	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent *string `pulumi:"authenticationProfileContent"`
	// The name of the authentication profile.
	AuthenticationProfileName *string `pulumi:"authenticationProfileName"`
}

type AuthenticationProfileState struct {
	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent pulumi.StringPtrInput
	// The name of the authentication profile.
	AuthenticationProfileName pulumi.StringPtrInput
}

func (AuthenticationProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationProfileState)(nil)).Elem()
}

type authenticationProfileArgs struct {
	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent string `pulumi:"authenticationProfileContent"`
	// The name of the authentication profile.
	AuthenticationProfileName string `pulumi:"authenticationProfileName"`
}

// The set of arguments for constructing a AuthenticationProfile resource.
type AuthenticationProfileArgs struct {
	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent pulumi.StringInput
	// The name of the authentication profile.
	AuthenticationProfileName pulumi.StringInput
}

func (AuthenticationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationProfileArgs)(nil)).Elem()
}

type AuthenticationProfileInput interface {
	pulumi.Input

	ToAuthenticationProfileOutput() AuthenticationProfileOutput
	ToAuthenticationProfileOutputWithContext(ctx context.Context) AuthenticationProfileOutput
}

func (*AuthenticationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationProfile)(nil)).Elem()
}

func (i *AuthenticationProfile) ToAuthenticationProfileOutput() AuthenticationProfileOutput {
	return i.ToAuthenticationProfileOutputWithContext(context.Background())
}

func (i *AuthenticationProfile) ToAuthenticationProfileOutputWithContext(ctx context.Context) AuthenticationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationProfileOutput)
}

// AuthenticationProfileArrayInput is an input type that accepts AuthenticationProfileArray and AuthenticationProfileArrayOutput values.
// You can construct a concrete instance of `AuthenticationProfileArrayInput` via:
//
//	AuthenticationProfileArray{ AuthenticationProfileArgs{...} }
type AuthenticationProfileArrayInput interface {
	pulumi.Input

	ToAuthenticationProfileArrayOutput() AuthenticationProfileArrayOutput
	ToAuthenticationProfileArrayOutputWithContext(context.Context) AuthenticationProfileArrayOutput
}

type AuthenticationProfileArray []AuthenticationProfileInput

func (AuthenticationProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationProfile)(nil)).Elem()
}

func (i AuthenticationProfileArray) ToAuthenticationProfileArrayOutput() AuthenticationProfileArrayOutput {
	return i.ToAuthenticationProfileArrayOutputWithContext(context.Background())
}

func (i AuthenticationProfileArray) ToAuthenticationProfileArrayOutputWithContext(ctx context.Context) AuthenticationProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationProfileArrayOutput)
}

// AuthenticationProfileMapInput is an input type that accepts AuthenticationProfileMap and AuthenticationProfileMapOutput values.
// You can construct a concrete instance of `AuthenticationProfileMapInput` via:
//
//	AuthenticationProfileMap{ "key": AuthenticationProfileArgs{...} }
type AuthenticationProfileMapInput interface {
	pulumi.Input

	ToAuthenticationProfileMapOutput() AuthenticationProfileMapOutput
	ToAuthenticationProfileMapOutputWithContext(context.Context) AuthenticationProfileMapOutput
}

type AuthenticationProfileMap map[string]AuthenticationProfileInput

func (AuthenticationProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationProfile)(nil)).Elem()
}

func (i AuthenticationProfileMap) ToAuthenticationProfileMapOutput() AuthenticationProfileMapOutput {
	return i.ToAuthenticationProfileMapOutputWithContext(context.Background())
}

func (i AuthenticationProfileMap) ToAuthenticationProfileMapOutputWithContext(ctx context.Context) AuthenticationProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationProfileMapOutput)
}

type AuthenticationProfileOutput struct{ *pulumi.OutputState }

func (AuthenticationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationProfile)(nil)).Elem()
}

func (o AuthenticationProfileOutput) ToAuthenticationProfileOutput() AuthenticationProfileOutput {
	return o
}

func (o AuthenticationProfileOutput) ToAuthenticationProfileOutputWithContext(ctx context.Context) AuthenticationProfileOutput {
	return o
}

// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
func (o AuthenticationProfileOutput) AuthenticationProfileContent() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationProfile) pulumi.StringOutput { return v.AuthenticationProfileContent }).(pulumi.StringOutput)
}

// The name of the authentication profile.
func (o AuthenticationProfileOutput) AuthenticationProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationProfile) pulumi.StringOutput { return v.AuthenticationProfileName }).(pulumi.StringOutput)
}

type AuthenticationProfileArrayOutput struct{ *pulumi.OutputState }

func (AuthenticationProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationProfile)(nil)).Elem()
}

func (o AuthenticationProfileArrayOutput) ToAuthenticationProfileArrayOutput() AuthenticationProfileArrayOutput {
	return o
}

func (o AuthenticationProfileArrayOutput) ToAuthenticationProfileArrayOutputWithContext(ctx context.Context) AuthenticationProfileArrayOutput {
	return o
}

func (o AuthenticationProfileArrayOutput) Index(i pulumi.IntInput) AuthenticationProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthenticationProfile {
		return vs[0].([]*AuthenticationProfile)[vs[1].(int)]
	}).(AuthenticationProfileOutput)
}

type AuthenticationProfileMapOutput struct{ *pulumi.OutputState }

func (AuthenticationProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationProfile)(nil)).Elem()
}

func (o AuthenticationProfileMapOutput) ToAuthenticationProfileMapOutput() AuthenticationProfileMapOutput {
	return o
}

func (o AuthenticationProfileMapOutput) ToAuthenticationProfileMapOutputWithContext(ctx context.Context) AuthenticationProfileMapOutput {
	return o
}

func (o AuthenticationProfileMapOutput) MapIndex(k pulumi.StringInput) AuthenticationProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthenticationProfile {
		return vs[0].(map[string]*AuthenticationProfile)[vs[1].(string)]
	}).(AuthenticationProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationProfileInput)(nil)).Elem(), &AuthenticationProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationProfileArrayInput)(nil)).Elem(), AuthenticationProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationProfileMapInput)(nil)).Elem(), AuthenticationProfileMap{})
	pulumi.RegisterOutputType(AuthenticationProfileOutput{})
	pulumi.RegisterOutputType(AuthenticationProfileArrayOutput{})
	pulumi.RegisterOutputType(AuthenticationProfileMapOutput{})
}
