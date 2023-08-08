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

__all__ = ['FolderArgs', 'Folder']

@pulumi.input_type
class FolderArgs:
    def __init__(__self__, *,
                 folder_id: pulumi.Input[str],
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 folder_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_folder_arn: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Folder resource.
        :param pulumi.Input[str] folder_id: Identifier for the folder.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] folder_type: The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
        :param pulumi.Input[str] name: Display name for the folder.
               
               The following arguments are optional:
        :param pulumi.Input[str] parent_folder_arn: The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
        :param pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]] permissions: A set of resource permissions on the folder. Maximum of 64 items. See permissions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "folder_id", folder_id)
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if folder_type is not None:
            pulumi.set(__self__, "folder_type", folder_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_folder_arn is not None:
            pulumi.set(__self__, "parent_folder_arn", parent_folder_arn)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Input[str]:
        """
        Identifier for the folder.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="folderType")
    def folder_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
        """
        return pulumi.get(self, "folder_type")

    @folder_type.setter
    def folder_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the folder.

        The following arguments are optional:
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentFolderArn")
    def parent_folder_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
        """
        return pulumi.get(self, "parent_folder_arn")

    @parent_folder_arn.setter
    def parent_folder_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_folder_arn", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]]]:
        """
        A set of resource permissions on the folder. Maximum of 64 items. See permissions.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _FolderState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 created_time: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 folder_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 folder_type: Optional[pulumi.Input[str]] = None,
                 last_updated_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_folder_arn: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Folder resources.
        :param pulumi.Input[str] arn: ARN of the folder.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] created_time: The time that the folder was created.
        :param pulumi.Input[str] folder_id: Identifier for the folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] folder_paths: An array of ancestor ARN strings for the folder. Empty for root-level folders.
        :param pulumi.Input[str] folder_type: The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
        :param pulumi.Input[str] last_updated_time: The time that the folder was last updated.
        :param pulumi.Input[str] name: Display name for the folder.
               
               The following arguments are optional:
        :param pulumi.Input[str] parent_folder_arn: The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
        :param pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]] permissions: A set of resource permissions on the folder. Maximum of 64 items. See permissions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if folder_paths is not None:
            pulumi.set(__self__, "folder_paths", folder_paths)
        if folder_type is not None:
            pulumi.set(__self__, "folder_type", folder_type)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_folder_arn is not None:
            pulumi.set(__self__, "parent_folder_arn", parent_folder_arn)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the folder.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time that the folder was created.
        """
        return pulumi.get(self, "created_time")

    @created_time.setter
    def created_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_time", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the folder.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter(name="folderPaths")
    def folder_paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of ancestor ARN strings for the folder. Empty for root-level folders.
        """
        return pulumi.get(self, "folder_paths")

    @folder_paths.setter
    def folder_paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "folder_paths", value)

    @property
    @pulumi.getter(name="folderType")
    def folder_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
        """
        return pulumi.get(self, "folder_type")

    @folder_type.setter
    def folder_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_type", value)

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time that the folder was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @last_updated_time.setter
    def last_updated_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the folder.

        The following arguments are optional:
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentFolderArn")
    def parent_folder_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
        """
        return pulumi.get(self, "parent_folder_arn")

    @parent_folder_arn.setter
    def parent_folder_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_folder_arn", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]]]:
        """
        A set of resource permissions on the folder. Maximum of 64 items. See permissions.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FolderPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Folder(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 folder_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_folder_arn: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for managing a QuickSight Folder.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.quicksight.Folder("example", folder_id="example-id")
        ```
        ### With Permissions

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.quicksight.Folder("example",
            folder_id="example-id",
            permissions=[aws.quicksight.FolderPermissionArgs(
                actions=[
                    "quicksight:CreateFolder",
                    "quicksight:DescribeFolder",
                    "quicksight:UpdateFolder",
                    "quicksight:DeleteFolder",
                    "quicksight:CreateFolderMembership",
                    "quicksight:DeleteFolderMembership",
                    "quicksight:DescribeFolderPermissions",
                    "quicksight:UpdateFolderPermissions",
                ],
                principal=aws_quicksight_user["example"]["arn"],
            )])
        ```
        ### With Parent Folder

        ```python
        import pulumi
        import pulumi_aws as aws

        parent = aws.quicksight.Folder("parent", folder_id="parent-id")
        example = aws.quicksight.Folder("example",
            folder_id="example-id",
            parent_folder_arn=parent.arn)
        ```

        ## Import

        terraform import {

         to = aws_quicksight_folder.example

         id = "123456789012,example-id" } Using `pulumi import`, import a QuickSight folder using the AWS account ID and folder ID name separated by a comma (`,`). For exampleconsole % pulumi import aws_quicksight_folder.example 123456789012,example-id

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] folder_id: Identifier for the folder.
        :param pulumi.Input[str] folder_type: The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
        :param pulumi.Input[str] name: Display name for the folder.
               
               The following arguments are optional:
        :param pulumi.Input[str] parent_folder_arn: The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionArgs']]]] permissions: A set of resource permissions on the folder. Maximum of 64 items. See permissions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing a QuickSight Folder.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.quicksight.Folder("example", folder_id="example-id")
        ```
        ### With Permissions

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.quicksight.Folder("example",
            folder_id="example-id",
            permissions=[aws.quicksight.FolderPermissionArgs(
                actions=[
                    "quicksight:CreateFolder",
                    "quicksight:DescribeFolder",
                    "quicksight:UpdateFolder",
                    "quicksight:DeleteFolder",
                    "quicksight:CreateFolderMembership",
                    "quicksight:DeleteFolderMembership",
                    "quicksight:DescribeFolderPermissions",
                    "quicksight:UpdateFolderPermissions",
                ],
                principal=aws_quicksight_user["example"]["arn"],
            )])
        ```
        ### With Parent Folder

        ```python
        import pulumi
        import pulumi_aws as aws

        parent = aws.quicksight.Folder("parent", folder_id="parent-id")
        example = aws.quicksight.Folder("example",
            folder_id="example-id",
            parent_folder_arn=parent.arn)
        ```

        ## Import

        terraform import {

         to = aws_quicksight_folder.example

         id = "123456789012,example-id" } Using `pulumi import`, import a QuickSight folder using the AWS account ID and folder ID name separated by a comma (`,`). For exampleconsole % pulumi import aws_quicksight_folder.example 123456789012,example-id

        :param str resource_name: The name of the resource.
        :param FolderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 folder_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_folder_arn: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FolderArgs.__new__(FolderArgs)

            __props__.__dict__["aws_account_id"] = aws_account_id
            if folder_id is None and not opts.urn:
                raise TypeError("Missing required property 'folder_id'")
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["folder_type"] = folder_type
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_folder_arn"] = parent_folder_arn
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["folder_paths"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["tags_all"] = None
        super(Folder, __self__).__init__(
            'aws:quicksight/folder:Folder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_account_id: Optional[pulumi.Input[str]] = None,
            created_time: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            folder_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            folder_type: Optional[pulumi.Input[str]] = None,
            last_updated_time: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_folder_arn: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Folder':
        """
        Get an existing Folder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the folder.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] created_time: The time that the folder was created.
        :param pulumi.Input[str] folder_id: Identifier for the folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] folder_paths: An array of ancestor ARN strings for the folder. Empty for root-level folders.
        :param pulumi.Input[str] folder_type: The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
        :param pulumi.Input[str] last_updated_time: The time that the folder was last updated.
        :param pulumi.Input[str] name: Display name for the folder.
               
               The following arguments are optional:
        :param pulumi.Input[str] parent_folder_arn: The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderPermissionArgs']]]] permissions: A set of resource permissions on the folder. Maximum of 64 items. See permissions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FolderState.__new__(_FolderState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["aws_account_id"] = aws_account_id
        __props__.__dict__["created_time"] = created_time
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["folder_paths"] = folder_paths
        __props__.__dict__["folder_type"] = folder_type
        __props__.__dict__["last_updated_time"] = last_updated_time
        __props__.__dict__["name"] = name
        __props__.__dict__["parent_folder_arn"] = parent_folder_arn
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Folder(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the folder.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        """
        AWS account ID.
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        The time that the folder was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        Identifier for the folder.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="folderPaths")
    def folder_paths(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of ancestor ARN strings for the folder. Empty for root-level folders.
        """
        return pulumi.get(self, "folder_paths")

    @property
    @pulumi.getter(name="folderType")
    def folder_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
        """
        return pulumi.get(self, "folder_type")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        """
        The time that the folder was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Display name for the folder.

        The following arguments are optional:
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentFolderArn")
    def parent_folder_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
        """
        return pulumi.get(self, "parent_folder_arn")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.FolderPermission']]]:
        """
        A set of resource permissions on the folder. Maximum of 64 items. See permissions.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

