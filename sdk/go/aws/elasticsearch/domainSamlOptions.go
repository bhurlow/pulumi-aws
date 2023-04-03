// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages SAML authentication options for an AWS Elasticsearch Domain.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticsearch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDomain, err := elasticsearch.NewDomain(ctx, "exampleDomain", &elasticsearch.DomainArgs{
//				ElasticsearchVersion: pulumi.String("1.5"),
//				ClusterConfig: &elasticsearch.DomainClusterConfigArgs{
//					InstanceType: pulumi.String("r4.large.elasticsearch"),
//				},
//				SnapshotOptions: &elasticsearch.DomainSnapshotOptionsArgs{
//					AutomatedSnapshotStartHour: pulumi.Int(23),
//				},
//				Tags: pulumi.StringMap{
//					"Domain": pulumi.String("TestDomain"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticsearch.NewDomainSamlOptions(ctx, "exampleDomainSamlOptions", &elasticsearch.DomainSamlOptionsArgs{
//				DomainName: exampleDomain.DomainName,
//				SamlOptions: &elasticsearch.DomainSamlOptionsSamlOptionsArgs{
//					Enabled: pulumi.Bool(true),
//					Idp: &elasticsearch.DomainSamlOptionsSamlOptionsIdpArgs{
//						EntityId:        pulumi.String("https://example.com"),
//						MetadataContent: readFileOrPanic("./saml-metadata.xml"),
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
//
// ## Import
//
// Elasticsearch domains can be imported using the `domain_name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:elasticsearch/domainSamlOptions:DomainSamlOptions example domain_name
//
// ```
type DomainSamlOptions struct {
	pulumi.CustomResourceState

	// Name of the domain.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The SAML authentication options for an AWS Elasticsearch Domain.
	SamlOptions DomainSamlOptionsSamlOptionsPtrOutput `pulumi:"samlOptions"`
}

// NewDomainSamlOptions registers a new resource with the given unique name, arguments, and options.
func NewDomainSamlOptions(ctx *pulumi.Context,
	name string, args *DomainSamlOptionsArgs, opts ...pulumi.ResourceOption) (*DomainSamlOptions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource DomainSamlOptions
	err := ctx.RegisterResource("aws:elasticsearch/domainSamlOptions:DomainSamlOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainSamlOptions gets an existing DomainSamlOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainSamlOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainSamlOptionsState, opts ...pulumi.ResourceOption) (*DomainSamlOptions, error) {
	var resource DomainSamlOptions
	err := ctx.ReadResource("aws:elasticsearch/domainSamlOptions:DomainSamlOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainSamlOptions resources.
type domainSamlOptionsState struct {
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// The SAML authentication options for an AWS Elasticsearch Domain.
	SamlOptions *DomainSamlOptionsSamlOptions `pulumi:"samlOptions"`
}

type DomainSamlOptionsState struct {
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// The SAML authentication options for an AWS Elasticsearch Domain.
	SamlOptions DomainSamlOptionsSamlOptionsPtrInput
}

func (DomainSamlOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainSamlOptionsState)(nil)).Elem()
}

type domainSamlOptionsArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// The SAML authentication options for an AWS Elasticsearch Domain.
	SamlOptions *DomainSamlOptionsSamlOptions `pulumi:"samlOptions"`
}

// The set of arguments for constructing a DomainSamlOptions resource.
type DomainSamlOptionsArgs struct {
	// Name of the domain.
	DomainName pulumi.StringInput
	// The SAML authentication options for an AWS Elasticsearch Domain.
	SamlOptions DomainSamlOptionsSamlOptionsPtrInput
}

func (DomainSamlOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainSamlOptionsArgs)(nil)).Elem()
}

type DomainSamlOptionsInput interface {
	pulumi.Input

	ToDomainSamlOptionsOutput() DomainSamlOptionsOutput
	ToDomainSamlOptionsOutputWithContext(ctx context.Context) DomainSamlOptionsOutput
}

func (*DomainSamlOptions) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainSamlOptions)(nil)).Elem()
}

func (i *DomainSamlOptions) ToDomainSamlOptionsOutput() DomainSamlOptionsOutput {
	return i.ToDomainSamlOptionsOutputWithContext(context.Background())
}

func (i *DomainSamlOptions) ToDomainSamlOptionsOutputWithContext(ctx context.Context) DomainSamlOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSamlOptionsOutput)
}

// DomainSamlOptionsArrayInput is an input type that accepts DomainSamlOptionsArray and DomainSamlOptionsArrayOutput values.
// You can construct a concrete instance of `DomainSamlOptionsArrayInput` via:
//
//	DomainSamlOptionsArray{ DomainSamlOptionsArgs{...} }
type DomainSamlOptionsArrayInput interface {
	pulumi.Input

	ToDomainSamlOptionsArrayOutput() DomainSamlOptionsArrayOutput
	ToDomainSamlOptionsArrayOutputWithContext(context.Context) DomainSamlOptionsArrayOutput
}

