// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.inspector2.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrganizationConfigurationAutoEnable {
    /**
     * @return Whether Amazon EC2 scans are automatically enabled for new members of your Amazon Inspector organization.
     * 
     */
    private Boolean ec2;
    /**
     * @return Whether Amazon ECR scans are automatically enabled for new members of your Amazon Inspector organization.
     * 
     */
    private Boolean ecr;
    /**
     * @return Whether Lambda Function scans are automatically enabled for new members of your Amazon Inspector organization.
     * 
     */
    private @Nullable Boolean lambda;
    /**
     * @return Whether AWS Lambda code scans are automatically enabled for new members of your Amazon Inspector organization. **Note:** Lambda code scanning requires Lambda standard scanning to be activated. Consequently, if you are setting this argument to `true`, you must also set the `lambda` argument to `true`. See [Scanning AWS Lambda functions with Amazon Inspector](https://docs.aws.amazon.com/inspector/latest/user/scanning-lambda.html#lambda-code-scans) for more information.
     * 
     */
    private @Nullable Boolean lambdaCode;

    private OrganizationConfigurationAutoEnable() {}
    /**
     * @return Whether Amazon EC2 scans are automatically enabled for new members of your Amazon Inspector organization.
     * 
     */
    public Boolean ec2() {
        return this.ec2;
    }
    /**
     * @return Whether Amazon ECR scans are automatically enabled for new members of your Amazon Inspector organization.
     * 
     */
    public Boolean ecr() {
        return this.ecr;
    }
    /**
     * @return Whether Lambda Function scans are automatically enabled for new members of your Amazon Inspector organization.
     * 
     */
    public Optional<Boolean> lambda() {
        return Optional.ofNullable(this.lambda);
    }
    /**
     * @return Whether AWS Lambda code scans are automatically enabled for new members of your Amazon Inspector organization. **Note:** Lambda code scanning requires Lambda standard scanning to be activated. Consequently, if you are setting this argument to `true`, you must also set the `lambda` argument to `true`. See [Scanning AWS Lambda functions with Amazon Inspector](https://docs.aws.amazon.com/inspector/latest/user/scanning-lambda.html#lambda-code-scans) for more information.
     * 
     */
    public Optional<Boolean> lambdaCode() {
        return Optional.ofNullable(this.lambdaCode);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationConfigurationAutoEnable defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean ec2;
        private Boolean ecr;
        private @Nullable Boolean lambda;
        private @Nullable Boolean lambdaCode;
        public Builder() {}
        public Builder(OrganizationConfigurationAutoEnable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ec2 = defaults.ec2;
    	      this.ecr = defaults.ecr;
    	      this.lambda = defaults.lambda;
    	      this.lambdaCode = defaults.lambdaCode;
        }

        @CustomType.Setter
        public Builder ec2(Boolean ec2) {
            this.ec2 = Objects.requireNonNull(ec2);
            return this;
        }
        @CustomType.Setter
        public Builder ecr(Boolean ecr) {
            this.ecr = Objects.requireNonNull(ecr);
            return this;
        }
        @CustomType.Setter
        public Builder lambda(@Nullable Boolean lambda) {
            this.lambda = lambda;
            return this;
        }
        @CustomType.Setter
        public Builder lambdaCode(@Nullable Boolean lambdaCode) {
            this.lambdaCode = lambdaCode;
            return this;
        }
        public OrganizationConfigurationAutoEnable build() {
            final var o = new OrganizationConfigurationAutoEnable();
            o.ec2 = ec2;
            o.ecr = ecr;
            o.lambda = lambda;
            o.lambdaCode = lambdaCode;
            return o;
        }
    }
}
