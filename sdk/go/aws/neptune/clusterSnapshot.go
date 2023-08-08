// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Neptune database cluster snapshot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/neptune"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := neptune.NewClusterSnapshot(ctx, "example", &neptune.ClusterSnapshotArgs{
//				DbClusterIdentifier:         pulumi.Any(aws_neptune_cluster.Example.Id),
//				DbClusterSnapshotIdentifier: pulumi.String("resourcetestsnapshot1234"),
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
//	to = aws_neptune_cluster_snapshot.example
//
//	id = "my-cluster-snapshot" } Using `pulumi import`, import `aws_neptune_cluster_snapshot` using the cluster snapshot identifier. For exampleconsole % pulumi import aws_neptune_cluster_snapshot.example my-cluster-snapshot
type ClusterSnapshot struct {
	pulumi.CustomResourceState

	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntOutput `pulumi:"allocatedStorage"`
	// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier pulumi.StringOutput `pulumi:"dbClusterIdentifier"`
	// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
	DbClusterSnapshotArn pulumi.StringOutput `pulumi:"dbClusterSnapshotArn"`
	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier pulumi.StringOutput `pulumi:"dbClusterSnapshotIdentifier"`
	// Specifies the name of the database engine.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Version of the database engine for this DB cluster snapshot.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// If storageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// License model information for the restored DB cluster.
	LicenseModel pulumi.StringOutput `pulumi:"licenseModel"`
	// Port that the DB cluster was listening on at the time of the snapshot.
	Port                       pulumi.IntOutput    `pulumi:"port"`
	SnapshotType               pulumi.StringOutput `pulumi:"snapshotType"`
	SourceDbClusterSnapshotArn pulumi.StringOutput `pulumi:"sourceDbClusterSnapshotArn"`
	// The status of this DB Cluster Snapshot.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies whether the DB cluster snapshot is encrypted.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`
	// The VPC ID associated with the DB cluster snapshot.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewClusterSnapshot registers a new resource with the given unique name, arguments, and options.
func NewClusterSnapshot(ctx *pulumi.Context,
	name string, args *ClusterSnapshotArgs, opts ...pulumi.ResourceOption) (*ClusterSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterIdentifier'")
	}
	if args.DbClusterSnapshotIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterSnapshotIdentifier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterSnapshot
	err := ctx.RegisterResource("aws:neptune/clusterSnapshot:ClusterSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterSnapshot gets an existing ClusterSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterSnapshotState, opts ...pulumi.ResourceOption) (*ClusterSnapshot, error) {
	var resource ClusterSnapshot
	err := ctx.ReadResource("aws:neptune/clusterSnapshot:ClusterSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterSnapshot resources.
type clusterSnapshotState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage *int `pulumi:"allocatedStorage"`
	// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier *string `pulumi:"dbClusterIdentifier"`
	// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
	DbClusterSnapshotArn *string `pulumi:"dbClusterSnapshotArn"`
	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier *string `pulumi:"dbClusterSnapshotIdentifier"`
	// Specifies the name of the database engine.
	Engine *string `pulumi:"engine"`
	// Version of the database engine for this DB cluster snapshot.
	EngineVersion *string `pulumi:"engineVersion"`
	// If storageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// License model information for the restored DB cluster.
	LicenseModel *string `pulumi:"licenseModel"`
	// Port that the DB cluster was listening on at the time of the snapshot.
	Port                       *int    `pulumi:"port"`
	SnapshotType               *string `pulumi:"snapshotType"`
	SourceDbClusterSnapshotArn *string `pulumi:"sourceDbClusterSnapshotArn"`
	// The status of this DB Cluster Snapshot.
	Status *string `pulumi:"status"`
	// Specifies whether the DB cluster snapshot is encrypted.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
	// The VPC ID associated with the DB cluster snapshot.
	VpcId *string `pulumi:"vpcId"`
}

type ClusterSnapshotState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntPtrInput
	// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
	AvailabilityZones pulumi.StringArrayInput
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
	DbClusterSnapshotArn pulumi.StringPtrInput
	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier pulumi.StringPtrInput
	// Specifies the name of the database engine.
	Engine pulumi.StringPtrInput
	// Version of the database engine for this DB cluster snapshot.
	EngineVersion pulumi.StringPtrInput
	// If storageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
	KmsKeyId pulumi.StringPtrInput
	// License model information for the restored DB cluster.
	LicenseModel pulumi.StringPtrInput
	// Port that the DB cluster was listening on at the time of the snapshot.
	Port                       pulumi.IntPtrInput
	SnapshotType               pulumi.StringPtrInput
	SourceDbClusterSnapshotArn pulumi.StringPtrInput
	// The status of this DB Cluster Snapshot.
	Status pulumi.StringPtrInput
	// Specifies whether the DB cluster snapshot is encrypted.
	StorageEncrypted pulumi.BoolPtrInput
	// The VPC ID associated with the DB cluster snapshot.
	VpcId pulumi.StringPtrInput
}

func (ClusterSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSnapshotState)(nil)).Elem()
}

