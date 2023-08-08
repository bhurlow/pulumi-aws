// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesis.inputs;

import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationCloudwatchLoggingOptionsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs Empty = new FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs();

    /**
     * The ARN of the S3 bucket
     * 
     */
    @Import(name="bucketArn", required=true)
    private Output<String> bucketArn;

    /**
     * @return The ARN of the S3 bucket
     * 
     */
    public Output<String> bucketArn() {
        return this.bucketArn;
    }

    /**
     * Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
     * 
     */
    @Import(name="bufferingInterval")
    private @Nullable Output<Integer> bufferingInterval;

    /**
     * @return Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
     * 
     */
    public Optional<Output<Integer>> bufferingInterval() {
        return Optional.ofNullable(this.bufferingInterval);
    }

    /**
     * Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
     * We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
     * 
     */
    @Import(name="bufferingSize")
    private @Nullable Output<Integer> bufferingSize;

    /**
     * @return Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
     * We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
     * 
     */
    public Optional<Output<Integer>> bufferingSize() {
        return Optional.ofNullable(this.bufferingSize);
    }

    /**
     * The CloudWatch Logging Options for the delivery stream. More details are given below
     * 
     */
    @Import(name="cloudwatchLoggingOptions")
    private @Nullable Output<FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationCloudwatchLoggingOptionsArgs> cloudwatchLoggingOptions;

    /**
     * @return The CloudWatch Logging Options for the delivery stream. More details are given below
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationCloudwatchLoggingOptionsArgs>> cloudwatchLoggingOptions() {
        return Optional.ofNullable(this.cloudwatchLoggingOptions);
    }

    /**
     * The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
     * 
     */
    @Import(name="compressionFormat")
    private @Nullable Output<String> compressionFormat;

    /**
     * @return The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
     * 
     */
    public Optional<Output<String>> compressionFormat() {
        return Optional.ofNullable(this.compressionFormat);
    }

    /**
     * Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
     * 
     */
    @Import(name="errorOutputPrefix")
    private @Nullable Output<String> errorOutputPrefix;

    /**
     * @return Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
     * 
     */
    public Optional<Output<String>> errorOutputPrefix() {
        return Optional.ofNullable(this.errorOutputPrefix);
    }

    /**
     * Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
     * be used.
     * 
     */
    @Import(name="kmsKeyArn")
    private @Nullable Output<String> kmsKeyArn;

    /**
     * @return Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
     * be used.
     * 
     */
    public Optional<Output<String>> kmsKeyArn() {
        return Optional.ofNullable(this.kmsKeyArn);
    }

    /**
     * The &#34;YYYY/MM/DD/HH&#34; time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
     * 
     */
    @Import(name="prefix")
    private @Nullable Output<String> prefix;

    /**
     * @return The &#34;YYYY/MM/DD/HH&#34; time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
     * 
     */
    public Optional<Output<String>> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    /**
     * The ARN of the role that provides access to the source Kinesis stream.
     * 
     */
    @Import(name="roleArn", required=true)
    private Output<String> roleArn;

    /**
     * @return The ARN of the role that provides access to the source Kinesis stream.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }

    private FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs() {}

    private FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs(FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs $) {
        this.bucketArn = $.bucketArn;
        this.bufferingInterval = $.bufferingInterval;
        this.bufferingSize = $.bufferingSize;
        this.cloudwatchLoggingOptions = $.cloudwatchLoggingOptions;
        this.compressionFormat = $.compressionFormat;
        this.errorOutputPrefix = $.errorOutputPrefix;
        this.kmsKeyArn = $.kmsKeyArn;
        this.prefix = $.prefix;
        this.roleArn = $.roleArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs $;

        public Builder() {
            $ = new FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs();
        }

        public Builder(FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs defaults) {
            $ = new FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucketArn The ARN of the S3 bucket
         * 
         * @return builder
         * 
         */
        public Builder bucketArn(Output<String> bucketArn) {
            $.bucketArn = bucketArn;
            return this;
        }

        /**
         * @param bucketArn The ARN of the S3 bucket
         * 
         * @return builder
         * 
         */
        public Builder bucketArn(String bucketArn) {
            return bucketArn(Output.of(bucketArn));
        }

        /**
         * @param bufferingInterval Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
         * 
         * @return builder
         * 
         */
        public Builder bufferingInterval(@Nullable Output<Integer> bufferingInterval) {
            $.bufferingInterval = bufferingInterval;
            return this;
        }

        /**
         * @param bufferingInterval Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
         * 
         * @return builder
         * 
         */
        public Builder bufferingInterval(Integer bufferingInterval) {
            return bufferingInterval(Output.of(bufferingInterval));
        }

        /**
         * @param bufferingSize Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
         * We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
         * 
         * @return builder
         * 
         */
        public Builder bufferingSize(@Nullable Output<Integer> bufferingSize) {
            $.bufferingSize = bufferingSize;
            return this;
        }

        /**
         * @param bufferingSize Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
         * We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
         * 
         * @return builder
         * 
         */
        public Builder bufferingSize(Integer bufferingSize) {
            return bufferingSize(Output.of(bufferingSize));
        }

        /**
         * @param cloudwatchLoggingOptions The CloudWatch Logging Options for the delivery stream. More details are given below
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchLoggingOptions(@Nullable Output<FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationCloudwatchLoggingOptionsArgs> cloudwatchLoggingOptions) {
            $.cloudwatchLoggingOptions = cloudwatchLoggingOptions;
            return this;
        }

        /**
         * @param cloudwatchLoggingOptions The CloudWatch Logging Options for the delivery stream. More details are given below
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchLoggingOptions(FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationCloudwatchLoggingOptionsArgs cloudwatchLoggingOptions) {
            return cloudwatchLoggingOptions(Output.of(cloudwatchLoggingOptions));
        }

        /**
         * @param compressionFormat The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
         * 
         * @return builder
         * 
         */
        public Builder compressionFormat(@Nullable Output<String> compressionFormat) {
            $.compressionFormat = compressionFormat;
            return this;
        }

        /**
         * @param compressionFormat The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
         * 
         * @return builder
         * 
         */
        public Builder compressionFormat(String compressionFormat) {
            return compressionFormat(Output.of(compressionFormat));
        }

        /**
         * @param errorOutputPrefix Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
         * 
         * @return builder
         * 
         */
        public Builder errorOutputPrefix(@Nullable Output<String> errorOutputPrefix) {
            $.errorOutputPrefix = errorOutputPrefix;
            return this;
        }

        /**
         * @param errorOutputPrefix Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
         * 
         * @return builder
         * 
         */
        public Builder errorOutputPrefix(String errorOutputPrefix) {
            return errorOutputPrefix(Output.of(errorOutputPrefix));
        }

        /**
         * @param kmsKeyArn Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
         * be used.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(@Nullable Output<String> kmsKeyArn) {
            $.kmsKeyArn = kmsKeyArn;
            return this;
        }

        /**
         * @param kmsKeyArn Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
         * be used.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(String kmsKeyArn) {
            return kmsKeyArn(Output.of(kmsKeyArn));
        }

        /**
         * @param prefix The &#34;YYYY/MM/DD/HH&#34; time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
         * 
         * @return builder
         * 
         */
        public Builder prefix(@Nullable Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        /**
         * @param prefix The &#34;YYYY/MM/DD/HH&#34; time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
         * 
         * @return builder
         * 
         */
        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        /**
         * @param roleArn The ARN of the role that provides access to the source Kinesis stream.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The ARN of the role that provides access to the source Kinesis stream.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        public FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs build() {
            $.bucketArn = Objects.requireNonNull($.bucketArn, "expected parameter 'bucketArn' to be non-null");
            $.roleArn = Objects.requireNonNull($.roleArn, "expected parameter 'roleArn' to be non-null");
            return $;
        }
    }

}
