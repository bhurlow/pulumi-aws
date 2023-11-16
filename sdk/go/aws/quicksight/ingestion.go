// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS QuickSight Ingestion.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := quicksight.NewIngestion(ctx, "example", &quicksight.IngestionArgs{
//				DataSetId:     pulumi.Any(aws_quicksight_data_set.Example.Data_set_id),
//				IngestionId:   pulumi.String("example-id"),
//				IngestionType: pulumi.String("FULL_REFRESH"),
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
// Using `pulumi import`, import QuickSight Ingestion using the AWS account ID, data set ID, and ingestion ID separated by commas (`,`). For example:
//
// ```sh
//
//	$ pulumi import aws:quicksight/ingestion:Ingestion example 123456789012,example-dataset-id,example-ingestion-id
//
// ```
type Ingestion struct {
	pulumi.CustomResourceState

	// ARN of the Ingestion.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// ID of the dataset used in the ingestion.
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// ID for the ingestion.
	IngestionId pulumi.StringOutput `pulumi:"ingestionId"`
	// Ingestion status.
	IngestionStatus pulumi.StringOutput `pulumi:"ingestionStatus"`
	// Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
	//
	// The following arguments are optional:
	IngestionType pulumi.StringOutput `pulumi:"ingestionType"`
}

// NewIngestion registers a new resource with the given unique name, arguments, and options.
func NewIngestion(ctx *pulumi.Context,
	name string, args *IngestionArgs, opts ...pulumi.ResourceOption) (*Ingestion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.IngestionId == nil {
		return nil, errors.New("invalid value for required argument 'IngestionId'")
	}
	if args.IngestionType == nil {
		return nil, errors.New("invalid value for required argument 'IngestionType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ingestion
	err := ctx.RegisterResource("aws:quicksight/ingestion:Ingestion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngestion gets an existing Ingestion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngestion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngestionState, opts ...pulumi.ResourceOption) (*Ingestion, error) {
	var resource Ingestion
	err := ctx.ReadResource("aws:quicksight/ingestion:Ingestion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ingestion resources.
type ingestionState struct {
	// ARN of the Ingestion.
	Arn *string `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// ID of the dataset used in the ingestion.
	DataSetId *string `pulumi:"dataSetId"`
	// ID for the ingestion.
	IngestionId *string `pulumi:"ingestionId"`
	// Ingestion status.
	IngestionStatus *string `pulumi:"ingestionStatus"`
	// Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
	//
	// The following arguments are optional:
	IngestionType *string `pulumi:"ingestionType"`
}

type IngestionState struct {
	// ARN of the Ingestion.
	Arn pulumi.StringPtrInput
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// ID of the dataset used in the ingestion.
	DataSetId pulumi.StringPtrInput
	// ID for the ingestion.
	IngestionId pulumi.StringPtrInput
	// Ingestion status.
	IngestionStatus pulumi.StringPtrInput
	// Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
	//
	// The following arguments are optional:
	IngestionType pulumi.StringPtrInput
}

func (IngestionState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionState)(nil)).Elem()
}

type ingestionArgs struct {
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// ID of the dataset used in the ingestion.
	DataSetId string `pulumi:"dataSetId"`
	// ID for the ingestion.
	IngestionId string `pulumi:"ingestionId"`
	// Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
	//
	// The following arguments are optional:
	IngestionType string `pulumi:"ingestionType"`
}

// The set of arguments for constructing a Ingestion resource.
type IngestionArgs struct {
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// ID of the dataset used in the ingestion.
	DataSetId pulumi.StringInput
	// ID for the ingestion.
	IngestionId pulumi.StringInput
	// Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
	//
	// The following arguments are optional:
	IngestionType pulumi.StringInput
}

func (IngestionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionArgs)(nil)).Elem()
}

type IngestionInput interface {
	pulumi.Input

	ToIngestionOutput() IngestionOutput
	ToIngestionOutputWithContext(ctx context.Context) IngestionOutput
}

func (*Ingestion) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingestion)(nil)).Elem()
}

func (i *Ingestion) ToIngestionOutput() IngestionOutput {
	return i.ToIngestionOutputWithContext(context.Background())
}

func (i *Ingestion) ToIngestionOutputWithContext(ctx context.Context) IngestionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionOutput)
}

// IngestionArrayInput is an input type that accepts IngestionArray and IngestionArrayOutput values.
// You can construct a concrete instance of `IngestionArrayInput` via:
//
//	IngestionArray{ IngestionArgs{...} }
type IngestionArrayInput interface {
	pulumi.Input

	ToIngestionArrayOutput() IngestionArrayOutput
	ToIngestionArrayOutputWithContext(context.Context) IngestionArrayOutput
}

type IngestionArray []IngestionInput

func (IngestionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ingestion)(nil)).Elem()
}

func (i IngestionArray) ToIngestionArrayOutput() IngestionArrayOutput {
	return i.ToIngestionArrayOutputWithContext(context.Background())
}

func (i IngestionArray) ToIngestionArrayOutputWithContext(ctx context.Context) IngestionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionArrayOutput)
}

// IngestionMapInput is an input type that accepts IngestionMap and IngestionMapOutput values.
// You can construct a concrete instance of `IngestionMapInput` via:
//
//	IngestionMap{ "key": IngestionArgs{...} }
type IngestionMapInput interface {
	pulumi.Input

	ToIngestionMapOutput() IngestionMapOutput
	ToIngestionMapOutputWithContext(context.Context) IngestionMapOutput
}

type IngestionMap map[string]IngestionInput

func (IngestionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ingestion)(nil)).Elem()
}

func (i IngestionMap) ToIngestionMapOutput() IngestionMapOutput {
	return i.ToIngestionMapOutputWithContext(context.Background())
}

func (i IngestionMap) ToIngestionMapOutputWithContext(ctx context.Context) IngestionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionMapOutput)
}

type IngestionOutput struct{ *pulumi.OutputState }

func (IngestionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingestion)(nil)).Elem()
}

func (o IngestionOutput) ToIngestionOutput() IngestionOutput {
	return o
}

func (o IngestionOutput) ToIngestionOutputWithContext(ctx context.Context) IngestionOutput {
	return o
}

// ARN of the Ingestion.
func (o IngestionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// AWS account ID.
func (o IngestionOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

// ID of the dataset used in the ingestion.
func (o IngestionOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// ID for the ingestion.
func (o IngestionOutput) IngestionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.IngestionId }).(pulumi.StringOutput)
}

// Ingestion status.
func (o IngestionOutput) IngestionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.IngestionStatus }).(pulumi.StringOutput)
}

// Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
//
// The following arguments are optional:
func (o IngestionOutput) IngestionType() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.IngestionType }).(pulumi.StringOutput)
}

type IngestionArrayOutput struct{ *pulumi.OutputState }

func (IngestionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ingestion)(nil)).Elem()
}

func (o IngestionArrayOutput) ToIngestionArrayOutput() IngestionArrayOutput {
	return o
}

func (o IngestionArrayOutput) ToIngestionArrayOutputWithContext(ctx context.Context) IngestionArrayOutput {
	return o
}

func (o IngestionArrayOutput) Index(i pulumi.IntInput) IngestionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ingestion {
		return vs[0].([]*Ingestion)[vs[1].(int)]
	}).(IngestionOutput)
}

type IngestionMapOutput struct{ *pulumi.OutputState }

func (IngestionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ingestion)(nil)).Elem()
}

func (o IngestionMapOutput) ToIngestionMapOutput() IngestionMapOutput {
	return o
}

func (o IngestionMapOutput) ToIngestionMapOutputWithContext(ctx context.Context) IngestionMapOutput {
	return o
}

func (o IngestionMapOutput) MapIndex(k pulumi.StringInput) IngestionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ingestion {
		return vs[0].(map[string]*Ingestion)[vs[1].(string)]
	}).(IngestionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionInput)(nil)).Elem(), &Ingestion{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionArrayInput)(nil)).Elem(), IngestionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionMapInput)(nil)).Elem(), IngestionMap{})
	pulumi.RegisterOutputType(IngestionOutput{})
	pulumi.RegisterOutputType(IngestionArrayOutput{})
	pulumi.RegisterOutputType(IngestionMapOutput{})
}
