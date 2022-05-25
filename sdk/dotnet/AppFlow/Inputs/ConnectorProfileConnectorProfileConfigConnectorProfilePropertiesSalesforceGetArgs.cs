// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the Salesforce resource.
        /// </summary>
        [Input("instanceUrl")]
        public Input<string>? InstanceUrl { get; set; }

        /// <summary>
        /// Indicates whether the connector profile applies to a sandbox or production environment.
        /// </summary>
        [Input("isSandboxEnvironment")]
        public Input<bool>? IsSandboxEnvironment { get; set; }

        public ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforceGetArgs()
        {
        }
    }
}
