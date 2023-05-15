// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appintegrations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon AppIntegrations Data Integration resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/appintegrations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appintegrations.NewDataIntegration(ctx, "example", &appintegrations.DataIntegrationArgs{
//				Description: pulumi.String("example"),
//				KmsKey:      pulumi.Any(aws_kms_key.Test.Arn),
//				SourceUri:   pulumi.String("Salesforce://AppFlow/example"),
//				ScheduleConfig: &appintegrations.DataIntegrationScheduleConfigArgs{
//					FirstExecutionFrom: pulumi.String("1439788442681"),
//					Object:             pulumi.String("Account"),
//					ScheduleExpression: pulumi.String("rate(1 hour)"),
//				},
//				Tags: pulumi.StringMap{
//					"Key1": pulumi.String("Value1"),
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
// Amazon AppIntegrations Data Integrations can be imported using the `id` e.g.,
//
// ```sh
//
//	$ pulumi import aws:appintegrations/dataIntegration:DataIntegration example 12345678-1234-1234-1234-123456789123
//
// ```
type DataIntegration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Data Integration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies the description of the Data Integration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
	KmsKey pulumi.StringOutput `pulumi:"kmsKey"`
	// Specifies the name of the Data Integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
	ScheduleConfig DataIntegrationScheduleConfigOutput `pulumi:"scheduleConfig"`
	// Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
	SourceUri pulumi.StringOutput `pulumi:"sourceUri"`
	// Tags to apply to the Data Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDataIntegration registers a new resource with the given unique name, arguments, and options.
func NewDataIntegration(ctx *pulumi.Context,
	name string, args *DataIntegrationArgs, opts ...pulumi.ResourceOption) (*DataIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KmsKey == nil {
		return nil, errors.New("invalid value for required argument 'KmsKey'")
	}
	if args.ScheduleConfig == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleConfig'")
	}
	if args.SourceUri == nil {
		return nil, errors.New("invalid value for required argument 'SourceUri'")
	}
	var resource DataIntegration
	err := ctx.RegisterResource("aws:appintegrations/dataIntegration:DataIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataIntegration gets an existing DataIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataIntegrationState, opts ...pulumi.ResourceOption) (*DataIntegration, error) {
	var resource DataIntegration
	err := ctx.ReadResource("aws:appintegrations/dataIntegration:DataIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataIntegration resources.
type dataIntegrationState struct {
	// The Amazon Resource Name (ARN) of the Data Integration.
	Arn *string `pulumi:"arn"`
	// Specifies the description of the Data Integration.
	Description *string `pulumi:"description"`
	// Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
	KmsKey *string `pulumi:"kmsKey"`
	// Specifies the name of the Data Integration.
	Name *string `pulumi:"name"`
	// A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
	ScheduleConfig *DataIntegrationScheduleConfig `pulumi:"scheduleConfig"`
	// Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
	SourceUri *string `pulumi:"sourceUri"`
	// Tags to apply to the Data Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DataIntegrationState struct {
	// The Amazon Resource Name (ARN) of the Data Integration.
	Arn pulumi.StringPtrInput
	// Specifies the description of the Data Integration.
	Description pulumi.StringPtrInput
	// Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
	KmsKey pulumi.StringPtrInput
	// Specifies the name of the Data Integration.
	Name pulumi.StringPtrInput
	// A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
	ScheduleConfig DataIntegrationScheduleConfigPtrInput
	// Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
	SourceUri pulumi.StringPtrInput
	// Tags to apply to the Data Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (DataIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataIntegrationState)(nil)).Elem()
}

type dataIntegrationArgs struct {
	// Specifies the description of the Data Integration.
	Description *string `pulumi:"description"`
	// Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
	KmsKey string `pulumi:"kmsKey"`
	// Specifies the name of the Data Integration.
	Name *string `pulumi:"name"`
	// A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
	ScheduleConfig DataIntegrationScheduleConfig `pulumi:"scheduleConfig"`
	// Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
	SourceUri string `pulumi:"sourceUri"`
	// Tags to apply to the Data Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataIntegration resource.
type DataIntegrationArgs struct {
	// Specifies the description of the Data Integration.
	Description pulumi.StringPtrInput
	// Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
	KmsKey pulumi.StringInput
	// Specifies the name of the Data Integration.
	Name pulumi.StringPtrInput
	// A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
	ScheduleConfig DataIntegrationScheduleConfigInput
	// Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
	SourceUri pulumi.StringInput
	// Tags to apply to the Data Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (DataIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataIntegrationArgs)(nil)).Elem()
}

type DataIntegrationInput interface {
	pulumi.Input

	ToDataIntegrationOutput() DataIntegrationOutput
	ToDataIntegrationOutputWithContext(ctx context.Context) DataIntegrationOutput
}

func (*DataIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**DataIntegration)(nil)).Elem()
}

func (i *DataIntegration) ToDataIntegrationOutput() DataIntegrationOutput {
	return i.ToDataIntegrationOutputWithContext(context.Background())
}

func (i *DataIntegration) ToDataIntegrationOutputWithContext(ctx context.Context) DataIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataIntegrationOutput)
}

