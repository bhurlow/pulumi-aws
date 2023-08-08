// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Imports a disk image from S3 as a Snapshot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ebs.NewSnapshotImport(ctx, "example", &ebs.SnapshotImportArgs{
//				DiskContainer: &ebs.SnapshotImportDiskContainerArgs{
//					Format: pulumi.String("VHD"),
//					UserBucket: &ebs.SnapshotImportDiskContainerUserBucketArgs{
//						S3Bucket: pulumi.String("disk-images"),
//						S3Key:    pulumi.String("source.vhd"),
//					},
//				},
//				RoleName: pulumi.String("disk-image-import"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("HelloWorld"),
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
type SnapshotImport struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The client-specific data. Detailed below.
	ClientData SnapshotImportClientDataPtrOutput `pulumi:"clientData"`
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId pulumi.StringOutput `pulumi:"dataEncryptionKeyId"`
	// The description string for the import snapshot task.
	Description pulumi.StringOutput `pulumi:"description"`
	// Information about the disk container. Detailed below.
	DiskContainer SnapshotImportDiskContainerOutput `pulumi:"diskContainer"`
	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted pulumi.BoolPtrOutput `pulumi:"encrypted"`
	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	KmsKeyId   pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	OutpostArn pulumi.StringOutput    `pulumi:"outpostArn"`
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias pulumi.StringOutput `pulumi:"ownerAlias"`
	// The AWS account ID of the EBS snapshot owner.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore pulumi.BoolPtrOutput `pulumi:"permanentRestore"`
	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier pulumi.StringOutput `pulumi:"storageTier"`
	// A map of tags to assign to the snapshot.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays pulumi.IntPtrOutput `pulumi:"temporaryRestoreDays"`
	VolumeId             pulumi.StringOutput `pulumi:"volumeId"`
	// The size of the drive in GiBs.
	VolumeSize pulumi.IntOutput `pulumi:"volumeSize"`
}

// NewSnapshotImport registers a new resource with the given unique name, arguments, and options.
func NewSnapshotImport(ctx *pulumi.Context,
	name string, args *SnapshotImportArgs, opts ...pulumi.ResourceOption) (*SnapshotImport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskContainer == nil {
		return nil, errors.New("invalid value for required argument 'DiskContainer'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SnapshotImport
	err := ctx.RegisterResource("aws:ebs/snapshotImport:SnapshotImport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotImport gets an existing SnapshotImport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotImport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotImportState, opts ...pulumi.ResourceOption) (*SnapshotImport, error) {
	var resource SnapshotImport
	err := ctx.ReadResource("aws:ebs/snapshotImport:SnapshotImport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotImport resources.
type snapshotImportState struct {
	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn *string `pulumi:"arn"`
	// The client-specific data. Detailed below.
	ClientData *SnapshotImportClientData `pulumi:"clientData"`
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId *string `pulumi:"dataEncryptionKeyId"`
	// The description string for the import snapshot task.
	Description *string `pulumi:"description"`
	// Information about the disk container. Detailed below.
	DiskContainer *SnapshotImportDiskContainer `pulumi:"diskContainer"`
	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted *bool `pulumi:"encrypted"`
	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	KmsKeyId   *string `pulumi:"kmsKeyId"`
	OutpostArn *string `pulumi:"outpostArn"`
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias *string `pulumi:"ownerAlias"`
	// The AWS account ID of the EBS snapshot owner.
	OwnerId *string `pulumi:"ownerId"`
	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore *bool `pulumi:"permanentRestore"`
	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
	RoleName *string `pulumi:"roleName"`
	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier *string `pulumi:"storageTier"`
	// A map of tags to assign to the snapshot.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays *int    `pulumi:"temporaryRestoreDays"`
	VolumeId             *string `pulumi:"volumeId"`
	// The size of the drive in GiBs.
	VolumeSize *int `pulumi:"volumeSize"`
}

type SnapshotImportState struct {
	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn pulumi.StringPtrInput
	// The client-specific data. Detailed below.
	ClientData SnapshotImportClientDataPtrInput
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId pulumi.StringPtrInput
	// The description string for the import snapshot task.
	Description pulumi.StringPtrInput
	// Information about the disk container. Detailed below.
	DiskContainer SnapshotImportDiskContainerPtrInput
	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted pulumi.BoolPtrInput
	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	KmsKeyId   pulumi.StringPtrInput
	OutpostArn pulumi.StringPtrInput
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias pulumi.StringPtrInput
	// The AWS account ID of the EBS snapshot owner.
	OwnerId pulumi.StringPtrInput
	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore pulumi.BoolPtrInput
	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
	RoleName pulumi.StringPtrInput
	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier pulumi.StringPtrInput
	// A map of tags to assign to the snapshot.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays pulumi.IntPtrInput
	VolumeId             pulumi.StringPtrInput
	// The size of the drive in GiBs.
	VolumeSize pulumi.IntPtrInput
}

func (SnapshotImportState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotImportState)(nil)).Elem()
}

type snapshotImportArgs struct {
	// The client-specific data. Detailed below.
	ClientData *SnapshotImportClientData `pulumi:"clientData"`
	// The description string for the import snapshot task.
	Description *string `pulumi:"description"`
	// Information about the disk container. Detailed below.
	DiskContainer SnapshotImportDiskContainer `pulumi:"diskContainer"`
	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted *bool `pulumi:"encrypted"`
	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore *bool `pulumi:"permanentRestore"`
	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
	RoleName *string `pulumi:"roleName"`
	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier *string `pulumi:"storageTier"`
	// A map of tags to assign to the snapshot.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays *int `pulumi:"temporaryRestoreDays"`
}

// The set of arguments for constructing a SnapshotImport resource.
type SnapshotImportArgs struct {
	// The client-specific data. Detailed below.
	ClientData SnapshotImportClientDataPtrInput
	// The description string for the import snapshot task.
	Description pulumi.StringPtrInput
	// Information about the disk container. Detailed below.
	DiskContainer SnapshotImportDiskContainerInput
	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted pulumi.BoolPtrInput
	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	KmsKeyId pulumi.StringPtrInput
	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore pulumi.BoolPtrInput
	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
	RoleName pulumi.StringPtrInput
	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier pulumi.StringPtrInput
	// A map of tags to assign to the snapshot.
	Tags pulumi.StringMapInput
	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays pulumi.IntPtrInput
}

func (SnapshotImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotImportArgs)(nil)).Elem()
}

type SnapshotImportInput interface {
	pulumi.Input

	ToSnapshotImportOutput() SnapshotImportOutput
	ToSnapshotImportOutputWithContext(ctx context.Context) SnapshotImportOutput
}

func (*SnapshotImport) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotImport)(nil)).Elem()
}

