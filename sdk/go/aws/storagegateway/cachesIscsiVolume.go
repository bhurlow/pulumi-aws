// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS Storage Gateway cached iSCSI volume.
//
// > **NOTE:** The gateway must have cache added (e.g., via the `storagegateway.Cache` resource) before creating volumes otherwise the Storage Gateway API will return an error.
//
// > **NOTE:** The gateway must have an upload buffer added (e.g., via the `storagegateway.UploadBuffer` resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `UPLOAD BUFFER NOT CONFIGURED`.
//
// ## Example Usage
//
// > **NOTE:** These examples are referencing the `storagegateway.Cache` resource `gatewayArn` attribute to ensure this provider properly adds cache before creating the volume. If you are not using this method, you may need to declare an expicit dependency (e.g. via `dependsOn = [aws_storagegateway_cache.example]`) to ensure proper ordering.
// ### Create Empty Cached iSCSI Volume
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
//			_, err := storagegateway.NewCachesIscsiVolume(ctx, "example", &storagegateway.CachesIscsiVolumeArgs{
//				GatewayArn:         pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
//				NetworkInterfaceId: pulumi.Any(aws_instance.Example.Private_ip),
//				TargetName:         pulumi.String("example"),
//				VolumeSizeInBytes:  pulumi.Int(5368709120),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create Cached iSCSI Volume From Snapshot
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
//			_, err := storagegateway.NewCachesIscsiVolume(ctx, "example", &storagegateway.CachesIscsiVolumeArgs{
//				GatewayArn:         pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
//				NetworkInterfaceId: pulumi.Any(aws_instance.Example.Private_ip),
//				SnapshotId:         pulumi.Any(aws_ebs_snapshot.Example.Id),
//				TargetName:         pulumi.String("example"),
//				VolumeSizeInBytes:  aws_ebs_snapshot.Example.Volume_size * 1024 * 1024 * 1024,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create Cached iSCSI Volume From Source Volume
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
//			_, err := storagegateway.NewCachesIscsiVolume(ctx, "example", &storagegateway.CachesIscsiVolumeArgs{
//				GatewayArn:         pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
//				NetworkInterfaceId: pulumi.Any(aws_instance.Example.Private_ip),
//				SourceVolumeArn:    pulumi.Any(aws_storagegateway_cached_iscsi_volume.Existing.Arn),
//				TargetName:         pulumi.String("example"),
//				VolumeSizeInBytes:  pulumi.Any(aws_storagegateway_cached_iscsi_volume.Existing.Volume_size_in_bytes),
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
// `aws_storagegateway_cached_iscsi_volume` can be imported by using the volume Amazon Resource Name (ARN), e.g.,
//
// ```sh
//
//	$ pulumi import aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678
//
// ```
type CachesIscsiVolume struct {
	pulumi.CustomResourceState

	// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled pulumi.BoolOutput `pulumi:"chapEnabled"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
	// Set to `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3.
	KmsEncrypted pulumi.BoolPtrOutput `pulumi:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. Is required when `kmsEncrypted` is set.
	KmsKey pulumi.StringPtrOutput `pulumi:"kmsKey"`
	// Logical disk number.
	LunNumber pulumi.IntOutput `pulumi:"lunNumber"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort pulumi.IntOutput `pulumi:"networkInterfacePort"`
	// The snapshot ID of the snapshot to restore as the new cached volumeE.g., `snap-1122aabb`.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn pulumi.StringPtrOutput `pulumi:"sourceVolumeArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Target Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringOutput `pulumi:"targetName"`
	// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	VolumeArn pulumi.StringOutput `pulumi:"volumeArn"`
	// Volume ID, e.g., `vol-12345678`.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// The size of the volume in bytes.
	VolumeSizeInBytes pulumi.IntOutput `pulumi:"volumeSizeInBytes"`
}

