// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a DB Snapshot for use when provisioning DB instances
//
// > **NOTE:** This data source does not apply to snapshots created on Aurora DB clusters.
// See the `rds.ClusterSnapshot` data source for DB Cluster snapshots.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			prod, err := rds.NewInstance(ctx, "prod", &rds.InstanceArgs{
//				AllocatedStorage:   pulumi.Int(10),
//				Engine:             pulumi.String("mysql"),
//				EngineVersion:      pulumi.String("5.6.17"),
//				InstanceClass:      pulumi.String("db.t2.micro"),
//				DbName:             pulumi.String("mydb"),
//				Username:           pulumi.String("foo"),
//				Password:           pulumi.String("bar"),
//				DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
//				ParameterGroupName: pulumi.String("default.mysql5.6"),
//			})
//			if err != nil {
//				return err
//			}
//			latestProdSnapshot := rds.LookupSnapshotOutput(ctx, rds.GetSnapshotOutputArgs{
//				DbInstanceIdentifier: prod.Identifier,
//				MostRecent:           pulumi.Bool(true),
//			}, nil)
//			_, err = rds.NewInstance(ctx, "dev", &rds.InstanceArgs{
//				InstanceClass: pulumi.String("db.t2.micro"),
//				DbName:        pulumi.String("mydbdev"),
//				SnapshotIdentifier: latestProdSnapshot.ApplyT(func(latestProdSnapshot rds.GetSnapshotResult) (*string, error) {
//					return &latestProdSnapshot.Id, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSnapshotResult
	err := ctx.Invoke("aws:rds/getSnapshot:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotArgs struct {
	// Returns the list of snapshots created by the specific db_instance
	DbInstanceIdentifier *string `pulumi:"dbInstanceIdentifier"`
	// Returns information on a specific snapshot_id.
	DbSnapshotIdentifier *string `pulumi:"dbSnapshotIdentifier"`
	// Set this value to true to include manual DB snapshots that are public and can be
	// copied or restored by any AWS account, otherwise set this value to false. The default is `false`.
	// `tags` - (Optional) Mapping of tags, each pair of which must exactly match
	// a pair on the desired DB snapshot.
	IncludePublic *bool `pulumi:"includePublic"`
	// Set this value to true to include shared manual DB snapshots from other
	// AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false.
	// The default is `false`.
	IncludeShared *bool `pulumi:"includeShared"`
	// If more than one result is returned, use the most
	// recent Snapshot.
	MostRecent *bool `pulumi:"mostRecent"`
	// Type of snapshots to be returned. If you don't specify a SnapshotType
	// value, then both automated and manual snapshots are returned. Shared and public DB snapshots are not
	// included in the returned results by default. Possible values are, `automated`, `manual`, `shared`, `public` and `awsbackup`.
	SnapshotType *string           `pulumi:"snapshotType"`
	Tags         map[string]string `pulumi:"tags"`
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResult struct {
	// Allocated storage size in gigabytes (GB).
	AllocatedStorage int `pulumi:"allocatedStorage"`
	// Name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone     string  `pulumi:"availabilityZone"`
	DbInstanceIdentifier *string `pulumi:"dbInstanceIdentifier"`
	// ARN for the DB snapshot.
	DbSnapshotArn        string  `pulumi:"dbSnapshotArn"`
	DbSnapshotIdentifier *string `pulumi:"dbSnapshotIdentifier"`
	// Whether the DB snapshot is encrypted.
	Encrypted bool `pulumi:"encrypted"`
	// Name of the database engine.
	Engine string `pulumi:"engine"`
	// Version of the database engine.
	EngineVersion string `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IncludePublic *bool  `pulumi:"includePublic"`
	IncludeShared *bool  `pulumi:"includeShared"`
	// Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops int `pulumi:"iops"`
	// ARN for the KMS encryption key.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// License model information for the restored DB instance.
	LicenseModel string `pulumi:"licenseModel"`
	MostRecent   *bool  `pulumi:"mostRecent"`
	// Provides the option group name for the DB snapshot.
	OptionGroupName string `pulumi:"optionGroupName"`
	Port            int    `pulumi:"port"`
	// Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).
	SnapshotCreateTime string  `pulumi:"snapshotCreateTime"`
	SnapshotType       *string `pulumi:"snapshotType"`
	// DB snapshot ARN that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDbSnapshotIdentifier string `pulumi:"sourceDbSnapshotIdentifier"`
	// Region that the DB snapshot was created in or copied from.
	SourceRegion string `pulumi:"sourceRegion"`
	// Status of this DB snapshot.
	Status string `pulumi:"status"`
	// Storage type associated with DB snapshot.
	StorageType string            `pulumi:"storageType"`
	Tags        map[string]string `pulumi:"tags"`
	// ID of the VPC associated with the DB snapshot.
	VpcId string `pulumi:"vpcId"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			var s LookupSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotResultOutput)
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotOutputArgs struct {
	// Returns the list of snapshots created by the specific db_instance
	DbInstanceIdentifier pulumi.StringPtrInput `pulumi:"dbInstanceIdentifier"`
	// Returns information on a specific snapshot_id.
	DbSnapshotIdentifier pulumi.StringPtrInput `pulumi:"dbSnapshotIdentifier"`
	// Set this value to true to include manual DB snapshots that are public and can be
	// copied or restored by any AWS account, otherwise set this value to false. The default is `false`.
	// `tags` - (Optional) Mapping of tags, each pair of which must exactly match
	// a pair on the desired DB snapshot.
	IncludePublic pulumi.BoolPtrInput `pulumi:"includePublic"`
	// Set this value to true to include shared manual DB snapshots from other
	// AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false.
	// The default is `false`.
	IncludeShared pulumi.BoolPtrInput `pulumi:"includeShared"`
	// If more than one result is returned, use the most
	// recent Snapshot.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// Type of snapshots to be returned. If you don't specify a SnapshotType
	// value, then both automated and manual snapshots are returned. Shared and public DB snapshots are not
	// included in the returned results by default. Possible values are, `automated`, `manual`, `shared`, `public` and `awsbackup`.
	SnapshotType pulumi.StringPtrInput `pulumi:"snapshotType"`
	Tags         pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

// Allocated storage size in gigabytes (GB).
func (o LookupSnapshotResultOutput) AllocatedStorage() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSnapshotResult) int { return v.AllocatedStorage }).(pulumi.IntOutput)
}

