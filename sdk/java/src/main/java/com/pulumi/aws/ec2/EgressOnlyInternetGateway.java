// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.EgressOnlyInternetGatewayArgs;
import com.pulumi.aws.ec2.inputs.EgressOnlyInternetGatewayState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * [IPv6 only] Creates an egress-only Internet gateway for your VPC.
 * An egress-only Internet gateway is used to enable outbound communication
 * over IPv6 from instances in your VPC to the Internet, and prevents hosts
 * outside of your VPC from initiating an IPv6 connection with your instance.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.EgressOnlyInternetGateway;
 * import com.pulumi.aws.ec2.EgressOnlyInternetGatewayArgs;
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
 *         var exampleVpc = new Vpc(&#34;exampleVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .assignGeneratedIpv6CidrBlock(true)
 *             .build());
 * 
 *         var exampleEgressOnlyInternetGateway = new EgressOnlyInternetGateway(&#34;exampleEgressOnlyInternetGateway&#34;, EgressOnlyInternetGatewayArgs.builder()        
 *             .vpcId(exampleVpc.id())
 *             .tags(Map.of(&#34;Name&#34;, &#34;main&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Egress-only Internet gateways can be imported using the `id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway example eigw-015e0e244e24dfe8a
 * ```
 * 
 */
@ResourceType(type="aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway")
public class EgressOnlyInternetGateway extends com.pulumi.resources.CustomResource {
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * The VPC ID to create in.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The VPC ID to create in.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EgressOnlyInternetGateway(String name) {
        this(name, EgressOnlyInternetGatewayArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EgressOnlyInternetGateway(String name, EgressOnlyInternetGatewayArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EgressOnlyInternetGateway(String name, EgressOnlyInternetGatewayArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, args == null ? EgressOnlyInternetGatewayArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EgressOnlyInternetGateway(String name, Output<String> id, @Nullable EgressOnlyInternetGatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, state, makeResourceOptions(options, id));
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
    public static EgressOnlyInternetGateway get(String name, Output<String> id, @Nullable EgressOnlyInternetGatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EgressOnlyInternetGateway(name, id, state, options);
    }
}