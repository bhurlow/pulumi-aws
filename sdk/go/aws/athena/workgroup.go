// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an Athena Workgroup.
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
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := athena.NewWorkgroup(ctx, "example", &athena.WorkgroupArgs{
//				Configuration: &athena.WorkgroupConfigurationArgs{
//					EnforceWorkgroupConfiguration:   pulumi.Bool(true),
//					PublishCloudwatchMetricsEnabled: pulumi.Bool(true),
//					ResultConfiguration: &athena.WorkgroupConfigurationResultConfigurationArgs{
//						OutputLocation: pulumi.String(fmt.Sprintf("s3://%v/output/", aws_s3_bucket.Example.Bucket)),
//						EncryptionConfiguration: &athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs{
//							EncryptionOption: pulumi.String("SSE_KMS"),
//							KmsKeyArn:        pulumi.Any(aws_kms_key.Example.Arn),
//						},
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
// Using `pulumi import`, import Athena Workgroups using their name. For example:
//
// ```sh
//
//	$ pulumi import aws:athena/workgroup:Workgroup example example
//
// ```
type Workgroup struct {
	pulumi.CustomResourceState

	// ARN of the workgroup
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration WorkgroupConfigurationPtrOutput `pulumi:"configuration"`
	// Description of the workgroup.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Name of the workgroup.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Key-value map of resource tags for the workgroup. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewWorkgroup registers a new resource with the given unique name, arguments, and options.
func NewWorkgroup(ctx *pulumi.Context,
	name string, args *WorkgroupArgs, opts ...pulumi.ResourceOption) (*Workgroup, error) {
	if args == nil {
		args = &WorkgroupArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Workgroup
	err := ctx.RegisterResource("aws:athena/workgroup:Workgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkgroup gets an existing Workgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkgroupState, opts ...pulumi.ResourceOption) (*Workgroup, error) {
	var resource Workgroup
	err := ctx.ReadResource("aws:athena/workgroup:Workgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workgroup resources.
type workgroupState struct {
	// ARN of the workgroup
	Arn *string `pulumi:"arn"`
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration *WorkgroupConfiguration `pulumi:"configuration"`
	// Description of the workgroup.
	Description *string `pulumi:"description"`
	// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Name of the workgroup.
	Name *string `pulumi:"name"`
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags for the workgroup. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type WorkgroupState struct {
	// ARN of the workgroup
	Arn pulumi.StringPtrInput
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration WorkgroupConfigurationPtrInput
	// Description of the workgroup.
	Description pulumi.StringPtrInput
	// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy pulumi.BoolPtrInput
	// Name of the workgroup.
	Name pulumi.StringPtrInput
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags for the workgroup. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (WorkgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workgroupState)(nil)).Elem()
}

type workgroupArgs struct {
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration *WorkgroupConfiguration `pulumi:"configuration"`
	// Description of the workgroup.
	Description *string `pulumi:"description"`
	// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Name of the workgroup.
	Name *string `pulumi:"name"`
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags for the workgroup. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workgroup resource.
type WorkgroupArgs struct {
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration WorkgroupConfigurationPtrInput
	// Description of the workgroup.
	Description pulumi.StringPtrInput
	// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy pulumi.BoolPtrInput
	// Name of the workgroup.
	Name pulumi.StringPtrInput
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags for the workgroup. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (WorkgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workgroupArgs)(nil)).Elem()
}

type WorkgroupInput interface {
	pulumi.Input

	ToWorkgroupOutput() WorkgroupOutput
	ToWorkgroupOutputWithContext(ctx context.Context) WorkgroupOutput
}

func (*Workgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Workgroup)(nil)).Elem()
}

func (i *Workgroup) ToWorkgroupOutput() WorkgroupOutput {
	return i.ToWorkgroupOutputWithContext(context.Background())
}

func (i *Workgroup) ToWorkgroupOutputWithContext(ctx context.Context) WorkgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkgroupOutput)
}

func (i *Workgroup) ToOutput(ctx context.Context) pulumix.Output[*Workgroup] {
	return pulumix.Output[*Workgroup]{
		OutputState: i.ToWorkgroupOutputWithContext(ctx).OutputState,
	}
}

// WorkgroupArrayInput is an input type that accepts WorkgroupArray and WorkgroupArrayOutput values.
// You can construct a concrete instance of `WorkgroupArrayInput` via:
//
//	WorkgroupArray{ WorkgroupArgs{...} }
type WorkgroupArrayInput interface {
	pulumi.Input

	ToWorkgroupArrayOutput() WorkgroupArrayOutput
	ToWorkgroupArrayOutputWithContext(context.Context) WorkgroupArrayOutput
}

type WorkgroupArray []WorkgroupInput

func (WorkgroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workgroup)(nil)).Elem()
}

func (i WorkgroupArray) ToWorkgroupArrayOutput() WorkgroupArrayOutput {
	return i.ToWorkgroupArrayOutputWithContext(context.Background())
}

func (i WorkgroupArray) ToWorkgroupArrayOutputWithContext(ctx context.Context) WorkgroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkgroupArrayOutput)
}

