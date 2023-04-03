// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Route53 Traffic Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.NewTrafficPolicy(ctx, "example", &route53.TrafficPolicyArgs{
//				Comment:  pulumi.String("example comment"),
//				Document: pulumi.String("{\n  \"AWSPolicyFormatVersion\": \"2015-10-01\",\n  \"RecordType\": \"A\",\n  \"Endpoints\": {\n    \"endpoint-start-NkPh\": {\n      \"Type\": \"value\",\n      \"Value\": \"10.0.0.2\"\n    }\n  },\n  \"StartEndpoint\": \"endpoint-start-NkPh\"\n}\n\n"),
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
// Route53 Traffic Policy can be imported using the `id` and `version`, e.g.
//
// ```sh
//
//	$ pulumi import aws:route53/trafficPolicy:TrafficPolicy example 01a52019-d16f-422a-ae72-c306d2b6df7e/1
//
// ```
type TrafficPolicy struct {
	pulumi.CustomResourceState

	// Comment for the traffic policy.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the [AWS Route53 Traffic Policy document format](https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)
	Document pulumi.StringOutput `pulumi:"document"`
	// Name of the traffic policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.
	Type pulumi.StringOutput `pulumi:"type"`
	// Version number of the traffic policy. This value is automatically incremented by AWS after each update of this resource.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewTrafficPolicy registers a new resource with the given unique name, arguments, and options.
func NewTrafficPolicy(ctx *pulumi.Context,
	name string, args *TrafficPolicyArgs, opts ...pulumi.ResourceOption) (*TrafficPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Document == nil {
		return nil, errors.New("invalid value for required argument 'Document'")
	}
	var resource TrafficPolicy
	err := ctx.RegisterResource("aws:route53/trafficPolicy:TrafficPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficPolicy gets an existing TrafficPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficPolicyState, opts ...pulumi.ResourceOption) (*TrafficPolicy, error) {
	var resource TrafficPolicy
	err := ctx.ReadResource("aws:route53/trafficPolicy:TrafficPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficPolicy resources.
type trafficPolicyState struct {
	// Comment for the traffic policy.
	Comment *string `pulumi:"comment"`
	// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the [AWS Route53 Traffic Policy document format](https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)
	Document *string `pulumi:"document"`
	// Name of the traffic policy.
	Name *string `pulumi:"name"`
	// DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.
	Type *string `pulumi:"type"`
	// Version number of the traffic policy. This value is automatically incremented by AWS after each update of this resource.
	Version *int `pulumi:"version"`
}

type TrafficPolicyState struct {
	// Comment for the traffic policy.
	Comment pulumi.StringPtrInput
	// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the [AWS Route53 Traffic Policy document format](https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)
	Document pulumi.StringPtrInput
	// Name of the traffic policy.
	Name pulumi.StringPtrInput
	// DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.
	Type pulumi.StringPtrInput
	// Version number of the traffic policy. This value is automatically incremented by AWS after each update of this resource.
	Version pulumi.IntPtrInput
}

func (TrafficPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficPolicyState)(nil)).Elem()
}

type trafficPolicyArgs struct {
	// Comment for the traffic policy.
	Comment *string `pulumi:"comment"`
	// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the [AWS Route53 Traffic Policy document format](https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)
	Document string `pulumi:"document"`
	// Name of the traffic policy.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a TrafficPolicy resource.
type TrafficPolicyArgs struct {
	// Comment for the traffic policy.
	Comment pulumi.StringPtrInput
	// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the [AWS Route53 Traffic Policy document format](https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)
	Document pulumi.StringInput
	// Name of the traffic policy.
	Name pulumi.StringPtrInput
}

func (TrafficPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficPolicyArgs)(nil)).Elem()
}

type TrafficPolicyInput interface {
	pulumi.Input

	ToTrafficPolicyOutput() TrafficPolicyOutput
	ToTrafficPolicyOutputWithContext(ctx context.Context) TrafficPolicyOutput
}

func (*TrafficPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficPolicy)(nil)).Elem()
}

func (i *TrafficPolicy) ToTrafficPolicyOutput() TrafficPolicyOutput {
	return i.ToTrafficPolicyOutputWithContext(context.Background())
}

func (i *TrafficPolicy) ToTrafficPolicyOutputWithContext(ctx context.Context) TrafficPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficPolicyOutput)
}

