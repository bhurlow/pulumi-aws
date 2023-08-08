// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fsx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Amazon FSx for OpenZFS volume.
// See the [FSx OpenZFS User Guide](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/what-is-fsx.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fsx.NewOpenZfsVolume(ctx, "test", &fsx.OpenZfsVolumeArgs{
//				ParentVolumeId: pulumi.Any(aws_fsx_openzfs_file_system.Test.Root_volume_id),
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
// terraform import {
//
//	to = aws_fsx_openzfs_volume.example
//
//	id = "fsvol-543ab12b1ca672f33" } Using `pulumi import`, import FSx Volumes using the `id`. For exampleconsole % pulumi import aws_fsx_openzfs_volume.example fsvol-543ab12b1ca672f33
type OpenZfsVolume struct {
	pulumi.CustomResourceState

	// Amazon Resource Name of the file system.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToSnapshots pulumi.BoolPtrOutput `pulumi:"copyTagsToSnapshots"`
	// Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
	DataCompressionType pulumi.StringPtrOutput `pulumi:"dataCompressionType"`
	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name pulumi.StringOutput `pulumi:"name"`
	// NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
	NfsExports OpenZfsVolumeNfsExportsPtrOutput `pulumi:"nfsExports"`
	// The ARN of the source snapshot to create the volume from.
	OriginSnapshot OpenZfsVolumeOriginSnapshotPtrOutput `pulumi:"originSnapshot"`
	// The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `fsx.OpenZfsVolume`.
	ParentVolumeId pulumi.StringOutput `pulumi:"parentVolumeId"`
	// specifies whether the volume is read-only. Default is false.
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
	RecordSizeKib pulumi.IntPtrOutput `pulumi:"recordSizeKib"`
	// The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
	StorageCapacityQuotaGib pulumi.IntOutput `pulumi:"storageCapacityQuotaGib"`
	// The amount of storage in gibibytes (GiB) to reserve from the parent volume.
	StorageCapacityReservationGib pulumi.IntOutput `pulumi:"storageCapacityReservationGib"`
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
	UserAndGroupQuotas OpenZfsVolumeUserAndGroupQuotaArrayOutput `pulumi:"userAndGroupQuotas"`
	VolumeType         pulumi.StringPtrOutput                    `pulumi:"volumeType"`
}

