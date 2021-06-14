// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class UserPoolLambdaConfigCustomEmailSenderGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// he Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send SMS notifications to users.
        /// </summary>
        [Input("lambdaArn", required: true)]
        public Input<string> LambdaArn { get; set; } = null!;

        /// <summary>
        /// The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom SMS Lambda function. The only supported value is `V1_0`.
        /// </summary>
        [Input("lambdaVersion", required: true)]
        public Input<string> LambdaVersion { get; set; } = null!;

        public UserPoolLambdaConfigCustomEmailSenderGetArgs()
        {
        }
    }
}
