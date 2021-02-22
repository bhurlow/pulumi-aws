# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'DistributionConfigurationDistributionArgs',
    'DistributionConfigurationDistributionAmiDistributionConfigurationArgs',
    'DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs',
    'ImageImageTestsConfigurationArgs',
    'ImageOutputResourceArgs',
    'ImageOutputResourceAmiArgs',
    'ImagePipelineImageTestsConfigurationArgs',
    'ImagePipelineScheduleArgs',
    'ImageRecipeBlockDeviceMappingArgs',
    'ImageRecipeBlockDeviceMappingEbsArgs',
    'ImageRecipeComponentArgs',
    'InfrastructureConfigurationLoggingArgs',
    'InfrastructureConfigurationLoggingS3LogsArgs',
]

@pulumi.input_type
class DistributionConfigurationDistributionArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 ami_distribution_configuration: Optional[pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationArgs']] = None,
                 license_configuration_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] region: AWS Region for the distribution.
        :param pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationArgs'] ami_distribution_configuration: Configuration block with Amazon Machine Image (AMI) distribution settings. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] license_configuration_arns: Set of Amazon Resource Names (ARNs) of License Manager License Configurations.
        """
        pulumi.set(__self__, "region", region)
        if ami_distribution_configuration is not None:
            pulumi.set(__self__, "ami_distribution_configuration", ami_distribution_configuration)
        if license_configuration_arns is not None:
            pulumi.set(__self__, "license_configuration_arns", license_configuration_arns)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        AWS Region for the distribution.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="amiDistributionConfiguration")
    def ami_distribution_configuration(self) -> Optional[pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationArgs']]:
        """
        Configuration block with Amazon Machine Image (AMI) distribution settings. Detailed below.
        """
        return pulumi.get(self, "ami_distribution_configuration")

    @ami_distribution_configuration.setter
    def ami_distribution_configuration(self, value: Optional[pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationArgs']]):
        pulumi.set(self, "ami_distribution_configuration", value)

    @property
    @pulumi.getter(name="licenseConfigurationArns")
    def license_configuration_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of Amazon Resource Names (ARNs) of License Manager License Configurations.
        """
        return pulumi.get(self, "license_configuration_arns")

    @license_configuration_arns.setter
    def license_configuration_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "license_configuration_arns", value)


