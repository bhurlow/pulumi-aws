// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sesv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource for managing an AWS SESv2 (Simple Email V2) Dedicated IP Pool.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewDedicatedIpPool(ctx, "example", &sesv2.DedicatedIpPoolArgs{
//				PoolName: pulumi.String("my-pool"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Managed Pool
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewDedicatedIpPool(ctx, "example", &sesv2.DedicatedIpPoolArgs{
//				PoolName:    pulumi.String("my-managed-pool"),
//				ScalingMode: pulumi.String("MANAGED"),
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
// Using `pulumi import`, import SESv2 (Simple Email V2) Dedicated IP Pool using the `pool_name`. For example:
//
// ```sh
//
//	$ pulumi import aws:sesv2/dedicatedIpPool:DedicatedIpPool example my-pool
//
// ```
type DedicatedIpPool struct {
	pulumi.CustomResourceState

	// ARN of the Dedicated IP Pool.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the dedicated IP pool.
	//
	// The following arguments are optional:
	PoolName pulumi.StringOutput `pulumi:"poolName"`
	// IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
	ScalingMode pulumi.StringOutput `pulumi:"scalingMode"`
	// A map of tags to assign to the pool. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDedicatedIpPool registers a new resource with the given unique name, arguments, and options.
func NewDedicatedIpPool(ctx *pulumi.Context,
	name string, args *DedicatedIpPoolArgs, opts ...pulumi.ResourceOption) (*DedicatedIpPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DedicatedIpPool
	err := ctx.RegisterResource("aws:sesv2/dedicatedIpPool:DedicatedIpPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedIpPool gets an existing DedicatedIpPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedIpPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedIpPoolState, opts ...pulumi.ResourceOption) (*DedicatedIpPool, error) {
	var resource DedicatedIpPool
	err := ctx.ReadResource("aws:sesv2/dedicatedIpPool:DedicatedIpPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedIpPool resources.
type dedicatedIpPoolState struct {
	// ARN of the Dedicated IP Pool.
	Arn *string `pulumi:"arn"`
	// Name of the dedicated IP pool.
	//
	// The following arguments are optional:
	PoolName *string `pulumi:"poolName"`
	// IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
	ScalingMode *string `pulumi:"scalingMode"`
	// A map of tags to assign to the pool. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DedicatedIpPoolState struct {
	// ARN of the Dedicated IP Pool.
	Arn pulumi.StringPtrInput
	// Name of the dedicated IP pool.
	//
	// The following arguments are optional:
	PoolName pulumi.StringPtrInput
	// IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
	ScalingMode pulumi.StringPtrInput
	// A map of tags to assign to the pool. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (DedicatedIpPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedIpPoolState)(nil)).Elem()
}

type dedicatedIpPoolArgs struct {
	// Name of the dedicated IP pool.
	//
	// The following arguments are optional:
	PoolName string `pulumi:"poolName"`
	// IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
	ScalingMode *string `pulumi:"scalingMode"`
	// A map of tags to assign to the pool. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DedicatedIpPool resource.
type DedicatedIpPoolArgs struct {
	// Name of the dedicated IP pool.
	//
	// The following arguments are optional:
	PoolName pulumi.StringInput
	// IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
	ScalingMode pulumi.StringPtrInput
	// A map of tags to assign to the pool. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (DedicatedIpPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedIpPoolArgs)(nil)).Elem()
}

type DedicatedIpPoolInput interface {
	pulumi.Input

	ToDedicatedIpPoolOutput() DedicatedIpPoolOutput
	ToDedicatedIpPoolOutputWithContext(ctx context.Context) DedicatedIpPoolOutput
}

func (*DedicatedIpPool) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedIpPool)(nil)).Elem()
}

func (i *DedicatedIpPool) ToDedicatedIpPoolOutput() DedicatedIpPoolOutput {
	return i.ToDedicatedIpPoolOutputWithContext(context.Background())
}

func (i *DedicatedIpPool) ToDedicatedIpPoolOutputWithContext(ctx context.Context) DedicatedIpPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedIpPoolOutput)
}

func (i *DedicatedIpPool) ToOutput(ctx context.Context) pulumix.Output[*DedicatedIpPool] {
	return pulumix.Output[*DedicatedIpPool]{
		OutputState: i.ToDedicatedIpPoolOutputWithContext(ctx).OutputState,
	}
}

// DedicatedIpPoolArrayInput is an input type that accepts DedicatedIpPoolArray and DedicatedIpPoolArrayOutput values.
// You can construct a concrete instance of `DedicatedIpPoolArrayInput` via:
//
//	DedicatedIpPoolArray{ DedicatedIpPoolArgs{...} }
type DedicatedIpPoolArrayInput interface {
	pulumi.Input

	ToDedicatedIpPoolArrayOutput() DedicatedIpPoolArrayOutput
	ToDedicatedIpPoolArrayOutputWithContext(context.Context) DedicatedIpPoolArrayOutput
}

type DedicatedIpPoolArray []DedicatedIpPoolInput

func (DedicatedIpPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedIpPool)(nil)).Elem()
}

func (i DedicatedIpPoolArray) ToDedicatedIpPoolArrayOutput() DedicatedIpPoolArrayOutput {
	return i.ToDedicatedIpPoolArrayOutputWithContext(context.Background())
}

