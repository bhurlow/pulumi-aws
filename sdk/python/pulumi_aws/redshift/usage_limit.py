# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UsageLimitArgs', 'UsageLimit']

@pulumi.input_type
class UsageLimitArgs:
    def __init__(__self__, *,
                 amount: pulumi.Input[int],
                 cluster_identifier: pulumi.Input[str],
                 feature_type: pulumi.Input[str],
                 limit_type: pulumi.Input[str],
                 breach_action: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a UsageLimit resource.
        :param pulumi.Input[int] amount: The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
        :param pulumi.Input[str] cluster_identifier: The identifier of the cluster that you want to limit usage.
        :param pulumi.Input[str] feature_type: The Amazon Redshift feature that you want to limit. Valid values are `spectrum`, `concurrency-scaling`, and `cross-region-datasharing`.
        :param pulumi.Input[str] limit_type: The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is `spectrum`, then LimitType must be `data-scanned`. If FeatureType is `concurrency-scaling`, then LimitType must be `time`. If FeatureType is `cross-region-datasharing`, then LimitType must be `data-scanned`. Valid values are `data-scanned`, and `time`.
        :param pulumi.Input[str] breach_action: The action that Amazon Redshift takes when the limit is reached. The default is `log`. Valid values are `log`, `emit-metric`, and `disable`.
        :param pulumi.Input[str] period: The time period that the amount applies to. A weekly period begins on Sunday. The default is `monthly`. Valid values are `daily`, `weekly`, and `monthly`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        pulumi.set(__self__, "amount", amount)
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        pulumi.set(__self__, "feature_type", feature_type)
        pulumi.set(__self__, "limit_type", limit_type)
        if breach_action is not None:
            pulumi.set(__self__, "breach_action", breach_action)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Input[int]:
        """
        The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
        """
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: pulumi.Input[int]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the cluster that you want to limit usage.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_identifier", value)

    @property
    @pulumi.getter(name="featureType")
    def feature_type(self) -> pulumi.Input[str]:
        """
        The Amazon Redshift feature that you want to limit. Valid values are `spectrum`, `concurrency-scaling`, and `cross-region-datasharing`.
        """
        return pulumi.get(self, "feature_type")

    @feature_type.setter
    def feature_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "feature_type", value)

    @property
    @pulumi.getter(name="limitType")
    def limit_type(self) -> pulumi.Input[str]:
        """
        The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is `spectrum`, then LimitType must be `data-scanned`. If FeatureType is `concurrency-scaling`, then LimitType must be `time`. If FeatureType is `cross-region-datasharing`, then LimitType must be `data-scanned`. Valid values are `data-scanned`, and `time`.
        """
        return pulumi.get(self, "limit_type")

    @limit_type.setter
    def limit_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "limit_type", value)

    @property
    @pulumi.getter(name="breachAction")
    def breach_action(self) -> Optional[pulumi.Input[str]]:
        """
        The action that Amazon Redshift takes when the limit is reached. The default is `log`. Valid values are `log`, `emit-metric`, and `disable`.
        """
        return pulumi.get(self, "breach_action")

    @breach_action.setter
    def breach_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "breach_action", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[str]]:
        """
        The time period that the amount applies to. A weekly period begins on Sunday. The default is `monthly`. Valid values are `daily`, `weekly`, and `monthly`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _UsageLimitState:
    def __init__(__self__, *,
                 amount: Optional[pulumi.Input[int]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 breach_action: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 feature_type: Optional[pulumi.Input[str]] = None,
                 limit_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering UsageLimit resources.
        :param pulumi.Input[int] amount: The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Redshift Usage Limit.
        :param pulumi.Input[str] breach_action: The action that Amazon Redshift takes when the limit is reached. The default is `log`. Valid values are `log`, `emit-metric`, and `disable`.
        :param pulumi.Input[str] cluster_identifier: The identifier of the cluster that you want to limit usage.
        :param pulumi.Input[str] feature_type: The Amazon Redshift feature that you want to limit. Valid values are `spectrum`, `concurrency-scaling`, and `cross-region-datasharing`.
        :param pulumi.Input[str] limit_type: The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is `spectrum`, then LimitType must be `data-scanned`. If FeatureType is `concurrency-scaling`, then LimitType must be `time`. If FeatureType is `cross-region-datasharing`, then LimitType must be `data-scanned`. Valid values are `data-scanned`, and `time`.
        :param pulumi.Input[str] period: The time period that the amount applies to. A weekly period begins on Sunday. The default is `monthly`. Valid values are `daily`, `weekly`, and `monthly`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        if amount is not None:
            pulumi.set(__self__, "amount", amount)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if breach_action is not None:
            pulumi.set(__self__, "breach_action", breach_action)
        if cluster_identifier is not None:
            pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        if feature_type is not None:
            pulumi.set(__self__, "feature_type", feature_type)
        if limit_type is not None:
            pulumi.set(__self__, "limit_type", limit_type)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def amount(self) -> Optional[pulumi.Input[int]]:
        """
        The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
        """
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Redshift Usage Limit.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="breachAction")
    def breach_action(self) -> Optional[pulumi.Input[str]]:
        """
        The action that Amazon Redshift takes when the limit is reached. The default is `log`. Valid values are `log`, `emit-metric`, and `disable`.
        """
        return pulumi.get(self, "breach_action")

    @breach_action.setter
    def breach_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "breach_action", value)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of the cluster that you want to limit usage.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_identifier", value)

    @property
    @pulumi.getter(name="featureType")
    def feature_type(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Redshift feature that you want to limit. Valid values are `spectrum`, `concurrency-scaling`, and `cross-region-datasharing`.
        """
        return pulumi.get(self, "feature_type")

    @feature_type.setter
    def feature_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_type", value)

    @property
    @pulumi.getter(name="limitType")
    def limit_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is `spectrum`, then LimitType must be `data-scanned`. If FeatureType is `concurrency-scaling`, then LimitType must be `time`. If FeatureType is `cross-region-datasharing`, then LimitType must be `data-scanned`. Valid values are `data-scanned`, and `time`.
        """
        return pulumi.get(self, "limit_type")

    @limit_type.setter
    def limit_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_type", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[str]]:
        """
        The time period that the amount applies to. A weekly period begins on Sunday. The default is `monthly`. Valid values are `daily`, `weekly`, and `monthly`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class UsageLimit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount: Optional[pulumi.Input[int]] = None,
                 breach_action: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 feature_type: Optional[pulumi.Input[str]] = None,
                 limit_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates a new Amazon Redshift Usage Limit.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.redshift.UsageLimit("example",
            cluster_identifier=aws_redshift_cluster["example"]["id"],
            feature_type="concurrency-scaling",
            limit_type="time",
            amount=60)
        ```

        ## Import

        Redshift usage limits can be imported using the `id`, e.g.,

        ```sh
         $ pulumi import aws:redshift/usageLimit:UsageLimit example example-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] amount: The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
        :param pulumi.Input[str] breach_action: The action that Amazon Redshift takes when the limit is reached. The default is `log`. Valid values are `log`, `emit-metric`, and `disable`.
        :param pulumi.Input[str] cluster_identifier: The identifier of the cluster that you want to limit usage.
        :param pulumi.Input[str] feature_type: The Amazon Redshift feature that you want to limit. Valid values are `spectrum`, `concurrency-scaling`, and `cross-region-datasharing`.
        :param pulumi.Input[str] limit_type: The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is `spectrum`, then LimitType must be `data-scanned`. If FeatureType is `concurrency-scaling`, then LimitType must be `time`. If FeatureType is `cross-region-datasharing`, then LimitType must be `data-scanned`. Valid values are `data-scanned`, and `time`.
        :param pulumi.Input[str] period: The time period that the amount applies to. A weekly period begins on Sunday. The default is `monthly`. Valid values are `daily`, `weekly`, and `monthly`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UsageLimitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Amazon Redshift Usage Limit.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.redshift.UsageLimit("example",
            cluster_identifier=aws_redshift_cluster["example"]["id"],
            feature_type="concurrency-scaling",
            limit_type="time",
            amount=60)
        ```

        ## Import

        Redshift usage limits can be imported using the `id`, e.g.,

        ```sh
         $ pulumi import aws:redshift/usageLimit:UsageLimit example example-id
        ```

        :param str resource_name: The name of the resource.
        :param UsageLimitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UsageLimitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount: Optional[pulumi.Input[int]] = None,
                 breach_action: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 feature_type: Optional[pulumi.Input[str]] = None,
                 limit_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = UsageLimitArgs.__new__(UsageLimitArgs)

            if amount is None and not opts.urn:
                raise TypeError("Missing required property 'amount'")
            __props__.__dict__["amount"] = amount
            __props__.__dict__["breach_action"] = breach_action
            if cluster_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_identifier'")
            __props__.__dict__["cluster_identifier"] = cluster_identifier
            if feature_type is None and not opts.urn:
                raise TypeError("Missing required property 'feature_type'")
            __props__.__dict__["feature_type"] = feature_type
            if limit_type is None and not opts.urn:
                raise TypeError("Missing required property 'limit_type'")
            __props__.__dict__["limit_type"] = limit_type
            __props__.__dict__["period"] = period
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["arn"] = None
        super(UsageLimit, __self__).__init__(
            'aws:redshift/usageLimit:UsageLimit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            amount: Optional[pulumi.Input[int]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            breach_action: Optional[pulumi.Input[str]] = None,
            cluster_identifier: Optional[pulumi.Input[str]] = None,
            feature_type: Optional[pulumi.Input[str]] = None,
            limit_type: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'UsageLimit':
        """
        Get an existing UsageLimit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] amount: The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Redshift Usage Limit.
        :param pulumi.Input[str] breach_action: The action that Amazon Redshift takes when the limit is reached. The default is `log`. Valid values are `log`, `emit-metric`, and `disable`.
        :param pulumi.Input[str] cluster_identifier: The identifier of the cluster that you want to limit usage.
        :param pulumi.Input[str] feature_type: The Amazon Redshift feature that you want to limit. Valid values are `spectrum`, `concurrency-scaling`, and `cross-region-datasharing`.
        :param pulumi.Input[str] limit_type: The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is `spectrum`, then LimitType must be `data-scanned`. If FeatureType is `concurrency-scaling`, then LimitType must be `time`. If FeatureType is `cross-region-datasharing`, then LimitType must be `data-scanned`. Valid values are `data-scanned`, and `time`.
        :param pulumi.Input[str] period: The time period that the amount applies to. A weekly period begins on Sunday. The default is `monthly`. Valid values are `daily`, `weekly`, and `monthly`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UsageLimitState.__new__(_UsageLimitState)

        __props__.__dict__["amount"] = amount
        __props__.__dict__["arn"] = arn
        __props__.__dict__["breach_action"] = breach_action
        __props__.__dict__["cluster_identifier"] = cluster_identifier
        __props__.__dict__["feature_type"] = feature_type
        __props__.__dict__["limit_type"] = limit_type
        __props__.__dict__["period"] = period
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return UsageLimit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Output[int]:
        """
        The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
        """
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Redshift Usage Limit.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="breachAction")
    def breach_action(self) -> pulumi.Output[Optional[str]]:
        """
        The action that Amazon Redshift takes when the limit is reached. The default is `log`. Valid values are `log`, `emit-metric`, and `disable`.
        """
        return pulumi.get(self, "breach_action")

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Output[str]:
        """
        The identifier of the cluster that you want to limit usage.
        """
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter(name="featureType")
    def feature_type(self) -> pulumi.Output[str]:
        """
        The Amazon Redshift feature that you want to limit. Valid values are `spectrum`, `concurrency-scaling`, and `cross-region-datasharing`.
        """
        return pulumi.get(self, "feature_type")

    @property
    @pulumi.getter(name="limitType")
    def limit_type(self) -> pulumi.Output[str]:
        """
        The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is `spectrum`, then LimitType must be `data-scanned`. If FeatureType is `concurrency-scaling`, then LimitType must be `time`. If FeatureType is `cross-region-datasharing`, then LimitType must be `data-scanned`. Valid values are `data-scanned`, and `time`.
        """
        return pulumi.get(self, "limit_type")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[str]]:
        """
        The time period that the amount applies to. A weekly period begins on Sunday. The default is `monthly`. Valid values are `daily`, `weekly`, and `monthly`.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        return pulumi.get(self, "tags_all")

