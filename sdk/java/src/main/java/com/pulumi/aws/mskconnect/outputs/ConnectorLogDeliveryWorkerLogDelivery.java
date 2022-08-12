// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.mskconnect.outputs;

import com.pulumi.aws.mskconnect.outputs.ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs;
import com.pulumi.aws.mskconnect.outputs.ConnectorLogDeliveryWorkerLogDeliveryFirehose;
import com.pulumi.aws.mskconnect.outputs.ConnectorLogDeliveryWorkerLogDeliveryS3;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConnectorLogDeliveryWorkerLogDelivery {
    /**
     * @return Details about delivering logs to Amazon CloudWatch Logs. See below.
     * 
     */
    private final @Nullable ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs cloudwatchLogs;
    /**
     * @return Details about delivering logs to Amazon Kinesis Data Firehose. See below.
     * 
     */
    private final @Nullable ConnectorLogDeliveryWorkerLogDeliveryFirehose firehose;
    /**
     * @return Details about delivering logs to Amazon S3. See below.
     * 
     */
    private final @Nullable ConnectorLogDeliveryWorkerLogDeliveryS3 s3;

    @CustomType.Constructor
    private ConnectorLogDeliveryWorkerLogDelivery(
        @CustomType.Parameter("cloudwatchLogs") @Nullable ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs cloudwatchLogs,
        @CustomType.Parameter("firehose") @Nullable ConnectorLogDeliveryWorkerLogDeliveryFirehose firehose,
        @CustomType.Parameter("s3") @Nullable ConnectorLogDeliveryWorkerLogDeliveryS3 s3) {
        this.cloudwatchLogs = cloudwatchLogs;
        this.firehose = firehose;
        this.s3 = s3;
    }

    /**
     * @return Details about delivering logs to Amazon CloudWatch Logs. See below.
     * 
     */
    public Optional<ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs> cloudwatchLogs() {
        return Optional.ofNullable(this.cloudwatchLogs);
    }
    /**
     * @return Details about delivering logs to Amazon Kinesis Data Firehose. See below.
     * 
     */
    public Optional<ConnectorLogDeliveryWorkerLogDeliveryFirehose> firehose() {
        return Optional.ofNullable(this.firehose);
    }
    /**
     * @return Details about delivering logs to Amazon S3. See below.
     * 
     */
    public Optional<ConnectorLogDeliveryWorkerLogDeliveryS3> s3() {
        return Optional.ofNullable(this.s3);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConnectorLogDeliveryWorkerLogDelivery defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs cloudwatchLogs;
        private @Nullable ConnectorLogDeliveryWorkerLogDeliveryFirehose firehose;
        private @Nullable ConnectorLogDeliveryWorkerLogDeliveryS3 s3;

        public Builder() {
    	      // Empty
        }

        public Builder(ConnectorLogDeliveryWorkerLogDelivery defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cloudwatchLogs = defaults.cloudwatchLogs;
    	      this.firehose = defaults.firehose;
    	      this.s3 = defaults.s3;
        }

        public Builder cloudwatchLogs(@Nullable ConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs cloudwatchLogs) {
            this.cloudwatchLogs = cloudwatchLogs;
            return this;
        }
        public Builder firehose(@Nullable ConnectorLogDeliveryWorkerLogDeliveryFirehose firehose) {
            this.firehose = firehose;
            return this;
        }
        public Builder s3(@Nullable ConnectorLogDeliveryWorkerLogDeliveryS3 s3) {
            this.s3 = s3;
            return this;
        }        public ConnectorLogDeliveryWorkerLogDelivery build() {
            return new ConnectorLogDeliveryWorkerLogDelivery(cloudwatchLogs, firehose, s3);
        }
    }
}