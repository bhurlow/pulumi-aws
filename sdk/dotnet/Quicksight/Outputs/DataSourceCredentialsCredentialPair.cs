// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Outputs
{

    [OutputType]
    public sealed class DataSourceCredentialsCredentialPair
    {
        /// <summary>
        /// Password, maximum length of 1024 characters.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// User name, maximum length of 64 characters.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private DataSourceCredentialsCredentialPair(
            string password,

            string username)
        {
            Password = password;
            Username = username;
        }
    }
}