func (i *SnapshotImport) ToSnapshotImportOutput() SnapshotImportOutput {
	return i.ToSnapshotImportOutputWithContext(context.Background())
}

func (i *SnapshotImport) ToSnapshotImportOutputWithContext(ctx context.Context) SnapshotImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotImportOutput)
}

// SnapshotImportArrayInput is an input type that accepts SnapshotImportArray and SnapshotImportArrayOutput values.
// You can construct a concrete instance of `SnapshotImportArrayInput` via:
//
//	SnapshotImportArray{ SnapshotImportArgs{...} }
type SnapshotImportArrayInput interface {
	pulumi.Input

	ToSnapshotImportArrayOutput() SnapshotImportArrayOutput
	ToSnapshotImportArrayOutputWithContext(context.Context) SnapshotImportArrayOutput
}

type SnapshotImportArray []SnapshotImportInput

func (SnapshotImportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotImport)(nil)).Elem()
}

func (i SnapshotImportArray) ToSnapshotImportArrayOutput() SnapshotImportArrayOutput {
	return i.ToSnapshotImportArrayOutputWithContext(context.Background())
}

func (i SnapshotImportArray) ToSnapshotImportArrayOutputWithContext(ctx context.Context) SnapshotImportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotImportArrayOutput)
}

// SnapshotImportMapInput is an input type that accepts SnapshotImportMap and SnapshotImportMapOutput values.
// You can construct a concrete instance of `SnapshotImportMapInput` via:
//
//	SnapshotImportMap{ "key": SnapshotImportArgs{...} }
type SnapshotImportMapInput interface {
	pulumi.Input

	ToSnapshotImportMapOutput() SnapshotImportMapOutput
	ToSnapshotImportMapOutputWithContext(context.Context) SnapshotImportMapOutput
}

type SnapshotImportMap map[string]SnapshotImportInput

func (SnapshotImportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotImport)(nil)).Elem()
}

func (i SnapshotImportMap) ToSnapshotImportMapOutput() SnapshotImportMapOutput {
	return i.ToSnapshotImportMapOutputWithContext(context.Background())
}

