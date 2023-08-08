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

__all__ = ['GraphQLApiArgs', 'GraphQLApi']

@pulumi.input_type
class GraphQLApiArgs:
    def __init__(__self__, *,
                 authentication_type: pulumi.Input[str],
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]] = None,
                 lambda_authorizer_config: Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']] = None,
                 log_config: Optional[pulumi.Input['GraphQLApiLogConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 openid_connect_config: Optional[pulumi.Input['GraphQLApiOpenidConnectConfigArgs']] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_pool_config: Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GraphQLApi resource.
        :param pulumi.Input[str] authentication_type: Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        :param pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs'] lambda_authorizer_config: Nested argument containing Lambda authorizer configuration. Defined below.
        :param pulumi.Input['GraphQLApiLogConfigArgs'] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: User-supplied name for the GraphqlApi.
        :param pulumi.Input['GraphQLApiOpenidConnectConfigArgs'] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input['GraphQLApiUserPoolConfigArgs'] user_pool_config: Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[str] visibility: Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.
        """
        pulumi.set(__self__, "authentication_type", authentication_type)
        if additional_authentication_providers is not None:
            pulumi.set(__self__, "additional_authentication_providers", additional_authentication_providers)
        if lambda_authorizer_config is not None:
            pulumi.set(__self__, "lambda_authorizer_config", lambda_authorizer_config)
        if log_config is not None:
            pulumi.set(__self__, "log_config", log_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if openid_connect_config is not None:
            pulumi.set(__self__, "openid_connect_config", openid_connect_config)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if user_pool_config is not None:
            pulumi.set(__self__, "user_pool_config", user_pool_config)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)
        if xray_enabled is not None:
            pulumi.set(__self__, "xray_enabled", xray_enabled)

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Input[str]:
        """
        Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        """
        return pulumi.get(self, "authentication_type")

    @authentication_type.setter
    def authentication_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authentication_type", value)

    @property
    @pulumi.getter(name="additionalAuthenticationProviders")
    def additional_authentication_providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]]:
        """
        One or more additional authentication providers for the GraphqlApi. Defined below.
        """
        return pulumi.get(self, "additional_authentication_providers")

    @additional_authentication_providers.setter
    def additional_authentication_providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]]):
        pulumi.set(self, "additional_authentication_providers", value)

    @property
    @pulumi.getter(name="lambdaAuthorizerConfig")
    def lambda_authorizer_config(self) -> Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']]:
        """
        Nested argument containing Lambda authorizer configuration. Defined below.
        """
        return pulumi.get(self, "lambda_authorizer_config")

    @lambda_authorizer_config.setter
    def lambda_authorizer_config(self, value: Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']]):
        pulumi.set(self, "lambda_authorizer_config", value)

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> Optional[pulumi.Input['GraphQLApiLogConfigArgs']]:
        """
        Nested argument containing logging configuration. Defined below.
        """
        return pulumi.get(self, "log_config")

    @log_config.setter
    def log_config(self, value: Optional[pulumi.Input['GraphQLApiLogConfigArgs']]):
        pulumi.set(self, "log_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-supplied name for the GraphqlApi.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openidConnectConfig")
    def openid_connect_config(self) -> Optional[pulumi.Input['GraphQLApiOpenidConnectConfigArgs']]:
        """
        Nested argument containing OpenID Connect configuration. Defined below.
        """
        return pulumi.get(self, "openid_connect_config")

    @openid_connect_config.setter
    def openid_connect_config(self, value: Optional[pulumi.Input['GraphQLApiOpenidConnectConfigArgs']]):
        pulumi.set(self, "openid_connect_config", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="userPoolConfig")
    def user_pool_config(self) -> Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']]:
        """
        Amazon Cognito User Pool configuration. Defined below.
        """
        return pulumi.get(self, "user_pool_config")

    @user_pool_config.setter
    def user_pool_config(self, value: Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']]):
        pulumi.set(self, "user_pool_config", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)

    @property
    @pulumi.getter(name="xrayEnabled")
    def xray_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether tracing with X-ray is enabled. Defaults to false.
        """
        return pulumi.get(self, "xray_enabled")

    @xray_enabled.setter
    def xray_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_enabled", value)


@pulumi.input_type
class _GraphQLApiState:
    def __init__(__self__, *,
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 authentication_type: Optional[pulumi.Input[str]] = None,
                 lambda_authorizer_config: Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']] = None,
                 log_config: Optional[pulumi.Input['GraphQLApiLogConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 openid_connect_config: Optional[pulumi.Input['GraphQLApiOpenidConnectConfigArgs']] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 uris: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_pool_config: Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering GraphQLApi resources.
        :param pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input[str] arn: ARN
        :param pulumi.Input[str] authentication_type: Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        :param pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs'] lambda_authorizer_config: Nested argument containing Lambda authorizer configuration. Defined below.
        :param pulumi.Input['GraphQLApiLogConfigArgs'] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: User-supplied name for the GraphqlApi.
        :param pulumi.Input['GraphQLApiOpenidConnectConfigArgs'] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] uris: Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        :param pulumi.Input['GraphQLApiUserPoolConfigArgs'] user_pool_config: Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[str] visibility: Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.
        """
        if additional_authentication_providers is not None:
            pulumi.set(__self__, "additional_authentication_providers", additional_authentication_providers)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if authentication_type is not None:
            pulumi.set(__self__, "authentication_type", authentication_type)
        if lambda_authorizer_config is not None:
            pulumi.set(__self__, "lambda_authorizer_config", lambda_authorizer_config)
        if log_config is not None:
            pulumi.set(__self__, "log_config", log_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if openid_connect_config is not None:
            pulumi.set(__self__, "openid_connect_config", openid_connect_config)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if uris is not None:
            pulumi.set(__self__, "uris", uris)
        if user_pool_config is not None:
            pulumi.set(__self__, "user_pool_config", user_pool_config)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)
        if xray_enabled is not None:
            pulumi.set(__self__, "xray_enabled", xray_enabled)

    @property
    @pulumi.getter(name="additionalAuthenticationProviders")
    def additional_authentication_providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]]:
        """
        One or more additional authentication providers for the GraphqlApi. Defined below.
        """
        return pulumi.get(self, "additional_authentication_providers")

    @additional_authentication_providers.setter
    def additional_authentication_providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]]):
        pulumi.set(self, "additional_authentication_providers", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        """
        return pulumi.get(self, "authentication_type")

    @authentication_type.setter
    def authentication_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_type", value)

    @property
    @pulumi.getter(name="lambdaAuthorizerConfig")
    def lambda_authorizer_config(self) -> Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']]:
        """
        Nested argument containing Lambda authorizer configuration. Defined below.
        """
        return pulumi.get(self, "lambda_authorizer_config")

    @lambda_authorizer_config.setter
    def lambda_authorizer_config(self, value: Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']]):
        pulumi.set(self, "lambda_authorizer_config", value)

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> Optional[pulumi.Input['GraphQLApiLogConfigArgs']]:
        """
        Nested argument containing logging configuration. Defined below.
        """
        return pulumi.get(self, "log_config")

    @log_config.setter
    def log_config(self, value: Optional[pulumi.Input['GraphQLApiLogConfigArgs']]):
        pulumi.set(self, "log_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-supplied name for the GraphqlApi.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openidConnectConfig")
    def openid_connect_config(self) -> Optional[pulumi.Input['GraphQLApiOpenidConnectConfigArgs']]:
        """
        Nested argument containing OpenID Connect configuration. Defined below.
        """
        return pulumi.get(self, "openid_connect_config")

    @openid_connect_config.setter
    def openid_connect_config(self, value: Optional[pulumi.Input['GraphQLApiOpenidConnectConfigArgs']]):
        pulumi.set(self, "openid_connect_config", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter
    def uris(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        """
        return pulumi.get(self, "uris")

    @uris.setter
    def uris(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "uris", value)

    @property
    @pulumi.getter(name="userPoolConfig")
    def user_pool_config(self) -> Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']]:
        """
        Amazon Cognito User Pool configuration. Defined below.
        """
        return pulumi.get(self, "user_pool_config")

    @user_pool_config.setter
    def user_pool_config(self, value: Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']]):
        pulumi.set(self, "user_pool_config", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)

    @property
    @pulumi.getter(name="xrayEnabled")
    def xray_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether tracing with X-ray is enabled. Defaults to false.
        """
        return pulumi.get(self, "xray_enabled")

    @xray_enabled.setter
    def xray_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_enabled", value)


class GraphQLApi(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]]] = None,
                 authentication_type: Optional[pulumi.Input[str]] = None,
                 lambda_authorizer_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLambdaAuthorizerConfigArgs']]] = None,
                 log_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 openid_connect_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_pool_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a GraphQLApi resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input[str] authentication_type: Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        :param pulumi.Input[pulumi.InputType['GraphQLApiLambdaAuthorizerConfigArgs']] lambda_authorizer_config: Nested argument containing Lambda authorizer configuration. Defined below.
        :param pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: User-supplied name for the GraphqlApi.
        :param pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']] user_pool_config: Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[str] visibility: Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GraphQLApiArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a GraphQLApi resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GraphQLApiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GraphQLApiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]]] = None,
                 authentication_type: Optional[pulumi.Input[str]] = None,
                 lambda_authorizer_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLambdaAuthorizerConfigArgs']]] = None,
                 log_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 openid_connect_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_pool_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GraphQLApiArgs.__new__(GraphQLApiArgs)

            __props__.__dict__["additional_authentication_providers"] = additional_authentication_providers
            if authentication_type is None and not opts.urn:
                raise TypeError("Missing required property 'authentication_type'")
            __props__.__dict__["authentication_type"] = authentication_type
            __props__.__dict__["lambda_authorizer_config"] = lambda_authorizer_config
            __props__.__dict__["log_config"] = log_config
            __props__.__dict__["name"] = name
            __props__.__dict__["openid_connect_config"] = openid_connect_config
            __props__.__dict__["schema"] = schema
            __props__.__dict__["tags"] = tags
            __props__.__dict__["user_pool_config"] = user_pool_config
            __props__.__dict__["visibility"] = visibility
            __props__.__dict__["xray_enabled"] = xray_enabled
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
            __props__.__dict__["uris"] = None
        super(GraphQLApi, __self__).__init__(
            'aws:appsync/graphQLApi:GraphQLApi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            authentication_type: Optional[pulumi.Input[str]] = None,
            lambda_authorizer_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLambdaAuthorizerConfigArgs']]] = None,
            log_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            openid_connect_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']]] = None,
            schema: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            uris: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            user_pool_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']]] = None,
            visibility: Optional[pulumi.Input[str]] = None,
            xray_enabled: Optional[pulumi.Input[bool]] = None) -> 'GraphQLApi':
        """
        Get an existing GraphQLApi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input[str] arn: ARN
        :param pulumi.Input[str] authentication_type: Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        :param pulumi.Input[pulumi.InputType['GraphQLApiLambdaAuthorizerConfigArgs']] lambda_authorizer_config: Nested argument containing Lambda authorizer configuration. Defined below.
        :param pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: User-supplied name for the GraphqlApi.
        :param pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] uris: Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        :param pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']] user_pool_config: Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[str] visibility: Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GraphQLApiState.__new__(_GraphQLApiState)

        __props__.__dict__["additional_authentication_providers"] = additional_authentication_providers
        __props__.__dict__["arn"] = arn
        __props__.__dict__["authentication_type"] = authentication_type
        __props__.__dict__["lambda_authorizer_config"] = lambda_authorizer_config
        __props__.__dict__["log_config"] = log_config
        __props__.__dict__["name"] = name
        __props__.__dict__["openid_connect_config"] = openid_connect_config
        __props__.__dict__["schema"] = schema
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["uris"] = uris
        __props__.__dict__["user_pool_config"] = user_pool_config
        __props__.__dict__["visibility"] = visibility
        __props__.__dict__["xray_enabled"] = xray_enabled
        return GraphQLApi(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalAuthenticationProviders")
    def additional_authentication_providers(self) -> pulumi.Output[Optional[Sequence['outputs.GraphQLApiAdditionalAuthenticationProvider']]]:
        """
        One or more additional authentication providers for the GraphqlApi. Defined below.
        """
        return pulumi.get(self, "additional_authentication_providers")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Output[str]:
        """
        Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        """
        return pulumi.get(self, "authentication_type")

    @property
    @pulumi.getter(name="lambdaAuthorizerConfig")
    def lambda_authorizer_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiLambdaAuthorizerConfig']]:
        """
        Nested argument containing Lambda authorizer configuration. Defined below.
        """
        return pulumi.get(self, "lambda_authorizer_config")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiLogConfig']]:
        """
        Nested argument containing logging configuration. Defined below.
        """
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User-supplied name for the GraphqlApi.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openidConnectConfig")
    def openid_connect_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiOpenidConnectConfig']]:
        """
        Nested argument containing OpenID Connect configuration. Defined below.
        """
        return pulumi.get(self, "openid_connect_config")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[Optional[str]]:
        """
        Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter
    def uris(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        """
        return pulumi.get(self, "uris")

    @property
    @pulumi.getter(name="userPoolConfig")
    def user_pool_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiUserPoolConfig']]:
        """
        Amazon Cognito User Pool configuration. Defined below.
        """
        return pulumi.get(self, "user_pool_config")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[Optional[str]]:
        """
        Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
        """
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter(name="xrayEnabled")
    def xray_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether tracing with X-ray is enabled. Defaults to false.
        """
        return pulumi.get(self, "xray_enabled")

