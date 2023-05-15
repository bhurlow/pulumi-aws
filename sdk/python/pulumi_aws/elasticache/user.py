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

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 access_string: pulumi.Input[str],
                 engine: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 authentication_mode: Optional[pulumi.Input['UserAuthenticationModeArgs']] = None,
                 no_password_required: Optional[pulumi.Input[bool]] = None,
                 passwords: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] access_string: Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
        :param pulumi.Input[str] engine: The current supported value is `REDIS`.
        :param pulumi.Input[str] user_id: The ID of the user.
        :param pulumi.Input[str] user_name: The username of the user.
        :param pulumi.Input['UserAuthenticationModeArgs'] authentication_mode: Denotes the user's authentication properties. Detailed below.
        :param pulumi.Input[bool] no_password_required: Indicates a password is not required for this user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] passwords: Passwords used for this user. You can create up to two passwords for each user.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A list of tags to be added to this resource. A tag is a key-value pair.
        """
        pulumi.set(__self__, "access_string", access_string)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "user_name", user_name)
        if authentication_mode is not None:
            pulumi.set(__self__, "authentication_mode", authentication_mode)
        if no_password_required is not None:
            pulumi.set(__self__, "no_password_required", no_password_required)
        if passwords is not None:
            pulumi.set(__self__, "passwords", passwords)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessString")
    def access_string(self) -> pulumi.Input[str]:
        """
        Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
        """
        return pulumi.get(self, "access_string")

    @access_string.setter
    def access_string(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_string", value)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        The current supported value is `REDIS`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        The ID of the user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        The username of the user.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="authenticationMode")
    def authentication_mode(self) -> Optional[pulumi.Input['UserAuthenticationModeArgs']]:
        """
        Denotes the user's authentication properties. Detailed below.
        """
        return pulumi.get(self, "authentication_mode")

    @authentication_mode.setter
    def authentication_mode(self, value: Optional[pulumi.Input['UserAuthenticationModeArgs']]):
        pulumi.set(self, "authentication_mode", value)

    @property
    @pulumi.getter(name="noPasswordRequired")
    def no_password_required(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates a password is not required for this user.
        """
        return pulumi.get(self, "no_password_required")

    @no_password_required.setter
    def no_password_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_password_required", value)

    @property
    @pulumi.getter
    def passwords(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Passwords used for this user. You can create up to two passwords for each user.
        """
        return pulumi.get(self, "passwords")

    @passwords.setter
    def passwords(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "passwords", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A list of tags to be added to this resource. A tag is a key-value pair.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 access_string: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 authentication_mode: Optional[pulumi.Input['UserAuthenticationModeArgs']] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 no_password_required: Optional[pulumi.Input[bool]] = None,
                 passwords: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[str] access_string: Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
        :param pulumi.Input[str] arn: The ARN of the created ElastiCache User.
        :param pulumi.Input['UserAuthenticationModeArgs'] authentication_mode: Denotes the user's authentication properties. Detailed below.
        :param pulumi.Input[str] engine: The current supported value is `REDIS`.
        :param pulumi.Input[bool] no_password_required: Indicates a password is not required for this user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] passwords: Passwords used for this user. You can create up to two passwords for each user.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A list of tags to be added to this resource. A tag is a key-value pair.
        :param pulumi.Input[str] user_id: The ID of the user.
        :param pulumi.Input[str] user_name: The username of the user.
        """
        if access_string is not None:
            pulumi.set(__self__, "access_string", access_string)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if authentication_mode is not None:
            pulumi.set(__self__, "authentication_mode", authentication_mode)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if no_password_required is not None:
            pulumi.set(__self__, "no_password_required", no_password_required)
        if passwords is not None:
            pulumi.set(__self__, "passwords", passwords)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessString")
    def access_string(self) -> Optional[pulumi.Input[str]]:
        """
        Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
        """
        return pulumi.get(self, "access_string")

    @access_string.setter
    def access_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_string", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the created ElastiCache User.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="authenticationMode")
    def authentication_mode(self) -> Optional[pulumi.Input['UserAuthenticationModeArgs']]:
        """
        Denotes the user's authentication properties. Detailed below.
        """
        return pulumi.get(self, "authentication_mode")

    @authentication_mode.setter
    def authentication_mode(self, value: Optional[pulumi.Input['UserAuthenticationModeArgs']]):
        pulumi.set(self, "authentication_mode", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        The current supported value is `REDIS`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="noPasswordRequired")
    def no_password_required(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates a password is not required for this user.
        """
        return pulumi.get(self, "no_password_required")

    @no_password_required.setter
    def no_password_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_password_required", value)

    @property
    @pulumi.getter
    def passwords(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Passwords used for this user. You can create up to two passwords for each user.
        """
        return pulumi.get(self, "passwords")

    @passwords.setter
    def passwords(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "passwords", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A list of tags to be added to this resource. A tag is a key-value pair.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the user.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_string: Optional[pulumi.Input[str]] = None,
                 authentication_mode: Optional[pulumi.Input[pulumi.InputType['UserAuthenticationModeArgs']]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 no_password_required: Optional[pulumi.Input[bool]] = None,
                 passwords: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an ElastiCache user resource.

        > **Note:** All arguments including the username and passwords will be stored in the raw state as plain-text.
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.elasticache.User("test",
            access_string="on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
            engine="REDIS",
            passwords=["password123456789"],
            user_id="testUserId",
            user_name="testUserName")
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.elasticache.User("test",
            access_string="on ~* +@all",
            authentication_mode=aws.elasticache.UserAuthenticationModeArgs(
                type="iam",
            ),
            engine="REDIS",
            user_id="testUserId",
            user_name="testUserName")
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.elasticache.User("test",
            access_string="on ~* +@all",
            authentication_mode=aws.elasticache.UserAuthenticationModeArgs(
                passwords=[
                    "password1",
                    "password2",
                ],
                type="password",
            ),
            engine="REDIS",
            user_id="testUserId",
            user_name="testUserName")
        ```

        ## Import

        ElastiCache users can be imported using the `user_id`, e.g.,

        ```sh
         $ pulumi import aws:elasticache/user:User my_user userId1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_string: Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
        :param pulumi.Input[pulumi.InputType['UserAuthenticationModeArgs']] authentication_mode: Denotes the user's authentication properties. Detailed below.
        :param pulumi.Input[str] engine: The current supported value is `REDIS`.
        :param pulumi.Input[bool] no_password_required: Indicates a password is not required for this user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] passwords: Passwords used for this user. You can create up to two passwords for each user.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A list of tags to be added to this resource. A tag is a key-value pair.
        :param pulumi.Input[str] user_id: The ID of the user.
        :param pulumi.Input[str] user_name: The username of the user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an ElastiCache user resource.

        > **Note:** All arguments including the username and passwords will be stored in the raw state as plain-text.
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.elasticache.User("test",
            access_string="on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
            engine="REDIS",
            passwords=["password123456789"],
            user_id="testUserId",
            user_name="testUserName")
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.elasticache.User("test",
            access_string="on ~* +@all",
            authentication_mode=aws.elasticache.UserAuthenticationModeArgs(
                type="iam",
            ),
            engine="REDIS",
            user_id="testUserId",
            user_name="testUserName")
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.elasticache.User("test",
            access_string="on ~* +@all",
            authentication_mode=aws.elasticache.UserAuthenticationModeArgs(
                passwords=[
                    "password1",
                    "password2",
                ],
                type="password",
            ),
            engine="REDIS",
            user_id="testUserId",
            user_name="testUserName")
        ```

        ## Import

        ElastiCache users can be imported using the `user_id`, e.g.,

        ```sh
         $ pulumi import aws:elasticache/user:User my_user userId1
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_string: Optional[pulumi.Input[str]] = None,
                 authentication_mode: Optional[pulumi.Input[pulumi.InputType['UserAuthenticationModeArgs']]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 no_password_required: Optional[pulumi.Input[bool]] = None,
                 passwords: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            if access_string is None and not opts.urn:
                raise TypeError("Missing required property 'access_string'")
            __props__.__dict__["access_string"] = access_string
            __props__.__dict__["authentication_mode"] = authentication_mode
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            __props__.__dict__["no_password_required"] = no_password_required
            __props__.__dict__["passwords"] = None if passwords is None else pulumi.Output.secret(passwords)
            __props__.__dict__["tags"] = tags
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["passwords"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(User, __self__).__init__(
            'aws:elasticache/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_string: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            authentication_mode: Optional[pulumi.Input[pulumi.InputType['UserAuthenticationModeArgs']]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            no_password_required: Optional[pulumi.Input[bool]] = None,
            passwords: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            user_id: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_string: Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
        :param pulumi.Input[str] arn: The ARN of the created ElastiCache User.
        :param pulumi.Input[pulumi.InputType['UserAuthenticationModeArgs']] authentication_mode: Denotes the user's authentication properties. Detailed below.
        :param pulumi.Input[str] engine: The current supported value is `REDIS`.
        :param pulumi.Input[bool] no_password_required: Indicates a password is not required for this user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] passwords: Passwords used for this user. You can create up to two passwords for each user.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A list of tags to be added to this resource. A tag is a key-value pair.
        :param pulumi.Input[str] user_id: The ID of the user.
        :param pulumi.Input[str] user_name: The username of the user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["access_string"] = access_string
        __props__.__dict__["arn"] = arn
        __props__.__dict__["authentication_mode"] = authentication_mode
        __props__.__dict__["engine"] = engine
        __props__.__dict__["no_password_required"] = no_password_required
        __props__.__dict__["passwords"] = passwords
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["user_id"] = user_id
        __props__.__dict__["user_name"] = user_name
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessString")
    def access_string(self) -> pulumi.Output[str]:
        """
        Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
        """
        return pulumi.get(self, "access_string")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the created ElastiCache User.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationMode")
    def authentication_mode(self) -> pulumi.Output['outputs.UserAuthenticationMode']:
        """
        Denotes the user's authentication properties. Detailed below.
        """
        return pulumi.get(self, "authentication_mode")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        The current supported value is `REDIS`.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="noPasswordRequired")
    def no_password_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates a password is not required for this user.
        """
        return pulumi.get(self, "no_password_required")

    @property
    @pulumi.getter
    def passwords(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Passwords used for this user. You can create up to two passwords for each user.
        """
        return pulumi.get(self, "passwords")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A list of tags to be added to this resource. A tag is a key-value pair.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The ID of the user.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        The username of the user.
        """
        return pulumi.get(self, "user_name")

