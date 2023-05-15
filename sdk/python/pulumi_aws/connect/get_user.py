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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, arn=None, directory_user_id=None, hierarchy_group_id=None, id=None, identity_infos=None, instance_id=None, name=None, phone_configs=None, routing_profile_id=None, security_profile_ids=None, tags=None, user_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if directory_user_id and not isinstance(directory_user_id, str):
            raise TypeError("Expected argument 'directory_user_id' to be a str")
        pulumi.set(__self__, "directory_user_id", directory_user_id)
        if hierarchy_group_id and not isinstance(hierarchy_group_id, str):
            raise TypeError("Expected argument 'hierarchy_group_id' to be a str")
        pulumi.set(__self__, "hierarchy_group_id", hierarchy_group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_infos and not isinstance(identity_infos, list):
            raise TypeError("Expected argument 'identity_infos' to be a list")
        pulumi.set(__self__, "identity_infos", identity_infos)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if phone_configs and not isinstance(phone_configs, list):
            raise TypeError("Expected argument 'phone_configs' to be a list")
        pulumi.set(__self__, "phone_configs", phone_configs)
        if routing_profile_id and not isinstance(routing_profile_id, str):
            raise TypeError("Expected argument 'routing_profile_id' to be a str")
        pulumi.set(__self__, "routing_profile_id", routing_profile_id)
        if security_profile_ids and not isinstance(security_profile_ids, list):
            raise TypeError("Expected argument 'security_profile_ids' to be a list")
        pulumi.set(__self__, "security_profile_ids", security_profile_ids)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the User.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="directoryUserId")
    def directory_user_id(self) -> str:
        """
        The identifier of the user account in the directory used for identity management.
        """
        return pulumi.get(self, "directory_user_id")

    @property
    @pulumi.getter(name="hierarchyGroupId")
    def hierarchy_group_id(self) -> str:
        """
        The identifier of the hierarchy group for the user.
        """
        return pulumi.get(self, "hierarchy_group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityInfos")
    def identity_infos(self) -> Sequence['outputs.GetUserIdentityInfoResult']:
        """
        A block that contains information about the identity of the user. Documented below.
        """
        return pulumi.get(self, "identity_infos")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        Specifies the identifier of the hosting Amazon Connect Instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="phoneConfigs")
    def phone_configs(self) -> Sequence['outputs.GetUserPhoneConfigResult']:
        """
        A block that contains information about the phone settings for the user. Documented below.
        """
        return pulumi.get(self, "phone_configs")

    @property
    @pulumi.getter(name="routingProfileId")
    def routing_profile_id(self) -> str:
        """
        The identifier of the routing profile for the user.
        """
        return pulumi.get(self, "routing_profile_id")

    @property
    @pulumi.getter(name="securityProfileIds")
    def security_profile_ids(self) -> Sequence[str]:
        """
        A list of identifiers for the security profiles for the user.
        """
        return pulumi.get(self, "security_profile_ids")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags to assign to the User.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        return pulumi.get(self, "user_id")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            arn=self.arn,
            directory_user_id=self.directory_user_id,
            hierarchy_group_id=self.hierarchy_group_id,
            id=self.id,
            identity_infos=self.identity_infos,
            instance_id=self.instance_id,
            name=self.name,
            phone_configs=self.phone_configs,
            routing_profile_id=self.routing_profile_id,
            security_profile_ids=self.security_profile_ids,
            tags=self.tags,
            user_id=self.user_id)


def get_user(instance_id: Optional[str] = None,
             name: Optional[str] = None,
             tags: Optional[Mapping[str, str]] = None,
             user_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Provides details about a specific Amazon Connect User.

    ## Example Usage

    By `name`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
        name="Example")
    ```

    By `user_id`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
        user_id="cccccccc-bbbb-cccc-dddd-111111111111")
    ```


    :param str instance_id: Reference to the hosting Amazon Connect Instance
    :param str name: Returns information on a specific User by name
    :param Mapping[str, str] tags: A map of tags to assign to the User.
    :param str user_id: Returns information on a specific User by User id
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['tags'] = tags
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:connect/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        arn=__ret__.arn,
        directory_user_id=__ret__.directory_user_id,
        hierarchy_group_id=__ret__.hierarchy_group_id,
        id=__ret__.id,
        identity_infos=__ret__.identity_infos,
        instance_id=__ret__.instance_id,
        name=__ret__.name,
        phone_configs=__ret__.phone_configs,
        routing_profile_id=__ret__.routing_profile_id,
        security_profile_ids=__ret__.security_profile_ids,
        tags=__ret__.tags,
        user_id=__ret__.user_id)


@_utilities.lift_output_func(get_user)
def get_user_output(instance_id: Optional[pulumi.Input[str]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                    user_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Provides details about a specific Amazon Connect User.

    ## Example Usage

    By `name`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
        name="Example")
    ```

    By `user_id`

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.connect.get_user(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
        user_id="cccccccc-bbbb-cccc-dddd-111111111111")
    ```


    :param str instance_id: Reference to the hosting Amazon Connect Instance
    :param str name: Returns information on a specific User by name
    :param Mapping[str, str] tags: A map of tags to assign to the User.
    :param str user_id: Returns information on a specific User by User id
    """
    ...