// NewCachesIscsiVolume registers a new resource with the given unique name, arguments, and options.
func NewCachesIscsiVolume(ctx *pulumi.Context,
	name string, args *CachesIscsiVolumeArgs, opts ...pulumi.ResourceOption) (*CachesIscsiVolume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'GatewayArn'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	if args.TargetName == nil {
		return nil, errors.New("invalid value for required argument 'TargetName'")
	}
	if args.VolumeSizeInBytes == nil {
		return nil, errors.New("invalid value for required argument 'VolumeSizeInBytes'")
	}
	var resource CachesIscsiVolume
	err := ctx.RegisterResource("aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCachesIscsiVolume gets an existing CachesIscsiVolume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCachesIscsiVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CachesIscsiVolumeState, opts ...pulumi.ResourceOption) (*CachesIscsiVolume, error) {
	var resource CachesIscsiVolume
	err := ctx.ReadResource("aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CachesIscsiVolume resources.
type cachesIscsiVolumeState struct {
	// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn *string `pulumi:"arn"`
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled *bool `pulumi:"chapEnabled"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
	// Set to `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. Is required when `kmsEncrypted` is set.
	KmsKey *string `pulumi:"kmsKey"`
	// Logical disk number.
	LunNumber *int `pulumi:"lunNumber"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort *int `pulumi:"networkInterfacePort"`
	// The snapshot ID of the snapshot to restore as the new cached volumeE.g., `snap-1122aabb`.
	SnapshotId *string `pulumi:"snapshotId"`
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn *string `pulumi:"sourceVolumeArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Target Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn *string `pulumi:"targetArn"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName *string `pulumi:"targetName"`
	// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	VolumeArn *string `pulumi:"volumeArn"`
	// Volume ID, e.g., `vol-12345678`.
	VolumeId *string `pulumi:"volumeId"`
	// The size of the volume in bytes.
	VolumeSizeInBytes *int `pulumi:"volumeSizeInBytes"`
}

type CachesIscsiVolumeState struct {
	// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn pulumi.StringPtrInput
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringPtrInput
	// Set to `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3.
	KmsEncrypted pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. Is required when `kmsEncrypted` is set.
	KmsKey pulumi.StringPtrInput
	// Logical disk number.
	LunNumber pulumi.IntPtrInput
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringPtrInput
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort pulumi.IntPtrInput
	// The snapshot ID of the snapshot to restore as the new cached volumeE.g., `snap-1122aabb`.
	SnapshotId pulumi.StringPtrInput
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Target Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn pulumi.StringPtrInput
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringPtrInput
	// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	VolumeArn pulumi.StringPtrInput
	// Volume ID, e.g., `vol-12345678`.
	VolumeId pulumi.StringPtrInput
	// The size of the volume in bytes.
	VolumeSizeInBytes pulumi.IntPtrInput
}

func (CachesIscsiVolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*cachesIscsiVolumeState)(nil)).Elem()
}

type cachesIscsiVolumeArgs struct {
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
	// Set to `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. Is required when `kmsEncrypted` is set.
	KmsKey *string `pulumi:"kmsKey"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The snapshot ID of the snapshot to restore as the new cached volumeE.g., `snap-1122aabb`.
	SnapshotId *string `pulumi:"snapshotId"`
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn *string `pulumi:"sourceVolumeArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName string `pulumi:"targetName"`
	// The size of the volume in bytes.
	VolumeSizeInBytes int `pulumi:"volumeSizeInBytes"`
}

// The set of arguments for constructing a CachesIscsiVolume resource.
type CachesIscsiVolumeArgs struct {
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringInput
	// Set to `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3.
	KmsEncrypted pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. Is required when `kmsEncrypted` is set.
	KmsKey pulumi.StringPtrInput
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringInput
	// The snapshot ID of the snapshot to restore as the new cached volumeE.g., `snap-1122aabb`.
	SnapshotId pulumi.StringPtrInput
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringInput
	// The size of the volume in bytes.
	VolumeSizeInBytes pulumi.IntInput
}

func (CachesIscsiVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cachesIscsiVolumeArgs)(nil)).Elem()
}

type CachesIscsiVolumeInput interface {
	pulumi.Input

	ToCachesIscsiVolumeOutput() CachesIscsiVolumeOutput
	ToCachesIscsiVolumeOutputWithContext(ctx context.Context) CachesIscsiVolumeOutput
}

