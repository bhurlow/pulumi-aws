// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a S3 Bucket Notification Configuration. For additional information, see the [Configuring S3 Event Notifications section in the Amazon S3 Developer Guide](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html).
//
// > **NOTE:** S3 Buckets only support a single notification configuration. Declaring multiple `s3.BucketNotification` resources to the same S3 Bucket will cause a perpetual difference in configuration. See the example "Trigger multiple Lambda functions" for an option.
//
// ## Example Usage
// ### Add notification configuration to SNS Topic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
//			if err != nil {
//				return err
//			}
//			topicPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Principals: iam.GetPolicyDocumentStatementPrincipalArray{
//							&iam.GetPolicyDocumentStatementPrincipalArgs{
//								Type: pulumi.String("Service"),
//								Identifiers: pulumi.StringArray{
//									pulumi.String("s3.amazonaws.com"),
//								},
//							},
//						},
//						Actions: pulumi.StringArray{
//							pulumi.String("SNS:Publish"),
//						},
//						Resources: pulumi.StringArray{
//							pulumi.String("arn:aws:sns:*:*:s3-event-notification-topic"),
//						},
//						Conditions: iam.GetPolicyDocumentStatementConditionArray{
//							&iam.GetPolicyDocumentStatementConditionArgs{
//								Test:     pulumi.String("ArnLike"),
//								Variable: pulumi.String("aws:SourceArn"),
//								Values: pulumi.StringArray{
//									bucket.Arn,
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			topicTopic, err := sns.NewTopic(ctx, "topicTopic", &sns.TopicArgs{
//				Policy: topicPolicyDocument.ApplyT(func(topicPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
//					return &topicPolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
//				Bucket: bucket.ID(),
//				Topics: s3.BucketNotificationTopicArray{
//					&s3.BucketNotificationTopicArgs{
//						TopicArn: topicTopic.Arn,
//						Events: pulumi.StringArray{
//							pulumi.String("s3:ObjectCreated:*"),
//						},
//						FilterSuffix: pulumi.String(".log"),
//					},
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
// ### Add notification configuration to SQS Queue
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
//			if err != nil {
//				return err
//			}
//			queuePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Principals: iam.GetPolicyDocumentStatementPrincipalArray{
//							&iam.GetPolicyDocumentStatementPrincipalArgs{
//								Type: pulumi.String("*"),
//								Identifiers: pulumi.StringArray{
//									pulumi.String("*"),
//								},
//							},
//						},
//						Actions: pulumi.StringArray{
//							pulumi.String("sqs:SendMessage"),
//						},
//						Resources: pulumi.StringArray{
//							pulumi.String("arn:aws:sqs:*:*:s3-event-notification-queue"),
//						},
//						Conditions: iam.GetPolicyDocumentStatementConditionArray{
//							&iam.GetPolicyDocumentStatementConditionArgs{
//								Test:     pulumi.String("ArnEquals"),
//								Variable: pulumi.String("aws:SourceArn"),
//								Values: pulumi.StringArray{
//									bucket.Arn,
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			queueQueue, err := sqs.NewQueue(ctx, "queueQueue", &sqs.QueueArgs{
//				Policy: queuePolicyDocument.ApplyT(func(queuePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
//					return &queuePolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
//				Bucket: bucket.ID(),
//				Queues: s3.BucketNotificationQueueArray{
//					&s3.BucketNotificationQueueArgs{
//						QueueArn: queueQueue.Arn,
//						Events: pulumi.StringArray{
//							pulumi.String("s3:ObjectCreated:*"),
//						},
//						FilterSuffix: pulumi.String(".log"),
//					},
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
// ### Add notification configuration to Lambda Function
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"lambda.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			iamForLambda, err := iam.NewRole(ctx, "iamForLambda", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(assumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = lambda.NewFunction(ctx, "func", &lambda.FunctionArgs{
//				Code:    pulumi.NewFileArchive("your-function.zip"),
//				Role:    iamForLambda.Arn,
//				Handler: pulumi.String("exports.example"),
//				Runtime: pulumi.String("go1.x"),
//			})
//			if err != nil {
//				return err
//			}
//			bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
//			if err != nil {
//				return err
//			}
//			allowBucket, err := lambda.NewPermission(ctx, "allowBucket", &lambda.PermissionArgs{
//				Action:    pulumi.String("lambda:InvokeFunction"),
//				Function:  _func.Arn,
//				Principal: pulumi.String("s3.amazonaws.com"),
//				SourceArn: bucket.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
//				Bucket: bucket.ID(),
//				LambdaFunctions: s3.BucketNotificationLambdaFunctionArray{
//					&s3.BucketNotificationLambdaFunctionArgs{
//						LambdaFunctionArn: _func.Arn,
//						Events: pulumi.StringArray{
//							pulumi.String("s3:ObjectCreated:*"),
//						},
//						FilterPrefix: pulumi.String("AWSLogs/"),
//						FilterSuffix: pulumi.String(".log"),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				allowBucket,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Add multiple notification configurations to SQS Queue
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
//			if err != nil {
//				return err
//			}
//			queuePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Principals: iam.GetPolicyDocumentStatementPrincipalArray{
//							&iam.GetPolicyDocumentStatementPrincipalArgs{
//								Type: pulumi.String("*"),
//								Identifiers: pulumi.StringArray{
//									pulumi.String("*"),
//								},
//							},
//						},
//						Actions: pulumi.StringArray{
//							pulumi.String("sqs:SendMessage"),
//						},
//						Resources: pulumi.StringArray{
//							pulumi.String("arn:aws:sqs:*:*:s3-event-notification-queue"),
//						},
//						Conditions: iam.GetPolicyDocumentStatementConditionArray{
//							&iam.GetPolicyDocumentStatementConditionArgs{
//								Test:     pulumi.String("ArnEquals"),
//								Variable: pulumi.String("aws:SourceArn"),
//								Values: pulumi.StringArray{
//									bucket.Arn,
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			queueQueue, err := sqs.NewQueue(ctx, "queueQueue", &sqs.QueueArgs{
//				Policy: queuePolicyDocument.ApplyT(func(queuePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
//					return &queuePolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
//				Bucket: bucket.ID(),
//				Queues: s3.BucketNotificationQueueArray{
//					&s3.BucketNotificationQueueArgs{
//						Id:       pulumi.String("image-upload-event"),
//						QueueArn: queueQueue.Arn,
//						Events: pulumi.StringArray{
//							pulumi.String("s3:ObjectCreated:*"),
//						},
//						FilterPrefix: pulumi.String("images/"),
//					},
//					&s3.BucketNotificationQueueArgs{
//						Id:       pulumi.String("video-upload-event"),
//						QueueArn: queueQueue.Arn,
//						Events: pulumi.StringArray{
//							pulumi.String("s3:ObjectCreated:*"),
//						},
//						FilterPrefix: pulumi.String("videos/"),
//					},
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
// For JSON syntax, use an array instead of defining the `queue` key twice.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// ### Emit events to EventBridge
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
//				Bucket:      bucket.ID(),
//				Eventbridge: pulumi.Bool(true),
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
//	to = aws_s3_bucket_notification.bucket_notification
//
//	id = "bucket-name" } Using `pulumi import`, import S3 bucket notification using the `bucket`. For exampleconsole % pulumi import aws_s3_bucket_notification.bucket_notification bucket-name
type BucketNotification struct {
	pulumi.CustomResourceState

	// Name of the bucket for notification configuration.
	//
	// The following arguments are optional:
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Whether to enable Amazon EventBridge notifications. Defaults to `false`.
	Eventbridge pulumi.BoolPtrOutput `pulumi:"eventbridge"`
	// Used to configure notifications to a Lambda Function. See below.
	LambdaFunctions BucketNotificationLambdaFunctionArrayOutput `pulumi:"lambdaFunctions"`
	// Notification configuration to SQS Queue. See below.
	Queues BucketNotificationQueueArrayOutput `pulumi:"queues"`
	// Notification configuration to SNS Topic. See below.
	Topics BucketNotificationTopicArrayOutput `pulumi:"topics"`
}

// NewBucketNotification registers a new resource with the given unique name, arguments, and options.
func NewBucketNotification(ctx *pulumi.Context,
	name string, args *BucketNotificationArgs, opts ...pulumi.ResourceOption) (*BucketNotification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketNotification
	err := ctx.RegisterResource("aws:s3/bucketNotification:BucketNotification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketNotification gets an existing BucketNotification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketNotificationState, opts ...pulumi.ResourceOption) (*BucketNotification, error) {
	var resource BucketNotification
	err := ctx.ReadResource("aws:s3/bucketNotification:BucketNotification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketNotification resources.
type bucketNotificationState struct {
	// Name of the bucket for notification configuration.
	//
	// The following arguments are optional:
	Bucket *string `pulumi:"bucket"`
	// Whether to enable Amazon EventBridge notifications. Defaults to `false`.
	Eventbridge *bool `pulumi:"eventbridge"`
	// Used to configure notifications to a Lambda Function. See below.
	LambdaFunctions []BucketNotificationLambdaFunction `pulumi:"lambdaFunctions"`
	// Notification configuration to SQS Queue. See below.
	Queues []BucketNotificationQueue `pulumi:"queues"`
	// Notification configuration to SNS Topic. See below.
	Topics []BucketNotificationTopic `pulumi:"topics"`
}

type BucketNotificationState struct {
	// Name of the bucket for notification configuration.
	//
	// The following arguments are optional:
	Bucket pulumi.StringPtrInput
	// Whether to enable Amazon EventBridge notifications. Defaults to `false`.
	Eventbridge pulumi.BoolPtrInput
	// Used to configure notifications to a Lambda Function. See below.
	LambdaFunctions BucketNotificationLambdaFunctionArrayInput
	// Notification configuration to SQS Queue. See below.
	Queues BucketNotificationQueueArrayInput
	// Notification configuration to SNS Topic. See below.
	Topics BucketNotificationTopicArrayInput
}

func (BucketNotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketNotificationState)(nil)).Elem()
}

type bucketNotificationArgs struct {
	// Name of the bucket for notification configuration.
	//
	// The following arguments are optional:
	Bucket string `pulumi:"bucket"`
	// Whether to enable Amazon EventBridge notifications. Defaults to `false`.
	Eventbridge *bool `pulumi:"eventbridge"`
	// Used to configure notifications to a Lambda Function. See below.
	LambdaFunctions []BucketNotificationLambdaFunction `pulumi:"lambdaFunctions"`
	// Notification configuration to SQS Queue. See below.
	Queues []BucketNotificationQueue `pulumi:"queues"`
	// Notification configuration to SNS Topic. See below.
	Topics []BucketNotificationTopic `pulumi:"topics"`
}

// The set of arguments for constructing a BucketNotification resource.
type BucketNotificationArgs struct {
	// Name of the bucket for notification configuration.
	//
	// The following arguments are optional:
	Bucket pulumi.StringInput
	// Whether to enable Amazon EventBridge notifications. Defaults to `false`.
	Eventbridge pulumi.BoolPtrInput
	// Used to configure notifications to a Lambda Function. See below.
	LambdaFunctions BucketNotificationLambdaFunctionArrayInput
	// Notification configuration to SQS Queue. See below.
	Queues BucketNotificationQueueArrayInput
	// Notification configuration to SNS Topic. See below.
	Topics BucketNotificationTopicArrayInput
}

func (BucketNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketNotificationArgs)(nil)).Elem()
}

type BucketNotificationInput interface {
	pulumi.Input

	ToBucketNotificationOutput() BucketNotificationOutput
	ToBucketNotificationOutputWithContext(ctx context.Context) BucketNotificationOutput
}

func (*BucketNotification) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketNotification)(nil)).Elem()
}

func (i *BucketNotification) ToBucketNotificationOutput() BucketNotificationOutput {
	return i.ToBucketNotificationOutputWithContext(context.Background())
}

func (i *BucketNotification) ToBucketNotificationOutputWithContext(ctx context.Context) BucketNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketNotificationOutput)
}

