// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class DomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSetting
    {
        /// <summary>
        /// The name of the data source that you're connecting to. Canvas currently supports OAuth for Snowflake and Salesforce Data Cloud. Valid values are `SalesforceGenie` and `Snowflake`.
        /// </summary>
        public readonly string? DataSourceName;
        /// <summary>
        /// The ARN of an Amazon Web Services Secrets Manager secret that stores the credentials from your identity provider, such as the client ID and secret, authorization URL, and token URL.
        /// </summary>
        public readonly string SecretArn;
        /// <summary>
        /// Describes whether OAuth for a data source is enabled or disabled in the Canvas application. Valid values are `ENABLED` and `DISABLED`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private DomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSetting(
            string? dataSourceName,

            string secretArn,

            string? status)
        {
            DataSourceName = dataSourceName;
            SecretArn = secretArn;
            Status = status;
        }
    }
}
