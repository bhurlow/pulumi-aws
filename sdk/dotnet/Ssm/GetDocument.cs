// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm
{
    public static class GetDocument
    {
        /// <summary>
        /// Gets the contents of the specified Systems Manager document.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// To get the contents of the document owned by AWS.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Aws.Ssm.GetDocument.InvokeAsync(new Aws.Ssm.GetDocumentArgs
        ///         {
        ///             DocumentFormat = "YAML",
        ///             Name = "AWS-GatherSoftwareInventory",
        ///         }));
        ///         this.Content = foo.Apply(foo =&gt; foo.Content);
        ///     }
        /// 
        ///     [Output("content")]
        ///     public Output&lt;string&gt; Content { get; set; }
        /// }
        /// ```
        /// 
        /// To get the contents of the custom document.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Aws.Ssm.GetDocument.InvokeAsync(new Aws.Ssm.GetDocumentArgs
        ///         {
        ///             DocumentFormat = "JSON",
        ///             Name = aws_ssm_document.Test.Name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDocumentResult> InvokeAsync(GetDocumentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDocumentResult>("aws:ssm/getDocument:getDocument", args ?? new GetDocumentArgs(), options.WithVersion());
    }


    public sealed class GetDocumentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns the document in the specified format. The document format can be either JSON or YAML. JSON is the default format.
        /// </summary>
        [Input("documentFormat")]
        public string? DocumentFormat { get; set; }

        /// <summary>
        /// The document version for which you want information.
        /// </summary>
        [Input("documentVersion")]
        public string? DocumentVersion { get; set; }

        /// <summary>
        /// The name of the Systems Manager document.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDocumentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDocumentResult
    {
        /// <summary>
        /// The ARN of the document.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The contents of the document.
        /// </summary>
        public readonly string Content;
        public readonly string? DocumentFormat;
        /// <summary>
        /// The type of the document.
        /// </summary>
        public readonly string DocumentType;
        public readonly string? DocumentVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetDocumentResult(
            string arn,

            string content,

            string? documentFormat,

            string documentType,

            string? documentVersion,

            string id,

            string name)
        {
            Arn = arn;
            Content = content;
            DocumentFormat = documentFormat;
            DocumentType = documentType;
            DocumentVersion = documentVersion;
            Id = id;
            Name = name;
        }
    }
}
