# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FrameworkShareArgs', 'FrameworkShare']

@pulumi.input_type
class FrameworkShareArgs:
    def __init__(__self__, *,
                 destination_account: pulumi.Input[str],
                 destination_region: pulumi.Input[str],
                 framework_id: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FrameworkShare resource.
        :param pulumi.Input[str] destination_account: Amazon Web Services account of the recipient.
        :param pulumi.Input[str] destination_region: Amazon Web Services region of the recipient.
        :param pulumi.Input[str] framework_id: Unique identifier for the shared custom framework.
               
               The following arguments are optional:
        :param pulumi.Input[str] comment: Comment from the sender about the share request.
        """
        pulumi.set(__self__, "destination_account", destination_account)
        pulumi.set(__self__, "destination_region", destination_region)
        pulumi.set(__self__, "framework_id", framework_id)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)

    @property
    @pulumi.getter(name="destinationAccount")
    def destination_account(self) -> pulumi.Input[str]:
        """
        Amazon Web Services account of the recipient.
        """
        return pulumi.get(self, "destination_account")

    @destination_account.setter
    def destination_account(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_account", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> pulumi.Input[str]:
        """
        Amazon Web Services region of the recipient.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="frameworkId")
    def framework_id(self) -> pulumi.Input[str]:
        """
        Unique identifier for the shared custom framework.

        The following arguments are optional:
        """
        return pulumi.get(self, "framework_id")

    @framework_id.setter
    def framework_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "framework_id", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment from the sender about the share request.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)


@pulumi.input_type
class _FrameworkShareState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 destination_account: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 framework_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FrameworkShare resources.
        :param pulumi.Input[str] comment: Comment from the sender about the share request.
        :param pulumi.Input[str] destination_account: Amazon Web Services account of the recipient.
        :param pulumi.Input[str] destination_region: Amazon Web Services region of the recipient.
        :param pulumi.Input[str] framework_id: Unique identifier for the shared custom framework.
               
               The following arguments are optional:
        :param pulumi.Input[str] status: Status of the share request.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if destination_account is not None:
            pulumi.set(__self__, "destination_account", destination_account)
        if destination_region is not None:
            pulumi.set(__self__, "destination_region", destination_region)
        if framework_id is not None:
            pulumi.set(__self__, "framework_id", framework_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment from the sender about the share request.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="destinationAccount")
    def destination_account(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Web Services account of the recipient.
        """
        return pulumi.get(self, "destination_account")

    @destination_account.setter
    def destination_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_account", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Web Services region of the recipient.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="frameworkId")
    def framework_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the shared custom framework.

        The following arguments are optional:
        """
        return pulumi.get(self, "framework_id")

    @framework_id.setter
    def framework_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "framework_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the share request.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class FrameworkShare(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 destination_account: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 framework_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource for managing an AWS Audit Manager Framework Share.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.auditmanager.FrameworkShare("example",
            destination_account="012345678901",
            destination_region="us-east-1",
            framework_id=aws_auditmanager_framework["example"]["id"])
        ```

        ## Import

        terraform import {

         to = aws_auditmanager_framework_share.example

         id = "abcdef-123456" } Using `pulumi import`, import Audit Manager Framework Share using the `id`. For exampleconsole % pulumi import aws_auditmanager_framework_share.example abcdef-123456

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment from the sender about the share request.
        :param pulumi.Input[str] destination_account: Amazon Web Services account of the recipient.
        :param pulumi.Input[str] destination_region: Amazon Web Services region of the recipient.
        :param pulumi.Input[str] framework_id: Unique identifier for the shared custom framework.
               
               The following arguments are optional:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FrameworkShareArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing an AWS Audit Manager Framework Share.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.auditmanager.FrameworkShare("example",
            destination_account="012345678901",
            destination_region="us-east-1",
            framework_id=aws_auditmanager_framework["example"]["id"])
        ```

        ## Import

        terraform import {

         to = aws_auditmanager_framework_share.example

         id = "abcdef-123456" } Using `pulumi import`, import Audit Manager Framework Share using the `id`. For exampleconsole % pulumi import aws_auditmanager_framework_share.example abcdef-123456

        :param str resource_name: The name of the resource.
        :param FrameworkShareArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FrameworkShareArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 destination_account: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 framework_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FrameworkShareArgs.__new__(FrameworkShareArgs)

            __props__.__dict__["comment"] = comment
            if destination_account is None and not opts.urn:
                raise TypeError("Missing required property 'destination_account'")
            __props__.__dict__["destination_account"] = destination_account
            if destination_region is None and not opts.urn:
                raise TypeError("Missing required property 'destination_region'")
            __props__.__dict__["destination_region"] = destination_region
            if framework_id is None and not opts.urn:
                raise TypeError("Missing required property 'framework_id'")
            __props__.__dict__["framework_id"] = framework_id
            __props__.__dict__["status"] = None
        super(FrameworkShare, __self__).__init__(
            'aws:auditmanager/frameworkShare:FrameworkShare',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            destination_account: Optional[pulumi.Input[str]] = None,
            destination_region: Optional[pulumi.Input[str]] = None,
            framework_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'FrameworkShare':
        """
        Get an existing FrameworkShare resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment from the sender about the share request.
        :param pulumi.Input[str] destination_account: Amazon Web Services account of the recipient.
        :param pulumi.Input[str] destination_region: Amazon Web Services region of the recipient.
        :param pulumi.Input[str] framework_id: Unique identifier for the shared custom framework.
               
               The following arguments are optional:
        :param pulumi.Input[str] status: Status of the share request.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FrameworkShareState.__new__(_FrameworkShareState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["destination_account"] = destination_account
        __props__.__dict__["destination_region"] = destination_region
        __props__.__dict__["framework_id"] = framework_id
        __props__.__dict__["status"] = status
        return FrameworkShare(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment from the sender about the share request.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="destinationAccount")
    def destination_account(self) -> pulumi.Output[str]:
        """
        Amazon Web Services account of the recipient.
        """
        return pulumi.get(self, "destination_account")

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> pulumi.Output[str]:
        """
        Amazon Web Services region of the recipient.
        """
        return pulumi.get(self, "destination_region")

    @property
    @pulumi.getter(name="frameworkId")
    def framework_id(self) -> pulumi.Output[str]:
        """
        Unique identifier for the shared custom framework.

        The following arguments are optional:
        """
        return pulumi.get(self, "framework_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the share request.
        """
        return pulumi.get(self, "status")

