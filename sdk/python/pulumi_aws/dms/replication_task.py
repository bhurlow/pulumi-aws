# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReplicationTaskArgs', 'ReplicationTask']

@pulumi.input_type
class ReplicationTaskArgs:
    def __init__(__self__, *,
                 migration_type: pulumi.Input[str],
                 replication_instance_arn: pulumi.Input[str],
                 replication_task_id: pulumi.Input[str],
                 source_endpoint_arn: pulumi.Input[str],
                 table_mappings: pulumi.Input[str],
                 target_endpoint_arn: pulumi.Input[str],
                 cdc_start_position: Optional[pulumi.Input[str]] = None,
                 cdc_start_time: Optional[pulumi.Input[str]] = None,
                 replication_task_settings: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ReplicationTask resource.
        :param pulumi.Input[str] migration_type: The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        :param pulumi.Input[str] replication_instance_arn: The Amazon Resource Name (ARN) of the replication instance.
        :param pulumi.Input[str] replication_task_id: The replication task identifier.
        :param pulumi.Input[str] source_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        :param pulumi.Input[str] table_mappings: An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        :param pulumi.Input[str] target_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        :param pulumi.Input[str] cdc_start_position: Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
        :param pulumi.Input[str] cdc_start_time: The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        :param pulumi.Input[str] replication_task_settings: An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "migration_type", migration_type)
        pulumi.set(__self__, "replication_instance_arn", replication_instance_arn)
        pulumi.set(__self__, "replication_task_id", replication_task_id)
        pulumi.set(__self__, "source_endpoint_arn", source_endpoint_arn)
        pulumi.set(__self__, "table_mappings", table_mappings)
        pulumi.set(__self__, "target_endpoint_arn", target_endpoint_arn)
        if cdc_start_position is not None:
            pulumi.set(__self__, "cdc_start_position", cdc_start_position)
        if cdc_start_time is not None:
            pulumi.set(__self__, "cdc_start_time", cdc_start_time)
        if replication_task_settings is not None:
            pulumi.set(__self__, "replication_task_settings", replication_task_settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="migrationType")
    def migration_type(self) -> pulumi.Input[str]:
        """
        The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        """
        return pulumi.get(self, "migration_type")

    @migration_type.setter
    def migration_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "migration_type", value)

    @property
    @pulumi.getter(name="replicationInstanceArn")
    def replication_instance_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the replication instance.
        """
        return pulumi.get(self, "replication_instance_arn")

    @replication_instance_arn.setter
    def replication_instance_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "replication_instance_arn", value)

    @property
    @pulumi.getter(name="replicationTaskId")
    def replication_task_id(self) -> pulumi.Input[str]:
        """
        The replication task identifier.
        """
        return pulumi.get(self, "replication_task_id")

    @replication_task_id.setter
    def replication_task_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "replication_task_id", value)

    @property
    @pulumi.getter(name="sourceEndpointArn")
    def source_endpoint_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        """
        return pulumi.get(self, "source_endpoint_arn")

    @source_endpoint_arn.setter
    def source_endpoint_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_endpoint_arn", value)

    @property
    @pulumi.getter(name="tableMappings")
    def table_mappings(self) -> pulumi.Input[str]:
        """
        An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        """
        return pulumi.get(self, "table_mappings")

    @table_mappings.setter
    def table_mappings(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_mappings", value)

    @property
    @pulumi.getter(name="targetEndpointArn")
    def target_endpoint_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        """
        return pulumi.get(self, "target_endpoint_arn")

    @target_endpoint_arn.setter
    def target_endpoint_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_endpoint_arn", value)

    @property
    @pulumi.getter(name="cdcStartPosition")
    def cdc_start_position(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
        """
        return pulumi.get(self, "cdc_start_position")

    @cdc_start_position.setter
    def cdc_start_position(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cdc_start_position", value)

    @property
    @pulumi.getter(name="cdcStartTime")
    def cdc_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        """
        return pulumi.get(self, "cdc_start_time")

    @cdc_start_time.setter
    def cdc_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cdc_start_time", value)

    @property
    @pulumi.getter(name="replicationTaskSettings")
    def replication_task_settings(self) -> Optional[pulumi.Input[str]]:
        """
        An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        """
        return pulumi.get(self, "replication_task_settings")

    @replication_task_settings.setter
    def replication_task_settings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_task_settings", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ReplicationTaskState:
    def __init__(__self__, *,
                 cdc_start_position: Optional[pulumi.Input[str]] = None,
                 cdc_start_time: Optional[pulumi.Input[str]] = None,
                 migration_type: Optional[pulumi.Input[str]] = None,
                 replication_instance_arn: Optional[pulumi.Input[str]] = None,
                 replication_task_arn: Optional[pulumi.Input[str]] = None,
                 replication_task_id: Optional[pulumi.Input[str]] = None,
                 replication_task_settings: Optional[pulumi.Input[str]] = None,
                 source_endpoint_arn: Optional[pulumi.Input[str]] = None,
                 table_mappings: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_endpoint_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReplicationTask resources.
        :param pulumi.Input[str] cdc_start_position: Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
        :param pulumi.Input[str] cdc_start_time: The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        :param pulumi.Input[str] migration_type: The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        :param pulumi.Input[str] replication_instance_arn: The Amazon Resource Name (ARN) of the replication instance.
        :param pulumi.Input[str] replication_task_arn: The Amazon Resource Name (ARN) for the replication task.
        :param pulumi.Input[str] replication_task_id: The replication task identifier.
        :param pulumi.Input[str] replication_task_settings: An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        :param pulumi.Input[str] source_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        :param pulumi.Input[str] table_mappings: An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] target_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        """
        if cdc_start_position is not None:
            pulumi.set(__self__, "cdc_start_position", cdc_start_position)
        if cdc_start_time is not None:
            pulumi.set(__self__, "cdc_start_time", cdc_start_time)
        if migration_type is not None:
            pulumi.set(__self__, "migration_type", migration_type)
        if replication_instance_arn is not None:
            pulumi.set(__self__, "replication_instance_arn", replication_instance_arn)
        if replication_task_arn is not None:
            pulumi.set(__self__, "replication_task_arn", replication_task_arn)
        if replication_task_id is not None:
            pulumi.set(__self__, "replication_task_id", replication_task_id)
        if replication_task_settings is not None:
            pulumi.set(__self__, "replication_task_settings", replication_task_settings)
        if source_endpoint_arn is not None:
            pulumi.set(__self__, "source_endpoint_arn", source_endpoint_arn)
        if table_mappings is not None:
            pulumi.set(__self__, "table_mappings", table_mappings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if target_endpoint_arn is not None:
            pulumi.set(__self__, "target_endpoint_arn", target_endpoint_arn)

    @property
    @pulumi.getter(name="cdcStartPosition")
    def cdc_start_position(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
        """
        return pulumi.get(self, "cdc_start_position")

    @cdc_start_position.setter
    def cdc_start_position(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cdc_start_position", value)

    @property
    @pulumi.getter(name="cdcStartTime")
    def cdc_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        """
        return pulumi.get(self, "cdc_start_time")

    @cdc_start_time.setter
    def cdc_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cdc_start_time", value)

    @property
    @pulumi.getter(name="migrationType")
    def migration_type(self) -> Optional[pulumi.Input[str]]:
        """
        The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        """
        return pulumi.get(self, "migration_type")

    @migration_type.setter
    def migration_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "migration_type", value)

    @property
    @pulumi.getter(name="replicationInstanceArn")
    def replication_instance_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the replication instance.
        """
        return pulumi.get(self, "replication_instance_arn")

    @replication_instance_arn.setter
    def replication_instance_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_instance_arn", value)

    @property
    @pulumi.getter(name="replicationTaskArn")
    def replication_task_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) for the replication task.
        """
        return pulumi.get(self, "replication_task_arn")

    @replication_task_arn.setter
    def replication_task_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_task_arn", value)

    @property
    @pulumi.getter(name="replicationTaskId")
    def replication_task_id(self) -> Optional[pulumi.Input[str]]:
        """
        The replication task identifier.
        """
        return pulumi.get(self, "replication_task_id")

    @replication_task_id.setter
    def replication_task_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_task_id", value)

    @property
    @pulumi.getter(name="replicationTaskSettings")
    def replication_task_settings(self) -> Optional[pulumi.Input[str]]:
        """
        An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        """
        return pulumi.get(self, "replication_task_settings")

    @replication_task_settings.setter
    def replication_task_settings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_task_settings", value)

    @property
    @pulumi.getter(name="sourceEndpointArn")
    def source_endpoint_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        """
        return pulumi.get(self, "source_endpoint_arn")

    @source_endpoint_arn.setter
    def source_endpoint_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_endpoint_arn", value)

    @property
    @pulumi.getter(name="tableMappings")
    def table_mappings(self) -> Optional[pulumi.Input[str]]:
        """
        An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        """
        return pulumi.get(self, "table_mappings")

    @table_mappings.setter
    def table_mappings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_mappings", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="targetEndpointArn")
    def target_endpoint_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        """
        return pulumi.get(self, "target_endpoint_arn")

    @target_endpoint_arn.setter
    def target_endpoint_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_endpoint_arn", value)


class ReplicationTask(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cdc_start_position: Optional[pulumi.Input[str]] = None,
                 cdc_start_time: Optional[pulumi.Input[str]] = None,
                 migration_type: Optional[pulumi.Input[str]] = None,
                 replication_instance_arn: Optional[pulumi.Input[str]] = None,
                 replication_task_id: Optional[pulumi.Input[str]] = None,
                 replication_task_settings: Optional[pulumi.Input[str]] = None,
                 source_endpoint_arn: Optional[pulumi.Input[str]] = None,
                 table_mappings: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_endpoint_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new replication task
        test = aws.dms.ReplicationTask("test",
            cdc_start_time="1484346880",
            migration_type="full-load",
            replication_instance_arn=aws_dms_replication_instance["test-dms-replication-instance-tf"]["replication_instance_arn"],
            replication_task_id="test-dms-replication-task-tf",
            replication_task_settings="...",
            source_endpoint_arn=aws_dms_endpoint["test-dms-source-endpoint-tf"]["endpoint_arn"],
            table_mappings="{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}",
            tags={
                "Name": "test",
            },
            target_endpoint_arn=aws_dms_endpoint["test-dms-target-endpoint-tf"]["endpoint_arn"])
        ```

        ## Import

        Replication tasks can be imported using the `replication_task_id`, e.g.

        ```sh
         $ pulumi import aws:dms/replicationTask:ReplicationTask test test-dms-replication-task-tf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdc_start_position: Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
        :param pulumi.Input[str] cdc_start_time: The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        :param pulumi.Input[str] migration_type: The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        :param pulumi.Input[str] replication_instance_arn: The Amazon Resource Name (ARN) of the replication instance.
        :param pulumi.Input[str] replication_task_id: The replication task identifier.
        :param pulumi.Input[str] replication_task_settings: An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        :param pulumi.Input[str] source_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        :param pulumi.Input[str] table_mappings: An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] target_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicationTaskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new replication task
        test = aws.dms.ReplicationTask("test",
            cdc_start_time="1484346880",
            migration_type="full-load",
            replication_instance_arn=aws_dms_replication_instance["test-dms-replication-instance-tf"]["replication_instance_arn"],
            replication_task_id="test-dms-replication-task-tf",
            replication_task_settings="...",
            source_endpoint_arn=aws_dms_endpoint["test-dms-source-endpoint-tf"]["endpoint_arn"],
            table_mappings="{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}",
            tags={
                "Name": "test",
            },
            target_endpoint_arn=aws_dms_endpoint["test-dms-target-endpoint-tf"]["endpoint_arn"])
        ```

        ## Import

        Replication tasks can be imported using the `replication_task_id`, e.g.

        ```sh
         $ pulumi import aws:dms/replicationTask:ReplicationTask test test-dms-replication-task-tf
        ```

        :param str resource_name: The name of the resource.
        :param ReplicationTaskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicationTaskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cdc_start_position: Optional[pulumi.Input[str]] = None,
                 cdc_start_time: Optional[pulumi.Input[str]] = None,
                 migration_type: Optional[pulumi.Input[str]] = None,
                 replication_instance_arn: Optional[pulumi.Input[str]] = None,
                 replication_task_id: Optional[pulumi.Input[str]] = None,
                 replication_task_settings: Optional[pulumi.Input[str]] = None,
                 source_endpoint_arn: Optional[pulumi.Input[str]] = None,
                 table_mappings: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_endpoint_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicationTaskArgs.__new__(ReplicationTaskArgs)

            __props__.__dict__["cdc_start_position"] = cdc_start_position
            __props__.__dict__["cdc_start_time"] = cdc_start_time
            if migration_type is None and not opts.urn:
                raise TypeError("Missing required property 'migration_type'")
            __props__.__dict__["migration_type"] = migration_type
            if replication_instance_arn is None and not opts.urn:
                raise TypeError("Missing required property 'replication_instance_arn'")
            __props__.__dict__["replication_instance_arn"] = replication_instance_arn
            if replication_task_id is None and not opts.urn:
                raise TypeError("Missing required property 'replication_task_id'")
            __props__.__dict__["replication_task_id"] = replication_task_id
            __props__.__dict__["replication_task_settings"] = replication_task_settings
            if source_endpoint_arn is None and not opts.urn:
                raise TypeError("Missing required property 'source_endpoint_arn'")
            __props__.__dict__["source_endpoint_arn"] = source_endpoint_arn
            if table_mappings is None and not opts.urn:
                raise TypeError("Missing required property 'table_mappings'")
            __props__.__dict__["table_mappings"] = table_mappings
            __props__.__dict__["tags"] = tags
            if target_endpoint_arn is None and not opts.urn:
                raise TypeError("Missing required property 'target_endpoint_arn'")
            __props__.__dict__["target_endpoint_arn"] = target_endpoint_arn
            __props__.__dict__["replication_task_arn"] = None
            __props__.__dict__["tags_all"] = None
        super(ReplicationTask, __self__).__init__(
            'aws:dms/replicationTask:ReplicationTask',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cdc_start_position: Optional[pulumi.Input[str]] = None,
            cdc_start_time: Optional[pulumi.Input[str]] = None,
            migration_type: Optional[pulumi.Input[str]] = None,
            replication_instance_arn: Optional[pulumi.Input[str]] = None,
            replication_task_arn: Optional[pulumi.Input[str]] = None,
            replication_task_id: Optional[pulumi.Input[str]] = None,
            replication_task_settings: Optional[pulumi.Input[str]] = None,
            source_endpoint_arn: Optional[pulumi.Input[str]] = None,
            table_mappings: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_endpoint_arn: Optional[pulumi.Input[str]] = None) -> 'ReplicationTask':
        """
        Get an existing ReplicationTask resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdc_start_position: Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
        :param pulumi.Input[str] cdc_start_time: The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        :param pulumi.Input[str] migration_type: The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        :param pulumi.Input[str] replication_instance_arn: The Amazon Resource Name (ARN) of the replication instance.
        :param pulumi.Input[str] replication_task_arn: The Amazon Resource Name (ARN) for the replication task.
        :param pulumi.Input[str] replication_task_id: The replication task identifier.
        :param pulumi.Input[str] replication_task_settings: An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        :param pulumi.Input[str] source_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        :param pulumi.Input[str] table_mappings: An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] target_endpoint_arn: The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReplicationTaskState.__new__(_ReplicationTaskState)

        __props__.__dict__["cdc_start_position"] = cdc_start_position
        __props__.__dict__["cdc_start_time"] = cdc_start_time
        __props__.__dict__["migration_type"] = migration_type
        __props__.__dict__["replication_instance_arn"] = replication_instance_arn
        __props__.__dict__["replication_task_arn"] = replication_task_arn
        __props__.__dict__["replication_task_id"] = replication_task_id
        __props__.__dict__["replication_task_settings"] = replication_task_settings
        __props__.__dict__["source_endpoint_arn"] = source_endpoint_arn
        __props__.__dict__["table_mappings"] = table_mappings
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["target_endpoint_arn"] = target_endpoint_arn
        return ReplicationTask(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cdcStartPosition")
    def cdc_start_position(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
        """
        return pulumi.get(self, "cdc_start_position")

    @property
    @pulumi.getter(name="cdcStartTime")
    def cdc_start_time(self) -> pulumi.Output[Optional[str]]:
        """
        The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
        """
        return pulumi.get(self, "cdc_start_time")

    @property
    @pulumi.getter(name="migrationType")
    def migration_type(self) -> pulumi.Output[str]:
        """
        The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
        """
        return pulumi.get(self, "migration_type")

    @property
    @pulumi.getter(name="replicationInstanceArn")
    def replication_instance_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the replication instance.
        """
        return pulumi.get(self, "replication_instance_arn")

    @property
    @pulumi.getter(name="replicationTaskArn")
    def replication_task_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) for the replication task.
        """
        return pulumi.get(self, "replication_task_arn")

    @property
    @pulumi.getter(name="replicationTaskId")
    def replication_task_id(self) -> pulumi.Output[str]:
        """
        The replication task identifier.
        """
        return pulumi.get(self, "replication_task_id")

    @property
    @pulumi.getter(name="replicationTaskSettings")
    def replication_task_settings(self) -> pulumi.Output[Optional[str]]:
        """
        An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
        """
        return pulumi.get(self, "replication_task_settings")

    @property
    @pulumi.getter(name="sourceEndpointArn")
    def source_endpoint_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
        """
        return pulumi.get(self, "source_endpoint_arn")

    @property
    @pulumi.getter(name="tableMappings")
    def table_mappings(self) -> pulumi.Output[str]:
        """
        An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
        """
        return pulumi.get(self, "table_mappings")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="targetEndpointArn")
    def target_endpoint_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
        """
        return pulumi.get(self, "target_endpoint_arn")

