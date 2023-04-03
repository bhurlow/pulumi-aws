// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IoT role alias.
//
// ## Import
//
// IOT Role Alias can be imported via the alias, e.g.,
//
// ```sh
//
//	$ pulumi import aws:iot/roleAlias:RoleAlias example myalias
//
// ```
type RoleAlias struct {
	pulumi.CustomResourceState

	// The name of the role alias.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The ARN assigned by AWS to this role alias.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
	CredentialDuration pulumi.IntPtrOutput `pulumi:"credentialDuration"`
	// The identity of the role to which the alias refers.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewRoleAlias registers a new resource with the given unique name, arguments, and options.
func NewRoleAlias(ctx *pulumi.Context,
	name string, args *RoleAliasArgs, opts ...pulumi.ResourceOption) (*RoleAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource RoleAlias
	err := ctx.RegisterResource("aws:iot/roleAlias:RoleAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAlias gets an existing RoleAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAliasState, opts ...pulumi.ResourceOption) (*RoleAlias, error) {
	var resource RoleAlias
	err := ctx.ReadResource("aws:iot/roleAlias:RoleAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAlias resources.
type roleAliasState struct {
	// The name of the role alias.
	Alias *string `pulumi:"alias"`
	// The ARN assigned by AWS to this role alias.
	Arn *string `pulumi:"arn"`
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
	CredentialDuration *int `pulumi:"credentialDuration"`
	// The identity of the role to which the alias refers.
	RoleArn *string `pulumi:"roleArn"`
}

type RoleAliasState struct {
	// The name of the role alias.
	Alias pulumi.StringPtrInput
	// The ARN assigned by AWS to this role alias.
	Arn pulumi.StringPtrInput
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
	CredentialDuration pulumi.IntPtrInput
	// The identity of the role to which the alias refers.
	RoleArn pulumi.StringPtrInput
}

func (RoleAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAliasState)(nil)).Elem()
}

type roleAliasArgs struct {
	// The name of the role alias.
	Alias string `pulumi:"alias"`
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
	CredentialDuration *int `pulumi:"credentialDuration"`
	// The identity of the role to which the alias refers.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a RoleAlias resource.
type RoleAliasArgs struct {
	// The name of the role alias.
	Alias pulumi.StringInput
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
	CredentialDuration pulumi.IntPtrInput
	// The identity of the role to which the alias refers.
	RoleArn pulumi.StringInput
}

func (RoleAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAliasArgs)(nil)).Elem()
}

type RoleAliasInput interface {
	pulumi.Input

	ToRoleAliasOutput() RoleAliasOutput
	ToRoleAliasOutputWithContext(ctx context.Context) RoleAliasOutput
}

func (*RoleAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAlias)(nil)).Elem()
}

func (i *RoleAlias) ToRoleAliasOutput() RoleAliasOutput {
	return i.ToRoleAliasOutputWithContext(context.Background())
}

func (i *RoleAlias) ToRoleAliasOutputWithContext(ctx context.Context) RoleAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAliasOutput)
}

// RoleAliasArrayInput is an input type that accepts RoleAliasArray and RoleAliasArrayOutput values.
// You can construct a concrete instance of `RoleAliasArrayInput` via:
//
//	RoleAliasArray{ RoleAliasArgs{...} }
type RoleAliasArrayInput interface {
	pulumi.Input

	ToRoleAliasArrayOutput() RoleAliasArrayOutput
	ToRoleAliasArrayOutputWithContext(context.Context) RoleAliasArrayOutput
}

type RoleAliasArray []RoleAliasInput

func (RoleAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAlias)(nil)).Elem()
}

func (i RoleAliasArray) ToRoleAliasArrayOutput() RoleAliasArrayOutput {
	return i.ToRoleAliasArrayOutputWithContext(context.Background())
}

func (i RoleAliasArray) ToRoleAliasArrayOutputWithContext(ctx context.Context) RoleAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAliasArrayOutput)
}

// RoleAliasMapInput is an input type that accepts RoleAliasMap and RoleAliasMapOutput values.
// You can construct a concrete instance of `RoleAliasMapInput` via:
//
//	RoleAliasMap{ "key": RoleAliasArgs{...} }
type RoleAliasMapInput interface {
	pulumi.Input

	ToRoleAliasMapOutput() RoleAliasMapOutput
	ToRoleAliasMapOutputWithContext(context.Context) RoleAliasMapOutput
}

type RoleAliasMap map[string]RoleAliasInput

func (RoleAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAlias)(nil)).Elem()
}

func (i RoleAliasMap) ToRoleAliasMapOutput() RoleAliasMapOutput {
	return i.ToRoleAliasMapOutputWithContext(context.Background())
}

func (i RoleAliasMap) ToRoleAliasMapOutputWithContext(ctx context.Context) RoleAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAliasMapOutput)
}

type RoleAliasOutput struct{ *pulumi.OutputState }

func (RoleAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAlias)(nil)).Elem()
}

func (o RoleAliasOutput) ToRoleAliasOutput() RoleAliasOutput {
	return o
}

func (o RoleAliasOutput) ToRoleAliasOutputWithContext(ctx context.Context) RoleAliasOutput {
	return o
}

// The name of the role alias.
func (o RoleAliasOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAlias) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// The ARN assigned by AWS to this role alias.
func (o RoleAliasOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAlias) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 43200 seconds (12 hours).
func (o RoleAliasOutput) CredentialDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RoleAlias) pulumi.IntPtrOutput { return v.CredentialDuration }).(pulumi.IntPtrOutput)
}

// The identity of the role to which the alias refers.
func (o RoleAliasOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAlias) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

type RoleAliasArrayOutput struct{ *pulumi.OutputState }

func (RoleAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAlias)(nil)).Elem()
}

func (o RoleAliasArrayOutput) ToRoleAliasArrayOutput() RoleAliasArrayOutput {
	return o
}

func (o RoleAliasArrayOutput) ToRoleAliasArrayOutputWithContext(ctx context.Context) RoleAliasArrayOutput {
	return o
}

func (o RoleAliasArrayOutput) Index(i pulumi.IntInput) RoleAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleAlias {
		return vs[0].([]*RoleAlias)[vs[1].(int)]
	}).(RoleAliasOutput)
}

type RoleAliasMapOutput struct{ *pulumi.OutputState }

func (RoleAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAlias)(nil)).Elem()
}

func (o RoleAliasMapOutput) ToRoleAliasMapOutput() RoleAliasMapOutput {
	return o
}

func (o RoleAliasMapOutput) ToRoleAliasMapOutputWithContext(ctx context.Context) RoleAliasMapOutput {
	return o
}

func (o RoleAliasMapOutput) MapIndex(k pulumi.StringInput) RoleAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleAlias {
		return vs[0].(map[string]*RoleAlias)[vs[1].(string)]
	}).(RoleAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAliasInput)(nil)).Elem(), &RoleAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAliasArrayInput)(nil)).Elem(), RoleAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAliasMapInput)(nil)).Elem(), RoleAliasMap{})
	pulumi.RegisterOutputType(RoleAliasOutput{})
	pulumi.RegisterOutputType(RoleAliasArrayOutput{})
	pulumi.RegisterOutputType(RoleAliasMapOutput{})
}
