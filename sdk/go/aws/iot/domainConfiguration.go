// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages an AWS IoT domain configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iot.NewDomainConfiguration(ctx, "iot", &iot.DomainConfigurationArgs{
//				DomainName:  pulumi.String("iot.example.com"),
//				ServiceType: pulumi.String("DATA"),
//				ServerCertificateArns: pulumi.StringArray{
//					aws_acm_certificate.Cert.Arn,
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
// Using `pulumi import`, import domain configurations using the name. For example:
//
// ```sh
//
//	$ pulumi import aws:iot/domainConfiguration:DomainConfiguration example example
//
// ```
type DomainConfiguration struct {
	pulumi.CustomResourceState

	// The ARN of the domain configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// An object that specifies the authorization service for a domain. See below.
	AuthorizerConfig DomainConfigurationAuthorizerConfigPtrOutput `pulumi:"authorizerConfig"`
	// Fully-qualified domain name.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// The type of the domain.
	DomainType pulumi.StringOutput `pulumi:"domainType"`
	// The name of the domain configuration. This value must be unique to a region.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARNs of the certificates that IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for Amazon Web Services-managed domains. When using a custom `domainName`, the cert must include it.
	ServerCertificateArns pulumi.StringArrayOutput `pulumi:"serverCertificateArns"`
	// The type of service delivered by the endpoint. Note: Amazon Web Services IoT Core currently supports only the `DATA` service type.
	ServiceType pulumi.StringPtrOutput `pulumi:"serviceType"`
	Status      pulumi.StringPtrOutput `pulumi:"status"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// An object that specifies the TLS configuration for a domain. See below.
	TlsConfig DomainConfigurationTlsConfigOutput `pulumi:"tlsConfig"`
	// The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for Amazon Web Services-managed domains.
	ValidationCertificateArn pulumi.StringPtrOutput `pulumi:"validationCertificateArn"`
}

// NewDomainConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDomainConfiguration(ctx *pulumi.Context,
	name string, args *DomainConfigurationArgs, opts ...pulumi.ResourceOption) (*DomainConfiguration, error) {
	if args == nil {
		args = &DomainConfigurationArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainConfiguration
	err := ctx.RegisterResource("aws:iot/domainConfiguration:DomainConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainConfiguration gets an existing DomainConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainConfigurationState, opts ...pulumi.ResourceOption) (*DomainConfiguration, error) {
	var resource DomainConfiguration
	err := ctx.ReadResource("aws:iot/domainConfiguration:DomainConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainConfiguration resources.
type domainConfigurationState struct {
	// The ARN of the domain configuration.
	Arn *string `pulumi:"arn"`
	// An object that specifies the authorization service for a domain. See below.
	AuthorizerConfig *DomainConfigurationAuthorizerConfig `pulumi:"authorizerConfig"`
	// Fully-qualified domain name.
	DomainName *string `pulumi:"domainName"`
	// The type of the domain.
	DomainType *string `pulumi:"domainType"`
	// The name of the domain configuration. This value must be unique to a region.
	Name *string `pulumi:"name"`
	// The ARNs of the certificates that IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for Amazon Web Services-managed domains. When using a custom `domainName`, the cert must include it.
	ServerCertificateArns []string `pulumi:"serverCertificateArns"`
	// The type of service delivered by the endpoint. Note: Amazon Web Services IoT Core currently supports only the `DATA` service type.
	ServiceType *string `pulumi:"serviceType"`
	Status      *string `pulumi:"status"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// An object that specifies the TLS configuration for a domain. See below.
	TlsConfig *DomainConfigurationTlsConfig `pulumi:"tlsConfig"`
	// The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for Amazon Web Services-managed domains.
	ValidationCertificateArn *string `pulumi:"validationCertificateArn"`
}

type DomainConfigurationState struct {
	// The ARN of the domain configuration.
	Arn pulumi.StringPtrInput
	// An object that specifies the authorization service for a domain. See below.
	AuthorizerConfig DomainConfigurationAuthorizerConfigPtrInput
	// Fully-qualified domain name.
	DomainName pulumi.StringPtrInput
	// The type of the domain.
	DomainType pulumi.StringPtrInput
	// The name of the domain configuration. This value must be unique to a region.
	Name pulumi.StringPtrInput
	// The ARNs of the certificates that IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for Amazon Web Services-managed domains. When using a custom `domainName`, the cert must include it.
	ServerCertificateArns pulumi.StringArrayInput
	// The type of service delivered by the endpoint. Note: Amazon Web Services IoT Core currently supports only the `DATA` service type.
	ServiceType pulumi.StringPtrInput
	Status      pulumi.StringPtrInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// An object that specifies the TLS configuration for a domain. See below.
	TlsConfig DomainConfigurationTlsConfigPtrInput
	// The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for Amazon Web Services-managed domains.
	ValidationCertificateArn pulumi.StringPtrInput
}

