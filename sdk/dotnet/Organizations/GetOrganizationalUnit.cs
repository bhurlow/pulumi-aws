// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Organizations
{
    public static class GetOrganizationalUnit
    {
        /// <summary>
        /// Data source for getting an AWS Organizations Organizational Unit.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var org = Aws.Organizations.GetOrganization.Invoke();
        /// 
        ///     var ou = Aws.Organizations.GetOrganizationalUnit.Invoke(new()
        ///     {
        ///         ParentId = org.Apply(getOrganizationResult =&gt; getOrganizationResult.Roots[0]?.Id),
        ///         Name = "dev",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrganizationalUnitResult> InvokeAsync(GetOrganizationalUnitArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationalUnitResult>("aws:organizations/getOrganizationalUnit:getOrganizationalUnit", args ?? new GetOrganizationalUnitArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for getting an AWS Organizations Organizational Unit.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var org = Aws.Organizations.GetOrganization.Invoke();
        /// 
        ///     var ou = Aws.Organizations.GetOrganizationalUnit.Invoke(new()
        ///     {
        ///         ParentId = org.Apply(getOrganizationResult =&gt; getOrganizationResult.Roots[0]?.Id),
        ///         Name = "dev",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrganizationalUnitResult> Invoke(GetOrganizationalUnitInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationalUnitResult>("aws:organizations/getOrganizationalUnit:getOrganizationalUnit", args ?? new GetOrganizationalUnitInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationalUnitArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the organizational unit
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Parent ID of the organizational unit.
        /// </summary>
        [Input("parentId", required: true)]
        public string ParentId { get; set; } = null!;

        public GetOrganizationalUnitArgs()
        {
        }
        public static new GetOrganizationalUnitArgs Empty => new GetOrganizationalUnitArgs();
    }

    public sealed class GetOrganizationalUnitInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the organizational unit
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Parent ID of the organizational unit.
        /// </summary>
        [Input("parentId", required: true)]
        public Input<string> ParentId { get; set; } = null!;

        public GetOrganizationalUnitInvokeArgs()
        {
        }
        public static new GetOrganizationalUnitInvokeArgs Empty => new GetOrganizationalUnitInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationalUnitResult
    {
        /// <summary>
        /// ARN of the organizational unit
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ParentId;

        [OutputConstructor]
        private GetOrganizationalUnitResult(
            string arn,

            string id,

            string name,

            string parentId)
        {
            Arn = arn;
            Id = id;
            Name = name;
            ParentId = parentId;
        }
    }
}
