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
    public sealed class ConnectorProfileConnectorProfileConfigConnectorProfileCredentials
    {
        /// <summary>
        /// The connector-specific credentials required when using Amplitude. See Amplitude Connector Profile Credentials for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude? Amplitude;
        /// <summary>
        /// The connector-specific profile properties required when using the custom connector. See Custom Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector? CustomConnector;
        /// <summary>
        /// The connector-specific properties required when using Datadog. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog? Datadog;
        /// <summary>
        /// The connector-specific properties required when using Dynatrace. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace? Dynatrace;
        /// <summary>
        /// The connector-specific credentials required when using Google Analytics. See Google Analytics Connector Profile Credentials for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics? GoogleAnalytics;
        /// <summary>
        /// The connector-specific credentials required when using Amazon Honeycode. See Honeycode Connector Profile Credentials for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode? Honeycode;
        /// <summary>
        /// The connector-specific properties required when using Infor Nexus. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus? InforNexus;
        /// <summary>
        /// The connector-specific properties required when using Marketo. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo? Marketo;
        /// <summary>
        /// The connector-specific properties required when using Amazon Redshift. See Redshift Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift? Redshift;
        /// <summary>
        /// The connector-specific properties required when using Salesforce. See Salesforce Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce? Salesforce;
        /// <summary>
        /// The connector-specific properties required when using SAPOData. See SAPOData Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData? SapoData;
        /// <summary>
        /// The connector-specific properties required when using ServiceNow. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow? ServiceNow;
        /// <summary>
        /// The connector-specific credentials required when using Singular. See Singular Connector Profile Credentials for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular? Singular;
        /// <summary>
        /// The connector-specific properties required when using Slack. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack? Slack;
        /// <summary>
        /// The connector-specific properties required when using Snowflake. See Snowflake Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake? Snowflake;
        /// <summary>
        /// The connector-specific credentials required when using Trend Micro. See Trend Micro Connector Profile Credentials for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro? Trendmicro;
        /// <summary>
        /// The connector-specific properties required when using Veeva. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva? Veeva;
        /// <summary>
        /// The connector-specific properties required when using Zendesk. See Generic Connector Profile Properties for more details.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk? Zendesk;

        [OutputConstructor]
        private ConnectorProfileConnectorProfileConfigConnectorProfileCredentials(
            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude? amplitude,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector? customConnector,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog? datadog,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace? dynatrace,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics? googleAnalytics,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode? honeycode,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus? inforNexus,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo? marketo,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift? redshift,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce? salesforce,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData? sapoData,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow? serviceNow,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular? singular,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack? slack,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake? snowflake,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro? trendmicro,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva? veeva,

            Outputs.ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk? zendesk)
        {
            Amplitude = amplitude;
            CustomConnector = customConnector;
            Datadog = datadog;
            Dynatrace = dynatrace;
            GoogleAnalytics = googleAnalytics;
            Honeycode = honeycode;
            InforNexus = inforNexus;
            Marketo = marketo;
            Redshift = redshift;
            Salesforce = salesforce;
            SapoData = sapoData;
            ServiceNow = serviceNow;
            Singular = singular;
            Slack = slack;
            Snowflake = snowflake;
            Trendmicro = trendmicro;
            Veeva = veeva;
            Zendesk = zendesk;
        }
    }
}
