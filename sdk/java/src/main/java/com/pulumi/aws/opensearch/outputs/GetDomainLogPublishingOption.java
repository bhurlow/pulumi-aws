// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDomainLogPublishingOption {
    /**
     * @return CloudWatch Log Group where the logs are published.
     * 
     */
    private String cloudwatchLogGroupArn;
    /**
     * @return Enabled disabled toggle for off-peak update window
     * 
     */
    private Boolean enabled;
    /**
     * @return Type of OpenSearch log being published.
     * 
     */
    private String logType;

    private GetDomainLogPublishingOption() {}
    /**
     * @return CloudWatch Log Group where the logs are published.
     * 
     */
    public String cloudwatchLogGroupArn() {
        return this.cloudwatchLogGroupArn;
    }
    /**
     * @return Enabled disabled toggle for off-peak update window
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return Type of OpenSearch log being published.
     * 
     */
    public String logType() {
        return this.logType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainLogPublishingOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cloudwatchLogGroupArn;
        private Boolean enabled;
        private String logType;
        public Builder() {}
        public Builder(GetDomainLogPublishingOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cloudwatchLogGroupArn = defaults.cloudwatchLogGroupArn;
    	      this.enabled = defaults.enabled;
    	      this.logType = defaults.logType;
        }

        @CustomType.Setter
        public Builder cloudwatchLogGroupArn(String cloudwatchLogGroupArn) {
            this.cloudwatchLogGroupArn = Objects.requireNonNull(cloudwatchLogGroupArn);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder logType(String logType) {
            this.logType = Objects.requireNonNull(logType);
            return this;
        }
        public GetDomainLogPublishingOption build() {
            final var o = new GetDomainLogPublishingOption();
            o.cloudwatchLogGroupArn = cloudwatchLogGroupArn;
            o.enabled = enabled;
            o.logType = logType;
            return o;
        }
    }
}
