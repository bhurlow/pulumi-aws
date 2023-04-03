// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds permission to create volumes off of a given EBS Snapshot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ebs"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
//				AvailabilityZone: pulumi.String("us-west-2a"),
//				Size:             pulumi.Int(40),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSnapshot, err := ebs.NewSnapshot(ctx, "exampleSnapshot", &ebs.SnapshotArgs{
//				VolumeId: example.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewSnapshotCreateVolumePermission(ctx, "examplePerm", &ec2.SnapshotCreateVolumePermissionArgs{
//				SnapshotId: exampleSnapshot.ID(),
//				AccountId:  pulumi.String("12345678"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SnapshotCreateVolumePermission struct {
	pulumi.CustomResourceState

	// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// A snapshot ID
	SnapshotId pulumi.StringOutput `pulumi:"snapshotId"`
}

// NewSnapshotCreateVolumePermission registers a new resource with the given unique name, arguments, and options.
func NewSnapshotCreateVolumePermission(ctx *pulumi.Context,
	name string, args *SnapshotCreateVolumePermissionArgs, opts ...pulumi.ResourceOption) (*SnapshotCreateVolumePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.SnapshotId == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotId'")
	}
	var resource SnapshotCreateVolumePermission
	err := ctx.RegisterResource("aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotCreateVolumePermission gets an existing SnapshotCreateVolumePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotCreateVolumePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotCreateVolumePermissionState, opts ...pulumi.ResourceOption) (*SnapshotCreateVolumePermission, error) {
	var resource SnapshotCreateVolumePermission
	err := ctx.ReadResource("aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotCreateVolumePermission resources.
type snapshotCreateVolumePermissionState struct {
	// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
	AccountId *string `pulumi:"accountId"`
	// A snapshot ID
	SnapshotId *string `pulumi:"snapshotId"`
}

type SnapshotCreateVolumePermissionState struct {
	// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
	AccountId pulumi.StringPtrInput
	// A snapshot ID
	SnapshotId pulumi.StringPtrInput
}

func (SnapshotCreateVolumePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCreateVolumePermissionState)(nil)).Elem()
}

type snapshotCreateVolumePermissionArgs struct {
	// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
	AccountId string `pulumi:"accountId"`
	// A snapshot ID
	SnapshotId string `pulumi:"snapshotId"`
}

// The set of arguments for constructing a SnapshotCreateVolumePermission resource.
type SnapshotCreateVolumePermissionArgs struct {
	// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
	AccountId pulumi.StringInput
	// A snapshot ID
	SnapshotId pulumi.StringInput
}

func (SnapshotCreateVolumePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCreateVolumePermissionArgs)(nil)).Elem()
}

type SnapshotCreateVolumePermissionInput interface {
	pulumi.Input

	ToSnapshotCreateVolumePermissionOutput() SnapshotCreateVolumePermissionOutput
	ToSnapshotCreateVolumePermissionOutputWithContext(ctx context.Context) SnapshotCreateVolumePermissionOutput
}

func (*SnapshotCreateVolumePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotCreateVolumePermission)(nil)).Elem()
}

func (i *SnapshotCreateVolumePermission) ToSnapshotCreateVolumePermissionOutput() SnapshotCreateVolumePermissionOutput {
	return i.ToSnapshotCreateVolumePermissionOutputWithContext(context.Background())
}

func (i *SnapshotCreateVolumePermission) ToSnapshotCreateVolumePermissionOutputWithContext(ctx context.Context) SnapshotCreateVolumePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotCreateVolumePermissionOutput)
}

// SnapshotCreateVolumePermissionArrayInput is an input type that accepts SnapshotCreateVolumePermissionArray and SnapshotCreateVolumePermissionArrayOutput values.
// You can construct a concrete instance of `SnapshotCreateVolumePermissionArrayInput` via:
//
//	SnapshotCreateVolumePermissionArray{ SnapshotCreateVolumePermissionArgs{...} }
type SnapshotCreateVolumePermissionArrayInput interface {
	pulumi.Input

	ToSnapshotCreateVolumePermissionArrayOutput() SnapshotCreateVolumePermissionArrayOutput
	ToSnapshotCreateVolumePermissionArrayOutputWithContext(context.Context) SnapshotCreateVolumePermissionArrayOutput
}

type SnapshotCreateVolumePermissionArray []SnapshotCreateVolumePermissionInput

func (SnapshotCreateVolumePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotCreateVolumePermission)(nil)).Elem()
}

func (i SnapshotCreateVolumePermissionArray) ToSnapshotCreateVolumePermissionArrayOutput() SnapshotCreateVolumePermissionArrayOutput {
	return i.ToSnapshotCreateVolumePermissionArrayOutputWithContext(context.Background())
}

