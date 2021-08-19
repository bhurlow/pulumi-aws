// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a Service Catalog Portfolio.
//
// ## Example Usage
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
// 		_, err := servicecatalog.NewPortfolio(ctx, "portfolio", &servicecatalog.PortfolioArgs{
// 			Description:  pulumi.String("List of my organizations apps"),
// 			ProviderName: pulumi.String("Brett"),
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
// Service Catalog Portfolios can be imported using the `service catalog portfolio id`, e.g.
//
// ```sh
//  $ pulumi import aws:servicecatalog/portfolio:Portfolio testfolio port-12344321
// ```
type Portfolio struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput `pulumi:"arn"`
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Description of the portfolio
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the portfolio.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the person or organization who owns the portfolio.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
	// Tags to apply to the connection. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewPortfolio registers a new resource with the given unique name, arguments, and options.
func NewPortfolio(ctx *pulumi.Context,
	name string, args *PortfolioArgs, opts ...pulumi.ResourceOption) (*Portfolio, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	var resource Portfolio
	err := ctx.RegisterResource("aws:servicecatalog/portfolio:Portfolio", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortfolio gets an existing Portfolio resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortfolio(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortfolioState, opts ...pulumi.ResourceOption) (*Portfolio, error) {
	var resource Portfolio
	err := ctx.ReadResource("aws:servicecatalog/portfolio:Portfolio", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Portfolio resources.
type portfolioState struct {
	Arn         *string `pulumi:"arn"`
	CreatedTime *string `pulumi:"createdTime"`
	// Description of the portfolio
	Description *string `pulumi:"description"`
	// The name of the portfolio.
	Name *string `pulumi:"name"`
	// Name of the person or organization who owns the portfolio.
	ProviderName *string `pulumi:"providerName"`
	// Tags to apply to the connection. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type PortfolioState struct {
	Arn         pulumi.StringPtrInput
	CreatedTime pulumi.StringPtrInput
	// Description of the portfolio
	Description pulumi.StringPtrInput
	// The name of the portfolio.
	Name pulumi.StringPtrInput
	// Name of the person or organization who owns the portfolio.
	ProviderName pulumi.StringPtrInput
	// Tags to apply to the connection. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (PortfolioState) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioState)(nil)).Elem()
}

type portfolioArgs struct {
	// Description of the portfolio
	Description *string `pulumi:"description"`
	// The name of the portfolio.
	Name *string `pulumi:"name"`
	// Name of the person or organization who owns the portfolio.
	ProviderName string `pulumi:"providerName"`
	// Tags to apply to the connection. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Portfolio resource.
type PortfolioArgs struct {
	// Description of the portfolio
	Description pulumi.StringPtrInput
	// The name of the portfolio.
	Name pulumi.StringPtrInput
	// Name of the person or organization who owns the portfolio.
	ProviderName pulumi.StringInput
	// Tags to apply to the connection. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (PortfolioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioArgs)(nil)).Elem()
}

type PortfolioInput interface {
	pulumi.Input

	ToPortfolioOutput() PortfolioOutput
	ToPortfolioOutputWithContext(ctx context.Context) PortfolioOutput
}

func (*Portfolio) ElementType() reflect.Type {
	return reflect.TypeOf((*Portfolio)(nil))
}

func (i *Portfolio) ToPortfolioOutput() PortfolioOutput {
	return i.ToPortfolioOutputWithContext(context.Background())
}

func (i *Portfolio) ToPortfolioOutputWithContext(ctx context.Context) PortfolioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioOutput)
}

func (i *Portfolio) ToPortfolioPtrOutput() PortfolioPtrOutput {
	return i.ToPortfolioPtrOutputWithContext(context.Background())
}

func (i *Portfolio) ToPortfolioPtrOutputWithContext(ctx context.Context) PortfolioPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioPtrOutput)
}

type PortfolioPtrInput interface {
	pulumi.Input

	ToPortfolioPtrOutput() PortfolioPtrOutput
	ToPortfolioPtrOutputWithContext(ctx context.Context) PortfolioPtrOutput
}

type portfolioPtrType PortfolioArgs

func (*portfolioPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Portfolio)(nil))
}

func (i *portfolioPtrType) ToPortfolioPtrOutput() PortfolioPtrOutput {
	return i.ToPortfolioPtrOutputWithContext(context.Background())
}

func (i *portfolioPtrType) ToPortfolioPtrOutputWithContext(ctx context.Context) PortfolioPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioPtrOutput)
}

