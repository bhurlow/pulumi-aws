# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ScheduleGroupArgs', 'ScheduleGroup']

@pulumi.input_type
class ScheduleGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ScheduleGroup resource.
        :param pulumi.Input[str] name: Name of the schedule group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the schedule group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

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
class _ScheduleGroupState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 creation_date: Optional[pulumi.Input[str]] = None,
                 last_modification_date: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ScheduleGroup resources.
        :param pulumi.Input[str] arn: ARN of the schedule group.
        :param pulumi.Input[str] creation_date: Time at which the schedule group was created.
        :param pulumi.Input[str] last_modification_date: Time at which the schedule group was last modified.
        :param pulumi.Input[str] name: Name of the schedule group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] state: State of the schedule group. Can be `ACTIVE` or `DELETING`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if creation_date is not None:
            pulumi.set(__self__, "creation_date", creation_date)
        if last_modification_date is not None:
            pulumi.set(__self__, "last_modification_date", last_modification_date)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if state is not None:
            pulumi.set(__self__, "state", state)
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
        ARN of the schedule group.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[pulumi.Input[str]]:
        """
        Time at which the schedule group was created.
        """
        return pulumi.get(self, "creation_date")

    @creation_date.setter
    def creation_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_date", value)

    @property
    @pulumi.getter(name="lastModificationDate")
    def last_modification_date(self) -> Optional[pulumi.Input[str]]:
        """
        Time at which the schedule group was last modified.
        """
        return pulumi.get(self, "last_modification_date")

    @last_modification_date.setter
    def last_modification_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modification_date", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the schedule group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the schedule group. Can be `ACTIVE` or `DELETING`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

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
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class ScheduleGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an EventBridge Scheduler Schedule Group resource.

        You can find out more about EventBridge Scheduler in the [User Guide](https://docs.aws.amazon.com/scheduler/latest/UserGuide/what-is-scheduler.html).

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.scheduler.ScheduleGroup("example")
        ```

        ## Import

        Using `pulumi import`, import schedule groups using the `name`. For example:

        ```sh
         $ pulumi import aws:scheduler/scheduleGroup:ScheduleGroup example my-schedule-group
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the schedule group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ScheduleGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an EventBridge Scheduler Schedule Group resource.

        You can find out more about EventBridge Scheduler in the [User Guide](https://docs.aws.amazon.com/scheduler/latest/UserGuide/what-is-scheduler.html).

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.scheduler.ScheduleGroup("example")
        ```

        ## Import

        Using `pulumi import`, import schedule groups using the `name`. For example:

        ```sh
         $ pulumi import aws:scheduler/scheduleGroup:ScheduleGroup example my-schedule-group
        ```

        :param str resource_name: The name of the resource.
        :param ScheduleGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduleGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScheduleGroupArgs.__new__(ScheduleGroupArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_date"] = None
            __props__.__dict__["last_modification_date"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ScheduleGroup, __self__).__init__(
            'aws:scheduler/scheduleGroup:ScheduleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            creation_date: Optional[pulumi.Input[str]] = None,
            last_modification_date: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ScheduleGroup':
        """
        Get an existing ScheduleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the schedule group.
        :param pulumi.Input[str] creation_date: Time at which the schedule group was created.
        :param pulumi.Input[str] last_modification_date: Time at which the schedule group was last modified.
        :param pulumi.Input[str] name: Name of the schedule group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] state: State of the schedule group. Can be `ACTIVE` or `DELETING`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduleGroupState.__new__(_ScheduleGroupState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["creation_date"] = creation_date
        __props__.__dict__["last_modification_date"] = last_modification_date
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        __props__.__dict__["state"] = state
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return ScheduleGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the schedule group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        Time at which the schedule group was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="lastModificationDate")
    def last_modification_date(self) -> pulumi.Output[str]:
        """
        Time at which the schedule group was last modified.
        """
        return pulumi.get(self, "last_modification_date")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the schedule group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the schedule group. Can be `ACTIVE` or `DELETING`.
        """
        return pulumi.get(self, "state")

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
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

