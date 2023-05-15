// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.EcrPublic
{
    /// <summary>
    /// Provides an Elastic Container Registry Public Repository Policy.
    /// 
    /// Note that currently only one policy may be applied to a repository.
    /// 
    /// &gt; **NOTE:** This resource can only be used in the `us-east-1` region.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleRepository = new Aws.EcrPublic.Repository("exampleRepository", new()
    ///     {
    ///         RepositoryName = "example",
    ///     });
    /// 
    ///     var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Sid = "new policy",
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "123456789012",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "ecr:GetDownloadUrlForLayer",
    ///                     "ecr:BatchGetImage",
    ///                     "ecr:BatchCheckLayerAvailability",
    ///                     "ecr:PutImage",
    ///                     "ecr:InitiateLayerUpload",
    ///                     "ecr:UploadLayerPart",
    ///                     "ecr:CompleteLayerUpload",
    ///                     "ecr:DescribeRepositories",
    ///                     "ecr:GetRepositoryPolicy",
    ///                     "ecr:ListImages",
    ///                     "ecr:DeleteRepository",
    ///                     "ecr:BatchDeleteImage",
    ///                     "ecr:SetRepositoryPolicy",
    ///                     "ecr:DeleteRepositoryPolicy",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleRepositoryPolicy = new Aws.EcrPublic.RepositoryPolicy("exampleRepositoryPolicy", new()
    ///     {
    ///         RepositoryName = exampleRepository.RepositoryName,
    ///         Policy = examplePolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECR Public Repository Policy can be imported using the repository name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecrpublic/repositoryPolicy:RepositoryPolicy example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecrpublic/repositoryPolicy:RepositoryPolicy")]
    public partial class RepositoryPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// Name of the repository to apply the policy.
        /// </summary>
        [Output("repositoryName")]
        public Output<string> RepositoryName { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPolicy(string name, RepositoryPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:ecrpublic/repositoryPolicy:RepositoryPolicy", name, args ?? new RepositoryPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPolicy(string name, Input<string> id, RepositoryPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecrpublic/repositoryPolicy:RepositoryPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RepositoryPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPolicy Get(string name, Input<string> id, RepositoryPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryPolicy(name, id, state, options);
        }
    }

    public sealed class RepositoryPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// Name of the repository to apply the policy.
        /// </summary>
        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        public RepositoryPolicyArgs()
        {
        }
        public static new RepositoryPolicyArgs Empty => new RepositoryPolicyArgs();
    }

    public sealed class RepositoryPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// Name of the repository to apply the policy.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        public RepositoryPolicyState()
        {
        }
        public static new RepositoryPolicyState Empty => new RepositoryPolicyState();
    }
}
