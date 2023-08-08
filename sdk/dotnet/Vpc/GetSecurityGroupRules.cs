// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Vpc
{
    public static class GetSecurityGroupRules
    {
        /// <summary>
        /// This resource can be useful for getting back a set of security group rule IDs.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.Vpc.GetSecurityGroupRules.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Aws.Vpc.Inputs.GetSecurityGroupRulesFilterInputArgs
        ///             {
        ///                 Name = "group-id",
        ///                 Values = new[]
        ///                 {
        ///                     @var.Security_group_id,
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecurityGroupRulesResult> InvokeAsync(GetSecurityGroupRulesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupRulesResult>("aws:vpc/getSecurityGroupRules:getSecurityGroupRules", args ?? new GetSecurityGroupRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This resource can be useful for getting back a set of security group rule IDs.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.Vpc.GetSecurityGroupRules.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Aws.Vpc.Inputs.GetSecurityGroupRulesFilterInputArgs
        ///             {
        ///                 Name = "group-id",
        ///                 Values = new[]
        ///                 {
        ///                     @var.Security_group_id,
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSecurityGroupRulesResult> Invoke(GetSecurityGroupRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupRulesResult>("aws:vpc/getSecurityGroupRules:getSecurityGroupRules", args ?? new GetSecurityGroupRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupRulesArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSecurityGroupRulesFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetSecurityGroupRulesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSecurityGroupRulesFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of tags, each pair of which must exactly match
        /// a pair on the desired security group rule.
        /// 
        /// More complex filters can be expressed using one or more `filter` sub-blocks,
        /// which take the following arguments:
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetSecurityGroupRulesArgs()
        {
        }
        public static new GetSecurityGroupRulesArgs Empty => new GetSecurityGroupRulesArgs();
    }

    public sealed class GetSecurityGroupRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetSecurityGroupRulesFilterInputArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public InputList<Inputs.GetSecurityGroupRulesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetSecurityGroupRulesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags, each pair of which must exactly match
        /// a pair on the desired security group rule.
        /// 
        /// More complex filters can be expressed using one or more `filter` sub-blocks,
        /// which take the following arguments:
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetSecurityGroupRulesInvokeArgs()
        {
        }
        public static new GetSecurityGroupRulesInvokeArgs Empty => new GetSecurityGroupRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityGroupRulesResult
    {
        public readonly ImmutableArray<Outputs.GetSecurityGroupRulesFilterResult> Filters;
        public readonly string Id;
        /// <summary>
        /// List of all the security group rule IDs found.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetSecurityGroupRulesResult(
            ImmutableArray<Outputs.GetSecurityGroupRulesFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableDictionary<string, string>? tags)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Tags = tags;
        }
    }
}
