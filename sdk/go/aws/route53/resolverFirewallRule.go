// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Route 53 Resolver DNS Firewall rule resource.
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
//			exampleResolverFirewallDomainList, err := route53.NewResolverFirewallDomainList(ctx, "exampleResolverFirewallDomainList", &route53.ResolverFirewallDomainListArgs{
//				Domains: pulumi.StringArray{
//					pulumi.String("example.com"),
//				},
//				Tags: nil,
//			})
//			if err != nil {
//				return err
//			}
//			exampleResolverFirewallRuleGroup, err := route53.NewResolverFirewallRuleGroup(ctx, "exampleResolverFirewallRuleGroup", &route53.ResolverFirewallRuleGroupArgs{
//				Tags: nil,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewResolverFirewallRule(ctx, "exampleResolverFirewallRule", &route53.ResolverFirewallRuleArgs{
//				Action:               pulumi.String("BLOCK"),
//				BlockOverrideDnsType: pulumi.String("CNAME"),
//				BlockOverrideDomain:  pulumi.String("example.com"),
//				BlockOverrideTtl:     pulumi.Int(1),
//				BlockResponse:        pulumi.String("OVERRIDE"),
//				FirewallDomainListId: exampleResolverFirewallDomainList.ID(),
//				FirewallRuleGroupId:  exampleResolverFirewallRuleGroup.ID(),
//				Priority:             pulumi.Int(100),
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
// # Using `pulumi import`, import
//
// Route 53 Resolver DNS Firewall rules using the Route 53 Resolver DNS Firewall rule group ID and domain list ID separated by ':'. For example:
//
// ```sh
//
//	$ pulumi import aws:route53/resolverFirewallRule:ResolverFirewallRule example rslvr-frg-0123456789abcdef:rslvr-fdl-0123456789abcdef
//
// ```
type ResolverFirewallRule struct {
	pulumi.CustomResourceState

	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list. Valid values: `ALLOW`, `BLOCK`, `ALERT`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The DNS record's type. This determines the format of the record value that you provided in BlockOverrideDomain. Value values: `CNAME`.
	BlockOverrideDnsType pulumi.StringPtrOutput `pulumi:"blockOverrideDnsType"`
	// The custom DNS record to send back in response to the query.
	BlockOverrideDomain pulumi.StringPtrOutput `pulumi:"blockOverrideDomain"`
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Minimum value of 0. Maximum value of 604800.
	BlockOverrideTtl pulumi.IntPtrOutput `pulumi:"blockOverrideTtl"`
	// The way that you want DNS Firewall to block the request. Valid values: `NODATA`, `NXDOMAIN`, `OVERRIDE`.
	BlockResponse pulumi.StringPtrOutput `pulumi:"blockResponse"`
	// The ID of the domain list that you want to use in the rule.
	FirewallDomainListId pulumi.StringOutput `pulumi:"firewallDomainListId"`
	// The unique identifier of the firewall rule group where you want to create the rule.
	FirewallRuleGroupId pulumi.StringOutput `pulumi:"firewallRuleGroupId"`
	// A name that lets you identify the rule, to manage and use it.
	Name pulumi.StringOutput `pulumi:"name"`
	// The setting that determines the processing order of the rule in the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority pulumi.IntOutput `pulumi:"priority"`
}

// NewResolverFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewResolverFirewallRule(ctx *pulumi.Context,
	name string, args *ResolverFirewallRuleArgs, opts ...pulumi.ResourceOption) (*ResolverFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.FirewallDomainListId == nil {
		return nil, errors.New("invalid value for required argument 'FirewallDomainListId'")
	}
	if args.FirewallRuleGroupId == nil {
		return nil, errors.New("invalid value for required argument 'FirewallRuleGroupId'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverFirewallRule
	err := ctx.RegisterResource("aws:route53/resolverFirewallRule:ResolverFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverFirewallRule gets an existing ResolverFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverFirewallRuleState, opts ...pulumi.ResourceOption) (*ResolverFirewallRule, error) {
	var resource ResolverFirewallRule
	err := ctx.ReadResource("aws:route53/resolverFirewallRule:ResolverFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverFirewallRule resources.
type resolverFirewallRuleState struct {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list. Valid values: `ALLOW`, `BLOCK`, `ALERT`.
	Action *string `pulumi:"action"`
	// The DNS record's type. This determines the format of the record value that you provided in BlockOverrideDomain. Value values: `CNAME`.
	BlockOverrideDnsType *string `pulumi:"blockOverrideDnsType"`
	// The custom DNS record to send back in response to the query.
	BlockOverrideDomain *string `pulumi:"blockOverrideDomain"`
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Minimum value of 0. Maximum value of 604800.
	BlockOverrideTtl *int `pulumi:"blockOverrideTtl"`
	// The way that you want DNS Firewall to block the request. Valid values: `NODATA`, `NXDOMAIN`, `OVERRIDE`.
	BlockResponse *string `pulumi:"blockResponse"`
	// The ID of the domain list that you want to use in the rule.
	FirewallDomainListId *string `pulumi:"firewallDomainListId"`
	// The unique identifier of the firewall rule group where you want to create the rule.
	FirewallRuleGroupId *string `pulumi:"firewallRuleGroupId"`
	// A name that lets you identify the rule, to manage and use it.
	Name *string `pulumi:"name"`
	// The setting that determines the processing order of the rule in the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority *int `pulumi:"priority"`
}

type ResolverFirewallRuleState struct {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list. Valid values: `ALLOW`, `BLOCK`, `ALERT`.
	Action pulumi.StringPtrInput
	// The DNS record's type. This determines the format of the record value that you provided in BlockOverrideDomain. Value values: `CNAME`.
	BlockOverrideDnsType pulumi.StringPtrInput
	// The custom DNS record to send back in response to the query.
	BlockOverrideDomain pulumi.StringPtrInput
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Minimum value of 0. Maximum value of 604800.
	BlockOverrideTtl pulumi.IntPtrInput
	// The way that you want DNS Firewall to block the request. Valid values: `NODATA`, `NXDOMAIN`, `OVERRIDE`.
	BlockResponse pulumi.StringPtrInput
	// The ID of the domain list that you want to use in the rule.
	FirewallDomainListId pulumi.StringPtrInput
	// The unique identifier of the firewall rule group where you want to create the rule.
	FirewallRuleGroupId pulumi.StringPtrInput
	// A name that lets you identify the rule, to manage and use it.
	Name pulumi.StringPtrInput
	// The setting that determines the processing order of the rule in the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority pulumi.IntPtrInput
}

func (ResolverFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallRuleState)(nil)).Elem()
}

type resolverFirewallRuleArgs struct {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list. Valid values: `ALLOW`, `BLOCK`, `ALERT`.
	Action string `pulumi:"action"`
	// The DNS record's type. This determines the format of the record value that you provided in BlockOverrideDomain. Value values: `CNAME`.
	BlockOverrideDnsType *string `pulumi:"blockOverrideDnsType"`
	// The custom DNS record to send back in response to the query.
	BlockOverrideDomain *string `pulumi:"blockOverrideDomain"`
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Minimum value of 0. Maximum value of 604800.
	BlockOverrideTtl *int `pulumi:"blockOverrideTtl"`
	// The way that you want DNS Firewall to block the request. Valid values: `NODATA`, `NXDOMAIN`, `OVERRIDE`.
	BlockResponse *string `pulumi:"blockResponse"`
	// The ID of the domain list that you want to use in the rule.
	FirewallDomainListId string `pulumi:"firewallDomainListId"`
	// The unique identifier of the firewall rule group where you want to create the rule.
	FirewallRuleGroupId string `pulumi:"firewallRuleGroupId"`
	// A name that lets you identify the rule, to manage and use it.
	Name *string `pulumi:"name"`
	// The setting that determines the processing order of the rule in the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority int `pulumi:"priority"`
}

// The set of arguments for constructing a ResolverFirewallRule resource.
type ResolverFirewallRuleArgs struct {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list. Valid values: `ALLOW`, `BLOCK`, `ALERT`.
	Action pulumi.StringInput
	// The DNS record's type. This determines the format of the record value that you provided in BlockOverrideDomain. Value values: `CNAME`.
	BlockOverrideDnsType pulumi.StringPtrInput
	// The custom DNS record to send back in response to the query.
	BlockOverrideDomain pulumi.StringPtrInput
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Minimum value of 0. Maximum value of 604800.
	BlockOverrideTtl pulumi.IntPtrInput
	// The way that you want DNS Firewall to block the request. Valid values: `NODATA`, `NXDOMAIN`, `OVERRIDE`.
	BlockResponse pulumi.StringPtrInput
	// The ID of the domain list that you want to use in the rule.
	FirewallDomainListId pulumi.StringInput
	// The unique identifier of the firewall rule group where you want to create the rule.
	FirewallRuleGroupId pulumi.StringInput
	// A name that lets you identify the rule, to manage and use it.
	Name pulumi.StringPtrInput
	// The setting that determines the processing order of the rule in the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority pulumi.IntInput
}

func (ResolverFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallRuleArgs)(nil)).Elem()
}

type ResolverFirewallRuleInput interface {
	pulumi.Input

	ToResolverFirewallRuleOutput() ResolverFirewallRuleOutput
	ToResolverFirewallRuleOutputWithContext(ctx context.Context) ResolverFirewallRuleOutput
}

func (*ResolverFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallRule)(nil)).Elem()
}

