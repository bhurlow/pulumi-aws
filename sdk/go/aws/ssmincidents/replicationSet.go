// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmincidents

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for managing a replication set in AWS Systems Manager Incident Manager.
//
// > **NOTE:** Deleting a replication set also deletes all Incident Manager related data including response plans, incident records, contacts and escalation plans.
//
// ## Example Usage
// ### Basic Usage
//
// Create a replication set.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
//				Regions: ssmincidents.ReplicationSetRegionArray{
//					&ssmincidents.ReplicationSetRegionArgs{
//						Name: pulumi.String("us-west-2"),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"exampleTag": pulumi.String("exampleValue"),
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
// Add a Region to a replication set. (You can add only one Region at a time.)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
//				Regions: ssmincidents.ReplicationSetRegionArray{
//					&ssmincidents.ReplicationSetRegionArgs{
//						Name: pulumi.String("us-west-2"),
//					},
//					&ssmincidents.ReplicationSetRegionArgs{
//						Name: pulumi.String("ap-southeast-2"),
//					},
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
// Delete a Region from a replication set. (You can delete only one Region at a time.)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
//				Regions: ssmincidents.ReplicationSetRegionArray{
//					&ssmincidents.ReplicationSetRegionArgs{
//						Name: pulumi.String("us-west-2"),
//					},
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
// ## Basic Usage with an AWS Customer Managed Key
//
// Create a replication set with an AWS Key Management Service (AWS KMS) customer manager key:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleKey, err := kms.NewKey(ctx, "exampleKey", nil)
//			if err != nil {
//				return err
//			}
//			_, err = ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
//				Regions: ssmincidents.ReplicationSetRegionArray{
//					&ssmincidents.ReplicationSetRegionArgs{
//						Name:      pulumi.String("us-west-2"),
//						KmsKeyArn: exampleKey.Arn,
//					},
//				},
//				Tags: pulumi.StringMap{
//					"exampleTag": pulumi.String("exampleValue"),
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
// terraform import {
//
//	to = aws_ssmincidents_replication_set.replicationSetName
//
//	id = "import" } Using `pulumi import`, import an Incident Manager replication. For exampleconsole % pulumi import aws_ssmincidents_replication_set.replicationSetName import
type ReplicationSet struct {
	pulumi.CustomResourceState

	// The ARN of the replication set.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the user who created the replication set.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// If `true`, the last region in a replication set cannot be deleted.
	DeletionProtected pulumi.BoolOutput `pulumi:"deletionProtected"`
	// A timestamp showing when the replication set was last modified.
	LastModifiedBy pulumi.StringOutput             `pulumi:"lastModifiedBy"`
	Regions        ReplicationSetRegionArrayOutput `pulumi:"regions"`
	// The current status of the Region.
	// * Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags applied to the replication set.
	//
	// For information about the maximum allowed number of Regions and tag value constraints, see [CreateReplicationSet in the *AWS Systems Manager Incident Manager API Reference*](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateReplicationSet.html).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewReplicationSet registers a new resource with the given unique name, arguments, and options.
func NewReplicationSet(ctx *pulumi.Context,
	name string, args *ReplicationSetArgs, opts ...pulumi.ResourceOption) (*ReplicationSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplicationSet
	err := ctx.RegisterResource("aws:ssmincidents/replicationSet:ReplicationSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationSet gets an existing ReplicationSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationSetState, opts ...pulumi.ResourceOption) (*ReplicationSet, error) {
	var resource ReplicationSet
	err := ctx.ReadResource("aws:ssmincidents/replicationSet:ReplicationSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationSet resources.
type replicationSetState struct {
	// The ARN of the replication set.
	Arn *string `pulumi:"arn"`
	// The ARN of the user who created the replication set.
	CreatedBy *string `pulumi:"createdBy"`
	// If `true`, the last region in a replication set cannot be deleted.
	DeletionProtected *bool `pulumi:"deletionProtected"`
	// A timestamp showing when the replication set was last modified.
	LastModifiedBy *string                `pulumi:"lastModifiedBy"`
	Regions        []ReplicationSetRegion `pulumi:"regions"`
	// The current status of the Region.
	// * Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
	Status *string `pulumi:"status"`
	// Tags applied to the replication set.
	//
	// For information about the maximum allowed number of Regions and tag value constraints, see [CreateReplicationSet in the *AWS Systems Manager Incident Manager API Reference*](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateReplicationSet.html).
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ReplicationSetState struct {
	// The ARN of the replication set.
	Arn pulumi.StringPtrInput
	// The ARN of the user who created the replication set.
	CreatedBy pulumi.StringPtrInput
	// If `true`, the last region in a replication set cannot be deleted.
	DeletionProtected pulumi.BoolPtrInput
	// A timestamp showing when the replication set was last modified.
	LastModifiedBy pulumi.StringPtrInput
	Regions        ReplicationSetRegionArrayInput
	// The current status of the Region.
	// * Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
	Status pulumi.StringPtrInput
	// Tags applied to the replication set.
	//
	// For information about the maximum allowed number of Regions and tag value constraints, see [CreateReplicationSet in the *AWS Systems Manager Incident Manager API Reference*](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateReplicationSet.html).
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ReplicationSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSetState)(nil)).Elem()
}

type replicationSetArgs struct {
	Regions []ReplicationSetRegion `pulumi:"regions"`
	// Tags applied to the replication set.
	//
	// For information about the maximum allowed number of Regions and tag value constraints, see [CreateReplicationSet in the *AWS Systems Manager Incident Manager API Reference*](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateReplicationSet.html).
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ReplicationSet resource.
type ReplicationSetArgs struct {
	Regions ReplicationSetRegionArrayInput
	// Tags applied to the replication set.
	//
	// For information about the maximum allowed number of Regions and tag value constraints, see [CreateReplicationSet in the *AWS Systems Manager Incident Manager API Reference*](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateReplicationSet.html).
	Tags pulumi.StringMapInput
}

func (ReplicationSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSetArgs)(nil)).Elem()
}

type ReplicationSetInput interface {
	pulumi.Input

	ToReplicationSetOutput() ReplicationSetOutput
	ToReplicationSetOutputWithContext(ctx context.Context) ReplicationSetOutput
}

func (*ReplicationSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSet)(nil)).Elem()
}

