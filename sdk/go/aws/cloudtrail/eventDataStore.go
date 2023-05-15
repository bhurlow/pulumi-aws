// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudTrail Event Data Store.
//
// More information about event data stores can be found in the [Event Data Store User Guide](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html).
//
// > **Tip:** For an organization event data store you must create this resource in the management account.
//
// ## Example Usage
// ### Basic
//
// The most simple event data store configuration requires us to only set the `name` attribute. The event data store will automatically capture all management events. To capture management events from all the regions, `multiRegionEnabled` must be `true`.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudtrail.NewEventDataStore(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Data Event Logging
//
// CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:
//
// - [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html)
// ### Log all DynamoDB PutEvent actions for a specific DynamoDB table
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			table, err := dynamodb.LookupTable(ctx, &dynamodb.LookupTableArgs{
//				Name: "not-important-dynamodb-table",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudtrail.NewEventDataStore(ctx, "example", &cloudtrail.EventDataStoreArgs{
//				AdvancedEventSelectors: cloudtrail.EventDataStoreAdvancedEventSelectorArray{
//					&cloudtrail.EventDataStoreAdvancedEventSelectorArgs{
//						Name: pulumi.String("Log all DynamoDB PutEvent actions for a specific DynamoDB table"),
//						FieldSelectors: cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArray{
//							&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
//								Field: pulumi.String("eventCategory"),
//								Equals: pulumi.StringArray{
//									pulumi.String("Data"),
//								},
//							},
//							&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
//								Field: pulumi.String("resources.type"),
//								Equals: pulumi.StringArray{
//									pulumi.String("AWS::DynamoDB::Table"),
//								},
//							},
//							&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
//								Field: pulumi.String("eventName"),
//								Equals: pulumi.StringArray{
//									pulumi.String("PutItem"),
//								},
//							},
//							&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
//								Field: pulumi.String("resources.ARN"),
//								Equals: pulumi.StringArray{
//									*pulumi.String(table.Arn),
//								},
//							},
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
// Event data stores can be imported using their `arn`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:cloudtrail/eventDataStore:EventDataStore example arn:aws:cloudtrail:us-east-1:123456789123:eventdatastore/22333815-4414-412c-b155-dd254033gfhf
//
// ```
type EventDataStore struct {
	pulumi.CustomResourceState

	// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	AdvancedEventSelectors EventDataStoreAdvancedEventSelectorArrayOutput `pulumi:"advancedEventSelectors"`
	// ARN of the event data store.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
	MultiRegionEnabled pulumi.BoolPtrOutput `pulumi:"multiRegionEnabled"`
	// The name of the event data store.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
	OrganizationEnabled pulumi.BoolPtrOutput `pulumi:"organizationEnabled"`
	// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
	RetentionPeriod pulumi.IntPtrOutput `pulumi:"retentionPeriod"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
	TerminationProtectionEnabled pulumi.BoolPtrOutput `pulumi:"terminationProtectionEnabled"`
}

// NewEventDataStore registers a new resource with the given unique name, arguments, and options.
func NewEventDataStore(ctx *pulumi.Context,
	name string, args *EventDataStoreArgs, opts ...pulumi.ResourceOption) (*EventDataStore, error) {
	if args == nil {
		args = &EventDataStoreArgs{}
	}

	var resource EventDataStore
	err := ctx.RegisterResource("aws:cloudtrail/eventDataStore:EventDataStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventDataStore gets an existing EventDataStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventDataStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventDataStoreState, opts ...pulumi.ResourceOption) (*EventDataStore, error) {
	var resource EventDataStore
	err := ctx.ReadResource("aws:cloudtrail/eventDataStore:EventDataStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventDataStore resources.
type eventDataStoreState struct {
	// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	AdvancedEventSelectors []EventDataStoreAdvancedEventSelector `pulumi:"advancedEventSelectors"`
	// ARN of the event data store.
	Arn *string `pulumi:"arn"`
	// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
	MultiRegionEnabled *bool `pulumi:"multiRegionEnabled"`
	// The name of the event data store.
	Name *string `pulumi:"name"`
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
	OrganizationEnabled *bool `pulumi:"organizationEnabled"`
	// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
	TerminationProtectionEnabled *bool `pulumi:"terminationProtectionEnabled"`
}

type EventDataStoreState struct {
	// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	AdvancedEventSelectors EventDataStoreAdvancedEventSelectorArrayInput
	// ARN of the event data store.
	Arn pulumi.StringPtrInput
	// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	KmsKeyId pulumi.StringPtrInput
	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
	MultiRegionEnabled pulumi.BoolPtrInput
	// The name of the event data store.
	Name pulumi.StringPtrInput
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
	OrganizationEnabled pulumi.BoolPtrInput
	// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
	RetentionPeriod pulumi.IntPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
	TerminationProtectionEnabled pulumi.BoolPtrInput
}

func (EventDataStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventDataStoreState)(nil)).Elem()
}

type eventDataStoreArgs struct {
	// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	AdvancedEventSelectors []EventDataStoreAdvancedEventSelector `pulumi:"advancedEventSelectors"`
	// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
	MultiRegionEnabled *bool `pulumi:"multiRegionEnabled"`
	// The name of the event data store.
	Name *string `pulumi:"name"`
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
	OrganizationEnabled *bool `pulumi:"organizationEnabled"`
	// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
	TerminationProtectionEnabled *bool `pulumi:"terminationProtectionEnabled"`
}

// The set of arguments for constructing a EventDataStore resource.
type EventDataStoreArgs struct {
	// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	AdvancedEventSelectors EventDataStoreAdvancedEventSelectorArrayInput
	// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	KmsKeyId pulumi.StringPtrInput
	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
	MultiRegionEnabled pulumi.BoolPtrInput
	// The name of the event data store.
	Name pulumi.StringPtrInput
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
	OrganizationEnabled pulumi.BoolPtrInput
	// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
	RetentionPeriod pulumi.IntPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
	TerminationProtectionEnabled pulumi.BoolPtrInput
}

func (EventDataStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventDataStoreArgs)(nil)).Elem()
}

type EventDataStoreInput interface {
	pulumi.Input

	ToEventDataStoreOutput() EventDataStoreOutput
	ToEventDataStoreOutputWithContext(ctx context.Context) EventDataStoreOutput
}

func (*EventDataStore) ElementType() reflect.Type {
	return reflect.TypeOf((**EventDataStore)(nil)).Elem()
}

func (i *EventDataStore) ToEventDataStoreOutput() EventDataStoreOutput {
	return i.ToEventDataStoreOutputWithContext(context.Background())
}

func (i *EventDataStore) ToEventDataStoreOutputWithContext(ctx context.Context) EventDataStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventDataStoreOutput)
}

// EventDataStoreArrayInput is an input type that accepts EventDataStoreArray and EventDataStoreArrayOutput values.
// You can construct a concrete instance of `EventDataStoreArrayInput` via:
//
//	EventDataStoreArray{ EventDataStoreArgs{...} }
type EventDataStoreArrayInput interface {
	pulumi.Input

	ToEventDataStoreArrayOutput() EventDataStoreArrayOutput
	ToEventDataStoreArrayOutputWithContext(context.Context) EventDataStoreArrayOutput
}

type EventDataStoreArray []EventDataStoreInput

func (EventDataStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventDataStore)(nil)).Elem()
}

func (i EventDataStoreArray) ToEventDataStoreArrayOutput() EventDataStoreArrayOutput {
	return i.ToEventDataStoreArrayOutputWithContext(context.Background())
}

func (i EventDataStoreArray) ToEventDataStoreArrayOutputWithContext(ctx context.Context) EventDataStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventDataStoreArrayOutput)
}

// EventDataStoreMapInput is an input type that accepts EventDataStoreMap and EventDataStoreMapOutput values.
// You can construct a concrete instance of `EventDataStoreMapInput` via:
//
//	EventDataStoreMap{ "key": EventDataStoreArgs{...} }
type EventDataStoreMapInput interface {
	pulumi.Input

	ToEventDataStoreMapOutput() EventDataStoreMapOutput
	ToEventDataStoreMapOutputWithContext(context.Context) EventDataStoreMapOutput
}

type EventDataStoreMap map[string]EventDataStoreInput

func (EventDataStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventDataStore)(nil)).Elem()
}

func (i EventDataStoreMap) ToEventDataStoreMapOutput() EventDataStoreMapOutput {
	return i.ToEventDataStoreMapOutputWithContext(context.Background())
}

func (i EventDataStoreMap) ToEventDataStoreMapOutputWithContext(ctx context.Context) EventDataStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventDataStoreMapOutput)
}

type EventDataStoreOutput struct{ *pulumi.OutputState }

func (EventDataStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventDataStore)(nil)).Elem()
}

func (o EventDataStoreOutput) ToEventDataStoreOutput() EventDataStoreOutput {
	return o
}

func (o EventDataStoreOutput) ToEventDataStoreOutputWithContext(ctx context.Context) EventDataStoreOutput {
	return o
}

// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
func (o EventDataStoreOutput) AdvancedEventSelectors() EventDataStoreAdvancedEventSelectorArrayOutput {
	return o.ApplyT(func(v *EventDataStore) EventDataStoreAdvancedEventSelectorArrayOutput {
		return v.AdvancedEventSelectors
	}).(EventDataStoreAdvancedEventSelectorArrayOutput)
}

// ARN of the event data store.
func (o EventDataStoreOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
func (o EventDataStoreOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
func (o EventDataStoreOutput) MultiRegionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.BoolPtrOutput { return v.MultiRegionEnabled }).(pulumi.BoolPtrOutput)
}

// The name of the event data store.
func (o EventDataStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
func (o EventDataStoreOutput) OrganizationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.BoolPtrOutput { return v.OrganizationEnabled }).(pulumi.BoolPtrOutput)
}

// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
func (o EventDataStoreOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.IntPtrOutput { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EventDataStoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o EventDataStoreOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
func (o EventDataStoreOutput) TerminationProtectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventDataStore) pulumi.BoolPtrOutput { return v.TerminationProtectionEnabled }).(pulumi.BoolPtrOutput)
}

type EventDataStoreArrayOutput struct{ *pulumi.OutputState }

func (EventDataStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventDataStore)(nil)).Elem()
}

func (o EventDataStoreArrayOutput) ToEventDataStoreArrayOutput() EventDataStoreArrayOutput {
	return o
}

func (o EventDataStoreArrayOutput) ToEventDataStoreArrayOutputWithContext(ctx context.Context) EventDataStoreArrayOutput {
	return o
}

func (o EventDataStoreArrayOutput) Index(i pulumi.IntInput) EventDataStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventDataStore {
		return vs[0].([]*EventDataStore)[vs[1].(int)]
	}).(EventDataStoreOutput)
}

type EventDataStoreMapOutput struct{ *pulumi.OutputState }

func (EventDataStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventDataStore)(nil)).Elem()
}

func (o EventDataStoreMapOutput) ToEventDataStoreMapOutput() EventDataStoreMapOutput {
	return o
}

func (o EventDataStoreMapOutput) ToEventDataStoreMapOutputWithContext(ctx context.Context) EventDataStoreMapOutput {
	return o
}

func (o EventDataStoreMapOutput) MapIndex(k pulumi.StringInput) EventDataStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventDataStore {
		return vs[0].(map[string]*EventDataStore)[vs[1].(string)]
	}).(EventDataStoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventDataStoreInput)(nil)).Elem(), &EventDataStore{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventDataStoreArrayInput)(nil)).Elem(), EventDataStoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventDataStoreMapInput)(nil)).Elem(), EventDataStoreMap{})
	pulumi.RegisterOutputType(EventDataStoreOutput{})
	pulumi.RegisterOutputType(EventDataStoreArrayOutput{})
	pulumi.RegisterOutputType(EventDataStoreMapOutput{})
}
