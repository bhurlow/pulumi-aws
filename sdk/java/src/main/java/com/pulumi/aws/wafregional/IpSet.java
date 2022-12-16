// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafregional;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.wafregional.IpSetArgs;
import com.pulumi.aws.wafregional.inputs.IpSetState;
import com.pulumi.aws.wafregional.outputs.IpSetIpSetDescriptor;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a WAF Regional IPSet Resource for use with Application Load Balancer.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.wafregional.IpSet;
 * import com.pulumi.aws.wafregional.IpSetArgs;
 * import com.pulumi.aws.wafregional.inputs.IpSetIpSetDescriptorArgs;
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
 *         var ipset = new IpSet(&#34;ipset&#34;, IpSetArgs.builder()        
 *             .ipSetDescriptors(            
 *                 IpSetIpSetDescriptorArgs.builder()
 *                     .type(&#34;IPV4&#34;)
 *                     .value(&#34;192.0.7.0/24&#34;)
 *                     .build(),
 *                 IpSetIpSetDescriptorArgs.builder()
 *                     .type(&#34;IPV4&#34;)
 *                     .value(&#34;10.16.16.0/16&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * WAF Regional IPSets can be imported using their ID, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:wafregional/ipSet:IpSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
 * ```
 * 
 */
@ResourceType(type="aws:wafregional/ipSet:IpSet")
public class IpSet extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the WAF IPSet.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the WAF IPSet.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
     * 
     */
    @Export(name="ipSetDescriptors", refs={List.class,IpSetIpSetDescriptor.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IpSetIpSetDescriptor>> ipSetDescriptors;

    /**
     * @return One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
     * 
     */
    public Output<Optional<List<IpSetIpSetDescriptor>>> ipSetDescriptors() {
        return Codegen.optional(this.ipSetDescriptors);
    }
    /**
     * The name or description of the IPSet.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name or description of the IPSet.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpSet(String name) {
        this(name, IpSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpSet(String name, @Nullable IpSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpSet(String name, @Nullable IpSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:wafregional/ipSet:IpSet", name, args == null ? IpSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpSet(String name, Output<String> id, @Nullable IpSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:wafregional/ipSet:IpSet", name, state, makeResourceOptions(options, id));
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
    public static IpSet get(String name, Output<String> id, @Nullable IpSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpSet(name, id, state, options);
    }
}
