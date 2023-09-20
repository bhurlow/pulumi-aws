// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package acm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The ACM certificate resource allows requesting and management of certificates
// from the Amazon Certificate Manager.
//
// ACM certificates can be created in three ways:
// Amazon-issued, where AWS provides the certificate authority and automatically manages renewal;
// imported certificates, issued by another certificate authority;
// and private certificates, issued using an ACM Private Certificate Authority.
//
// ## Amazon-Issued Certificates
//
// For Amazon-issued certificates, this resource deals with requesting certificates and managing their attributes and life-cycle.
// This resource does not deal with validation of a certificate but can provide inputs
// for other resources implementing the validation.
// It does not wait for a certificate to be issued.
// Use a `acm.CertificateValidation` resource for this.
//
// Most commonly, this resource is used together with `route53.Record` and
// `acm.CertificateValidation` to request a DNS validated certificate,
// deploy the required validation records and wait for validation to complete.
//
// Domain validation through email is also supported but should be avoided as it requires a manual step outside of this provider.
//
// ## Certificates Imported from Other Certificate Authority
//
// Imported certificates can be used to make certificates created with an external certificate authority available for AWS services.
//
// As they are not managed by AWS, imported certificates are not eligible for automatic renewal.
// New certificate materials can be supplied to an existing imported certificate to update it in place.
//
// ## Private Certificates
//
// Private certificates are issued by an ACM Private Cerificate Authority, which can be created using the resource type `acmpca.CertificateAuthority`.
//
// Private certificates created using this resource are eligible for managed renewal if they have been exported or associated with another AWS service.
// See [managed renewal documentation](https://docs.aws.amazon.com/acm/latest/userguide/managed-renewal.html) for more information.
// By default, a certificate is valid for 395 days and the managed renewal process will start 60 days before expiration.
// To renew the certificate earlier than 60 days before expiration, configure `earlyRenewalDuration`.
//
// ## Example Usage
// ### Create Certificate
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := acm.NewCertificate(ctx, "cert", &acm.CertificateArgs{
//				DomainName: pulumi.String("example.com"),
//				Tags: pulumi.StringMap{
//					"Environment": pulumi.String("test"),
//				},
//				ValidationMethod: pulumi.String("DNS"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Custom Domain Validation Options
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := acm.NewCertificate(ctx, "cert", &acm.CertificateArgs{
//				DomainName:       pulumi.String("testing.example.com"),
//				ValidationMethod: pulumi.String("EMAIL"),
//				ValidationOptions: acm.CertificateValidationOptionArray{
//					&acm.CertificateValidationOptionArgs{
//						DomainName:       pulumi.String("testing.example.com"),
//						ValidationDomain: pulumi.String("example.com"),
//					},
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
// ### Existing Certificate Body Import
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
//	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			examplePrivateKey, err := tls.NewPrivateKey(ctx, "examplePrivateKey", &tls.PrivateKeyArgs{
//				Algorithm: pulumi.String("RSA"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSelfSignedCert, err := tls.NewSelfSignedCert(ctx, "exampleSelfSignedCert", &tls.SelfSignedCertArgs{
//				KeyAlgorithm:  pulumi.String("RSA"),
//				PrivateKeyPem: examplePrivateKey.PrivateKeyPem,
//				Subject: &tls.SelfSignedCertSubjectArgs{
//					CommonName:   pulumi.String("example.com"),
//					Organization: pulumi.String("ACME Examples, Inc"),
//				},
//				ValidityPeriodHours: pulumi.Int(12),
//				AllowedUses: pulumi.StringArray{
//					pulumi.String("key_encipherment"),
//					pulumi.String("digital_signature"),
//					pulumi.String("server_auth"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = acm.NewCertificate(ctx, "cert", &acm.CertificateArgs{
//				PrivateKey:      examplePrivateKey.PrivateKeyPem,
//				CertificateBody: exampleSelfSignedCert.CertPem,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Referencing domainValidationOptions With forEach Based Resources
//
// See the `acm.CertificateValidation` resource for a full example of performing DNS validation.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			var example []*route53.Record
//			for key0, val0 := range "TODO: For expression" {
//				__res, err := route53.NewRecord(ctx, fmt.Sprintf("example-%v", key0), &route53.RecordArgs{
//					AllowOverwrite: pulumi.Bool(true),
//					Name:           pulumi.Any(val0),
//					Records: pulumi.StringArray{
//						val0,
//					},
//					Ttl:    pulumi.Int(60),
//					Type:   route53.RecordType(val0),
//					ZoneId: pulumi.Any(aws_route53_zone.Example.Zone_id),
//				})
//				if err != nil {
//					return err
//				}
//				example = append(example, __res)
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import certificates using their ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:acm/certificate:Certificate cert arn:aws:acm:eu-central-1:123456789012:certificate/7e7a28d2-163f-4b8f-b9cd-822f96c08d6a
//
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// ARN of the certificate
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ARN of an ACM PCA
	CertificateAuthorityArn pulumi.StringPtrOutput `pulumi:"certificateAuthorityArn"`
	// Certificate's PEM-formatted public key
	CertificateBody pulumi.StringPtrOutput `pulumi:"certificateBody"`
	// Certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain pulumi.StringPtrOutput `pulumi:"certificateChain"`
	// Fully qualified domain name (FQDN) in the certificate.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Set of domain validation objects which can be used to complete certificate validation.
	// Can have more than one element, e.g., if SANs are defined.
	// Only set if `DNS`-validation was used.
	DomainValidationOptions CertificateDomainValidationOptionArrayOutput `pulumi:"domainValidationOptions"`
	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
	// or a string such as `2160h`.
	EarlyRenewalDuration pulumi.StringPtrOutput `pulumi:"earlyRenewalDuration"`
	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// Expiration date and time of the certificate.
	NotAfter pulumi.StringOutput `pulumi:"notAfter"`
	// Start of the validity period of the certificate.
	NotBefore pulumi.StringOutput `pulumi:"notBefore"`
	// Configuration block used to set certificate options. Detailed below.
	Options CertificateOptionsOutput `pulumi:"options"`
	// `true` if a Private certificate eligible for managed renewal is within the `earlyRenewalDuration` period.
	PendingRenewal pulumi.BoolOutput `pulumi:"pendingRenewal"`
	// Certificate's PEM-formatted private key
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// Whether the certificate is eligible for managed renewal.
	RenewalEligibility pulumi.StringOutput `pulumi:"renewalEligibility"`
	// Contains information about the status of ACM's [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the certificate.
	RenewalSummaries CertificateRenewalSummaryArrayOutput `pulumi:"renewalSummaries"`
	// Status of the certificate.
	Status pulumi.StringOutput `pulumi:"status"`
	// Set of domains that should be SANs in the issued certificate.
	// To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Source of the certificate.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of addresses that received a validation email. Only set if `EMAIL` validation was used.
	ValidationEmails pulumi.StringArrayOutput `pulumi:"validationEmails"`
	// Which method to use for validation. `DNS` or `EMAIL` are valid. This parameter must not be set for certificates that were imported into ACM and then into Pulumi.
	ValidationMethod pulumi.StringOutput `pulumi:"validationMethod"`
	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	// * Importing an existing certificate
	ValidationOptions CertificateValidationOptionArrayOutput `pulumi:"validationOptions"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		args = &CertificateArgs{}
	}

	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// ARN of the certificate
	Arn *string `pulumi:"arn"`
	// ARN of an ACM PCA
	CertificateAuthorityArn *string `pulumi:"certificateAuthorityArn"`
	// Certificate's PEM-formatted public key
	CertificateBody *string `pulumi:"certificateBody"`
	// Certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain *string `pulumi:"certificateChain"`
	// Fully qualified domain name (FQDN) in the certificate.
	DomainName *string `pulumi:"domainName"`
	// Set of domain validation objects which can be used to complete certificate validation.
	// Can have more than one element, e.g., if SANs are defined.
	// Only set if `DNS`-validation was used.
	DomainValidationOptions []CertificateDomainValidationOption `pulumi:"domainValidationOptions"`
	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
	// or a string such as `2160h`.
	EarlyRenewalDuration *string `pulumi:"earlyRenewalDuration"`
	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// Expiration date and time of the certificate.
	NotAfter *string `pulumi:"notAfter"`
	// Start of the validity period of the certificate.
	NotBefore *string `pulumi:"notBefore"`
	// Configuration block used to set certificate options. Detailed below.
	Options *CertificateOptions `pulumi:"options"`
	// `true` if a Private certificate eligible for managed renewal is within the `earlyRenewalDuration` period.
	PendingRenewal *bool `pulumi:"pendingRenewal"`
	// Certificate's PEM-formatted private key
	PrivateKey *string `pulumi:"privateKey"`
	// Whether the certificate is eligible for managed renewal.
	RenewalEligibility *string `pulumi:"renewalEligibility"`
	// Contains information about the status of ACM's [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the certificate.
	RenewalSummaries []CertificateRenewalSummary `pulumi:"renewalSummaries"`
	// Status of the certificate.
	Status *string `pulumi:"status"`
	// Set of domains that should be SANs in the issued certificate.
	// To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Source of the certificate.
	Type *string `pulumi:"type"`
	// List of addresses that received a validation email. Only set if `EMAIL` validation was used.
	ValidationEmails []string `pulumi:"validationEmails"`
	// Which method to use for validation. `DNS` or `EMAIL` are valid. This parameter must not be set for certificates that were imported into ACM and then into Pulumi.
	ValidationMethod *string `pulumi:"validationMethod"`
	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	// * Importing an existing certificate
	ValidationOptions []CertificateValidationOption `pulumi:"validationOptions"`
}

type CertificateState struct {
	// ARN of the certificate
	Arn pulumi.StringPtrInput
	// ARN of an ACM PCA
	CertificateAuthorityArn pulumi.StringPtrInput
	// Certificate's PEM-formatted public key
	CertificateBody pulumi.StringPtrInput
	// Certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain pulumi.StringPtrInput
	// Fully qualified domain name (FQDN) in the certificate.
	DomainName pulumi.StringPtrInput
	// Set of domain validation objects which can be used to complete certificate validation.
	// Can have more than one element, e.g., if SANs are defined.
	// Only set if `DNS`-validation was used.
	DomainValidationOptions CertificateDomainValidationOptionArrayInput
	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
	// or a string such as `2160h`.
	EarlyRenewalDuration pulumi.StringPtrInput
	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
	KeyAlgorithm pulumi.StringPtrInput
	// Expiration date and time of the certificate.
	NotAfter pulumi.StringPtrInput
	// Start of the validity period of the certificate.
	NotBefore pulumi.StringPtrInput
	// Configuration block used to set certificate options. Detailed below.
	Options CertificateOptionsPtrInput
	// `true` if a Private certificate eligible for managed renewal is within the `earlyRenewalDuration` period.
	PendingRenewal pulumi.BoolPtrInput
	// Certificate's PEM-formatted private key
	PrivateKey pulumi.StringPtrInput
	// Whether the certificate is eligible for managed renewal.
	RenewalEligibility pulumi.StringPtrInput
	// Contains information about the status of ACM's [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the certificate.
	RenewalSummaries CertificateRenewalSummaryArrayInput
	// Status of the certificate.
	Status pulumi.StringPtrInput
	// Set of domains that should be SANs in the issued certificate.
	// To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
	SubjectAlternativeNames pulumi.StringArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Source of the certificate.
	Type pulumi.StringPtrInput
	// List of addresses that received a validation email. Only set if `EMAIL` validation was used.
	ValidationEmails pulumi.StringArrayInput
	// Which method to use for validation. `DNS` or `EMAIL` are valid. This parameter must not be set for certificates that were imported into ACM and then into Pulumi.
	ValidationMethod pulumi.StringPtrInput
	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	// * Importing an existing certificate
	ValidationOptions CertificateValidationOptionArrayInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// ARN of an ACM PCA
	CertificateAuthorityArn *string `pulumi:"certificateAuthorityArn"`
	// Certificate's PEM-formatted public key
	CertificateBody *string `pulumi:"certificateBody"`
	// Certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain *string `pulumi:"certificateChain"`
	// Fully qualified domain name (FQDN) in the certificate.
	DomainName *string `pulumi:"domainName"`
	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
	// or a string such as `2160h`.
	EarlyRenewalDuration *string `pulumi:"earlyRenewalDuration"`
	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// Configuration block used to set certificate options. Detailed below.
	Options *CertificateOptions `pulumi:"options"`
	// Certificate's PEM-formatted private key
	PrivateKey *string `pulumi:"privateKey"`
	// Set of domains that should be SANs in the issued certificate.
	// To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Which method to use for validation. `DNS` or `EMAIL` are valid. This parameter must not be set for certificates that were imported into ACM and then into Pulumi.
	ValidationMethod *string `pulumi:"validationMethod"`
	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	// * Importing an existing certificate
	ValidationOptions []CertificateValidationOption `pulumi:"validationOptions"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// ARN of an ACM PCA
	CertificateAuthorityArn pulumi.StringPtrInput
	// Certificate's PEM-formatted public key
	CertificateBody pulumi.StringPtrInput
	// Certificate's PEM-formatted chain
	// * Creating a private CA issued certificate
	CertificateChain pulumi.StringPtrInput
	// Fully qualified domain name (FQDN) in the certificate.
	DomainName pulumi.StringPtrInput
	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
	// or a string such as `2160h`.
	EarlyRenewalDuration pulumi.StringPtrInput
	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
	KeyAlgorithm pulumi.StringPtrInput
	// Configuration block used to set certificate options. Detailed below.
	Options CertificateOptionsPtrInput
	// Certificate's PEM-formatted private key
	PrivateKey pulumi.StringPtrInput
	// Set of domains that should be SANs in the issued certificate.
	// To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
	SubjectAlternativeNames pulumi.StringArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Which method to use for validation. `DNS` or `EMAIL` are valid. This parameter must not be set for certificates that were imported into ACM and then into Pulumi.
	ValidationMethod pulumi.StringPtrInput
	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	// * Importing an existing certificate
	ValidationOptions CertificateValidationOptionArrayInput
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
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

func (i *Certificate) ToOutput(ctx context.Context) pulumix.Output[*Certificate] {
	return pulumix.Output[*Certificate]{
		OutputState: i.ToCertificateOutputWithContext(ctx).OutputState,
	}
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//	CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

func (i CertificateArray) ToOutput(ctx context.Context) pulumix.Output[[]*Certificate] {
	return pulumix.Output[[]*Certificate]{
		OutputState: i.ToCertificateArrayOutputWithContext(ctx).OutputState,
	}
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//	CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

func (i CertificateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Certificate] {
	return pulumix.Output[map[string]*Certificate]{
		OutputState: i.ToCertificateMapOutputWithContext(ctx).OutputState,
	}
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) ToOutput(ctx context.Context) pulumix.Output[*Certificate] {
	return pulumix.Output[*Certificate]{
		OutputState: o.OutputState,
	}
}

// ARN of the certificate
func (o CertificateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// ARN of an ACM PCA
func (o CertificateOutput) CertificateAuthorityArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.CertificateAuthorityArn }).(pulumi.StringPtrOutput)
}

// Certificate's PEM-formatted public key
func (o CertificateOutput) CertificateBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.CertificateBody }).(pulumi.StringPtrOutput)
}

// Certificate's PEM-formatted chain
// * Creating a private CA issued certificate
func (o CertificateOutput) CertificateChain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.CertificateChain }).(pulumi.StringPtrOutput)
}

