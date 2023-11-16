// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SES domain DKIM generation resource.
//
// # Domain ownership needs to be confirmed first using sesDomainIdentity Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDomainIdentity, err := ses.NewDomainIdentity(ctx, "exampleDomainIdentity", &ses.DomainIdentityArgs{
//				Domain: pulumi.String("example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleDomainDkim, err := ses.NewDomainDkim(ctx, "exampleDomainDkim", &ses.DomainDkimArgs{
//				Domain: exampleDomainIdentity.Domain,
//			})
//			if err != nil {
//				return err
//			}
//			var exampleAmazonsesDkimRecord []*route53.Record
//			for index := 0; index < 3; index++ {
//				key0 := index
//				val0 := index
//				__res, err := route53.NewRecord(ctx, fmt.Sprintf("exampleAmazonsesDkimRecord-%v", key0), &route53.RecordArgs{
//					ZoneId: pulumi.String("ABCDEFGHIJ123"),
//					Name: exampleDomainDkim.DkimTokens.ApplyT(func(dkimTokens []string) (string, error) {
//						return fmt.Sprintf("%v._domainkey", dkimTokens[val0]), nil
//					}).(pulumi.StringOutput),
//					Type: pulumi.String("CNAME"),
//					Ttl:  pulumi.Int(600),
//					Records: pulumi.StringArray{
//						exampleDomainDkim.DkimTokens.ApplyT(func(dkimTokens []string) (string, error) {
//							return fmt.Sprintf("%v.dkim.amazonses.com", dkimTokens[val0]), nil
//						}).(pulumi.StringOutput),
//					},
//				})
//				if err != nil {
//					return err
//				}
//				exampleAmazonsesDkimRecord = append(exampleAmazonsesDkimRecord, __res)
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import DKIM tokens using the `domain` attribute. For example:
//
// ```sh
//
//	$ pulumi import aws:ses/domainDkim:DomainDkim example example.com
//
// ```
type DomainDkim struct {
	pulumi.CustomResourceState

	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens pulumi.StringArrayOutput `pulumi:"dkimTokens"`
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringOutput `pulumi:"domain"`
}

// NewDomainDkim registers a new resource with the given unique name, arguments, and options.
func NewDomainDkim(ctx *pulumi.Context,
	name string, args *DomainDkimArgs, opts ...pulumi.ResourceOption) (*DomainDkim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainDkim
	err := ctx.RegisterResource("aws:ses/domainDkim:DomainDkim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainDkim gets an existing DomainDkim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainDkim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainDkimState, opts ...pulumi.ResourceOption) (*DomainDkim, error) {
	var resource DomainDkim
	err := ctx.ReadResource("aws:ses/domainDkim:DomainDkim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainDkim resources.
type domainDkimState struct {
	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens []string `pulumi:"dkimTokens"`
	// Verified domain name to generate DKIM tokens for.
	Domain *string `pulumi:"domain"`
}

type DomainDkimState struct {
	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens pulumi.StringArrayInput
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringPtrInput
}

func (DomainDkimState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainDkimState)(nil)).Elem()
}

type domainDkimArgs struct {
	// Verified domain name to generate DKIM tokens for.
	Domain string `pulumi:"domain"`
}

// The set of arguments for constructing a DomainDkim resource.
type DomainDkimArgs struct {
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringInput
}

func (DomainDkimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainDkimArgs)(nil)).Elem()
}

type DomainDkimInput interface {
	pulumi.Input

	ToDomainDkimOutput() DomainDkimOutput
	ToDomainDkimOutputWithContext(ctx context.Context) DomainDkimOutput
}

func (*DomainDkim) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainDkim)(nil)).Elem()
}

func (i *DomainDkim) ToDomainDkimOutput() DomainDkimOutput {
	return i.ToDomainDkimOutputWithContext(context.Background())
}

func (i *DomainDkim) ToDomainDkimOutputWithContext(ctx context.Context) DomainDkimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimOutput)
}

