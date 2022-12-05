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
    public sealed class ClassifierCsvClassifier
    {
        /// <summary>
        /// Enables the processing of files that contain only one column.
        /// </summary>
        public readonly bool? AllowSingleColumn;
        /// <summary>
        /// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
        /// </summary>
        public readonly string? ContainsHeader;
        /// <summary>
        /// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
        /// </summary>
        public readonly bool? CustomDatatypeConfigured;
        /// <summary>
        /// A list of supported custom datatypes. Valid values are `BINARY`, `BOOLEAN`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `LONG`, `SHORT`, `STRING`, `TIMESTAMP`.
        /// </summary>
        public readonly ImmutableArray<string> CustomDatatypes;
        /// <summary>
        /// The delimiter used in the Csv to separate columns.
        /// </summary>
        public readonly string? Delimiter;
        /// <summary>
        /// Specifies whether to trim column values.
        /// </summary>
        public readonly bool? DisableValueTrimming;
        /// <summary>
        /// A list of strings representing column names.
        /// </summary>
        public readonly ImmutableArray<string> Headers;
        /// <summary>
        /// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
        /// </summary>
        public readonly string? QuoteSymbol;

        [OutputConstructor]
        private ClassifierCsvClassifier(
            bool? allowSingleColumn,

            string? containsHeader,

            bool? customDatatypeConfigured,

            ImmutableArray<string> customDatatypes,

            string? delimiter,

            bool? disableValueTrimming,

            ImmutableArray<string> headers,

            string? quoteSymbol)
        {
            AllowSingleColumn = allowSingleColumn;
            ContainsHeader = containsHeader;
            CustomDatatypeConfigured = customDatatypeConfigured;
            CustomDatatypes = customDatatypes;
            Delimiter = delimiter;
            DisableValueTrimming = disableValueTrimming;
            Headers = headers;
            QuoteSymbol = quoteSymbol;
        }
    }
}
