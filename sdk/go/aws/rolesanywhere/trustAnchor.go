// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rolesanywhere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing a Roles Anywhere Trust Anchor.
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
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/acmpca"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rolesanywhere"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "exampleCertificateAuthority", &acmpca.CertificateAuthorityArgs{
//				PermanentDeletionTimeInDays: pulumi.Int(7),
//				Type:                        pulumi.String("ROOT"),
//				CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
//					KeyAlgorithm:     pulumi.String("RSA_4096"),
//					SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
//					Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
//						CommonName: pulumi.String("example.com"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			current, err := aws.GetPartition(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = acmpca.NewCertificate(ctx, "testCertificate", &acmpca.CertificateArgs{
//				CertificateAuthorityArn:   exampleCertificateAuthority.Arn,
//				CertificateSigningRequest: exampleCertificateAuthority.CertificateSigningRequest,
//				SigningAlgorithm:          pulumi.String("SHA512WITHRSA"),
//				TemplateArn:               pulumi.String(fmt.Sprintf("arn:%v:acm-pca:::template/RootCACertificate/V1", current.Partition)),
//				Validity: &acmpca.CertificateValidityArgs{
//					Type:  pulumi.String("YEARS"),
//					Value: pulumi.String("1"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleCertificateAuthorityCertificate, err := acmpca.NewCertificateAuthorityCertificate(ctx, "exampleCertificateAuthorityCertificate", &acmpca.CertificateAuthorityCertificateArgs{
//				CertificateAuthorityArn: exampleCertificateAuthority.Arn,
//				Certificate:             pulumi.Any(aws_acmpca_certificate.Example.Certificate),
//				CertificateChain:        pulumi.Any(aws_acmpca_certificate.Example.Certificate_chain),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rolesanywhere.NewTrustAnchor(ctx, "testTrustAnchor", &rolesanywhere.TrustAnchorArgs{
//				Source: &rolesanywhere.TrustAnchorSourceArgs{
//					SourceData: &rolesanywhere.TrustAnchorSourceSourceDataArgs{
//						AcmPcaArn: exampleCertificateAuthority.Arn,
//					},
//					SourceType: pulumi.String("AWS_ACM_PCA"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleCertificateAuthorityCertificate,
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
// `aws_rolesanywhere_trust_anchor` can be imported using its `id`, e.g.
//
// ```sh
//
//	$ pulumi import aws:rolesanywhere/trustAnchor:TrustAnchor example 92b2fbbb-984d-41a3-a765-e3cbdb69ebb1
//
// ```
type TrustAnchor struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Trust Anchor
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether or not the Trust Anchor should be enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The name of the Trust Anchor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The source of trust, documented below
	Source TrustAnchorSourceOutput `pulumi:"source"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTrustAnchor registers a new resource with the given unique name, arguments, and options.
func NewTrustAnchor(ctx *pulumi.Context,
	name string, args *TrustAnchorArgs, opts ...pulumi.ResourceOption) (*TrustAnchor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource TrustAnchor
	err := ctx.RegisterResource("aws:rolesanywhere/trustAnchor:TrustAnchor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrustAnchor gets an existing TrustAnchor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrustAnchor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrustAnchorState, opts ...pulumi.ResourceOption) (*TrustAnchor, error) {
	var resource TrustAnchor
	err := ctx.ReadResource("aws:rolesanywhere/trustAnchor:TrustAnchor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrustAnchor resources.
type trustAnchorState struct {
	// Amazon Resource Name (ARN) of the Trust Anchor
	Arn *string `pulumi:"arn"`
	// Whether or not the Trust Anchor should be enabled.
	Enabled *bool `pulumi:"enabled"`
	// The name of the Trust Anchor.
	Name *string `pulumi:"name"`
	// The source of trust, documented below
	Source *TrustAnchorSource `pulumi:"source"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type TrustAnchorState struct {
	// Amazon Resource Name (ARN) of the Trust Anchor
	Arn pulumi.StringPtrInput
	// Whether or not the Trust Anchor should be enabled.
	Enabled pulumi.BoolPtrInput
	// The name of the Trust Anchor.
	Name pulumi.StringPtrInput
	// The source of trust, documented below
	Source TrustAnchorSourcePtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (TrustAnchorState) ElementType() reflect.Type {
	return reflect.TypeOf((*trustAnchorState)(nil)).Elem()
}

type trustAnchorArgs struct {
	// Whether or not the Trust Anchor should be enabled.
	Enabled *bool `pulumi:"enabled"`
	// The name of the Trust Anchor.
	Name *string `pulumi:"name"`
	// The source of trust, documented below
	Source TrustAnchorSource `pulumi:"source"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a TrustAnchor resource.
type TrustAnchorArgs struct {
	// Whether or not the Trust Anchor should be enabled.
	Enabled pulumi.BoolPtrInput
	// The name of the Trust Anchor.
	Name pulumi.StringPtrInput
	// The source of trust, documented below
	Source TrustAnchorSourceInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (TrustAnchorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trustAnchorArgs)(nil)).Elem()
}

type TrustAnchorInput interface {
	pulumi.Input

	ToTrustAnchorOutput() TrustAnchorOutput
	ToTrustAnchorOutputWithContext(ctx context.Context) TrustAnchorOutput
}

func (*TrustAnchor) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchor)(nil)).Elem()
}

func (i *TrustAnchor) ToTrustAnchorOutput() TrustAnchorOutput {
	return i.ToTrustAnchorOutputWithContext(context.Background())
}

func (i *TrustAnchor) ToTrustAnchorOutputWithContext(ctx context.Context) TrustAnchorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorOutput)
}

// TrustAnchorArrayInput is an input type that accepts TrustAnchorArray and TrustAnchorArrayOutput values.
// You can construct a concrete instance of `TrustAnchorArrayInput` via:
//
//	TrustAnchorArray{ TrustAnchorArgs{...} }
type TrustAnchorArrayInput interface {
	pulumi.Input

	ToTrustAnchorArrayOutput() TrustAnchorArrayOutput
	ToTrustAnchorArrayOutputWithContext(context.Context) TrustAnchorArrayOutput
}

type TrustAnchorArray []TrustAnchorInput

func (TrustAnchorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrustAnchor)(nil)).Elem()
}

func (i TrustAnchorArray) ToTrustAnchorArrayOutput() TrustAnchorArrayOutput {
	return i.ToTrustAnchorArrayOutputWithContext(context.Background())
}

func (i TrustAnchorArray) ToTrustAnchorArrayOutputWithContext(ctx context.Context) TrustAnchorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorArrayOutput)
}

// TrustAnchorMapInput is an input type that accepts TrustAnchorMap and TrustAnchorMapOutput values.
// You can construct a concrete instance of `TrustAnchorMapInput` via:
//
//	TrustAnchorMap{ "key": TrustAnchorArgs{...} }
type TrustAnchorMapInput interface {
	pulumi.Input

	ToTrustAnchorMapOutput() TrustAnchorMapOutput
	ToTrustAnchorMapOutputWithContext(context.Context) TrustAnchorMapOutput
}

type TrustAnchorMap map[string]TrustAnchorInput

func (TrustAnchorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrustAnchor)(nil)).Elem()
}

func (i TrustAnchorMap) ToTrustAnchorMapOutput() TrustAnchorMapOutput {
	return i.ToTrustAnchorMapOutputWithContext(context.Background())
}

func (i TrustAnchorMap) ToTrustAnchorMapOutputWithContext(ctx context.Context) TrustAnchorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorMapOutput)
}

type TrustAnchorOutput struct{ *pulumi.OutputState }

func (TrustAnchorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchor)(nil)).Elem()
}

func (o TrustAnchorOutput) ToTrustAnchorOutput() TrustAnchorOutput {
	return o
}

func (o TrustAnchorOutput) ToTrustAnchorOutputWithContext(ctx context.Context) TrustAnchorOutput {
	return o
}

// Amazon Resource Name (ARN) of the Trust Anchor
func (o TrustAnchorOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Whether or not the Trust Anchor should be enabled.
func (o TrustAnchorOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The name of the Trust Anchor.
func (o TrustAnchorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The source of trust, documented below
func (o TrustAnchorOutput) Source() TrustAnchorSourceOutput {
	return o.ApplyT(func(v *TrustAnchor) TrustAnchorSourceOutput { return v.Source }).(TrustAnchorSourceOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TrustAnchorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o TrustAnchorOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type TrustAnchorArrayOutput struct{ *pulumi.OutputState }

func (TrustAnchorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrustAnchor)(nil)).Elem()
}

func (o TrustAnchorArrayOutput) ToTrustAnchorArrayOutput() TrustAnchorArrayOutput {
	return o
}

func (o TrustAnchorArrayOutput) ToTrustAnchorArrayOutputWithContext(ctx context.Context) TrustAnchorArrayOutput {
	return o
}

func (o TrustAnchorArrayOutput) Index(i pulumi.IntInput) TrustAnchorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrustAnchor {
		return vs[0].([]*TrustAnchor)[vs[1].(int)]
	}).(TrustAnchorOutput)
}

type TrustAnchorMapOutput struct{ *pulumi.OutputState }

func (TrustAnchorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrustAnchor)(nil)).Elem()
}

func (o TrustAnchorMapOutput) ToTrustAnchorMapOutput() TrustAnchorMapOutput {
	return o
}

func (o TrustAnchorMapOutput) ToTrustAnchorMapOutputWithContext(ctx context.Context) TrustAnchorMapOutput {
	return o
}

func (o TrustAnchorMapOutput) MapIndex(k pulumi.StringInput) TrustAnchorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrustAnchor {
		return vs[0].(map[string]*TrustAnchor)[vs[1].(string)]
	}).(TrustAnchorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorInput)(nil)).Elem(), &TrustAnchor{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorArrayInput)(nil)).Elem(), TrustAnchorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorMapInput)(nil)).Elem(), TrustAnchorMap{})
	pulumi.RegisterOutputType(TrustAnchorOutput{})
	pulumi.RegisterOutputType(TrustAnchorArrayOutput{})
	pulumi.RegisterOutputType(TrustAnchorMapOutput{})
}
