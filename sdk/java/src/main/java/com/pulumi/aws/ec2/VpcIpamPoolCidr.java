// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcIpamPoolCidrArgs;
import com.pulumi.aws.ec2.inputs.VpcIpamPoolCidrState;
import com.pulumi.aws.ec2.outputs.VpcIpamPoolCidrCidrAuthorizationContext;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provisions a CIDR from an IPAM address pool.
 * 
 * &gt; **NOTE:** Provisioning Public IPv4 or Public IPv6 require [steps outside the scope of this resource](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html#prepare-for-byoip). The resource accepts `message` and `signature` as part of the `cidr_authorization_context` attribute but those must be generated ahead of time. Public IPv6 CIDRs that are provisioned into a Pool with `publicly_advertisable = true` and all public IPv4 CIDRs also require creating a Route Origin Authorization (ROA) object in your Regional Internet Registry (RIR).
 * 
 * &gt; **NOTE:** In order to deprovision CIDRs all Allocations must be released. Allocations created by a VPC take up to 30 minutes to be released. However, for IPAM to properly manage the removal of allocation records created by VPCs and other resources, you must [grant it permissions](https://docs.aws.amazon.com/vpc/latest/ipam/choose-single-user-or-orgs-ipam.html) in
 * either a single account or organizationally. If you are unable to deprovision a cidr after waiting over 30 minutes, you may be missing the Service Linked Role.
 * 
 * ## Example Usage
 * 
 * Basic usage:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetRegionArgs;
 * import com.pulumi.aws.ec2.VpcIpam;
 * import com.pulumi.aws.ec2.VpcIpamArgs;
 * import com.pulumi.aws.ec2.inputs.VpcIpamOperatingRegionArgs;
 * import com.pulumi.aws.ec2.VpcIpamPool;
 * import com.pulumi.aws.ec2.VpcIpamPoolArgs;
 * import com.pulumi.aws.ec2.VpcIpamPoolCidr;
 * import com.pulumi.aws.ec2.VpcIpamPoolCidrArgs;
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
 *         final var current = AwsFunctions.getRegion();
 * 
 *         var exampleVpcIpam = new VpcIpam(&#34;exampleVpcIpam&#34;, VpcIpamArgs.builder()        
 *             .operatingRegions(VpcIpamOperatingRegionArgs.builder()
 *                 .regionName(current.applyValue(getRegionResult -&gt; getRegionResult.name()))
 *                 .build())
 *             .build());
 * 
 *         var exampleVpcIpamPool = new VpcIpamPool(&#34;exampleVpcIpamPool&#34;, VpcIpamPoolArgs.builder()        
 *             .addressFamily(&#34;ipv4&#34;)
 *             .ipamScopeId(exampleVpcIpam.privateDefaultScopeId())
 *             .locale(current.applyValue(getRegionResult -&gt; getRegionResult.name()))
 *             .build());
 * 
 *         var exampleVpcIpamPoolCidr = new VpcIpamPoolCidr(&#34;exampleVpcIpamPoolCidr&#34;, VpcIpamPoolCidrArgs.builder()        
 *             .ipamPoolId(exampleVpcIpamPool.id())
 *             .cidr(&#34;172.2.0.0/16&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Provision Public IPv6 Pool CIDRs:
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetRegionArgs;
 * import com.pulumi.aws.ec2.VpcIpam;
 * import com.pulumi.aws.ec2.VpcIpamArgs;
 * import com.pulumi.aws.ec2.inputs.VpcIpamOperatingRegionArgs;
 * import com.pulumi.aws.ec2.VpcIpamPool;
 * import com.pulumi.aws.ec2.VpcIpamPoolArgs;
 * import com.pulumi.aws.ec2.VpcIpamPoolCidr;
 * import com.pulumi.aws.ec2.VpcIpamPoolCidrArgs;
 * import com.pulumi.aws.ec2.inputs.VpcIpamPoolCidrCidrAuthorizationContextArgs;
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
 *         final var current = AwsFunctions.getRegion();
 * 
 *         var example = new VpcIpam(&#34;example&#34;, VpcIpamArgs.builder()        
 *             .operatingRegions(VpcIpamOperatingRegionArgs.builder()
 *                 .regionName(current.applyValue(getRegionResult -&gt; getRegionResult.name()))
 *                 .build())
 *             .build());
 * 
 *         var ipv6TestPublicVpcIpamPool = new VpcIpamPool(&#34;ipv6TestPublicVpcIpamPool&#34;, VpcIpamPoolArgs.builder()        
 *             .addressFamily(&#34;ipv6&#34;)
 *             .ipamScopeId(example.publicDefaultScopeId())
 *             .locale(&#34;us-east-1&#34;)
 *             .description(&#34;public ipv6&#34;)
 *             .advertisable(false)
 *             .awsService(&#34;ec2&#34;)
 *             .build());
 * 
 *         var ipv6TestPublicVpcIpamPoolCidr = new VpcIpamPoolCidr(&#34;ipv6TestPublicVpcIpamPoolCidr&#34;, VpcIpamPoolCidrArgs.builder()        
 *             .ipamPoolId(ipv6TestPublicVpcIpamPool.id())
 *             .cidr(var_.ipv6_cidr())
 *             .cidrAuthorizationContext(VpcIpamPoolCidrCidrAuthorizationContextArgs.builder()
 *                 .message(var_.message())
 *                 .signature(var_.signature())
 *                 .build())
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
 *  to = aws_vpc_ipam_pool_cidr.example
 * 
 *  id = &#34;172.2.0.0/24_ipam-pool-0e634f5a1517cccdc&#34; } Using `pulumi import`, import IPAMs using the `&lt;cidr&gt;_&lt;ipam-pool-id&gt;`. For exampleconsole % pulumi import aws_vpc_ipam_pool_cidr.example 172.2.0.0/24_ipam-pool-0e634f5a1517cccdc
 * 
 */
@ResourceType(type="aws:ec2/vpcIpamPoolCidr:VpcIpamPoolCidr")
public class VpcIpamPoolCidr extends com.pulumi.resources.CustomResource {
    /**
     * The CIDR you want to assign to the pool. Conflicts with `netmask_length`.
     * 
     */
    @Export(name="cidr", refs={String.class}, tree="[0]")
    private Output<String> cidr;

    /**
     * @return The CIDR you want to assign to the pool. Conflicts with `netmask_length`.
     * 
     */
    public Output<String> cidr() {
        return this.cidr;
    }
    /**
     * A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP. This is not stored in the state file. See cidr_authorization_context for more information.
     * 
     */
    @Export(name="cidrAuthorizationContext", refs={VpcIpamPoolCidrCidrAuthorizationContext.class}, tree="[0]")
    private Output</* @Nullable */ VpcIpamPoolCidrCidrAuthorizationContext> cidrAuthorizationContext;

    /**
     * @return A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP. This is not stored in the state file. See cidr_authorization_context for more information.
     * 
     */
    public Output<Optional<VpcIpamPoolCidrCidrAuthorizationContext>> cidrAuthorizationContext() {
        return Codegen.optional(this.cidrAuthorizationContext);
    }
    /**
     * The unique ID generated by AWS for the pool cidr. Typically this is the resource `id` but this attribute was added to the API calls after the fact and is therefore not used as the resource id.
     * 
     */
    @Export(name="ipamPoolCidrId", refs={String.class}, tree="[0]")
    private Output<String> ipamPoolCidrId;

    /**
     * @return The unique ID generated by AWS for the pool cidr. Typically this is the resource `id` but this attribute was added to the API calls after the fact and is therefore not used as the resource id.
     * 
     */
    public Output<String> ipamPoolCidrId() {
        return this.ipamPoolCidrId;
    }
    /**
     * The ID of the pool to which you want to assign a CIDR.
     * 
     */
    @Export(name="ipamPoolId", refs={String.class}, tree="[0]")
    private Output<String> ipamPoolId;

    /**
     * @return The ID of the pool to which you want to assign a CIDR.
     * 
     */
    public Output<String> ipamPoolId() {
        return this.ipamPoolId;
    }
    /**
     * If provided, the cidr provisioned into the specified pool will be the next available cidr given this declared netmask length. Conflicts with `cidr`.
     * 
     */
    @Export(name="netmaskLength", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> netmaskLength;

    /**
     * @return If provided, the cidr provisioned into the specified pool will be the next available cidr given this declared netmask length. Conflicts with `cidr`.
     * 
     */
    public Output<Optional<Integer>> netmaskLength() {
        return Codegen.optional(this.netmaskLength);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcIpamPoolCidr(String name) {
        this(name, VpcIpamPoolCidrArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcIpamPoolCidr(String name, VpcIpamPoolCidrArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcIpamPoolCidr(String name, VpcIpamPoolCidrArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcIpamPoolCidr:VpcIpamPoolCidr", name, args == null ? VpcIpamPoolCidrArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcIpamPoolCidr(String name, Output<String> id, @Nullable VpcIpamPoolCidrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcIpamPoolCidr:VpcIpamPoolCidr", name, state, makeResourceOptions(options, id));
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
    public static VpcIpamPoolCidr get(String name, Output<String> id, @Nullable VpcIpamPoolCidrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcIpamPoolCidr(name, id, state, options);
    }
}
