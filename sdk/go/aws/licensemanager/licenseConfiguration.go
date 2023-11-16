// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a License Manager license configuration resource.
//
// > **Note:** Removing the `licenseCount` attribute is not supported by the License Manager API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/licensemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := licensemanager.NewLicenseConfiguration(ctx, "example", &licensemanager.LicenseConfigurationArgs{
//				Description:           pulumi.String("Example"),
//				LicenseCount:          pulumi.Int(10),
//				LicenseCountHardLimit: pulumi.Bool(true),
//				LicenseCountingType:   pulumi.String("Socket"),
//				LicenseRules: pulumi.StringArray{
//					pulumi.String("#minimumSockets=2"),
//				},
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("barr"),
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
// ## Rules
//
// License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:
//
// * `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
// * `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
// * `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
// * `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
// * `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
// * `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
// * `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`
//
// ## Import
//
// Using `pulumi import`, import license configurations using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:licensemanager/licenseConfiguration:LicenseConfiguration example arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
//
// ```
type LicenseConfiguration struct {
	pulumi.CustomResourceState

	// The license configuration ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the license configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Number of licenses managed by the license configuration.
	LicenseCount pulumi.IntPtrOutput `pulumi:"licenseCount"`
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit pulumi.BoolPtrOutput `pulumi:"licenseCountHardLimit"`
	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType pulumi.StringOutput `pulumi:"licenseCountingType"`
	// Array of configured License Manager rules.
	LicenseRules pulumi.StringArrayOutput `pulumi:"licenseRules"`
	// Name of the license configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Account ID of the owner of the license configuration.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewLicenseConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLicenseConfiguration(ctx *pulumi.Context,
	name string, args *LicenseConfigurationArgs, opts ...pulumi.ResourceOption) (*LicenseConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LicenseCountingType == nil {
		return nil, errors.New("invalid value for required argument 'LicenseCountingType'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LicenseConfiguration
	err := ctx.RegisterResource("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLicenseConfiguration gets an existing LicenseConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicenseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LicenseConfigurationState, opts ...pulumi.ResourceOption) (*LicenseConfiguration, error) {
	var resource LicenseConfiguration
	err := ctx.ReadResource("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LicenseConfiguration resources.
type licenseConfigurationState struct {
	// The license configuration ARN.
	Arn *string `pulumi:"arn"`
	// Description of the license configuration.
	Description *string `pulumi:"description"`
	// Number of licenses managed by the license configuration.
	LicenseCount *int `pulumi:"licenseCount"`
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool `pulumi:"licenseCountHardLimit"`
	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType *string `pulumi:"licenseCountingType"`
	// Array of configured License Manager rules.
	LicenseRules []string `pulumi:"licenseRules"`
	// Name of the license configuration.
	Name *string `pulumi:"name"`
	// Account ID of the owner of the license configuration.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type LicenseConfigurationState struct {
	// The license configuration ARN.
	Arn pulumi.StringPtrInput
	// Description of the license configuration.
	Description pulumi.StringPtrInput
	// Number of licenses managed by the license configuration.
	LicenseCount pulumi.IntPtrInput
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit pulumi.BoolPtrInput
	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType pulumi.StringPtrInput
	// Array of configured License Manager rules.
	LicenseRules pulumi.StringArrayInput
	// Name of the license configuration.
	Name pulumi.StringPtrInput
	// Account ID of the owner of the license configuration.
	OwnerAccountId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (LicenseConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseConfigurationState)(nil)).Elem()
}

type licenseConfigurationArgs struct {
	// Description of the license configuration.
	Description *string `pulumi:"description"`
	// Number of licenses managed by the license configuration.
	LicenseCount *int `pulumi:"licenseCount"`
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool `pulumi:"licenseCountHardLimit"`
	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType string `pulumi:"licenseCountingType"`
	// Array of configured License Manager rules.
	LicenseRules []string `pulumi:"licenseRules"`
	// Name of the license configuration.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LicenseConfiguration resource.
type LicenseConfigurationArgs struct {
	// Description of the license configuration.
	Description pulumi.StringPtrInput
	// Number of licenses managed by the license configuration.
	LicenseCount pulumi.IntPtrInput
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit pulumi.BoolPtrInput
	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType pulumi.StringInput
	// Array of configured License Manager rules.
	LicenseRules pulumi.StringArrayInput
	// Name of the license configuration.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (LicenseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseConfigurationArgs)(nil)).Elem()
}

type LicenseConfigurationInput interface {
	pulumi.Input

	ToLicenseConfigurationOutput() LicenseConfigurationOutput
	ToLicenseConfigurationOutputWithContext(ctx context.Context) LicenseConfigurationOutput
}

func (*LicenseConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseConfiguration)(nil)).Elem()
}

func (i *LicenseConfiguration) ToLicenseConfigurationOutput() LicenseConfigurationOutput {
	return i.ToLicenseConfigurationOutputWithContext(context.Background())
}

func (i *LicenseConfiguration) ToLicenseConfigurationOutputWithContext(ctx context.Context) LicenseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseConfigurationOutput)
}

// LicenseConfigurationArrayInput is an input type that accepts LicenseConfigurationArray and LicenseConfigurationArrayOutput values.
// You can construct a concrete instance of `LicenseConfigurationArrayInput` via:
//
//	LicenseConfigurationArray{ LicenseConfigurationArgs{...} }
type LicenseConfigurationArrayInput interface {
	pulumi.Input

	ToLicenseConfigurationArrayOutput() LicenseConfigurationArrayOutput
	ToLicenseConfigurationArrayOutputWithContext(context.Context) LicenseConfigurationArrayOutput
}

type LicenseConfigurationArray []LicenseConfigurationInput

func (LicenseConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LicenseConfiguration)(nil)).Elem()
}

