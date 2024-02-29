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

__all__ = [
    'GetUserGroupsResult',
    'AwaitableGetUserGroupsResult',
    'get_user_groups',
    'get_user_groups_output',
]

@pulumi.output_type
class GetUserGroupsResult:
    """
    A collection of values returned by getUserGroups.
    """
    def __init__(__self__, groups=None, id=None, user_pool_id=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if user_pool_id and not isinstance(user_pool_id, str):
            raise TypeError("Expected argument 'user_pool_id' to be a str")
        pulumi.set(__self__, "user_pool_id", user_pool_id)

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence['outputs.GetUserGroupsGroupResult']]:
        """
        List of groups. See `groups` below.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        User pool identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> str:
        return pulumi.get(self, "user_pool_id")


class AwaitableGetUserGroupsResult(GetUserGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserGroupsResult(
            groups=self.groups,
            id=self.id,
            user_pool_id=self.user_pool_id)


def get_user_groups(groups: Optional[Sequence[pulumi.InputType['GetUserGroupsGroupArgs']]] = None,
                    user_pool_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserGroupsResult:
    """
    Data source for managing AWS Cognito IDP (Identity Provider) User Groups.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.cognito.get_user_groups(user_pool_id="us-west-2_aaaaaaaaa")
    ```


    :param Sequence[pulumi.InputType['GetUserGroupsGroupArgs']] groups: List of groups. See `groups` below.
    :param str user_pool_id: User pool the client belongs to.
    """
    __args__ = dict()
    __args__['groups'] = groups
    __args__['userPoolId'] = user_pool_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:cognito/getUserGroups:getUserGroups', __args__, opts=opts, typ=GetUserGroupsResult).value

    return AwaitableGetUserGroupsResult(
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        user_pool_id=pulumi.get(__ret__, 'user_pool_id'))


@_utilities.lift_output_func(get_user_groups)
def get_user_groups_output(groups: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetUserGroupsGroupArgs']]]]] = None,
                           user_pool_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserGroupsResult]:
    """
    Data source for managing AWS Cognito IDP (Identity Provider) User Groups.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.cognito.get_user_groups(user_pool_id="us-west-2_aaaaaaaaa")
    ```


    :param Sequence[pulumi.InputType['GetUserGroupsGroupArgs']] groups: List of groups. See `groups` below.
    :param str user_pool_id: User pool the client belongs to.
    """
    ...