func (i SnapshotCreateVolumePermissionArray) ToSnapshotCreateVolumePermissionArrayOutputWithContext(ctx context.Context) SnapshotCreateVolumePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotCreateVolumePermissionArrayOutput)
}

// SnapshotCreateVolumePermissionMapInput is an input type that accepts SnapshotCreateVolumePermissionMap and SnapshotCreateVolumePermissionMapOutput values.
// You can construct a concrete instance of `SnapshotCreateVolumePermissionMapInput` via:
//
//	SnapshotCreateVolumePermissionMap{ "key": SnapshotCreateVolumePermissionArgs{...} }
type SnapshotCreateVolumePermissionMapInput interface {
	pulumi.Input

	ToSnapshotCreateVolumePermissionMapOutput() SnapshotCreateVolumePermissionMapOutput
	ToSnapshotCreateVolumePermissionMapOutputWithContext(context.Context) SnapshotCreateVolumePermissionMapOutput
}

type SnapshotCreateVolumePermissionMap map[string]SnapshotCreateVolumePermissionInput

func (SnapshotCreateVolumePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotCreateVolumePermission)(nil)).Elem()
}

func (i SnapshotCreateVolumePermissionMap) ToSnapshotCreateVolumePermissionMapOutput() SnapshotCreateVolumePermissionMapOutput {
	return i.ToSnapshotCreateVolumePermissionMapOutputWithContext(context.Background())
}

func (i SnapshotCreateVolumePermissionMap) ToSnapshotCreateVolumePermissionMapOutputWithContext(ctx context.Context) SnapshotCreateVolumePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotCreateVolumePermissionMapOutput)
}

type SnapshotCreateVolumePermissionOutput struct{ *pulumi.OutputState }

func (SnapshotCreateVolumePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotCreateVolumePermission)(nil)).Elem()
}

func (o SnapshotCreateVolumePermissionOutput) ToSnapshotCreateVolumePermissionOutput() SnapshotCreateVolumePermissionOutput {
	return o
}

func (o SnapshotCreateVolumePermissionOutput) ToSnapshotCreateVolumePermissionOutputWithContext(ctx context.Context) SnapshotCreateVolumePermissionOutput {
	return o
}

// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
func (o SnapshotCreateVolumePermissionOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCreateVolumePermission) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// A snapshot ID
func (o SnapshotCreateVolumePermissionOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotCreateVolumePermission) pulumi.StringOutput { return v.SnapshotId }).(pulumi.StringOutput)
}

type SnapshotCreateVolumePermissionArrayOutput struct{ *pulumi.OutputState }

func (SnapshotCreateVolumePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotCreateVolumePermission)(nil)).Elem()
}

func (o SnapshotCreateVolumePermissionArrayOutput) ToSnapshotCreateVolumePermissionArrayOutput() SnapshotCreateVolumePermissionArrayOutput {
	return o
}

func (o SnapshotCreateVolumePermissionArrayOutput) ToSnapshotCreateVolumePermissionArrayOutputWithContext(ctx context.Context) SnapshotCreateVolumePermissionArrayOutput {
	return o
}

func (o SnapshotCreateVolumePermissionArrayOutput) Index(i pulumi.IntInput) SnapshotCreateVolumePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotCreateVolumePermission {
		return vs[0].([]*SnapshotCreateVolumePermission)[vs[1].(int)]
	}).(SnapshotCreateVolumePermissionOutput)
}

type SnapshotCreateVolumePermissionMapOutput struct{ *pulumi.OutputState }

func (SnapshotCreateVolumePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotCreateVolumePermission)(nil)).Elem()
}

func (o SnapshotCreateVolumePermissionMapOutput) ToSnapshotCreateVolumePermissionMapOutput() SnapshotCreateVolumePermissionMapOutput {
	return o
}

func (o SnapshotCreateVolumePermissionMapOutput) ToSnapshotCreateVolumePermissionMapOutputWithContext(ctx context.Context) SnapshotCreateVolumePermissionMapOutput {
	return o
}

func (o SnapshotCreateVolumePermissionMapOutput) MapIndex(k pulumi.StringInput) SnapshotCreateVolumePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotCreateVolumePermission {
		return vs[0].(map[string]*SnapshotCreateVolumePermission)[vs[1].(string)]
	}).(SnapshotCreateVolumePermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotCreateVolumePermissionInput)(nil)).Elem(), &SnapshotCreateVolumePermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotCreateVolumePermissionArrayInput)(nil)).Elem(), SnapshotCreateVolumePermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotCreateVolumePermissionMapInput)(nil)).Elem(), SnapshotCreateVolumePermissionMap{})
	pulumi.RegisterOutputType(SnapshotCreateVolumePermissionOutput{})
	pulumi.RegisterOutputType(SnapshotCreateVolumePermissionArrayOutput{})
	pulumi.RegisterOutputType(SnapshotCreateVolumePermissionMapOutput{})
}
