// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"deadLetterTargetArn": aws_sqs_queue.Queue_deadletter.Arn,
// 			"maxReceiveCount":     4,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := sqs.NewQueue(ctx, "queue", &sqs.QueueArgs{
// 			DelaySeconds:            pulumi.Int(90),
// 			MaxMessageSize:          pulumi.Int(2048),
// 			MessageRetentionSeconds: pulumi.Int(86400),
// 			ReceiveWaitTimeSeconds:  pulumi.Int(10),
// 			RedrivePolicy:           pulumi.String(json0),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## FIFO queue
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sqs.NewQueue(ctx, "queue", &sqs.QueueArgs{
// 			ContentBasedDeduplication: pulumi.Bool(true),
// 			FifoQueue:                 pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Server-side encryption (SSE)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sqs.NewQueue(ctx, "queue", &sqs.QueueArgs{
// 			KmsDataKeyReusePeriodSeconds: pulumi.Int(300),
// 			KmsMasterKeyId:               pulumi.String("alias/aws/sqs"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// SQS Queues can be imported using the `queue url`, e.g.
//
// ```sh
//  $ pulumi import aws:sqs/queue:Queue public_queue https://queue.amazonaws.com/80398EXAMPLE/MyQueue
// ```
type Queue struct {
	pulumi.CustomResourceState

	// The ARN of the SQS queue
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication pulumi.BoolPtrOutput `pulumi:"contentBasedDeduplication"`
	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds pulumi.IntPtrOutput `pulumi:"delaySeconds"`
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue pulumi.BoolPtrOutput `pulumi:"fifoQueue"`
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds pulumi.IntOutput `pulumi:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId pulumi.StringPtrOutput `pulumi:"kmsMasterKeyId"`
	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize pulumi.IntPtrOutput `pulumi:"maxMessageSize"`
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds pulumi.IntPtrOutput `pulumi:"messageRetentionSeconds"`
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The JSON policy for the SQS queue.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds pulumi.IntPtrOutput `pulumi:"receiveWaitTimeSeconds"`
	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy pulumi.StringPtrOutput `pulumi:"redrivePolicy"`
	// A map of tags to assign to the queue.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds pulumi.IntPtrOutput `pulumi:"visibilityTimeoutSeconds"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		args = &QueueArgs{}
	}

	var resource Queue
	err := ctx.RegisterResource("aws:sqs/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("aws:sqs/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The ARN of the SQS queue
	Arn *string `pulumi:"arn"`
	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds *int `pulumi:"delaySeconds"`
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds *int `pulumi:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId *string `pulumi:"kmsMasterKeyId"`
	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize *int `pulumi:"maxMessageSize"`
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds *int `pulumi:"messageRetentionSeconds"`
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`
	NamePrefix *string `pulumi:"namePrefix"`
	// The JSON policy for the SQS queue.
	Policy *string `pulumi:"policy"`
	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds *int `pulumi:"receiveWaitTimeSeconds"`
	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy *string `pulumi:"redrivePolicy"`
	// A map of tags to assign to the queue.
	Tags map[string]string `pulumi:"tags"`
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

type QueueState struct {
	// The ARN of the SQS queue
	Arn pulumi.StringPtrInput
	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication pulumi.BoolPtrInput
	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds pulumi.IntPtrInput
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue pulumi.BoolPtrInput
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds pulumi.IntPtrInput
	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId pulumi.StringPtrInput
	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize pulumi.IntPtrInput
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds pulumi.IntPtrInput
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`
	NamePrefix pulumi.StringPtrInput
	// The JSON policy for the SQS queue.
	Policy pulumi.StringPtrInput
	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds pulumi.IntPtrInput
	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy pulumi.StringPtrInput
	// A map of tags to assign to the queue.
	Tags pulumi.StringMapInput
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds *int `pulumi:"delaySeconds"`
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds *int `pulumi:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId *string `pulumi:"kmsMasterKeyId"`
	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize *int `pulumi:"maxMessageSize"`
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds *int `pulumi:"messageRetentionSeconds"`
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`
	NamePrefix *string `pulumi:"namePrefix"`
	// The JSON policy for the SQS queue.
	Policy *string `pulumi:"policy"`
	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds *int `pulumi:"receiveWaitTimeSeconds"`
	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy *string `pulumi:"redrivePolicy"`
	// A map of tags to assign to the queue.
	Tags map[string]string `pulumi:"tags"`
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication pulumi.BoolPtrInput
	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds pulumi.IntPtrInput
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue pulumi.BoolPtrInput
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds pulumi.IntPtrInput
	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId pulumi.StringPtrInput
	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize pulumi.IntPtrInput
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds pulumi.IntPtrInput
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`
	NamePrefix pulumi.StringPtrInput
	// The JSON policy for the SQS queue.
	Policy pulumi.StringPtrInput
	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds pulumi.IntPtrInput
	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy pulumi.StringPtrInput
	// A map of tags to assign to the queue.
	Tags pulumi.StringMapInput
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

func (i *Queue) ToQueuePtrOutput() QueuePtrOutput {
	return i.ToQueuePtrOutputWithContext(context.Background())
}

func (i *Queue) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePtrOutput)
}

type QueuePtrInput interface {
	pulumi.Input

	ToQueuePtrOutput() QueuePtrOutput
	ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput
}

type queuePtrType QueueArgs

func (*queuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil))
}

func (i *queuePtrType) ToQueuePtrOutput() QueuePtrOutput {
	return i.ToQueuePtrOutputWithContext(context.Background())
}

func (i *queuePtrType) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePtrOutput)
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//          QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Queue)(nil))
}

func (i QueueArray) ToQueueArrayOutput() QueueArrayOutput {
	return i.ToQueueArrayOutputWithContext(context.Background())
}

func (i QueueArray) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueArrayOutput)
}

// QueueMapInput is an input type that accepts QueueMap and QueueMapOutput values.
// You can construct a concrete instance of `QueueMapInput` via:
//
//          QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Queue)(nil))
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

type QueueOutput struct {
	*pulumi.OutputState
}

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func (o QueueOutput) ToQueuePtrOutput() QueuePtrOutput {
	return o.ToQueuePtrOutputWithContext(context.Background())
}

func (o QueueOutput) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return o.ApplyT(func(v Queue) *Queue {
		return &v
	}).(QueuePtrOutput)
}

type QueuePtrOutput struct {
	*pulumi.OutputState
}

func (QueuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil))
}

func (o QueuePtrOutput) ToQueuePtrOutput() QueuePtrOutput {
	return o
}

func (o QueuePtrOutput) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return o
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Queue)(nil))
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Queue {
		return vs[0].([]Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Queue)(nil))
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Queue {
		return vs[0].(map[string]Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueuePtrOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
