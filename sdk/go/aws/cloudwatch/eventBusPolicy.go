// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create an EventBridge resource policy to support cross-account events.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// > **Note:** The EventBridge bus policy resource  (`cloudwatch.EventBusPolicy`) is incompatible with the EventBridge permission resource (`cloudwatch.EventPermission`) and will overwrite permissions.
//
// ## Example Usage
// ### Account Access
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Sid:    pulumi.StringRef("DevAccountAccess"),
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"events:PutEvents",
//						},
//						Resources: []string{
//							"arn:aws:events:eu-west-1:123456789012:event-bus/default",
//						},
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "AWS",
//								Identifiers: []string{
//									"123456789012",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudwatch.NewEventBusPolicy(ctx, "testEventBusPolicy", &cloudwatch.EventBusPolicyArgs{
//				Policy:       *pulumi.String(testPolicyDocument.Json),
//				EventBusName: pulumi.Any(aws_cloudwatch_event_bus.Test.Name),
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
// EventBridge permissions can be imported using the `event_bus_name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:cloudwatch/eventBusPolicy:EventBusPolicy DevAccountAccess example-event-bus
//
// ```
type EventBusPolicy struct {
	pulumi.CustomResourceState

	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the `default` event bus.
	EventBusName pulumi.StringPtrOutput `pulumi:"eventBusName"`
	// The text of the policy.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewEventBusPolicy registers a new resource with the given unique name, arguments, and options.
func NewEventBusPolicy(ctx *pulumi.Context,
	name string, args *EventBusPolicyArgs, opts ...pulumi.ResourceOption) (*EventBusPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource EventBusPolicy
	err := ctx.RegisterResource("aws:cloudwatch/eventBusPolicy:EventBusPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventBusPolicy gets an existing EventBusPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventBusPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventBusPolicyState, opts ...pulumi.ResourceOption) (*EventBusPolicy, error) {
	var resource EventBusPolicy
	err := ctx.ReadResource("aws:cloudwatch/eventBusPolicy:EventBusPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventBusPolicy resources.
type eventBusPolicyState struct {
	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the `default` event bus.
	EventBusName *string `pulumi:"eventBusName"`
	// The text of the policy.
	Policy *string `pulumi:"policy"`
}

type EventBusPolicyState struct {
	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the `default` event bus.
	EventBusName pulumi.StringPtrInput
	// The text of the policy.
	Policy pulumi.StringPtrInput
}

func (EventBusPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventBusPolicyState)(nil)).Elem()
}

type eventBusPolicyArgs struct {
	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the `default` event bus.
	EventBusName *string `pulumi:"eventBusName"`
	// The text of the policy.
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a EventBusPolicy resource.
type EventBusPolicyArgs struct {
	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the `default` event bus.
	EventBusName pulumi.StringPtrInput
	// The text of the policy.
	Policy pulumi.StringInput
}

func (EventBusPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventBusPolicyArgs)(nil)).Elem()
}

type EventBusPolicyInput interface {
	pulumi.Input

	ToEventBusPolicyOutput() EventBusPolicyOutput
	ToEventBusPolicyOutputWithContext(ctx context.Context) EventBusPolicyOutput
}

func (*EventBusPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EventBusPolicy)(nil)).Elem()
}

func (i *EventBusPolicy) ToEventBusPolicyOutput() EventBusPolicyOutput {
	return i.ToEventBusPolicyOutputWithContext(context.Background())
}

func (i *EventBusPolicy) ToEventBusPolicyOutputWithContext(ctx context.Context) EventBusPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventBusPolicyOutput)
}

// EventBusPolicyArrayInput is an input type that accepts EventBusPolicyArray and EventBusPolicyArrayOutput values.
// You can construct a concrete instance of `EventBusPolicyArrayInput` via:
//
//	EventBusPolicyArray{ EventBusPolicyArgs{...} }
type EventBusPolicyArrayInput interface {
	pulumi.Input

	ToEventBusPolicyArrayOutput() EventBusPolicyArrayOutput
	ToEventBusPolicyArrayOutputWithContext(context.Context) EventBusPolicyArrayOutput
}

