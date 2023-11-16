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

__all__ = ['DataSourceArgs', 'DataSource']

@pulumi.input_type
class DataSourceArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 dynamodb_config: Optional[pulumi.Input['DataSourceDynamodbConfigArgs']] = None,
                 elasticsearch_config: Optional[pulumi.Input['DataSourceElasticsearchConfigArgs']] = None,
                 event_bridge_config: Optional[pulumi.Input['DataSourceEventBridgeConfigArgs']] = None,
                 http_config: Optional[pulumi.Input['DataSourceHttpConfigArgs']] = None,
                 lambda_config: Optional[pulumi.Input['DataSourceLambdaConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 opensearchservice_config: Optional[pulumi.Input['DataSourceOpensearchserviceConfigArgs']] = None,
                 relational_database_config: Optional[pulumi.Input['DataSourceRelationalDatabaseConfigArgs']] = None,
                 service_role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataSource resource.
        :param pulumi.Input[str] api_id: API ID for the GraphQL API for the data source.
        :param pulumi.Input[str] type: Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        :param pulumi.Input[str] description: Description of the data source.
        :param pulumi.Input['DataSourceDynamodbConfigArgs'] dynamodb_config: DynamoDB settings. See DynamoDB Config
        :param pulumi.Input['DataSourceElasticsearchConfigArgs'] elasticsearch_config: Amazon Elasticsearch settings. See ElasticSearch Config
        :param pulumi.Input['DataSourceEventBridgeConfigArgs'] event_bridge_config: AWS EventBridge settings. See Event Bridge Config
        :param pulumi.Input['DataSourceHttpConfigArgs'] http_config: HTTP settings. See HTTP Config
        :param pulumi.Input['DataSourceLambdaConfigArgs'] lambda_config: AWS Lambda settings. See Lambda Config
        :param pulumi.Input[str] name: User-supplied name for the data source.
        :param pulumi.Input['DataSourceOpensearchserviceConfigArgs'] opensearchservice_config: Amazon OpenSearch Service settings. See OpenSearch Service Config
        :param pulumi.Input['DataSourceRelationalDatabaseConfigArgs'] relational_database_config: AWS RDS settings. See Relational Database Config
        :param pulumi.Input[str] service_role_arn: IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamodb_config is not None:
            pulumi.set(__self__, "dynamodb_config", dynamodb_config)
        if elasticsearch_config is not None:
            pulumi.set(__self__, "elasticsearch_config", elasticsearch_config)
        if event_bridge_config is not None:
            pulumi.set(__self__, "event_bridge_config", event_bridge_config)
        if http_config is not None:
            pulumi.set(__self__, "http_config", http_config)
        if lambda_config is not None:
            pulumi.set(__self__, "lambda_config", lambda_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if opensearchservice_config is not None:
            pulumi.set(__self__, "opensearchservice_config", opensearchservice_config)
        if relational_database_config is not None:
            pulumi.set(__self__, "relational_database_config", relational_database_config)
        if service_role_arn is not None:
            pulumi.set(__self__, "service_role_arn", service_role_arn)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        API ID for the GraphQL API for the data source.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the data source.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dynamodbConfig")
    def dynamodb_config(self) -> Optional[pulumi.Input['DataSourceDynamodbConfigArgs']]:
        """
        DynamoDB settings. See DynamoDB Config
        """
        return pulumi.get(self, "dynamodb_config")

    @dynamodb_config.setter
    def dynamodb_config(self, value: Optional[pulumi.Input['DataSourceDynamodbConfigArgs']]):
        pulumi.set(self, "dynamodb_config", value)

    @property
    @pulumi.getter(name="elasticsearchConfig")
    def elasticsearch_config(self) -> Optional[pulumi.Input['DataSourceElasticsearchConfigArgs']]:
        """
        Amazon Elasticsearch settings. See ElasticSearch Config
        """
        return pulumi.get(self, "elasticsearch_config")

    @elasticsearch_config.setter
    def elasticsearch_config(self, value: Optional[pulumi.Input['DataSourceElasticsearchConfigArgs']]):
        pulumi.set(self, "elasticsearch_config", value)

    @property
    @pulumi.getter(name="eventBridgeConfig")
    def event_bridge_config(self) -> Optional[pulumi.Input['DataSourceEventBridgeConfigArgs']]:
        """
        AWS EventBridge settings. See Event Bridge Config
        """
        return pulumi.get(self, "event_bridge_config")

    @event_bridge_config.setter
    def event_bridge_config(self, value: Optional[pulumi.Input['DataSourceEventBridgeConfigArgs']]):
        pulumi.set(self, "event_bridge_config", value)

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> Optional[pulumi.Input['DataSourceHttpConfigArgs']]:
        """
        HTTP settings. See HTTP Config
        """
        return pulumi.get(self, "http_config")

    @http_config.setter
    def http_config(self, value: Optional[pulumi.Input['DataSourceHttpConfigArgs']]):
        pulumi.set(self, "http_config", value)

    @property
    @pulumi.getter(name="lambdaConfig")
    def lambda_config(self) -> Optional[pulumi.Input['DataSourceLambdaConfigArgs']]:
        """
        AWS Lambda settings. See Lambda Config
        """
        return pulumi.get(self, "lambda_config")

    @lambda_config.setter
    def lambda_config(self, value: Optional[pulumi.Input['DataSourceLambdaConfigArgs']]):
        pulumi.set(self, "lambda_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-supplied name for the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="opensearchserviceConfig")
    def opensearchservice_config(self) -> Optional[pulumi.Input['DataSourceOpensearchserviceConfigArgs']]:
        """
        Amazon OpenSearch Service settings. See OpenSearch Service Config
        """
        return pulumi.get(self, "opensearchservice_config")

    @opensearchservice_config.setter
    def opensearchservice_config(self, value: Optional[pulumi.Input['DataSourceOpensearchserviceConfigArgs']]):
        pulumi.set(self, "opensearchservice_config", value)

    @property
    @pulumi.getter(name="relationalDatabaseConfig")
    def relational_database_config(self) -> Optional[pulumi.Input['DataSourceRelationalDatabaseConfigArgs']]:
        """
        AWS RDS settings. See Relational Database Config
        """
        return pulumi.get(self, "relational_database_config")

    @relational_database_config.setter
    def relational_database_config(self, value: Optional[pulumi.Input['DataSourceRelationalDatabaseConfigArgs']]):
        pulumi.set(self, "relational_database_config", value)

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        """
        return pulumi.get(self, "service_role_arn")

    @service_role_arn.setter
    def service_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_role_arn", value)


@pulumi.input_type
class _DataSourceState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamodb_config: Optional[pulumi.Input['DataSourceDynamodbConfigArgs']] = None,
                 elasticsearch_config: Optional[pulumi.Input['DataSourceElasticsearchConfigArgs']] = None,
                 event_bridge_config: Optional[pulumi.Input['DataSourceEventBridgeConfigArgs']] = None,
                 http_config: Optional[pulumi.Input['DataSourceHttpConfigArgs']] = None,
                 lambda_config: Optional[pulumi.Input['DataSourceLambdaConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 opensearchservice_config: Optional[pulumi.Input['DataSourceOpensearchserviceConfigArgs']] = None,
                 relational_database_config: Optional[pulumi.Input['DataSourceRelationalDatabaseConfigArgs']] = None,
                 service_role_arn: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataSource resources.
        :param pulumi.Input[str] api_id: API ID for the GraphQL API for the data source.
        :param pulumi.Input[str] arn: ARN
        :param pulumi.Input[str] description: Description of the data source.
        :param pulumi.Input['DataSourceDynamodbConfigArgs'] dynamodb_config: DynamoDB settings. See DynamoDB Config
        :param pulumi.Input['DataSourceElasticsearchConfigArgs'] elasticsearch_config: Amazon Elasticsearch settings. See ElasticSearch Config
        :param pulumi.Input['DataSourceEventBridgeConfigArgs'] event_bridge_config: AWS EventBridge settings. See Event Bridge Config
        :param pulumi.Input['DataSourceHttpConfigArgs'] http_config: HTTP settings. See HTTP Config
        :param pulumi.Input['DataSourceLambdaConfigArgs'] lambda_config: AWS Lambda settings. See Lambda Config
        :param pulumi.Input[str] name: User-supplied name for the data source.
        :param pulumi.Input['DataSourceOpensearchserviceConfigArgs'] opensearchservice_config: Amazon OpenSearch Service settings. See OpenSearch Service Config
        :param pulumi.Input['DataSourceRelationalDatabaseConfigArgs'] relational_database_config: AWS RDS settings. See Relational Database Config
        :param pulumi.Input[str] service_role_arn: IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        :param pulumi.Input[str] type: Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamodb_config is not None:
            pulumi.set(__self__, "dynamodb_config", dynamodb_config)
        if elasticsearch_config is not None:
            pulumi.set(__self__, "elasticsearch_config", elasticsearch_config)
        if event_bridge_config is not None:
            pulumi.set(__self__, "event_bridge_config", event_bridge_config)
        if http_config is not None:
            pulumi.set(__self__, "http_config", http_config)
        if lambda_config is not None:
            pulumi.set(__self__, "lambda_config", lambda_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if opensearchservice_config is not None:
            pulumi.set(__self__, "opensearchservice_config", opensearchservice_config)
        if relational_database_config is not None:
            pulumi.set(__self__, "relational_database_config", relational_database_config)
        if service_role_arn is not None:
            pulumi.set(__self__, "service_role_arn", service_role_arn)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        API ID for the GraphQL API for the data source.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the data source.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dynamodbConfig")
    def dynamodb_config(self) -> Optional[pulumi.Input['DataSourceDynamodbConfigArgs']]:
        """
        DynamoDB settings. See DynamoDB Config
        """
        return pulumi.get(self, "dynamodb_config")

    @dynamodb_config.setter
    def dynamodb_config(self, value: Optional[pulumi.Input['DataSourceDynamodbConfigArgs']]):
        pulumi.set(self, "dynamodb_config", value)

    @property
    @pulumi.getter(name="elasticsearchConfig")
    def elasticsearch_config(self) -> Optional[pulumi.Input['DataSourceElasticsearchConfigArgs']]:
        """
        Amazon Elasticsearch settings. See ElasticSearch Config
        """
        return pulumi.get(self, "elasticsearch_config")

    @elasticsearch_config.setter
    def elasticsearch_config(self, value: Optional[pulumi.Input['DataSourceElasticsearchConfigArgs']]):
        pulumi.set(self, "elasticsearch_config", value)

    @property
    @pulumi.getter(name="eventBridgeConfig")
    def event_bridge_config(self) -> Optional[pulumi.Input['DataSourceEventBridgeConfigArgs']]:
        """
        AWS EventBridge settings. See Event Bridge Config
        """
        return pulumi.get(self, "event_bridge_config")

    @event_bridge_config.setter
    def event_bridge_config(self, value: Optional[pulumi.Input['DataSourceEventBridgeConfigArgs']]):
        pulumi.set(self, "event_bridge_config", value)

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> Optional[pulumi.Input['DataSourceHttpConfigArgs']]:
        """
        HTTP settings. See HTTP Config
        """
        return pulumi.get(self, "http_config")

    @http_config.setter
    def http_config(self, value: Optional[pulumi.Input['DataSourceHttpConfigArgs']]):
        pulumi.set(self, "http_config", value)

    @property
    @pulumi.getter(name="lambdaConfig")
    def lambda_config(self) -> Optional[pulumi.Input['DataSourceLambdaConfigArgs']]:
        """
        AWS Lambda settings. See Lambda Config
        """
        return pulumi.get(self, "lambda_config")

    @lambda_config.setter
    def lambda_config(self, value: Optional[pulumi.Input['DataSourceLambdaConfigArgs']]):
        pulumi.set(self, "lambda_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-supplied name for the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="opensearchserviceConfig")
    def opensearchservice_config(self) -> Optional[pulumi.Input['DataSourceOpensearchserviceConfigArgs']]:
        """
        Amazon OpenSearch Service settings. See OpenSearch Service Config
        """
        return pulumi.get(self, "opensearchservice_config")

    @opensearchservice_config.setter
    def opensearchservice_config(self, value: Optional[pulumi.Input['DataSourceOpensearchserviceConfigArgs']]):
        pulumi.set(self, "opensearchservice_config", value)

    @property
    @pulumi.getter(name="relationalDatabaseConfig")
    def relational_database_config(self) -> Optional[pulumi.Input['DataSourceRelationalDatabaseConfigArgs']]:
        """
        AWS RDS settings. See Relational Database Config
        """
        return pulumi.get(self, "relational_database_config")

    @relational_database_config.setter
    def relational_database_config(self, value: Optional[pulumi.Input['DataSourceRelationalDatabaseConfigArgs']]):
        pulumi.set(self, "relational_database_config", value)

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        """
        return pulumi.get(self, "service_role_arn")

    @service_role_arn.setter
    def service_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_role_arn", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class DataSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamodb_config: Optional[pulumi.Input[pulumi.InputType['DataSourceDynamodbConfigArgs']]] = None,
                 elasticsearch_config: Optional[pulumi.Input[pulumi.InputType['DataSourceElasticsearchConfigArgs']]] = None,
                 event_bridge_config: Optional[pulumi.Input[pulumi.InputType['DataSourceEventBridgeConfigArgs']]] = None,
                 http_config: Optional[pulumi.Input[pulumi.InputType['DataSourceHttpConfigArgs']]] = None,
                 lambda_config: Optional[pulumi.Input[pulumi.InputType['DataSourceLambdaConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 opensearchservice_config: Optional[pulumi.Input[pulumi.InputType['DataSourceOpensearchserviceConfigArgs']]] = None,
                 relational_database_config: Optional[pulumi.Input[pulumi.InputType['DataSourceRelationalDatabaseConfigArgs']]] = None,
                 service_role_arn: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AppSync Data Source.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_table = aws.dynamodb.Table("exampleTable",
            read_capacity=1,
            write_capacity=1,
            hash_key="UserId",
            attributes=[aws.dynamodb.TableAttributeArgs(
                name="UserId",
                type="S",
            )])
        assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="Service",
                identifiers=["appsync.amazonaws.com"],
            )],
            actions=["sts:AssumeRole"],
        )])
        example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
        example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            actions=["dynamodb:*"],
            resources=[example_table.arn],
        )])
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=example_policy_document.json)
        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
        example_data_source = aws.appsync.DataSource("exampleDataSource",
            api_id=example_graph_ql_api.id,
            name="my_appsync_example",
            service_role_arn=example_role.arn,
            type="AMAZON_DYNAMODB",
            dynamodb_config=aws.appsync.DataSourceDynamodbConfigArgs(
                table_name=example_table.name,
            ))
        ```

        ## Import

        Using `pulumi import`, import `aws_appsync_datasource` using the `api_id`, a hyphen, and `name`. For example:

        ```sh
         $ pulumi import aws:appsync/dataSource:DataSource example abcdef123456-example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: API ID for the GraphQL API for the data source.
        :param pulumi.Input[str] description: Description of the data source.
        :param pulumi.Input[pulumi.InputType['DataSourceDynamodbConfigArgs']] dynamodb_config: DynamoDB settings. See DynamoDB Config
        :param pulumi.Input[pulumi.InputType['DataSourceElasticsearchConfigArgs']] elasticsearch_config: Amazon Elasticsearch settings. See ElasticSearch Config
        :param pulumi.Input[pulumi.InputType['DataSourceEventBridgeConfigArgs']] event_bridge_config: AWS EventBridge settings. See Event Bridge Config
        :param pulumi.Input[pulumi.InputType['DataSourceHttpConfigArgs']] http_config: HTTP settings. See HTTP Config
        :param pulumi.Input[pulumi.InputType['DataSourceLambdaConfigArgs']] lambda_config: AWS Lambda settings. See Lambda Config
        :param pulumi.Input[str] name: User-supplied name for the data source.
        :param pulumi.Input[pulumi.InputType['DataSourceOpensearchserviceConfigArgs']] opensearchservice_config: Amazon OpenSearch Service settings. See OpenSearch Service Config
        :param pulumi.Input[pulumi.InputType['DataSourceRelationalDatabaseConfigArgs']] relational_database_config: AWS RDS settings. See Relational Database Config
        :param pulumi.Input[str] service_role_arn: IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        :param pulumi.Input[str] type: Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AppSync Data Source.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_table = aws.dynamodb.Table("exampleTable",
            read_capacity=1,
            write_capacity=1,
            hash_key="UserId",
            attributes=[aws.dynamodb.TableAttributeArgs(
                name="UserId",
                type="S",
            )])
        assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="Service",
                identifiers=["appsync.amazonaws.com"],
            )],
            actions=["sts:AssumeRole"],
        )])
        example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
        example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            actions=["dynamodb:*"],
            resources=[example_table.arn],
        )])
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=example_policy_document.json)
        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
        example_data_source = aws.appsync.DataSource("exampleDataSource",
            api_id=example_graph_ql_api.id,
            name="my_appsync_example",
            service_role_arn=example_role.arn,
            type="AMAZON_DYNAMODB",
            dynamodb_config=aws.appsync.DataSourceDynamodbConfigArgs(
                table_name=example_table.name,
            ))
        ```

        ## Import

        Using `pulumi import`, import `aws_appsync_datasource` using the `api_id`, a hyphen, and `name`. For example:

        ```sh
         $ pulumi import aws:appsync/dataSource:DataSource example abcdef123456-example
        ```

        :param str resource_name: The name of the resource.
        :param DataSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamodb_config: Optional[pulumi.Input[pulumi.InputType['DataSourceDynamodbConfigArgs']]] = None,
                 elasticsearch_config: Optional[pulumi.Input[pulumi.InputType['DataSourceElasticsearchConfigArgs']]] = None,
                 event_bridge_config: Optional[pulumi.Input[pulumi.InputType['DataSourceEventBridgeConfigArgs']]] = None,
                 http_config: Optional[pulumi.Input[pulumi.InputType['DataSourceHttpConfigArgs']]] = None,
                 lambda_config: Optional[pulumi.Input[pulumi.InputType['DataSourceLambdaConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 opensearchservice_config: Optional[pulumi.Input[pulumi.InputType['DataSourceOpensearchserviceConfigArgs']]] = None,
                 relational_database_config: Optional[pulumi.Input[pulumi.InputType['DataSourceRelationalDatabaseConfigArgs']]] = None,
                 service_role_arn: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataSourceArgs.__new__(DataSourceArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["description"] = description
            __props__.__dict__["dynamodb_config"] = dynamodb_config
            __props__.__dict__["elasticsearch_config"] = elasticsearch_config
            __props__.__dict__["event_bridge_config"] = event_bridge_config
            __props__.__dict__["http_config"] = http_config
            __props__.__dict__["lambda_config"] = lambda_config
            __props__.__dict__["name"] = name
            __props__.__dict__["opensearchservice_config"] = opensearchservice_config
            __props__.__dict__["relational_database_config"] = relational_database_config
            __props__.__dict__["service_role_arn"] = service_role_arn
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["arn"] = None
        super(DataSource, __self__).__init__(
            'aws:appsync/dataSource:DataSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamodb_config: Optional[pulumi.Input[pulumi.InputType['DataSourceDynamodbConfigArgs']]] = None,
            elasticsearch_config: Optional[pulumi.Input[pulumi.InputType['DataSourceElasticsearchConfigArgs']]] = None,
            event_bridge_config: Optional[pulumi.Input[pulumi.InputType['DataSourceEventBridgeConfigArgs']]] = None,
            http_config: Optional[pulumi.Input[pulumi.InputType['DataSourceHttpConfigArgs']]] = None,
            lambda_config: Optional[pulumi.Input[pulumi.InputType['DataSourceLambdaConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            opensearchservice_config: Optional[pulumi.Input[pulumi.InputType['DataSourceOpensearchserviceConfigArgs']]] = None,
            relational_database_config: Optional[pulumi.Input[pulumi.InputType['DataSourceRelationalDatabaseConfigArgs']]] = None,
            service_role_arn: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'DataSource':
        """
        Get an existing DataSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: API ID for the GraphQL API for the data source.
        :param pulumi.Input[str] arn: ARN
        :param pulumi.Input[str] description: Description of the data source.
        :param pulumi.Input[pulumi.InputType['DataSourceDynamodbConfigArgs']] dynamodb_config: DynamoDB settings. See DynamoDB Config
        :param pulumi.Input[pulumi.InputType['DataSourceElasticsearchConfigArgs']] elasticsearch_config: Amazon Elasticsearch settings. See ElasticSearch Config
        :param pulumi.Input[pulumi.InputType['DataSourceEventBridgeConfigArgs']] event_bridge_config: AWS EventBridge settings. See Event Bridge Config
        :param pulumi.Input[pulumi.InputType['DataSourceHttpConfigArgs']] http_config: HTTP settings. See HTTP Config
        :param pulumi.Input[pulumi.InputType['DataSourceLambdaConfigArgs']] lambda_config: AWS Lambda settings. See Lambda Config
        :param pulumi.Input[str] name: User-supplied name for the data source.
        :param pulumi.Input[pulumi.InputType['DataSourceOpensearchserviceConfigArgs']] opensearchservice_config: Amazon OpenSearch Service settings. See OpenSearch Service Config
        :param pulumi.Input[pulumi.InputType['DataSourceRelationalDatabaseConfigArgs']] relational_database_config: AWS RDS settings. See Relational Database Config
        :param pulumi.Input[str] service_role_arn: IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        :param pulumi.Input[str] type: Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSourceState.__new__(_DataSourceState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["dynamodb_config"] = dynamodb_config
        __props__.__dict__["elasticsearch_config"] = elasticsearch_config
        __props__.__dict__["event_bridge_config"] = event_bridge_config
        __props__.__dict__["http_config"] = http_config
        __props__.__dict__["lambda_config"] = lambda_config
        __props__.__dict__["name"] = name
        __props__.__dict__["opensearchservice_config"] = opensearchservice_config
        __props__.__dict__["relational_database_config"] = relational_database_config
        __props__.__dict__["service_role_arn"] = service_role_arn
        __props__.__dict__["type"] = type
        return DataSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        API ID for the GraphQL API for the data source.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the data source.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamodbConfig")
    def dynamodb_config(self) -> pulumi.Output[Optional['outputs.DataSourceDynamodbConfig']]:
        """
        DynamoDB settings. See DynamoDB Config
        """
        return pulumi.get(self, "dynamodb_config")

    @property
    @pulumi.getter(name="elasticsearchConfig")
    def elasticsearch_config(self) -> pulumi.Output[Optional['outputs.DataSourceElasticsearchConfig']]:
        """
        Amazon Elasticsearch settings. See ElasticSearch Config
        """
        return pulumi.get(self, "elasticsearch_config")

    @property
    @pulumi.getter(name="eventBridgeConfig")
    def event_bridge_config(self) -> pulumi.Output[Optional['outputs.DataSourceEventBridgeConfig']]:
        """
        AWS EventBridge settings. See Event Bridge Config
        """
        return pulumi.get(self, "event_bridge_config")

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> pulumi.Output[Optional['outputs.DataSourceHttpConfig']]:
        """
        HTTP settings. See HTTP Config
        """
        return pulumi.get(self, "http_config")

    @property
    @pulumi.getter(name="lambdaConfig")
    def lambda_config(self) -> pulumi.Output[Optional['outputs.DataSourceLambdaConfig']]:
        """
        AWS Lambda settings. See Lambda Config
        """
        return pulumi.get(self, "lambda_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User-supplied name for the data source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="opensearchserviceConfig")
    def opensearchservice_config(self) -> pulumi.Output[Optional['outputs.DataSourceOpensearchserviceConfig']]:
        """
        Amazon OpenSearch Service settings. See OpenSearch Service Config
        """
        return pulumi.get(self, "opensearchservice_config")

    @property
    @pulumi.getter(name="relationalDatabaseConfig")
    def relational_database_config(self) -> pulumi.Output[Optional['outputs.DataSourceRelationalDatabaseConfig']]:
        """
        AWS RDS settings. See Relational Database Config
        """
        return pulumi.get(self, "relational_database_config")

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        """
        return pulumi.get(self, "service_role_arn")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        """
        return pulumi.get(self, "type")

