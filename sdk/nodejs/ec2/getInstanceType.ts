// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Get characteristics for a single EC2 Instance Type.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ec2.getInstanceType({
 *     instanceType: "t2.micro",
 * }));
 * ```
 */
export function getInstanceType(args: GetInstanceTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTypeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getInstanceType:getInstanceType", {
        "defaultCores": args.defaultCores,
        "defaultThreadsPerCore": args.defaultThreadsPerCore,
        "fpgas": args.fpgas,
        "gpuses": args.gpuses,
        "hypervisor": args.hypervisor,
        "inferenceAccelerators": args.inferenceAccelerators,
        "instanceDisks": args.instanceDisks,
        "instanceType": args.instanceType,
        "maximumIpv6AddressesPerInterface": args.maximumIpv6AddressesPerInterface,
        "totalFpgaMemory": args.totalFpgaMemory,
        "totalGpuMemory": args.totalGpuMemory,
        "totalInstanceStorage": args.totalInstanceStorage,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceType.
 */
export interface GetInstanceTypeArgs {
    /**
     * The default number of cores for the instance type.
     */
    defaultCores?: number;
    /**
     * The  default  number of threads per core for the instance type.
     */
    defaultThreadsPerCore?: number;
    /**
     * Describes the FPGA accelerator settings for the instance type.
     * * `fpgas.#.count` - The count of FPGA accelerators for the instance type.
     * * `fpgas.#.manufacturer` - The manufacturer of the FPGA accelerator.
     * * `fpgas.#.memory_size` - The size (in MiB) for the memory available to the FPGA accelerator.
     * * `fpgas.#.name` - The name of the FPGA accelerator.
     */
    fpgas?: inputs.ec2.GetInstanceTypeFpga[];
    /**
     * Describes the GPU accelerators for the instance type.
     * * `gpus.#.count` - The number of GPUs for the instance type.
     * * `gpus.#.manufacturer` - The manufacturer of the GPU accelerator.
     * * `gpus.#.memory_size` - The size (in MiB) for the memory available to the GPU accelerator.
     * * `gpus.#.name` - The name of the GPU accelerator.
     */
    gpuses?: inputs.ec2.GetInstanceTypeGpus[];
    /**
     * Indicates the hypervisor used for the instance type.
     * * `inferenceAccelerators` Describes the Inference accelerators for the instance type.
     * * `inference_accelerators.#.count` - The number of Inference accelerators for the instance type.
     * * `inference_accelerators.#.manufacturer` - The manufacturer of the Inference accelerator.
     * * `inference_accelerators.#.name` - The name of the Inference accelerator.
     */
    hypervisor?: string;
    inferenceAccelerators?: inputs.ec2.GetInstanceTypeInferenceAccelerator[];
    /**
     * Describes the disks for the instance type.
     * * `instance_disks.#.count` - The number of disks with this configuration.
     * * `instance_disks.#.size` - The size of the disk in GB.
     * * `instance_disks.#.type` - The type of disk.
     */
    instanceDisks?: inputs.ec2.GetInstanceTypeInstanceDisk[];
    /**
     * Instance
     */
    instanceType: string;
    /**
     * The maximum number of IPv6 addresses per network interface.
     */
    maximumIpv6AddressesPerInterface?: number;
    /**
     * The total memory of all FPGA accelerators for the instance type (in MiB).
     */
    totalFpgaMemory?: number;
    /**
     * The total size of the memory for the GPU accelerators for the instance type (in MiB).
     */
    totalGpuMemory?: number;
    /**
     * The total size of the instance disks, in GB.
     */
    totalInstanceStorage?: number;
}

/**
 * A collection of values returned by getInstanceType.
 */
