// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.securitylake.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class SubscriberSubscriberIdentityArgs extends com.pulumi.resources.ResourceArgs {

    public static final SubscriberSubscriberIdentityArgs Empty = new SubscriberSubscriberIdentityArgs();

    /**
     * The AWS Regions where Security Lake is automatically enabled.
     * 
     */
    @Import(name="externalId", required=true)
    private Output<String> externalId;

    /**
     * @return The AWS Regions where Security Lake is automatically enabled.
     * 
     */
    public Output<String> externalId() {
        return this.externalId;
    }

    /**
     * Provides encryption details of Amazon Security Lake object.
     * 
     */
    @Import(name="principal", required=true)
    private Output<String> principal;

    /**
     * @return Provides encryption details of Amazon Security Lake object.
     * 
     */
    public Output<String> principal() {
        return this.principal;
    }

    private SubscriberSubscriberIdentityArgs() {}

    private SubscriberSubscriberIdentityArgs(SubscriberSubscriberIdentityArgs $) {
        this.externalId = $.externalId;
        this.principal = $.principal;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubscriberSubscriberIdentityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubscriberSubscriberIdentityArgs $;

        public Builder() {
            $ = new SubscriberSubscriberIdentityArgs();
        }

        public Builder(SubscriberSubscriberIdentityArgs defaults) {
            $ = new SubscriberSubscriberIdentityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param externalId The AWS Regions where Security Lake is automatically enabled.
         * 
         * @return builder
         * 
         */
        public Builder externalId(Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId The AWS Regions where Security Lake is automatically enabled.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param principal Provides encryption details of Amazon Security Lake object.
         * 
         * @return builder
         * 
         */
        public Builder principal(Output<String> principal) {
            $.principal = principal;
            return this;
        }

        /**
         * @param principal Provides encryption details of Amazon Security Lake object.
         * 
         * @return builder
         * 
         */
        public Builder principal(String principal) {
            return principal(Output.of(principal));
        }

        public SubscriberSubscriberIdentityArgs build() {
            if ($.externalId == null) {
                throw new MissingRequiredPropertyException("SubscriberSubscriberIdentityArgs", "externalId");
            }
            if ($.principal == null) {
                throw new MissingRequiredPropertyException("SubscriberSubscriberIdentityArgs", "principal");
            }
            return $;
        }
    }

}
