// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose extends com.pulumi.resources.InvokeArgs {

    public static final GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose Empty = new GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose();

    /**
     * Name of the Kinesis Firehose Delivery Stream to send findings to.
     * 
     */
    @Import(name="deliveryStream", required=true)
    private String deliveryStream;

    /**
     * @return Name of the Kinesis Firehose Delivery Stream to send findings to.
     * 
     */
    public String deliveryStream() {
        return this.deliveryStream;
    }

    private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose() {}

    private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose $) {
        this.deliveryStream = $.deliveryStream;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose $;

        public Builder() {
            $ = new GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose();
        }

        public Builder(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose defaults) {
            $ = new GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose(Objects.requireNonNull(defaults));
        }

        /**
         * @param deliveryStream Name of the Kinesis Firehose Delivery Stream to send findings to.
         * 
         * @return builder
         * 
         */
        public Builder deliveryStream(String deliveryStream) {
            $.deliveryStream = deliveryStream;
            return this;
        }

        public GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose build() {
            $.deliveryStream = Objects.requireNonNull($.deliveryStream, "expected parameter 'deliveryStream' to be non-null");
            return $;
        }
    }

}
