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
    public sealed class UserPoolAdminCreateUserConfigInviteMessageTemplate
    {
        /// <summary>
        /// Email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
        /// </summary>
        public readonly string? EmailMessage;
        /// <summary>
        /// Subject line for the email message template. Conflicts with `email_verification_subject` argument.
        /// </summary>
        public readonly string? EmailSubject;
        /// <summary>
        /// SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.
        /// </summary>
        public readonly string? SmsMessage;

        [OutputConstructor]
        private UserPoolAdminCreateUserConfigInviteMessageTemplate(
            string? emailMessage,

            string? emailSubject,

            string? smsMessage)
        {
            EmailMessage = emailMessage;
            EmailSubject = emailSubject;
            SmsMessage = smsMessage;
        }
    }
}
