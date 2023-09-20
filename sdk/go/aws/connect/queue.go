// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an Amazon Connect Queue resource. For more information see
// [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewQueue(ctx, "test", &connect.QueueArgs{
//				Description:        pulumi.String("Example Description"),
//				HoursOfOperationId: pulumi.String("12345678-1234-1234-1234-123456789012"),
//				InstanceId:         pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("Example Queue"),
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
// ### With Quick Connect IDs
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewQueue(ctx, "test", &connect.QueueArgs{
//				Description:        pulumi.String("Example Description"),
//				HoursOfOperationId: pulumi.String("12345678-1234-1234-1234-123456789012"),
//				InstanceId:         pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
//				QuickConnectIds: pulumi.StringArray{
//					pulumi.String("12345678-abcd-1234-abcd-123456789012"),
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("Example Queue with Quick Connect IDs"),
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
// ### With Outbound Caller Config
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewQueue(ctx, "test", &connect.QueueArgs{
//				Description:        pulumi.String("Example Description"),
//				HoursOfOperationId: pulumi.String("12345678-1234-1234-1234-123456789012"),
//				InstanceId:         pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
//				OutboundCallerConfig: &connect.QueueOutboundCallerConfigArgs{
//					OutboundCallerIdName:     pulumi.String("example"),
//					OutboundCallerIdNumberId: pulumi.String("12345678-abcd-1234-abcd-123456789012"),
//					OutboundFlowId:           pulumi.String("87654321-defg-1234-defg-987654321234"),
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("Example Queue with Outbound Caller Config"),
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
// Using `pulumi import`, import Amazon Connect Queues using the `instance_id` and `queue_id` separated by a colon (`:`). For example:
//
// ```sh
//
//	$ pulumi import aws:connect/queue:Queue example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
//
// ```
type Queue struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Queue.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies the description of the Queue.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the identifier of the Hours of Operation.
	HoursOfOperationId pulumi.StringOutput `pulumi:"hoursOfOperationId"`
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts pulumi.IntPtrOutput `pulumi:"maxContacts"`
	// Specifies the name of the Queue.
	Name pulumi.StringOutput `pulumi:"name"`
	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig QueueOutboundCallerConfigPtrOutput `pulumi:"outboundCallerConfig"`
	// The identifier for the Queue.
	QueueId pulumi.StringOutput `pulumi:"queueId"`
	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	QuickConnectIds pulumi.StringArrayOutput `pulumi:"quickConnectIds"`
	// Specifies the description of the Queue. Valid values are `ENABLED`, `DISABLED`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags to apply to the Queue. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HoursOfOperationId == nil {
		return nil, errors.New("invalid value for required argument 'HoursOfOperationId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Queue
	err := ctx.RegisterResource("aws:connect/queue:Queue", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:connect/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The Amazon Resource Name (ARN) of the Queue.
	Arn *string `pulumi:"arn"`
	// Specifies the description of the Queue.
	Description *string `pulumi:"description"`
	// Specifies the identifier of the Hours of Operation.
	HoursOfOperationId *string `pulumi:"hoursOfOperationId"`
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts *int `pulumi:"maxContacts"`
	// Specifies the name of the Queue.
	Name *string `pulumi:"name"`
	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig *QueueOutboundCallerConfig `pulumi:"outboundCallerConfig"`
	// The identifier for the Queue.
	QueueId *string `pulumi:"queueId"`
	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	QuickConnectIds []string `pulumi:"quickConnectIds"`
	// Specifies the description of the Queue. Valid values are `ENABLED`, `DISABLED`.
	Status *string `pulumi:"status"`
	// Tags to apply to the Queue. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type QueueState struct {
	// The Amazon Resource Name (ARN) of the Queue.
	Arn pulumi.StringPtrInput
	// Specifies the description of the Queue.
	Description pulumi.StringPtrInput
	// Specifies the identifier of the Hours of Operation.
	HoursOfOperationId pulumi.StringPtrInput
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId pulumi.StringPtrInput
	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts pulumi.IntPtrInput
	// Specifies the name of the Queue.
	Name pulumi.StringPtrInput
	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig QueueOutboundCallerConfigPtrInput
	// The identifier for the Queue.
	QueueId pulumi.StringPtrInput
	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	QuickConnectIds pulumi.StringArrayInput
	// Specifies the description of the Queue. Valid values are `ENABLED`, `DISABLED`.
	Status pulumi.StringPtrInput
	// Tags to apply to the Queue. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// Specifies the description of the Queue.
	Description *string `pulumi:"description"`
	// Specifies the identifier of the Hours of Operation.
	HoursOfOperationId string `pulumi:"hoursOfOperationId"`
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts *int `pulumi:"maxContacts"`
	// Specifies the name of the Queue.
	Name *string `pulumi:"name"`
	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig *QueueOutboundCallerConfig `pulumi:"outboundCallerConfig"`
	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	QuickConnectIds []string `pulumi:"quickConnectIds"`
	// Specifies the description of the Queue. Valid values are `ENABLED`, `DISABLED`.
	Status *string `pulumi:"status"`
	// Tags to apply to the Queue. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Specifies the description of the Queue.
	Description pulumi.StringPtrInput
	// Specifies the identifier of the Hours of Operation.
	HoursOfOperationId pulumi.StringInput
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId pulumi.StringInput
	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts pulumi.IntPtrInput
	// Specifies the name of the Queue.
	Name pulumi.StringPtrInput
	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig QueueOutboundCallerConfigPtrInput
	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	QuickConnectIds pulumi.StringArrayInput
	// Specifies the description of the Queue. Valid values are `ENABLED`, `DISABLED`.
	Status pulumi.StringPtrInput
	// Tags to apply to the Queue. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
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
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

func (i *Queue) ToOutput(ctx context.Context) pulumix.Output[*Queue] {
	return pulumix.Output[*Queue]{
		OutputState: i.ToQueueOutputWithContext(ctx).OutputState,
	}
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//	QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (i QueueArray) ToQueueArrayOutput() QueueArrayOutput {
	return i.ToQueueArrayOutputWithContext(context.Background())
}

func (i QueueArray) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueArrayOutput)
}

func (i QueueArray) ToOutput(ctx context.Context) pulumix.Output[[]*Queue] {
	return pulumix.Output[[]*Queue]{
		OutputState: i.ToQueueArrayOutputWithContext(ctx).OutputState,
	}
}

// QueueMapInput is an input type that accepts QueueMap and QueueMapOutput values.
// You can construct a concrete instance of `QueueMapInput` via:
//
//	QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

func (i QueueMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Queue] {
	return pulumix.Output[map[string]*Queue]{
		OutputState: i.ToQueueMapOutputWithContext(ctx).OutputState,
	}
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func (o QueueOutput) ToOutput(ctx context.Context) pulumix.Output[*Queue] {
	return pulumix.Output[*Queue]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the Queue.
func (o QueueOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies the description of the Queue.
func (o QueueOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the identifier of the Hours of Operation.
func (o QueueOutput) HoursOfOperationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.HoursOfOperationId }).(pulumi.StringOutput)
}

// Specifies the identifier of the hosting Amazon Connect Instance.
func (o QueueOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
func (o QueueOutput) MaxContacts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.MaxContacts }).(pulumi.IntPtrOutput)
}

// Specifies the name of the Queue.
func (o QueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
func (o QueueOutput) OutboundCallerConfig() QueueOutboundCallerConfigPtrOutput {
	return o.ApplyT(func(v *Queue) QueueOutboundCallerConfigPtrOutput { return v.OutboundCallerConfig }).(QueueOutboundCallerConfigPtrOutput)
}

// The identifier for the Queue.
func (o QueueOutput) QueueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.QueueId }).(pulumi.StringOutput)
}

// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
func (o QueueOutput) QuickConnectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringArrayOutput { return v.QuickConnectIds }).(pulumi.StringArrayOutput)
}

// Specifies the description of the Queue. Valid values are `ENABLED`, `DISABLED`.
func (o QueueOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags to apply to the Queue. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o QueueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o QueueOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Queue] {
	return pulumix.Output[[]*Queue]{
		OutputState: o.OutputState,
	}
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].([]*Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Queue] {
	return pulumix.Output[map[string]*Queue]{
		OutputState: o.OutputState,
	}
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].(map[string]*Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueInput)(nil)).Elem(), &Queue{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueArrayInput)(nil)).Elem(), QueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueMapInput)(nil)).Elem(), QueueMap{})
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
