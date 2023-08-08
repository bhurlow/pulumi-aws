// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2transitgateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2transitgateway.ConnectArgs;
import com.pulumi.aws.ec2transitgateway.inputs.ConnectState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an EC2 Transit Gateway Connect.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2transitgateway.VpcAttachment;
 * import com.pulumi.aws.ec2transitgateway.VpcAttachmentArgs;
 * import com.pulumi.aws.ec2transitgateway.Connect;
 * import com.pulumi.aws.ec2transitgateway.ConnectArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new VpcAttachment(&#34;example&#34;, VpcAttachmentArgs.builder()        
 *             .subnetIds(aws_subnet.example().id())
 *             .transitGatewayId(aws_ec2_transit_gateway.example().id())
 *             .vpcId(aws_vpc.example().id())
 *             .build());
 * 
 *         var attachment = new Connect(&#34;attachment&#34;, ConnectArgs.builder()        
 *             .transportAttachmentId(example.id())
 *             .transitGatewayId(aws_ec2_transit_gateway.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_ec2_transit_gateway_connect.example
 * 
 *  id = &#34;tgw-attach-12345678&#34; } Using `pulumi import`, import `aws_ec2_transit_gateway_connect` using the EC2 Transit Gateway Connect identifier. For exampleconsole % pulumi import aws_ec2_transit_gateway_connect.example tgw-attach-12345678
 * 
 */
@ResourceType(type="aws:ec2transitgateway/connect:Connect")
public class Connect extends com.pulumi.resources.CustomResource {
    /**
     * The tunnel protocol. Valid values: `gre`. Default is `gre`.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> protocol;

    /**
     * @return The tunnel protocol. Valid values: `gre`. Default is `gre`.
     * 
     */
    public Output<Optional<String>> protocol() {
        return Codegen.optional(this.protocol);
    }
    /**
     * Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    @Export(name="transitGatewayDefaultRouteTableAssociation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> transitGatewayDefaultRouteTableAssociation;

    /**
     * @return Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    public Output<Optional<Boolean>> transitGatewayDefaultRouteTableAssociation() {
        return Codegen.optional(this.transitGatewayDefaultRouteTableAssociation);
    }
    /**
     * Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    @Export(name="transitGatewayDefaultRouteTablePropagation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> transitGatewayDefaultRouteTablePropagation;

    /**
     * @return Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    public Output<Optional<Boolean>> transitGatewayDefaultRouteTablePropagation() {
        return Codegen.optional(this.transitGatewayDefaultRouteTablePropagation);
    }
    /**
     * Identifier of EC2 Transit Gateway.
     * 
     */
    @Export(name="transitGatewayId", refs={String.class}, tree="[0]")
    private Output<String> transitGatewayId;

    /**
     * @return Identifier of EC2 Transit Gateway.
     * 
     */
    public Output<String> transitGatewayId() {
        return this.transitGatewayId;
    }
    /**
     * The underlaying VPC attachment
     * 
     */
    @Export(name="transportAttachmentId", refs={String.class}, tree="[0]")
    private Output<String> transportAttachmentId;

    /**
     * @return The underlaying VPC attachment
     * 
     */
    public Output<String> transportAttachmentId() {
        return this.transportAttachmentId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connect(String name) {
        this(name, ConnectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connect(String name, ConnectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connect(String name, ConnectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2transitgateway/connect:Connect", name, args == null ? ConnectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connect(String name, Output<String> id, @Nullable ConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2transitgateway/connect:Connect", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Connect get(String name, Output<String> id, @Nullable ConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connect(name, id, state, options);
    }
}
