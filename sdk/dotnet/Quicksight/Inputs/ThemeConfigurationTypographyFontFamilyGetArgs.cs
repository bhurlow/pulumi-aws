// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Inputs
{

    public sealed class ThemeConfigurationTypographyFontFamilyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Font family name.
        /// </summary>
        [Input("fontFamily")]
        public Input<string>? FontFamily { get; set; }

        public ThemeConfigurationTypographyFontFamilyGetArgs()
        {
        }
        public static new ThemeConfigurationTypographyFontFamilyGetArgs Empty => new ThemeConfigurationTypographyFontFamilyGetArgs();
    }
}
