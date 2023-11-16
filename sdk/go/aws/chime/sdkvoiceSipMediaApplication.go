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

// A ChimeSDKVoice SIP Media Application is a managed object that passes values from a SIP rule to a target AWS Lambda function.
//
// ## Example Usage
// ### Basic Usage
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
//			_, err := chime.NewSdkvoiceSipMediaApplication(ctx, "example", &chime.SdkvoiceSipMediaApplicationArgs{
//				AwsRegion: pulumi.String("us-east-1"),
//				Endpoints: &chime.SdkvoiceSipMediaApplicationEndpointsArgs{
//					LambdaArn: pulumi.Any(aws_lambda_function.Test.Arn),
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
// Using `pulumi import`, import a ChimeSDKVoice SIP Media Application using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:chime/sdkvoiceSipMediaApplication:SdkvoiceSipMediaApplication example abcdef123456
//
// ```
type SdkvoiceSipMediaApplication struct {
	pulumi.CustomResourceState

	// ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
	AwsRegion pulumi.StringOutput `pulumi:"awsRegion"`
	// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
	Endpoints SdkvoiceSipMediaApplicationEndpointsOutput `pulumi:"endpoints"`
	// The name of the AWS Chime SDK Voice Sip Media Application.
	//
	// The following arguments are optional:
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewSdkvoiceSipMediaApplication registers a new resource with the given unique name, arguments, and options.
func NewSdkvoiceSipMediaApplication(ctx *pulumi.Context,
	name string, args *SdkvoiceSipMediaApplicationArgs, opts ...pulumi.ResourceOption) (*SdkvoiceSipMediaApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsRegion == nil {
		return nil, errors.New("invalid value for required argument 'AwsRegion'")
	}
	if args.Endpoints == nil {
		return nil, errors.New("invalid value for required argument 'Endpoints'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SdkvoiceSipMediaApplication
	err := ctx.RegisterResource("aws:chime/sdkvoiceSipMediaApplication:SdkvoiceSipMediaApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSdkvoiceSipMediaApplication gets an existing SdkvoiceSipMediaApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSdkvoiceSipMediaApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SdkvoiceSipMediaApplicationState, opts ...pulumi.ResourceOption) (*SdkvoiceSipMediaApplication, error) {
	var resource SdkvoiceSipMediaApplication
	err := ctx.ReadResource("aws:chime/sdkvoiceSipMediaApplication:SdkvoiceSipMediaApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SdkvoiceSipMediaApplication resources.
type sdkvoiceSipMediaApplicationState struct {
	// ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
	Arn *string `pulumi:"arn"`
	// The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
	AwsRegion *string `pulumi:"awsRegion"`
	// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
	Endpoints *SdkvoiceSipMediaApplicationEndpoints `pulumi:"endpoints"`
	// The name of the AWS Chime SDK Voice Sip Media Application.
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type SdkvoiceSipMediaApplicationState struct {
	// ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
	Arn pulumi.StringPtrInput
	// The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
	AwsRegion pulumi.StringPtrInput
	// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
	Endpoints SdkvoiceSipMediaApplicationEndpointsPtrInput
	// The name of the AWS Chime SDK Voice Sip Media Application.
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (SdkvoiceSipMediaApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*sdkvoiceSipMediaApplicationState)(nil)).Elem()
}

type sdkvoiceSipMediaApplicationArgs struct {
	// The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
	AwsRegion string `pulumi:"awsRegion"`
	// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
	Endpoints SdkvoiceSipMediaApplicationEndpoints `pulumi:"endpoints"`
	// The name of the AWS Chime SDK Voice Sip Media Application.
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SdkvoiceSipMediaApplication resource.
type SdkvoiceSipMediaApplicationArgs struct {
	// The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
	AwsRegion pulumi.StringInput
	// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
	Endpoints SdkvoiceSipMediaApplicationEndpointsInput
	// The name of the AWS Chime SDK Voice Sip Media Application.
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (SdkvoiceSipMediaApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sdkvoiceSipMediaApplicationArgs)(nil)).Elem()
}

type SdkvoiceSipMediaApplicationInput interface {
	pulumi.Input

	ToSdkvoiceSipMediaApplicationOutput() SdkvoiceSipMediaApplicationOutput
	ToSdkvoiceSipMediaApplicationOutputWithContext(ctx context.Context) SdkvoiceSipMediaApplicationOutput
}

func (*SdkvoiceSipMediaApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**SdkvoiceSipMediaApplication)(nil)).Elem()
}

func (i *SdkvoiceSipMediaApplication) ToSdkvoiceSipMediaApplicationOutput() SdkvoiceSipMediaApplicationOutput {
	return i.ToSdkvoiceSipMediaApplicationOutputWithContext(context.Background())
}

func (i *SdkvoiceSipMediaApplication) ToSdkvoiceSipMediaApplicationOutputWithContext(ctx context.Context) SdkvoiceSipMediaApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdkvoiceSipMediaApplicationOutput)
}

// SdkvoiceSipMediaApplicationArrayInput is an input type that accepts SdkvoiceSipMediaApplicationArray and SdkvoiceSipMediaApplicationArrayOutput values.
// You can construct a concrete instance of `SdkvoiceSipMediaApplicationArrayInput` via:
//
//	SdkvoiceSipMediaApplicationArray{ SdkvoiceSipMediaApplicationArgs{...} }
type SdkvoiceSipMediaApplicationArrayInput interface {
	pulumi.Input

	ToSdkvoiceSipMediaApplicationArrayOutput() SdkvoiceSipMediaApplicationArrayOutput
	ToSdkvoiceSipMediaApplicationArrayOutputWithContext(context.Context) SdkvoiceSipMediaApplicationArrayOutput
}

type SdkvoiceSipMediaApplicationArray []SdkvoiceSipMediaApplicationInput

func (SdkvoiceSipMediaApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SdkvoiceSipMediaApplication)(nil)).Elem()
}

