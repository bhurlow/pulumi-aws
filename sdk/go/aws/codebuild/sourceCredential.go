// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeBuild Source Credentials Resource.
//
// > **NOTE:**
// [Codebuild only allows a single credential per given server type in a given region](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_codebuild.GitHubSourceCredentials.html). Therefore, when you define `codebuild.SourceCredential`, `codebuild.Project` resource defined in the same module will use it.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/codebuild"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := codebuild.NewSourceCredential(ctx, "example", &codebuild.SourceCredentialArgs{
//				AuthType:   pulumi.String("PERSONAL_ACCESS_TOKEN"),
//				ServerType: pulumi.String("GITHUB"),
//				Token:      pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Bitbucket Server Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/codebuild"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := codebuild.NewSourceCredential(ctx, "example", &codebuild.SourceCredentialArgs{
//				AuthType:   pulumi.String("BASIC_AUTH"),
//				ServerType: pulumi.String("BITBUCKET"),
//				Token:      pulumi.String("example"),
//				UserName:   pulumi.String("test-user"),
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
// CodeBuild Source Credential can be imported using the CodeBuild Source Credential arn, e.g.,
//
// ```sh
//
//	$ pulumi import aws:codebuild/sourceCredential:SourceCredential example arn:aws:codebuild:us-west-2:123456789:token:github
//
// ```
type SourceCredential struct {
	pulumi.CustomResourceState

	// The ARN of Source Credential.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// The source provider used for this project.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token pulumi.StringOutput `pulumi:"token"`
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewSourceCredential registers a new resource with the given unique name, arguments, and options.
func NewSourceCredential(ctx *pulumi.Context,
	name string, args *SourceCredentialArgs, opts ...pulumi.ResourceOption) (*SourceCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthType == nil {
		return nil, errors.New("invalid value for required argument 'AuthType'")
	}
	if args.ServerType == nil {
		return nil, errors.New("invalid value for required argument 'ServerType'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	var resource SourceCredential
	err := ctx.RegisterResource("aws:codebuild/sourceCredential:SourceCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCredential gets an existing SourceCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCredentialState, opts ...pulumi.ResourceOption) (*SourceCredential, error) {
	var resource SourceCredential
	err := ctx.ReadResource("aws:codebuild/sourceCredential:SourceCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCredential resources.
type sourceCredentialState struct {
	// The ARN of Source Credential.
	Arn *string `pulumi:"arn"`
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType *string `pulumi:"authType"`
	// The source provider used for this project.
	ServerType *string `pulumi:"serverType"`
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token *string `pulumi:"token"`
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName *string `pulumi:"userName"`
}

type SourceCredentialState struct {
	// The ARN of Source Credential.
	Arn pulumi.StringPtrInput
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType pulumi.StringPtrInput
	// The source provider used for this project.
	ServerType pulumi.StringPtrInput
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token pulumi.StringPtrInput
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName pulumi.StringPtrInput
}

func (SourceCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCredentialState)(nil)).Elem()
}

type sourceCredentialArgs struct {
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType string `pulumi:"authType"`
	// The source provider used for this project.
	ServerType string `pulumi:"serverType"`
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token string `pulumi:"token"`
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a SourceCredential resource.
type SourceCredentialArgs struct {
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType pulumi.StringInput
	// The source provider used for this project.
	ServerType pulumi.StringInput
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token pulumi.StringInput
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName pulumi.StringPtrInput
}

func (SourceCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCredentialArgs)(nil)).Elem()
}

type SourceCredentialInput interface {
	pulumi.Input

	ToSourceCredentialOutput() SourceCredentialOutput
	ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput
}

func (*SourceCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCredential)(nil)).Elem()
}

func (i *SourceCredential) ToSourceCredentialOutput() SourceCredentialOutput {
	return i.ToSourceCredentialOutputWithContext(context.Background())
}

func (i *SourceCredential) ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCredentialOutput)
}

