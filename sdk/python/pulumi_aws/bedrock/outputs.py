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

__all__ = [
    'CustomModelOutputDataConfig',
    'CustomModelTimeouts',
    'CustomModelTrainingDataConfig',
    'CustomModelTrainingMetric',
    'CustomModelValidationDataConfig',
    'CustomModelValidationDataConfigValidator',
    'CustomModelValidationMetric',
    'CustomModelVpcConfig',
    'ProvisionedModelThroughputTimeouts',
    'GetCustomModelOutputDataConfigResult',
    'GetCustomModelTrainingDataConfigResult',
    'GetCustomModelTrainingMetricResult',
    'GetCustomModelValidationDataConfigResult',
    'GetCustomModelValidationMetricResult',
    'GetCustomModelsModelSummaryResult',
]

@pulumi.output_type
class CustomModelOutputDataConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Uri":
            suggest = "s3_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomModelOutputDataConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomModelOutputDataConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomModelOutputDataConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s3_uri: str):
        """
        :param str s3_uri: The S3 URI where the validation data is stored.
        """
        pulumi.set(__self__, "s3_uri", s3_uri)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> str:
        """
        The S3 URI where the validation data is stored.
        """
        return pulumi.get(self, "s3_uri")


@pulumi.output_type
class CustomModelTimeouts(dict):
    def __init__(__self__, *,
                 create: Optional[str] = None,
                 delete: Optional[str] = None):
        """
        :param str create: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        :param str delete: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
        """
        if create is not None:
            pulumi.set(__self__, "create", create)
        if delete is not None:
            pulumi.set(__self__, "delete", delete)

    @property
    @pulumi.getter
    def create(self) -> Optional[str]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        """
        return pulumi.get(self, "create")

    @property
    @pulumi.getter
    def delete(self) -> Optional[str]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
        """
        return pulumi.get(self, "delete")


@pulumi.output_type
class CustomModelTrainingDataConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Uri":
            suggest = "s3_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomModelTrainingDataConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomModelTrainingDataConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomModelTrainingDataConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s3_uri: str):
        """
        :param str s3_uri: The S3 URI where the validation data is stored.
        """
        pulumi.set(__self__, "s3_uri", s3_uri)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> str:
        """
        The S3 URI where the validation data is stored.
        """
        return pulumi.get(self, "s3_uri")


@pulumi.output_type
class CustomModelTrainingMetric(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "trainingLoss":
            suggest = "training_loss"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomModelTrainingMetric. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomModelTrainingMetric.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomModelTrainingMetric.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 training_loss: float):
        """
        :param float training_loss: Loss metric associated with the customization job.
        """
        pulumi.set(__self__, "training_loss", training_loss)

    @property
    @pulumi.getter(name="trainingLoss")
    def training_loss(self) -> float:
        """
        Loss metric associated with the customization job.
        """
        return pulumi.get(self, "training_loss")


@pulumi.output_type
class CustomModelValidationDataConfig(dict):
    def __init__(__self__, *,
                 validators: Optional[Sequence['outputs.CustomModelValidationDataConfigValidator']] = None):
        """
        :param Sequence['CustomModelValidationDataConfigValidatorArgs'] validators: Information about the validators.
        """
        if validators is not None:
            pulumi.set(__self__, "validators", validators)

    @property
    @pulumi.getter
    def validators(self) -> Optional[Sequence['outputs.CustomModelValidationDataConfigValidator']]:
        """
        Information about the validators.
        """
        return pulumi.get(self, "validators")


