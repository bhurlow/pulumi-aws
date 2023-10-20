// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3
{
    /// <summary>
    /// Provides an independent configuration resource for S3 bucket [lifecycle configuration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html).
    /// 
    /// An S3 Lifecycle configuration consists of one or more Lifecycle rules. Each rule consists of the following:
    /// 
    /// * Rule metadata (`id` and `status`)
    /// * Filter identifying objects to which the rule applies
    /// * One or more transition or expiration actions
    /// 
    /// For more information see the Amazon S3 User Guide on [`Lifecycle Configuration Elements`](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html).
    /// 
    /// &gt; **NOTE:** S3 Buckets only support a single lifecycle configuration. Declaring multiple `aws.s3.BucketLifecycleConfigurationV2` resources to the same S3 Bucket will cause a perpetual difference in configuration.
    /// 
    /// &gt; **NOTE:** Lifecycle configurations may take some time to fully propagate to all AWS S3 systems.
    /// Running Pulumi operations shortly after creating a lifecycle configuration may result in changes that affect configuration idempotence.
    /// See the Amazon S3 User Guide on [setting lifecycle configuration on a bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/how-to-set-lifecycle-configuration-intro.html).
    /// 
    /// ## Example Usage
    /// ### With neither a filter nor prefix specified
    /// 
    /// The Lifecycle rule applies to a subset of objects based on the key name prefix (`""`).
    /// 
    /// This configuration is intended to replicate the default behavior of the `lifecycle_rule`
    /// parameter in the AWS Provider `aws.s3.BucketV2` resource prior to `v4.0`.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Specifying an empty filter
    /// 
    /// The Lifecycle rule applies to all objects in the bucket.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = null,
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Specifying a filter using key prefixes
    /// 
    /// The Lifecycle rule applies to a subset of objects based on the key name prefix (`logs/`).
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     Prefix = "logs/",
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// If you want to apply a Lifecycle action to a subset of objects based on different key name prefixes, specify separate rules.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     Prefix = "logs/",
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-2",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     Prefix = "tmp/",
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Specifying a filter based on an object tag
    /// 
    /// The Lifecycle rule specifies a filter based on a tag key and value. The rule then applies only to a subset of objects with the specific tag.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     Tag = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterTagArgs
    ///                     {
    ///                         Key = "Name",
    ///                         Value = "Staging",
    ///                     },
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Specifying a filter based on multiple tags
    /// 
    /// The Lifecycle rule directs Amazon S3 to perform lifecycle actions on objects with two tags (with the specific tag keys and values). Notice `tags` is wrapped in the `and` configuration block.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
    ///                     {
    ///                         Tags = 
    ///                         {
    ///                             { "Key1", "Value1" },
    ///                             { "Key2", "Value2" },
    ///                         },
    ///                     },
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Specifying a filter based on both prefix and one or more tags
    /// 
    /// The Lifecycle rule directs Amazon S3 to perform lifecycle actions on objects with the specified prefix and two tags (with the specific tag keys and values). Notice both `prefix` and `tags` are wrapped in the `and` configuration block.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
    ///                     {
    ///                         Prefix = "logs/",
    ///                         Tags = 
    ///                         {
    ///                             { "Key1", "Value1" },
    ///                             { "Key2", "Value2" },
    ///                         },
    ///                     },
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Specifying a filter based on object size
    /// 
    /// Object size values are in bytes. Maximum filter size is 5TB. Some storage classes have minimum object size limitations, for more information, see [Comparing the Amazon S3 storage classes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html#sc-compare).
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     ObjectSizeGreaterThan = "500",
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Specifying a filter based on object size range and prefix
    /// 
    /// The `object_size_greater_than` must be less than the `object_size_less_than`. Notice both the object size range and prefix are wrapped in the `and` configuration block.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    ///     {
    ///         Bucket = aws_s3_bucket.Bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "rule-1",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
    ///                     {
    ///                         Prefix = "logs/",
    ///                         ObjectSizeGreaterThan = 500,
    ///                         ObjectSizeLessThan = 64000,
    ///                     },
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Creating a Lifecycle Configuration for a bucket with versioning
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bucket = new Aws.S3.BucketV2("bucket");
    /// 
    ///     var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    ///     {
    ///         Bucket = bucket.Id,
    ///         Acl = "private",
    ///     });
    /// 
    ///     var bucket_config = new Aws.S3.BucketLifecycleConfigurationV2("bucket-config", new()
    ///     {
    ///         Bucket = bucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "log",
    ///                 Expiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleExpirationArgs
    ///                 {
    ///                     Days = 90,
    ///                 },
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
    ///                     {
    ///                         Prefix = "log/",
    ///                         Tags = 
    ///                         {
    ///                             { "rule", "log" },
    ///                             { "autoclean", "true" },
    ///                         },
    ///                     },
    ///                 },
    ///                 Status = "Enabled",
    ///                 Transitions = new[]
    ///                 {
    ///                     new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleTransitionArgs
    ///                     {
    ///                         Days = 30,
    ///                         StorageClass = "STANDARD_IA",
    ///                     },
    ///                     new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleTransitionArgs
    ///                     {
    ///                         Days = 60,
    ///                         StorageClass = "GLACIER",
    ///                     },
    ///                 },
    ///             },
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "tmp",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     Prefix = "tmp/",
    ///                 },
    ///                 Expiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleExpirationArgs
    ///                 {
    ///                     Date = "2023-01-13T00:00:00Z",
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var versioningBucket = new Aws.S3.BucketV2("versioningBucket");
    /// 
    ///     var versioningBucketAcl = new Aws.S3.BucketAclV2("versioningBucketAcl", new()
    ///     {
    ///         Bucket = versioningBucket.Id,
    ///         Acl = "private",
    ///     });
    /// 
    ///     var versioning = new Aws.S3.BucketVersioningV2("versioning", new()
    ///     {
    ///         Bucket = versioningBucket.Id,
    ///         VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningV2VersioningConfigurationArgs
    ///         {
    ///             Status = "Enabled",
    ///         },
    ///     });
    /// 
    ///     var versioning_bucket_config = new Aws.S3.BucketLifecycleConfigurationV2("versioning-bucket-config", new()
    ///     {
    ///         Bucket = versioningBucket.Id,
    ///         Rules = new[]
    ///         {
    ///             new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
    ///             {
    ///                 Id = "config",
    ///                 Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
    ///                 {
    ///                     Prefix = "config/",
    ///                 },
    ///                 NoncurrentVersionExpiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleNoncurrentVersionExpirationArgs
    ///                 {
    ///                     NoncurrentDays = 90,
    ///                 },
    ///                 NoncurrentVersionTransitions = new[]
    ///                 {
    ///                     new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs
    ///                     {
    ///                         NoncurrentDays = 30,
    ///                         StorageClass = "STANDARD_IA",
    ///                     },
    ///                     new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs
    ///                     {
    ///                         NoncurrentDays = 60,
    ///                         StorageClass = "GLACIER",
    ///                     },
    ///                 },
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             versioning,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):
    /// 
    /// __Using `pulumi import` to import__ S3 bucket lifecycle configuration using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:
    /// 
    /// If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2 example bucket-name
    /// ```
    ///  If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2 example bucket-name,123456789012
    /// ```
    /// </summary>
    [AwsResourceType("aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2")]
    public partial class BucketLifecycleConfigurationV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the source S3 bucket you want Amazon S3 to monitor.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
        /// </summary>
        [Output("expectedBucketOwner")]
        public Output<string?> ExpectedBucketOwner { get; private set; } = null!;

        /// <summary>
        /// List of configuration blocks describing the rules managing the replication. See below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.BucketLifecycleConfigurationV2Rule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a BucketLifecycleConfigurationV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketLifecycleConfigurationV2(string name, BucketLifecycleConfigurationV2Args args, CustomResourceOptions? options = null)
            : base("aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2", name, args ?? new BucketLifecycleConfigurationV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private BucketLifecycleConfigurationV2(string name, Input<string> id, BucketLifecycleConfigurationV2State? state = null, CustomResourceOptions? options = null)
            : base("aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketLifecycleConfigurationV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketLifecycleConfigurationV2 Get(string name, Input<string> id, BucketLifecycleConfigurationV2State? state = null, CustomResourceOptions? options = null)
        {
            return new BucketLifecycleConfigurationV2(name, id, state, options);
        }
    }

    public sealed class BucketLifecycleConfigurationV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the source S3 bucket you want Amazon S3 to monitor.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
        /// </summary>
        [Input("expectedBucketOwner")]
        public Input<string>? ExpectedBucketOwner { get; set; }

        [Input("rules", required: true)]
        private InputList<Inputs.BucketLifecycleConfigurationV2RuleArgs>? _rules;

        /// <summary>
        /// List of configuration blocks describing the rules managing the replication. See below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleConfigurationV2RuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketLifecycleConfigurationV2RuleArgs>());
            set => _rules = value;
        }

        public BucketLifecycleConfigurationV2Args()
        {
        }
        public static new BucketLifecycleConfigurationV2Args Empty => new BucketLifecycleConfigurationV2Args();
    }

    public sealed class BucketLifecycleConfigurationV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the source S3 bucket you want Amazon S3 to monitor.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
        /// </summary>
        [Input("expectedBucketOwner")]
        public Input<string>? ExpectedBucketOwner { get; set; }

        [Input("rules")]
        private InputList<Inputs.BucketLifecycleConfigurationV2RuleGetArgs>? _rules;

        /// <summary>
        /// List of configuration blocks describing the rules managing the replication. See below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleConfigurationV2RuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketLifecycleConfigurationV2RuleGetArgs>());
            set => _rules = value;
        }

        public BucketLifecycleConfigurationV2State()
        {
        }
        public static new BucketLifecycleConfigurationV2State Empty => new BucketLifecycleConfigurationV2State();
    }
}
