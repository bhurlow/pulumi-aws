// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package macie2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a resource to manage an [Amazon Macie Member](https://docs.aws.amazon.com/macie/latest/APIReference/members-id.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleAccount, err := macie2.NewAccount(ctx, "exampleAccount", nil)
//			if err != nil {
//				return err
//			}
//			_, err = macie2.NewMember(ctx, "exampleMember", &macie2.MemberArgs{
//				AccountId:                          pulumi.String("AWS ACCOUNT ID"),
//				Email:                              pulumi.String("EMAIL"),
//				Invite:                             pulumi.Bool(true),
//				InvitationMessage:                  pulumi.String("Message of the invitation"),
//				InvitationDisableEmailNotification: pulumi.Bool(true),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleAccount,
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
// Using `pulumi import`, import `aws_macie2_member` using the account ID of the member account. For example:
//
// ```sh
//
//	$ pulumi import aws:macie2/member:Member example 123456789012
//
// ```
type Member struct {
	pulumi.CustomResourceState

	// The AWS account ID for the account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The AWS account ID for the administrator account.
	AdministratorAccountId pulumi.StringOutput `pulumi:"administratorAccountId"`
	// The Amazon Resource Name (ARN) of the account.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The email address for the account.
	Email pulumi.StringOutput `pulumi:"email"`
	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
	InvitationDisableEmailNotification pulumi.BoolPtrOutput `pulumi:"invitationDisableEmailNotification"`
	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage pulumi.StringPtrOutput `pulumi:"invitationMessage"`
	// Send an invitation to a member
	Invite pulumi.BoolOutput `pulumi:"invite"`
	// The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
	InvitedAt       pulumi.StringOutput `pulumi:"invitedAt"`
	MasterAccountId pulumi.StringOutput `pulumi:"masterAccountId"`
	// The current status of the relationship between the account and the administrator account.
	RelationshipStatus pulumi.StringOutput `pulumi:"relationshipStatus"`
	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
	Status pulumi.StringOutput `pulumi:"status"`
	// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewMember registers a new resource with the given unique name, arguments, and options.
func NewMember(ctx *pulumi.Context,
	name string, args *MemberArgs, opts ...pulumi.ResourceOption) (*Member, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Member
	err := ctx.RegisterResource("aws:macie2/member:Member", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMember gets an existing Member resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberState, opts ...pulumi.ResourceOption) (*Member, error) {
	var resource Member
	err := ctx.ReadResource("aws:macie2/member:Member", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Member resources.
type memberState struct {
	// The AWS account ID for the account.
	AccountId *string `pulumi:"accountId"`
	// The AWS account ID for the administrator account.
	AdministratorAccountId *string `pulumi:"administratorAccountId"`
	// The Amazon Resource Name (ARN) of the account.
	Arn *string `pulumi:"arn"`
	// The email address for the account.
	Email *string `pulumi:"email"`
	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
	InvitationDisableEmailNotification *bool `pulumi:"invitationDisableEmailNotification"`
	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage *string `pulumi:"invitationMessage"`
	// Send an invitation to a member
	Invite *bool `pulumi:"invite"`
	// The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
	InvitedAt       *string `pulumi:"invitedAt"`
	MasterAccountId *string `pulumi:"masterAccountId"`
	// The current status of the relationship between the account and the administrator account.
	RelationshipStatus *string `pulumi:"relationshipStatus"`
	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
	Status *string `pulumi:"status"`
	// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type MemberState struct {
	// The AWS account ID for the account.
	AccountId pulumi.StringPtrInput
	// The AWS account ID for the administrator account.
	AdministratorAccountId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the account.
	Arn pulumi.StringPtrInput
	// The email address for the account.
	Email pulumi.StringPtrInput
	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
	InvitationDisableEmailNotification pulumi.BoolPtrInput
	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage pulumi.StringPtrInput
	// Send an invitation to a member
	Invite pulumi.BoolPtrInput
	// The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
	InvitedAt       pulumi.StringPtrInput
	MasterAccountId pulumi.StringPtrInput
	// The current status of the relationship between the account and the administrator account.
	RelationshipStatus pulumi.StringPtrInput
	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
	Status pulumi.StringPtrInput
	// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
	UpdatedAt pulumi.StringPtrInput
}

func (MemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*memberState)(nil)).Elem()
}

type memberArgs struct {
	// The AWS account ID for the account.
	AccountId string `pulumi:"accountId"`
	// The email address for the account.
	Email string `pulumi:"email"`
	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
	InvitationDisableEmailNotification *bool `pulumi:"invitationDisableEmailNotification"`
	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage *string `pulumi:"invitationMessage"`
	// Send an invitation to a member
	Invite *bool `pulumi:"invite"`
	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
	Status *string `pulumi:"status"`
	// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Member resource.
type MemberArgs struct {
	// The AWS account ID for the account.
	AccountId pulumi.StringInput
	// The email address for the account.
	Email pulumi.StringInput
	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
	InvitationDisableEmailNotification pulumi.BoolPtrInput
	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage pulumi.StringPtrInput
	// Send an invitation to a member
	Invite pulumi.BoolPtrInput
	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
	Status pulumi.StringPtrInput
	// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
	Tags pulumi.StringMapInput
}

func (MemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*memberArgs)(nil)).Elem()
}

type MemberInput interface {
	pulumi.Input

	ToMemberOutput() MemberOutput
	ToMemberOutputWithContext(ctx context.Context) MemberOutput
}

func (*Member) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (i *Member) ToMemberOutput() MemberOutput {
	return i.ToMemberOutputWithContext(context.Background())
}

func (i *Member) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberOutput)
}

