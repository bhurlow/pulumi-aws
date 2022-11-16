// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS Storage Gateway Tape Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/storagegateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := storagegateway.NewTapePool(ctx, "example", &storagegateway.TapePoolArgs{
//				PoolName:     pulumi.String("example"),
//				StorageClass: pulumi.String("GLACIER"),
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
// `aws_storagegateway_tape_pool` can be imported by using the volume Amazon Resource Name (ARN), e.g.,
//
// ```sh
//
//	$ pulumi import aws:storagegateway/tapePool:TapePool example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678
//
// ```
type TapePool struct {
	pulumi.CustomResourceState

	// Volume Amazon Resource Name (ARN), e.g., `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the new custom tape pool.
	PoolName pulumi.StringOutput `pulumi:"poolName"`
	// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
	RetentionLockTimeInDays pulumi.IntPtrOutput `pulumi:"retentionLockTimeInDays"`
	// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
	RetentionLockType pulumi.StringPtrOutput `pulumi:"retentionLockType"`
	// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTapePool registers a new resource with the given unique name, arguments, and options.
func NewTapePool(ctx *pulumi.Context,
	name string, args *TapePoolArgs, opts ...pulumi.ResourceOption) (*TapePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.StorageClass == nil {
		return nil, errors.New("invalid value for required argument 'StorageClass'")
	}
	var resource TapePool
	err := ctx.RegisterResource("aws:storagegateway/tapePool:TapePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTapePool gets an existing TapePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTapePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TapePoolState, opts ...pulumi.ResourceOption) (*TapePool, error) {
	var resource TapePool
	err := ctx.ReadResource("aws:storagegateway/tapePool:TapePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TapePool resources.
type tapePoolState struct {
	// Volume Amazon Resource Name (ARN), e.g., `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
	Arn *string `pulumi:"arn"`
	// The name of the new custom tape pool.
	PoolName *string `pulumi:"poolName"`
	// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
	RetentionLockTimeInDays *int `pulumi:"retentionLockTimeInDays"`
	// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
	RetentionLockType *string `pulumi:"retentionLockType"`
	// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
	StorageClass *string `pulumi:"storageClass"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type TapePoolState struct {
	// Volume Amazon Resource Name (ARN), e.g., `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
	Arn pulumi.StringPtrInput
	// The name of the new custom tape pool.
	PoolName pulumi.StringPtrInput
	// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
	RetentionLockTimeInDays pulumi.IntPtrInput
	// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
	RetentionLockType pulumi.StringPtrInput
	// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
	StorageClass pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (TapePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*tapePoolState)(nil)).Elem()
}

type tapePoolArgs struct {
	// The name of the new custom tape pool.
	PoolName string `pulumi:"poolName"`
	// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
	RetentionLockTimeInDays *int `pulumi:"retentionLockTimeInDays"`
	// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
	RetentionLockType *string `pulumi:"retentionLockType"`
	// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
	StorageClass string `pulumi:"storageClass"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a TapePool resource.
type TapePoolArgs struct {
	// The name of the new custom tape pool.
	PoolName pulumi.StringInput
	// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
	RetentionLockTimeInDays pulumi.IntPtrInput
	// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
	RetentionLockType pulumi.StringPtrInput
	// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
	StorageClass pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (TapePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tapePoolArgs)(nil)).Elem()
}

type TapePoolInput interface {
	pulumi.Input

	ToTapePoolOutput() TapePoolOutput
	ToTapePoolOutputWithContext(ctx context.Context) TapePoolOutput
}

func (*TapePool) ElementType() reflect.Type {
	return reflect.TypeOf((**TapePool)(nil)).Elem()
}

func (i *TapePool) ToTapePoolOutput() TapePoolOutput {
	return i.ToTapePoolOutputWithContext(context.Background())
}

func (i *TapePool) ToTapePoolOutputWithContext(ctx context.Context) TapePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TapePoolOutput)
}

// TapePoolArrayInput is an input type that accepts TapePoolArray and TapePoolArrayOutput values.
// You can construct a concrete instance of `TapePoolArrayInput` via:
//
//	TapePoolArray{ TapePoolArgs{...} }
type TapePoolArrayInput interface {
	pulumi.Input

	ToTapePoolArrayOutput() TapePoolArrayOutput
	ToTapePoolArrayOutputWithContext(context.Context) TapePoolArrayOutput
}

type TapePoolArray []TapePoolInput

func (TapePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TapePool)(nil)).Elem()
}

func (i TapePoolArray) ToTapePoolArrayOutput() TapePoolArrayOutput {
	return i.ToTapePoolArrayOutputWithContext(context.Background())
}

func (i TapePoolArray) ToTapePoolArrayOutputWithContext(ctx context.Context) TapePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TapePoolArrayOutput)
}

// TapePoolMapInput is an input type that accepts TapePoolMap and TapePoolMapOutput values.
// You can construct a concrete instance of `TapePoolMapInput` via:
//
//	TapePoolMap{ "key": TapePoolArgs{...} }
type TapePoolMapInput interface {
	pulumi.Input

	ToTapePoolMapOutput() TapePoolMapOutput
	ToTapePoolMapOutputWithContext(context.Context) TapePoolMapOutput
}

type TapePoolMap map[string]TapePoolInput

func (TapePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TapePool)(nil)).Elem()
}

func (i TapePoolMap) ToTapePoolMapOutput() TapePoolMapOutput {
	return i.ToTapePoolMapOutputWithContext(context.Background())
}

func (i TapePoolMap) ToTapePoolMapOutputWithContext(ctx context.Context) TapePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TapePoolMapOutput)
}

type TapePoolOutput struct{ *pulumi.OutputState }

func (TapePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TapePool)(nil)).Elem()
}

func (o TapePoolOutput) ToTapePoolOutput() TapePoolOutput {
	return o
}

func (o TapePoolOutput) ToTapePoolOutputWithContext(ctx context.Context) TapePoolOutput {
	return o
}

// Volume Amazon Resource Name (ARN), e.g., `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
func (o TapePoolOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TapePool) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the new custom tape pool.
func (o TapePoolOutput) PoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *TapePool) pulumi.StringOutput { return v.PoolName }).(pulumi.StringOutput)
}

// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
func (o TapePoolOutput) RetentionLockTimeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TapePool) pulumi.IntPtrOutput { return v.RetentionLockTimeInDays }).(pulumi.IntPtrOutput)
}

// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
func (o TapePoolOutput) RetentionLockType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TapePool) pulumi.StringPtrOutput { return v.RetentionLockType }).(pulumi.StringPtrOutput)
}

// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
func (o TapePoolOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v *TapePool) pulumi.StringOutput { return v.StorageClass }).(pulumi.StringOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TapePoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TapePool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o TapePoolOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TapePool) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type TapePoolArrayOutput struct{ *pulumi.OutputState }

func (TapePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TapePool)(nil)).Elem()
}

func (o TapePoolArrayOutput) ToTapePoolArrayOutput() TapePoolArrayOutput {
	return o
}

func (o TapePoolArrayOutput) ToTapePoolArrayOutputWithContext(ctx context.Context) TapePoolArrayOutput {
	return o
}

func (o TapePoolArrayOutput) Index(i pulumi.IntInput) TapePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TapePool {
		return vs[0].([]*TapePool)[vs[1].(int)]
	}).(TapePoolOutput)
}

type TapePoolMapOutput struct{ *pulumi.OutputState }

func (TapePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TapePool)(nil)).Elem()
}

func (o TapePoolMapOutput) ToTapePoolMapOutput() TapePoolMapOutput {
	return o
}

func (o TapePoolMapOutput) ToTapePoolMapOutputWithContext(ctx context.Context) TapePoolMapOutput {
	return o
}

func (o TapePoolMapOutput) MapIndex(k pulumi.StringInput) TapePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TapePool {
		return vs[0].(map[string]*TapePool)[vs[1].(string)]
	}).(TapePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TapePoolInput)(nil)).Elem(), &TapePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*TapePoolArrayInput)(nil)).Elem(), TapePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TapePoolMapInput)(nil)).Elem(), TapePoolMap{})
	pulumi.RegisterOutputType(TapePoolOutput{})
	pulumi.RegisterOutputType(TapePoolArrayOutput{})
	pulumi.RegisterOutputType(TapePoolMapOutput{})
}
