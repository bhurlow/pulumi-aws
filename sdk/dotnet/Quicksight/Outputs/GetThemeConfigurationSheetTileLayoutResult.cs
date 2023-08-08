// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Outputs
{

    [OutputType]
    public sealed class GetThemeConfigurationSheetTileLayoutResult
    {
        /// <summary>
        /// The gutter settings that apply between tiles. See gutter.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetThemeConfigurationSheetTileLayoutGutterResult> Gutters;
        /// <summary>
        /// The margin settings that apply around the outside edge of sheets. See margin.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetThemeConfigurationSheetTileLayoutMarginResult> Margins;

        [OutputConstructor]
        private GetThemeConfigurationSheetTileLayoutResult(
            ImmutableArray<Outputs.GetThemeConfigurationSheetTileLayoutGutterResult> gutters,

            ImmutableArray<Outputs.GetThemeConfigurationSheetTileLayoutMarginResult> margins)
        {
            Gutters = gutters;
            Margins = margins;
        }
    }
}
