// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcDhcpOptionsArgs;
import com.pulumi.aws.ec2.inputs.VpcDhcpOptionsState;
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
 * Provides a VPC DHCP Options resource.
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
 * import com.pulumi.aws.ec2.VpcDhcpOptions;
 * import com.pulumi.aws.ec2.VpcDhcpOptionsArgs;
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
 *         var dnsResolver = new VpcDhcpOptions(&#34;dnsResolver&#34;, VpcDhcpOptionsArgs.builder()        
 *             .domainNameServers(            
 *                 &#34;8.8.8.8&#34;,
 *                 &#34;8.8.4.4&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Full usage:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcDhcpOptions;
 * import com.pulumi.aws.ec2.VpcDhcpOptionsArgs;
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
 *         var foo = new VpcDhcpOptions(&#34;foo&#34;, VpcDhcpOptionsArgs.builder()        
 *             .domainName(&#34;service.consul&#34;)
 *             .domainNameServers(            
 *                 &#34;127.0.0.1&#34;,
 *                 &#34;10.0.0.2&#34;)
 *             .netbiosNameServers(&#34;127.0.0.1&#34;)
 *             .netbiosNodeType(2)
 *             .ntpServers(&#34;127.0.0.1&#34;)
 *             .tags(Map.of(&#34;Name&#34;, &#34;foo-name&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Remarks
 * 
 * * Notice that all arguments are optional but you have to specify at least one argument.
 * * `domain_name_servers`, `netbios_name_servers`, `ntp_servers` are limited by AWS to maximum four servers only.
 * * To actually use the DHCP Options Set you need to associate it to a VPC using `aws.ec2.VpcDhcpOptionsAssociation`.
 * * If you delete a DHCP Options Set, all VPCs using it will be associated to AWS&#39;s `default` DHCP Option Set.
 * * In most cases unless you&#39;re configuring your own DNS you&#39;ll want to set `domain_name_servers` to `AmazonProvidedDNS`.
 * 
 * ## Import
 * 
 * Using `pulumi import`, import VPC DHCP Options using the DHCP Options `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:ec2/vpcDhcpOptions:VpcDhcpOptions my_options dopt-d9070ebb
 * ```
 * 
 */
@ResourceType(type="aws:ec2/vpcDhcpOptions:VpcDhcpOptions")
public class VpcDhcpOptions extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the DHCP Options Set.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the DHCP Options Set.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> domainName;

    /**
     * @return the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
     * 
     */
    public Output<Optional<String>> domainName() {
        return Codegen.optional(this.domainName);
    }
    /**
     * List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
     * 
     */
    @Export(name="domainNameServers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> domainNameServers;

    /**
     * @return List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
     * 
     */
    public Output<Optional<List<String>>> domainNameServers() {
        return Codegen.optional(this.domainNameServers);
    }
    /**
     * List of NETBIOS name servers.
     * 
     */
    @Export(name="netbiosNameServers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> netbiosNameServers;

    /**
     * @return List of NETBIOS name servers.
     * 
     */
    public Output<Optional<List<String>>> netbiosNameServers() {
        return Codegen.optional(this.netbiosNameServers);
    }
    /**
     * The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
     * 
     */
    @Export(name="netbiosNodeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> netbiosNodeType;

    /**
     * @return The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
     * 
     */
    public Output<Optional<String>> netbiosNodeType() {
        return Codegen.optional(this.netbiosNodeType);
    }
    /**
     * List of NTP servers to configure.
     * 
     */
    @Export(name="ntpServers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ntpServers;

    /**
     * @return List of NTP servers to configure.
     * 
     */
    public Output<Optional<List<String>>> ntpServers() {
        return Codegen.optional(this.ntpServers);
    }
    /**
     * The ID of the AWS account that owns the DHCP options set.
     * 
     */
    @Export(name="ownerId", refs={String.class}, tree="[0]")
    private Output<String> ownerId;

    /**
     * @return The ID of the AWS account that owns the DHCP options set.
     * 
     */
    public Output<String> ownerId() {
        return this.ownerId;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcDhcpOptions(String name) {
        this(name, VpcDhcpOptionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcDhcpOptions(String name, @Nullable VpcDhcpOptionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcDhcpOptions(String name, @Nullable VpcDhcpOptionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcDhcpOptions:VpcDhcpOptions", name, args == null ? VpcDhcpOptionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcDhcpOptions(String name, Output<String> id, @Nullable VpcDhcpOptionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcDhcpOptions:VpcDhcpOptions", name, state, makeResourceOptions(options, id));
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
    public static VpcDhcpOptions get(String name, Output<String> id, @Nullable VpcDhcpOptionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcDhcpOptions(name, id, state, options);
    }
}
