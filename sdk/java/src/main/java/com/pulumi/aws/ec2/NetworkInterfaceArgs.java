// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.ec2.inputs.NetworkInterfaceAttachmentArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkInterfaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkInterfaceArgs Empty = new NetworkInterfaceArgs();

    /**
     * Configuration block to define the attachment of the ENI. See Attachment below for more details!
     * 
     */
    @Import(name="attachments")
    private @Nullable Output<List<NetworkInterfaceAttachmentArgs>> attachments;

    /**
     * @return Configuration block to define the attachment of the ENI. See Attachment below for more details!
     * 
     */
    public Optional<Output<List<NetworkInterfaceAttachmentArgs>>> attachments() {
        return Optional.ofNullable(this.attachments);
    }

    /**
     * Description for the network interface.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the network interface.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
     * 
     */
    @Import(name="interfaceType")
    private @Nullable Output<String> interfaceType;

    /**
     * @return Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
     * 
     */
    public Optional<Output<String>> interfaceType() {
        return Optional.ofNullable(this.interfaceType);
    }

    /**
     * Number of IPv4 prefixes that AWS automatically assigns to the network interface.
     * 
     */
    @Import(name="ipv4PrefixCount")
    private @Nullable Output<Integer> ipv4PrefixCount;

    /**
     * @return Number of IPv4 prefixes that AWS automatically assigns to the network interface.
     * 
     */
    public Optional<Output<Integer>> ipv4PrefixCount() {
        return Optional.ofNullable(this.ipv4PrefixCount);
    }

    /**
     * One or more IPv4 prefixes assigned to the network interface.
     * 
     */
    @Import(name="ipv4Prefixes")
    private @Nullable Output<List<String>> ipv4Prefixes;

    /**
     * @return One or more IPv4 prefixes assigned to the network interface.
     * 
     */
    public Optional<Output<List<String>>> ipv4Prefixes() {
        return Optional.ofNullable(this.ipv4Prefixes);
    }

    /**
     * Number of IPv6 addresses to assign to a network interface. You can&#39;t use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
     * 
     */
    @Import(name="ipv6AddressCount")
    private @Nullable Output<Integer> ipv6AddressCount;

    /**
     * @return Number of IPv6 addresses to assign to a network interface. You can&#39;t use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
     * 
     */
    public Optional<Output<Integer>> ipv6AddressCount() {
        return Optional.ofNullable(this.ipv6AddressCount);
    }

    /**
     * Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
     * 
     */
    @Import(name="ipv6AddressListEnabled")
    private @Nullable Output<Boolean> ipv6AddressListEnabled;

    /**
     * @return Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
     * 
     */
    public Optional<Output<Boolean>> ipv6AddressListEnabled() {
        return Optional.ofNullable(this.ipv6AddressListEnabled);
    }

    /**
     * List of private IPs to assign to the ENI in sequential order.
     * 
     */
    @Import(name="ipv6AddressLists")
    private @Nullable Output<List<String>> ipv6AddressLists;

    /**
     * @return List of private IPs to assign to the ENI in sequential order.
     * 
     */
    public Optional<Output<List<String>>> ipv6AddressLists() {
        return Optional.ofNullable(this.ipv6AddressLists);
    }

    /**
     * One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can&#39;t use this option if you&#39;re specifying `ipv6_address_count`.
     * 
     */
    @Import(name="ipv6Addresses")
    private @Nullable Output<List<String>> ipv6Addresses;

    /**
     * @return One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can&#39;t use this option if you&#39;re specifying `ipv6_address_count`.
     * 
     */
    public Optional<Output<List<String>>> ipv6Addresses() {
        return Optional.ofNullable(this.ipv6Addresses);
    }

    /**
     * Number of IPv6 prefixes that AWS automatically assigns to the network interface.
     * 
     */
    @Import(name="ipv6PrefixCount")
    private @Nullable Output<Integer> ipv6PrefixCount;

    /**
     * @return Number of IPv6 prefixes that AWS automatically assigns to the network interface.
     * 
     */
    public Optional<Output<Integer>> ipv6PrefixCount() {
        return Optional.ofNullable(this.ipv6PrefixCount);
    }

    /**
     * One or more IPv6 prefixes assigned to the network interface.
     * 
     */
    @Import(name="ipv6Prefixes")
    private @Nullable Output<List<String>> ipv6Prefixes;

    /**
     * @return One or more IPv6 prefixes assigned to the network interface.
     * 
     */
    public Optional<Output<List<String>>> ipv6Prefixes() {
        return Optional.ofNullable(this.ipv6Prefixes);
    }

    @Import(name="privateIp")
    private @Nullable Output<String> privateIp;

    public Optional<Output<String>> privateIp() {
        return Optional.ofNullable(this.privateIp);
    }

    /**
     * Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
     * 
     */
    @Import(name="privateIpListEnabled")
    private @Nullable Output<Boolean> privateIpListEnabled;

    /**
     * @return Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
     * 
     */
    public Optional<Output<Boolean>> privateIpListEnabled() {
        return Optional.ofNullable(this.privateIpListEnabled);
    }

    /**
     * List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
     * 
     */
    @Import(name="privateIpLists")
    private @Nullable Output<List<String>> privateIpLists;

    /**
     * @return List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
     * 
     */
    public Optional<Output<List<String>>> privateIpLists() {
        return Optional.ofNullable(this.privateIpLists);
    }

    /**
     * List of private IPs to assign to the ENI without regard to order.
     * 
     */
    @Import(name="privateIps")
    private @Nullable Output<List<String>> privateIps;

    /**
     * @return List of private IPs to assign to the ENI without regard to order.
     * 
     */
    public Optional<Output<List<String>>> privateIps() {
        return Optional.ofNullable(this.privateIps);
    }

    /**
     * Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
     * 
     */
    @Import(name="privateIpsCount")
    private @Nullable Output<Integer> privateIpsCount;

    /**
     * @return Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
     * 
     */
    public Optional<Output<Integer>> privateIpsCount() {
        return Optional.ofNullable(this.privateIpsCount);
    }

    /**
     * List of security group IDs to assign to the ENI.
     * 
     */
    @Import(name="securityGroups")
    private @Nullable Output<List<String>> securityGroups;

    /**
     * @return List of security group IDs to assign to the ENI.
     * 
     */
    public Optional<Output<List<String>>> securityGroups() {
        return Optional.ofNullable(this.securityGroups);
    }

    /**
     * Whether to enable source destination checking for the ENI. Default true.
     * 
     */
    @Import(name="sourceDestCheck")
    private @Nullable Output<Boolean> sourceDestCheck;

    /**
     * @return Whether to enable source destination checking for the ENI. Default true.
     * 
     */
    public Optional<Output<Boolean>> sourceDestCheck() {
        return Optional.ofNullable(this.sourceDestCheck);
    }

    /**
     * Subnet ID to create the ENI in.
     * 
     */
    @Import(name="subnetId", required=true)
    private Output<String> subnetId;

    /**
     * @return Subnet ID to create the ENI in.
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }

    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private NetworkInterfaceArgs() {}

    private NetworkInterfaceArgs(NetworkInterfaceArgs $) {
        this.attachments = $.attachments;
        this.description = $.description;
        this.interfaceType = $.interfaceType;
        this.ipv4PrefixCount = $.ipv4PrefixCount;
        this.ipv4Prefixes = $.ipv4Prefixes;
        this.ipv6AddressCount = $.ipv6AddressCount;
        this.ipv6AddressListEnabled = $.ipv6AddressListEnabled;
        this.ipv6AddressLists = $.ipv6AddressLists;
        this.ipv6Addresses = $.ipv6Addresses;
        this.ipv6PrefixCount = $.ipv6PrefixCount;
        this.ipv6Prefixes = $.ipv6Prefixes;
        this.privateIp = $.privateIp;
        this.privateIpListEnabled = $.privateIpListEnabled;
        this.privateIpLists = $.privateIpLists;
        this.privateIps = $.privateIps;
        this.privateIpsCount = $.privateIpsCount;
        this.securityGroups = $.securityGroups;
        this.sourceDestCheck = $.sourceDestCheck;
        this.subnetId = $.subnetId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkInterfaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkInterfaceArgs $;

        public Builder() {
            $ = new NetworkInterfaceArgs();
        }

        public Builder(NetworkInterfaceArgs defaults) {
            $ = new NetworkInterfaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attachments Configuration block to define the attachment of the ENI. See Attachment below for more details!
         * 
         * @return builder
         * 
         */
        public Builder attachments(@Nullable Output<List<NetworkInterfaceAttachmentArgs>> attachments) {
            $.attachments = attachments;
            return this;
        }

        /**
         * @param attachments Configuration block to define the attachment of the ENI. See Attachment below for more details!
         * 
         * @return builder
         * 
         */
        public Builder attachments(List<NetworkInterfaceAttachmentArgs> attachments) {
            return attachments(Output.of(attachments));
        }

        /**
         * @param attachments Configuration block to define the attachment of the ENI. See Attachment below for more details!
         * 
         * @return builder
         * 
         */
        public Builder attachments(NetworkInterfaceAttachmentArgs... attachments) {
            return attachments(List.of(attachments));
        }

        /**
         * @param description Description for the network interface.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the network interface.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param interfaceType Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
         * 
         * @return builder
         * 
         */
        public Builder interfaceType(@Nullable Output<String> interfaceType) {
            $.interfaceType = interfaceType;
            return this;
        }

        /**
         * @param interfaceType Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
         * 
         * @return builder
         * 
         */
        public Builder interfaceType(String interfaceType) {
            return interfaceType(Output.of(interfaceType));
        }

        /**
         * @param ipv4PrefixCount Number of IPv4 prefixes that AWS automatically assigns to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv4PrefixCount(@Nullable Output<Integer> ipv4PrefixCount) {
            $.ipv4PrefixCount = ipv4PrefixCount;
            return this;
        }

        /**
         * @param ipv4PrefixCount Number of IPv4 prefixes that AWS automatically assigns to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv4PrefixCount(Integer ipv4PrefixCount) {
            return ipv4PrefixCount(Output.of(ipv4PrefixCount));
        }

        /**
         * @param ipv4Prefixes One or more IPv4 prefixes assigned to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv4Prefixes(@Nullable Output<List<String>> ipv4Prefixes) {
            $.ipv4Prefixes = ipv4Prefixes;
            return this;
        }

        /**
         * @param ipv4Prefixes One or more IPv4 prefixes assigned to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv4Prefixes(List<String> ipv4Prefixes) {
            return ipv4Prefixes(Output.of(ipv4Prefixes));
        }

        /**
         * @param ipv4Prefixes One or more IPv4 prefixes assigned to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv4Prefixes(String... ipv4Prefixes) {
            return ipv4Prefixes(List.of(ipv4Prefixes));
        }

        /**
         * @param ipv6AddressCount Number of IPv6 addresses to assign to a network interface. You can&#39;t use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressCount(@Nullable Output<Integer> ipv6AddressCount) {
            $.ipv6AddressCount = ipv6AddressCount;
            return this;
        }

        /**
         * @param ipv6AddressCount Number of IPv6 addresses to assign to a network interface. You can&#39;t use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressCount(Integer ipv6AddressCount) {
            return ipv6AddressCount(Output.of(ipv6AddressCount));
        }

        /**
         * @param ipv6AddressListEnabled Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressListEnabled(@Nullable Output<Boolean> ipv6AddressListEnabled) {
            $.ipv6AddressListEnabled = ipv6AddressListEnabled;
            return this;
        }

        /**
         * @param ipv6AddressListEnabled Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressListEnabled(Boolean ipv6AddressListEnabled) {
            return ipv6AddressListEnabled(Output.of(ipv6AddressListEnabled));
        }

        /**
         * @param ipv6AddressLists List of private IPs to assign to the ENI in sequential order.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressLists(@Nullable Output<List<String>> ipv6AddressLists) {
            $.ipv6AddressLists = ipv6AddressLists;
            return this;
        }

        /**
         * @param ipv6AddressLists List of private IPs to assign to the ENI in sequential order.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressLists(List<String> ipv6AddressLists) {
            return ipv6AddressLists(Output.of(ipv6AddressLists));
        }

        /**
         * @param ipv6AddressLists List of private IPs to assign to the ENI in sequential order.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressLists(String... ipv6AddressLists) {
            return ipv6AddressLists(List.of(ipv6AddressLists));
        }

        /**
         * @param ipv6Addresses One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can&#39;t use this option if you&#39;re specifying `ipv6_address_count`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Addresses(@Nullable Output<List<String>> ipv6Addresses) {
            $.ipv6Addresses = ipv6Addresses;
            return this;
        }

        /**
         * @param ipv6Addresses One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can&#39;t use this option if you&#39;re specifying `ipv6_address_count`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Addresses(List<String> ipv6Addresses) {
            return ipv6Addresses(Output.of(ipv6Addresses));
        }

        /**
         * @param ipv6Addresses One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can&#39;t use this option if you&#39;re specifying `ipv6_address_count`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Addresses(String... ipv6Addresses) {
            return ipv6Addresses(List.of(ipv6Addresses));
        }

        /**
         * @param ipv6PrefixCount Number of IPv6 prefixes that AWS automatically assigns to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv6PrefixCount(@Nullable Output<Integer> ipv6PrefixCount) {
            $.ipv6PrefixCount = ipv6PrefixCount;
            return this;
        }

        /**
         * @param ipv6PrefixCount Number of IPv6 prefixes that AWS automatically assigns to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv6PrefixCount(Integer ipv6PrefixCount) {
            return ipv6PrefixCount(Output.of(ipv6PrefixCount));
        }

        /**
         * @param ipv6Prefixes One or more IPv6 prefixes assigned to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Prefixes(@Nullable Output<List<String>> ipv6Prefixes) {
            $.ipv6Prefixes = ipv6Prefixes;
            return this;
        }

        /**
         * @param ipv6Prefixes One or more IPv6 prefixes assigned to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Prefixes(List<String> ipv6Prefixes) {
            return ipv6Prefixes(Output.of(ipv6Prefixes));
        }

        /**
         * @param ipv6Prefixes One or more IPv6 prefixes assigned to the network interface.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Prefixes(String... ipv6Prefixes) {
            return ipv6Prefixes(List.of(ipv6Prefixes));
        }

        public Builder privateIp(@Nullable Output<String> privateIp) {
            $.privateIp = privateIp;
            return this;
        }

        public Builder privateIp(String privateIp) {
            return privateIp(Output.of(privateIp));
        }

        /**
         * @param privateIpListEnabled Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
         * 
         * @return builder
         * 
         */
        public Builder privateIpListEnabled(@Nullable Output<Boolean> privateIpListEnabled) {
            $.privateIpListEnabled = privateIpListEnabled;
            return this;
        }

        /**
         * @param privateIpListEnabled Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
         * 
         * @return builder
         * 
         */
        public Builder privateIpListEnabled(Boolean privateIpListEnabled) {
            return privateIpListEnabled(Output.of(privateIpListEnabled));
        }

        /**
         * @param privateIpLists List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
         * 
         * @return builder
         * 
         */
        public Builder privateIpLists(@Nullable Output<List<String>> privateIpLists) {
            $.privateIpLists = privateIpLists;
            return this;
        }

        /**
         * @param privateIpLists List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
         * 
         * @return builder
         * 
         */
        public Builder privateIpLists(List<String> privateIpLists) {
            return privateIpLists(Output.of(privateIpLists));
        }

        /**
         * @param privateIpLists List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
         * 
         * @return builder
         * 
         */
        public Builder privateIpLists(String... privateIpLists) {
            return privateIpLists(List.of(privateIpLists));
        }

        /**
         * @param privateIps List of private IPs to assign to the ENI without regard to order.
         * 
         * @return builder
         * 
         */
        public Builder privateIps(@Nullable Output<List<String>> privateIps) {
            $.privateIps = privateIps;
            return this;
        }

        /**
         * @param privateIps List of private IPs to assign to the ENI without regard to order.
         * 
         * @return builder
         * 
         */
        public Builder privateIps(List<String> privateIps) {
            return privateIps(Output.of(privateIps));
        }

        /**
         * @param privateIps List of private IPs to assign to the ENI without regard to order.
         * 
         * @return builder
         * 
         */
        public Builder privateIps(String... privateIps) {
            return privateIps(List.of(privateIps));
        }

        /**
         * @param privateIpsCount Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
         * 
         * @return builder
         * 
         */
        public Builder privateIpsCount(@Nullable Output<Integer> privateIpsCount) {
            $.privateIpsCount = privateIpsCount;
            return this;
        }

        /**
         * @param privateIpsCount Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
         * 
         * @return builder
         * 
         */
        public Builder privateIpsCount(Integer privateIpsCount) {
            return privateIpsCount(Output.of(privateIpsCount));
        }

        /**
         * @param securityGroups List of security group IDs to assign to the ENI.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(@Nullable Output<List<String>> securityGroups) {
            $.securityGroups = securityGroups;
            return this;
        }

        /**
         * @param securityGroups List of security group IDs to assign to the ENI.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(List<String> securityGroups) {
            return securityGroups(Output.of(securityGroups));
        }

        /**
         * @param securityGroups List of security group IDs to assign to the ENI.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(String... securityGroups) {
            return securityGroups(List.of(securityGroups));
        }

        /**
         * @param sourceDestCheck Whether to enable source destination checking for the ENI. Default true.
         * 
         * @return builder
         * 
         */
        public Builder sourceDestCheck(@Nullable Output<Boolean> sourceDestCheck) {
            $.sourceDestCheck = sourceDestCheck;
            return this;
        }

        /**
         * @param sourceDestCheck Whether to enable source destination checking for the ENI. Default true.
         * 
         * @return builder
         * 
         */
        public Builder sourceDestCheck(Boolean sourceDestCheck) {
            return sourceDestCheck(Output.of(sourceDestCheck));
        }

        /**
         * @param subnetId Subnet ID to create the ENI in.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId Subnet ID to create the ENI in.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public NetworkInterfaceArgs build() {
            $.subnetId = Objects.requireNonNull($.subnetId, "expected parameter 'subnetId' to be non-null");
            return $;
        }
    }

}
