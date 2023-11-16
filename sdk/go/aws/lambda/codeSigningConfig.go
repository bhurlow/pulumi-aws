// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Lambda Code Signing Config resource. A code signing configuration defines a list of allowed signing profiles and defines the code-signing validation policy (action to be taken if deployment validation checks fail).
//
// For information about Lambda code signing configurations and how to use them, see [configuring code signing for Lambda functions](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := lambda.NewCodeSigningConfig(ctx, "newCsc", &lambda.CodeSigningConfigArgs{
//				AllowedPublishers: &lambda.CodeSigningConfigAllowedPublishersArgs{
//					SigningProfileVersionArns: pulumi.StringArray{
//						aws_signer_signing_profile.Example1.Arn,
//						aws_signer_signing_profile.Example2.Arn,
//					},
//				},
//				Policies: &lambda.CodeSigningConfigPoliciesArgs{
//					UntrustedArtifactOnDeployment: pulumi.String("Warn"),
//				},
//				Description: pulumi.String("My awesome code signing config."),
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
// Using `pulumi import`, import Code Signing Configs using their ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:lambda/codeSigningConfig:CodeSigningConfig imported_csc arn:aws:lambda:us-west-2:123456789012:code-signing-config:csc-0f6c334abcdea4d8b
//
// ```
type CodeSigningConfig struct {
	pulumi.CustomResourceState

	// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
	AllowedPublishers CodeSigningConfigAllowedPublishersOutput `pulumi:"allowedPublishers"`
	// The Amazon Resource Name (ARN) of the code signing configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Unique identifier for the code signing configuration.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// Descriptive name for this code signing configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The date and time that the code signing configuration was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
	Policies CodeSigningConfigPoliciesOutput `pulumi:"policies"`
}

// NewCodeSigningConfig registers a new resource with the given unique name, arguments, and options.
func NewCodeSigningConfig(ctx *pulumi.Context,
	name string, args *CodeSigningConfigArgs, opts ...pulumi.ResourceOption) (*CodeSigningConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedPublishers == nil {
		return nil, errors.New("invalid value for required argument 'AllowedPublishers'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CodeSigningConfig
	err := ctx.RegisterResource("aws:lambda/codeSigningConfig:CodeSigningConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCodeSigningConfig gets an existing CodeSigningConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCodeSigningConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodeSigningConfigState, opts ...pulumi.ResourceOption) (*CodeSigningConfig, error) {
	var resource CodeSigningConfig
	err := ctx.ReadResource("aws:lambda/codeSigningConfig:CodeSigningConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CodeSigningConfig resources.
type codeSigningConfigState struct {
	// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
	AllowedPublishers *CodeSigningConfigAllowedPublishers `pulumi:"allowedPublishers"`
	// The Amazon Resource Name (ARN) of the code signing configuration.
	Arn *string `pulumi:"arn"`
	// Unique identifier for the code signing configuration.
	ConfigId *string `pulumi:"configId"`
	// Descriptive name for this code signing configuration.
	Description *string `pulumi:"description"`
	// The date and time that the code signing configuration was last modified.
	LastModified *string `pulumi:"lastModified"`
	// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
	Policies *CodeSigningConfigPolicies `pulumi:"policies"`
}

type CodeSigningConfigState struct {
	// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
	AllowedPublishers CodeSigningConfigAllowedPublishersPtrInput
	// The Amazon Resource Name (ARN) of the code signing configuration.
	Arn pulumi.StringPtrInput
	// Unique identifier for the code signing configuration.
	ConfigId pulumi.StringPtrInput
	// Descriptive name for this code signing configuration.
	Description pulumi.StringPtrInput
	// The date and time that the code signing configuration was last modified.
	LastModified pulumi.StringPtrInput
	// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
	Policies CodeSigningConfigPoliciesPtrInput
}

func (CodeSigningConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*codeSigningConfigState)(nil)).Elem()
}

type codeSigningConfigArgs struct {
	// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
	AllowedPublishers CodeSigningConfigAllowedPublishers `pulumi:"allowedPublishers"`
	// Descriptive name for this code signing configuration.
	Description *string `pulumi:"description"`
	// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
	Policies *CodeSigningConfigPolicies `pulumi:"policies"`
}

// The set of arguments for constructing a CodeSigningConfig resource.
type CodeSigningConfigArgs struct {
	// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
	AllowedPublishers CodeSigningConfigAllowedPublishersInput
	// Descriptive name for this code signing configuration.
	Description pulumi.StringPtrInput
	// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
	Policies CodeSigningConfigPoliciesPtrInput
}

func (CodeSigningConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codeSigningConfigArgs)(nil)).Elem()
}

type CodeSigningConfigInput interface {
	pulumi.Input

	ToCodeSigningConfigOutput() CodeSigningConfigOutput
	ToCodeSigningConfigOutputWithContext(ctx context.Context) CodeSigningConfigOutput
}

func (*CodeSigningConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeSigningConfig)(nil)).Elem()
}

func (i *CodeSigningConfig) ToCodeSigningConfigOutput() CodeSigningConfigOutput {
	return i.ToCodeSigningConfigOutputWithContext(context.Background())
}

func (i *CodeSigningConfig) ToCodeSigningConfigOutputWithContext(ctx context.Context) CodeSigningConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeSigningConfigOutput)
}

