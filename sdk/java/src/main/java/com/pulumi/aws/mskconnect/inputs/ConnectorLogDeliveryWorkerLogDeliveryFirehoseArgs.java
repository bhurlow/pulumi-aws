// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.mskconnect.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs Empty = new ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs();

    /**
     * The name of the Kinesis Data Firehose delivery stream that is the destination for log delivery.
     * 
     */
    @Import(name="deliveryStream")
    private @Nullable Output<String> deliveryStream;

    /**
     * @return The name of the Kinesis Data Firehose delivery stream that is the destination for log delivery.
     * 
     */
    public Optional<Output<String>> deliveryStream() {
        return Optional.ofNullable(this.deliveryStream);
    }

    /**
     * Specifies whether connector logs get sent to the specified Amazon S3 destination.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Specifies whether connector logs get sent to the specified Amazon S3 destination.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    private ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs() {}

    private ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs(ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs $) {
        this.deliveryStream = $.deliveryStream;
        this.enabled = $.enabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs $;

        public Builder() {
            $ = new ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs();
        }

        public Builder(ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs defaults) {
            $ = new ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deliveryStream The name of the Kinesis Data Firehose delivery stream that is the destination for log delivery.
         * 
         * @return builder
         * 
         */
        public Builder deliveryStream(@Nullable Output<String> deliveryStream) {
            $.deliveryStream = deliveryStream;
            return this;
        }

        /**
         * @param deliveryStream The name of the Kinesis Data Firehose delivery stream that is the destination for log delivery.
         * 
         * @return builder
         * 
         */
        public Builder deliveryStream(String deliveryStream) {
            return deliveryStream(Output.of(deliveryStream));
        }

        /**
         * @param enabled Specifies whether connector logs get sent to the specified Amazon S3 destination.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether connector logs get sent to the specified Amazon S3 destination.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public ConnectorLogDeliveryWorkerLogDeliveryFirehoseArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            return $;
        }
    }

}