export interface GetInstanceTypeResult {
    /**
     * `true` if auto recovery is supported.
     */
    readonly autoRecoverySupported: boolean;
    /**
     * `true` if it is a bare metal instance type.
     */
    readonly bareMetal: boolean;
    /**
     * `true` if the instance type is a burstable performance instance type.
     */
    readonly burstablePerformanceSupported: boolean;
    /**
     * `true`  if the instance type is a current generation.
     */
    readonly currentGeneration: boolean;
    /**
     * `true` if Dedicated Hosts are supported on the instance type.
     */
    readonly dedicatedHostsSupported: boolean;
    /**
     * The default number of cores for the instance type.
     */
    readonly defaultCores: number;
    /**
     * The  default  number of threads per core for the instance type.
     */
    readonly defaultThreadsPerCore: number;
    /**
     * The default number of vCPUs for the instance type.
     */
    readonly defaultVcpus: number;
    /**
     * Indicates whether Amazon EBS encryption is supported.
     */
    readonly ebsEncryptionSupport: string;
    /**
     * Indicates whether non-volatile memory express (NVMe) is supported.
     */
    readonly ebsNvmeSupport: string;
    /**
     * Indicates that the instance type is Amazon EBS-optimized.
     */
    readonly ebsOptimizedSupport: string;
    /**
     * The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.
     */
    readonly ebsPerformanceBaselineBandwidth: number;
    /**
     * The baseline input/output storage operations per seconds for an EBS-optimized instance type.
     */
    readonly ebsPerformanceBaselineIops: number;
    /**
     * The baseline throughput performance for an EBS-optimized instance type, in MBps.
     */
    readonly ebsPerformanceBaselineThroughput: number;
    /**
     * The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.
     */
    readonly ebsPerformanceMaximumBandwidth: number;
    /**
     * The maximum input/output storage operations per second for an EBS-optimized instance type.
     */
    readonly ebsPerformanceMaximumIops: number;
    /**
     * The maximum throughput performance for an EBS-optimized instance type, in MBps.
     */
    readonly ebsPerformanceMaximumThroughput: number;
    /**
     * Indicates whether Elastic Fabric Adapter (EFA) is supported.
     */
    readonly efaSupported: boolean;
    /**
     * Indicates whether Elastic Network Adapter (ENA) is supported.
     */
    readonly enaSupport: string;
    /**
     * Describes the FPGA accelerator settings for the instance type.
     * * `fpgas.#.count` - The count of FPGA accelerators for the instance type.
     * * `fpgas.#.manufacturer` - The manufacturer of the FPGA accelerator.
     * * `fpgas.#.memory_size` - The size (in MiB) for the memory available to the FPGA accelerator.
     * * `fpgas.#.name` - The name of the FPGA accelerator.
     */
    readonly fpgas: outputs.ec2.GetInstanceTypeFpga[];
    /**
     * `true` if the instance type is eligible for the free tier.
     */
    readonly freeTierEligible: boolean;
    /**
     * Describes the GPU accelerators for the instance type.
     * * `gpus.#.count` - The number of GPUs for the instance type.
     * * `gpus.#.manufacturer` - The manufacturer of the GPU accelerator.
     * * `gpus.#.memory_size` - The size (in MiB) for the memory available to the GPU accelerator.
     * * `gpus.#.name` - The name of the GPU accelerator.
     */
    readonly gpuses: outputs.ec2.GetInstanceTypeGpus[];
    /**
     * `true` if On-Demand hibernation is supported.
     */
    readonly hibernationSupported: boolean;
    /**
     * Indicates the hypervisor used for the instance type.
     * * `inferenceAccelerators` Describes the Inference accelerators for the instance type.
     * * `inference_accelerators.#.count` - The number of Inference accelerators for the instance type.
     * * `inference_accelerators.#.manufacturer` - The manufacturer of the Inference accelerator.
     * * `inference_accelerators.#.name` - The name of the Inference accelerator.
     */
    readonly hypervisor: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly inferenceAccelerators: outputs.ec2.GetInstanceTypeInferenceAccelerator[];
    /**
     * Describes the disks for the instance type.
     * * `instance_disks.#.count` - The number of disks with this configuration.
     * * `instance_disks.#.size` - The size of the disk in GB.
     * * `instance_disks.#.type` - The type of disk.
     */
    readonly instanceDisks: outputs.ec2.GetInstanceTypeInstanceDisk[];
    /**
     * `true` if instance storage is supported.
     */
    readonly instanceStorageSupported: boolean;
    readonly instanceType: string;
    /**
     * `true` if IPv6 is supported.
     */
    readonly ipv6Supported: boolean;
    /**
     * The maximum number of IPv4 addresses per network interface.
     */
    readonly maximumIpv4AddressesPerInterface: number;
    /**
     * The maximum number of IPv6 addresses per network interface.
     */
    readonly maximumIpv6AddressesPerInterface: number;
    /**
     * The maximum number of network interfaces for the instance type.
     */
    readonly maximumNetworkInterfaces: number;
    /**
     * Size of the instance memory, in MiB.
     */
    readonly memorySize: number;
    /**
     * Describes the network performance.
     */
    readonly networkPerformance: string;
    /**
     * A list of architectures supported by the instance type.
     */
    readonly supportedArchitectures: string[];
    /**
     * A list of supported placement groups types.
     */
    readonly supportedPlacementStrategies: string[];
    /**
     * Indicates the supported root device types.
     */
    readonly supportedRootDeviceTypes: string[];
    /**
     * Indicates whether the instance type is offered for spot or On-Demand.
     */
    readonly supportedUsagesClasses: string[];
    /**
     * The supported virtualization types.
     */
    readonly supportedVirtualizationTypes: string[];
    /**
     * The speed of the processor, in GHz.
     */
    readonly sustainedClockSpeed: number;
    /**
     * The total memory of all FPGA accelerators for the instance type (in MiB).
     */
    readonly totalFpgaMemory: number;
    /**
     * The total size of the memory for the GPU accelerators for the instance type (in MiB).
     */
    readonly totalGpuMemory: number;
    /**
     * The total size of the instance disks, in GB.
     */
    readonly totalInstanceStorage: number;
    /**
     * List of the valid number of cores that can be configured for the instance type.
     */
    readonly validCores: number[];
    /**
     * List of the valid number of threads per core that can be configured for the instance type.
     */
    readonly validThreadsPerCores: number[];
}
