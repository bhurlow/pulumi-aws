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
    public sealed class TableGlobalSecondaryIndex
    {
        /// <summary>
        /// Name of the hash key in the index; must be defined as an attribute in the resource.
        /// </summary>
        public readonly string HashKey;
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
        public readonly string? RangeKey;
        /// <summary>
        /// Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        /// </summary>
        public readonly int? ReadCapacity;
        /// <summary>
        /// Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
        /// </summary>
        public readonly int? WriteCapacity;

        [OutputConstructor]
        private TableGlobalSecondaryIndex(
            string hashKey,

            string name,

            ImmutableArray<string> nonKeyAttributes,

            string projectionType,

            string? rangeKey,

            int? readCapacity,

            int? writeCapacity)
        {
            HashKey = hashKey;
            Name = name;
            NonKeyAttributes = nonKeyAttributes;
            ProjectionType = projectionType;
            RangeKey = rangeKey;
            ReadCapacity = readCapacity;
            WriteCapacity = writeCapacity;
        }
    }
}