// NewOpenZfsVolume registers a new resource with the given unique name, arguments, and options.
func NewOpenZfsVolume(ctx *pulumi.Context,
	name string, args *OpenZfsVolumeArgs, opts ...pulumi.ResourceOption) (*OpenZfsVolume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentVolumeId == nil {
		return nil, errors.New("invalid value for required argument 'ParentVolumeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OpenZfsVolume
	err := ctx.RegisterResource("aws:fsx/openZfsVolume:OpenZfsVolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenZfsVolume gets an existing OpenZfsVolume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenZfsVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenZfsVolumeState, opts ...pulumi.ResourceOption) (*OpenZfsVolume, error) {
	var resource OpenZfsVolume
	err := ctx.ReadResource("aws:fsx/openZfsVolume:OpenZfsVolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenZfsVolume resources.
type openZfsVolumeState struct {
	// Amazon Resource Name of the file system.
	Arn *string `pulumi:"arn"`
	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToSnapshots *bool `pulumi:"copyTagsToSnapshots"`
	// Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
	DataCompressionType *string `pulumi:"dataCompressionType"`
	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name *string `pulumi:"name"`
	// NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
	NfsExports *OpenZfsVolumeNfsExports `pulumi:"nfsExports"`
	// The ARN of the source snapshot to create the volume from.
	OriginSnapshot *OpenZfsVolumeOriginSnapshot `pulumi:"originSnapshot"`
	// The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `fsx.OpenZfsVolume`.
	ParentVolumeId *string `pulumi:"parentVolumeId"`
	// specifies whether the volume is read-only. Default is false.
	ReadOnly *bool `pulumi:"readOnly"`
	// The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
	RecordSizeKib *int `pulumi:"recordSizeKib"`
	// The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
	StorageCapacityQuotaGib *int `pulumi:"storageCapacityQuotaGib"`
	// The amount of storage in gibibytes (GiB) to reserve from the parent volume.
	StorageCapacityReservationGib *int `pulumi:"storageCapacityReservationGib"`
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
	UserAndGroupQuotas []OpenZfsVolumeUserAndGroupQuota `pulumi:"userAndGroupQuotas"`
	VolumeType         *string                          `pulumi:"volumeType"`
}

type OpenZfsVolumeState struct {
	// Amazon Resource Name of the file system.
	Arn pulumi.StringPtrInput
	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToSnapshots pulumi.BoolPtrInput
	// Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
	DataCompressionType pulumi.StringPtrInput
	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name pulumi.StringPtrInput
	// NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
	NfsExports OpenZfsVolumeNfsExportsPtrInput
	// The ARN of the source snapshot to create the volume from.
	OriginSnapshot OpenZfsVolumeOriginSnapshotPtrInput
	// The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `fsx.OpenZfsVolume`.
	ParentVolumeId pulumi.StringPtrInput
	// specifies whether the volume is read-only. Default is false.
	ReadOnly pulumi.BoolPtrInput
	// The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
	RecordSizeKib pulumi.IntPtrInput
	// The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
	StorageCapacityQuotaGib pulumi.IntPtrInput
	// The amount of storage in gibibytes (GiB) to reserve from the parent volume.
	StorageCapacityReservationGib pulumi.IntPtrInput
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
	UserAndGroupQuotas OpenZfsVolumeUserAndGroupQuotaArrayInput
	VolumeType         pulumi.StringPtrInput
}

func (OpenZfsVolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*openZfsVolumeState)(nil)).Elem()
}

type openZfsVolumeArgs struct {
	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToSnapshots *bool `pulumi:"copyTagsToSnapshots"`
	// Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
	DataCompressionType *string `pulumi:"dataCompressionType"`
	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name *string `pulumi:"name"`
	// NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
	NfsExports *OpenZfsVolumeNfsExports `pulumi:"nfsExports"`
	// The ARN of the source snapshot to create the volume from.
	OriginSnapshot *OpenZfsVolumeOriginSnapshot `pulumi:"originSnapshot"`
	// The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `fsx.OpenZfsVolume`.
	ParentVolumeId string `pulumi:"parentVolumeId"`
	// specifies whether the volume is read-only. Default is false.
	ReadOnly *bool `pulumi:"readOnly"`
	// The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
	RecordSizeKib *int `pulumi:"recordSizeKib"`
	// The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
	StorageCapacityQuotaGib *int `pulumi:"storageCapacityQuotaGib"`
	// The amount of storage in gibibytes (GiB) to reserve from the parent volume.
	StorageCapacityReservationGib *int `pulumi:"storageCapacityReservationGib"`
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
	UserAndGroupQuotas []OpenZfsVolumeUserAndGroupQuota `pulumi:"userAndGroupQuotas"`
	VolumeType         *string                          `pulumi:"volumeType"`
}

// The set of arguments for constructing a OpenZfsVolume resource.
type OpenZfsVolumeArgs struct {
	// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
	CopyTagsToSnapshots pulumi.BoolPtrInput
	// Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
	DataCompressionType pulumi.StringPtrInput
	// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
	Name pulumi.StringPtrInput
	// NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
	NfsExports OpenZfsVolumeNfsExportsPtrInput
	// The ARN of the source snapshot to create the volume from.
	OriginSnapshot OpenZfsVolumeOriginSnapshotPtrInput
	// The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `fsx.OpenZfsVolume`.
	ParentVolumeId pulumi.StringInput
	// specifies whether the volume is read-only. Default is false.
	ReadOnly pulumi.BoolPtrInput
	// The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
	RecordSizeKib pulumi.IntPtrInput
	// The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
	StorageCapacityQuotaGib pulumi.IntPtrInput
	// The amount of storage in gibibytes (GiB) to reserve from the parent volume.
	StorageCapacityReservationGib pulumi.IntPtrInput
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
	UserAndGroupQuotas OpenZfsVolumeUserAndGroupQuotaArrayInput
	VolumeType         pulumi.StringPtrInput
}

func (OpenZfsVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openZfsVolumeArgs)(nil)).Elem()
}

type OpenZfsVolumeInput interface {
	pulumi.Input

	ToOpenZfsVolumeOutput() OpenZfsVolumeOutput
	ToOpenZfsVolumeOutputWithContext(ctx context.Context) OpenZfsVolumeOutput
}

func (*OpenZfsVolume) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenZfsVolume)(nil)).Elem()
}

