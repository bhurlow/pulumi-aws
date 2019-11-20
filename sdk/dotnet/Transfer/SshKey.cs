// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer
{
    /// <summary>
    /// Provides a AWS Transfer User SSH Key resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_ssh_key.html.markdown.
    /// </summary>
    public partial class SshKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The public key portion of an SSH key pair.
        /// </summary>
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        /// <summary>
        /// The Server ID of the Transfer Server (e.g. `s-12345678`)
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// The name of the user account that is assigned to one or more servers.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a SshKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SshKey(string name, SshKeyArgs args, CustomResourceOptions? options = null)
            : base("aws:transfer/sshKey:SshKey", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SshKey(string name, Input<string> id, SshKeyState? state = null, CustomResourceOptions? options = null)
            : base("aws:transfer/sshKey:SshKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SshKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SshKey Get(string name, Input<string> id, SshKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new SshKey(name, id, state, options);
        }
    }

    public sealed class SshKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public key portion of an SSH key pair.
        /// </summary>
        [Input("body", required: true)]
        public Input<string> Body { get; set; } = null!;

        /// <summary>
        /// The Server ID of the Transfer Server (e.g. `s-12345678`)
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        /// <summary>
        /// The name of the user account that is assigned to one or more servers.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public SshKeyArgs()
        {
        }
    }

    public sealed class SshKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public key portion of an SSH key pair.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// The Server ID of the Transfer Server (e.g. `s-12345678`)
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// The name of the user account that is assigned to one or more servers.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public SshKeyState()
        {
        }
    }
}
