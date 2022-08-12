// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesisanalyticsv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs Empty = new ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs();

    /**
     * The ARN of the input Kinesis data stream to read.
     * 
     */
    @Import(name="resourceArn", required=true)
    private Output<String> resourceArn;

    /**
     * @return The ARN of the input Kinesis data stream to read.
     * 
     */
    public Output<String> resourceArn() {
        return this.resourceArn;
    }

    private ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs() {}

    private ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs(ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs $) {
        this.resourceArn = $.resourceArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs $;

        public Builder() {
            $ = new ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs();
        }

        public Builder(ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs defaults) {
            $ = new ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param resourceArn The ARN of the input Kinesis data stream to read.
         * 
         * @return builder
         * 
         */
        public Builder resourceArn(Output<String> resourceArn) {
            $.resourceArn = resourceArn;
            return this;
        }

        /**
         * @param resourceArn The ARN of the input Kinesis data stream to read.
         * 
         * @return builder
         * 
         */
        public Builder resourceArn(String resourceArn) {
            return resourceArn(Output.of(resourceArn));
        }

        public ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs build() {
            $.resourceArn = Objects.requireNonNull($.resourceArn, "expected parameter 'resourceArn' to be non-null");
            return $;
        }
    }

}