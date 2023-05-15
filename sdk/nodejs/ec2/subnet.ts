// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an VPC subnet resource.
 *
 * > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), subnets associated with Lambda Functions can take up to 45 minutes to successfully delete.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.ec2.Subnet("main", {
 *     vpcId: aws_vpc.main.id,
 *     cidrBlock: "10.0.1.0/24",
 *     tags: {
 *         Name: "Main",
 *     },
 * });
 * ```
 * ### Subnets In Secondary VPC CIDR Blocks
 *
 * When managing subnets in one of a VPC's secondary CIDR blocks created using a `aws.ec2.VpcIpv4CidrBlockAssociation`
 * resource, it is recommended to reference that resource's `vpcId` attribute to ensure correct dependency ordering.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const secondaryCidr = new aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", {
 *     vpcId: aws_vpc.main.id,
 *     cidrBlock: "172.2.0.0/16",
 * });
 * const inSecondaryCidr = new aws.ec2.Subnet("inSecondaryCidr", {
 *     vpcId: secondaryCidr.vpcId,
 *     cidrBlock: "172.2.0.0/24",
 * });
 * ```
 *
 * ## Import
 *
 * Subnets can be imported using the `subnet id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:ec2/subnet:Subnet public_subnet subnet-9d4a7b6c
 * ```
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetState, opts?: pulumi.CustomResourceOptions): Subnet {
        return new Subnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/subnet:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    /**
     * The ARN of the subnet.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     */
    public readonly assignIpv6AddressOnCreation!: pulumi.Output<boolean | undefined>;
    /**
     * AZ for the subnet.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availabilityZone` instead.
     */
    public readonly availabilityZoneId!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR block for the subnet.
     */
    public readonly cidrBlock!: pulumi.Output<string | undefined>;
    /**
     * The customer owned IPv4 address pool. Typically used with the `mapCustomerOwnedIpOnLaunch` argument. The `outpostArn` argument must be specified when configured.
     */
    public readonly customerOwnedIpv4Pool!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
     */
    public readonly enableDns64!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
     */
    public readonly enableLniAtDeviceIndex!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
     */
    public readonly enableResourceNameDnsARecordOnLaunch!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
     */
    public readonly enableResourceNameDnsAaaaRecordOnLaunch!: pulumi.Output<boolean | undefined>;
    /**
     * The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     */
    public readonly ipv6CidrBlock!: pulumi.Output<string | undefined>;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    public /*out*/ readonly ipv6CidrBlockAssociationId!: pulumi.Output<string>;
    /**
     * Indicates whether to create an IPv6-only subnet. Default: `false`.
     */
    public readonly ipv6Native!: pulumi.Output<boolean | undefined>;
    /**
     * Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customerOwnedIpv4Pool` and `outpostArn` arguments must be specified when set to `true`. Default is `false`.
     */
    public readonly mapCustomerOwnedIpOnLaunch!: pulumi.Output<boolean | undefined>;
    /**
     * Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     */
    public readonly mapPublicIpOnLaunch!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    public readonly outpostArn!: pulumi.Output<string | undefined>;
    /**
     * The ID of the AWS account that owns the subnet.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
     */
    public readonly privateDnsHostnameTypeOnLaunch!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetArgs | SubnetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubnetState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["assignIpv6AddressOnCreation"] = state ? state.assignIpv6AddressOnCreation : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["availabilityZoneId"] = state ? state.availabilityZoneId : undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["customerOwnedIpv4Pool"] = state ? state.customerOwnedIpv4Pool : undefined;
            resourceInputs["enableDns64"] = state ? state.enableDns64 : undefined;
            resourceInputs["enableLniAtDeviceIndex"] = state ? state.enableLniAtDeviceIndex : undefined;
            resourceInputs["enableResourceNameDnsARecordOnLaunch"] = state ? state.enableResourceNameDnsARecordOnLaunch : undefined;
            resourceInputs["enableResourceNameDnsAaaaRecordOnLaunch"] = state ? state.enableResourceNameDnsAaaaRecordOnLaunch : undefined;
            resourceInputs["ipv6CidrBlock"] = state ? state.ipv6CidrBlock : undefined;
            resourceInputs["ipv6CidrBlockAssociationId"] = state ? state.ipv6CidrBlockAssociationId : undefined;
            resourceInputs["ipv6Native"] = state ? state.ipv6Native : undefined;
            resourceInputs["mapCustomerOwnedIpOnLaunch"] = state ? state.mapCustomerOwnedIpOnLaunch : undefined;
            resourceInputs["mapPublicIpOnLaunch"] = state ? state.mapPublicIpOnLaunch : undefined;
            resourceInputs["outpostArn"] = state ? state.outpostArn : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["privateDnsHostnameTypeOnLaunch"] = state ? state.privateDnsHostnameTypeOnLaunch : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as SubnetArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["assignIpv6AddressOnCreation"] = args ? args.assignIpv6AddressOnCreation : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["availabilityZoneId"] = args ? args.availabilityZoneId : undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["customerOwnedIpv4Pool"] = args ? args.customerOwnedIpv4Pool : undefined;
            resourceInputs["enableDns64"] = args ? args.enableDns64 : undefined;
            resourceInputs["enableLniAtDeviceIndex"] = args ? args.enableLniAtDeviceIndex : undefined;
            resourceInputs["enableResourceNameDnsARecordOnLaunch"] = args ? args.enableResourceNameDnsARecordOnLaunch : undefined;
            resourceInputs["enableResourceNameDnsAaaaRecordOnLaunch"] = args ? args.enableResourceNameDnsAaaaRecordOnLaunch : undefined;
            resourceInputs["ipv6CidrBlock"] = args ? args.ipv6CidrBlock : undefined;
            resourceInputs["ipv6Native"] = args ? args.ipv6Native : undefined;
            resourceInputs["mapCustomerOwnedIpOnLaunch"] = args ? args.mapCustomerOwnedIpOnLaunch : undefined;
            resourceInputs["mapPublicIpOnLaunch"] = args ? args.mapPublicIpOnLaunch : undefined;
            resourceInputs["outpostArn"] = args ? args.outpostArn : undefined;
            resourceInputs["privateDnsHostnameTypeOnLaunch"] = args ? args.privateDnsHostnameTypeOnLaunch : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["ipv6CidrBlockAssociationId"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subnet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnet resources.
 */
export interface SubnetState {
    /**
     * The ARN of the subnet.
     */
    arn?: pulumi.Input<string>;
    /**
     * Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     */
    assignIpv6AddressOnCreation?: pulumi.Input<boolean>;
    /**
     * AZ for the subnet.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availabilityZone` instead.
     */
    availabilityZoneId?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR block for the subnet.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The customer owned IPv4 address pool. Typically used with the `mapCustomerOwnedIpOnLaunch` argument. The `outpostArn` argument must be specified when configured.
     */
    customerOwnedIpv4Pool?: pulumi.Input<string>;
    /**
     * Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
     */
    enableDns64?: pulumi.Input<boolean>;
    /**
     * Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
     */
    enableLniAtDeviceIndex?: pulumi.Input<number>;
    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
     */
    enableResourceNameDnsARecordOnLaunch?: pulumi.Input<boolean>;
    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
     */
    enableResourceNameDnsAaaaRecordOnLaunch?: pulumi.Input<boolean>;
    /**
     * The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     */
    ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    ipv6CidrBlockAssociationId?: pulumi.Input<string>;
    /**
     * Indicates whether to create an IPv6-only subnet. Default: `false`.
     */
    ipv6Native?: pulumi.Input<boolean>;
    /**
     * Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customerOwnedIpv4Pool` and `outpostArn` arguments must be specified when set to `true`. Default is `false`.
     */
    mapCustomerOwnedIpOnLaunch?: pulumi.Input<boolean>;
    /**
     * Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     */
    mapPublicIpOnLaunch?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    outpostArn?: pulumi.Input<string>;
    /**
     * The ID of the AWS account that owns the subnet.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
     */
    privateDnsHostnameTypeOnLaunch?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VPC ID.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    /**
     * Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     */
    assignIpv6AddressOnCreation?: pulumi.Input<boolean>;
    /**
     * AZ for the subnet.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availabilityZone` instead.
     */
    availabilityZoneId?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR block for the subnet.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The customer owned IPv4 address pool. Typically used with the `mapCustomerOwnedIpOnLaunch` argument. The `outpostArn` argument must be specified when configured.
     */
    customerOwnedIpv4Pool?: pulumi.Input<string>;
    /**
     * Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
     */
    enableDns64?: pulumi.Input<boolean>;
    /**
     * Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
     */
    enableLniAtDeviceIndex?: pulumi.Input<number>;
    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
     */
    enableResourceNameDnsARecordOnLaunch?: pulumi.Input<boolean>;
    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
     */
    enableResourceNameDnsAaaaRecordOnLaunch?: pulumi.Input<boolean>;
    /**
     * The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     */
    ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * Indicates whether to create an IPv6-only subnet. Default: `false`.
     */
    ipv6Native?: pulumi.Input<boolean>;
    /**
     * Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customerOwnedIpv4Pool` and `outpostArn` arguments must be specified when set to `true`. Default is `false`.
     */
    mapCustomerOwnedIpOnLaunch?: pulumi.Input<boolean>;
    /**
     * Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     */
    mapPublicIpOnLaunch?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    outpostArn?: pulumi.Input<string>;
    /**
     * The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
     */
    privateDnsHostnameTypeOnLaunch?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VPC ID.
     */
    vpcId: pulumi.Input<string>;
}
