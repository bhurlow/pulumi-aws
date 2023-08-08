// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.emrcontainers.outputs;

import com.pulumi.aws.emrcontainers.outputs.JobTemplateJobTemplateDataJobDriverSparkSqlJobDriver;
import com.pulumi.aws.emrcontainers.outputs.JobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class JobTemplateJobTemplateDataJobDriver {
    /**
     * @return The job driver for job type.
     * 
     */
    private @Nullable JobTemplateJobTemplateDataJobDriverSparkSqlJobDriver sparkSqlJobDriver;
    /**
     * @return The job driver parameters specified for spark submit.
     * 
     */
    private @Nullable JobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver sparkSubmitJobDriver;

    private JobTemplateJobTemplateDataJobDriver() {}
    /**
     * @return The job driver for job type.
     * 
     */
    public Optional<JobTemplateJobTemplateDataJobDriverSparkSqlJobDriver> sparkSqlJobDriver() {
        return Optional.ofNullable(this.sparkSqlJobDriver);
    }
    /**
     * @return The job driver parameters specified for spark submit.
     * 
     */
    public Optional<JobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver> sparkSubmitJobDriver() {
        return Optional.ofNullable(this.sparkSubmitJobDriver);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(JobTemplateJobTemplateDataJobDriver defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable JobTemplateJobTemplateDataJobDriverSparkSqlJobDriver sparkSqlJobDriver;
        private @Nullable JobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver sparkSubmitJobDriver;
        public Builder() {}
        public Builder(JobTemplateJobTemplateDataJobDriver defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.sparkSqlJobDriver = defaults.sparkSqlJobDriver;
    	      this.sparkSubmitJobDriver = defaults.sparkSubmitJobDriver;
        }

        @CustomType.Setter
        public Builder sparkSqlJobDriver(@Nullable JobTemplateJobTemplateDataJobDriverSparkSqlJobDriver sparkSqlJobDriver) {
            this.sparkSqlJobDriver = sparkSqlJobDriver;
            return this;
        }
        @CustomType.Setter
        public Builder sparkSubmitJobDriver(@Nullable JobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver sparkSubmitJobDriver) {
            this.sparkSubmitJobDriver = sparkSubmitJobDriver;
            return this;
        }
        public JobTemplateJobTemplateDataJobDriver build() {
            final var o = new JobTemplateJobTemplateDataJobDriver();
            o.sparkSqlJobDriver = sparkSqlJobDriver;
            o.sparkSubmitJobDriver = sparkSubmitJobDriver;
            return o;
        }
    }
}