// PortfolioArrayInput is an input type that accepts PortfolioArray and PortfolioArrayOutput values.
// You can construct a concrete instance of `PortfolioArrayInput` via:
//
//          PortfolioArray{ PortfolioArgs{...} }
type PortfolioArrayInput interface {
	pulumi.Input

	ToPortfolioArrayOutput() PortfolioArrayOutput
	ToPortfolioArrayOutputWithContext(context.Context) PortfolioArrayOutput
}

type PortfolioArray []PortfolioInput

func (PortfolioArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Portfolio)(nil)).Elem()
}

func (i PortfolioArray) ToPortfolioArrayOutput() PortfolioArrayOutput {
	return i.ToPortfolioArrayOutputWithContext(context.Background())
}

func (i PortfolioArray) ToPortfolioArrayOutputWithContext(ctx context.Context) PortfolioArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioArrayOutput)
}

// PortfolioMapInput is an input type that accepts PortfolioMap and PortfolioMapOutput values.
// You can construct a concrete instance of `PortfolioMapInput` via:
//
//          PortfolioMap{ "key": PortfolioArgs{...} }
type PortfolioMapInput interface {
	pulumi.Input

	ToPortfolioMapOutput() PortfolioMapOutput
	ToPortfolioMapOutputWithContext(context.Context) PortfolioMapOutput
}

type PortfolioMap map[string]PortfolioInput

func (PortfolioMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Portfolio)(nil)).Elem()
}

func (i PortfolioMap) ToPortfolioMapOutput() PortfolioMapOutput {
	return i.ToPortfolioMapOutputWithContext(context.Background())
}

func (i PortfolioMap) ToPortfolioMapOutputWithContext(ctx context.Context) PortfolioMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioMapOutput)
}

type PortfolioOutput struct{ *pulumi.OutputState }

func (PortfolioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Portfolio)(nil))
}

func (o PortfolioOutput) ToPortfolioOutput() PortfolioOutput {
	return o
}

func (o PortfolioOutput) ToPortfolioOutputWithContext(ctx context.Context) PortfolioOutput {
	return o
}

func (o PortfolioOutput) ToPortfolioPtrOutput() PortfolioPtrOutput {
	return o.ToPortfolioPtrOutputWithContext(context.Background())
}

func (o PortfolioOutput) ToPortfolioPtrOutputWithContext(ctx context.Context) PortfolioPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Portfolio) *Portfolio {
		return &v
	}).(PortfolioPtrOutput)
}

type PortfolioPtrOutput struct{ *pulumi.OutputState }

func (PortfolioPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Portfolio)(nil))
}

func (o PortfolioPtrOutput) ToPortfolioPtrOutput() PortfolioPtrOutput {
	return o
}

func (o PortfolioPtrOutput) ToPortfolioPtrOutputWithContext(ctx context.Context) PortfolioPtrOutput {
	return o
}

func (o PortfolioPtrOutput) Elem() PortfolioOutput {
	return o.ApplyT(func(v *Portfolio) Portfolio {
		if v != nil {
			return *v
		}
		var ret Portfolio
		return ret
	}).(PortfolioOutput)
}

type PortfolioArrayOutput struct{ *pulumi.OutputState }

func (PortfolioArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Portfolio)(nil))
}

func (o PortfolioArrayOutput) ToPortfolioArrayOutput() PortfolioArrayOutput {
	return o
}

func (o PortfolioArrayOutput) ToPortfolioArrayOutputWithContext(ctx context.Context) PortfolioArrayOutput {
	return o
}

func (o PortfolioArrayOutput) Index(i pulumi.IntInput) PortfolioOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Portfolio {
		return vs[0].([]Portfolio)[vs[1].(int)]
	}).(PortfolioOutput)
}

type PortfolioMapOutput struct{ *pulumi.OutputState }

func (PortfolioMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Portfolio)(nil))
}

func (o PortfolioMapOutput) ToPortfolioMapOutput() PortfolioMapOutput {
	return o
}

func (o PortfolioMapOutput) ToPortfolioMapOutputWithContext(ctx context.Context) PortfolioMapOutput {
	return o
}

func (o PortfolioMapOutput) MapIndex(k pulumi.StringInput) PortfolioOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Portfolio {
		return vs[0].(map[string]Portfolio)[vs[1].(string)]
	}).(PortfolioOutput)
}

func init() {
	pulumi.RegisterOutputType(PortfolioOutput{})
	pulumi.RegisterOutputType(PortfolioPtrOutput{})
	pulumi.RegisterOutputType(PortfolioArrayOutput{})
	pulumi.RegisterOutputType(PortfolioMapOutput{})
}
