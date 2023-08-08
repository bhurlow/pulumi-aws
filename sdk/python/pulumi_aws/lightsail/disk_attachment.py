# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Disk_attachmentArgs', 'Disk_attachment']

@pulumi.input_type
class Disk_attachmentArgs:
    def __init__(__self__, *,
                 disk_name: pulumi.Input[str],
                 disk_path: pulumi.Input[str],
                 instance_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a Disk_attachment resource.
        :param pulumi.Input[str] disk_name: The name of the Lightsail Disk.
        :param pulumi.Input[str] disk_path: The disk path to expose to the instance.
        :param pulumi.Input[str] instance_name: The name of the Lightsail Instance to attach to.
        """
        pulumi.set(__self__, "disk_name", disk_name)
        pulumi.set(__self__, "disk_path", disk_path)
        pulumi.set(__self__, "instance_name", instance_name)

    @property
    @pulumi.getter(name="diskName")
    def disk_name(self) -> pulumi.Input[str]:
        """
        The name of the Lightsail Disk.
        """
        return pulumi.get(self, "disk_name")

    @disk_name.setter
    def disk_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_name", value)

    @property
    @pulumi.getter(name="diskPath")
    def disk_path(self) -> pulumi.Input[str]:
        """
        The disk path to expose to the instance.
        """
        return pulumi.get(self, "disk_path")

    @disk_path.setter
    def disk_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_path", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[str]:
        """
        The name of the Lightsail Instance to attach to.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_name", value)


@pulumi.input_type
class _Disk_attachmentState:
    def __init__(__self__, *,
                 disk_name: Optional[pulumi.Input[str]] = None,
                 disk_path: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Disk_attachment resources.
        :param pulumi.Input[str] disk_name: The name of the Lightsail Disk.
        :param pulumi.Input[str] disk_path: The disk path to expose to the instance.
        :param pulumi.Input[str] instance_name: The name of the Lightsail Instance to attach to.
        """
        if disk_name is not None:
            pulumi.set(__self__, "disk_name", disk_name)
        if disk_path is not None:
            pulumi.set(__self__, "disk_path", disk_path)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)

    @property
    @pulumi.getter(name="diskName")
    def disk_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Lightsail Disk.
        """
        return pulumi.get(self, "disk_name")

    @disk_name.setter
    def disk_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_name", value)

    @property
    @pulumi.getter(name="diskPath")
    def disk_path(self) -> Optional[pulumi.Input[str]]:
        """
        The disk path to expose to the instance.
        """
        return pulumi.get(self, "disk_path")

    @disk_path.setter
    def disk_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_path", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Lightsail Instance to attach to.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)


class Disk_attachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_name: Optional[pulumi.Input[str]] = None,
                 disk_path: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attaches a Lightsail disk to a Lightsail Instance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        available = aws.get_availability_zones(state="available",
            filters=[aws.GetAvailabilityZonesFilterArgs(
                name="opt-in-status",
                values=["opt-in-not-required"],
            )])
        test_disk = aws.lightsail.Disk("testDisk",
            size_in_gb=8,
            availability_zone=available.names[0])
        test_instance = aws.lightsail.Instance("testInstance",
            availability_zone=available.names[0],
            blueprint_id="amazon_linux_2",
            bundle_id="nano_1_0")
        test_disk_attachment = aws.lightsail.Disk_attachment("testDisk_attachment",
            disk_name=test_disk.name,
            instance_name=test_instance.name,
            disk_path="/dev/xvdf")
        ```

        ## Import

        terraform import {

         to = aws_lightsail_disk_attachment.test

         id = "test-disk,test-instance" } Using `pulumi import`, import `aws_lightsail_disk` using the id attribute. For exampleconsole % pulumi import aws_lightsail_disk_attachment.test test-disk,test-instance

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_name: The name of the Lightsail Disk.
        :param pulumi.Input[str] disk_path: The disk path to expose to the instance.
        :param pulumi.Input[str] instance_name: The name of the Lightsail Instance to attach to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Disk_attachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attaches a Lightsail disk to a Lightsail Instance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        available = aws.get_availability_zones(state="available",
            filters=[aws.GetAvailabilityZonesFilterArgs(
                name="opt-in-status",
                values=["opt-in-not-required"],
            )])
        test_disk = aws.lightsail.Disk("testDisk",
            size_in_gb=8,
            availability_zone=available.names[0])
        test_instance = aws.lightsail.Instance("testInstance",
            availability_zone=available.names[0],
            blueprint_id="amazon_linux_2",
            bundle_id="nano_1_0")
        test_disk_attachment = aws.lightsail.Disk_attachment("testDisk_attachment",
            disk_name=test_disk.name,
            instance_name=test_instance.name,
            disk_path="/dev/xvdf")
        ```

        ## Import

        terraform import {

         to = aws_lightsail_disk_attachment.test

         id = "test-disk,test-instance" } Using `pulumi import`, import `aws_lightsail_disk` using the id attribute. For exampleconsole % pulumi import aws_lightsail_disk_attachment.test test-disk,test-instance

        :param str resource_name: The name of the resource.
        :param Disk_attachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Disk_attachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_name: Optional[pulumi.Input[str]] = None,
                 disk_path: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Disk_attachmentArgs.__new__(Disk_attachmentArgs)

            if disk_name is None and not opts.urn:
                raise TypeError("Missing required property 'disk_name'")
            __props__.__dict__["disk_name"] = disk_name
            if disk_path is None and not opts.urn:
                raise TypeError("Missing required property 'disk_path'")
            __props__.__dict__["disk_path"] = disk_path
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
        super(Disk_attachment, __self__).__init__(
            'aws:lightsail/disk_attachment:Disk_attachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disk_name: Optional[pulumi.Input[str]] = None,
            disk_path: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None) -> 'Disk_attachment':
        """
        Get an existing Disk_attachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_name: The name of the Lightsail Disk.
        :param pulumi.Input[str] disk_path: The disk path to expose to the instance.
        :param pulumi.Input[str] instance_name: The name of the Lightsail Instance to attach to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Disk_attachmentState.__new__(_Disk_attachmentState)

        __props__.__dict__["disk_name"] = disk_name
        __props__.__dict__["disk_path"] = disk_path
        __props__.__dict__["instance_name"] = instance_name
        return Disk_attachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="diskName")
    def disk_name(self) -> pulumi.Output[str]:
        """
        The name of the Lightsail Disk.
        """
        return pulumi.get(self, "disk_name")

    @property
    @pulumi.getter(name="diskPath")
    def disk_path(self) -> pulumi.Output[str]:
        """
        The disk path to expose to the instance.
        """
        return pulumi.get(self, "disk_path")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        The name of the Lightsail Instance to attach to.
        """
        return pulumi.get(self, "instance_name")

