// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows you to set a policy of an SQS Queue
// while referencing ARN of the queue within the policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// queue, err := sqs.NewQueue(ctx, "queue", nil)
// if err != nil {
// return err
// }
// testPolicyDocument := queue.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
// return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
// Statements: []iam.GetPolicyDocumentStatement{
// {
// Sid: "First",
// Effect: "Allow",
// Principals: []iam.GetPolicyDocumentStatementPrincipal{
// {
// Type: "*",
// Identifiers: []string{
// "*",
// },
// },
// },
// Actions: []string{
// "sqs:SendMessage",
// },
// Resources: []string{
// arn,
// },
// Conditions: []iam.GetPolicyDocumentStatementCondition{
// {
// Test: "ArnEquals",
// Variable: "aws:SourceArn",
// Values: interface{}{
// aws_sns_topic.Example.Arn,
// },
// },
// },
// },
// },
// }, nil), nil
// }).(iam.GetPolicyDocumentResultOutput)
// _, err = sqs.NewQueuePolicy(ctx, "testQueuePolicy", &sqs.QueuePolicyArgs{
// QueueUrl: queue.ID(),
// Policy: testPolicyDocument.ApplyT(func(testPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
// return &testPolicyDocument.Json, nil
// }).(pulumi.StringPtrOutput),
// })
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
//
// ## Import
//
// terraform import {
//
//	to = aws_sqs_queue_policy.test
//
//	id = "https://queue.amazonaws.com/0123456789012/myqueue" } Using `pulumi import`, import SQS Queue Policies using the queue URL. For exampleconsole % pulumi import aws_sqs_queue_policy.test https://queue.amazonaws.com/0123456789012/myqueue
type QueuePolicy struct {
	pulumi.CustomResourceState

	// The JSON policy for the SQS queue.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringOutput `pulumi:"queueUrl"`
}

// NewQueuePolicy registers a new resource with the given unique name, arguments, and options.
func NewQueuePolicy(ctx *pulumi.Context,
	name string, args *QueuePolicyArgs, opts ...pulumi.ResourceOption) (*QueuePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.QueueUrl == nil {
		return nil, errors.New("invalid value for required argument 'QueueUrl'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueuePolicy
	err := ctx.RegisterResource("aws:sqs/queuePolicy:QueuePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueuePolicy gets an existing QueuePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueuePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueuePolicyState, opts ...pulumi.ResourceOption) (*QueuePolicy, error) {
	var resource QueuePolicy
	err := ctx.ReadResource("aws:sqs/queuePolicy:QueuePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueuePolicy resources.
type queuePolicyState struct {
	// The JSON policy for the SQS queue.
	Policy interface{} `pulumi:"policy"`
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl *string `pulumi:"queueUrl"`
}

type QueuePolicyState struct {
	// The JSON policy for the SQS queue.
	Policy pulumi.Input
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringPtrInput
}

func (QueuePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*queuePolicyState)(nil)).Elem()
}

type queuePolicyArgs struct {
	// The JSON policy for the SQS queue.
	Policy interface{} `pulumi:"policy"`
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl string `pulumi:"queueUrl"`
}

// The set of arguments for constructing a QueuePolicy resource.
type QueuePolicyArgs struct {
	// The JSON policy for the SQS queue.
	Policy pulumi.Input
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringInput
}

func (QueuePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queuePolicyArgs)(nil)).Elem()
}

type QueuePolicyInput interface {
	pulumi.Input

	ToQueuePolicyOutput() QueuePolicyOutput
	ToQueuePolicyOutputWithContext(ctx context.Context) QueuePolicyOutput
}

func (*QueuePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuePolicy)(nil)).Elem()
}

func (i *QueuePolicy) ToQueuePolicyOutput() QueuePolicyOutput {
	return i.ToQueuePolicyOutputWithContext(context.Background())
}

func (i *QueuePolicy) ToQueuePolicyOutputWithContext(ctx context.Context) QueuePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePolicyOutput)
}

