// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Route 53 Resolver DNS Firewall domain list resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.NewResolverFirewallDomainList(ctx, "example", nil)
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
// # In TODO v1.5.0 and later, use an `import` block to import
//
// Route 53 Resolver DNS Firewall domain lists using the Route 53 Resolver DNS Firewall domain list ID. For exampleterraform import {
//
//	to = aws_route53_resolver_firewall_domain_list.example
//
//	id = "rslvr-fdl-0123456789abcdef" } Using `TODO import`, import
//
// Route 53 Resolver DNS Firewall domain lists using the Route 53 Resolver DNS Firewall domain list ID. For exampleconsole % TODO import aws_route53_resolver_firewall_domain_list.example rslvr-fdl-0123456789abcdef
type ResolverFirewallDomainList struct {
	pulumi.CustomResourceState

	// The ARN (Amazon Resource Name) of the domain list.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A array of domains for the firewall domain list.
	Domains pulumi.StringArrayOutput `pulumi:"domains"`
	// A name that lets you identify the domain list, to manage and use it.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource. f configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewResolverFirewallDomainList registers a new resource with the given unique name, arguments, and options.
func NewResolverFirewallDomainList(ctx *pulumi.Context,
	name string, args *ResolverFirewallDomainListArgs, opts ...pulumi.ResourceOption) (*ResolverFirewallDomainList, error) {
	if args == nil {
		args = &ResolverFirewallDomainListArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverFirewallDomainList
	err := ctx.RegisterResource("aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverFirewallDomainList gets an existing ResolverFirewallDomainList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverFirewallDomainList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverFirewallDomainListState, opts ...pulumi.ResourceOption) (*ResolverFirewallDomainList, error) {
	var resource ResolverFirewallDomainList
	err := ctx.ReadResource("aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverFirewallDomainList resources.
type resolverFirewallDomainListState struct {
	// The ARN (Amazon Resource Name) of the domain list.
	Arn *string `pulumi:"arn"`
	// A array of domains for the firewall domain list.
	Domains []string `pulumi:"domains"`
	// A name that lets you identify the domain list, to manage and use it.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. f configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ResolverFirewallDomainListState struct {
	// The ARN (Amazon Resource Name) of the domain list.
	Arn pulumi.StringPtrInput
	// A array of domains for the firewall domain list.
	Domains pulumi.StringArrayInput
	// A name that lets you identify the domain list, to manage and use it.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. f configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (ResolverFirewallDomainListState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallDomainListState)(nil)).Elem()
}

type resolverFirewallDomainListArgs struct {
	// A array of domains for the firewall domain list.
	Domains []string `pulumi:"domains"`
	// A name that lets you identify the domain list, to manage and use it.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. f configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResolverFirewallDomainList resource.
type ResolverFirewallDomainListArgs struct {
	// A array of domains for the firewall domain list.
	Domains pulumi.StringArrayInput
	// A name that lets you identify the domain list, to manage and use it.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. f configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ResolverFirewallDomainListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallDomainListArgs)(nil)).Elem()
}

type ResolverFirewallDomainListInput interface {
	pulumi.Input

	ToResolverFirewallDomainListOutput() ResolverFirewallDomainListOutput
	ToResolverFirewallDomainListOutputWithContext(ctx context.Context) ResolverFirewallDomainListOutput
}

func (*ResolverFirewallDomainList) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallDomainList)(nil)).Elem()
}

func (i *ResolverFirewallDomainList) ToResolverFirewallDomainListOutput() ResolverFirewallDomainListOutput {
	return i.ToResolverFirewallDomainListOutputWithContext(context.Background())
}

func (i *ResolverFirewallDomainList) ToResolverFirewallDomainListOutputWithContext(ctx context.Context) ResolverFirewallDomainListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallDomainListOutput)
}

func (i *ResolverFirewallDomainList) ToOutput(ctx context.Context) pulumix.Output[*ResolverFirewallDomainList] {
	return pulumix.Output[*ResolverFirewallDomainList]{
		OutputState: i.ToResolverFirewallDomainListOutputWithContext(ctx).OutputState,
	}
}

// ResolverFirewallDomainListArrayInput is an input type that accepts ResolverFirewallDomainListArray and ResolverFirewallDomainListArrayOutput values.
// You can construct a concrete instance of `ResolverFirewallDomainListArrayInput` via:
//
//	ResolverFirewallDomainListArray{ ResolverFirewallDomainListArgs{...} }
type ResolverFirewallDomainListArrayInput interface {
	pulumi.Input

	ToResolverFirewallDomainListArrayOutput() ResolverFirewallDomainListArrayOutput
	ToResolverFirewallDomainListArrayOutputWithContext(context.Context) ResolverFirewallDomainListArrayOutput
}

type ResolverFirewallDomainListArray []ResolverFirewallDomainListInput

func (ResolverFirewallDomainListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverFirewallDomainList)(nil)).Elem()
}

func (i ResolverFirewallDomainListArray) ToResolverFirewallDomainListArrayOutput() ResolverFirewallDomainListArrayOutput {
	return i.ToResolverFirewallDomainListArrayOutputWithContext(context.Background())
}

func (i ResolverFirewallDomainListArray) ToResolverFirewallDomainListArrayOutputWithContext(ctx context.Context) ResolverFirewallDomainListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallDomainListArrayOutput)
}

func (i ResolverFirewallDomainListArray) ToOutput(ctx context.Context) pulumix.Output[[]*ResolverFirewallDomainList] {
	return pulumix.Output[[]*ResolverFirewallDomainList]{
		OutputState: i.ToResolverFirewallDomainListArrayOutputWithContext(ctx).OutputState,
	}
}

