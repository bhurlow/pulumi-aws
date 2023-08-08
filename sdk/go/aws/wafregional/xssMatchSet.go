// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Regional XSS Match Set Resource for use with Application Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafregional"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wafregional.NewXssMatchSet(ctx, "xssMatchSet", &wafregional.XssMatchSetArgs{
//				XssMatchTuples: wafregional.XssMatchSetXssMatchTupleArray{
//					&wafregional.XssMatchSetXssMatchTupleArgs{
//						FieldToMatch: &wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs{
//							Type: pulumi.String("URI"),
//						},
//						TextTransformation: pulumi.String("NONE"),
//					},
//					&wafregional.XssMatchSetXssMatchTupleArgs{
//						FieldToMatch: &wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs{
//							Type: pulumi.String("QUERY_STRING"),
//						},
//						TextTransformation: pulumi.String("NONE"),
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
//	to = aws_wafregional_xss_match_set.example
//
//	id = "12345abcde" } Using `pulumi import`, import AWS WAF Regional XSS Match using the `id`. For exampleconsole % pulumi import aws_wafregional_xss_match_set.example 12345abcde
type XssMatchSet struct {
	pulumi.CustomResourceState

	// The name of the set
	Name pulumi.StringOutput `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayOutput `pulumi:"xssMatchTuples"`
}

// NewXssMatchSet registers a new resource with the given unique name, arguments, and options.
func NewXssMatchSet(ctx *pulumi.Context,
	name string, args *XssMatchSetArgs, opts ...pulumi.ResourceOption) (*XssMatchSet, error) {
	if args == nil {
		args = &XssMatchSetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource XssMatchSet
	err := ctx.RegisterResource("aws:wafregional/xssMatchSet:XssMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetXssMatchSet gets an existing XssMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetXssMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *XssMatchSetState, opts ...pulumi.ResourceOption) (*XssMatchSet, error) {
	var resource XssMatchSet
	err := ctx.ReadResource("aws:wafregional/xssMatchSet:XssMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering XssMatchSet resources.
type xssMatchSetState struct {
	// The name of the set
	Name *string `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []XssMatchSetXssMatchTuple `pulumi:"xssMatchTuples"`
}

type XssMatchSetState struct {
	// The name of the set
	Name pulumi.StringPtrInput
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayInput
}

func (XssMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*xssMatchSetState)(nil)).Elem()
}

