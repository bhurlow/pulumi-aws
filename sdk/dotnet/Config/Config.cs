// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Aws
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("aws");

        private static readonly __Value<string?> _accessKey = new __Value<string?>(() => __config.Get("accessKey"));
        /// <summary>
        /// The access key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of the AWS console.
        /// </summary>
        public static string? AccessKey
        {
            get => _accessKey.Get();
            set => _accessKey.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _allowedAccountIds = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("allowedAccountIds"));
        public static ImmutableArray<string> AllowedAccountIds
        {
            get => _allowedAccountIds.Get();
            set => _allowedAccountIds.Set(value);
        }

        private static readonly __Value<Pulumi.Aws.Config.Types.AssumeRole?> _assumeRole = new __Value<Pulumi.Aws.Config.Types.AssumeRole?>(() => __config.GetObject<Pulumi.Aws.Config.Types.AssumeRole>("assumeRole"));
        public static Pulumi.Aws.Config.Types.AssumeRole? AssumeRole
        {
            get => _assumeRole.Get();
            set => _assumeRole.Set(value);
        }

        private static readonly __Value<Pulumi.Aws.Config.Types.AssumeRoleWithWebIdentity?> _assumeRoleWithWebIdentity = new __Value<Pulumi.Aws.Config.Types.AssumeRoleWithWebIdentity?>(() => __config.GetObject<Pulumi.Aws.Config.Types.AssumeRoleWithWebIdentity>("assumeRoleWithWebIdentity"));
        public static Pulumi.Aws.Config.Types.AssumeRoleWithWebIdentity? AssumeRoleWithWebIdentity
        {
            get => _assumeRoleWithWebIdentity.Get();
            set => _assumeRoleWithWebIdentity.Set(value);
        }

        private static readonly __Value<string?> _customCaBundle = new __Value<string?>(() => __config.Get("customCaBundle"));
        /// <summary>
        /// File containing custom root and intermediate certificates. Can also be configured using the `AWS_CA_BUNDLE` environment
        /// variable. (Setting `ca_bundle` in the shared config file is not supported.)
        /// </summary>
        public static string? CustomCaBundle
        {
            get => _customCaBundle.Get();
            set => _customCaBundle.Set(value);
        }

        private static readonly __Value<Pulumi.Aws.Config.Types.DefaultTags?> _defaultTags = new __Value<Pulumi.Aws.Config.Types.DefaultTags?>(() => __config.GetObject<Pulumi.Aws.Config.Types.DefaultTags>("defaultTags"));
        /// <summary>
        /// Configuration block with settings to default resource tags across all resources.
        /// </summary>
        public static Pulumi.Aws.Config.Types.DefaultTags? DefaultTags
        {
            get => _defaultTags.Get();
            set => _defaultTags.Set(value);
        }

        private static readonly __Value<string?> _ec2MetadataServiceEndpoint = new __Value<string?>(() => __config.Get("ec2MetadataServiceEndpoint"));
        /// <summary>
        /// Address of the EC2 metadata service endpoint to use. Can also be configured using the
        /// `AWS_EC2_METADATA_SERVICE_ENDPOINT` environment variable.
        /// </summary>
        public static string? Ec2MetadataServiceEndpoint
        {
            get => _ec2MetadataServiceEndpoint.Get();
            set => _ec2MetadataServiceEndpoint.Set(value);
        }

        private static readonly __Value<string?> _ec2MetadataServiceEndpointMode = new __Value<string?>(() => __config.Get("ec2MetadataServiceEndpointMode"));
        /// <summary>
        /// Protocol to use with EC2 metadata service endpoint.Valid values are `IPv4` and `IPv6`. Can also be configured using the
        /// `AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE` environment variable.
        /// </summary>
        public static string? Ec2MetadataServiceEndpointMode
        {
            get => _ec2MetadataServiceEndpointMode.Get();
            set => _ec2MetadataServiceEndpointMode.Set(value);
        }

        private static readonly __Value<ImmutableArray<Pulumi.Aws.Config.Types.Endpoints>> _endpoints = new __Value<ImmutableArray<Pulumi.Aws.Config.Types.Endpoints>>(() => __config.GetObject<ImmutableArray<Pulumi.Aws.Config.Types.Endpoints>>("endpoints"));
        public static ImmutableArray<Pulumi.Aws.Config.Types.Endpoints> Endpoints
        {
            get => _endpoints.Get();
            set => _endpoints.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _forbiddenAccountIds = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("forbiddenAccountIds"));
        public static ImmutableArray<string> ForbiddenAccountIds
        {
            get => _forbiddenAccountIds.Get();
            set => _forbiddenAccountIds.Set(value);
        }

        private static readonly __Value<string?> _httpProxy = new __Value<string?>(() => __config.Get("httpProxy"));
        /// <summary>
        /// URL of a proxy to use for HTTP requests when accessing the AWS API. Can also be set using the `HTTP_PROXY` or
        /// `http_proxy` environment variables.
        /// </summary>
        public static string? HttpProxy
        {
            get => _httpProxy.Get();
            set => _httpProxy.Set(value);
        }

        private static readonly __Value<string?> _httpsProxy = new __Value<string?>(() => __config.Get("httpsProxy"));
        /// <summary>
        /// URL of a proxy to use for HTTPS requests when accessing the AWS API. Can also be set using the `HTTPS_PROXY` or
        /// `https_proxy` environment variables.
        /// </summary>
        public static string? HttpsProxy
        {
            get => _httpsProxy.Get();
            set => _httpsProxy.Set(value);
        }

        private static readonly __Value<Pulumi.Aws.Config.Types.IgnoreTags?> _ignoreTags = new __Value<Pulumi.Aws.Config.Types.IgnoreTags?>(() => __config.GetObject<Pulumi.Aws.Config.Types.IgnoreTags>("ignoreTags"));
        /// <summary>
        /// Configuration block with settings to ignore resource tags across all resources.
        /// </summary>
        public static Pulumi.Aws.Config.Types.IgnoreTags? IgnoreTags
        {
            get => _ignoreTags.Get();
            set => _ignoreTags.Set(value);
        }

        private static readonly __Value<bool?> _insecure = new __Value<bool?>(() => __config.GetBoolean("insecure"));
        /// <summary>
        /// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`
        /// </summary>
        public static bool? Insecure
        {
            get => _insecure.Get();
            set => _insecure.Set(value);
        }

        private static readonly __Value<int?> _maxRetries = new __Value<int?>(() => __config.GetInt32("maxRetries"));
        /// <summary>
        /// The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
        /// </summary>
        public static int? MaxRetries
        {
            get => _maxRetries.Get();
            set => _maxRetries.Set(value);
        }

        private static readonly __Value<string?> _noProxy = new __Value<string?>(() => __config.Get("noProxy"));
        /// <summary>
        /// Comma-separated list of hosts that should not use HTTP or HTTPS proxies. Can also be set using the `NO_PROXY` or
        /// `no_proxy` environment variables.
        /// </summary>
        public static string? NoProxy
        {
            get => _noProxy.Get();
            set => _noProxy.Set(value);
        }

        private static readonly __Value<string?> _profile = new __Value<string?>(() => __config.Get("profile"));
        /// <summary>
        /// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
        /// </summary>
        public static string? Profile
        {
            get => _profile.Get();
            set => _profile.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region") ?? Utilities.GetEnv("AWS_REGION", "AWS_DEFAULT_REGION"));
        /// <summary>
        /// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
        /// </summary>
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _retryMode = new __Value<string?>(() => __config.Get("retryMode"));
        /// <summary>
        /// Specifies how retries are attempted. Valid values are `standard` and `adaptive`. Can also be configured using the
        /// `AWS_RETRY_MODE` environment variable.
        /// </summary>
        public static string? RetryMode
        {
            get => _retryMode.Get();
            set => _retryMode.Set(value);
        }

        private static readonly __Value<string?> _s3UsEast1RegionalEndpoint = new __Value<string?>(() => __config.Get("s3UsEast1RegionalEndpoint"));
        /// <summary>
        /// Specifies whether S3 API calls in the `us-east-1` region use the legacy global endpoint or a regional endpoint. Valid
        /// values are `legacy` or `regional`. Can also be configured using the `AWS_S3_US_EAST_1_REGIONAL_ENDPOINT` environment
        /// variable or the `s3_us_east_1_regional_endpoint` shared config file parameter
        /// </summary>
        public static string? S3UsEast1RegionalEndpoint
        {
            get => _s3UsEast1RegionalEndpoint.Get();
            set => _s3UsEast1RegionalEndpoint.Set(value);
        }

        private static readonly __Value<bool?> _s3UsePathStyle = new __Value<bool?>(() => __config.GetBoolean("s3UsePathStyle"));
        /// <summary>
        /// Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By
        /// default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY).
        /// Specific to the Amazon S3 service.
        /// </summary>
        public static bool? S3UsePathStyle
        {
            get => _s3UsePathStyle.Get();
            set => _s3UsePathStyle.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey"));
        /// <summary>
        /// The secret key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of the AWS console.
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _sharedConfigFiles = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("sharedConfigFiles"));
        /// <summary>
        /// List of paths to shared config files. If not set, defaults to [~/.aws/config].
        /// </summary>
        public static ImmutableArray<string> SharedConfigFiles
        {
            get => _sharedConfigFiles.Get();
            set => _sharedConfigFiles.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _sharedCredentialsFiles = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("sharedCredentialsFiles"));
        /// <summary>
        /// List of paths to shared credentials files. If not set, defaults to [~/.aws/credentials].
        /// </summary>
        public static ImmutableArray<string> SharedCredentialsFiles
        {
            get => _sharedCredentialsFiles.Get();
            set => _sharedCredentialsFiles.Set(value);
        }

        private static readonly __Value<bool?> _skipCredentialsValidation = new __Value<bool?>(() => __config.GetBoolean("skipCredentialsValidation") ?? false);
        /// <summary>
        /// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
        /// available/implemented.
        /// </summary>
        public static bool? SkipCredentialsValidation
        {
            get => _skipCredentialsValidation.Get();
            set => _skipCredentialsValidation.Set(value);
        }

        private static readonly __Value<bool?> _skipMetadataApiCheck = new __Value<bool?>(() => __config.GetBoolean("skipMetadataApiCheck") ?? true);
        /// <summary>
        /// Skip the AWS Metadata API check. Used for AWS API implementations that do not have a metadata api endpoint.
        /// </summary>
        public static bool? SkipMetadataApiCheck
        {
            get => _skipMetadataApiCheck.Get();
            set => _skipMetadataApiCheck.Set(value);
        }

        private static readonly __Value<bool?> _skipRegionValidation = new __Value<bool?>(() => __config.GetBoolean("skipRegionValidation") ?? true);
        /// <summary>
        /// Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
        /// not public (yet).
        /// </summary>
        public static bool? SkipRegionValidation
        {
            get => _skipRegionValidation.Get();
            set => _skipRegionValidation.Set(value);
        }

        private static readonly __Value<bool?> _skipRequestingAccountId = new __Value<bool?>(() => __config.GetBoolean("skipRequestingAccountId"));
        /// <summary>
        /// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
        /// </summary>
        public static bool? SkipRequestingAccountId
        {
            get => _skipRequestingAccountId.Get();
            set => _skipRequestingAccountId.Set(value);
        }

        private static readonly __Value<string?> _stsRegion = new __Value<string?>(() => __config.Get("stsRegion"));
        /// <summary>
        /// The region where AWS STS operations will take place. Examples are us-east-1 and us-west-2.
        /// </summary>
        public static string? StsRegion
        {
            get => _stsRegion.Get();
            set => _stsRegion.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        /// <summary>
        /// session token. A session token is only required if you are using temporary security credentials.
        /// </summary>
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

        private static readonly __Value<bool?> _useDualstackEndpoint = new __Value<bool?>(() => __config.GetBoolean("useDualstackEndpoint"));
        /// <summary>
        /// Resolve an endpoint with DualStack capability
        /// </summary>
        public static bool? UseDualstackEndpoint
        {
            get => _useDualstackEndpoint.Get();
            set => _useDualstackEndpoint.Set(value);
        }

        private static readonly __Value<bool?> _useFipsEndpoint = new __Value<bool?>(() => __config.GetBoolean("useFipsEndpoint"));
        /// <summary>
        /// Resolve an endpoint with FIPS capability
        /// </summary>
        public static bool? UseFipsEndpoint
        {
            get => _useFipsEndpoint.Get();
            set => _useFipsEndpoint.Set(value);
        }

        public static class Types
        {

             public class AssumeRole
             {
                public string? Duration { get; set; } = null!;
                public string? ExternalId { get; set; } = null!;
                public string? Policy { get; set; } = null!;
                public ImmutableArray<string> PolicyArns { get; set; }
                public string? RoleArn { get; set; } = null!;
                public string? SessionName { get; set; } = null!;
                public string? SourceIdentity { get; set; } = null!;
                public ImmutableDictionary<string, string>? Tags { get; set; } = null!;
                public ImmutableArray<string> TransitiveTagKeys { get; set; }
            }

             public class AssumeRoleWithWebIdentity
             {
                public string? Duration { get; set; } = null!;
                public string? Policy { get; set; } = null!;
                public ImmutableArray<string> PolicyArns { get; set; }
                public string? RoleArn { get; set; } = null!;
                public string? SessionName { get; set; } = null!;
                public string? WebIdentityToken { get; set; } = null!;
                public string? WebIdentityTokenFile { get; set; } = null!;
            }

             public class DefaultTags
             {
                public ImmutableDictionary<string, string>? Tags { get; set; } = null!;
            }

             public class Endpoints
             {
                public string? Accessanalyzer { get; set; } = null!;
                public string? Account { get; set; } = null!;
                public string? Acm { get; set; } = null!;
                public string? Acmpca { get; set; } = null!;
                public string? Amg { get; set; } = null!;
                public string? Amp { get; set; } = null!;
                public string? Amplify { get; set; } = null!;
                public string? Apigateway { get; set; } = null!;
                public string? Apigatewayv2 { get; set; } = null!;
                public string? Appautoscaling { get; set; } = null!;
                public string? Appconfig { get; set; } = null!;
                public string? Appfabric { get; set; } = null!;
                public string? Appflow { get; set; } = null!;
                public string? Appintegrations { get; set; } = null!;
                public string? Appintegrationsservice { get; set; } = null!;
                public string? Applicationautoscaling { get; set; } = null!;
                public string? Applicationinsights { get; set; } = null!;
                public string? Appmesh { get; set; } = null!;
                public string? Apprunner { get; set; } = null!;
                public string? Appstream { get; set; } = null!;
                public string? Appsync { get; set; } = null!;
                public string? Athena { get; set; } = null!;
                public string? Auditmanager { get; set; } = null!;
                public string? Autoscaling { get; set; } = null!;
                public string? Autoscalingplans { get; set; } = null!;
                public string? Backup { get; set; } = null!;
                public string? Batch { get; set; } = null!;
                public string? Beanstalk { get; set; } = null!;
                public string? Bedrock { get; set; } = null!;
                public string? Budgets { get; set; } = null!;
                public string? Ce { get; set; } = null!;
                public string? Chime { get; set; } = null!;
                public string? Chimesdkmediapipelines { get; set; } = null!;
                public string? Chimesdkvoice { get; set; } = null!;
                public string? Cleanrooms { get; set; } = null!;
                public string? Cloud9 { get; set; } = null!;
                public string? Cloudcontrol { get; set; } = null!;
                public string? Cloudcontrolapi { get; set; } = null!;
                public string? Cloudformation { get; set; } = null!;
                public string? Cloudfront { get; set; } = null!;
                public string? Cloudhsm { get; set; } = null!;
                public string? Cloudhsmv2 { get; set; } = null!;
                public string? Cloudsearch { get; set; } = null!;
                public string? Cloudtrail { get; set; } = null!;
                public string? Cloudwatch { get; set; } = null!;
                public string? Cloudwatchevents { get; set; } = null!;
                public string? Cloudwatchevidently { get; set; } = null!;
                public string? Cloudwatchlog { get; set; } = null!;
                public string? Cloudwatchlogs { get; set; } = null!;
                public string? Cloudwatchobservabilityaccessmanager { get; set; } = null!;
                public string? Cloudwatchrum { get; set; } = null!;
                public string? Codeartifact { get; set; } = null!;
                public string? Codebuild { get; set; } = null!;
                public string? Codecatalyst { get; set; } = null!;
                public string? Codecommit { get; set; } = null!;
                public string? Codedeploy { get; set; } = null!;
                public string? Codeguruprofiler { get; set; } = null!;
                public string? Codegurureviewer { get; set; } = null!;
                public string? Codepipeline { get; set; } = null!;
                public string? Codestarconnections { get; set; } = null!;
                public string? Codestarnotifications { get; set; } = null!;
                public string? Cognitoidentity { get; set; } = null!;
                public string? Cognitoidentityprovider { get; set; } = null!;
                public string? Cognitoidp { get; set; } = null!;
                public string? Comprehend { get; set; } = null!;
                public string? Computeoptimizer { get; set; } = null!;
                public string? Config { get; set; } = null!;
                public string? Configservice { get; set; } = null!;
                public string? Connect { get; set; } = null!;
                public string? Connectcases { get; set; } = null!;
                public string? Controltower { get; set; } = null!;
                public string? Costandusagereportservice { get; set; } = null!;
                public string? Costexplorer { get; set; } = null!;
                public string? Cur { get; set; } = null!;
                public string? Customerprofiles { get; set; } = null!;
                public string? Databasemigration { get; set; } = null!;
                public string? Databasemigrationservice { get; set; } = null!;
                public string? Dataexchange { get; set; } = null!;
                public string? Datapipeline { get; set; } = null!;
                public string? Datasync { get; set; } = null!;
                public string? Dax { get; set; } = null!;
                public string? Deploy { get; set; } = null!;
                public string? Detective { get; set; } = null!;
                public string? Devicefarm { get; set; } = null!;
                public string? Directconnect { get; set; } = null!;
                public string? Directoryservice { get; set; } = null!;
                public string? Dlm { get; set; } = null!;
                public string? Dms { get; set; } = null!;
                public string? Docdb { get; set; } = null!;
                public string? Docdbelastic { get; set; } = null!;
                public string? Ds { get; set; } = null!;
                public string? Dynamodb { get; set; } = null!;
                public string? Ec2 { get; set; } = null!;
                public string? Ecr { get; set; } = null!;
                public string? Ecrpublic { get; set; } = null!;
                public string? Ecs { get; set; } = null!;
                public string? Efs { get; set; } = null!;
                public string? Eks { get; set; } = null!;
                public string? Elasticache { get; set; } = null!;
                public string? Elasticbeanstalk { get; set; } = null!;
                public string? Elasticloadbalancing { get; set; } = null!;
                public string? Elasticloadbalancingv2 { get; set; } = null!;
                public string? Elasticsearch { get; set; } = null!;
                public string? Elasticsearchservice { get; set; } = null!;
                public string? Elastictranscoder { get; set; } = null!;
                public string? Elb { get; set; } = null!;
                public string? Elbv2 { get; set; } = null!;
                public string? Emr { get; set; } = null!;
                public string? Emrcontainers { get; set; } = null!;
                public string? Emrserverless { get; set; } = null!;
                public string? Es { get; set; } = null!;
                public string? Eventbridge { get; set; } = null!;
                public string? Events { get; set; } = null!;
                public string? Evidently { get; set; } = null!;
                public string? Finspace { get; set; } = null!;
                public string? Firehose { get; set; } = null!;
                public string? Fis { get; set; } = null!;
                public string? Fms { get; set; } = null!;
                public string? Fsx { get; set; } = null!;
                public string? Gamelift { get; set; } = null!;
                public string? Glacier { get; set; } = null!;
                public string? Globalaccelerator { get; set; } = null!;
                public string? Glue { get; set; } = null!;
                public string? Grafana { get; set; } = null!;
                public string? Greengrass { get; set; } = null!;
                public string? Guardduty { get; set; } = null!;
                public string? Healthlake { get; set; } = null!;
                public string? Iam { get; set; } = null!;
                public string? Identitystore { get; set; } = null!;
                public string? Imagebuilder { get; set; } = null!;
                public string? Inspector { get; set; } = null!;
                public string? Inspector2 { get; set; } = null!;
                public string? Inspectorv2 { get; set; } = null!;
                public string? Internetmonitor { get; set; } = null!;
                public string? Iot { get; set; } = null!;
                public string? Iotanalytics { get; set; } = null!;
                public string? Iotevents { get; set; } = null!;
                public string? Ivs { get; set; } = null!;
                public string? Ivschat { get; set; } = null!;
                public string? Kafka { get; set; } = null!;
                public string? Kafkaconnect { get; set; } = null!;
                public string? Kendra { get; set; } = null!;
                public string? Keyspaces { get; set; } = null!;
                public string? Kinesis { get; set; } = null!;
                public string? Kinesisanalytics { get; set; } = null!;
                public string? Kinesisanalyticsv2 { get; set; } = null!;
                public string? Kinesisvideo { get; set; } = null!;
                public string? Kms { get; set; } = null!;
                public string? Lakeformation { get; set; } = null!;
                public string? Lambda { get; set; } = null!;
                public string? Lex { get; set; } = null!;
                public string? Lexmodelbuilding { get; set; } = null!;
                public string? Lexmodelbuildingservice { get; set; } = null!;
                public string? Lexmodels { get; set; } = null!;
                public string? Lexmodelsv2 { get; set; } = null!;
                public string? Lexv2models { get; set; } = null!;
                public string? Licensemanager { get; set; } = null!;
                public string? Lightsail { get; set; } = null!;
                public string? Location { get; set; } = null!;
                public string? Locationservice { get; set; } = null!;
                public string? Logs { get; set; } = null!;
                public string? Lookoutmetrics { get; set; } = null!;
                public string? Macie2 { get; set; } = null!;
                public string? Managedgrafana { get; set; } = null!;
                public string? Mediaconnect { get; set; } = null!;
                public string? Mediaconvert { get; set; } = null!;
                public string? Medialive { get; set; } = null!;
                public string? Mediapackage { get; set; } = null!;
                public string? Mediapackagev2 { get; set; } = null!;
                public string? Mediastore { get; set; } = null!;
                public string? Memorydb { get; set; } = null!;
                public string? Mq { get; set; } = null!;
                public string? Msk { get; set; } = null!;
                public string? Mwaa { get; set; } = null!;
                public string? Neptune { get; set; } = null!;
                public string? Networkfirewall { get; set; } = null!;
                public string? Networkmanager { get; set; } = null!;
                public string? Oam { get; set; } = null!;
                public string? Opensearch { get; set; } = null!;
                public string? Opensearchingestion { get; set; } = null!;
                public string? Opensearchserverless { get; set; } = null!;
                public string? Opensearchservice { get; set; } = null!;
                public string? Opsworks { get; set; } = null!;
                public string? Organizations { get; set; } = null!;
                public string? Osis { get; set; } = null!;
                public string? Outposts { get; set; } = null!;
                public string? Pinpoint { get; set; } = null!;
                public string? Pipes { get; set; } = null!;
                public string? Polly { get; set; } = null!;
                public string? Pricing { get; set; } = null!;
                public string? Prometheus { get; set; } = null!;
                public string? Prometheusservice { get; set; } = null!;
                public string? Qldb { get; set; } = null!;
                public string? Quicksight { get; set; } = null!;
                public string? Ram { get; set; } = null!;
                public string? Rbin { get; set; } = null!;
                public string? Rds { get; set; } = null!;
                public string? Recyclebin { get; set; } = null!;
                public string? Redshift { get; set; } = null!;
                public string? Redshiftdata { get; set; } = null!;
                public string? Redshiftdataapiservice { get; set; } = null!;
                public string? Redshiftserverless { get; set; } = null!;
                public string? Resourceexplorer2 { get; set; } = null!;
                public string? Resourcegroups { get; set; } = null!;
                public string? Resourcegroupstagging { get; set; } = null!;
                public string? Resourcegroupstaggingapi { get; set; } = null!;
                public string? Rolesanywhere { get; set; } = null!;
                public string? Route53 { get; set; } = null!;
                public string? Route53domains { get; set; } = null!;
                public string? Route53recoverycontrolconfig { get; set; } = null!;
                public string? Route53recoveryreadiness { get; set; } = null!;
                public string? Route53resolver { get; set; } = null!;
                public string? Rum { get; set; } = null!;
                public string? S3 { get; set; } = null!;
                public string? S3api { get; set; } = null!;
                public string? S3control { get; set; } = null!;
                public string? S3outposts { get; set; } = null!;
                public string? Sagemaker { get; set; } = null!;
                public string? Scheduler { get; set; } = null!;
                public string? Schemas { get; set; } = null!;
                public string? Sdb { get; set; } = null!;
                public string? Secretsmanager { get; set; } = null!;
                public string? Securityhub { get; set; } = null!;
                public string? Securitylake { get; set; } = null!;
                public string? Serverlessapplicationrepository { get; set; } = null!;
                public string? Serverlessapprepo { get; set; } = null!;
                public string? Serverlessrepo { get; set; } = null!;
                public string? Servicecatalog { get; set; } = null!;
                public string? Servicediscovery { get; set; } = null!;
                public string? Servicequotas { get; set; } = null!;
                public string? Ses { get; set; } = null!;
                public string? Sesv2 { get; set; } = null!;
                public string? Sfn { get; set; } = null!;
                public string? Shield { get; set; } = null!;
                public string? Signer { get; set; } = null!;
                public string? Simpledb { get; set; } = null!;
                public string? Sns { get; set; } = null!;
                public string? Sqs { get; set; } = null!;
                public string? Ssm { get; set; } = null!;
                public string? Ssmcontacts { get; set; } = null!;
                public string? Ssmincidents { get; set; } = null!;
                public string? Sso { get; set; } = null!;
                public string? Ssoadmin { get; set; } = null!;
                public string? Stepfunctions { get; set; } = null!;
                public string? Storagegateway { get; set; } = null!;
                public string? Sts { get; set; } = null!;
                public string? Swf { get; set; } = null!;
                public string? Synthetics { get; set; } = null!;
                public string? Timestreamwrite { get; set; } = null!;
                public string? Transcribe { get; set; } = null!;
                public string? Transcribeservice { get; set; } = null!;
                public string? Transfer { get; set; } = null!;
                public string? Verifiedpermissions { get; set; } = null!;
                public string? Vpclattice { get; set; } = null!;
                public string? Waf { get; set; } = null!;
                public string? Wafregional { get; set; } = null!;
                public string? Wafv2 { get; set; } = null!;
                public string? Worklink { get; set; } = null!;
                public string? Workspaces { get; set; } = null!;
                public string? Xray { get; set; } = null!;
            }

             public class IgnoreTags
             {
                public ImmutableArray<string> KeyPrefixes { get; set; }
                public ImmutableArray<string> Keys { get; set; }
            }
        }
    }
}
