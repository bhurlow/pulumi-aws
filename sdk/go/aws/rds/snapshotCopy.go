// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an RDS database instance snapshot copy. For managing RDS database cluster snapshots, see the `rds.ClusterSnapshot` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := rds.NewInstance(ctx, "exampleInstance", &rds.InstanceArgs{
//				AllocatedStorage:      pulumi.Int(10),
//				Engine:                pulumi.String("mysql"),
//				EngineVersion:         pulumi.String("5.6.21"),
//				InstanceClass:         pulumi.String("db.t2.micro"),
//				Name:                  pulumi.String("baz"),
//				Password:              pulumi.String("barbarbarbar"),
//				Username:              pulumi.String("foo"),
//				MaintenanceWindow:     pulumi.String("Fri:09:00-Fri:09:30"),
//				BackupRetentionPeriod: pulumi.Int(0),
//				ParameterGroupName:    pulumi.String("default.mysql5.6"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSnapshot, err := rds.NewSnapshot(ctx, "exampleSnapshot", &rds.SnapshotArgs{
//				DbInstanceIdentifier: exampleInstance.ID(),
//				DbSnapshotIdentifier: pulumi.String("testsnapshot1234"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewSnapshotCopy(ctx, "exampleSnapshotCopy", &rds.SnapshotCopyArgs{
//				SourceDbSnapshotIdentifier: exampleSnapshot.DbSnapshotArn,
//				TargetDbSnapshotIdentifier: pulumi.String("testsnapshot1234-copy"),
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
// `aws_db_snapshot_copy` can be imported by using the snapshot identifier, e.g.,
//
// ```sh
//
//	$ pulumi import aws:rds/snapshotCopy:SnapshotCopy example my-snapshot
//
// ```
type SnapshotCopy struct {
	pulumi.CustomResourceState

	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntOutput `pulumi:"allocatedStorage"`
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Whether to copy existing tags. Defaults to `false`.
	CopyTags pulumi.BoolPtrOutput `pulumi:"copyTags"`
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn pulumi.StringOutput `pulumi:"dbSnapshotArn"`
	// The Destination region to place snapshot copy.
	DestinationRegion pulumi.StringPtrOutput `pulumi:"destinationRegion"`
	// Specifies whether the DB snapshot is encrypted.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`
	// Specifies the name of the database engine.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Specifies the version of the database engine.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops pulumi.IntOutput `pulumi:"iops"`
	// KMS key ID.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// License model information for the restored DB instance.
	LicenseModel pulumi.StringOutput `pulumi:"licenseModel"`
	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName pulumi.StringOutput `pulumi:"optionGroupName"`
	Port            pulumi.IntOutput    `pulumi:"port"`
	// he URL that contains a Signature Version 4 signed request.
	PresignedUrl pulumi.StringPtrOutput `pulumi:"presignedUrl"`
	SnapshotType pulumi.StringOutput    `pulumi:"snapshotType"`
	// Snapshot identifier of the source snapshot.
	SourceDbSnapshotIdentifier pulumi.StringOutput `pulumi:"sourceDbSnapshotIdentifier"`
	// The region that the DB snapshot was created in or copied from.
	SourceRegion pulumi.StringOutput `pulumi:"sourceRegion"`
	// Specifies the storage type associated with DB snapshot.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The external custom Availability Zone.
	TargetCustomAvailabilityZone pulumi.StringPtrOutput `pulumi:"targetCustomAvailabilityZone"`
	// The Identifier for the snapshot.
	TargetDbSnapshotIdentifier pulumi.StringOutput `pulumi:"targetDbSnapshotIdentifier"`
	// Provides the VPC ID associated with the DB snapshot.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSnapshotCopy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotCopy(ctx *pulumi.Context,
	name string, args *SnapshotCopyArgs, opts ...pulumi.ResourceOption) (*SnapshotCopy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceDbSnapshotIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'SourceDbSnapshotIdentifier'")
	}
	if args.TargetDbSnapshotIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'TargetDbSnapshotIdentifier'")
	}
	var resource SnapshotCopy
	err := ctx.RegisterResource("aws:rds/snapshotCopy:SnapshotCopy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotCopy gets an existing SnapshotCopy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotCopy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotCopyState, opts ...pulumi.ResourceOption) (*SnapshotCopy, error) {
	var resource SnapshotCopy
	err := ctx.ReadResource("aws:rds/snapshotCopy:SnapshotCopy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotCopy resources.
type snapshotCopyState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage *int `pulumi:"allocatedStorage"`
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Whether to copy existing tags. Defaults to `false`.
	CopyTags *bool `pulumi:"copyTags"`
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn *string `pulumi:"dbSnapshotArn"`
	// The Destination region to place snapshot copy.
	DestinationRegion *string `pulumi:"destinationRegion"`
	// Specifies whether the DB snapshot is encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// Specifies the name of the database engine.
	Engine *string `pulumi:"engine"`
	// Specifies the version of the database engine.
	EngineVersion *string `pulumi:"engineVersion"`
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops *int `pulumi:"iops"`
	// KMS key ID.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// License model information for the restored DB instance.
	LicenseModel *string `pulumi:"licenseModel"`
	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName *string `pulumi:"optionGroupName"`
	Port            *int    `pulumi:"port"`
	// he URL that contains a Signature Version 4 signed request.
	PresignedUrl *string `pulumi:"presignedUrl"`
	SnapshotType *string `pulumi:"snapshotType"`
	// Snapshot identifier of the source snapshot.
	SourceDbSnapshotIdentifier *string `pulumi:"sourceDbSnapshotIdentifier"`
	// The region that the DB snapshot was created in or copied from.
	SourceRegion *string `pulumi:"sourceRegion"`
	// Specifies the storage type associated with DB snapshot.
	StorageType *string `pulumi:"storageType"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The external custom Availability Zone.
	TargetCustomAvailabilityZone *string `pulumi:"targetCustomAvailabilityZone"`
	// The Identifier for the snapshot.
	TargetDbSnapshotIdentifier *string `pulumi:"targetDbSnapshotIdentifier"`
	// Provides the VPC ID associated with the DB snapshot.
	VpcId *string `pulumi:"vpcId"`
}

type SnapshotCopyState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntPtrInput
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone pulumi.StringPtrInput
	// Whether to copy existing tags. Defaults to `false`.
	CopyTags pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn pulumi.StringPtrInput
	// The Destination region to place snapshot copy.
	DestinationRegion pulumi.StringPtrInput
	// Specifies whether the DB snapshot is encrypted.
	Encrypted pulumi.BoolPtrInput
	// Specifies the name of the database engine.
	Engine pulumi.StringPtrInput
	// Specifies the version of the database engine.
	EngineVersion pulumi.StringPtrInput
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops pulumi.IntPtrInput
	// KMS key ID.
	KmsKeyId pulumi.StringPtrInput
	// License model information for the restored DB instance.
	LicenseModel pulumi.StringPtrInput
	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName pulumi.StringPtrInput
	Port            pulumi.IntPtrInput
	// he URL that contains a Signature Version 4 signed request.
	PresignedUrl pulumi.StringPtrInput
	SnapshotType pulumi.StringPtrInput
	// Snapshot identifier of the source snapshot.
	SourceDbSnapshotIdentifier pulumi.StringPtrInput
	// The region that the DB snapshot was created in or copied from.
	SourceRegion pulumi.StringPtrInput
	// Specifies the storage type associated with DB snapshot.
	StorageType pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The external custom Availability Zone.
	TargetCustomAvailabilityZone pulumi.StringPtrInput
	// The Identifier for the snapshot.
	TargetDbSnapshotIdentifier pulumi.StringPtrInput
	// Provides the VPC ID associated with the DB snapshot.
	VpcId pulumi.StringPtrInput
}

func (SnapshotCopyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCopyState)(nil)).Elem()
}

type snapshotCopyArgs struct {
	// Whether to copy existing tags. Defaults to `false`.
	CopyTags *bool `pulumi:"copyTags"`
	// The Destination region to place snapshot copy.
	DestinationRegion *string `pulumi:"destinationRegion"`
	// KMS key ID.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName *string `pulumi:"optionGroupName"`
	// he URL that contains a Signature Version 4 signed request.
	PresignedUrl *string `pulumi:"presignedUrl"`
	// Snapshot identifier of the source snapshot.
	SourceDbSnapshotIdentifier string `pulumi:"sourceDbSnapshotIdentifier"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The external custom Availability Zone.
	TargetCustomAvailabilityZone *string `pulumi:"targetCustomAvailabilityZone"`
	// The Identifier for the snapshot.
	TargetDbSnapshotIdentifier string `pulumi:"targetDbSnapshotIdentifier"`
}

// The set of arguments for constructing a SnapshotCopy resource.
type SnapshotCopyArgs struct {
	// Whether to copy existing tags. Defaults to `false`.
	CopyTags pulumi.BoolPtrInput
	// The Destination region to place snapshot copy.
	DestinationRegion pulumi.StringPtrInput
	// KMS key ID.
	KmsKeyId pulumi.StringPtrInput
	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName pulumi.StringPtrInput
	// he URL that contains a Signature Version 4 signed request.
	PresignedUrl pulumi.StringPtrInput
	// Snapshot identifier of the source snapshot.
	SourceDbSnapshotIdentifier pulumi.StringInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The external custom Availability Zone.
	TargetCustomAvailabilityZone pulumi.StringPtrInput
	// The Identifier for the snapshot.
	TargetDbSnapshotIdentifier pulumi.StringInput
}

func (SnapshotCopyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCopyArgs)(nil)).Elem()
}

