// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package acm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ACM certificate resource allows requesting and management of certificates
// from the Amazon Certificate Manager.
//
// It deals with requesting certificates and managing their attributes and life-cycle.
// This resource does not deal with validation of a certificate but can provide inputs
// for other resources implementing the validation. It does not wait for a certificate to be issued.
// Use a `acm.CertificateValidation` resource for this.
//
// Most commonly, this resource is used together with `route53.Record` and
// `acm.CertificateValidation` to request a DNS validated certificate,
// deploy the required validation records and wait for validation to complete.
//
// Domain validation through E-Mail is also supported but should be avoided as it requires a manual step outside
// of this provider.
//
// It's recommended to specify `createBeforeDestroy = true` in a [lifecycle](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html) block to replace a certificate
// which is currently in use (eg, by `lb.Listener`).
//
// ## Example Usage
// ### Certificate creation
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/acm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := acm.NewCertificate(ctx, "cert", &acm.CertificateArgs{
// 			DomainName: pulumi.String("example.com"),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("test"),
// 			},
// 			ValidationMethod: pulumi.String("DNS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Importing an existing certificate
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/acm"
// 	"github.com/pulumi/pulumi-tls/sdk/v2/go/tls"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		examplePrivateKey, err := tls.NewPrivateKey(ctx, "examplePrivateKey", &tls.PrivateKeyArgs{
// 			Algorithm: pulumi.String("RSA"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSelfSignedCert, err := tls.NewSelfSignedCert(ctx, "exampleSelfSignedCert", &tls.SelfSignedCertArgs{
// 			KeyAlgorithm:  pulumi.String("RSA"),
// 			PrivateKeyPem: examplePrivateKey.PrivateKeyPem,
// 			Subjects: tls.SelfSignedCertSubjectArray{
// 				&tls.SelfSignedCertSubjectArgs{
// 					CommonName:   pulumi.String("example.com"),
// 					Organization: pulumi.String("ACME Examples, Inc"),
// 				},
// 			},
// 			ValidityPeriodHours: pulumi.Int(12),
// 			AllowedUses: pulumi.StringArray{
// 				pulumi.String("key_encipherment"),
// 				pulumi.String("digital_signature"),
// 				pulumi.String("server_auth"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = acm.NewCertificate(ctx, "cert", &acm.CertificateArgs{
// 			PrivateKey:      examplePrivateKey.PrivateKeyPem,
// 			CertificateBody: exampleSelfSignedCert.CertPem,
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
// Certificates can be imported using their ARN, e.g.
//
// ```sh
//  $ pulumi import aws:acm/certificate:Certificate cert arn:aws:acm:eu-central-1:123456789012:certificate/7e7a28d2-163f-4b8f-b9cd-822f96c08d6a
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// The ARN of the certificate
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ARN of an ACMPCA
	CertificateAuthorityArn pulumi.StringPtrOutput `pulumi:"certificateAuthorityArn"`
	// The certificate's PEM-formatted public key
	CertificateBody pulumi.StringPtrOutput `pulumi:"certificateBody"`
	// The certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain pulumi.StringPtrOutput `pulumi:"certificateChain"`
	// A domain name for which the certificate should be issued
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g. if SANs are defined. Only set if `DNS`-validation was used.
	DomainValidationOptions CertificateDomainValidationOptionArrayOutput `pulumi:"domainValidationOptions"`
	// Configuration block used to set certificate options. Detailed below.
	// * Importing an existing certificate
	Options CertificateOptionsPtrOutput `pulumi:"options"`
	// The certificate's PEM-formatted private key
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// Status of the certificate.
	Status pulumi.StringOutput `pulumi:"status"`
	// Set of domains that should be SANs in the issued certificate. To remove all elements of a previously configured list, set this value equal to an empty list (`[]`) to trigger recreation.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of addresses that received a validation E-Mail. Only set if `EMAIL`-validation was used.
	ValidationEmails pulumi.StringArrayOutput `pulumi:"validationEmails"`
	// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
	ValidationMethod pulumi.StringOutput `pulumi:"validationMethod"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		args = &CertificateArgs{}
	}

	var resource Certificate
	err := ctx.RegisterResource("aws:acm/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("aws:acm/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// The ARN of the certificate
	Arn *string `pulumi:"arn"`
	// ARN of an ACMPCA
	CertificateAuthorityArn *string `pulumi:"certificateAuthorityArn"`
	// The certificate's PEM-formatted public key
	CertificateBody *string `pulumi:"certificateBody"`
	// The certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain *string `pulumi:"certificateChain"`
	// A domain name for which the certificate should be issued
	DomainName *string `pulumi:"domainName"`
	// Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g. if SANs are defined. Only set if `DNS`-validation was used.
	DomainValidationOptions []CertificateDomainValidationOption `pulumi:"domainValidationOptions"`
	// Configuration block used to set certificate options. Detailed below.
	// * Importing an existing certificate
	Options *CertificateOptions `pulumi:"options"`
	// The certificate's PEM-formatted private key
	PrivateKey *string `pulumi:"privateKey"`
	// Status of the certificate.
	Status *string `pulumi:"status"`
	// Set of domains that should be SANs in the issued certificate. To remove all elements of a previously configured list, set this value equal to an empty list (`[]`) to trigger recreation.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of addresses that received a validation E-Mail. Only set if `EMAIL`-validation was used.
	ValidationEmails []string `pulumi:"validationEmails"`
	// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
	ValidationMethod *string `pulumi:"validationMethod"`
}

type CertificateState struct {
	// The ARN of the certificate
	Arn pulumi.StringPtrInput
	// ARN of an ACMPCA
	CertificateAuthorityArn pulumi.StringPtrInput
	// The certificate's PEM-formatted public key
	CertificateBody pulumi.StringPtrInput
	// The certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain pulumi.StringPtrInput
	// A domain name for which the certificate should be issued
	DomainName pulumi.StringPtrInput
	// Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g. if SANs are defined. Only set if `DNS`-validation was used.
	DomainValidationOptions CertificateDomainValidationOptionArrayInput
	// Configuration block used to set certificate options. Detailed below.
	// * Importing an existing certificate
	Options CertificateOptionsPtrInput
	// The certificate's PEM-formatted private key
	PrivateKey pulumi.StringPtrInput
	// Status of the certificate.
	Status pulumi.StringPtrInput
	// Set of domains that should be SANs in the issued certificate. To remove all elements of a previously configured list, set this value equal to an empty list (`[]`) to trigger recreation.
	SubjectAlternativeNames pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of addresses that received a validation E-Mail. Only set if `EMAIL`-validation was used.
	ValidationEmails pulumi.StringArrayInput
	// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
	ValidationMethod pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// ARN of an ACMPCA
	CertificateAuthorityArn *string `pulumi:"certificateAuthorityArn"`
	// The certificate's PEM-formatted public key
	CertificateBody *string `pulumi:"certificateBody"`
	// The certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain *string `pulumi:"certificateChain"`
	// A domain name for which the certificate should be issued
	DomainName *string `pulumi:"domainName"`
	// Configuration block used to set certificate options. Detailed below.
	// * Importing an existing certificate
	Options *CertificateOptions `pulumi:"options"`
	// The certificate's PEM-formatted private key
	PrivateKey *string `pulumi:"privateKey"`
	// Set of domains that should be SANs in the issued certificate. To remove all elements of a previously configured list, set this value equal to an empty list (`[]`) to trigger recreation.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
	ValidationMethod *string `pulumi:"validationMethod"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// ARN of an ACMPCA
	CertificateAuthorityArn pulumi.StringPtrInput
	// The certificate's PEM-formatted public key
	CertificateBody pulumi.StringPtrInput
	// The certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain pulumi.StringPtrInput
	// A domain name for which the certificate should be issued
	DomainName pulumi.StringPtrInput
	// Configuration block used to set certificate options. Detailed below.
	// * Importing an existing certificate
	Options CertificateOptionsPtrInput
	// The certificate's PEM-formatted private key
	PrivateKey pulumi.StringPtrInput
	// Set of domains that should be SANs in the issued certificate. To remove all elements of a previously configured list, set this value equal to an empty list (`[]`) to trigger recreation.
	SubjectAlternativeNames pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
	ValidationMethod pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
