// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync
{
    /// <summary>
    /// Provides an AppSync Data Source.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleTable = new Aws.DynamoDB.Table("exampleTable", new()
    ///     {
    ///         ReadCapacity = 1,
    ///         WriteCapacity = 1,
    ///         HashKey = "UserId",
    ///         Attributes = new[]
    ///         {
    ///             new Aws.DynamoDB.Inputs.TableAttributeArgs
    ///             {
    ///                 Name = "UserId",
    ///                 Type = "S",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "Service",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "appsync.amazonaws.com",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "sts:AssumeRole",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleRole = new Aws.Iam.Role("exampleRole", new()
    ///     {
    ///         AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    ///     var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Effect = "Allow",
    ///                 Actions = new[]
    ///                 {
    ///                     "dynamodb:*",
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     exampleTable.Arn,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new()
    ///     {
    ///         Role = exampleRole.Id,
    ///         Policy = examplePolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    ///     var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new()
    ///     {
    ///         AuthenticationType = "API_KEY",
    ///     });
    /// 
    ///     var exampleDataSource = new Aws.AppSync.DataSource("exampleDataSource", new()
    ///     {
    ///         ApiId = exampleGraphQLApi.Id,
    ///         Name = "my_appsync_example",
    ///         ServiceRoleArn = exampleRole.Arn,
    ///         Type = "AMAZON_DYNAMODB",
    ///         DynamodbConfig = new Aws.AppSync.Inputs.DataSourceDynamodbConfigArgs
    ///         {
    ///             TableName = exampleTable.Name,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_appsync_datasource` using the `api_id`, a hyphen, and `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:appsync/dataSource:DataSource example abcdef123456-example
    /// ```
    /// </summary>
    [AwsResourceType("aws:appsync/dataSource:DataSource")]
    public partial class DataSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// API ID for the GraphQL API for the data source.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// ARN
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the data source.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// DynamoDB settings. See DynamoDB Config
        /// </summary>
        [Output("dynamodbConfig")]
        public Output<Outputs.DataSourceDynamodbConfig?> DynamodbConfig { get; private set; } = null!;

        /// <summary>
        /// Amazon Elasticsearch settings. See ElasticSearch Config
        /// </summary>
        [Output("elasticsearchConfig")]
        public Output<Outputs.DataSourceElasticsearchConfig?> ElasticsearchConfig { get; private set; } = null!;

        /// <summary>
        /// AWS EventBridge settings. See Event Bridge Config
        /// </summary>
        [Output("eventBridgeConfig")]
        public Output<Outputs.DataSourceEventBridgeConfig?> EventBridgeConfig { get; private set; } = null!;

        /// <summary>
        /// HTTP settings. See HTTP Config
        /// </summary>
        [Output("httpConfig")]
        public Output<Outputs.DataSourceHttpConfig?> HttpConfig { get; private set; } = null!;

        /// <summary>
        /// AWS Lambda settings. See Lambda Config
        /// </summary>
        [Output("lambdaConfig")]
        public Output<Outputs.DataSourceLambdaConfig?> LambdaConfig { get; private set; } = null!;

        /// <summary>
        /// User-supplied name for the data source.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Amazon OpenSearch Service settings. See OpenSearch Service Config
        /// </summary>
        [Output("opensearchserviceConfig")]
        public Output<Outputs.DataSourceOpensearchserviceConfig?> OpensearchserviceConfig { get; private set; } = null!;

        /// <summary>
        /// AWS RDS settings. See Relational Database Config
        /// </summary>
        [Output("relationalDatabaseConfig")]
        public Output<Outputs.DataSourceRelationalDatabaseConfig?> RelationalDatabaseConfig { get; private set; } = null!;

        /// <summary>
        /// IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        /// </summary>
        [Output("serviceRoleArn")]
        public Output<string?> ServiceRoleArn { get; private set; } = null!;

        /// <summary>
        /// Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DataSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSource(string name, DataSourceArgs args, CustomResourceOptions? options = null)
            : base("aws:appsync/dataSource:DataSource", name, args ?? new DataSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSource(string name, Input<string> id, DataSourceState? state = null, CustomResourceOptions? options = null)
            : base("aws:appsync/dataSource:DataSource", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSource Get(string name, Input<string> id, DataSourceState? state = null, CustomResourceOptions? options = null)
        {
            return new DataSource(name, id, state, options);
        }
    }

    public sealed class DataSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API ID for the GraphQL API for the data source.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Description of the data source.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// DynamoDB settings. See DynamoDB Config
        /// </summary>
        [Input("dynamodbConfig")]
        public Input<Inputs.DataSourceDynamodbConfigArgs>? DynamodbConfig { get; set; }

        /// <summary>
        /// Amazon Elasticsearch settings. See ElasticSearch Config
        /// </summary>
        [Input("elasticsearchConfig")]
        public Input<Inputs.DataSourceElasticsearchConfigArgs>? ElasticsearchConfig { get; set; }

        /// <summary>
        /// AWS EventBridge settings. See Event Bridge Config
        /// </summary>
        [Input("eventBridgeConfig")]
        public Input<Inputs.DataSourceEventBridgeConfigArgs>? EventBridgeConfig { get; set; }

        /// <summary>
        /// HTTP settings. See HTTP Config
        /// </summary>
        [Input("httpConfig")]
        public Input<Inputs.DataSourceHttpConfigArgs>? HttpConfig { get; set; }

        /// <summary>
        /// AWS Lambda settings. See Lambda Config
        /// </summary>
        [Input("lambdaConfig")]
        public Input<Inputs.DataSourceLambdaConfigArgs>? LambdaConfig { get; set; }

        /// <summary>
        /// User-supplied name for the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Amazon OpenSearch Service settings. See OpenSearch Service Config
        /// </summary>
        [Input("opensearchserviceConfig")]
        public Input<Inputs.DataSourceOpensearchserviceConfigArgs>? OpensearchserviceConfig { get; set; }

        /// <summary>
        /// AWS RDS settings. See Relational Database Config
        /// </summary>
        [Input("relationalDatabaseConfig")]
        public Input<Inputs.DataSourceRelationalDatabaseConfigArgs>? RelationalDatabaseConfig { get; set; }

        /// <summary>
        /// IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        /// </summary>
        [Input("serviceRoleArn")]
        public Input<string>? ServiceRoleArn { get; set; }

        /// <summary>
        /// Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DataSourceArgs()
        {
        }
        public static new DataSourceArgs Empty => new DataSourceArgs();
    }

    public sealed class DataSourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API ID for the GraphQL API for the data source.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// ARN
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the data source.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// DynamoDB settings. See DynamoDB Config
        /// </summary>
        [Input("dynamodbConfig")]
        public Input<Inputs.DataSourceDynamodbConfigGetArgs>? DynamodbConfig { get; set; }

        /// <summary>
        /// Amazon Elasticsearch settings. See ElasticSearch Config
        /// </summary>
        [Input("elasticsearchConfig")]
        public Input<Inputs.DataSourceElasticsearchConfigGetArgs>? ElasticsearchConfig { get; set; }

        /// <summary>
        /// AWS EventBridge settings. See Event Bridge Config
        /// </summary>
        [Input("eventBridgeConfig")]
        public Input<Inputs.DataSourceEventBridgeConfigGetArgs>? EventBridgeConfig { get; set; }

        /// <summary>
        /// HTTP settings. See HTTP Config
        /// </summary>
        [Input("httpConfig")]
        public Input<Inputs.DataSourceHttpConfigGetArgs>? HttpConfig { get; set; }

        /// <summary>
        /// AWS Lambda settings. See Lambda Config
        /// </summary>
        [Input("lambdaConfig")]
        public Input<Inputs.DataSourceLambdaConfigGetArgs>? LambdaConfig { get; set; }

        /// <summary>
        /// User-supplied name for the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Amazon OpenSearch Service settings. See OpenSearch Service Config
        /// </summary>
        [Input("opensearchserviceConfig")]
        public Input<Inputs.DataSourceOpensearchserviceConfigGetArgs>? OpensearchserviceConfig { get; set; }

        /// <summary>
        /// AWS RDS settings. See Relational Database Config
        /// </summary>
        [Input("relationalDatabaseConfig")]
        public Input<Inputs.DataSourceRelationalDatabaseConfigGetArgs>? RelationalDatabaseConfig { get; set; }

        /// <summary>
        /// IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
        /// </summary>
        [Input("serviceRoleArn")]
        public Input<string>? ServiceRoleArn { get; set; }

        /// <summary>
        /// Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DataSourceState()
        {
        }
        public static new DataSourceState Empty => new DataSourceState();
    }
}
