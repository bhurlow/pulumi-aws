// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsoAdmin.Inputs
{

    public sealed class TrustedTokenIssuerTrustedTokenIssuerConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A block that describes the settings for a trusted token issuer that works with OpenID Connect (OIDC) by using JSON Web Tokens (JWT). See Documented below below.
        /// </summary>
        [Input("oidcJwtConfiguration")]
        public Input<Inputs.TrustedTokenIssuerTrustedTokenIssuerConfigurationOidcJwtConfigurationGetArgs>? OidcJwtConfiguration { get; set; }

        public TrustedTokenIssuerTrustedTokenIssuerConfigurationGetArgs()
        {
        }
        public static new TrustedTokenIssuerTrustedTokenIssuerConfigurationGetArgs Empty => new TrustedTokenIssuerTrustedTokenIssuerConfigurationGetArgs();
    }
}
