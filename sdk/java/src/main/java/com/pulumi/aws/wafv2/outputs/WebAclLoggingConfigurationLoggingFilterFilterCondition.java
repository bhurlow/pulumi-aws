// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.outputs;

import com.pulumi.aws.wafv2.outputs.WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition;
import com.pulumi.aws.wafv2.outputs.WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WebAclLoggingConfigurationLoggingFilterFilterCondition {
    /**
     * @return Configuration for a single action condition. See Action Condition below for more details.
     * 
     */
    private @Nullable WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition actionCondition;
    /**
     * @return Condition for a single label name. See Label Name Condition below for more details.
     * 
     */
    private @Nullable WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition labelNameCondition;

    private WebAclLoggingConfigurationLoggingFilterFilterCondition() {}
    /**
     * @return Configuration for a single action condition. See Action Condition below for more details.
     * 
     */
    public Optional<WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition> actionCondition() {
        return Optional.ofNullable(this.actionCondition);
    }
    /**
     * @return Condition for a single label name. See Label Name Condition below for more details.
     * 
     */
    public Optional<WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition> labelNameCondition() {
        return Optional.ofNullable(this.labelNameCondition);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WebAclLoggingConfigurationLoggingFilterFilterCondition defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition actionCondition;
        private @Nullable WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition labelNameCondition;
        public Builder() {}
        public Builder(WebAclLoggingConfigurationLoggingFilterFilterCondition defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actionCondition = defaults.actionCondition;
    	      this.labelNameCondition = defaults.labelNameCondition;
        }

        @CustomType.Setter
        public Builder actionCondition(@Nullable WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition actionCondition) {
            this.actionCondition = actionCondition;
            return this;
        }
        @CustomType.Setter
        public Builder labelNameCondition(@Nullable WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition labelNameCondition) {
            this.labelNameCondition = labelNameCondition;
            return this;
        }
        public WebAclLoggingConfigurationLoggingFilterFilterCondition build() {
            final var o = new WebAclLoggingConfigurationLoggingFilterFilterCondition();
            o.actionCondition = actionCondition;
            o.labelNameCondition = labelNameCondition;
            return o;
        }
    }
}
