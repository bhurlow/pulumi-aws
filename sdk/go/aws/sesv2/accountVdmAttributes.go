// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sesv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS SESv2 (Simple Email V2) Account VDM Attributes.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewAccountVdmAttributes(ctx, "example", &sesv2.AccountVdmAttributesArgs{
//				DashboardAttributes: &sesv2.AccountVdmAttributesDashboardAttributesArgs{
//					EngagementMetrics: pulumi.String("ENABLED"),
//				},
//				GuardianAttributes: &sesv2.AccountVdmAttributesGuardianAttributesArgs{
//					OptimizedSharedDelivery: pulumi.String("ENABLED"),
//				},
//				VdmEnabled: pulumi.String("ENABLED"),
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
// Using `pulumi import`, import SESv2 (Simple Email V2) Account VDM Attributes using the word `ses-account-vdm-attributes`. For example:
//
// ```sh
//
//	$ pulumi import aws:sesv2/accountVdmAttributes:AccountVdmAttributes example ses-account-vdm-attributes
//
// ```
type AccountVdmAttributes struct {
	pulumi.CustomResourceState

	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardAttributes AccountVdmAttributesDashboardAttributesOutput `pulumi:"dashboardAttributes"`
	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianAttributes AccountVdmAttributesGuardianAttributesOutput `pulumi:"guardianAttributes"`
	// Specifies the status of your VDM configuration. Valid values: `ENABLED`, `DISABLED`.
	//
	// The following arguments are optional:
	VdmEnabled pulumi.StringOutput `pulumi:"vdmEnabled"`
}

// NewAccountVdmAttributes registers a new resource with the given unique name, arguments, and options.
func NewAccountVdmAttributes(ctx *pulumi.Context,
	name string, args *AccountVdmAttributesArgs, opts ...pulumi.ResourceOption) (*AccountVdmAttributes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VdmEnabled == nil {
		return nil, errors.New("invalid value for required argument 'VdmEnabled'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountVdmAttributes
	err := ctx.RegisterResource("aws:sesv2/accountVdmAttributes:AccountVdmAttributes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountVdmAttributes gets an existing AccountVdmAttributes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountVdmAttributes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountVdmAttributesState, opts ...pulumi.ResourceOption) (*AccountVdmAttributes, error) {
	var resource AccountVdmAttributes
	err := ctx.ReadResource("aws:sesv2/accountVdmAttributes:AccountVdmAttributes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountVdmAttributes resources.
type accountVdmAttributesState struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardAttributes *AccountVdmAttributesDashboardAttributes `pulumi:"dashboardAttributes"`
	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianAttributes *AccountVdmAttributesGuardianAttributes `pulumi:"guardianAttributes"`
	// Specifies the status of your VDM configuration. Valid values: `ENABLED`, `DISABLED`.
	//
	// The following arguments are optional:
	VdmEnabled *string `pulumi:"vdmEnabled"`
}

type AccountVdmAttributesState struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardAttributes AccountVdmAttributesDashboardAttributesPtrInput
	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianAttributes AccountVdmAttributesGuardianAttributesPtrInput
	// Specifies the status of your VDM configuration. Valid values: `ENABLED`, `DISABLED`.
	//
	// The following arguments are optional:
	VdmEnabled pulumi.StringPtrInput
}

func (AccountVdmAttributesState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountVdmAttributesState)(nil)).Elem()
}

type accountVdmAttributesArgs struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardAttributes *AccountVdmAttributesDashboardAttributes `pulumi:"dashboardAttributes"`
	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianAttributes *AccountVdmAttributesGuardianAttributes `pulumi:"guardianAttributes"`
	// Specifies the status of your VDM configuration. Valid values: `ENABLED`, `DISABLED`.
	//
	// The following arguments are optional:
	VdmEnabled string `pulumi:"vdmEnabled"`
}

// The set of arguments for constructing a AccountVdmAttributes resource.
type AccountVdmAttributesArgs struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardAttributes AccountVdmAttributesDashboardAttributesPtrInput
	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianAttributes AccountVdmAttributesGuardianAttributesPtrInput
	// Specifies the status of your VDM configuration. Valid values: `ENABLED`, `DISABLED`.
	//
	// The following arguments are optional:
	VdmEnabled pulumi.StringInput
}

func (AccountVdmAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountVdmAttributesArgs)(nil)).Elem()
}

type AccountVdmAttributesInput interface {
	pulumi.Input

	ToAccountVdmAttributesOutput() AccountVdmAttributesOutput
	ToAccountVdmAttributesOutputWithContext(ctx context.Context) AccountVdmAttributesOutput
}

func (*AccountVdmAttributes) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountVdmAttributes)(nil)).Elem()
}

func (i *AccountVdmAttributes) ToAccountVdmAttributesOutput() AccountVdmAttributesOutput {
	return i.ToAccountVdmAttributesOutputWithContext(context.Background())
}

func (i *AccountVdmAttributes) ToAccountVdmAttributesOutputWithContext(ctx context.Context) AccountVdmAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountVdmAttributesOutput)
}

// AccountVdmAttributesArrayInput is an input type that accepts AccountVdmAttributesArray and AccountVdmAttributesArrayOutput values.
// You can construct a concrete instance of `AccountVdmAttributesArrayInput` via:
//
//	AccountVdmAttributesArray{ AccountVdmAttributesArgs{...} }
type AccountVdmAttributesArrayInput interface {
	pulumi.Input

	ToAccountVdmAttributesArrayOutput() AccountVdmAttributesArrayOutput
	ToAccountVdmAttributesArrayOutputWithContext(context.Context) AccountVdmAttributesArrayOutput
}

