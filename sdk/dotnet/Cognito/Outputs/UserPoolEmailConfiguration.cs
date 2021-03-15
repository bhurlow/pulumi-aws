// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolEmailConfiguration
    {
        /// <summary>
        /// Email configuration set name from SES.
        /// </summary>
        public readonly string? ConfigurationSet;
        /// <summary>
        /// Email delivery method to use. `COGNITO_DEFAULT` for the default email functionality built into Cognito or `DEVELOPER` to use your Amazon SES configuration.
        /// </summary>
        public readonly string? EmailSendingAccount;
        /// <summary>
        /// Sender’s email address or sender’s display name with their email address (e.g. `john@example.com`, `John Smith &lt;john@example.com&gt;` or `\"John Smith Ph.D.\" &lt;john@example.com&gt;`). Escaped double quotes are required around display names that contain certain characters as specified in [RFC 5322](https://tools.ietf.org/html/rfc5322).
        /// </summary>
        public readonly string? FromEmailAddress;
        /// <summary>
        /// REPLY-TO email address.
        /// </summary>
        public readonly string? ReplyToEmailAddress;
        /// <summary>
        /// ARN of the SES verified email identity to to use. Required if `email_sending_account` is set to `DEVELOPER`.
        /// </summary>
        public readonly string? SourceArn;

        [OutputConstructor]
        private UserPoolEmailConfiguration(
            string? configurationSet,

            string? emailSendingAccount,

            string? fromEmailAddress,

            string? replyToEmailAddress,

            string? sourceArn)
        {
            ConfigurationSet = configurationSet;
            EmailSendingAccount = emailSendingAccount;
            FromEmailAddress = fromEmailAddress;
            ReplyToEmailAddress = replyToEmailAddress;
            SourceArn = sourceArn;
        }
    }
}
