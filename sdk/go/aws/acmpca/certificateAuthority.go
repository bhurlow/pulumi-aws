// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package acmpca

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).
//
// > **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificateSigningRequest` attribute. The `acmpca.CertificateAuthorityCertificate` resource can be used for this purpose.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/acmpca"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := acmpca.NewCertificateAuthority(ctx, "example", &acmpca.CertificateAuthorityArgs{
//				CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
//					KeyAlgorithm:     pulumi.String("RSA_4096"),
//					SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
//					Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
//						CommonName: pulumi.String("example.com"),
//					},
//				},
//				PermanentDeletionTimeInDays: pulumi.Int(7),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Short-lived certificate
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/acmpca"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := acmpca.NewCertificateAuthority(ctx, "example", &acmpca.CertificateAuthorityArgs{
//				CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
//					KeyAlgorithm:     pulumi.String("RSA_4096"),
//					SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
//					Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
//						CommonName: pulumi.String("example.com"),
//					},
//				},
//				UsageMode: pulumi.String("SHORT_LIVED_CERTIFICATE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Enable Certificate Revocation List
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/acmpca"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			acmpcaBucketAccess := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Actions: pulumi.StringArray{
//							pulumi.String("s3:GetBucketAcl"),
//							pulumi.String("s3:GetBucketLocation"),
//							pulumi.String("s3:PutObject"),
//							pulumi.String("s3:PutObjectAcl"),
//						},
//						Resources: pulumi.StringArray{
//							exampleBucketV2.Arn,
//							exampleBucketV2.Arn.ApplyT(func(arn string) (string, error) {
//								return fmt.Sprintf("%v/*", arn), nil
//							}).(pulumi.StringOutput),
//						},
//						Principals: iam.GetPolicyDocumentStatementPrincipalArray{
//							&iam.GetPolicyDocumentStatementPrincipalArgs{
//								Identifiers: pulumi.StringArray{
//									pulumi.String("acm-pca.amazonaws.com"),
//								},
//								Type: pulumi.String("Service"),
//							},
//						},
//					},
//				},
//			}, nil)
//			exampleBucketPolicy, err := s3.NewBucketPolicy(ctx, "exampleBucketPolicy", &s3.BucketPolicyArgs{
//				Bucket: exampleBucketV2.ID(),
//				Policy: acmpcaBucketAccess.ApplyT(func(acmpcaBucketAccess iam.GetPolicyDocumentResult) (*string, error) {
//					return &acmpcaBucketAccess.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = acmpca.NewCertificateAuthority(ctx, "exampleCertificateAuthority", &acmpca.CertificateAuthorityArgs{
//				CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
//					KeyAlgorithm:     pulumi.String("RSA_4096"),
//					SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
//					Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
//						CommonName: pulumi.String("example.com"),
//					},
//				},
//				RevocationConfiguration: &acmpca.CertificateAuthorityRevocationConfigurationArgs{
//					CrlConfiguration: &acmpca.CertificateAuthorityRevocationConfigurationCrlConfigurationArgs{
//						CustomCname:      pulumi.String("crl.example.com"),
//						Enabled:          pulumi.Bool(true),
//						ExpirationInDays: pulumi.Int(7),
//						S3BucketName:     exampleBucketV2.ID(),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleBucketPolicy,
//			}))
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
// `aws_acmpca_certificate_authority` can be imported by using the certificate authority ARN, e.g.,
//
// ```sh
//
//	$ pulumi import aws:acmpca/certificateAuthority:CertificateAuthority example arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
//
// ```
type CertificateAuthority struct {
	pulumi.CustomResourceState

	// ARN of the certificate authority.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfigurationOutput `pulumi:"certificateAuthorityConfiguration"`
	// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
	CertificateChain pulumi.StringOutput `pulumi:"certificateChain"`
	// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
	CertificateSigningRequest pulumi.StringOutput `pulumi:"certificateSigningRequest"`
	// Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
	KeyStorageSecurityStandard pulumi.StringOutput `pulumi:"keyStorageSecurityStandard"`
	// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotAfter pulumi.StringOutput `pulumi:"notAfter"`
	// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotBefore pulumi.StringOutput `pulumi:"notBefore"`
	// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays pulumi.IntPtrOutput `pulumi:"permanentDeletionTimeInDays"`
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration CertificateAuthorityRevocationConfigurationPtrOutput `pulumi:"revocationConfiguration"`
	// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
	Serial pulumi.StringOutput `pulumi:"serial"`
	// (**Deprecated** use the `enabled` attribute instead) Status of the certificate authority.
	//
	// Deprecated: The reported value of the "status" attribute is often inaccurate. Use the resource's "enabled" attribute to explicitly set status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
	UsageMode pulumi.StringOutput `pulumi:"usageMode"`
}

