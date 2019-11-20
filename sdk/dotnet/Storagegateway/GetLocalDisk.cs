// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.StorageGateway
{
    public static partial class Invokes
    {
        /// <summary>
        /// Retrieve information about a Storage Gateway local disk. The disk identifier is useful for adding the disk as a cache or upload buffer to a gateway.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/storagegateway_local_disk.html.markdown.
        /// </summary>
        public static Task<GetLocalDiskResult> GetLocalDisk(GetLocalDiskArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLocalDiskResult>("aws:storagegateway/getLocalDisk:getLocalDisk", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetLocalDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device node of the local disk to retrieve. For example, `/dev/sdb`.
        /// </summary>
        [Input("diskNode")]
        public Input<string>? DiskNode { get; set; }

        /// <summary>
        /// The device path of the local disk to retrieve. For example, `/dev/xvdb` or `/dev/nvme1n1`.
        /// </summary>
        [Input("diskPath")]
        public Input<string>? DiskPath { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Input("gatewayArn", required: true)]
        public Input<string> GatewayArn { get; set; } = null!;

        public GetLocalDiskArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetLocalDiskResult
    {
        /// <summary>
        /// The disk identifier. e.g. `pci-0000:03:00.0-scsi-0:0:0:0`
        /// </summary>
        public readonly string DiskId;
        public readonly string? DiskNode;
        public readonly string? DiskPath;
        public readonly string GatewayArn;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetLocalDiskResult(
            string diskId,
            string? diskNode,
            string? diskPath,
            string gatewayArn,
            string id)
        {
            DiskId = diskId;
            DiskNode = diskNode;
            DiskPath = diskPath;
            GatewayArn = gatewayArn;
            Id = id;
        }
    }
}