// CodeSigningConfigArrayInput is an input type that accepts CodeSigningConfigArray and CodeSigningConfigArrayOutput values.
// You can construct a concrete instance of `CodeSigningConfigArrayInput` via:
//
//	CodeSigningConfigArray{ CodeSigningConfigArgs{...} }
type CodeSigningConfigArrayInput interface {
	pulumi.Input

	ToCodeSigningConfigArrayOutput() CodeSigningConfigArrayOutput
	ToCodeSigningConfigArrayOutputWithContext(context.Context) CodeSigningConfigArrayOutput
}

type CodeSigningConfigArray []CodeSigningConfigInput

func (CodeSigningConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CodeSigningConfig)(nil)).Elem()
}

func (i CodeSigningConfigArray) ToCodeSigningConfigArrayOutput() CodeSigningConfigArrayOutput {
	return i.ToCodeSigningConfigArrayOutputWithContext(context.Background())
}

func (i CodeSigningConfigArray) ToCodeSigningConfigArrayOutputWithContext(ctx context.Context) CodeSigningConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeSigningConfigArrayOutput)
}

// CodeSigningConfigMapInput is an input type that accepts CodeSigningConfigMap and CodeSigningConfigMapOutput values.
// You can construct a concrete instance of `CodeSigningConfigMapInput` via:
//
//	CodeSigningConfigMap{ "key": CodeSigningConfigArgs{...} }
type CodeSigningConfigMapInput interface {
	pulumi.Input

	ToCodeSigningConfigMapOutput() CodeSigningConfigMapOutput
	ToCodeSigningConfigMapOutputWithContext(context.Context) CodeSigningConfigMapOutput
}

type CodeSigningConfigMap map[string]CodeSigningConfigInput

func (CodeSigningConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CodeSigningConfig)(nil)).Elem()
}

func (i CodeSigningConfigMap) ToCodeSigningConfigMapOutput() CodeSigningConfigMapOutput {
	return i.ToCodeSigningConfigMapOutputWithContext(context.Background())
}

func (i CodeSigningConfigMap) ToCodeSigningConfigMapOutputWithContext(ctx context.Context) CodeSigningConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeSigningConfigMapOutput)
}

type CodeSigningConfigOutput struct{ *pulumi.OutputState }

func (CodeSigningConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeSigningConfig)(nil)).Elem()
}

func (o CodeSigningConfigOutput) ToCodeSigningConfigOutput() CodeSigningConfigOutput {
	return o
}

func (o CodeSigningConfigOutput) ToCodeSigningConfigOutputWithContext(ctx context.Context) CodeSigningConfigOutput {
	return o
}

// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
func (o CodeSigningConfigOutput) AllowedPublishers() CodeSigningConfigAllowedPublishersOutput {
	return o.ApplyT(func(v *CodeSigningConfig) CodeSigningConfigAllowedPublishersOutput { return v.AllowedPublishers }).(CodeSigningConfigAllowedPublishersOutput)
}

// The Amazon Resource Name (ARN) of the code signing configuration.
func (o CodeSigningConfigOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeSigningConfig) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Unique identifier for the code signing configuration.
func (o CodeSigningConfigOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeSigningConfig) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// Descriptive name for this code signing configuration.
func (o CodeSigningConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeSigningConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The date and time that the code signing configuration was last modified.
func (o CodeSigningConfigOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeSigningConfig) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
func (o CodeSigningConfigOutput) Policies() CodeSigningConfigPoliciesOutput {
	return o.ApplyT(func(v *CodeSigningConfig) CodeSigningConfigPoliciesOutput { return v.Policies }).(CodeSigningConfigPoliciesOutput)
}

type CodeSigningConfigArrayOutput struct{ *pulumi.OutputState }

func (CodeSigningConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CodeSigningConfig)(nil)).Elem()
}

func (o CodeSigningConfigArrayOutput) ToCodeSigningConfigArrayOutput() CodeSigningConfigArrayOutput {
	return o
}

func (o CodeSigningConfigArrayOutput) ToCodeSigningConfigArrayOutputWithContext(ctx context.Context) CodeSigningConfigArrayOutput {
	return o
}

func (o CodeSigningConfigArrayOutput) Index(i pulumi.IntInput) CodeSigningConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CodeSigningConfig {
		return vs[0].([]*CodeSigningConfig)[vs[1].(int)]
	}).(CodeSigningConfigOutput)
}

type CodeSigningConfigMapOutput struct{ *pulumi.OutputState }

func (CodeSigningConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CodeSigningConfig)(nil)).Elem()
}

func (o CodeSigningConfigMapOutput) ToCodeSigningConfigMapOutput() CodeSigningConfigMapOutput {
	return o
}

func (o CodeSigningConfigMapOutput) ToCodeSigningConfigMapOutputWithContext(ctx context.Context) CodeSigningConfigMapOutput {
	return o
}

func (o CodeSigningConfigMapOutput) MapIndex(k pulumi.StringInput) CodeSigningConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CodeSigningConfig {
		return vs[0].(map[string]*CodeSigningConfig)[vs[1].(string)]
	}).(CodeSigningConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CodeSigningConfigInput)(nil)).Elem(), &CodeSigningConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodeSigningConfigArrayInput)(nil)).Elem(), CodeSigningConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodeSigningConfigMapInput)(nil)).Elem(), CodeSigningConfigMap{})
	pulumi.RegisterOutputType(CodeSigningConfigOutput{})
	pulumi.RegisterOutputType(CodeSigningConfigArrayOutput{})
	pulumi.RegisterOutputType(CodeSigningConfigMapOutput{})
}