type SnapshotCopyInput interface {
	pulumi.Input

	ToSnapshotCopyOutput() SnapshotCopyOutput
	ToSnapshotCopyOutputWithContext(ctx context.Context) SnapshotCopyOutput
}

func (*SnapshotCopy) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotCopy)(nil)).Elem()
}

func (i *SnapshotCopy) ToSnapshotCopyOutput() SnapshotCopyOutput {
	return i.ToSnapshotCopyOutputWithContext(context.Background())
}

func (i *SnapshotCopy) ToSnapshotCopyOutputWithContext(ctx context.Context) SnapshotCopyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotCopyOutput)
}

// SnapshotCopyArrayInput is an input type that accepts SnapshotCopyArray and SnapshotCopyArrayOutput values.
// You can construct a concrete instance of `SnapshotCopyArrayInput` via:
//
//	SnapshotCopyArray{ SnapshotCopyArgs{...} }
type SnapshotCopyArrayInput interface {
	pulumi.Input

	ToSnapshotCopyArrayOutput() SnapshotCopyArrayOutput
	ToSnapshotCopyArrayOutputWithContext(context.Context) SnapshotCopyArrayOutput
}

type SnapshotCopyArray []SnapshotCopyInput

func (SnapshotCopyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotCopy)(nil)).Elem()
}

