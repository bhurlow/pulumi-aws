// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepterArgs;
import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionAccepterState;
import com.pulumi.aws.ec2.outputs.VpcPeeringConnectionAccepterAccepter;
import com.pulumi.aws.ec2.outputs.VpcPeeringConnectionAccepterRequester;
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
 * Provides a resource to manage the accepter&#39;s side of a VPC Peering Connection.
 * 
 * When a cross-account (requester&#39;s AWS account differs from the accepter&#39;s AWS account) or an inter-region
 * VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
 * accepter&#39;s account.
 * The requester can use the `aws.ec2.VpcPeeringConnection` resource to manage its side of the connection
 * and the accepter can use the `aws.ec2.VpcPeeringConnectionAccepter` resource to &#34;adopt&#34; its side of the
 * connection into management.
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
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetCallerIdentityArgs;
 * import com.pulumi.aws.ec2.VpcPeeringConnection;
 * import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
 * import com.pulumi.aws.ec2.VpcPeeringConnectionAccepter;
 * import com.pulumi.aws.ec2.VpcPeeringConnectionAccepterArgs;
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
 *         var peer = new Provider(&#34;peer&#34;, ProviderArgs.builder()        
 *             .region(&#34;us-west-2&#34;)
 *             .build());
 * 
 *         var main = new Vpc(&#34;main&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.0.0.0/16&#34;)
 *             .build());
 * 
 *         var peerVpc = new Vpc(&#34;peerVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.peer())
 *                 .build());
 * 
 *         final var peerCallerIdentity = AwsFunctions.getCallerIdentity();
 * 
 *         var peerVpcPeeringConnection = new VpcPeeringConnection(&#34;peerVpcPeeringConnection&#34;, VpcPeeringConnectionArgs.builder()        
 *             .vpcId(main.id())
 *             .peerVpcId(peerVpc.id())
 *             .peerOwnerId(peerCallerIdentity.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId()))
 *             .peerRegion(&#34;us-west-2&#34;)
 *             .autoAccept(false)
 *             .tags(Map.of(&#34;Side&#34;, &#34;Requester&#34;))
 *             .build());
 * 
 *         var peerVpcPeeringConnectionAccepter = new VpcPeeringConnectionAccepter(&#34;peerVpcPeeringConnectionAccepter&#34;, VpcPeeringConnectionAccepterArgs.builder()        
 *             .vpcPeeringConnectionId(peerVpcPeeringConnection.id())
 *             .autoAccept(true)
 *             .tags(Map.of(&#34;Side&#34;, &#34;Accepter&#34;))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.peer())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Peering Connection Accepters can be imported by using the Peering Connection ID, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter example pcx-12345678
 * ```
 * 
 *  Certain resource arguments, like `auto_accept`, do not have an EC2 API method for reading the information after peering connection creation. If the argument is set in the provider configuration on an imported resource, this provder will always show a difference. To workaround this behavior, either omit the argument from the configuration or use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to hide the difference, e.g. terraform resource &#34;aws_vpc_peering_connection_accepter&#34; &#34;example&#34; {
 * 
 * # ... other configuration ...
 * 
 * # There is no AWS EC2 API for reading auto_accept
 * 
 *  lifecycle {
 * 
 *  ignore_changes = [auto_accept]
 * 
 *  } }
 * 
 */
@ResourceType(type="aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter")
public class VpcPeeringConnectionAccepter extends com.pulumi.resources.CustomResource {
    /**
     * The status of the VPC Peering Connection request.
     * 
     */
    @Export(name="acceptStatus", refs={String.class}, tree="[0]")
    private Output<String> acceptStatus;

    /**
     * @return The status of the VPC Peering Connection request.
     * 
     */
    public Output<String> acceptStatus() {
        return this.acceptStatus;
    }
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     * 
     */
    @Export(name="accepter", refs={VpcPeeringConnectionAccepterAccepter.class}, tree="[0]")
    private Output<VpcPeeringConnectionAccepterAccepter> accepter;

    /**
     * @return A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     * 
     */
    public Output<VpcPeeringConnectionAccepterAccepter> accepter() {
        return this.accepter;
    }
    /**
     * Whether or not to accept the peering request. Defaults to `false`.
     * 
     */
    @Export(name="autoAccept", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoAccept;

    /**
     * @return Whether or not to accept the peering request. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> autoAccept() {
        return Codegen.optional(this.autoAccept);
    }
    /**
     * The AWS account ID of the owner of the requester VPC.
     * 
     */
    @Export(name="peerOwnerId", refs={String.class}, tree="[0]")
    private Output<String> peerOwnerId;

    /**
     * @return The AWS account ID of the owner of the requester VPC.
     * 
     */
    public Output<String> peerOwnerId() {
        return this.peerOwnerId;
    }
    /**
     * The region of the accepter VPC.
     * 
     */
    @Export(name="peerRegion", refs={String.class}, tree="[0]")
    private Output<String> peerRegion;

    /**
     * @return The region of the accepter VPC.
     * 
     */
    public Output<String> peerRegion() {
        return this.peerRegion;
    }
    /**
     * The ID of the requester VPC.
     * 
     */
    @Export(name="peerVpcId", refs={String.class}, tree="[0]")
    private Output<String> peerVpcId;

    /**
     * @return The ID of the requester VPC.
     * 
     */
    public Output<String> peerVpcId() {
        return this.peerVpcId;
    }
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     * 
     */
    @Export(name="requester", refs={VpcPeeringConnectionAccepterRequester.class}, tree="[0]")
    private Output<VpcPeeringConnectionAccepterRequester> requester;

    /**
     * @return A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     * 
     */
    public Output<VpcPeeringConnectionAccepterRequester> requester() {
        return this.requester;
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
     * The ID of the accepter VPC.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the accepter VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The VPC Peering Connection ID to manage.
     * 
     */
    @Export(name="vpcPeeringConnectionId", refs={String.class}, tree="[0]")
    private Output<String> vpcPeeringConnectionId;

    /**
     * @return The VPC Peering Connection ID to manage.
     * 
     */
    public Output<String> vpcPeeringConnectionId() {
        return this.vpcPeeringConnectionId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcPeeringConnectionAccepter(String name) {
        this(name, VpcPeeringConnectionAccepterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcPeeringConnectionAccepter(String name, VpcPeeringConnectionAccepterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcPeeringConnectionAccepter(String name, VpcPeeringConnectionAccepterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, args == null ? VpcPeeringConnectionAccepterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcPeeringConnectionAccepter(String name, Output<String> id, @Nullable VpcPeeringConnectionAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, state, makeResourceOptions(options, id));
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
    public static VpcPeeringConnectionAccepter get(String name, Output<String> id, @Nullable VpcPeeringConnectionAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcPeeringConnectionAccepter(name, id, state, options);
    }
}