type DomainSamlOptionsArray []DomainSamlOptionsInput

func (DomainSamlOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainSamlOptions)(nil)).Elem()
}

func (i DomainSamlOptionsArray) ToDomainSamlOptionsArrayOutput() DomainSamlOptionsArrayOutput {
	return i.ToDomainSamlOptionsArrayOutputWithContext(context.Background())
}

func (i DomainSamlOptionsArray) ToDomainSamlOptionsArrayOutputWithContext(ctx context.Context) DomainSamlOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSamlOptionsArrayOutput)
}

// DomainSamlOptionsMapInput is an input type that accepts DomainSamlOptionsMap and DomainSamlOptionsMapOutput values.
// You can construct a concrete instance of `DomainSamlOptionsMapInput` via:
//
//	DomainSamlOptionsMap{ "key": DomainSamlOptionsArgs{...} }
type DomainSamlOptionsMapInput interface {
	pulumi.Input

	ToDomainSamlOptionsMapOutput() DomainSamlOptionsMapOutput
	ToDomainSamlOptionsMapOutputWithContext(context.Context) DomainSamlOptionsMapOutput
}

type DomainSamlOptionsMap map[string]DomainSamlOptionsInput

func (DomainSamlOptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainSamlOptions)(nil)).Elem()
}

func (i DomainSamlOptionsMap) ToDomainSamlOptionsMapOutput() DomainSamlOptionsMapOutput {
	return i.ToDomainSamlOptionsMapOutputWithContext(context.Background())
}

func (i DomainSamlOptionsMap) ToDomainSamlOptionsMapOutputWithContext(ctx context.Context) DomainSamlOptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSamlOptionsMapOutput)
}

type DomainSamlOptionsOutput struct{ *pulumi.OutputState }

func (DomainSamlOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainSamlOptions)(nil)).Elem()
}

func (o DomainSamlOptionsOutput) ToDomainSamlOptionsOutput() DomainSamlOptionsOutput {
	return o
}

func (o DomainSamlOptionsOutput) ToDomainSamlOptionsOutputWithContext(ctx context.Context) DomainSamlOptionsOutput {
	return o
}

// Name of the domain.
func (o DomainSamlOptionsOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainSamlOptions) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The SAML authentication options for an AWS Elasticsearch Domain.
func (o DomainSamlOptionsOutput) SamlOptions() DomainSamlOptionsSamlOptionsPtrOutput {
	return o.ApplyT(func(v *DomainSamlOptions) DomainSamlOptionsSamlOptionsPtrOutput { return v.SamlOptions }).(DomainSamlOptionsSamlOptionsPtrOutput)
}

type DomainSamlOptionsArrayOutput struct{ *pulumi.OutputState }

func (DomainSamlOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainSamlOptions)(nil)).Elem()
}

func (o DomainSamlOptionsArrayOutput) ToDomainSamlOptionsArrayOutput() DomainSamlOptionsArrayOutput {
	return o
}

func (o DomainSamlOptionsArrayOutput) ToDomainSamlOptionsArrayOutputWithContext(ctx context.Context) DomainSamlOptionsArrayOutput {
	return o
}

func (o DomainSamlOptionsArrayOutput) Index(i pulumi.IntInput) DomainSamlOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainSamlOptions {
		return vs[0].([]*DomainSamlOptions)[vs[1].(int)]
	}).(DomainSamlOptionsOutput)
}

type DomainSamlOptionsMapOutput struct{ *pulumi.OutputState }

func (DomainSamlOptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainSamlOptions)(nil)).Elem()
}

func (o DomainSamlOptionsMapOutput) ToDomainSamlOptionsMapOutput() DomainSamlOptionsMapOutput {
	return o
}

func (o DomainSamlOptionsMapOutput) ToDomainSamlOptionsMapOutputWithContext(ctx context.Context) DomainSamlOptionsMapOutput {
	return o
}

func (o DomainSamlOptionsMapOutput) MapIndex(k pulumi.StringInput) DomainSamlOptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainSamlOptions {
		return vs[0].(map[string]*DomainSamlOptions)[vs[1].(string)]
	}).(DomainSamlOptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainSamlOptionsInput)(nil)).Elem(), &DomainSamlOptions{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainSamlOptionsArrayInput)(nil)).Elem(), DomainSamlOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainSamlOptionsMapInput)(nil)).Elem(), DomainSamlOptionsMap{})
	pulumi.RegisterOutputType(DomainSamlOptionsOutput{})
	pulumi.RegisterOutputType(DomainSamlOptionsArrayOutput{})
	pulumi.RegisterOutputType(DomainSamlOptionsMapOutput{})
}
