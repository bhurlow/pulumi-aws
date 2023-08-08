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

// Provides a SQS Queue Redrive Allow Policy resource.
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
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleQueue, err := sqs.NewQueue(ctx, "exampleQueue", nil)
//			if err != nil {
//				return err
//			}
//			src, err := sqs.NewQueue(ctx, "src", &sqs.QueueArgs{
//				RedrivePolicy: exampleQueue.Arn.ApplyT(func(arn string) (pulumi.String, error) {
//					var _zero pulumi.String
//					tmpJSON0, err := json.Marshal(map[string]interface{}{
//						"deadLetterTargetArn": arn,
//						"maxReceiveCount":     4,
//					})
//					if err != nil {
//						return _zero, err
//					}
//					json0 := string(tmpJSON0)
//					return pulumi.String(json0), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sqs.NewRedriveAllowPolicy(ctx, "exampleRedriveAllowPolicy", &sqs.RedriveAllowPolicyArgs{
//				QueueUrl: exampleQueue.ID(),
//				RedriveAllowPolicy: src.Arn.ApplyT(func(arn string) (pulumi.String, error) {
//					var _zero pulumi.String
//					tmpJSON1, err := json.Marshal(map[string]interface{}{
//						"redrivePermission": "byQueue",
//						"sourceQueueArns": []string{
//							arn,
//						},
//					})
//					if err != nil {
//						return _zero, err
//					}
//					json1 := string(tmpJSON1)
//					return pulumi.String(json1), nil
//				}).(pulumi.StringOutput),
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
//	to = aws_sqs_queue_redrive_allow_policy.test
//
//	id = "https://queue.amazonaws.com/0123456789012/myqueue" } Using `pulumi import`, import SQS Queue Redrive Allow Policies using the queue URL. For exampleconsole % pulumi import aws_sqs_queue_redrive_allow_policy.test https://queue.amazonaws.com/0123456789012/myqueue
type RedriveAllowPolicy struct {
	pulumi.CustomResourceState

	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringOutput `pulumi:"queueUrl"`
	// The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
	RedriveAllowPolicy pulumi.StringOutput `pulumi:"redriveAllowPolicy"`
}

// NewRedriveAllowPolicy registers a new resource with the given unique name, arguments, and options.
func NewRedriveAllowPolicy(ctx *pulumi.Context,
	name string, args *RedriveAllowPolicyArgs, opts ...pulumi.ResourceOption) (*RedriveAllowPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.QueueUrl == nil {
		return nil, errors.New("invalid value for required argument 'QueueUrl'")
	}
	if args.RedriveAllowPolicy == nil {
		return nil, errors.New("invalid value for required argument 'RedriveAllowPolicy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RedriveAllowPolicy
	err := ctx.RegisterResource("aws:sqs/redriveAllowPolicy:RedriveAllowPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRedriveAllowPolicy gets an existing RedriveAllowPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRedriveAllowPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedriveAllowPolicyState, opts ...pulumi.ResourceOption) (*RedriveAllowPolicy, error) {
	var resource RedriveAllowPolicy
	err := ctx.ReadResource("aws:sqs/redriveAllowPolicy:RedriveAllowPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RedriveAllowPolicy resources.
type redriveAllowPolicyState struct {
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl *string `pulumi:"queueUrl"`
	// The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
	RedriveAllowPolicy *string `pulumi:"redriveAllowPolicy"`
}

type RedriveAllowPolicyState struct {
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringPtrInput
	// The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
	RedriveAllowPolicy pulumi.StringPtrInput
}

func (RedriveAllowPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*redriveAllowPolicyState)(nil)).Elem()
}

type redriveAllowPolicyArgs struct {
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl string `pulumi:"queueUrl"`
	// The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
	RedriveAllowPolicy string `pulumi:"redriveAllowPolicy"`
}

// The set of arguments for constructing a RedriveAllowPolicy resource.
type RedriveAllowPolicyArgs struct {
	// The URL of the SQS Queue to which to attach the policy
	QueueUrl pulumi.StringInput
	// The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
	RedriveAllowPolicy pulumi.StringInput
}

func (RedriveAllowPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redriveAllowPolicyArgs)(nil)).Elem()
}

type RedriveAllowPolicyInput interface {
	pulumi.Input

	ToRedriveAllowPolicyOutput() RedriveAllowPolicyOutput
	ToRedriveAllowPolicyOutputWithContext(ctx context.Context) RedriveAllowPolicyOutput
}

func (*RedriveAllowPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RedriveAllowPolicy)(nil)).Elem()
}

func (i *RedriveAllowPolicy) ToRedriveAllowPolicyOutput() RedriveAllowPolicyOutput {
	return i.ToRedriveAllowPolicyOutputWithContext(context.Background())
}

func (i *RedriveAllowPolicy) ToRedriveAllowPolicyOutputWithContext(ctx context.Context) RedriveAllowPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedriveAllowPolicyOutput)
}

