# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EventRuleArgs', 'EventRule']

@pulumi.input_type
class EventRuleArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a EventRule resource.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] event_bus_name: The name or ARN of the event bus to associate with this rule.
               If you omit this, the `default` event bus is used.
        :param pulumi.Input[str] event_pattern: The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. **Note**: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
        :param pulumi.Input[bool] is_enabled: Whether the rule should be enabled (defaults to `true`).
        :param pulumi.Input[str] name: The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        :param pulumi.Input[str] schedule_expression: The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if event_bus_name is not None:
            pulumi.set(__self__, "event_bus_name", event_bus_name)
        if event_pattern is not None:
            pulumi.set(__self__, "event_pattern", event_pattern)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if schedule_expression is not None:
            pulumi.set(__self__, "schedule_expression", schedule_expression)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or ARN of the event bus to associate with this rule.
        If you omit this, the `default` event bus is used.
        """
        return pulumi.get(self, "event_bus_name")

    @event_bus_name.setter
    def event_bus_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_bus_name", value)

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. **Note**: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
        """
        return pulumi.get(self, "event_pattern")

    @event_pattern.setter
    def event_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_pattern", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the rule should be enabled (defaults to `true`).
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> Optional[pulumi.Input[str]]:
        """
        The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
        """
        return pulumi.get(self, "schedule_expression")

    @schedule_expression.setter
    def schedule_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_expression", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EventRuleState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering EventRule resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the rule.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] event_bus_name: The name or ARN of the event bus to associate with this rule.
               If you omit this, the `default` event bus is used.
        :param pulumi.Input[str] event_pattern: The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. **Note**: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
        :param pulumi.Input[bool] is_enabled: Whether the rule should be enabled (defaults to `true`).
        :param pulumi.Input[str] name: The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        :param pulumi.Input[str] schedule_expression: The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if event_bus_name is not None:
            pulumi.set(__self__, "event_bus_name", event_bus_name)
        if event_pattern is not None:
            pulumi.set(__self__, "event_pattern", event_pattern)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if schedule_expression is not None:
            pulumi.set(__self__, "schedule_expression", schedule_expression)
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
        The Amazon Resource Name (ARN) of the rule.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or ARN of the event bus to associate with this rule.
        If you omit this, the `default` event bus is used.
        """
        return pulumi.get(self, "event_bus_name")

    @event_bus_name.setter
    def event_bus_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_bus_name", value)

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. **Note**: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
        """
        return pulumi.get(self, "event_pattern")

    @event_pattern.setter
    def event_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_pattern", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the rule should be enabled (defaults to `true`).
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> Optional[pulumi.Input[str]]:
        """
        The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
        """
        return pulumi.get(self, "schedule_expression")

    @schedule_expression.setter
    def schedule_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_expression", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class EventRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an EventBridge Rule resource.

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        console = aws.cloudwatch.EventRule("console",
            description="Capture each AWS Console Sign In",
            event_pattern=json.dumps({
                "detail-type": ["AWS Console Sign In via CloudTrail"],
            }))
        aws_logins = aws.sns.Topic("awsLogins")
        sns = aws.cloudwatch.EventTarget("sns",
            rule=console.name,
            arn=aws_logins.arn)
        sns_topic_policy = aws_logins.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            actions=["SNS:Publish"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="Service",
                identifiers=["events.amazonaws.com"],
            )],
            resources=[arn],
        )]))
        default = aws.sns.TopicPolicy("default",
            arn=aws_logins.arn,
            policy=sns_topic_policy.json)
        ```

        ## Import

        Using `pulumi import`, import EventBridge Rules using the `event_bus_name/rule_name` (if you omit `event_bus_name`, the `default` event bus will be used). For example:

        ```sh
         $ pulumi import aws:cloudwatch/eventRule:EventRule console example-event-bus/capture-console-sign-in
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] event_bus_name: The name or ARN of the event bus to associate with this rule.
               If you omit this, the `default` event bus is used.
        :param pulumi.Input[str] event_pattern: The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. **Note**: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
        :param pulumi.Input[bool] is_enabled: Whether the rule should be enabled (defaults to `true`).
        :param pulumi.Input[str] name: The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        :param pulumi.Input[str] schedule_expression: The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EventRuleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an EventBridge Rule resource.

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        console = aws.cloudwatch.EventRule("console",
            description="Capture each AWS Console Sign In",
            event_pattern=json.dumps({
                "detail-type": ["AWS Console Sign In via CloudTrail"],
            }))
        aws_logins = aws.sns.Topic("awsLogins")
        sns = aws.cloudwatch.EventTarget("sns",
            rule=console.name,
            arn=aws_logins.arn)
        sns_topic_policy = aws_logins.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            actions=["SNS:Publish"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="Service",
                identifiers=["events.amazonaws.com"],
            )],
            resources=[arn],
        )]))
        default = aws.sns.TopicPolicy("default",
            arn=aws_logins.arn,
            policy=sns_topic_policy.json)
        ```

        ## Import

        Using `pulumi import`, import EventBridge Rules using the `event_bus_name/rule_name` (if you omit `event_bus_name`, the `default` event bus will be used). For example:

        ```sh
         $ pulumi import aws:cloudwatch/eventRule:EventRule console example-event-bus/capture-console-sign-in
        ```

        :param str resource_name: The name of the resource.
        :param EventRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventRuleArgs.__new__(EventRuleArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["event_bus_name"] = event_bus_name
            __props__.__dict__["event_pattern"] = event_pattern
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["schedule_expression"] = schedule_expression
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(EventRule, __self__).__init__(
            'aws:cloudwatch/eventRule:EventRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            event_bus_name: Optional[pulumi.Input[str]] = None,
            event_pattern: Optional[pulumi.Input[str]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            schedule_expression: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'EventRule':
        """
        Get an existing EventRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the rule.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] event_bus_name: The name or ARN of the event bus to associate with this rule.
               If you omit this, the `default` event bus is used.
        :param pulumi.Input[str] event_pattern: The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. **Note**: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
        :param pulumi.Input[bool] is_enabled: Whether the rule should be enabled (defaults to `true`).
        :param pulumi.Input[str] name: The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        :param pulumi.Input[str] schedule_expression: The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventRuleState.__new__(_EventRuleState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["event_bus_name"] = event_bus_name
        __props__.__dict__["event_pattern"] = event_pattern
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["schedule_expression"] = schedule_expression
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return EventRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the rule.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name or ARN of the event bus to associate with this rule.
        If you omit this, the `default` event bus is used.
        """
        return pulumi.get(self, "event_bus_name")

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> pulumi.Output[Optional[str]]:
        """
        The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details. **Note**: The event pattern size is 2048 by default but it is adjustable up to 4096 characters by submitting a service quota increase request. See [Amazon EventBridge quotas](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-quota.html) for details.
        """
        return pulumi.get(self, "event_pattern")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the rule should be enabled (defaults to `true`).
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> pulumi.Output[Optional[str]]:
        """
        The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
        """
        return pulumi.get(self, "schedule_expression")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

