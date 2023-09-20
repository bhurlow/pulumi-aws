// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.
//
// > **Note:** AWS requires a special IAM role called `dms-vpc-role` when using this resource. See the example below to create it as part of your configuration.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dms.NewReplicationSubnetGroup(ctx, "example", &dms.ReplicationSubnetGroupArgs{
//				ReplicationSubnetGroupDescription: pulumi.String("Example replication subnet group"),
//				ReplicationSubnetGroupId:          pulumi.String("example-dms-replication-subnet-group-tf"),
//				SubnetIds: pulumi.StringArray{
//					pulumi.String("subnet-12345678"),
//					pulumi.String("subnet-12345679"),
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("example"),
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
// ### Creating special IAM role
//
// If your account does not already include the `dms-vpc-role` IAM role, you will need to create it to allow DMS to manage subnets in the VPC.
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Version": "2012-10-17",
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Effect": "Allow",
//						"Principal": map[string]interface{}{
//							"Service": "dms.amazonaws.com",
//						},
//						"Action": "sts:AssumeRole",
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = iam.NewRole(ctx, "dms-vpc-role", &iam.RoleArgs{
//				Description:      pulumi.String("Allows DMS to manage VPC"),
//				AssumeRolePolicy: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			exampleRolePolicyAttachment, err := iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
//				Role:      dms_vpc_role.Name,
//				PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dms.NewReplicationSubnetGroup(ctx, "exampleReplicationSubnetGroup", &dms.ReplicationSubnetGroupArgs{
//				ReplicationSubnetGroupDescription: pulumi.String("Example"),
//				ReplicationSubnetGroupId:          pulumi.String("example-id"),
//				SubnetIds: pulumi.StringArray{
//					pulumi.String("subnet-12345678"),
//					pulumi.String("subnet-12345679"),
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("example-id"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleRolePolicyAttachment,
//			}))
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
// Using `pulumi import`, import replication subnet groups using the `replication_subnet_group_id`. For example:
//
// ```sh
//
//	$ pulumi import aws:dms/replicationSubnetGroup:ReplicationSubnetGroup test test-dms-replication-subnet-group-tf
//
// ```
type ReplicationSubnetGroup struct {
	pulumi.CustomResourceState

	ReplicationSubnetGroupArn pulumi.StringOutput `pulumi:"replicationSubnetGroupArn"`
	// Description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringOutput `pulumi:"replicationSubnetGroupDescription"`
	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId pulumi.StringOutput `pulumi:"replicationSubnetGroupId"`
	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ID of the VPC the subnet group is in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewReplicationSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewReplicationSubnetGroup(ctx *pulumi.Context,
	name string, args *ReplicationSubnetGroupArgs, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReplicationSubnetGroupDescription == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationSubnetGroupDescription'")
	}
	if args.ReplicationSubnetGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationSubnetGroupId'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplicationSubnetGroup
	err := ctx.RegisterResource("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationSubnetGroup gets an existing ReplicationSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationSubnetGroupState, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	var resource ReplicationSubnetGroup
	err := ctx.ReadResource("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationSubnetGroup resources.
type replicationSubnetGroupState struct {
	ReplicationSubnetGroupArn *string `pulumi:"replicationSubnetGroupArn"`
	// Description for the subnet group.
	ReplicationSubnetGroupDescription *string `pulumi:"replicationSubnetGroupDescription"`
	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId *string `pulumi:"replicationSubnetGroupId"`
	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds []string `pulumi:"subnetIds"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ID of the VPC the subnet group is in.
	VpcId *string `pulumi:"vpcId"`
}

type ReplicationSubnetGroupState struct {
	ReplicationSubnetGroupArn pulumi.StringPtrInput
	// Description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringPtrInput
	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId pulumi.StringPtrInput
	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds pulumi.StringArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The ID of the VPC the subnet group is in.
	VpcId pulumi.StringPtrInput
}

func (ReplicationSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupState)(nil)).Elem()
}

type replicationSubnetGroupArgs struct {
	// Description for the subnet group.
	ReplicationSubnetGroupDescription string `pulumi:"replicationSubnetGroupDescription"`
	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId string `pulumi:"replicationSubnetGroupId"`
	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds []string `pulumi:"subnetIds"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ReplicationSubnetGroup resource.
type ReplicationSubnetGroupArgs struct {
	// Description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringInput
	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId pulumi.StringInput
	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds pulumi.StringArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ReplicationSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupArgs)(nil)).Elem()
}

type ReplicationSubnetGroupInput interface {
	pulumi.Input

	ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput
	ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput
}

func (*ReplicationSubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSubnetGroup)(nil)).Elem()
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput {
	return i.ToReplicationSubnetGroupOutputWithContext(context.Background())
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupOutput)
}

func (i *ReplicationSubnetGroup) ToOutput(ctx context.Context) pulumix.Output[*ReplicationSubnetGroup] {
	return pulumix.Output[*ReplicationSubnetGroup]{
		OutputState: i.ToReplicationSubnetGroupOutputWithContext(ctx).OutputState,
	}
}

// ReplicationSubnetGroupArrayInput is an input type that accepts ReplicationSubnetGroupArray and ReplicationSubnetGroupArrayOutput values.
// You can construct a concrete instance of `ReplicationSubnetGroupArrayInput` via:
//
//	ReplicationSubnetGroupArray{ ReplicationSubnetGroupArgs{...} }
type ReplicationSubnetGroupArrayInput interface {
	pulumi.Input

	ToReplicationSubnetGroupArrayOutput() ReplicationSubnetGroupArrayOutput
	ToReplicationSubnetGroupArrayOutputWithContext(context.Context) ReplicationSubnetGroupArrayOutput
}

type ReplicationSubnetGroupArray []ReplicationSubnetGroupInput

func (ReplicationSubnetGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationSubnetGroup)(nil)).Elem()
}

func (i ReplicationSubnetGroupArray) ToReplicationSubnetGroupArrayOutput() ReplicationSubnetGroupArrayOutput {
	return i.ToReplicationSubnetGroupArrayOutputWithContext(context.Background())
}

func (i ReplicationSubnetGroupArray) ToReplicationSubnetGroupArrayOutputWithContext(ctx context.Context) ReplicationSubnetGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupArrayOutput)
}