type AccountVdmAttributesArray []AccountVdmAttributesInput

func (AccountVdmAttributesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountVdmAttributes)(nil)).Elem()
}

func (i AccountVdmAttributesArray) ToAccountVdmAttributesArrayOutput() AccountVdmAttributesArrayOutput {
	return i.ToAccountVdmAttributesArrayOutputWithContext(context.Background())
}

func (i AccountVdmAttributesArray) ToAccountVdmAttributesArrayOutputWithContext(ctx context.Context) AccountVdmAttributesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountVdmAttributesArrayOutput)
}

// AccountVdmAttributesMapInput is an input type that accepts AccountVdmAttributesMap and AccountVdmAttributesMapOutput values.
// You can construct a concrete instance of `AccountVdmAttributesMapInput` via:
//
//	AccountVdmAttributesMap{ "key": AccountVdmAttributesArgs{...} }
type AccountVdmAttributesMapInput interface {
	pulumi.Input

	ToAccountVdmAttributesMapOutput() AccountVdmAttributesMapOutput
	ToAccountVdmAttributesMapOutputWithContext(context.Context) AccountVdmAttributesMapOutput
}

type AccountVdmAttributesMap map[string]AccountVdmAttributesInput

func (AccountVdmAttributesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountVdmAttributes)(nil)).Elem()
}

func (i AccountVdmAttributesMap) ToAccountVdmAttributesMapOutput() AccountVdmAttributesMapOutput {
	return i.ToAccountVdmAttributesMapOutputWithContext(context.Background())
}

func (i AccountVdmAttributesMap) ToAccountVdmAttributesMapOutputWithContext(ctx context.Context) AccountVdmAttributesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountVdmAttributesMapOutput)
}

type AccountVdmAttributesOutput struct{ *pulumi.OutputState }

func (AccountVdmAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountVdmAttributes)(nil)).Elem()
}

func (o AccountVdmAttributesOutput) ToAccountVdmAttributesOutput() AccountVdmAttributesOutput {
	return o
}

func (o AccountVdmAttributesOutput) ToAccountVdmAttributesOutputWithContext(ctx context.Context) AccountVdmAttributesOutput {
	return o
}

// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
func (o AccountVdmAttributesOutput) DashboardAttributes() AccountVdmAttributesDashboardAttributesOutput {
	return o.ApplyT(func(v *AccountVdmAttributes) AccountVdmAttributesDashboardAttributesOutput {
		return v.DashboardAttributes
	}).(AccountVdmAttributesDashboardAttributesOutput)
}

// Specifies additional settings for your VDM configuration as applicable to the Guardian.
func (o AccountVdmAttributesOutput) GuardianAttributes() AccountVdmAttributesGuardianAttributesOutput {
	return o.ApplyT(func(v *AccountVdmAttributes) AccountVdmAttributesGuardianAttributesOutput {
		return v.GuardianAttributes
	}).(AccountVdmAttributesGuardianAttributesOutput)
}

// Specifies the status of your VDM configuration. Valid values: `ENABLED`, `DISABLED`.
//
// The following arguments are optional:
func (o AccountVdmAttributesOutput) VdmEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountVdmAttributes) pulumi.StringOutput { return v.VdmEnabled }).(pulumi.StringOutput)
}

type AccountVdmAttributesArrayOutput struct{ *pulumi.OutputState }

func (AccountVdmAttributesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountVdmAttributes)(nil)).Elem()
}

func (o AccountVdmAttributesArrayOutput) ToAccountVdmAttributesArrayOutput() AccountVdmAttributesArrayOutput {
	return o
}

func (o AccountVdmAttributesArrayOutput) ToAccountVdmAttributesArrayOutputWithContext(ctx context.Context) AccountVdmAttributesArrayOutput {
	return o
}

func (o AccountVdmAttributesArrayOutput) Index(i pulumi.IntInput) AccountVdmAttributesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountVdmAttributes {
		return vs[0].([]*AccountVdmAttributes)[vs[1].(int)]
	}).(AccountVdmAttributesOutput)
}

type AccountVdmAttributesMapOutput struct{ *pulumi.OutputState }

func (AccountVdmAttributesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountVdmAttributes)(nil)).Elem()
}

func (o AccountVdmAttributesMapOutput) ToAccountVdmAttributesMapOutput() AccountVdmAttributesMapOutput {
	return o
}

func (o AccountVdmAttributesMapOutput) ToAccountVdmAttributesMapOutputWithContext(ctx context.Context) AccountVdmAttributesMapOutput {
	return o
}

func (o AccountVdmAttributesMapOutput) MapIndex(k pulumi.StringInput) AccountVdmAttributesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountVdmAttributes {
		return vs[0].(map[string]*AccountVdmAttributes)[vs[1].(string)]
	}).(AccountVdmAttributesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountVdmAttributesInput)(nil)).Elem(), &AccountVdmAttributes{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountVdmAttributesArrayInput)(nil)).Elem(), AccountVdmAttributesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountVdmAttributesMapInput)(nil)).Elem(), AccountVdmAttributesMap{})
	pulumi.RegisterOutputType(AccountVdmAttributesOutput{})
	pulumi.RegisterOutputType(AccountVdmAttributesArrayOutput{})
	pulumi.RegisterOutputType(AccountVdmAttributesMapOutput{})
}
