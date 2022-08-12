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

__all__ = [
    'GetUserHierarchyGroupResult',
    'AwaitableGetUserHierarchyGroupResult',
    'get_user_hierarchy_group',
    'get_user_hierarchy_group_output',
]

@pulumi.output_type
class GetUserHierarchyGroupResult:
    """
    A collection of values returned by getUserHierarchyGroup.
    """
    def __init__(__self__, arn=None, hierarchy_group_id=None, hierarchy_paths=None, id=None, instance_id=None, level_id=None, name=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if hierarchy_group_id and not isinstance(hierarchy_group_id, str):
            raise TypeError("Expected argument 'hierarchy_group_id' to be a str")
        pulumi.set(__self__, "hierarchy_group_id", hierarchy_group_id)
        if hierarchy_paths and not isinstance(hierarchy_paths, list):
            raise TypeError("Expected argument 'hierarchy_paths' to be a list")
        pulumi.set(__self__, "hierarchy_paths", hierarchy_paths)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if level_id and not isinstance(level_id, str):
            raise TypeError("Expected argument 'level_id' to be a str")
        pulumi.set(__self__, "level_id", level_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the hierarchy group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="hierarchyGroupId")
    def hierarchy_group_id(self) -> str:
        return pulumi.get(self, "hierarchy_group_id")

    @property
    @pulumi.getter(name="hierarchyPaths")
    def hierarchy_paths(self) -> Sequence['outputs.GetUserHierarchyGroupHierarchyPathResult']:
        """
        A block that contains information about the levels in the hierarchy group. The `hierarchy_path` block is documented below.
        """
        return pulumi.get(self, "hierarchy_paths")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="levelId")
    def level_id(self) -> str:
        """
        The identifier of the level in the hierarchy group.
        """
        return pulumi.get(self, "level_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the hierarchy group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags to assign to the hierarchy group.
        """
        return pulumi.get(self, "tags")


class AwaitableGetUserHierarchyGroupResult(GetUserHierarchyGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserHierarchyGroupResult(
            arn=self.arn,
            hierarchy_group_id=self.hierarchy_group_id,
            hierarchy_paths=self.hierarchy_paths,
            id=self.id,
            instance_id=self.instance_id,
            level_id=self.level_id,
            name=self.name,
            tags=self.tags)


def get_user_hierarchy_group(hierarchy_group_id: Optional[str] = None,
                             instance_id: Optional[str] = None,
                             name: Optional[str] = None,
                             tags: Optional[Mapping[str, str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserHierarchyGroupResult:
    """
    Provides details about a specific Amazon Connect User Hierarchy Group.

    ## Example Usage

    By `name`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user_hierarchy_group(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
        name="Example")
    ```

    By `hierarchy_group_id`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user_hierarchy_group(hierarchy_group_id="cccccccc-bbbb-cccc-dddd-111111111111",
        instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111")
    ```


    :param str hierarchy_group_id: Returns information on a specific hierarchy group by hierarchy group id
    :param str instance_id: Reference to the hosting Amazon Connect Instance
    :param str name: Returns information on a specific hierarchy group by name
    :param Mapping[str, str] tags: A map of tags to assign to the hierarchy group.
    """
    __args__ = dict()
    __args__['hierarchyGroupId'] = hierarchy_group_id
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:connect/getUserHierarchyGroup:getUserHierarchyGroup', __args__, opts=opts, typ=GetUserHierarchyGroupResult).value

    return AwaitableGetUserHierarchyGroupResult(
        arn=__ret__.arn,
        hierarchy_group_id=__ret__.hierarchy_group_id,
        hierarchy_paths=__ret__.hierarchy_paths,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        level_id=__ret__.level_id,
        name=__ret__.name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_user_hierarchy_group)
def get_user_hierarchy_group_output(hierarchy_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    instance_id: Optional[pulumi.Input[str]] = None,
                                    name: Optional[pulumi.Input[Optional[str]]] = None,
                                    tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserHierarchyGroupResult]:
    """
    Provides details about a specific Amazon Connect User Hierarchy Group.

    ## Example Usage

    By `name`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user_hierarchy_group(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
        name="Example")
    ```

    By `hierarchy_group_id`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user_hierarchy_group(hierarchy_group_id="cccccccc-bbbb-cccc-dddd-111111111111",
        instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111")
    ```


    :param str hierarchy_group_id: Returns information on a specific hierarchy group by hierarchy group id
    :param str instance_id: Reference to the hosting Amazon Connect Instance
    :param str name: Returns information on a specific hierarchy group by name
    :param Mapping[str, str] tags: A map of tags to assign to the hierarchy group.
    """
    ...