func (i *ResolverFirewallRule) ToResolverFirewallRuleOutput() ResolverFirewallRuleOutput {
	return i.ToResolverFirewallRuleOutputWithContext(context.Background())
}

func (i *ResolverFirewallRule) ToResolverFirewallRuleOutputWithContext(ctx context.Context) ResolverFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleOutput)
}

// ResolverFirewallRuleArrayInput is an input type that accepts ResolverFirewallRuleArray and ResolverFirewallRuleArrayOutput values.
// You can construct a concrete instance of `ResolverFirewallRuleArrayInput` via:
//
//	ResolverFirewallRuleArray{ ResolverFirewallRuleArgs{...} }
type ResolverFirewallRuleArrayInput interface {
	pulumi.Input

	ToResolverFirewallRuleArrayOutput() ResolverFirewallRuleArrayOutput
	ToResolverFirewallRuleArrayOutputWithContext(context.Context) ResolverFirewallRuleArrayOutput
}

type ResolverFirewallRuleArray []ResolverFirewallRuleInput

func (ResolverFirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverFirewallRule)(nil)).Elem()
}

func (i ResolverFirewallRuleArray) ToResolverFirewallRuleArrayOutput() ResolverFirewallRuleArrayOutput {
	return i.ToResolverFirewallRuleArrayOutputWithContext(context.Background())
}

func (i ResolverFirewallRuleArray) ToResolverFirewallRuleArrayOutputWithContext(ctx context.Context) ResolverFirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleArrayOutput)
}

// ResolverFirewallRuleMapInput is an input type that accepts ResolverFirewallRuleMap and ResolverFirewallRuleMapOutput values.
// You can construct a concrete instance of `ResolverFirewallRuleMapInput` via:
//
//	ResolverFirewallRuleMap{ "key": ResolverFirewallRuleArgs{...} }
type ResolverFirewallRuleMapInput interface {
	pulumi.Input

	ToResolverFirewallRuleMapOutput() ResolverFirewallRuleMapOutput
	ToResolverFirewallRuleMapOutputWithContext(context.Context) ResolverFirewallRuleMapOutput
}

type ResolverFirewallRuleMap map[string]ResolverFirewallRuleInput

func (ResolverFirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverFirewallRule)(nil)).Elem()
}

func (i ResolverFirewallRuleMap) ToResolverFirewallRuleMapOutput() ResolverFirewallRuleMapOutput {
	return i.ToResolverFirewallRuleMapOutputWithContext(context.Background())
}

func (i ResolverFirewallRuleMap) ToResolverFirewallRuleMapOutputWithContext(ctx context.Context) ResolverFirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleMapOutput)
}

