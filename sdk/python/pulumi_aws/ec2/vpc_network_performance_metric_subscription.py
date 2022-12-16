# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcNetworkPerformanceMetricSubscriptionArgs', 'VpcNetworkPerformanceMetricSubscription']

@pulumi.input_type
class VpcNetworkPerformanceMetricSubscriptionArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str],
                 source: pulumi.Input[str],
                 metric: Optional[pulumi.Input[str]] = None,
                 statistic: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcNetworkPerformanceMetricSubscription resource.
        :param pulumi.Input[str] destination: The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
        :param pulumi.Input[str] source: The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
        :param pulumi.Input[str] metric: The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
        :param pulumi.Input[str] statistic: The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "source", source)
        if metric is not None:
            pulumi.set(__self__, "metric", metric)
        if statistic is not None:
            pulumi.set(__self__, "statistic", statistic)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def metric(self) -> Optional[pulumi.Input[str]]:
        """
        The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter
    def statistic(self) -> Optional[pulumi.Input[str]]:
        """
        The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
        """
        return pulumi.get(self, "statistic")

    @statistic.setter
    def statistic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "statistic", value)


@pulumi.input_type
class _VpcNetworkPerformanceMetricSubscriptionState:
    def __init__(__self__, *,
                 destination: Optional[pulumi.Input[str]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 statistic: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcNetworkPerformanceMetricSubscription resources.
        :param pulumi.Input[str] destination: The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
        :param pulumi.Input[str] metric: The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
        :param pulumi.Input[str] period: The data aggregation time for the subscription.
        :param pulumi.Input[str] source: The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
        :param pulumi.Input[str] statistic: The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
        """
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if metric is not None:
            pulumi.set(__self__, "metric", metric)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if statistic is not None:
            pulumi.set(__self__, "statistic", statistic)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input[str]]:
        """
        The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def metric(self) -> Optional[pulumi.Input[str]]:
        """
        The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[str]]:
        """
        The data aggregation time for the subscription.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def statistic(self) -> Optional[pulumi.Input[str]]:
        """
        The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
        """
        return pulumi.get(self, "statistic")

    @statistic.setter
    def statistic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "statistic", value)


class VpcNetworkPerformanceMetricSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 statistic: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage an Infrastructure Performance subscription.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcNetworkPerformanceMetricSubscription("example",
            destination="us-west-1",
            source="us-east-1")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination: The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
        :param pulumi.Input[str] metric: The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
        :param pulumi.Input[str] source: The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
        :param pulumi.Input[str] statistic: The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcNetworkPerformanceMetricSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage an Infrastructure Performance subscription.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcNetworkPerformanceMetricSubscription("example",
            destination="us-west-1",
            source="us-east-1")
        ```

        :param str resource_name: The name of the resource.
        :param VpcNetworkPerformanceMetricSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcNetworkPerformanceMetricSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 metric: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 statistic: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcNetworkPerformanceMetricSubscriptionArgs.__new__(VpcNetworkPerformanceMetricSubscriptionArgs)

            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__.__dict__["destination"] = destination
            __props__.__dict__["metric"] = metric
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["statistic"] = statistic
            __props__.__dict__["period"] = None
        super(VpcNetworkPerformanceMetricSubscription, __self__).__init__(
            'aws:ec2/vpcNetworkPerformanceMetricSubscription:VpcNetworkPerformanceMetricSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination: Optional[pulumi.Input[str]] = None,
            metric: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            statistic: Optional[pulumi.Input[str]] = None) -> 'VpcNetworkPerformanceMetricSubscription':
        """
        Get an existing VpcNetworkPerformanceMetricSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination: The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
        :param pulumi.Input[str] metric: The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
        :param pulumi.Input[str] period: The data aggregation time for the subscription.
        :param pulumi.Input[str] source: The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
        :param pulumi.Input[str] statistic: The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcNetworkPerformanceMetricSubscriptionState.__new__(_VpcNetworkPerformanceMetricSubscriptionState)

        __props__.__dict__["destination"] = destination
        __props__.__dict__["metric"] = metric
        __props__.__dict__["period"] = period
        __props__.__dict__["source"] = source
        __props__.__dict__["statistic"] = statistic
        return VpcNetworkPerformanceMetricSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[str]:
        """
        The target Region or Availability Zone that the metric subscription is enabled for. For example, `eu-west-1`.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def metric(self) -> pulumi.Output[Optional[str]]:
        """
        The metric used for the enabled subscription. Valid values: `aggregate-latency`. Default: `aggregate-latency`.
        """
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[str]:
        """
        The data aggregation time for the subscription.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        The source Region or Availability Zone that the metric subscription is enabled for. For example, `us-east-1`.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def statistic(self) -> pulumi.Output[Optional[str]]:
        """
        The statistic used for the enabled subscription. Valid values: `p50`. Default: `p50`.
        """
        return pulumi.get(self, "statistic")