// TrafficPolicyArrayInput is an input type that accepts TrafficPolicyArray and TrafficPolicyArrayOutput values.
// You can construct a concrete instance of `TrafficPolicyArrayInput` via:
//
//	TrafficPolicyArray{ TrafficPolicyArgs{...} }
type TrafficPolicyArrayInput interface {
	pulumi.Input

	ToTrafficPolicyArrayOutput() TrafficPolicyArrayOutput
	ToTrafficPolicyArrayOutputWithContext(context.Context) TrafficPolicyArrayOutput
}

type TrafficPolicyArray []TrafficPolicyInput

func (TrafficPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficPolicy)(nil)).Elem()
}

func (i TrafficPolicyArray) ToTrafficPolicyArrayOutput() TrafficPolicyArrayOutput {
	return i.ToTrafficPolicyArrayOutputWithContext(context.Background())
}

func (i TrafficPolicyArray) ToTrafficPolicyArrayOutputWithContext(ctx context.Context) TrafficPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficPolicyArrayOutput)
}

// TrafficPolicyMapInput is an input type that accepts TrafficPolicyMap and TrafficPolicyMapOutput values.
// You can construct a concrete instance of `TrafficPolicyMapInput` via:
//
//	TrafficPolicyMap{ "key": TrafficPolicyArgs{...} }
type TrafficPolicyMapInput interface {
	pulumi.Input

	ToTrafficPolicyMapOutput() TrafficPolicyMapOutput
	ToTrafficPolicyMapOutputWithContext(context.Context) TrafficPolicyMapOutput
}

type TrafficPolicyMap map[string]TrafficPolicyInput

func (TrafficPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficPolicy)(nil)).Elem()
}

func (i TrafficPolicyMap) ToTrafficPolicyMapOutput() TrafficPolicyMapOutput {
	return i.ToTrafficPolicyMapOutputWithContext(context.Background())
}

func (i TrafficPolicyMap) ToTrafficPolicyMapOutputWithContext(ctx context.Context) TrafficPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficPolicyMapOutput)
}

type TrafficPolicyOutput struct{ *pulumi.OutputState }

func (TrafficPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficPolicy)(nil)).Elem()
}

func (o TrafficPolicyOutput) ToTrafficPolicyOutput() TrafficPolicyOutput {
	return o
}

func (o TrafficPolicyOutput) ToTrafficPolicyOutputWithContext(ctx context.Context) TrafficPolicyOutput {
	return o
}

// Comment for the traffic policy.
func (o TrafficPolicyOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficPolicy) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the [AWS Route53 Traffic Policy document format](https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)
func (o TrafficPolicyOutput) Document() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficPolicy) pulumi.StringOutput { return v.Document }).(pulumi.StringOutput)
}

// Name of the traffic policy.
func (o TrafficPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.
func (o TrafficPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version number of the traffic policy. This value is automatically incremented by AWS after each update of this resource.
func (o TrafficPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type TrafficPolicyArrayOutput struct{ *pulumi.OutputState }

func (TrafficPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficPolicy)(nil)).Elem()
}

func (o TrafficPolicyArrayOutput) ToTrafficPolicyArrayOutput() TrafficPolicyArrayOutput {
	return o
}

func (o TrafficPolicyArrayOutput) ToTrafficPolicyArrayOutputWithContext(ctx context.Context) TrafficPolicyArrayOutput {
	return o
}

func (o TrafficPolicyArrayOutput) Index(i pulumi.IntInput) TrafficPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficPolicy {
		return vs[0].([]*TrafficPolicy)[vs[1].(int)]
	}).(TrafficPolicyOutput)
}

type TrafficPolicyMapOutput struct{ *pulumi.OutputState }

func (TrafficPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficPolicy)(nil)).Elem()
}

func (o TrafficPolicyMapOutput) ToTrafficPolicyMapOutput() TrafficPolicyMapOutput {
	return o
}

func (o TrafficPolicyMapOutput) ToTrafficPolicyMapOutputWithContext(ctx context.Context) TrafficPolicyMapOutput {
	return o
}

func (o TrafficPolicyMapOutput) MapIndex(k pulumi.StringInput) TrafficPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficPolicy {
		return vs[0].(map[string]*TrafficPolicy)[vs[1].(string)]
	}).(TrafficPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficPolicyInput)(nil)).Elem(), &TrafficPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficPolicyArrayInput)(nil)).Elem(), TrafficPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficPolicyMapInput)(nil)).Elem(), TrafficPolicyMap{})
	pulumi.RegisterOutputType(TrafficPolicyOutput{})
	pulumi.RegisterOutputType(TrafficPolicyArrayOutput{})
	pulumi.RegisterOutputType(TrafficPolicyMapOutput{})
}
