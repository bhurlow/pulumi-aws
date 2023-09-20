// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an Amazon AppIntegrations Event Integration resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appconfig.NewEventIntegration(ctx, "example", &appconfig.EventIntegrationArgs{
//				Description: pulumi.String("Example Description"),
//				EventFilter: &appconfig.EventIntegrationEventFilterArgs{
//					Source: pulumi.String("aws.partner/examplepartner.com"),
//				},
//				EventbridgeBus: pulumi.String("default"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("Example Event Integration"),
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
// Using `pulumi import`, import Amazon AppIntegrations Event Integrations using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:appconfig/eventIntegration:EventIntegration example example-name
//
// ```
type EventIntegration struct {
	pulumi.CustomResourceState

	// ARN of the Event Integration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the Event Integration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter EventIntegrationEventFilterOutput `pulumi:"eventFilter"`
	// EventBridge bus.
	EventbridgeBus pulumi.StringOutput `pulumi:"eventbridgeBus"`
	// Name of the Event Integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags to apply to the Event Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewEventIntegration registers a new resource with the given unique name, arguments, and options.
func NewEventIntegration(ctx *pulumi.Context,
	name string, args *EventIntegrationArgs, opts ...pulumi.ResourceOption) (*EventIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventFilter == nil {
		return nil, errors.New("invalid value for required argument 'EventFilter'")
	}
	if args.EventbridgeBus == nil {
		return nil, errors.New("invalid value for required argument 'EventbridgeBus'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventIntegration
	err := ctx.RegisterResource("aws:appconfig/eventIntegration:EventIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventIntegration gets an existing EventIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventIntegrationState, opts ...pulumi.ResourceOption) (*EventIntegration, error) {
	var resource EventIntegration
	err := ctx.ReadResource("aws:appconfig/eventIntegration:EventIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventIntegration resources.
type eventIntegrationState struct {
	// ARN of the Event Integration.
	Arn *string `pulumi:"arn"`
	// Description of the Event Integration.
	Description *string `pulumi:"description"`
	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter *EventIntegrationEventFilter `pulumi:"eventFilter"`
	// EventBridge bus.
	EventbridgeBus *string `pulumi:"eventbridgeBus"`
	// Name of the Event Integration.
	Name *string `pulumi:"name"`
	// Tags to apply to the Event Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type EventIntegrationState struct {
	// ARN of the Event Integration.
	Arn pulumi.StringPtrInput
	// Description of the Event Integration.
	Description pulumi.StringPtrInput
	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter EventIntegrationEventFilterPtrInput
	// EventBridge bus.
	EventbridgeBus pulumi.StringPtrInput
	// Name of the Event Integration.
	Name pulumi.StringPtrInput
	// Tags to apply to the Event Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (EventIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventIntegrationState)(nil)).Elem()
}

type eventIntegrationArgs struct {
	// Description of the Event Integration.
	Description *string `pulumi:"description"`
	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter EventIntegrationEventFilter `pulumi:"eventFilter"`
	// EventBridge bus.
	EventbridgeBus string `pulumi:"eventbridgeBus"`
	// Name of the Event Integration.
	Name *string `pulumi:"name"`
	// Tags to apply to the Event Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventIntegration resource.
type EventIntegrationArgs struct {
	// Description of the Event Integration.
	Description pulumi.StringPtrInput
	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter EventIntegrationEventFilterInput
	// EventBridge bus.
	EventbridgeBus pulumi.StringInput
	// Name of the Event Integration.
	Name pulumi.StringPtrInput
	// Tags to apply to the Event Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (EventIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventIntegrationArgs)(nil)).Elem()
}

type EventIntegrationInput interface {
	pulumi.Input

	ToEventIntegrationOutput() EventIntegrationOutput
	ToEventIntegrationOutputWithContext(ctx context.Context) EventIntegrationOutput
}

func (*EventIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**EventIntegration)(nil)).Elem()
}

func (i *EventIntegration) ToEventIntegrationOutput() EventIntegrationOutput {
	return i.ToEventIntegrationOutputWithContext(context.Background())
}

func (i *EventIntegration) ToEventIntegrationOutputWithContext(ctx context.Context) EventIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventIntegrationOutput)
}

func (i *EventIntegration) ToOutput(ctx context.Context) pulumix.Output[*EventIntegration] {
	return pulumix.Output[*EventIntegration]{
		OutputState: i.ToEventIntegrationOutputWithContext(ctx).OutputState,
	}
}

// EventIntegrationArrayInput is an input type that accepts EventIntegrationArray and EventIntegrationArrayOutput values.
// You can construct a concrete instance of `EventIntegrationArrayInput` via:
//
//	EventIntegrationArray{ EventIntegrationArgs{...} }
type EventIntegrationArrayInput interface {
	pulumi.Input

	ToEventIntegrationArrayOutput() EventIntegrationArrayOutput
	ToEventIntegrationArrayOutputWithContext(context.Context) EventIntegrationArrayOutput
}

type EventIntegrationArray []EventIntegrationInput

func (EventIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventIntegration)(nil)).Elem()
}

func (i EventIntegrationArray) ToEventIntegrationArrayOutput() EventIntegrationArrayOutput {
	return i.ToEventIntegrationArrayOutputWithContext(context.Background())
}

func (i EventIntegrationArray) ToEventIntegrationArrayOutputWithContext(ctx context.Context) EventIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventIntegrationArrayOutput)
}

