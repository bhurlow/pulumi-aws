// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a replica of an existing EFS file system in the same or another region. Creating this resource causes the source EFS file system to be replicated to a new read-only destination EFS file system. Deleting this resource will cause the replication from source to destination to stop and the destination file system will no longer be read only.
//
// > **NOTE:** Deleting this resource does **not** delete the destination file system that was created.
//
// ## Example Usage
//
// Will create a replica using regional storage in us-west-2 that will be encrypted by the default EFS KMS key `/aws/elasticfilesystem`.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/efs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleFileSystem, err := efs.NewFileSystem(ctx, "exampleFileSystem", nil)
//			if err != nil {
//				return err
//			}
//			_, err = efs.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &efs.ReplicationConfigurationArgs{
//				SourceFileSystemId: exampleFileSystem.ID(),
//				Destination: &efs.ReplicationConfigurationDestinationArgs{
//					Region: pulumi.String("us-west-2"),
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
// Replica will be created as One Zone storage in the us-west-2b Availability Zone and encrypted with the specified KMS key.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/efs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleFileSystem, err := efs.NewFileSystem(ctx, "exampleFileSystem", nil)
//			if err != nil {
//				return err
//			}
//			_, err = efs.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &efs.ReplicationConfigurationArgs{
//				SourceFileSystemId: exampleFileSystem.ID(),
//				Destination: &efs.ReplicationConfigurationDestinationArgs{
//					AvailabilityZoneName: pulumi.String("us-west-2b"),
//					KmsKeyId:             pulumi.String("1234abcd-12ab-34cd-56ef-1234567890ab"),
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
// EFS Replication Configurations can be imported using the file system ID of either the source or destination file system. When importing, the `availability_zone_name` and `kms_key_id` attributes must **not** be set in the configuration. The AWS API does not return these values when querying the replication configuration and their presence will therefore show as a diff in a subsequent plan.
//
// ```sh
//
//	$ pulumi import aws:efs/replicationConfiguration:ReplicationConfiguration example fs-id
//
// ```
type ReplicationConfiguration struct {
	pulumi.CustomResourceState

	// When the replication configuration was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// A destination configuration block (documented below).
	Destination ReplicationConfigurationDestinationOutput `pulumi:"destination"`
	// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in the replication configuration.
	OriginalSourceFileSystemArn pulumi.StringOutput `pulumi:"originalSourceFileSystemArn"`
	// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
	SourceFileSystemArn pulumi.StringOutput `pulumi:"sourceFileSystemArn"`
	// The ID of the file system that is to be replicated.
	SourceFileSystemId pulumi.StringOutput `pulumi:"sourceFileSystemId"`
	// The AWS Region in which the source Amazon EFS file system is located.
	// * `destination[0].file_system_id` - The fs ID of the replica.
	// * `destination[0].status` - The status of the replication.
	SourceFileSystemRegion pulumi.StringOutput `pulumi:"sourceFileSystemRegion"`
}

// NewReplicationConfiguration registers a new resource with the given unique name, arguments, and options.
func NewReplicationConfiguration(ctx *pulumi.Context,
	name string, args *ReplicationConfigurationArgs, opts ...pulumi.ResourceOption) (*ReplicationConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.SourceFileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'SourceFileSystemId'")
	}
	var resource ReplicationConfiguration
	err := ctx.RegisterResource("aws:efs/replicationConfiguration:ReplicationConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationConfiguration gets an existing ReplicationConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationConfigurationState, opts ...pulumi.ResourceOption) (*ReplicationConfiguration, error) {
	var resource ReplicationConfiguration
	err := ctx.ReadResource("aws:efs/replicationConfiguration:ReplicationConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationConfiguration resources.
type replicationConfigurationState struct {
	// When the replication configuration was created.
	CreationTime *string `pulumi:"creationTime"`
	// A destination configuration block (documented below).
	Destination *ReplicationConfigurationDestination `pulumi:"destination"`
	// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in the replication configuration.
	OriginalSourceFileSystemArn *string `pulumi:"originalSourceFileSystemArn"`
	// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
	SourceFileSystemArn *string `pulumi:"sourceFileSystemArn"`
	// The ID of the file system that is to be replicated.
	SourceFileSystemId *string `pulumi:"sourceFileSystemId"`
	// The AWS Region in which the source Amazon EFS file system is located.
	// * `destination[0].file_system_id` - The fs ID of the replica.
	// * `destination[0].status` - The status of the replication.
	SourceFileSystemRegion *string `pulumi:"sourceFileSystemRegion"`
}

type ReplicationConfigurationState struct {
	// When the replication configuration was created.
	CreationTime pulumi.StringPtrInput
	// A destination configuration block (documented below).
	Destination ReplicationConfigurationDestinationPtrInput
	// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in the replication configuration.
	OriginalSourceFileSystemArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
	SourceFileSystemArn pulumi.StringPtrInput
	// The ID of the file system that is to be replicated.
	SourceFileSystemId pulumi.StringPtrInput
	// The AWS Region in which the source Amazon EFS file system is located.
	// * `destination[0].file_system_id` - The fs ID of the replica.
	// * `destination[0].status` - The status of the replication.
	SourceFileSystemRegion pulumi.StringPtrInput
}

func (ReplicationConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigurationState)(nil)).Elem()
}

