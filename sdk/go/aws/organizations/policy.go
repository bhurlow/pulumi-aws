// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an [AWS Organizations policy](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"*",
//						},
//						Resources: []string{
//							"*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = organizations.NewPolicy(ctx, "examplePolicy", &organizations.PolicyArgs{
//				Content: *pulumi.String(examplePolicyDocument.Json),
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
// Using `pulumi import`, import `aws_organizations_policy` using the policy ID. For example:
//
// ```sh
//
//	$ pulumi import aws:organizations/policy:Policy example p-12345678
//
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The policy content to add to the new policy. For example, if you create a [service control policy (SCP)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html), this string must be JSON text that specifies the permissions that admins in attached accounts can delegate to their users, groups, and roles. For more information about the SCP syntax, see the [Service Control Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html) and for more information on the Tag Policy syntax, see the [Tag Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_example-tag-policies.html).
	Content pulumi.StringOutput `pulumi:"content"`
	// A description to assign to the policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The friendly name to assign to the policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// If set to `true`, destroy will **not** delete the policy and instead just remove the resource from state. This can be useful in situations where the policies (and the associated attachment) must be preserved to meet the AWS minimum requirement of 1 attached policy.
	SkipDestroy pulumi.BoolPtrOutput `pulumi:"skipDestroy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of policy to create. Valid values are `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY` (SCP), and `TAG_POLICY`. Defaults to `SERVICE_CONTROL_POLICY`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("aws:organizations/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("aws:organizations/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Amazon Resource Name (ARN) of the policy.
	Arn *string `pulumi:"arn"`
	// The policy content to add to the new policy. For example, if you create a [service control policy (SCP)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html), this string must be JSON text that specifies the permissions that admins in attached accounts can delegate to their users, groups, and roles. For more information about the SCP syntax, see the [Service Control Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html) and for more information on the Tag Policy syntax, see the [Tag Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_example-tag-policies.html).
	Content *string `pulumi:"content"`
	// A description to assign to the policy.
	Description *string `pulumi:"description"`
	// The friendly name to assign to the policy.
	Name *string `pulumi:"name"`
	// If set to `true`, destroy will **not** delete the policy and instead just remove the resource from state. This can be useful in situations where the policies (and the associated attachment) must be preserved to meet the AWS minimum requirement of 1 attached policy.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of policy to create. Valid values are `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY` (SCP), and `TAG_POLICY`. Defaults to `SERVICE_CONTROL_POLICY`.
	Type *string `pulumi:"type"`
}

type PolicyState struct {
	// Amazon Resource Name (ARN) of the policy.
	Arn pulumi.StringPtrInput
	// The policy content to add to the new policy. For example, if you create a [service control policy (SCP)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html), this string must be JSON text that specifies the permissions that admins in attached accounts can delegate to their users, groups, and roles. For more information about the SCP syntax, see the [Service Control Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html) and for more information on the Tag Policy syntax, see the [Tag Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_example-tag-policies.html).
	Content pulumi.StringPtrInput
	// A description to assign to the policy.
	Description pulumi.StringPtrInput
	// The friendly name to assign to the policy.
	Name pulumi.StringPtrInput
	// If set to `true`, destroy will **not** delete the policy and instead just remove the resource from state. This can be useful in situations where the policies (and the associated attachment) must be preserved to meet the AWS minimum requirement of 1 attached policy.
	SkipDestroy pulumi.BoolPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The type of policy to create. Valid values are `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY` (SCP), and `TAG_POLICY`. Defaults to `SERVICE_CONTROL_POLICY`.
	Type pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// The policy content to add to the new policy. For example, if you create a [service control policy (SCP)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html), this string must be JSON text that specifies the permissions that admins in attached accounts can delegate to their users, groups, and roles. For more information about the SCP syntax, see the [Service Control Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html) and for more information on the Tag Policy syntax, see the [Tag Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_example-tag-policies.html).
	Content string `pulumi:"content"`
	// A description to assign to the policy.
	Description *string `pulumi:"description"`
	// The friendly name to assign to the policy.
	Name *string `pulumi:"name"`
	// If set to `true`, destroy will **not** delete the policy and instead just remove the resource from state. This can be useful in situations where the policies (and the associated attachment) must be preserved to meet the AWS minimum requirement of 1 attached policy.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The type of policy to create. Valid values are `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY` (SCP), and `TAG_POLICY`. Defaults to `SERVICE_CONTROL_POLICY`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// The policy content to add to the new policy. For example, if you create a [service control policy (SCP)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html), this string must be JSON text that specifies the permissions that admins in attached accounts can delegate to their users, groups, and roles. For more information about the SCP syntax, see the [Service Control Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html) and for more information on the Tag Policy syntax, see the [Tag Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_example-tag-policies.html).
	Content pulumi.StringInput
	// A description to assign to the policy.
	Description pulumi.StringPtrInput
	// The friendly name to assign to the policy.
	Name pulumi.StringPtrInput
	// If set to `true`, destroy will **not** delete the policy and instead just remove the resource from state. This can be useful in situations where the policies (and the associated attachment) must be preserved to meet the AWS minimum requirement of 1 attached policy.
	SkipDestroy pulumi.BoolPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The type of policy to create. Valid values are `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY` (SCP), and `TAG_POLICY`. Defaults to `SERVICE_CONTROL_POLICY`.
	Type pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Amazon Resource Name (ARN) of the policy.
func (o PolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The policy content to add to the new policy. For example, if you create a [service control policy (SCP)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html), this string must be JSON text that specifies the permissions that admins in attached accounts can delegate to their users, groups, and roles. For more information about the SCP syntax, see the [Service Control Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html) and for more information on the Tag Policy syntax, see the [Tag Policy Syntax documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_example-tag-policies.html).
func (o PolicyOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// A description to assign to the policy.
func (o PolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The friendly name to assign to the policy.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If set to `true`, destroy will **not** delete the policy and instead just remove the resource from state. This can be useful in situations where the policies (and the associated attachment) must be preserved to meet the AWS minimum requirement of 1 attached policy.
func (o PolicyOutput) SkipDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolPtrOutput { return v.SkipDestroy }).(pulumi.BoolPtrOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o PolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o PolicyOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of policy to create. Valid values are `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY` (SCP), and `TAG_POLICY`. Defaults to `SERVICE_CONTROL_POLICY`.
func (o PolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
