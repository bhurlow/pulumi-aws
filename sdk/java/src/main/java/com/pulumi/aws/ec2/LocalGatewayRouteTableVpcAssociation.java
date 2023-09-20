// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.LocalGatewayRouteTableVpcAssociationArgs;
import com.pulumi.aws.ec2.inputs.LocalGatewayRouteTableVpcAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an EC2 Local Gateway Route Table VPC Association. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-local-gateways.html#vpc-associations).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.Ec2Functions;
 * import com.pulumi.aws.ec2.inputs.GetLocalGatewayRouteTableArgs;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.LocalGatewayRouteTableVpcAssociation;
 * import com.pulumi.aws.ec2.LocalGatewayRouteTableVpcAssociationArgs;
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
 *         final var exampleLocalGatewayRouteTable = Ec2Functions.getLocalGatewayRouteTable(GetLocalGatewayRouteTableArgs.builder()
 *             .outpostArn(&#34;arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef&#34;)
 *             .build());
 * 
 *         var exampleVpc = new Vpc(&#34;exampleVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.0.0.0/16&#34;)
 *             .build());
 * 
 *         var exampleLocalGatewayRouteTableVpcAssociation = new LocalGatewayRouteTableVpcAssociation(&#34;exampleLocalGatewayRouteTableVpcAssociation&#34;, LocalGatewayRouteTableVpcAssociationArgs.builder()        
 *             .localGatewayRouteTableId(exampleLocalGatewayRouteTable.applyValue(getLocalGatewayRouteTableResult -&gt; getLocalGatewayRouteTableResult.id()))
 *             .vpcId(exampleVpc.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import `aws_ec2_local_gateway_route_table_vpc_association` using the Local Gateway Route Table VPC Association identifier. For example:
 * 
 * ```sh
 *  $ pulumi import aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation example lgw-vpc-assoc-1234567890abcdef
 * ```
 * 
 */
@ResourceType(type="aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation")
public class LocalGatewayRouteTableVpcAssociation extends com.pulumi.resources.CustomResource {
    @Export(name="localGatewayId", refs={String.class}, tree="[0]")
    private Output<String> localGatewayId;

    public Output<String> localGatewayId() {
        return this.localGatewayId;
    }
    /**
     * Identifier of EC2 Local Gateway Route Table.
     * 
     */
    @Export(name="localGatewayRouteTableId", refs={String.class}, tree="[0]")
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
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
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
     * Identifier of EC2 VPC.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return Identifier of EC2 VPC.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LocalGatewayRouteTableVpcAssociation(String name) {
        this(name, LocalGatewayRouteTableVpcAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LocalGatewayRouteTableVpcAssociation(String name, LocalGatewayRouteTableVpcAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LocalGatewayRouteTableVpcAssociation(String name, LocalGatewayRouteTableVpcAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation", name, args == null ? LocalGatewayRouteTableVpcAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LocalGatewayRouteTableVpcAssociation(String name, Output<String> id, @Nullable LocalGatewayRouteTableVpcAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static LocalGatewayRouteTableVpcAssociation get(String name, Output<String> id, @Nullable LocalGatewayRouteTableVpcAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LocalGatewayRouteTableVpcAssociation(name, id, state, options);
    }
}
