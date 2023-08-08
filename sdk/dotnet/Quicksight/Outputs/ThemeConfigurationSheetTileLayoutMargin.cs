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
    public sealed class ThemeConfigurationSheetTileLayoutMargin
    {
        /// <summary>
        /// This Boolean value controls whether to display sheet margins.
        /// </summary>
        public readonly bool? Show;

        [OutputConstructor]
        private ThemeConfigurationSheetTileLayoutMargin(bool? show)
        {
            Show = show;
        }
    }
}