type clusterSnapshotArgs struct {
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier string `pulumi:"dbClusterIdentifier"`
	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier string `pulumi:"dbClusterSnapshotIdentifier"`
}

// The set of arguments for constructing a ClusterSnapshot resource.
type ClusterSnapshotArgs struct {
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier pulumi.StringInput
	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier pulumi.StringInput
}

func (ClusterSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSnapshotArgs)(nil)).Elem()
}

type ClusterSnapshotInput interface {
	pulumi.Input

	ToClusterSnapshotOutput() ClusterSnapshotOutput
	ToClusterSnapshotOutputWithContext(ctx context.Context) ClusterSnapshotOutput
}

func (*ClusterSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSnapshot)(nil)).Elem()
}

func (i *ClusterSnapshot) ToClusterSnapshotOutput() ClusterSnapshotOutput {
	return i.ToClusterSnapshotOutputWithContext(context.Background())
}

func (i *ClusterSnapshot) ToClusterSnapshotOutputWithContext(ctx context.Context) ClusterSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSnapshotOutput)
}

// ClusterSnapshotArrayInput is an input type that accepts ClusterSnapshotArray and ClusterSnapshotArrayOutput values.
// You can construct a concrete instance of `ClusterSnapshotArrayInput` via:
//
//	ClusterSnapshotArray{ ClusterSnapshotArgs{...} }
type ClusterSnapshotArrayInput interface {
	pulumi.Input

	ToClusterSnapshotArrayOutput() ClusterSnapshotArrayOutput
	ToClusterSnapshotArrayOutputWithContext(context.Context) ClusterSnapshotArrayOutput
}

type ClusterSnapshotArray []ClusterSnapshotInput

func (ClusterSnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterSnapshot)(nil)).Elem()
}

func (i ClusterSnapshotArray) ToClusterSnapshotArrayOutput() ClusterSnapshotArrayOutput {
	return i.ToClusterSnapshotArrayOutputWithContext(context.Background())
}

func (i ClusterSnapshotArray) ToClusterSnapshotArrayOutputWithContext(ctx context.Context) ClusterSnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSnapshotArrayOutput)
}

// ClusterSnapshotMapInput is an input type that accepts ClusterSnapshotMap and ClusterSnapshotMapOutput values.
// You can construct a concrete instance of `ClusterSnapshotMapInput` via:
//
//	ClusterSnapshotMap{ "key": ClusterSnapshotArgs{...} }
type ClusterSnapshotMapInput interface {
	pulumi.Input

	ToClusterSnapshotMapOutput() ClusterSnapshotMapOutput
	ToClusterSnapshotMapOutputWithContext(context.Context) ClusterSnapshotMapOutput
}

type ClusterSnapshotMap map[string]ClusterSnapshotInput

func (ClusterSnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterSnapshot)(nil)).Elem()
}

func (i ClusterSnapshotMap) ToClusterSnapshotMapOutput() ClusterSnapshotMapOutput {
	return i.ToClusterSnapshotMapOutputWithContext(context.Background())
}

func (i ClusterSnapshotMap) ToClusterSnapshotMapOutputWithContext(ctx context.Context) ClusterSnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSnapshotMapOutput)
}

