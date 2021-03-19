// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an AWS Storage Gateway file, tape, or volume gateway in the provider region.
 *
 * > NOTE: The Storage Gateway API requires the gateway to be connected to properly return information after activation. If you are receiving `The specified gateway is not connected` errors during resource creation (gateway activation), ensure your gateway instance meets the [Storage Gateway requirements](https://docs.aws.amazon.com/storagegateway/latest/userguide/Requirements.html).
 *
 * ## Example Usage
 * ### File Gateway
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.storagegateway.Gateway("example", {
 *     gatewayIpAddress: "1.2.3.4",
 *     gatewayName: "example",
 *     gatewayTimezone: "GMT",
 *     gatewayType: "FILE_S3",
 * });
 * ```
 * ### Tape Gateway
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.storagegateway.Gateway("example", {
 *     gatewayIpAddress: "1.2.3.4",
 *     gatewayName: "example",
 *     gatewayTimezone: "GMT",
 *     gatewayType: "VTL",
 *     mediumChangerType: "AWS-Gateway-VTL",
 *     tapeDriveType: "IBM-ULT3580-TD5",
 * });
 * ```
 * ### Volume Gateway (Cached)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.storagegateway.Gateway("example", {
 *     gatewayIpAddress: "1.2.3.4",
 *     gatewayName: "example",
 *     gatewayTimezone: "GMT",
 *     gatewayType: "CACHED",
 * });
 * ```
 * ### Volume Gateway (Stored)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.storagegateway.Gateway("example", {
 *     gatewayIpAddress: "1.2.3.4",
 *     gatewayName: "example",
 *     gatewayTimezone: "GMT",
 *     gatewayType: "STORED",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_storagegateway_gateway` can be imported by using the gateway Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:storagegateway/gateway:Gateway example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678
 * ```
 *
 *  <<<<<<< HEAD Certain resource arguments, like `gateway_ip_address` do not have a Storage Gateway API method for reading the information after creation, either omit the argument from the provider configuration or use `ignoreChanges` to hide the difference. ======= Certain resource arguments, like `gateway_ip_address` do not have a Storage Gateway API method for reading the information after creation, either omit the argument from the provider configuration or use `ignore_changes` to hide the difference, e.g. terraform resource "aws_storagegateway_gateway" "example" {
 *
 * # ... other configuration ...
 *
 *  gateway_ip_address = aws_instance.sgw.private_ip
 *
 * # There is no Storage Gateway API for reading gateway_ip_address
 *
 *  lifecycle {
 *
 *  ignore_changes = ["gateway_ip_address"]
 *
 *  } } >>>>>>> v3.33.0
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayState, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:storagegateway/gateway:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
     */
    public readonly activationKey!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the gateway.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
     */
    public readonly averageDownloadRateLimitInBitsPerSec!: pulumi.Output<number | undefined>;
    /**
     * The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
     */
    public readonly averageUploadRateLimitInBitsPerSec!: pulumi.Output<number | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
     */
    public readonly cloudwatchLogGroupArn!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Amazon EC2 instance that was used to launch the gateway.
     */
    public /*out*/ readonly ec2InstanceId!: pulumi.Output<string>;
    /**
     * The type of endpoint for your gateway.
     */
    public /*out*/ readonly endpointType!: pulumi.Output<string>;
    /**
     * Identifier of the gateway.
     */
    public /*out*/ readonly gatewayId!: pulumi.Output<string>;
    /**
     * Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
     */
    public readonly gatewayIpAddress!: pulumi.Output<string>;
    /**
     * Name of the gateway.
     */
    public readonly gatewayName!: pulumi.Output<string>;
    /**
     * An array that contains descriptions of the gateway network interfaces. See Gateway Network Interface.
     */
    public /*out*/ readonly gatewayNetworkInterfaces!: pulumi.Output<outputs.storagegateway.GatewayGatewayNetworkInterface[]>;
    /**
     * Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
     */
    public readonly gatewayTimezone!: pulumi.Output<string>;
    /**
     * Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
     */
    public readonly gatewayType!: pulumi.Output<string | undefined>;
    /**
     * VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
     */
    public readonly gatewayVpcEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The type of hypervisor environment used by the host.
     */
    public /*out*/ readonly hostEnvironment!: pulumi.Output<string>;
    /**
     * Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
     */
    public readonly mediumChangerType!: pulumi.Output<string | undefined>;
    /**
     * Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
     */
    public readonly smbActiveDirectorySettings!: pulumi.Output<outputs.storagegateway.GatewaySmbActiveDirectorySettings | undefined>;
    /**
     * Specifies whether the shares on this gateway appear when listing shares.
     */
    public readonly smbFileShareVisibility!: pulumi.Output<boolean | undefined>;
    /**
     * Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
     */
    public readonly smbGuestPassword!: pulumi.Output<string | undefined>;
    /**
     * Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
     */
    public readonly smbSecurityStrategy!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
     */
    public readonly tapeDriveType!: pulumi.Output<string | undefined>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayArgs | GatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayState | undefined;
            inputs["activationKey"] = state ? state.activationKey : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["averageDownloadRateLimitInBitsPerSec"] = state ? state.averageDownloadRateLimitInBitsPerSec : undefined;
            inputs["averageUploadRateLimitInBitsPerSec"] = state ? state.averageUploadRateLimitInBitsPerSec : undefined;
            inputs["cloudwatchLogGroupArn"] = state ? state.cloudwatchLogGroupArn : undefined;
            inputs["ec2InstanceId"] = state ? state.ec2InstanceId : undefined;
            inputs["endpointType"] = state ? state.endpointType : undefined;
            inputs["gatewayId"] = state ? state.gatewayId : undefined;
            inputs["gatewayIpAddress"] = state ? state.gatewayIpAddress : undefined;
            inputs["gatewayName"] = state ? state.gatewayName : undefined;
            inputs["gatewayNetworkInterfaces"] = state ? state.gatewayNetworkInterfaces : undefined;
            inputs["gatewayTimezone"] = state ? state.gatewayTimezone : undefined;
            inputs["gatewayType"] = state ? state.gatewayType : undefined;
            inputs["gatewayVpcEndpoint"] = state ? state.gatewayVpcEndpoint : undefined;
            inputs["hostEnvironment"] = state ? state.hostEnvironment : undefined;
            inputs["mediumChangerType"] = state ? state.mediumChangerType : undefined;
            inputs["smbActiveDirectorySettings"] = state ? state.smbActiveDirectorySettings : undefined;
            inputs["smbFileShareVisibility"] = state ? state.smbFileShareVisibility : undefined;
            inputs["smbGuestPassword"] = state ? state.smbGuestPassword : undefined;
            inputs["smbSecurityStrategy"] = state ? state.smbSecurityStrategy : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tapeDriveType"] = state ? state.tapeDriveType : undefined;
        } else {
            const args = argsOrState as GatewayArgs | undefined;
            if ((!args || args.gatewayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayName'");
            }
            if ((!args || args.gatewayTimezone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayTimezone'");
            }
            inputs["activationKey"] = args ? args.activationKey : undefined;
            inputs["averageDownloadRateLimitInBitsPerSec"] = args ? args.averageDownloadRateLimitInBitsPerSec : undefined;
            inputs["averageUploadRateLimitInBitsPerSec"] = args ? args.averageUploadRateLimitInBitsPerSec : undefined;
            inputs["cloudwatchLogGroupArn"] = args ? args.cloudwatchLogGroupArn : undefined;
            inputs["gatewayIpAddress"] = args ? args.gatewayIpAddress : undefined;
            inputs["gatewayName"] = args ? args.gatewayName : undefined;
            inputs["gatewayTimezone"] = args ? args.gatewayTimezone : undefined;
            inputs["gatewayType"] = args ? args.gatewayType : undefined;
            inputs["gatewayVpcEndpoint"] = args ? args.gatewayVpcEndpoint : undefined;
            inputs["mediumChangerType"] = args ? args.mediumChangerType : undefined;
            inputs["smbActiveDirectorySettings"] = args ? args.smbActiveDirectorySettings : undefined;
            inputs["smbFileShareVisibility"] = args ? args.smbFileShareVisibility : undefined;
            inputs["smbGuestPassword"] = args ? args.smbGuestPassword : undefined;
            inputs["smbSecurityStrategy"] = args ? args.smbSecurityStrategy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tapeDriveType"] = args ? args.tapeDriveType : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ec2InstanceId"] = undefined /*out*/;
            inputs["endpointType"] = undefined /*out*/;
            inputs["gatewayId"] = undefined /*out*/;
            inputs["gatewayNetworkInterfaces"] = undefined /*out*/;
            inputs["hostEnvironment"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Gateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Gateway resources.
 */
export interface GatewayState {
    /**
     * Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
     */
    readonly activationKey?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the gateway.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
     */
    readonly averageDownloadRateLimitInBitsPerSec?: pulumi.Input<number>;
    /**
     * The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
     */
    readonly averageUploadRateLimitInBitsPerSec?: pulumi.Input<number>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
     */
    readonly cloudwatchLogGroupArn?: pulumi.Input<string>;
    /**
     * The ID of the Amazon EC2 instance that was used to launch the gateway.
     */
    readonly ec2InstanceId?: pulumi.Input<string>;
    /**
     * The type of endpoint for your gateway.
     */
    readonly endpointType?: pulumi.Input<string>;
    /**
     * Identifier of the gateway.
     */
    readonly gatewayId?: pulumi.Input<string>;
    /**
     * Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
     */
    readonly gatewayIpAddress?: pulumi.Input<string>;
    /**
     * Name of the gateway.
     */
    readonly gatewayName?: pulumi.Input<string>;
    /**
     * An array that contains descriptions of the gateway network interfaces. See Gateway Network Interface.
     */
    readonly gatewayNetworkInterfaces?: pulumi.Input<pulumi.Input<inputs.storagegateway.GatewayGatewayNetworkInterface>[]>;
    /**
     * Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
     */
    readonly gatewayTimezone?: pulumi.Input<string>;
    /**
     * Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
     */
    readonly gatewayType?: pulumi.Input<string>;
    /**
     * VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
     */
    readonly gatewayVpcEndpoint?: pulumi.Input<string>;
    /**
     * The type of hypervisor environment used by the host.
     */
    readonly hostEnvironment?: pulumi.Input<string>;
    /**
     * Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
     */
    readonly mediumChangerType?: pulumi.Input<string>;
    /**
     * Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
     */
    readonly smbActiveDirectorySettings?: pulumi.Input<inputs.storagegateway.GatewaySmbActiveDirectorySettings>;
    /**
     * Specifies whether the shares on this gateway appear when listing shares.
     */
    readonly smbFileShareVisibility?: pulumi.Input<boolean>;
    /**
     * Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
     */
    readonly smbGuestPassword?: pulumi.Input<string>;
    /**
     * Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
     */
    readonly smbSecurityStrategy?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
     */
    readonly tapeDriveType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
     */
    readonly activationKey?: pulumi.Input<string>;
    /**
     * The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
     */
    readonly averageDownloadRateLimitInBitsPerSec?: pulumi.Input<number>;
    /**
     * The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
     */
    readonly averageUploadRateLimitInBitsPerSec?: pulumi.Input<number>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
     */
    readonly cloudwatchLogGroupArn?: pulumi.Input<string>;
    /**
     * Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
     */
    readonly gatewayIpAddress?: pulumi.Input<string>;
    /**
     * Name of the gateway.
     */
    readonly gatewayName: pulumi.Input<string>;
    /**
     * Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
     */
    readonly gatewayTimezone: pulumi.Input<string>;
    /**
     * Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
     */
    readonly gatewayType?: pulumi.Input<string>;
    /**
     * VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
     */
    readonly gatewayVpcEndpoint?: pulumi.Input<string>;
    /**
     * Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
     */
    readonly mediumChangerType?: pulumi.Input<string>;
    /**
     * Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
     */
    readonly smbActiveDirectorySettings?: pulumi.Input<inputs.storagegateway.GatewaySmbActiveDirectorySettings>;
    /**
     * Specifies whether the shares on this gateway appear when listing shares.
     */
    readonly smbFileShareVisibility?: pulumi.Input<boolean>;
    /**
     * Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
     */
    readonly smbGuestPassword?: pulumi.Input<string>;
    /**
     * Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
     */
    readonly smbSecurityStrategy?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
     */
    readonly tapeDriveType?: pulumi.Input<string>;
}
