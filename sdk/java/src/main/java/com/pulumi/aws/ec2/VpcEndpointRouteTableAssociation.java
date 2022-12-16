// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcEndpointRouteTableAssociationArgs;
import com.pulumi.aws.ec2.inputs.VpcEndpointRouteTableAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a VPC Endpoint Route Table Association
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcEndpointRouteTableAssociation;
 * import com.pulumi.aws.ec2.VpcEndpointRouteTableAssociationArgs;
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
 *         var example = new VpcEndpointRouteTableAssociation(&#34;example&#34;, VpcEndpointRouteTableAssociationArgs.builder()        
 *             .routeTableId(aws_route_table.example().id())
 *             .vpcEndpointId(aws_vpc_endpoint.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Endpoint Route Table Associations can be imported using `vpc_endpoint_id` together with `route_table_id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation example vpce-aaaaaaaa/rtb-bbbbbbbb
 * ```
 * 
 */
@ResourceType(type="aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation")
public class VpcEndpointRouteTableAssociation extends com.pulumi.resources.CustomResource {
    /**
     * Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
     * 
     */
    @Export(name="routeTableId", refs={String.class}, tree="[0]")
    private Output<String> routeTableId;

    /**
     * @return Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }
    /**
     * Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
     * 
     */
    @Export(name="vpcEndpointId", refs={String.class}, tree="[0]")
    private Output<String> vpcEndpointId;

    /**
     * @return Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
     * 
     */
    public Output<String> vpcEndpointId() {
        return this.vpcEndpointId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpointRouteTableAssociation(String name) {
        this(name, VpcEndpointRouteTableAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpointRouteTableAssociation(String name, VpcEndpointRouteTableAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpointRouteTableAssociation(String name, VpcEndpointRouteTableAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, args == null ? VpcEndpointRouteTableAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcEndpointRouteTableAssociation(String name, Output<String> id, @Nullable VpcEndpointRouteTableAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, state, makeResourceOptions(options, id));
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
    public static VpcEndpointRouteTableAssociation get(String name, Output<String> id, @Nullable VpcEndpointRouteTableAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpointRouteTableAssociation(name, id, state, options);
    }
}
