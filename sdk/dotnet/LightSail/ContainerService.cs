// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LightSail
{
    /// <summary>
    /// An Amazon Lightsail container service is a highly scalable compute and networking resource on which you can deploy, run,
    /// and manage containers. For more information, see
    /// [Container services in Amazon Lightsail](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-container-services).
    /// 
    /// &gt; **Note:** For more information about the AWS Regions in which you can create Amazon Lightsail container services,
    /// see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail).
    /// 
    /// ## Example Usage
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
    ///     var myContainerService = new Aws.LightSail.ContainerService("myContainerService", new()
    ///     {
    ///         IsDisabled = false,
    ///         Power = "nano",
    ///         Scale = 1,
    ///         Tags = 
    ///         {
    ///             { "foo1", "bar1" },
    ///             { "foo2", "" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Public Domain Names
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myContainerService = new Aws.LightSail.ContainerService("myContainerService", new()
    ///     {
    ///         PublicDomainNames = new Aws.LightSail.Inputs.ContainerServicePublicDomainNamesArgs
    ///         {
    ///             Certificates = new[]
    ///             {
    ///                 new Aws.LightSail.Inputs.ContainerServicePublicDomainNamesCertificateArgs
    ///                 {
    ///                     CertificateName = "example-certificate",
    ///                     DomainNames = new[]
    ///                     {
    ///                         "www.example.com",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Private Registry Access
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // ... other configuration ...
    ///     var defaultContainerService = new Aws.LightSail.ContainerService("defaultContainerService", new()
    ///     {
    ///         PrivateRegistryAccess = new Aws.LightSail.Inputs.ContainerServicePrivateRegistryAccessArgs
    ///         {
    ///             EcrImagePullerRole = new Aws.LightSail.Inputs.ContainerServicePrivateRegistryAccessEcrImagePullerRoleArgs
    ///             {
    ///                 IsActive = true,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             defaultContainerService.PrivateRegistryAccess.EcrImagePullerRole?.PrincipalArn,
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "ecr:BatchGetImage",
    ///                     "ecr:GetDownloadUrlForLayer",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultRepositoryPolicy = new Aws.Ecr.RepositoryPolicy("defaultRepositoryPolicy", new()
    ///     {
    ///         Repository = aws_ecr_repository.Default.Name,
    ///         Policy = defaultPolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Lightsail Container Service using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:lightsail/containerService:ContainerService my_container_service container-service-1
    /// ```
    /// </summary>
    [AwsResourceType("aws:lightsail/containerService:ContainerService")]
    public partial class ContainerService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the container service.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone. Follows the format us-east-2a (case-sensitive).
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A Boolean value indicating whether the container service is disabled. Defaults to `false`.
        /// </summary>
        [Output("isDisabled")]
        public Output<bool?> IsDisabled { get; private set; } = null!;

        /// <summary>
        /// The name for the container service. Names must be of length 1 to 63, and be
        /// unique within each AWS Region in your Lightsail account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The power specification for the container service. The power specifies the amount of memory,
        /// the number of vCPUs, and the monthly price of each node of the container service.
        /// Possible values: `nano`, `micro`, `small`, `medium`, `large`, `xlarge`.
        /// </summary>
        [Output("power")]
        public Output<string> Power { get; private set; } = null!;

        /// <summary>
        /// The ID of the power of the container service.
        /// </summary>
        [Output("powerId")]
        public Output<string> PowerId { get; private set; } = null!;

        /// <summary>
        /// The principal ARN of the container service. The principal ARN can be used to create a trust
        /// relationship between your standard AWS account and your Lightsail container service. This allows you to give your
        /// service permission to access resources in your standard AWS account.
        /// </summary>
        [Output("principalArn")]
        public Output<string> PrincipalArn { get; private set; } = null!;

        /// <summary>
        /// The private domain name of the container service. The private domain name is accessible only
        /// by other resources within the default virtual private cloud (VPC) of your Lightsail account.
        /// </summary>
        [Output("privateDomainName")]
        public Output<string> PrivateDomainName { get; private set; } = null!;

        /// <summary>
        /// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See Private Registry Access below for more details.
        /// </summary>
        [Output("privateRegistryAccess")]
        public Output<Outputs.ContainerServicePrivateRegistryAccess> PrivateRegistryAccess { get; private set; } = null!;

        /// <summary>
        /// The public domain names to use with the container service, such as example.com
        /// and www.example.com. You can specify up to four public domain names for a container service. The domain names that you
        /// specify are used when you create a deployment with a container configured as the public endpoint of your container
        /// service. If you don't specify public domain names, then you can use the default domain of the container service.
        /// Defined below.
        /// </summary>
        [Output("publicDomainNames")]
        public Output<Outputs.ContainerServicePublicDomainNames?> PublicDomainNames { get; private set; } = null!;

        /// <summary>
        /// The Lightsail resource type of the container service (i.e., ContainerService).
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// The scale specification for the container service. The scale specifies the allocated compute
        /// nodes of the container service.
        /// </summary>
        [Output("scale")]
        public Output<int> Scale { get; private set; } = null!;

        /// <summary>
        /// The current state of the container service.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Map of container service tags. To tag at launch, specify the tags in the Launch Template. If
        /// configured with a provider
        /// `default_tags` configuration block
        /// present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider
        /// `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible URL of the container service. If no public endpoint is specified in the
        /// currentDeployment, this URL returns a 404 response.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerService(string name, ContainerServiceArgs args, CustomResourceOptions? options = null)
            : base("aws:lightsail/containerService:ContainerService", name, args ?? new ContainerServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerService(string name, Input<string> id, ContainerServiceState? state = null, CustomResourceOptions? options = null)
            : base("aws:lightsail/containerService:ContainerService", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ContainerService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerService Get(string name, Input<string> id, ContainerServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerService(name, id, state, options);
        }
    }

    public sealed class ContainerServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Boolean value indicating whether the container service is disabled. Defaults to `false`.
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// The name for the container service. Names must be of length 1 to 63, and be
        /// unique within each AWS Region in your Lightsail account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The power specification for the container service. The power specifies the amount of memory,
        /// the number of vCPUs, and the monthly price of each node of the container service.
        /// Possible values: `nano`, `micro`, `small`, `medium`, `large`, `xlarge`.
        /// </summary>
        [Input("power", required: true)]
        public Input<string> Power { get; set; } = null!;

        /// <summary>
        /// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See Private Registry Access below for more details.
        /// </summary>
        [Input("privateRegistryAccess")]
        public Input<Inputs.ContainerServicePrivateRegistryAccessArgs>? PrivateRegistryAccess { get; set; }

        /// <summary>
        /// The public domain names to use with the container service, such as example.com
        /// and www.example.com. You can specify up to four public domain names for a container service. The domain names that you
        /// specify are used when you create a deployment with a container configured as the public endpoint of your container
        /// service. If you don't specify public domain names, then you can use the default domain of the container service.
        /// Defined below.
        /// </summary>
        [Input("publicDomainNames")]
        public Input<Inputs.ContainerServicePublicDomainNamesArgs>? PublicDomainNames { get; set; }

        /// <summary>
        /// The scale specification for the container service. The scale specifies the allocated compute
        /// nodes of the container service.
        /// </summary>
        [Input("scale", required: true)]
        public Input<int> Scale { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of container service tags. To tag at launch, specify the tags in the Launch Template. If
        /// configured with a provider
        /// `default_tags` configuration block
        /// present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ContainerServiceArgs()
        {
        }
        public static new ContainerServiceArgs Empty => new ContainerServiceArgs();
    }

    public sealed class ContainerServiceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the container service.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Availability Zone. Follows the format us-east-2a (case-sensitive).
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A Boolean value indicating whether the container service is disabled. Defaults to `false`.
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// The name for the container service. Names must be of length 1 to 63, and be
        /// unique within each AWS Region in your Lightsail account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The power specification for the container service. The power specifies the amount of memory,
        /// the number of vCPUs, and the monthly price of each node of the container service.
        /// Possible values: `nano`, `micro`, `small`, `medium`, `large`, `xlarge`.
        /// </summary>
        [Input("power")]
        public Input<string>? Power { get; set; }

        /// <summary>
        /// The ID of the power of the container service.
        /// </summary>
        [Input("powerId")]
        public Input<string>? PowerId { get; set; }

        /// <summary>
        /// The principal ARN of the container service. The principal ARN can be used to create a trust
        /// relationship between your standard AWS account and your Lightsail container service. This allows you to give your
        /// service permission to access resources in your standard AWS account.
        /// </summary>
        [Input("principalArn")]
        public Input<string>? PrincipalArn { get; set; }

        /// <summary>
        /// The private domain name of the container service. The private domain name is accessible only
        /// by other resources within the default virtual private cloud (VPC) of your Lightsail account.
        /// </summary>
        [Input("privateDomainName")]
        public Input<string>? PrivateDomainName { get; set; }

        /// <summary>
        /// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See Private Registry Access below for more details.
        /// </summary>
        [Input("privateRegistryAccess")]
        public Input<Inputs.ContainerServicePrivateRegistryAccessGetArgs>? PrivateRegistryAccess { get; set; }

        /// <summary>
        /// The public domain names to use with the container service, such as example.com
        /// and www.example.com. You can specify up to four public domain names for a container service. The domain names that you
        /// specify are used when you create a deployment with a container configured as the public endpoint of your container
        /// service. If you don't specify public domain names, then you can use the default domain of the container service.
        /// Defined below.
        /// </summary>
        [Input("publicDomainNames")]
        public Input<Inputs.ContainerServicePublicDomainNamesGetArgs>? PublicDomainNames { get; set; }

        /// <summary>
        /// The Lightsail resource type of the container service (i.e., ContainerService).
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The scale specification for the container service. The scale specifies the allocated compute
        /// nodes of the container service.
        /// </summary>
        [Input("scale")]
        public Input<int>? Scale { get; set; }

        /// <summary>
        /// The current state of the container service.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of container service tags. To tag at launch, specify the tags in the Launch Template. If
        /// configured with a provider
        /// `default_tags` configuration block
        /// present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider
        /// `default_tags` configuration block.
        /// </summary>
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The publicly accessible URL of the container service. If no public endpoint is specified in the
        /// currentDeployment, this URL returns a 404 response.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ContainerServiceState()
        {
        }
        public static new ContainerServiceState Empty => new ContainerServiceState();
    }
}
