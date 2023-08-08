// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.directconnect;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.directconnect.HostedPrivateVirtualInterfaceAccepterArgs;
import com.pulumi.aws.directconnect.inputs.HostedPrivateVirtualInterfaceAccepterState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage the accepter&#39;s side of a Direct Connect hosted private virtual interface.
 * This resource accepts ownership of a private virtual interface created by another AWS account.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.Provider;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetCallerIdentityArgs;
 * import com.pulumi.aws.ec2.VpnGateway;
 * import com.pulumi.aws.ec2.VpnGatewayArgs;
 * import com.pulumi.aws.directconnect.HostedPrivateVirtualInterface;
 * import com.pulumi.aws.directconnect.HostedPrivateVirtualInterfaceArgs;
 * import com.pulumi.aws.directconnect.HostedPrivateVirtualInterfaceAccepter;
 * import com.pulumi.aws.directconnect.HostedPrivateVirtualInterfaceAccepterArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var accepter = new Provider(&#34;accepter&#34;);
 * 
 *         final var accepterCallerIdentity = AwsFunctions.getCallerIdentity();
 * 
 *         var vpnGw = new VpnGateway(&#34;vpnGw&#34;, VpnGatewayArgs.Empty, CustomResourceOptions.builder()
 *             .provider(aws.accepter())
 *             .build());
 * 
 *         var creator = new HostedPrivateVirtualInterface(&#34;creator&#34;, HostedPrivateVirtualInterfaceArgs.builder()        
 *             .connectionId(&#34;dxcon-zzzzzzzz&#34;)
 *             .ownerAccountId(accepterCallerIdentity.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId()))
 *             .vlan(4094)
 *             .addressFamily(&#34;ipv4&#34;)
 *             .bgpAsn(65352)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpnGw)
 *                 .build());
 * 
 *         var accepterHostedPrivateVirtualInterfaceAccepter = new HostedPrivateVirtualInterfaceAccepter(&#34;accepterHostedPrivateVirtualInterfaceAccepter&#34;, HostedPrivateVirtualInterfaceAccepterArgs.builder()        
 *             .virtualInterfaceId(creator.id())
 *             .vpnGatewayId(vpnGw.id())
 *             .tags(Map.of(&#34;Side&#34;, &#34;Accepter&#34;))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.accepter())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_dx_hosted_private_virtual_interface_accepter.test
 * 
 *  id = &#34;dxvif-33cc44dd&#34; } Using `pulumi import`, import Direct Connect hosted private virtual interfaces using the VIF `id`. For exampleconsole % pulumi import aws_dx_hosted_private_virtual_interface_accepter.test dxvif-33cc44dd
 * 
 */
@ResourceType(type="aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter")
public class HostedPrivateVirtualInterfaceAccepter extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the virtual interface.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the virtual interface.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     * 
     */
    @Export(name="dxGatewayId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dxGatewayId;

    /**
     * @return The ID of the Direct Connect gateway to which to connect the virtual interface.
     * 
     */
    public Output<Optional<String>> dxGatewayId() {
        return Codegen.optional(this.dxGatewayId);
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The ID of the Direct Connect virtual interface to accept.
     * 
     */
    @Export(name="virtualInterfaceId", refs={String.class}, tree="[0]")
    private Output<String> virtualInterfaceId;

    /**
     * @return The ID of the Direct Connect virtual interface to accept.
     * 
     */
    public Output<String> virtualInterfaceId() {
        return this.virtualInterfaceId;
    }
    /**
     * The ID of the virtual private gateway to which to connect the virtual interface.
     * 
     */
    @Export(name="vpnGatewayId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpnGatewayId;

    /**
     * @return The ID of the virtual private gateway to which to connect the virtual interface.
     * 
     */
    public Output<Optional<String>> vpnGatewayId() {
        return Codegen.optional(this.vpnGatewayId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostedPrivateVirtualInterfaceAccepter(String name) {
        this(name, HostedPrivateVirtualInterfaceAccepterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostedPrivateVirtualInterfaceAccepter(String name, HostedPrivateVirtualInterfaceAccepterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostedPrivateVirtualInterfaceAccepter(String name, HostedPrivateVirtualInterfaceAccepterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, args == null ? HostedPrivateVirtualInterfaceAccepterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HostedPrivateVirtualInterfaceAccepter(String name, Output<String> id, @Nullable HostedPrivateVirtualInterfaceAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, state, makeResourceOptions(options, id));
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
    public static HostedPrivateVirtualInterfaceAccepter get(String name, Output<String> id, @Nullable HostedPrivateVirtualInterfaceAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostedPrivateVirtualInterfaceAccepter(name, id, state, options);
    }
}