func (i DedicatedIpPoolArray) ToDedicatedIpPoolArrayOutputWithContext(ctx context.Context) DedicatedIpPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedIpPoolArrayOutput)
}

func (i DedicatedIpPoolArray) ToOutput(ctx context.Context) pulumix.Output[[]*DedicatedIpPool] {
	return pulumix.Output[[]*DedicatedIpPool]{
		OutputState: i.ToDedicatedIpPoolArrayOutputWithContext(ctx).OutputState,
	}
}

// DedicatedIpPoolMapInput is an input type that accepts DedicatedIpPoolMap and DedicatedIpPoolMapOutput values.
// You can construct a concrete instance of `DedicatedIpPoolMapInput` via:
//
//	DedicatedIpPoolMap{ "key": DedicatedIpPoolArgs{...} }
type DedicatedIpPoolMapInput interface {
	pulumi.Input

	ToDedicatedIpPoolMapOutput() DedicatedIpPoolMapOutput
	ToDedicatedIpPoolMapOutputWithContext(context.Context) DedicatedIpPoolMapOutput
}

type DedicatedIpPoolMap map[string]DedicatedIpPoolInput

func (DedicatedIpPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedIpPool)(nil)).Elem()
}

func (i DedicatedIpPoolMap) ToDedicatedIpPoolMapOutput() DedicatedIpPoolMapOutput {
	return i.ToDedicatedIpPoolMapOutputWithContext(context.Background())
}

func (i DedicatedIpPoolMap) ToDedicatedIpPoolMapOutputWithContext(ctx context.Context) DedicatedIpPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedIpPoolMapOutput)
}

func (i DedicatedIpPoolMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DedicatedIpPool] {
	return pulumix.Output[map[string]*DedicatedIpPool]{
		OutputState: i.ToDedicatedIpPoolMapOutputWithContext(ctx).OutputState,
	}
}

type DedicatedIpPoolOutput struct{ *pulumi.OutputState }

func (DedicatedIpPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedIpPool)(nil)).Elem()
}

func (o DedicatedIpPoolOutput) ToDedicatedIpPoolOutput() DedicatedIpPoolOutput {
	return o
}

func (o DedicatedIpPoolOutput) ToDedicatedIpPoolOutputWithContext(ctx context.Context) DedicatedIpPoolOutput {
	return o
}

func (o DedicatedIpPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*DedicatedIpPool] {
	return pulumix.Output[*DedicatedIpPool]{
		OutputState: o.OutputState,
	}
}

// ARN of the Dedicated IP Pool.
func (o DedicatedIpPoolOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedIpPool) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of the dedicated IP pool.
//
// The following arguments are optional:
func (o DedicatedIpPoolOutput) PoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedIpPool) pulumi.StringOutput { return v.PoolName }).(pulumi.StringOutput)
}

// IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
func (o DedicatedIpPoolOutput) ScalingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedIpPool) pulumi.StringOutput { return v.ScalingMode }).(pulumi.StringOutput)
}

// A map of tags to assign to the pool. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DedicatedIpPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedIpPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o DedicatedIpPoolOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedIpPool) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type DedicatedIpPoolArrayOutput struct{ *pulumi.OutputState }

func (DedicatedIpPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedIpPool)(nil)).Elem()
}

func (o DedicatedIpPoolArrayOutput) ToDedicatedIpPoolArrayOutput() DedicatedIpPoolArrayOutput {
	return o
}

func (o DedicatedIpPoolArrayOutput) ToDedicatedIpPoolArrayOutputWithContext(ctx context.Context) DedicatedIpPoolArrayOutput {
	return o
}

func (o DedicatedIpPoolArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DedicatedIpPool] {
	return pulumix.Output[[]*DedicatedIpPool]{
		OutputState: o.OutputState,
	}
}

func (o DedicatedIpPoolArrayOutput) Index(i pulumi.IntInput) DedicatedIpPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedIpPool {
		return vs[0].([]*DedicatedIpPool)[vs[1].(int)]
	}).(DedicatedIpPoolOutput)
}

type DedicatedIpPoolMapOutput struct{ *pulumi.OutputState }

func (DedicatedIpPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedIpPool)(nil)).Elem()
}

func (o DedicatedIpPoolMapOutput) ToDedicatedIpPoolMapOutput() DedicatedIpPoolMapOutput {
	return o
}

func (o DedicatedIpPoolMapOutput) ToDedicatedIpPoolMapOutputWithContext(ctx context.Context) DedicatedIpPoolMapOutput {
	return o
}

func (o DedicatedIpPoolMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DedicatedIpPool] {
	return pulumix.Output[map[string]*DedicatedIpPool]{
		OutputState: o.OutputState,
	}
}

func (o DedicatedIpPoolMapOutput) MapIndex(k pulumi.StringInput) DedicatedIpPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedIpPool {
		return vs[0].(map[string]*DedicatedIpPool)[vs[1].(string)]
	}).(DedicatedIpPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedIpPoolInput)(nil)).Elem(), &DedicatedIpPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedIpPoolArrayInput)(nil)).Elem(), DedicatedIpPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedIpPoolMapInput)(nil)).Elem(), DedicatedIpPoolMap{})
	pulumi.RegisterOutputType(DedicatedIpPoolOutput{})
	pulumi.RegisterOutputType(DedicatedIpPoolArrayOutput{})
	pulumi.RegisterOutputType(DedicatedIpPoolMapOutput{})
}
