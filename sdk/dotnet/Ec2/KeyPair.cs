// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides an [EC2 key pair](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) resource. A key pair is used to control login access to EC2 instances.
    /// 
    /// Currently this resource requires an existing user-supplied key pair. This key pair's public key will be registered with AWS to allow logging-in to EC2 instances.
    /// 
    /// When importing an existing key pair the public key material may be in any format supported by AWS. Supported formats (per the [AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html#how-to-generate-your-own-key-and-import-it-to-aws)) are:
    /// 
    /// * OpenSSH public key format (the format in ~/.ssh/authorized_keys)
    /// * Base64 encoded DER format
    /// * SSH public key file format as specified in RFC4716
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var deployer = new Aws.Ec2.KeyPair("deployer", new Aws.Ec2.KeyPairArgs
    ///         {
    ///             PublicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD3F6tyPEFEzV0LX3X8BsXdMsQz1x2cEikKDEY0aIj41qgxMCP/iteneqXSIFZBp5vizPvaoIR3Um9xK7PGoW8giupGn+EPuxIA4cDM4vzOqOkiMPhz5XK0whEjkVzTo4+S0puvDZuwIsdiW9mxhJc7tgBNL0cYlWSYVkz4G/fslNfRPW5mYAM49f4fhtxPb5ok4Q2Lg9dPKVHO/Bgeu5woMc7RY0p1ej6D4CKFE6lymSDJpW0YHX/wqE9+cfEauh7xZcG0q9t2ta6F6fmX0agvpFyZo8aFbXeUBr7osSCJNgvavWbM/06niWrOvYX2xwWdhXmXSrbX8ZbabVohBK41 email@example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Key Pairs can be imported using the `key_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/keyPair:KeyPair deployer deployer-key
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/keyPair:KeyPair")]
    public partial class KeyPair : Pulumi.CustomResource
    {
        /// <summary>
        /// The key pair ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The name for the key pair.
        /// </summary>
        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `key_name`.
        /// </summary>
        [Output("keyNamePrefix")]
        public Output<string?> KeyNamePrefix { get; private set; } = null!;

        /// <summary>
        /// The key pair ID.
        /// </summary>
        [Output("keyPairId")]
        public Output<string> KeyPairId { get; private set; } = null!;

        /// <summary>
        /// The public key material.
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a KeyPair resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyPair(string name, KeyPairArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/keyPair:KeyPair", name, args ?? new KeyPairArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyPair(string name, Input<string> id, KeyPairState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/keyPair:KeyPair", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyPair resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyPair Get(string name, Input<string> id, KeyPairState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyPair(name, id, state, options);
        }
    }

    public sealed class KeyPairArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for the key pair.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `key_name`.
        /// </summary>
        [Input("keyNamePrefix")]
        public Input<string>? KeyNamePrefix { get; set; }

        /// <summary>
        /// The public key material.
        /// </summary>
        [Input("publicKey", required: true)]
        public Input<string> PublicKey { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public KeyPairArgs()
        {
        }
    }

    public sealed class KeyPairState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key pair ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The name for the key pair.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `key_name`.
        /// </summary>
        [Input("keyNamePrefix")]
        public Input<string>? KeyNamePrefix { get; set; }

        /// <summary>
        /// The key pair ID.
        /// </summary>
        [Input("keyPairId")]
        public Input<string>? KeyPairId { get; set; }

        /// <summary>
        /// The public key material.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public KeyPairState()
        {
        }
    }
}
