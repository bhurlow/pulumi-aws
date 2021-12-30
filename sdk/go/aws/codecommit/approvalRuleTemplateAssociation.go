// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associates a CodeCommit Approval Rule Template with a Repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codecommit"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := codecommit.NewApprovalRuleTemplateAssociation(ctx, "example", &codecommit.ApprovalRuleTemplateAssociationArgs{
// 			ApprovalRuleTemplateName: pulumi.Any(aws_codecommit_approval_rule_template.Example.Name),
// 			RepositoryName:           pulumi.Any(aws_codecommit_repository.Example.Repository_name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// CodeCommit approval rule template associations can be imported using the `approval_rule_template_name` and `repository_name` separated by a comma (`,`), e.g.
//
// ```sh
//  $ pulumi import aws:codecommit/approvalRuleTemplateAssociation:ApprovalRuleTemplateAssociation example approver-rule-for-example,MyExampleRepo
// ```
type ApprovalRuleTemplateAssociation struct {
	pulumi.CustomResourceState

	// The name for the approval rule template.
	ApprovalRuleTemplateName pulumi.StringOutput `pulumi:"approvalRuleTemplateName"`
	// The name of the repository that you want to associate with the template.
	RepositoryName pulumi.StringOutput `pulumi:"repositoryName"`
}

// NewApprovalRuleTemplateAssociation registers a new resource with the given unique name, arguments, and options.
func NewApprovalRuleTemplateAssociation(ctx *pulumi.Context,
	name string, args *ApprovalRuleTemplateAssociationArgs, opts ...pulumi.ResourceOption) (*ApprovalRuleTemplateAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApprovalRuleTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'ApprovalRuleTemplateName'")
	}
	if args.RepositoryName == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryName'")
	}
	var resource ApprovalRuleTemplateAssociation
	err := ctx.RegisterResource("aws:codecommit/approvalRuleTemplateAssociation:ApprovalRuleTemplateAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApprovalRuleTemplateAssociation gets an existing ApprovalRuleTemplateAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApprovalRuleTemplateAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApprovalRuleTemplateAssociationState, opts ...pulumi.ResourceOption) (*ApprovalRuleTemplateAssociation, error) {
	var resource ApprovalRuleTemplateAssociation
	err := ctx.ReadResource("aws:codecommit/approvalRuleTemplateAssociation:ApprovalRuleTemplateAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApprovalRuleTemplateAssociation resources.
type approvalRuleTemplateAssociationState struct {
	// The name for the approval rule template.
	ApprovalRuleTemplateName *string `pulumi:"approvalRuleTemplateName"`
	// The name of the repository that you want to associate with the template.
	RepositoryName *string `pulumi:"repositoryName"`
}

type ApprovalRuleTemplateAssociationState struct {
	// The name for the approval rule template.
	ApprovalRuleTemplateName pulumi.StringPtrInput
	// The name of the repository that you want to associate with the template.
	RepositoryName pulumi.StringPtrInput
}

func (ApprovalRuleTemplateAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*approvalRuleTemplateAssociationState)(nil)).Elem()
}

type approvalRuleTemplateAssociationArgs struct {
	// The name for the approval rule template.
	ApprovalRuleTemplateName string `pulumi:"approvalRuleTemplateName"`
	// The name of the repository that you want to associate with the template.
	RepositoryName string `pulumi:"repositoryName"`
}

// The set of arguments for constructing a ApprovalRuleTemplateAssociation resource.
type ApprovalRuleTemplateAssociationArgs struct {
	// The name for the approval rule template.
	ApprovalRuleTemplateName pulumi.StringInput
	// The name of the repository that you want to associate with the template.
	RepositoryName pulumi.StringInput
}

func (ApprovalRuleTemplateAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*approvalRuleTemplateAssociationArgs)(nil)).Elem()
}

type ApprovalRuleTemplateAssociationInput interface {
	pulumi.Input

	ToApprovalRuleTemplateAssociationOutput() ApprovalRuleTemplateAssociationOutput
	ToApprovalRuleTemplateAssociationOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationOutput
}

func (*ApprovalRuleTemplateAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ApprovalRuleTemplateAssociation)(nil))
}

func (i *ApprovalRuleTemplateAssociation) ToApprovalRuleTemplateAssociationOutput() ApprovalRuleTemplateAssociationOutput {
	return i.ToApprovalRuleTemplateAssociationOutputWithContext(context.Background())
}