// QueuePolicyArrayInput is an input type that accepts QueuePolicyArray and QueuePolicyArrayOutput values.
// You can construct a concrete instance of `QueuePolicyArrayInput` via:
//
//	QueuePolicyArray{ QueuePolicyArgs{...} }
type QueuePolicyArrayInput interface {
	pulumi.Input

	ToQueuePolicyArrayOutput() QueuePolicyArrayOutput
	ToQueuePolicyArrayOutputWithContext(context.Context) QueuePolicyArrayOutput
}

type QueuePolicyArray []QueuePolicyInput

func (QueuePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuePolicy)(nil)).Elem()
}

func (i QueuePolicyArray) ToQueuePolicyArrayOutput() QueuePolicyArrayOutput {
	return i.ToQueuePolicyArrayOutputWithContext(context.Background())
}

func (i QueuePolicyArray) ToQueuePolicyArrayOutputWithContext(ctx context.Context) QueuePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePolicyArrayOutput)
}

// QueuePolicyMapInput is an input type that accepts QueuePolicyMap and QueuePolicyMapOutput values.
// You can construct a concrete instance of `QueuePolicyMapInput` via:
//
//	QueuePolicyMap{ "key": QueuePolicyArgs{...} }
type QueuePolicyMapInput interface {
	pulumi.Input

	ToQueuePolicyMapOutput() QueuePolicyMapOutput
	ToQueuePolicyMapOutputWithContext(context.Context) QueuePolicyMapOutput
}

type QueuePolicyMap map[string]QueuePolicyInput

func (QueuePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuePolicy)(nil)).Elem()
}

func (i QueuePolicyMap) ToQueuePolicyMapOutput() QueuePolicyMapOutput {
	return i.ToQueuePolicyMapOutputWithContext(context.Background())
}

func (i QueuePolicyMap) ToQueuePolicyMapOutputWithContext(ctx context.Context) QueuePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePolicyMapOutput)
}

type QueuePolicyOutput struct{ *pulumi.OutputState }

func (QueuePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuePolicy)(nil)).Elem()
}

func (o QueuePolicyOutput) ToQueuePolicyOutput() QueuePolicyOutput {
	return o
}

func (o QueuePolicyOutput) ToQueuePolicyOutputWithContext(ctx context.Context) QueuePolicyOutput {
	return o
}

// The JSON policy for the SQS queue.
func (o QueuePolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuePolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The URL of the SQS Queue to which to attach the policy
func (o QueuePolicyOutput) QueueUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuePolicy) pulumi.StringOutput { return v.QueueUrl }).(pulumi.StringOutput)
}

type QueuePolicyArrayOutput struct{ *pulumi.OutputState }

func (QueuePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuePolicy)(nil)).Elem()
}

func (o QueuePolicyArrayOutput) ToQueuePolicyArrayOutput() QueuePolicyArrayOutput {
	return o
}

func (o QueuePolicyArrayOutput) ToQueuePolicyArrayOutputWithContext(ctx context.Context) QueuePolicyArrayOutput {
	return o
}

func (o QueuePolicyArrayOutput) Index(i pulumi.IntInput) QueuePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QueuePolicy {
		return vs[0].([]*QueuePolicy)[vs[1].(int)]
	}).(QueuePolicyOutput)
}

type QueuePolicyMapOutput struct{ *pulumi.OutputState }

func (QueuePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuePolicy)(nil)).Elem()
}

func (o QueuePolicyMapOutput) ToQueuePolicyMapOutput() QueuePolicyMapOutput {
	return o
}

func (o QueuePolicyMapOutput) ToQueuePolicyMapOutputWithContext(ctx context.Context) QueuePolicyMapOutput {
	return o
}

func (o QueuePolicyMapOutput) MapIndex(k pulumi.StringInput) QueuePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QueuePolicy {
		return vs[0].(map[string]*QueuePolicy)[vs[1].(string)]
	}).(QueuePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueuePolicyInput)(nil)).Elem(), &QueuePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuePolicyArrayInput)(nil)).Elem(), QueuePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuePolicyMapInput)(nil)).Elem(), QueuePolicyMap{})
	pulumi.RegisterOutputType(QueuePolicyOutput{})
	pulumi.RegisterOutputType(QueuePolicyArrayOutput{})
	pulumi.RegisterOutputType(QueuePolicyMapOutput{})
}
