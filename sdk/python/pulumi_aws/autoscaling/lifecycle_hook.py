# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LifecycleHookArgs', 'LifecycleHook']

@pulumi.input_type
class LifecycleHookArgs:
    def __init__(__self__, *,
                 autoscaling_group_name: pulumi.Input[str],
                 lifecycle_transition: pulumi.Input[str],
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LifecycleHook resource.
        :param pulumi.Input[str] autoscaling_group_name: Name of the Auto Scaling group to which you want to assign the lifecycle hook
        :param pulumi.Input[str] lifecycle_transition: Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        :param pulumi.Input[str] name: Name of the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        :param pulumi.Input[str] role_arn: ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        pulumi.set(__self__, "autoscaling_group_name", autoscaling_group_name)
        pulumi.set(__self__, "lifecycle_transition", lifecycle_transition)
        if default_result is not None:
            pulumi.set(__self__, "default_result", default_result)
        if heartbeat_timeout is not None:
            pulumi.set(__self__, "heartbeat_timeout", heartbeat_timeout)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_metadata is not None:
            pulumi.set(__self__, "notification_metadata", notification_metadata)
        if notification_target_arn is not None:
            pulumi.set(__self__, "notification_target_arn", notification_target_arn)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="autoscalingGroupName")
    def autoscaling_group_name(self) -> pulumi.Input[str]:
        """
        Name of the Auto Scaling group to which you want to assign the lifecycle hook
        """
        return pulumi.get(self, "autoscaling_group_name")

    @autoscaling_group_name.setter
    def autoscaling_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "autoscaling_group_name", value)

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> pulumi.Input[str]:
        """
        Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        """
        return pulumi.get(self, "lifecycle_transition")

    @lifecycle_transition.setter
    def lifecycle_transition(self, value: pulumi.Input[str]):
        pulumi.set(self, "lifecycle_transition", value)

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        """
        return pulumi.get(self, "default_result")

    @default_result.setter
    def default_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_result", value)

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        """
        return pulumi.get(self, "heartbeat_timeout")

    @heartbeat_timeout.setter
    def heartbeat_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "heartbeat_timeout", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the lifecycle hook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @notification_metadata.setter
    def notification_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_metadata", value)

    @property
    @pulumi.getter(name="notificationTargetArn")
    def notification_target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        """
        return pulumi.get(self, "notification_target_arn")

    @notification_target_arn.setter
    def notification_target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_target_arn", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


@pulumi.input_type
class _LifecycleHookState:
    def __init__(__self__, *,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LifecycleHook resources.
        :param pulumi.Input[str] autoscaling_group_name: Name of the Auto Scaling group to which you want to assign the lifecycle hook
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        :param pulumi.Input[str] lifecycle_transition: Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        :param pulumi.Input[str] name: Name of the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        :param pulumi.Input[str] role_arn: ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        if autoscaling_group_name is not None:
            pulumi.set(__self__, "autoscaling_group_name", autoscaling_group_name)
        if default_result is not None:
            pulumi.set(__self__, "default_result", default_result)
        if heartbeat_timeout is not None:
            pulumi.set(__self__, "heartbeat_timeout", heartbeat_timeout)
        if lifecycle_transition is not None:
            pulumi.set(__self__, "lifecycle_transition", lifecycle_transition)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_metadata is not None:
            pulumi.set(__self__, "notification_metadata", notification_metadata)
        if notification_target_arn is not None:
            pulumi.set(__self__, "notification_target_arn", notification_target_arn)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="autoscalingGroupName")
    def autoscaling_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Auto Scaling group to which you want to assign the lifecycle hook
        """
        return pulumi.get(self, "autoscaling_group_name")

    @autoscaling_group_name.setter
    def autoscaling_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "autoscaling_group_name", value)

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        """
        return pulumi.get(self, "default_result")

    @default_result.setter
    def default_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_result", value)

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        """
        return pulumi.get(self, "heartbeat_timeout")

    @heartbeat_timeout.setter
    def heartbeat_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "heartbeat_timeout", value)

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> Optional[pulumi.Input[str]]:
        """
        Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        """
        return pulumi.get(self, "lifecycle_transition")

    @lifecycle_transition.setter
    def lifecycle_transition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_transition", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the lifecycle hook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @notification_metadata.setter
    def notification_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_metadata", value)

    @property
    @pulumi.getter(name="notificationTargetArn")
    def notification_target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        """
        return pulumi.get(self, "notification_target_arn")

    @notification_target_arn.setter
    def notification_target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_target_arn", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


class LifecycleHook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AutoScaling Lifecycle Hook resource.

        > **NOTE:** This provider has two types of ways you can add lifecycle hooks - via
        the `initial_lifecycle_hook` attribute from the
        `autoscaling.Group`
        resource, or via this one. Hooks added via this resource will not be added
        until the autoscaling group has been created, and depending on your
        capacity
        settings, after the initial instances have been launched, creating unintended
        behavior. If you need hooks to run on all instances, add them with
        `initial_lifecycle_hook` in
        `autoscaling.Group`,
        but take care to not duplicate those hooks with this resource.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        foobar_group = aws.autoscaling.Group("foobarGroup",
            availability_zones=["us-west-2a"],
            health_check_type="EC2",
            termination_policies=["OldestInstance"],
            tags=[aws.autoscaling.GroupTagArgs(
                key="Foo",
                value="foo-bar",
                propagate_at_launch=True,
            )])
        foobar_lifecycle_hook = aws.autoscaling.LifecycleHook("foobarLifecycleHook",
            autoscaling_group_name=foobar_group.name,
            default_result="CONTINUE",
            heartbeat_timeout=2000,
            lifecycle_transition="autoscaling:EC2_INSTANCE_LAUNCHING",
            notification_metadata=json.dumps({
                "foo": "bar",
            }),
            notification_target_arn="arn:aws:sqs:us-east-1:444455556666:queue1*",
            role_arn="arn:aws:iam::123456789012:role/S3Access")
        ```

        ## Import

        terraform import {

         to = aws_autoscaling_lifecycle_hook.test-lifecycle-hook

         id = "asg-name/lifecycle-hook-name" } Using `pulumi import`, import AutoScaling Lifecycle Hooks using the role autoscaling_group_name and name separated by `/`. For exampleconsole % pulumi import aws_autoscaling_lifecycle_hook.test-lifecycle-hook asg-name/lifecycle-hook-name

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscaling_group_name: Name of the Auto Scaling group to which you want to assign the lifecycle hook
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        :param pulumi.Input[str] lifecycle_transition: Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        :param pulumi.Input[str] name: Name of the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        :param pulumi.Input[str] role_arn: ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LifecycleHookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AutoScaling Lifecycle Hook resource.

        > **NOTE:** This provider has two types of ways you can add lifecycle hooks - via
        the `initial_lifecycle_hook` attribute from the
        `autoscaling.Group`
        resource, or via this one. Hooks added via this resource will not be added
        until the autoscaling group has been created, and depending on your
        capacity
        settings, after the initial instances have been launched, creating unintended
        behavior. If you need hooks to run on all instances, add them with
        `initial_lifecycle_hook` in
        `autoscaling.Group`,
        but take care to not duplicate those hooks with this resource.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        foobar_group = aws.autoscaling.Group("foobarGroup",
            availability_zones=["us-west-2a"],
            health_check_type="EC2",
            termination_policies=["OldestInstance"],
            tags=[aws.autoscaling.GroupTagArgs(
                key="Foo",
                value="foo-bar",
                propagate_at_launch=True,
            )])
        foobar_lifecycle_hook = aws.autoscaling.LifecycleHook("foobarLifecycleHook",
            autoscaling_group_name=foobar_group.name,
            default_result="CONTINUE",
            heartbeat_timeout=2000,
            lifecycle_transition="autoscaling:EC2_INSTANCE_LAUNCHING",
            notification_metadata=json.dumps({
                "foo": "bar",
            }),
            notification_target_arn="arn:aws:sqs:us-east-1:444455556666:queue1*",
            role_arn="arn:aws:iam::123456789012:role/S3Access")
        ```

        ## Import

        terraform import {

         to = aws_autoscaling_lifecycle_hook.test-lifecycle-hook

         id = "asg-name/lifecycle-hook-name" } Using `pulumi import`, import AutoScaling Lifecycle Hooks using the role autoscaling_group_name and name separated by `/`. For exampleconsole % pulumi import aws_autoscaling_lifecycle_hook.test-lifecycle-hook asg-name/lifecycle-hook-name

        :param str resource_name: The name of the resource.
        :param LifecycleHookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LifecycleHookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LifecycleHookArgs.__new__(LifecycleHookArgs)

            if autoscaling_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'autoscaling_group_name'")
            __props__.__dict__["autoscaling_group_name"] = autoscaling_group_name
            __props__.__dict__["default_result"] = default_result
            __props__.__dict__["heartbeat_timeout"] = heartbeat_timeout
            if lifecycle_transition is None and not opts.urn:
                raise TypeError("Missing required property 'lifecycle_transition'")
            __props__.__dict__["lifecycle_transition"] = lifecycle_transition
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_metadata"] = notification_metadata
            __props__.__dict__["notification_target_arn"] = notification_target_arn
            __props__.__dict__["role_arn"] = role_arn
        super(LifecycleHook, __self__).__init__(
            'aws:autoscaling/lifecycleHook:LifecycleHook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscaling_group_name: Optional[pulumi.Input[str]] = None,
            default_result: Optional[pulumi.Input[str]] = None,
            heartbeat_timeout: Optional[pulumi.Input[int]] = None,
            lifecycle_transition: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_metadata: Optional[pulumi.Input[str]] = None,
            notification_target_arn: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None) -> 'LifecycleHook':
        """
        Get an existing LifecycleHook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscaling_group_name: Name of the Auto Scaling group to which you want to assign the lifecycle hook
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        :param pulumi.Input[str] lifecycle_transition: Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        :param pulumi.Input[str] name: Name of the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        :param pulumi.Input[str] role_arn: ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LifecycleHookState.__new__(_LifecycleHookState)

        __props__.__dict__["autoscaling_group_name"] = autoscaling_group_name
        __props__.__dict__["default_result"] = default_result
        __props__.__dict__["heartbeat_timeout"] = heartbeat_timeout
        __props__.__dict__["lifecycle_transition"] = lifecycle_transition
        __props__.__dict__["name"] = name
        __props__.__dict__["notification_metadata"] = notification_metadata
        __props__.__dict__["notification_target_arn"] = notification_target_arn
        __props__.__dict__["role_arn"] = role_arn
        return LifecycleHook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoscalingGroupName")
    def autoscaling_group_name(self) -> pulumi.Output[str]:
        """
        Name of the Auto Scaling group to which you want to assign the lifecycle hook
        """
        return pulumi.get(self, "autoscaling_group_name")

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> pulumi.Output[str]:
        """
        Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        """
        return pulumi.get(self, "default_result")

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        """
        return pulumi.get(self, "heartbeat_timeout")

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> pulumi.Output[str]:
        """
        Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        """
        return pulumi.get(self, "lifecycle_transition")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the lifecycle hook.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> pulumi.Output[Optional[str]]:
        """
        Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @property
    @pulumi.getter(name="notificationTargetArn")
    def notification_target_arn(self) -> pulumi.Output[Optional[str]]:
        """
        ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        """
        return pulumi.get(self, "notification_target_arn")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        return pulumi.get(self, "role_arn")