// NewCertificateAuthority registers a new resource with the given unique name, arguments, and options.
func NewCertificateAuthority(ctx *pulumi.Context,
	name string, args *CertificateAuthorityArgs, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAuthorityConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityConfiguration'")
	}
	var resource CertificateAuthority
	err := ctx.RegisterResource("aws:acmpca/certificateAuthority:CertificateAuthority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateAuthority gets an existing CertificateAuthority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateAuthorityState, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	var resource CertificateAuthority
	err := ctx.ReadResource("aws:acmpca/certificateAuthority:CertificateAuthority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateAuthority resources.
type certificateAuthorityState struct {
	// ARN of the certificate authority.
	Arn *string `pulumi:"arn"`
	// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
	Certificate *string `pulumi:"certificate"`
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration *CertificateAuthorityCertificateAuthorityConfiguration `pulumi:"certificateAuthorityConfiguration"`
	// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
	CertificateChain *string `pulumi:"certificateChain"`
	// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
	CertificateSigningRequest *string `pulumi:"certificateSigningRequest"`
	// Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
	Enabled *bool `pulumi:"enabled"`
	// Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
	KeyStorageSecurityStandard *string `pulumi:"keyStorageSecurityStandard"`
	// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotAfter *string `pulumi:"notAfter"`
	// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotBefore *string `pulumi:"notBefore"`
	// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays *int `pulumi:"permanentDeletionTimeInDays"`
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration *CertificateAuthorityRevocationConfiguration `pulumi:"revocationConfiguration"`
	// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
	Serial *string `pulumi:"serial"`
	// (**Deprecated** use the `enabled` attribute instead) Status of the certificate authority.
	//
	// Deprecated: The reported value of the "status" attribute is often inaccurate. Use the resource's "enabled" attribute to explicitly set status.
	Status *string `pulumi:"status"`
	// Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type *string `pulumi:"type"`
	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
	UsageMode *string `pulumi:"usageMode"`
}

type CertificateAuthorityState struct {
	// ARN of the certificate authority.
	Arn pulumi.StringPtrInput
	// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
	Certificate pulumi.StringPtrInput
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfigurationPtrInput
	// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
	CertificateChain pulumi.StringPtrInput
	// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
	CertificateSigningRequest pulumi.StringPtrInput
	// Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
	Enabled pulumi.BoolPtrInput
	// Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
	KeyStorageSecurityStandard pulumi.StringPtrInput
	// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotAfter pulumi.StringPtrInput
	// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotBefore pulumi.StringPtrInput
	// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays pulumi.IntPtrInput
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration CertificateAuthorityRevocationConfigurationPtrInput
	// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
	Serial pulumi.StringPtrInput
	// (**Deprecated** use the `enabled` attribute instead) Status of the certificate authority.
	//
	// Deprecated: The reported value of the "status" attribute is often inaccurate. Use the resource's "enabled" attribute to explicitly set status.
	Status pulumi.StringPtrInput
	// Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type pulumi.StringPtrInput
	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
	UsageMode pulumi.StringPtrInput
}

func (CertificateAuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityState)(nil)).Elem()
}

type certificateAuthorityArgs struct {
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfiguration `pulumi:"certificateAuthorityConfiguration"`
	// Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
	Enabled *bool `pulumi:"enabled"`
	// Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
	KeyStorageSecurityStandard *string `pulumi:"keyStorageSecurityStandard"`
	// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays *int `pulumi:"permanentDeletionTimeInDays"`
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration *CertificateAuthorityRevocationConfiguration `pulumi:"revocationConfiguration"`
	// Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type *string `pulumi:"type"`
	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
	UsageMode *string `pulumi:"usageMode"`
}

