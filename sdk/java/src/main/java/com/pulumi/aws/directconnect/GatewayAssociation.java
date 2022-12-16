// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.directconnect;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.directconnect.GatewayAssociationArgs;
import com.pulumi.aws.directconnect.inputs.GatewayAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Associates a Direct Connect Gateway with a VGW or transit gateway.
 * 
 * To create a cross-account association, create an `aws.directconnect.GatewayAssociationProposal` resource
 * in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
 * by creating an `aws.directconnect.GatewayAssociation` resource with the `proposal_id` and `associated_gateway_owner_account_id` attributes set.
 * 
 * ## Example Usage
 * ### VPN Gateway Association
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.directconnect.Gateway;
 * import com.pulumi.aws.directconnect.GatewayArgs;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.VpnGateway;
 * import com.pulumi.aws.ec2.VpnGatewayArgs;
 * import com.pulumi.aws.directconnect.GatewayAssociation;
 * import com.pulumi.aws.directconnect.GatewayAssociationArgs;
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
 *         var exampleGateway = new Gateway(&#34;exampleGateway&#34;, GatewayArgs.builder()        
 *             .amazonSideAsn(&#34;64512&#34;)
 *             .build());
 * 
 *         var exampleVpc = new Vpc(&#34;exampleVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.255.255.0/28&#34;)
 *             .build());
 * 
 *         var exampleVpnGateway = new VpnGateway(&#34;exampleVpnGateway&#34;, VpnGatewayArgs.builder()        
 *             .vpcId(exampleVpc.id())
 *             .build());
 * 
 *         var exampleGatewayAssociation = new GatewayAssociation(&#34;exampleGatewayAssociation&#34;, GatewayAssociationArgs.builder()        
 *             .dxGatewayId(exampleGateway.id())
 *             .associatedGatewayId(exampleVpnGateway.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Transit Gateway Association
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.directconnect.Gateway;
 * import com.pulumi.aws.directconnect.GatewayArgs;
 * import com.pulumi.aws.ec2transitgateway.TransitGateway;
 * import com.pulumi.aws.directconnect.GatewayAssociation;
 * import com.pulumi.aws.directconnect.GatewayAssociationArgs;
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
 *         var exampleGateway = new Gateway(&#34;exampleGateway&#34;, GatewayArgs.builder()        
 *             .amazonSideAsn(&#34;64512&#34;)
 *             .build());
 * 
 *         var exampleTransitGateway = new TransitGateway(&#34;exampleTransitGateway&#34;);
 * 
 *         var exampleGatewayAssociation = new GatewayAssociation(&#34;exampleGatewayAssociation&#34;, GatewayAssociationArgs.builder()        
 *             .dxGatewayId(exampleGateway.id())
 *             .associatedGatewayId(exampleTransitGateway.id())
 *             .allowedPrefixes(            
 *                 &#34;10.255.255.0/30&#34;,
 *                 &#34;10.255.255.8/30&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Allowed Prefixes
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.directconnect.Gateway;
 * import com.pulumi.aws.directconnect.GatewayArgs;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.VpnGateway;
 * import com.pulumi.aws.ec2.VpnGatewayArgs;
 * import com.pulumi.aws.directconnect.GatewayAssociation;
 * import com.pulumi.aws.directconnect.GatewayAssociationArgs;
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
 *         var exampleGateway = new Gateway(&#34;exampleGateway&#34;, GatewayArgs.builder()        
 *             .amazonSideAsn(&#34;64512&#34;)
 *             .build());
 * 
 *         var exampleVpc = new Vpc(&#34;exampleVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.255.255.0/28&#34;)
 *             .build());
 * 
 *         var exampleVpnGateway = new VpnGateway(&#34;exampleVpnGateway&#34;, VpnGatewayArgs.builder()        
 *             .vpcId(exampleVpc.id())
 *             .build());
 * 
 *         var exampleGatewayAssociation = new GatewayAssociation(&#34;exampleGatewayAssociation&#34;, GatewayAssociationArgs.builder()        
 *             .dxGatewayId(exampleGateway.id())
 *             .associatedGatewayId(exampleVpnGateway.id())
 *             .allowedPrefixes(            
 *                 &#34;210.52.109.0/24&#34;,
 *                 &#34;175.45.176.0/22&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Direct Connect gateway associations can be imported using `dx_gateway_id` together with `associated_gateway_id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:directconnect/gatewayAssociation:GatewayAssociation example 345508c3-7215-4aef-9832-07c125d5bd0f/vgw-98765432
 * ```
 * 
 */
@ResourceType(type="aws:directconnect/gatewayAssociation:GatewayAssociation")
public class GatewayAssociation extends com.pulumi.resources.CustomResource {
    /**
     * VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     * 
     */
    @Export(name="allowedPrefixes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedPrefixes;

    /**
     * @return VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     * 
     */
    public Output<List<String>> allowedPrefixes() {
        return this.allowedPrefixes;
    }
    /**
     * The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for single account Direct Connect gateway associations.
     * 
     */
    @Export(name="associatedGatewayId", refs={String.class}, tree="[0]")
    private Output<String> associatedGatewayId;

    /**
     * @return The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for single account Direct Connect gateway associations.
     * 
     */
    public Output<String> associatedGatewayId() {
        return this.associatedGatewayId;
    }
    /**
     * The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for cross-account Direct Connect gateway associations.
     * 
     */
    @Export(name="associatedGatewayOwnerAccountId", refs={String.class}, tree="[0]")
    private Output<String> associatedGatewayOwnerAccountId;

    /**
     * @return The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
     * Used for cross-account Direct Connect gateway associations.
     * 
     */
    public Output<String> associatedGatewayOwnerAccountId() {
        return this.associatedGatewayOwnerAccountId;
    }
    /**
     * The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
     * 
     */
    @Export(name="associatedGatewayType", refs={String.class}, tree="[0]")
    private Output<String> associatedGatewayType;

    /**
     * @return The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
     * 
     */
    public Output<String> associatedGatewayType() {
        return this.associatedGatewayType;
    }
    /**
     * The ID of the Direct Connect gateway association.
     * 
     */
    @Export(name="dxGatewayAssociationId", refs={String.class}, tree="[0]")
    private Output<String> dxGatewayAssociationId;

    /**
     * @return The ID of the Direct Connect gateway association.
     * 
     */
    public Output<String> dxGatewayAssociationId() {
        return this.dxGatewayAssociationId;
    }
    /**
     * The ID of the Direct Connect gateway.
     * 
     */
    @Export(name="dxGatewayId", refs={String.class}, tree="[0]")
    private Output<String> dxGatewayId;

    /**
     * @return The ID of the Direct Connect gateway.
     * 
     */
    public Output<String> dxGatewayId() {
        return this.dxGatewayId;
    }
    /**
     * The ID of the AWS account that owns the Direct Connect gateway.
     * 
     */
    @Export(name="dxGatewayOwnerAccountId", refs={String.class}, tree="[0]")
    private Output<String> dxGatewayOwnerAccountId;

    /**
     * @return The ID of the AWS account that owns the Direct Connect gateway.
     * 
     */
    public Output<String> dxGatewayOwnerAccountId() {
        return this.dxGatewayOwnerAccountId;
    }
    /**
     * The ID of the Direct Connect gateway association proposal.
     * Used for cross-account Direct Connect gateway associations.
     * 
     */
    @Export(name="proposalId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proposalId;

    /**
     * @return The ID of the Direct Connect gateway association proposal.
     * Used for cross-account Direct Connect gateway associations.
     * 
     */
    public Output<Optional<String>> proposalId() {
        return Codegen.optional(this.proposalId);
    }
    /**
     * @deprecated
     * use &#39;associated_gateway_id&#39; argument instead
     * 
     */
    @Deprecated /* use 'associated_gateway_id' argument instead */
    @Export(name="vpnGatewayId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpnGatewayId;

    public Output<Optional<String>> vpnGatewayId() {
        return Codegen.optional(this.vpnGatewayId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GatewayAssociation(String name) {
        this(name, GatewayAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GatewayAssociation(String name, GatewayAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GatewayAssociation(String name, GatewayAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:directconnect/gatewayAssociation:GatewayAssociation", name, args == null ? GatewayAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GatewayAssociation(String name, Output<String> id, @Nullable GatewayAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:directconnect/gatewayAssociation:GatewayAssociation", name, state, makeResourceOptions(options, id));
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
    public static GatewayAssociation get(String name, Output<String> id, @Nullable GatewayAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GatewayAssociation(name, id, state, options);
    }
}