func (i SdkvoiceSipMediaApplicationArray) ToSdkvoiceSipMediaApplicationArrayOutput() SdkvoiceSipMediaApplicationArrayOutput {
	return i.ToSdkvoiceSipMediaApplicationArrayOutputWithContext(context.Background())
}

func (i SdkvoiceSipMediaApplicationArray) ToSdkvoiceSipMediaApplicationArrayOutputWithContext(ctx context.Context) SdkvoiceSipMediaApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdkvoiceSipMediaApplicationArrayOutput)
}

// SdkvoiceSipMediaApplicationMapInput is an input type that accepts SdkvoiceSipMediaApplicationMap and SdkvoiceSipMediaApplicationMapOutput values.
// You can construct a concrete instance of `SdkvoiceSipMediaApplicationMapInput` via:
//
//	SdkvoiceSipMediaApplicationMap{ "key": SdkvoiceSipMediaApplicationArgs{...} }
type SdkvoiceSipMediaApplicationMapInput interface {
	pulumi.Input

	ToSdkvoiceSipMediaApplicationMapOutput() SdkvoiceSipMediaApplicationMapOutput
	ToSdkvoiceSipMediaApplicationMapOutputWithContext(context.Context) SdkvoiceSipMediaApplicationMapOutput
}

type SdkvoiceSipMediaApplicationMap map[string]SdkvoiceSipMediaApplicationInput

func (SdkvoiceSipMediaApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SdkvoiceSipMediaApplication)(nil)).Elem()
}

func (i SdkvoiceSipMediaApplicationMap) ToSdkvoiceSipMediaApplicationMapOutput() SdkvoiceSipMediaApplicationMapOutput {
	return i.ToSdkvoiceSipMediaApplicationMapOutputWithContext(context.Background())
}

func (i SdkvoiceSipMediaApplicationMap) ToSdkvoiceSipMediaApplicationMapOutputWithContext(ctx context.Context) SdkvoiceSipMediaApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdkvoiceSipMediaApplicationMapOutput)
}

type SdkvoiceSipMediaApplicationOutput struct{ *pulumi.OutputState }

func (SdkvoiceSipMediaApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SdkvoiceSipMediaApplication)(nil)).Elem()
}

