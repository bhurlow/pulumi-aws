// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorApiKey
    {
        /// <summary>
        /// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// The Secret Access Key portion of the credentials.
        /// </summary>
        public readonly string? ApiSecretKey;

        [OutputConstructor]
        private ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorApiKey(
            string apiKey,

            string? apiSecretKey)
        {
            ApiKey = apiKey;
            ApiSecretKey = apiSecretKey;
        }
    }
}
