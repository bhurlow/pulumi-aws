// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Inputs
{

    public sealed class ThemeConfigurationSheetTileGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The border around a tile. See border.
        /// </summary>
        [Input("border")]
        public Input<Inputs.ThemeConfigurationSheetTileBorderGetArgs>? Border { get; set; }

        public ThemeConfigurationSheetTileGetArgs()
        {
        }
        public static new ThemeConfigurationSheetTileGetArgs Empty => new ThemeConfigurationSheetTileGetArgs();
    }
}
