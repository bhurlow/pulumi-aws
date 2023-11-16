// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Regional Regex Pattern Set Resource
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
//			_, err := wafregional.NewRegexPatternSet(ctx, "example", &wafregional.RegexPatternSetArgs{
//				RegexPatternStrings: pulumi.StringArray{
//					pulumi.String("one"),
//					pulumi.String("two"),
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
// Using `pulumi import`, import WAF Regional Regex Pattern Set using the id. For example:
//
// ```sh
//
//	$ pulumi import aws:wafregional/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
//
// ```
type RegexPatternSet struct {
	pulumi.CustomResourceState

	// The name or description of the Regex Pattern Set.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings pulumi.StringArrayOutput `pulumi:"regexPatternStrings"`
}

// NewRegexPatternSet registers a new resource with the given unique name, arguments, and options.
func NewRegexPatternSet(ctx *pulumi.Context,
	name string, args *RegexPatternSetArgs, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	if args == nil {
		args = &RegexPatternSetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegexPatternSet
	err := ctx.RegisterResource("aws:wafregional/regexPatternSet:RegexPatternSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexPatternSet gets an existing RegexPatternSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexPatternSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexPatternSetState, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	var resource RegexPatternSet
	err := ctx.ReadResource("aws:wafregional/regexPatternSet:RegexPatternSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexPatternSet resources.
type regexPatternSetState struct {
	// The name or description of the Regex Pattern Set.
	Name *string `pulumi:"name"`
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings []string `pulumi:"regexPatternStrings"`
}

type RegexPatternSetState struct {
	// The name or description of the Regex Pattern Set.
	Name pulumi.StringPtrInput
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings pulumi.StringArrayInput
}

func (RegexPatternSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetState)(nil)).Elem()
}

type regexPatternSetArgs struct {
	// The name or description of the Regex Pattern Set.
	Name *string `pulumi:"name"`
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings []string `pulumi:"regexPatternStrings"`
}

// The set of arguments for constructing a RegexPatternSet resource.
type RegexPatternSetArgs struct {
	// The name or description of the Regex Pattern Set.
	Name pulumi.StringPtrInput
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings pulumi.StringArrayInput
}

func (RegexPatternSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetArgs)(nil)).Elem()
}

type RegexPatternSetInput interface {
	pulumi.Input

	ToRegexPatternSetOutput() RegexPatternSetOutput
	ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput
}

func (*RegexPatternSet) ElementType() reflect.Type {
	return reflect.TypeOf((**RegexPatternSet)(nil)).Elem()
}

func (i *RegexPatternSet) ToRegexPatternSetOutput() RegexPatternSetOutput {
	return i.ToRegexPatternSetOutputWithContext(context.Background())
}

func (i *RegexPatternSet) ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexPatternSetOutput)
}

// RegexPatternSetArrayInput is an input type that accepts RegexPatternSetArray and RegexPatternSetArrayOutput values.
// You can construct a concrete instance of `RegexPatternSetArrayInput` via:
//
//	RegexPatternSetArray{ RegexPatternSetArgs{...} }
type RegexPatternSetArrayInput interface {
	pulumi.Input

	ToRegexPatternSetArrayOutput() RegexPatternSetArrayOutput
	ToRegexPatternSetArrayOutputWithContext(context.Context) RegexPatternSetArrayOutput
}

type RegexPatternSetArray []RegexPatternSetInput

func (RegexPatternSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegexPatternSet)(nil)).Elem()
}

func (i RegexPatternSetArray) ToRegexPatternSetArrayOutput() RegexPatternSetArrayOutput {
	return i.ToRegexPatternSetArrayOutputWithContext(context.Background())
}

func (i RegexPatternSetArray) ToRegexPatternSetArrayOutputWithContext(ctx context.Context) RegexPatternSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexPatternSetArrayOutput)
}

// RegexPatternSetMapInput is an input type that accepts RegexPatternSetMap and RegexPatternSetMapOutput values.
// You can construct a concrete instance of `RegexPatternSetMapInput` via:
//
//	RegexPatternSetMap{ "key": RegexPatternSetArgs{...} }
type RegexPatternSetMapInput interface {
	pulumi.Input

	ToRegexPatternSetMapOutput() RegexPatternSetMapOutput
	ToRegexPatternSetMapOutputWithContext(context.Context) RegexPatternSetMapOutput
}

type RegexPatternSetMap map[string]RegexPatternSetInput

func (RegexPatternSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegexPatternSet)(nil)).Elem()
}

func (i RegexPatternSetMap) ToRegexPatternSetMapOutput() RegexPatternSetMapOutput {
	return i.ToRegexPatternSetMapOutputWithContext(context.Background())
}

func (i RegexPatternSetMap) ToRegexPatternSetMapOutputWithContext(ctx context.Context) RegexPatternSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexPatternSetMapOutput)
}

type RegexPatternSetOutput struct{ *pulumi.OutputState }

func (RegexPatternSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegexPatternSet)(nil)).Elem()
}

func (o RegexPatternSetOutput) ToRegexPatternSetOutput() RegexPatternSetOutput {
	return o
}

func (o RegexPatternSetOutput) ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput {
	return o
}

// The name or description of the Regex Pattern Set.
func (o RegexPatternSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegexPatternSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
func (o RegexPatternSetOutput) RegexPatternStrings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RegexPatternSet) pulumi.StringArrayOutput { return v.RegexPatternStrings }).(pulumi.StringArrayOutput)
}

type RegexPatternSetArrayOutput struct{ *pulumi.OutputState }

func (RegexPatternSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegexPatternSet)(nil)).Elem()
}

func (o RegexPatternSetArrayOutput) ToRegexPatternSetArrayOutput() RegexPatternSetArrayOutput {
	return o
}

func (o RegexPatternSetArrayOutput) ToRegexPatternSetArrayOutputWithContext(ctx context.Context) RegexPatternSetArrayOutput {
	return o
}

func (o RegexPatternSetArrayOutput) Index(i pulumi.IntInput) RegexPatternSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegexPatternSet {
		return vs[0].([]*RegexPatternSet)[vs[1].(int)]
	}).(RegexPatternSetOutput)
}

type RegexPatternSetMapOutput struct{ *pulumi.OutputState }

func (RegexPatternSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegexPatternSet)(nil)).Elem()
}

func (o RegexPatternSetMapOutput) ToRegexPatternSetMapOutput() RegexPatternSetMapOutput {
	return o
}

func (o RegexPatternSetMapOutput) ToRegexPatternSetMapOutputWithContext(ctx context.Context) RegexPatternSetMapOutput {
	return o
}

func (o RegexPatternSetMapOutput) MapIndex(k pulumi.StringInput) RegexPatternSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegexPatternSet {
		return vs[0].(map[string]*RegexPatternSet)[vs[1].(string)]
	}).(RegexPatternSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegexPatternSetInput)(nil)).Elem(), &RegexPatternSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegexPatternSetArrayInput)(nil)).Elem(), RegexPatternSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegexPatternSetMapInput)(nil)).Elem(), RegexPatternSetMap{})
	pulumi.RegisterOutputType(RegexPatternSetOutput{})
	pulumi.RegisterOutputType(RegexPatternSetArrayOutput{})
	pulumi.RegisterOutputType(RegexPatternSetMapOutput{})
}