func (i SnapshotCopyArray) ToSnapshotCopyArrayOutput() SnapshotCopyArrayOutput {
	return i.ToSnapshotCopyArrayOutputWithContext(context.Background())
}

func (i SnapshotCopyArray) ToSnapshotCopyArrayOutputWithContext(ctx context.Context) SnapshotCopyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotCopyArrayOutput)
}

// SnapshotCopyMapInput is an input type that accepts SnapshotCopyMap and SnapshotCopyMapOutput values.
// You can construct a concrete instance of `SnapshotCopyMapInput` via:
//
//	SnapshotCopyMap{ "key": SnapshotCopyArgs{...} }
type SnapshotCopyMapInput interface {
	pulumi.Input

	ToSnapshotCopyMapOutput() SnapshotCopyMapOutput
	ToSnapshotCopyMapOutputWithContext(context.Context) SnapshotCopyMapOutput
}

type SnapshotCopyMap map[string]SnapshotCopyInput

func (SnapshotCopyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotCopy)(nil)).Elem()
}

func (i SnapshotCopyMap) ToSnapshotCopyMapOutput() SnapshotCopyMapOutput {
	return i.ToSnapshotCopyMapOutputWithContext(context.Background())
}

func (i SnapshotCopyMap) ToSnapshotCopyMapOutputWithContext(ctx context.Context) SnapshotCopyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotCopyMapOutput)
}

type SnapshotCopyOutput struct{ *pulumi.OutputState }

func (SnapshotCopyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotCopy)(nil)).Elem()
}

func (o SnapshotCopyOutput) ToSnapshotCopyOutput() SnapshotCopyOutput {
	return o
}

func (o SnapshotCopyOutput) ToSnapshotCopyOutputWithContext(ctx context.Context) SnapshotCopyOutput {
	return o
}

// Specifies the allocated storage size in gigabytes (GB).
func (o SnapshotCopyOutput) AllocatedStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.IntOutput { return v.AllocatedStorage }).(pulumi.IntOutput)
}

// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
func (o SnapshotCopyOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Whether to copy existing tags. Defaults to `false`.
func (o SnapshotCopyOutput) CopyTags() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.BoolPtrOutput { return v.CopyTags }).(pulumi.BoolPtrOutput)
}