func (i LicenseConfigurationArray) ToLicenseConfigurationArrayOutput() LicenseConfigurationArrayOutput {
	return i.ToLicenseConfigurationArrayOutputWithContext(context.Background())
}

func (i LicenseConfigurationArray) ToLicenseConfigurationArrayOutputWithContext(ctx context.Context) LicenseConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseConfigurationArrayOutput)
}

// LicenseConfigurationMapInput is an input type that accepts LicenseConfigurationMap and LicenseConfigurationMapOutput values.
// You can construct a concrete instance of `LicenseConfigurationMapInput` via:
//
//	LicenseConfigurationMap{ "key": LicenseConfigurationArgs{...} }
type LicenseConfigurationMapInput interface {
	pulumi.Input

	ToLicenseConfigurationMapOutput() LicenseConfigurationMapOutput
	ToLicenseConfigurationMapOutputWithContext(context.Context) LicenseConfigurationMapOutput
}

type LicenseConfigurationMap map[string]LicenseConfigurationInput

func (LicenseConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LicenseConfiguration)(nil)).Elem()
}

func (i LicenseConfigurationMap) ToLicenseConfigurationMapOutput() LicenseConfigurationMapOutput {
	return i.ToLicenseConfigurationMapOutputWithContext(context.Background())
}

func (i LicenseConfigurationMap) ToLicenseConfigurationMapOutputWithContext(ctx context.Context) LicenseConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseConfigurationMapOutput)
}

type LicenseConfigurationOutput struct{ *pulumi.OutputState }

func (LicenseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseConfiguration)(nil)).Elem()
}

func (o LicenseConfigurationOutput) ToLicenseConfigurationOutput() LicenseConfigurationOutput {
	return o
}

func (o LicenseConfigurationOutput) ToLicenseConfigurationOutputWithContext(ctx context.Context) LicenseConfigurationOutput {
	return o
}

// The license configuration ARN.
func (o LicenseConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the license configuration.
func (o LicenseConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Number of licenses managed by the license configuration.
func (o LicenseConfigurationOutput) LicenseCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.IntPtrOutput { return v.LicenseCount }).(pulumi.IntPtrOutput)
}

// Sets the number of available licenses as a hard limit.
func (o LicenseConfigurationOutput) LicenseCountHardLimit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.BoolPtrOutput { return v.LicenseCountHardLimit }).(pulumi.BoolPtrOutput)
}

// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
func (o LicenseConfigurationOutput) LicenseCountingType() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringOutput { return v.LicenseCountingType }).(pulumi.StringOutput)
}

// Array of configured License Manager rules.
func (o LicenseConfigurationOutput) LicenseRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringArrayOutput { return v.LicenseRules }).(pulumi.StringArrayOutput)
}

// Name of the license configuration.
func (o LicenseConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Account ID of the owner of the license configuration.
func (o LicenseConfigurationOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o LicenseConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o LicenseConfigurationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LicenseConfiguration) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type LicenseConfigurationArrayOutput struct{ *pulumi.OutputState }

func (LicenseConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LicenseConfiguration)(nil)).Elem()
}

func (o LicenseConfigurationArrayOutput) ToLicenseConfigurationArrayOutput() LicenseConfigurationArrayOutput {
	return o
}

func (o LicenseConfigurationArrayOutput) ToLicenseConfigurationArrayOutputWithContext(ctx context.Context) LicenseConfigurationArrayOutput {
	return o
}

func (o LicenseConfigurationArrayOutput) Index(i pulumi.IntInput) LicenseConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LicenseConfiguration {
		return vs[0].([]*LicenseConfiguration)[vs[1].(int)]
	}).(LicenseConfigurationOutput)
}

type LicenseConfigurationMapOutput struct{ *pulumi.OutputState }

func (LicenseConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LicenseConfiguration)(nil)).Elem()
}

func (o LicenseConfigurationMapOutput) ToLicenseConfigurationMapOutput() LicenseConfigurationMapOutput {
	return o
}

func (o LicenseConfigurationMapOutput) ToLicenseConfigurationMapOutputWithContext(ctx context.Context) LicenseConfigurationMapOutput {
	return o
}

func (o LicenseConfigurationMapOutput) MapIndex(k pulumi.StringInput) LicenseConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LicenseConfiguration {
		return vs[0].(map[string]*LicenseConfiguration)[vs[1].(string)]
	}).(LicenseConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseConfigurationInput)(nil)).Elem(), &LicenseConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseConfigurationArrayInput)(nil)).Elem(), LicenseConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseConfigurationMapInput)(nil)).Elem(), LicenseConfigurationMap{})
	pulumi.RegisterOutputType(LicenseConfigurationOutput{})
	pulumi.RegisterOutputType(LicenseConfigurationArrayOutput{})
	pulumi.RegisterOutputType(LicenseConfigurationMapOutput{})
}
