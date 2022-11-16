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

__all__ = ['CodeRepositoryArgs', 'CodeRepository']

@pulumi.input_type
class CodeRepositoryArgs:
    def __init__(__self__, *,
                 code_repository_name: pulumi.Input[str],
                 git_config: pulumi.Input['CodeRepositoryGitConfigArgs'],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CodeRepository resource.
        :param pulumi.Input[str] code_repository_name: The name of the Code Repository (must be unique).
        :param pulumi.Input['CodeRepositoryGitConfigArgs'] git_config: Specifies details about the repository. see Git Config details below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "code_repository_name", code_repository_name)
        pulumi.set(__self__, "git_config", git_config)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="codeRepositoryName")
    def code_repository_name(self) -> pulumi.Input[str]:
        """
        The name of the Code Repository (must be unique).
        """
        return pulumi.get(self, "code_repository_name")

    @code_repository_name.setter
    def code_repository_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "code_repository_name", value)

    @property
    @pulumi.getter(name="gitConfig")
    def git_config(self) -> pulumi.Input['CodeRepositoryGitConfigArgs']:
        """
        Specifies details about the repository. see Git Config details below.
        """
        return pulumi.get(self, "git_config")

    @git_config.setter
    def git_config(self, value: pulumi.Input['CodeRepositoryGitConfigArgs']):
        pulumi.set(self, "git_config", value)

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
class _CodeRepositoryState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 code_repository_name: Optional[pulumi.Input[str]] = None,
                 git_config: Optional[pulumi.Input['CodeRepositoryGitConfigArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering CodeRepository resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Code Repository.
        :param pulumi.Input[str] code_repository_name: The name of the Code Repository (must be unique).
        :param pulumi.Input['CodeRepositoryGitConfigArgs'] git_config: Specifies details about the repository. see Git Config details below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if code_repository_name is not None:
            pulumi.set(__self__, "code_repository_name", code_repository_name)
        if git_config is not None:
            pulumi.set(__self__, "git_config", git_config)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Code Repository.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="codeRepositoryName")
    def code_repository_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Code Repository (must be unique).
        """
        return pulumi.get(self, "code_repository_name")

    @code_repository_name.setter
    def code_repository_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code_repository_name", value)

    @property
    @pulumi.getter(name="gitConfig")
    def git_config(self) -> Optional[pulumi.Input['CodeRepositoryGitConfigArgs']]:
        """
        Specifies details about the repository. see Git Config details below.
        """
        return pulumi.get(self, "git_config")

    @git_config.setter
    def git_config(self, value: Optional[pulumi.Input['CodeRepositoryGitConfigArgs']]):
        pulumi.set(self, "git_config", value)

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


class CodeRepository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 code_repository_name: Optional[pulumi.Input[str]] = None,
                 git_config: Optional[pulumi.Input[pulumi.InputType['CodeRepositoryGitConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a SageMaker Code Repository resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.CodeRepository("example",
            code_repository_name="example",
            git_config=aws.sagemaker.CodeRepositoryGitConfigArgs(
                repository_url="https://github.com/github/docs.git",
            ))
        ```
        ### Example with Secret

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example_secret = aws.secretsmanager.Secret("exampleSecret")
        example_secret_version = aws.secretsmanager.SecretVersion("exampleSecretVersion",
            secret_id=example_secret.id,
            secret_string=json.dumps({
                "username": "example",
                "password": "example",
            }))
        example_code_repository = aws.sagemaker.CodeRepository("exampleCodeRepository",
            code_repository_name="example",
            git_config=aws.sagemaker.CodeRepositoryGitConfigArgs(
                repository_url="https://github.com/github/docs.git",
                secret_arn=example_secret.arn,
            ),
            opts=pulumi.ResourceOptions(depends_on=[example_secret_version]))
        ```

        ## Import

        SageMaker Code Repositories can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:sagemaker/codeRepository:CodeRepository test_code_repository my-code-repo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] code_repository_name: The name of the Code Repository (must be unique).
        :param pulumi.Input[pulumi.InputType['CodeRepositoryGitConfigArgs']] git_config: Specifies details about the repository. see Git Config details below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CodeRepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a SageMaker Code Repository resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.CodeRepository("example",
            code_repository_name="example",
            git_config=aws.sagemaker.CodeRepositoryGitConfigArgs(
                repository_url="https://github.com/github/docs.git",
            ))
        ```
        ### Example with Secret

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example_secret = aws.secretsmanager.Secret("exampleSecret")
        example_secret_version = aws.secretsmanager.SecretVersion("exampleSecretVersion",
            secret_id=example_secret.id,
            secret_string=json.dumps({
                "username": "example",
                "password": "example",
            }))
        example_code_repository = aws.sagemaker.CodeRepository("exampleCodeRepository",
            code_repository_name="example",
            git_config=aws.sagemaker.CodeRepositoryGitConfigArgs(
                repository_url="https://github.com/github/docs.git",
                secret_arn=example_secret.arn,
            ),
            opts=pulumi.ResourceOptions(depends_on=[example_secret_version]))
        ```

        ## Import

        SageMaker Code Repositories can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:sagemaker/codeRepository:CodeRepository test_code_repository my-code-repo
        ```

        :param str resource_name: The name of the resource.
        :param CodeRepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CodeRepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 code_repository_name: Optional[pulumi.Input[str]] = None,
                 git_config: Optional[pulumi.Input[pulumi.InputType['CodeRepositoryGitConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CodeRepositoryArgs.__new__(CodeRepositoryArgs)

            if code_repository_name is None and not opts.urn:
                raise TypeError("Missing required property 'code_repository_name'")
            __props__.__dict__["code_repository_name"] = code_repository_name
            if git_config is None and not opts.urn:
                raise TypeError("Missing required property 'git_config'")
            __props__.__dict__["git_config"] = git_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(CodeRepository, __self__).__init__(
            'aws:sagemaker/codeRepository:CodeRepository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            code_repository_name: Optional[pulumi.Input[str]] = None,
            git_config: Optional[pulumi.Input[pulumi.InputType['CodeRepositoryGitConfigArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'CodeRepository':
        """
        Get an existing CodeRepository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Code Repository.
        :param pulumi.Input[str] code_repository_name: The name of the Code Repository (must be unique).
        :param pulumi.Input[pulumi.InputType['CodeRepositoryGitConfigArgs']] git_config: Specifies details about the repository. see Git Config details below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CodeRepositoryState.__new__(_CodeRepositoryState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["code_repository_name"] = code_repository_name
        __props__.__dict__["git_config"] = git_config
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return CodeRepository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Code Repository.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="codeRepositoryName")
    def code_repository_name(self) -> pulumi.Output[str]:
        """
        The name of the Code Repository (must be unique).
        """
        return pulumi.get(self, "code_repository_name")

    @property
    @pulumi.getter(name="gitConfig")
    def git_config(self) -> pulumi.Output['outputs.CodeRepositoryGitConfig']:
        """
        Specifies details about the repository. see Git Config details below.
        """
        return pulumi.get(self, "git_config")

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

