// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages an outbound (egress) rule for a security group.
//
// When specifying an outbound rule for your security group in a VPC, the configuration must include a destination for the traffic.
//
// > **NOTE on Security Groups and Security Group Rules:** this provider currently provides a Security Group resource with `ingress` and `egress` rules defined in-line and a Security Group Rule resource which manages one or more `ingress` or
// `egress` rules. Both of these resource were added before AWS assigned a [security group rule unique ID](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules.html), and they do not work well in all scenarios using the`description` and `tags` attributes, which rely on the unique ID.
// The `vpc.SecurityGroupEgressRule` resource has been added to address these limitations and should be used for all new security group rules.
// You should not use the `vpc.SecurityGroupEgressRule` resource in conjunction with an `ec2.SecurityGroup` resource with in-line rules or with `ec2.SecurityGroupRule` resources defined for the same Security Group, as rule conflicts may occur and rules will be overwritten.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewSecurityGroupEgressRule(ctx, "example", &vpc.SecurityGroupEgressRuleArgs{
//				SecurityGroupId: pulumi.Any(aws_security_group.Example.Id),
//				CidrIpv4:        pulumi.String("10.0.0.0/8"),
//				FromPort:        pulumi.Int(80),
//				IpProtocol:      pulumi.String("tcp"),
//				ToPort:          pulumi.Int(80),
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
// Using `pulumi import`, import security group egress rules using the `security_group_rule_id`. For example:
//
// ```sh
//
//	$ pulumi import aws:vpc/securityGroupEgressRule:SecurityGroupEgressRule example sgr-02108b27edd666983
//
// ```
type SecurityGroupEgressRule struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the security group rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The destination IPv4 CIDR range.
	CidrIpv4 pulumi.StringPtrOutput `pulumi:"cidrIpv4"`
	// The destination IPv6 CIDR range.
	CidrIpv6 pulumi.StringPtrOutput `pulumi:"cidrIpv6"`
	// The security group rule description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort pulumi.IntPtrOutput `pulumi:"fromPort"`
	// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ipProtocol` is set to `-1`, it translates to all protocols, all port ranges, and `fromPort` and `toPort` values should not be defined.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The ID of the destination prefix list.
	PrefixListId pulumi.StringPtrOutput `pulumi:"prefixListId"`
	// The destination security group that is referenced in the rule.
	ReferencedSecurityGroupId pulumi.StringPtrOutput `pulumi:"referencedSecurityGroupId"`
	// The ID of the security group.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The ID of the security group rule.
	SecurityGroupRuleId pulumi.StringOutput `pulumi:"securityGroupRuleId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort pulumi.IntPtrOutput `pulumi:"toPort"`
}

// NewSecurityGroupEgressRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroupEgressRule(ctx *pulumi.Context,
	name string, args *SecurityGroupEgressRuleArgs, opts ...pulumi.ResourceOption) (*SecurityGroupEgressRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpProtocol == nil {
		return nil, errors.New("invalid value for required argument 'IpProtocol'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityGroupEgressRule
	err := ctx.RegisterResource("aws:vpc/securityGroupEgressRule:SecurityGroupEgressRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroupEgressRule gets an existing SecurityGroupEgressRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroupEgressRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupEgressRuleState, opts ...pulumi.ResourceOption) (*SecurityGroupEgressRule, error) {
	var resource SecurityGroupEgressRule
	err := ctx.ReadResource("aws:vpc/securityGroupEgressRule:SecurityGroupEgressRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroupEgressRule resources.
type securityGroupEgressRuleState struct {
	// The Amazon Resource Name (ARN) of the security group rule.
	Arn *string `pulumi:"arn"`
	// The destination IPv4 CIDR range.
	CidrIpv4 *string `pulumi:"cidrIpv4"`
	// The destination IPv6 CIDR range.
	CidrIpv6 *string `pulumi:"cidrIpv6"`
	// The security group rule description.
	Description *string `pulumi:"description"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort *int `pulumi:"fromPort"`
	// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ipProtocol` is set to `-1`, it translates to all protocols, all port ranges, and `fromPort` and `toPort` values should not be defined.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The ID of the destination prefix list.
	PrefixListId *string `pulumi:"prefixListId"`
	// The destination security group that is referenced in the rule.
	ReferencedSecurityGroupId *string `pulumi:"referencedSecurityGroupId"`
	// The ID of the security group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The ID of the security group rule.
	SecurityGroupRuleId *string `pulumi:"securityGroupRuleId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort *int `pulumi:"toPort"`
}

type SecurityGroupEgressRuleState struct {
	// The Amazon Resource Name (ARN) of the security group rule.
	Arn pulumi.StringPtrInput
	// The destination IPv4 CIDR range.
	CidrIpv4 pulumi.StringPtrInput
	// The destination IPv6 CIDR range.
	CidrIpv6 pulumi.StringPtrInput
	// The security group rule description.
	Description pulumi.StringPtrInput
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort pulumi.IntPtrInput
	// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ipProtocol` is set to `-1`, it translates to all protocols, all port ranges, and `fromPort` and `toPort` values should not be defined.
	IpProtocol pulumi.StringPtrInput
	// The ID of the destination prefix list.
	PrefixListId pulumi.StringPtrInput
	// The destination security group that is referenced in the rule.
	ReferencedSecurityGroupId pulumi.StringPtrInput
	// The ID of the security group.
	SecurityGroupId pulumi.StringPtrInput
	// The ID of the security group rule.
	SecurityGroupRuleId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort pulumi.IntPtrInput
}

