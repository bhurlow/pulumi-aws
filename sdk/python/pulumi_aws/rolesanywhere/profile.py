# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 role_arns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 require_instance_properties: Optional[pulumi.Input[bool]] = None,
                 session_policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_arns: A list of IAM roles that this profile can assume
        :param pulumi.Input[int] duration_seconds: The number of seconds the vended session credentials are valid for. Defaults to 3600.
        :param pulumi.Input[bool] enabled: Whether or not the Profile is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] managed_policy_arns: A list of managed policy ARNs that apply to the vended session credentials.
        :param pulumi.Input[str] name: The name of the Profile.
        :param pulumi.Input[bool] require_instance_properties: Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        :param pulumi.Input[str] session_policy: A session policy that applies to the trust boundary of the vended session credentials.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "role_arns", role_arns)
        if duration_seconds is not None:
            pulumi.set(__self__, "duration_seconds", duration_seconds)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if managed_policy_arns is not None:
            pulumi.set(__self__, "managed_policy_arns", managed_policy_arns)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if require_instance_properties is not None:
            pulumi.set(__self__, "require_instance_properties", require_instance_properties)
        if session_policy is not None:
            pulumi.set(__self__, "session_policy", session_policy)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="roleArns")
    def role_arns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of IAM roles that this profile can assume
        """
        return pulumi.get(self, "role_arns")

    @role_arns.setter
    def role_arns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "role_arns", value)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds the vended session credentials are valid for. Defaults to 3600.
        """
        return pulumi.get(self, "duration_seconds")

    @duration_seconds.setter
    def duration_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_seconds", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the Profile is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of managed policy ARNs that apply to the vended session credentials.
        """
        return pulumi.get(self, "managed_policy_arns")

    @managed_policy_arns.setter
    def managed_policy_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "managed_policy_arns", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requireInstanceProperties")
    def require_instance_properties(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        """
        return pulumi.get(self, "require_instance_properties")

    @require_instance_properties.setter
    def require_instance_properties(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_instance_properties", value)

    @property
    @pulumi.getter(name="sessionPolicy")
    def session_policy(self) -> Optional[pulumi.Input[str]]:
        """
        A session policy that applies to the trust boundary of the vended session credentials.
        """
        return pulumi.get(self, "session_policy")

    @session_policy.setter
    def session_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_policy", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ProfileState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 require_instance_properties: Optional[pulumi.Input[bool]] = None,
                 role_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 session_policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Profile resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Profile
        :param pulumi.Input[int] duration_seconds: The number of seconds the vended session credentials are valid for. Defaults to 3600.
        :param pulumi.Input[bool] enabled: Whether or not the Profile is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] managed_policy_arns: A list of managed policy ARNs that apply to the vended session credentials.
        :param pulumi.Input[str] name: The name of the Profile.
        :param pulumi.Input[bool] require_instance_properties: Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_arns: A list of IAM roles that this profile can assume
        :param pulumi.Input[str] session_policy: A session policy that applies to the trust boundary of the vended session credentials.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if duration_seconds is not None:
            pulumi.set(__self__, "duration_seconds", duration_seconds)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if managed_policy_arns is not None:
            pulumi.set(__self__, "managed_policy_arns", managed_policy_arns)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if require_instance_properties is not None:
            pulumi.set(__self__, "require_instance_properties", require_instance_properties)
        if role_arns is not None:
            pulumi.set(__self__, "role_arns", role_arns)
        if session_policy is not None:
            pulumi.set(__self__, "session_policy", session_policy)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Profile
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds the vended session credentials are valid for. Defaults to 3600.
        """
        return pulumi.get(self, "duration_seconds")

    @duration_seconds.setter
    def duration_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_seconds", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the Profile is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of managed policy ARNs that apply to the vended session credentials.
        """
        return pulumi.get(self, "managed_policy_arns")

    @managed_policy_arns.setter
    def managed_policy_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "managed_policy_arns", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requireInstanceProperties")
    def require_instance_properties(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        """
        return pulumi.get(self, "require_instance_properties")

    @require_instance_properties.setter
    def require_instance_properties(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_instance_properties", value)

    @property
    @pulumi.getter(name="roleArns")
    def role_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of IAM roles that this profile can assume
        """
        return pulumi.get(self, "role_arns")

    @role_arns.setter
    def role_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_arns", value)

    @property
    @pulumi.getter(name="sessionPolicy")
    def session_policy(self) -> Optional[pulumi.Input[str]]:
        """
        A session policy that applies to the trust boundary of the vended session credentials.
        """
        return pulumi.get(self, "session_policy")

    @session_policy.setter
    def session_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_policy", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 require_instance_properties: Optional[pulumi.Input[bool]] = None,
                 role_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 session_policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for managing a Roles Anywhere Profile.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        test_role = aws.iam.Role("testRole",
            path="/",
            assume_role_policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Action": [
                        "sts:AssumeRole",
                        "sts:TagSession",
                        "sts:SetSourceIdentity",
                    ],
                    "Principal": {
                        "Service": "rolesanywhere.amazonaws.com",
                    },
                    "Effect": "Allow",
                    "Sid": "",
                }],
            }))
        test_profile = aws.rolesanywhere.Profile("testProfile", role_arns=[test_role.arn])
        ```

        ## Import

        `aws_rolesanywhere_profile` can be imported using its `id`, e.g.

        ```sh
         $ pulumi import aws:rolesanywhere/profile:Profile example db138a85-8925-4f9f-a409-08231233cacf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration_seconds: The number of seconds the vended session credentials are valid for. Defaults to 3600.
        :param pulumi.Input[bool] enabled: Whether or not the Profile is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] managed_policy_arns: A list of managed policy ARNs that apply to the vended session credentials.
        :param pulumi.Input[str] name: The name of the Profile.
        :param pulumi.Input[bool] require_instance_properties: Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_arns: A list of IAM roles that this profile can assume
        :param pulumi.Input[str] session_policy: A session policy that applies to the trust boundary of the vended session credentials.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing a Roles Anywhere Profile.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        test_role = aws.iam.Role("testRole",
            path="/",
            assume_role_policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Action": [
                        "sts:AssumeRole",
                        "sts:TagSession",
                        "sts:SetSourceIdentity",
                    ],
                    "Principal": {
                        "Service": "rolesanywhere.amazonaws.com",
                    },
                    "Effect": "Allow",
                    "Sid": "",
                }],
            }))
        test_profile = aws.rolesanywhere.Profile("testProfile", role_arns=[test_role.arn])
        ```

        ## Import

        `aws_rolesanywhere_profile` can be imported using its `id`, e.g.

        ```sh
         $ pulumi import aws:rolesanywhere/profile:Profile example db138a85-8925-4f9f-a409-08231233cacf
        ```

        :param str resource_name: The name of the resource.
        :param ProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration_seconds: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 require_instance_properties: Optional[pulumi.Input[bool]] = None,
                 role_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 session_policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            __props__.__dict__["duration_seconds"] = duration_seconds
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["managed_policy_arns"] = managed_policy_arns
            __props__.__dict__["name"] = name
            __props__.__dict__["require_instance_properties"] = require_instance_properties
            if role_arns is None and not opts.urn:
                raise TypeError("Missing required property 'role_arns'")
            __props__.__dict__["role_arns"] = role_arns
            __props__.__dict__["session_policy"] = session_policy
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(Profile, __self__).__init__(
            'aws:rolesanywhere/profile:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            duration_seconds: Optional[pulumi.Input[int]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            require_instance_properties: Optional[pulumi.Input[bool]] = None,
            role_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            session_policy: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Profile
        :param pulumi.Input[int] duration_seconds: The number of seconds the vended session credentials are valid for. Defaults to 3600.
        :param pulumi.Input[bool] enabled: Whether or not the Profile is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] managed_policy_arns: A list of managed policy ARNs that apply to the vended session credentials.
        :param pulumi.Input[str] name: The name of the Profile.
        :param pulumi.Input[bool] require_instance_properties: Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_arns: A list of IAM roles that this profile can assume
        :param pulumi.Input[str] session_policy: A session policy that applies to the trust boundary of the vended session credentials.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProfileState.__new__(_ProfileState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["duration_seconds"] = duration_seconds
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["managed_policy_arns"] = managed_policy_arns
        __props__.__dict__["name"] = name
        __props__.__dict__["require_instance_properties"] = require_instance_properties
        __props__.__dict__["role_arns"] = role_arns
        __props__.__dict__["session_policy"] = session_policy
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Profile
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> pulumi.Output[int]:
        """
        The number of seconds the vended session credentials are valid for. Defaults to 3600.
        """
        return pulumi.get(self, "duration_seconds")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not the Profile is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of managed policy ARNs that apply to the vended session credentials.
        """
        return pulumi.get(self, "managed_policy_arns")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Profile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="requireInstanceProperties")
    def require_instance_properties(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        """
        return pulumi.get(self, "require_instance_properties")

    @property
    @pulumi.getter(name="roleArns")
    def role_arns(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of IAM roles that this profile can assume
        """
        return pulumi.get(self, "role_arns")

    @property
    @pulumi.getter(name="sessionPolicy")
    def session_policy(self) -> pulumi.Output[Optional[str]]:
        """
        A session policy that applies to the trust boundary of the vended session credentials.
        """
        return pulumi.get(self, "session_policy")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

