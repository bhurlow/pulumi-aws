// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.secretsmanager;

import com.pulumi.aws.secretsmanager.inputs.SecretRotationRotationRulesArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class SecretRotationArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretRotationArgs Empty = new SecretRotationArgs();

    /**
     * Specifies the ARN of the Lambda function that can rotate the secret.
     * 
     */
    @Import(name="rotationLambdaArn", required=true)
    private Output<String> rotationLambdaArn;

    /**
     * @return Specifies the ARN of the Lambda function that can rotate the secret.
     * 
     */
    public Output<String> rotationLambdaArn() {
        return this.rotationLambdaArn;
    }

    /**
     * A structure that defines the rotation configuration for this secret. Defined below.
     * 
     */
    @Import(name="rotationRules", required=true)
    private Output<SecretRotationRotationRulesArgs> rotationRules;

    /**
     * @return A structure that defines the rotation configuration for this secret. Defined below.
     * 
     */
    public Output<SecretRotationRotationRulesArgs> rotationRules() {
        return this.rotationRules;
    }

    /**
     * Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
     * 
     */
    @Import(name="secretId", required=true)
    private Output<String> secretId;

    /**
     * @return Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
     * 
     */
    public Output<String> secretId() {
        return this.secretId;
    }

    private SecretRotationArgs() {}

    private SecretRotationArgs(SecretRotationArgs $) {
        this.rotationLambdaArn = $.rotationLambdaArn;
        this.rotationRules = $.rotationRules;
        this.secretId = $.secretId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretRotationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretRotationArgs $;

        public Builder() {
            $ = new SecretRotationArgs();
        }

        public Builder(SecretRotationArgs defaults) {
            $ = new SecretRotationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param rotationLambdaArn Specifies the ARN of the Lambda function that can rotate the secret.
         * 
         * @return builder
         * 
         */
        public Builder rotationLambdaArn(Output<String> rotationLambdaArn) {
            $.rotationLambdaArn = rotationLambdaArn;
            return this;
        }

        /**
         * @param rotationLambdaArn Specifies the ARN of the Lambda function that can rotate the secret.
         * 
         * @return builder
         * 
         */
        public Builder rotationLambdaArn(String rotationLambdaArn) {
            return rotationLambdaArn(Output.of(rotationLambdaArn));
        }

        /**
         * @param rotationRules A structure that defines the rotation configuration for this secret. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder rotationRules(Output<SecretRotationRotationRulesArgs> rotationRules) {
            $.rotationRules = rotationRules;
            return this;
        }

        /**
         * @param rotationRules A structure that defines the rotation configuration for this secret. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder rotationRules(SecretRotationRotationRulesArgs rotationRules) {
            return rotationRules(Output.of(rotationRules));
        }

        /**
         * @param secretId Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
         * 
         * @return builder
         * 
         */
        public Builder secretId(Output<String> secretId) {
            $.secretId = secretId;
            return this;
        }

        /**
         * @param secretId Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
         * 
         * @return builder
         * 
         */
        public Builder secretId(String secretId) {
            return secretId(Output.of(secretId));
        }

        public SecretRotationArgs build() {
            $.rotationLambdaArn = Objects.requireNonNull($.rotationLambdaArn, "expected parameter 'rotationLambdaArn' to be non-null");
            $.rotationRules = Objects.requireNonNull($.rotationRules, "expected parameter 'rotationRules' to be non-null");
            $.secretId = Objects.requireNonNull($.secretId, "expected parameter 'secretId' to be non-null");
            return $;
        }
    }

}
