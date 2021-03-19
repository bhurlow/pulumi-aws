# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SnapshotSchedule']


class SnapshotSchedule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.redshift.SnapshotSchedule("default",
            definitions=["rate(12 hours)"],
            identifier="tf-redshift-snapshot-schedule")
        ```

        ## Import

        Redshift Snapshot Schedule can be imported using the `identifier`, e.g.

        ```sh
         $ pulumi import aws:redshift/snapshotSchedule:SnapshotSchedule default tf-redshift-snapshot-schedule
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        :param pulumi.Input[str] description: The description of the snapshot schedule.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        :param pulumi.Input[str] identifier: The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique
               identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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

            if definitions is None and not opts.urn:
                raise TypeError("Missing required property 'definitions'")
            __props__['definitions'] = definitions
            __props__['description'] = description
            __props__['force_destroy'] = force_destroy
            __props__['identifier'] = identifier
            __props__['identifier_prefix'] = identifier_prefix
            __props__['tags'] = tags
            __props__['arn'] = None
        super(SnapshotSchedule, __self__).__init__(
            'aws:redshift/snapshotSchedule:SnapshotSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            identifier_prefix: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SnapshotSchedule':
        """
        Get an existing SnapshotSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Redshift Snapshot Schedule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        :param pulumi.Input[str] description: The description of the snapshot schedule.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        :param pulumi.Input[str] identifier: The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique
               identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["definitions"] = definitions
        __props__["description"] = description
        __props__["force_destroy"] = force_destroy
        __props__["identifier"] = identifier
        __props__["identifier_prefix"] = identifier_prefix
        __props__["tags"] = tags
        return SnapshotSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Redshift Snapshot Schedule.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def definitions(self) -> pulumi.Output[Sequence[str]]:
        """
        The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
        """
        return pulumi.get(self, "definitions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the snapshot schedule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        """
        The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter(name="identifierPrefix")
    def identifier_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique
        identifier beginning with the specified prefix. Conflicts with `identifier`.
        """
        return pulumi.get(self, "identifier_prefix")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

