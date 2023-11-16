// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Resource Access Manager (RAM) principal association. Depending if [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), the RAM behavior with different principal types changes.
//
// When RAM Sharing with AWS Organizations is enabled:
//
// - For AWS Account ID, Organization, and Organizational Unit principals within the same AWS Organization, no resource share invitation is sent and resources become available automatically after creating the association.
// - For AWS Account ID principals outside the AWS Organization, a resource share invitation is sent and must be accepted before resources become available. See the `ram.ResourceShareAccepter` resource to accept these invitations.
//
// When RAM Sharing with AWS Organizations is not enabled:
//
// - Organization and Organizational Unit principals cannot be used.
// - For AWS Account ID principals, a resource share invitation is sent and must be accepted before resources become available. See the `ram.ResourceShareAccepter` resource to accept these invitations.
//
// ## Example Usage
// ### AWS Account ID
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleResourceShare, err := ram.NewResourceShare(ctx, "exampleResourceShare", &ram.ResourceShareArgs{
//				AllowExternalPrincipals: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewPrincipalAssociation(ctx, "examplePrincipalAssociation", &ram.PrincipalAssociationArgs{
//				Principal:        pulumi.String("111111111111"),
//				ResourceShareArn: exampleResourceShare.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### AWS Organization
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ram.NewPrincipalAssociation(ctx, "example", &ram.PrincipalAssociationArgs{
//				Principal:        pulumi.Any(aws_organizations_organization.Example.Arn),
//				ResourceShareArn: pulumi.Any(aws_ram_resource_share.Example.Arn),
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
// Using `pulumi import`, import RAM Principal Associations using their Resource Share ARN and the `principal` separated by a comma. For example:
//
// ```sh
//
//	$ pulumi import aws:ram/principalAssociation:PrincipalAssociation example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,123456789012
//
// ```
type PrincipalAssociation struct {
	pulumi.CustomResourceState

	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn pulumi.StringOutput `pulumi:"resourceShareArn"`
}

// NewPrincipalAssociation registers a new resource with the given unique name, arguments, and options.
func NewPrincipalAssociation(ctx *pulumi.Context,
	name string, args *PrincipalAssociationArgs, opts ...pulumi.ResourceOption) (*PrincipalAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.ResourceShareArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceShareArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrincipalAssociation
	err := ctx.RegisterResource("aws:ram/principalAssociation:PrincipalAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrincipalAssociation gets an existing PrincipalAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrincipalAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrincipalAssociationState, opts ...pulumi.ResourceOption) (*PrincipalAssociation, error) {
	var resource PrincipalAssociation
	err := ctx.ReadResource("aws:ram/principalAssociation:PrincipalAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrincipalAssociation resources.
type principalAssociationState struct {
	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal *string `pulumi:"principal"`
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string `pulumi:"resourceShareArn"`
}

type PrincipalAssociationState struct {
	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn pulumi.StringPtrInput
}

func (PrincipalAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*principalAssociationState)(nil)).Elem()
}

type principalAssociationArgs struct {
	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal string `pulumi:"principal"`
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn string `pulumi:"resourceShareArn"`
}

// The set of arguments for constructing a PrincipalAssociation resource.
type PrincipalAssociationArgs struct {
	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal pulumi.StringInput
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn pulumi.StringInput
}

func (PrincipalAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*principalAssociationArgs)(nil)).Elem()
}

type PrincipalAssociationInput interface {
	pulumi.Input

	ToPrincipalAssociationOutput() PrincipalAssociationOutput
	ToPrincipalAssociationOutputWithContext(ctx context.Context) PrincipalAssociationOutput
}

func (*PrincipalAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalAssociation)(nil)).Elem()
}

func (i *PrincipalAssociation) ToPrincipalAssociationOutput() PrincipalAssociationOutput {
	return i.ToPrincipalAssociationOutputWithContext(context.Background())
}

func (i *PrincipalAssociation) ToPrincipalAssociationOutputWithContext(ctx context.Context) PrincipalAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalAssociationOutput)
}

// PrincipalAssociationArrayInput is an input type that accepts PrincipalAssociationArray and PrincipalAssociationArrayOutput values.
// You can construct a concrete instance of `PrincipalAssociationArrayInput` via:
//
//	PrincipalAssociationArray{ PrincipalAssociationArgs{...} }
type PrincipalAssociationArrayInput interface {
	pulumi.Input

	ToPrincipalAssociationArrayOutput() PrincipalAssociationArrayOutput
	ToPrincipalAssociationArrayOutputWithContext(context.Context) PrincipalAssociationArrayOutput
}

