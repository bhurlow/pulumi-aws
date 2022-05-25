// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// Application keys, in conjunction with your API key, give you full access to Datadog’s programmatic API. Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.
        /// </summary>
        [Input("applicationKey", required: true)]
        public Input<string> ApplicationKey { get; set; } = null!;

        public ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogGetArgs()
        {
        }
    }
}
