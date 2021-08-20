// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class WorkforceOidcConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OIDC IdP authorization endpoint used to configure your private workforce.
        /// </summary>
        [Input("authorizationEndpoint", required: true)]
        public Input<string> AuthorizationEndpoint { get; set; } = null!;

        /// <summary>
        /// The OIDC IdP client ID used to configure your private workforce.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The OIDC IdP client secret used to configure your private workforce.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The OIDC IdP issuer used to configure your private workforce.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.
        /// </summary>
        [Input("jwksUri", required: true)]
        public Input<string> JwksUri { get; set; } = null!;

        /// <summary>
        /// The OIDC IdP logout endpoint used to configure your private workforce.
        /// </summary>
        [Input("logoutEndpoint", required: true)]
        public Input<string> LogoutEndpoint { get; set; } = null!;

        /// <summary>
        /// The OIDC IdP token endpoint used to configure your private workforce.
        /// </summary>
        [Input("tokenEndpoint", required: true)]
        public Input<string> TokenEndpoint { get; set; } = null!;

        /// <summary>
        /// The OIDC IdP user information endpoint used to configure your private workforce.
        /// </summary>
        [Input("userInfoEndpoint", required: true)]
        public Input<string> UserInfoEndpoint { get; set; } = null!;

        public WorkforceOidcConfigArgs()
        {
        }
    }
}