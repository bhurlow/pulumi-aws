// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sesv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetEmailIdentityMailFromAttributesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEmailIdentityMailFromAttributesArgs Empty = new GetEmailIdentityMailFromAttributesArgs();

    /**
     * The name of the email identity.
     * 
     */
    @Import(name="emailIdentity", required=true)
    private Output<String> emailIdentity;

    /**
     * @return The name of the email identity.
     * 
     */
    public Output<String> emailIdentity() {
        return this.emailIdentity;
    }

    private GetEmailIdentityMailFromAttributesArgs() {}

    private GetEmailIdentityMailFromAttributesArgs(GetEmailIdentityMailFromAttributesArgs $) {
        this.emailIdentity = $.emailIdentity;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEmailIdentityMailFromAttributesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEmailIdentityMailFromAttributesArgs $;

        public Builder() {
            $ = new GetEmailIdentityMailFromAttributesArgs();
        }

        public Builder(GetEmailIdentityMailFromAttributesArgs defaults) {
            $ = new GetEmailIdentityMailFromAttributesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param emailIdentity The name of the email identity.
         * 
         * @return builder
         * 
         */
        public Builder emailIdentity(Output<String> emailIdentity) {
            $.emailIdentity = emailIdentity;
            return this;
        }

        /**
         * @param emailIdentity The name of the email identity.
         * 
         * @return builder
         * 
         */
        public Builder emailIdentity(String emailIdentity) {
            return emailIdentity(Output.of(emailIdentity));
        }

        public GetEmailIdentityMailFromAttributesArgs build() {
            $.emailIdentity = Objects.requireNonNull($.emailIdentity, "expected parameter 'emailIdentity' to be non-null");
            return $;
        }
    }

}