// SourceCredentialArrayInput is an input type that accepts SourceCredentialArray and SourceCredentialArrayOutput values.
// You can construct a concrete instance of `SourceCredentialArrayInput` via:
//
//	SourceCredentialArray{ SourceCredentialArgs{...} }
type SourceCredentialArrayInput interface {
	pulumi.Input

	ToSourceCredentialArrayOutput() SourceCredentialArrayOutput
	ToSourceCredentialArrayOutputWithContext(context.Context) SourceCredentialArrayOutput
}

type SourceCredentialArray []SourceCredentialInput

func (SourceCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCredential)(nil)).Elem()
}

func (i SourceCredentialArray) ToSourceCredentialArrayOutput() SourceCredentialArrayOutput {
	return i.ToSourceCredentialArrayOutputWithContext(context.Background())
}

func (i SourceCredentialArray) ToSourceCredentialArrayOutputWithContext(ctx context.Context) SourceCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCredentialArrayOutput)
}

// SourceCredentialMapInput is an input type that accepts SourceCredentialMap and SourceCredentialMapOutput values.
// You can construct a concrete instance of `SourceCredentialMapInput` via:
//
//	SourceCredentialMap{ "key": SourceCredentialArgs{...} }
type SourceCredentialMapInput interface {
	pulumi.Input

	ToSourceCredentialMapOutput() SourceCredentialMapOutput
	ToSourceCredentialMapOutputWithContext(context.Context) SourceCredentialMapOutput
}

type SourceCredentialMap map[string]SourceCredentialInput

func (SourceCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCredential)(nil)).Elem()
}

func (i SourceCredentialMap) ToSourceCredentialMapOutput() SourceCredentialMapOutput {
	return i.ToSourceCredentialMapOutputWithContext(context.Background())
}

func (i SourceCredentialMap) ToSourceCredentialMapOutputWithContext(ctx context.Context) SourceCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCredentialMapOutput)
}

type SourceCredentialOutput struct{ *pulumi.OutputState }

func (SourceCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCredential)(nil)).Elem()
}

func (o SourceCredentialOutput) ToSourceCredentialOutput() SourceCredentialOutput {
	return o
}

func (o SourceCredentialOutput) ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput {
	return o
}

// The ARN of Source Credential.
func (o SourceCredentialOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
func (o SourceCredentialOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// The source provider used for this project.
func (o SourceCredentialOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
func (o SourceCredentialOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
func (o SourceCredentialOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

type SourceCredentialArrayOutput struct{ *pulumi.OutputState }

func (SourceCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCredential)(nil)).Elem()
}

func (o SourceCredentialArrayOutput) ToSourceCredentialArrayOutput() SourceCredentialArrayOutput {
	return o
}

func (o SourceCredentialArrayOutput) ToSourceCredentialArrayOutputWithContext(ctx context.Context) SourceCredentialArrayOutput {
	return o
}

func (o SourceCredentialArrayOutput) Index(i pulumi.IntInput) SourceCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceCredential {
		return vs[0].([]*SourceCredential)[vs[1].(int)]
	}).(SourceCredentialOutput)
}

type SourceCredentialMapOutput struct{ *pulumi.OutputState }

func (SourceCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCredential)(nil)).Elem()
}

func (o SourceCredentialMapOutput) ToSourceCredentialMapOutput() SourceCredentialMapOutput {
	return o
}

func (o SourceCredentialMapOutput) ToSourceCredentialMapOutputWithContext(ctx context.Context) SourceCredentialMapOutput {
	return o
}

func (o SourceCredentialMapOutput) MapIndex(k pulumi.StringInput) SourceCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceCredential {
		return vs[0].(map[string]*SourceCredential)[vs[1].(string)]
	}).(SourceCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCredentialInput)(nil)).Elem(), &SourceCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCredentialArrayInput)(nil)).Elem(), SourceCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCredentialMapInput)(nil)).Elem(), SourceCredentialMap{})
	pulumi.RegisterOutputType(SourceCredentialOutput{})
	pulumi.RegisterOutputType(SourceCredentialArrayOutput{})
	pulumi.RegisterOutputType(SourceCredentialMapOutput{})
}