// RedriveAllowPolicyArrayInput is an input type that accepts RedriveAllowPolicyArray and RedriveAllowPolicyArrayOutput values.
// You can construct a concrete instance of `RedriveAllowPolicyArrayInput` via:
//
//	RedriveAllowPolicyArray{ RedriveAllowPolicyArgs{...} }
type RedriveAllowPolicyArrayInput interface {
	pulumi.Input

	ToRedriveAllowPolicyArrayOutput() RedriveAllowPolicyArrayOutput
	ToRedriveAllowPolicyArrayOutputWithContext(context.Context) RedriveAllowPolicyArrayOutput
}

type RedriveAllowPolicyArray []RedriveAllowPolicyInput

func (RedriveAllowPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RedriveAllowPolicy)(nil)).Elem()
}

func (i RedriveAllowPolicyArray) ToRedriveAllowPolicyArrayOutput() RedriveAllowPolicyArrayOutput {
	return i.ToRedriveAllowPolicyArrayOutputWithContext(context.Background())
}

func (i RedriveAllowPolicyArray) ToRedriveAllowPolicyArrayOutputWithContext(ctx context.Context) RedriveAllowPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedriveAllowPolicyArrayOutput)
}

// RedriveAllowPolicyMapInput is an input type that accepts RedriveAllowPolicyMap and RedriveAllowPolicyMapOutput values.
// You can construct a concrete instance of `RedriveAllowPolicyMapInput` via:
//
//	RedriveAllowPolicyMap{ "key": RedriveAllowPolicyArgs{...} }
type RedriveAllowPolicyMapInput interface {
	pulumi.Input

	ToRedriveAllowPolicyMapOutput() RedriveAllowPolicyMapOutput
	ToRedriveAllowPolicyMapOutputWithContext(context.Context) RedriveAllowPolicyMapOutput
}

type RedriveAllowPolicyMap map[string]RedriveAllowPolicyInput

func (RedriveAllowPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RedriveAllowPolicy)(nil)).Elem()
}

func (i RedriveAllowPolicyMap) ToRedriveAllowPolicyMapOutput() RedriveAllowPolicyMapOutput {
	return i.ToRedriveAllowPolicyMapOutputWithContext(context.Background())
}

func (i RedriveAllowPolicyMap) ToRedriveAllowPolicyMapOutputWithContext(ctx context.Context) RedriveAllowPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedriveAllowPolicyMapOutput)
}

type RedriveAllowPolicyOutput struct{ *pulumi.OutputState }

func (RedriveAllowPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedriveAllowPolicy)(nil)).Elem()
}

func (o RedriveAllowPolicyOutput) ToRedriveAllowPolicyOutput() RedriveAllowPolicyOutput {
	return o
}

func (o RedriveAllowPolicyOutput) ToRedriveAllowPolicyOutputWithContext(ctx context.Context) RedriveAllowPolicyOutput {
	return o
}

// The URL of the SQS Queue to which to attach the policy
func (o RedriveAllowPolicyOutput) QueueUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *RedriveAllowPolicy) pulumi.StringOutput { return v.QueueUrl }).(pulumi.StringOutput)
}

// The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).
func (o RedriveAllowPolicyOutput) RedriveAllowPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *RedriveAllowPolicy) pulumi.StringOutput { return v.RedriveAllowPolicy }).(pulumi.StringOutput)
}

type RedriveAllowPolicyArrayOutput struct{ *pulumi.OutputState }

func (RedriveAllowPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RedriveAllowPolicy)(nil)).Elem()
}

func (o RedriveAllowPolicyArrayOutput) ToRedriveAllowPolicyArrayOutput() RedriveAllowPolicyArrayOutput {
	return o
}

func (o RedriveAllowPolicyArrayOutput) ToRedriveAllowPolicyArrayOutputWithContext(ctx context.Context) RedriveAllowPolicyArrayOutput {
	return o
}

func (o RedriveAllowPolicyArrayOutput) Index(i pulumi.IntInput) RedriveAllowPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RedriveAllowPolicy {
		return vs[0].([]*RedriveAllowPolicy)[vs[1].(int)]
	}).(RedriveAllowPolicyOutput)
}

type RedriveAllowPolicyMapOutput struct{ *pulumi.OutputState }

func (RedriveAllowPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RedriveAllowPolicy)(nil)).Elem()
}

func (o RedriveAllowPolicyMapOutput) ToRedriveAllowPolicyMapOutput() RedriveAllowPolicyMapOutput {
	return o
}

func (o RedriveAllowPolicyMapOutput) ToRedriveAllowPolicyMapOutputWithContext(ctx context.Context) RedriveAllowPolicyMapOutput {
	return o
}

func (o RedriveAllowPolicyMapOutput) MapIndex(k pulumi.StringInput) RedriveAllowPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RedriveAllowPolicy {
		return vs[0].(map[string]*RedriveAllowPolicy)[vs[1].(string)]
	}).(RedriveAllowPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RedriveAllowPolicyInput)(nil)).Elem(), &RedriveAllowPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedriveAllowPolicyArrayInput)(nil)).Elem(), RedriveAllowPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedriveAllowPolicyMapInput)(nil)).Elem(), RedriveAllowPolicyMap{})
	pulumi.RegisterOutputType(RedriveAllowPolicyOutput{})
	pulumi.RegisterOutputType(RedriveAllowPolicyArrayOutput{})
	pulumi.RegisterOutputType(RedriveAllowPolicyMapOutput{})
}
