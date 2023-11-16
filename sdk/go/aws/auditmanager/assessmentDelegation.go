// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package auditmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS Audit Manager Assessment Delegation.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := auditmanager.NewAssessmentDelegation(ctx, "example", &auditmanager.AssessmentDelegationArgs{
//				AssessmentId: pulumi.Any(aws_auditmanager_assessment.Example.Id),
//				RoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
//				RoleType:     pulumi.String("RESOURCE_OWNER"),
//				ControlSetId: pulumi.String("example"),
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
// Using `pulumi import`, import Audit Manager Assessment Delegation using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:auditmanager/assessmentDelegation:AssessmentDelegation example abcdef-123456,arn:aws:iam::012345678901:role/example,example
//
// ```
type AssessmentDelegation struct {
	pulumi.CustomResourceState

	// Identifier for the assessment.
	AssessmentId pulumi.StringOutput `pulumi:"assessmentId"`
	// Comment describing the delegation request.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
	ControlSetId pulumi.StringOutput `pulumi:"controlSetId"`
	// Unique identifier for the delegation.
	DelegationId pulumi.StringOutput `pulumi:"delegationId"`
	// Amazon Resource Name (ARN) of the IAM role.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.
	//
	// The following arguments are optional:
	RoleType pulumi.StringOutput `pulumi:"roleType"`
	// Status of the delegation.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAssessmentDelegation registers a new resource with the given unique name, arguments, and options.
func NewAssessmentDelegation(ctx *pulumi.Context,
	name string, args *AssessmentDelegationArgs, opts ...pulumi.ResourceOption) (*AssessmentDelegation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssessmentId == nil {
		return nil, errors.New("invalid value for required argument 'AssessmentId'")
	}
	if args.ControlSetId == nil {
		return nil, errors.New("invalid value for required argument 'ControlSetId'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.RoleType == nil {
		return nil, errors.New("invalid value for required argument 'RoleType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AssessmentDelegation
	err := ctx.RegisterResource("aws:auditmanager/assessmentDelegation:AssessmentDelegation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessmentDelegation gets an existing AssessmentDelegation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentDelegation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentDelegationState, opts ...pulumi.ResourceOption) (*AssessmentDelegation, error) {
	var resource AssessmentDelegation
	err := ctx.ReadResource("aws:auditmanager/assessmentDelegation:AssessmentDelegation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssessmentDelegation resources.
type assessmentDelegationState struct {
	// Identifier for the assessment.
	AssessmentId *string `pulumi:"assessmentId"`
	// Comment describing the delegation request.
	Comment *string `pulumi:"comment"`
	// Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
	ControlSetId *string `pulumi:"controlSetId"`
	// Unique identifier for the delegation.
	DelegationId *string `pulumi:"delegationId"`
	// Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string `pulumi:"roleArn"`
	// Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.
	//
	// The following arguments are optional:
	RoleType *string `pulumi:"roleType"`
	// Status of the delegation.
	Status *string `pulumi:"status"`
}

type AssessmentDelegationState struct {
	// Identifier for the assessment.
	AssessmentId pulumi.StringPtrInput
	// Comment describing the delegation request.
	Comment pulumi.StringPtrInput
	// Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
	ControlSetId pulumi.StringPtrInput
	// Unique identifier for the delegation.
	DelegationId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM role.
	RoleArn pulumi.StringPtrInput
	// Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.
	//
	// The following arguments are optional:
	RoleType pulumi.StringPtrInput
	// Status of the delegation.
	Status pulumi.StringPtrInput
}

func (AssessmentDelegationState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentDelegationState)(nil)).Elem()
}

type assessmentDelegationArgs struct {
	// Identifier for the assessment.
	AssessmentId string `pulumi:"assessmentId"`
	// Comment describing the delegation request.
	Comment *string `pulumi:"comment"`
	// Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
	ControlSetId string `pulumi:"controlSetId"`
	// Amazon Resource Name (ARN) of the IAM role.
	RoleArn string `pulumi:"roleArn"`
	// Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.
	//
	// The following arguments are optional:
	RoleType string `pulumi:"roleType"`
}

// The set of arguments for constructing a AssessmentDelegation resource.
type AssessmentDelegationArgs struct {
	// Identifier for the assessment.
	AssessmentId pulumi.StringInput
	// Comment describing the delegation request.
	Comment pulumi.StringPtrInput
	// Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
	ControlSetId pulumi.StringInput
	// Amazon Resource Name (ARN) of the IAM role.
	RoleArn pulumi.StringInput
	// Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.
	//
	// The following arguments are optional:
	RoleType pulumi.StringInput
}

func (AssessmentDelegationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentDelegationArgs)(nil)).Elem()
}

type AssessmentDelegationInput interface {
	pulumi.Input

	ToAssessmentDelegationOutput() AssessmentDelegationOutput
	ToAssessmentDelegationOutputWithContext(ctx context.Context) AssessmentDelegationOutput
}

func (*AssessmentDelegation) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentDelegation)(nil)).Elem()
}

func (i *AssessmentDelegation) ToAssessmentDelegationOutput() AssessmentDelegationOutput {
	return i.ToAssessmentDelegationOutputWithContext(context.Background())
}

func (i *AssessmentDelegation) ToAssessmentDelegationOutputWithContext(ctx context.Context) AssessmentDelegationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentDelegationOutput)
}

