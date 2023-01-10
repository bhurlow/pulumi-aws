// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Outputs
{

    [OutputType]
    public sealed class GetCatalogTableStorageDescriptorColumnResult
    {
        /// <summary>
        /// Free-form text comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Name of the table.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Map of initialization parameters for the SerDe, in key-value form.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Parameters;
        /// <summary>
        /// Datatype of data in the Column.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCatalogTableStorageDescriptorColumnResult(
            string comment,

            string name,

            ImmutableDictionary<string, string> parameters,

            string type)
        {
            Comment = comment;
            Name = name;
            Parameters = parameters;
            Type = type;
        }
    }
}
