// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourcePolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourcePolicyArgs Empty = new ResourcePolicyArgs();

    /**
     * Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
     * 
     */
    @Import(name="enableHybrid")
    private @Nullable Output<String> enableHybrid;

    /**
     * @return Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
     * 
     */
    public Optional<Output<String>> enableHybrid() {
        return Optional.ofNullable(this.enableHybrid);
    }

    /**
     * The policy to be applied to the aws glue data catalog.
     * 
     */
    @Import(name="policy", required=true)
    private Output<String> policy;

    /**
     * @return The policy to be applied to the aws glue data catalog.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    private ResourcePolicyArgs() {}

    private ResourcePolicyArgs(ResourcePolicyArgs $) {
        this.enableHybrid = $.enableHybrid;
        this.policy = $.policy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourcePolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourcePolicyArgs $;

        public Builder() {
            $ = new ResourcePolicyArgs();
        }

        public Builder(ResourcePolicyArgs defaults) {
            $ = new ResourcePolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableHybrid Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
         * 
         * @return builder
         * 
         */
        public Builder enableHybrid(@Nullable Output<String> enableHybrid) {
            $.enableHybrid = enableHybrid;
            return this;
        }

        /**
         * @param enableHybrid Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
         * 
         * @return builder
         * 
         */
        public Builder enableHybrid(String enableHybrid) {
            return enableHybrid(Output.of(enableHybrid));
        }

        /**
         * @param policy The policy to be applied to the aws glue data catalog.
         * 
         * @return builder
         * 
         */
        public Builder policy(Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The policy to be applied to the aws glue data catalog.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public ResourcePolicyArgs build() {
            $.policy = Objects.requireNonNull($.policy, "expected parameter 'policy' to be non-null");
            return $;
        }
    }

}
