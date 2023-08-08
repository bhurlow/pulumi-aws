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

__all__ = ['MonitoringSubscriptionArgs', 'MonitoringSubscription']

@pulumi.input_type
class MonitoringSubscriptionArgs:
    def __init__(__self__, *,
                 distribution_id: pulumi.Input[str],
                 monitoring_subscription: pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs']):
        """
        The set of arguments for constructing a MonitoringSubscription resource.
        :param pulumi.Input[str] distribution_id: The ID of the distribution that you are enabling metrics for.
        :param pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs'] monitoring_subscription: A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
        """
        pulumi.set(__self__, "distribution_id", distribution_id)
        pulumi.set(__self__, "monitoring_subscription", monitoring_subscription)

    @property
    @pulumi.getter(name="distributionId")
    def distribution_id(self) -> pulumi.Input[str]:
        """
        The ID of the distribution that you are enabling metrics for.
        """
        return pulumi.get(self, "distribution_id")

    @distribution_id.setter
    def distribution_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "distribution_id", value)

    @property
    @pulumi.getter(name="monitoringSubscription")
    def monitoring_subscription(self) -> pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs']:
        """
        A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
        """
        return pulumi.get(self, "monitoring_subscription")

    @monitoring_subscription.setter
    def monitoring_subscription(self, value: pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs']):
        pulumi.set(self, "monitoring_subscription", value)


@pulumi.input_type
class _MonitoringSubscriptionState:
    def __init__(__self__, *,
                 distribution_id: Optional[pulumi.Input[str]] = None,
                 monitoring_subscription: Optional[pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs']] = None):
        """
        Input properties used for looking up and filtering MonitoringSubscription resources.
        :param pulumi.Input[str] distribution_id: The ID of the distribution that you are enabling metrics for.
        :param pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs'] monitoring_subscription: A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
        """
        if distribution_id is not None:
            pulumi.set(__self__, "distribution_id", distribution_id)
        if monitoring_subscription is not None:
            pulumi.set(__self__, "monitoring_subscription", monitoring_subscription)

    @property
    @pulumi.getter(name="distributionId")
    def distribution_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the distribution that you are enabling metrics for.
        """
        return pulumi.get(self, "distribution_id")

    @distribution_id.setter
    def distribution_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distribution_id", value)

    @property
    @pulumi.getter(name="monitoringSubscription")
    def monitoring_subscription(self) -> Optional[pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs']]:
        """
        A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
        """
        return pulumi.get(self, "monitoring_subscription")

    @monitoring_subscription.setter
    def monitoring_subscription(self, value: Optional[pulumi.Input['MonitoringSubscriptionMonitoringSubscriptionArgs']]):
        pulumi.set(self, "monitoring_subscription", value)


class MonitoringSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distribution_id: Optional[pulumi.Input[str]] = None,
                 monitoring_subscription: Optional[pulumi.Input[pulumi.InputType['MonitoringSubscriptionMonitoringSubscriptionArgs']]] = None,
                 __props__=None):
        """
        Provides a CloudFront real-time log configuration resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudfront.MonitoringSubscription("example",
            distribution_id=aws_cloudfront_distribution["example"]["id"],
            monitoring_subscription=aws.cloudfront.MonitoringSubscriptionMonitoringSubscriptionArgs(
                realtime_metrics_subscription_config=aws.cloudfront.MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs(
                    realtime_metrics_subscription_status="Enabled",
                ),
            ))
        ```

        ## Import

        terraform import {

         to = aws_cloudfront_monitoring_subscription.example

         id = "E3QYSUHO4VYRGB" } Using `pulumi import`, import CloudFront monitoring subscription using the id. For exampleconsole % pulumi import aws_cloudfront_monitoring_subscription.example E3QYSUHO4VYRGB

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distribution_id: The ID of the distribution that you are enabling metrics for.
        :param pulumi.Input[pulumi.InputType['MonitoringSubscriptionMonitoringSubscriptionArgs']] monitoring_subscription: A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MonitoringSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CloudFront real-time log configuration resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudfront.MonitoringSubscription("example",
            distribution_id=aws_cloudfront_distribution["example"]["id"],
            monitoring_subscription=aws.cloudfront.MonitoringSubscriptionMonitoringSubscriptionArgs(
                realtime_metrics_subscription_config=aws.cloudfront.MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs(
                    realtime_metrics_subscription_status="Enabled",
                ),
            ))
        ```

        ## Import

        terraform import {

         to = aws_cloudfront_monitoring_subscription.example

         id = "E3QYSUHO4VYRGB" } Using `pulumi import`, import CloudFront monitoring subscription using the id. For exampleconsole % pulumi import aws_cloudfront_monitoring_subscription.example E3QYSUHO4VYRGB

        :param str resource_name: The name of the resource.
        :param MonitoringSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MonitoringSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distribution_id: Optional[pulumi.Input[str]] = None,
                 monitoring_subscription: Optional[pulumi.Input[pulumi.InputType['MonitoringSubscriptionMonitoringSubscriptionArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MonitoringSubscriptionArgs.__new__(MonitoringSubscriptionArgs)

            if distribution_id is None and not opts.urn:
                raise TypeError("Missing required property 'distribution_id'")
            __props__.__dict__["distribution_id"] = distribution_id
            if monitoring_subscription is None and not opts.urn:
                raise TypeError("Missing required property 'monitoring_subscription'")
            __props__.__dict__["monitoring_subscription"] = monitoring_subscription
        super(MonitoringSubscription, __self__).__init__(
            'aws:cloudfront/monitoringSubscription:MonitoringSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            distribution_id: Optional[pulumi.Input[str]] = None,
            monitoring_subscription: Optional[pulumi.Input[pulumi.InputType['MonitoringSubscriptionMonitoringSubscriptionArgs']]] = None) -> 'MonitoringSubscription':
        """
        Get an existing MonitoringSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distribution_id: The ID of the distribution that you are enabling metrics for.
        :param pulumi.Input[pulumi.InputType['MonitoringSubscriptionMonitoringSubscriptionArgs']] monitoring_subscription: A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MonitoringSubscriptionState.__new__(_MonitoringSubscriptionState)

        __props__.__dict__["distribution_id"] = distribution_id
        __props__.__dict__["monitoring_subscription"] = monitoring_subscription
        return MonitoringSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="distributionId")
    def distribution_id(self) -> pulumi.Output[str]:
        """
        The ID of the distribution that you are enabling metrics for.
        """
        return pulumi.get(self, "distribution_id")

    @property
    @pulumi.getter(name="monitoringSubscription")
    def monitoring_subscription(self) -> pulumi.Output['outputs.MonitoringSubscriptionMonitoringSubscription']:
        """
        A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
        """
        return pulumi.get(self, "monitoring_subscription")

