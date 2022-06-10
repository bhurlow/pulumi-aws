# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnapshotCreateVolumePermissionArgs', 'SnapshotCreateVolumePermission']

@pulumi.input_type
class SnapshotCreateVolumePermissionArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 snapshot_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a SnapshotCreateVolumePermission resource.
        :param pulumi.Input[str] account_id: An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
        :param pulumi.Input[str] snapshot_id: A snapshot ID
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "snapshot_id", snapshot_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        """
        An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Input[str]:
        """
        A snapshot ID
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_id", value)


@pulumi.input_type
class _SnapshotCreateVolumePermissionState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnapshotCreateVolumePermission resources.
        :param pulumi.Input[str] account_id: An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
        :param pulumi.Input[str] snapshot_id: A snapshot ID
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        A snapshot ID
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)


class SnapshotCreateVolumePermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Adds permission to create volumes off of a given EBS Snapshot.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ebs.Volume("example",
            availability_zone="us-west-2a",
            size=40)
        example_snapshot = aws.ebs.Snapshot("exampleSnapshot", volume_id=example.id)
        example_perm = aws.ec2.SnapshotCreateVolumePermission("examplePerm",
            snapshot_id=example_snapshot.id,
            account_id="12345678")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
        :param pulumi.Input[str] snapshot_id: A snapshot ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotCreateVolumePermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Adds permission to create volumes off of a given EBS Snapshot.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ebs.Volume("example",
            availability_zone="us-west-2a",
            size=40)
        example_snapshot = aws.ebs.Snapshot("exampleSnapshot", volume_id=example.id)
        example_perm = aws.ec2.SnapshotCreateVolumePermission("examplePerm",
            snapshot_id=example_snapshot.id,
            account_id="12345678")
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotCreateVolumePermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotCreateVolumePermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotCreateVolumePermissionArgs.__new__(SnapshotCreateVolumePermissionArgs)

            if account_id is None and not opts.urn:
                raise TypeError("Missing required property 'account_id'")
            __props__.__dict__["account_id"] = account_id
            if snapshot_id is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_id'")
            __props__.__dict__["snapshot_id"] = snapshot_id
        super(SnapshotCreateVolumePermission, __self__).__init__(
            'aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None) -> 'SnapshotCreateVolumePermission':
        """
        Get an existing SnapshotCreateVolumePermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
        :param pulumi.Input[str] snapshot_id: A snapshot ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotCreateVolumePermissionState.__new__(_SnapshotCreateVolumePermissionState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["snapshot_id"] = snapshot_id
        return SnapshotCreateVolumePermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[str]:
        """
        A snapshot ID
        """
        return pulumi.get(self, "snapshot_id")