func (SecurityGroupEgressRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupEgressRuleState)(nil)).Elem()
}

type securityGroupEgressRuleArgs struct {
	// The destination IPv4 CIDR range.
	CidrIpv4 *string `pulumi:"cidrIpv4"`
	// The destination IPv6 CIDR range.
	CidrIpv6 *string `pulumi:"cidrIpv6"`
	// The security group rule description.
	Description *string `pulumi:"description"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort *int `pulumi:"fromPort"`
	// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ipProtocol` is set to `-1`, it translates to all protocols, all port ranges, and `fromPort` and `toPort` values should not be defined.
	IpProtocol string `pulumi:"ipProtocol"`
	// The ID of the destination prefix list.
	PrefixListId *string `pulumi:"prefixListId"`
	// The destination security group that is referenced in the rule.
	ReferencedSecurityGroupId *string `pulumi:"referencedSecurityGroupId"`
	// The ID of the security group.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort *int `pulumi:"toPort"`
}

// The set of arguments for constructing a SecurityGroupEgressRule resource.
type SecurityGroupEgressRuleArgs struct {
	// The destination IPv4 CIDR range.
	CidrIpv4 pulumi.StringPtrInput
	// The destination IPv6 CIDR range.
	CidrIpv6 pulumi.StringPtrInput
	// The security group rule description.
	Description pulumi.StringPtrInput
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort pulumi.IntPtrInput
	// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ipProtocol` is set to `-1`, it translates to all protocols, all port ranges, and `fromPort` and `toPort` values should not be defined.
	IpProtocol pulumi.StringInput
	// The ID of the destination prefix list.
	PrefixListId pulumi.StringPtrInput
	// The destination security group that is referenced in the rule.
	ReferencedSecurityGroupId pulumi.StringPtrInput
	// The ID of the security group.
	SecurityGroupId pulumi.StringInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort pulumi.IntPtrInput
}

func (SecurityGroupEgressRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupEgressRuleArgs)(nil)).Elem()
}

type SecurityGroupEgressRuleInput interface {
	pulumi.Input

	ToSecurityGroupEgressRuleOutput() SecurityGroupEgressRuleOutput
	ToSecurityGroupEgressRuleOutputWithContext(ctx context.Context) SecurityGroupEgressRuleOutput
}

func (*SecurityGroupEgressRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupEgressRule)(nil)).Elem()
}

func (i *SecurityGroupEgressRule) ToSecurityGroupEgressRuleOutput() SecurityGroupEgressRuleOutput {
	return i.ToSecurityGroupEgressRuleOutputWithContext(context.Background())
}

func (i *SecurityGroupEgressRule) ToSecurityGroupEgressRuleOutputWithContext(ctx context.Context) SecurityGroupEgressRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupEgressRuleOutput)
}

