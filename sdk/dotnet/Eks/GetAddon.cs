// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Eks
{
    public static class GetAddon
    {
        /// <summary>
        /// Retrieve information about an EKS add-on.
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
        ///     var example = Aws.Eks.GetAddon.Invoke(new()
        ///     {
        ///         AddonName = "vpc-cni",
        ///         ClusterName = aws_eks_cluster.Example.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["eksAddonOutputs"] = aws_eks_addon.Example,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAddonResult> InvokeAsync(GetAddonArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAddonResult>("aws:eks/getAddon:getAddon", args ?? new GetAddonArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve information about an EKS add-on.
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
        ///     var example = Aws.Eks.GetAddon.Invoke(new()
        ///     {
        ///         AddonName = "vpc-cni",
        ///         ClusterName = aws_eks_cluster.Example.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["eksAddonOutputs"] = aws_eks_addon.Example,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAddonResult> Invoke(GetAddonInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddonResult>("aws:eks/getAddon:getAddon", args ?? new GetAddonInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddonArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the EKS add-on. The name must match one of
        /// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
        /// </summary>
        [Input("addonName", required: true)]
        public string AddonName { get; set; } = null!;

        /// <summary>
        /// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetAddonArgs()
        {
        }
        public static new GetAddonArgs Empty => new GetAddonArgs();
    }

    public sealed class GetAddonInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the EKS add-on. The name must match one of
        /// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
        /// </summary>
        [Input("addonName", required: true)]
        public Input<string> AddonName { get; set; } = null!;

        /// <summary>
        /// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetAddonInvokeArgs()
        {
        }
        public static new GetAddonInvokeArgs Empty => new GetAddonInvokeArgs();
    }


    [OutputType]
    public sealed class GetAddonResult
    {
        public readonly string AddonName;
        /// <summary>
        /// Version of EKS add-on.
        /// </summary>
        public readonly string AddonVersion;
        /// <summary>
        /// ARN of the EKS add-on.
        /// </summary>
        public readonly string Arn;
        public readonly string ClusterName;
        /// <summary>
        /// Configuration values for the addon with a single JSON string.
        /// </summary>
        public readonly string ConfigurationValues;
        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was updated.
        /// </summary>
        public readonly string ModifiedAt;
        /// <summary>
        /// ARN of IAM role used for EKS add-on. If value is empty -
        /// then add-on uses the IAM role assigned to the EKS Cluster node.
        /// </summary>
        public readonly string ServiceAccountRoleArn;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetAddonResult(
            string addonName,

            string addonVersion,

            string arn,

            string clusterName,

            string configurationValues,

            string createdAt,

            string id,

            string modifiedAt,

            string serviceAccountRoleArn,

            ImmutableDictionary<string, string> tags)
        {
            AddonName = addonName;
            AddonVersion = addonVersion;
            Arn = arn;
            ClusterName = clusterName;
            ConfigurationValues = configurationValues;
            CreatedAt = createdAt;
            Id = id;
            ModifiedAt = modifiedAt;
            ServiceAccountRoleArn = serviceAccountRoleArn;
            Tags = tags;
        }
    }
}