@pulumi.input_type
class DistributionConfigurationDistributionAmiDistributionConfigurationArgs:
    def __init__(__self__, *,
                 ami_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 launch_permission: Optional[pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] ami_tags: Key-value map of tags to apply to the distributed AMI.
        :param pulumi.Input[str] description: Description to apply to the distributed AMI.
        :param pulumi.Input[str] kms_key_id: Amazon Resource Name (ARN) of the Key Management Service (KMS) Key to encrypt the distributed AMI.
        :param pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs'] launch_permission: Configuration block of EC2 launch permissions to apply to the distributed AMI. Detailed below.
        :param pulumi.Input[str] name: Name to apply to the distributed AMI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_account_ids: Set of AWS Account identifiers to distribute the AMI.
        """
        if ami_tags is not None:
            pulumi.set(__self__, "ami_tags", ami_tags)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if launch_permission is not None:
            pulumi.set(__self__, "launch_permission", launch_permission)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if target_account_ids is not None:
            pulumi.set(__self__, "target_account_ids", target_account_ids)

    @property
    @pulumi.getter(name="amiTags")
    def ami_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of tags to apply to the distributed AMI.
        """
        return pulumi.get(self, "ami_tags")

    @ami_tags.setter
    def ami_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "ami_tags", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description to apply to the distributed AMI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Key Management Service (KMS) Key to encrypt the distributed AMI.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="launchPermission")
    def launch_permission(self) -> Optional[pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs']]:
        """
        Configuration block of EC2 launch permissions to apply to the distributed AMI. Detailed below.
        """
        return pulumi.get(self, "launch_permission")

    @launch_permission.setter
    def launch_permission(self, value: Optional[pulumi.Input['DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs']]):
        pulumi.set(self, "launch_permission", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name to apply to the distributed AMI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="targetAccountIds")
    def target_account_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of AWS Account identifiers to distribute the AMI.
        """
        return pulumi.get(self, "target_account_ids")

    @target_account_ids.setter
    def target_account_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_account_ids", value)


@pulumi.input_type
class DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs:
    def __init__(__self__, *,
                 user_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_groups: Set of EC2 launch permission user groups to assign. Use `all` to distribute a public AMI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: Set of AWS Account identifiers to assign.
        """
        if user_groups is not None:
            pulumi.set(__self__, "user_groups", user_groups)
        if user_ids is not None:
            pulumi.set(__self__, "user_ids", user_ids)

    @property
    @pulumi.getter(name="userGroups")
    def user_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of EC2 launch permission user groups to assign. Use `all` to distribute a public AMI.
        """
        return pulumi.get(self, "user_groups")

    @user_groups.setter
    def user_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_groups", value)

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of AWS Account identifiers to assign.
        """
        return pulumi.get(self, "user_ids")

    @user_ids.setter
    def user_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_ids", value)


@pulumi.input_type
class ImageImageTestsConfigurationArgs:
    def __init__(__self__, *,
                 image_tests_enabled: Optional[pulumi.Input[bool]] = None,
                 timeout_minutes: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] image_tests_enabled: Whether image tests are enabled. Defaults to `true`.
        :param pulumi.Input[int] timeout_minutes: Number of minutes before image tests time out. Valid values are between `60` and `1440`. Defaults to `720`.
        """
        if image_tests_enabled is not None:
            pulumi.set(__self__, "image_tests_enabled", image_tests_enabled)
        if timeout_minutes is not None:
            pulumi.set(__self__, "timeout_minutes", timeout_minutes)

    @property
    @pulumi.getter(name="imageTestsEnabled")
    def image_tests_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether image tests are enabled. Defaults to `true`.
        """
        return pulumi.get(self, "image_tests_enabled")

    @image_tests_enabled.setter
    def image_tests_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "image_tests_enabled", value)

    @property
    @pulumi.getter(name="timeoutMinutes")
    def timeout_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of minutes before image tests time out. Valid values are between `60` and `1440`. Defaults to `720`.
        """
        return pulumi.get(self, "timeout_minutes")

    @timeout_minutes.setter
    def timeout_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_minutes", value)


@pulumi.input_type
class ImageOutputResourceArgs:
    def __init__(__self__, *,
                 amis: Optional[pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceAmiArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceAmiArgs']]] amis: Set of objects with each Amazon Machine Image (AMI) created.
        """
        if amis is not None:
            pulumi.set(__self__, "amis", amis)

    @property
    @pulumi.getter
    def amis(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceAmiArgs']]]]:
        """
        Set of objects with each Amazon Machine Image (AMI) created.
        """
        return pulumi.get(self, "amis")

    @amis.setter
    def amis(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceAmiArgs']]]]):
        pulumi.set(self, "amis", value)


@pulumi.input_type
class ImageOutputResourceAmiArgs:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] account_id: Account identifier of the AMI.
        :param pulumi.Input[str] description: Description of the AMI.
        :param pulumi.Input[str] image: Identifier of the AMI.
        :param pulumi.Input[str] name: Name of the AMI.
        :param pulumi.Input[str] region: Region of the AMI.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if image is not None:
            pulumi.set(__self__, "image", image)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Account identifier of the AMI.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the AMI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def image(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the AMI.
        """
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the AMI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of the AMI.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class ImagePipelineImageTestsConfigurationArgs:
    def __init__(__self__, *,
                 image_tests_enabled: Optional[pulumi.Input[bool]] = None,
                 timeout_minutes: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] image_tests_enabled: Whether image tests are enabled. Defaults to `true`.
        :param pulumi.Input[int] timeout_minutes: Number of minutes before image tests time out. Valid values are between `60` and `1440`. Defaults to `720`.
        """
        if image_tests_enabled is not None:
            pulumi.set(__self__, "image_tests_enabled", image_tests_enabled)
        if timeout_minutes is not None:
            pulumi.set(__self__, "timeout_minutes", timeout_minutes)

    @property
    @pulumi.getter(name="imageTestsEnabled")
    def image_tests_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether image tests are enabled. Defaults to `true`.
        """
        return pulumi.get(self, "image_tests_enabled")

    @image_tests_enabled.setter
    def image_tests_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "image_tests_enabled", value)

    @property
    @pulumi.getter(name="timeoutMinutes")
    def timeout_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of minutes before image tests time out. Valid values are between `60` and `1440`. Defaults to `720`.
        """
        return pulumi.get(self, "timeout_minutes")

    @timeout_minutes.setter
    def timeout_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_minutes", value)


@pulumi.input_type
class ImagePipelineScheduleArgs:
    def __init__(__self__, *,
                 schedule_expression: pulumi.Input[str],
                 pipeline_execution_start_condition: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] schedule_expression: Cron expression of how often the pipeline start condition is evaluated. For example, `cron(0 0 * * ? *)` is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as `cron(0 0 * * *)`, must be updated to the six field syntax. For more information, see the [Image Builder User Guide](https://docs.aws.amazon.com/imagebuilder/latest/userguide/cron-expressions.html).
        :param pulumi.Input[str] pipeline_execution_start_condition: Condition when the pipeline should trigger a new image build. Valid values are `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE` and `EXPRESSION_MATCH_ONLY`. Defaults to `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE`.
        """
        pulumi.set(__self__, "schedule_expression", schedule_expression)
        if pipeline_execution_start_condition is not None:
            pulumi.set(__self__, "pipeline_execution_start_condition", pipeline_execution_start_condition)

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> pulumi.Input[str]:
        """
        Cron expression of how often the pipeline start condition is evaluated. For example, `cron(0 0 * * ? *)` is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as `cron(0 0 * * *)`, must be updated to the six field syntax. For more information, see the [Image Builder User Guide](https://docs.aws.amazon.com/imagebuilder/latest/userguide/cron-expressions.html).
        """
        return pulumi.get(self, "schedule_expression")

    @schedule_expression.setter
    def schedule_expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule_expression", value)

    @property
    @pulumi.getter(name="pipelineExecutionStartCondition")
    def pipeline_execution_start_condition(self) -> Optional[pulumi.Input[str]]:
        """
        Condition when the pipeline should trigger a new image build. Valid values are `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE` and `EXPRESSION_MATCH_ONLY`. Defaults to `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE`.
        """
        return pulumi.get(self, "pipeline_execution_start_condition")

    @pipeline_execution_start_condition.setter
    def pipeline_execution_start_condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_execution_start_condition", value)


@pulumi.input_type
class ImageRecipeBlockDeviceMappingArgs:
    def __init__(__self__, *,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ebs: Optional[pulumi.Input['ImageRecipeBlockDeviceMappingEbsArgs']] = None,
                 no_device: Optional[pulumi.Input[bool]] = None,
                 virtual_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] device_name: Name of the device. For example, `/dev/sda` or `/dev/xvdb`.
        :param pulumi.Input['ImageRecipeBlockDeviceMappingEbsArgs'] ebs: Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
        :param pulumi.Input[bool] no_device: Set to `true` to remove a mapping from the parent image.
        :param pulumi.Input[str] virtual_name: Virtual device name. For example, `ephemeral0`. Instance store volumes are numbered starting from 0.
        """
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if ebs is not None:
            pulumi.set(__self__, "ebs", ebs)
        if no_device is not None:
            pulumi.set(__self__, "no_device", no_device)
        if virtual_name is not None:
            pulumi.set(__self__, "virtual_name", virtual_name)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the device. For example, `/dev/sda` or `/dev/xvdb`.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def ebs(self) -> Optional[pulumi.Input['ImageRecipeBlockDeviceMappingEbsArgs']]:
        """
        Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
        """
        return pulumi.get(self, "ebs")

    @ebs.setter
    def ebs(self, value: Optional[pulumi.Input['ImageRecipeBlockDeviceMappingEbsArgs']]):
        pulumi.set(self, "ebs", value)

    @property
    @pulumi.getter(name="noDevice")
    def no_device(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to remove a mapping from the parent image.
        """
        return pulumi.get(self, "no_device")

    @no_device.setter
    def no_device(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_device", value)

    @property
    @pulumi.getter(name="virtualName")
    def virtual_name(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual device name. For example, `ephemeral0`. Instance store volumes are numbered starting from 0.
        """
        return pulumi.get(self, "virtual_name")

    @virtual_name.setter
    def virtual_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_name", value)


@pulumi.input_type
class ImageRecipeBlockDeviceMappingEbsArgs:
    def __init__(__self__, *,
                 delete_on_termination: Optional[pulumi.Input[str]] = None,
                 encrypted: Optional[pulumi.Input[str]] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 volume_size: Optional[pulumi.Input[int]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] delete_on_termination: Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
        :param pulumi.Input[str] encrypted: Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
        :param pulumi.Input[int] iops: Number of Input/Output (I/O) operations per second to provision for an `io1` or `io2` volume.
        :param pulumi.Input[str] kms_key_id: Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
        :param pulumi.Input[str] snapshot_id: Identifier of the EC2 Volume Snapshot.
        :param pulumi.Input[int] volume_size: Size of the volume, in GiB.
        :param pulumi.Input[str] volume_type: Type of the volume. For example, `gp2` or `io2`.
        """
        if delete_on_termination is not None:
            pulumi.set(__self__, "delete_on_termination", delete_on_termination)
        if encrypted is not None:
            pulumi.set(__self__, "encrypted", encrypted)
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="deleteOnTermination")
    def delete_on_termination(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
        """
        return pulumi.get(self, "delete_on_termination")

    @delete_on_termination.setter
    def delete_on_termination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_on_termination", value)

    @property
    @pulumi.getter
    def encrypted(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
        """
        return pulumi.get(self, "encrypted")

    @encrypted.setter
    def encrypted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encrypted", value)

    @property
    @pulumi.getter
    def iops(self) -> Optional[pulumi.Input[int]]:
        """
        Number of Input/Output (I/O) operations per second to provision for an `io1` or `io2` volume.
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the EC2 Volume Snapshot.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[pulumi.Input[int]]:
        """
        Size of the volume, in GiB.
        """
        return pulumi.get(self, "volume_size")

    @volume_size.setter
    def volume_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_size", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the volume. For example, `gp2` or `io2`.
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)


@pulumi.input_type
class ImageRecipeComponentArgs:
    def __init__(__self__, *,
                 component_arn: pulumi.Input[str]):
        """
        :param pulumi.Input[str] component_arn: Amazon Resource Name (ARN) of the Image Builder Component to associate.
        """
        pulumi.set(__self__, "component_arn", component_arn)

    @property
    @pulumi.getter(name="componentArn")
    def component_arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the Image Builder Component to associate.
        """
        return pulumi.get(self, "component_arn")

    @component_arn.setter
    def component_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "component_arn", value)


@pulumi.input_type
class InfrastructureConfigurationLoggingArgs:
    def __init__(__self__, *,
                 s3_logs: pulumi.Input['InfrastructureConfigurationLoggingS3LogsArgs']):
        """
        :param pulumi.Input['InfrastructureConfigurationLoggingS3LogsArgs'] s3_logs: Configuration block with S3 logging settings. Detailed below.
        """
        pulumi.set(__self__, "s3_logs", s3_logs)

    @property
    @pulumi.getter(name="s3Logs")
    def s3_logs(self) -> pulumi.Input['InfrastructureConfigurationLoggingS3LogsArgs']:
        """
        Configuration block with S3 logging settings. Detailed below.
        """
        return pulumi.get(self, "s3_logs")

    @s3_logs.setter
    def s3_logs(self, value: pulumi.Input['InfrastructureConfigurationLoggingS3LogsArgs']):
        pulumi.set(self, "s3_logs", value)


@pulumi.input_type
class InfrastructureConfigurationLoggingS3LogsArgs:
    def __init__(__self__, *,
                 s3_bucket_name: pulumi.Input[str],
                 s3_key_prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] s3_bucket_name: Name of the S3 Bucket.
        :param pulumi.Input[str] s3_key_prefix: Prefix to use for S3 logs. Defaults to `/`.
        """
        pulumi.set(__self__, "s3_bucket_name", s3_bucket_name)
        if s3_key_prefix is not None:
            pulumi.set(__self__, "s3_key_prefix", s3_key_prefix)

    @property
    @pulumi.getter(name="s3BucketName")
    def s3_bucket_name(self) -> pulumi.Input[str]:
        """
        Name of the S3 Bucket.
        """
        return pulumi.get(self, "s3_bucket_name")

    @s3_bucket_name.setter
    def s3_bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_bucket_name", value)

    @property
    @pulumi.getter(name="s3KeyPrefix")
    def s3_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Prefix to use for S3 logs. Defaults to `/`.
        """
        return pulumi.get(self, "s3_key_prefix")

    @s3_key_prefix.setter
    def s3_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_key_prefix", value)


