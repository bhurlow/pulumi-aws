// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch
{
    /// <summary>
    /// Manages an AWS Elasticsearch Domain.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.ElasticSearch.Domain("example", new()
    ///     {
    ///         ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
    ///         {
    ///             InstanceType = "r4.large.elasticsearch",
    ///         },
    ///         ElasticsearchVersion = "7.10",
    ///         Tags = 
    ///         {
    ///             { "Domain", "TestDomain" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Access Policy
    /// 
    /// &gt; See also: `aws.elasticsearch.DomainPolicy` resource
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var domain = config.Get("domain") ?? "tf-test";
    ///     var currentRegion = Aws.GetRegion.Invoke();
    /// 
    ///     var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();
    /// 
    ///     var example = new Aws.ElasticSearch.Domain("example", new()
    ///     {
    ///         AccessPolicies = Output.Tuple(currentRegion, currentCallerIdentity).Apply(values =&gt;
    ///         {
    ///             var currentRegion = values.Item1;
    ///             var currentCallerIdentity = values.Item2;
    ///             return @$"{{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {{
    ///       ""Action"": ""es:*"",
    ///       ""Principal"": ""*"",
    ///       ""Effect"": ""Allow"",
    ///       ""Resource"": ""arn:aws:es:{currentRegion.Apply(getRegionResult =&gt; getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult =&gt; getCallerIdentityResult.AccountId)}:domain/{domain}/*"",
    ///       ""Condition"": {{
    ///         ""IpAddress"": {{""aws:SourceIp"": [""66.193.100.22/32""]}}
    ///       }}
    ///     }}
    ///   ]
    /// }}
    /// ";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// ### Log Publishing to CloudWatch Logs
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");
    /// 
    ///     var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
    ///                             "es.amazonaws.com",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "logs:PutLogEvents",
    ///                     "logs:PutLogEventsBatch",
    ///                     "logs:CreateLogStream",
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     "arn:aws:logs:*",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("exampleLogResourcePolicy", new()
    ///     {
    ///         PolicyName = "example",
    ///         PolicyDocument = examplePolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    ///     // .. other configuration ...
    ///     var exampleDomain = new Aws.ElasticSearch.Domain("exampleDomain", new()
    ///     {
    ///         LogPublishingOptions = new[]
    ///         {
    ///             new Aws.ElasticSearch.Inputs.DomainLogPublishingOptionArgs
    ///             {
    ///                 CloudwatchLogGroupArn = exampleLogGroup.Arn,
    ///                 LogType = "INDEX_SLOW_LOGS",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### VPC based ES
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var vpc = config.RequireObject&lt;dynamic&gt;("vpc");
    ///     var domain = config.Get("domain") ?? "tf-test";
    ///     var selectedVpc = Aws.Ec2.GetVpc.Invoke(new()
    ///     {
    ///         Tags = 
    ///         {
    ///             { "Name", vpc },
    ///         },
    ///     });
    /// 
    ///     var selectedSubnets = Aws.Ec2.GetSubnets.Invoke(new()
    ///     {
    ///         Filters = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.GetSubnetsFilterInputArgs
    ///             {
    ///                 Name = "vpc-id",
    ///                 Values = new[]
    ///                 {
    ///                     selectedVpc.Apply(getVpcResult =&gt; getVpcResult.Id),
    ///                 },
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Tier", "private" },
    ///         },
    ///     });
    /// 
    ///     var currentRegion = Aws.GetRegion.Invoke();
    /// 
    ///     var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();
    /// 
    ///     var esSecurityGroup = new Aws.Ec2.SecurityGroup("esSecurityGroup", new()
    ///     {
    ///         Description = "Managed by Pulumi",
    ///         VpcId = selectedVpc.Apply(getVpcResult =&gt; getVpcResult.Id),
    ///         Ingress = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.SecurityGroupIngressArgs
    ///             {
    ///                 FromPort = 443,
    ///                 ToPort = 443,
    ///                 Protocol = "tcp",
    ///                 CidrBlocks = new[]
    ///                 {
    ///                     selectedVpc.Apply(getVpcResult =&gt; getVpcResult.CidrBlock),
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var esServiceLinkedRole = new Aws.Iam.ServiceLinkedRole("esServiceLinkedRole", new()
    ///     {
    ///         AwsServiceName = "opensearchservice.amazonaws.com",
    ///     });
    /// 
    ///     var esDomain = new Aws.ElasticSearch.Domain("esDomain", new()
    ///     {
    ///         ElasticsearchVersion = "6.3",
    ///         ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
    ///         {
    ///             InstanceType = "m4.large.elasticsearch",
    ///             ZoneAwarenessEnabled = true,
    ///         },
    ///         VpcOptions = new Aws.ElasticSearch.Inputs.DomainVpcOptionsArgs
    ///         {
    ///             SubnetIds = new[]
    ///             {
    ///                 selectedSubnets.Apply(getSubnetsResult =&gt; getSubnetsResult.Ids[0]),
    ///                 selectedSubnets.Apply(getSubnetsResult =&gt; getSubnetsResult.Ids[1]),
    ///             },
    ///             SecurityGroupIds = new[]
    ///             {
    ///                 esSecurityGroup.Id,
    ///             },
    ///         },
    ///         AdvancedOptions = 
    ///         {
    ///             { "rest.action.multi.allow_explicit_index", "true" },
    ///         },
    ///         AccessPolicies = Output.Tuple(currentRegion, currentCallerIdentity).Apply(values =&gt;
    ///         {
    ///             var currentRegion = values.Item1;
    ///             var currentCallerIdentity = values.Item2;
    ///             return @$"{{
    /// 	""Version"": ""2012-10-17"",
    /// 	""Statement"": [
    /// 		{{
    /// 			""Action"": ""es:*"",
    /// 			""Principal"": ""*"",
    /// 			""Effect"": ""Allow"",
    /// 			""Resource"": ""arn:aws:es:{currentRegion.Apply(getRegionResult =&gt; getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult =&gt; getCallerIdentityResult.AccountId)}:domain/{domain}/*""
    /// 		}}
    /// 	]
    /// }}
    /// ";
    ///         }),
    ///         Tags = 
    ///         {
    ///             { "Domain", "TestDomain" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             esServiceLinkedRole,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_elasticsearch_domain.example
    /// 
    ///  id = "domain_name" } Using `pulumi import`, import Elasticsearch domains using the `domain_name`. For exampleconsole % pulumi import aws_elasticsearch_domain.example domain_name
    /// </summary>
    [AwsResourceType("aws:elasticsearch/domain:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IAM policy document specifying the access policies for the domain.
        /// </summary>
        [Output("accessPolicies")]
        public Output<string> AccessPolicies { get; private set; } = null!;

        /// <summary>
        /// Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your Elasticsearch domain on every apply.
        /// </summary>
        [Output("advancedOptions")]
        public Output<ImmutableDictionary<string, string>> AdvancedOptions { get; private set; } = null!;

        /// <summary>
        /// Configuration block for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). Detailed below.
        /// </summary>
        [Output("advancedSecurityOptions")]
        public Output<Outputs.DomainAdvancedSecurityOptions> AdvancedSecurityOptions { get; private set; } = null!;

        /// <summary>
        /// ARN of the domain.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the Auto-Tune options of the domain. Detailed below.
        /// </summary>
        [Output("autoTuneOptions")]
        public Output<Outputs.DomainAutoTuneOptions> AutoTuneOptions { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the cluster of the domain. Detailed below.
        /// </summary>
        [Output("clusterConfig")]
        public Output<Outputs.DomainClusterConfig> ClusterConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration block for authenticating Kibana with Cognito. Detailed below.
        /// </summary>
        [Output("cognitoOptions")]
        public Output<Outputs.DomainCognitoOptions?> CognitoOptions { get; private set; } = null!;

        /// <summary>
        /// Configuration block for domain endpoint HTTP(S) related options. Detailed below.
        /// </summary>
        [Output("domainEndpointOptions")]
        public Output<Outputs.DomainDomainEndpointOptions> DomainEndpointOptions { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the domain.
        /// </summary>
        [Output("domainId")]
        public Output<string> DomainId { get; private set; } = null!;

        /// <summary>
        /// Name of the domain.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). Detailed below.
        /// </summary>
        [Output("ebsOptions")]
        public Output<Outputs.DomainEbsOptions> EbsOptions { get; private set; } = null!;

        /// <summary>
        /// Version of Elasticsearch to deploy. Defaults to `1.5`.
        /// </summary>
        [Output("elasticsearchVersion")]
        public Output<string?> ElasticsearchVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration block for encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). Detailed below.
        /// </summary>
        [Output("encryptAtRest")]
        public Output<Outputs.DomainEncryptAtRest> EncryptAtRest { get; private set; } = null!;

        /// <summary>
        /// Domain-specific endpoint used to submit index, search, and data upload requests.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Domain-specific endpoint for kibana without https scheme.
        /// </summary>
        [Output("kibanaEndpoint")]
        public Output<string> KibanaEndpoint { get; private set; } = null!;

        /// <summary>
        /// Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
        /// </summary>
        [Output("logPublishingOptions")]
        public Output<ImmutableArray<Outputs.DomainLogPublishingOption>> LogPublishingOptions { get; private set; } = null!;

        /// <summary>
        /// Configuration block for node-to-node encryption options. Detailed below.
        /// </summary>
        [Output("nodeToNodeEncryption")]
        public Output<Outputs.DomainNodeToNodeEncryption> NodeToNodeEncryption { get; private set; } = null!;

        /// <summary>
        /// Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running Elasticsearch 5.3 and later, Amazon ES takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions of Elasticsearch, Amazon ES takes daily automated snapshots.
        /// </summary>
        [Output("snapshotOptions")]
        public Output<Outputs.DomainSnapshotOptions?> SnapshotOptions { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
        /// * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)). Detailed below.
        /// </summary>
        [Output("vpcOptions")]
        public Output<Outputs.DomainVpcOptions?> VpcOptions { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:elasticsearch/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticsearch/domain:Domain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAM policy document specifying the access policies for the domain.
        /// </summary>
        [Input("accessPolicies")]
        public Input<string>? AccessPolicies { get; set; }

        [Input("advancedOptions")]
        private InputMap<string>? _advancedOptions;

        /// <summary>
        /// Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your Elasticsearch domain on every apply.
        /// </summary>
        public InputMap<string> AdvancedOptions
        {
            get => _advancedOptions ?? (_advancedOptions = new InputMap<string>());
            set => _advancedOptions = value;
        }

        /// <summary>
        /// Configuration block for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). Detailed below.
        /// </summary>
        [Input("advancedSecurityOptions")]
        public Input<Inputs.DomainAdvancedSecurityOptionsArgs>? AdvancedSecurityOptions { get; set; }

        /// <summary>
        /// Configuration block for the Auto-Tune options of the domain. Detailed below.
        /// </summary>
        [Input("autoTuneOptions")]
        public Input<Inputs.DomainAutoTuneOptionsArgs>? AutoTuneOptions { get; set; }

        /// <summary>
        /// Configuration block for the cluster of the domain. Detailed below.
        /// </summary>
        [Input("clusterConfig")]
        public Input<Inputs.DomainClusterConfigArgs>? ClusterConfig { get; set; }

        /// <summary>
        /// Configuration block for authenticating Kibana with Cognito. Detailed below.
        /// </summary>
        [Input("cognitoOptions")]
        public Input<Inputs.DomainCognitoOptionsArgs>? CognitoOptions { get; set; }

        /// <summary>
        /// Configuration block for domain endpoint HTTP(S) related options. Detailed below.
        /// </summary>
        [Input("domainEndpointOptions")]
        public Input<Inputs.DomainDomainEndpointOptionsArgs>? DomainEndpointOptions { get; set; }

        /// <summary>
        /// Name of the domain.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). Detailed below.
        /// </summary>
        [Input("ebsOptions")]
        public Input<Inputs.DomainEbsOptionsArgs>? EbsOptions { get; set; }

        /// <summary>
        /// Version of Elasticsearch to deploy. Defaults to `1.5`.
        /// </summary>
        [Input("elasticsearchVersion")]
        public Input<string>? ElasticsearchVersion { get; set; }

        /// <summary>
        /// Configuration block for encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). Detailed below.
        /// </summary>
        [Input("encryptAtRest")]
        public Input<Inputs.DomainEncryptAtRestArgs>? EncryptAtRest { get; set; }

        [Input("logPublishingOptions")]
        private InputList<Inputs.DomainLogPublishingOptionArgs>? _logPublishingOptions;

        /// <summary>
        /// Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
        /// </summary>
        public InputList<Inputs.DomainLogPublishingOptionArgs> LogPublishingOptions
        {
            get => _logPublishingOptions ?? (_logPublishingOptions = new InputList<Inputs.DomainLogPublishingOptionArgs>());
            set => _logPublishingOptions = value;
        }

        /// <summary>
        /// Configuration block for node-to-node encryption options. Detailed below.
        /// </summary>
        [Input("nodeToNodeEncryption")]
        public Input<Inputs.DomainNodeToNodeEncryptionArgs>? NodeToNodeEncryption { get; set; }

        /// <summary>
        /// Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running Elasticsearch 5.3 and later, Amazon ES takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions of Elasticsearch, Amazon ES takes daily automated snapshots.
        /// </summary>
        [Input("snapshotOptions")]
        public Input<Inputs.DomainSnapshotOptionsArgs>? SnapshotOptions { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)). Detailed below.
        /// </summary>
        [Input("vpcOptions")]
        public Input<Inputs.DomainVpcOptionsArgs>? VpcOptions { get; set; }

        public DomainArgs()
        {
        }
        public static new DomainArgs Empty => new DomainArgs();
    }

    public sealed class DomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAM policy document specifying the access policies for the domain.
        /// </summary>
        [Input("accessPolicies")]
        public Input<string>? AccessPolicies { get; set; }

        [Input("advancedOptions")]
        private InputMap<string>? _advancedOptions;

        /// <summary>
        /// Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your Elasticsearch domain on every apply.
        /// </summary>
        public InputMap<string> AdvancedOptions
        {
            get => _advancedOptions ?? (_advancedOptions = new InputMap<string>());
            set => _advancedOptions = value;
        }

        /// <summary>
        /// Configuration block for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). Detailed below.
        /// </summary>
        [Input("advancedSecurityOptions")]
        public Input<Inputs.DomainAdvancedSecurityOptionsGetArgs>? AdvancedSecurityOptions { get; set; }

        /// <summary>
        /// ARN of the domain.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Configuration block for the Auto-Tune options of the domain. Detailed below.
        /// </summary>
        [Input("autoTuneOptions")]
        public Input<Inputs.DomainAutoTuneOptionsGetArgs>? AutoTuneOptions { get; set; }

        /// <summary>
        /// Configuration block for the cluster of the domain. Detailed below.
        /// </summary>
        [Input("clusterConfig")]
        public Input<Inputs.DomainClusterConfigGetArgs>? ClusterConfig { get; set; }

        /// <summary>
        /// Configuration block for authenticating Kibana with Cognito. Detailed below.
        /// </summary>
        [Input("cognitoOptions")]
        public Input<Inputs.DomainCognitoOptionsGetArgs>? CognitoOptions { get; set; }

        /// <summary>
        /// Configuration block for domain endpoint HTTP(S) related options. Detailed below.
        /// </summary>
        [Input("domainEndpointOptions")]
        public Input<Inputs.DomainDomainEndpointOptionsGetArgs>? DomainEndpointOptions { get; set; }

        /// <summary>
        /// Unique identifier for the domain.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Name of the domain.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). Detailed below.
        /// </summary>
        [Input("ebsOptions")]
        public Input<Inputs.DomainEbsOptionsGetArgs>? EbsOptions { get; set; }

        /// <summary>
        /// Version of Elasticsearch to deploy. Defaults to `1.5`.
        /// </summary>
        [Input("elasticsearchVersion")]
        public Input<string>? ElasticsearchVersion { get; set; }

        /// <summary>
        /// Configuration block for encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). Detailed below.
        /// </summary>
        [Input("encryptAtRest")]
        public Input<Inputs.DomainEncryptAtRestGetArgs>? EncryptAtRest { get; set; }

        /// <summary>
        /// Domain-specific endpoint used to submit index, search, and data upload requests.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Domain-specific endpoint for kibana without https scheme.
        /// </summary>
        [Input("kibanaEndpoint")]
        public Input<string>? KibanaEndpoint { get; set; }

        [Input("logPublishingOptions")]
        private InputList<Inputs.DomainLogPublishingOptionGetArgs>? _logPublishingOptions;

        /// <summary>
        /// Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
        /// </summary>
        public InputList<Inputs.DomainLogPublishingOptionGetArgs> LogPublishingOptions
        {
            get => _logPublishingOptions ?? (_logPublishingOptions = new InputList<Inputs.DomainLogPublishingOptionGetArgs>());
            set => _logPublishingOptions = value;
        }

        /// <summary>
        /// Configuration block for node-to-node encryption options. Detailed below.
        /// </summary>
        [Input("nodeToNodeEncryption")]
        public Input<Inputs.DomainNodeToNodeEncryptionGetArgs>? NodeToNodeEncryption { get; set; }

        /// <summary>
        /// Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running Elasticsearch 5.3 and later, Amazon ES takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions of Elasticsearch, Amazon ES takes daily automated snapshots.
        /// </summary>
        [Input("snapshotOptions")]
        public Input<Inputs.DomainSnapshotOptionsGetArgs>? SnapshotOptions { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
        /// * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)). Detailed below.
        /// </summary>
        [Input("vpcOptions")]
        public Input<Inputs.DomainVpcOptionsGetArgs>? VpcOptions { get; set; }

        public DomainState()
        {
        }
        public static new DomainState Empty => new DomainState();
    }
}