// Fully qualified domain name (FQDN) in the certificate.
func (o CertificateOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Set of domain validation objects which can be used to complete certificate validation.
// Can have more than one element, e.g., if SANs are defined.
// Only set if `DNS`-validation was used.
func (o CertificateOutput) DomainValidationOptions() CertificateDomainValidationOptionArrayOutput {
	return o.ApplyT(func(v *Certificate) CertificateDomainValidationOptionArrayOutput { return v.DomainValidationOptions }).(CertificateDomainValidationOptionArrayOutput)
}

// Amount of time to start automatic renewal process before expiration.
// Has no effect if less than 60 days.
// Represented by either
// a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
// or a string such as `2160h`.
func (o CertificateOutput) EarlyRenewalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.EarlyRenewalDuration }).(pulumi.StringPtrOutput)
}

// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
func (o CertificateOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// Expiration date and time of the certificate.
func (o CertificateOutput) NotAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.NotAfter }).(pulumi.StringOutput)
}

// Start of the validity period of the certificate.
func (o CertificateOutput) NotBefore() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.NotBefore }).(pulumi.StringOutput)
}

// Configuration block used to set certificate options. Detailed below.
func (o CertificateOutput) Options() CertificateOptionsOutput {
	return o.ApplyT(func(v *Certificate) CertificateOptionsOutput { return v.Options }).(CertificateOptionsOutput)
}

