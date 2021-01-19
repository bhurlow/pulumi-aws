// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssoadmin

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM managed policy for a Single Sign-On (SSO) Permission Set resource
//
// > **NOTE:** Creating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.
//
// ## Import
//
// SSO Managed Policy Attachments can be imported using the `managed_policy_arn`, `permission_set_arn`, and `instance_arn` separated by a comma (`,`) e.g.
//
// ```sh
//  $ pulumi import aws:ssoadmin/managedPolicyAttachment:ManagedPolicyAttachment example arn:aws:iam::aws:policy/AlexaForBusinessDeviceSetup,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
// ```
type ManagedPolicyAttachment struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn pulumi.StringOutput `pulumi:"managedPolicyArn"`
	// The name of the IAM Managed Policy.
	ManagedPolicyName pulumi.StringOutput `pulumi:"managedPolicyName"`
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn pulumi.StringOutput `pulumi:"permissionSetArn"`
}

// NewManagedPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewManagedPolicyAttachment(ctx *pulumi.Context,
	name string, args *ManagedPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*ManagedPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	if args.ManagedPolicyArn == nil {
		return nil, errors.New("invalid value for required argument 'ManagedPolicyArn'")
	}
	if args.PermissionSetArn == nil {
		return nil, errors.New("invalid value for required argument 'PermissionSetArn'")
	}
	var resource ManagedPolicyAttachment
	err := ctx.RegisterResource("aws:ssoadmin/managedPolicyAttachment:ManagedPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPolicyAttachment gets an existing ManagedPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPolicyAttachmentState, opts ...pulumi.ResourceOption) (*ManagedPolicyAttachment, error) {
	var resource ManagedPolicyAttachment
	err := ctx.ReadResource("aws:ssoadmin/managedPolicyAttachment:ManagedPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPolicyAttachment resources.
type managedPolicyAttachmentState struct {
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn *string `pulumi:"instanceArn"`
	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn *string `pulumi:"managedPolicyArn"`
	// The name of the IAM Managed Policy.
	ManagedPolicyName *string `pulumi:"managedPolicyName"`
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn *string `pulumi:"permissionSetArn"`
}

type ManagedPolicyAttachmentState struct {
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringPtrInput
	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn pulumi.StringPtrInput
	// The name of the IAM Managed Policy.
	ManagedPolicyName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn pulumi.StringPtrInput
}

func (ManagedPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPolicyAttachmentState)(nil)).Elem()
}

type managedPolicyAttachmentArgs struct {
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `pulumi:"instanceArn"`
	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn string `pulumi:"managedPolicyArn"`
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn string `pulumi:"permissionSetArn"`
}

// The set of arguments for constructing a ManagedPolicyAttachment resource.
type ManagedPolicyAttachmentArgs struct {
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringInput
	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn pulumi.StringInput
}

func (ManagedPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPolicyAttachmentArgs)(nil)).Elem()
}

type ManagedPolicyAttachmentInput interface {
	pulumi.Input

	ToManagedPolicyAttachmentOutput() ManagedPolicyAttachmentOutput
	ToManagedPolicyAttachmentOutputWithContext(ctx context.Context) ManagedPolicyAttachmentOutput
}

func (*ManagedPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPolicyAttachment)(nil))
}

func (i *ManagedPolicyAttachment) ToManagedPolicyAttachmentOutput() ManagedPolicyAttachmentOutput {
	return i.ToManagedPolicyAttachmentOutputWithContext(context.Background())
}

func (i *ManagedPolicyAttachment) ToManagedPolicyAttachmentOutputWithContext(ctx context.Context) ManagedPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPolicyAttachmentOutput)
}

type ManagedPolicyAttachmentOutput struct {
	*pulumi.OutputState
}

func (ManagedPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPolicyAttachment)(nil))
}

func (o ManagedPolicyAttachmentOutput) ToManagedPolicyAttachmentOutput() ManagedPolicyAttachmentOutput {
	return o
}

func (o ManagedPolicyAttachmentOutput) ToManagedPolicyAttachmentOutputWithContext(ctx context.Context) ManagedPolicyAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedPolicyAttachmentOutput{})
}