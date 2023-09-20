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

__all__ = ['EventDataStoreArgs', 'EventDataStore']

@pulumi.input_type
class EventDataStoreArgs:
    def __init__(__self__, *,
                 advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 multi_region_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_enabled: Optional[pulumi.Input[bool]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 termination_protection_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a EventDataStore resource.
        :param pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]] advanced_event_selectors: The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
        :param pulumi.Input[str] kms_key_id: Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        :param pulumi.Input[bool] multi_region_enabled: Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
        :param pulumi.Input[str] name: The name of the event data store.
        :param pulumi.Input[bool] organization_enabled: Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
        :param pulumi.Input[int] retention_period: The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[bool] termination_protection_enabled: Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
        """
        if advanced_event_selectors is not None:
            pulumi.set(__self__, "advanced_event_selectors", advanced_event_selectors)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if multi_region_enabled is not None:
            pulumi.set(__self__, "multi_region_enabled", multi_region_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_enabled is not None:
            pulumi.set(__self__, "organization_enabled", organization_enabled)
        if retention_period is not None:
            pulumi.set(__self__, "retention_period", retention_period)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if termination_protection_enabled is not None:
            pulumi.set(__self__, "termination_protection_enabled", termination_protection_enabled)

    @property
    @pulumi.getter(name="advancedEventSelectors")
    def advanced_event_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]]]:
        """
        The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
        """
        return pulumi.get(self, "advanced_event_selectors")

    @advanced_event_selectors.setter
    def advanced_event_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]]]):
        pulumi.set(self, "advanced_event_selectors", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="multiRegionEnabled")
    def multi_region_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
        """
        return pulumi.get(self, "multi_region_enabled")

    @multi_region_enabled.setter
    def multi_region_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_region_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the event data store.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationEnabled")
    def organization_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
        """
        return pulumi.get(self, "organization_enabled")

    @organization_enabled.setter
    def organization_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "organization_enabled", value)

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
        """
        return pulumi.get(self, "retention_period")

    @retention_period.setter
    def retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="terminationProtectionEnabled")
    def termination_protection_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
        """
        return pulumi.get(self, "termination_protection_enabled")

    @termination_protection_enabled.setter
    def termination_protection_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "termination_protection_enabled", value)


@pulumi.input_type
class _EventDataStoreState:
    def __init__(__self__, *,
                 advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 multi_region_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_enabled: Optional[pulumi.Input[bool]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 termination_protection_enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering EventDataStore resources.
        :param pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]] advanced_event_selectors: The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
        :param pulumi.Input[str] arn: ARN of the event data store.
        :param pulumi.Input[str] kms_key_id: Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        :param pulumi.Input[bool] multi_region_enabled: Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
        :param pulumi.Input[str] name: The name of the event data store.
        :param pulumi.Input[bool] organization_enabled: Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
        :param pulumi.Input[int] retention_period: The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[bool] termination_protection_enabled: Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
        """
        if advanced_event_selectors is not None:
            pulumi.set(__self__, "advanced_event_selectors", advanced_event_selectors)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if multi_region_enabled is not None:
            pulumi.set(__self__, "multi_region_enabled", multi_region_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_enabled is not None:
            pulumi.set(__self__, "organization_enabled", organization_enabled)
        if retention_period is not None:
            pulumi.set(__self__, "retention_period", retention_period)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if termination_protection_enabled is not None:
            pulumi.set(__self__, "termination_protection_enabled", termination_protection_enabled)

    @property
    @pulumi.getter(name="advancedEventSelectors")
    def advanced_event_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]]]:
        """
        The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
        """
        return pulumi.get(self, "advanced_event_selectors")

    @advanced_event_selectors.setter
    def advanced_event_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventDataStoreAdvancedEventSelectorArgs']]]]):
        pulumi.set(self, "advanced_event_selectors", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the event data store.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="multiRegionEnabled")
    def multi_region_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
        """
        return pulumi.get(self, "multi_region_enabled")

    @multi_region_enabled.setter
    def multi_region_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_region_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the event data store.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationEnabled")
    def organization_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
        """
        return pulumi.get(self, "organization_enabled")

    @organization_enabled.setter
    def organization_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "organization_enabled", value)

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
        """
        return pulumi.get(self, "retention_period")

    @retention_period.setter
    def retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="terminationProtectionEnabled")
    def termination_protection_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
        """
        return pulumi.get(self, "termination_protection_enabled")

    @termination_protection_enabled.setter
    def termination_protection_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "termination_protection_enabled", value)


class EventDataStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventDataStoreAdvancedEventSelectorArgs']]]]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 multi_region_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_enabled: Optional[pulumi.Input[bool]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 termination_protection_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a CloudTrail Event Data Store.

        More information about event data stores can be found in the [Event Data Store User Guide](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html).

        > **Tip:** For an organization event data store you must create this resource in the management account.

        ## Example Usage
        ### Basic

        The most simple event data store configuration requires us to only set the `name` attribute. The event data store will automatically capture all management events. To capture management events from all the regions, `multi_region_enabled` must be `true`.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudtrail.EventDataStore("example")
        ```
        ### Data Event Logging

        CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:

        - [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html)
        ### Log all DynamoDB PutEvent actions for a specific DynamoDB table

        ```python
        import pulumi
        import pulumi_aws as aws

        table = aws.dynamodb.get_table(name="not-important-dynamodb-table")
        # ... other configuration ...
        example = aws.cloudtrail.EventDataStore("example", advanced_event_selectors=[aws.cloudtrail.EventDataStoreAdvancedEventSelectorArgs(
            name="Log all DynamoDB PutEvent actions for a specific DynamoDB table",
            field_selectors=[
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="eventCategory",
                    equals=["Data"],
                ),
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="resources.type",
                    equals=["AWS::DynamoDB::Table"],
                ),
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="eventName",
                    equals=["PutItem"],
                ),
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="resources.ARN",
                    equals=[table.arn],
                ),
            ],
        )])
        ```

        ## Import

        Using `pulumi import`, import event data stores using their `arn`. For example:

        ```sh
         $ pulumi import aws:cloudtrail/eventDataStore:EventDataStore example arn:aws:cloudtrail:us-east-1:123456789123:eventdatastore/22333815-4414-412c-b155-dd254033gfhf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventDataStoreAdvancedEventSelectorArgs']]]] advanced_event_selectors: The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
        :param pulumi.Input[str] kms_key_id: Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        :param pulumi.Input[bool] multi_region_enabled: Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
        :param pulumi.Input[str] name: The name of the event data store.
        :param pulumi.Input[bool] organization_enabled: Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
        :param pulumi.Input[int] retention_period: The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[bool] termination_protection_enabled: Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EventDataStoreArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CloudTrail Event Data Store.

        More information about event data stores can be found in the [Event Data Store User Guide](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html).

        > **Tip:** For an organization event data store you must create this resource in the management account.

        ## Example Usage
        ### Basic

        The most simple event data store configuration requires us to only set the `name` attribute. The event data store will automatically capture all management events. To capture management events from all the regions, `multi_region_enabled` must be `true`.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudtrail.EventDataStore("example")
        ```
        ### Data Event Logging

        CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:

        - [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html)
        ### Log all DynamoDB PutEvent actions for a specific DynamoDB table

        ```python
        import pulumi
        import pulumi_aws as aws

        table = aws.dynamodb.get_table(name="not-important-dynamodb-table")
        # ... other configuration ...
        example = aws.cloudtrail.EventDataStore("example", advanced_event_selectors=[aws.cloudtrail.EventDataStoreAdvancedEventSelectorArgs(
            name="Log all DynamoDB PutEvent actions for a specific DynamoDB table",
            field_selectors=[
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="eventCategory",
                    equals=["Data"],
                ),
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="resources.type",
                    equals=["AWS::DynamoDB::Table"],
                ),
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="eventName",
                    equals=["PutItem"],
                ),
                aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
                    field="resources.ARN",
                    equals=[table.arn],
                ),
            ],
        )])
        ```

        ## Import

        Using `pulumi import`, import event data stores using their `arn`. For example:

        ```sh
         $ pulumi import aws:cloudtrail/eventDataStore:EventDataStore example arn:aws:cloudtrail:us-east-1:123456789123:eventdatastore/22333815-4414-412c-b155-dd254033gfhf
        ```

        :param str resource_name: The name of the resource.
        :param EventDataStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventDataStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventDataStoreAdvancedEventSelectorArgs']]]]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 multi_region_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_enabled: Optional[pulumi.Input[bool]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 termination_protection_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventDataStoreArgs.__new__(EventDataStoreArgs)

            __props__.__dict__["advanced_event_selectors"] = advanced_event_selectors
            __props__.__dict__["kms_key_id"] = kms_key_id
            __props__.__dict__["multi_region_enabled"] = multi_region_enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["organization_enabled"] = organization_enabled
            __props__.__dict__["retention_period"] = retention_period
            __props__.__dict__["tags"] = tags
            __props__.__dict__["termination_protection_enabled"] = termination_protection_enabled
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(EventDataStore, __self__).__init__(
            'aws:cloudtrail/eventDataStore:EventDataStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventDataStoreAdvancedEventSelectorArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            multi_region_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_enabled: Optional[pulumi.Input[bool]] = None,
            retention_period: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            termination_protection_enabled: Optional[pulumi.Input[bool]] = None) -> 'EventDataStore':
        """
        Get an existing EventDataStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventDataStoreAdvancedEventSelectorArgs']]]] advanced_event_selectors: The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
        :param pulumi.Input[str] arn: ARN of the event data store.
        :param pulumi.Input[str] kms_key_id: Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        :param pulumi.Input[bool] multi_region_enabled: Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
        :param pulumi.Input[str] name: The name of the event data store.
        :param pulumi.Input[bool] organization_enabled: Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
        :param pulumi.Input[int] retention_period: The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[bool] termination_protection_enabled: Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventDataStoreState.__new__(_EventDataStoreState)

        __props__.__dict__["advanced_event_selectors"] = advanced_event_selectors
        __props__.__dict__["arn"] = arn
        __props__.__dict__["kms_key_id"] = kms_key_id
        __props__.__dict__["multi_region_enabled"] = multi_region_enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_enabled"] = organization_enabled
        __props__.__dict__["retention_period"] = retention_period
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["termination_protection_enabled"] = termination_protection_enabled
        return EventDataStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="advancedEventSelectors")
    def advanced_event_selectors(self) -> pulumi.Output[Sequence['outputs.EventDataStoreAdvancedEventSelector']]:
        """
        The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
        """
        return pulumi.get(self, "advanced_event_selectors")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the event data store.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="multiRegionEnabled")
    def multi_region_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
        """
        return pulumi.get(self, "multi_region_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the event data store.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationEnabled")
    def organization_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
        """
        return pulumi.get(self, "organization_enabled")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> pulumi.Output[Optional[int]]:
        """
        The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
        """
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="terminationProtectionEnabled")
    def termination_protection_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
        """
        return pulumi.get(self, "termination_protection_enabled")

