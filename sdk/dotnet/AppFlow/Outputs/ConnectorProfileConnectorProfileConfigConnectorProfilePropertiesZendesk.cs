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
    public sealed class ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk
    {
        /// <summary>
        /// The location of the Salesforce resource.
        /// </summary>
        public readonly string InstanceUrl;

        [OutputConstructor]
        private ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk(string instanceUrl)
        {
            InstanceUrl = instanceUrl;
        }
    }
}