func (i *SecurityGroupEgressRule) ToOutput(ctx context.Context) pulumix.Output[*SecurityGroupEgressRule] {
	return pulumix.Output[*SecurityGroupEgressRule]{
		OutputState: i.ToSecurityGroupEgressRuleOutputWithContext(ctx).OutputState,
	}
}

// SecurityGroupEgressRuleArrayInput is an input type that accepts SecurityGroupEgressRuleArray and SecurityGroupEgressRuleArrayOutput values.
// You can construct a concrete instance of `SecurityGroupEgressRuleArrayInput` via:
//
//	SecurityGroupEgressRuleArray{ SecurityGroupEgressRuleArgs{...} }
type SecurityGroupEgressRuleArrayInput interface {
	pulumi.Input

	ToSecurityGroupEgressRuleArrayOutput() SecurityGroupEgressRuleArrayOutput
	ToSecurityGroupEgressRuleArrayOutputWithContext(context.Context) SecurityGroupEgressRuleArrayOutput
}

type SecurityGroupEgressRuleArray []SecurityGroupEgressRuleInput

func (SecurityGroupEgressRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroupEgressRule)(nil)).Elem()
}

func (i SecurityGroupEgressRuleArray) ToSecurityGroupEgressRuleArrayOutput() SecurityGroupEgressRuleArrayOutput {
	return i.ToSecurityGroupEgressRuleArrayOutputWithContext(context.Background())
}

func (i SecurityGroupEgressRuleArray) ToSecurityGroupEgressRuleArrayOutputWithContext(ctx context.Context) SecurityGroupEgressRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupEgressRuleArrayOutput)
}

func (i SecurityGroupEgressRuleArray) ToOutput(ctx context.Context) pulumix.Output[[]*SecurityGroupEgressRule] {
	return pulumix.Output[[]*SecurityGroupEgressRule]{
		OutputState: i.ToSecurityGroupEgressRuleArrayOutputWithContext(ctx).OutputState,
	}
}

// SecurityGroupEgressRuleMapInput is an input type that accepts SecurityGroupEgressRuleMap and SecurityGroupEgressRuleMapOutput values.
// You can construct a concrete instance of `SecurityGroupEgressRuleMapInput` via:
//
//	SecurityGroupEgressRuleMap{ "key": SecurityGroupEgressRuleArgs{...} }
type SecurityGroupEgressRuleMapInput interface {
	pulumi.Input

	ToSecurityGroupEgressRuleMapOutput() SecurityGroupEgressRuleMapOutput
	ToSecurityGroupEgressRuleMapOutputWithContext(context.Context) SecurityGroupEgressRuleMapOutput
}

type SecurityGroupEgressRuleMap map[string]SecurityGroupEgressRuleInput

func (SecurityGroupEgressRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroupEgressRule)(nil)).Elem()
}

func (i SecurityGroupEgressRuleMap) ToSecurityGroupEgressRuleMapOutput() SecurityGroupEgressRuleMapOutput {
	return i.ToSecurityGroupEgressRuleMapOutputWithContext(context.Background())
}

func (i SecurityGroupEgressRuleMap) ToSecurityGroupEgressRuleMapOutputWithContext(ctx context.Context) SecurityGroupEgressRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupEgressRuleMapOutput)
}

func (i SecurityGroupEgressRuleMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecurityGroupEgressRule] {
	return pulumix.Output[map[string]*SecurityGroupEgressRule]{
		OutputState: i.ToSecurityGroupEgressRuleMapOutputWithContext(ctx).OutputState,
	}
}

type SecurityGroupEgressRuleOutput struct{ *pulumi.OutputState }

func (SecurityGroupEgressRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupEgressRule)(nil)).Elem()
}

func (o SecurityGroupEgressRuleOutput) ToSecurityGroupEgressRuleOutput() SecurityGroupEgressRuleOutput {
	return o
}

func (o SecurityGroupEgressRuleOutput) ToSecurityGroupEgressRuleOutputWithContext(ctx context.Context) SecurityGroupEgressRuleOutput {
	return o
}

