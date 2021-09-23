// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieve information about the AWS Direct Connect locations in the current AWS Region.
 * These are the locations that can be specified when configuring [`aws.directconnect.Connection`](https://www.terraform.io/docs/providers/aws/r/dx_connection.html) or [`aws.directconnect.LinkAggregationGroup`](https://www.terraform.io/docs/providers/aws/r/dx_lag.html) resources.
 *
 * > **Note:** This data source is different from the [`aws.directconnect.getLocation`](https://www.terraform.io/docs/providers/aws/d/dx_location.html) data source which retrieves information about a specific AWS Direct Connect location in the current AWS Region.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const available = pulumi.output(aws.directconnect.getLocations());
 * ```
 */
export function getLocations(opts?: pulumi.InvokeOptions): Promise<GetLocationsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:directconnect/getLocations:getLocations", {
    }, opts);
}

/**
 * A collection of values returned by getLocations.
 */
export interface GetLocationsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The code for the locations.
     */
    readonly locationCodes: string[];
}
