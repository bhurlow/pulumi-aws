// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Efs
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about an Elastic File System (EFS).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/efs_file_system.html.markdown.
        /// </summary>
        public static Task<GetFileSystemResult> GetFileSystem(GetFileSystemArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFileSystemResult>("aws:efs/getFileSystem:getFileSystem", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetFileSystemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Restricts the list to the file system with this creation token.
        /// </summary>
        [Input("creationToken")]
        public Input<string>? CreationToken { get; set; }

        /// <summary>
        /// The ID that identifies the file system (e.g. fs-ccfc0d65).
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetFileSystemArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetFileSystemResult
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        public readonly string Arn;
        public readonly string CreationToken;
        /// <summary>
        /// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        /// </summary>
        public readonly string DnsName;
        /// <summary>
        /// Whether EFS is encrypted.
        /// </summary>
        public readonly bool Encrypted;
        public readonly string FileSystemId;
        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// The PerformanceMode of the file system.
        /// </summary>
        public readonly string PerformanceMode;
        /// <summary>
        /// The list of tags assigned to the file system.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFileSystemResult(
            string arn,
            string creationToken,
            string dnsName,
            bool encrypted,
            string fileSystemId,
            string kmsKeyId,
            string performanceMode,
            ImmutableDictionary<string, object> tags,
            string id)
        {
            Arn = arn;
            CreationToken = creationToken;
            DnsName = dnsName;
            Encrypted = encrypted;
            FileSystemId = fileSystemId;
            KmsKeyId = kmsKeyId;
            PerformanceMode = performanceMode;
            Tags = tags;
            Id = id;
        }
    }
}
