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

// Provides an SES domain identity resource
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ses.NewDomainIdentity(ctx, "example", &ses.DomainIdentityArgs{
//				Domain: pulumi.String("example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With Route53 Record
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ses.NewDomainIdentity(ctx, "example", &ses.DomainIdentityArgs{
//				Domain: pulumi.String("example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewRecord(ctx, "exampleAmazonsesVerificationRecord", &route53.RecordArgs{
//				ZoneId: pulumi.String("ABCDEFGHIJ123"),
//				Name:   pulumi.String("_amazonses.example.com"),
//				Type:   pulumi.String("TXT"),
//				Ttl:    pulumi.Int(600),
//				Records: pulumi.StringArray{
//					example.VerificationToken,
//				},
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
// terraform import {
//
//	to = aws_ses_domain_identity.example
//
//	id = "example.com" } Using `pulumi import`, import SES domain identities using the domain name. For exampleconsole % pulumi import aws_ses_domain_identity.example example.com
type DomainIdentity struct {
	pulumi.CustomResourceState

	// The ARN of the domain identity.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The domain name to assign to SES
	Domain pulumi.StringOutput `pulumi:"domain"`
	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done. See the With Route53 Record example
	// for how this might be achieved when the domain is hosted in Route 53 and
	// managed by this provider.  Find out more about verifying domains in Amazon
	// SES in the [AWS SES
	// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
	VerificationToken pulumi.StringOutput `pulumi:"verificationToken"`
}

// NewDomainIdentity registers a new resource with the given unique name, arguments, and options.
func NewDomainIdentity(ctx *pulumi.Context,
	name string, args *DomainIdentityArgs, opts ...pulumi.ResourceOption) (*DomainIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainIdentity
	err := ctx.RegisterResource("aws:ses/domainIdentity:DomainIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainIdentity gets an existing DomainIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainIdentityState, opts ...pulumi.ResourceOption) (*DomainIdentity, error) {
	var resource DomainIdentity
	err := ctx.ReadResource("aws:ses/domainIdentity:DomainIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainIdentity resources.
type domainIdentityState struct {
	// The ARN of the domain identity.
	Arn *string `pulumi:"arn"`
	// The domain name to assign to SES
	Domain *string `pulumi:"domain"`
	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done. See the With Route53 Record example
	// for how this might be achieved when the domain is hosted in Route 53 and
	// managed by this provider.  Find out more about verifying domains in Amazon
	// SES in the [AWS SES
	// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
	VerificationToken *string `pulumi:"verificationToken"`
}

type DomainIdentityState struct {
	// The ARN of the domain identity.
	Arn pulumi.StringPtrInput
	// The domain name to assign to SES
	Domain pulumi.StringPtrInput
	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done. See the With Route53 Record example
	// for how this might be achieved when the domain is hosted in Route 53 and
	// managed by this provider.  Find out more about verifying domains in Amazon
	// SES in the [AWS SES
	// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
	VerificationToken pulumi.StringPtrInput
}

func (DomainIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainIdentityState)(nil)).Elem()
}

type domainIdentityArgs struct {
	// The domain name to assign to SES
	Domain string `pulumi:"domain"`
}

// The set of arguments for constructing a DomainIdentity resource.
type DomainIdentityArgs struct {
	// The domain name to assign to SES
	Domain pulumi.StringInput
}

func (DomainIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainIdentityArgs)(nil)).Elem()
}

type DomainIdentityInput interface {
	pulumi.Input

	ToDomainIdentityOutput() DomainIdentityOutput
	ToDomainIdentityOutputWithContext(ctx context.Context) DomainIdentityOutput
}

func (*DomainIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainIdentity)(nil)).Elem()
}

func (i *DomainIdentity) ToDomainIdentityOutput() DomainIdentityOutput {
	return i.ToDomainIdentityOutputWithContext(context.Background())
}

func (i *DomainIdentity) ToDomainIdentityOutputWithContext(ctx context.Context) DomainIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityOutput)
}

// DomainIdentityArrayInput is an input type that accepts DomainIdentityArray and DomainIdentityArrayOutput values.
// You can construct a concrete instance of `DomainIdentityArrayInput` via:
//
//	DomainIdentityArray{ DomainIdentityArgs{...} }
type DomainIdentityArrayInput interface {
	pulumi.Input

	ToDomainIdentityArrayOutput() DomainIdentityArrayOutput
	ToDomainIdentityArrayOutputWithContext(context.Context) DomainIdentityArrayOutput
}

