// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Service Catalog Portfolio Share. Shares the specified portfolio with the specified account or organization node. You can share portfolios to an organization, an organizational unit, or a specific account.
//
// If the portfolio share with the specified account or organization node already exists, using this resource to re-create the share will have no effect and will not return an error. You can then use this resource to update the share.
//
// > **NOTE:** Shares to an organization node can only be created by the management account of an organization or by a delegated administrator. If a delegated admin is de-registered, they can no longer create portfolio shares.
//
// > **NOTE:** AWSOrganizationsAccess must be enabled in order to create a portfolio share to an organization node.
//
// > **NOTE:** You can't share a shared resource, including portfolios that contain a shared product.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicecatalog"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicecatalog.NewPortfolioShare(ctx, "example", &servicecatalog.PortfolioShareArgs{
// 			PrincipalId: pulumi.String("012128675309"),
// 			PortfolioId: pulumi.Any(aws_servicecatalog_portfolio.Example.Id),
// 			Type:        pulumi.String("ACCOUNT"),
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
// `aws_servicecatalog_portfolio_share` can be imported using the portfolio share ID, e.g.
//
// ```sh
//  $ pulumi import aws:servicecatalog/portfolioShare:PortfolioShare example port-12344321:ACCOUNT:123456789012
// ```
type PortfolioShare struct {
	pulumi.CustomResourceState

	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// Whether the shared portfolio is imported by the recipient account. If the recipient is organizational, the share is automatically imported, and the field is always set to true.
	Accepted pulumi.BoolOutput `pulumi:"accepted"`
	// Portfolio identifier.
	PortfolioId pulumi.StringOutput `pulumi:"portfolioId"`
	// Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// Whether to enable sharing of `servicecatalog.TagOption` resources when creating the portfolio share.
	ShareTagOptions pulumi.BoolPtrOutput `pulumi:"shareTagOptions"`
	// Type of portfolio share. Valid values are `ACCOUNT` (an external account), `ORGANIZATION` (a share to every account in an organization), `ORGANIZATIONAL_UNIT`, `ORGANIZATION_MEMBER_ACCOUNT` (a share to an account in an organization).
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.
	WaitForAcceptance pulumi.BoolPtrOutput `pulumi:"waitForAcceptance"`
}

// NewPortfolioShare registers a new resource with the given unique name, arguments, and options.
func NewPortfolioShare(ctx *pulumi.Context,
	name string, args *PortfolioShareArgs, opts ...pulumi.ResourceOption) (*PortfolioShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource PortfolioShare
	err := ctx.RegisterResource("aws:servicecatalog/portfolioShare:PortfolioShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortfolioShare gets an existing PortfolioShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortfolioShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortfolioShareState, opts ...pulumi.ResourceOption) (*PortfolioShare, error) {
	var resource PortfolioShare
	err := ctx.ReadResource("aws:servicecatalog/portfolioShare:PortfolioShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortfolioShare resources.
type portfolioShareState struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// Whether the shared portfolio is imported by the recipient account. If the recipient is organizational, the share is automatically imported, and the field is always set to true.
	Accepted *bool `pulumi:"accepted"`
	// Portfolio identifier.
	PortfolioId *string `pulumi:"portfolioId"`
	// Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
	PrincipalId *string `pulumi:"principalId"`
	// Whether to enable sharing of `servicecatalog.TagOption` resources when creating the portfolio share.
	ShareTagOptions *bool `pulumi:"shareTagOptions"`
	// Type of portfolio share. Valid values are `ACCOUNT` (an external account), `ORGANIZATION` (a share to every account in an organization), `ORGANIZATIONAL_UNIT`, `ORGANIZATION_MEMBER_ACCOUNT` (a share to an account in an organization).
	Type *string `pulumi:"type"`
	// Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.
	WaitForAcceptance *bool `pulumi:"waitForAcceptance"`
}

type PortfolioShareState struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrInput
	// Whether the shared portfolio is imported by the recipient account. If the recipient is organizational, the share is automatically imported, and the field is always set to true.
	Accepted pulumi.BoolPtrInput
	// Portfolio identifier.
	PortfolioId pulumi.StringPtrInput
	// Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
	PrincipalId pulumi.StringPtrInput
	// Whether to enable sharing of `servicecatalog.TagOption` resources when creating the portfolio share.
	ShareTagOptions pulumi.BoolPtrInput
	// Type of portfolio share. Valid values are `ACCOUNT` (an external account), `ORGANIZATION` (a share to every account in an organization), `ORGANIZATIONAL_UNIT`, `ORGANIZATION_MEMBER_ACCOUNT` (a share to an account in an organization).
	Type pulumi.StringPtrInput
	// Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.
	WaitForAcceptance pulumi.BoolPtrInput
}

func (PortfolioShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioShareState)(nil)).Elem()
}

type portfolioShareArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// Portfolio identifier.
	PortfolioId string `pulumi:"portfolioId"`
	// Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
	PrincipalId string `pulumi:"principalId"`
	// Whether to enable sharing of `servicecatalog.TagOption` resources when creating the portfolio share.
	ShareTagOptions *bool `pulumi:"shareTagOptions"`
	// Type of portfolio share. Valid values are `ACCOUNT` (an external account), `ORGANIZATION` (a share to every account in an organization), `ORGANIZATIONAL_UNIT`, `ORGANIZATION_MEMBER_ACCOUNT` (a share to an account in an organization).
	Type string `pulumi:"type"`
	// Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.
	WaitForAcceptance *bool `pulumi:"waitForAcceptance"`
}

// The set of arguments for constructing a PortfolioShare resource.
type PortfolioShareArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrInput
	// Portfolio identifier.
	PortfolioId pulumi.StringInput
	// Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
	PrincipalId pulumi.StringInput
	// Whether to enable sharing of `servicecatalog.TagOption` resources when creating the portfolio share.
	ShareTagOptions pulumi.BoolPtrInput
	// Type of portfolio share. Valid values are `ACCOUNT` (an external account), `ORGANIZATION` (a share to every account in an organization), `ORGANIZATIONAL_UNIT`, `ORGANIZATION_MEMBER_ACCOUNT` (a share to an account in an organization).
	Type pulumi.StringInput
	// Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.
	WaitForAcceptance pulumi.BoolPtrInput
}

func (PortfolioShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioShareArgs)(nil)).Elem()
}

type PortfolioShareInput interface {
	pulumi.Input

	ToPortfolioShareOutput() PortfolioShareOutput
	ToPortfolioShareOutputWithContext(ctx context.Context) PortfolioShareOutput
}

func (*PortfolioShare) ElementType() reflect.Type {
	return reflect.TypeOf((*PortfolioShare)(nil))
}

func (i *PortfolioShare) ToPortfolioShareOutput() PortfolioShareOutput {
	return i.ToPortfolioShareOutputWithContext(context.Background())
}

func (i *PortfolioShare) ToPortfolioShareOutputWithContext(ctx context.Context) PortfolioShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioShareOutput)
}

func (i *PortfolioShare) ToPortfolioSharePtrOutput() PortfolioSharePtrOutput {
	return i.ToPortfolioSharePtrOutputWithContext(context.Background())
}

func (i *PortfolioShare) ToPortfolioSharePtrOutputWithContext(ctx context.Context) PortfolioSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioSharePtrOutput)
}

type PortfolioSharePtrInput interface {
	pulumi.Input

	ToPortfolioSharePtrOutput() PortfolioSharePtrOutput
	ToPortfolioSharePtrOutputWithContext(ctx context.Context) PortfolioSharePtrOutput
}

type portfolioSharePtrType PortfolioShareArgs

func (*portfolioSharePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PortfolioShare)(nil))
}

func (i *portfolioSharePtrType) ToPortfolioSharePtrOutput() PortfolioSharePtrOutput {
	return i.ToPortfolioSharePtrOutputWithContext(context.Background())
}

func (i *portfolioSharePtrType) ToPortfolioSharePtrOutputWithContext(ctx context.Context) PortfolioSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioSharePtrOutput)
}

// PortfolioShareArrayInput is an input type that accepts PortfolioShareArray and PortfolioShareArrayOutput values.
// You can construct a concrete instance of `PortfolioShareArrayInput` via:
//
//          PortfolioShareArray{ PortfolioShareArgs{...} }
type PortfolioShareArrayInput interface {
	pulumi.Input

	ToPortfolioShareArrayOutput() PortfolioShareArrayOutput
	ToPortfolioShareArrayOutputWithContext(context.Context) PortfolioShareArrayOutput
}

type PortfolioShareArray []PortfolioShareInput

func (PortfolioShareArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortfolioShare)(nil)).Elem()
}

func (i PortfolioShareArray) ToPortfolioShareArrayOutput() PortfolioShareArrayOutput {
	return i.ToPortfolioShareArrayOutputWithContext(context.Background())
}