// BucketNotificationArrayInput is an input type that accepts BucketNotificationArray and BucketNotificationArrayOutput values.
// You can construct a concrete instance of `BucketNotificationArrayInput` via:
//
//	BucketNotificationArray{ BucketNotificationArgs{...} }
type BucketNotificationArrayInput interface {
	pulumi.Input

	ToBucketNotificationArrayOutput() BucketNotificationArrayOutput
	ToBucketNotificationArrayOutputWithContext(context.Context) BucketNotificationArrayOutput
}

type BucketNotificationArray []BucketNotificationInput

func (BucketNotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketNotification)(nil)).Elem()
}

func (i BucketNotificationArray) ToBucketNotificationArrayOutput() BucketNotificationArrayOutput {
	return i.ToBucketNotificationArrayOutputWithContext(context.Background())
}

func (i BucketNotificationArray) ToBucketNotificationArrayOutputWithContext(ctx context.Context) BucketNotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketNotificationArrayOutput)
}

// BucketNotificationMapInput is an input type that accepts BucketNotificationMap and BucketNotificationMapOutput values.
// You can construct a concrete instance of `BucketNotificationMapInput` via:
//
//	BucketNotificationMap{ "key": BucketNotificationArgs{...} }
type BucketNotificationMapInput interface {
	pulumi.Input

	ToBucketNotificationMapOutput() BucketNotificationMapOutput
	ToBucketNotificationMapOutputWithContext(context.Context) BucketNotificationMapOutput
}