func (o SecurityGroupEgressRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*SecurityGroupEgressRule] {
	return pulumix.Output[*SecurityGroupEgressRule]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the security group rule.
func (o SecurityGroupEgressRuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The destination IPv4 CIDR range.
func (o SecurityGroupEgressRuleOutput) CidrIpv4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringPtrOutput { return v.CidrIpv4 }).(pulumi.StringPtrOutput)
}

// The destination IPv6 CIDR range.
func (o SecurityGroupEgressRuleOutput) CidrIpv6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringPtrOutput { return v.CidrIpv6 }).(pulumi.StringPtrOutput)
}

// The security group rule description.
func (o SecurityGroupEgressRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
func (o SecurityGroupEgressRuleOutput) FromPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.IntPtrOutput { return v.FromPort }).(pulumi.IntPtrOutput)
}

// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ipProtocol` is set to `-1`, it translates to all protocols, all port ranges, and `fromPort` and `toPort` values should not be defined.
func (o SecurityGroupEgressRuleOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringOutput { return v.IpProtocol }).(pulumi.StringOutput)
}

// The ID of the destination prefix list.
func (o SecurityGroupEgressRuleOutput) PrefixListId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringPtrOutput { return v.PrefixListId }).(pulumi.StringPtrOutput)
}

// The destination security group that is referenced in the rule.
func (o SecurityGroupEgressRuleOutput) ReferencedSecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringPtrOutput { return v.ReferencedSecurityGroupId }).(pulumi.StringPtrOutput)
}

// The ID of the security group.
func (o SecurityGroupEgressRuleOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The ID of the security group rule.
func (o SecurityGroupEgressRuleOutput) SecurityGroupRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringOutput { return v.SecurityGroupRuleId }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SecurityGroupEgressRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o SecurityGroupEgressRuleOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
func (o SecurityGroupEgressRuleOutput) ToPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgressRule) pulumi.IntPtrOutput { return v.ToPort }).(pulumi.IntPtrOutput)
}

type SecurityGroupEgressRuleArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupEgressRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroupEgressRule)(nil)).Elem()
}

func (o SecurityGroupEgressRuleArrayOutput) ToSecurityGroupEgressRuleArrayOutput() SecurityGroupEgressRuleArrayOutput {
	return o
}

func (o SecurityGroupEgressRuleArrayOutput) ToSecurityGroupEgressRuleArrayOutputWithContext(ctx context.Context) SecurityGroupEgressRuleArrayOutput {
	return o
}

func (o SecurityGroupEgressRuleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SecurityGroupEgressRule] {
	return pulumix.Output[[]*SecurityGroupEgressRule]{
		OutputState: o.OutputState,
	}
}

func (o SecurityGroupEgressRuleArrayOutput) Index(i pulumi.IntInput) SecurityGroupEgressRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityGroupEgressRule {
		return vs[0].([]*SecurityGroupEgressRule)[vs[1].(int)]
	}).(SecurityGroupEgressRuleOutput)
}

type SecurityGroupEgressRuleMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupEgressRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroupEgressRule)(nil)).Elem()
}

func (o SecurityGroupEgressRuleMapOutput) ToSecurityGroupEgressRuleMapOutput() SecurityGroupEgressRuleMapOutput {
	return o
}

func (o SecurityGroupEgressRuleMapOutput) ToSecurityGroupEgressRuleMapOutputWithContext(ctx context.Context) SecurityGroupEgressRuleMapOutput {
	return o
}

func (o SecurityGroupEgressRuleMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecurityGroupEgressRule] {
	return pulumix.Output[map[string]*SecurityGroupEgressRule]{
		OutputState: o.OutputState,
	}
}

func (o SecurityGroupEgressRuleMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupEgressRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityGroupEgressRule {
		return vs[0].(map[string]*SecurityGroupEgressRule)[vs[1].(string)]
	}).(SecurityGroupEgressRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupEgressRuleInput)(nil)).Elem(), &SecurityGroupEgressRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupEgressRuleArrayInput)(nil)).Elem(), SecurityGroupEgressRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupEgressRuleMapInput)(nil)).Elem(), SecurityGroupEgressRuleMap{})
	pulumi.RegisterOutputType(SecurityGroupEgressRuleOutput{})
	pulumi.RegisterOutputType(SecurityGroupEgressRuleArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupEgressRuleMapOutput{})
}