func (*CachesIscsiVolume) ElementType() reflect.Type {
	return reflect.TypeOf((**CachesIscsiVolume)(nil)).Elem()
}

func (i *CachesIscsiVolume) ToCachesIscsiVolumeOutput() CachesIscsiVolumeOutput {
	return i.ToCachesIscsiVolumeOutputWithContext(context.Background())
}

func (i *CachesIscsiVolume) ToCachesIscsiVolumeOutputWithContext(ctx context.Context) CachesIscsiVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachesIscsiVolumeOutput)
}

// CachesIscsiVolumeArrayInput is an input type that accepts CachesIscsiVolumeArray and CachesIscsiVolumeArrayOutput values.
// You can construct a concrete instance of `CachesIscsiVolumeArrayInput` via:
//
//	CachesIscsiVolumeArray{ CachesIscsiVolumeArgs{...} }
type CachesIscsiVolumeArrayInput interface {
	pulumi.Input

	ToCachesIscsiVolumeArrayOutput() CachesIscsiVolumeArrayOutput
	ToCachesIscsiVolumeArrayOutputWithContext(context.Context) CachesIscsiVolumeArrayOutput
}

type CachesIscsiVolumeArray []CachesIscsiVolumeInput

func (CachesIscsiVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CachesIscsiVolume)(nil)).Elem()
}

func (i CachesIscsiVolumeArray) ToCachesIscsiVolumeArrayOutput() CachesIscsiVolumeArrayOutput {
	return i.ToCachesIscsiVolumeArrayOutputWithContext(context.Background())
}

func (i CachesIscsiVolumeArray) ToCachesIscsiVolumeArrayOutputWithContext(ctx context.Context) CachesIscsiVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachesIscsiVolumeArrayOutput)
}

// CachesIscsiVolumeMapInput is an input type that accepts CachesIscsiVolumeMap and CachesIscsiVolumeMapOutput values.
// You can construct a concrete instance of `CachesIscsiVolumeMapInput` via:
//
//	CachesIscsiVolumeMap{ "key": CachesIscsiVolumeArgs{...} }
type CachesIscsiVolumeMapInput interface {
	pulumi.Input

	ToCachesIscsiVolumeMapOutput() CachesIscsiVolumeMapOutput
	ToCachesIscsiVolumeMapOutputWithContext(context.Context) CachesIscsiVolumeMapOutput
}

type CachesIscsiVolumeMap map[string]CachesIscsiVolumeInput

func (CachesIscsiVolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CachesIscsiVolume)(nil)).Elem()
}

func (i CachesIscsiVolumeMap) ToCachesIscsiVolumeMapOutput() CachesIscsiVolumeMapOutput {
	return i.ToCachesIscsiVolumeMapOutputWithContext(context.Background())
}

func (i CachesIscsiVolumeMap) ToCachesIscsiVolumeMapOutputWithContext(ctx context.Context) CachesIscsiVolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachesIscsiVolumeMapOutput)
}

type CachesIscsiVolumeOutput struct{ *pulumi.OutputState }

func (CachesIscsiVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CachesIscsiVolume)(nil)).Elem()
}

func (o CachesIscsiVolumeOutput) ToCachesIscsiVolumeOutput() CachesIscsiVolumeOutput {
	return o
}

func (o CachesIscsiVolumeOutput) ToCachesIscsiVolumeOutputWithContext(ctx context.Context) CachesIscsiVolumeOutput {
	return o
}

// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
func (o CachesIscsiVolumeOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Whether mutual CHAP is enabled for the iSCSI target.
func (o CachesIscsiVolumeOutput) ChapEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.BoolOutput { return v.ChapEnabled }).(pulumi.BoolOutput)
}

// The Amazon Resource Name (ARN) of the gateway.
func (o CachesIscsiVolumeOutput) GatewayArn() pulumi.StringOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringOutput { return v.GatewayArn }).(pulumi.StringOutput)
}

// Set to `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3.
func (o CachesIscsiVolumeOutput) KmsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.BoolPtrOutput { return v.KmsEncrypted }).(pulumi.BoolPtrOutput)
}

// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. Is required when `kmsEncrypted` is set.
func (o CachesIscsiVolumeOutput) KmsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringPtrOutput { return v.KmsKey }).(pulumi.StringPtrOutput)
}

// Logical disk number.
func (o CachesIscsiVolumeOutput) LunNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.IntOutput { return v.LunNumber }).(pulumi.IntOutput)
}

// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
func (o CachesIscsiVolumeOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// The port used to communicate with iSCSI targets.
func (o CachesIscsiVolumeOutput) NetworkInterfacePort() pulumi.IntOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.IntOutput { return v.NetworkInterfacePort }).(pulumi.IntOutput)
}

// The snapshot ID of the snapshot to restore as the new cached volumeE.g., `snap-1122aabb`.
func (o CachesIscsiVolumeOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
func (o CachesIscsiVolumeOutput) SourceVolumeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringPtrOutput { return v.SourceVolumeArn }).(pulumi.StringPtrOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CachesIscsiVolumeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o CachesIscsiVolumeOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Target Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
func (o CachesIscsiVolumeOutput) TargetArn() pulumi.StringOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringOutput { return v.TargetArn }).(pulumi.StringOutput)
}

// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
func (o CachesIscsiVolumeOutput) TargetName() pulumi.StringOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringOutput { return v.TargetName }).(pulumi.StringOutput)
}

// Volume Amazon Resource Name (ARN), e.g., `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
func (o CachesIscsiVolumeOutput) VolumeArn() pulumi.StringOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringOutput { return v.VolumeArn }).(pulumi.StringOutput)
}

// Volume ID, e.g., `vol-12345678`.
func (o CachesIscsiVolumeOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

// The size of the volume in bytes.
func (o CachesIscsiVolumeOutput) VolumeSizeInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *CachesIscsiVolume) pulumi.IntOutput { return v.VolumeSizeInBytes }).(pulumi.IntOutput)
}

type CachesIscsiVolumeArrayOutput struct{ *pulumi.OutputState }

func (CachesIscsiVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CachesIscsiVolume)(nil)).Elem()
}

func (o CachesIscsiVolumeArrayOutput) ToCachesIscsiVolumeArrayOutput() CachesIscsiVolumeArrayOutput {
	return o
}

func (o CachesIscsiVolumeArrayOutput) ToCachesIscsiVolumeArrayOutputWithContext(ctx context.Context) CachesIscsiVolumeArrayOutput {
	return o
}

func (o CachesIscsiVolumeArrayOutput) Index(i pulumi.IntInput) CachesIscsiVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CachesIscsiVolume {
		return vs[0].([]*CachesIscsiVolume)[vs[1].(int)]
	}).(CachesIscsiVolumeOutput)
}

type CachesIscsiVolumeMapOutput struct{ *pulumi.OutputState }

func (CachesIscsiVolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CachesIscsiVolume)(nil)).Elem()
}

func (o CachesIscsiVolumeMapOutput) ToCachesIscsiVolumeMapOutput() CachesIscsiVolumeMapOutput {
	return o
}

func (o CachesIscsiVolumeMapOutput) ToCachesIscsiVolumeMapOutputWithContext(ctx context.Context) CachesIscsiVolumeMapOutput {
	return o
}

func (o CachesIscsiVolumeMapOutput) MapIndex(k pulumi.StringInput) CachesIscsiVolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CachesIscsiVolume {
		return vs[0].(map[string]*CachesIscsiVolume)[vs[1].(string)]
	}).(CachesIscsiVolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CachesIscsiVolumeInput)(nil)).Elem(), &CachesIscsiVolume{})
	pulumi.RegisterInputType(reflect.TypeOf((*CachesIscsiVolumeArrayInput)(nil)).Elem(), CachesIscsiVolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CachesIscsiVolumeMapInput)(nil)).Elem(), CachesIscsiVolumeMap{})
	pulumi.RegisterOutputType(CachesIscsiVolumeOutput{})
	pulumi.RegisterOutputType(CachesIscsiVolumeArrayOutput{})
	pulumi.RegisterOutputType(CachesIscsiVolumeMapOutput{})
}
