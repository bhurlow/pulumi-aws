# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RestApi(pulumi.CustomResource):
    api_key_source: pulumi.Output[str]
    """
    The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN)
    """
    binary_media_types: pulumi.Output[list]
    """
    The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
    """
    body: pulumi.Output[str]
    """
    An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
    """
    created_date: pulumi.Output[str]
    """
    The creation date of the REST API
    """
    description: pulumi.Output[str]
    """
    The description of the REST API
    """
    endpoint_configuration: pulumi.Output[dict]
    """
    Nested argument defining API endpoint configuration including endpoint type. Defined below.

      * `types` (`str`) - A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE`, `REGIONAL` or `PRIVATE`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
      * `vpcEndpointIds` (`list`) - A list of VPC Endpoint Ids. It is only supported for PRIVATE endpoint type.
    """
    execution_arn: pulumi.Output[str]
    """
    The execution ARN part to be used in `lambda_permission`'s `source_arn`
    when allowing API Gateway to invoke a Lambda function,
    e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
    """
    minimum_compression_size: pulumi.Output[float]
    """
    Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
    """
    name: pulumi.Output[str]
    """
    The name of the REST API
    """
    policy: pulumi.Output[str]
    """
    JSON formatted policy document that controls access to the API Gateway.
    """
    root_resource_id: pulumi.Output[str]
    """
    The resource ID of the REST API's root
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of resource tags
    """
    def __init__(__self__, resource_name, opts=None, api_key_source=None, binary_media_types=None, body=None, description=None, endpoint_configuration=None, minimum_compression_size=None, name=None, policy=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway REST API.

        > **Note:** Amazon API Gateway Version 1 resources are used for creating and deploying REST APIs. To create and deploy WebSocket and HTTP APIs, use Amazon API Gateway Version 2.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
        ```

        ### Regional Endpoint Type

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigateway.RestApi("example", endpoint_configuration={
            "types": "REGIONAL",
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key_source: The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
        :param pulumi.Input[list] binary_media_types: The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
        :param pulumi.Input[str] body: An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
        :param pulumi.Input[str] description: The description of the REST API
        :param pulumi.Input[dict] endpoint_configuration: Nested argument defining API endpoint configuration including endpoint type. Defined below.
        :param pulumi.Input[float] minimum_compression_size: Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
        :param pulumi.Input[str] name: The name of the REST API
        :param pulumi.Input[str] policy: JSON formatted policy document that controls access to the API Gateway.
        :param pulumi.Input[dict] tags: Key-value map of resource tags

        The **endpoint_configuration** object supports the following:

          * `types` (`pulumi.Input[str]`) - A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE`, `REGIONAL` or `PRIVATE`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
          * `vpcEndpointIds` (`pulumi.Input[list]`) - A list of VPC Endpoint Ids. It is only supported for PRIVATE endpoint type.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['api_key_source'] = api_key_source
            __props__['binary_media_types'] = binary_media_types
            __props__['body'] = body
            __props__['description'] = description
            __props__['endpoint_configuration'] = endpoint_configuration
            __props__['minimum_compression_size'] = minimum_compression_size
            __props__['name'] = name
            __props__['policy'] = policy
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['created_date'] = None
            __props__['execution_arn'] = None
            __props__['root_resource_id'] = None
        super(RestApi, __self__).__init__(
            'aws:apigateway/restApi:RestApi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_key_source=None, arn=None, binary_media_types=None, body=None, created_date=None, description=None, endpoint_configuration=None, execution_arn=None, minimum_compression_size=None, name=None, policy=None, root_resource_id=None, tags=None):
        """
        Get an existing RestApi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key_source: The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[list] binary_media_types: The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
        :param pulumi.Input[str] body: An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
        :param pulumi.Input[str] created_date: The creation date of the REST API
        :param pulumi.Input[str] description: The description of the REST API
        :param pulumi.Input[dict] endpoint_configuration: Nested argument defining API endpoint configuration including endpoint type. Defined below.
        :param pulumi.Input[str] execution_arn: The execution ARN part to be used in `lambda_permission`'s `source_arn`
               when allowing API Gateway to invoke a Lambda function,
               e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
        :param pulumi.Input[float] minimum_compression_size: Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
        :param pulumi.Input[str] name: The name of the REST API
        :param pulumi.Input[str] policy: JSON formatted policy document that controls access to the API Gateway.
        :param pulumi.Input[str] root_resource_id: The resource ID of the REST API's root
        :param pulumi.Input[dict] tags: Key-value map of resource tags

        The **endpoint_configuration** object supports the following:

          * `types` (`pulumi.Input[str]`) - A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE`, `REGIONAL` or `PRIVATE`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
          * `vpcEndpointIds` (`pulumi.Input[list]`) - A list of VPC Endpoint Ids. It is only supported for PRIVATE endpoint type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_key_source"] = api_key_source
        __props__["arn"] = arn
        __props__["binary_media_types"] = binary_media_types
        __props__["body"] = body
        __props__["created_date"] = created_date
        __props__["description"] = description
        __props__["endpoint_configuration"] = endpoint_configuration
        __props__["execution_arn"] = execution_arn
        __props__["minimum_compression_size"] = minimum_compression_size
        __props__["name"] = name
        __props__["policy"] = policy
        __props__["root_resource_id"] = root_resource_id
        __props__["tags"] = tags
        return RestApi(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