type replicationConfigurationArgs struct {
	// A destination configuration block (documented below).
	Destination ReplicationConfigurationDestination `pulumi:"destination"`
	// The ID of the file system that is to be replicated.
	SourceFileSystemId string `pulumi:"sourceFileSystemId"`
}

// The set of arguments for constructing a ReplicationConfiguration resource.
type ReplicationConfigurationArgs struct {
	// A destination configuration block (documented below).
	Destination ReplicationConfigurationDestinationInput
	// The ID of the file system that is to be replicated.
	SourceFileSystemId pulumi.StringInput
}

func (ReplicationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigurationArgs)(nil)).Elem()
}

type ReplicationConfigurationInput interface {
	pulumi.Input

	ToReplicationConfigurationOutput() ReplicationConfigurationOutput
	ToReplicationConfigurationOutputWithContext(ctx context.Context) ReplicationConfigurationOutput
}

func (*ReplicationConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfiguration)(nil)).Elem()
}

func (i *ReplicationConfiguration) ToReplicationConfigurationOutput() ReplicationConfigurationOutput {
	return i.ToReplicationConfigurationOutputWithContext(context.Background())
}

func (i *ReplicationConfiguration) ToReplicationConfigurationOutputWithContext(ctx context.Context) ReplicationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationOutput)
}

// ReplicationConfigurationArrayInput is an input type that accepts ReplicationConfigurationArray and ReplicationConfigurationArrayOutput values.
// You can construct a concrete instance of `ReplicationConfigurationArrayInput` via:
//
//	ReplicationConfigurationArray{ ReplicationConfigurationArgs{...} }
type ReplicationConfigurationArrayInput interface {
	pulumi.Input

	ToReplicationConfigurationArrayOutput() ReplicationConfigurationArrayOutput
	ToReplicationConfigurationArrayOutputWithContext(context.Context) ReplicationConfigurationArrayOutput
}

type ReplicationConfigurationArray []ReplicationConfigurationInput

func (ReplicationConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationConfiguration)(nil)).Elem()
}

func (i ReplicationConfigurationArray) ToReplicationConfigurationArrayOutput() ReplicationConfigurationArrayOutput {
	return i.ToReplicationConfigurationArrayOutputWithContext(context.Background())
}

func (i ReplicationConfigurationArray) ToReplicationConfigurationArrayOutputWithContext(ctx context.Context) ReplicationConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationArrayOutput)
}

// ReplicationConfigurationMapInput is an input type that accepts ReplicationConfigurationMap and ReplicationConfigurationMapOutput values.
// You can construct a concrete instance of `ReplicationConfigurationMapInput` via:
//
//	ReplicationConfigurationMap{ "key": ReplicationConfigurationArgs{...} }
type ReplicationConfigurationMapInput interface {
	pulumi.Input

	ToReplicationConfigurationMapOutput() ReplicationConfigurationMapOutput
	ToReplicationConfigurationMapOutputWithContext(context.Context) ReplicationConfigurationMapOutput
}