type ResolverFirewallRuleOutput struct{ *pulumi.OutputState }

func (ResolverFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallRule)(nil)).Elem()
}

func (o ResolverFirewallRuleOutput) ToResolverFirewallRuleOutput() ResolverFirewallRuleOutput {
	return o
}

func (o ResolverFirewallRuleOutput) ToResolverFirewallRuleOutputWithContext(ctx context.Context) ResolverFirewallRuleOutput {
	return o
}

// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list. Valid values: `ALLOW`, `BLOCK`, `ALERT`.
func (o ResolverFirewallRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The DNS record's type. This determines the format of the record value that you provided in BlockOverrideDomain. Value values: `CNAME`.
func (o ResolverFirewallRuleOutput) BlockOverrideDnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.StringPtrOutput { return v.BlockOverrideDnsType }).(pulumi.StringPtrOutput)
}

// The custom DNS record to send back in response to the query.
func (o ResolverFirewallRuleOutput) BlockOverrideDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.StringPtrOutput { return v.BlockOverrideDomain }).(pulumi.StringPtrOutput)
}

// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Minimum value of 0. Maximum value of 604800.
func (o ResolverFirewallRuleOutput) BlockOverrideTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.IntPtrOutput { return v.BlockOverrideTtl }).(pulumi.IntPtrOutput)
}

// The way that you want DNS Firewall to block the request. Valid values: `NODATA`, `NXDOMAIN`, `OVERRIDE`.
func (o ResolverFirewallRuleOutput) BlockResponse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.StringPtrOutput { return v.BlockResponse }).(pulumi.StringPtrOutput)
}

// The ID of the domain list that you want to use in the rule.
func (o ResolverFirewallRuleOutput) FirewallDomainListId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.StringOutput { return v.FirewallDomainListId }).(pulumi.StringOutput)
}

// The unique identifier of the firewall rule group where you want to create the rule.
func (o ResolverFirewallRuleOutput) FirewallRuleGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.StringOutput { return v.FirewallRuleGroupId }).(pulumi.StringOutput)
}

// A name that lets you identify the rule, to manage and use it.
func (o ResolverFirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The setting that determines the processing order of the rule in the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
func (o ResolverFirewallRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *ResolverFirewallRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

type ResolverFirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (ResolverFirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverFirewallRule)(nil)).Elem()
}

func (o ResolverFirewallRuleArrayOutput) ToResolverFirewallRuleArrayOutput() ResolverFirewallRuleArrayOutput {
	return o
}

func (o ResolverFirewallRuleArrayOutput) ToResolverFirewallRuleArrayOutputWithContext(ctx context.Context) ResolverFirewallRuleArrayOutput {
	return o
}

func (o ResolverFirewallRuleArrayOutput) Index(i pulumi.IntInput) ResolverFirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResolverFirewallRule {
		return vs[0].([]*ResolverFirewallRule)[vs[1].(int)]
	}).(ResolverFirewallRuleOutput)
}

type ResolverFirewallRuleMapOutput struct{ *pulumi.OutputState }

func (ResolverFirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverFirewallRule)(nil)).Elem()
}

func (o ResolverFirewallRuleMapOutput) ToResolverFirewallRuleMapOutput() ResolverFirewallRuleMapOutput {
	return o
}

func (o ResolverFirewallRuleMapOutput) ToResolverFirewallRuleMapOutputWithContext(ctx context.Context) ResolverFirewallRuleMapOutput {
	return o
}

func (o ResolverFirewallRuleMapOutput) MapIndex(k pulumi.StringInput) ResolverFirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResolverFirewallRule {
		return vs[0].(map[string]*ResolverFirewallRule)[vs[1].(string)]
	}).(ResolverFirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallRuleInput)(nil)).Elem(), &ResolverFirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallRuleArrayInput)(nil)).Elem(), ResolverFirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallRuleMapInput)(nil)).Elem(), ResolverFirewallRuleMap{})
	pulumi.RegisterOutputType(ResolverFirewallRuleOutput{})
	pulumi.RegisterOutputType(ResolverFirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(ResolverFirewallRuleMapOutput{})
}