type BucketNotificationMap map[string]BucketNotificationInput

func (BucketNotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketNotification)(nil)).Elem()
}

func (i BucketNotificationMap) ToBucketNotificationMapOutput() BucketNotificationMapOutput {
	return i.ToBucketNotificationMapOutputWithContext(context.Background())
}

func (i BucketNotificationMap) ToBucketNotificationMapOutputWithContext(ctx context.Context) BucketNotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketNotificationMapOutput)
}

type BucketNotificationOutput struct{ *pulumi.OutputState }

func (BucketNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketNotification)(nil)).Elem()
}

func (o BucketNotificationOutput) ToBucketNotificationOutput() BucketNotificationOutput {
	return o
}

func (o BucketNotificationOutput) ToBucketNotificationOutputWithContext(ctx context.Context) BucketNotificationOutput {
	return o
}

// Name of the bucket for notification configuration.
//
// The following arguments are optional:
func (o BucketNotificationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketNotification) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Whether to enable Amazon EventBridge notifications. Defaults to `false`.
func (o BucketNotificationOutput) Eventbridge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketNotification) pulumi.BoolPtrOutput { return v.Eventbridge }).(pulumi.BoolPtrOutput)
}

// Used to configure notifications to a Lambda Function. See below.
func (o BucketNotificationOutput) LambdaFunctions() BucketNotificationLambdaFunctionArrayOutput {
	return o.ApplyT(func(v *BucketNotification) BucketNotificationLambdaFunctionArrayOutput { return v.LambdaFunctions }).(BucketNotificationLambdaFunctionArrayOutput)
}

