// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kms
{
    public static partial class Invokes
    {
        /// <summary>
        /// Decrypt multiple secrets from data encrypted with the AWS KMS service.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_secrets.html.markdown.
        /// </summary>
        public static Task<GetSecretsResult> GetSecrets(GetSecretsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretsResult>("aws:kms/getSecrets:getSecrets", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetSecretsArgs : Pulumi.ResourceArgs
    {
        [Input("secrets", required: true)]
        private InputList<Inputs.GetSecretsSecretsArgs>? _secrets;

        /// <summary>
        /// One or more encrypted payload definitions from the KMS service. See the Secret Definitions below.
        /// </summary>
        public InputList<Inputs.GetSecretsSecretsArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.GetSecretsSecretsArgs>());
            set => _secrets = value;
        }

        public GetSecretsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSecretsResult
    {
        /// <summary>
        /// Map containing each `secret` `name` as the key with its decrypted plaintext value
        /// </summary>
        public readonly ImmutableDictionary<string, string> Plaintext;
        public readonly ImmutableArray<Outputs.GetSecretsSecretsResult> Secrets;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSecretsResult(
            ImmutableDictionary<string, string> plaintext,
            ImmutableArray<Outputs.GetSecretsSecretsResult> secrets,
            string id)
        {
            Plaintext = plaintext;
            Secrets = secrets;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetSecretsSecretsArgs : Pulumi.ResourceArgs
    {
        [Input("context")]
        private InputMap<string>? _context;

        /// <summary>
        /// An optional mapping that makes up the Encryption Context for the secret.
        /// </summary>
        public InputMap<string> Context
        {
            get => _context ?? (_context = new InputMap<string>());
            set => _context = value;
        }

        [Input("grantTokens")]
        private InputList<string>? _grantTokens;

        /// <summary>
        /// An optional list of Grant Tokens for the secret.
        /// </summary>
        public InputList<string> GrantTokens
        {
            get => _grantTokens ?? (_grantTokens = new InputList<string>());
            set => _grantTokens = value;
        }

        /// <summary>
        /// The name to export this secret under in the attributes.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Base64 encoded payload, as returned from a KMS encrypt operation.
        /// </summary>
        [Input("payload", required: true)]
        public Input<string> Payload { get; set; } = null!;

        public GetSecretsSecretsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetSecretsSecretsResult
    {
        /// <summary>
        /// An optional mapping that makes up the Encryption Context for the secret.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Context;
        /// <summary>
        /// An optional list of Grant Tokens for the secret.
        /// </summary>
        public readonly ImmutableArray<string> GrantTokens;
        /// <summary>
        /// The name to export this secret under in the attributes.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Base64 encoded payload, as returned from a KMS encrypt operation.
        /// </summary>
        public readonly string Payload;

        [OutputConstructor]
        private GetSecretsSecretsResult(
            ImmutableDictionary<string, string>? context,
            ImmutableArray<string> grantTokens,
            string name,
            string payload)
        {
            Context = context;
            GrantTokens = grantTokens;
            Name = name;
            Payload = payload;
        }
    }
    }
}
