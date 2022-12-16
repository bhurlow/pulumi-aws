// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2transitgateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2transitgateway.PeeringAttachmentArgs;
import com.pulumi.aws.ec2transitgateway.inputs.PeeringAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an EC2 Transit Gateway Peering Attachment.
 * For examples of custom route table association and propagation, see the [EC2 Transit Gateway Networking Examples Guide](https://docs.aws.amazon.com/vpc/latest/tgw/TGW_Scenarios.html).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.Provider;
 * import com.pulumi.aws.ProviderArgs;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetRegionArgs;
 * import com.pulumi.aws.ec2transitgateway.TransitGateway;
 * import com.pulumi.aws.ec2transitgateway.TransitGatewayArgs;
 * import com.pulumi.aws.ec2transitgateway.PeeringAttachment;
 * import com.pulumi.aws.ec2transitgateway.PeeringAttachmentArgs;
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
 *         var local = new Provider(&#34;local&#34;, ProviderArgs.builder()        
 *             .region(&#34;us-east-1&#34;)
 *             .build());
 * 
 *         var peer = new Provider(&#34;peer&#34;, ProviderArgs.builder()        
 *             .region(&#34;us-west-2&#34;)
 *             .build());
 * 
 *         final var peerRegion = AwsFunctions.getRegion();
 * 
 *         var localTransitGateway = new TransitGateway(&#34;localTransitGateway&#34;, TransitGatewayArgs.builder()        
 *             .tags(Map.of(&#34;Name&#34;, &#34;Local TGW&#34;))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.local())
 *                 .build());
 * 
 *         var peerTransitGateway = new TransitGateway(&#34;peerTransitGateway&#34;, TransitGatewayArgs.builder()        
 *             .tags(Map.of(&#34;Name&#34;, &#34;Peer TGW&#34;))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.peer())
 *                 .build());
 * 
 *         var example = new PeeringAttachment(&#34;example&#34;, PeeringAttachmentArgs.builder()        
 *             .peerAccountId(peerTransitGateway.ownerId())
 *             .peerRegion(peerRegion.applyValue(getRegionResult -&gt; getRegionResult.name()))
 *             .peerTransitGatewayId(peerTransitGateway.id())
 *             .transitGatewayId(localTransitGateway.id())
 *             .tags(Map.of(&#34;Name&#34;, &#34;TGW Peering Requestor&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_ec2_transit_gateway_peering_attachment` can be imported by using the EC2 Transit Gateway Attachment identifier, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2transitgateway/peeringAttachment:PeeringAttachment example tgw-attach-12345678
 * ```
 * 
 */
@ResourceType(type="aws:ec2transitgateway/peeringAttachment:PeeringAttachment")
public class PeeringAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the AWS provider is currently connected to.
     * 
     */
    @Export(name="peerAccountId", refs={String.class}, tree="[0]")
    private Output<String> peerAccountId;

    /**
     * @return Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the AWS provider is currently connected to.
     * 
     */
    public Output<String> peerAccountId() {
        return this.peerAccountId;
    }
    /**
     * Region of EC2 Transit Gateway to peer with.
     * 
     */
    @Export(name="peerRegion", refs={String.class}, tree="[0]")
    private Output<String> peerRegion;

    /**
     * @return Region of EC2 Transit Gateway to peer with.
     * 
     */
    public Output<String> peerRegion() {
        return this.peerRegion;
    }
    /**
     * Identifier of EC2 Transit Gateway to peer with.
     * 
     */
    @Export(name="peerTransitGatewayId", refs={String.class}, tree="[0]")
    private Output<String> peerTransitGatewayId;

    /**
     * @return Identifier of EC2 Transit Gateway to peer with.
     * 
     */
    public Output<String> peerTransitGatewayId() {
        return this.peerTransitGatewayId;
    }
    /**
     * Key-value tags for the EC2 Transit Gateway Peering Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value tags for the EC2 Transit Gateway Peering Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PeeringAttachment(String name) {
        this(name, PeeringAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PeeringAttachment(String name, PeeringAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PeeringAttachment(String name, PeeringAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2transitgateway/peeringAttachment:PeeringAttachment", name, args == null ? PeeringAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PeeringAttachment(String name, Output<String> id, @Nullable PeeringAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2transitgateway/peeringAttachment:PeeringAttachment", name, state, makeResourceOptions(options, id));
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
    public static PeeringAttachment get(String name, Output<String> id, @Nullable PeeringAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PeeringAttachment(name, id, state, options);
    }
}