// The set of arguments for constructing a CertificateAuthority resource.
type CertificateAuthorityArgs struct {
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfigurationInput
	// Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
	Enabled pulumi.BoolPtrInput
	// Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
	KeyStorageSecurityStandard pulumi.StringPtrInput
	// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays pulumi.IntPtrInput
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration CertificateAuthorityRevocationConfigurationPtrInput
	// Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type pulumi.StringPtrInput
	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
	UsageMode pulumi.StringPtrInput
}

func (CertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityArgs)(nil)).Elem()
}

type CertificateAuthorityInput interface {
	pulumi.Input

	ToCertificateAuthorityOutput() CertificateAuthorityOutput
	ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput
}

func (*CertificateAuthority) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthority)(nil)).Elem()
}

func (i *CertificateAuthority) ToCertificateAuthorityOutput() CertificateAuthorityOutput {
	return i.ToCertificateAuthorityOutputWithContext(context.Background())
}

func (i *CertificateAuthority) ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityOutput)
}

// CertificateAuthorityArrayInput is an input type that accepts CertificateAuthorityArray and CertificateAuthorityArrayOutput values.
// You can construct a concrete instance of `CertificateAuthorityArrayInput` via:
//
//	CertificateAuthorityArray{ CertificateAuthorityArgs{...} }
type CertificateAuthorityArrayInput interface {
	pulumi.Input

	ToCertificateAuthorityArrayOutput() CertificateAuthorityArrayOutput
	ToCertificateAuthorityArrayOutputWithContext(context.Context) CertificateAuthorityArrayOutput
}

type CertificateAuthorityArray []CertificateAuthorityInput

func (CertificateAuthorityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateAuthority)(nil)).Elem()
}

func (i CertificateAuthorityArray) ToCertificateAuthorityArrayOutput() CertificateAuthorityArrayOutput {
	return i.ToCertificateAuthorityArrayOutputWithContext(context.Background())
}

func (i CertificateAuthorityArray) ToCertificateAuthorityArrayOutputWithContext(ctx context.Context) CertificateAuthorityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityArrayOutput)
}

// CertificateAuthorityMapInput is an input type that accepts CertificateAuthorityMap and CertificateAuthorityMapOutput values.
// You can construct a concrete instance of `CertificateAuthorityMapInput` via:
//
//	CertificateAuthorityMap{ "key": CertificateAuthorityArgs{...} }
type CertificateAuthorityMapInput interface {
	pulumi.Input

	ToCertificateAuthorityMapOutput() CertificateAuthorityMapOutput
	ToCertificateAuthorityMapOutputWithContext(context.Context) CertificateAuthorityMapOutput
}

type CertificateAuthorityMap map[string]CertificateAuthorityInput

func (CertificateAuthorityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateAuthority)(nil)).Elem()
}

func (i CertificateAuthorityMap) ToCertificateAuthorityMapOutput() CertificateAuthorityMapOutput {
	return i.ToCertificateAuthorityMapOutputWithContext(context.Background())
}

func (i CertificateAuthorityMap) ToCertificateAuthorityMapOutputWithContext(ctx context.Context) CertificateAuthorityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityMapOutput)
}

type CertificateAuthorityOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthority)(nil)).Elem()
}

func (o CertificateAuthorityOutput) ToCertificateAuthorityOutput() CertificateAuthorityOutput {
	return o
}

func (o CertificateAuthorityOutput) ToCertificateAuthorityOutputWithContext(ctx context.Context) CertificateAuthorityOutput {
	return o
}

// ARN of the certificate authority.
func (o CertificateAuthorityOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
func (o CertificateAuthorityOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Nested argument containing algorithms and certificate subject information. Defined below.
func (o CertificateAuthorityOutput) CertificateAuthorityConfiguration() CertificateAuthorityCertificateAuthorityConfigurationOutput {
	return o.ApplyT(func(v *CertificateAuthority) CertificateAuthorityCertificateAuthorityConfigurationOutput {
		return v.CertificateAuthorityConfiguration
	}).(CertificateAuthorityCertificateAuthorityConfigurationOutput)
}

// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
func (o CertificateAuthorityOutput) CertificateChain() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.CertificateChain }).(pulumi.StringOutput)
}

// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
func (o CertificateAuthorityOutput) CertificateSigningRequest() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.CertificateSigningRequest }).(pulumi.StringOutput)
}

// Whether the certificate authority is enabled or disabled. Defaults to `true`. Can only be disabled if the CA is in an `ACTIVE` state.
func (o CertificateAuthorityOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Cryptographic key management compliance standard used for handling CA keys. Defaults to `FIPS_140_2_LEVEL_3_OR_HIGHER`. Valid values: `FIPS_140_2_LEVEL_3_OR_HIGHER` and `FIPS_140_2_LEVEL_2_OR_HIGHER`. Supported standard for each region can be found in the [Storage and security compliance of AWS Private CA private keys Documentation](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
func (o CertificateAuthorityOutput) KeyStorageSecurityStandard() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.KeyStorageSecurityStandard }).(pulumi.StringOutput)
}

// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
func (o CertificateAuthorityOutput) NotAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.NotAfter }).(pulumi.StringOutput)
}

// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
func (o CertificateAuthorityOutput) NotBefore() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.NotBefore }).(pulumi.StringOutput)
}

// Number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
func (o CertificateAuthorityOutput) PermanentDeletionTimeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.IntPtrOutput { return v.PermanentDeletionTimeInDays }).(pulumi.IntPtrOutput)
}

// Nested argument containing revocation configuration. Defined below.
func (o CertificateAuthorityOutput) RevocationConfiguration() CertificateAuthorityRevocationConfigurationPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) CertificateAuthorityRevocationConfigurationPtrOutput {
		return v.RevocationConfiguration
	}).(CertificateAuthorityRevocationConfigurationPtrOutput)
}

// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
func (o CertificateAuthorityOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Serial }).(pulumi.StringOutput)
}

// (**Deprecated** use the `enabled` attribute instead) Status of the certificate authority.
//
// Deprecated: The reported value of the "status" attribute is often inaccurate. Use the resource's "enabled" attribute to explicitly set status.
func (o CertificateAuthorityOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Key-value map of user-defined tags that are attached to the certificate authority. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CertificateAuthorityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o CertificateAuthorityOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
func (o CertificateAuthorityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. Defaults to `GENERAL_PURPOSE`. Valid values: `GENERAL_PURPOSE` and `SHORT_LIVED_CERTIFICATE`.
func (o CertificateAuthorityOutput) UsageMode() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthority) pulumi.StringOutput { return v.UsageMode }).(pulumi.StringOutput)
}

type CertificateAuthorityArrayOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateAuthority)(nil)).Elem()
}

func (o CertificateAuthorityArrayOutput) ToCertificateAuthorityArrayOutput() CertificateAuthorityArrayOutput {
	return o
}

func (o CertificateAuthorityArrayOutput) ToCertificateAuthorityArrayOutputWithContext(ctx context.Context) CertificateAuthorityArrayOutput {
	return o
}

func (o CertificateAuthorityArrayOutput) Index(i pulumi.IntInput) CertificateAuthorityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateAuthority {
		return vs[0].([]*CertificateAuthority)[vs[1].(int)]
	}).(CertificateAuthorityOutput)
}

type CertificateAuthorityMapOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateAuthority)(nil)).Elem()
}

func (o CertificateAuthorityMapOutput) ToCertificateAuthorityMapOutput() CertificateAuthorityMapOutput {
	return o
}

func (o CertificateAuthorityMapOutput) ToCertificateAuthorityMapOutputWithContext(ctx context.Context) CertificateAuthorityMapOutput {
	return o
}

func (o CertificateAuthorityMapOutput) MapIndex(k pulumi.StringInput) CertificateAuthorityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateAuthority {
		return vs[0].(map[string]*CertificateAuthority)[vs[1].(string)]
	}).(CertificateAuthorityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateAuthorityInput)(nil)).Elem(), &CertificateAuthority{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateAuthorityArrayInput)(nil)).Elem(), CertificateAuthorityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateAuthorityMapInput)(nil)).Elem(), CertificateAuthorityMap{})
	pulumi.RegisterOutputType(CertificateAuthorityOutput{})
	pulumi.RegisterOutputType(CertificateAuthorityArrayOutput{})
	pulumi.RegisterOutputType(CertificateAuthorityMapOutput{})
}
