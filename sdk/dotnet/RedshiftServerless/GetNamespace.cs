// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedshiftServerless
{
    public static class GetNamespace
    {
        /// <summary>
        /// Data source for managing an AWS Redshift Serverless Namespace.
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
        ///     var example = Aws.RedshiftServerless.GetNamespace.Invoke(new()
        ///     {
        ///         NamespaceName = "example-namespace",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNamespaceResult> InvokeAsync(GetNamespaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceResult>("aws:redshiftserverless/getNamespace:getNamespace", args ?? new GetNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for managing an AWS Redshift Serverless Namespace.
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
        ///     var example = Aws.RedshiftServerless.GetNamespace.Invoke(new()
        ///     {
        ///         NamespaceName = "example-namespace",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNamespaceResult> Invoke(GetNamespaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNamespaceResult>("aws:redshiftserverless/getNamespace:getNamespace", args ?? new GetNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNamespaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        public GetNamespaceArgs()
        {
        }
        public static new GetNamespaceArgs Empty => new GetNamespaceArgs();
    }

    public sealed class GetNamespaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        public GetNamespaceInvokeArgs()
        {
        }
        public static new GetNamespaceInvokeArgs Empty => new GetNamespaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetNamespaceResult
    {
        /// <summary>
        /// The username of the administrator for the first database created in the namespace.
        /// </summary>
        public readonly string AdminUsername;
        /// <summary>
        /// Amazon Resource Name (ARN) of the Redshift Serverless Namespace.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The name of the first database created in the namespace.
        /// </summary>
        public readonly string DbName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying `default_iam_role_arn`, it also must be part of `iam_roles`.
        /// </summary>
        public readonly string DefaultIamRoleArn;
        /// <summary>
        /// A list of IAM roles to associate with the namespace.
        /// </summary>
        public readonly ImmutableArray<string> IamRoles;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// The types of logs the namespace can export. Available export types are `userlog`, `connectionlog`, and `useractivitylog`.
        /// </summary>
        public readonly ImmutableArray<string> LogExports;
        /// <summary>
        /// The Redshift Namespace ID.
        /// </summary>
        public readonly string NamespaceId;
        public readonly string NamespaceName;

        [OutputConstructor]
        private GetNamespaceResult(
            string adminUsername,

            string arn,

            string dbName,

            string defaultIamRoleArn,

            ImmutableArray<string> iamRoles,

            string id,

            string kmsKeyId,

            ImmutableArray<string> logExports,

            string namespaceId,

            string namespaceName)
        {
            AdminUsername = adminUsername;
            Arn = arn;
            DbName = dbName;
            DefaultIamRoleArn = defaultIamRoleArn;
            IamRoles = iamRoles;
            Id = id;
            KmsKeyId = kmsKeyId;
            LogExports = logExports;
            NamespaceId = namespaceId;
            NamespaceName = namespaceName;
        }
    }
}
