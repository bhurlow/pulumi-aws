// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Shield
{
    /// <summary>
    /// Creates an association between a Route53 Health Check and a Shield Advanced protected resource.
    /// This association uses the health of your applications to improve responsiveness and accuracy in attack detection and mitigation.
    /// 
    /// Blog post: [AWS Shield Advanced now supports Health Based Detection](https://aws.amazon.com/about-aws/whats-new/2020/02/aws-shield-advanced-now-supports-health-based-detection/)
    /// 
    /// ## Example Usage
    /// ### Create an association between a protected EIP and a Route53 Health Check
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var currentRegion = Aws.GetRegion.Invoke();
    /// 
    ///     var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();
    /// 
    ///     var currentPartition = Aws.GetPartition.Invoke();
    /// 
    ///     var exampleEip = new Aws.Ec2.Eip("exampleEip", new()
    ///     {
    ///         Domain = "vpc",
    ///         Tags = 
    ///         {
    ///             { "Name", "example" },
    ///         },
    ///     });
    /// 
    ///     var exampleProtection = new Aws.Shield.Protection("exampleProtection", new()
    ///     {
    ///         ResourceArn = Output.Tuple(currentPartition, currentRegion, currentCallerIdentity, exampleEip.Id).Apply(values =&gt;
    ///         {
    ///             var currentPartition = values.Item1;
    ///             var currentRegion = values.Item2;
    ///             var currentCallerIdentity = values.Item3;
    ///             var id = values.Item4;
    ///             return $"arn:{currentPartition.Apply(getPartitionResult =&gt; getPartitionResult.Partition)}:ec2:{currentRegion.Apply(getRegionResult =&gt; getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult =&gt; getCallerIdentityResult.AccountId)}:eip-allocation/{id}";
    ///         }),
    ///     });
    /// 
    ///     var exampleHealthCheck = new Aws.Route53.HealthCheck("exampleHealthCheck", new()
    ///     {
    ///         IpAddress = exampleEip.PublicIp,
    ///         Port = 80,
    ///         Type = "HTTP",
    ///         ResourcePath = "/ready",
    ///         FailureThreshold = 3,
    ///         RequestInterval = 30,
    ///         Tags = 
    ///         {
    ///             { "Name", "tf-example-health-check" },
    ///         },
    ///     });
    /// 
    ///     var exampleProtectionHealthCheckAssociation = new Aws.Shield.ProtectionHealthCheckAssociation("exampleProtectionHealthCheckAssociation", new()
    ///     {
    ///         HealthCheckArn = exampleHealthCheck.Arn,
    ///         ShieldProtectionId = exampleProtection.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_shield_protection_health_check_association.example
    /// 
    ///  id = "ff9592dc-22f3-4e88-afa1-7b29fde9669a+arn:aws:route53:::healthcheck/3742b175-edb9-46bc-9359-f53e3b794b1b" } Using `pulumi import`, import Shield protection health check association resources using the `shield_protection_id` and `health_check_arn`. For exampleconsole % pulumi import aws_shield_protection_health_check_association.example ff9592dc-22f3-4e88-afa1-7b29fde9669a+arn:aws:route53:::healthcheck/3742b175-edb9-46bc-9359-f53e3b794b1b
    /// </summary>
    [AwsResourceType("aws:shield/protectionHealthCheckAssociation:ProtectionHealthCheckAssociation")]
    public partial class ProtectionHealthCheckAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the Route53 Health Check resource which will be associated to the protected resource.
        /// </summary>
        [Output("healthCheckArn")]
        public Output<string> HealthCheckArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the protected resource.
        /// </summary>
        [Output("shieldProtectionId")]
        public Output<string> ShieldProtectionId { get; private set; } = null!;


        /// <summary>
        /// Create a ProtectionHealthCheckAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProtectionHealthCheckAssociation(string name, ProtectionHealthCheckAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:shield/protectionHealthCheckAssociation:ProtectionHealthCheckAssociation", name, args ?? new ProtectionHealthCheckAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProtectionHealthCheckAssociation(string name, Input<string> id, ProtectionHealthCheckAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:shield/protectionHealthCheckAssociation:ProtectionHealthCheckAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProtectionHealthCheckAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProtectionHealthCheckAssociation Get(string name, Input<string> id, ProtectionHealthCheckAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new ProtectionHealthCheckAssociation(name, id, state, options);
        }
    }

    public sealed class ProtectionHealthCheckAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the Route53 Health Check resource which will be associated to the protected resource.
        /// </summary>
        [Input("healthCheckArn", required: true)]
        public Input<string> HealthCheckArn { get; set; } = null!;

        /// <summary>
        /// The ID of the protected resource.
        /// </summary>
        [Input("shieldProtectionId", required: true)]
        public Input<string> ShieldProtectionId { get; set; } = null!;

        public ProtectionHealthCheckAssociationArgs()
        {
        }
        public static new ProtectionHealthCheckAssociationArgs Empty => new ProtectionHealthCheckAssociationArgs();
    }

    public sealed class ProtectionHealthCheckAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the Route53 Health Check resource which will be associated to the protected resource.
        /// </summary>
        [Input("healthCheckArn")]
        public Input<string>? HealthCheckArn { get; set; }

        /// <summary>
        /// The ID of the protected resource.
        /// </summary>
        [Input("shieldProtectionId")]
        public Input<string>? ShieldProtectionId { get; set; }

        public ProtectionHealthCheckAssociationState()
        {
        }
        public static new ProtectionHealthCheckAssociationState Empty => new ProtectionHealthCheckAssociationState();
    }
}
