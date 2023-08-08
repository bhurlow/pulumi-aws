// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.
//
// > **NOTE:** Changing most arguments will stop the task if it is running. You can set `startReplicationTask` to resume the task afterwards.
//
// ## Example Usage
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
//			_, err := dms.NewReplicationTask(ctx, "test", &dms.ReplicationTaskArgs{
//				CdcStartTime:            pulumi.String("1484346880"),
//				MigrationType:           pulumi.String("full-load"),
//				ReplicationInstanceArn:  pulumi.Any(aws_dms_replication_instance.TestDmsReplicationInstanceTf.Replication_instance_arn),
//				ReplicationTaskId:       pulumi.String("test-dms-replication-task-tf"),
//				ReplicationTaskSettings: pulumi.String("..."),
//				SourceEndpointArn:       pulumi.Any(aws_dms_endpoint.TestDmsSourceEndpointTf.Endpoint_arn),
//				TableMappings:           pulumi.String("{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("test"),
//				},
//				TargetEndpointArn: pulumi.Any(aws_dms_endpoint.TestDmsTargetEndpointTf.Endpoint_arn),
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
//	to = aws_dms_replication_task.test
//
//	id = "test-dms-replication-task-tf" } Using `pulumi import`, import replication tasks using the `replication_task_id`. For exampleconsole % pulumi import aws_dms_replication_task.test test-dms-replication-task-tf
type ReplicationTask struct {
	pulumi.CustomResourceState

	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
	CdcStartPosition pulumi.StringOutput `pulumi:"cdcStartPosition"`
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringPtrOutput `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringOutput `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringOutput `pulumi:"replicationInstanceArn"`
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn pulumi.StringOutput `pulumi:"replicationTaskArn"`
	// The replication task identifier.
	//
	// - Must contain from 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen.
	// - Cannot contain two consecutive hyphens.
	ReplicationTaskId pulumi.StringOutput `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringPtrOutput `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringOutput `pulumi:"sourceEndpointArn"`
	// Whether to run or stop the replication task.
	StartReplicationTask pulumi.BoolPtrOutput `pulumi:"startReplicationTask"`
	// Replication Task status.
	Status pulumi.StringOutput `pulumi:"status"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringOutput `pulumi:"tableMappings"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringOutput `pulumi:"targetEndpointArn"`
}

