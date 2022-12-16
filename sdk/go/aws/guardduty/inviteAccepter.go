// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to accept a pending GuardDuty invite on creation, ensure the detector has the correct primary account on read, and disassociate with the primary account upon removal.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/guardduty"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.NewProvider(ctx, "primary", nil)
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewProvider(ctx, "member", nil)
//			if err != nil {
//				return err
//			}
//			primaryDetector, err := guardduty.NewDetector(ctx, "primaryDetector", nil, pulumi.Provider(aws.Primary))
//			if err != nil {
//				return err
//			}
//			memberDetector, err := guardduty.NewDetector(ctx, "memberDetector", nil, pulumi.Provider(aws.Member))
//			if err != nil {
//				return err
//			}
//			memberMember, err := guardduty.NewMember(ctx, "memberMember", &guardduty.MemberArgs{
//				AccountId:  memberDetector.AccountId,
//				DetectorId: primaryDetector.ID(),
//				Email:      pulumi.String("required@example.com"),
//				Invite:     pulumi.Bool(true),
//			}, pulumi.Provider(aws.Primary))
//			if err != nil {
//				return err
//			}
//			_, err = guardduty.NewInviteAccepter(ctx, "memberInviteAccepter", &guardduty.InviteAccepterArgs{
//				DetectorId:      memberDetector.ID(),
//				MasterAccountId: primaryDetector.AccountId,
//			}, pulumi.Provider(aws.Member), pulumi.DependsOn([]pulumi.Resource{
//				memberMember,
//			}))
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
// `aws_guardduty_invite_accepter` can be imported using the member GuardDuty detector ID, e.g.,
//
// ```sh
//
//	$ pulumi import aws:guardduty/inviteAccepter:InviteAccepter member 00b00fd5aecc0ab60a708659477e9617
//
// ```
type InviteAccepter struct {
	pulumi.CustomResourceState

	// The detector ID of the member GuardDuty account.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// AWS account ID for primary account.
	MasterAccountId pulumi.StringOutput `pulumi:"masterAccountId"`
}

// NewInviteAccepter registers a new resource with the given unique name, arguments, and options.
func NewInviteAccepter(ctx *pulumi.Context,
	name string, args *InviteAccepterArgs, opts ...pulumi.ResourceOption) (*InviteAccepter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.MasterAccountId == nil {
		return nil, errors.New("invalid value for required argument 'MasterAccountId'")
	}
	var resource InviteAccepter
	err := ctx.RegisterResource("aws:guardduty/inviteAccepter:InviteAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInviteAccepter gets an existing InviteAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInviteAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InviteAccepterState, opts ...pulumi.ResourceOption) (*InviteAccepter, error) {
	var resource InviteAccepter
	err := ctx.ReadResource("aws:guardduty/inviteAccepter:InviteAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InviteAccepter resources.
type inviteAccepterState struct {
	// The detector ID of the member GuardDuty account.
	DetectorId *string `pulumi:"detectorId"`
	// AWS account ID for primary account.
	MasterAccountId *string `pulumi:"masterAccountId"`
}

type InviteAccepterState struct {
	// The detector ID of the member GuardDuty account.
	DetectorId pulumi.StringPtrInput
	// AWS account ID for primary account.
	MasterAccountId pulumi.StringPtrInput
}

func (InviteAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*inviteAccepterState)(nil)).Elem()
}

type inviteAccepterArgs struct {
	// The detector ID of the member GuardDuty account.
	DetectorId string `pulumi:"detectorId"`
	// AWS account ID for primary account.
	MasterAccountId string `pulumi:"masterAccountId"`
}

// The set of arguments for constructing a InviteAccepter resource.
type InviteAccepterArgs struct {
	// The detector ID of the member GuardDuty account.
	DetectorId pulumi.StringInput
	// AWS account ID for primary account.
	MasterAccountId pulumi.StringInput
}

func (InviteAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inviteAccepterArgs)(nil)).Elem()
}

type InviteAccepterInput interface {
	pulumi.Input

	ToInviteAccepterOutput() InviteAccepterOutput
	ToInviteAccepterOutputWithContext(ctx context.Context) InviteAccepterOutput
}

func (*InviteAccepter) ElementType() reflect.Type {
	return reflect.TypeOf((**InviteAccepter)(nil)).Elem()
}

func (i *InviteAccepter) ToInviteAccepterOutput() InviteAccepterOutput {
	return i.ToInviteAccepterOutputWithContext(context.Background())
}

func (i *InviteAccepter) ToInviteAccepterOutputWithContext(ctx context.Context) InviteAccepterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterOutput)
}

