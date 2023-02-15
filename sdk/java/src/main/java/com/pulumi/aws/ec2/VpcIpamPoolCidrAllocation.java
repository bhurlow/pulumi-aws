// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocationArgs;
import com.pulumi.aws.ec2.inputs.VpcIpamPoolCidrAllocationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allocates (reserves) a CIDR from an IPAM address pool, preventing usage by IPAM. Only works for private IPv4.
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
 * import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocation;
 * import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocationArgs;
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
 *         var exampleVpcIpamPoolCidrAllocation = new VpcIpamPoolCidrAllocation(&#34;exampleVpcIpamPoolCidrAllocation&#34;, VpcIpamPoolCidrAllocationArgs.builder()        
 *             .ipamPoolId(exampleVpcIpamPool.id())
 *             .cidr(&#34;172.2.0.0/24&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(exampleVpcIpamPoolCidr)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * With the `disallowed_cidrs` attribute:
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
 * import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocation;
 * import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocationArgs;
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
 *         var exampleVpcIpamPoolCidrAllocation = new VpcIpamPoolCidrAllocation(&#34;exampleVpcIpamPoolCidrAllocation&#34;, VpcIpamPoolCidrAllocationArgs.builder()        
 *             .ipamPoolId(exampleVpcIpamPool.id())
 *             .netmaskLength(28)
 *             .disallowedCidrs(&#34;172.2.0.0/28&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(exampleVpcIpamPoolCidr)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * IPAM allocations can be imported using the `allocation id` and `pool id`, separated by `_`, e.g.
 * 
 * ```sh
 *  $ pulumi import aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation example ipam-pool-alloc-0dc6d196509c049ba8b549ff99f639736_ipam-pool-07cfb559e0921fcbe
 * ```
 * 
 */
@ResourceType(type="aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation")
public class VpcIpamPoolCidrAllocation extends com.pulumi.resources.CustomResource {
    /**
     * The CIDR you want to assign to the pool.
     * 
     */
    @Export(name="cidr", refs={String.class}, tree="[0]")
    private Output<String> cidr;

    /**
     * @return The CIDR you want to assign to the pool.
     * 
     */
    public Output<String> cidr() {
        return this.cidr;
    }
    /**
     * The description for the allocation.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description for the allocation.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Exclude a particular CIDR range from being returned by the pool.
     * 
     */
    @Export(name="disallowedCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> disallowedCidrs;

    /**
     * @return Exclude a particular CIDR range from being returned by the pool.
     * 
     */
    public Output<Optional<List<String>>> disallowedCidrs() {
        return Codegen.optional(this.disallowedCidrs);
    }
    @Export(name="ipamPoolAllocationId", refs={String.class}, tree="[0]")
    private Output<String> ipamPoolAllocationId;

    public Output<String> ipamPoolAllocationId() {
        return this.ipamPoolAllocationId;
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
     * The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
     * 
     */
    @Export(name="netmaskLength", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> netmaskLength;

    /**
     * @return The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
     * 
     */
    public Output<Optional<Integer>> netmaskLength() {
        return Codegen.optional(this.netmaskLength);
    }
    /**
     * The ID of the resource.
     * 
     */
    @Export(name="resourceId", refs={String.class}, tree="[0]")
    private Output<String> resourceId;

    /**
     * @return The ID of the resource.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * The owner of the resource.
     * 
     */
    @Export(name="resourceOwner", refs={String.class}, tree="[0]")
    private Output<String> resourceOwner;

    /**
     * @return The owner of the resource.
     * 
     */
    public Output<String> resourceOwner() {
        return this.resourceOwner;
    }
    /**
     * The type of the resource.
     * 
     */
    @Export(name="resourceType", refs={String.class}, tree="[0]")
    private Output<String> resourceType;

    /**
     * @return The type of the resource.
     * 
     */
    public Output<String> resourceType() {
        return this.resourceType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcIpamPoolCidrAllocation(String name) {
        this(name, VpcIpamPoolCidrAllocationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcIpamPoolCidrAllocation(String name, VpcIpamPoolCidrAllocationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcIpamPoolCidrAllocation(String name, VpcIpamPoolCidrAllocationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation", name, args == null ? VpcIpamPoolCidrAllocationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcIpamPoolCidrAllocation(String name, Output<String> id, @Nullable VpcIpamPoolCidrAllocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation", name, state, makeResourceOptions(options, id));
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
    public static VpcIpamPoolCidrAllocation get(String name, Output<String> id, @Nullable VpcIpamPoolCidrAllocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcIpamPoolCidrAllocation(name, id, state, options);
    }
}