func (o SdkvoiceSipMediaApplicationOutput) ToSdkvoiceSipMediaApplicationOutput() SdkvoiceSipMediaApplicationOutput {
	return o
}

func (o SdkvoiceSipMediaApplicationOutput) ToSdkvoiceSipMediaApplicationOutputWithContext(ctx context.Context) SdkvoiceSipMediaApplicationOutput {
	return o
}

// ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
func (o SdkvoiceSipMediaApplicationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SdkvoiceSipMediaApplication) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
func (o SdkvoiceSipMediaApplicationOutput) AwsRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *SdkvoiceSipMediaApplication) pulumi.StringOutput { return v.AwsRegion }).(pulumi.StringOutput)
}

// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
func (o SdkvoiceSipMediaApplicationOutput) Endpoints() SdkvoiceSipMediaApplicationEndpointsOutput {
	return o.ApplyT(func(v *SdkvoiceSipMediaApplication) SdkvoiceSipMediaApplicationEndpointsOutput { return v.Endpoints }).(SdkvoiceSipMediaApplicationEndpointsOutput)
}

// The name of the AWS Chime SDK Voice Sip Media Application.
//
// The following arguments are optional:
func (o SdkvoiceSipMediaApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SdkvoiceSipMediaApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SdkvoiceSipMediaApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SdkvoiceSipMediaApplication) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o SdkvoiceSipMediaApplicationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SdkvoiceSipMediaApplication) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type SdkvoiceSipMediaApplicationArrayOutput struct{ *pulumi.OutputState }

func (SdkvoiceSipMediaApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SdkvoiceSipMediaApplication)(nil)).Elem()
}

func (o SdkvoiceSipMediaApplicationArrayOutput) ToSdkvoiceSipMediaApplicationArrayOutput() SdkvoiceSipMediaApplicationArrayOutput {
	return o
}

func (o SdkvoiceSipMediaApplicationArrayOutput) ToSdkvoiceSipMediaApplicationArrayOutputWithContext(ctx context.Context) SdkvoiceSipMediaApplicationArrayOutput {
	return o
}

func (o SdkvoiceSipMediaApplicationArrayOutput) Index(i pulumi.IntInput) SdkvoiceSipMediaApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SdkvoiceSipMediaApplication {
		return vs[0].([]*SdkvoiceSipMediaApplication)[vs[1].(int)]
	}).(SdkvoiceSipMediaApplicationOutput)
}

type SdkvoiceSipMediaApplicationMapOutput struct{ *pulumi.OutputState }

func (SdkvoiceSipMediaApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SdkvoiceSipMediaApplication)(nil)).Elem()
}

func (o SdkvoiceSipMediaApplicationMapOutput) ToSdkvoiceSipMediaApplicationMapOutput() SdkvoiceSipMediaApplicationMapOutput {
	return o
}

func (o SdkvoiceSipMediaApplicationMapOutput) ToSdkvoiceSipMediaApplicationMapOutputWithContext(ctx context.Context) SdkvoiceSipMediaApplicationMapOutput {
	return o
}

func (o SdkvoiceSipMediaApplicationMapOutput) MapIndex(k pulumi.StringInput) SdkvoiceSipMediaApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SdkvoiceSipMediaApplication {
		return vs[0].(map[string]*SdkvoiceSipMediaApplication)[vs[1].(string)]
	}).(SdkvoiceSipMediaApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SdkvoiceSipMediaApplicationInput)(nil)).Elem(), &SdkvoiceSipMediaApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*SdkvoiceSipMediaApplicationArrayInput)(nil)).Elem(), SdkvoiceSipMediaApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SdkvoiceSipMediaApplicationMapInput)(nil)).Elem(), SdkvoiceSipMediaApplicationMap{})
	pulumi.RegisterOutputType(SdkvoiceSipMediaApplicationOutput{})
	pulumi.RegisterOutputType(SdkvoiceSipMediaApplicationArrayOutput{})
	pulumi.RegisterOutputType(SdkvoiceSipMediaApplicationMapOutput{})
}
