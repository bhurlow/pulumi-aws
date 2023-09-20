// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2transitgateway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TransitGatewayState extends com.pulumi.resources.ResourceArgs {

    public static final TransitGatewayState Empty = new TransitGatewayState();

    /**
     * Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
     * 
     * &gt; **NOTE:** Modifying `amazon_side_asn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazon_side_asn`.
     * 
     */
    @Import(name="amazonSideAsn")
    private @Nullable Output<Integer> amazonSideAsn;

    /**
     * @return Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
     * 
     * &gt; **NOTE:** Modifying `amazon_side_asn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazon_side_asn`.
     * 
     */
    public Optional<Output<Integer>> amazonSideAsn() {
        return Optional.ofNullable(this.amazonSideAsn);
    }

    /**
     * EC2 Transit Gateway Amazon Resource Name (ARN)
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return EC2 Transit Gateway Amazon Resource Name (ARN)
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Identifier of the default association route table
     * 
     */
    @Import(name="associationDefaultRouteTableId")
    private @Nullable Output<String> associationDefaultRouteTableId;

    /**
     * @return Identifier of the default association route table
     * 
     */
    public Optional<Output<String>> associationDefaultRouteTableId() {
        return Optional.ofNullable(this.associationDefaultRouteTableId);
    }

    /**
     * Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    @Import(name="autoAcceptSharedAttachments")
    private @Nullable Output<String> autoAcceptSharedAttachments;

    /**
     * @return Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    public Optional<Output<String>> autoAcceptSharedAttachments() {
        return Optional.ofNullable(this.autoAcceptSharedAttachments);
    }

    /**
     * Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    @Import(name="defaultRouteTableAssociation")
    private @Nullable Output<String> defaultRouteTableAssociation;

    /**
     * @return Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    public Optional<Output<String>> defaultRouteTableAssociation() {
        return Optional.ofNullable(this.defaultRouteTableAssociation);
    }

    /**
     * Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    @Import(name="defaultRouteTablePropagation")
    private @Nullable Output<String> defaultRouteTablePropagation;

    /**
     * @return Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    public Optional<Output<String>> defaultRouteTablePropagation() {
        return Optional.ofNullable(this.defaultRouteTablePropagation);
    }

    /**
     * Description of the EC2 Transit Gateway.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the EC2 Transit Gateway.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    @Import(name="dnsSupport")
    private @Nullable Output<String> dnsSupport;

    /**
     * @return Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    public Optional<Output<String>> dnsSupport() {
        return Optional.ofNullable(this.dnsSupport);
    }

    /**
     * Whether Multicast support is enabled. Required to use `ec2_transit_gateway_multicast_domain`. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    @Import(name="multicastSupport")
    private @Nullable Output<String> multicastSupport;

    /**
     * @return Whether Multicast support is enabled. Required to use `ec2_transit_gateway_multicast_domain`. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    public Optional<Output<String>> multicastSupport() {
        return Optional.ofNullable(this.multicastSupport);
    }

    /**
     * Identifier of the AWS account that owns the EC2 Transit Gateway
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<String> ownerId;

    /**
     * @return Identifier of the AWS account that owns the EC2 Transit Gateway
     * 
     */
    public Optional<Output<String>> ownerId() {
        return Optional.ofNullable(this.ownerId);
    }

    /**
     * Identifier of the default propagation route table
     * 
     */
    @Import(name="propagationDefaultRouteTableId")
    private @Nullable Output<String> propagationDefaultRouteTableId;

    /**
     * @return Identifier of the default propagation route table
     * 
     */
    public Optional<Output<String>> propagationDefaultRouteTableId() {
        return Optional.ofNullable(this.propagationDefaultRouteTableId);
    }

    /**
     * Key-value tags for the EC2 Transit Gateway. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value tags for the EC2 Transit Gateway. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
     * 
     */
    @Import(name="transitGatewayCidrBlocks")
    private @Nullable Output<List<String>> transitGatewayCidrBlocks;

    /**
     * @return One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
     * 
     */
    public Optional<Output<List<String>>> transitGatewayCidrBlocks() {
        return Optional.ofNullable(this.transitGatewayCidrBlocks);
    }

    /**
     * Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    @Import(name="vpnEcmpSupport")
    private @Nullable Output<String> vpnEcmpSupport;

    /**
     * @return Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
     * 
     */
    public Optional<Output<String>> vpnEcmpSupport() {
        return Optional.ofNullable(this.vpnEcmpSupport);
    }

    private TransitGatewayState() {}

    private TransitGatewayState(TransitGatewayState $) {
        this.amazonSideAsn = $.amazonSideAsn;
        this.arn = $.arn;
        this.associationDefaultRouteTableId = $.associationDefaultRouteTableId;
        this.autoAcceptSharedAttachments = $.autoAcceptSharedAttachments;
        this.defaultRouteTableAssociation = $.defaultRouteTableAssociation;
        this.defaultRouteTablePropagation = $.defaultRouteTablePropagation;
        this.description = $.description;
        this.dnsSupport = $.dnsSupport;
        this.multicastSupport = $.multicastSupport;
        this.ownerId = $.ownerId;
        this.propagationDefaultRouteTableId = $.propagationDefaultRouteTableId;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.transitGatewayCidrBlocks = $.transitGatewayCidrBlocks;
        this.vpnEcmpSupport = $.vpnEcmpSupport;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransitGatewayState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransitGatewayState $;

        public Builder() {
            $ = new TransitGatewayState();
        }

        public Builder(TransitGatewayState defaults) {
            $ = new TransitGatewayState(Objects.requireNonNull(defaults));
        }

        /**
         * @param amazonSideAsn Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
         * 
         * &gt; **NOTE:** Modifying `amazon_side_asn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazon_side_asn`.
         * 
         * @return builder
         * 
         */
        public Builder amazonSideAsn(@Nullable Output<Integer> amazonSideAsn) {
            $.amazonSideAsn = amazonSideAsn;
            return this;
        }

        /**
         * @param amazonSideAsn Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
         * 
         * &gt; **NOTE:** Modifying `amazon_side_asn` on a Transit Gateway with active BGP sessions is [not allowed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTransitGatewayOptions.html). You must first delete all Transit Gateway attachments that have BGP configured prior to modifying `amazon_side_asn`.
         * 
         * @return builder
         * 
         */
        public Builder amazonSideAsn(Integer amazonSideAsn) {
            return amazonSideAsn(Output.of(amazonSideAsn));
        }

        /**
         * @param arn EC2 Transit Gateway Amazon Resource Name (ARN)
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn EC2 Transit Gateway Amazon Resource Name (ARN)
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param associationDefaultRouteTableId Identifier of the default association route table
         * 
         * @return builder
         * 
         */
        public Builder associationDefaultRouteTableId(@Nullable Output<String> associationDefaultRouteTableId) {
            $.associationDefaultRouteTableId = associationDefaultRouteTableId;
            return this;
        }

        /**
         * @param associationDefaultRouteTableId Identifier of the default association route table
         * 
         * @return builder
         * 
         */
        public Builder associationDefaultRouteTableId(String associationDefaultRouteTableId) {
            return associationDefaultRouteTableId(Output.of(associationDefaultRouteTableId));
        }

        /**
         * @param autoAcceptSharedAttachments Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder autoAcceptSharedAttachments(@Nullable Output<String> autoAcceptSharedAttachments) {
            $.autoAcceptSharedAttachments = autoAcceptSharedAttachments;
            return this;
        }

        /**
         * @param autoAcceptSharedAttachments Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder autoAcceptSharedAttachments(String autoAcceptSharedAttachments) {
            return autoAcceptSharedAttachments(Output.of(autoAcceptSharedAttachments));
        }

        /**
         * @param defaultRouteTableAssociation Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder defaultRouteTableAssociation(@Nullable Output<String> defaultRouteTableAssociation) {
            $.defaultRouteTableAssociation = defaultRouteTableAssociation;
            return this;
        }

        /**
         * @param defaultRouteTableAssociation Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder defaultRouteTableAssociation(String defaultRouteTableAssociation) {
            return defaultRouteTableAssociation(Output.of(defaultRouteTableAssociation));
        }

        /**
         * @param defaultRouteTablePropagation Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder defaultRouteTablePropagation(@Nullable Output<String> defaultRouteTablePropagation) {
            $.defaultRouteTablePropagation = defaultRouteTablePropagation;
            return this;
        }

        /**
         * @param defaultRouteTablePropagation Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder defaultRouteTablePropagation(String defaultRouteTablePropagation) {
            return defaultRouteTablePropagation(Output.of(defaultRouteTablePropagation));
        }

        /**
         * @param description Description of the EC2 Transit Gateway.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the EC2 Transit Gateway.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dnsSupport Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder dnsSupport(@Nullable Output<String> dnsSupport) {
            $.dnsSupport = dnsSupport;
            return this;
        }

        /**
         * @param dnsSupport Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder dnsSupport(String dnsSupport) {
            return dnsSupport(Output.of(dnsSupport));
        }

        /**
         * @param multicastSupport Whether Multicast support is enabled. Required to use `ec2_transit_gateway_multicast_domain`. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder multicastSupport(@Nullable Output<String> multicastSupport) {
            $.multicastSupport = multicastSupport;
            return this;
        }

        /**
         * @param multicastSupport Whether Multicast support is enabled. Required to use `ec2_transit_gateway_multicast_domain`. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder multicastSupport(String multicastSupport) {
            return multicastSupport(Output.of(multicastSupport));
        }

        /**
         * @param ownerId Identifier of the AWS account that owns the EC2 Transit Gateway
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<String> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId Identifier of the AWS account that owns the EC2 Transit Gateway
         * 
         * @return builder
         * 
         */
        public Builder ownerId(String ownerId) {
            return ownerId(Output.of(ownerId));
        }

        /**
         * @param propagationDefaultRouteTableId Identifier of the default propagation route table
         * 
         * @return builder
         * 
         */
        public Builder propagationDefaultRouteTableId(@Nullable Output<String> propagationDefaultRouteTableId) {
            $.propagationDefaultRouteTableId = propagationDefaultRouteTableId;
            return this;
        }

        /**
         * @param propagationDefaultRouteTableId Identifier of the default propagation route table
         * 
         * @return builder
         * 
         */
        public Builder propagationDefaultRouteTableId(String propagationDefaultRouteTableId) {
            return propagationDefaultRouteTableId(Output.of(propagationDefaultRouteTableId));
        }

        /**
         * @param tags Key-value tags for the EC2 Transit Gateway. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value tags for the EC2 Transit Gateway. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param transitGatewayCidrBlocks One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayCidrBlocks(@Nullable Output<List<String>> transitGatewayCidrBlocks) {
            $.transitGatewayCidrBlocks = transitGatewayCidrBlocks;
            return this;
        }

        /**
         * @param transitGatewayCidrBlocks One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayCidrBlocks(List<String> transitGatewayCidrBlocks) {
            return transitGatewayCidrBlocks(Output.of(transitGatewayCidrBlocks));
        }

        /**
         * @param transitGatewayCidrBlocks One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayCidrBlocks(String... transitGatewayCidrBlocks) {
            return transitGatewayCidrBlocks(List.of(transitGatewayCidrBlocks));
        }

        /**
         * @param vpnEcmpSupport Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder vpnEcmpSupport(@Nullable Output<String> vpnEcmpSupport) {
            $.vpnEcmpSupport = vpnEcmpSupport;
            return this;
        }

        /**
         * @param vpnEcmpSupport Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
         * 
         * @return builder
         * 
         */
        public Builder vpnEcmpSupport(String vpnEcmpSupport) {
            return vpnEcmpSupport(Output.of(vpnEcmpSupport));
        }

        public TransitGatewayState build() {
            return $;
        }
    }

}
