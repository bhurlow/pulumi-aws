// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Qldb
{
    public static class GetLedger
    {
        /// <summary>
        /// Use this data source to fetch information about a Quantum Ledger Database.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Qldb.GetLedger.InvokeAsync(new Aws.Qldb.GetLedgerArgs
        ///         {
        ///             Name = "an_example_ledger",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLedgerResult> InvokeAsync(GetLedgerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLedgerResult>("aws:qldb/getLedger:getLedger", args ?? new GetLedgerArgs(), options.WithVersion());
    }


    public sealed class GetLedgerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly name of the ledger to match.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetLedgerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLedgerResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the ledger.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Deletion protection on the QLDB Ledger instance. Set to `true` by default.
        /// </summary>
        public readonly bool DeletionProtection;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetLedgerResult(
            string arn,

            bool deletionProtection,

            string id,

            string name)
        {
            Arn = arn;
            DeletionProtection = deletionProtection;
            Id = id;
            Name = name;
        }
    }
}
