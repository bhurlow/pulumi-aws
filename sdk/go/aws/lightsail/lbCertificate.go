// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Lightsail load balancer Certificate resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testLb, err := lightsail.NewLb(ctx, "testLb", &lightsail.LbArgs{
//				HealthCheckPath: pulumi.String("/"),
//				InstancePort:    pulumi.Int(80),
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = lightsail.NewLbCertificate(ctx, "testLbCertificate", &lightsail.LbCertificateArgs{
//				LbName:     testLb.ID(),
//				DomainName: pulumi.String("test.com"),
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
// Using `pulumi import`, import `aws_lightsail_lb_certificate` using the id attribute. For example:
//
// ```sh
//
//	$ pulumi import aws:lightsail/lbCertificate:LbCertificate test example-load-balancer,example-load-balancer-certificate
//
// ```
type LbCertificate struct {
	pulumi.CustomResourceState

	// The ARN of the lightsail certificate.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The timestamp when the instance was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName              pulumi.StringOutput                            `pulumi:"domainName"`
	DomainValidationRecords LbCertificateDomainValidationRecordArrayOutput `pulumi:"domainValidationRecords"`
	// The load balancer name where you want to create the SSL/TLS certificate.
	LbName pulumi.StringOutput `pulumi:"lbName"`
	// The SSL/TLS certificate name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set of domains that should be SANs in the issued certificate. `domainName` attribute is automatically added as a Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	SupportCode             pulumi.StringOutput      `pulumi:"supportCode"`
}

// NewLbCertificate registers a new resource with the given unique name, arguments, and options.
func NewLbCertificate(ctx *pulumi.Context,
	name string, args *LbCertificateArgs, opts ...pulumi.ResourceOption) (*LbCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LbName == nil {
		return nil, errors.New("invalid value for required argument 'LbName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LbCertificate
	err := ctx.RegisterResource("aws:lightsail/lbCertificate:LbCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbCertificate gets an existing LbCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbCertificateState, opts ...pulumi.ResourceOption) (*LbCertificate, error) {
	var resource LbCertificate
	err := ctx.ReadResource("aws:lightsail/lbCertificate:LbCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbCertificate resources.
type lbCertificateState struct {
	// The ARN of the lightsail certificate.
	Arn *string `pulumi:"arn"`
	// The timestamp when the instance was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName              *string                               `pulumi:"domainName"`
	DomainValidationRecords []LbCertificateDomainValidationRecord `pulumi:"domainValidationRecords"`
	// The load balancer name where you want to create the SSL/TLS certificate.
	LbName *string `pulumi:"lbName"`
	// The SSL/TLS certificate name.
	Name *string `pulumi:"name"`
	// Set of domains that should be SANs in the issued certificate. `domainName` attribute is automatically added as a Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	SupportCode             *string  `pulumi:"supportCode"`
}

type LbCertificateState struct {
	// The ARN of the lightsail certificate.
	Arn pulumi.StringPtrInput
	// The timestamp when the instance was created.
	CreatedAt pulumi.StringPtrInput
	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName              pulumi.StringPtrInput
	DomainValidationRecords LbCertificateDomainValidationRecordArrayInput
	// The load balancer name where you want to create the SSL/TLS certificate.
	LbName pulumi.StringPtrInput
	// The SSL/TLS certificate name.
	Name pulumi.StringPtrInput
	// Set of domains that should be SANs in the issued certificate. `domainName` attribute is automatically added as a Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayInput
	SupportCode             pulumi.StringPtrInput
}

func (LbCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbCertificateState)(nil)).Elem()
}

type lbCertificateArgs struct {
	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName *string `pulumi:"domainName"`
	// The load balancer name where you want to create the SSL/TLS certificate.
	LbName string `pulumi:"lbName"`
	// The SSL/TLS certificate name.
	Name *string `pulumi:"name"`
	// Set of domains that should be SANs in the issued certificate. `domainName` attribute is automatically added as a Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
}

// The set of arguments for constructing a LbCertificate resource.
type LbCertificateArgs struct {
	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName pulumi.StringPtrInput
	// The load balancer name where you want to create the SSL/TLS certificate.
	LbName pulumi.StringInput
	// The SSL/TLS certificate name.
	Name pulumi.StringPtrInput
	// Set of domains that should be SANs in the issued certificate. `domainName` attribute is automatically added as a Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayInput
}

func (LbCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbCertificateArgs)(nil)).Elem()
}

type LbCertificateInput interface {
	pulumi.Input

	ToLbCertificateOutput() LbCertificateOutput
	ToLbCertificateOutputWithContext(ctx context.Context) LbCertificateOutput
}

func (*LbCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**LbCertificate)(nil)).Elem()
}

func (i *LbCertificate) ToLbCertificateOutput() LbCertificateOutput {
	return i.ToLbCertificateOutputWithContext(context.Background())
}

func (i *LbCertificate) ToLbCertificateOutputWithContext(ctx context.Context) LbCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbCertificateOutput)
}

