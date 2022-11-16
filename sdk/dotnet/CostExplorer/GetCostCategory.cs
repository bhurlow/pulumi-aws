// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CostExplorer
{
    public static class GetCostCategory
    {
        /// <summary>
        /// Provides details about a specific CostExplorer Cost Category.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.CostExplorer.GetCostCategory.Invoke(new()
        ///     {
        ///         CostCategoryArn = "costCategoryARN",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCostCategoryResult> InvokeAsync(GetCostCategoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCostCategoryResult>("aws:costexplorer/getCostCategory:getCostCategory", args ?? new GetCostCategoryArgs(), options.WithDefaults());

        /// <summary>
        /// Provides details about a specific CostExplorer Cost Category.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.CostExplorer.GetCostCategory.Invoke(new()
        ///     {
        ///         CostCategoryArn = "costCategoryARN",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCostCategoryResult> Invoke(GetCostCategoryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCostCategoryResult>("aws:costexplorer/getCostCategory:getCostCategory", args ?? new GetCostCategoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCostCategoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique name for the Cost Category.
        /// </summary>
        [Input("costCategoryArn", required: true)]
        public string CostCategoryArn { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Configuration block for the specific `Tag` to use for `Expression`. See below.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetCostCategoryArgs()
        {
        }
        public static new GetCostCategoryArgs Empty => new GetCostCategoryArgs();
    }

    public sealed class GetCostCategoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique name for the Cost Category.
        /// </summary>
        [Input("costCategoryArn", required: true)]
        public Input<string> CostCategoryArn { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Configuration block for the specific `Tag` to use for `Expression`. See below.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetCostCategoryInvokeArgs()
        {
        }
        public static new GetCostCategoryInvokeArgs Empty => new GetCostCategoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetCostCategoryResult
    {
        public readonly string CostCategoryArn;
        /// <summary>
        /// Effective end data of your Cost Category.
        /// </summary>
        public readonly string EffectiveEnd;
        /// <summary>
        /// Effective state data of your Cost Category.
        /// </summary>
        public readonly string EffectiveStart;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// Rule schema version in this particular Cost Category.
        /// </summary>
        public readonly string RuleVersion;
        /// <summary>
        /// Configuration block for the `Expression` object used to categorize costs. See below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCostCategoryRuleResult> Rules;
        /// <summary>
        /// Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCostCategorySplitChargeRuleResult> SplitChargeRules;
        /// <summary>
        /// Configuration block for the specific `Tag` to use for `Expression`. See below.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetCostCategoryResult(
            string costCategoryArn,

            string effectiveEnd,

            string effectiveStart,

            string id,

            string name,

            string ruleVersion,

            ImmutableArray<Outputs.GetCostCategoryRuleResult> rules,

            ImmutableArray<Outputs.GetCostCategorySplitChargeRuleResult> splitChargeRules,

            ImmutableDictionary<string, string> tags)
        {
            CostCategoryArn = costCategoryArn;
            EffectiveEnd = effectiveEnd;
            EffectiveStart = effectiveStart;
            Id = id;
            Name = name;
            RuleVersion = ruleVersion;
            Rules = rules;
            SplitChargeRules = splitChargeRules;
            Tags = tags;
        }
    }
}
