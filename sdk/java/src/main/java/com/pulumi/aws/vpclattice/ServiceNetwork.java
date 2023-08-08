// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpclattice;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.vpclattice.ServiceNetworkArgs;
import com.pulumi.aws.vpclattice.inputs.ServiceNetworkState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS VPC Lattice Service Network.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.vpclattice.ServiceNetwork;
 * import com.pulumi.aws.vpclattice.ServiceNetworkArgs;
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
 *         var example = new ServiceNetwork(&#34;example&#34;, ServiceNetworkArgs.builder()        
 *             .authType(&#34;AWS_IAM&#34;)
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
 *  to = aws_vpclattice_service_network.example
 * 
 *  id = &#34;sn-0158f91c1e3358dba&#34; } Using `pulumi import`, import VPC Lattice Service Network using the `id`. For exampleconsole % pulumi import aws_vpclattice_service_network.example sn-0158f91c1e3358dba
 * 
 */
@ResourceType(type="aws:vpclattice/serviceNetwork:ServiceNetwork")
public class ServiceNetwork extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the Service Network.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the Service Network.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Type of IAM policy. Either `NONE` or `AWS_IAM`.
     * 
     */
    @Export(name="authType", refs={String.class}, tree="[0]")
    private Output<String> authType;

    /**
     * @return Type of IAM policy. Either `NONE` or `AWS_IAM`.
     * 
     */
    public Output<String> authType() {
        return this.authType;
    }
    /**
     * Name of the service network
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the service network
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceNetwork(String name) {
        this(name, ServiceNetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceNetwork(String name, @Nullable ServiceNetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceNetwork(String name, @Nullable ServiceNetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:vpclattice/serviceNetwork:ServiceNetwork", name, args == null ? ServiceNetworkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceNetwork(String name, Output<String> id, @Nullable ServiceNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:vpclattice/serviceNetwork:ServiceNetwork", name, state, makeResourceOptions(options, id));
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
    public static ServiceNetwork get(String name, Output<String> id, @Nullable ServiceNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceNetwork(name, id, state, options);
    }
}