// LbCertificateArrayInput is an input type that accepts LbCertificateArray and LbCertificateArrayOutput values.
// You can construct a concrete instance of `LbCertificateArrayInput` via:
//
//	LbCertificateArray{ LbCertificateArgs{...} }
type LbCertificateArrayInput interface {
	pulumi.Input

	ToLbCertificateArrayOutput() LbCertificateArrayOutput
	ToLbCertificateArrayOutputWithContext(context.Context) LbCertificateArrayOutput
}

type LbCertificateArray []LbCertificateInput

func (LbCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbCertificate)(nil)).Elem()
}

func (i LbCertificateArray) ToLbCertificateArrayOutput() LbCertificateArrayOutput {
	return i.ToLbCertificateArrayOutputWithContext(context.Background())
}

func (i LbCertificateArray) ToLbCertificateArrayOutputWithContext(ctx context.Context) LbCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbCertificateArrayOutput)
}

// LbCertificateMapInput is an input type that accepts LbCertificateMap and LbCertificateMapOutput values.
// You can construct a concrete instance of `LbCertificateMapInput` via:
//
//	LbCertificateMap{ "key": LbCertificateArgs{...} }
type LbCertificateMapInput interface {
	pulumi.Input

	ToLbCertificateMapOutput() LbCertificateMapOutput
	ToLbCertificateMapOutputWithContext(context.Context) LbCertificateMapOutput
}

type LbCertificateMap map[string]LbCertificateInput

func (LbCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbCertificate)(nil)).Elem()
}

func (i LbCertificateMap) ToLbCertificateMapOutput() LbCertificateMapOutput {
	return i.ToLbCertificateMapOutputWithContext(context.Background())
}

func (i LbCertificateMap) ToLbCertificateMapOutputWithContext(ctx context.Context) LbCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbCertificateMapOutput)
}

type LbCertificateOutput struct{ *pulumi.OutputState }

func (LbCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbCertificate)(nil)).Elem()
}

func (o LbCertificateOutput) ToLbCertificateOutput() LbCertificateOutput {
	return o
}

func (o LbCertificateOutput) ToLbCertificateOutputWithContext(ctx context.Context) LbCertificateOutput {
	return o
}

// The ARN of the lightsail certificate.
func (o LbCertificateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The timestamp when the instance was created.
func (o LbCertificateOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The domain name (e.g., example.com) for your SSL/TLS certificate.
func (o LbCertificateOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o LbCertificateOutput) DomainValidationRecords() LbCertificateDomainValidationRecordArrayOutput {
	return o.ApplyT(func(v *LbCertificate) LbCertificateDomainValidationRecordArrayOutput {
		return v.DomainValidationRecords
	}).(LbCertificateDomainValidationRecordArrayOutput)
}

// The load balancer name where you want to create the SSL/TLS certificate.
func (o LbCertificateOutput) LbName() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.LbName }).(pulumi.StringOutput)
}

// The SSL/TLS certificate name.
func (o LbCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set of domains that should be SANs in the issued certificate. `domainName` attribute is automatically added as a Subject Alternative Name.
func (o LbCertificateOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringArrayOutput { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

func (o LbCertificateOutput) SupportCode() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.SupportCode }).(pulumi.StringOutput)
}

type LbCertificateArrayOutput struct{ *pulumi.OutputState }

func (LbCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbCertificate)(nil)).Elem()
}

func (o LbCertificateArrayOutput) ToLbCertificateArrayOutput() LbCertificateArrayOutput {
	return o
}

func (o LbCertificateArrayOutput) ToLbCertificateArrayOutputWithContext(ctx context.Context) LbCertificateArrayOutput {
	return o
}

func (o LbCertificateArrayOutput) Index(i pulumi.IntInput) LbCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbCertificate {
		return vs[0].([]*LbCertificate)[vs[1].(int)]
	}).(LbCertificateOutput)
}

type LbCertificateMapOutput struct{ *pulumi.OutputState }

func (LbCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbCertificate)(nil)).Elem()
}

func (o LbCertificateMapOutput) ToLbCertificateMapOutput() LbCertificateMapOutput {
	return o
}

func (o LbCertificateMapOutput) ToLbCertificateMapOutputWithContext(ctx context.Context) LbCertificateMapOutput {
	return o
}

func (o LbCertificateMapOutput) MapIndex(k pulumi.StringInput) LbCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbCertificate {
		return vs[0].(map[string]*LbCertificate)[vs[1].(string)]
	}).(LbCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbCertificateInput)(nil)).Elem(), &LbCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbCertificateArrayInput)(nil)).Elem(), LbCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbCertificateMapInput)(nil)).Elem(), LbCertificateMap{})
	pulumi.RegisterOutputType(LbCertificateOutput{})
	pulumi.RegisterOutputType(LbCertificateArrayOutput{})
	pulumi.RegisterOutputType(LbCertificateMapOutput{})
}