func (i PortfolioShareArray) ToPortfolioShareArrayOutputWithContext(ctx context.Context) PortfolioShareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioShareArrayOutput)
}

// PortfolioShareMapInput is an input type that accepts PortfolioShareMap and PortfolioShareMapOutput values.
// You can construct a concrete instance of `PortfolioShareMapInput` via:
//
//          PortfolioShareMap{ "key": PortfolioShareArgs{...} }
type PortfolioShareMapInput interface {
	pulumi.Input

	ToPortfolioShareMapOutput() PortfolioShareMapOutput
	ToPortfolioShareMapOutputWithContext(context.Context) PortfolioShareMapOutput
}

type PortfolioShareMap map[string]PortfolioShareInput

func (PortfolioShareMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortfolioShare)(nil)).Elem()
}

func (i PortfolioShareMap) ToPortfolioShareMapOutput() PortfolioShareMapOutput {
	return i.ToPortfolioShareMapOutputWithContext(context.Background())
}

func (i PortfolioShareMap) ToPortfolioShareMapOutputWithContext(ctx context.Context) PortfolioShareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioShareMapOutput)
}

type PortfolioShareOutput struct{ *pulumi.OutputState }

func (PortfolioShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortfolioShare)(nil))
}

func (o PortfolioShareOutput) ToPortfolioShareOutput() PortfolioShareOutput {
	return o
}

func (o PortfolioShareOutput) ToPortfolioShareOutputWithContext(ctx context.Context) PortfolioShareOutput {
	return o
}

func (o PortfolioShareOutput) ToPortfolioSharePtrOutput() PortfolioSharePtrOutput {
	return o.ToPortfolioSharePtrOutputWithContext(context.Background())
}

func (o PortfolioShareOutput) ToPortfolioSharePtrOutputWithContext(ctx context.Context) PortfolioSharePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PortfolioShare) *PortfolioShare {
		return &v
	}).(PortfolioSharePtrOutput)
}

type PortfolioSharePtrOutput struct{ *pulumi.OutputState }

func (PortfolioSharePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortfolioShare)(nil))
}

func (o PortfolioSharePtrOutput) ToPortfolioSharePtrOutput() PortfolioSharePtrOutput {
	return o
}

func (o PortfolioSharePtrOutput) ToPortfolioSharePtrOutputWithContext(ctx context.Context) PortfolioSharePtrOutput {
	return o
}

func (o PortfolioSharePtrOutput) Elem() PortfolioShareOutput {
	return o.ApplyT(func(v *PortfolioShare) PortfolioShare {
		if v != nil {
			return *v
		}
		var ret PortfolioShare
		return ret
	}).(PortfolioShareOutput)
}

type PortfolioShareArrayOutput struct{ *pulumi.OutputState }

func (PortfolioShareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PortfolioShare)(nil))
}

func (o PortfolioShareArrayOutput) ToPortfolioShareArrayOutput() PortfolioShareArrayOutput {
	return o
}

func (o PortfolioShareArrayOutput) ToPortfolioShareArrayOutputWithContext(ctx context.Context) PortfolioShareArrayOutput {
	return o
}

func (o PortfolioShareArrayOutput) Index(i pulumi.IntInput) PortfolioShareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PortfolioShare {
		return vs[0].([]PortfolioShare)[vs[1].(int)]
	}).(PortfolioShareOutput)
}

type PortfolioShareMapOutput struct{ *pulumi.OutputState }

func (PortfolioShareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PortfolioShare)(nil))
}

func (o PortfolioShareMapOutput) ToPortfolioShareMapOutput() PortfolioShareMapOutput {
	return o
}

func (o PortfolioShareMapOutput) ToPortfolioShareMapOutputWithContext(ctx context.Context) PortfolioShareMapOutput {
	return o
}

func (o PortfolioShareMapOutput) MapIndex(k pulumi.StringInput) PortfolioShareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PortfolioShare {
		return vs[0].(map[string]PortfolioShare)[vs[1].(string)]
	}).(PortfolioShareOutput)
}

func init() {
	pulumi.RegisterOutputType(PortfolioShareOutput{})
	pulumi.RegisterOutputType(PortfolioSharePtrOutput{})
	pulumi.RegisterOutputType(PortfolioShareArrayOutput{})
	pulumi.RegisterOutputType(PortfolioShareMapOutput{})
}