type EventBusPolicyArray []EventBusPolicyInput

func (EventBusPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventBusPolicy)(nil)).Elem()
}

func (i EventBusPolicyArray) ToEventBusPolicyArrayOutput() EventBusPolicyArrayOutput {
	return i.ToEventBusPolicyArrayOutputWithContext(context.Background())
}

func (i EventBusPolicyArray) ToEventBusPolicyArrayOutputWithContext(ctx context.Context) EventBusPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventBusPolicyArrayOutput)
}

// EventBusPolicyMapInput is an input type that accepts EventBusPolicyMap and EventBusPolicyMapOutput values.
// You can construct a concrete instance of `EventBusPolicyMapInput` via:
//
//	EventBusPolicyMap{ "key": EventBusPolicyArgs{...} }
type EventBusPolicyMapInput interface {
	pulumi.Input

	ToEventBusPolicyMapOutput() EventBusPolicyMapOutput
	ToEventBusPolicyMapOutputWithContext(context.Context) EventBusPolicyMapOutput
}

type EventBusPolicyMap map[string]EventBusPolicyInput

func (EventBusPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventBusPolicy)(nil)).Elem()
}

func (i EventBusPolicyMap) ToEventBusPolicyMapOutput() EventBusPolicyMapOutput {
	return i.ToEventBusPolicyMapOutputWithContext(context.Background())
}

func (i EventBusPolicyMap) ToEventBusPolicyMapOutputWithContext(ctx context.Context) EventBusPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventBusPolicyMapOutput)
}

type EventBusPolicyOutput struct{ *pulumi.OutputState }

func (EventBusPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventBusPolicy)(nil)).Elem()
}

func (o EventBusPolicyOutput) ToEventBusPolicyOutput() EventBusPolicyOutput {
	return o
}

func (o EventBusPolicyOutput) ToEventBusPolicyOutputWithContext(ctx context.Context) EventBusPolicyOutput {
	return o
}

// The name of the event bus to set the permissions on.
// If you omit this, the permissions are set on the `default` event bus.
func (o EventBusPolicyOutput) EventBusName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventBusPolicy) pulumi.StringPtrOutput { return v.EventBusName }).(pulumi.StringPtrOutput)
}

// The text of the policy.
func (o EventBusPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *EventBusPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type EventBusPolicyArrayOutput struct{ *pulumi.OutputState }

func (EventBusPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventBusPolicy)(nil)).Elem()
}

func (o EventBusPolicyArrayOutput) ToEventBusPolicyArrayOutput() EventBusPolicyArrayOutput {
	return o
}

func (o EventBusPolicyArrayOutput) ToEventBusPolicyArrayOutputWithContext(ctx context.Context) EventBusPolicyArrayOutput {
	return o
}

func (o EventBusPolicyArrayOutput) Index(i pulumi.IntInput) EventBusPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventBusPolicy {
		return vs[0].([]*EventBusPolicy)[vs[1].(int)]
	}).(EventBusPolicyOutput)
}

type EventBusPolicyMapOutput struct{ *pulumi.OutputState }

func (EventBusPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventBusPolicy)(nil)).Elem()
}

func (o EventBusPolicyMapOutput) ToEventBusPolicyMapOutput() EventBusPolicyMapOutput {
	return o
}

func (o EventBusPolicyMapOutput) ToEventBusPolicyMapOutputWithContext(ctx context.Context) EventBusPolicyMapOutput {
	return o
}

func (o EventBusPolicyMapOutput) MapIndex(k pulumi.StringInput) EventBusPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventBusPolicy {
		return vs[0].(map[string]*EventBusPolicy)[vs[1].(string)]
	}).(EventBusPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventBusPolicyInput)(nil)).Elem(), &EventBusPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventBusPolicyArrayInput)(nil)).Elem(), EventBusPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventBusPolicyMapInput)(nil)).Elem(), EventBusPolicyMap{})
	pulumi.RegisterOutputType(EventBusPolicyOutput{})
	pulumi.RegisterOutputType(EventBusPolicyArrayOutput{})
	pulumi.RegisterOutputType(EventBusPolicyMapOutput{})
}
