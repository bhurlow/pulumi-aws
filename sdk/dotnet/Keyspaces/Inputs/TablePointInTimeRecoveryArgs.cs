// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Keyspaces.Inputs
{

    public sealed class TablePointInTimeRecoveryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid values: `ENABLED`, `DISABLED`. The default value is `DISABLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public TablePointInTimeRecoveryArgs()
        {
        }
    }
}
