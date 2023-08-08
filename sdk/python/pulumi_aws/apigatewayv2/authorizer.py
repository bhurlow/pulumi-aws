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

__all__ = ['AuthorizerArgs', 'Authorizer']

@pulumi.input_type
class AuthorizerArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 authorizer_type: pulumi.Input[str],
                 authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[bool]] = None,
                 identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 jwt_configuration: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Authorizer resource.
        :param pulumi.Input[str] api_id: API identifier.
        :param pulumi.Input[str] authorizer_type: Authorizer type. Valid values: `JWT`, `REQUEST`.
               Specify `REQUEST` for a Lambda function using incoming request parameters.
               For HTTP APIs, specify `JWT` to use JSON Web Tokens.
        :param pulumi.Input[str] authorizer_credentials_arn: Required credentials as an IAM role for API Gateway to invoke the authorizer.
               Supported only for `REQUEST` authorizers.
        :param pulumi.Input[str] authorizer_payload_format_version: Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
               Valid values: `1.0`, `2.0`.
        :param pulumi.Input[int] authorizer_result_ttl_in_seconds: Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
               If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
               Supported only for HTTP API Lambda authorizers.
        :param pulumi.Input[str] authorizer_uri: Authorizer's Uniform Resource Identifier (URI).
               For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `lambda.Function` resource.
               Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
        :param pulumi.Input[bool] enable_simple_responses: Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
               Supported only for HTTP APIs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] identity_sources: Identity sources for which authorization is requested.
               For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
               For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
        :param pulumi.Input['AuthorizerJwtConfigurationArgs'] jwt_configuration: Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
               Supported only for HTTP APIs.
        :param pulumi.Input[str] name: Name of the authorizer. Must be between 1 and 128 characters in length.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "authorizer_type", authorizer_type)
        if authorizer_credentials_arn is not None:
            pulumi.set(__self__, "authorizer_credentials_arn", authorizer_credentials_arn)
        if authorizer_payload_format_version is not None:
            pulumi.set(__self__, "authorizer_payload_format_version", authorizer_payload_format_version)
        if authorizer_result_ttl_in_seconds is not None:
            pulumi.set(__self__, "authorizer_result_ttl_in_seconds", authorizer_result_ttl_in_seconds)
        if authorizer_uri is not None:
            pulumi.set(__self__, "authorizer_uri", authorizer_uri)
        if enable_simple_responses is not None:
            pulumi.set(__self__, "enable_simple_responses", enable_simple_responses)
        if identity_sources is not None:
            pulumi.set(__self__, "identity_sources", identity_sources)
        if jwt_configuration is not None:
            pulumi.set(__self__, "jwt_configuration", jwt_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        API identifier.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> pulumi.Input[str]:
        """
        Authorizer type. Valid values: `JWT`, `REQUEST`.
        Specify `REQUEST` for a Lambda function using incoming request parameters.
        For HTTP APIs, specify `JWT` to use JSON Web Tokens.
        """
        return pulumi.get(self, "authorizer_type")

    @authorizer_type.setter
    def authorizer_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorizer_type", value)

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Required credentials as an IAM role for API Gateway to invoke the authorizer.
        Supported only for `REQUEST` authorizers.
        """
        return pulumi.get(self, "authorizer_credentials_arn")

    @authorizer_credentials_arn.setter
    def authorizer_credentials_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_credentials_arn", value)

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> Optional[pulumi.Input[str]]:
        """
        Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
        Valid values: `1.0`, `2.0`.
        """
        return pulumi.get(self, "authorizer_payload_format_version")

    @authorizer_payload_format_version.setter
    def authorizer_payload_format_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_payload_format_version", value)

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
        If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
        Supported only for HTTP API Lambda authorizers.
        """
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @authorizer_result_ttl_in_seconds.setter
    def authorizer_result_ttl_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "authorizer_result_ttl_in_seconds", value)

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Authorizer's Uniform Resource Identifier (URI).
        For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `lambda.Function` resource.
        Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
        """
        return pulumi.get(self, "authorizer_uri")

    @authorizer_uri.setter
    def authorizer_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_uri", value)

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
        Supported only for HTTP APIs.
        """
        return pulumi.get(self, "enable_simple_responses")

    @enable_simple_responses.setter
    def enable_simple_responses(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_simple_responses", value)

    @property
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Identity sources for which authorization is requested.
        For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
        For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
        """
        return pulumi.get(self, "identity_sources")

    @identity_sources.setter
    def identity_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "identity_sources", value)

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]:
        """
        Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
        Supported only for HTTP APIs.
        """
        return pulumi.get(self, "jwt_configuration")

    @jwt_configuration.setter
    def jwt_configuration(self, value: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]):
        pulumi.set(self, "jwt_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the authorizer. Must be between 1 and 128 characters in length.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AuthorizerState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
                 authorizer_type: Optional[pulumi.Input[str]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[bool]] = None,
                 identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 jwt_configuration: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Authorizer resources.
        :param pulumi.Input[str] api_id: API identifier.
        :param pulumi.Input[str] authorizer_credentials_arn: Required credentials as an IAM role for API Gateway to invoke the authorizer.
               Supported only for `REQUEST` authorizers.
        :param pulumi.Input[str] authorizer_payload_format_version: Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
               Valid values: `1.0`, `2.0`.
        :param pulumi.Input[int] authorizer_result_ttl_in_seconds: Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
               If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
               Supported only for HTTP API Lambda authorizers.
        :param pulumi.Input[str] authorizer_type: Authorizer type. Valid values: `JWT`, `REQUEST`.
               Specify `REQUEST` for a Lambda function using incoming request parameters.
               For HTTP APIs, specify `JWT` to use JSON Web Tokens.
        :param pulumi.Input[str] authorizer_uri: Authorizer's Uniform Resource Identifier (URI).
               For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `lambda.Function` resource.
               Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
        :param pulumi.Input[bool] enable_simple_responses: Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
               Supported only for HTTP APIs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] identity_sources: Identity sources for which authorization is requested.
               For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
               For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
        :param pulumi.Input['AuthorizerJwtConfigurationArgs'] jwt_configuration: Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
               Supported only for HTTP APIs.
        :param pulumi.Input[str] name: Name of the authorizer. Must be between 1 and 128 characters in length.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if authorizer_credentials_arn is not None:
            pulumi.set(__self__, "authorizer_credentials_arn", authorizer_credentials_arn)
        if authorizer_payload_format_version is not None:
            pulumi.set(__self__, "authorizer_payload_format_version", authorizer_payload_format_version)
        if authorizer_result_ttl_in_seconds is not None:
            pulumi.set(__self__, "authorizer_result_ttl_in_seconds", authorizer_result_ttl_in_seconds)
        if authorizer_type is not None:
            pulumi.set(__self__, "authorizer_type", authorizer_type)
        if authorizer_uri is not None:
            pulumi.set(__self__, "authorizer_uri", authorizer_uri)
        if enable_simple_responses is not None:
            pulumi.set(__self__, "enable_simple_responses", enable_simple_responses)
        if identity_sources is not None:
            pulumi.set(__self__, "identity_sources", identity_sources)
        if jwt_configuration is not None:
            pulumi.set(__self__, "jwt_configuration", jwt_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        API identifier.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Required credentials as an IAM role for API Gateway to invoke the authorizer.
        Supported only for `REQUEST` authorizers.
        """
        return pulumi.get(self, "authorizer_credentials_arn")

    @authorizer_credentials_arn.setter
    def authorizer_credentials_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_credentials_arn", value)

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> Optional[pulumi.Input[str]]:
        """
        Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
        Valid values: `1.0`, `2.0`.
        """
        return pulumi.get(self, "authorizer_payload_format_version")

    @authorizer_payload_format_version.setter
    def authorizer_payload_format_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_payload_format_version", value)

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
        If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
        Supported only for HTTP API Lambda authorizers.
        """
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @authorizer_result_ttl_in_seconds.setter
    def authorizer_result_ttl_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "authorizer_result_ttl_in_seconds", value)

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> Optional[pulumi.Input[str]]:
        """
        Authorizer type. Valid values: `JWT`, `REQUEST`.
        Specify `REQUEST` for a Lambda function using incoming request parameters.
        For HTTP APIs, specify `JWT` to use JSON Web Tokens.
        """
        return pulumi.get(self, "authorizer_type")

    @authorizer_type.setter
    def authorizer_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_type", value)

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Authorizer's Uniform Resource Identifier (URI).
        For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `lambda.Function` resource.
        Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
        """
        return pulumi.get(self, "authorizer_uri")

    @authorizer_uri.setter
    def authorizer_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_uri", value)

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
        Supported only for HTTP APIs.
        """
        return pulumi.get(self, "enable_simple_responses")

    @enable_simple_responses.setter
    def enable_simple_responses(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_simple_responses", value)

    @property
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Identity sources for which authorization is requested.
        For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
        For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
        """
        return pulumi.get(self, "identity_sources")

    @identity_sources.setter
    def identity_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "identity_sources", value)

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]:
        """
        Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
        Supported only for HTTP APIs.
        """
        return pulumi.get(self, "jwt_configuration")

    @jwt_configuration.setter
    def jwt_configuration(self, value: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]):
        pulumi.set(self, "jwt_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the authorizer. Must be between 1 and 128 characters in length.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Authorizer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
                 authorizer_type: Optional[pulumi.Input[str]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[bool]] = None,
                 identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 jwt_configuration: Optional[pulumi.Input[pulumi.InputType['AuthorizerJwtConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Amazon API Gateway Version 2 authorizer.
        More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

        ## Example Usage
        ### Basic WebSocket API

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Authorizer("example",
            api_id=aws_apigatewayv2_api["example"]["id"],
            authorizer_type="REQUEST",
            authorizer_uri=aws_lambda_function["example"]["invoke_arn"],
            identity_sources=["route.request.header.Auth"])
        ```
        ### Basic HTTP API

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Authorizer("example",
            api_id=aws_apigatewayv2_api["example"]["id"],
            authorizer_type="REQUEST",
            authorizer_uri=aws_lambda_function["example"]["invoke_arn"],
            identity_sources=["$request.header.Authorization"],
            authorizer_payload_format_version="2.0")
        ```

        ## Import

        terraform import {

         to = aws_apigatewayv2_authorizer.example

         id = "aabbccddee/1122334" } Using `pulumi import`, import `aws_apigatewayv2_authorizer` using the API identifier and authorizer identifier. For exampleconsole % pulumi import aws_apigatewayv2_authorizer.example aabbccddee/1122334

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: API identifier.
        :param pulumi.Input[str] authorizer_credentials_arn: Required credentials as an IAM role for API Gateway to invoke the authorizer.
               Supported only for `REQUEST` authorizers.
        :param pulumi.Input[str] authorizer_payload_format_version: Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
               Valid values: `1.0`, `2.0`.
        :param pulumi.Input[int] authorizer_result_ttl_in_seconds: Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
               If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
               Supported only for HTTP API Lambda authorizers.
        :param pulumi.Input[str] authorizer_type: Authorizer type. Valid values: `JWT`, `REQUEST`.
               Specify `REQUEST` for a Lambda function using incoming request parameters.
               For HTTP APIs, specify `JWT` to use JSON Web Tokens.
        :param pulumi.Input[str] authorizer_uri: Authorizer's Uniform Resource Identifier (URI).
               For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `lambda.Function` resource.
               Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
        :param pulumi.Input[bool] enable_simple_responses: Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
               Supported only for HTTP APIs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] identity_sources: Identity sources for which authorization is requested.
               For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
               For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
        :param pulumi.Input[pulumi.InputType['AuthorizerJwtConfigurationArgs']] jwt_configuration: Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
               Supported only for HTTP APIs.
        :param pulumi.Input[str] name: Name of the authorizer. Must be between 1 and 128 characters in length.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Amazon API Gateway Version 2 authorizer.
        More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

        ## Example Usage
        ### Basic WebSocket API

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Authorizer("example",
            api_id=aws_apigatewayv2_api["example"]["id"],
            authorizer_type="REQUEST",
            authorizer_uri=aws_lambda_function["example"]["invoke_arn"],
            identity_sources=["route.request.header.Auth"])
        ```
        ### Basic HTTP API

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Authorizer("example",
            api_id=aws_apigatewayv2_api["example"]["id"],
            authorizer_type="REQUEST",
            authorizer_uri=aws_lambda_function["example"]["invoke_arn"],
            identity_sources=["$request.header.Authorization"],
            authorizer_payload_format_version="2.0")
        ```

        ## Import

        terraform import {

         to = aws_apigatewayv2_authorizer.example

         id = "aabbccddee/1122334" } Using `pulumi import`, import `aws_apigatewayv2_authorizer` using the API identifier and authorizer identifier. For exampleconsole % pulumi import aws_apigatewayv2_authorizer.example aabbccddee/1122334

        :param str resource_name: The name of the resource.
        :param AuthorizerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
                 authorizer_type: Optional[pulumi.Input[str]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[bool]] = None,
                 identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 jwt_configuration: Optional[pulumi.Input[pulumi.InputType['AuthorizerJwtConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthorizerArgs.__new__(AuthorizerArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["authorizer_credentials_arn"] = authorizer_credentials_arn
            __props__.__dict__["authorizer_payload_format_version"] = authorizer_payload_format_version
            __props__.__dict__["authorizer_result_ttl_in_seconds"] = authorizer_result_ttl_in_seconds
            if authorizer_type is None and not opts.urn:
                raise TypeError("Missing required property 'authorizer_type'")
            __props__.__dict__["authorizer_type"] = authorizer_type
            __props__.__dict__["authorizer_uri"] = authorizer_uri
            __props__.__dict__["enable_simple_responses"] = enable_simple_responses
            __props__.__dict__["identity_sources"] = identity_sources
            __props__.__dict__["jwt_configuration"] = jwt_configuration
            __props__.__dict__["name"] = name
        super(Authorizer, __self__).__init__(
            'aws:apigatewayv2/authorizer:Authorizer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
            authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
            authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
            authorizer_type: Optional[pulumi.Input[str]] = None,
            authorizer_uri: Optional[pulumi.Input[str]] = None,
            enable_simple_responses: Optional[pulumi.Input[bool]] = None,
            identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            jwt_configuration: Optional[pulumi.Input[pulumi.InputType['AuthorizerJwtConfigurationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Authorizer':
        """
        Get an existing Authorizer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: API identifier.
        :param pulumi.Input[str] authorizer_credentials_arn: Required credentials as an IAM role for API Gateway to invoke the authorizer.
               Supported only for `REQUEST` authorizers.
        :param pulumi.Input[str] authorizer_payload_format_version: Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
               Valid values: `1.0`, `2.0`.
        :param pulumi.Input[int] authorizer_result_ttl_in_seconds: Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
               If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
               Supported only for HTTP API Lambda authorizers.
        :param pulumi.Input[str] authorizer_type: Authorizer type. Valid values: `JWT`, `REQUEST`.
               Specify `REQUEST` for a Lambda function using incoming request parameters.
               For HTTP APIs, specify `JWT` to use JSON Web Tokens.
        :param pulumi.Input[str] authorizer_uri: Authorizer's Uniform Resource Identifier (URI).
               For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `lambda.Function` resource.
               Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
        :param pulumi.Input[bool] enable_simple_responses: Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
               Supported only for HTTP APIs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] identity_sources: Identity sources for which authorization is requested.
               For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
               For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
        :param pulumi.Input[pulumi.InputType['AuthorizerJwtConfigurationArgs']] jwt_configuration: Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
               Supported only for HTTP APIs.
        :param pulumi.Input[str] name: Name of the authorizer. Must be between 1 and 128 characters in length.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthorizerState.__new__(_AuthorizerState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["authorizer_credentials_arn"] = authorizer_credentials_arn
        __props__.__dict__["authorizer_payload_format_version"] = authorizer_payload_format_version
        __props__.__dict__["authorizer_result_ttl_in_seconds"] = authorizer_result_ttl_in_seconds
        __props__.__dict__["authorizer_type"] = authorizer_type
        __props__.__dict__["authorizer_uri"] = authorizer_uri
        __props__.__dict__["enable_simple_responses"] = enable_simple_responses
        __props__.__dict__["identity_sources"] = identity_sources
        __props__.__dict__["jwt_configuration"] = jwt_configuration
        __props__.__dict__["name"] = name
        return Authorizer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        API identifier.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Required credentials as an IAM role for API Gateway to invoke the authorizer.
        Supported only for `REQUEST` authorizers.
        """
        return pulumi.get(self, "authorizer_credentials_arn")

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> pulumi.Output[Optional[str]]:
        """
        Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
        Valid values: `1.0`, `2.0`.
        """
        return pulumi.get(self, "authorizer_payload_format_version")

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> pulumi.Output[int]:
        """
        Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
        If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
        Supported only for HTTP API Lambda authorizers.
        """
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> pulumi.Output[str]:
        """
        Authorizer type. Valid values: `JWT`, `REQUEST`.
        Specify `REQUEST` for a Lambda function using incoming request parameters.
        For HTTP APIs, specify `JWT` to use JSON Web Tokens.
        """
        return pulumi.get(self, "authorizer_type")

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> pulumi.Output[Optional[str]]:
        """
        Authorizer's Uniform Resource Identifier (URI).
        For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invoke_arn` attribute of the `lambda.Function` resource.
        Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
        """
        return pulumi.get(self, "authorizer_uri")

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
        Supported only for HTTP APIs.
        """
        return pulumi.get(self, "enable_simple_responses")

    @property
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Identity sources for which authorization is requested.
        For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
        For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
        """
        return pulumi.get(self, "identity_sources")

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> pulumi.Output[Optional['outputs.AuthorizerJwtConfiguration']]:
        """
        Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
        Supported only for HTTP APIs.
        """
        return pulumi.get(self, "jwt_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the authorizer. Must be between 1 and 128 characters in length.
        """
        return pulumi.get(self, "name")

