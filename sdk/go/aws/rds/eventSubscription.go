// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a DB event subscription resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultInstance, err := rds.NewInstance(ctx, "defaultInstance", &rds.InstanceArgs{
//				AllocatedStorage:   pulumi.Int(10),
//				Engine:             pulumi.String("mysql"),
//				EngineVersion:      pulumi.String("5.6.17"),
//				InstanceClass:      pulumi.String("db.t2.micro"),
//				DbName:             pulumi.String("mydb"),
//				Username:           pulumi.String("foo"),
//				Password:           pulumi.String("bar"),
//				DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
//				ParameterGroupName: pulumi.String("default.mysql5.6"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewEventSubscription(ctx, "defaultEventSubscription", &rds.EventSubscriptionArgs{
//				SnsTopic:   defaultTopic.Arn,
//				SourceType: pulumi.String("db-instance"),
//				SourceIds: pulumi.StringArray{
//					defaultInstance.Identifier,
//				},
//				EventCategories: pulumi.StringArray{
//					pulumi.String("availability"),
//					pulumi.String("deletion"),
//					pulumi.String("failover"),
//					pulumi.String("failure"),
//					pulumi.String("low storage"),
//					pulumi.String("maintenance"),
//					pulumi.String("notification"),
//					pulumi.String("read replica"),
//					pulumi.String("recovery"),
//					pulumi.String("restoration"),
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
// Using `pulumi import`, import DB Event Subscriptions using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:rds/eventSubscription:EventSubscription default rds-event-sub
//
// ```
type EventSubscription struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the RDS event notification subscription
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The AWS customer account associated with the RDS event notification subscription
	CustomerAwsId pulumi.StringOutput `pulumi:"customerAwsId"`
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
	EventCategories pulumi.StringArrayOutput `pulumi:"eventCategories"`
	// The name of the DB event subscription. By default generated by this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the DB event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The SNS topic to send events to.
	SnsTopic pulumi.StringOutput `pulumi:"snsTopic"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds pulumi.StringArrayOutput `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SnsTopic == nil {
		return nil, errors.New("invalid value for required argument 'SnsTopic'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventSubscription
	err := ctx.RegisterResource("aws:rds/eventSubscription:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("aws:rds/eventSubscription:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
	// The Amazon Resource Name of the RDS event notification subscription
	Arn *string `pulumi:"arn"`
	// The AWS customer account associated with the RDS event notification subscription
	CustomerAwsId *string `pulumi:"customerAwsId"`
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
	EventCategories []string `pulumi:"eventCategories"`
	// The name of the DB event subscription. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name of the DB event subscription. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The SNS topic to send events to.
	SnsTopic *string `pulumi:"snsTopic"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds []string `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
	SourceType *string `pulumi:"sourceType"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type EventSubscriptionState struct {
	// The Amazon Resource Name of the RDS event notification subscription
	Arn pulumi.StringPtrInput
	// The AWS customer account associated with the RDS event notification subscription
	CustomerAwsId pulumi.StringPtrInput
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
	EventCategories pulumi.StringArrayInput
	// The name of the DB event subscription. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name of the DB event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The SNS topic to send events to.
	SnsTopic pulumi.StringPtrInput
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds pulumi.StringArrayInput
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
	EventCategories []string `pulumi:"eventCategories"`
	// The name of the DB event subscription. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name of the DB event subscription. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The SNS topic to send events to.
	SnsTopic string `pulumi:"snsTopic"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds []string `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
	SourceType *string `pulumi:"sourceType"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
	EventCategories pulumi.StringArrayInput
	// The name of the DB event subscription. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name of the DB event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The SNS topic to send events to.
	SnsTopic pulumi.StringInput
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds pulumi.StringArrayInput
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

type EventSubscriptionInput interface {
	pulumi.Input

	ToEventSubscriptionOutput() EventSubscriptionOutput
	ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput
}

func (*EventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (i *EventSubscription) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return i.ToEventSubscriptionOutputWithContext(context.Background())
}

func (i *EventSubscription) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionOutput)
}

func (i *EventSubscription) ToOutput(ctx context.Context) pulumix.Output[*EventSubscription] {
	return pulumix.Output[*EventSubscription]{
		OutputState: i.ToEventSubscriptionOutputWithContext(ctx).OutputState,
	}
}

// EventSubscriptionArrayInput is an input type that accepts EventSubscriptionArray and EventSubscriptionArrayOutput values.
// You can construct a concrete instance of `EventSubscriptionArrayInput` via:
//
//	EventSubscriptionArray{ EventSubscriptionArgs{...} }
type EventSubscriptionArrayInput interface {
	pulumi.Input

	ToEventSubscriptionArrayOutput() EventSubscriptionArrayOutput
	ToEventSubscriptionArrayOutputWithContext(context.Context) EventSubscriptionArrayOutput
}

type EventSubscriptionArray []EventSubscriptionInput

func (EventSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventSubscription)(nil)).Elem()
}

