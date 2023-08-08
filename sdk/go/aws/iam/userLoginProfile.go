// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an IAM User Login Profile with limited support for password creation during this provider resource creation. Uses PGP to encrypt the password for safe transport to the user. PGP keys can be obtained from Keybase.
//
// > To reset an IAM User login password via this provider, you can use delete and recreate this resource or change any of the arguments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleUser, err := iam.NewUser(ctx, "exampleUser", &iam.UserArgs{
//				Path:         pulumi.String("/"),
//				ForceDestroy: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleUserLoginProfile, err := iam.NewUserLoginProfile(ctx, "exampleUserLoginProfile", &iam.UserLoginProfileArgs{
//				User:   exampleUser.Name,
//				PgpKey: pulumi.String("keybase:some_person_that_exists"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("password", exampleUserLoginProfile.EncryptedPassword)
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
//	to = aws_iam_user_login_profile.example
//
//	id = "myusername" } Using `pulumi import`, import IAM User Login Profiles without password information via the IAM User name. For exampleconsole % pulumi import aws_iam_user_login_profile.example myusername Since TODO has no method to read the PGP or password information during import, use the TODO resource `lifecycle` configuration block `ignore_changes` argument to ignore them (unless you want to recreate a password). For exampleterraform resource "aws_iam_user_login_profile" "example" {
//
// # ... other configuration ...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	password_length,
//
//	password_reset_required,
//
//	pgp_key,
//
//	]
//
//	} }
type UserLoginProfile struct {
	pulumi.CustomResourceState

	// The encrypted password, base64 encoded. Only available if password was handled on resource creation, not import.
	EncryptedPassword pulumi.StringOutput `pulumi:"encryptedPassword"`
	// The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
	KeyFingerprint pulumi.StringOutput `pulumi:"keyFingerprint"`
	// The plain text password, only available when `pgpKey` is not provided.
	Password pulumi.StringOutput `pulumi:"password"`
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is `20`.
	PasswordLength pulumi.IntPtrOutput `pulumi:"passwordLength"`
	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
	PasswordResetRequired pulumi.BoolOutput `pulumi:"passwordResetRequired"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey pulumi.StringPtrOutput `pulumi:"pgpKey"`
	// The IAM user's name.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewUserLoginProfile registers a new resource with the given unique name, arguments, and options.
func NewUserLoginProfile(ctx *pulumi.Context,
	name string, args *UserLoginProfileArgs, opts ...pulumi.ResourceOption) (*UserLoginProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserLoginProfile
	err := ctx.RegisterResource("aws:iam/userLoginProfile:UserLoginProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserLoginProfile gets an existing UserLoginProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserLoginProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserLoginProfileState, opts ...pulumi.ResourceOption) (*UserLoginProfile, error) {
	var resource UserLoginProfile
	err := ctx.ReadResource("aws:iam/userLoginProfile:UserLoginProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserLoginProfile resources.
type userLoginProfileState struct {
	// The encrypted password, base64 encoded. Only available if password was handled on resource creation, not import.
	EncryptedPassword *string `pulumi:"encryptedPassword"`
	// The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
	KeyFingerprint *string `pulumi:"keyFingerprint"`
	// The plain text password, only available when `pgpKey` is not provided.
	Password *string `pulumi:"password"`
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is `20`.
	PasswordLength *int `pulumi:"passwordLength"`
	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
	PasswordResetRequired *bool `pulumi:"passwordResetRequired"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey *string `pulumi:"pgpKey"`
	// The IAM user's name.
	User *string `pulumi:"user"`
}

type UserLoginProfileState struct {
	// The encrypted password, base64 encoded. Only available if password was handled on resource creation, not import.
	EncryptedPassword pulumi.StringPtrInput
	// The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
	KeyFingerprint pulumi.StringPtrInput
	// The plain text password, only available when `pgpKey` is not provided.
	Password pulumi.StringPtrInput
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is `20`.
	PasswordLength pulumi.IntPtrInput
	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
	PasswordResetRequired pulumi.BoolPtrInput
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey pulumi.StringPtrInput
	// The IAM user's name.
	User pulumi.StringPtrInput
}

func (UserLoginProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*userLoginProfileState)(nil)).Elem()
}