// NewReplicationTask registers a new resource with the given unique name, arguments, and options.
func NewReplicationTask(ctx *pulumi.Context,
	name string, args *ReplicationTaskArgs, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrationType == nil {
		return nil, errors.New("invalid value for required argument 'MigrationType'")
	}
	if args.ReplicationInstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationInstanceArn'")
	}
	if args.ReplicationTaskId == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationTaskId'")
	}
	if args.SourceEndpointArn == nil {
		return nil, errors.New("invalid value for required argument 'SourceEndpointArn'")
	}
	if args.TableMappings == nil {
		return nil, errors.New("invalid value for required argument 'TableMappings'")
	}
	if args.TargetEndpointArn == nil {
		return nil, errors.New("invalid value for required argument 'TargetEndpointArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplicationTask
	err := ctx.RegisterResource("aws:dms/replicationTask:ReplicationTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationTask gets an existing ReplicationTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationTaskState, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	var resource ReplicationTask
	err := ctx.ReadResource("aws:dms/replicationTask:ReplicationTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationTask resources.
type replicationTaskState struct {
	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
	CdcStartPosition *string `pulumi:"cdcStartPosition"`
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime *string `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType *string `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string `pulumi:"replicationInstanceArn"`
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn *string `pulumi:"replicationTaskArn"`
	// The replication task identifier.
	//
	// - Must contain from 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen.
	// - Cannot contain two consecutive hyphens.
	ReplicationTaskId *string `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings *string `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn *string `pulumi:"sourceEndpointArn"`
	// Whether to run or stop the replication task.
	StartReplicationTask *bool `pulumi:"startReplicationTask"`
	// Replication Task status.
	Status *string `pulumi:"status"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings *string `pulumi:"tableMappings"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn *string `pulumi:"targetEndpointArn"`
}

type ReplicationTaskState struct {
	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
	CdcStartPosition pulumi.StringPtrInput
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringPtrInput
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn pulumi.StringPtrInput
	// The replication task identifier.
	//
	// - Must contain from 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen.
	// - Cannot contain two consecutive hyphens.
	ReplicationTaskId pulumi.StringPtrInput
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringPtrInput
	// Whether to run or stop the replication task.
	StartReplicationTask pulumi.BoolPtrInput
	// Replication Task status.
	Status pulumi.StringPtrInput
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringPtrInput
}

func (ReplicationTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationTaskState)(nil)).Elem()
}

type replicationTaskArgs struct {
	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
	CdcStartPosition *string `pulumi:"cdcStartPosition"`
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime *string `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType string `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn string `pulumi:"replicationInstanceArn"`
	// The replication task identifier.
	//
	// - Must contain from 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen.
	// - Cannot contain two consecutive hyphens.
	ReplicationTaskId string `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings *string `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn string `pulumi:"sourceEndpointArn"`
	// Whether to run or stop the replication task.
	StartReplicationTask *bool `pulumi:"startReplicationTask"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings string `pulumi:"tableMappings"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn string `pulumi:"targetEndpointArn"`
}

// The set of arguments for constructing a ReplicationTask resource.
type ReplicationTaskArgs struct {
	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
	CdcStartPosition pulumi.StringPtrInput
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringPtrInput
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringInput
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringInput
	// The replication task identifier.
	//
	// - Must contain from 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen.
	// - Cannot contain two consecutive hyphens.
	ReplicationTaskId pulumi.StringInput
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringInput
	// Whether to run or stop the replication task.
	StartReplicationTask pulumi.BoolPtrInput
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringInput
}

func (ReplicationTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationTaskArgs)(nil)).Elem()
}

type ReplicationTaskInput interface {
	pulumi.Input

	ToReplicationTaskOutput() ReplicationTaskOutput
	ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput
}

func (*ReplicationTask) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationTask)(nil)).Elem()
}

func (i *ReplicationTask) ToReplicationTaskOutput() ReplicationTaskOutput {
	return i.ToReplicationTaskOutputWithContext(context.Background())
}

func (i *ReplicationTask) ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationTaskOutput)
}

// ReplicationTaskArrayInput is an input type that accepts ReplicationTaskArray and ReplicationTaskArrayOutput values.
// You can construct a concrete instance of `ReplicationTaskArrayInput` via:
//
//	ReplicationTaskArray{ ReplicationTaskArgs{...} }
type ReplicationTaskArrayInput interface {
	pulumi.Input

	ToReplicationTaskArrayOutput() ReplicationTaskArrayOutput
	ToReplicationTaskArrayOutputWithContext(context.Context) ReplicationTaskArrayOutput
}

type ReplicationTaskArray []ReplicationTaskInput

func (ReplicationTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationTask)(nil)).Elem()
}

func (i ReplicationTaskArray) ToReplicationTaskArrayOutput() ReplicationTaskArrayOutput {
	return i.ToReplicationTaskArrayOutputWithContext(context.Background())
}

func (i ReplicationTaskArray) ToReplicationTaskArrayOutputWithContext(ctx context.Context) ReplicationTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationTaskArrayOutput)
}

// ReplicationTaskMapInput is an input type that accepts ReplicationTaskMap and ReplicationTaskMapOutput values.
// You can construct a concrete instance of `ReplicationTaskMapInput` via:
//
//	ReplicationTaskMap{ "key": ReplicationTaskArgs{...} }
type ReplicationTaskMapInput interface {
	pulumi.Input

	ToReplicationTaskMapOutput() ReplicationTaskMapOutput
	ToReplicationTaskMapOutputWithContext(context.Context) ReplicationTaskMapOutput
}

type ReplicationTaskMap map[string]ReplicationTaskInput

func (ReplicationTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationTask)(nil)).Elem()
}

func (i ReplicationTaskMap) ToReplicationTaskMapOutput() ReplicationTaskMapOutput {
	return i.ToReplicationTaskMapOutputWithContext(context.Background())
}

