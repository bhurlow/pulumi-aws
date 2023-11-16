// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SageMaker endpoint configuration resource.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sagemaker.NewEndpointConfiguration(ctx, "ec", &sagemaker.EndpointConfigurationArgs{
//				ProductionVariants: sagemaker.EndpointConfigurationProductionVariantArray{
//					&sagemaker.EndpointConfigurationProductionVariantArgs{
//						VariantName:          pulumi.String("variant-1"),
//						ModelName:            pulumi.Any(aws_sagemaker_model.M.Name),
//						InitialInstanceCount: pulumi.Int(1),
//						InstanceType:         pulumi.String("ml.t2.medium"),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("foo"),
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
// Using `pulumi import`, import endpoint configurations using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:sagemaker/endpointConfiguration:EndpointConfiguration test_endpoint_config endpoint-config-foo
//
// ```
type EndpointConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig EndpointConfigurationAsyncInferenceConfigPtrOutput `pulumi:"asyncInferenceConfig"`
	// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
	DataCaptureConfig EndpointConfigurationDataCaptureConfigPtrOutput `pulumi:"dataCaptureConfig"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrOutput `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayOutput `pulumi:"productionVariants"`
	// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
	ShadowProductionVariants EndpointConfigurationShadowProductionVariantArrayOutput `pulumi:"shadowProductionVariants"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewEndpointConfiguration registers a new resource with the given unique name, arguments, and options.
func NewEndpointConfiguration(ctx *pulumi.Context,
	name string, args *EndpointConfigurationArgs, opts ...pulumi.ResourceOption) (*EndpointConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductionVariants == nil {
		return nil, errors.New("invalid value for required argument 'ProductionVariants'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointConfiguration
	err := ctx.RegisterResource("aws:sagemaker/endpointConfiguration:EndpointConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointConfiguration gets an existing EndpointConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointConfigurationState, opts ...pulumi.ResourceOption) (*EndpointConfiguration, error) {
	var resource EndpointConfiguration
	err := ctx.ReadResource("aws:sagemaker/endpointConfiguration:EndpointConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointConfiguration resources.
type endpointConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn *string `pulumi:"arn"`
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig *EndpointConfigurationAsyncInferenceConfig `pulumi:"asyncInferenceConfig"`
	// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
	DataCaptureConfig *EndpointConfigurationDataCaptureConfig `pulumi:"dataCaptureConfig"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
	ProductionVariants []EndpointConfigurationProductionVariant `pulumi:"productionVariants"`
	// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
	ShadowProductionVariants []EndpointConfigurationShadowProductionVariant `pulumi:"shadowProductionVariants"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type EndpointConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn pulumi.StringPtrInput
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig EndpointConfigurationAsyncInferenceConfigPtrInput
	// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
	DataCaptureConfig EndpointConfigurationDataCaptureConfigPtrInput
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrInput
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayInput
	// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
	ShadowProductionVariants EndpointConfigurationShadowProductionVariantArrayInput
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (EndpointConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointConfigurationState)(nil)).Elem()
}

type endpointConfigurationArgs struct {
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig *EndpointConfigurationAsyncInferenceConfig `pulumi:"asyncInferenceConfig"`
	// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
	DataCaptureConfig *EndpointConfigurationDataCaptureConfig `pulumi:"dataCaptureConfig"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
	ProductionVariants []EndpointConfigurationProductionVariant `pulumi:"productionVariants"`
	// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
	ShadowProductionVariants []EndpointConfigurationShadowProductionVariant `pulumi:"shadowProductionVariants"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EndpointConfiguration resource.
type EndpointConfigurationArgs struct {
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig EndpointConfigurationAsyncInferenceConfigPtrInput
	// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
	DataCaptureConfig EndpointConfigurationDataCaptureConfigPtrInput
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrInput
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayInput
	// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
	ShadowProductionVariants EndpointConfigurationShadowProductionVariantArrayInput
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (EndpointConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointConfigurationArgs)(nil)).Elem()
}

type EndpointConfigurationInput interface {
	pulumi.Input

	ToEndpointConfigurationOutput() EndpointConfigurationOutput
	ToEndpointConfigurationOutputWithContext(ctx context.Context) EndpointConfigurationOutput
}

func (*EndpointConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointConfiguration)(nil)).Elem()
}

func (i *EndpointConfiguration) ToEndpointConfigurationOutput() EndpointConfigurationOutput {
	return i.ToEndpointConfigurationOutputWithContext(context.Background())
}

func (i *EndpointConfiguration) ToEndpointConfigurationOutputWithContext(ctx context.Context) EndpointConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationOutput)
}

// EndpointConfigurationArrayInput is an input type that accepts EndpointConfigurationArray and EndpointConfigurationArrayOutput values.
// You can construct a concrete instance of `EndpointConfigurationArrayInput` via:
//
//	EndpointConfigurationArray{ EndpointConfigurationArgs{...} }
type EndpointConfigurationArrayInput interface {
	pulumi.Input

	ToEndpointConfigurationArrayOutput() EndpointConfigurationArrayOutput
	ToEndpointConfigurationArrayOutputWithContext(context.Context) EndpointConfigurationArrayOutput
}

type EndpointConfigurationArray []EndpointConfigurationInput

func (EndpointConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointConfiguration)(nil)).Elem()
}

func (i EndpointConfigurationArray) ToEndpointConfigurationArrayOutput() EndpointConfigurationArrayOutput {
	return i.ToEndpointConfigurationArrayOutputWithContext(context.Background())
}

func (i EndpointConfigurationArray) ToEndpointConfigurationArrayOutputWithContext(ctx context.Context) EndpointConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationArrayOutput)
}

// EndpointConfigurationMapInput is an input type that accepts EndpointConfigurationMap and EndpointConfigurationMapOutput values.
// You can construct a concrete instance of `EndpointConfigurationMapInput` via:
//
//	EndpointConfigurationMap{ "key": EndpointConfigurationArgs{...} }
type EndpointConfigurationMapInput interface {
	pulumi.Input

	ToEndpointConfigurationMapOutput() EndpointConfigurationMapOutput
	ToEndpointConfigurationMapOutputWithContext(context.Context) EndpointConfigurationMapOutput
}

type EndpointConfigurationMap map[string]EndpointConfigurationInput

func (EndpointConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointConfiguration)(nil)).Elem()
}

func (i EndpointConfigurationMap) ToEndpointConfigurationMapOutput() EndpointConfigurationMapOutput {
	return i.ToEndpointConfigurationMapOutputWithContext(context.Background())
}

func (i EndpointConfigurationMap) ToEndpointConfigurationMapOutputWithContext(ctx context.Context) EndpointConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationMapOutput)
}

type EndpointConfigurationOutput struct{ *pulumi.OutputState }

func (EndpointConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointConfiguration)(nil)).Elem()
}

func (o EndpointConfigurationOutput) ToEndpointConfigurationOutput() EndpointConfigurationOutput {
	return o
}

func (o EndpointConfigurationOutput) ToEndpointConfigurationOutputWithContext(ctx context.Context) EndpointConfigurationOutput {
	return o
}

// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
func (o EndpointConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies configuration for how an endpoint performs asynchronous inference.
func (o EndpointConfigurationOutput) AsyncInferenceConfig() EndpointConfigurationAsyncInferenceConfigPtrOutput {
	return o.ApplyT(func(v *EndpointConfiguration) EndpointConfigurationAsyncInferenceConfigPtrOutput {
		return v.AsyncInferenceConfig
	}).(EndpointConfigurationAsyncInferenceConfigPtrOutput)
}

// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
func (o EndpointConfigurationOutput) DataCaptureConfig() EndpointConfigurationDataCaptureConfigPtrOutput {
	return o.ApplyT(func(v *EndpointConfiguration) EndpointConfigurationDataCaptureConfigPtrOutput {
		return v.DataCaptureConfig
	}).(EndpointConfigurationDataCaptureConfigPtrOutput)
}

// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
func (o EndpointConfigurationOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointConfiguration) pulumi.StringPtrOutput { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
func (o EndpointConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
func (o EndpointConfigurationOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointConfiguration) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
func (o EndpointConfigurationOutput) ProductionVariants() EndpointConfigurationProductionVariantArrayOutput {
	return o.ApplyT(func(v *EndpointConfiguration) EndpointConfigurationProductionVariantArrayOutput {
		return v.ProductionVariants
	}).(EndpointConfigurationProductionVariantArrayOutput)
}

// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
func (o EndpointConfigurationOutput) ShadowProductionVariants() EndpointConfigurationShadowProductionVariantArrayOutput {
	return o.ApplyT(func(v *EndpointConfiguration) EndpointConfigurationShadowProductionVariantArrayOutput {
		return v.ShadowProductionVariants
	}).(EndpointConfigurationShadowProductionVariantArrayOutput)
}

// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EndpointConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EndpointConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o EndpointConfigurationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EndpointConfiguration) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type EndpointConfigurationArrayOutput struct{ *pulumi.OutputState }

func (EndpointConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointConfiguration)(nil)).Elem()
}

func (o EndpointConfigurationArrayOutput) ToEndpointConfigurationArrayOutput() EndpointConfigurationArrayOutput {
	return o
}

func (o EndpointConfigurationArrayOutput) ToEndpointConfigurationArrayOutputWithContext(ctx context.Context) EndpointConfigurationArrayOutput {
	return o
}

func (o EndpointConfigurationArrayOutput) Index(i pulumi.IntInput) EndpointConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointConfiguration {
		return vs[0].([]*EndpointConfiguration)[vs[1].(int)]
	}).(EndpointConfigurationOutput)
}

type EndpointConfigurationMapOutput struct{ *pulumi.OutputState }

func (EndpointConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointConfiguration)(nil)).Elem()
}

func (o EndpointConfigurationMapOutput) ToEndpointConfigurationMapOutput() EndpointConfigurationMapOutput {
	return o
}

func (o EndpointConfigurationMapOutput) ToEndpointConfigurationMapOutputWithContext(ctx context.Context) EndpointConfigurationMapOutput {
	return o
}

func (o EndpointConfigurationMapOutput) MapIndex(k pulumi.StringInput) EndpointConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointConfiguration {
		return vs[0].(map[string]*EndpointConfiguration)[vs[1].(string)]
	}).(EndpointConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointConfigurationInput)(nil)).Elem(), &EndpointConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointConfigurationArrayInput)(nil)).Elem(), EndpointConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointConfigurationMapInput)(nil)).Elem(), EndpointConfigurationMap{})
	pulumi.RegisterOutputType(EndpointConfigurationOutput{})
	pulumi.RegisterOutputType(EndpointConfigurationArrayOutput{})
	pulumi.RegisterOutputType(EndpointConfigurationMapOutput{})
}
