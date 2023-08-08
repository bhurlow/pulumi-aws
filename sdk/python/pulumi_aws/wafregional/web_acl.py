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

__all__ = ['WebAclArgs', 'WebAcl']

@pulumi.input_type
class WebAclArgs:
    def __init__(__self__, *,
                 default_action: pulumi.Input['WebAclDefaultActionArgs'],
                 metric_name: pulumi.Input[str],
                 logging_configuration: Optional[pulumi.Input['WebAclLoggingConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a WebAcl resource.
        :param pulumi.Input['WebAclDefaultActionArgs'] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input['WebAclLoggingConfigurationArgs'] logging_configuration: Configuration block to enable WAF logging. Detailed below.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]] rules: Set of configuration blocks containing rules for the web ACL. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "default_action", default_action)
        pulumi.set(__self__, "metric_name", metric_name)
        if logging_configuration is not None:
            pulumi.set(__self__, "logging_configuration", logging_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Input['WebAclDefaultActionArgs']:
        """
        The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: pulumi.Input['WebAclDefaultActionArgs']):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Input[str]:
        """
        The name or description for the Amazon CloudWatch metric of this web ACL.
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> Optional[pulumi.Input['WebAclLoggingConfigurationArgs']]:
        """
        Configuration block to enable WAF logging. Detailed below.
        """
        return pulumi.get(self, "logging_configuration")

    @logging_configuration.setter
    def logging_configuration(self, value: Optional[pulumi.Input['WebAclLoggingConfigurationArgs']]):
        pulumi.set(self, "logging_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the web ACL.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]]:
        """
        Set of configuration blocks containing rules for the web ACL. Detailed below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]]):
        pulumi.set(self, "rules", value)

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
class _WebAclState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input['WebAclDefaultActionArgs']] = None,
                 logging_configuration: Optional[pulumi.Input['WebAclLoggingConfigurationArgs']] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering WebAcl resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the WAF Regional WebACL.
        :param pulumi.Input['WebAclDefaultActionArgs'] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input['WebAclLoggingConfigurationArgs'] logging_configuration: Configuration block to enable WAF logging. Detailed below.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]] rules: Set of configuration blocks containing rules for the web ACL. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if logging_configuration is not None:
            pulumi.set(__self__, "logging_configuration", logging_configuration)
        if metric_name is not None:
            pulumi.set(__self__, "metric_name", metric_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the WAF Regional WebACL.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input['WebAclDefaultActionArgs']]:
        """
        The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input['WebAclDefaultActionArgs']]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> Optional[pulumi.Input['WebAclLoggingConfigurationArgs']]:
        """
        Configuration block to enable WAF logging. Detailed below.
        """
        return pulumi.get(self, "logging_configuration")

    @logging_configuration.setter
    def logging_configuration(self, value: Optional[pulumi.Input['WebAclLoggingConfigurationArgs']]):
        pulumi.set(self, "logging_configuration", value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description for the Amazon CloudWatch metric of this web ACL.
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the web ACL.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]]:
        """
        Set of configuration blocks containing rules for the web ACL. Detailed below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]]):
        pulumi.set(self, "rules", value)

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
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class WebAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']]] = None,
                 logging_configuration: Optional[pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationArgs']]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.

        ## Example Usage
        ### Regular Rule

        ```python
        import pulumi
        import pulumi_aws as aws

        ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[aws.wafregional.IpSetIpSetDescriptorArgs(
            type="IPV4",
            value="192.0.7.0/24",
        )])
        wafrule = aws.wafregional.Rule("wafrule",
            metric_name="tfWAFRule",
            predicates=[aws.wafregional.RulePredicateArgs(
                data_id=ipset.id,
                negated=False,
                type="IPMatch",
            )])
        wafacl = aws.wafregional.WebAcl("wafacl",
            metric_name="tfWebACL",
            default_action=aws.wafregional.WebAclDefaultActionArgs(
                type="ALLOW",
            ),
            rules=[aws.wafregional.WebAclRuleArgs(
                action=aws.wafregional.WebAclRuleActionArgs(
                    type="BLOCK",
                ),
                priority=1,
                rule_id=wafrule.id,
                type="REGULAR",
            )])
        ```
        ### Group Rule

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafregional.WebAcl("example",
            metric_name="example",
            default_action=aws.wafregional.WebAclDefaultActionArgs(
                type="ALLOW",
            ),
            rules=[aws.wafregional.WebAclRuleArgs(
                priority=1,
                rule_id=aws_wafregional_rule_group["example"]["id"],
                type="GROUP",
                override_action=aws.wafregional.WebAclRuleOverrideActionArgs(
                    type="NONE",
                ),
            )])
        ```
        ### Logging

        > *NOTE:* The Kinesis Firehose Delivery Stream name must begin with `aws-waf-logs-`. See the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) for more information about enabling WAF logging.

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configuration ...
        example = aws.wafregional.WebAcl("example", logging_configuration=aws.wafregional.WebAclLoggingConfigurationArgs(
            log_destination=aws_kinesis_firehose_delivery_stream["example"]["arn"],
            redacted_fields=aws.wafregional.WebAclLoggingConfigurationRedactedFieldsArgs(
                field_to_matches=[
                    aws.wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
                        type="URI",
                    ),
                    aws.wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
                        data="referer",
                        type="HEADER",
                    ),
                ],
            ),
        ))
        ```

        ## Import

        terraform import {

         to = aws_wafregional_web_acl.wafacl

         id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import WAF Regional Web ACL using the id. For exampleconsole % pulumi import aws_wafregional_web_acl.wafacl a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationArgs']] logging_configuration: Configuration block to enable WAF logging. Detailed below.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]] rules: Set of configuration blocks containing rules for the web ACL. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.

        ## Example Usage
        ### Regular Rule

        ```python
        import pulumi
        import pulumi_aws as aws

        ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[aws.wafregional.IpSetIpSetDescriptorArgs(
            type="IPV4",
            value="192.0.7.0/24",
        )])
        wafrule = aws.wafregional.Rule("wafrule",
            metric_name="tfWAFRule",
            predicates=[aws.wafregional.RulePredicateArgs(
                data_id=ipset.id,
                negated=False,
                type="IPMatch",
            )])
        wafacl = aws.wafregional.WebAcl("wafacl",
            metric_name="tfWebACL",
            default_action=aws.wafregional.WebAclDefaultActionArgs(
                type="ALLOW",
            ),
            rules=[aws.wafregional.WebAclRuleArgs(
                action=aws.wafregional.WebAclRuleActionArgs(
                    type="BLOCK",
                ),
                priority=1,
                rule_id=wafrule.id,
                type="REGULAR",
            )])
        ```
        ### Group Rule

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafregional.WebAcl("example",
            metric_name="example",
            default_action=aws.wafregional.WebAclDefaultActionArgs(
                type="ALLOW",
            ),
            rules=[aws.wafregional.WebAclRuleArgs(
                priority=1,
                rule_id=aws_wafregional_rule_group["example"]["id"],
                type="GROUP",
                override_action=aws.wafregional.WebAclRuleOverrideActionArgs(
                    type="NONE",
                ),
            )])
        ```
        ### Logging

        > *NOTE:* The Kinesis Firehose Delivery Stream name must begin with `aws-waf-logs-`. See the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) for more information about enabling WAF logging.

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configuration ...
        example = aws.wafregional.WebAcl("example", logging_configuration=aws.wafregional.WebAclLoggingConfigurationArgs(
            log_destination=aws_kinesis_firehose_delivery_stream["example"]["arn"],
            redacted_fields=aws.wafregional.WebAclLoggingConfigurationRedactedFieldsArgs(
                field_to_matches=[
                    aws.wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
                        type="URI",
                    ),
                    aws.wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
                        data="referer",
                        type="HEADER",
                    ),
                ],
            ),
        ))
        ```

        ## Import

        terraform import {

         to = aws_wafregional_web_acl.wafacl

         id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import WAF Regional Web ACL using the id. For exampleconsole % pulumi import aws_wafregional_web_acl.wafacl a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc

        :param str resource_name: The name of the resource.
        :param WebAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']]] = None,
                 logging_configuration: Optional[pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationArgs']]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebAclArgs.__new__(WebAclArgs)

            if default_action is None and not opts.urn:
                raise TypeError("Missing required property 'default_action'")
            __props__.__dict__["default_action"] = default_action
            __props__.__dict__["logging_configuration"] = logging_configuration
            if metric_name is None and not opts.urn:
                raise TypeError("Missing required property 'metric_name'")
            __props__.__dict__["metric_name"] = metric_name
            __props__.__dict__["name"] = name
            __props__.__dict__["rules"] = rules
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(WebAcl, __self__).__init__(
            'aws:wafregional/webAcl:WebAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            default_action: Optional[pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']]] = None,
            logging_configuration: Optional[pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationArgs']]] = None,
            metric_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'WebAcl':
        """
        Get an existing WebAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the WAF Regional WebACL.
        :param pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationArgs']] logging_configuration: Configuration block to enable WAF logging. Detailed below.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]] rules: Set of configuration blocks containing rules for the web ACL. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebAclState.__new__(_WebAclState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["default_action"] = default_action
        __props__.__dict__["logging_configuration"] = logging_configuration
        __props__.__dict__["metric_name"] = metric_name
        __props__.__dict__["name"] = name
        __props__.__dict__["rules"] = rules
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return WebAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the WAF Regional WebACL.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output['outputs.WebAclDefaultAction']:
        """
        The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> pulumi.Output[Optional['outputs.WebAclLoggingConfiguration']]:
        """
        Configuration block to enable WAF logging. Detailed below.
        """
        return pulumi.get(self, "logging_configuration")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Output[str]:
        """
        The name or description for the Amazon CloudWatch metric of this web ACL.
        """
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or description of the web ACL.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.WebAclRule']]]:
        """
        Set of configuration blocks containing rules for the web ACL. Detailed below.
        """
        return pulumi.get(self, "rules")

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
        return pulumi.get(self, "tags_all")