type ClusterSnapshotOutput struct{ *pulumi.OutputState }

func (ClusterSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSnapshot)(nil)).Elem()
}

func (o ClusterSnapshotOutput) ToClusterSnapshotOutput() ClusterSnapshotOutput {
	return o
}

func (o ClusterSnapshotOutput) ToClusterSnapshotOutputWithContext(ctx context.Context) ClusterSnapshotOutput {
	return o
}

// Specifies the allocated storage size in gigabytes (GB).
func (o ClusterSnapshotOutput) AllocatedStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.IntOutput { return v.AllocatedStorage }).(pulumi.IntOutput)
}

// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
func (o ClusterSnapshotOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The DB Cluster Identifier from which to take the snapshot.
func (o ClusterSnapshotOutput) DbClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.DbClusterIdentifier }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
func (o ClusterSnapshotOutput) DbClusterSnapshotArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.DbClusterSnapshotArn }).(pulumi.StringOutput)
}

// The Identifier for the snapshot.
func (o ClusterSnapshotOutput) DbClusterSnapshotIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.DbClusterSnapshotIdentifier }).(pulumi.StringOutput)
}

// Specifies the name of the database engine.
func (o ClusterSnapshotOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Version of the database engine for this DB cluster snapshot.
func (o ClusterSnapshotOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// If storageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
func (o ClusterSnapshotOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

// License model information for the restored DB cluster.
func (o ClusterSnapshotOutput) LicenseModel() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.LicenseModel }).(pulumi.StringOutput)
}

// Port that the DB cluster was listening on at the time of the snapshot.
func (o ClusterSnapshotOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o ClusterSnapshotOutput) SnapshotType() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.SnapshotType }).(pulumi.StringOutput)
}

func (o ClusterSnapshotOutput) SourceDbClusterSnapshotArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.SourceDbClusterSnapshotArn }).(pulumi.StringOutput)
}

// The status of this DB Cluster Snapshot.
func (o ClusterSnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies whether the DB cluster snapshot is encrypted.
func (o ClusterSnapshotOutput) StorageEncrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.BoolOutput { return v.StorageEncrypted }).(pulumi.BoolOutput)
}

// The VPC ID associated with the DB cluster snapshot.
func (o ClusterSnapshotOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSnapshot) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ClusterSnapshotArrayOutput struct{ *pulumi.OutputState }

func (ClusterSnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterSnapshot)(nil)).Elem()
}

func (o ClusterSnapshotArrayOutput) ToClusterSnapshotArrayOutput() ClusterSnapshotArrayOutput {
	return o
}

func (o ClusterSnapshotArrayOutput) ToClusterSnapshotArrayOutputWithContext(ctx context.Context) ClusterSnapshotArrayOutput {
	return o
}

func (o ClusterSnapshotArrayOutput) Index(i pulumi.IntInput) ClusterSnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterSnapshot {
		return vs[0].([]*ClusterSnapshot)[vs[1].(int)]
	}).(ClusterSnapshotOutput)
}

type ClusterSnapshotMapOutput struct{ *pulumi.OutputState }

func (ClusterSnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterSnapshot)(nil)).Elem()
}

func (o ClusterSnapshotMapOutput) ToClusterSnapshotMapOutput() ClusterSnapshotMapOutput {
	return o
}

func (o ClusterSnapshotMapOutput) ToClusterSnapshotMapOutputWithContext(ctx context.Context) ClusterSnapshotMapOutput {
	return o
}

func (o ClusterSnapshotMapOutput) MapIndex(k pulumi.StringInput) ClusterSnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterSnapshot {
		return vs[0].(map[string]*ClusterSnapshot)[vs[1].(string)]
	}).(ClusterSnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSnapshotInput)(nil)).Elem(), &ClusterSnapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSnapshotArrayInput)(nil)).Elem(), ClusterSnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSnapshotMapInput)(nil)).Elem(), ClusterSnapshotMap{})
	pulumi.RegisterOutputType(ClusterSnapshotOutput{})
	pulumi.RegisterOutputType(ClusterSnapshotArrayOutput{})
	pulumi.RegisterOutputType(ClusterSnapshotMapOutput{})
}
