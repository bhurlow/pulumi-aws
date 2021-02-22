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
    public sealed class UserPoolClientAnalyticsConfiguration
    {
        /// <summary>
        /// The application ARN for an Amazon Pinpoint application. Conflicts with `external_id` and `role_arn`.
        /// </summary>
        public readonly string? ApplicationArn;
        /// <summary>
        /// The application ID for an Amazon Pinpoint application.
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// An ID for the Analytics Configuration. Conflicts with `application_arn`.
        /// </summary>
        public readonly string? ExternalId;
        /// <summary>
        /// The ARN of an IAM role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics. Conflicts with `application_arn`.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// If set to `true`, Amazon Cognito will include user data in the events it publishes to Amazon Pinpoint analytics.
        /// </summary>
        public readonly bool? UserDataShared;

        [OutputConstructor]
        private UserPoolClientAnalyticsConfiguration(
            string? applicationArn,

            string? applicationId,

            string? externalId,

            string? roleArn,

            bool? userDataShared)
        {
            ApplicationArn = applicationArn;
            ApplicationId = applicationId;
            ExternalId = externalId;
            RoleArn = roleArn;
            UserDataShared = userDataShared;
        }
    }
}