func (i *Member) ToOutput(ctx context.Context) pulumix.Output[*Member] {
	return pulumix.Output[*Member]{
		OutputState: i.ToMemberOutputWithContext(ctx).OutputState,
	}
}

// MemberArrayInput is an input type that accepts MemberArray and MemberArrayOutput values.
// You can construct a concrete instance of `MemberArrayInput` via:
//
//	MemberArray{ MemberArgs{...} }
type MemberArrayInput interface {
	pulumi.Input

	ToMemberArrayOutput() MemberArrayOutput
	ToMemberArrayOutputWithContext(context.Context) MemberArrayOutput
}

type MemberArray []MemberInput

func (MemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (i MemberArray) ToMemberArrayOutput() MemberArrayOutput {
	return i.ToMemberArrayOutputWithContext(context.Background())
}

func (i MemberArray) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberArrayOutput)
}

func (i MemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*Member] {
	return pulumix.Output[[]*Member]{
		OutputState: i.ToMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// MemberMapInput is an input type that accepts MemberMap and MemberMapOutput values.
// You can construct a concrete instance of `MemberMapInput` via:
//
//	MemberMap{ "key": MemberArgs{...} }
type MemberMapInput interface {
	pulumi.Input

	ToMemberMapOutput() MemberMapOutput
	ToMemberMapOutputWithContext(context.Context) MemberMapOutput
}

type MemberMap map[string]MemberInput

func (MemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (i MemberMap) ToMemberMapOutput() MemberMapOutput {
	return i.ToMemberMapOutputWithContext(context.Background())
}

func (i MemberMap) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberMapOutput)
}

func (i MemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Member] {
	return pulumix.Output[map[string]*Member]{
		OutputState: i.ToMemberMapOutputWithContext(ctx).OutputState,
	}
}

type MemberOutput struct{ *pulumi.OutputState }

func (MemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (o MemberOutput) ToMemberOutput() MemberOutput {
	return o
}

func (o MemberOutput) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return o
}

func (o MemberOutput) ToOutput(ctx context.Context) pulumix.Output[*Member] {
	return pulumix.Output[*Member]{
		OutputState: o.OutputState,
	}
}

// The AWS account ID for the account.
func (o MemberOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The AWS account ID for the administrator account.
func (o MemberOutput) AdministratorAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.AdministratorAccountId }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the account.
func (o MemberOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The email address for the account.
func (o MemberOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
func (o MemberOutput) InvitationDisableEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Member) pulumi.BoolPtrOutput { return v.InvitationDisableEmailNotification }).(pulumi.BoolPtrOutput)
}

// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
func (o MemberOutput) InvitationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Member) pulumi.StringPtrOutput { return v.InvitationMessage }).(pulumi.StringPtrOutput)
}

// Send an invitation to a member
func (o MemberOutput) Invite() pulumi.BoolOutput {
	return o.ApplyT(func(v *Member) pulumi.BoolOutput { return v.Invite }).(pulumi.BoolOutput)
}

// The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
func (o MemberOutput) InvitedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.InvitedAt }).(pulumi.StringOutput)
}

func (o MemberOutput) MasterAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.MasterAccountId }).(pulumi.StringOutput)
}

// The current status of the relationship between the account and the administrator account.
func (o MemberOutput) RelationshipStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.RelationshipStatus }).(pulumi.StringOutput)
}

// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
func (o MemberOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
func (o MemberOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Member) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o MemberOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Member) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
func (o MemberOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type MemberArrayOutput struct{ *pulumi.OutputState }

func (MemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (o MemberArrayOutput) ToMemberArrayOutput() MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Member] {
	return pulumix.Output[[]*Member]{
		OutputState: o.OutputState,
	}
}

func (o MemberArrayOutput) Index(i pulumi.IntInput) MemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Member {
		return vs[0].([]*Member)[vs[1].(int)]
	}).(MemberOutput)
}

type MemberMapOutput struct{ *pulumi.OutputState }

func (MemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (o MemberMapOutput) ToMemberMapOutput() MemberMapOutput {
	return o
}

func (o MemberMapOutput) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return o
}

func (o MemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Member] {
	return pulumix.Output[map[string]*Member]{
		OutputState: o.OutputState,
	}
}

func (o MemberMapOutput) MapIndex(k pulumi.StringInput) MemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Member {
		return vs[0].(map[string]*Member)[vs[1].(string)]
	}).(MemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MemberInput)(nil)).Elem(), &Member{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberArrayInput)(nil)).Elem(), MemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberMapInput)(nil)).Elem(), MemberMap{})
	pulumi.RegisterOutputType(MemberOutput{})
	pulumi.RegisterOutputType(MemberArrayOutput{})
	pulumi.RegisterOutputType(MemberMapOutput{})
}