// DataIntegrationArrayInput is an input type that accepts DataIntegrationArray and DataIntegrationArrayOutput values.
// You can construct a concrete instance of `DataIntegrationArrayInput` via:
//
//	DataIntegrationArray{ DataIntegrationArgs{...} }
type DataIntegrationArrayInput interface {
	pulumi.Input

	ToDataIntegrationArrayOutput() DataIntegrationArrayOutput
	ToDataIntegrationArrayOutputWithContext(context.Context) DataIntegrationArrayOutput
}

type DataIntegrationArray []DataIntegrationInput

func (DataIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataIntegration)(nil)).Elem()
}

func (i DataIntegrationArray) ToDataIntegrationArrayOutput() DataIntegrationArrayOutput {
	return i.ToDataIntegrationArrayOutputWithContext(context.Background())
}

func (i DataIntegrationArray) ToDataIntegrationArrayOutputWithContext(ctx context.Context) DataIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataIntegrationArrayOutput)
}

// DataIntegrationMapInput is an input type that accepts DataIntegrationMap and DataIntegrationMapOutput values.
// You can construct a concrete instance of `DataIntegrationMapInput` via:
//
//	DataIntegrationMap{ "key": DataIntegrationArgs{...} }
type DataIntegrationMapInput interface {
	pulumi.Input

	ToDataIntegrationMapOutput() DataIntegrationMapOutput
	ToDataIntegrationMapOutputWithContext(context.Context) DataIntegrationMapOutput
}

type DataIntegrationMap map[string]DataIntegrationInput

func (DataIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataIntegration)(nil)).Elem()
}

func (i DataIntegrationMap) ToDataIntegrationMapOutput() DataIntegrationMapOutput {
	return i.ToDataIntegrationMapOutputWithContext(context.Background())
}

func (i DataIntegrationMap) ToDataIntegrationMapOutputWithContext(ctx context.Context) DataIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataIntegrationMapOutput)
}

type DataIntegrationOutput struct{ *pulumi.OutputState }

func (DataIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataIntegration)(nil)).Elem()
}

func (o DataIntegrationOutput) ToDataIntegrationOutput() DataIntegrationOutput {
	return o
}

func (o DataIntegrationOutput) ToDataIntegrationOutputWithContext(ctx context.Context) DataIntegrationOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Data Integration.
func (o DataIntegrationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DataIntegration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies the description of the Data Integration.
func (o DataIntegrationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataIntegration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
func (o DataIntegrationOutput) KmsKey() pulumi.StringOutput {
	return o.ApplyT(func(v *DataIntegration) pulumi.StringOutput { return v.KmsKey }).(pulumi.StringOutput)
}

// Specifies the name of the Data Integration.
func (o DataIntegrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataIntegration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
func (o DataIntegrationOutput) ScheduleConfig() DataIntegrationScheduleConfigOutput {
	return o.ApplyT(func(v *DataIntegration) DataIntegrationScheduleConfigOutput { return v.ScheduleConfig }).(DataIntegrationScheduleConfigOutput)
}

// Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
func (o DataIntegrationOutput) SourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DataIntegration) pulumi.StringOutput { return v.SourceUri }).(pulumi.StringOutput)
}

// Tags to apply to the Data Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DataIntegrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataIntegration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o DataIntegrationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataIntegration) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type DataIntegrationArrayOutput struct{ *pulumi.OutputState }

func (DataIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataIntegration)(nil)).Elem()
}

func (o DataIntegrationArrayOutput) ToDataIntegrationArrayOutput() DataIntegrationArrayOutput {
	return o
}

func (o DataIntegrationArrayOutput) ToDataIntegrationArrayOutputWithContext(ctx context.Context) DataIntegrationArrayOutput {
	return o
}

func (o DataIntegrationArrayOutput) Index(i pulumi.IntInput) DataIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataIntegration {
		return vs[0].([]*DataIntegration)[vs[1].(int)]
	}).(DataIntegrationOutput)
}

type DataIntegrationMapOutput struct{ *pulumi.OutputState }

func (DataIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataIntegration)(nil)).Elem()
}

func (o DataIntegrationMapOutput) ToDataIntegrationMapOutput() DataIntegrationMapOutput {
	return o
}

func (o DataIntegrationMapOutput) ToDataIntegrationMapOutputWithContext(ctx context.Context) DataIntegrationMapOutput {
	return o
}

func (o DataIntegrationMapOutput) MapIndex(k pulumi.StringInput) DataIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataIntegration {
		return vs[0].(map[string]*DataIntegration)[vs[1].(string)]
	}).(DataIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataIntegrationInput)(nil)).Elem(), &DataIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataIntegrationArrayInput)(nil)).Elem(), DataIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataIntegrationMapInput)(nil)).Elem(), DataIntegrationMap{})
	pulumi.RegisterOutputType(DataIntegrationOutput{})
	pulumi.RegisterOutputType(DataIntegrationArrayOutput{})
	pulumi.RegisterOutputType(DataIntegrationMapOutput{})
}
