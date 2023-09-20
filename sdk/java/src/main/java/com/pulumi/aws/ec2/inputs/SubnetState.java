// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SubnetState extends com.pulumi.resources.ResourceArgs {

    public static final SubnetState Empty = new SubnetState();

    /**
     * The ARN of the subnet.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the subnet.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     * 
     */
    @Import(name="assignIpv6AddressOnCreation")
    private @Nullable Output<Boolean> assignIpv6AddressOnCreation;

    /**
     * @return Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     * 
     */
    public Optional<Output<Boolean>> assignIpv6AddressOnCreation() {
        return Optional.ofNullable(this.assignIpv6AddressOnCreation);
    }

    /**
     * AZ for the subnet.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return AZ for the subnet.
     * 
     */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
     * 
     */
    @Import(name="availabilityZoneId")
    private @Nullable Output<String> availabilityZoneId;

    /**
     * @return AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
     * 
     */
    public Optional<Output<String>> availabilityZoneId() {
        return Optional.ofNullable(this.availabilityZoneId);
    }

    /**
     * The IPv4 CIDR block for the subnet.
     * 
     */
    @Import(name="cidrBlock")
    private @Nullable Output<String> cidrBlock;

    /**
     * @return The IPv4 CIDR block for the subnet.
     * 
     */
    public Optional<Output<String>> cidrBlock() {
        return Optional.ofNullable(this.cidrBlock);
    }

    /**
     * The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
     * 
     */
    @Import(name="customerOwnedIpv4Pool")
    private @Nullable Output<String> customerOwnedIpv4Pool;

    /**
     * @return The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
     * 
     */
    public Optional<Output<String>> customerOwnedIpv4Pool() {
        return Optional.ofNullable(this.customerOwnedIpv4Pool);
    }

    /**
     * Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
     * 
     */
    @Import(name="enableDns64")
    private @Nullable Output<Boolean> enableDns64;

    /**
     * @return Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> enableDns64() {
        return Optional.ofNullable(this.enableDns64);
    }

    /**
     * Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
     * 
     */
    @Import(name="enableLniAtDeviceIndex")
    private @Nullable Output<Integer> enableLniAtDeviceIndex;

    /**
     * @return Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
     * 
     */
    public Optional<Output<Integer>> enableLniAtDeviceIndex() {
        return Optional.ofNullable(this.enableLniAtDeviceIndex);
    }

    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
     * 
     */
    @Import(name="enableResourceNameDnsARecordOnLaunch")
    private @Nullable Output<Boolean> enableResourceNameDnsARecordOnLaunch;

    /**
     * @return Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> enableResourceNameDnsARecordOnLaunch() {
        return Optional.ofNullable(this.enableResourceNameDnsARecordOnLaunch);
    }

    /**
     * Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
     * 
     */
    @Import(name="enableResourceNameDnsAaaaRecordOnLaunch")
    private @Nullable Output<Boolean> enableResourceNameDnsAaaaRecordOnLaunch;

    /**
     * @return Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> enableResourceNameDnsAaaaRecordOnLaunch() {
        return Optional.ofNullable(this.enableResourceNameDnsAaaaRecordOnLaunch);
    }

    /**
     * The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     * 
     */
    @Import(name="ipv6CidrBlock")
    private @Nullable Output<String> ipv6CidrBlock;

    /**
     * @return The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     * 
     */
    public Optional<Output<String>> ipv6CidrBlock() {
        return Optional.ofNullable(this.ipv6CidrBlock);
    }

    /**
     * The association ID for the IPv6 CIDR block.
     * 
     */
    @Import(name="ipv6CidrBlockAssociationId")
    private @Nullable Output<String> ipv6CidrBlockAssociationId;

    /**
     * @return The association ID for the IPv6 CIDR block.
     * 
     */
    public Optional<Output<String>> ipv6CidrBlockAssociationId() {
        return Optional.ofNullable(this.ipv6CidrBlockAssociationId);
    }

    /**
     * Indicates whether to create an IPv6-only subnet. Default: `false`.
     * 
     */
    @Import(name="ipv6Native")
    private @Nullable Output<Boolean> ipv6Native;

    /**
     * @return Indicates whether to create an IPv6-only subnet. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> ipv6Native() {
        return Optional.ofNullable(this.ipv6Native);
    }

    /**
     * Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
     * 
     */
    @Import(name="mapCustomerOwnedIpOnLaunch")
    private @Nullable Output<Boolean> mapCustomerOwnedIpOnLaunch;

    /**
     * @return Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> mapCustomerOwnedIpOnLaunch() {
        return Optional.ofNullable(this.mapCustomerOwnedIpOnLaunch);
    }

    /**
     * Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     * 
     */
    @Import(name="mapPublicIpOnLaunch")
    private @Nullable Output<Boolean> mapPublicIpOnLaunch;

    /**
     * @return Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> mapPublicIpOnLaunch() {
        return Optional.ofNullable(this.mapPublicIpOnLaunch);
    }

    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     * 
     */
    @Import(name="outpostArn")
    private @Nullable Output<String> outpostArn;

    /**
     * @return The Amazon Resource Name (ARN) of the Outpost.
     * 
     */
    public Optional<Output<String>> outpostArn() {
        return Optional.ofNullable(this.outpostArn);
    }

    /**
     * The ID of the AWS account that owns the subnet.
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<String> ownerId;

    /**
     * @return The ID of the AWS account that owns the subnet.
     * 
     */
    public Optional<Output<String>> ownerId() {
        return Optional.ofNullable(this.ownerId);
    }

    /**
     * The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
     * 
     */
    @Import(name="privateDnsHostnameTypeOnLaunch")
    private @Nullable Output<String> privateDnsHostnameTypeOnLaunch;

    /**
     * @return The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
     * 
     */
    public Optional<Output<String>> privateDnsHostnameTypeOnLaunch() {
        return Optional.ofNullable(this.privateDnsHostnameTypeOnLaunch);
    }

    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The VPC ID.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The VPC ID.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private SubnetState() {}

    private SubnetState(SubnetState $) {
        this.arn = $.arn;
        this.assignIpv6AddressOnCreation = $.assignIpv6AddressOnCreation;
        this.availabilityZone = $.availabilityZone;
        this.availabilityZoneId = $.availabilityZoneId;
        this.cidrBlock = $.cidrBlock;
        this.customerOwnedIpv4Pool = $.customerOwnedIpv4Pool;
        this.enableDns64 = $.enableDns64;
        this.enableLniAtDeviceIndex = $.enableLniAtDeviceIndex;
        this.enableResourceNameDnsARecordOnLaunch = $.enableResourceNameDnsARecordOnLaunch;
        this.enableResourceNameDnsAaaaRecordOnLaunch = $.enableResourceNameDnsAaaaRecordOnLaunch;
        this.ipv6CidrBlock = $.ipv6CidrBlock;
        this.ipv6CidrBlockAssociationId = $.ipv6CidrBlockAssociationId;
        this.ipv6Native = $.ipv6Native;
        this.mapCustomerOwnedIpOnLaunch = $.mapCustomerOwnedIpOnLaunch;
        this.mapPublicIpOnLaunch = $.mapPublicIpOnLaunch;
        this.outpostArn = $.outpostArn;
        this.ownerId = $.ownerId;
        this.privateDnsHostnameTypeOnLaunch = $.privateDnsHostnameTypeOnLaunch;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubnetState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubnetState $;

        public Builder() {
            $ = new SubnetState();
        }

        public Builder(SubnetState defaults) {
            $ = new SubnetState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the subnet.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the subnet.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param assignIpv6AddressOnCreation Specify true to indicate
         * that network interfaces created in the specified subnet should be
         * assigned an IPv6 address. Default is `false`
         * 
         * @return builder
         * 
         */
        public Builder assignIpv6AddressOnCreation(@Nullable Output<Boolean> assignIpv6AddressOnCreation) {
            $.assignIpv6AddressOnCreation = assignIpv6AddressOnCreation;
            return this;
        }

        /**
         * @param assignIpv6AddressOnCreation Specify true to indicate
         * that network interfaces created in the specified subnet should be
         * assigned an IPv6 address. Default is `false`
         * 
         * @return builder
         * 
         */
        public Builder assignIpv6AddressOnCreation(Boolean assignIpv6AddressOnCreation) {
            return assignIpv6AddressOnCreation(Output.of(assignIpv6AddressOnCreation));
        }

        /**
         * @param availabilityZone AZ for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone AZ for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param availabilityZoneId AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZoneId(@Nullable Output<String> availabilityZoneId) {
            $.availabilityZoneId = availabilityZoneId;
            return this;
        }

        /**
         * @param availabilityZoneId AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZoneId(String availabilityZoneId) {
            return availabilityZoneId(Output.of(availabilityZoneId));
        }

        /**
         * @param cidrBlock The IPv4 CIDR block for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(@Nullable Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock The IPv4 CIDR block for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        /**
         * @param customerOwnedIpv4Pool The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
         * 
         * @return builder
         * 
         */
        public Builder customerOwnedIpv4Pool(@Nullable Output<String> customerOwnedIpv4Pool) {
            $.customerOwnedIpv4Pool = customerOwnedIpv4Pool;
            return this;
        }

        /**
         * @param customerOwnedIpv4Pool The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
         * 
         * @return builder
         * 
         */
        public Builder customerOwnedIpv4Pool(String customerOwnedIpv4Pool) {
            return customerOwnedIpv4Pool(Output.of(customerOwnedIpv4Pool));
        }

        /**
         * @param enableDns64 Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableDns64(@Nullable Output<Boolean> enableDns64) {
            $.enableDns64 = enableDns64;
            return this;
        }

        /**
         * @param enableDns64 Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableDns64(Boolean enableDns64) {
            return enableDns64(Output.of(enableDns64));
        }

        /**
         * @param enableLniAtDeviceIndex Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
         * 
         * @return builder
         * 
         */
        public Builder enableLniAtDeviceIndex(@Nullable Output<Integer> enableLniAtDeviceIndex) {
            $.enableLniAtDeviceIndex = enableLniAtDeviceIndex;
            return this;
        }

        /**
         * @param enableLniAtDeviceIndex Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
         * 
         * @return builder
         * 
         */
        public Builder enableLniAtDeviceIndex(Integer enableLniAtDeviceIndex) {
            return enableLniAtDeviceIndex(Output.of(enableLniAtDeviceIndex));
        }

        /**
         * @param enableResourceNameDnsARecordOnLaunch Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableResourceNameDnsARecordOnLaunch(@Nullable Output<Boolean> enableResourceNameDnsARecordOnLaunch) {
            $.enableResourceNameDnsARecordOnLaunch = enableResourceNameDnsARecordOnLaunch;
            return this;
        }

        /**
         * @param enableResourceNameDnsARecordOnLaunch Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableResourceNameDnsARecordOnLaunch(Boolean enableResourceNameDnsARecordOnLaunch) {
            return enableResourceNameDnsARecordOnLaunch(Output.of(enableResourceNameDnsARecordOnLaunch));
        }

        /**
         * @param enableResourceNameDnsAaaaRecordOnLaunch Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableResourceNameDnsAaaaRecordOnLaunch(@Nullable Output<Boolean> enableResourceNameDnsAaaaRecordOnLaunch) {
            $.enableResourceNameDnsAaaaRecordOnLaunch = enableResourceNameDnsAaaaRecordOnLaunch;
            return this;
        }

        /**
         * @param enableResourceNameDnsAaaaRecordOnLaunch Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableResourceNameDnsAaaaRecordOnLaunch(Boolean enableResourceNameDnsAaaaRecordOnLaunch) {
            return enableResourceNameDnsAaaaRecordOnLaunch(Output.of(enableResourceNameDnsAaaaRecordOnLaunch));
        }

        /**
         * @param ipv6CidrBlock The IPv6 network range for the subnet,
         * in CIDR notation. The subnet size must use a /64 prefix length.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlock(@Nullable Output<String> ipv6CidrBlock) {
            $.ipv6CidrBlock = ipv6CidrBlock;
            return this;
        }

        /**
         * @param ipv6CidrBlock The IPv6 network range for the subnet,
         * in CIDR notation. The subnet size must use a /64 prefix length.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlock(String ipv6CidrBlock) {
            return ipv6CidrBlock(Output.of(ipv6CidrBlock));
        }

        /**
         * @param ipv6CidrBlockAssociationId The association ID for the IPv6 CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlockAssociationId(@Nullable Output<String> ipv6CidrBlockAssociationId) {
            $.ipv6CidrBlockAssociationId = ipv6CidrBlockAssociationId;
            return this;
        }

        /**
         * @param ipv6CidrBlockAssociationId The association ID for the IPv6 CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlockAssociationId(String ipv6CidrBlockAssociationId) {
            return ipv6CidrBlockAssociationId(Output.of(ipv6CidrBlockAssociationId));
        }

        /**
         * @param ipv6Native Indicates whether to create an IPv6-only subnet. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Native(@Nullable Output<Boolean> ipv6Native) {
            $.ipv6Native = ipv6Native;
            return this;
        }

        /**
         * @param ipv6Native Indicates whether to create an IPv6-only subnet. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Native(Boolean ipv6Native) {
            return ipv6Native(Output.of(ipv6Native));
        }

        /**
         * @param mapCustomerOwnedIpOnLaunch Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder mapCustomerOwnedIpOnLaunch(@Nullable Output<Boolean> mapCustomerOwnedIpOnLaunch) {
            $.mapCustomerOwnedIpOnLaunch = mapCustomerOwnedIpOnLaunch;
            return this;
        }

        /**
         * @param mapCustomerOwnedIpOnLaunch Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder mapCustomerOwnedIpOnLaunch(Boolean mapCustomerOwnedIpOnLaunch) {
            return mapCustomerOwnedIpOnLaunch(Output.of(mapCustomerOwnedIpOnLaunch));
        }

        /**
         * @param mapPublicIpOnLaunch Specify true to indicate
         * that instances launched into the subnet should be assigned
         * a public IP address. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder mapPublicIpOnLaunch(@Nullable Output<Boolean> mapPublicIpOnLaunch) {
            $.mapPublicIpOnLaunch = mapPublicIpOnLaunch;
            return this;
        }

        /**
         * @param mapPublicIpOnLaunch Specify true to indicate
         * that instances launched into the subnet should be assigned
         * a public IP address. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder mapPublicIpOnLaunch(Boolean mapPublicIpOnLaunch) {
            return mapPublicIpOnLaunch(Output.of(mapPublicIpOnLaunch));
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the Outpost.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(@Nullable Output<String> outpostArn) {
            $.outpostArn = outpostArn;
            return this;
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the Outpost.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(String outpostArn) {
            return outpostArn(Output.of(outpostArn));
        }

        /**
         * @param ownerId The ID of the AWS account that owns the subnet.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<String> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId The ID of the AWS account that owns the subnet.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(String ownerId) {
            return ownerId(Output.of(ownerId));
        }

        /**
         * @param privateDnsHostnameTypeOnLaunch The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
         * 
         * @return builder
         * 
         */
        public Builder privateDnsHostnameTypeOnLaunch(@Nullable Output<String> privateDnsHostnameTypeOnLaunch) {
            $.privateDnsHostnameTypeOnLaunch = privateDnsHostnameTypeOnLaunch;
            return this;
        }

        /**
         * @param privateDnsHostnameTypeOnLaunch The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
         * 
         * @return builder
         * 
         */
        public Builder privateDnsHostnameTypeOnLaunch(String privateDnsHostnameTypeOnLaunch) {
            return privateDnsHostnameTypeOnLaunch(Output.of(privateDnsHostnameTypeOnLaunch));
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public SubnetState build() {
            return $;
        }
    }

}
