// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LocalGatewayRouteTableVpcAssociationArgs extends com.pulumi.resources.ResourceArgs {

    public static final LocalGatewayRouteTableVpcAssociationArgs Empty = new LocalGatewayRouteTableVpcAssociationArgs();

    /**
     * Identifier of EC2 Local Gateway Route Table.
     * 
     */
    @Import(name="localGatewayRouteTableId", required=true)
    private Output<String> localGatewayRouteTableId;

    /**
     * @return Identifier of EC2 Local Gateway Route Table.
     * 
     */
    public Output<String> localGatewayRouteTableId() {
        return this.localGatewayRouteTableId;
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Identifier of EC2 VPC.
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return Identifier of EC2 VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    private LocalGatewayRouteTableVpcAssociationArgs() {}

    private LocalGatewayRouteTableVpcAssociationArgs(LocalGatewayRouteTableVpcAssociationArgs $) {
        this.localGatewayRouteTableId = $.localGatewayRouteTableId;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LocalGatewayRouteTableVpcAssociationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LocalGatewayRouteTableVpcAssociationArgs $;

        public Builder() {
            $ = new LocalGatewayRouteTableVpcAssociationArgs();
        }

        public Builder(LocalGatewayRouteTableVpcAssociationArgs defaults) {
            $ = new LocalGatewayRouteTableVpcAssociationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param localGatewayRouteTableId Identifier of EC2 Local Gateway Route Table.
         * 
         * @return builder
         * 
         */
        public Builder localGatewayRouteTableId(Output<String> localGatewayRouteTableId) {
            $.localGatewayRouteTableId = localGatewayRouteTableId;
            return this;
        }

        /**
         * @param localGatewayRouteTableId Identifier of EC2 Local Gateway Route Table.
         * 
         * @return builder
         * 
         */
        public Builder localGatewayRouteTableId(String localGatewayRouteTableId) {
            return localGatewayRouteTableId(Output.of(localGatewayRouteTableId));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId Identifier of EC2 VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId Identifier of EC2 VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public LocalGatewayRouteTableVpcAssociationArgs build() {
            $.localGatewayRouteTableId = Objects.requireNonNull($.localGatewayRouteTableId, "expected parameter 'localGatewayRouteTableId' to be non-null");
            $.vpcId = Objects.requireNonNull($.vpcId, "expected parameter 'vpcId' to be non-null");
            return $;
        }
    }

}