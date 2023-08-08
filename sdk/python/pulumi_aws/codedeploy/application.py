# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 compute_platform: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] compute_platform: The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        :param pulumi.Input[str] name: The name of the application.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if compute_platform is not None:
            pulumi.set(__self__, "compute_platform", compute_platform)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="computePlatform")
    def compute_platform(self) -> Optional[pulumi.Input[str]]:
        """
        The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        """
        return pulumi.get(self, "compute_platform")

    @compute_platform.setter
    def compute_platform(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_platform", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the application.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ApplicationState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 compute_platform: Optional[pulumi.Input[str]] = None,
                 github_account_name: Optional[pulumi.Input[str]] = None,
                 linked_to_github: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Application resources.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] arn: The ARN of the CodeDeploy application.
        :param pulumi.Input[str] compute_platform: The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        :param pulumi.Input[str] github_account_name: The name for a connection to a GitHub account.
        :param pulumi.Input[bool] linked_to_github: Whether the user has authenticated with GitHub for the specified application.
        :param pulumi.Input[str] name: The name of the application.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if compute_platform is not None:
            pulumi.set(__self__, "compute_platform", compute_platform)
        if github_account_name is not None:
            pulumi.set(__self__, "github_account_name", github_account_name)
        if linked_to_github is not None:
            pulumi.set(__self__, "linked_to_github", linked_to_github)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the CodeDeploy application.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="computePlatform")
    def compute_platform(self) -> Optional[pulumi.Input[str]]:
        """
        The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        """
        return pulumi.get(self, "compute_platform")

    @compute_platform.setter
    def compute_platform(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_platform", value)

    @property
    @pulumi.getter(name="githubAccountName")
    def github_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for a connection to a GitHub account.
        """
        return pulumi.get(self, "github_account_name")

    @github_account_name.setter
    def github_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "github_account_name", value)

    @property
    @pulumi.getter(name="linkedToGithub")
    def linked_to_github(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the user has authenticated with GitHub for the specified application.
        """
        return pulumi.get(self, "linked_to_github")

    @linked_to_github.setter
    def linked_to_github(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "linked_to_github", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the application.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_platform: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a CodeDeploy application to be used as a basis for deployments

        ## Example Usage
        ### ECS Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codedeploy.Application("example", compute_platform="ECS")
        ```
        ### Lambda Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codedeploy.Application("example", compute_platform="Lambda")
        ```
        ### Server Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codedeploy.Application("example", compute_platform="Server")
        ```

        ## Import

        terraform import {

         to = aws_codedeploy_app.example

         id = "my-application" } Using `pulumi import`, import CodeDeploy Applications using the `name`. For exampleconsole % pulumi import aws_codedeploy_app.example my-application

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_platform: The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        :param pulumi.Input[str] name: The name of the application.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ApplicationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CodeDeploy application to be used as a basis for deployments

        ## Example Usage
        ### ECS Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codedeploy.Application("example", compute_platform="ECS")
        ```
        ### Lambda Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codedeploy.Application("example", compute_platform="Lambda")
        ```
        ### Server Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codedeploy.Application("example", compute_platform="Server")
        ```

        ## Import

        terraform import {

         to = aws_codedeploy_app.example

         id = "my-application" } Using `pulumi import`, import CodeDeploy Applications using the `name`. For exampleconsole % pulumi import aws_codedeploy_app.example my-application

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_platform: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["compute_platform"] = compute_platform
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["application_id"] = None
            __props__.__dict__["arn"] = None
            __props__.__dict__["github_account_name"] = None
            __props__.__dict__["linked_to_github"] = None
            __props__.__dict__["tags_all"] = None
        super(Application, __self__).__init__(
            'aws:codedeploy/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            compute_platform: Optional[pulumi.Input[str]] = None,
            github_account_name: Optional[pulumi.Input[str]] = None,
            linked_to_github: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] arn: The ARN of the CodeDeploy application.
        :param pulumi.Input[str] compute_platform: The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        :param pulumi.Input[str] github_account_name: The name for a connection to a GitHub account.
        :param pulumi.Input[bool] linked_to_github: Whether the user has authenticated with GitHub for the specified application.
        :param pulumi.Input[str] name: The name of the application.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationState.__new__(_ApplicationState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["compute_platform"] = compute_platform
        __props__.__dict__["github_account_name"] = github_account_name
        __props__.__dict__["linked_to_github"] = linked_to_github
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the CodeDeploy application.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="computePlatform")
    def compute_platform(self) -> pulumi.Output[Optional[str]]:
        """
        The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        """
        return pulumi.get(self, "compute_platform")

    @property
    @pulumi.getter(name="githubAccountName")
    def github_account_name(self) -> pulumi.Output[str]:
        """
        The name for a connection to a GitHub account.
        """
        return pulumi.get(self, "github_account_name")

    @property
    @pulumi.getter(name="linkedToGithub")
    def linked_to_github(self) -> pulumi.Output[bool]:
        """
        Whether the user has authenticated with GitHub for the specified application.
        """
        return pulumi.get(self, "linked_to_github")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the application.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

