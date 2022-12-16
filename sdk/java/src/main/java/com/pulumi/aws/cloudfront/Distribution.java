// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudfront;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cloudfront.DistributionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionState;
import com.pulumi.aws.cloudfront.outputs.DistributionCustomErrorResponse;
import com.pulumi.aws.cloudfront.outputs.DistributionDefaultCacheBehavior;
import com.pulumi.aws.cloudfront.outputs.DistributionLoggingConfig;
import com.pulumi.aws.cloudfront.outputs.DistributionOrderedCacheBehavior;
import com.pulumi.aws.cloudfront.outputs.DistributionOrigin;
import com.pulumi.aws.cloudfront.outputs.DistributionOriginGroup;
import com.pulumi.aws.cloudfront.outputs.DistributionRestrictions;
import com.pulumi.aws.cloudfront.outputs.DistributionTrustedKeyGroup;
import com.pulumi.aws.cloudfront.outputs.DistributionTrustedSigner;
import com.pulumi.aws.cloudfront.outputs.DistributionViewerCertificate;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates an Amazon CloudFront web distribution.
 * 
 * For information about CloudFront distributions, see the
 * [Amazon CloudFront Developer Guide](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html). For specific information about creating
 * CloudFront web distributions, see the [POST Distribution](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html) page in the Amazon
 * CloudFront API Reference.
 * 
 * &gt; **NOTE:** CloudFront distributions take about 15 minutes to reach a deployed
 * state after creation or modification. During this time, deletes to resources will
 * be blocked. If you need to delete a distribution that is enabled and you do not
 * want to wait, you need to use the `retain_on_delete` flag.
 * 
 * ## Example Usage
 * 
 * The following example below creates a CloudFront distribution with an S3 origin.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.BucketV2Args;
 * import com.pulumi.aws.s3.BucketAclV2;
 * import com.pulumi.aws.s3.BucketAclV2Args;
 * import com.pulumi.aws.cloudfront.Distribution;
 * import com.pulumi.aws.cloudfront.DistributionArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOriginArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionLoggingConfigArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorForwardedValuesArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOrderedCacheBehaviorArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOrderedCacheBehaviorForwardedValuesArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionRestrictionsArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionRestrictionsGeoRestrictionArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionViewerCertificateArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var bucketV2 = new BucketV2(&#34;bucketV2&#34;, BucketV2Args.builder()        
 *             .tags(Map.of(&#34;Name&#34;, &#34;My bucket&#34;))
 *             .build());
 * 
 *         var bAcl = new BucketAclV2(&#34;bAcl&#34;, BucketAclV2Args.builder()        
 *             .bucket(bucketV2.id())
 *             .acl(&#34;private&#34;)
 *             .build());
 * 
 *         final var s3OriginId = &#34;myS3Origin&#34;;
 * 
 *         var s3Distribution = new Distribution(&#34;s3Distribution&#34;, DistributionArgs.builder()        
 *             .origins(DistributionOriginArgs.builder()
 *                 .domainName(bucketV2.bucketRegionalDomainName())
 *                 .originAccessControlId(aws_cloudfront_origin_access_control.default().id())
 *                 .originId(s3OriginId)
 *                 .build())
 *             .enabled(true)
 *             .isIpv6Enabled(true)
 *             .comment(&#34;Some comment&#34;)
 *             .defaultRootObject(&#34;index.html&#34;)
 *             .loggingConfig(DistributionLoggingConfigArgs.builder()
 *                 .includeCookies(false)
 *                 .bucket(&#34;mylogs.s3.amazonaws.com&#34;)
 *                 .prefix(&#34;myprefix&#34;)
 *                 .build())
 *             .aliases(            
 *                 &#34;mysite.example.com&#34;,
 *                 &#34;yoursite.example.com&#34;)
 *             .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
 *                 .allowedMethods(                
 *                     &#34;DELETE&#34;,
 *                     &#34;GET&#34;,
 *                     &#34;HEAD&#34;,
 *                     &#34;OPTIONS&#34;,
 *                     &#34;PATCH&#34;,
 *                     &#34;POST&#34;,
 *                     &#34;PUT&#34;)
 *                 .cachedMethods(                
 *                     &#34;GET&#34;,
 *                     &#34;HEAD&#34;)
 *                 .targetOriginId(s3OriginId)
 *                 .forwardedValues(DistributionDefaultCacheBehaviorForwardedValuesArgs.builder()
 *                     .queryString(false)
 *                     .cookies(DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs.builder()
 *                         .forward(&#34;none&#34;)
 *                         .build())
 *                     .build())
 *                 .viewerProtocolPolicy(&#34;allow-all&#34;)
 *                 .minTtl(0)
 *                 .defaultTtl(3600)
 *                 .maxTtl(86400)
 *                 .build())
 *             .orderedCacheBehaviors(            
 *                 DistributionOrderedCacheBehaviorArgs.builder()
 *                     .pathPattern(&#34;/content/immutable/*&#34;)
 *                     .allowedMethods(                    
 *                         &#34;GET&#34;,
 *                         &#34;HEAD&#34;,
 *                         &#34;OPTIONS&#34;)
 *                     .cachedMethods(                    
 *                         &#34;GET&#34;,
 *                         &#34;HEAD&#34;,
 *                         &#34;OPTIONS&#34;)
 *                     .targetOriginId(s3OriginId)
 *                     .forwardedValues(DistributionOrderedCacheBehaviorForwardedValuesArgs.builder()
 *                         .queryString(false)
 *                         .headers(&#34;Origin&#34;)
 *                         .cookies(DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs.builder()
 *                             .forward(&#34;none&#34;)
 *                             .build())
 *                         .build())
 *                     .minTtl(0)
 *                     .defaultTtl(86400)
 *                     .maxTtl(31536000)
 *                     .compress(true)
 *                     .viewerProtocolPolicy(&#34;redirect-to-https&#34;)
 *                     .build(),
 *                 DistributionOrderedCacheBehaviorArgs.builder()
 *                     .pathPattern(&#34;/content/*&#34;)
 *                     .allowedMethods(                    
 *                         &#34;GET&#34;,
 *                         &#34;HEAD&#34;,
 *                         &#34;OPTIONS&#34;)
 *                     .cachedMethods(                    
 *                         &#34;GET&#34;,
 *                         &#34;HEAD&#34;)
 *                     .targetOriginId(s3OriginId)
 *                     .forwardedValues(DistributionOrderedCacheBehaviorForwardedValuesArgs.builder()
 *                         .queryString(false)
 *                         .cookies(DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs.builder()
 *                             .forward(&#34;none&#34;)
 *                             .build())
 *                         .build())
 *                     .minTtl(0)
 *                     .defaultTtl(3600)
 *                     .maxTtl(86400)
 *                     .compress(true)
 *                     .viewerProtocolPolicy(&#34;redirect-to-https&#34;)
 *                     .build())
 *             .priceClass(&#34;PriceClass_200&#34;)
 *             .restrictions(DistributionRestrictionsArgs.builder()
 *                 .geoRestriction(DistributionRestrictionsGeoRestrictionArgs.builder()
 *                     .restrictionType(&#34;whitelist&#34;)
 *                     .locations(                    
 *                         &#34;US&#34;,
 *                         &#34;CA&#34;,
 *                         &#34;GB&#34;,
 *                         &#34;DE&#34;)
 *                     .build())
 *                 .build())
 *             .tags(Map.of(&#34;Environment&#34;, &#34;production&#34;))
 *             .viewerCertificate(DistributionViewerCertificateArgs.builder()
 *                 .cloudfrontDefaultCertificate(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * The following example below creates a Cloudfront distribution with an origin group for failover routing:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudfront.Distribution;
 * import com.pulumi.aws.cloudfront.DistributionArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOriginGroupArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOriginGroupFailoverCriteriaArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOriginArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionOriginS3OriginConfigArgs;
 * import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var s3Distribution = new Distribution(&#34;s3Distribution&#34;, DistributionArgs.builder()        
 *             .originGroups(DistributionOriginGroupArgs.builder()
 *                 .originId(&#34;groupS3&#34;)
 *                 .failoverCriteria(DistributionOriginGroupFailoverCriteriaArgs.builder()
 *                     .statusCodes(                    
 *                         403,
 *                         404,
 *                         500,
 *                         502)
 *                     .build())
 *                 .members(                
 *                     DistributionOriginGroupMemberArgs.builder()
 *                         .originId(&#34;primaryS3&#34;)
 *                         .build(),
 *                     DistributionOriginGroupMemberArgs.builder()
 *                         .originId(&#34;failoverS3&#34;)
 *                         .build())
 *                 .build())
 *             .origins(            
 *                 DistributionOriginArgs.builder()
 *                     .domainName(aws_s3_bucket.primary().bucket_regional_domain_name())
 *                     .originId(&#34;primaryS3&#34;)
 *                     .s3OriginConfig(DistributionOriginS3OriginConfigArgs.builder()
 *                         .originAccessIdentity(aws_cloudfront_origin_access_identity.default().cloudfront_access_identity_path())
 *                         .build())
 *                     .build(),
 *                 DistributionOriginArgs.builder()
 *                     .domainName(aws_s3_bucket.failover().bucket_regional_domain_name())
 *                     .originId(&#34;failoverS3&#34;)
 *                     .s3OriginConfig(DistributionOriginS3OriginConfigArgs.builder()
 *                         .originAccessIdentity(aws_cloudfront_origin_access_identity.default().cloudfront_access_identity_path())
 *                         .build())
 *                     .build())
 *             .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
 *                 .targetOriginId(&#34;groupS3&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloudfront Distributions can be imported using the `id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:cloudfront/distribution:Distribution distribution E74FTE3EXAMPLE
 * ```
 * 
 */
@ResourceType(type="aws:cloudfront/distribution:Distribution")
public class Distribution extends com.pulumi.resources.CustomResource {
    /**
     * Extra CNAMEs (alternate domain names), if any, for
     * this distribution.
     * 
     */
    @Export(name="aliases", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> aliases;

    /**
     * @return Extra CNAMEs (alternate domain names), if any, for
     * this distribution.
     * 
     */
    public Output<Optional<List<String>>> aliases() {
        return Codegen.optional(this.aliases);
    }
    /**
     * The ARN (Amazon Resource Name) for the distribution. For example: `arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5`, where `123456789012` is your AWS account ID.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN (Amazon Resource Name) for the distribution. For example: `arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5`, where `123456789012` is your AWS account ID.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Internal value used by CloudFront to allow future
     * updates to the distribution configuration.
     * 
     */
    @Export(name="callerReference", refs={String.class}, tree="[0]")
    private Output<String> callerReference;

    /**
     * @return Internal value used by CloudFront to allow future
     * updates to the distribution configuration.
     * 
     */
    public Output<String> callerReference() {
        return this.callerReference;
    }
    /**
     * Any comments you want to include about the
     * distribution.
     * 
     */
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> comment;

    /**
     * @return Any comments you want to include about the
     * distribution.
     * 
     */
    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    /**
     * One or more custom error response elements (multiples allowed).
     * 
     */
    @Export(name="customErrorResponses", refs={List.class,DistributionCustomErrorResponse.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DistributionCustomErrorResponse>> customErrorResponses;

    /**
     * @return One or more custom error response elements (multiples allowed).
     * 
     */
    public Output<Optional<List<DistributionCustomErrorResponse>>> customErrorResponses() {
        return Codegen.optional(this.customErrorResponses);
    }
    /**
     * The default cache behavior for this distribution (maximum
     * one).
     * 
     */
    @Export(name="defaultCacheBehavior", refs={DistributionDefaultCacheBehavior.class}, tree="[0]")
    private Output<DistributionDefaultCacheBehavior> defaultCacheBehavior;

    /**
     * @return The default cache behavior for this distribution (maximum
     * one).
     * 
     */
    public Output<DistributionDefaultCacheBehavior> defaultCacheBehavior() {
        return this.defaultCacheBehavior;
    }
    /**
     * The object that you want CloudFront to
     * return (for example, index.html) when an end user requests the root URL.
     * 
     */
    @Export(name="defaultRootObject", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultRootObject;

    /**
     * @return The object that you want CloudFront to
     * return (for example, index.html) when an end user requests the root URL.
     * 
     */
    public Output<Optional<String>> defaultRootObject() {
        return Codegen.optional(this.defaultRootObject);
    }
    /**
     * The DNS domain name of either the S3 bucket, or
     * web site of your custom origin.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The DNS domain name of either the S3 bucket, or
     * web site of your custom origin.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * A flag that specifies whether Origin Shield is enabled.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return A flag that specifies whether Origin Shield is enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The current version of the distribution&#39;s information. For example:
     * `E2QWRUHAPOMQZL`.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return The current version of the distribution&#39;s information. For example:
     * `E2QWRUHAPOMQZL`.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The CloudFront Route 53 zone ID that can be used to
     * route an [Alias Resource Record Set](http://docs.aws.amazon.com/Route53/latest/APIReference/CreateAliasRRSAPI.html) to. This attribute is simply an
     * alias for the zone ID `Z2FDTNDATAQYW2`.
     * 
     */
    @Export(name="hostedZoneId", refs={String.class}, tree="[0]")
    private Output<String> hostedZoneId;

    /**
     * @return The CloudFront Route 53 zone ID that can be used to
     * route an [Alias Resource Record Set](http://docs.aws.amazon.com/Route53/latest/APIReference/CreateAliasRRSAPI.html) to. This attribute is simply an
     * alias for the zone ID `Z2FDTNDATAQYW2`.
     * 
     */
    public Output<String> hostedZoneId() {
        return this.hostedZoneId;
    }
    /**
     * The maximum HTTP version to support on the
     * distribution. Allowed values are `http1.1`, `http2`, `http2and3` and `http3`. The default is
     * `http2`.
     * 
     */
    @Export(name="httpVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> httpVersion;

    /**
     * @return The maximum HTTP version to support on the
     * distribution. Allowed values are `http1.1`, `http2`, `http2and3` and `http3`. The default is
     * `http2`.
     * 
     */
    public Output<Optional<String>> httpVersion() {
        return Codegen.optional(this.httpVersion);
    }
    /**
     * The number of invalidation batches
     * currently in progress.
     * 
     */
    @Export(name="inProgressValidationBatches", refs={Integer.class}, tree="[0]")
    private Output<Integer> inProgressValidationBatches;

    /**
     * @return The number of invalidation batches
     * currently in progress.
     * 
     */
    public Output<Integer> inProgressValidationBatches() {
        return this.inProgressValidationBatches;
    }
    /**
     * Whether the IPv6 is enabled for the distribution.
     * 
     */
    @Export(name="isIpv6Enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isIpv6Enabled;

    /**
     * @return Whether the IPv6 is enabled for the distribution.
     * 
     */
    public Output<Optional<Boolean>> isIpv6Enabled() {
        return Codegen.optional(this.isIpv6Enabled);
    }
    /**
     * The date and time the distribution was last modified.
     * 
     */
    @Export(name="lastModifiedTime", refs={String.class}, tree="[0]")
    private Output<String> lastModifiedTime;

    /**
     * @return The date and time the distribution was last modified.
     * 
     */
    public Output<String> lastModifiedTime() {
        return this.lastModifiedTime;
    }
    /**
     * The logging
     * configuration that controls how logs are written
     * to your distribution (maximum one).
     * 
     */
    @Export(name="loggingConfig", refs={DistributionLoggingConfig.class}, tree="[0]")
    private Output</* @Nullable */ DistributionLoggingConfig> loggingConfig;

    /**
     * @return The logging
     * configuration that controls how logs are written
     * to your distribution (maximum one).
     * 
     */
    public Output<Optional<DistributionLoggingConfig>> loggingConfig() {
        return Codegen.optional(this.loggingConfig);
    }
    /**
     * An ordered list of cache behaviors
     * resource for this distribution. List from top to bottom
     * in order of precedence. The topmost cache behavior will have precedence 0.
     * 
     */
    @Export(name="orderedCacheBehaviors", refs={List.class,DistributionOrderedCacheBehavior.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DistributionOrderedCacheBehavior>> orderedCacheBehaviors;

    /**
     * @return An ordered list of cache behaviors
     * resource for this distribution. List from top to bottom
     * in order of precedence. The topmost cache behavior will have precedence 0.
     * 
     */
    public Output<Optional<List<DistributionOrderedCacheBehavior>>> orderedCacheBehaviors() {
        return Codegen.optional(this.orderedCacheBehaviors);
    }
    /**
     * One or more origin_group for this
     * distribution (multiples allowed).
     * 
     */
    @Export(name="originGroups", refs={List.class,DistributionOriginGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DistributionOriginGroup>> originGroups;

    /**
     * @return One or more origin_group for this
     * distribution (multiples allowed).
     * 
     */
    public Output<Optional<List<DistributionOriginGroup>>> originGroups() {
        return Codegen.optional(this.originGroups);
    }
    /**
     * One or more origins for this
     * distribution (multiples allowed).
     * 
     */
    @Export(name="origins", refs={List.class,DistributionOrigin.class}, tree="[0,1]")
    private Output<List<DistributionOrigin>> origins;

    /**
     * @return One or more origins for this
     * distribution (multiples allowed).
     * 
     */
    public Output<List<DistributionOrigin>> origins() {
        return this.origins;
    }
    /**
     * The price class for this distribution. One of
     * `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
     * 
     */
    @Export(name="priceClass", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> priceClass;

    /**
     * @return The price class for this distribution. One of
     * `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
     * 
     */
    public Output<Optional<String>> priceClass() {
        return Codegen.optional(this.priceClass);
    }
    /**
     * The restriction
     * configuration for this distribution (maximum one).
     * 
     */
    @Export(name="restrictions", refs={DistributionRestrictions.class}, tree="[0]")
    private Output<DistributionRestrictions> restrictions;

    /**
     * @return The restriction
     * configuration for this distribution (maximum one).
     * 
     */
    public Output<DistributionRestrictions> restrictions() {
        return this.restrictions;
    }
    /**
     * Disables the distribution instead of
     * deleting it when destroying the resource. If this is set,
     * the distribution needs to be deleted manually afterwards. Default: `false`.
     * 
     */
    @Export(name="retainOnDelete", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> retainOnDelete;

    /**
     * @return Disables the distribution instead of
     * deleting it when destroying the resource. If this is set,
     * the distribution needs to be deleted manually afterwards. Default: `false`.
     * 
     */
    public Output<Optional<Boolean>> retainOnDelete() {
        return Codegen.optional(this.retainOnDelete);
    }
    /**
     * The current status of the distribution. `Deployed` if the
     * distribution&#39;s information is fully propagated throughout the Amazon
     * CloudFront system.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The current status of the distribution. `Deployed` if the
     * distribution&#39;s information is fully propagated throughout the Amazon
     * CloudFront system.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * A list of key group IDs that CloudFront can use to validate signed URLs or signed cookies.
     * See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
     * 
     */
    @Export(name="trustedKeyGroups", refs={List.class,DistributionTrustedKeyGroup.class}, tree="[0,1]")
    private Output<List<DistributionTrustedKeyGroup>> trustedKeyGroups;

    /**
     * @return A list of key group IDs that CloudFront can use to validate signed URLs or signed cookies.
     * See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
     * 
     */
    public Output<List<DistributionTrustedKeyGroup>> trustedKeyGroups() {
        return this.trustedKeyGroups;
    }
    /**
     * List of AWS account IDs (or `self`) that you want to allow to create signed URLs for private content.
     * See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
     * 
     */
    @Export(name="trustedSigners", refs={List.class,DistributionTrustedSigner.class}, tree="[0,1]")
    private Output<List<DistributionTrustedSigner>> trustedSigners;

    /**
     * @return List of AWS account IDs (or `self`) that you want to allow to create signed URLs for private content.
     * See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
     * 
     */
    public Output<List<DistributionTrustedSigner>> trustedSigners() {
        return this.trustedSigners;
    }
    /**
     * The SSL
     * configuration for this distribution (maximum
     * one).
     * 
     */
    @Export(name="viewerCertificate", refs={DistributionViewerCertificate.class}, tree="[0]")
    private Output<DistributionViewerCertificate> viewerCertificate;

    /**
     * @return The SSL
     * configuration for this distribution (maximum
     * one).
     * 
     */
    public Output<DistributionViewerCertificate> viewerCertificate() {
        return this.viewerCertificate;
    }
    /**
     * If enabled, the resource will wait for
     * the distribution status to change from `InProgress` to `Deployed`. Setting
     * this to`false` will skip the process. Default: `true`.
     * 
     */
    @Export(name="waitForDeployment", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> waitForDeployment;

    /**
     * @return If enabled, the resource will wait for
     * the distribution status to change from `InProgress` to `Deployed`. Setting
     * this to`false` will skip the process. Default: `true`.
     * 
     */
    public Output<Optional<Boolean>> waitForDeployment() {
        return Codegen.optional(this.waitForDeployment);
    }
    /**
     * A unique identifier that specifies the AWS WAF web ACL,
     * if any, to associate with this distribution.
     * To specify a web ACL created using the latest version of AWS WAF (WAFv2), use the ACL ARN,
     * for example `aws_wafv2_web_acl.example.arn`. To specify a web
     * ACL created using AWS WAF Classic, use the ACL ID, for example `aws_waf_web_acl.example.id`.
     * The WAF Web ACL must exist in the WAF Global (CloudFront) region and the
     * credentials configuring this argument must have `waf:GetWebACL` permissions assigned.
     * 
     */
    @Export(name="webAclId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> webAclId;

    /**
     * @return A unique identifier that specifies the AWS WAF web ACL,
     * if any, to associate with this distribution.
     * To specify a web ACL created using the latest version of AWS WAF (WAFv2), use the ACL ARN,
     * for example `aws_wafv2_web_acl.example.arn`. To specify a web
     * ACL created using AWS WAF Classic, use the ACL ID, for example `aws_waf_web_acl.example.id`.
     * The WAF Web ACL must exist in the WAF Global (CloudFront) region and the
     * credentials configuring this argument must have `waf:GetWebACL` permissions assigned.
     * 
     */
    public Output<Optional<String>> webAclId() {
        return Codegen.optional(this.webAclId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Distribution(String name) {
        this(name, DistributionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Distribution(String name, DistributionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Distribution(String name, DistributionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudfront/distribution:Distribution", name, args == null ? DistributionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Distribution(String name, Output<String> id, @Nullable DistributionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudfront/distribution:Distribution", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Distribution get(String name, Output<String> id, @Nullable DistributionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Distribution(name, id, state, options);
    }
}