// DomainDkimArrayInput is an input type that accepts DomainDkimArray and DomainDkimArrayOutput values.
// You can construct a concrete instance of `DomainDkimArrayInput` via:
//
//	DomainDkimArray{ DomainDkimArgs{...} }
type DomainDkimArrayInput interface {
	pulumi.Input

	ToDomainDkimArrayOutput() DomainDkimArrayOutput
	ToDomainDkimArrayOutputWithContext(context.Context) DomainDkimArrayOutput
}

type DomainDkimArray []DomainDkimInput

func (DomainDkimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainDkim)(nil)).Elem()
}

func (i DomainDkimArray) ToDomainDkimArrayOutput() DomainDkimArrayOutput {
	return i.ToDomainDkimArrayOutputWithContext(context.Background())
}

func (i DomainDkimArray) ToDomainDkimArrayOutputWithContext(ctx context.Context) DomainDkimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimArrayOutput)
}

// DomainDkimMapInput is an input type that accepts DomainDkimMap and DomainDkimMapOutput values.
// You can construct a concrete instance of `DomainDkimMapInput` via:
//
//	DomainDkimMap{ "key": DomainDkimArgs{...} }
type DomainDkimMapInput interface {
	pulumi.Input

	ToDomainDkimMapOutput() DomainDkimMapOutput
	ToDomainDkimMapOutputWithContext(context.Context) DomainDkimMapOutput
}

type DomainDkimMap map[string]DomainDkimInput

func (DomainDkimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainDkim)(nil)).Elem()
}

func (i DomainDkimMap) ToDomainDkimMapOutput() DomainDkimMapOutput {
	return i.ToDomainDkimMapOutputWithContext(context.Background())
}

func (i DomainDkimMap) ToDomainDkimMapOutputWithContext(ctx context.Context) DomainDkimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimMapOutput)
}

type DomainDkimOutput struct{ *pulumi.OutputState }

func (DomainDkimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainDkim)(nil)).Elem()
}

func (o DomainDkimOutput) ToDomainDkimOutput() DomainDkimOutput {
	return o
}

func (o DomainDkimOutput) ToDomainDkimOutputWithContext(ctx context.Context) DomainDkimOutput {
	return o
}

// DKIM tokens generated by SES.
// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
// See below for an example of how this might be achieved
// when the domain is hosted in Route 53 and managed by this provider.
// Find out more about verifying domains in Amazon SES
// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
func (o DomainDkimOutput) DkimTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainDkim) pulumi.StringArrayOutput { return v.DkimTokens }).(pulumi.StringArrayOutput)
}

// Verified domain name to generate DKIM tokens for.
func (o DomainDkimOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainDkim) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

type DomainDkimArrayOutput struct{ *pulumi.OutputState }

func (DomainDkimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainDkim)(nil)).Elem()
}

func (o DomainDkimArrayOutput) ToDomainDkimArrayOutput() DomainDkimArrayOutput {
	return o
}

func (o DomainDkimArrayOutput) ToDomainDkimArrayOutputWithContext(ctx context.Context) DomainDkimArrayOutput {
	return o
}

func (o DomainDkimArrayOutput) Index(i pulumi.IntInput) DomainDkimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainDkim {
		return vs[0].([]*DomainDkim)[vs[1].(int)]
	}).(DomainDkimOutput)
}

type DomainDkimMapOutput struct{ *pulumi.OutputState }

func (DomainDkimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainDkim)(nil)).Elem()
}

func (o DomainDkimMapOutput) ToDomainDkimMapOutput() DomainDkimMapOutput {
	return o
}

func (o DomainDkimMapOutput) ToDomainDkimMapOutputWithContext(ctx context.Context) DomainDkimMapOutput {
	return o
}

func (o DomainDkimMapOutput) MapIndex(k pulumi.StringInput) DomainDkimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainDkim {
		return vs[0].(map[string]*DomainDkim)[vs[1].(string)]
	}).(DomainDkimOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainDkimInput)(nil)).Elem(), &DomainDkim{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainDkimArrayInput)(nil)).Elem(), DomainDkimArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainDkimMapInput)(nil)).Elem(), DomainDkimMap{})
	pulumi.RegisterOutputType(DomainDkimOutput{})
	pulumi.RegisterOutputType(DomainDkimArrayOutput{})
	pulumi.RegisterOutputType(DomainDkimMapOutput{})
}