func (i EventSubscriptionArray) ToEventSubscriptionArrayOutput() EventSubscriptionArrayOutput {
	return i.ToEventSubscriptionArrayOutputWithContext(context.Background())
}

func (i EventSubscriptionArray) ToEventSubscriptionArrayOutputWithContext(ctx context.Context) EventSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionArrayOutput)
}

func (i EventSubscriptionArray) ToOutput(ctx context.Context) pulumix.Output[[]*EventSubscription] {
	return pulumix.Output[[]*EventSubscription]{
		OutputState: i.ToEventSubscriptionArrayOutputWithContext(ctx).OutputState,
	}
}

// EventSubscriptionMapInput is an input type that accepts EventSubscriptionMap and EventSubscriptionMapOutput values.
// You can construct a concrete instance of `EventSubscriptionMapInput` via:
//
//	EventSubscriptionMap{ "key": EventSubscriptionArgs{...} }
type EventSubscriptionMapInput interface {
	pulumi.Input

	ToEventSubscriptionMapOutput() EventSubscriptionMapOutput
	ToEventSubscriptionMapOutputWithContext(context.Context) EventSubscriptionMapOutput
}

type EventSubscriptionMap map[string]EventSubscriptionInput

func (EventSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventSubscription)(nil)).Elem()
}

func (i EventSubscriptionMap) ToEventSubscriptionMapOutput() EventSubscriptionMapOutput {
	return i.ToEventSubscriptionMapOutputWithContext(context.Background())
}

func (i EventSubscriptionMap) ToEventSubscriptionMapOutputWithContext(ctx context.Context) EventSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionMapOutput)
}

func (i EventSubscriptionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EventSubscription] {
	return pulumix.Output[map[string]*EventSubscription]{
		OutputState: i.ToEventSubscriptionMapOutputWithContext(ctx).OutputState,
	}
}

type EventSubscriptionOutput struct{ *pulumi.OutputState }

func (EventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*EventSubscription] {
	return pulumix.Output[*EventSubscription]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name of the RDS event notification subscription
func (o EventSubscriptionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The AWS customer account associated with the RDS event notification subscription
func (o EventSubscriptionOutput) CustomerAwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.CustomerAwsId }).(pulumi.StringOutput)
}

// A boolean flag to enable/disable the subscription. Defaults to true.
func (o EventSubscriptionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
func (o EventSubscriptionOutput) EventCategories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringArrayOutput { return v.EventCategories }).(pulumi.StringArrayOutput)
}

// The name of the DB event subscription. By default generated by this provider.
func (o EventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the DB event subscription. Conflicts with `name`.
func (o EventSubscriptionOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// The SNS topic to send events to.
func (o EventSubscriptionOutput) SnsTopic() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.SnsTopic }).(pulumi.StringOutput)
}

// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
func (o EventSubscriptionOutput) SourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringArrayOutput { return v.SourceIds }).(pulumi.StringArrayOutput)
}

// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
func (o EventSubscriptionOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EventSubscriptionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o EventSubscriptionOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type EventSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (EventSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventSubscription)(nil)).Elem()
}

func (o EventSubscriptionArrayOutput) ToEventSubscriptionArrayOutput() EventSubscriptionArrayOutput {
	return o
}

func (o EventSubscriptionArrayOutput) ToEventSubscriptionArrayOutputWithContext(ctx context.Context) EventSubscriptionArrayOutput {
	return o
}

func (o EventSubscriptionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EventSubscription] {
	return pulumix.Output[[]*EventSubscription]{
		OutputState: o.OutputState,
	}
}

func (o EventSubscriptionArrayOutput) Index(i pulumi.IntInput) EventSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventSubscription {
		return vs[0].([]*EventSubscription)[vs[1].(int)]
	}).(EventSubscriptionOutput)
}

type EventSubscriptionMapOutput struct{ *pulumi.OutputState }

func (EventSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventSubscription)(nil)).Elem()
}

func (o EventSubscriptionMapOutput) ToEventSubscriptionMapOutput() EventSubscriptionMapOutput {
	return o
}

func (o EventSubscriptionMapOutput) ToEventSubscriptionMapOutputWithContext(ctx context.Context) EventSubscriptionMapOutput {
	return o
}

func (o EventSubscriptionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EventSubscription] {
	return pulumix.Output[map[string]*EventSubscription]{
		OutputState: o.OutputState,
	}
}

func (o EventSubscriptionMapOutput) MapIndex(k pulumi.StringInput) EventSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventSubscription {
		return vs[0].(map[string]*EventSubscription)[vs[1].(string)]
	}).(EventSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionInput)(nil)).Elem(), &EventSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionArrayInput)(nil)).Elem(), EventSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionMapInput)(nil)).Elem(), EventSubscriptionMap{})
	pulumi.RegisterOutputType(EventSubscriptionOutput{})
	pulumi.RegisterOutputType(EventSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(EventSubscriptionMapOutput{})
}