type xssMatchSetArgs struct {
	// The name of the set
	Name *string `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []XssMatchSetXssMatchTuple `pulumi:"xssMatchTuples"`
}

// The set of arguments for constructing a XssMatchSet resource.
type XssMatchSetArgs struct {
	// The name of the set
	Name pulumi.StringPtrInput
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayInput
}

func (XssMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*xssMatchSetArgs)(nil)).Elem()
}

type XssMatchSetInput interface {
	pulumi.Input

	ToXssMatchSetOutput() XssMatchSetOutput
	ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput
}

func (*XssMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((**XssMatchSet)(nil)).Elem()
}

func (i *XssMatchSet) ToXssMatchSetOutput() XssMatchSetOutput {
	return i.ToXssMatchSetOutputWithContext(context.Background())
}

func (i *XssMatchSet) ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetOutput)
}

// XssMatchSetArrayInput is an input type that accepts XssMatchSetArray and XssMatchSetArrayOutput values.
// You can construct a concrete instance of `XssMatchSetArrayInput` via:
//
//	XssMatchSetArray{ XssMatchSetArgs{...} }
type XssMatchSetArrayInput interface {
	pulumi.Input

	ToXssMatchSetArrayOutput() XssMatchSetArrayOutput
	ToXssMatchSetArrayOutputWithContext(context.Context) XssMatchSetArrayOutput
}

type XssMatchSetArray []XssMatchSetInput

func (XssMatchSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*XssMatchSet)(nil)).Elem()
}

func (i XssMatchSetArray) ToXssMatchSetArrayOutput() XssMatchSetArrayOutput {
	return i.ToXssMatchSetArrayOutputWithContext(context.Background())
}

func (i XssMatchSetArray) ToXssMatchSetArrayOutputWithContext(ctx context.Context) XssMatchSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetArrayOutput)
}

// XssMatchSetMapInput is an input type that accepts XssMatchSetMap and XssMatchSetMapOutput values.
// You can construct a concrete instance of `XssMatchSetMapInput` via:
//
//	XssMatchSetMap{ "key": XssMatchSetArgs{...} }
type XssMatchSetMapInput interface {
	pulumi.Input

	ToXssMatchSetMapOutput() XssMatchSetMapOutput
	ToXssMatchSetMapOutputWithContext(context.Context) XssMatchSetMapOutput
}

type XssMatchSetMap map[string]XssMatchSetInput

func (XssMatchSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*XssMatchSet)(nil)).Elem()
}

func (i XssMatchSetMap) ToXssMatchSetMapOutput() XssMatchSetMapOutput {
	return i.ToXssMatchSetMapOutputWithContext(context.Background())
}

func (i XssMatchSetMap) ToXssMatchSetMapOutputWithContext(ctx context.Context) XssMatchSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetMapOutput)
}

type XssMatchSetOutput struct{ *pulumi.OutputState }

func (XssMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**XssMatchSet)(nil)).Elem()
}

func (o XssMatchSetOutput) ToXssMatchSetOutput() XssMatchSetOutput {
	return o
}

func (o XssMatchSetOutput) ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput {
	return o
}

// The name of the set
func (o XssMatchSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *XssMatchSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parts of web requests that you want to inspect for cross-site scripting attacks.
func (o XssMatchSetOutput) XssMatchTuples() XssMatchSetXssMatchTupleArrayOutput {
	return o.ApplyT(func(v *XssMatchSet) XssMatchSetXssMatchTupleArrayOutput { return v.XssMatchTuples }).(XssMatchSetXssMatchTupleArrayOutput)
}

type XssMatchSetArrayOutput struct{ *pulumi.OutputState }

func (XssMatchSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*XssMatchSet)(nil)).Elem()
}

func (o XssMatchSetArrayOutput) ToXssMatchSetArrayOutput() XssMatchSetArrayOutput {
	return o
}

func (o XssMatchSetArrayOutput) ToXssMatchSetArrayOutputWithContext(ctx context.Context) XssMatchSetArrayOutput {
	return o
}

func (o XssMatchSetArrayOutput) Index(i pulumi.IntInput) XssMatchSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *XssMatchSet {
		return vs[0].([]*XssMatchSet)[vs[1].(int)]
	}).(XssMatchSetOutput)
}

type XssMatchSetMapOutput struct{ *pulumi.OutputState }

func (XssMatchSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*XssMatchSet)(nil)).Elem()
}

func (o XssMatchSetMapOutput) ToXssMatchSetMapOutput() XssMatchSetMapOutput {
	return o
}

func (o XssMatchSetMapOutput) ToXssMatchSetMapOutputWithContext(ctx context.Context) XssMatchSetMapOutput {
	return o
}

func (o XssMatchSetMapOutput) MapIndex(k pulumi.StringInput) XssMatchSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *XssMatchSet {
		return vs[0].(map[string]*XssMatchSet)[vs[1].(string)]
	}).(XssMatchSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*XssMatchSetInput)(nil)).Elem(), &XssMatchSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*XssMatchSetArrayInput)(nil)).Elem(), XssMatchSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*XssMatchSetMapInput)(nil)).Elem(), XssMatchSetMap{})
	pulumi.RegisterOutputType(XssMatchSetOutput{})
	pulumi.RegisterOutputType(XssMatchSetArrayOutput{})
	pulumi.RegisterOutputType(XssMatchSetMapOutput{})
}
