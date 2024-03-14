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

__all__ = ['DetectorFeatureArgs', 'DetectorFeature']

@pulumi.input_type
class DetectorFeatureArgs:
    def __init__(__self__, *,
                 detector_id: pulumi.Input[str],
                 status: pulumi.Input[str],
                 additional_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DetectorFeature resource.
        :param pulumi.Input[str] detector_id: Amazon GuardDuty detector ID.
        :param pulumi.Input[str] status: The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
        :param pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]] additional_configurations: Additional feature configuration block. See below.
        :param pulumi.Input[str] name: The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`.
        """
        pulumi.set(__self__, "detector_id", detector_id)
        pulumi.set(__self__, "status", status)
        if additional_configurations is not None:
            pulumi.set(__self__, "additional_configurations", additional_configurations)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Input[str]:
        """
        Amazon GuardDuty detector ID.
        """
        return pulumi.get(self, "detector_id")

    @detector_id.setter
    def detector_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "detector_id", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="additionalConfigurations")
    def additional_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]]]:
        """
        Additional feature configuration block. See below.
        """
        return pulumi.get(self, "additional_configurations")

    @additional_configurations.setter
    def additional_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]]]):
        pulumi.set(self, "additional_configurations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DetectorFeatureState:
    def __init__(__self__, *,
                 additional_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DetectorFeature resources.
        :param pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]] additional_configurations: Additional feature configuration block. See below.
        :param pulumi.Input[str] detector_id: Amazon GuardDuty detector ID.
        :param pulumi.Input[str] name: The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`.
        :param pulumi.Input[str] status: The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
        """
        if additional_configurations is not None:
            pulumi.set(__self__, "additional_configurations", additional_configurations)
        if detector_id is not None:
            pulumi.set(__self__, "detector_id", detector_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="additionalConfigurations")
    def additional_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]]]:
        """
        Additional feature configuration block. See below.
        """
        return pulumi.get(self, "additional_configurations")

    @additional_configurations.setter
    def additional_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DetectorFeatureAdditionalConfigurationArgs']]]]):
        pulumi.set(self, "additional_configurations", value)

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon GuardDuty detector ID.
        """
        return pulumi.get(self, "detector_id")

    @detector_id.setter
    def detector_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "detector_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class DetectorFeature(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorFeatureAdditionalConfigurationArgs']]]]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage a single Amazon GuardDuty [detector feature](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-features-activation-model.html#guardduty-features).

        > **NOTE:** Deleting this resource does not disable the detector feature, the resource in simply removed from state instead.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.guardduty.Detector("example", enable=True)
        eks_runtime_monitoring = aws.guardduty.DetectorFeature("eks_runtime_monitoring",
            detector_id=example.id,
            name="EKS_RUNTIME_MONITORING",
            status="ENABLED",
            additional_configurations=[aws.guardduty.DetectorFeatureAdditionalConfigurationArgs(
                name="EKS_ADDON_MANAGEMENT",
                status="ENABLED",
            )])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorFeatureAdditionalConfigurationArgs']]]] additional_configurations: Additional feature configuration block. See below.
        :param pulumi.Input[str] detector_id: Amazon GuardDuty detector ID.
        :param pulumi.Input[str] name: The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`.
        :param pulumi.Input[str] status: The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DetectorFeatureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a single Amazon GuardDuty [detector feature](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-features-activation-model.html#guardduty-features).

        > **NOTE:** Deleting this resource does not disable the detector feature, the resource in simply removed from state instead.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.guardduty.Detector("example", enable=True)
        eks_runtime_monitoring = aws.guardduty.DetectorFeature("eks_runtime_monitoring",
            detector_id=example.id,
            name="EKS_RUNTIME_MONITORING",
            status="ENABLED",
            additional_configurations=[aws.guardduty.DetectorFeatureAdditionalConfigurationArgs(
                name="EKS_ADDON_MANAGEMENT",
                status="ENABLED",
            )])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param DetectorFeatureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DetectorFeatureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorFeatureAdditionalConfigurationArgs']]]]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DetectorFeatureArgs.__new__(DetectorFeatureArgs)

            __props__.__dict__["additional_configurations"] = additional_configurations
            if detector_id is None and not opts.urn:
                raise TypeError("Missing required property 'detector_id'")
            __props__.__dict__["detector_id"] = detector_id
            __props__.__dict__["name"] = name
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
        super(DetectorFeature, __self__).__init__(
            'aws:guardduty/detectorFeature:DetectorFeature',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorFeatureAdditionalConfigurationArgs']]]]] = None,
            detector_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'DetectorFeature':
        """
        Get an existing DetectorFeature resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorFeatureAdditionalConfigurationArgs']]]] additional_configurations: Additional feature configuration block. See below.
        :param pulumi.Input[str] detector_id: Amazon GuardDuty detector ID.
        :param pulumi.Input[str] name: The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`.
        :param pulumi.Input[str] status: The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DetectorFeatureState.__new__(_DetectorFeatureState)

        __props__.__dict__["additional_configurations"] = additional_configurations
        __props__.__dict__["detector_id"] = detector_id
        __props__.__dict__["name"] = name
        __props__.__dict__["status"] = status
        return DetectorFeature(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalConfigurations")
    def additional_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.DetectorFeatureAdditionalConfiguration']]]:
        """
        Additional feature configuration block. See below.
        """
        return pulumi.get(self, "additional_configurations")

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Output[str]:
        """
        Amazon GuardDuty detector ID.
        """
        return pulumi.get(self, "detector_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "status")