func (i ReplicationSubnetGroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*ReplicationSubnetGroup] {
	return pulumix.Output[[]*ReplicationSubnetGroup]{
		OutputState: i.ToReplicationSubnetGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// ReplicationSubnetGroupMapInput is an input type that accepts ReplicationSubnetGroupMap and ReplicationSubnetGroupMapOutput values.
// You can construct a concrete instance of `ReplicationSubnetGroupMapInput` via:
//
//	ReplicationSubnetGroupMap{ "key": ReplicationSubnetGroupArgs{...} }
type ReplicationSubnetGroupMapInput interface {
	pulumi.Input

	ToReplicationSubnetGroupMapOutput() ReplicationSubnetGroupMapOutput
	ToReplicationSubnetGroupMapOutputWithContext(context.Context) ReplicationSubnetGroupMapOutput
}

type ReplicationSubnetGroupMap map[string]ReplicationSubnetGroupInput

func (ReplicationSubnetGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationSubnetGroup)(nil)).Elem()
}

func (i ReplicationSubnetGroupMap) ToReplicationSubnetGroupMapOutput() ReplicationSubnetGroupMapOutput {
	return i.ToReplicationSubnetGroupMapOutputWithContext(context.Background())
}

func (i ReplicationSubnetGroupMap) ToReplicationSubnetGroupMapOutputWithContext(ctx context.Context) ReplicationSubnetGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupMapOutput)
}

func (i ReplicationSubnetGroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ReplicationSubnetGroup] {
	return pulumix.Output[map[string]*ReplicationSubnetGroup]{
		OutputState: i.ToReplicationSubnetGroupMapOutputWithContext(ctx).OutputState,
	}
}

type ReplicationSubnetGroupOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSubnetGroup)(nil)).Elem()
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput {
	return o
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput {
	return o
}

func (o ReplicationSubnetGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*ReplicationSubnetGroup] {
	return pulumix.Output[*ReplicationSubnetGroup]{
		OutputState: o.OutputState,
	}
}

func (o ReplicationSubnetGroupOutput) ReplicationSubnetGroupArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) pulumi.StringOutput { return v.ReplicationSubnetGroupArn }).(pulumi.StringOutput)
}

// Description for the subnet group.
func (o ReplicationSubnetGroupOutput) ReplicationSubnetGroupDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) pulumi.StringOutput { return v.ReplicationSubnetGroupDescription }).(pulumi.StringOutput)
}

// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
func (o ReplicationSubnetGroupOutput) ReplicationSubnetGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) pulumi.StringOutput { return v.ReplicationSubnetGroupId }).(pulumi.StringOutput)
}

// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
func (o ReplicationSubnetGroupOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ReplicationSubnetGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ReplicationSubnetGroupOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The ID of the VPC the subnet group is in.
func (o ReplicationSubnetGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ReplicationSubnetGroupArrayOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationSubnetGroup)(nil)).Elem()
}

func (o ReplicationSubnetGroupArrayOutput) ToReplicationSubnetGroupArrayOutput() ReplicationSubnetGroupArrayOutput {
	return o
}

func (o ReplicationSubnetGroupArrayOutput) ToReplicationSubnetGroupArrayOutputWithContext(ctx context.Context) ReplicationSubnetGroupArrayOutput {
	return o
}

func (o ReplicationSubnetGroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ReplicationSubnetGroup] {
	return pulumix.Output[[]*ReplicationSubnetGroup]{
		OutputState: o.OutputState,
	}
}

func (o ReplicationSubnetGroupArrayOutput) Index(i pulumi.IntInput) ReplicationSubnetGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationSubnetGroup {
		return vs[0].([]*ReplicationSubnetGroup)[vs[1].(int)]
	}).(ReplicationSubnetGroupOutput)
}

type ReplicationSubnetGroupMapOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationSubnetGroup)(nil)).Elem()
}

func (o ReplicationSubnetGroupMapOutput) ToReplicationSubnetGroupMapOutput() ReplicationSubnetGroupMapOutput {
	return o
}

func (o ReplicationSubnetGroupMapOutput) ToReplicationSubnetGroupMapOutputWithContext(ctx context.Context) ReplicationSubnetGroupMapOutput {
	return o
}

func (o ReplicationSubnetGroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ReplicationSubnetGroup] {
	return pulumix.Output[map[string]*ReplicationSubnetGroup]{
		OutputState: o.OutputState,
	}
}

func (o ReplicationSubnetGroupMapOutput) MapIndex(k pulumi.StringInput) ReplicationSubnetGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationSubnetGroup {
		return vs[0].(map[string]*ReplicationSubnetGroup)[vs[1].(string)]
	}).(ReplicationSubnetGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSubnetGroupInput)(nil)).Elem(), &ReplicationSubnetGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSubnetGroupArrayInput)(nil)).Elem(), ReplicationSubnetGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSubnetGroupMapInput)(nil)).Elem(), ReplicationSubnetGroupMap{})
	pulumi.RegisterOutputType(ReplicationSubnetGroupOutput{})
	pulumi.RegisterOutputType(ReplicationSubnetGroupArrayOutput{})
	pulumi.RegisterOutputType(ReplicationSubnetGroupMapOutput{})
}
