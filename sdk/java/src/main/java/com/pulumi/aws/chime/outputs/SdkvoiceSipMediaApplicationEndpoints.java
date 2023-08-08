// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.chime.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SdkvoiceSipMediaApplicationEndpoints {
    /**
     * @return Valid Amazon Resource Name (ARN) of the Lambda function, version, or alias. The function must be created in the same AWS Region as the SIP media application.
     * 
     */
    private String lambdaArn;

    private SdkvoiceSipMediaApplicationEndpoints() {}
    /**
     * @return Valid Amazon Resource Name (ARN) of the Lambda function, version, or alias. The function must be created in the same AWS Region as the SIP media application.
     * 
     */
    public String lambdaArn() {
        return this.lambdaArn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SdkvoiceSipMediaApplicationEndpoints defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String lambdaArn;
        public Builder() {}
        public Builder(SdkvoiceSipMediaApplicationEndpoints defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.lambdaArn = defaults.lambdaArn;
        }

        @CustomType.Setter
        public Builder lambdaArn(String lambdaArn) {
            this.lambdaArn = Objects.requireNonNull(lambdaArn);
            return this;
        }
        public SdkvoiceSipMediaApplicationEndpoints build() {
            final var o = new SdkvoiceSipMediaApplicationEndpoints();
            o.lambdaArn = lambdaArn;
            return o;
        }
    }
}