@pulumi.output_type
class CustomModelValidationDataConfigValidator(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Uri":
            suggest = "s3_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomModelValidationDataConfigValidator. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomModelValidationDataConfigValidator.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomModelValidationDataConfigValidator.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s3_uri: str):
        """
        :param str s3_uri: The S3 URI where the validation data is stored.
        """
        pulumi.set(__self__, "s3_uri", s3_uri)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> str:
        """
        The S3 URI where the validation data is stored.
        """
        return pulumi.get(self, "s3_uri")


@pulumi.output_type
class CustomModelValidationMetric(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "validationLoss":
            suggest = "validation_loss"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomModelValidationMetric. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomModelValidationMetric.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomModelValidationMetric.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 validation_loss: float):
        """
        :param float validation_loss: The validation loss associated with the validator.
        """
        pulumi.set(__self__, "validation_loss", validation_loss)

    @property
    @pulumi.getter(name="validationLoss")
    def validation_loss(self) -> float:
        """
        The validation loss associated with the validator.
        """
        return pulumi.get(self, "validation_loss")


@pulumi.output_type
class CustomModelVpcConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroupIds":
            suggest = "security_group_ids"
        elif key == "subnetIds":
            suggest = "subnet_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomModelVpcConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomModelVpcConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomModelVpcConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_group_ids: Sequence[str],
                 subnet_ids: Sequence[str]):
        """
        :param Sequence[str] security_group_ids: VPC configuration security group IDs.
        :param Sequence[str] subnet_ids: VPC configuration subnets.
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[str]:
        """
        VPC configuration security group IDs.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Sequence[str]:
        """
        VPC configuration subnets.
        """
        return pulumi.get(self, "subnet_ids")


@pulumi.output_type
class ProvisionedModelThroughputTimeouts(dict):
    def __init__(__self__, *,
                 create: Optional[str] = None):
        """
        :param str create: A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        """
        if create is not None:
            pulumi.set(__self__, "create", create)

    @property
    @pulumi.getter
    def create(self) -> Optional[str]:
        """
        A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        """
        return pulumi.get(self, "create")


@pulumi.output_type
class GetCustomModelOutputDataConfigResult(dict):
    def __init__(__self__, *,
                 s3_uri: str):
        """
        :param str s3_uri: The S3 URI where the validation data is stored..
        """
        pulumi.set(__self__, "s3_uri", s3_uri)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> str:
        """
        The S3 URI where the validation data is stored..
        """
        return pulumi.get(self, "s3_uri")


@pulumi.output_type
class GetCustomModelTrainingDataConfigResult(dict):
    def __init__(__self__, *,
                 s3_uri: str):
        """
        :param str s3_uri: The S3 URI where the validation data is stored..
        """
        pulumi.set(__self__, "s3_uri", s3_uri)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> str:
        """
        The S3 URI where the validation data is stored..
        """
        return pulumi.get(self, "s3_uri")


@pulumi.output_type
class GetCustomModelTrainingMetricResult(dict):
    def __init__(__self__, *,
                 training_loss: float):
        """
        :param float training_loss: Loss metric associated with the customization job.
        """
        pulumi.set(__self__, "training_loss", training_loss)

    @property
    @pulumi.getter(name="trainingLoss")
    def training_loss(self) -> float:
        """
        Loss metric associated with the customization job.
        """
        return pulumi.get(self, "training_loss")


@pulumi.output_type
class GetCustomModelValidationDataConfigResult(dict):
    def __init__(__self__, *,
                 validators: Sequence[Any]):
        """
        :param Sequence[Any] validators: Information about the validators.
        """
        pulumi.set(__self__, "validators", validators)

    @property
    @pulumi.getter
    def validators(self) -> Sequence[Any]:
        """
        Information about the validators.
        """
        return pulumi.get(self, "validators")


@pulumi.output_type
class GetCustomModelValidationMetricResult(dict):
    def __init__(__self__, *,
                 validation_loss: float):
        """
        :param float validation_loss: The validation loss associated with the validator.
        """
        pulumi.set(__self__, "validation_loss", validation_loss)

    @property
    @pulumi.getter(name="validationLoss")
    def validation_loss(self) -> float:
        """
        The validation loss associated with the validator.
        """
        return pulumi.get(self, "validation_loss")


@pulumi.output_type
class GetCustomModelsModelSummaryResult(dict):
    def __init__(__self__, *,
                 creation_time: str,
                 model_arn: str,
                 model_name: str):
        """
        :param str creation_time: Creation time of the model.
        :param str model_arn: The ARN of the custom model.
        :param str model_name: The name of the custom model.
        """
        pulumi.set(__self__, "creation_time", creation_time)
        pulumi.set(__self__, "model_arn", model_arn)
        pulumi.set(__self__, "model_name", model_name)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Creation time of the model.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="modelArn")
    def model_arn(self) -> str:
        """
        The ARN of the custom model.
        """
        return pulumi.get(self, "model_arn")

    @property
    @pulumi.getter(name="modelName")
    def model_name(self) -> str:
        """
        The name of the custom model.
        """
        return pulumi.get(self, "model_name")