// Name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
func (o LookupSnapshotResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) DbInstanceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.DbInstanceIdentifier }).(pulumi.StringPtrOutput)
}

// ARN for the DB snapshot.
func (o LookupSnapshotResultOutput) DbSnapshotArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.DbSnapshotArn }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) DbSnapshotIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.DbSnapshotIdentifier }).(pulumi.StringPtrOutput)
}

// Whether the DB snapshot is encrypted.
func (o LookupSnapshotResultOutput) Encrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.Encrypted }).(pulumi.BoolOutput)
}

// Name of the database engine.
func (o LookupSnapshotResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Engine }).(pulumi.StringOutput)
}

// Version of the database engine.
func (o LookupSnapshotResultOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.EngineVersion }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) IncludePublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *bool { return v.IncludePublic }).(pulumi.BoolPtrOutput)
}

func (o LookupSnapshotResultOutput) IncludeShared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *bool { return v.IncludeShared }).(pulumi.BoolPtrOutput)
}

// Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
func (o LookupSnapshotResultOutput) Iops() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSnapshotResult) int { return v.Iops }).(pulumi.IntOutput)
}

// ARN for the KMS encryption key.
func (o LookupSnapshotResultOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.KmsKeyId }).(pulumi.StringOutput)
}

// License model information for the restored DB instance.
func (o LookupSnapshotResultOutput) LicenseModel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.LicenseModel }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// Provides the option group name for the DB snapshot.
func (o LookupSnapshotResultOutput) OptionGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.OptionGroupName }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSnapshotResult) int { return v.Port }).(pulumi.IntOutput)
}

// Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).
func (o LookupSnapshotResultOutput) SnapshotCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.SnapshotCreateTime }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

// DB snapshot ARN that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
func (o LookupSnapshotResultOutput) SourceDbSnapshotIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.SourceDbSnapshotIdentifier }).(pulumi.StringOutput)
}

// Region that the DB snapshot was created in or copied from.
func (o LookupSnapshotResultOutput) SourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.SourceRegion }).(pulumi.StringOutput)
}

// Status of this DB snapshot.
func (o LookupSnapshotResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Status }).(pulumi.StringOutput)
}

// Storage type associated with DB snapshot.
func (o LookupSnapshotResultOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.StorageType }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSnapshotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// ID of the VPC associated with the DB snapshot.
func (o LookupSnapshotResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
