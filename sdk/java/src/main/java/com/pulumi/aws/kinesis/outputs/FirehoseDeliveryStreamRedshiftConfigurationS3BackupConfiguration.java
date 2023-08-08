// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesis.outputs;

import com.pulumi.aws.kinesis.outputs.FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration {
    /**
     * @return The ARN of the S3 bucket
     * 
     */
    private String bucketArn;
    /**
     * @return Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
     * 
     */
    private @Nullable Integer bufferingInterval;
    /**
     * @return Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
     * We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
     * 
     */
    private @Nullable Integer bufferingSize;
    /**
     * @return The CloudWatch Logging Options for the delivery stream. More details are given below
     * 
     */
    private @Nullable FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions cloudwatchLoggingOptions;
    /**
     * @return The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
     * 
     */
    private @Nullable String compressionFormat;
    /**
     * @return Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
     * 
     */
    private @Nullable String errorOutputPrefix;
    /**
     * @return Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
     * be used.
     * 
     */
    private @Nullable String kmsKeyArn;
    /**
     * @return The &#34;YYYY/MM/DD/HH&#34; time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
     * 
     */
    private @Nullable String prefix;
    /**
     * @return The ARN of the role that provides access to the source Kinesis stream.
     * 
     */
    private String roleArn;

    private FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration() {}
    /**
     * @return The ARN of the S3 bucket
     * 
     */
    public String bucketArn() {
        return this.bucketArn;
    }
    /**
     * @return Buffer incoming data for the specified period of time, in seconds between 60 to 900, before delivering it to the destination.  The default value is 300s.
     * 
     */
    public Optional<Integer> bufferingInterval() {
        return Optional.ofNullable(this.bufferingInterval);
    }
    /**
     * @return Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
     * We recommend setting SizeInMBs to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec set SizeInMBs to be 10 MB or higher.
     * 
     */
    public Optional<Integer> bufferingSize() {
        return Optional.ofNullable(this.bufferingSize);
    }
    /**
     * @return The CloudWatch Logging Options for the delivery stream. More details are given below
     * 
     */
    public Optional<FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions> cloudwatchLoggingOptions() {
        return Optional.ofNullable(this.cloudwatchLoggingOptions);
    }
    /**
     * @return The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, &amp; `HADOOP_SNAPPY`.
     * 
     */
    public Optional<String> compressionFormat() {
        return Optional.ofNullable(this.compressionFormat);
    }
    /**
     * @return Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
     * 
     */
    public Optional<String> errorOutputPrefix() {
        return Optional.ofNullable(this.errorOutputPrefix);
    }
    /**
     * @return Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
     * be used.
     * 
     */
    public Optional<String> kmsKeyArn() {
        return Optional.ofNullable(this.kmsKeyArn);
    }
    /**
     * @return The &#34;YYYY/MM/DD/HH&#34; time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
     * 
     */
    public Optional<String> prefix() {
        return Optional.ofNullable(this.prefix);
    }
    /**
     * @return The ARN of the role that provides access to the source Kinesis stream.
     * 
     */
    public String roleArn() {
        return this.roleArn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bucketArn;
        private @Nullable Integer bufferingInterval;
        private @Nullable Integer bufferingSize;
        private @Nullable FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions cloudwatchLoggingOptions;
        private @Nullable String compressionFormat;
        private @Nullable String errorOutputPrefix;
        private @Nullable String kmsKeyArn;
        private @Nullable String prefix;
        private String roleArn;
        public Builder() {}
        public Builder(FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bucketArn = defaults.bucketArn;
    	      this.bufferingInterval = defaults.bufferingInterval;
    	      this.bufferingSize = defaults.bufferingSize;
    	      this.cloudwatchLoggingOptions = defaults.cloudwatchLoggingOptions;
    	      this.compressionFormat = defaults.compressionFormat;
    	      this.errorOutputPrefix = defaults.errorOutputPrefix;
    	      this.kmsKeyArn = defaults.kmsKeyArn;
    	      this.prefix = defaults.prefix;
    	      this.roleArn = defaults.roleArn;
        }

        @CustomType.Setter
        public Builder bucketArn(String bucketArn) {
            this.bucketArn = Objects.requireNonNull(bucketArn);
            return this;
        }
        @CustomType.Setter
        public Builder bufferingInterval(@Nullable Integer bufferingInterval) {
            this.bufferingInterval = bufferingInterval;
            return this;
        }
        @CustomType.Setter
        public Builder bufferingSize(@Nullable Integer bufferingSize) {
            this.bufferingSize = bufferingSize;
            return this;
        }
        @CustomType.Setter
        public Builder cloudwatchLoggingOptions(@Nullable FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions cloudwatchLoggingOptions) {
            this.cloudwatchLoggingOptions = cloudwatchLoggingOptions;
            return this;
        }
        @CustomType.Setter
        public Builder compressionFormat(@Nullable String compressionFormat) {
            this.compressionFormat = compressionFormat;
            return this;
        }
        @CustomType.Setter
        public Builder errorOutputPrefix(@Nullable String errorOutputPrefix) {
            this.errorOutputPrefix = errorOutputPrefix;
            return this;
        }
        @CustomType.Setter
        public Builder kmsKeyArn(@Nullable String kmsKeyArn) {
            this.kmsKeyArn = kmsKeyArn;
            return this;
        }
        @CustomType.Setter
        public Builder prefix(@Nullable String prefix) {
            this.prefix = prefix;
            return this;
        }
        @CustomType.Setter
        public Builder roleArn(String roleArn) {
            this.roleArn = Objects.requireNonNull(roleArn);
            return this;
        }
        public FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration build() {
            final var o = new FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration();
            o.bucketArn = bucketArn;
            o.bufferingInterval = bufferingInterval;
            o.bufferingSize = bufferingSize;
            o.cloudwatchLoggingOptions = cloudwatchLoggingOptions;
            o.compressionFormat = compressionFormat;
            o.errorOutputPrefix = errorOutputPrefix;
            o.kmsKeyArn = kmsKeyArn;
            o.prefix = prefix;
            o.roleArn = roleArn;
            return o;
        }
    }
}
