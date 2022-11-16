// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Outputs
{

    [OutputType]
    public sealed class TableLocalSecondaryIndex
    {
        /// <summary>
        /// Name of the index
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        /// </summary>
        public readonly ImmutableArray<string> NonKeyAttributes;
        /// <summary>
        /// One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        /// </summary>
        public readonly string ProjectionType;
        /// <summary>
        /// Name of the range key.
        /// </summary>
        public readonly string RangeKey;

        [OutputConstructor]
        private TableLocalSecondaryIndex(
            string name,

            ImmutableArray<string> nonKeyAttributes,

            string projectionType,

            string rangeKey)
        {
            Name = name;
            NonKeyAttributes = nonKeyAttributes;
            ProjectionType = projectionType;
            RangeKey = rangeKey;
        }
    }
}