func (DomainConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainConfigurationState)(nil)).Elem()
}

type domainConfigurationArgs struct {
	// An object that specifies the authorization service for a domain. See below.
	AuthorizerConfig *DomainConfigurationAuthorizerConfig `pulumi:"authorizerConfig"`
	// Fully-qualified domain name.
	DomainName *string `pulumi:"domainName"`
	// The name of the domain configuration. This value must be unique to a region.
	Name *string `pulumi:"name"`
	// The ARNs of the certificates that IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for Amazon Web Services-managed domains. When using a custom `domainName`, the cert must include it.
	ServerCertificateArns []string `pulumi:"serverCertificateArns"`
	// The type of service delivered by the endpoint. Note: Amazon Web Services IoT Core currently supports only the `DATA` service type.
	ServiceType *string `pulumi:"serviceType"`
	Status      *string `pulumi:"status"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// An object that specifies the TLS configuration for a domain. See below.
	TlsConfig *DomainConfigurationTlsConfig `pulumi:"tlsConfig"`
	// The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for Amazon Web Services-managed domains.
	ValidationCertificateArn *string `pulumi:"validationCertificateArn"`
}

// The set of arguments for constructing a DomainConfiguration resource.
type DomainConfigurationArgs struct {
	// An object that specifies the authorization service for a domain. See below.
	AuthorizerConfig DomainConfigurationAuthorizerConfigPtrInput
	// Fully-qualified domain name.
	DomainName pulumi.StringPtrInput
	// The name of the domain configuration. This value must be unique to a region.
	Name pulumi.StringPtrInput
	// The ARNs of the certificates that IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for Amazon Web Services-managed domains. When using a custom `domainName`, the cert must include it.
	ServerCertificateArns pulumi.StringArrayInput
	// The type of service delivered by the endpoint. Note: Amazon Web Services IoT Core currently supports only the `DATA` service type.
	ServiceType pulumi.StringPtrInput
	Status      pulumi.StringPtrInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// An object that specifies the TLS configuration for a domain. See below.
	TlsConfig DomainConfigurationTlsConfigPtrInput
	// The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for Amazon Web Services-managed domains.
	ValidationCertificateArn pulumi.StringPtrInput
}

func (DomainConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainConfigurationArgs)(nil)).Elem()
}

type DomainConfigurationInput interface {
	pulumi.Input

	ToDomainConfigurationOutput() DomainConfigurationOutput
	ToDomainConfigurationOutputWithContext(ctx context.Context) DomainConfigurationOutput
}

func (*DomainConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainConfiguration)(nil)).Elem()
}

func (i *DomainConfiguration) ToDomainConfigurationOutput() DomainConfigurationOutput {
	return i.ToDomainConfigurationOutputWithContext(context.Background())
}

func (i *DomainConfiguration) ToDomainConfigurationOutputWithContext(ctx context.Context) DomainConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigurationOutput)
}

// DomainConfigurationArrayInput is an input type that accepts DomainConfigurationArray and DomainConfigurationArrayOutput values.
// You can construct a concrete instance of `DomainConfigurationArrayInput` via:
//
//	DomainConfigurationArray{ DomainConfigurationArgs{...} }
type DomainConfigurationArrayInput interface {
	pulumi.Input

	ToDomainConfigurationArrayOutput() DomainConfigurationArrayOutput
	ToDomainConfigurationArrayOutputWithContext(context.Context) DomainConfigurationArrayOutput
}

type DomainConfigurationArray []DomainConfigurationInput

func (DomainConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainConfiguration)(nil)).Elem()
}

func (i DomainConfigurationArray) ToDomainConfigurationArrayOutput() DomainConfigurationArrayOutput {
	return i.ToDomainConfigurationArrayOutputWithContext(context.Background())
}

func (i DomainConfigurationArray) ToDomainConfigurationArrayOutputWithContext(ctx context.Context) DomainConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigurationArrayOutput)
}

// DomainConfigurationMapInput is an input type that accepts DomainConfigurationMap and DomainConfigurationMapOutput values.
// You can construct a concrete instance of `DomainConfigurationMapInput` via:
//
//	DomainConfigurationMap{ "key": DomainConfigurationArgs{...} }
type DomainConfigurationMapInput interface {
	pulumi.Input

	ToDomainConfigurationMapOutput() DomainConfigurationMapOutput
	ToDomainConfigurationMapOutputWithContext(context.Context) DomainConfigurationMapOutput
}

type DomainConfigurationMap map[string]DomainConfigurationInput

func (DomainConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainConfiguration)(nil)).Elem()
}

func (i DomainConfigurationMap) ToDomainConfigurationMapOutput() DomainConfigurationMapOutput {
	return i.ToDomainConfigurationMapOutputWithContext(context.Background())
}

func (i DomainConfigurationMap) ToDomainConfigurationMapOutputWithContext(ctx context.Context) DomainConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigurationMapOutput)
}

type DomainConfigurationOutput struct{ *pulumi.OutputState }

func (DomainConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainConfiguration)(nil)).Elem()
}

func (o DomainConfigurationOutput) ToDomainConfigurationOutput() DomainConfigurationOutput {
	return o
}

func (o DomainConfigurationOutput) ToDomainConfigurationOutputWithContext(ctx context.Context) DomainConfigurationOutput {
	return o
}

// The ARN of the domain configuration.
func (o DomainConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// An object that specifies the authorization service for a domain. See below.
func (o DomainConfigurationOutput) AuthorizerConfig() DomainConfigurationAuthorizerConfigPtrOutput {
	return o.ApplyT(func(v *DomainConfiguration) DomainConfigurationAuthorizerConfigPtrOutput { return v.AuthorizerConfig }).(DomainConfigurationAuthorizerConfigPtrOutput)
}

// Fully-qualified domain name.
func (o DomainConfigurationOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The type of the domain.
func (o DomainConfigurationOutput) DomainType() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringOutput { return v.DomainType }).(pulumi.StringOutput)
}

// The name of the domain configuration. This value must be unique to a region.
func (o DomainConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARNs of the certificates that IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for Amazon Web Services-managed domains. When using a custom `domainName`, the cert must include it.
func (o DomainConfigurationOutput) ServerCertificateArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringArrayOutput { return v.ServerCertificateArns }).(pulumi.StringArrayOutput)
}

// The type of service delivered by the endpoint. Note: Amazon Web Services IoT Core currently supports only the `DATA` service type.
func (o DomainConfigurationOutput) ServiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringPtrOutput { return v.ServiceType }).(pulumi.StringPtrOutput)
}

func (o DomainConfigurationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DomainConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o DomainConfigurationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// An object that specifies the TLS configuration for a domain. See below.
func (o DomainConfigurationOutput) TlsConfig() DomainConfigurationTlsConfigOutput {
	return o.ApplyT(func(v *DomainConfiguration) DomainConfigurationTlsConfigOutput { return v.TlsConfig }).(DomainConfigurationTlsConfigOutput)
}

// The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for Amazon Web Services-managed domains.
func (o DomainConfigurationOutput) ValidationCertificateArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainConfiguration) pulumi.StringPtrOutput { return v.ValidationCertificateArn }).(pulumi.StringPtrOutput)
}

type DomainConfigurationArrayOutput struct{ *pulumi.OutputState }

func (DomainConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainConfiguration)(nil)).Elem()
}

func (o DomainConfigurationArrayOutput) ToDomainConfigurationArrayOutput() DomainConfigurationArrayOutput {
	return o
}

func (o DomainConfigurationArrayOutput) ToDomainConfigurationArrayOutputWithContext(ctx context.Context) DomainConfigurationArrayOutput {
	return o
}

func (o DomainConfigurationArrayOutput) Index(i pulumi.IntInput) DomainConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainConfiguration {
		return vs[0].([]*DomainConfiguration)[vs[1].(int)]
	}).(DomainConfigurationOutput)
}

type DomainConfigurationMapOutput struct{ *pulumi.OutputState }

func (DomainConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainConfiguration)(nil)).Elem()
}

func (o DomainConfigurationMapOutput) ToDomainConfigurationMapOutput() DomainConfigurationMapOutput {
	return o
}

func (o DomainConfigurationMapOutput) ToDomainConfigurationMapOutputWithContext(ctx context.Context) DomainConfigurationMapOutput {
	return o
}

func (o DomainConfigurationMapOutput) MapIndex(k pulumi.StringInput) DomainConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainConfiguration {
		return vs[0].(map[string]*DomainConfiguration)[vs[1].(string)]
	}).(DomainConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigurationInput)(nil)).Elem(), &DomainConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigurationArrayInput)(nil)).Elem(), DomainConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigurationMapInput)(nil)).Elem(), DomainConfigurationMap{})
	pulumi.RegisterOutputType(DomainConfigurationOutput{})
	pulumi.RegisterOutputType(DomainConfigurationArrayOutput{})
	pulumi.RegisterOutputType(DomainConfigurationMapOutput{})
}
