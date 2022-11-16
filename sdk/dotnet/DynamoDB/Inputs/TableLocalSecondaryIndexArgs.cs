// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Inputs
{

    public sealed class TableLocalSecondaryIndexArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the index
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("nonKeyAttributes")]
        private InputList<string>? _nonKeyAttributes;

        /// <summary>
        /// Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
        /// </summary>
        public InputList<string> NonKeyAttributes
        {
            get => _nonKeyAttributes ?? (_nonKeyAttributes = new InputList<string>());
            set => _nonKeyAttributes = value;
        }

        /// <summary>
        /// One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
        /// </summary>
        [Input("projectionType", required: true)]
        public Input<string> ProjectionType { get; set; } = null!;

        /// <summary>
        /// Name of the range key.
        /// </summary>
        [Input("rangeKey", required: true)]
        public Input<string> RangeKey { get; set; } = null!;

        public TableLocalSecondaryIndexArgs()
        {
        }
        public static new TableLocalSecondaryIndexArgs Empty => new TableLocalSecondaryIndexArgs();
    }
}