func (i *ReplicationSet) ToReplicationSetOutput() ReplicationSetOutput {
	return i.ToReplicationSetOutputWithContext(context.Background())
}

func (i *ReplicationSet) ToReplicationSetOutputWithContext(ctx context.Context) ReplicationSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSetOutput)
}

// ReplicationSetArrayInput is an input type that accepts ReplicationSetArray and ReplicationSetArrayOutput values.
// You can construct a concrete instance of `ReplicationSetArrayInput` via:
//
//	ReplicationSetArray{ ReplicationSetArgs{...} }
type ReplicationSetArrayInput interface {
	pulumi.Input

	ToReplicationSetArrayOutput() ReplicationSetArrayOutput
	ToReplicationSetArrayOutputWithContext(context.Context) ReplicationSetArrayOutput
}

type ReplicationSetArray []ReplicationSetInput

func (ReplicationSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationSet)(nil)).Elem()
}

func (i ReplicationSetArray) ToReplicationSetArrayOutput() ReplicationSetArrayOutput {
	return i.ToReplicationSetArrayOutputWithContext(context.Background())
}

func (i ReplicationSetArray) ToReplicationSetArrayOutputWithContext(ctx context.Context) ReplicationSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSetArrayOutput)
}

// ReplicationSetMapInput is an input type that accepts ReplicationSetMap and ReplicationSetMapOutput values.
// You can construct a concrete instance of `ReplicationSetMapInput` via:
//
//	ReplicationSetMap{ "key": ReplicationSetArgs{...} }
type ReplicationSetMapInput interface {
	pulumi.Input

	ToReplicationSetMapOutput() ReplicationSetMapOutput
	ToReplicationSetMapOutputWithContext(context.Context) ReplicationSetMapOutput
}