func (i *ApprovalRuleTemplateAssociation) ToApprovalRuleTemplateAssociationOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApprovalRuleTemplateAssociationOutput)
}

func (i *ApprovalRuleTemplateAssociation) ToApprovalRuleTemplateAssociationPtrOutput() ApprovalRuleTemplateAssociationPtrOutput {
	return i.ToApprovalRuleTemplateAssociationPtrOutputWithContext(context.Background())
}

func (i *ApprovalRuleTemplateAssociation) ToApprovalRuleTemplateAssociationPtrOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApprovalRuleTemplateAssociationPtrOutput)
}

type ApprovalRuleTemplateAssociationPtrInput interface {
	pulumi.Input

	ToApprovalRuleTemplateAssociationPtrOutput() ApprovalRuleTemplateAssociationPtrOutput
	ToApprovalRuleTemplateAssociationPtrOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationPtrOutput
}

type approvalRuleTemplateAssociationPtrType ApprovalRuleTemplateAssociationArgs

func (*approvalRuleTemplateAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApprovalRuleTemplateAssociation)(nil))
}

func (i *approvalRuleTemplateAssociationPtrType) ToApprovalRuleTemplateAssociationPtrOutput() ApprovalRuleTemplateAssociationPtrOutput {
	return i.ToApprovalRuleTemplateAssociationPtrOutputWithContext(context.Background())
}

func (i *approvalRuleTemplateAssociationPtrType) ToApprovalRuleTemplateAssociationPtrOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApprovalRuleTemplateAssociationPtrOutput)
}

// ApprovalRuleTemplateAssociationArrayInput is an input type that accepts ApprovalRuleTemplateAssociationArray and ApprovalRuleTemplateAssociationArrayOutput values.
// You can construct a concrete instance of `ApprovalRuleTemplateAssociationArrayInput` via:
//
//          ApprovalRuleTemplateAssociationArray{ ApprovalRuleTemplateAssociationArgs{...} }
type ApprovalRuleTemplateAssociationArrayInput interface {
	pulumi.Input

	ToApprovalRuleTemplateAssociationArrayOutput() ApprovalRuleTemplateAssociationArrayOutput
	ToApprovalRuleTemplateAssociationArrayOutputWithContext(context.Context) ApprovalRuleTemplateAssociationArrayOutput
}

type ApprovalRuleTemplateAssociationArray []ApprovalRuleTemplateAssociationInput

func (ApprovalRuleTemplateAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApprovalRuleTemplateAssociation)(nil)).Elem()
}

func (i ApprovalRuleTemplateAssociationArray) ToApprovalRuleTemplateAssociationArrayOutput() ApprovalRuleTemplateAssociationArrayOutput {
	return i.ToApprovalRuleTemplateAssociationArrayOutputWithContext(context.Background())
}

func (i ApprovalRuleTemplateAssociationArray) ToApprovalRuleTemplateAssociationArrayOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApprovalRuleTemplateAssociationArrayOutput)
}

// ApprovalRuleTemplateAssociationMapInput is an input type that accepts ApprovalRuleTemplateAssociationMap and ApprovalRuleTemplateAssociationMapOutput values.
// You can construct a concrete instance of `ApprovalRuleTemplateAssociationMapInput` via:
//
//          ApprovalRuleTemplateAssociationMap{ "key": ApprovalRuleTemplateAssociationArgs{...} }
type ApprovalRuleTemplateAssociationMapInput interface {
	pulumi.Input

	ToApprovalRuleTemplateAssociationMapOutput() ApprovalRuleTemplateAssociationMapOutput
	ToApprovalRuleTemplateAssociationMapOutputWithContext(context.Context) ApprovalRuleTemplateAssociationMapOutput
}

type ApprovalRuleTemplateAssociationMap map[string]ApprovalRuleTemplateAssociationInput

func (ApprovalRuleTemplateAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApprovalRuleTemplateAssociation)(nil)).Elem()
}

func (i ApprovalRuleTemplateAssociationMap) ToApprovalRuleTemplateAssociationMapOutput() ApprovalRuleTemplateAssociationMapOutput {
	return i.ToApprovalRuleTemplateAssociationMapOutputWithContext(context.Background())
}

func (i ApprovalRuleTemplateAssociationMap) ToApprovalRuleTemplateAssociationMapOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApprovalRuleTemplateAssociationMapOutput)
}