func (i SnapshotImportMap) ToSnapshotImportMapOutputWithContext(ctx context.Context) SnapshotImportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotImportMapOutput)
}

type SnapshotImportOutput struct{ *pulumi.OutputState }

func (SnapshotImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotImport)(nil)).Elem()
}

func (o SnapshotImportOutput) ToSnapshotImportOutput() SnapshotImportOutput {
	return o
}

func (o SnapshotImportOutput) ToSnapshotImportOutputWithContext(ctx context.Context) SnapshotImportOutput {
	return o
}

// Amazon Resource Name (ARN) of the EBS Snapshot.
func (o SnapshotImportOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The client-specific data. Detailed below.
func (o SnapshotImportOutput) ClientData() SnapshotImportClientDataPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) SnapshotImportClientDataPtrOutput { return v.ClientData }).(SnapshotImportClientDataPtrOutput)
}

// The data encryption key identifier for the snapshot.
func (o SnapshotImportOutput) DataEncryptionKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.DataEncryptionKeyId }).(pulumi.StringOutput)
}

// The description string for the import snapshot task.
func (o SnapshotImportOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Information about the disk container. Detailed below.
func (o SnapshotImportOutput) DiskContainer() SnapshotImportDiskContainerOutput {
	return o.ApplyT(func(v *SnapshotImport) SnapshotImportDiskContainerOutput { return v.DiskContainer }).(SnapshotImportDiskContainerOutput)
}

// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
func (o SnapshotImportOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.BoolPtrOutput { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
func (o SnapshotImportOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o SnapshotImportOutput) OutpostArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.OutpostArn }).(pulumi.StringOutput)
}

// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
func (o SnapshotImportOutput) OwnerAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.OwnerAlias }).(pulumi.StringOutput)
}

// The AWS account ID of the EBS snapshot owner.
func (o SnapshotImportOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Indicates whether to permanently restore an archived snapshot.
func (o SnapshotImportOutput) PermanentRestore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.BoolPtrOutput { return v.PermanentRestore }).(pulumi.BoolPtrOutput)
}

// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
func (o SnapshotImportOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringPtrOutput { return v.RoleName }).(pulumi.StringPtrOutput)
}

// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
func (o SnapshotImportOutput) StorageTier() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.StorageTier }).(pulumi.StringOutput)
}

// A map of tags to assign to the snapshot.
func (o SnapshotImportOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o SnapshotImportOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
func (o SnapshotImportOutput) TemporaryRestoreDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.IntPtrOutput { return v.TemporaryRestoreDays }).(pulumi.IntPtrOutput)
}

func (o SnapshotImportOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

// The size of the drive in GiBs.
func (o SnapshotImportOutput) VolumeSize() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotImport) pulumi.IntOutput { return v.VolumeSize }).(pulumi.IntOutput)
}

type SnapshotImportArrayOutput struct{ *pulumi.OutputState }

func (SnapshotImportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotImport)(nil)).Elem()
}

func (o SnapshotImportArrayOutput) ToSnapshotImportArrayOutput() SnapshotImportArrayOutput {
	return o
}

func (o SnapshotImportArrayOutput) ToSnapshotImportArrayOutputWithContext(ctx context.Context) SnapshotImportArrayOutput {
	return o
}

func (o SnapshotImportArrayOutput) Index(i pulumi.IntInput) SnapshotImportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotImport {
		return vs[0].([]*SnapshotImport)[vs[1].(int)]
	}).(SnapshotImportOutput)
}

type SnapshotImportMapOutput struct{ *pulumi.OutputState }

func (SnapshotImportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotImport)(nil)).Elem()
}

func (o SnapshotImportMapOutput) ToSnapshotImportMapOutput() SnapshotImportMapOutput {
	return o
}

func (o SnapshotImportMapOutput) ToSnapshotImportMapOutputWithContext(ctx context.Context) SnapshotImportMapOutput {
	return o
}

func (o SnapshotImportMapOutput) MapIndex(k pulumi.StringInput) SnapshotImportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotImport {
		return vs[0].(map[string]*SnapshotImport)[vs[1].(string)]
	}).(SnapshotImportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotImportInput)(nil)).Elem(), &SnapshotImport{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotImportArrayInput)(nil)).Elem(), SnapshotImportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotImportMapInput)(nil)).Elem(), SnapshotImportMap{})
	pulumi.RegisterOutputType(SnapshotImportOutput{})
	pulumi.RegisterOutputType(SnapshotImportArrayOutput{})
	pulumi.RegisterOutputType(SnapshotImportMapOutput{})
}
