// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Retrieve information about an EKS add-on.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.eks.getAddon({
 *     addonName: "vpc-cni",
 *     clusterName: aws_eks_cluster.example.name,
 * });
 * export const eksAddonOutputs = aws_eks_addon.example;
 * ```
 */
export function getAddon(args: GetAddonArgs, opts?: pulumi.InvokeOptions): Promise<GetAddonResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:eks/getAddon:getAddon", {
        "addonName": args.addonName,
        "clusterName": args.clusterName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddon.
 */
export interface GetAddonArgs {
    /**
     * Name of the EKS add-on. The name must match one of
     * the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
     */
    readonly addonName: string;
    /**
     * Name of the EKS Cluster.
     */
    readonly clusterName: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getAddon.
 */
export interface GetAddonResult {
    readonly addonName: string;
    /**
     * The version of EKS add-on.
     */
    readonly addonVersion: string;
    /**
     * Amazon Resource Name (ARN) of the EKS add-on.
     */
    readonly arn: string;
    readonly clusterName: string;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was created.
     */
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was updated.
     */
    readonly modifiedAt: string;
    /**
     * ARN of IAM role used for EKS add-on. If value is empty -
     * then add-on uses the IAM role assigned to the EKS Cluster node.
     */
    readonly serviceAccountRoleArn: string;
    readonly tags: {[key: string]: string};
}
