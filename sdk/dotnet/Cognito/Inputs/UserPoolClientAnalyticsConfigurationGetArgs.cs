// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class UserPoolClientAnalyticsConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application ARN for an Amazon Pinpoint application. Conflicts with `external_id` and `role_arn`.
        /// </summary>
        [Input("applicationArn")]
        public Input<string>? ApplicationArn { get; set; }

        /// <summary>
        /// The application ID for an Amazon Pinpoint application.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// An ID for the Analytics Configuration. Conflicts with `application_arn`.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The ARN of an IAM role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics. Conflicts with `application_arn`.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// If set to `true`, Amazon Cognito will include user data in the events it publishes to Amazon Pinpoint analytics.
        /// </summary>
        [Input("userDataShared")]
        public Input<bool>? UserDataShared { get; set; }

        public UserPoolClientAnalyticsConfigurationGetArgs()
        {
        }
    }
}