// Notification configuration to SQS Queue. See below.
func (o BucketNotificationOutput) Queues() BucketNotificationQueueArrayOutput {
	return o.ApplyT(func(v *BucketNotification) BucketNotificationQueueArrayOutput { return v.Queues }).(BucketNotificationQueueArrayOutput)
}

// Notification configuration to SNS Topic. See below.
func (o BucketNotificationOutput) Topics() BucketNotificationTopicArrayOutput {
	return o.ApplyT(func(v *BucketNotification) BucketNotificationTopicArrayOutput { return v.Topics }).(BucketNotificationTopicArrayOutput)
}

type BucketNotificationArrayOutput struct{ *pulumi.OutputState }

func (BucketNotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketNotification)(nil)).Elem()
}

func (o BucketNotificationArrayOutput) ToBucketNotificationArrayOutput() BucketNotificationArrayOutput {
	return o
}

func (o BucketNotificationArrayOutput) ToBucketNotificationArrayOutputWithContext(ctx context.Context) BucketNotificationArrayOutput {
	return o
}

func (o BucketNotificationArrayOutput) Index(i pulumi.IntInput) BucketNotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketNotification {
		return vs[0].([]*BucketNotification)[vs[1].(int)]
	}).(BucketNotificationOutput)
}

type BucketNotificationMapOutput struct{ *pulumi.OutputState }

func (BucketNotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketNotification)(nil)).Elem()
}

func (o BucketNotificationMapOutput) ToBucketNotificationMapOutput() BucketNotificationMapOutput {
	return o
}

func (o BucketNotificationMapOutput) ToBucketNotificationMapOutputWithContext(ctx context.Context) BucketNotificationMapOutput {
	return o
}

func (o BucketNotificationMapOutput) MapIndex(k pulumi.StringInput) BucketNotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketNotification {
		return vs[0].(map[string]*BucketNotification)[vs[1].(string)]
	}).(BucketNotificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketNotificationInput)(nil)).Elem(), &BucketNotification{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketNotificationArrayInput)(nil)).Elem(), BucketNotificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketNotificationMapInput)(nil)).Elem(), BucketNotificationMap{})
	pulumi.RegisterOutputType(BucketNotificationOutput{})
	pulumi.RegisterOutputType(BucketNotificationArrayOutput{})
	pulumi.RegisterOutputType(BucketNotificationMapOutput{})
}
