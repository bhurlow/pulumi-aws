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

__all__ = ['CostCategoryArgs', 'CostCategory']

@pulumi.input_type
class CostCategoryArgs:
    def __init__(__self__, *,
                 rule_version: pulumi.Input[str],
                 rules: pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]],
                 default_value: Optional[pulumi.Input[str]] = None,
                 effective_start: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 split_charge_rules: Optional[pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CostCategory resource.
        :param pulumi.Input[str] rule_version: Rule schema version in this particular Cost Category.
        :param pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]] rules: Configuration block for the Cost Category rules used to categorize costs. See below.
        :param pulumi.Input[str] default_value: Default value for the cost category.
        :param pulumi.Input[str] effective_start: The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.
               
               The following arguments are optional:
        :param pulumi.Input[str] name: Unique name for the Cost Category.
        :param pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]] split_charge_rules: Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "rule_version", rule_version)
        pulumi.set(__self__, "rules", rules)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if effective_start is not None:
            pulumi.set(__self__, "effective_start", effective_start)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if split_charge_rules is not None:
            pulumi.set(__self__, "split_charge_rules", split_charge_rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="ruleVersion")
    def rule_version(self) -> pulumi.Input[str]:
        """
        Rule schema version in this particular Cost Category.
        """
        return pulumi.get(self, "rule_version")

    @rule_version.setter
    def rule_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_version", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]]:
        """
        Configuration block for the Cost Category rules used to categorize costs. See below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[str]]:
        """
        Default value for the cost category.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter(name="effectiveStart")
    def effective_start(self) -> Optional[pulumi.Input[str]]:
        """
        The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.

        The following arguments are optional:
        """
        return pulumi.get(self, "effective_start")

    @effective_start.setter
    def effective_start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "effective_start", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name for the Cost Category.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="splitChargeRules")
    def split_charge_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]]]:
        """
        Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        """
        return pulumi.get(self, "split_charge_rules")

    @split_charge_rules.setter
    def split_charge_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]]]):
        pulumi.set(self, "split_charge_rules", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _CostCategoryState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 default_value: Optional[pulumi.Input[str]] = None,
                 effective_end: Optional[pulumi.Input[str]] = None,
                 effective_start: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rule_version: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]]] = None,
                 split_charge_rules: Optional[pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering CostCategory resources.
        :param pulumi.Input[str] arn: ARN of the cost category.
        :param pulumi.Input[str] default_value: Default value for the cost category.
        :param pulumi.Input[str] effective_end: Effective end data of your Cost Category.
        :param pulumi.Input[str] effective_start: The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.
               
               The following arguments are optional:
        :param pulumi.Input[str] name: Unique name for the Cost Category.
        :param pulumi.Input[str] rule_version: Rule schema version in this particular Cost Category.
        :param pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]] rules: Configuration block for the Cost Category rules used to categorize costs. See below.
        :param pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]] split_charge_rules: Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if effective_end is not None:
            pulumi.set(__self__, "effective_end", effective_end)
        if effective_start is not None:
            pulumi.set(__self__, "effective_start", effective_start)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rule_version is not None:
            pulumi.set(__self__, "rule_version", rule_version)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if split_charge_rules is not None:
            pulumi.set(__self__, "split_charge_rules", split_charge_rules)
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
        ARN of the cost category.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[str]]:
        """
        Default value for the cost category.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter(name="effectiveEnd")
    def effective_end(self) -> Optional[pulumi.Input[str]]:
        """
        Effective end data of your Cost Category.
        """
        return pulumi.get(self, "effective_end")

    @effective_end.setter
    def effective_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "effective_end", value)

    @property
    @pulumi.getter(name="effectiveStart")
    def effective_start(self) -> Optional[pulumi.Input[str]]:
        """
        The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.

        The following arguments are optional:
        """
        return pulumi.get(self, "effective_start")

    @effective_start.setter
    def effective_start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "effective_start", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name for the Cost Category.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ruleVersion")
    def rule_version(self) -> Optional[pulumi.Input[str]]:
        """
        Rule schema version in this particular Cost Category.
        """
        return pulumi.get(self, "rule_version")

    @rule_version.setter
    def rule_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_version", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]]]:
        """
        Configuration block for the Cost Category rules used to categorize costs. See below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CostCategoryRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="splitChargeRules")
    def split_charge_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]]]:
        """
        Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        """
        return pulumi.get(self, "split_charge_rules")

    @split_charge_rules.setter
    def split_charge_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CostCategorySplitChargeRuleArgs']]]]):
        pulumi.set(self, "split_charge_rules", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class CostCategory(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_value: Optional[pulumi.Input[str]] = None,
                 effective_start: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rule_version: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategoryRuleArgs']]]]] = None,
                 split_charge_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategorySplitChargeRuleArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a CE Cost Category.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.costexplorer.CostCategory("test",
            rules=[
                aws.costexplorer.CostCategoryRuleArgs(
                    rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                        dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                            key="LINKED_ACCOUNT_NAME",
                            match_options=["ENDS_WITH"],
                            values=["-prod"],
                        ),
                    ),
                    value="production",
                ),
                aws.costexplorer.CostCategoryRuleArgs(
                    rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                        dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                            key="LINKED_ACCOUNT_NAME",
                            match_options=["ENDS_WITH"],
                            values=["-stg"],
                        ),
                    ),
                    value="staging",
                ),
                aws.costexplorer.CostCategoryRuleArgs(
                    rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                        dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                            key="LINKED_ACCOUNT_NAME",
                            match_options=["ENDS_WITH"],
                            values=["-dev"],
                        ),
                    ),
                    value="testing",
                ),
            ],
            rule_version="CostCategoryExpression.v1")
        ```

        ## Import

        Using `pulumi import`, import `aws_ce_cost_category` using the id. For example:

        ```sh
         $ pulumi import aws:costexplorer/costCategory:CostCategory example costCategoryARN
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_value: Default value for the cost category.
        :param pulumi.Input[str] effective_start: The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.
               
               The following arguments are optional:
        :param pulumi.Input[str] name: Unique name for the Cost Category.
        :param pulumi.Input[str] rule_version: Rule schema version in this particular Cost Category.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategoryRuleArgs']]]] rules: Configuration block for the Cost Category rules used to categorize costs. See below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategorySplitChargeRuleArgs']]]] split_charge_rules: Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CostCategoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CE Cost Category.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.costexplorer.CostCategory("test",
            rules=[
                aws.costexplorer.CostCategoryRuleArgs(
                    rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                        dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                            key="LINKED_ACCOUNT_NAME",
                            match_options=["ENDS_WITH"],
                            values=["-prod"],
                        ),
                    ),
                    value="production",
                ),
                aws.costexplorer.CostCategoryRuleArgs(
                    rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                        dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                            key="LINKED_ACCOUNT_NAME",
                            match_options=["ENDS_WITH"],
                            values=["-stg"],
                        ),
                    ),
                    value="staging",
                ),
                aws.costexplorer.CostCategoryRuleArgs(
                    rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                        dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                            key="LINKED_ACCOUNT_NAME",
                            match_options=["ENDS_WITH"],
                            values=["-dev"],
                        ),
                    ),
                    value="testing",
                ),
            ],
            rule_version="CostCategoryExpression.v1")
        ```

        ## Import

        Using `pulumi import`, import `aws_ce_cost_category` using the id. For example:

        ```sh
         $ pulumi import aws:costexplorer/costCategory:CostCategory example costCategoryARN
        ```

        :param str resource_name: The name of the resource.
        :param CostCategoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CostCategoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_value: Optional[pulumi.Input[str]] = None,
                 effective_start: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rule_version: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategoryRuleArgs']]]]] = None,
                 split_charge_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategorySplitChargeRuleArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CostCategoryArgs.__new__(CostCategoryArgs)

            __props__.__dict__["default_value"] = default_value
            __props__.__dict__["effective_start"] = effective_start
            __props__.__dict__["name"] = name
            if rule_version is None and not opts.urn:
                raise TypeError("Missing required property 'rule_version'")
            __props__.__dict__["rule_version"] = rule_version
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            __props__.__dict__["split_charge_rules"] = split_charge_rules
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["effective_end"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(CostCategory, __self__).__init__(
            'aws:costexplorer/costCategory:CostCategory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            default_value: Optional[pulumi.Input[str]] = None,
            effective_end: Optional[pulumi.Input[str]] = None,
            effective_start: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rule_version: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategoryRuleArgs']]]]] = None,
            split_charge_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategorySplitChargeRuleArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'CostCategory':
        """
        Get an existing CostCategory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the cost category.
        :param pulumi.Input[str] default_value: Default value for the cost category.
        :param pulumi.Input[str] effective_end: Effective end data of your Cost Category.
        :param pulumi.Input[str] effective_start: The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.
               
               The following arguments are optional:
        :param pulumi.Input[str] name: Unique name for the Cost Category.
        :param pulumi.Input[str] rule_version: Rule schema version in this particular Cost Category.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategoryRuleArgs']]]] rules: Configuration block for the Cost Category rules used to categorize costs. See below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CostCategorySplitChargeRuleArgs']]]] split_charge_rules: Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CostCategoryState.__new__(_CostCategoryState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["default_value"] = default_value
        __props__.__dict__["effective_end"] = effective_end
        __props__.__dict__["effective_start"] = effective_start
        __props__.__dict__["name"] = name
        __props__.__dict__["rule_version"] = rule_version
        __props__.__dict__["rules"] = rules
        __props__.__dict__["split_charge_rules"] = split_charge_rules
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return CostCategory(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the cost category.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> pulumi.Output[Optional[str]]:
        """
        Default value for the cost category.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter(name="effectiveEnd")
    def effective_end(self) -> pulumi.Output[str]:
        """
        Effective end data of your Cost Category.
        """
        return pulumi.get(self, "effective_end")

    @property
    @pulumi.getter(name="effectiveStart")
    def effective_start(self) -> pulumi.Output[str]:
        """
        The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.

        The following arguments are optional:
        """
        return pulumi.get(self, "effective_start")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name for the Cost Category.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ruleVersion")
    def rule_version(self) -> pulumi.Output[str]:
        """
        Rule schema version in this particular Cost Category.
        """
        return pulumi.get(self, "rule_version")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.CostCategoryRule']]:
        """
        Configuration block for the Cost Category rules used to categorize costs. See below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="splitChargeRules")
    def split_charge_rules(self) -> pulumi.Output[Optional[Sequence['outputs.CostCategorySplitChargeRule']]]:
        """
        Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        """
        return pulumi.get(self, "split_charge_rules")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

