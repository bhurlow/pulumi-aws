# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['CompositeAlarm']


class CompositeAlarm(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions_enabled: Optional[pulumi.Input[bool]] = None,
                 alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_description: Optional[pulumi.Input[str]] = None,
                 alarm_name: Optional[pulumi.Input[str]] = None,
                 alarm_rule: Optional[pulumi.Input[str]] = None,
                 insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Composite Alarm resource.

        > **NOTE:** An alarm (composite or metric) cannot be destroyed when there are other composite alarms depending on it. This can lead to a cyclical dependency on update, as the provider will unsuccessfully attempt to destroy alarms before updating the rule. Consider using `depends_on`, references to alarm names, and two-stage updates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudwatch.CompositeAlarm("example",
            alarm_description="This is a composite alarm!",
            alarm_name="example-composite-alarm",
            alarm_actions=aws_sns_topic["example"]["arn"],
            ok_actions=aws_sns_topic["example"]["arn"],
            alarm_rule=f\"\"\"ALARM({aws_cloudwatch_metric_alarm["alpha"]["alarm_name"]}) OR
        ALARM({aws_cloudwatch_metric_alarm["bravo"]["alarm_name"]})
        \"\"\")
        ```

        ## Import

        Use the `alarm_name` to import a CloudWatch Composite Alarm. For example

        ```sh
         $ pulumi import aws:cloudwatch/compositeAlarm:CompositeAlarm test my-alarm
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[str] alarm_description: The description for the composite alarm.
        :param pulumi.Input[str] alarm_name: The name for the composite alarm. This name must be unique within the region.
        :param pulumi.Input[str] alarm_rule: An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['actions_enabled'] = actions_enabled
            __props__['alarm_actions'] = alarm_actions
            __props__['alarm_description'] = alarm_description
            if alarm_name is None and not opts.urn:
                raise TypeError("Missing required property 'alarm_name'")
            __props__['alarm_name'] = alarm_name
            if alarm_rule is None and not opts.urn:
                raise TypeError("Missing required property 'alarm_rule'")
            __props__['alarm_rule'] = alarm_rule
            __props__['insufficient_data_actions'] = insufficient_data_actions
            __props__['ok_actions'] = ok_actions
            __props__['tags'] = tags
            __props__['arn'] = None
        super(CompositeAlarm, __self__).__init__(
            'aws:cloudwatch/compositeAlarm:CompositeAlarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actions_enabled: Optional[pulumi.Input[bool]] = None,
            alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            alarm_description: Optional[pulumi.Input[str]] = None,
            alarm_name: Optional[pulumi.Input[str]] = None,
            alarm_rule: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'CompositeAlarm':
        """
        Get an existing CompositeAlarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[str] alarm_description: The description for the composite alarm.
        :param pulumi.Input[str] alarm_name: The name for the composite alarm. This name must be unique within the region.
        :param pulumi.Input[str] alarm_rule: An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        :param pulumi.Input[str] arn: The ARN of the composite alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions_enabled"] = actions_enabled
        __props__["alarm_actions"] = alarm_actions
        __props__["alarm_description"] = alarm_description
        __props__["alarm_name"] = alarm_name
        __props__["alarm_rule"] = alarm_rule
        __props__["arn"] = arn
        __props__["insufficient_data_actions"] = insufficient_data_actions
        __props__["ok_actions"] = ok_actions
        __props__["tags"] = tags
        return CompositeAlarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionsEnabled")
    def actions_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        """
        return pulumi.get(self, "actions_enabled")

    @property
    @pulumi.getter(name="alarmActions")
    def alarm_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "alarm_actions")

    @property
    @pulumi.getter(name="alarmDescription")
    def alarm_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the composite alarm.
        """
        return pulumi.get(self, "alarm_description")

    @property
    @pulumi.getter(name="alarmName")
    def alarm_name(self) -> pulumi.Output[str]:
        """
        The name for the composite alarm. This name must be unique within the region.
        """
        return pulumi.get(self, "alarm_name")

    @property
    @pulumi.getter(name="alarmRule")
    def alarm_rule(self) -> pulumi.Output[str]:
        """
        An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        """
        return pulumi.get(self, "alarm_rule")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the composite alarm.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="insufficientDataActions")
    def insufficient_data_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "insufficient_data_actions")

    @property
    @pulumi.getter(name="okActions")
    def ok_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "ok_actions")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

