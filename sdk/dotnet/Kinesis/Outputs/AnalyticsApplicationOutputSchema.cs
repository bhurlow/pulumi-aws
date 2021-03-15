// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Outputs
{

    [OutputType]
    public sealed class AnalyticsApplicationOutputSchema
    {
        /// <summary>
        /// The Format Type of the records on the output stream. Can be `CSV` or `JSON`.
        /// </summary>
        public readonly string RecordFormatType;

        [OutputConstructor]
        private AnalyticsApplicationOutputSchema(string recordFormatType)
        {
            RecordFormatType = recordFormatType;
        }
    }
}
