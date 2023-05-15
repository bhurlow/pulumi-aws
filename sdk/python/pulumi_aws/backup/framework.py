# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FrameworkArgs', 'Framework']

@pulumi.input_type
class FrameworkArgs:
    def __init__(__self__, *,
                 controls: pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Framework resource.
        :param pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]] controls: One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
        :param pulumi.Input[str] description: The description of the framework with a maximum of 1,024 characters
        :param pulumi.Input[str] name: The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "controls", controls)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def controls(self) -> pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]]:
        """
        One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
        """
        return pulumi.get(self, "controls")

    @controls.setter
    def controls(self, value: pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]]):
        pulumi.set(self, "controls", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the framework with a maximum of 1,024 characters
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _FrameworkState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 controls: Optional[pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 deployment_status: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Framework resources.
        :param pulumi.Input[str] arn: The ARN of the backup framework.
        :param pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]] controls: One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
        :param pulumi.Input[str] creation_time: The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC).
        :param pulumi.Input[str] deployment_status: The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS` | `UPDATE_IN_PROGRESS` | `DELETE_IN_PROGRESS` | `COMPLETED` | `FAILED`.
        :param pulumi.Input[str] description: The description of the framework with a maximum of 1,024 characters
        :param pulumi.Input[str] name: The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
        :param pulumi.Input[str] status: A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. For more information refer to the [AWS documentation for Framework Status](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeFramework.html#Backup-DescribeFramework-response-FrameworkStatus)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if controls is not None:
            pulumi.set(__self__, "controls", controls)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if deployment_status is not None:
            pulumi.set(__self__, "deployment_status", deployment_status)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the backup framework.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]]]:
        """
        One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
        """
        return pulumi.get(self, "controls")

    @controls.setter
    def controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]]]):
        pulumi.set(self, "controls", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC).
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter(name="deploymentStatus")
    def deployment_status(self) -> Optional[pulumi.Input[str]]:
        """
        The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS` | `UPDATE_IN_PROGRESS` | `DELETE_IN_PROGRESS` | `COMPLETED` | `FAILED`.
        """
        return pulumi.get(self, "deployment_status")

    @deployment_status.setter
    def deployment_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_status", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the framework with a maximum of 1,024 characters
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. For more information refer to the [AWS documentation for Framework Status](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeFramework.html#Backup-DescribeFramework-response-FrameworkStatus)
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Framework(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an AWS Backup Framework resource.

        > **Note:** For the Deployment Status of the Framework to be successful, please turn on resource tracking to enable AWS Config recording to track configuration changes of your backup resources. This can be done from the AWS Console.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.backup.Framework("example",
            controls=[
                aws.backup.FrameworkControlArgs(
                    input_parameters=[aws.backup.FrameworkControlInputParameterArgs(
                        name="requiredRetentionDays",
                        value="35",
                    )],
                    name="BACKUP_RECOVERY_POINT_MINIMUM_RETENTION_CHECK",
                ),
                aws.backup.FrameworkControlArgs(
                    input_parameters=[
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="requiredFrequencyUnit",
                            value="hours",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="requiredRetentionDays",
                            value="35",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="requiredFrequencyValue",
                            value="1",
                        ),
                    ],
                    name="BACKUP_PLAN_MIN_FREQUENCY_AND_MIN_RETENTION_CHECK",
                ),
                aws.backup.FrameworkControlArgs(
                    name="BACKUP_RECOVERY_POINT_ENCRYPTED",
                ),
                aws.backup.FrameworkControlArgs(
                    name="BACKUP_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
                    scope=aws.backup.FrameworkControlScopeArgs(
                        compliance_resource_types=["EBS"],
                    ),
                ),
                aws.backup.FrameworkControlArgs(
                    name="BACKUP_RECOVERY_POINT_MANUAL_DELETION_DISABLED",
                ),
                aws.backup.FrameworkControlArgs(
                    input_parameters=[
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="maxRetentionDays",
                            value="100",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="minRetentionDays",
                            value="1",
                        ),
                    ],
                    name="BACKUP_RESOURCES_PROTECTED_BY_BACKUP_VAULT_LOCK",
                    scope=aws.backup.FrameworkControlScopeArgs(
                        compliance_resource_types=["EBS"],
                    ),
                ),
                aws.backup.FrameworkControlArgs(
                    input_parameters=[
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="recoveryPointAgeUnit",
                            value="days",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="recoveryPointAgeValue",
                            value="1",
                        ),
                    ],
                    name="BACKUP_LAST_RECOVERY_POINT_CREATED",
                    scope=aws.backup.FrameworkControlScopeArgs(
                        compliance_resource_types=["EBS"],
                    ),
                ),
            ],
            description="this is an example framework",
            tags={
                "Name": "Example Framework",
            })
        ```

        ## Import

        Backup Framework can be imported using the `id` which corresponds to the name of the Backup Framework, e.g.,

        ```sh
         $ pulumi import aws:backup/framework:Framework test <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]] controls: One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
        :param pulumi.Input[str] description: The description of the framework with a maximum of 1,024 characters
        :param pulumi.Input[str] name: The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FrameworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Backup Framework resource.

        > **Note:** For the Deployment Status of the Framework to be successful, please turn on resource tracking to enable AWS Config recording to track configuration changes of your backup resources. This can be done from the AWS Console.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.backup.Framework("example",
            controls=[
                aws.backup.FrameworkControlArgs(
                    input_parameters=[aws.backup.FrameworkControlInputParameterArgs(
                        name="requiredRetentionDays",
                        value="35",
                    )],
                    name="BACKUP_RECOVERY_POINT_MINIMUM_RETENTION_CHECK",
                ),
                aws.backup.FrameworkControlArgs(
                    input_parameters=[
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="requiredFrequencyUnit",
                            value="hours",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="requiredRetentionDays",
                            value="35",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="requiredFrequencyValue",
                            value="1",
                        ),
                    ],
                    name="BACKUP_PLAN_MIN_FREQUENCY_AND_MIN_RETENTION_CHECK",
                ),
                aws.backup.FrameworkControlArgs(
                    name="BACKUP_RECOVERY_POINT_ENCRYPTED",
                ),
                aws.backup.FrameworkControlArgs(
                    name="BACKUP_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
                    scope=aws.backup.FrameworkControlScopeArgs(
                        compliance_resource_types=["EBS"],
                    ),
                ),
                aws.backup.FrameworkControlArgs(
                    name="BACKUP_RECOVERY_POINT_MANUAL_DELETION_DISABLED",
                ),
                aws.backup.FrameworkControlArgs(
                    input_parameters=[
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="maxRetentionDays",
                            value="100",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="minRetentionDays",
                            value="1",
                        ),
                    ],
                    name="BACKUP_RESOURCES_PROTECTED_BY_BACKUP_VAULT_LOCK",
                    scope=aws.backup.FrameworkControlScopeArgs(
                        compliance_resource_types=["EBS"],
                    ),
                ),
                aws.backup.FrameworkControlArgs(
                    input_parameters=[
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="recoveryPointAgeUnit",
                            value="days",
                        ),
                        aws.backup.FrameworkControlInputParameterArgs(
                            name="recoveryPointAgeValue",
                            value="1",
                        ),
                    ],
                    name="BACKUP_LAST_RECOVERY_POINT_CREATED",
                    scope=aws.backup.FrameworkControlScopeArgs(
                        compliance_resource_types=["EBS"],
                    ),
                ),
            ],
            description="this is an example framework",
            tags={
                "Name": "Example Framework",
            })
        ```

        ## Import

        Backup Framework can be imported using the `id` which corresponds to the name of the Backup Framework, e.g.,

        ```sh
         $ pulumi import aws:backup/framework:Framework test <id>
        ```

        :param str resource_name: The name of the resource.
        :param FrameworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FrameworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FrameworkArgs.__new__(FrameworkArgs)

            if controls is None and not opts.urn:
                raise TypeError("Missing required property 'controls'")
            __props__.__dict__["controls"] = controls
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["deployment_status"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["tags_all"] = None
        super(Framework, __self__).__init__(
            'aws:backup/framework:Framework',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            deployment_status: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Framework':
        """
        Get an existing Framework resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the backup framework.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]] controls: One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
        :param pulumi.Input[str] creation_time: The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC).
        :param pulumi.Input[str] deployment_status: The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS` | `UPDATE_IN_PROGRESS` | `DELETE_IN_PROGRESS` | `COMPLETED` | `FAILED`.
        :param pulumi.Input[str] description: The description of the framework with a maximum of 1,024 characters
        :param pulumi.Input[str] name: The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
        :param pulumi.Input[str] status: A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. For more information refer to the [AWS documentation for Framework Status](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeFramework.html#Backup-DescribeFramework-response-FrameworkStatus)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FrameworkState.__new__(_FrameworkState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["controls"] = controls
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["deployment_status"] = deployment_status
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Framework(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the backup framework.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def controls(self) -> pulumi.Output[Sequence['outputs.FrameworkControl']]:
        """
        One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
        """
        return pulumi.get(self, "controls")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC).
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="deploymentStatus")
    def deployment_status(self) -> pulumi.Output[str]:
        """
        The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS` | `UPDATE_IN_PROGRESS` | `DELETE_IN_PROGRESS` | `COMPLETED` | `FAILED`.
        """
        return pulumi.get(self, "deployment_status")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the framework with a maximum of 1,024 characters
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. For more information refer to the [AWS documentation for Framework Status](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeFramework.html#Backup-DescribeFramework-response-FrameworkStatus)
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Metadata that you can assign to help organize the frameworks you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