// InviteAccepterArrayInput is an input type that accepts InviteAccepterArray and InviteAccepterArrayOutput values.
// You can construct a concrete instance of `InviteAccepterArrayInput` via:
//
//	InviteAccepterArray{ InviteAccepterArgs{...} }
type InviteAccepterArrayInput interface {
	pulumi.Input

	ToInviteAccepterArrayOutput() InviteAccepterArrayOutput
	ToInviteAccepterArrayOutputWithContext(context.Context) InviteAccepterArrayOutput
}

type InviteAccepterArray []InviteAccepterInput

func (InviteAccepterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InviteAccepter)(nil)).Elem()
}

func (i InviteAccepterArray) ToInviteAccepterArrayOutput() InviteAccepterArrayOutput {
	return i.ToInviteAccepterArrayOutputWithContext(context.Background())
}

func (i InviteAccepterArray) ToInviteAccepterArrayOutputWithContext(ctx context.Context) InviteAccepterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterArrayOutput)
}

// InviteAccepterMapInput is an input type that accepts InviteAccepterMap and InviteAccepterMapOutput values.
// You can construct a concrete instance of `InviteAccepterMapInput` via:
//
//	InviteAccepterMap{ "key": InviteAccepterArgs{...} }
type InviteAccepterMapInput interface {
	pulumi.Input

	ToInviteAccepterMapOutput() InviteAccepterMapOutput
	ToInviteAccepterMapOutputWithContext(context.Context) InviteAccepterMapOutput
}

type InviteAccepterMap map[string]InviteAccepterInput

func (InviteAccepterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InviteAccepter)(nil)).Elem()
}

func (i InviteAccepterMap) ToInviteAccepterMapOutput() InviteAccepterMapOutput {
	return i.ToInviteAccepterMapOutputWithContext(context.Background())
}

func (i InviteAccepterMap) ToInviteAccepterMapOutputWithContext(ctx context.Context) InviteAccepterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterMapOutput)
}

type InviteAccepterOutput struct{ *pulumi.OutputState }

func (InviteAccepterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InviteAccepter)(nil)).Elem()
}

func (o InviteAccepterOutput) ToInviteAccepterOutput() InviteAccepterOutput {
	return o
}

func (o InviteAccepterOutput) ToInviteAccepterOutputWithContext(ctx context.Context) InviteAccepterOutput {
	return o
}

// The detector ID of the member GuardDuty account.
func (o InviteAccepterOutput) DetectorId() pulumi.StringOutput {
	return o.ApplyT(func(v *InviteAccepter) pulumi.StringOutput { return v.DetectorId }).(pulumi.StringOutput)
}

// AWS account ID for primary account.
func (o InviteAccepterOutput) MasterAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *InviteAccepter) pulumi.StringOutput { return v.MasterAccountId }).(pulumi.StringOutput)
}

type InviteAccepterArrayOutput struct{ *pulumi.OutputState }

func (InviteAccepterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InviteAccepter)(nil)).Elem()
}

func (o InviteAccepterArrayOutput) ToInviteAccepterArrayOutput() InviteAccepterArrayOutput {
	return o
}

func (o InviteAccepterArrayOutput) ToInviteAccepterArrayOutputWithContext(ctx context.Context) InviteAccepterArrayOutput {
	return o
}

func (o InviteAccepterArrayOutput) Index(i pulumi.IntInput) InviteAccepterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InviteAccepter {
		return vs[0].([]*InviteAccepter)[vs[1].(int)]
	}).(InviteAccepterOutput)
}

type InviteAccepterMapOutput struct{ *pulumi.OutputState }

func (InviteAccepterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InviteAccepter)(nil)).Elem()
}

func (o InviteAccepterMapOutput) ToInviteAccepterMapOutput() InviteAccepterMapOutput {
	return o
}

func (o InviteAccepterMapOutput) ToInviteAccepterMapOutputWithContext(ctx context.Context) InviteAccepterMapOutput {
	return o
}

func (o InviteAccepterMapOutput) MapIndex(k pulumi.StringInput) InviteAccepterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InviteAccepter {
		return vs[0].(map[string]*InviteAccepter)[vs[1].(string)]
	}).(InviteAccepterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InviteAccepterInput)(nil)).Elem(), &InviteAccepter{})
	pulumi.RegisterInputType(reflect.TypeOf((*InviteAccepterArrayInput)(nil)).Elem(), InviteAccepterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InviteAccepterMapInput)(nil)).Elem(), InviteAccepterMap{})
	pulumi.RegisterOutputType(InviteAccepterOutput{})
	pulumi.RegisterOutputType(InviteAccepterArrayOutput{})
	pulumi.RegisterOutputType(InviteAccepterMapOutput{})
}
