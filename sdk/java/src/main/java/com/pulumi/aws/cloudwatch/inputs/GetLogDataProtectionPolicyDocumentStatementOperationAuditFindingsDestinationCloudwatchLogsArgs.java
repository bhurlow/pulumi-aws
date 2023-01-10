// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs Empty = new GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs();

    /**
     * Name of the CloudWatch Log Group to send findings to.
     * 
     */
    @Import(name="logGroup", required=true)
    private Output<String> logGroup;

    /**
     * @return Name of the CloudWatch Log Group to send findings to.
     * 
     */
    public Output<String> logGroup() {
        return this.logGroup;
    }

    private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs() {}

    private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs $) {
        this.logGroup = $.logGroup;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs $;

        public Builder() {
            $ = new GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs();
        }

        public Builder(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs defaults) {
            $ = new GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param logGroup Name of the CloudWatch Log Group to send findings to.
         * 
         * @return builder
         * 
         */
        public Builder logGroup(Output<String> logGroup) {
            $.logGroup = logGroup;
            return this;
        }

        /**
         * @param logGroup Name of the CloudWatch Log Group to send findings to.
         * 
         * @return builder
         * 
         */
        public Builder logGroup(String logGroup) {
            return logGroup(Output.of(logGroup));
        }

        public GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs build() {
            $.logGroup = Objects.requireNonNull($.logGroup, "expected parameter 'logGroup' to be non-null");
            return $;
        }
    }

}