type PrincipalAssociationArray []PrincipalAssociationInput

func (PrincipalAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrincipalAssociation)(nil)).Elem()
}

func (i PrincipalAssociationArray) ToPrincipalAssociationArrayOutput() PrincipalAssociationArrayOutput {
	return i.ToPrincipalAssociationArrayOutputWithContext(context.Background())
}

func (i PrincipalAssociationArray) ToPrincipalAssociationArrayOutputWithContext(ctx context.Context) PrincipalAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalAssociationArrayOutput)
}

// PrincipalAssociationMapInput is an input type that accepts PrincipalAssociationMap and PrincipalAssociationMapOutput values.
// You can construct a concrete instance of `PrincipalAssociationMapInput` via:
//
//	PrincipalAssociationMap{ "key": PrincipalAssociationArgs{...} }
type PrincipalAssociationMapInput interface {
	pulumi.Input

	ToPrincipalAssociationMapOutput() PrincipalAssociationMapOutput
	ToPrincipalAssociationMapOutputWithContext(context.Context) PrincipalAssociationMapOutput
}

type PrincipalAssociationMap map[string]PrincipalAssociationInput

func (PrincipalAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrincipalAssociation)(nil)).Elem()
}

func (i PrincipalAssociationMap) ToPrincipalAssociationMapOutput() PrincipalAssociationMapOutput {
	return i.ToPrincipalAssociationMapOutputWithContext(context.Background())
}

func (i PrincipalAssociationMap) ToPrincipalAssociationMapOutputWithContext(ctx context.Context) PrincipalAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalAssociationMapOutput)
}

type PrincipalAssociationOutput struct{ *pulumi.OutputState }

func (PrincipalAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalAssociation)(nil)).Elem()
}

func (o PrincipalAssociationOutput) ToPrincipalAssociationOutput() PrincipalAssociationOutput {
	return o
}

func (o PrincipalAssociationOutput) ToPrincipalAssociationOutputWithContext(ctx context.Context) PrincipalAssociationOutput {
	return o
}

// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
func (o PrincipalAssociationOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *PrincipalAssociation) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the resource share.
func (o PrincipalAssociationOutput) ResourceShareArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PrincipalAssociation) pulumi.StringOutput { return v.ResourceShareArn }).(pulumi.StringOutput)
}

type PrincipalAssociationArrayOutput struct{ *pulumi.OutputState }

func (PrincipalAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrincipalAssociation)(nil)).Elem()
}

func (o PrincipalAssociationArrayOutput) ToPrincipalAssociationArrayOutput() PrincipalAssociationArrayOutput {
	return o
}

func (o PrincipalAssociationArrayOutput) ToPrincipalAssociationArrayOutputWithContext(ctx context.Context) PrincipalAssociationArrayOutput {
	return o
}

func (o PrincipalAssociationArrayOutput) Index(i pulumi.IntInput) PrincipalAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrincipalAssociation {
		return vs[0].([]*PrincipalAssociation)[vs[1].(int)]
	}).(PrincipalAssociationOutput)
}

type PrincipalAssociationMapOutput struct{ *pulumi.OutputState }

func (PrincipalAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrincipalAssociation)(nil)).Elem()
}

func (o PrincipalAssociationMapOutput) ToPrincipalAssociationMapOutput() PrincipalAssociationMapOutput {
	return o
}

func (o PrincipalAssociationMapOutput) ToPrincipalAssociationMapOutputWithContext(ctx context.Context) PrincipalAssociationMapOutput {
	return o
}

func (o PrincipalAssociationMapOutput) MapIndex(k pulumi.StringInput) PrincipalAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrincipalAssociation {
		return vs[0].(map[string]*PrincipalAssociation)[vs[1].(string)]
	}).(PrincipalAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalAssociationInput)(nil)).Elem(), &PrincipalAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalAssociationArrayInput)(nil)).Elem(), PrincipalAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalAssociationMapInput)(nil)).Elem(), PrincipalAssociationMap{})
	pulumi.RegisterOutputType(PrincipalAssociationOutput{})
	pulumi.RegisterOutputType(PrincipalAssociationArrayOutput{})
	pulumi.RegisterOutputType(PrincipalAssociationMapOutput{})
}
