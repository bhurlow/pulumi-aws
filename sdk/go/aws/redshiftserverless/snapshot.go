// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshiftserverless

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Amazon Redshift Serverless Snapshot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/redshiftserverless"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := redshiftserverless.NewSnapshot(ctx, "example", &redshiftserverless.SnapshotArgs{
//				NamespaceName: pulumi.Any(aws_redshiftserverless_workgroup.Example.Namespace_name),
//				SnapshotName:  pulumi.String("example"),
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
// Redshift Serverless Snapshots can be imported using the `snapshot_name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:redshiftserverless/snapshot:Snapshot example example
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// All of the Amazon Web Services accounts that have access to restore a snapshot to a provisioned cluster.
	AccountsWithProvisionedRestoreAccesses pulumi.StringArrayOutput `pulumi:"accountsWithProvisionedRestoreAccesses"`
	// All of the Amazon Web Services accounts that have access to restore a snapshot to a namespace.
	AccountsWithRestoreAccesses pulumi.StringArrayOutput `pulumi:"accountsWithRestoreAccesses"`
	// The username of the database within a snapshot.
	AdminUsername pulumi.StringOutput `pulumi:"adminUsername"`
	// The Amazon Resource Name (ARN) of the snapshot.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The unique identifier of the KMS key used to encrypt the snapshot.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// The Amazon Resource Name (ARN) of the namespace the snapshot was created from.
	NamespaceArn pulumi.StringOutput `pulumi:"namespaceArn"`
	// The namespace to create a snapshot for.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The owner Amazon Web Services; account of the snapshot.
	OwnerAccount pulumi.StringOutput `pulumi:"ownerAccount"`
	// How long to retain the created snapshot. Default value is `-1`.
	RetentionPeriod pulumi.IntPtrOutput `pulumi:"retentionPeriod"`
	// The name of the snapshot.
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.SnapshotName == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotName'")
	}
	var resource Snapshot
	err := ctx.RegisterResource("aws:redshiftserverless/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("aws:redshiftserverless/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// All of the Amazon Web Services accounts that have access to restore a snapshot to a provisioned cluster.
	AccountsWithProvisionedRestoreAccesses []string `pulumi:"accountsWithProvisionedRestoreAccesses"`
	// All of the Amazon Web Services accounts that have access to restore a snapshot to a namespace.
	AccountsWithRestoreAccesses []string `pulumi:"accountsWithRestoreAccesses"`
	// The username of the database within a snapshot.
	AdminUsername *string `pulumi:"adminUsername"`
	// The Amazon Resource Name (ARN) of the snapshot.
	Arn *string `pulumi:"arn"`
	// The unique identifier of the KMS key used to encrypt the snapshot.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The Amazon Resource Name (ARN) of the namespace the snapshot was created from.
	NamespaceArn *string `pulumi:"namespaceArn"`
	// The namespace to create a snapshot for.
	NamespaceName *string `pulumi:"namespaceName"`
	// The owner Amazon Web Services; account of the snapshot.
	OwnerAccount *string `pulumi:"ownerAccount"`
	// How long to retain the created snapshot. Default value is `-1`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The name of the snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
}

type SnapshotState struct {
	// All of the Amazon Web Services accounts that have access to restore a snapshot to a provisioned cluster.
	AccountsWithProvisionedRestoreAccesses pulumi.StringArrayInput
	// All of the Amazon Web Services accounts that have access to restore a snapshot to a namespace.
	AccountsWithRestoreAccesses pulumi.StringArrayInput
	// The username of the database within a snapshot.
	AdminUsername pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the snapshot.
	Arn pulumi.StringPtrInput
	// The unique identifier of the KMS key used to encrypt the snapshot.
	KmsKeyId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the namespace the snapshot was created from.
	NamespaceArn pulumi.StringPtrInput
	// The namespace to create a snapshot for.
	NamespaceName pulumi.StringPtrInput
	// The owner Amazon Web Services; account of the snapshot.
	OwnerAccount pulumi.StringPtrInput
	// How long to retain the created snapshot. Default value is `-1`.
	RetentionPeriod pulumi.IntPtrInput
	// The name of the snapshot.
	SnapshotName pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The namespace to create a snapshot for.
	NamespaceName string `pulumi:"namespaceName"`
	// How long to retain the created snapshot. Default value is `-1`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The name of the snapshot.
	SnapshotName string `pulumi:"snapshotName"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The namespace to create a snapshot for.
	NamespaceName pulumi.StringInput
	// How long to retain the created snapshot. Default value is `-1`.
	RetentionPeriod pulumi.IntPtrInput
	// The name of the snapshot.
	SnapshotName pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// All of the Amazon Web Services accounts that have access to restore a snapshot to a provisioned cluster.
func (o SnapshotOutput) AccountsWithProvisionedRestoreAccesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.AccountsWithProvisionedRestoreAccesses }).(pulumi.StringArrayOutput)
}

// All of the Amazon Web Services accounts that have access to restore a snapshot to a namespace.
func (o SnapshotOutput) AccountsWithRestoreAccesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.AccountsWithRestoreAccesses }).(pulumi.StringArrayOutput)
}

// The username of the database within a snapshot.
func (o SnapshotOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.AdminUsername }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the snapshot.
func (o SnapshotOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The unique identifier of the KMS key used to encrypt the snapshot.
func (o SnapshotOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the namespace the snapshot was created from.
func (o SnapshotOutput) NamespaceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.NamespaceArn }).(pulumi.StringOutput)
}

// The namespace to create a snapshot for.
func (o SnapshotOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.NamespaceName }).(pulumi.StringOutput)
}

// The owner Amazon Web Services; account of the snapshot.
func (o SnapshotOutput) OwnerAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OwnerAccount }).(pulumi.StringOutput)
}

// How long to retain the created snapshot. Default value is `-1`.
func (o SnapshotOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntPtrOutput { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

// The name of the snapshot.
func (o SnapshotOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