// The Amazon Resource Name (ARN) for the DB snapshot.
func (o SnapshotCopyOutput) DbSnapshotArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.DbSnapshotArn }).(pulumi.StringOutput)
}

// The Destination region to place snapshot copy.
func (o SnapshotCopyOutput) DestinationRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringPtrOutput { return v.DestinationRegion }).(pulumi.StringPtrOutput)
}

// Specifies whether the DB snapshot is encrypted.
func (o SnapshotCopyOutput) Encrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.BoolOutput { return v.Encrypted }).(pulumi.BoolOutput)
}

// Specifies the name of the database engine.
func (o SnapshotCopyOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Specifies the version of the database engine.
func (o SnapshotCopyOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
func (o SnapshotCopyOutput) Iops() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.IntOutput { return v.Iops }).(pulumi.IntOutput)
}

// KMS key ID.
func (o SnapshotCopyOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// License model information for the restored DB instance.
func (o SnapshotCopyOutput) LicenseModel() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.LicenseModel }).(pulumi.StringOutput)
}

// The name of an option group to associate with the copy of the snapshot.
func (o SnapshotCopyOutput) OptionGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.OptionGroupName }).(pulumi.StringOutput)
}

func (o SnapshotCopyOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// he URL that contains a Signature Version 4 signed request.
func (o SnapshotCopyOutput) PresignedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringPtrOutput { return v.PresignedUrl }).(pulumi.StringPtrOutput)
}

func (o SnapshotCopyOutput) SnapshotType() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.SnapshotType }).(pulumi.StringOutput)
}

// Snapshot identifier of the source snapshot.
func (o SnapshotCopyOutput) SourceDbSnapshotIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.SourceDbSnapshotIdentifier }).(pulumi.StringOutput)
}

// The region that the DB snapshot was created in or copied from.
func (o SnapshotCopyOutput) SourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.SourceRegion }).(pulumi.StringOutput)
}

// Specifies the storage type associated with DB snapshot.
func (o SnapshotCopyOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SnapshotCopyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o SnapshotCopyOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The external custom Availability Zone.
func (o SnapshotCopyOutput) TargetCustomAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringPtrOutput { return v.TargetCustomAvailabilityZone }).(pulumi.StringPtrOutput)
}

// The Identifier for the snapshot.
func (o SnapshotCopyOutput) TargetDbSnapshotIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.TargetDbSnapshotIdentifier }).(pulumi.StringOutput)
}

// Provides the VPC ID associated with the DB snapshot.
func (o SnapshotCopyOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCopy) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type SnapshotCopyArrayOutput struct{ *pulumi.OutputState }

func (SnapshotCopyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotCopy)(nil)).Elem()
}

func (o SnapshotCopyArrayOutput) ToSnapshotCopyArrayOutput() SnapshotCopyArrayOutput {
	return o
}

func (o SnapshotCopyArrayOutput) ToSnapshotCopyArrayOutputWithContext(ctx context.Context) SnapshotCopyArrayOutput {
	return o
}

func (o SnapshotCopyArrayOutput) Index(i pulumi.IntInput) SnapshotCopyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotCopy {
		return vs[0].([]*SnapshotCopy)[vs[1].(int)]
	}).(SnapshotCopyOutput)
}

type SnapshotCopyMapOutput struct{ *pulumi.OutputState }

func (SnapshotCopyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotCopy)(nil)).Elem()
}

func (o SnapshotCopyMapOutput) ToSnapshotCopyMapOutput() SnapshotCopyMapOutput {
	return o
}

func (o SnapshotCopyMapOutput) ToSnapshotCopyMapOutputWithContext(ctx context.Context) SnapshotCopyMapOutput {
	return o
}

func (o SnapshotCopyMapOutput) MapIndex(k pulumi.StringInput) SnapshotCopyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotCopy {
		return vs[0].(map[string]*SnapshotCopy)[vs[1].(string)]
	}).(SnapshotCopyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotCopyInput)(nil)).Elem(), &SnapshotCopy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotCopyArrayInput)(nil)).Elem(), SnapshotCopyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotCopyMapInput)(nil)).Elem(), SnapshotCopyMap{})
	pulumi.RegisterOutputType(SnapshotCopyOutput{})
	pulumi.RegisterOutputType(SnapshotCopyArrayOutput{})
	pulumi.RegisterOutputType(SnapshotCopyMapOutput{})
}
