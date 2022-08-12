// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetArnArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetArnArgs Empty = new GetArnArgs();

    /**
     * The ARN to parse.
     * 
     */
    @Import(name="arn", required=true)
    private Output<String> arn;

    /**
     * @return The ARN to parse.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }

    private GetArnArgs() {}

    private GetArnArgs(GetArnArgs $) {
        this.arn = $.arn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetArnArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetArnArgs $;

        public Builder() {
            $ = new GetArnArgs();
        }

        public Builder(GetArnArgs defaults) {
            $ = new GetArnArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN to parse.
         * 
         * @return builder
         * 
         */
        public Builder arn(Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN to parse.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        public GetArnArgs build() {
            $.arn = Objects.requireNonNull($.arn, "expected parameter 'arn' to be non-null");
            return $;
        }
    }

}