func (i ReplicationTaskMap) ToReplicationTaskMapOutputWithContext(ctx context.Context) ReplicationTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationTaskMapOutput)
}

type ReplicationTaskOutput struct{ *pulumi.OutputState }

func (ReplicationTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationTask)(nil)).Elem()
}

func (o ReplicationTaskOutput) ToReplicationTaskOutput() ReplicationTaskOutput {
	return o
}

func (o ReplicationTaskOutput) ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput {
	return o
}

// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
func (o ReplicationTaskOutput) CdcStartPosition() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.CdcStartPosition }).(pulumi.StringOutput)
}

// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
func (o ReplicationTaskOutput) CdcStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringPtrOutput { return v.CdcStartTime }).(pulumi.StringPtrOutput)
}

// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
func (o ReplicationTaskOutput) MigrationType() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.MigrationType }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the replication instance.
func (o ReplicationTaskOutput) ReplicationInstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.ReplicationInstanceArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the replication task.
func (o ReplicationTaskOutput) ReplicationTaskArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.ReplicationTaskArn }).(pulumi.StringOutput)
}

// The replication task identifier.
//
// - Must contain from 1 to 255 alphanumeric characters or hyphens.
// - First character must be a letter.
// - Cannot end with a hyphen.
// - Cannot contain two consecutive hyphens.
func (o ReplicationTaskOutput) ReplicationTaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.ReplicationTaskId }).(pulumi.StringOutput)
}

// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
func (o ReplicationTaskOutput) ReplicationTaskSettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringPtrOutput { return v.ReplicationTaskSettings }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
func (o ReplicationTaskOutput) SourceEndpointArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.SourceEndpointArn }).(pulumi.StringOutput)
}

// Whether to run or stop the replication task.
func (o ReplicationTaskOutput) StartReplicationTask() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.BoolPtrOutput { return v.StartReplicationTask }).(pulumi.BoolPtrOutput)
}

// Replication Task status.
func (o ReplicationTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
func (o ReplicationTaskOutput) TableMappings() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.TableMappings }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ReplicationTaskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ReplicationTaskOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
func (o ReplicationTaskOutput) TargetEndpointArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationTask) pulumi.StringOutput { return v.TargetEndpointArn }).(pulumi.StringOutput)
}

type ReplicationTaskArrayOutput struct{ *pulumi.OutputState }

func (ReplicationTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationTask)(nil)).Elem()
}

func (o ReplicationTaskArrayOutput) ToReplicationTaskArrayOutput() ReplicationTaskArrayOutput {
	return o
}

func (o ReplicationTaskArrayOutput) ToReplicationTaskArrayOutputWithContext(ctx context.Context) ReplicationTaskArrayOutput {
	return o
}

func (o ReplicationTaskArrayOutput) Index(i pulumi.IntInput) ReplicationTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationTask {
		return vs[0].([]*ReplicationTask)[vs[1].(int)]
	}).(ReplicationTaskOutput)
}

type ReplicationTaskMapOutput struct{ *pulumi.OutputState }

func (ReplicationTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationTask)(nil)).Elem()
}

func (o ReplicationTaskMapOutput) ToReplicationTaskMapOutput() ReplicationTaskMapOutput {
	return o
}

func (o ReplicationTaskMapOutput) ToReplicationTaskMapOutputWithContext(ctx context.Context) ReplicationTaskMapOutput {
	return o
}

func (o ReplicationTaskMapOutput) MapIndex(k pulumi.StringInput) ReplicationTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationTask {
		return vs[0].(map[string]*ReplicationTask)[vs[1].(string)]
	}).(ReplicationTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationTaskInput)(nil)).Elem(), &ReplicationTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationTaskArrayInput)(nil)).Elem(), ReplicationTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationTaskMapInput)(nil)).Elem(), ReplicationTaskMap{})
	pulumi.RegisterOutputType(ReplicationTaskOutput{})
	pulumi.RegisterOutputType(ReplicationTaskArrayOutput{})
	pulumi.RegisterOutputType(ReplicationTaskMapOutput{})
}