type userLoginProfileArgs struct {
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is `20`.
	PasswordLength *int `pulumi:"passwordLength"`
	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
	PasswordResetRequired *bool `pulumi:"passwordResetRequired"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey *string `pulumi:"pgpKey"`
	// The IAM user's name.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a UserLoginProfile resource.
type UserLoginProfileArgs struct {
	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is `20`.
	PasswordLength pulumi.IntPtrInput
	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
	PasswordResetRequired pulumi.BoolPtrInput
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
	PgpKey pulumi.StringPtrInput
	// The IAM user's name.
	User pulumi.StringInput
}

func (UserLoginProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userLoginProfileArgs)(nil)).Elem()
}

type UserLoginProfileInput interface {
	pulumi.Input

	ToUserLoginProfileOutput() UserLoginProfileOutput
	ToUserLoginProfileOutputWithContext(ctx context.Context) UserLoginProfileOutput
}

func (*UserLoginProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**UserLoginProfile)(nil)).Elem()
}

func (i *UserLoginProfile) ToUserLoginProfileOutput() UserLoginProfileOutput {
	return i.ToUserLoginProfileOutputWithContext(context.Background())
}

func (i *UserLoginProfile) ToUserLoginProfileOutputWithContext(ctx context.Context) UserLoginProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLoginProfileOutput)
}

// UserLoginProfileArrayInput is an input type that accepts UserLoginProfileArray and UserLoginProfileArrayOutput values.
// You can construct a concrete instance of `UserLoginProfileArrayInput` via:
//
//	UserLoginProfileArray{ UserLoginProfileArgs{...} }
type UserLoginProfileArrayInput interface {
	pulumi.Input

	ToUserLoginProfileArrayOutput() UserLoginProfileArrayOutput
	ToUserLoginProfileArrayOutputWithContext(context.Context) UserLoginProfileArrayOutput
}

type UserLoginProfileArray []UserLoginProfileInput

func (UserLoginProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserLoginProfile)(nil)).Elem()
}

func (i UserLoginProfileArray) ToUserLoginProfileArrayOutput() UserLoginProfileArrayOutput {
	return i.ToUserLoginProfileArrayOutputWithContext(context.Background())
}

func (i UserLoginProfileArray) ToUserLoginProfileArrayOutputWithContext(ctx context.Context) UserLoginProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLoginProfileArrayOutput)
}

// UserLoginProfileMapInput is an input type that accepts UserLoginProfileMap and UserLoginProfileMapOutput values.
// You can construct a concrete instance of `UserLoginProfileMapInput` via:
//
//	UserLoginProfileMap{ "key": UserLoginProfileArgs{...} }
type UserLoginProfileMapInput interface {
	pulumi.Input

	ToUserLoginProfileMapOutput() UserLoginProfileMapOutput
	ToUserLoginProfileMapOutputWithContext(context.Context) UserLoginProfileMapOutput
}

type UserLoginProfileMap map[string]UserLoginProfileInput

func (UserLoginProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserLoginProfile)(nil)).Elem()
}

func (i UserLoginProfileMap) ToUserLoginProfileMapOutput() UserLoginProfileMapOutput {
	return i.ToUserLoginProfileMapOutputWithContext(context.Background())
}

func (i UserLoginProfileMap) ToUserLoginProfileMapOutputWithContext(ctx context.Context) UserLoginProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLoginProfileMapOutput)
}

type UserLoginProfileOutput struct{ *pulumi.OutputState }

func (UserLoginProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserLoginProfile)(nil)).Elem()
}

func (o UserLoginProfileOutput) ToUserLoginProfileOutput() UserLoginProfileOutput {
	return o
}

func (o UserLoginProfileOutput) ToUserLoginProfileOutputWithContext(ctx context.Context) UserLoginProfileOutput {
	return o
}

// The encrypted password, base64 encoded. Only available if password was handled on resource creation, not import.
func (o UserLoginProfileOutput) EncryptedPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *UserLoginProfile) pulumi.StringOutput { return v.EncryptedPassword }).(pulumi.StringOutput)
}

// The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
func (o UserLoginProfileOutput) KeyFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *UserLoginProfile) pulumi.StringOutput { return v.KeyFingerprint }).(pulumi.StringOutput)
}

// The plain text password, only available when `pgpKey` is not provided.
func (o UserLoginProfileOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *UserLoginProfile) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is `20`.
func (o UserLoginProfileOutput) PasswordLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserLoginProfile) pulumi.IntPtrOutput { return v.PasswordLength }).(pulumi.IntPtrOutput)
}

// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
func (o UserLoginProfileOutput) PasswordResetRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserLoginProfile) pulumi.BoolOutput { return v.PasswordResetRequired }).(pulumi.BoolOutput)
}

// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
func (o UserLoginProfileOutput) PgpKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserLoginProfile) pulumi.StringPtrOutput { return v.PgpKey }).(pulumi.StringPtrOutput)
}

// The IAM user's name.
func (o UserLoginProfileOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *UserLoginProfile) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type UserLoginProfileArrayOutput struct{ *pulumi.OutputState }

func (UserLoginProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserLoginProfile)(nil)).Elem()
}

func (o UserLoginProfileArrayOutput) ToUserLoginProfileArrayOutput() UserLoginProfileArrayOutput {
	return o
}

func (o UserLoginProfileArrayOutput) ToUserLoginProfileArrayOutputWithContext(ctx context.Context) UserLoginProfileArrayOutput {
	return o
}

func (o UserLoginProfileArrayOutput) Index(i pulumi.IntInput) UserLoginProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserLoginProfile {
		return vs[0].([]*UserLoginProfile)[vs[1].(int)]
	}).(UserLoginProfileOutput)
}

type UserLoginProfileMapOutput struct{ *pulumi.OutputState }

func (UserLoginProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserLoginProfile)(nil)).Elem()
}

func (o UserLoginProfileMapOutput) ToUserLoginProfileMapOutput() UserLoginProfileMapOutput {
	return o
}

func (o UserLoginProfileMapOutput) ToUserLoginProfileMapOutputWithContext(ctx context.Context) UserLoginProfileMapOutput {
	return o
}

func (o UserLoginProfileMapOutput) MapIndex(k pulumi.StringInput) UserLoginProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserLoginProfile {
		return vs[0].(map[string]*UserLoginProfile)[vs[1].(string)]
	}).(UserLoginProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserLoginProfileInput)(nil)).Elem(), &UserLoginProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserLoginProfileArrayInput)(nil)).Elem(), UserLoginProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserLoginProfileMapInput)(nil)).Elem(), UserLoginProfileMap{})
	pulumi.RegisterOutputType(UserLoginProfileOutput{})
	pulumi.RegisterOutputType(UserLoginProfileArrayOutput{})
	pulumi.RegisterOutputType(UserLoginProfileMapOutput{})
}
