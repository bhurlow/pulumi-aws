// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.batch.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.Object;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetJobDefinitionRetryStrategy {
    /**
     * @return The number of times to move a job to the RUNNABLE status.
     * 
     */
    private Integer attempts;
    /**
     * @return Array of up to 5 objects that specify the conditions where jobs are retried or failed.
     * 
     */
    private List<Object> evaluateOnExits;

    private GetJobDefinitionRetryStrategy() {}
    /**
     * @return The number of times to move a job to the RUNNABLE status.
     * 
     */
    public Integer attempts() {
        return this.attempts;
    }
    /**
     * @return Array of up to 5 objects that specify the conditions where jobs are retried or failed.
     * 
     */
    public List<Object> evaluateOnExits() {
        return this.evaluateOnExits;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetJobDefinitionRetryStrategy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer attempts;
        private List<Object> evaluateOnExits;
        public Builder() {}
        public Builder(GetJobDefinitionRetryStrategy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attempts = defaults.attempts;
    	      this.evaluateOnExits = defaults.evaluateOnExits;
        }

        @CustomType.Setter
        public Builder attempts(Integer attempts) {
            if (attempts == null) {
              throw new MissingRequiredPropertyException("GetJobDefinitionRetryStrategy", "attempts");
            }
            this.attempts = attempts;
            return this;
        }
        @CustomType.Setter
        public Builder evaluateOnExits(List<Object> evaluateOnExits) {
            if (evaluateOnExits == null) {
              throw new MissingRequiredPropertyException("GetJobDefinitionRetryStrategy", "evaluateOnExits");
            }
            this.evaluateOnExits = evaluateOnExits;
            return this;
        }
        public Builder evaluateOnExits(Object... evaluateOnExits) {
            return evaluateOnExits(List.of(evaluateOnExits));
        }
        public GetJobDefinitionRetryStrategy build() {
            final var _resultValue = new GetJobDefinitionRetryStrategy();
            _resultValue.attempts = attempts;
            _resultValue.evaluateOnExits = evaluateOnExits;
            return _resultValue;
        }
    }
}