func (i WorkgroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*Workgroup] {
	return pulumix.Output[[]*Workgroup]{
		OutputState: i.ToWorkgroupArrayOutputWithContext(ctx).OutputState,
	}
}

// WorkgroupMapInput is an input type that accepts WorkgroupMap and WorkgroupMapOutput values.
// You can construct a concrete instance of `WorkgroupMapInput` via:
//
//	WorkgroupMap{ "key": WorkgroupArgs{...} }
type WorkgroupMapInput interface {
	pulumi.Input

	ToWorkgroupMapOutput() WorkgroupMapOutput
	ToWorkgroupMapOutputWithContext(context.Context) WorkgroupMapOutput
}

type WorkgroupMap map[string]WorkgroupInput

func (WorkgroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workgroup)(nil)).Elem()
}

func (i WorkgroupMap) ToWorkgroupMapOutput() WorkgroupMapOutput {
	return i.ToWorkgroupMapOutputWithContext(context.Background())
}

func (i WorkgroupMap) ToWorkgroupMapOutputWithContext(ctx context.Context) WorkgroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkgroupMapOutput)
}

func (i WorkgroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workgroup] {
	return pulumix.Output[map[string]*Workgroup]{
		OutputState: i.ToWorkgroupMapOutputWithContext(ctx).OutputState,
	}
}

type WorkgroupOutput struct{ *pulumi.OutputState }

func (WorkgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workgroup)(nil)).Elem()
}

func (o WorkgroupOutput) ToWorkgroupOutput() WorkgroupOutput {
	return o
}

func (o WorkgroupOutput) ToWorkgroupOutputWithContext(ctx context.Context) WorkgroupOutput {
	return o
}

func (o WorkgroupOutput) ToOutput(ctx context.Context) pulumix.Output[*Workgroup] {
	return pulumix.Output[*Workgroup]{
		OutputState: o.OutputState,
	}
}

// ARN of the workgroup
func (o WorkgroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Workgroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Configuration block with various settings for the workgroup. Documented below.
func (o WorkgroupOutput) Configuration() WorkgroupConfigurationPtrOutput {
	return o.ApplyT(func(v *Workgroup) WorkgroupConfigurationPtrOutput { return v.Configuration }).(WorkgroupConfigurationPtrOutput)
}

// Description of the workgroup.
func (o WorkgroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workgroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
func (o WorkgroupOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workgroup) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// Name of the workgroup.
func (o WorkgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
func (o WorkgroupOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workgroup) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Key-value map of resource tags for the workgroup. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o WorkgroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workgroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o WorkgroupOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workgroup) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type WorkgroupArrayOutput struct{ *pulumi.OutputState }

func (WorkgroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workgroup)(nil)).Elem()
}

func (o WorkgroupArrayOutput) ToWorkgroupArrayOutput() WorkgroupArrayOutput {
	return o
}

func (o WorkgroupArrayOutput) ToWorkgroupArrayOutputWithContext(ctx context.Context) WorkgroupArrayOutput {
	return o
}

func (o WorkgroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Workgroup] {
	return pulumix.Output[[]*Workgroup]{
		OutputState: o.OutputState,
	}
}

func (o WorkgroupArrayOutput) Index(i pulumi.IntInput) WorkgroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workgroup {
		return vs[0].([]*Workgroup)[vs[1].(int)]
	}).(WorkgroupOutput)
}

type WorkgroupMapOutput struct{ *pulumi.OutputState }

func (WorkgroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workgroup)(nil)).Elem()
}

func (o WorkgroupMapOutput) ToWorkgroupMapOutput() WorkgroupMapOutput {
	return o
}

func (o WorkgroupMapOutput) ToWorkgroupMapOutputWithContext(ctx context.Context) WorkgroupMapOutput {
	return o
}

func (o WorkgroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workgroup] {
	return pulumix.Output[map[string]*Workgroup]{
		OutputState: o.OutputState,
	}
}

func (o WorkgroupMapOutput) MapIndex(k pulumi.StringInput) WorkgroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workgroup {
		return vs[0].(map[string]*Workgroup)[vs[1].(string)]
	}).(WorkgroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkgroupInput)(nil)).Elem(), &Workgroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkgroupArrayInput)(nil)).Elem(), WorkgroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkgroupMapInput)(nil)).Elem(), WorkgroupMap{})
	pulumi.RegisterOutputType(WorkgroupOutput{})
	pulumi.RegisterOutputType(WorkgroupArrayOutput{})
	pulumi.RegisterOutputType(WorkgroupMapOutput{})
}
