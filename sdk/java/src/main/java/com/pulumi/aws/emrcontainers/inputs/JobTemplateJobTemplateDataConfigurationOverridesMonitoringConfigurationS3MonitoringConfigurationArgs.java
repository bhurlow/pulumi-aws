// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.emrcontainers.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs Empty = new JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs();

    /**
     * Amazon S3 destination URI for log publishing.
     * 
     */
    @Import(name="logUri", required=true)
    private Output<String> logUri;

    /**
     * @return Amazon S3 destination URI for log publishing.
     * 
     */
    public Output<String> logUri() {
        return this.logUri;
    }

    private JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs() {}

    private JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs(JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs $) {
        this.logUri = $.logUri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs $;

        public Builder() {
            $ = new JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs();
        }

        public Builder(JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs defaults) {
            $ = new JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param logUri Amazon S3 destination URI for log publishing.
         * 
         * @return builder
         * 
         */
        public Builder logUri(Output<String> logUri) {
            $.logUri = logUri;
            return this;
        }

        /**
         * @param logUri Amazon S3 destination URI for log publishing.
         * 
         * @return builder
         * 
         */
        public Builder logUri(String logUri) {
            return logUri(Output.of(logUri));
        }

        public JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfigurationArgs build() {
            $.logUri = Objects.requireNonNull($.logUri, "expected parameter 'logUri' to be non-null");
            return $;
        }
    }

}
