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

__all__ = ['RateBasedRuleArgs', 'RateBasedRule']

@pulumi.input_type
class RateBasedRuleArgs:
    def __init__(__self__, *,
                 metric_name: pulumi.Input[str],
                 rate_key: pulumi.Input[str],
                 rate_limit: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None,
                 predicates: Optional[pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RateBasedRule resource.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this rule.
        :param pulumi.Input[str] rate_key: Valid value is IP.
        :param pulumi.Input[int] rate_limit: The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        :param pulumi.Input[str] name: The name or description of the rule.
        :param pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]] predicates: The objects to include in a rule (documented below).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "metric_name", metric_name)
        pulumi.set(__self__, "rate_key", rate_key)
        pulumi.set(__self__, "rate_limit", rate_limit)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if predicates is not None:
            pulumi.set(__self__, "predicates", predicates)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Input[str]:
        """
        The name or description for the Amazon CloudWatch metric of this rule.
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter(name="rateKey")
    def rate_key(self) -> pulumi.Input[str]:
        """
        Valid value is IP.
        """
        return pulumi.get(self, "rate_key")

    @rate_key.setter
    def rate_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "rate_key", value)

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> pulumi.Input[int]:
        """
        The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        """
        return pulumi.get(self, "rate_limit")

    @rate_limit.setter
    def rate_limit(self, value: pulumi.Input[int]):
        pulumi.set(self, "rate_limit", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def predicates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]]]:
        """
        The objects to include in a rule (documented below).
        """
        return pulumi.get(self, "predicates")

    @predicates.setter
    def predicates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]]]):
        pulumi.set(self, "predicates", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RateBasedRuleState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 predicates: Optional[pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]]] = None,
                 rate_key: Optional[pulumi.Input[str]] = None,
                 rate_limit: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RateBasedRule resources.
        :param pulumi.Input[str] arn: The ARN of the WAF Regional Rate Based Rule.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this rule.
        :param pulumi.Input[str] name: The name or description of the rule.
        :param pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]] predicates: The objects to include in a rule (documented below).
        :param pulumi.Input[str] rate_key: Valid value is IP.
        :param pulumi.Input[int] rate_limit: The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if metric_name is not None:
            pulumi.set(__self__, "metric_name", metric_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if predicates is not None:
            pulumi.set(__self__, "predicates", predicates)
        if rate_key is not None:
            pulumi.set(__self__, "rate_key", rate_key)
        if rate_limit is not None:
            pulumi.set(__self__, "rate_limit", rate_limit)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the WAF Regional Rate Based Rule.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description for the Amazon CloudWatch metric of this rule.
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def predicates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]]]:
        """
        The objects to include in a rule (documented below).
        """
        return pulumi.get(self, "predicates")

    @predicates.setter
    def predicates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RateBasedRulePredicateArgs']]]]):
        pulumi.set(self, "predicates", value)

    @property
    @pulumi.getter(name="rateKey")
    def rate_key(self) -> Optional[pulumi.Input[str]]:
        """
        Valid value is IP.
        """
        return pulumi.get(self, "rate_key")

    @rate_key.setter
    def rate_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rate_key", value)

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        """
        return pulumi.get(self, "rate_limit")

    @rate_limit.setter
    def rate_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate_limit", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class RateBasedRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 predicates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RateBasedRulePredicateArgs']]]]] = None,
                 rate_key: Optional[pulumi.Input[str]] = None,
                 rate_limit: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a WAF Rate Based Rule Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[aws.wafregional.IpSetIpSetDescriptorArgs(
            type="IPV4",
            value="192.0.7.0/24",
        )])
        wafrule = aws.wafregional.RateBasedRule("wafrule",
            metric_name="tfWAFRule",
            rate_key="IP",
            rate_limit=100,
            predicates=[aws.wafregional.RateBasedRulePredicateArgs(
                data_id=ipset.id,
                negated=False,
                type="IPMatch",
            )],
            opts=pulumi.ResourceOptions(depends_on=[ipset]))
        ```

        ## Import

        Using `pulumi import`, import WAF Regional Rate Based Rule using the id. For example:

        ```sh
         $ pulumi import aws:wafregional/rateBasedRule:RateBasedRule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this rule.
        :param pulumi.Input[str] name: The name or description of the rule.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RateBasedRulePredicateArgs']]]] predicates: The objects to include in a rule (documented below).
        :param pulumi.Input[str] rate_key: Valid value is IP.
        :param pulumi.Input[int] rate_limit: The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RateBasedRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Rate Based Rule Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[aws.wafregional.IpSetIpSetDescriptorArgs(
            type="IPV4",
            value="192.0.7.0/24",
        )])
        wafrule = aws.wafregional.RateBasedRule("wafrule",
            metric_name="tfWAFRule",
            rate_key="IP",
            rate_limit=100,
            predicates=[aws.wafregional.RateBasedRulePredicateArgs(
                data_id=ipset.id,
                negated=False,
                type="IPMatch",
            )],
            opts=pulumi.ResourceOptions(depends_on=[ipset]))
        ```

        ## Import

        Using `pulumi import`, import WAF Regional Rate Based Rule using the id. For example:

        ```sh
         $ pulumi import aws:wafregional/rateBasedRule:RateBasedRule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param RateBasedRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RateBasedRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 predicates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RateBasedRulePredicateArgs']]]]] = None,
                 rate_key: Optional[pulumi.Input[str]] = None,
                 rate_limit: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RateBasedRuleArgs.__new__(RateBasedRuleArgs)

            if metric_name is None and not opts.urn:
                raise TypeError("Missing required property 'metric_name'")
            __props__.__dict__["metric_name"] = metric_name
            __props__.__dict__["name"] = name
            __props__.__dict__["predicates"] = predicates
            if rate_key is None and not opts.urn:
                raise TypeError("Missing required property 'rate_key'")
            __props__.__dict__["rate_key"] = rate_key
            if rate_limit is None and not opts.urn:
                raise TypeError("Missing required property 'rate_limit'")
            __props__.__dict__["rate_limit"] = rate_limit
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(RateBasedRule, __self__).__init__(
            'aws:wafregional/rateBasedRule:RateBasedRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            metric_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            predicates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RateBasedRulePredicateArgs']]]]] = None,
            rate_key: Optional[pulumi.Input[str]] = None,
            rate_limit: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'RateBasedRule':
        """
        Get an existing RateBasedRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the WAF Regional Rate Based Rule.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this rule.
        :param pulumi.Input[str] name: The name or description of the rule.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RateBasedRulePredicateArgs']]]] predicates: The objects to include in a rule (documented below).
        :param pulumi.Input[str] rate_key: Valid value is IP.
        :param pulumi.Input[int] rate_limit: The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RateBasedRuleState.__new__(_RateBasedRuleState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["metric_name"] = metric_name
        __props__.__dict__["name"] = name
        __props__.__dict__["predicates"] = predicates
        __props__.__dict__["rate_key"] = rate_key
        __props__.__dict__["rate_limit"] = rate_limit
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return RateBasedRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the WAF Regional Rate Based Rule.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Output[str]:
        """
        The name or description for the Amazon CloudWatch metric of this rule.
        """
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or description of the rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def predicates(self) -> pulumi.Output[Optional[Sequence['outputs.RateBasedRulePredicate']]]:
        """
        The objects to include in a rule (documented below).
        """
        return pulumi.get(self, "predicates")

    @property
    @pulumi.getter(name="rateKey")
    def rate_key(self) -> pulumi.Output[str]:
        """
        Valid value is IP.
        """
        return pulumi.get(self, "rate_key")

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> pulumi.Output[int]:
        """
        The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        """
        return pulumi.get(self, "rate_limit")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

