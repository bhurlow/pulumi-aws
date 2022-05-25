// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The code provided by the connector when it has been authenticated via the connected app.
        /// </summary>
        [Input("authCode")]
        public Input<string>? AuthCode { get; set; }

        /// <summary>
        /// The URL to which the authentication server redirects the browser after authorization has been granted.
        /// </summary>
        [Input("redirectUri")]
        public Input<string>? RedirectUri { get; set; }

        public ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestGetArgs()
        {
        }
    }
}
