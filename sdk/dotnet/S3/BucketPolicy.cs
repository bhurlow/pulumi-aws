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
    /// Attaches a policy to an S3 bucket resource.
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
    ///     var example = new Aws.S3.BucketV2("example");
    /// 
    ///     var allowAccessFromAnotherAccountPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "123456789012",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "s3:GetObject",
    ///                     "s3:ListBucket",
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     example.Arn,
    ///                     $"{example.Arn}/*",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var allowAccessFromAnotherAccountBucketPolicy = new Aws.S3.BucketPolicy("allowAccessFromAnotherAccountBucketPolicy", new()
    ///     {
    ///         Bucket = example.Id,
    ///         Policy = allowAccessFromAnotherAccountPolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_s3_bucket_policy.allow_access_from_another_account
    /// 
    ///  id = "my-tf-test-bucket" } Using `pulumi import`, import S3 bucket policies using the bucket name. For exampleconsole % pulumi import aws_s3_bucket_policy.allow_access_from_another_account my-tf-test-bucket
    /// </summary>
    [AwsResourceType("aws:s3/bucketPolicy:BucketPolicy")]
    public partial class BucketPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the bucket to which to apply the policy.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `aws.iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a BucketPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketPolicy(string name, BucketPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:s3/bucketPolicy:BucketPolicy", name, args ?? new BucketPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketPolicy(string name, Input<string> id, BucketPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3/bucketPolicy:BucketPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketPolicy Get(string name, Input<string> id, BucketPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketPolicy(name, id, state, options);
        }
    }

    public sealed class BucketPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the bucket to which to apply the policy.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `aws.iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public BucketPolicyArgs()
        {
        }
        public static new BucketPolicyArgs Empty => new BucketPolicyArgs();
    }

    public sealed class BucketPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the bucket to which to apply the policy.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `aws.iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public BucketPolicyState()
        {
        }
        public static new BucketPolicyState Empty => new BucketPolicyState();
    }
}