// AssessmentDelegationArrayInput is an input type that accepts AssessmentDelegationArray and AssessmentDelegationArrayOutput values.
// You can construct a concrete instance of `AssessmentDelegationArrayInput` via:
//
//	AssessmentDelegationArray{ AssessmentDelegationArgs{...} }
type AssessmentDelegationArrayInput interface {
	pulumi.Input

	ToAssessmentDelegationArrayOutput() AssessmentDelegationArrayOutput
	ToAssessmentDelegationArrayOutputWithContext(context.Context) AssessmentDelegationArrayOutput
}

type AssessmentDelegationArray []AssessmentDelegationInput

func (AssessmentDelegationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssessmentDelegation)(nil)).Elem()
}

func (i AssessmentDelegationArray) ToAssessmentDelegationArrayOutput() AssessmentDelegationArrayOutput {
	return i.ToAssessmentDelegationArrayOutputWithContext(context.Background())
}

func (i AssessmentDelegationArray) ToAssessmentDelegationArrayOutputWithContext(ctx context.Context) AssessmentDelegationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentDelegationArrayOutput)
}

// AssessmentDelegationMapInput is an input type that accepts AssessmentDelegationMap and AssessmentDelegationMapOutput values.
// You can construct a concrete instance of `AssessmentDelegationMapInput` via:
//
//	AssessmentDelegationMap{ "key": AssessmentDelegationArgs{...} }
type AssessmentDelegationMapInput interface {
	pulumi.Input

	ToAssessmentDelegationMapOutput() AssessmentDelegationMapOutput
	ToAssessmentDelegationMapOutputWithContext(context.Context) AssessmentDelegationMapOutput
}

type AssessmentDelegationMap map[string]AssessmentDelegationInput

func (AssessmentDelegationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssessmentDelegation)(nil)).Elem()
}

func (i AssessmentDelegationMap) ToAssessmentDelegationMapOutput() AssessmentDelegationMapOutput {
	return i.ToAssessmentDelegationMapOutputWithContext(context.Background())
}

func (i AssessmentDelegationMap) ToAssessmentDelegationMapOutputWithContext(ctx context.Context) AssessmentDelegationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentDelegationMapOutput)
}

type AssessmentDelegationOutput struct{ *pulumi.OutputState }

func (AssessmentDelegationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentDelegation)(nil)).Elem()
}

func (o AssessmentDelegationOutput) ToAssessmentDelegationOutput() AssessmentDelegationOutput {
	return o
}

func (o AssessmentDelegationOutput) ToAssessmentDelegationOutputWithContext(ctx context.Context) AssessmentDelegationOutput {
	return o
}

// Identifier for the assessment.
func (o AssessmentDelegationOutput) AssessmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentDelegation) pulumi.StringOutput { return v.AssessmentId }).(pulumi.StringOutput)
}

// Comment describing the delegation request.
func (o AssessmentDelegationOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentDelegation) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
func (o AssessmentDelegationOutput) ControlSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentDelegation) pulumi.StringOutput { return v.ControlSetId }).(pulumi.StringOutput)
}

// Unique identifier for the delegation.
func (o AssessmentDelegationOutput) DelegationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentDelegation) pulumi.StringOutput { return v.DelegationId }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the IAM role.
func (o AssessmentDelegationOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentDelegation) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.
//
// The following arguments are optional:
func (o AssessmentDelegationOutput) RoleType() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentDelegation) pulumi.StringOutput { return v.RoleType }).(pulumi.StringOutput)
}

// Status of the delegation.
func (o AssessmentDelegationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentDelegation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AssessmentDelegationArrayOutput struct{ *pulumi.OutputState }

func (AssessmentDelegationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssessmentDelegation)(nil)).Elem()
}

func (o AssessmentDelegationArrayOutput) ToAssessmentDelegationArrayOutput() AssessmentDelegationArrayOutput {
	return o
}

func (o AssessmentDelegationArrayOutput) ToAssessmentDelegationArrayOutputWithContext(ctx context.Context) AssessmentDelegationArrayOutput {
	return o
}

func (o AssessmentDelegationArrayOutput) Index(i pulumi.IntInput) AssessmentDelegationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AssessmentDelegation {
		return vs[0].([]*AssessmentDelegation)[vs[1].(int)]
	}).(AssessmentDelegationOutput)
}

type AssessmentDelegationMapOutput struct{ *pulumi.OutputState }

func (AssessmentDelegationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssessmentDelegation)(nil)).Elem()
}

func (o AssessmentDelegationMapOutput) ToAssessmentDelegationMapOutput() AssessmentDelegationMapOutput {
	return o
}

func (o AssessmentDelegationMapOutput) ToAssessmentDelegationMapOutputWithContext(ctx context.Context) AssessmentDelegationMapOutput {
	return o
}

func (o AssessmentDelegationMapOutput) MapIndex(k pulumi.StringInput) AssessmentDelegationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AssessmentDelegation {
		return vs[0].(map[string]*AssessmentDelegation)[vs[1].(string)]
	}).(AssessmentDelegationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentDelegationInput)(nil)).Elem(), &AssessmentDelegation{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentDelegationArrayInput)(nil)).Elem(), AssessmentDelegationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentDelegationMapInput)(nil)).Elem(), AssessmentDelegationMap{})
	pulumi.RegisterOutputType(AssessmentDelegationOutput{})
	pulumi.RegisterOutputType(AssessmentDelegationArrayOutput{})
	pulumi.RegisterOutputType(AssessmentDelegationMapOutput{})
}