// ResolverFirewallDomainListMapInput is an input type that accepts ResolverFirewallDomainListMap and ResolverFirewallDomainListMapOutput values.
// You can construct a concrete instance of `ResolverFirewallDomainListMapInput` via:
//
//	ResolverFirewallDomainListMap{ "key": ResolverFirewallDomainListArgs{...} }
type ResolverFirewallDomainListMapInput interface {
	pulumi.Input

	ToResolverFirewallDomainListMapOutput() ResolverFirewallDomainListMapOutput
	ToResolverFirewallDomainListMapOutputWithContext(context.Context) ResolverFirewallDomainListMapOutput
}

type ResolverFirewallDomainListMap map[string]ResolverFirewallDomainListInput

func (ResolverFirewallDomainListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverFirewallDomainList)(nil)).Elem()
}

func (i ResolverFirewallDomainListMap) ToResolverFirewallDomainListMapOutput() ResolverFirewallDomainListMapOutput {
	return i.ToResolverFirewallDomainListMapOutputWithContext(context.Background())
}

func (i ResolverFirewallDomainListMap) ToResolverFirewallDomainListMapOutputWithContext(ctx context.Context) ResolverFirewallDomainListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallDomainListMapOutput)
}

func (i ResolverFirewallDomainListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ResolverFirewallDomainList] {
	return pulumix.Output[map[string]*ResolverFirewallDomainList]{
		OutputState: i.ToResolverFirewallDomainListMapOutputWithContext(ctx).OutputState,
	}
}

type ResolverFirewallDomainListOutput struct{ *pulumi.OutputState }

func (ResolverFirewallDomainListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallDomainList)(nil)).Elem()
}

func (o ResolverFirewallDomainListOutput) ToResolverFirewallDomainListOutput() ResolverFirewallDomainListOutput {
	return o
}

func (o ResolverFirewallDomainListOutput) ToResolverFirewallDomainListOutputWithContext(ctx context.Context) ResolverFirewallDomainListOutput {
	return o
}

func (o ResolverFirewallDomainListOutput) ToOutput(ctx context.Context) pulumix.Output[*ResolverFirewallDomainList] {
	return pulumix.Output[*ResolverFirewallDomainList]{
		OutputState: o.OutputState,
	}
}

// The ARN (Amazon Resource Name) of the domain list.
func (o ResolverFirewallDomainListOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallDomainList) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A array of domains for the firewall domain list.
func (o ResolverFirewallDomainListOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResolverFirewallDomainList) pulumi.StringArrayOutput { return v.Domains }).(pulumi.StringArrayOutput)
}

// A name that lets you identify the domain list, to manage and use it.
func (o ResolverFirewallDomainListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallDomainList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. f configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ResolverFirewallDomainListOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResolverFirewallDomainList) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ResolverFirewallDomainListOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResolverFirewallDomainList) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type ResolverFirewallDomainListArrayOutput struct{ *pulumi.OutputState }

func (ResolverFirewallDomainListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverFirewallDomainList)(nil)).Elem()
}

func (o ResolverFirewallDomainListArrayOutput) ToResolverFirewallDomainListArrayOutput() ResolverFirewallDomainListArrayOutput {
	return o
}

func (o ResolverFirewallDomainListArrayOutput) ToResolverFirewallDomainListArrayOutputWithContext(ctx context.Context) ResolverFirewallDomainListArrayOutput {
	return o
}

func (o ResolverFirewallDomainListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ResolverFirewallDomainList] {
	return pulumix.Output[[]*ResolverFirewallDomainList]{
		OutputState: o.OutputState,
	}
}

func (o ResolverFirewallDomainListArrayOutput) Index(i pulumi.IntInput) ResolverFirewallDomainListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResolverFirewallDomainList {
		return vs[0].([]*ResolverFirewallDomainList)[vs[1].(int)]
	}).(ResolverFirewallDomainListOutput)
}

type ResolverFirewallDomainListMapOutput struct{ *pulumi.OutputState }

func (ResolverFirewallDomainListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverFirewallDomainList)(nil)).Elem()
}

func (o ResolverFirewallDomainListMapOutput) ToResolverFirewallDomainListMapOutput() ResolverFirewallDomainListMapOutput {
	return o
}

func (o ResolverFirewallDomainListMapOutput) ToResolverFirewallDomainListMapOutputWithContext(ctx context.Context) ResolverFirewallDomainListMapOutput {
	return o
}

func (o ResolverFirewallDomainListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ResolverFirewallDomainList] {
	return pulumix.Output[map[string]*ResolverFirewallDomainList]{
		OutputState: o.OutputState,
	}
}

func (o ResolverFirewallDomainListMapOutput) MapIndex(k pulumi.StringInput) ResolverFirewallDomainListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResolverFirewallDomainList {
		return vs[0].(map[string]*ResolverFirewallDomainList)[vs[1].(string)]
	}).(ResolverFirewallDomainListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallDomainListInput)(nil)).Elem(), &ResolverFirewallDomainList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallDomainListArrayInput)(nil)).Elem(), ResolverFirewallDomainListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallDomainListMapInput)(nil)).Elem(), ResolverFirewallDomainListMap{})
	pulumi.RegisterOutputType(ResolverFirewallDomainListOutput{})
	pulumi.RegisterOutputType(ResolverFirewallDomainListArrayOutput{})
	pulumi.RegisterOutputType(ResolverFirewallDomainListMapOutput{})
}
