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
    /// Provides a S3 bucket object resource.
    /// 
    /// ## Example Usage
    /// ### Encrypting with KMS Key
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var examplekms = new Aws.Kms.Key("examplekms", new Aws.Kms.KeyArgs
    ///         {
    ///             Description = "KMS key 1",
    ///             DeletionWindowInDays = 7,
    ///         });
    ///         var examplebucket = new Aws.S3.Bucket("examplebucket", new Aws.S3.BucketArgs
    ///         {
    ///             Acl = "private",
    ///         });
    ///         var examplebucketObject = new Aws.S3.BucketObject("examplebucketObject", new Aws.S3.BucketObjectArgs
    ///         {
    ///             Key = "someobject",
    ///             Bucket = examplebucket.Id,
    ///             Source = new FileAsset("index.html"),
    ///             KmsKeyId = examplekms.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Server Side Encryption with S3 Default Master Key
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var examplebucket = new Aws.S3.Bucket("examplebucket", new Aws.S3.BucketArgs
    ///         {
    ///             Acl = "private",
    ///         });
    ///         var examplebucketObject = new Aws.S3.BucketObject("examplebucketObject", new Aws.S3.BucketObjectArgs
    ///         {
    ///             Key = "someobject",
    ///             Bucket = examplebucket.Id,
    ///             Source = new FileAsset("index.html"),
    ///             ServerSideEncryption = "aws:kms",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Server Side Encryption with AWS-Managed Key
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var examplebucket = new Aws.S3.Bucket("examplebucket", new Aws.S3.BucketArgs
    ///         {
    ///             Acl = "private",
    ///         });
    ///         var examplebucketObject = new Aws.S3.BucketObject("examplebucketObject", new Aws.S3.BucketObjectArgs
    ///         {
    ///             Key = "someobject",
    ///             Bucket = examplebucket.Id,
    ///             Source = new FileAsset("index.html"),
    ///             ServerSideEncryption = "AES256",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### S3 Object Lock
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var examplebucket = new Aws.S3.Bucket("examplebucket", new Aws.S3.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Versioning = new Aws.S3.Inputs.BucketVersioningArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///             ObjectLockConfiguration = new Aws.S3.Inputs.BucketObjectLockConfigurationArgs
    ///             {
    ///                 ObjectLockEnabled = "Enabled",
    ///             },
    ///         });
    ///         var examplebucketObject = new Aws.S3.BucketObject("examplebucketObject", new Aws.S3.BucketObjectArgs
    ///         {
    ///             Key = "someobject",
    ///             Bucket = examplebucket.Id,
    ///             Source = new FileAsset("important.txt"),
    ///             ObjectLockLegalHoldStatus = "ON",
    ///             ObjectLockMode = "GOVERNANCE",
    ///             ObjectLockRetainUntilDate = "2021-12-31T23:59:60Z",
    ///             ForceDestroy = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Objects can be imported using the `id`. The `id` is the bucket name and the key together e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3/bucketObject:BucketObject object some-bucket-name/some/key.txt
    /// ```
    /// 
    ///  Additionally, s3 url syntax can be used, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3/bucketObject:BucketObject object s3://some-bucket-name/some/key.txt
    /// ```
    /// </summary>
    [AwsResourceType("aws:s3/bucketObject:BucketObject")]
    public partial class BucketObject : Pulumi.CustomResource
    {
        /// <summary>
        /// [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Defaults to `private`.
        /// </summary>
        [Output("acl")]
        public Output<string?> Acl { get; private set; } = null!;

        /// <summary>
        /// Name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
        /// </summary>
        [Output("bucketKeyEnabled")]
        public Output<bool> BucketKeyEnabled { get; private set; } = null!;

        /// <summary>
        /// Caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
        /// </summary>
        [Output("cacheControl")]
        public Output<string?> CacheControl { get; private set; } = null!;

        /// <summary>
        /// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
        /// </summary>
        [Output("contentBase64")]
        public Output<string?> ContentBase64 { get; private set; } = null!;

        /// <summary>
        /// Presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
        /// </summary>
        [Output("contentDisposition")]
        public Output<string?> ContentDisposition { get; private set; } = null!;

        /// <summary>
        /// Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
        /// </summary>
        [Output("contentEncoding")]
        public Output<string?> ContentEncoding { get; private set; } = null!;

        /// <summary>
        /// Language the content is in e.g. en-US or en-GB.
        /// </summary>
        [Output("contentLanguage")]
        public Output<string?> ContentLanguage { get; private set; } = null!;

        /// <summary>
        /// Standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// Triggers updates when the value changes. The only meaningful value is `filemd5("path/to/file")`. This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"` (see `source_hash` instead).
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Whether to allow the object to be deleted by removing any legal hold on any object version. Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// Name of the object once it is in the bucket.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// ARN of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the `aws.kms.Key` resource, use the `arn` attribute. If referencing the `aws.kms.Alias` data source or resource, use the `target_key_arn` attribute. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// [Legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
        /// </summary>
        [Output("objectLockLegalHoldStatus")]
        public Output<string?> ObjectLockLegalHoldStatus { get; private set; } = null!;

        /// <summary>
        /// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
        /// </summary>
        [Output("objectLockMode")]
        public Output<string?> ObjectLockMode { get; private set; } = null!;

        /// <summary>
        /// Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
        /// </summary>
        [Output("objectLockRetainUntilDate")]
        public Output<string?> ObjectLockRetainUntilDate { get; private set; } = null!;

        /// <summary>
        /// Server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
        /// </summary>
        [Output("serverSideEncryption")]
        public Output<string> ServerSideEncryption { get; private set; } = null!;

        /// <summary>
        /// Path to a file that will be read and uploaded as raw bytes for the object content.
        /// </summary>
        [Output("source")]
        public Output<AssetOrArchive?> Source { get; private set; } = null!;

        /// <summary>
        /// Triggers updates like `etag` but useful to address `etag` encryption limitations. Set using `filemd5("path/to/source")`. (The value is only stored in state and not saved by AWS.)
        /// </summary>
        [Output("sourceHash")]
        public Output<string?> SourceHash { get; private set; } = null!;

        /// <summary>
        /// [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
        /// </summary>
        [Output("storageClass")]
        public Output<string> StorageClass { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Unique version ID value for the object, if bucket versioning is enabled.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;

        /// <summary>
        /// Target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
        /// </summary>
        [Output("websiteRedirect")]
        public Output<string?> WebsiteRedirect { get; private set; } = null!;


        /// <summary>
        /// Create a BucketObject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketObject(string name, BucketObjectArgs args, CustomResourceOptions? options = null)
            : base("aws:s3/bucketObject:BucketObject", name, args ?? new BucketObjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketObject(string name, Input<string> id, BucketObjectState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3/bucketObject:BucketObject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketObject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketObject Get(string name, Input<string> id, BucketObjectState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketObject(name, id, state, options);
        }
    }

    public sealed class BucketObjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Defaults to `private`.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        /// <summary>
        /// Name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
        /// </summary>
        [Input("bucketKeyEnabled")]
        public Input<bool>? BucketKeyEnabled { get; set; }

        /// <summary>
        /// Caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
        /// </summary>
        [Input("cacheControl")]
        public Input<string>? CacheControl { get; set; }

        /// <summary>
        /// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
        /// </summary>
        [Input("contentBase64")]
        public Input<string>? ContentBase64 { get; set; }

        /// <summary>
        /// Presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
        /// </summary>
        [Input("contentDisposition")]
        public Input<string>? ContentDisposition { get; set; }

        /// <summary>
        /// Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// Language the content is in e.g. en-US or en-GB.
        /// </summary>
        [Input("contentLanguage")]
        public Input<string>? ContentLanguage { get; set; }

        /// <summary>
        /// Standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// Triggers updates when the value changes. The only meaningful value is `filemd5("path/to/file")`. This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"` (see `source_hash` instead).
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Whether to allow the object to be deleted by removing any legal hold on any object version. Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Name of the object once it is in the bucket.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// ARN of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the `aws.kms.Key` resource, use the `arn` attribute. If referencing the `aws.kms.Alias` data source or resource, use the `target_key_arn` attribute. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// [Legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
        /// </summary>
        [Input("objectLockLegalHoldStatus")]
        public Input<string>? ObjectLockLegalHoldStatus { get; set; }

        /// <summary>
        /// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
        /// </summary>
        [Input("objectLockMode")]
        public Input<string>? ObjectLockMode { get; set; }

        /// <summary>
        /// Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
        /// </summary>
        [Input("objectLockRetainUntilDate")]
        public Input<string>? ObjectLockRetainUntilDate { get; set; }

        /// <summary>
        /// Server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
        /// </summary>
        [Input("serverSideEncryption")]
        public Input<string>? ServerSideEncryption { get; set; }

        /// <summary>
        /// Path to a file that will be read and uploaded as raw bytes for the object content.
        /// </summary>
        [Input("source")]
        public Input<AssetOrArchive>? Source { get; set; }

        /// <summary>
        /// Triggers updates like `etag` but useful to address `etag` encryption limitations. Set using `filemd5("path/to/source")`. (The value is only stored in state and not saved by AWS.)
        /// </summary>
        [Input("sourceHash")]
        public Input<string>? SourceHash { get; set; }

        /// <summary>
        /// [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
        /// </summary>
        [Input("websiteRedirect")]
        public Input<string>? WebsiteRedirect { get; set; }

        public BucketObjectArgs()
        {
        }
    }

    public sealed class BucketObjectState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Defaults to `private`.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        /// <summary>
        /// Name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
        /// </summary>
        [Input("bucketKeyEnabled")]
        public Input<bool>? BucketKeyEnabled { get; set; }

        /// <summary>
        /// Caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
        /// </summary>
        [Input("cacheControl")]
        public Input<string>? CacheControl { get; set; }

        /// <summary>
        /// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
        /// </summary>
        [Input("contentBase64")]
        public Input<string>? ContentBase64 { get; set; }

        /// <summary>
        /// Presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
        /// </summary>
        [Input("contentDisposition")]
        public Input<string>? ContentDisposition { get; set; }

        /// <summary>
        /// Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// Language the content is in e.g. en-US or en-GB.
        /// </summary>
        [Input("contentLanguage")]
        public Input<string>? ContentLanguage { get; set; }

        /// <summary>
        /// Standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// Triggers updates when the value changes. The only meaningful value is `filemd5("path/to/file")`. This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"` (see `source_hash` instead).
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Whether to allow the object to be deleted by removing any legal hold on any object version. Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Name of the object once it is in the bucket.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// ARN of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the `aws.kms.Key` resource, use the `arn` attribute. If referencing the `aws.kms.Alias` data source or resource, use the `target_key_arn` attribute. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// [Legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
        /// </summary>
        [Input("objectLockLegalHoldStatus")]
        public Input<string>? ObjectLockLegalHoldStatus { get; set; }

        /// <summary>
        /// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
        /// </summary>
        [Input("objectLockMode")]
        public Input<string>? ObjectLockMode { get; set; }

        /// <summary>
        /// Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
        /// </summary>
        [Input("objectLockRetainUntilDate")]
        public Input<string>? ObjectLockRetainUntilDate { get; set; }

        /// <summary>
        /// Server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
        /// </summary>
        [Input("serverSideEncryption")]
        public Input<string>? ServerSideEncryption { get; set; }

        /// <summary>
        /// Path to a file that will be read and uploaded as raw bytes for the object content.
        /// </summary>
        [Input("source")]
        public Input<AssetOrArchive>? Source { get; set; }

        /// <summary>
        /// Triggers updates like `etag` but useful to address `etag` encryption limitations. Set using `filemd5("path/to/source")`. (The value is only stored in state and not saved by AWS.)
        /// </summary>
        [Input("sourceHash")]
        public Input<string>? SourceHash { get; set; }

        /// <summary>
        /// [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Unique version ID value for the object, if bucket versioning is enabled.
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        /// <summary>
        /// Target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
        /// </summary>
        [Input("websiteRedirect")]
        public Input<string>? WebsiteRedirect { get; set; }

        public BucketObjectState()
        {
        }
    }
}
