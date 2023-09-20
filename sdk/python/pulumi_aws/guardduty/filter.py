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

__all__ = ['FilterArgs', 'Filter']

@pulumi.input_type
class FilterArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 detector_id: pulumi.Input[str],
                 finding_criteria: pulumi.Input['FilterFindingCriteriaArgs'],
                 rank: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Filter resource.
        :param pulumi.Input[str] action: Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        :param pulumi.Input[str] detector_id: ID of a GuardDuty detector, attached to your account.
        :param pulumi.Input['FilterFindingCriteriaArgs'] finding_criteria: Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        :param pulumi.Input[int] rank: Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        :param pulumi.Input[str] description: Description of the filter.
        :param pulumi.Input[str] name: The name of your filter.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "detector_id", detector_id)
        pulumi.set(__self__, "finding_criteria", finding_criteria)
        pulumi.set(__self__, "rank", rank)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Input[str]:
        """
        ID of a GuardDuty detector, attached to your account.
        """
        return pulumi.get(self, "detector_id")

    @detector_id.setter
    def detector_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "detector_id", value)

    @property
    @pulumi.getter(name="findingCriteria")
    def finding_criteria(self) -> pulumi.Input['FilterFindingCriteriaArgs']:
        """
        Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        """
        return pulumi.get(self, "finding_criteria")

    @finding_criteria.setter
    def finding_criteria(self, value: pulumi.Input['FilterFindingCriteriaArgs']):
        pulumi.set(self, "finding_criteria", value)

    @property
    @pulumi.getter
    def rank(self) -> pulumi.Input[int]:
        """
        Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        """
        return pulumi.get(self, "rank")

    @rank.setter
    def rank(self, value: pulumi.Input[int]):
        pulumi.set(self, "rank", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the filter.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of your filter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _FilterState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 finding_criteria: Optional[pulumi.Input['FilterFindingCriteriaArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rank: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Filter resources.
        :param pulumi.Input[str] action: Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        :param pulumi.Input[str] arn: The ARN of the GuardDuty filter.
        :param pulumi.Input[str] description: Description of the filter.
        :param pulumi.Input[str] detector_id: ID of a GuardDuty detector, attached to your account.
        :param pulumi.Input['FilterFindingCriteriaArgs'] finding_criteria: Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        :param pulumi.Input[str] name: The name of your filter.
        :param pulumi.Input[int] rank: Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if detector_id is not None:
            pulumi.set(__self__, "detector_id", detector_id)
        if finding_criteria is not None:
            pulumi.set(__self__, "finding_criteria", finding_criteria)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rank is not None:
            pulumi.set(__self__, "rank", rank)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the GuardDuty filter.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the filter.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a GuardDuty detector, attached to your account.
        """
        return pulumi.get(self, "detector_id")

    @detector_id.setter
    def detector_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "detector_id", value)

    @property
    @pulumi.getter(name="findingCriteria")
    def finding_criteria(self) -> Optional[pulumi.Input['FilterFindingCriteriaArgs']]:
        """
        Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        """
        return pulumi.get(self, "finding_criteria")

    @finding_criteria.setter
    def finding_criteria(self, value: Optional[pulumi.Input['FilterFindingCriteriaArgs']]):
        pulumi.set(self, "finding_criteria", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of your filter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rank(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        """
        return pulumi.get(self, "rank")

    @rank.setter
    def rank(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rank", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class Filter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 finding_criteria: Optional[pulumi.Input[pulumi.InputType['FilterFindingCriteriaArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rank: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage a GuardDuty filter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_filter = aws.guardduty.Filter("myFilter",
            action="ARCHIVE",
            detector_id=aws_guardduty_detector["example"]["id"],
            rank=1,
            finding_criteria=aws.guardduty.FilterFindingCriteriaArgs(
                criterions=[
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="region",
                        equals=["eu-west-1"],
                    ),
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="service.additionalInfo.threatListName",
                        not_equals=[
                            "some-threat",
                            "another-threat",
                        ],
                    ),
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="updatedAt",
                        greater_than="2020-01-01T00:00:00Z",
                        less_than="2020-02-01T00:00:00Z",
                    ),
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="severity",
                        greater_than_or_equal="4",
                    ),
                ],
            ))
        ```

        ## Import

        Using `pulumi import`, import GuardDuty filters using the detector ID and filter's name separated by a colon. For example:

        ```sh
         $ pulumi import aws:guardduty/filter:Filter MyFilter 00b00fd5aecc0ab60a708659477e9617:MyFilter
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        :param pulumi.Input[str] description: Description of the filter.
        :param pulumi.Input[str] detector_id: ID of a GuardDuty detector, attached to your account.
        :param pulumi.Input[pulumi.InputType['FilterFindingCriteriaArgs']] finding_criteria: Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        :param pulumi.Input[str] name: The name of your filter.
        :param pulumi.Input[int] rank: Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FilterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a GuardDuty filter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_filter = aws.guardduty.Filter("myFilter",
            action="ARCHIVE",
            detector_id=aws_guardduty_detector["example"]["id"],
            rank=1,
            finding_criteria=aws.guardduty.FilterFindingCriteriaArgs(
                criterions=[
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="region",
                        equals=["eu-west-1"],
                    ),
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="service.additionalInfo.threatListName",
                        not_equals=[
                            "some-threat",
                            "another-threat",
                        ],
                    ),
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="updatedAt",
                        greater_than="2020-01-01T00:00:00Z",
                        less_than="2020-02-01T00:00:00Z",
                    ),
                    aws.guardduty.FilterFindingCriteriaCriterionArgs(
                        field="severity",
                        greater_than_or_equal="4",
                    ),
                ],
            ))
        ```

        ## Import

        Using `pulumi import`, import GuardDuty filters using the detector ID and filter's name separated by a colon. For example:

        ```sh
         $ pulumi import aws:guardduty/filter:Filter MyFilter 00b00fd5aecc0ab60a708659477e9617:MyFilter
        ```

        :param str resource_name: The name of the resource.
        :param FilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 finding_criteria: Optional[pulumi.Input[pulumi.InputType['FilterFindingCriteriaArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rank: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FilterArgs.__new__(FilterArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["description"] = description
            if detector_id is None and not opts.urn:
                raise TypeError("Missing required property 'detector_id'")
            __props__.__dict__["detector_id"] = detector_id
            if finding_criteria is None and not opts.urn:
                raise TypeError("Missing required property 'finding_criteria'")
            __props__.__dict__["finding_criteria"] = finding_criteria
            __props__.__dict__["name"] = name
            if rank is None and not opts.urn:
                raise TypeError("Missing required property 'rank'")
            __props__.__dict__["rank"] = rank
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Filter, __self__).__init__(
            'aws:guardduty/filter:Filter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            detector_id: Optional[pulumi.Input[str]] = None,
            finding_criteria: Optional[pulumi.Input[pulumi.InputType['FilterFindingCriteriaArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rank: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Filter':
        """
        Get an existing Filter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        :param pulumi.Input[str] arn: The ARN of the GuardDuty filter.
        :param pulumi.Input[str] description: Description of the filter.
        :param pulumi.Input[str] detector_id: ID of a GuardDuty detector, attached to your account.
        :param pulumi.Input[pulumi.InputType['FilterFindingCriteriaArgs']] finding_criteria: Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        :param pulumi.Input[str] name: The name of your filter.
        :param pulumi.Input[int] rank: Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FilterState.__new__(_FilterState)

        __props__.__dict__["action"] = action
        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["detector_id"] = detector_id
        __props__.__dict__["finding_criteria"] = finding_criteria
        __props__.__dict__["name"] = name
        __props__.__dict__["rank"] = rank
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Filter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the GuardDuty filter.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the filter.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Output[str]:
        """
        ID of a GuardDuty detector, attached to your account.
        """
        return pulumi.get(self, "detector_id")

    @property
    @pulumi.getter(name="findingCriteria")
    def finding_criteria(self) -> pulumi.Output['outputs.FilterFindingCriteria']:
        """
        Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        """
        return pulumi.get(self, "finding_criteria")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of your filter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rank(self) -> pulumi.Output[int]:
        """
        Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        """
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