type ReplicationConfigurationMap map[string]ReplicationConfigurationInput

func (ReplicationConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationConfiguration)(nil)).Elem()
}

func (i ReplicationConfigurationMap) ToReplicationConfigurationMapOutput() ReplicationConfigurationMapOutput {
	return i.ToReplicationConfigurationMapOutputWithContext(context.Background())
}

func (i ReplicationConfigurationMap) ToReplicationConfigurationMapOutputWithContext(ctx context.Context) ReplicationConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigurationMapOutput)
}

type ReplicationConfigurationOutput struct{ *pulumi.OutputState }

func (ReplicationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfiguration)(nil)).Elem()
}

func (o ReplicationConfigurationOutput) ToReplicationConfigurationOutput() ReplicationConfigurationOutput {
	return o
}

func (o ReplicationConfigurationOutput) ToReplicationConfigurationOutputWithContext(ctx context.Context) ReplicationConfigurationOutput {
	return o
}

// When the replication configuration was created.
func (o ReplicationConfigurationOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationConfiguration) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// A destination configuration block (documented below).
func (o ReplicationConfigurationOutput) Destination() ReplicationConfigurationDestinationOutput {
	return o.ApplyT(func(v *ReplicationConfiguration) ReplicationConfigurationDestinationOutput { return v.Destination }).(ReplicationConfigurationDestinationOutput)
}

// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in the replication configuration.
func (o ReplicationConfigurationOutput) OriginalSourceFileSystemArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationConfiguration) pulumi.StringOutput { return v.OriginalSourceFileSystemArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
func (o ReplicationConfigurationOutput) SourceFileSystemArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationConfiguration) pulumi.StringOutput { return v.SourceFileSystemArn }).(pulumi.StringOutput)
}

// The ID of the file system that is to be replicated.
func (o ReplicationConfigurationOutput) SourceFileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationConfiguration) pulumi.StringOutput { return v.SourceFileSystemId }).(pulumi.StringOutput)
}

// The AWS Region in which the source Amazon EFS file system is located.
// * `destination[0].file_system_id` - The fs ID of the replica.
// * `destination[0].status` - The status of the replication.
func (o ReplicationConfigurationOutput) SourceFileSystemRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationConfiguration) pulumi.StringOutput { return v.SourceFileSystemRegion }).(pulumi.StringOutput)
}

type ReplicationConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ReplicationConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationConfiguration)(nil)).Elem()
}

func (o ReplicationConfigurationArrayOutput) ToReplicationConfigurationArrayOutput() ReplicationConfigurationArrayOutput {
	return o
}

func (o ReplicationConfigurationArrayOutput) ToReplicationConfigurationArrayOutputWithContext(ctx context.Context) ReplicationConfigurationArrayOutput {
	return o
}

func (o ReplicationConfigurationArrayOutput) Index(i pulumi.IntInput) ReplicationConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationConfiguration {
		return vs[0].([]*ReplicationConfiguration)[vs[1].(int)]
	}).(ReplicationConfigurationOutput)
}

type ReplicationConfigurationMapOutput struct{ *pulumi.OutputState }

func (ReplicationConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationConfiguration)(nil)).Elem()
}

func (o ReplicationConfigurationMapOutput) ToReplicationConfigurationMapOutput() ReplicationConfigurationMapOutput {
	return o
}

func (o ReplicationConfigurationMapOutput) ToReplicationConfigurationMapOutputWithContext(ctx context.Context) ReplicationConfigurationMapOutput {
	return o
}

func (o ReplicationConfigurationMapOutput) MapIndex(k pulumi.StringInput) ReplicationConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationConfiguration {
		return vs[0].(map[string]*ReplicationConfiguration)[vs[1].(string)]
	}).(ReplicationConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationConfigurationInput)(nil)).Elem(), &ReplicationConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationConfigurationArrayInput)(nil)).Elem(), ReplicationConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationConfigurationMapInput)(nil)).Elem(), ReplicationConfigurationMap{})
	pulumi.RegisterOutputType(ReplicationConfigurationOutput{})
	pulumi.RegisterOutputType(ReplicationConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ReplicationConfigurationMapOutput{})
}