type ApprovalRuleTemplateAssociationOutput struct{ *pulumi.OutputState }

func (ApprovalRuleTemplateAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApprovalRuleTemplateAssociation)(nil))
}

func (o ApprovalRuleTemplateAssociationOutput) ToApprovalRuleTemplateAssociationOutput() ApprovalRuleTemplateAssociationOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationOutput) ToApprovalRuleTemplateAssociationOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationOutput) ToApprovalRuleTemplateAssociationPtrOutput() ApprovalRuleTemplateAssociationPtrOutput {
	return o.ToApprovalRuleTemplateAssociationPtrOutputWithContext(context.Background())
}

func (o ApprovalRuleTemplateAssociationOutput) ToApprovalRuleTemplateAssociationPtrOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApprovalRuleTemplateAssociation) *ApprovalRuleTemplateAssociation {
		return &v
	}).(ApprovalRuleTemplateAssociationPtrOutput)
}

type ApprovalRuleTemplateAssociationPtrOutput struct{ *pulumi.OutputState }

func (ApprovalRuleTemplateAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApprovalRuleTemplateAssociation)(nil))
}

func (o ApprovalRuleTemplateAssociationPtrOutput) ToApprovalRuleTemplateAssociationPtrOutput() ApprovalRuleTemplateAssociationPtrOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationPtrOutput) ToApprovalRuleTemplateAssociationPtrOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationPtrOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationPtrOutput) Elem() ApprovalRuleTemplateAssociationOutput {
	return o.ApplyT(func(v *ApprovalRuleTemplateAssociation) ApprovalRuleTemplateAssociation {
		if v != nil {
			return *v
		}
		var ret ApprovalRuleTemplateAssociation
		return ret
	}).(ApprovalRuleTemplateAssociationOutput)
}

type ApprovalRuleTemplateAssociationArrayOutput struct{ *pulumi.OutputState }

func (ApprovalRuleTemplateAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApprovalRuleTemplateAssociation)(nil))
}

func (o ApprovalRuleTemplateAssociationArrayOutput) ToApprovalRuleTemplateAssociationArrayOutput() ApprovalRuleTemplateAssociationArrayOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationArrayOutput) ToApprovalRuleTemplateAssociationArrayOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationArrayOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationArrayOutput) Index(i pulumi.IntInput) ApprovalRuleTemplateAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApprovalRuleTemplateAssociation {
		return vs[0].([]ApprovalRuleTemplateAssociation)[vs[1].(int)]
	}).(ApprovalRuleTemplateAssociationOutput)
}

type ApprovalRuleTemplateAssociationMapOutput struct{ *pulumi.OutputState }

func (ApprovalRuleTemplateAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApprovalRuleTemplateAssociation)(nil))
}

func (o ApprovalRuleTemplateAssociationMapOutput) ToApprovalRuleTemplateAssociationMapOutput() ApprovalRuleTemplateAssociationMapOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationMapOutput) ToApprovalRuleTemplateAssociationMapOutputWithContext(ctx context.Context) ApprovalRuleTemplateAssociationMapOutput {
	return o
}

func (o ApprovalRuleTemplateAssociationMapOutput) MapIndex(k pulumi.StringInput) ApprovalRuleTemplateAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApprovalRuleTemplateAssociation {
		return vs[0].(map[string]ApprovalRuleTemplateAssociation)[vs[1].(string)]
	}).(ApprovalRuleTemplateAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApprovalRuleTemplateAssociationInput)(nil)).Elem(), &ApprovalRuleTemplateAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApprovalRuleTemplateAssociationPtrInput)(nil)).Elem(), &ApprovalRuleTemplateAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApprovalRuleTemplateAssociationArrayInput)(nil)).Elem(), ApprovalRuleTemplateAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApprovalRuleTemplateAssociationMapInput)(nil)).Elem(), ApprovalRuleTemplateAssociationMap{})
	pulumi.RegisterOutputType(ApprovalRuleTemplateAssociationOutput{})
	pulumi.RegisterOutputType(ApprovalRuleTemplateAssociationPtrOutput{})
	pulumi.RegisterOutputType(ApprovalRuleTemplateAssociationArrayOutput{})
	pulumi.RegisterOutputType(ApprovalRuleTemplateAssociationMapOutput{})
}