// `true` if a Private certificate eligible for managed renewal is within the `earlyRenewalDuration` period.
func (o CertificateOutput) PendingRenewal() pulumi.BoolOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolOutput { return v.PendingRenewal }).(pulumi.BoolOutput)
}

// Certificate's PEM-formatted private key
func (o CertificateOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

// Whether the certificate is eligible for managed renewal.
func (o CertificateOutput) RenewalEligibility() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.RenewalEligibility }).(pulumi.StringOutput)
}

// Contains information about the status of ACM's [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the certificate.
func (o CertificateOutput) RenewalSummaries() CertificateRenewalSummaryArrayOutput {
	return o.ApplyT(func(v *Certificate) CertificateRenewalSummaryArrayOutput { return v.RenewalSummaries }).(CertificateRenewalSummaryArrayOutput)
}

// Status of the certificate.
func (o CertificateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Set of domains that should be SANs in the issued certificate.
// To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
func (o CertificateOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o CertificateOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Source of the certificate.
func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of addresses that received a validation email. Only set if `EMAIL` validation was used.
func (o CertificateOutput) ValidationEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.ValidationEmails }).(pulumi.StringArrayOutput)
}

// Which method to use for validation. `DNS` or `EMAIL` are valid. This parameter must not be set for certificates that were imported into ACM and then into Pulumi.
func (o CertificateOutput) ValidationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ValidationMethod }).(pulumi.StringOutput)
}

// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
// * Importing an existing certificate
func (o CertificateOutput) ValidationOptions() CertificateValidationOptionArrayOutput {
	return o.ApplyT(func(v *Certificate) CertificateValidationOptionArrayOutput { return v.ValidationOptions }).(CertificateValidationOptionArrayOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Certificate] {
	return pulumix.Output[[]*Certificate]{
		OutputState: o.OutputState,
	}
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Certificate] {
	return pulumix.Output[map[string]*Certificate]{
		OutputState: o.OutputState,
	}
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