func (i *OpenZfsVolume) ToOpenZfsVolumeOutput() OpenZfsVolumeOutput {
	return i.ToOpenZfsVolumeOutputWithContext(context.Background())
}

func (i *OpenZfsVolume) ToOpenZfsVolumeOutputWithContext(ctx context.Context) OpenZfsVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenZfsVolumeOutput)
}

// OpenZfsVolumeArrayInput is an input type that accepts OpenZfsVolumeArray and OpenZfsVolumeArrayOutput values.
// You can construct a concrete instance of `OpenZfsVolumeArrayInput` via:
//
//	OpenZfsVolumeArray{ OpenZfsVolumeArgs{...} }
type OpenZfsVolumeArrayInput interface {
	pulumi.Input

	ToOpenZfsVolumeArrayOutput() OpenZfsVolumeArrayOutput
	ToOpenZfsVolumeArrayOutputWithContext(context.Context) OpenZfsVolumeArrayOutput
}

type OpenZfsVolumeArray []OpenZfsVolumeInput

func (OpenZfsVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OpenZfsVolume)(nil)).Elem()
}

func (i OpenZfsVolumeArray) ToOpenZfsVolumeArrayOutput() OpenZfsVolumeArrayOutput {
	return i.ToOpenZfsVolumeArrayOutputWithContext(context.Background())
}

func (i OpenZfsVolumeArray) ToOpenZfsVolumeArrayOutputWithContext(ctx context.Context) OpenZfsVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenZfsVolumeArrayOutput)
}

// OpenZfsVolumeMapInput is an input type that accepts OpenZfsVolumeMap and OpenZfsVolumeMapOutput values.
// You can construct a concrete instance of `OpenZfsVolumeMapInput` via:
//
//	OpenZfsVolumeMap{ "key": OpenZfsVolumeArgs{...} }
type OpenZfsVolumeMapInput interface {
	pulumi.Input

	ToOpenZfsVolumeMapOutput() OpenZfsVolumeMapOutput
	ToOpenZfsVolumeMapOutputWithContext(context.Context) OpenZfsVolumeMapOutput
}

type OpenZfsVolumeMap map[string]OpenZfsVolumeInput

func (OpenZfsVolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OpenZfsVolume)(nil)).Elem()
}

func (i OpenZfsVolumeMap) ToOpenZfsVolumeMapOutput() OpenZfsVolumeMapOutput {
	return i.ToOpenZfsVolumeMapOutputWithContext(context.Background())
}

func (i OpenZfsVolumeMap) ToOpenZfsVolumeMapOutputWithContext(ctx context.Context) OpenZfsVolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenZfsVolumeMapOutput)
}

type OpenZfsVolumeOutput struct{ *pulumi.OutputState }

func (OpenZfsVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenZfsVolume)(nil)).Elem()
}

func (o OpenZfsVolumeOutput) ToOpenZfsVolumeOutput() OpenZfsVolumeOutput {
	return o
}

func (o OpenZfsVolumeOutput) ToOpenZfsVolumeOutputWithContext(ctx context.Context) OpenZfsVolumeOutput {
	return o
}

// Amazon Resource Name of the file system.
func (o OpenZfsVolumeOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A boolean flag indicating whether tags for the file system should be copied to snapshots. The default value is false.
func (o OpenZfsVolumeOutput) CopyTagsToSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.BoolPtrOutput { return v.CopyTagsToSnapshots }).(pulumi.BoolPtrOutput)
}

// Method used to compress the data on the volume. Valid values are `NONE` or `ZSTD`. Child volumes that don't specify compression option will inherit from parent volume. This option on file system applies to the root volume.
func (o OpenZfsVolumeOutput) DataCompressionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.StringPtrOutput { return v.DataCompressionType }).(pulumi.StringPtrOutput)
}

// The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
func (o OpenZfsVolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NFS export configuration for the root volume. Exactly 1 item. See NFS Exports Below.
func (o OpenZfsVolumeOutput) NfsExports() OpenZfsVolumeNfsExportsPtrOutput {
	return o.ApplyT(func(v *OpenZfsVolume) OpenZfsVolumeNfsExportsPtrOutput { return v.NfsExports }).(OpenZfsVolumeNfsExportsPtrOutput)
}

// The ARN of the source snapshot to create the volume from.
func (o OpenZfsVolumeOutput) OriginSnapshot() OpenZfsVolumeOriginSnapshotPtrOutput {
	return o.ApplyT(func(v *OpenZfsVolume) OpenZfsVolumeOriginSnapshotPtrOutput { return v.OriginSnapshot }).(OpenZfsVolumeOriginSnapshotPtrOutput)
}

// The volume id of volume that will be the parent volume for the volume being created, this could be the root volume created from the `fsx.OpenZfsFileSystem` resource with the `rootVolumeId` or the `id` property of another `fsx.OpenZfsVolume`.
func (o OpenZfsVolumeOutput) ParentVolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.StringOutput { return v.ParentVolumeId }).(pulumi.StringOutput)
}

