// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// The Secret Access Key portion of the credentials.
        /// </summary>
        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        public ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeGetArgs()
        {
        }
    }
}
