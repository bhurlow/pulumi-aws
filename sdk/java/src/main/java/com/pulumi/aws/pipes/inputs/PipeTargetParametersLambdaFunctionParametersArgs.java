// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class PipeTargetParametersLambdaFunctionParametersArgs extends com.pulumi.resources.ResourceArgs {

    public static final PipeTargetParametersLambdaFunctionParametersArgs Empty = new PipeTargetParametersLambdaFunctionParametersArgs();

    /**
     * Specify whether to invoke the function synchronously or asynchronously. Valid Values: REQUEST_RESPONSE, FIRE_AND_FORGET.
     * 
     */
    @Import(name="invocationType", required=true)
    private Output<String> invocationType;

    /**
     * @return Specify whether to invoke the function synchronously or asynchronously. Valid Values: REQUEST_RESPONSE, FIRE_AND_FORGET.
     * 
     */
    public Output<String> invocationType() {
        return this.invocationType;
    }

    private PipeTargetParametersLambdaFunctionParametersArgs() {}

    private PipeTargetParametersLambdaFunctionParametersArgs(PipeTargetParametersLambdaFunctionParametersArgs $) {
        this.invocationType = $.invocationType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipeTargetParametersLambdaFunctionParametersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipeTargetParametersLambdaFunctionParametersArgs $;

        public Builder() {
            $ = new PipeTargetParametersLambdaFunctionParametersArgs();
        }

        public Builder(PipeTargetParametersLambdaFunctionParametersArgs defaults) {
            $ = new PipeTargetParametersLambdaFunctionParametersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param invocationType Specify whether to invoke the function synchronously or asynchronously. Valid Values: REQUEST_RESPONSE, FIRE_AND_FORGET.
         * 
         * @return builder
         * 
         */
        public Builder invocationType(Output<String> invocationType) {
            $.invocationType = invocationType;
            return this;
        }

        /**
         * @param invocationType Specify whether to invoke the function synchronously or asynchronously. Valid Values: REQUEST_RESPONSE, FIRE_AND_FORGET.
         * 
         * @return builder
         * 
         */
        public Builder invocationType(String invocationType) {
            return invocationType(Output.of(invocationType));
        }

        public PipeTargetParametersLambdaFunctionParametersArgs build() {
            $.invocationType = Objects.requireNonNull($.invocationType, "expected parameter 'invocationType' to be non-null");
            return $;
        }
    }

}
