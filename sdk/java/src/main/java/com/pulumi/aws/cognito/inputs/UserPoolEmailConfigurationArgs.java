// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserPoolEmailConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserPoolEmailConfigurationArgs Empty = new UserPoolEmailConfigurationArgs();

    /**
     * Email configuration set name from SES.
     * 
     */
    @Import(name="configurationSet")
    private @Nullable Output<String> configurationSet;

    /**
     * @return Email configuration set name from SES.
     * 
     */
    public Optional<Output<String>> configurationSet() {
        return Optional.ofNullable(this.configurationSet);
    }

    /**
     * Email delivery method to use. `COGNITO_DEFAULT` for the default email functionality built into Cognito or `DEVELOPER` to use your Amazon SES configuration. Required to be `DEVELOPER` if `from_email_address` is set.
     * 
     */
    @Import(name="emailSendingAccount")
    private @Nullable Output<String> emailSendingAccount;

    /**
     * @return Email delivery method to use. `COGNITO_DEFAULT` for the default email functionality built into Cognito or `DEVELOPER` to use your Amazon SES configuration. Required to be `DEVELOPER` if `from_email_address` is set.
     * 
     */
    public Optional<Output<String>> emailSendingAccount() {
        return Optional.ofNullable(this.emailSendingAccount);
    }

    /**
     * Sender’s email address or sender’s display name with their email address (e.g., `john@example.com`, `John Smith &lt;john@example.com&gt;` or `\&#34;John Smith Ph.D.\&#34; &lt;john@example.com&gt;`). Escaped double quotes are required around display names that contain certain characters as specified in [RFC 5322](https://tools.ietf.org/html/rfc5322).
     * 
     */
    @Import(name="fromEmailAddress")
    private @Nullable Output<String> fromEmailAddress;

    /**
     * @return Sender’s email address or sender’s display name with their email address (e.g., `john@example.com`, `John Smith &lt;john@example.com&gt;` or `\&#34;John Smith Ph.D.\&#34; &lt;john@example.com&gt;`). Escaped double quotes are required around display names that contain certain characters as specified in [RFC 5322](https://tools.ietf.org/html/rfc5322).
     * 
     */
    public Optional<Output<String>> fromEmailAddress() {
        return Optional.ofNullable(this.fromEmailAddress);
    }

    /**
     * REPLY-TO email address.
     * 
     */
    @Import(name="replyToEmailAddress")
    private @Nullable Output<String> replyToEmailAddress;

    /**
     * @return REPLY-TO email address.
     * 
     */
    public Optional<Output<String>> replyToEmailAddress() {
        return Optional.ofNullable(this.replyToEmailAddress);
    }

    /**
     * ARN of the SES verified email identity to use. Required if `email_sending_account` is set to `DEVELOPER`.
     * 
     */
    @Import(name="sourceArn")
    private @Nullable Output<String> sourceArn;

    /**
     * @return ARN of the SES verified email identity to use. Required if `email_sending_account` is set to `DEVELOPER`.
     * 
     */
    public Optional<Output<String>> sourceArn() {
        return Optional.ofNullable(this.sourceArn);
    }

    private UserPoolEmailConfigurationArgs() {}

    private UserPoolEmailConfigurationArgs(UserPoolEmailConfigurationArgs $) {
        this.configurationSet = $.configurationSet;
        this.emailSendingAccount = $.emailSendingAccount;
        this.fromEmailAddress = $.fromEmailAddress;
        this.replyToEmailAddress = $.replyToEmailAddress;
        this.sourceArn = $.sourceArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserPoolEmailConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserPoolEmailConfigurationArgs $;

        public Builder() {
            $ = new UserPoolEmailConfigurationArgs();
        }

        public Builder(UserPoolEmailConfigurationArgs defaults) {
            $ = new UserPoolEmailConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configurationSet Email configuration set name from SES.
         * 
         * @return builder
         * 
         */
        public Builder configurationSet(@Nullable Output<String> configurationSet) {
            $.configurationSet = configurationSet;
            return this;
        }

        /**
         * @param configurationSet Email configuration set name from SES.
         * 
         * @return builder
         * 
         */
        public Builder configurationSet(String configurationSet) {
            return configurationSet(Output.of(configurationSet));
        }

        /**
         * @param emailSendingAccount Email delivery method to use. `COGNITO_DEFAULT` for the default email functionality built into Cognito or `DEVELOPER` to use your Amazon SES configuration. Required to be `DEVELOPER` if `from_email_address` is set.
         * 
         * @return builder
         * 
         */
        public Builder emailSendingAccount(@Nullable Output<String> emailSendingAccount) {
            $.emailSendingAccount = emailSendingAccount;
            return this;
        }

        /**
         * @param emailSendingAccount Email delivery method to use. `COGNITO_DEFAULT` for the default email functionality built into Cognito or `DEVELOPER` to use your Amazon SES configuration. Required to be `DEVELOPER` if `from_email_address` is set.
         * 
         * @return builder
         * 
         */
        public Builder emailSendingAccount(String emailSendingAccount) {
            return emailSendingAccount(Output.of(emailSendingAccount));
        }

        /**
         * @param fromEmailAddress Sender’s email address or sender’s display name with their email address (e.g., `john@example.com`, `John Smith &lt;john@example.com&gt;` or `\&#34;John Smith Ph.D.\&#34; &lt;john@example.com&gt;`). Escaped double quotes are required around display names that contain certain characters as specified in [RFC 5322](https://tools.ietf.org/html/rfc5322).
         * 
         * @return builder
         * 
         */
        public Builder fromEmailAddress(@Nullable Output<String> fromEmailAddress) {
            $.fromEmailAddress = fromEmailAddress;
            return this;
        }

        /**
         * @param fromEmailAddress Sender’s email address or sender’s display name with their email address (e.g., `john@example.com`, `John Smith &lt;john@example.com&gt;` or `\&#34;John Smith Ph.D.\&#34; &lt;john@example.com&gt;`). Escaped double quotes are required around display names that contain certain characters as specified in [RFC 5322](https://tools.ietf.org/html/rfc5322).
         * 
         * @return builder
         * 
         */
        public Builder fromEmailAddress(String fromEmailAddress) {
            return fromEmailAddress(Output.of(fromEmailAddress));
        }

        /**
         * @param replyToEmailAddress REPLY-TO email address.
         * 
         * @return builder
         * 
         */
        public Builder replyToEmailAddress(@Nullable Output<String> replyToEmailAddress) {
            $.replyToEmailAddress = replyToEmailAddress;
            return this;
        }

        /**
         * @param replyToEmailAddress REPLY-TO email address.
         * 
         * @return builder
         * 
         */
        public Builder replyToEmailAddress(String replyToEmailAddress) {
            return replyToEmailAddress(Output.of(replyToEmailAddress));
        }

        /**
         * @param sourceArn ARN of the SES verified email identity to use. Required if `email_sending_account` is set to `DEVELOPER`.
         * 
         * @return builder
         * 
         */
        public Builder sourceArn(@Nullable Output<String> sourceArn) {
            $.sourceArn = sourceArn;
            return this;
        }

        /**
         * @param sourceArn ARN of the SES verified email identity to use. Required if `email_sending_account` is set to `DEVELOPER`.
         * 
         * @return builder
         * 
         */
        public Builder sourceArn(String sourceArn) {
            return sourceArn(Output.of(sourceArn));
        }

        public UserPoolEmailConfigurationArgs build() {
            return $;
        }
    }

}
