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

__all__ = ['DetectorArgs', 'Detector']

@pulumi.input_type
class DetectorArgs:
    def __init__(__self__, *,
                 datasources: Optional[pulumi.Input['DetectorDatasourcesArgs']] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 finding_publishing_frequency: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Detector resource.
        :param pulumi.Input['DetectorDatasourcesArgs'] datasources: Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        :param pulumi.Input[bool] enable: Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
        :param pulumi.Input[str] finding_publishing_frequency: Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if datasources is not None:
            pulumi.set(__self__, "datasources", datasources)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if finding_publishing_frequency is not None:
            pulumi.set(__self__, "finding_publishing_frequency", finding_publishing_frequency)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def datasources(self) -> Optional[pulumi.Input['DetectorDatasourcesArgs']]:
        """
        Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        """
        return pulumi.get(self, "datasources")

    @datasources.setter
    def datasources(self, value: Optional[pulumi.Input['DetectorDatasourcesArgs']]):
        pulumi.set(self, "datasources", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="findingPublishingFrequency")
    def finding_publishing_frequency(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        """
        return pulumi.get(self, "finding_publishing_frequency")

    @finding_publishing_frequency.setter
    def finding_publishing_frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "finding_publishing_frequency", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DetectorState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 datasources: Optional[pulumi.Input['DetectorDatasourcesArgs']] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 finding_publishing_frequency: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Detector resources.
        :param pulumi.Input[str] account_id: The AWS account ID of the GuardDuty detector
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the GuardDuty detector
        :param pulumi.Input['DetectorDatasourcesArgs'] datasources: Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        :param pulumi.Input[bool] enable: Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
        :param pulumi.Input[str] finding_publishing_frequency: Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if datasources is not None:
            pulumi.set(__self__, "datasources", datasources)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if finding_publishing_frequency is not None:
            pulumi.set(__self__, "finding_publishing_frequency", finding_publishing_frequency)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS account ID of the GuardDuty detector
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the GuardDuty detector
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def datasources(self) -> Optional[pulumi.Input['DetectorDatasourcesArgs']]:
        """
        Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        """
        return pulumi.get(self, "datasources")

    @datasources.setter
    def datasources(self, value: Optional[pulumi.Input['DetectorDatasourcesArgs']]):
        pulumi.set(self, "datasources", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="findingPublishingFrequency")
    def finding_publishing_frequency(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        """
        return pulumi.get(self, "finding_publishing_frequency")

    @finding_publishing_frequency.setter
    def finding_publishing_frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "finding_publishing_frequency", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Detector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datasources: Optional[pulumi.Input[pulumi.InputType['DetectorDatasourcesArgs']]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 finding_publishing_frequency: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage a GuardDuty detector.

        > **NOTE:** Deleting this resource is equivalent to "disabling" GuardDuty for an AWS region, which removes all existing findings. You can set the `enable` attribute to `false` to instead "suspend" monitoring and feedback reporting while keeping existing data. See the [Suspending or Disabling Amazon GuardDuty documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_suspend-disable.html) for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_detector = aws.guardduty.Detector("myDetector",
            datasources=aws.guardduty.DetectorDatasourcesArgs(
                kubernetes=aws.guardduty.DetectorDatasourcesKubernetesArgs(
                    audit_logs=aws.guardduty.DetectorDatasourcesKubernetesAuditLogsArgs(
                        enable=False,
                    ),
                ),
                malware_protection=aws.guardduty.DetectorDatasourcesMalwareProtectionArgs(
                    scan_ec2_instance_with_findings=aws.guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs(
                        ebs_volumes=aws.guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs(
                            enable=True,
                        ),
                    ),
                ),
                s3_logs=aws.guardduty.DetectorDatasourcesS3LogsArgs(
                    enable=True,
                ),
            ),
            enable=True)
        ```

        ## Import

        Using `pulumi import`, import GuardDuty detectors using the detector ID. For example:

        ```sh
         $ pulumi import aws:guardduty/detector:Detector MyDetector 00b00fd5aecc0ab60a708659477e9617
        ```
         The ID of the detector can be retrieved via the [AWS CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/guardduty/list-detectors.html) using `aws guardduty list-detectors`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DetectorDatasourcesArgs']] datasources: Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        :param pulumi.Input[bool] enable: Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
        :param pulumi.Input[str] finding_publishing_frequency: Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DetectorArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a GuardDuty detector.

        > **NOTE:** Deleting this resource is equivalent to "disabling" GuardDuty for an AWS region, which removes all existing findings. You can set the `enable` attribute to `false` to instead "suspend" monitoring and feedback reporting while keeping existing data. See the [Suspending or Disabling Amazon GuardDuty documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_suspend-disable.html) for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_detector = aws.guardduty.Detector("myDetector",
            datasources=aws.guardduty.DetectorDatasourcesArgs(
                kubernetes=aws.guardduty.DetectorDatasourcesKubernetesArgs(
                    audit_logs=aws.guardduty.DetectorDatasourcesKubernetesAuditLogsArgs(
                        enable=False,
                    ),
                ),
                malware_protection=aws.guardduty.DetectorDatasourcesMalwareProtectionArgs(
                    scan_ec2_instance_with_findings=aws.guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs(
                        ebs_volumes=aws.guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs(
                            enable=True,
                        ),
                    ),
                ),
                s3_logs=aws.guardduty.DetectorDatasourcesS3LogsArgs(
                    enable=True,
                ),
            ),
            enable=True)
        ```

        ## Import

        Using `pulumi import`, import GuardDuty detectors using the detector ID. For example:

        ```sh
         $ pulumi import aws:guardduty/detector:Detector MyDetector 00b00fd5aecc0ab60a708659477e9617
        ```
         The ID of the detector can be retrieved via the [AWS CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/guardduty/list-detectors.html) using `aws guardduty list-detectors`.

        :param str resource_name: The name of the resource.
        :param DetectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DetectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datasources: Optional[pulumi.Input[pulumi.InputType['DetectorDatasourcesArgs']]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 finding_publishing_frequency: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DetectorArgs.__new__(DetectorArgs)

            __props__.__dict__["datasources"] = datasources
            __props__.__dict__["enable"] = enable
            __props__.__dict__["finding_publishing_frequency"] = finding_publishing_frequency
            __props__.__dict__["tags"] = tags
            __props__.__dict__["account_id"] = None
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Detector, __self__).__init__(
            'aws:guardduty/detector:Detector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            datasources: Optional[pulumi.Input[pulumi.InputType['DetectorDatasourcesArgs']]] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            finding_publishing_frequency: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Detector':
        """
        Get an existing Detector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The AWS account ID of the GuardDuty detector
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the GuardDuty detector
        :param pulumi.Input[pulumi.InputType['DetectorDatasourcesArgs']] datasources: Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        :param pulumi.Input[bool] enable: Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
        :param pulumi.Input[str] finding_publishing_frequency: Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DetectorState.__new__(_DetectorState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["datasources"] = datasources
        __props__.__dict__["enable"] = enable
        __props__.__dict__["finding_publishing_frequency"] = finding_publishing_frequency
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Detector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID of the GuardDuty detector
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the GuardDuty detector
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def datasources(self) -> pulumi.Output['outputs.DetectorDatasources']:
        """
        Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        """
        return pulumi.get(self, "datasources")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="findingPublishingFrequency")
    def finding_publishing_frequency(self) -> pulumi.Output[str]:
        """
        Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        """
        return pulumi.get(self, "finding_publishing_frequency")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

