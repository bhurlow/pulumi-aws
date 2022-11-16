// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionState;
import com.pulumi.aws.ec2.outputs.VpcPeeringConnectionAccepter;
import com.pulumi.aws.ec2.outputs.VpcPeeringConnectionRequester;
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
 * Provides a resource to manage a VPC peering connection.
 * 
 * &gt; **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
 * both a standalone VPC Peering Connection Options and a VPC Peering Connection
 * resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
 * connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
 * Doing so will cause a conflict of options and will overwrite the options.
 * Using a VPC Peering Connection Options resource decouples management of the connection options from
 * management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.
 * 
 * &gt; **Note:** For cross-account (requester&#39;s AWS account differs from the accepter&#39;s AWS account) or inter-region
 * VPC Peering Connections use the `aws.ec2.VpcPeeringConnection` resource to manage the requester&#39;s side of the
 * connection and use the `aws.ec2.VpcPeeringConnectionAccepter` resource to manage the accepter&#39;s side of the connection.
 * 
 * &gt; **Note:** Creating multiple `aws.ec2.VpcPeeringConnection` resources with the same `peer_vpc_id` and `vpc_id` will not produce an error. Instead, AWS will return the connection `id` that already exists, resulting in multiple `aws.ec2.VpcPeeringConnection` resources with the same `id`.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcPeeringConnection;
 * import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
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
 *         var foo = new VpcPeeringConnection(&#34;foo&#34;, VpcPeeringConnectionArgs.builder()        
 *             .peerOwnerId(var_.peer_owner_id())
 *             .peerVpcId(aws_vpc.bar().id())
 *             .vpcId(aws_vpc.foo().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Basic usage with connection options:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcPeeringConnection;
 * import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
 * import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionAccepterArgs;
 * import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionRequesterArgs;
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
 *         var foo = new VpcPeeringConnection(&#34;foo&#34;, VpcPeeringConnectionArgs.builder()        
 *             .peerOwnerId(var_.peer_owner_id())
 *             .peerVpcId(aws_vpc.bar().id())
 *             .vpcId(aws_vpc.foo().id())
 *             .accepter(VpcPeeringConnectionAccepterArgs.builder()
 *                 .allowRemoteVpcDnsResolution(true)
 *                 .build())
 *             .requester(VpcPeeringConnectionRequesterArgs.builder()
 *                 .allowRemoteVpcDnsResolution(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Basic usage with tags:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.VpcPeeringConnection;
 * import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
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
 *         var fooVpc = new Vpc(&#34;fooVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .build());
 * 
 *         var bar = new Vpc(&#34;bar&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.2.0.0/16&#34;)
 *             .build());
 * 
 *         var fooVpcPeeringConnection = new VpcPeeringConnection(&#34;fooVpcPeeringConnection&#34;, VpcPeeringConnectionArgs.builder()        
 *             .peerOwnerId(var_.peer_owner_id())
 *             .peerVpcId(bar.id())
 *             .vpcId(fooVpc.id())
 *             .autoAccept(true)
 *             .tags(Map.of(&#34;Name&#34;, &#34;VPC Peering between foo and bar&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Basic usage with region:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.VpcPeeringConnection;
 * import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
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
 *         var fooVpc = new Vpc(&#34;fooVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.us-west-2())
 *                 .build());
 * 
 *         var bar = new Vpc(&#34;bar&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.2.0.0/16&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.us-east-1())
 *                 .build());
 * 
 *         var fooVpcPeeringConnection = new VpcPeeringConnection(&#34;fooVpcPeeringConnection&#34;, VpcPeeringConnectionArgs.builder()        
 *             .peerOwnerId(var_.peer_owner_id())
 *             .peerVpcId(bar.id())
 *             .vpcId(fooVpc.id())
 *             .peerRegion(&#34;us-east-1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Notes
 * 
 * If both VPCs are not in the same AWS account and region do not enable the `auto_accept` attribute.
 * The accepter can manage its side of the connection using the `aws.ec2.VpcPeeringConnectionAccepter` resource
 * or accept the connection manually using the AWS Management Console, AWS CLI, through SDKs, etc.
 * 
 * ## Import
 * 
 * VPC Peering resources can be imported using the `vpc peering id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/vpcPeeringConnection:VpcPeeringConnection test_connection pcx-111aaa111
 * ```
 * 
 */
@ResourceType(type="aws:ec2/vpcPeeringConnection:VpcPeeringConnection")
public class VpcPeeringConnection extends com.pulumi.resources.CustomResource {
    /**
     * The status of the VPC Peering Connection request.
     * 
     */
    @Export(name="acceptStatus", type=String.class, parameters={})
    private Output<String> acceptStatus;

    /**
     * @return The status of the VPC Peering Connection request.
     * 
     */
    public Output<String> acceptStatus() {
        return this.acceptStatus;
    }
    /**
     * An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
     * the peering connection (a maximum of one).
     * 
     */
    @Export(name="accepter", type=VpcPeeringConnectionAccepter.class, parameters={})
    private Output<VpcPeeringConnectionAccepter> accepter;

    /**
     * @return An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
     * the peering connection (a maximum of one).
     * 
     */
    public Output<VpcPeeringConnectionAccepter> accepter() {
        return this.accepter;
    }
    /**
     * Accept the peering (both VPCs need to be in the same AWS account and region).
     * 
     */
    @Export(name="autoAccept", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoAccept;

    /**
     * @return Accept the peering (both VPCs need to be in the same AWS account and region).
     * 
     */
    public Output<Optional<Boolean>> autoAccept() {
        return Codegen.optional(this.autoAccept);
    }
    /**
     * The AWS account ID of the owner of the peer VPC.
     * Defaults to the account ID the AWS provider is currently connected to.
     * 
     */
    @Export(name="peerOwnerId", type=String.class, parameters={})
    private Output<String> peerOwnerId;

    /**
     * @return The AWS account ID of the owner of the peer VPC.
     * Defaults to the account ID the AWS provider is currently connected to.
     * 
     */
    public Output<String> peerOwnerId() {
        return this.peerOwnerId;
    }
    /**
     * The region of the accepter VPC of the VPC Peering Connection. `auto_accept` must be `false`,
     * and use the `aws.ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
     * 
     */
    @Export(name="peerRegion", type=String.class, parameters={})
    private Output<String> peerRegion;

    /**
     * @return The region of the accepter VPC of the VPC Peering Connection. `auto_accept` must be `false`,
     * and use the `aws.ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
     * 
     */
    public Output<String> peerRegion() {
        return this.peerRegion;
    }
    /**
     * The ID of the VPC with which you are creating the VPC Peering Connection.
     * 
     */
    @Export(name="peerVpcId", type=String.class, parameters={})
    private Output<String> peerVpcId;

    /**
     * @return The ID of the VPC with which you are creating the VPC Peering Connection.
     * 
     */
    public Output<String> peerVpcId() {
        return this.peerVpcId;
    }
    /**
     * A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
     * the peering connection (a maximum of one).
     * 
     */
    @Export(name="requester", type=VpcPeeringConnectionRequester.class, parameters={})
    private Output<VpcPeeringConnectionRequester> requester;

    /**
     * @return A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
     * the peering connection (a maximum of one).
     * 
     */
    public Output<VpcPeeringConnectionRequester> requester() {
        return this.requester;
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
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
     * The ID of the requester VPC.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the requester VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcPeeringConnection(String name) {
        this(name, VpcPeeringConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcPeeringConnection(String name, VpcPeeringConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcPeeringConnection(String name, VpcPeeringConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, args == null ? VpcPeeringConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcPeeringConnection(String name, Output<String> id, @Nullable VpcPeeringConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, state, makeResourceOptions(options, id));
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
    public static VpcPeeringConnection get(String name, Output<String> id, @Nullable VpcPeeringConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcPeeringConnection(name, id, state, options);
    }
}