type ReplicationSetMap map[string]ReplicationSetInput

func (ReplicationSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationSet)(nil)).Elem()
}

func (i ReplicationSetMap) ToReplicationSetMapOutput() ReplicationSetMapOutput {
	return i.ToReplicationSetMapOutputWithContext(context.Background())
}

func (i ReplicationSetMap) ToReplicationSetMapOutputWithContext(ctx context.Context) ReplicationSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSetMapOutput)
}

type ReplicationSetOutput struct{ *pulumi.OutputState }

func (ReplicationSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSet)(nil)).Elem()
}

func (o ReplicationSetOutput) ToReplicationSetOutput() ReplicationSetOutput {
	return o
}

func (o ReplicationSetOutput) ToReplicationSetOutputWithContext(ctx context.Context) ReplicationSetOutput {
	return o
}

// The ARN of the replication set.
func (o ReplicationSetOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSet) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ARN of the user who created the replication set.
func (o ReplicationSetOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSet) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// If `true`, the last region in a replication set cannot be deleted.
func (o ReplicationSetOutput) DeletionProtected() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReplicationSet) pulumi.BoolOutput { return v.DeletionProtected }).(pulumi.BoolOutput)
}

// A timestamp showing when the replication set was last modified.
func (o ReplicationSetOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSet) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o ReplicationSetOutput) Regions() ReplicationSetRegionArrayOutput {
	return o.ApplyT(func(v *ReplicationSet) ReplicationSetRegionArrayOutput { return v.Regions }).(ReplicationSetRegionArrayOutput)
}

// The current status of the Region.
// * Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
func (o ReplicationSetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSet) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags applied to the replication set.
//
// For information about the maximum allowed number of Regions and tag value constraints, see [CreateReplicationSet in the *AWS Systems Manager Incident Manager API Reference*](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateReplicationSet.html).
func (o ReplicationSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicationSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ReplicationSetOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicationSet) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type ReplicationSetArrayOutput struct{ *pulumi.OutputState }

func (ReplicationSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationSet)(nil)).Elem()
}

func (o ReplicationSetArrayOutput) ToReplicationSetArrayOutput() ReplicationSetArrayOutput {
	return o
}

func (o ReplicationSetArrayOutput) ToReplicationSetArrayOutputWithContext(ctx context.Context) ReplicationSetArrayOutput {
	return o
}

func (o ReplicationSetArrayOutput) Index(i pulumi.IntInput) ReplicationSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationSet {
		return vs[0].([]*ReplicationSet)[vs[1].(int)]
	}).(ReplicationSetOutput)
}

type ReplicationSetMapOutput struct{ *pulumi.OutputState }

func (ReplicationSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationSet)(nil)).Elem()
}

func (o ReplicationSetMapOutput) ToReplicationSetMapOutput() ReplicationSetMapOutput {
	return o
}

func (o ReplicationSetMapOutput) ToReplicationSetMapOutputWithContext(ctx context.Context) ReplicationSetMapOutput {
	return o
}

func (o ReplicationSetMapOutput) MapIndex(k pulumi.StringInput) ReplicationSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationSet {
		return vs[0].(map[string]*ReplicationSet)[vs[1].(string)]
	}).(ReplicationSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSetInput)(nil)).Elem(), &ReplicationSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSetArrayInput)(nil)).Elem(), ReplicationSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSetMapInput)(nil)).Elem(), ReplicationSetMap{})
	pulumi.RegisterOutputType(ReplicationSetOutput{})
	pulumi.RegisterOutputType(ReplicationSetArrayOutput{})
	pulumi.RegisterOutputType(ReplicationSetMapOutput{})
}
