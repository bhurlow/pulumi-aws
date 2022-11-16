// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3
{
    public static class GetObject
    {
        /// <summary>
        /// The S3 object data source allows access to the metadata and
        /// _optionally_ (see below) content of an object stored inside S3 bucket.
        /// 
        /// &gt; **Note:** The content of an object (`body` field) is available only for objects which have a human-readable `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favour of metadata.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example retrieves a text object (which must have a `Content-Type`
        /// value starting with `text/`) and uses it as the `user_data` for an EC2 instance:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var bootstrapScript = Aws.S3.GetObject.Invoke(new()
        ///     {
        ///         Bucket = "ourcorp-deploy-config",
        ///         Key = "ec2-bootstrap-script.sh",
        ///     });
        /// 
        ///     var example = new Aws.Ec2.Instance("example", new()
        ///     {
        ///         InstanceType = "t2.micro",
        ///         Ami = "ami-2757f631",
        ///         UserData = bootstrapScript.Apply(getObjectResult =&gt; getObjectResult.Body),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// The following, more-complex example retrieves only the metadata for a zip
        /// file stored in S3, which is then used to pass the most recent `version_id`
        /// to AWS Lambda for use as a function implementation. More information about
        /// Lambda functions is available in the documentation for
        /// `aws.lambda.Function`.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var lambda = Aws.S3.GetObject.Invoke(new()
        ///     {
        ///         Bucket = "ourcorp-lambda-functions",
        ///         Key = "hello-world.zip",
        ///     });
        /// 
        ///     var testLambda = new Aws.Lambda.Function("testLambda", new()
        ///     {
        ///         S3Bucket = lambda.Apply(getObjectResult =&gt; getObjectResult.Bucket),
        ///         S3Key = lambda.Apply(getObjectResult =&gt; getObjectResult.Key),
        ///         S3ObjectVersion = lambda.Apply(getObjectResult =&gt; getObjectResult.VersionId),
        ///         Role = aws_iam_role.Iam_for_lambda.Arn,
        ///         Handler = "exports.test",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetObjectResult> InvokeAsync(GetObjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetObjectResult>("aws:s3/getObject:getObject", args ?? new GetObjectArgs(), options.WithDefaults());

        /// <summary>
        /// The S3 object data source allows access to the metadata and
        /// _optionally_ (see below) content of an object stored inside S3 bucket.
        /// 
        /// &gt; **Note:** The content of an object (`body` field) is available only for objects which have a human-readable `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favour of metadata.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example retrieves a text object (which must have a `Content-Type`
        /// value starting with `text/`) and uses it as the `user_data` for an EC2 instance:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var bootstrapScript = Aws.S3.GetObject.Invoke(new()
        ///     {
        ///         Bucket = "ourcorp-deploy-config",
        ///         Key = "ec2-bootstrap-script.sh",
        ///     });
        /// 
        ///     var example = new Aws.Ec2.Instance("example", new()
        ///     {
        ///         InstanceType = "t2.micro",
        ///         Ami = "ami-2757f631",
        ///         UserData = bootstrapScript.Apply(getObjectResult =&gt; getObjectResult.Body),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// The following, more-complex example retrieves only the metadata for a zip
        /// file stored in S3, which is then used to pass the most recent `version_id`
        /// to AWS Lambda for use as a function implementation. More information about
        /// Lambda functions is available in the documentation for
        /// `aws.lambda.Function`.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var lambda = Aws.S3.GetObject.Invoke(new()
        ///     {
        ///         Bucket = "ourcorp-lambda-functions",
        ///         Key = "hello-world.zip",
        ///     });
        /// 
        ///     var testLambda = new Aws.Lambda.Function("testLambda", new()
        ///     {
        ///         S3Bucket = lambda.Apply(getObjectResult =&gt; getObjectResult.Bucket),
        ///         S3Key = lambda.Apply(getObjectResult =&gt; getObjectResult.Key),
        ///         S3ObjectVersion = lambda.Apply(getObjectResult =&gt; getObjectResult.VersionId),
        ///         Role = aws_iam_role.Iam_for_lambda.Arn,
        ///         Handler = "exports.test",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetObjectResult> Invoke(GetObjectInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetObjectResult>("aws:s3/getObject:getObject", args ?? new GetObjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetObjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the bucket to read the object from. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
        /// </summary>
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        /// <summary>
        /// Full path to the object inside the bucket
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("range")]
        public string? Range { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of tags assigned to the object.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specific version ID of the object returned (defaults to latest version)
        /// </summary>
        [Input("versionId")]
        public string? VersionId { get; set; }

        public GetObjectArgs()
        {
        }
        public static new GetObjectArgs Empty => new GetObjectArgs();
    }

    public sealed class GetObjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the bucket to read the object from. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Full path to the object inside the bucket
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("range")]
        public Input<string>? Range { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags assigned to the object.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specific version ID of the object returned (defaults to latest version)
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public GetObjectInvokeArgs()
        {
        }
        public static new GetObjectInvokeArgs Empty => new GetObjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetObjectResult
    {
        /// <summary>
        /// Object data (see **limitations above** to understand cases in which this field is actually available)
        /// </summary>
        public readonly string Body;
        public readonly string Bucket;
        /// <summary>
        /// (Optional) Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
        /// </summary>
        public readonly bool BucketKeyEnabled;
        /// <summary>
        /// Caching behavior along the request/reply chain.
        /// </summary>
        public readonly string CacheControl;
        /// <summary>
        /// Presentational information for the object.
        /// </summary>
        public readonly string ContentDisposition;
        /// <summary>
        /// What content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
        /// </summary>
        public readonly string ContentEncoding;
        /// <summary>
        /// Language the content is in.
        /// </summary>
        public readonly string ContentLanguage;
        /// <summary>
        /// Size of the body in bytes.
        /// </summary>
        public readonly int ContentLength;
        /// <summary>
        /// Standard MIME type describing the format of the object data.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
        /// </summary>
        public readonly string Expiration;
        /// <summary>
        /// Date and time at which the object is no longer cacheable.
        /// </summary>
        public readonly string Expires;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Key;
        /// <summary>
        /// Last modified date of the object in RFC1123 format (e.g., `Mon, 02 Jan 2006 15:04:05 MST`)
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// Map of metadata stored with the object in S3
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Indicates whether this object has an active [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds). This field is only returned if you have permission to view an object's legal hold status.
        /// </summary>
        public readonly string ObjectLockLegalHoldStatus;
        /// <summary>
        /// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) currently in place for this object.
        /// </summary>
        public readonly string ObjectLockMode;
        /// <summary>
        /// The date and time when this object's object lock will expire.
        /// </summary>
        public readonly string ObjectLockRetainUntilDate;
        public readonly string? Range;
        /// <summary>
        /// If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.
        /// </summary>
        public readonly string ServerSideEncryption;
        /// <summary>
        /// If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.
        /// </summary>
        public readonly string SseKmsKeyId;
        /// <summary>
        /// [Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for `Standard` storage class objects.
        /// </summary>
        public readonly string StorageClass;
        /// <summary>
        /// Map of tags assigned to the object.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Latest version ID of the object returned.
        /// </summary>
        public readonly string VersionId;
        /// <summary>
        /// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.
        /// </summary>
        public readonly string WebsiteRedirectLocation;

        [OutputConstructor]
        private GetObjectResult(
            string body,

            string bucket,

            bool bucketKeyEnabled,

            string cacheControl,

            string contentDisposition,

            string contentEncoding,

            string contentLanguage,

            int contentLength,

            string contentType,

            string etag,

            string expiration,

            string expires,

            string id,

            string key,

            string lastModified,

            ImmutableDictionary<string, string> metadata,

            string objectLockLegalHoldStatus,

            string objectLockMode,

            string objectLockRetainUntilDate,

            string? range,

            string serverSideEncryption,

            string sseKmsKeyId,

            string storageClass,

            ImmutableDictionary<string, string> tags,

            string versionId,

            string websiteRedirectLocation)
        {
            Body = body;
            Bucket = bucket;
            BucketKeyEnabled = bucketKeyEnabled;
            CacheControl = cacheControl;
            ContentDisposition = contentDisposition;
            ContentEncoding = contentEncoding;
            ContentLanguage = contentLanguage;
            ContentLength = contentLength;
            ContentType = contentType;
            Etag = etag;
            Expiration = expiration;
            Expires = expires;
            Id = id;
            Key = key;
            LastModified = lastModified;
            Metadata = metadata;
            ObjectLockLegalHoldStatus = objectLockLegalHoldStatus;
            ObjectLockMode = objectLockMode;
            ObjectLockRetainUntilDate = objectLockRetainUntilDate;
            Range = range;
            ServerSideEncryption = serverSideEncryption;
            SseKmsKeyId = sseKmsKeyId;
            StorageClass = storageClass;
            Tags = tags;
            VersionId = versionId;
            WebsiteRedirectLocation = websiteRedirectLocation;
        }
    }
}
