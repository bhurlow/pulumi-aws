// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Signer
{
    /// <summary>
    /// Creates a Signer Signing Profile. A signing profile contains information about the code signing configuration parameters that can be used by a given code signing user.
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
    ///     var testSp = new Aws.Signer.SigningProfile("testSp", new()
    ///     {
    ///         PlatformId = "AWSLambda-SHA384-ECDSA",
    ///     });
    /// 
    ///     var prodSp = new Aws.Signer.SigningProfile("prodSp", new()
    ///     {
    ///         NamePrefix = "prod_sp_",
    ///         PlatformId = "AWSLambda-SHA384-ECDSA",
    ///         SignatureValidityPeriod = new Aws.Signer.Inputs.SigningProfileSignatureValidityPeriodArgs
    ///         {
    ///             Type = "YEARS",
    ///             Value = 5,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "tag1", "value1" },
    ///             { "tag2", "value2" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Signer signing profiles using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:signer/signingProfile:SigningProfile test_signer_signing_profile test_sp_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK
    /// ```
    /// </summary>
    [AwsResourceType("aws:signer/signingProfile:SigningProfile")]
    public partial class SigningProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the signing profile.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// A human-readable name for the signing platform associated with the signing profile.
        /// </summary>
        [Output("platformDisplayName")]
        public Output<string> PlatformDisplayName { get; private set; } = null!;

        /// <summary>
        /// The ID of the platform that is used by the target signing profile.
        /// </summary>
        [Output("platformId")]
        public Output<string> PlatformId { get; private set; } = null!;

        /// <summary>
        /// Revocation information for a signing profile.
        /// </summary>
        [Output("revocationRecords")]
        public Output<ImmutableArray<Outputs.SigningProfileRevocationRecord>> RevocationRecords { get; private set; } = null!;

        /// <summary>
        /// The validity period for a signing job.
        /// </summary>
        [Output("signatureValidityPeriod")]
        public Output<Outputs.SigningProfileSignatureValidityPeriod> SignatureValidityPeriod { get; private set; } = null!;

        [Output("signingMaterial")]
        public Output<Outputs.SigningProfileSigningMaterial> SigningMaterial { get; private set; } = null!;

        /// <summary>
        /// The status of the target signing profile.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A list of tags associated with the signing profile. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The current version of the signing profile.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The signing profile ARN, including the profile version.
        /// </summary>
        [Output("versionArn")]
        public Output<string> VersionArn { get; private set; } = null!;


        /// <summary>
        /// Create a SigningProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SigningProfile(string name, SigningProfileArgs args, CustomResourceOptions? options = null)
            : base("aws:signer/signingProfile:SigningProfile", name, args ?? new SigningProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SigningProfile(string name, Input<string> id, SigningProfileState? state = null, CustomResourceOptions? options = null)
            : base("aws:signer/signingProfile:SigningProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SigningProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SigningProfile Get(string name, Input<string> id, SigningProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new SigningProfile(name, id, state, options);
        }
    }

    public sealed class SigningProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The ID of the platform that is used by the target signing profile.
        /// </summary>
        [Input("platformId", required: true)]
        public Input<string> PlatformId { get; set; } = null!;

        /// <summary>
        /// The validity period for a signing job.
        /// </summary>
        [Input("signatureValidityPeriod")]
        public Input<Inputs.SigningProfileSignatureValidityPeriodArgs>? SignatureValidityPeriod { get; set; }

        [Input("signingMaterial")]
        public Input<Inputs.SigningProfileSigningMaterialArgs>? SigningMaterial { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A list of tags associated with the signing profile. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SigningProfileArgs()
        {
        }
        public static new SigningProfileArgs Empty => new SigningProfileArgs();
    }

    public sealed class SigningProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the signing profile.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// A human-readable name for the signing platform associated with the signing profile.
        /// </summary>
        [Input("platformDisplayName")]
        public Input<string>? PlatformDisplayName { get; set; }

        /// <summary>
        /// The ID of the platform that is used by the target signing profile.
        /// </summary>
        [Input("platformId")]
        public Input<string>? PlatformId { get; set; }

        [Input("revocationRecords")]
        private InputList<Inputs.SigningProfileRevocationRecordGetArgs>? _revocationRecords;

        /// <summary>
        /// Revocation information for a signing profile.
        /// </summary>
        public InputList<Inputs.SigningProfileRevocationRecordGetArgs> RevocationRecords
        {
            get => _revocationRecords ?? (_revocationRecords = new InputList<Inputs.SigningProfileRevocationRecordGetArgs>());
            set => _revocationRecords = value;
        }

        /// <summary>
        /// The validity period for a signing job.
        /// </summary>
        [Input("signatureValidityPeriod")]
        public Input<Inputs.SigningProfileSignatureValidityPeriodGetArgs>? SignatureValidityPeriod { get; set; }

        [Input("signingMaterial")]
        public Input<Inputs.SigningProfileSigningMaterialGetArgs>? SigningMaterial { get; set; }

        /// <summary>
        /// The status of the target signing profile.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A list of tags associated with the signing profile. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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
        /// The current version of the signing profile.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The signing profile ARN, including the profile version.
        /// </summary>
        [Input("versionArn")]
        public Input<string>? VersionArn { get; set; }

        public SigningProfileState()
        {
        }
        public static new SigningProfileState Empty => new SigningProfileState();
    }
}
