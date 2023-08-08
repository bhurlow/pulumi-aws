// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Glue Trigger resource.
//
// ## Example Usage
// ### Conditional Trigger
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
//				Type: pulumi.String("CONDITIONAL"),
//				Actions: glue.TriggerActionArray{
//					&glue.TriggerActionArgs{
//						JobName: pulumi.Any(aws_glue_job.Example1.Name),
//					},
//				},
//				Predicate: &glue.TriggerPredicateArgs{
//					Conditions: glue.TriggerPredicateConditionArray{
//						&glue.TriggerPredicateConditionArgs{
//							JobName: pulumi.Any(aws_glue_job.Example2.Name),
//							State:   pulumi.String("SUCCEEDED"),
//						},
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
// ### On-Demand Trigger
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
//				Type: pulumi.String("ON_DEMAND"),
//				Actions: glue.TriggerActionArray{
//					&glue.TriggerActionArgs{
//						JobName: pulumi.Any(aws_glue_job.Example.Name),
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
// ### Scheduled Trigger
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
//				Schedule: pulumi.String("cron(15 12 * * ? *)"),
//				Type:     pulumi.String("SCHEDULED"),
//				Actions: glue.TriggerActionArray{
//					&glue.TriggerActionArgs{
//						JobName: pulumi.Any(aws_glue_job.Example.Name),
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
// ### Conditional Trigger with Crawler Action
//
// **Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
//				Type: pulumi.String("CONDITIONAL"),
//				Actions: glue.TriggerActionArray{
//					&glue.TriggerActionArgs{
//						CrawlerName: pulumi.Any(aws_glue_crawler.Example1.Name),
//					},
//				},
//				Predicate: &glue.TriggerPredicateArgs{
//					Conditions: glue.TriggerPredicateConditionArray{
//						&glue.TriggerPredicateConditionArgs{
//							JobName: pulumi.Any(aws_glue_job.Example2.Name),
//							State:   pulumi.String("SUCCEEDED"),
//						},
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
// ### Conditional Trigger with Crawler Condition
//
// **Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
//				Type: pulumi.String("CONDITIONAL"),
//				Actions: glue.TriggerActionArray{
//					&glue.TriggerActionArgs{
//						JobName: pulumi.Any(aws_glue_job.Example1.Name),
//					},
//				},
//				Predicate: &glue.TriggerPredicateArgs{
//					Conditions: glue.TriggerPredicateConditionArray{
//						&glue.TriggerPredicateConditionArgs{
//							CrawlerName: pulumi.Any(aws_glue_crawler.Example2.Name),
//							CrawlState:  pulumi.String("SUCCEEDED"),
//						},
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
// ## Import
//
// terraform import {
//
//	to = aws_glue_trigger.MyTrigger
//
//	id = "MyTrigger" } Using `pulumi import`, import Glue Triggers using `name`. For exampleconsole % pulumi import aws_glue_trigger.MyTrigger MyTrigger
type Trigger struct {
	pulumi.CustomResourceState

	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions TriggerActionArrayOutput `pulumi:"actions"`
	// Amazon Resource Name (ARN) of Glue Trigger
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the new trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Start the trigger. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingConditions TriggerEventBatchingConditionArrayOutput `pulumi:"eventBatchingConditions"`
	// The name of the trigger.
	Name pulumi.StringOutput `pulumi:"name"`
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate TriggerPredicatePtrOutput `pulumi:"predicate"`
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.
	StartOnCreation pulumi.BoolPtrOutput `pulumi:"startOnCreation"`
	// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `jobName` must also be specified. Conflicts with `crawlerState`.
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of trigger. Valid values are `CONDITIONAL`, `EVENT`, `ON_DEMAND`, and `SCHEDULED`.
	Type pulumi.StringOutput `pulumi:"type"`
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName pulumi.StringPtrOutput `pulumi:"workflowName"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trigger
	err := ctx.RegisterResource("aws:glue/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("aws:glue/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions []TriggerAction `pulumi:"actions"`
	// Amazon Resource Name (ARN) of Glue Trigger
	Arn *string `pulumi:"arn"`
	// A description of the new trigger.
	Description *string `pulumi:"description"`
	// Start the trigger. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingConditions []TriggerEventBatchingCondition `pulumi:"eventBatchingConditions"`
	// The name of the trigger.
	Name *string `pulumi:"name"`
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate *TriggerPredicate `pulumi:"predicate"`
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule *string `pulumi:"schedule"`
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.
	StartOnCreation *bool `pulumi:"startOnCreation"`
	// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `jobName` must also be specified. Conflicts with `crawlerState`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of trigger. Valid values are `CONDITIONAL`, `EVENT`, `ON_DEMAND`, and `SCHEDULED`.
	Type *string `pulumi:"type"`
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName *string `pulumi:"workflowName"`
}

type TriggerState struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions TriggerActionArrayInput
	// Amazon Resource Name (ARN) of Glue Trigger
	Arn pulumi.StringPtrInput
	// A description of the new trigger.
	Description pulumi.StringPtrInput
	// Start the trigger. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingConditions TriggerEventBatchingConditionArrayInput
	// The name of the trigger.
	Name pulumi.StringPtrInput
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate TriggerPredicatePtrInput
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule pulumi.StringPtrInput
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.
	StartOnCreation pulumi.BoolPtrInput
	// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `jobName` must also be specified. Conflicts with `crawlerState`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The type of trigger. Valid values are `CONDITIONAL`, `EVENT`, `ON_DEMAND`, and `SCHEDULED`.
	Type pulumi.StringPtrInput
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions []TriggerAction `pulumi:"actions"`
	// A description of the new trigger.
	Description *string `pulumi:"description"`
	// Start the trigger. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingConditions []TriggerEventBatchingCondition `pulumi:"eventBatchingConditions"`
	// The name of the trigger.
	Name *string `pulumi:"name"`
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate *TriggerPredicate `pulumi:"predicate"`
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule *string `pulumi:"schedule"`
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.
	StartOnCreation *bool `pulumi:"startOnCreation"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The type of trigger. Valid values are `CONDITIONAL`, `EVENT`, `ON_DEMAND`, and `SCHEDULED`.
	Type string `pulumi:"type"`
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName *string `pulumi:"workflowName"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions TriggerActionArrayInput
	// A description of the new trigger.
	Description pulumi.StringPtrInput
	// Start the trigger. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
	EventBatchingConditions TriggerEventBatchingConditionArrayInput
	// The name of the trigger.
	Name pulumi.StringPtrInput
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate TriggerPredicatePtrInput
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule pulumi.StringPtrInput
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.
	StartOnCreation pulumi.BoolPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The type of trigger. Valid values are `CONDITIONAL`, `EVENT`, `ON_DEMAND`, and `SCHEDULED`.
	Type pulumi.StringInput
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName pulumi.StringPtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

// TriggerArrayInput is an input type that accepts TriggerArray and TriggerArrayOutput values.
// You can construct a concrete instance of `TriggerArrayInput` via:
//
//	TriggerArray{ TriggerArgs{...} }
type TriggerArrayInput interface {
	pulumi.Input

	ToTriggerArrayOutput() TriggerArrayOutput
	ToTriggerArrayOutputWithContext(context.Context) TriggerArrayOutput
}

type TriggerArray []TriggerInput

func (TriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (i TriggerArray) ToTriggerArrayOutput() TriggerArrayOutput {
	return i.ToTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerArray) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerArrayOutput)
}

// TriggerMapInput is an input type that accepts TriggerMap and TriggerMapOutput values.
// You can construct a concrete instance of `TriggerMapInput` via:
//
//	TriggerMap{ "key": TriggerArgs{...} }
type TriggerMapInput interface {
	pulumi.Input

	ToTriggerMapOutput() TriggerMapOutput
	ToTriggerMapOutputWithContext(context.Context) TriggerMapOutput
}

type TriggerMap map[string]TriggerInput

func (TriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (i TriggerMap) ToTriggerMapOutput() TriggerMapOutput {
	return i.ToTriggerMapOutputWithContext(context.Background())
}

func (i TriggerMap) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMapOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

// List of actions initiated by this trigger when it fires. See Actions Below.
func (o TriggerOutput) Actions() TriggerActionArrayOutput {
	return o.ApplyT(func(v *Trigger) TriggerActionArrayOutput { return v.Actions }).(TriggerActionArrayOutput)
}

// Amazon Resource Name (ARN) of Glue Trigger
func (o TriggerOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A description of the new trigger.
func (o TriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Start the trigger. Defaults to `true`.
func (o TriggerOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires. See Event Batching Condition.
func (o TriggerOutput) EventBatchingConditions() TriggerEventBatchingConditionArrayOutput {
	return o.ApplyT(func(v *Trigger) TriggerEventBatchingConditionArrayOutput { return v.EventBatchingConditions }).(TriggerEventBatchingConditionArrayOutput)
}

// The name of the trigger.
func (o TriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
func (o TriggerOutput) Predicate() TriggerPredicatePtrOutput {
	return o.ApplyT(func(v *Trigger) TriggerPredicatePtrOutput { return v.Predicate }).(TriggerPredicatePtrOutput)
}

// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
func (o TriggerOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Schedule }).(pulumi.StringPtrOutput)
}

// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.
func (o TriggerOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.BoolPtrOutput { return v.StartOnCreation }).(pulumi.BoolPtrOutput)
}

// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `jobName` must also be specified. Conflicts with `crawlerState`.
func (o TriggerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TriggerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o TriggerOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of trigger. Valid values are `CONDITIONAL`, `EVENT`, `ON_DEMAND`, and `SCHEDULED`.
func (o TriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
func (o TriggerOutput) WorkflowName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.WorkflowName }).(pulumi.StringPtrOutput)
}

type TriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (o TriggerArrayOutput) ToTriggerArrayOutput() TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) Index(i pulumi.IntInput) TriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].([]*Trigger)[vs[1].(int)]
	}).(TriggerOutput)
}

type TriggerMapOutput struct{ *pulumi.OutputState }

func (TriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (o TriggerMapOutput) ToTriggerMapOutput() TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) MapIndex(k pulumi.StringInput) TriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].(map[string]*Trigger)[vs[1].(string)]
	}).(TriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerArrayInput)(nil)).Elem(), TriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerMapInput)(nil)).Elem(), TriggerMap{})
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerMapOutput{})
}
