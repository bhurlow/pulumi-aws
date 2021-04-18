// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class DistributionDefaultCacheBehaviorGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedMethods", required: true)]
        private InputList<string>? _allowedMethods;

        /// <summary>
        /// Controls which HTTP methods CloudFront
        /// processes and forwards to your Amazon S3 bucket or your custom origin.
        /// </summary>
        public InputList<string> AllowedMethods
        {
            get => _allowedMethods ?? (_allowedMethods = new InputList<string>());
            set => _allowedMethods = value;
        }

        /// <summary>
        /// The unique identifier of the cache policy that
        /// is attached to the cache behavior.
        /// </summary>
        [Input("cachePolicyId")]
        public Input<string>? CachePolicyId { get; set; }

        [Input("cachedMethods", required: true)]
        private InputList<string>? _cachedMethods;

        /// <summary>
        /// Controls whether CloudFront caches the
        /// response to requests using the specified HTTP methods.
        /// </summary>
        public InputList<string> CachedMethods
        {
            get => _cachedMethods ?? (_cachedMethods = new InputList<string>());
            set => _cachedMethods = value;
        }

        /// <summary>
        /// Whether you want CloudFront to automatically
        /// compress content for web requests that include `Accept-Encoding: gzip` in
        /// the request header (default: `false`).
        /// </summary>
        [Input("compress")]
        public Input<bool>? Compress { get; set; }

        /// <summary>
        /// The default amount of time (in seconds) that an
        /// object is in a CloudFront cache before CloudFront forwards another request
        /// in the absence of an `Cache-Control max-age` or `Expires` header.
        /// </summary>
        [Input("defaultTtl")]
        public Input<int>? DefaultTtl { get; set; }

        /// <summary>
        /// Field level encryption configuration ID
        /// </summary>
        [Input("fieldLevelEncryptionId")]
        public Input<string>? FieldLevelEncryptionId { get; set; }

        /// <summary>
        /// The forwarded values configuration that specifies how CloudFront
        /// handles query strings, cookies and headers (maximum one).
        /// </summary>
        [Input("forwardedValues")]
        public Input<Inputs.DistributionDefaultCacheBehaviorForwardedValuesGetArgs>? ForwardedValues { get; set; }

        [Input("lambdaFunctionAssociations")]
        private InputList<Inputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociationGetArgs>? _lambdaFunctionAssociations;

        /// <summary>
        /// A config block that triggers a lambda
        /// function with specific actions (maximum 4).
        /// </summary>
        public InputList<Inputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociationGetArgs> LambdaFunctionAssociations
        {
            get => _lambdaFunctionAssociations ?? (_lambdaFunctionAssociations = new InputList<Inputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociationGetArgs>());
            set => _lambdaFunctionAssociations = value;
        }

        /// <summary>
        /// The maximum amount of time (in seconds) that an
        /// object is in a CloudFront cache before CloudFront forwards another request
        /// to your origin to determine whether the object has been updated. Only
        /// effective in the presence of `Cache-Control max-age`, `Cache-Control
        /// s-maxage`, and `Expires` headers.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The minimum amount of time that you want objects to
        /// stay in CloudFront caches before CloudFront queries your origin to see
        /// whether the object has been updated. Defaults to 0 seconds.
        /// </summary>
        [Input("minTtl")]
        public Input<int>? MinTtl { get; set; }

        /// <summary>
        /// The unique identifier of the origin request policy
        /// that is attached to the behavior.
        /// </summary>
        [Input("originRequestPolicyId")]
        public Input<string>? OriginRequestPolicyId { get; set; }

        /// <summary>
        /// The ARN of the real-time log configuration
        /// that is attached to this cache behavior.
        /// </summary>
        [Input("realtimeLogConfigArn")]
        public Input<string>? RealtimeLogConfigArn { get; set; }

        /// <summary>
        /// Indicates whether you want to distribute
        /// media files in Microsoft Smooth Streaming format using the origin that is
        /// associated with this cache behavior.
        /// </summary>
        [Input("smoothStreaming")]
        public Input<bool>? SmoothStreaming { get; set; }

        /// <summary>
        /// The value of ID for the origin that you want
        /// CloudFront to route requests to when a request matches the path pattern
        /// either for a cache behavior or for the default cache behavior.
        /// </summary>
        [Input("targetOriginId", required: true)]
        public Input<string> TargetOriginId { get; set; } = null!;

        [Input("trustedKeyGroups")]
        private InputList<string>? _trustedKeyGroups;

        /// <summary>
        /// A list of key group IDs that CloudFront can use to validate signed URLs or signed cookies.
        /// See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
        /// </summary>
        public InputList<string> TrustedKeyGroups
        {
            get => _trustedKeyGroups ?? (_trustedKeyGroups = new InputList<string>());
            set => _trustedKeyGroups = value;
        }

        [Input("trustedSigners")]
        private InputList<string>? _trustedSigners;

        /// <summary>
        /// List of AWS account IDs (or `self`) that you want to allow to create signed URLs for private content.
        /// See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
        /// </summary>
        public InputList<string> TrustedSigners
        {
            get => _trustedSigners ?? (_trustedSigners = new InputList<string>());
            set => _trustedSigners = value;
        }

        /// <summary>
        /// Use this element to specify the
        /// protocol that users can use to access the files in the origin specified by
        /// TargetOriginId when a request matches the path pattern in PathPattern. One
        /// of `allow-all`, `https-only`, or `redirect-to-https`.
        /// </summary>
        [Input("viewerProtocolPolicy", required: true)]
        public Input<string> ViewerProtocolPolicy { get; set; } = null!;

        public DistributionDefaultCacheBehaviorGetArgs()
        {
        }
    }
}
