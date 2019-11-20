// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    /// <summary>
    /// The provider type for the aws package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/index.html.markdown.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, ResourceOptions? options = null)
            : base("aws", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private static ResourceOptions MakeResourceOptions(ResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of
        /// the AWS console.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        [Input("allowedAccountIds", json: true)]
        private InputList<string>? _allowedAccountIds;
        public InputList<string> AllowedAccountIds
        {
            get => _allowedAccountIds ?? (_allowedAccountIds = new InputList<string>());
            set => _allowedAccountIds = value;
        }

        [Input("assumeRole", json: true)]
        public Input<Inputs.ProviderAssumeRoleArgs>? AssumeRole { get; set; }

        [Input("endpoints", json: true)]
        private InputList<Inputs.ProviderEndpointsArgs>? _endpoints;
        public InputList<Inputs.ProviderEndpointsArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.ProviderEndpointsArgs>());
            set => _endpoints = value;
        }

        [Input("forbiddenAccountIds", json: true)]
        private InputList<string>? _forbiddenAccountIds;
        public InputList<string> ForbiddenAccountIds
        {
            get => _forbiddenAccountIds ?? (_forbiddenAccountIds = new InputList<string>());
            set => _forbiddenAccountIds = value;
        }

        [Input("ignoreTagPrefixes", json: true)]
        private InputList<string>? _ignoreTagPrefixes;

        /// <summary>
        /// Resource tag key prefixes to ignore across all resources.
        /// </summary>
        public InputList<string> IgnoreTagPrefixes
        {
            get => _ignoreTagPrefixes ?? (_ignoreTagPrefixes = new InputList<string>());
            set => _ignoreTagPrefixes = value;
        }

        [Input("ignoreTags", json: true)]
        private InputList<string>? _ignoreTags;

        /// <summary>
        /// Resource tag keys to ignore across all resources.
        /// </summary>
        public InputList<string> IgnoreTags
        {
            get => _ignoreTags ?? (_ignoreTags = new InputList<string>());
            set => _ignoreTags = value;
        }

        /// <summary>
        /// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`
        /// </summary>
        [Input("insecure", json: true)]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// The maximum number of times an AWS API request is being executed. If the API request still fails, an error
        /// is thrown.
        /// </summary>
        [Input("maxRetries", json: true)]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Set this to true to force the request to use path-style addressing, i.e.,
        /// http://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing when
        /// possible (http://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.
        /// </summary>
        [Input("s3ForcePathStyle", json: true)]
        public Input<bool>? S3ForcePathStyle { get; set; }

        /// <summary>
        /// The secret key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of
        /// the AWS console.
        /// </summary>
        [Input("secretKey")]
        public Input<string>? SecretKey { get; set; }

        /// <summary>
        /// The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.
        /// </summary>
        [Input("sharedCredentialsFile")]
        public Input<string>? SharedCredentialsFile { get; set; }

        /// <summary>
        /// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
        /// available/implemented.
        /// </summary>
        [Input("skipCredentialsValidation", json: true)]
        public Input<bool>? SkipCredentialsValidation { get; set; }

        /// <summary>
        /// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes
        /// permissions.
        /// </summary>
        [Input("skipGetEc2Platforms", json: true)]
        public Input<bool>? SkipGetEc2Platforms { get; set; }

        [Input("skipMetadataApiCheck", json: true)]
        public Input<bool>? SkipMetadataApiCheck { get; set; }

        /// <summary>
        /// Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to
        /// regions that are not public (yet).
        /// </summary>
        [Input("skipRegionValidation", json: true)]
        public Input<bool>? SkipRegionValidation { get; set; }

        /// <summary>
        /// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or
        /// metadata API.
        /// </summary>
        [Input("skipRequestingAccountId", json: true)]
        public Input<bool>? SkipRequestingAccountId { get; set; }

        /// <summary>
        /// session token. A session token is only required if you are using temporary security credentials.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public ProviderArgs()
        {
            Profile = Utilities.GetEnv("AWS_PROFILE");
            Region = Utilities.GetEnv("AWS_REGION", "AWS_DEFAULT_REGION");
        }
    }

    namespace Inputs
    {

    public sealed class ProviderAssumeRoleArgs : Pulumi.ResourceArgs
    {
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("sessionName")]
        public Input<string>? SessionName { get; set; }

        public ProviderAssumeRoleArgs()
        {
        }
    }

    public sealed class ProviderEndpointsArgs : Pulumi.ResourceArgs
    {
        [Input("acm")]
        public Input<string>? Acm { get; set; }

        [Input("acmpca")]
        public Input<string>? Acmpca { get; set; }

        [Input("amplify")]
        public Input<string>? Amplify { get; set; }

        [Input("apigateway")]
        public Input<string>? Apigateway { get; set; }

        [Input("applicationautoscaling")]
        public Input<string>? Applicationautoscaling { get; set; }

        [Input("applicationinsights")]
        public Input<string>? Applicationinsights { get; set; }

        [Input("appmesh")]
        public Input<string>? Appmesh { get; set; }

        [Input("appstream")]
        public Input<string>? Appstream { get; set; }

        [Input("appsync")]
        public Input<string>? Appsync { get; set; }

        [Input("athena")]
        public Input<string>? Athena { get; set; }

        [Input("autoscaling")]
        public Input<string>? Autoscaling { get; set; }

        [Input("autoscalingplans")]
        public Input<string>? Autoscalingplans { get; set; }

        [Input("backup")]
        public Input<string>? Backup { get; set; }

        [Input("batch")]
        public Input<string>? Batch { get; set; }

        [Input("budgets")]
        public Input<string>? Budgets { get; set; }

        [Input("cloud9")]
        public Input<string>? Cloud9 { get; set; }

        [Input("cloudformation")]
        public Input<string>? Cloudformation { get; set; }

        [Input("cloudfront")]
        public Input<string>? Cloudfront { get; set; }

        [Input("cloudhsm")]
        public Input<string>? Cloudhsm { get; set; }

        [Input("cloudsearch")]
        public Input<string>? Cloudsearch { get; set; }

        [Input("cloudtrail")]
        public Input<string>? Cloudtrail { get; set; }

        [Input("cloudwatch")]
        public Input<string>? Cloudwatch { get; set; }

        [Input("cloudwatchevents")]
        public Input<string>? Cloudwatchevents { get; set; }

        [Input("cloudwatchlogs")]
        public Input<string>? Cloudwatchlogs { get; set; }

        [Input("codebuild")]
        public Input<string>? Codebuild { get; set; }

        [Input("codecommit")]
        public Input<string>? Codecommit { get; set; }

        [Input("codedeploy")]
        public Input<string>? Codedeploy { get; set; }

        [Input("codepipeline")]
        public Input<string>? Codepipeline { get; set; }

        [Input("cognitoidentity")]
        public Input<string>? Cognitoidentity { get; set; }

        [Input("cognitoidp")]
        public Input<string>? Cognitoidp { get; set; }

        [Input("configservice")]
        public Input<string>? Configservice { get; set; }

        [Input("cur")]
        public Input<string>? Cur { get; set; }

        [Input("datapipeline")]
        public Input<string>? Datapipeline { get; set; }

        [Input("datasync")]
        public Input<string>? Datasync { get; set; }

        [Input("dax")]
        public Input<string>? Dax { get; set; }

        [Input("devicefarm")]
        public Input<string>? Devicefarm { get; set; }

        [Input("directconnect")]
        public Input<string>? Directconnect { get; set; }

        [Input("dlm")]
        public Input<string>? Dlm { get; set; }

        [Input("dms")]
        public Input<string>? Dms { get; set; }

        [Input("docdb")]
        public Input<string>? Docdb { get; set; }

        [Input("ds")]
        public Input<string>? Ds { get; set; }

        [Input("dynamodb")]
        public Input<string>? Dynamodb { get; set; }

        [Input("ec2")]
        public Input<string>? Ec2 { get; set; }

        [Input("ecr")]
        public Input<string>? Ecr { get; set; }

        [Input("ecs")]
        public Input<string>? Ecs { get; set; }

        [Input("efs")]
        public Input<string>? Efs { get; set; }

        [Input("eks")]
        public Input<string>? Eks { get; set; }

        [Input("elasticache")]
        public Input<string>? Elasticache { get; set; }

        [Input("elasticbeanstalk")]
        public Input<string>? Elasticbeanstalk { get; set; }

        [Input("elastictranscoder")]
        public Input<string>? Elastictranscoder { get; set; }

        [Input("elb")]
        public Input<string>? Elb { get; set; }

        [Input("emr")]
        public Input<string>? Emr { get; set; }

        [Input("es")]
        public Input<string>? Es { get; set; }

        [Input("firehose")]
        public Input<string>? Firehose { get; set; }

        [Input("fms")]
        public Input<string>? Fms { get; set; }

        [Input("forecast")]
        public Input<string>? Forecast { get; set; }

        [Input("fsx")]
        public Input<string>? Fsx { get; set; }

        [Input("gamelift")]
        public Input<string>? Gamelift { get; set; }

        [Input("glacier")]
        public Input<string>? Glacier { get; set; }

        [Input("globalaccelerator")]
        public Input<string>? Globalaccelerator { get; set; }

        [Input("glue")]
        public Input<string>? Glue { get; set; }

        [Input("greengrass")]
        public Input<string>? Greengrass { get; set; }

        [Input("guardduty")]
        public Input<string>? Guardduty { get; set; }

        [Input("iam")]
        public Input<string>? Iam { get; set; }

        [Input("inspector")]
        public Input<string>? Inspector { get; set; }

        [Input("iot")]
        public Input<string>? Iot { get; set; }

        [Input("iotanalytics")]
        public Input<string>? Iotanalytics { get; set; }

        [Input("iotevents")]
        public Input<string>? Iotevents { get; set; }

        [Input("kafka")]
        public Input<string>? Kafka { get; set; }

        [Input("kinesis")]
        public Input<string>? Kinesis { get; set; }

        [Input("kinesisAnalytics")]
        public Input<string>? KinesisAnalytics { get; set; }

        [Input("kinesisanalytics")]
        public Input<string>? Kinesisanalytics { get; set; }

        [Input("kinesisvideo")]
        public Input<string>? Kinesisvideo { get; set; }

        [Input("kms")]
        public Input<string>? Kms { get; set; }

        [Input("lakeformation")]
        public Input<string>? Lakeformation { get; set; }

        [Input("lambda")]
        public Input<string>? Lambda { get; set; }

        [Input("lexmodels")]
        public Input<string>? Lexmodels { get; set; }

        [Input("licensemanager")]
        public Input<string>? Licensemanager { get; set; }

        [Input("lightsail")]
        public Input<string>? Lightsail { get; set; }

        [Input("macie")]
        public Input<string>? Macie { get; set; }

        [Input("managedblockchain")]
        public Input<string>? Managedblockchain { get; set; }

        [Input("mediaconnect")]
        public Input<string>? Mediaconnect { get; set; }

        [Input("mediaconvert")]
        public Input<string>? Mediaconvert { get; set; }

        [Input("medialive")]
        public Input<string>? Medialive { get; set; }

        [Input("mediapackage")]
        public Input<string>? Mediapackage { get; set; }

        [Input("mediastore")]
        public Input<string>? Mediastore { get; set; }

        [Input("mediastoredata")]
        public Input<string>? Mediastoredata { get; set; }

        [Input("mq")]
        public Input<string>? Mq { get; set; }

        [Input("neptune")]
        public Input<string>? Neptune { get; set; }

        [Input("opsworks")]
        public Input<string>? Opsworks { get; set; }

        [Input("organizations")]
        public Input<string>? Organizations { get; set; }

        [Input("personalize")]
        public Input<string>? Personalize { get; set; }

        [Input("pinpoint")]
        public Input<string>? Pinpoint { get; set; }

        [Input("pricing")]
        public Input<string>? Pricing { get; set; }

        [Input("qldb")]
        public Input<string>? Qldb { get; set; }

        [Input("quicksight")]
        public Input<string>? Quicksight { get; set; }

        [Input("r53")]
        public Input<string>? R53 { get; set; }

        [Input("ram")]
        public Input<string>? Ram { get; set; }

        [Input("rds")]
        public Input<string>? Rds { get; set; }

        [Input("redshift")]
        public Input<string>? Redshift { get; set; }

        [Input("resourcegroups")]
        public Input<string>? Resourcegroups { get; set; }

        [Input("route53")]
        public Input<string>? Route53 { get; set; }

        [Input("route53resolver")]
        public Input<string>? Route53resolver { get; set; }

        [Input("s3")]
        public Input<string>? S3 { get; set; }

        [Input("s3control")]
        public Input<string>? S3control { get; set; }

        [Input("sagemaker")]
        public Input<string>? Sagemaker { get; set; }

        [Input("sdb")]
        public Input<string>? Sdb { get; set; }

        [Input("secretsmanager")]
        public Input<string>? Secretsmanager { get; set; }

        [Input("securityhub")]
        public Input<string>? Securityhub { get; set; }

        [Input("serverlessrepo")]
        public Input<string>? Serverlessrepo { get; set; }

        [Input("servicecatalog")]
        public Input<string>? Servicecatalog { get; set; }

        [Input("servicediscovery")]
        public Input<string>? Servicediscovery { get; set; }

        [Input("servicequotas")]
        public Input<string>? Servicequotas { get; set; }

        [Input("ses")]
        public Input<string>? Ses { get; set; }

        [Input("shield")]
        public Input<string>? Shield { get; set; }

        [Input("sns")]
        public Input<string>? Sns { get; set; }

        [Input("sqs")]
        public Input<string>? Sqs { get; set; }

        [Input("ssm")]
        public Input<string>? Ssm { get; set; }

        [Input("stepfunctions")]
        public Input<string>? Stepfunctions { get; set; }

        [Input("storagegateway")]
        public Input<string>? Storagegateway { get; set; }

        [Input("sts")]
        public Input<string>? Sts { get; set; }

        [Input("swf")]
        public Input<string>? Swf { get; set; }

        [Input("transfer")]
        public Input<string>? Transfer { get; set; }

        [Input("waf")]
        public Input<string>? Waf { get; set; }

        [Input("wafregional")]
        public Input<string>? Wafregional { get; set; }

        [Input("worklink")]
        public Input<string>? Worklink { get; set; }

        [Input("workspaces")]
        public Input<string>? Workspaces { get; set; }

        [Input("xray")]
        public Input<string>? Xray { get; set; }

        public ProviderEndpointsArgs()
        {
        }
    }
    }
}