// specifies whether the volume is read-only. Default is false.
func (o OpenZfsVolumeOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.BoolOutput { return v.ReadOnly }).(pulumi.BoolOutput)
}

// The record size of an OpenZFS volume, in kibibytes (KiB). Valid values are `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, or `1024` KiB. The default is `128` KiB.
func (o OpenZfsVolumeOutput) RecordSizeKib() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.IntPtrOutput { return v.RecordSizeKib }).(pulumi.IntPtrOutput)
}

// The maximum amount of storage in gibibytes (GiB) that the volume can use from its parent.
func (o OpenZfsVolumeOutput) StorageCapacityQuotaGib() pulumi.IntOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.IntOutput { return v.StorageCapacityQuotaGib }).(pulumi.IntOutput)
}

// The amount of storage in gibibytes (GiB) to reserve from the parent volume.
func (o OpenZfsVolumeOutput) StorageCapacityReservationGib() pulumi.IntOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.IntOutput { return v.StorageCapacityReservationGib }).(pulumi.IntOutput)
}

// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o OpenZfsVolumeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o OpenZfsVolumeOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Specify how much storage users or groups can use on the volume. Maximum of 100 items. See User and Group Quotas Below.
func (o OpenZfsVolumeOutput) UserAndGroupQuotas() OpenZfsVolumeUserAndGroupQuotaArrayOutput {
	return o.ApplyT(func(v *OpenZfsVolume) OpenZfsVolumeUserAndGroupQuotaArrayOutput { return v.UserAndGroupQuotas }).(OpenZfsVolumeUserAndGroupQuotaArrayOutput)
}

func (o OpenZfsVolumeOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenZfsVolume) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type OpenZfsVolumeArrayOutput struct{ *pulumi.OutputState }

func (OpenZfsVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OpenZfsVolume)(nil)).Elem()
}

func (o OpenZfsVolumeArrayOutput) ToOpenZfsVolumeArrayOutput() OpenZfsVolumeArrayOutput {
	return o
}

func (o OpenZfsVolumeArrayOutput) ToOpenZfsVolumeArrayOutputWithContext(ctx context.Context) OpenZfsVolumeArrayOutput {
	return o
}

func (o OpenZfsVolumeArrayOutput) Index(i pulumi.IntInput) OpenZfsVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OpenZfsVolume {
		return vs[0].([]*OpenZfsVolume)[vs[1].(int)]
	}).(OpenZfsVolumeOutput)
}

type OpenZfsVolumeMapOutput struct{ *pulumi.OutputState }

func (OpenZfsVolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OpenZfsVolume)(nil)).Elem()
}

func (o OpenZfsVolumeMapOutput) ToOpenZfsVolumeMapOutput() OpenZfsVolumeMapOutput {
	return o
}

func (o OpenZfsVolumeMapOutput) ToOpenZfsVolumeMapOutputWithContext(ctx context.Context) OpenZfsVolumeMapOutput {
	return o
}

func (o OpenZfsVolumeMapOutput) MapIndex(k pulumi.StringInput) OpenZfsVolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OpenZfsVolume {
		return vs[0].(map[string]*OpenZfsVolume)[vs[1].(string)]
	}).(OpenZfsVolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OpenZfsVolumeInput)(nil)).Elem(), &OpenZfsVolume{})
	pulumi.RegisterInputType(reflect.TypeOf((*OpenZfsVolumeArrayInput)(nil)).Elem(), OpenZfsVolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OpenZfsVolumeMapInput)(nil)).Elem(), OpenZfsVolumeMap{})
	pulumi.RegisterOutputType(OpenZfsVolumeOutput{})
	pulumi.RegisterOutputType(OpenZfsVolumeArrayOutput{})
	pulumi.RegisterOutputType(OpenZfsVolumeMapOutput{})
}
