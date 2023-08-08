// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpenSearch.Inputs
{

    public sealed class ServerlessSecurityConfigSamlOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group attribute for this SAML integration.
        /// </summary>
        [Input("groupAttribute")]
        public Input<string>? GroupAttribute { get; set; }

        /// <summary>
        /// The XML IdP metadata file generated from your identity provider.
        /// </summary>
        [Input("metadata", required: true)]
        public Input<string> Metadata { get; set; } = null!;

        /// <summary>
        /// Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.
        /// </summary>
        [Input("sessionTimeout")]
        public Input<int>? SessionTimeout { get; set; }

        /// <summary>
        /// User attribute for this SAML integration.
        /// </summary>
        [Input("userAttribute")]
        public Input<string>? UserAttribute { get; set; }

        public ServerlessSecurityConfigSamlOptionsGetArgs()
        {
        }
        public static new ServerlessSecurityConfigSamlOptionsGetArgs Empty => new ServerlessSecurityConfigSamlOptionsGetArgs();
    }
}