func (i EventIntegrationArray) ToOutput(ctx context.Context) pulumix.Output[[]*EventIntegration] {
	return pulumix.Output[[]*EventIntegration]{
		OutputState: i.ToEventIntegrationArrayOutputWithContext(ctx).OutputState,
	}
}

// EventIntegrationMapInput is an input type that accepts EventIntegrationMap and EventIntegrationMapOutput values.
// You can construct a concrete instance of `EventIntegrationMapInput` via:
//
//	EventIntegrationMap{ "key": EventIntegrationArgs{...} }
type EventIntegrationMapInput interface {
	pulumi.Input

	ToEventIntegrationMapOutput() EventIntegrationMapOutput
	ToEventIntegrationMapOutputWithContext(context.Context) EventIntegrationMapOutput
}

type EventIntegrationMap map[string]EventIntegrationInput

func (EventIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventIntegration)(nil)).Elem()
}

func (i EventIntegrationMap) ToEventIntegrationMapOutput() EventIntegrationMapOutput {
	return i.ToEventIntegrationMapOutputWithContext(context.Background())
}

func (i EventIntegrationMap) ToEventIntegrationMapOutputWithContext(ctx context.Context) EventIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventIntegrationMapOutput)
}

func (i EventIntegrationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EventIntegration] {
	return pulumix.Output[map[string]*EventIntegration]{
		OutputState: i.ToEventIntegrationMapOutputWithContext(ctx).OutputState,
	}
}

type EventIntegrationOutput struct{ *pulumi.OutputState }

func (EventIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventIntegration)(nil)).Elem()
}

func (o EventIntegrationOutput) ToEventIntegrationOutput() EventIntegrationOutput {
	return o
}

func (o EventIntegrationOutput) ToEventIntegrationOutputWithContext(ctx context.Context) EventIntegrationOutput {
	return o
}

func (o EventIntegrationOutput) ToOutput(ctx context.Context) pulumix.Output[*EventIntegration] {
	return pulumix.Output[*EventIntegration]{
		OutputState: o.OutputState,
	}
}

// ARN of the Event Integration.
func (o EventIntegrationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventIntegration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the Event Integration.
func (o EventIntegrationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventIntegration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
func (o EventIntegrationOutput) EventFilter() EventIntegrationEventFilterOutput {
	return o.ApplyT(func(v *EventIntegration) EventIntegrationEventFilterOutput { return v.EventFilter }).(EventIntegrationEventFilterOutput)
}

// EventBridge bus.
func (o EventIntegrationOutput) EventbridgeBus() pulumi.StringOutput {
	return o.ApplyT(func(v *EventIntegration) pulumi.StringOutput { return v.EventbridgeBus }).(pulumi.StringOutput)
}

// Name of the Event Integration.
func (o EventIntegrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventIntegration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Tags to apply to the Event Integration. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EventIntegrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventIntegration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o EventIntegrationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventIntegration) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type EventIntegrationArrayOutput struct{ *pulumi.OutputState }

func (EventIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventIntegration)(nil)).Elem()
}

func (o EventIntegrationArrayOutput) ToEventIntegrationArrayOutput() EventIntegrationArrayOutput {
	return o
}

func (o EventIntegrationArrayOutput) ToEventIntegrationArrayOutputWithContext(ctx context.Context) EventIntegrationArrayOutput {
	return o
}

func (o EventIntegrationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EventIntegration] {
	return pulumix.Output[[]*EventIntegration]{
		OutputState: o.OutputState,
	}
}

func (o EventIntegrationArrayOutput) Index(i pulumi.IntInput) EventIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventIntegration {
		return vs[0].([]*EventIntegration)[vs[1].(int)]
	}).(EventIntegrationOutput)
}

type EventIntegrationMapOutput struct{ *pulumi.OutputState }

func (EventIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventIntegration)(nil)).Elem()
}

func (o EventIntegrationMapOutput) ToEventIntegrationMapOutput() EventIntegrationMapOutput {
	return o
}

func (o EventIntegrationMapOutput) ToEventIntegrationMapOutputWithContext(ctx context.Context) EventIntegrationMapOutput {
	return o
}

func (o EventIntegrationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EventIntegration] {
	return pulumix.Output[map[string]*EventIntegration]{
		OutputState: o.OutputState,
	}
}

func (o EventIntegrationMapOutput) MapIndex(k pulumi.StringInput) EventIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventIntegration {
		return vs[0].(map[string]*EventIntegration)[vs[1].(string)]
	}).(EventIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventIntegrationInput)(nil)).Elem(), &EventIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventIntegrationArrayInput)(nil)).Elem(), EventIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventIntegrationMapInput)(nil)).Elem(), EventIntegrationMap{})
	pulumi.RegisterOutputType(EventIntegrationOutput{})
	pulumi.RegisterOutputType(EventIntegrationArrayOutput{})
	pulumi.RegisterOutputType(EventIntegrationMapOutput{})
}
