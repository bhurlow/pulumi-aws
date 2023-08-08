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
    public sealed class ThemeConfigurationSheetTile
    {
        /// <summary>
        /// The border around a tile. See border.
        /// </summary>
        public readonly Outputs.ThemeConfigurationSheetTileBorder? Border;

        [OutputConstructor]
        private ThemeConfigurationSheetTile(Outputs.ThemeConfigurationSheetTileBorder? border)
        {
            Border = border;
        }
    }
}
