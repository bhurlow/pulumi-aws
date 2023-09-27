// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieve information on FSx ONTAP Storage Virtual Machine (SVM).
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.fsx.getOntapStorageVirtualMachine({
 *     id: "svm-12345678",
 * });
 * ```
 * ### Filter Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.fsx.getOntapStorageVirtualMachine({
 *     filters: [{
 *         name: "file-system-id",
 *         values: ["fs-12345678"],
 *     }],
 * });
 * ```
 */
export function getOntapStorageVirtualMachine(args?: GetOntapStorageVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetOntapStorageVirtualMachineResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:fsx/getOntapStorageVirtualMachine:getOntapStorageVirtualMachine", {
        "filters": args.filters,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getOntapStorageVirtualMachine.
 */
export interface GetOntapStorageVirtualMachineArgs {
    /**
     * Configuration block. Detailed below.
     */
    filters?: inputs.fsx.GetOntapStorageVirtualMachineFilter[];
    /**
     * Identifier of the storage virtual machine (e.g. `svm-12345678`).
     */
    id?: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getOntapStorageVirtualMachine.
 */
export interface GetOntapStorageVirtualMachineResult {
    /**
     * The Microsoft Active Directory configuration to which the SVM is joined, if applicable. See Active Directory Configuration below.
     */
    readonly activeDirectoryConfigurations: outputs.fsx.GetOntapStorageVirtualMachineActiveDirectoryConfiguration[];
    /**
     * Amazon Resource Name of the SVM.
     */
    readonly arn: string;
    /**
     * The time that the SVM was created.
     */
    readonly creationTime: string;
    /**
     * The endpoints that are used to access data or to manage the SVM using the NetApp ONTAP CLI, REST API, or NetApp CloudManager. They are the Iscsi, Management, Nfs, and Smb endpoints. See SVM Endpoints below.
     */
    readonly endpoints: outputs.fsx.GetOntapStorageVirtualMachineEndpoint[];
    /**
     * Identifier of the file system (e.g. `fs-12345678`).
     */
    readonly fileSystemId: string;
    readonly filters?: outputs.fsx.GetOntapStorageVirtualMachineFilter[];
    /**
     * The SVM's system generated unique ID.
     */
    readonly id: string;
    /**
     * The SVM's lifecycle status.
     */
    readonly lifecycleStatus: string;
    /**
     * Describes why the SVM lifecycle state changed. See Lifecycle Transition Reason below.
     */
    readonly lifecycleTransitionReasons: outputs.fsx.GetOntapStorageVirtualMachineLifecycleTransitionReason[];
    /**
     * The name of the SVM, if provisioned.
     */
    readonly name: string;
    /**
     * The SVM's subtype.
     */
    readonly subtype: string;
    readonly tags: {[key: string]: string};
    /**
     * The SVM's UUID.
     */
    readonly uuid: string;
}
/**
 * Retrieve information on FSx ONTAP Storage Virtual Machine (SVM).
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.fsx.getOntapStorageVirtualMachine({
 *     id: "svm-12345678",
 * });
 * ```
 * ### Filter Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.fsx.getOntapStorageVirtualMachine({
 *     filters: [{
 *         name: "file-system-id",
 *         values: ["fs-12345678"],
 *     }],
 * });
 * ```
 */
export function getOntapStorageVirtualMachineOutput(args?: GetOntapStorageVirtualMachineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOntapStorageVirtualMachineResult> {
    return pulumi.output(args).apply((a: any) => getOntapStorageVirtualMachine(a, opts))
}

/**
 * A collection of arguments for invoking getOntapStorageVirtualMachine.
 */
export interface GetOntapStorageVirtualMachineOutputArgs {
    /**
     * Configuration block. Detailed below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.fsx.GetOntapStorageVirtualMachineFilterArgs>[]>;
    /**
     * Identifier of the storage virtual machine (e.g. `svm-12345678`).
     */
    id?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