type DomainIdentityArray []DomainIdentityInput

func (DomainIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainIdentity)(nil)).Elem()
}

func (i DomainIdentityArray) ToDomainIdentityArrayOutput() DomainIdentityArrayOutput {
	return i.ToDomainIdentityArrayOutputWithContext(context.Background())
}

func (i DomainIdentityArray) ToDomainIdentityArrayOutputWithContext(ctx context.Context) DomainIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityArrayOutput)
}

// DomainIdentityMapInput is an input type that accepts DomainIdentityMap and DomainIdentityMapOutput values.
// You can construct a concrete instance of `DomainIdentityMapInput` via:
//
//	DomainIdentityMap{ "key": DomainIdentityArgs{...} }
type DomainIdentityMapInput interface {
	pulumi.Input

	ToDomainIdentityMapOutput() DomainIdentityMapOutput
	ToDomainIdentityMapOutputWithContext(context.Context) DomainIdentityMapOutput
}

type DomainIdentityMap map[string]DomainIdentityInput

func (DomainIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainIdentity)(nil)).Elem()
}

func (i DomainIdentityMap) ToDomainIdentityMapOutput() DomainIdentityMapOutput {
	return i.ToDomainIdentityMapOutputWithContext(context.Background())
}

func (i DomainIdentityMap) ToDomainIdentityMapOutputWithContext(ctx context.Context) DomainIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityMapOutput)
}

type DomainIdentityOutput struct{ *pulumi.OutputState }

func (DomainIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainIdentity)(nil)).Elem()
}

func (o DomainIdentityOutput) ToDomainIdentityOutput() DomainIdentityOutput {
	return o
}

func (o DomainIdentityOutput) ToDomainIdentityOutputWithContext(ctx context.Context) DomainIdentityOutput {
	return o
}

// The ARN of the domain identity.
func (o DomainIdentityOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainIdentity) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The domain name to assign to SES
func (o DomainIdentityOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainIdentity) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// A code which when added to the domain as a TXT record
// will signal to SES that the owner of the domain has authorised SES to act on
// their behalf. The domain identity will be in state "verification pending"
// until this is done. See the With Route53 Record example
// for how this might be achieved when the domain is hosted in Route 53 and
// managed by this provider.  Find out more about verifying domains in Amazon
// SES in the [AWS SES
// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
func (o DomainIdentityOutput) VerificationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainIdentity) pulumi.StringOutput { return v.VerificationToken }).(pulumi.StringOutput)
}

type DomainIdentityArrayOutput struct{ *pulumi.OutputState }

func (DomainIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainIdentity)(nil)).Elem()
}

func (o DomainIdentityArrayOutput) ToDomainIdentityArrayOutput() DomainIdentityArrayOutput {
	return o
}

func (o DomainIdentityArrayOutput) ToDomainIdentityArrayOutputWithContext(ctx context.Context) DomainIdentityArrayOutput {
	return o
}

func (o DomainIdentityArrayOutput) Index(i pulumi.IntInput) DomainIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainIdentity {
		return vs[0].([]*DomainIdentity)[vs[1].(int)]
	}).(DomainIdentityOutput)
}

type DomainIdentityMapOutput struct{ *pulumi.OutputState }

func (DomainIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainIdentity)(nil)).Elem()
}

func (o DomainIdentityMapOutput) ToDomainIdentityMapOutput() DomainIdentityMapOutput {
	return o
}

func (o DomainIdentityMapOutput) ToDomainIdentityMapOutputWithContext(ctx context.Context) DomainIdentityMapOutput {
	return o
}

func (o DomainIdentityMapOutput) MapIndex(k pulumi.StringInput) DomainIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainIdentity {
		return vs[0].(map[string]*DomainIdentity)[vs[1].(string)]
	}).(DomainIdentityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIdentityInput)(nil)).Elem(), &DomainIdentity{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIdentityArrayInput)(nil)).Elem(), DomainIdentityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIdentityMapInput)(nil)).Elem(), DomainIdentityMap{})
	pulumi.RegisterOutputType(DomainIdentityOutput{})
	pulumi.RegisterOutputType(DomainIdentityArrayOutput{})
	pulumi.RegisterOutputType(DomainIdentityMapOutput{})
}
