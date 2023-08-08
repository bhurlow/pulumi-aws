# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MountTargetArgs', 'MountTarget']

@pulumi.input_type
class MountTargetArgs:
    def __init__(__self__, *,
                 file_system_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 ip_address: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a MountTarget resource.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which the mount target is intended.
        :param pulumi.Input[str] subnet_id: The ID of the subnet to add the mount target in.
        :param pulumi.Input[str] ip_address: The address (within the address range of the specified subnet) at
               which the file system may be mounted via the mount target.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: A list of up to 5 VPC security group IDs (that must
               be for the same VPC as subnet specified) in effect for the mount target.
        """
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if security_groups is not None:
            pulumi.set(__self__, "security_groups", security_groups)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        """
        The ID of the file system for which the mount target is intended.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The ID of the subnet to add the mount target in.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The address (within the address range of the specified subnet) at
        which the file system may be mounted via the mount target.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of up to 5 VPC security group IDs (that must
        be for the same VPC as subnet specified) in effect for the mount target.
        """
        return pulumi.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups", value)


@pulumi.input_type
class _MountTargetState:
    def __init__(__self__, *,
                 availability_zone_id: Optional[pulumi.Input[str]] = None,
                 availability_zone_name: Optional[pulumi.Input[str]] = None,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 file_system_arn: Optional[pulumi.Input[str]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 mount_target_dns_name: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MountTarget resources.
        :param pulumi.Input[str] availability_zone_id: The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
        :param pulumi.Input[str] availability_zone_name: The name of the Availability Zone (AZ) that the mount target resides in.
        :param pulumi.Input[str] dns_name: The DNS name for the EFS file system.
        :param pulumi.Input[str] file_system_arn: Amazon Resource Name of the file system.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which the mount target is intended.
        :param pulumi.Input[str] ip_address: The address (within the address range of the specified subnet) at
               which the file system may be mounted via the mount target.
        :param pulumi.Input[str] mount_target_dns_name: The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        :param pulumi.Input[str] network_interface_id: The ID of the network interface that Amazon EFS created when it created the mount target.
        :param pulumi.Input[str] owner_id: AWS account ID that owns the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: A list of up to 5 VPC security group IDs (that must
               be for the same VPC as subnet specified) in effect for the mount target.
        :param pulumi.Input[str] subnet_id: The ID of the subnet to add the mount target in.
        """
        if availability_zone_id is not None:
            pulumi.set(__self__, "availability_zone_id", availability_zone_id)
        if availability_zone_name is not None:
            pulumi.set(__self__, "availability_zone_name", availability_zone_name)
        if dns_name is not None:
            pulumi.set(__self__, "dns_name", dns_name)
        if file_system_arn is not None:
            pulumi.set(__self__, "file_system_arn", file_system_arn)
        if file_system_id is not None:
            pulumi.set(__self__, "file_system_id", file_system_id)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if mount_target_dns_name is not None:
            pulumi.set(__self__, "mount_target_dns_name", mount_target_dns_name)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if security_groups is not None:
            pulumi.set(__self__, "security_groups", security_groups)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="availabilityZoneId")
    def availability_zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
        """
        return pulumi.get(self, "availability_zone_id")

    @availability_zone_id.setter
    def availability_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone_id", value)

    @property
    @pulumi.getter(name="availabilityZoneName")
    def availability_zone_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Availability Zone (AZ) that the mount target resides in.
        """
        return pulumi.get(self, "availability_zone_name")

    @availability_zone_name.setter
    def availability_zone_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone_name", value)

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS name for the EFS file system.
        """
        return pulumi.get(self, "dns_name")

    @dns_name.setter
    def dns_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_name", value)

    @property
    @pulumi.getter(name="fileSystemArn")
    def file_system_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name of the file system.
        """
        return pulumi.get(self, "file_system_arn")

    @file_system_arn.setter
    def file_system_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_arn", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the file system for which the mount target is intended.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The address (within the address range of the specified subnet) at
        which the file system may be mounted via the mount target.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="mountTargetDnsName")
    def mount_target_dns_name(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        """
        return pulumi.get(self, "mount_target_dns_name")

    @mount_target_dns_name.setter
    def mount_target_dns_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_target_dns_name", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the network interface that Amazon EFS created when it created the mount target.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID that owns the resource.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of up to 5 VPC security group IDs (that must
        be for the same VPC as subnet specified) in effect for the mount target.
        """
        return pulumi.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the subnet to add the mount target in.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class MountTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Elastic File System (EFS) mount target.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ec2.Vpc("foo", cidr_block="10.0.0.0/16")
        alpha_subnet = aws.ec2.Subnet("alphaSubnet",
            vpc_id=foo.id,
            availability_zone="us-west-2a",
            cidr_block="10.0.1.0/24")
        alpha_mount_target = aws.efs.MountTarget("alphaMountTarget",
            file_system_id=aws_efs_file_system["foo"]["id"],
            subnet_id=alpha_subnet.id)
        ```

        ## Import

        terraform import {

         to = aws_efs_mount_target.alpha

         id = "fsmt-52a643fb" } Using `pulumi import`, import the EFS mount targets using the `id`. For exampleconsole % pulumi import aws_efs_mount_target.alpha fsmt-52a643fb

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which the mount target is intended.
        :param pulumi.Input[str] ip_address: The address (within the address range of the specified subnet) at
               which the file system may be mounted via the mount target.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: A list of up to 5 VPC security group IDs (that must
               be for the same VPC as subnet specified) in effect for the mount target.
        :param pulumi.Input[str] subnet_id: The ID of the subnet to add the mount target in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MountTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Elastic File System (EFS) mount target.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ec2.Vpc("foo", cidr_block="10.0.0.0/16")
        alpha_subnet = aws.ec2.Subnet("alphaSubnet",
            vpc_id=foo.id,
            availability_zone="us-west-2a",
            cidr_block="10.0.1.0/24")
        alpha_mount_target = aws.efs.MountTarget("alphaMountTarget",
            file_system_id=aws_efs_file_system["foo"]["id"],
            subnet_id=alpha_subnet.id)
        ```

        ## Import

        terraform import {

         to = aws_efs_mount_target.alpha

         id = "fsmt-52a643fb" } Using `pulumi import`, import the EFS mount targets using the `id`. For exampleconsole % pulumi import aws_efs_mount_target.alpha fsmt-52a643fb

        :param str resource_name: The name of the resource.
        :param MountTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MountTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MountTargetArgs.__new__(MountTargetArgs)

            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__.__dict__["file_system_id"] = file_system_id
            __props__.__dict__["ip_address"] = ip_address
            __props__.__dict__["security_groups"] = security_groups
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["availability_zone_id"] = None
            __props__.__dict__["availability_zone_name"] = None
            __props__.__dict__["dns_name"] = None
            __props__.__dict__["file_system_arn"] = None
            __props__.__dict__["mount_target_dns_name"] = None
            __props__.__dict__["network_interface_id"] = None
            __props__.__dict__["owner_id"] = None
        super(MountTarget, __self__).__init__(
            'aws:efs/mountTarget:MountTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone_id: Optional[pulumi.Input[str]] = None,
            availability_zone_name: Optional[pulumi.Input[str]] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            file_system_arn: Optional[pulumi.Input[str]] = None,
            file_system_id: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            mount_target_dns_name: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'MountTarget':
        """
        Get an existing MountTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone_id: The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
        :param pulumi.Input[str] availability_zone_name: The name of the Availability Zone (AZ) that the mount target resides in.
        :param pulumi.Input[str] dns_name: The DNS name for the EFS file system.
        :param pulumi.Input[str] file_system_arn: Amazon Resource Name of the file system.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which the mount target is intended.
        :param pulumi.Input[str] ip_address: The address (within the address range of the specified subnet) at
               which the file system may be mounted via the mount target.
        :param pulumi.Input[str] mount_target_dns_name: The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        :param pulumi.Input[str] network_interface_id: The ID of the network interface that Amazon EFS created when it created the mount target.
        :param pulumi.Input[str] owner_id: AWS account ID that owns the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: A list of up to 5 VPC security group IDs (that must
               be for the same VPC as subnet specified) in effect for the mount target.
        :param pulumi.Input[str] subnet_id: The ID of the subnet to add the mount target in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MountTargetState.__new__(_MountTargetState)

        __props__.__dict__["availability_zone_id"] = availability_zone_id
        __props__.__dict__["availability_zone_name"] = availability_zone_name
        __props__.__dict__["dns_name"] = dns_name
        __props__.__dict__["file_system_arn"] = file_system_arn
        __props__.__dict__["file_system_id"] = file_system_id
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["mount_target_dns_name"] = mount_target_dns_name
        __props__.__dict__["network_interface_id"] = network_interface_id
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["security_groups"] = security_groups
        __props__.__dict__["subnet_id"] = subnet_id
        return MountTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZoneId")
    def availability_zone_id(self) -> pulumi.Output[str]:
        """
        The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
        """
        return pulumi.get(self, "availability_zone_id")

    @property
    @pulumi.getter(name="availabilityZoneName")
    def availability_zone_name(self) -> pulumi.Output[str]:
        """
        The name of the Availability Zone (AZ) that the mount target resides in.
        """
        return pulumi.get(self, "availability_zone_name")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        """
        The DNS name for the EFS file system.
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="fileSystemArn")
    def file_system_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name of the file system.
        """
        return pulumi.get(self, "file_system_arn")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        """
        The ID of the file system for which the mount target is intended.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The address (within the address range of the specified subnet) at
        which the file system may be mounted via the mount target.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="mountTargetDnsName")
    def mount_target_dns_name(self) -> pulumi.Output[str]:
        """
        The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        """
        return pulumi.get(self, "mount_target_dns_name")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        The ID of the network interface that Amazon EFS created when it created the mount target.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        AWS account ID that owns the resource.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of up to 5 VPC security group IDs (that must
        be for the same VPC as subnet specified) in effect for the mount target.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the subnet to add the mount target in.
        """
        return pulumi.get(self, "subnet_id")

