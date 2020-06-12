// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3
{
    public static class GetBucketObjects
    {
        /// <summary>
        /// &gt; **NOTE on `max_keys`:** Retrieving very large numbers of keys can adversely affect this provider's performance.
        /// 
        /// The bucket-objects data source returns keys (i.e., file names) and other metadata about objects in an S3 bucket.
        /// </summary>
        public static Task<GetBucketObjectsResult> InvokeAsync(GetBucketObjectsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBucketObjectsResult>("aws:s3/getBucketObjects:getBucketObjects", args ?? new GetBucketObjectsArgs(), options.WithVersion());
    }


    public sealed class GetBucketObjectsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Lists object keys in this S3 bucket. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
        /// </summary>
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        /// <summary>
        /// A character used to group keys (Default: none)
        /// </summary>
        [Input("delimiter")]
        public string? Delimiter { get; set; }

        /// <summary>
        /// Encodes keys using this method (Default: none; besides none, only "url" can be used)
        /// </summary>
        [Input("encodingType")]
        public string? EncodingType { get; set; }

        /// <summary>
        /// Boolean specifying whether to populate the owner list (Default: false)
        /// </summary>
        [Input("fetchOwner")]
        public bool? FetchOwner { get; set; }

        /// <summary>
        /// Maximum object keys to return (Default: 1000)
        /// </summary>
        [Input("maxKeys")]
        public int? MaxKeys { get; set; }

        /// <summary>
        /// Limits results to object keys with this prefix (Default: none)
        /// </summary>
        [Input("prefix")]
        public string? Prefix { get; set; }

        /// <summary>
        /// Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
        /// </summary>
        [Input("startAfter")]
        public string? StartAfter { get; set; }

        public GetBucketObjectsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBucketObjectsResult
    {
        public readonly string Bucket;
        /// <summary>
        /// List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
        /// </summary>
        public readonly ImmutableArray<string> CommonPrefixes;
        public readonly string? Delimiter;
        public readonly string? EncodingType;
        public readonly bool? FetchOwner;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of strings representing object keys
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        public readonly int? MaxKeys;
        /// <summary>
        /// List of strings representing object owner IDs (see `fetch_owner` above)
        /// </summary>
        public readonly ImmutableArray<string> Owners;
        public readonly string? Prefix;
        public readonly string? StartAfter;

        [OutputConstructor]
        private GetBucketObjectsResult(
            string bucket,

            ImmutableArray<string> commonPrefixes,

            string? delimiter,

            string? encodingType,

            bool? fetchOwner,

            string id,

            ImmutableArray<string> keys,

            int? maxKeys,

            ImmutableArray<string> owners,

            string? prefix,

            string? startAfter)
        {
            Bucket = bucket;
            CommonPrefixes = commonPrefixes;
            Delimiter = delimiter;
            EncodingType = encodingType;
            FetchOwner = fetchOwner;
            Id = id;
            Keys = keys;
            MaxKeys = maxKeys;
            Owners = owners;
            Prefix = prefix;
            StartAfter = startAfter;
        }
    }
}
