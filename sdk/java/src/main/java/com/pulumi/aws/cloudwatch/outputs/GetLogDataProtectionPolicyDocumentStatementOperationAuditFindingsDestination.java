// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch.outputs;

import com.pulumi.aws.cloudwatch.outputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs;
import com.pulumi.aws.cloudwatch.outputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose;
import com.pulumi.aws.cloudwatch.outputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination {
    /**
     * @return Configures CloudWatch Logs as a findings destination.
     * 
     */
    private @Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs cloudwatchLogs;
    /**
     * @return Configures Kinesis Firehose as a findings destination.
     * 
     */
    private @Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose firehose;
    /**
     * @return Configures S3 as a findings destination.
     * 
     */
    private @Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3 s3;

    private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination() {}
    /**
     * @return Configures CloudWatch Logs as a findings destination.
     * 
     */
    public Optional<GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs> cloudwatchLogs() {
        return Optional.ofNullable(this.cloudwatchLogs);
    }
    /**
     * @return Configures Kinesis Firehose as a findings destination.
     * 
     */
    public Optional<GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose> firehose() {
        return Optional.ofNullable(this.firehose);
    }
    /**
     * @return Configures S3 as a findings destination.
     * 
     */
    public Optional<GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3> s3() {
        return Optional.ofNullable(this.s3);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs cloudwatchLogs;
        private @Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose firehose;
        private @Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3 s3;
        public Builder() {}
        public Builder(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cloudwatchLogs = defaults.cloudwatchLogs;
    	      this.firehose = defaults.firehose;
    	      this.s3 = defaults.s3;
        }

        @CustomType.Setter
        public Builder cloudwatchLogs(@Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs cloudwatchLogs) {
            this.cloudwatchLogs = cloudwatchLogs;
            return this;
        }
        @CustomType.Setter
        public Builder firehose(@Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose firehose) {
            this.firehose = firehose;
            return this;
        }
        @CustomType.Setter
        public Builder s3(@Nullable GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3 s3) {
            this.s3 = s3;
            return this;
        }
        public GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination build() {
            final var o = new GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination();
            o.cloudwatchLogs = cloudwatchLogs;
            o.firehose = firehose;
            o.s3 = s3;
            return o;
        }
    }
}
