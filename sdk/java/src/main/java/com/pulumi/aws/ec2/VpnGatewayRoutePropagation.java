// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpnGatewayRoutePropagationArgs;
import com.pulumi.aws.ec2.inputs.VpnGatewayRoutePropagationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Requests automatic route propagation between a VPN gateway and a route table.
 * 
 * &gt; **Note:** This resource should not be used with a route table that has
 * the `propagating_vgws` argument set. If that argument is set, any route
 * propagation not explicitly listed in its value will be removed.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpnGatewayRoutePropagation;
 * import com.pulumi.aws.ec2.VpnGatewayRoutePropagationArgs;
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
 *         var example = new VpnGatewayRoutePropagation(&#34;example&#34;, VpnGatewayRoutePropagationArgs.builder()        
 *             .vpnGatewayId(aws_vpn_gateway.example().id())
 *             .routeTableId(aws_route_table.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation")
public class VpnGatewayRoutePropagation extends com.pulumi.resources.CustomResource {
    /**
     * The id of the `aws.ec2.RouteTable` to propagate routes into.
     * 
     */
    @Export(name="routeTableId", refs={String.class}, tree="[0]")
    private Output<String> routeTableId;

    /**
     * @return The id of the `aws.ec2.RouteTable` to propagate routes into.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }
    /**
     * The id of the `aws.ec2.VpnGateway` to propagate routes from.
     * 
     */
    @Export(name="vpnGatewayId", refs={String.class}, tree="[0]")
    private Output<String> vpnGatewayId;

    /**
     * @return The id of the `aws.ec2.VpnGateway` to propagate routes from.
     * 
     */
    public Output<String> vpnGatewayId() {
        return this.vpnGatewayId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpnGatewayRoutePropagation(String name) {
        this(name, VpnGatewayRoutePropagationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpnGatewayRoutePropagation(String name, VpnGatewayRoutePropagationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpnGatewayRoutePropagation(String name, VpnGatewayRoutePropagationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation", name, args == null ? VpnGatewayRoutePropagationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpnGatewayRoutePropagation(String name, Output<String> id, @Nullable VpnGatewayRoutePropagationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation", name, state, makeResourceOptions(options, id));
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
    public static VpnGatewayRoutePropagation get(String name, Output<String> id, @Nullable VpnGatewayRoutePropagationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpnGatewayRoutePropagation(name, id, state, options);
    }
}
