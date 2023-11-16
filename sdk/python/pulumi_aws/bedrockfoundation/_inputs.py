# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetModelsModelSummaryArgs',
]

@pulumi.input_type
class GetModelsModelSummaryArgs:
    def __init__(__self__, *,
                 customizations_supporteds: Sequence[str],
                 inference_types_supporteds: Sequence[str],
                 input_modalities: Sequence[str],
                 model_arn: str,
                 model_id: str,
                 model_name: str,
                 output_modalities: Sequence[str],
                 provider_name: str,
                 response_streaming_supported: bool):
        """
        :param Sequence[str] customizations_supporteds: Customizations that the model supports.
        :param Sequence[str] inference_types_supporteds: Inference types that the model supports.
        :param Sequence[str] input_modalities: Input modalities that the model supports.
        :param str model_arn: Model ARN.
        :param str model_id: Model identifier.
        :param str model_name: Model name.
        :param Sequence[str] output_modalities: Output modalities that the model supports.
        :param str provider_name: Model provider name.
        :param bool response_streaming_supported: Indicates whether the model supports streaming.
        """
        pulumi.set(__self__, "customizations_supporteds", customizations_supporteds)
        pulumi.set(__self__, "inference_types_supporteds", inference_types_supporteds)
        pulumi.set(__self__, "input_modalities", input_modalities)
        pulumi.set(__self__, "model_arn", model_arn)
        pulumi.set(__self__, "model_id", model_id)
        pulumi.set(__self__, "model_name", model_name)
        pulumi.set(__self__, "output_modalities", output_modalities)
        pulumi.set(__self__, "provider_name", provider_name)
        pulumi.set(__self__, "response_streaming_supported", response_streaming_supported)

    @property
    @pulumi.getter(name="customizationsSupporteds")
    def customizations_supporteds(self) -> Sequence[str]:
        """
        Customizations that the model supports.
        """
        return pulumi.get(self, "customizations_supporteds")

    @customizations_supporteds.setter
    def customizations_supporteds(self, value: Sequence[str]):
        pulumi.set(self, "customizations_supporteds", value)

    @property
    @pulumi.getter(name="inferenceTypesSupporteds")
    def inference_types_supporteds(self) -> Sequence[str]:
        """
        Inference types that the model supports.
        """
        return pulumi.get(self, "inference_types_supporteds")

    @inference_types_supporteds.setter
    def inference_types_supporteds(self, value: Sequence[str]):
        pulumi.set(self, "inference_types_supporteds", value)

    @property
    @pulumi.getter(name="inputModalities")
    def input_modalities(self) -> Sequence[str]:
        """
        Input modalities that the model supports.
        """
        return pulumi.get(self, "input_modalities")

    @input_modalities.setter
    def input_modalities(self, value: Sequence[str]):
        pulumi.set(self, "input_modalities", value)

    @property
    @pulumi.getter(name="modelArn")
    def model_arn(self) -> str:
        """
        Model ARN.
        """
        return pulumi.get(self, "model_arn")

    @model_arn.setter
    def model_arn(self, value: str):
        pulumi.set(self, "model_arn", value)

    @property
    @pulumi.getter(name="modelId")
    def model_id(self) -> str:
        """
        Model identifier.
        """
        return pulumi.get(self, "model_id")

    @model_id.setter
    def model_id(self, value: str):
        pulumi.set(self, "model_id", value)

    @property
    @pulumi.getter(name="modelName")
    def model_name(self) -> str:
        """
        Model name.
        """
        return pulumi.get(self, "model_name")

    @model_name.setter
    def model_name(self, value: str):
        pulumi.set(self, "model_name", value)

    @property
    @pulumi.getter(name="outputModalities")
    def output_modalities(self) -> Sequence[str]:
        """
        Output modalities that the model supports.
        """
        return pulumi.get(self, "output_modalities")

    @output_modalities.setter
    def output_modalities(self, value: Sequence[str]):
        pulumi.set(self, "output_modalities", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> str:
        """
        Model provider name.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: str):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter(name="responseStreamingSupported")
    def response_streaming_supported(self) -> bool:
        """
        Indicates whether the model supports streaming.
        """
        return pulumi.get(self, "response_streaming_supported")

    @response_streaming_supported.setter
    def response_streaming_supported(self, value: bool):
        pulumi.set(self, "response_streaming_supported", value)


