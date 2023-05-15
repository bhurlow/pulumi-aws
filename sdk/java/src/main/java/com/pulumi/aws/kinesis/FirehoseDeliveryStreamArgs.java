// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesis;

import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamHttpEndpointConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamKinesisSourceConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamRedshiftConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamServerSideEncryptionArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamSplunkConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FirehoseDeliveryStreamArgs extends com.pulumi.resources.ResourceArgs {

    public static final FirehoseDeliveryStreamArgs Empty = new FirehoseDeliveryStreamArgs();

    /**
     * The Amazon Resource Name (ARN) specifying the Stream
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) specifying the Stream
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, `splunk`, `http_endpoint` and `opensearch`.
     * 
     */
    @Import(name="destination", required=true)
    private Output<String> destination;

    /**
     * @return This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, `splunk`, `http_endpoint` and `opensearch`.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }

    @Import(name="destinationId")
    private @Nullable Output<String> destinationId;

    public Optional<Output<String>> destinationId() {
        return Optional.ofNullable(this.destinationId);
    }

    /**
     * Configuration options if elasticsearch is the destination. More details are given below.
     * 
     */
    @Import(name="elasticsearchConfiguration")
    private @Nullable Output<FirehoseDeliveryStreamElasticsearchConfigurationArgs> elasticsearchConfiguration;

    /**
     * @return Configuration options if elasticsearch is the destination. More details are given below.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamElasticsearchConfigurationArgs>> elasticsearchConfiguration() {
        return Optional.ofNullable(this.elasticsearchConfiguration);
    }

    /**
     * Enhanced configuration options for the s3 destination. More details are given below.
     * 
     */
    @Import(name="extendedS3Configuration")
    private @Nullable Output<FirehoseDeliveryStreamExtendedS3ConfigurationArgs> extendedS3Configuration;

    /**
     * @return Enhanced configuration options for the s3 destination. More details are given below.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamExtendedS3ConfigurationArgs>> extendedS3Configuration() {
        return Optional.ofNullable(this.extendedS3Configuration);
    }

    /**
     * Configuration options if http_endpoint is the destination. requires the user to also specify a `s3_configuration` block.  More details are given below.
     * 
     */
    @Import(name="httpEndpointConfiguration")
    private @Nullable Output<FirehoseDeliveryStreamHttpEndpointConfigurationArgs> httpEndpointConfiguration;

    /**
     * @return Configuration options if http_endpoint is the destination. requires the user to also specify a `s3_configuration` block.  More details are given below.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamHttpEndpointConfigurationArgs>> httpEndpointConfiguration() {
        return Optional.ofNullable(this.httpEndpointConfiguration);
    }

    /**
     * Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
     * 
     */
    @Import(name="kinesisSourceConfiguration")
    private @Nullable Output<FirehoseDeliveryStreamKinesisSourceConfigurationArgs> kinesisSourceConfiguration;

    /**
     * @return Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamKinesisSourceConfigurationArgs>> kinesisSourceConfiguration() {
        return Optional.ofNullable(this.kinesisSourceConfiguration);
    }

    /**
     * A name to identify the stream. This is unique to the AWS account and region the Stream is created in. When using for WAF logging, name must be prefixed with `aws-waf-logs-`. See [AWS Documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-policies.html#waf-policies-logging-config) for more details.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name to identify the stream. This is unique to the AWS account and region the Stream is created in. When using for WAF logging, name must be prefixed with `aws-waf-logs-`. See [AWS Documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-policies.html#waf-policies-logging-config) for more details.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Configuration options if opensearch is the destination. More details are given below.
     * 
     */
    @Import(name="opensearchConfiguration")
    private @Nullable Output<FirehoseDeliveryStreamOpensearchConfigurationArgs> opensearchConfiguration;

    /**
     * @return Configuration options if opensearch is the destination. More details are given below.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamOpensearchConfigurationArgs>> opensearchConfiguration() {
        return Optional.ofNullable(this.opensearchConfiguration);
    }

    /**
     * Configuration options if redshift is the destination.
     * Using `redshift_configuration` requires the user to also specify a
     * `s3_configuration` block. More details are given below.
     * 
     */
    @Import(name="redshiftConfiguration")
    private @Nullable Output<FirehoseDeliveryStreamRedshiftConfigurationArgs> redshiftConfiguration;

    /**
     * @return Configuration options if redshift is the destination.
     * Using `redshift_configuration` requires the user to also specify a
     * `s3_configuration` block. More details are given below.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamRedshiftConfigurationArgs>> redshiftConfiguration() {
        return Optional.ofNullable(this.redshiftConfiguration);
    }

    /**
     * Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
     * is redshift). More details are given below.
     * 
     */
    @Import(name="s3Configuration")
    private @Nullable Output<FirehoseDeliveryStreamS3ConfigurationArgs> s3Configuration;

    /**
     * @return Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
     * is redshift). More details are given below.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamS3ConfigurationArgs>> s3Configuration() {
        return Optional.ofNullable(this.s3Configuration);
    }

    /**
     * Encrypt at rest options.
     * Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
     * 
     */
    @Import(name="serverSideEncryption")
    private @Nullable Output<FirehoseDeliveryStreamServerSideEncryptionArgs> serverSideEncryption;

    /**
     * @return Encrypt at rest options.
     * Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamServerSideEncryptionArgs>> serverSideEncryption() {
        return Optional.ofNullable(this.serverSideEncryption);
    }

    /**
     * Configuration options if splunk is the destination. More details are given below.
     * 
     */
    @Import(name="splunkConfiguration")
    private @Nullable Output<FirehoseDeliveryStreamSplunkConfigurationArgs> splunkConfiguration;

    /**
     * @return Configuration options if splunk is the destination. More details are given below.
     * 
     */
    public Optional<Output<FirehoseDeliveryStreamSplunkConfigurationArgs>> splunkConfiguration() {
        return Optional.ofNullable(this.splunkConfiguration);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Specifies the table version for the output data schema. Defaults to `LATEST`.
     * 
     */
    @Import(name="versionId")
    private @Nullable Output<String> versionId;

    /**
     * @return Specifies the table version for the output data schema. Defaults to `LATEST`.
     * 
     */
    public Optional<Output<String>> versionId() {
        return Optional.ofNullable(this.versionId);
    }

    private FirehoseDeliveryStreamArgs() {}

    private FirehoseDeliveryStreamArgs(FirehoseDeliveryStreamArgs $) {
        this.arn = $.arn;
        this.destination = $.destination;
        this.destinationId = $.destinationId;
        this.elasticsearchConfiguration = $.elasticsearchConfiguration;
        this.extendedS3Configuration = $.extendedS3Configuration;
        this.httpEndpointConfiguration = $.httpEndpointConfiguration;
        this.kinesisSourceConfiguration = $.kinesisSourceConfiguration;
        this.name = $.name;
        this.opensearchConfiguration = $.opensearchConfiguration;
        this.redshiftConfiguration = $.redshiftConfiguration;
        this.s3Configuration = $.s3Configuration;
        this.serverSideEncryption = $.serverSideEncryption;
        this.splunkConfiguration = $.splunkConfiguration;
        this.tags = $.tags;
        this.versionId = $.versionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FirehoseDeliveryStreamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FirehoseDeliveryStreamArgs $;

        public Builder() {
            $ = new FirehoseDeliveryStreamArgs();
        }

        public Builder(FirehoseDeliveryStreamArgs defaults) {
            $ = new FirehoseDeliveryStreamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) specifying the Stream
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) specifying the Stream
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param destination This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, `splunk`, `http_endpoint` and `opensearch`.
         * 
         * @return builder
         * 
         */
        public Builder destination(Output<String> destination) {
            $.destination = destination;
            return this;
        }

        /**
         * @param destination This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, `splunk`, `http_endpoint` and `opensearch`.
         * 
         * @return builder
         * 
         */
        public Builder destination(String destination) {
            return destination(Output.of(destination));
        }

        public Builder destinationId(@Nullable Output<String> destinationId) {
            $.destinationId = destinationId;
            return this;
        }

        public Builder destinationId(String destinationId) {
            return destinationId(Output.of(destinationId));
        }

        /**
         * @param elasticsearchConfiguration Configuration options if elasticsearch is the destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder elasticsearchConfiguration(@Nullable Output<FirehoseDeliveryStreamElasticsearchConfigurationArgs> elasticsearchConfiguration) {
            $.elasticsearchConfiguration = elasticsearchConfiguration;
            return this;
        }

        /**
         * @param elasticsearchConfiguration Configuration options if elasticsearch is the destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder elasticsearchConfiguration(FirehoseDeliveryStreamElasticsearchConfigurationArgs elasticsearchConfiguration) {
            return elasticsearchConfiguration(Output.of(elasticsearchConfiguration));
        }

        /**
         * @param extendedS3Configuration Enhanced configuration options for the s3 destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder extendedS3Configuration(@Nullable Output<FirehoseDeliveryStreamExtendedS3ConfigurationArgs> extendedS3Configuration) {
            $.extendedS3Configuration = extendedS3Configuration;
            return this;
        }

        /**
         * @param extendedS3Configuration Enhanced configuration options for the s3 destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs extendedS3Configuration) {
            return extendedS3Configuration(Output.of(extendedS3Configuration));
        }

        /**
         * @param httpEndpointConfiguration Configuration options if http_endpoint is the destination. requires the user to also specify a `s3_configuration` block.  More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder httpEndpointConfiguration(@Nullable Output<FirehoseDeliveryStreamHttpEndpointConfigurationArgs> httpEndpointConfiguration) {
            $.httpEndpointConfiguration = httpEndpointConfiguration;
            return this;
        }

        /**
         * @param httpEndpointConfiguration Configuration options if http_endpoint is the destination. requires the user to also specify a `s3_configuration` block.  More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder httpEndpointConfiguration(FirehoseDeliveryStreamHttpEndpointConfigurationArgs httpEndpointConfiguration) {
            return httpEndpointConfiguration(Output.of(httpEndpointConfiguration));
        }

        /**
         * @param kinesisSourceConfiguration Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
         * 
         * @return builder
         * 
         */
        public Builder kinesisSourceConfiguration(@Nullable Output<FirehoseDeliveryStreamKinesisSourceConfigurationArgs> kinesisSourceConfiguration) {
            $.kinesisSourceConfiguration = kinesisSourceConfiguration;
            return this;
        }

        /**
         * @param kinesisSourceConfiguration Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
         * 
         * @return builder
         * 
         */
        public Builder kinesisSourceConfiguration(FirehoseDeliveryStreamKinesisSourceConfigurationArgs kinesisSourceConfiguration) {
            return kinesisSourceConfiguration(Output.of(kinesisSourceConfiguration));
        }

        /**
         * @param name A name to identify the stream. This is unique to the AWS account and region the Stream is created in. When using for WAF logging, name must be prefixed with `aws-waf-logs-`. See [AWS Documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-policies.html#waf-policies-logging-config) for more details.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name to identify the stream. This is unique to the AWS account and region the Stream is created in. When using for WAF logging, name must be prefixed with `aws-waf-logs-`. See [AWS Documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-policies.html#waf-policies-logging-config) for more details.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param opensearchConfiguration Configuration options if opensearch is the destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder opensearchConfiguration(@Nullable Output<FirehoseDeliveryStreamOpensearchConfigurationArgs> opensearchConfiguration) {
            $.opensearchConfiguration = opensearchConfiguration;
            return this;
        }

        /**
         * @param opensearchConfiguration Configuration options if opensearch is the destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder opensearchConfiguration(FirehoseDeliveryStreamOpensearchConfigurationArgs opensearchConfiguration) {
            return opensearchConfiguration(Output.of(opensearchConfiguration));
        }

        /**
         * @param redshiftConfiguration Configuration options if redshift is the destination.
         * Using `redshift_configuration` requires the user to also specify a
         * `s3_configuration` block. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder redshiftConfiguration(@Nullable Output<FirehoseDeliveryStreamRedshiftConfigurationArgs> redshiftConfiguration) {
            $.redshiftConfiguration = redshiftConfiguration;
            return this;
        }

        /**
         * @param redshiftConfiguration Configuration options if redshift is the destination.
         * Using `redshift_configuration` requires the user to also specify a
         * `s3_configuration` block. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder redshiftConfiguration(FirehoseDeliveryStreamRedshiftConfigurationArgs redshiftConfiguration) {
            return redshiftConfiguration(Output.of(redshiftConfiguration));
        }

        /**
         * @param s3Configuration Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
         * is redshift). More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder s3Configuration(@Nullable Output<FirehoseDeliveryStreamS3ConfigurationArgs> s3Configuration) {
            $.s3Configuration = s3Configuration;
            return this;
        }

        /**
         * @param s3Configuration Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
         * is redshift). More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder s3Configuration(FirehoseDeliveryStreamS3ConfigurationArgs s3Configuration) {
            return s3Configuration(Output.of(s3Configuration));
        }

        /**
         * @param serverSideEncryption Encrypt at rest options.
         * Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
         * 
         * @return builder
         * 
         */
        public Builder serverSideEncryption(@Nullable Output<FirehoseDeliveryStreamServerSideEncryptionArgs> serverSideEncryption) {
            $.serverSideEncryption = serverSideEncryption;
            return this;
        }

        /**
         * @param serverSideEncryption Encrypt at rest options.
         * Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
         * 
         * @return builder
         * 
         */
        public Builder serverSideEncryption(FirehoseDeliveryStreamServerSideEncryptionArgs serverSideEncryption) {
            return serverSideEncryption(Output.of(serverSideEncryption));
        }

        /**
         * @param splunkConfiguration Configuration options if splunk is the destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder splunkConfiguration(@Nullable Output<FirehoseDeliveryStreamSplunkConfigurationArgs> splunkConfiguration) {
            $.splunkConfiguration = splunkConfiguration;
            return this;
        }

        /**
         * @param splunkConfiguration Configuration options if splunk is the destination. More details are given below.
         * 
         * @return builder
         * 
         */
        public Builder splunkConfiguration(FirehoseDeliveryStreamSplunkConfigurationArgs splunkConfiguration) {
            return splunkConfiguration(Output.of(splunkConfiguration));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param versionId Specifies the table version for the output data schema. Defaults to `LATEST`.
         * 
         * @return builder
         * 
         */
        public Builder versionId(@Nullable Output<String> versionId) {
            $.versionId = versionId;
            return this;
        }

        /**
         * @param versionId Specifies the table version for the output data schema. Defaults to `LATEST`.
         * 
         * @return builder
         * 
         */
        public Builder versionId(String versionId) {
            return versionId(Output.of(versionId));
        }

        public FirehoseDeliveryStreamArgs build() {
            $.destination = Objects.requireNonNull($.destination, "expected parameter 'destination' to be non-null");
            return $;
        }
    }

}
