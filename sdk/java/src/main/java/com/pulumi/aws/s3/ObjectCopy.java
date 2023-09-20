// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3.ObjectCopyArgs;
import com.pulumi.aws.s3.inputs.ObjectCopyState;
import com.pulumi.aws.s3.outputs.ObjectCopyGrant;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for copying an S3 object.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.ObjectCopy;
 * import com.pulumi.aws.s3.ObjectCopyArgs;
 * import com.pulumi.aws.s3.inputs.ObjectCopyGrantArgs;
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
 *         var test = new ObjectCopy(&#34;test&#34;, ObjectCopyArgs.builder()        
 *             .bucket(&#34;destination_bucket&#34;)
 *             .grants(ObjectCopyGrantArgs.builder()
 *                 .permissions(&#34;READ&#34;)
 *                 .type(&#34;Group&#34;)
 *                 .uri(&#34;http://acs.amazonaws.com/groups/global/AllUsers&#34;)
 *                 .build())
 *             .key(&#34;destination_key&#34;)
 *             .source(&#34;source_bucket/source_key&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="aws:s3/objectCopy:ObjectCopy")
public class ObjectCopy extends com.pulumi.resources.CustomResource {
    /**
     * [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `authenticated-read`, `aws-exec-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Conflicts with `grant`.
     * 
     */
    @Export(name="acl", refs={String.class}, tree="[0]")
    private Output<String> acl;

    /**
     * @return [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `authenticated-read`, `aws-exec-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Conflicts with `grant`.
     * 
     */
    public Output<String> acl() {
        return this.acl;
    }
    /**
     * Name of the bucket to put the file in.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return Name of the bucket to put the file in.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    @Export(name="bucketKeyEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> bucketKeyEnabled;

    public Output<Boolean> bucketKeyEnabled() {
        return this.bucketKeyEnabled;
    }
    /**
     * Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
     * 
     */
    @Export(name="cacheControl", refs={String.class}, tree="[0]")
    private Output<String> cacheControl;

    /**
     * @return Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
     * 
     */
    public Output<String> cacheControl() {
        return this.cacheControl;
    }
    /**
     * Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
     * 
     */
    @Export(name="contentDisposition", refs={String.class}, tree="[0]")
    private Output<String> contentDisposition;

    /**
     * @return Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
     * 
     */
    public Output<String> contentDisposition() {
        return this.contentDisposition;
    }
    /**
     * Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
     * 
     */
    @Export(name="contentEncoding", refs={String.class}, tree="[0]")
    private Output<String> contentEncoding;

    /**
     * @return Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
     * 
     */
    public Output<String> contentEncoding() {
        return this.contentEncoding;
    }
    /**
     * Language the content is in e.g., en-US or en-GB.
     * 
     */
    @Export(name="contentLanguage", refs={String.class}, tree="[0]")
    private Output<String> contentLanguage;

    /**
     * @return Language the content is in e.g., en-US or en-GB.
     * 
     */
    public Output<String> contentLanguage() {
        return this.contentLanguage;
    }
    /**
     * Standard MIME type describing the format of the object data, e.g., `application/octet-stream`. All Valid MIME Types are valid for this input.
     * 
     */
    @Export(name="contentType", refs={String.class}, tree="[0]")
    private Output<String> contentType;

    /**
     * @return Standard MIME type describing the format of the object data, e.g., `application/octet-stream`. All Valid MIME Types are valid for this input.
     * 
     */
    public Output<String> contentType() {
        return this.contentType;
    }
    /**
     * Copies the object if its entity tag (ETag) matches the specified tag.
     * 
     */
    @Export(name="copyIfMatch", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> copyIfMatch;

    /**
     * @return Copies the object if its entity tag (ETag) matches the specified tag.
     * 
     */
    public Output<Optional<String>> copyIfMatch() {
        return Codegen.optional(this.copyIfMatch);
    }
    /**
     * Copies the object if it has been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    @Export(name="copyIfModifiedSince", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> copyIfModifiedSince;

    /**
     * @return Copies the object if it has been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    public Output<Optional<String>> copyIfModifiedSince() {
        return Codegen.optional(this.copyIfModifiedSince);
    }
    /**
     * Copies the object if its entity tag (ETag) is different than the specified ETag.
     * 
     */
    @Export(name="copyIfNoneMatch", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> copyIfNoneMatch;

    /**
     * @return Copies the object if its entity tag (ETag) is different than the specified ETag.
     * 
     */
    public Output<Optional<String>> copyIfNoneMatch() {
        return Codegen.optional(this.copyIfNoneMatch);
    }
    /**
     * Copies the object if it hasn&#39;t been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    @Export(name="copyIfUnmodifiedSince", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> copyIfUnmodifiedSince;

    /**
     * @return Copies the object if it hasn&#39;t been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    public Output<Optional<String>> copyIfUnmodifiedSince() {
        return Codegen.optional(this.copyIfUnmodifiedSince);
    }
    /**
     * Specifies the algorithm to use to when encrypting the object (for example, AES256).
     * 
     */
    @Export(name="customerAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> customerAlgorithm;

    /**
     * @return Specifies the algorithm to use to when encrypting the object (for example, AES256).
     * 
     */
    public Output<String> customerAlgorithm() {
        return this.customerAlgorithm;
    }
    /**
     * Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data. This value is used to store the object and then it is discarded; Amazon S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the x-amz-server-side-encryption-customer-algorithm header.
     * 
     */
    @Export(name="customerKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customerKey;

    /**
     * @return Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data. This value is used to store the object and then it is discarded; Amazon S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the x-amz-server-side-encryption-customer-algorithm header.
     * 
     */
    public Output<Optional<String>> customerKey() {
        return Codegen.optional(this.customerKey);
    }
    /**
     * Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
     * 
     */
    @Export(name="customerKeyMd5", refs={String.class}, tree="[0]")
    private Output<String> customerKeyMd5;

    /**
     * @return Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
     * 
     */
    public Output<String> customerKeyMd5() {
        return this.customerKeyMd5;
    }
    /**
     * ETag generated for the object (an MD5 sum of the object content). For plaintext objects or objects encrypted with an AWS-managed key, the hash is an MD5 digest of the object data. For objects encrypted with a KMS key or objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5 digest, regardless of the method of encryption. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return ETag generated for the object (an MD5 sum of the object content). For plaintext objects or objects encrypted with an AWS-managed key, the hash is an MD5 digest of the object data. For objects encrypted with a KMS key or objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5 digest, regardless of the method of encryption. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Account id of the expected destination bucket owner. If the destination bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
     * 
     */
    @Export(name="expectedBucketOwner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expectedBucketOwner;

    /**
     * @return Account id of the expected destination bucket owner. If the destination bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
     * 
     */
    public Output<Optional<String>> expectedBucketOwner() {
        return Codegen.optional(this.expectedBucketOwner);
    }
    /**
     * Account id of the expected source bucket owner. If the source bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
     * 
     */
    @Export(name="expectedSourceBucketOwner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expectedSourceBucketOwner;

    /**
     * @return Account id of the expected source bucket owner. If the source bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
     * 
     */
    public Output<Optional<String>> expectedSourceBucketOwner() {
        return Codegen.optional(this.expectedSourceBucketOwner);
    }
    /**
     * If the object expiration is configured, this attribute will be set.
     * 
     */
    @Export(name="expiration", refs={String.class}, tree="[0]")
    private Output<String> expiration;

    /**
     * @return If the object expiration is configured, this attribute will be set.
     * 
     */
    public Output<String> expiration() {
        return this.expiration;
    }
    /**
     * Date and time at which the object is no longer cacheable, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    @Export(name="expires", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expires;

    /**
     * @return Date and time at which the object is no longer cacheable, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    public Output<Optional<String>> expires() {
        return Codegen.optional(this.expires);
    }
    /**
     * Allow the object to be deleted by removing any legal hold on any object version. Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return Allow the object to be deleted by removing any legal hold on any object version. Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * Configuration block for header grants. Documented below. Conflicts with `acl`.
     * 
     */
    @Export(name="grants", refs={List.class,ObjectCopyGrant.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ObjectCopyGrant>> grants;

    /**
     * @return Configuration block for header grants. Documented below. Conflicts with `acl`.
     * 
     */
    public Output<Optional<List<ObjectCopyGrant>>> grants() {
        return Codegen.optional(this.grants);
    }
    /**
     * Name of the object once it is in the bucket.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return Name of the object once it is in the bucket.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Specifies the AWS KMS Encryption Context to use for object encryption. The value is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
     * 
     */
    @Export(name="kmsEncryptionContext", refs={String.class}, tree="[0]")
    private Output<String> kmsEncryptionContext;

    /**
     * @return Specifies the AWS KMS Encryption Context to use for object encryption. The value is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
     * 
     */
    public Output<String> kmsEncryptionContext() {
        return this.kmsEncryptionContext;
    }
    /**
     * Specifies the AWS KMS Key ARN to use for object encryption. This value is a fully qualified **ARN** of the KMS Key. If using `aws.kms.Key`, use the exported `arn` attribute: `kms_key_id = aws_kms_key.foo.arn`
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyId;

    /**
     * @return Specifies the AWS KMS Key ARN to use for object encryption. This value is a fully qualified **ARN** of the KMS Key. If using `aws.kms.Key`, use the exported `arn` attribute: `kms_key_id = aws_kms_key.foo.arn`
     * 
     */
    public Output<String> kmsKeyId() {
        return this.kmsKeyId;
    }
    /**
     * Returns the date that the object was last modified, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    @Export(name="lastModified", refs={String.class}, tree="[0]")
    private Output<String> lastModified;

    /**
     * @return Returns the date that the object was last modified, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     * 
     */
    public Output<String> lastModified() {
        return this.lastModified;
    }
    /**
     * Map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> metadata;

    /**
     * @return Map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
     * 
     */
    public Output<Map<String,String>> metadata() {
        return this.metadata;
    }
    /**
     * Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Valid values are `COPY` and `REPLACE`.
     * 
     */
    @Export(name="metadataDirective", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> metadataDirective;

    /**
     * @return Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Valid values are `COPY` and `REPLACE`.
     * 
     */
    public Output<Optional<String>> metadataDirective() {
        return Codegen.optional(this.metadataDirective);
    }
    /**
     * The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
     * 
     */
    @Export(name="objectLockLegalHoldStatus", refs={String.class}, tree="[0]")
    private Output<String> objectLockLegalHoldStatus;

    /**
     * @return The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
     * 
     */
    public Output<String> objectLockLegalHoldStatus() {
        return this.objectLockLegalHoldStatus;
    }
    /**
     * Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
     * 
     */
    @Export(name="objectLockMode", refs={String.class}, tree="[0]")
    private Output<String> objectLockMode;

    /**
     * @return Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
     * 
     */
    public Output<String> objectLockMode() {
        return this.objectLockMode;
    }
    /**
     * Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object&#39;s object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
     * 
     */
    @Export(name="objectLockRetainUntilDate", refs={String.class}, tree="[0]")
    private Output<String> objectLockRetainUntilDate;

    /**
     * @return Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object&#39;s object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
     * 
     */
    public Output<String> objectLockRetainUntilDate() {
        return this.objectLockRetainUntilDate;
    }
    /**
     * If present, indicates that the requester was successfully charged for the request.
     * 
     */
    @Export(name="requestCharged", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> requestCharged;

    /**
     * @return If present, indicates that the requester was successfully charged for the request.
     * 
     */
    public Output<Boolean> requestCharged() {
        return this.requestCharged;
    }
    /**
     * Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. For information about downloading objects from requester pays buckets, see Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html) in the Amazon S3 Developer Guide. If included, the only valid value is `requester`.
     * 
     */
    @Export(name="requestPayer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestPayer;

    /**
     * @return Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. For information about downloading objects from requester pays buckets, see Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html) in the Amazon S3 Developer Guide. If included, the only valid value is `requester`.
     * 
     */
    public Output<Optional<String>> requestPayer() {
        return Codegen.optional(this.requestPayer);
    }
    /**
     * Specifies server-side encryption of the object in S3. Valid values are `AES256` and `aws:kms`.
     * 
     */
    @Export(name="serverSideEncryption", refs={String.class}, tree="[0]")
    private Output<String> serverSideEncryption;

    /**
     * @return Specifies server-side encryption of the object in S3. Valid values are `AES256` and `aws:kms`.
     * 
     */
    public Output<String> serverSideEncryption() {
        return this.serverSideEncryption;
    }
    /**
     * Specifies the source object for the copy operation. You specify the value in one of two formats. For objects not accessed through an access point, specify the name of the source bucket and the key of the source object, separated by a slash (`/`). For example, `testbucket/test1.json`. For objects accessed through access points, specify the ARN of the object as accessed through the access point, in the format `arn:aws:s3:&lt;Region&gt;:&lt;account-id&gt;:accesspoint/&lt;access-point-name&gt;/object/&lt;key&gt;`. For example, `arn:aws:s3:us-west-2:9999912999:accesspoint/my-access-point/object/testbucket/test1.json`.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return Specifies the source object for the copy operation. You specify the value in one of two formats. For objects not accessed through an access point, specify the name of the source bucket and the key of the source object, separated by a slash (`/`). For example, `testbucket/test1.json`. For objects accessed through access points, specify the ARN of the object as accessed through the access point, in the format `arn:aws:s3:&lt;Region&gt;:&lt;account-id&gt;:accesspoint/&lt;access-point-name&gt;/object/&lt;key&gt;`. For example, `arn:aws:s3:us-west-2:9999912999:accesspoint/my-access-point/object/testbucket/test1.json`.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * Specifies the algorithm to use when decrypting the source object (for example, AES256).
     * 
     */
    @Export(name="sourceCustomerAlgorithm", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceCustomerAlgorithm;

    /**
     * @return Specifies the algorithm to use when decrypting the source object (for example, AES256).
     * 
     */
    public Output<Optional<String>> sourceCustomerAlgorithm() {
        return Codegen.optional(this.sourceCustomerAlgorithm);
    }
    /**
     * Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object. The encryption key provided in this header must be one that was used when the source object was created.
     * 
     */
    @Export(name="sourceCustomerKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceCustomerKey;

    /**
     * @return Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object. The encryption key provided in this header must be one that was used when the source object was created.
     * 
     */
    public Output<Optional<String>> sourceCustomerKey() {
        return Codegen.optional(this.sourceCustomerKey);
    }
    /**
     * Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
     * 
     */
    @Export(name="sourceCustomerKeyMd5", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceCustomerKeyMd5;

    /**
     * @return Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
     * 
     */
    public Output<Optional<String>> sourceCustomerKeyMd5() {
        return Codegen.optional(this.sourceCustomerKeyMd5);
    }
    /**
     * Version of the copied object in the source bucket.
     * 
     */
    @Export(name="sourceVersionId", refs={String.class}, tree="[0]")
    private Output<String> sourceVersionId;

    /**
     * @return Version of the copied object in the source bucket.
     * 
     */
    public Output<String> sourceVersionId() {
        return this.sourceVersionId;
    }
    /**
     * Specifies the desired [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html#AmazonS3-CopyObject-request-header-StorageClass) for the object. Defaults to `STANDARD`.
     * 
     */
    @Export(name="storageClass", refs={String.class}, tree="[0]")
    private Output<String> storageClass;

    /**
     * @return Specifies the desired [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html#AmazonS3-CopyObject-request-header-StorageClass) for the object. Defaults to `STANDARD`.
     * 
     */
    public Output<String> storageClass() {
        return this.storageClass;
    }
    /**
     * Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Valid values are `COPY` and `REPLACE`.
     * 
     */
    @Export(name="taggingDirective", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> taggingDirective;

    /**
     * @return Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Valid values are `COPY` and `REPLACE`.
     * 
     */
    public Output<Optional<String>> taggingDirective() {
        return Codegen.optional(this.taggingDirective);
    }
    /**
     * Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * Version ID of the newly created copy.
     * 
     */
    @Export(name="versionId", refs={String.class}, tree="[0]")
    private Output<String> versionId;

    /**
     * @return Version ID of the newly created copy.
     * 
     */
    public Output<String> versionId() {
        return this.versionId;
    }
    /**
     * Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
     * 
     */
    @Export(name="websiteRedirect", refs={String.class}, tree="[0]")
    private Output<String> websiteRedirect;

    /**
     * @return Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
     * 
     */
    public Output<String> websiteRedirect() {
        return this.websiteRedirect;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ObjectCopy(String name) {
        this(name, ObjectCopyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ObjectCopy(String name, ObjectCopyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ObjectCopy(String name, ObjectCopyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/objectCopy:ObjectCopy", name, args == null ? ObjectCopyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ObjectCopy(String name, Output<String> id, @Nullable ObjectCopyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/objectCopy:ObjectCopy", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "customerKey",
                "kmsEncryptionContext",
                "kmsKeyId",
                "sourceCustomerKey",
                "tagsAll"
            ))
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
    public static ObjectCopy get(String name, Output<String> id, @Nullable ObjectCopyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ObjectCopy(name, id, state, options);
    }
}
