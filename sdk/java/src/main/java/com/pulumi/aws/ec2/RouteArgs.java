// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouteArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouteArgs Empty = new RouteArgs();

    /**
     * Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
     * 
     */
    @Import(name="carrierGatewayId")
    private @Nullable Output<String> carrierGatewayId;

    /**
     * @return Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
     * 
     */
    public Optional<Output<String>> carrierGatewayId() {
        return Optional.ofNullable(this.carrierGatewayId);
    }

    /**
     * The Amazon Resource Name (ARN) of a core network.
     * 
     */
    @Import(name="coreNetworkArn")
    private @Nullable Output<String> coreNetworkArn;

    /**
     * @return The Amazon Resource Name (ARN) of a core network.
     * 
     */
    public Optional<Output<String>> coreNetworkArn() {
        return Optional.ofNullable(this.coreNetworkArn);
    }

    /**
     * The destination CIDR block.
     * 
     */
    @Import(name="destinationCidrBlock")
    private @Nullable Output<String> destinationCidrBlock;

    /**
     * @return The destination CIDR block.
     * 
     */
    public Optional<Output<String>> destinationCidrBlock() {
        return Optional.ofNullable(this.destinationCidrBlock);
    }

    /**
     * The destination IPv6 CIDR block.
     * 
     */
    @Import(name="destinationIpv6CidrBlock")
    private @Nullable Output<String> destinationIpv6CidrBlock;

    /**
     * @return The destination IPv6 CIDR block.
     * 
     */
    public Optional<Output<String>> destinationIpv6CidrBlock() {
        return Optional.ofNullable(this.destinationIpv6CidrBlock);
    }

    /**
     * The ID of a managed prefix list destination.
     * 
     * One of the following target arguments must be supplied:
     * 
     */
    @Import(name="destinationPrefixListId")
    private @Nullable Output<String> destinationPrefixListId;

    /**
     * @return The ID of a managed prefix list destination.
     * 
     * One of the following target arguments must be supplied:
     * 
     */
    public Optional<Output<String>> destinationPrefixListId() {
        return Optional.ofNullable(this.destinationPrefixListId);
    }

    /**
     * Identifier of a VPC Egress Only Internet Gateway.
     * 
     */
    @Import(name="egressOnlyGatewayId")
    private @Nullable Output<String> egressOnlyGatewayId;

    /**
     * @return Identifier of a VPC Egress Only Internet Gateway.
     * 
     */
    public Optional<Output<String>> egressOnlyGatewayId() {
        return Optional.ofNullable(this.egressOnlyGatewayId);
    }

    /**
     * Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
     * 
     */
    @Import(name="gatewayId")
    private @Nullable Output<String> gatewayId;

    /**
     * @return Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
     * 
     */
    public Optional<Output<String>> gatewayId() {
        return Optional.ofNullable(this.gatewayId);
    }

    /**
     * Identifier of a Outpost local gateway.
     * 
     */
    @Import(name="localGatewayId")
    private @Nullable Output<String> localGatewayId;

    /**
     * @return Identifier of a Outpost local gateway.
     * 
     */
    public Optional<Output<String>> localGatewayId() {
        return Optional.ofNullable(this.localGatewayId);
    }

    /**
     * Identifier of a VPC NAT gateway.
     * 
     */
    @Import(name="natGatewayId")
    private @Nullable Output<String> natGatewayId;

    /**
     * @return Identifier of a VPC NAT gateway.
     * 
     */
    public Optional<Output<String>> natGatewayId() {
        return Optional.ofNullable(this.natGatewayId);
    }

    /**
     * Identifier of an EC2 network interface.
     * 
     */
    @Import(name="networkInterfaceId")
    private @Nullable Output<String> networkInterfaceId;

    /**
     * @return Identifier of an EC2 network interface.
     * 
     */
    public Optional<Output<String>> networkInterfaceId() {
        return Optional.ofNullable(this.networkInterfaceId);
    }

    /**
     * The ID of the routing table.
     * 
     * One of the following destination arguments must be supplied:
     * 
     */
    @Import(name="routeTableId", required=true)
    private Output<String> routeTableId;

    /**
     * @return The ID of the routing table.
     * 
     * One of the following destination arguments must be supplied:
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }

    /**
     * Identifier of an EC2 Transit Gateway.
     * 
     */
    @Import(name="transitGatewayId")
    private @Nullable Output<String> transitGatewayId;

    /**
     * @return Identifier of an EC2 Transit Gateway.
     * 
     */
    public Optional<Output<String>> transitGatewayId() {
        return Optional.ofNullable(this.transitGatewayId);
    }

    /**
     * Identifier of a VPC Endpoint.
     * 
     */
    @Import(name="vpcEndpointId")
    private @Nullable Output<String> vpcEndpointId;

    /**
     * @return Identifier of a VPC Endpoint.
     * 
     */
    public Optional<Output<String>> vpcEndpointId() {
        return Optional.ofNullable(this.vpcEndpointId);
    }

    /**
     * Identifier of a VPC peering connection.
     * 
     * Note that the default route, mapping the VPC&#39;s CIDR block to &#34;local&#34;, is created implicitly and cannot be specified.
     * 
     */
    @Import(name="vpcPeeringConnectionId")
    private @Nullable Output<String> vpcPeeringConnectionId;

    /**
     * @return Identifier of a VPC peering connection.
     * 
     * Note that the default route, mapping the VPC&#39;s CIDR block to &#34;local&#34;, is created implicitly and cannot be specified.
     * 
     */
    public Optional<Output<String>> vpcPeeringConnectionId() {
        return Optional.ofNullable(this.vpcPeeringConnectionId);
    }

    private RouteArgs() {}

    private RouteArgs(RouteArgs $) {
        this.carrierGatewayId = $.carrierGatewayId;
        this.coreNetworkArn = $.coreNetworkArn;
        this.destinationCidrBlock = $.destinationCidrBlock;
        this.destinationIpv6CidrBlock = $.destinationIpv6CidrBlock;
        this.destinationPrefixListId = $.destinationPrefixListId;
        this.egressOnlyGatewayId = $.egressOnlyGatewayId;
        this.gatewayId = $.gatewayId;
        this.localGatewayId = $.localGatewayId;
        this.natGatewayId = $.natGatewayId;
        this.networkInterfaceId = $.networkInterfaceId;
        this.routeTableId = $.routeTableId;
        this.transitGatewayId = $.transitGatewayId;
        this.vpcEndpointId = $.vpcEndpointId;
        this.vpcPeeringConnectionId = $.vpcPeeringConnectionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouteArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouteArgs $;

        public Builder() {
            $ = new RouteArgs();
        }

        public Builder(RouteArgs defaults) {
            $ = new RouteArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param carrierGatewayId Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
         * 
         * @return builder
         * 
         */
        public Builder carrierGatewayId(@Nullable Output<String> carrierGatewayId) {
            $.carrierGatewayId = carrierGatewayId;
            return this;
        }

        /**
         * @param carrierGatewayId Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
         * 
         * @return builder
         * 
         */
        public Builder carrierGatewayId(String carrierGatewayId) {
            return carrierGatewayId(Output.of(carrierGatewayId));
        }

        /**
         * @param coreNetworkArn The Amazon Resource Name (ARN) of a core network.
         * 
         * @return builder
         * 
         */
        public Builder coreNetworkArn(@Nullable Output<String> coreNetworkArn) {
            $.coreNetworkArn = coreNetworkArn;
            return this;
        }

        /**
         * @param coreNetworkArn The Amazon Resource Name (ARN) of a core network.
         * 
         * @return builder
         * 
         */
        public Builder coreNetworkArn(String coreNetworkArn) {
            return coreNetworkArn(Output.of(coreNetworkArn));
        }

        /**
         * @param destinationCidrBlock The destination CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidrBlock(@Nullable Output<String> destinationCidrBlock) {
            $.destinationCidrBlock = destinationCidrBlock;
            return this;
        }

        /**
         * @param destinationCidrBlock The destination CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidrBlock(String destinationCidrBlock) {
            return destinationCidrBlock(Output.of(destinationCidrBlock));
        }

        /**
         * @param destinationIpv6CidrBlock The destination IPv6 CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder destinationIpv6CidrBlock(@Nullable Output<String> destinationIpv6CidrBlock) {
            $.destinationIpv6CidrBlock = destinationIpv6CidrBlock;
            return this;
        }

        /**
         * @param destinationIpv6CidrBlock The destination IPv6 CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder destinationIpv6CidrBlock(String destinationIpv6CidrBlock) {
            return destinationIpv6CidrBlock(Output.of(destinationIpv6CidrBlock));
        }

        /**
         * @param destinationPrefixListId The ID of a managed prefix list destination.
         * 
         * One of the following target arguments must be supplied:
         * 
         * @return builder
         * 
         */
        public Builder destinationPrefixListId(@Nullable Output<String> destinationPrefixListId) {
            $.destinationPrefixListId = destinationPrefixListId;
            return this;
        }

        /**
         * @param destinationPrefixListId The ID of a managed prefix list destination.
         * 
         * One of the following target arguments must be supplied:
         * 
         * @return builder
         * 
         */
        public Builder destinationPrefixListId(String destinationPrefixListId) {
            return destinationPrefixListId(Output.of(destinationPrefixListId));
        }

        /**
         * @param egressOnlyGatewayId Identifier of a VPC Egress Only Internet Gateway.
         * 
         * @return builder
         * 
         */
        public Builder egressOnlyGatewayId(@Nullable Output<String> egressOnlyGatewayId) {
            $.egressOnlyGatewayId = egressOnlyGatewayId;
            return this;
        }

        /**
         * @param egressOnlyGatewayId Identifier of a VPC Egress Only Internet Gateway.
         * 
         * @return builder
         * 
         */
        public Builder egressOnlyGatewayId(String egressOnlyGatewayId) {
            return egressOnlyGatewayId(Output.of(egressOnlyGatewayId));
        }

        /**
         * @param gatewayId Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
         * 
         * @return builder
         * 
         */
        public Builder gatewayId(@Nullable Output<String> gatewayId) {
            $.gatewayId = gatewayId;
            return this;
        }

        /**
         * @param gatewayId Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
         * 
         * @return builder
         * 
         */
        public Builder gatewayId(String gatewayId) {
            return gatewayId(Output.of(gatewayId));
        }

        /**
         * @param localGatewayId Identifier of a Outpost local gateway.
         * 
         * @return builder
         * 
         */
        public Builder localGatewayId(@Nullable Output<String> localGatewayId) {
            $.localGatewayId = localGatewayId;
            return this;
        }

        /**
         * @param localGatewayId Identifier of a Outpost local gateway.
         * 
         * @return builder
         * 
         */
        public Builder localGatewayId(String localGatewayId) {
            return localGatewayId(Output.of(localGatewayId));
        }

        /**
         * @param natGatewayId Identifier of a VPC NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder natGatewayId(@Nullable Output<String> natGatewayId) {
            $.natGatewayId = natGatewayId;
            return this;
        }

        /**
         * @param natGatewayId Identifier of a VPC NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder natGatewayId(String natGatewayId) {
            return natGatewayId(Output.of(natGatewayId));
        }

        /**
         * @param networkInterfaceId Identifier of an EC2 network interface.
         * 
         * @return builder
         * 
         */
        public Builder networkInterfaceId(@Nullable Output<String> networkInterfaceId) {
            $.networkInterfaceId = networkInterfaceId;
            return this;
        }

        /**
         * @param networkInterfaceId Identifier of an EC2 network interface.
         * 
         * @return builder
         * 
         */
        public Builder networkInterfaceId(String networkInterfaceId) {
            return networkInterfaceId(Output.of(networkInterfaceId));
        }

        /**
         * @param routeTableId The ID of the routing table.
         * 
         * One of the following destination arguments must be supplied:
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(Output<String> routeTableId) {
            $.routeTableId = routeTableId;
            return this;
        }

        /**
         * @param routeTableId The ID of the routing table.
         * 
         * One of the following destination arguments must be supplied:
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(String routeTableId) {
            return routeTableId(Output.of(routeTableId));
        }

        /**
         * @param transitGatewayId Identifier of an EC2 Transit Gateway.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayId(@Nullable Output<String> transitGatewayId) {
            $.transitGatewayId = transitGatewayId;
            return this;
        }

        /**
         * @param transitGatewayId Identifier of an EC2 Transit Gateway.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayId(String transitGatewayId) {
            return transitGatewayId(Output.of(transitGatewayId));
        }

        /**
         * @param vpcEndpointId Identifier of a VPC Endpoint.
         * 
         * @return builder
         * 
         */
        public Builder vpcEndpointId(@Nullable Output<String> vpcEndpointId) {
            $.vpcEndpointId = vpcEndpointId;
            return this;
        }

        /**
         * @param vpcEndpointId Identifier of a VPC Endpoint.
         * 
         * @return builder
         * 
         */
        public Builder vpcEndpointId(String vpcEndpointId) {
            return vpcEndpointId(Output.of(vpcEndpointId));
        }

        /**
         * @param vpcPeeringConnectionId Identifier of a VPC peering connection.
         * 
         * Note that the default route, mapping the VPC&#39;s CIDR block to &#34;local&#34;, is created implicitly and cannot be specified.
         * 
         * @return builder
         * 
         */
        public Builder vpcPeeringConnectionId(@Nullable Output<String> vpcPeeringConnectionId) {
            $.vpcPeeringConnectionId = vpcPeeringConnectionId;
            return this;
        }

        /**
         * @param vpcPeeringConnectionId Identifier of a VPC peering connection.
         * 
         * Note that the default route, mapping the VPC&#39;s CIDR block to &#34;local&#34;, is created implicitly and cannot be specified.
         * 
         * @return builder
         * 
         */
        public Builder vpcPeeringConnectionId(String vpcPeeringConnectionId) {
            return vpcPeeringConnectionId(Output.of(vpcPeeringConnectionId));
        }

        public RouteArgs build() {
            $.routeTableId = Objects.requireNonNull($.routeTableId, "expected parameter 'routeTableId' to be non-null");
            return $;
        }
    }

}
