// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lightsail;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.lightsail.LbArgs;
import com.pulumi.aws.lightsail.inputs.LbState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a Lightsail load balancer resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.Lb;
 * import com.pulumi.aws.lightsail.LbArgs;
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
 *         var test = new Lb(&#34;test&#34;, LbArgs.builder()        
 *             .healthCheckPath(&#34;/&#34;)
 *             .instancePort(&#34;80&#34;)
 *             .tags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
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
 *  to = aws_lightsail_lb.test
 * 
 *  id = &#34;example-load-balancer&#34; } Using `pulumi import`, import `aws_lightsail_lb` using the name attribute. For exampleconsole % pulumi import aws_lightsail_lb.test example-load-balancer
 * 
 */
@ResourceType(type="aws:lightsail/lb:Lb")
public class Lb extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the Lightsail load balancer.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the Lightsail load balancer.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The timestamp when the load balancer was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The timestamp when the load balancer was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The DNS name of the load balancer.
     * 
     */
    @Export(name="dnsName", refs={String.class}, tree="[0]")
    private Output<String> dnsName;

    /**
     * @return The DNS name of the load balancer.
     * 
     */
    public Output<String> dnsName() {
        return this.dnsName;
    }
    /**
     * The health check path of the load balancer. Default value &#34;/&#34;.
     * 
     */
    @Export(name="healthCheckPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckPath;

    /**
     * @return The health check path of the load balancer. Default value &#34;/&#34;.
     * 
     */
    public Output<Optional<String>> healthCheckPath() {
        return Codegen.optional(this.healthCheckPath);
    }
    /**
     * The instance port the load balancer will connect.
     * 
     */
    @Export(name="instancePort", refs={Integer.class}, tree="[0]")
    private Output<Integer> instancePort;

    /**
     * @return The instance port the load balancer will connect.
     * 
     */
    public Output<Integer> instancePort() {
        return this.instancePort;
    }
    @Export(name="ipAddressType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipAddressType;

    public Output<Optional<String>> ipAddressType() {
        return Codegen.optional(this.ipAddressType);
    }
    /**
     * The name of the Lightsail load balancer.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Lightsail load balancer.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The protocol of the load balancer.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The protocol of the load balancer.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The public ports of the load balancer.
     * 
     */
    @Export(name="publicPorts", refs={List.class,Integer.class}, tree="[0,1]")
    private Output<List<Integer>> publicPorts;

    /**
     * @return The public ports of the load balancer.
     * 
     */
    public Output<List<Integer>> publicPorts() {
        return this.publicPorts;
    }
    /**
     * The support code for the database. Include this code in your email to support when you have questions about a database in Lightsail. This code enables our support team to look up your Lightsail information more easily.
     * 
     */
    @Export(name="supportCode", refs={String.class}, tree="[0]")
    private Output<String> supportCode;

    /**
     * @return The support code for the database. Include this code in your email to support when you have questions about a database in Lightsail. This code enables our support team to look up your Lightsail information more easily.
     * 
     */
    public Output<String> supportCode() {
        return this.supportCode;
    }
    /**
     * A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Lb(String name) {
        this(name, LbArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Lb(String name, LbArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Lb(String name, LbArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/lb:Lb", name, args == null ? LbArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Lb(String name, Output<String> id, @Nullable LbState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/lb:Lb", name, state, makeResourceOptions(options, id));
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
    public static Lb get(String name, Output<String> id, @Nullable LbState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Lb(name, id, state, options);
    }
}
