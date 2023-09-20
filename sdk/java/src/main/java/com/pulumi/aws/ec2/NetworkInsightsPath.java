// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.NetworkInsightsPathArgs;
import com.pulumi.aws.ec2.inputs.NetworkInsightsPathState;
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
 * Provides a Network Insights Path resource. Part of the &#34;Reachability Analyzer&#34; service in the AWS VPC console.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.NetworkInsightsPath;
 * import com.pulumi.aws.ec2.NetworkInsightsPathArgs;
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
 *         var test = new NetworkInsightsPath(&#34;test&#34;, NetworkInsightsPathArgs.builder()        
 *             .source(aws_network_interface.source().id())
 *             .destination(aws_network_interface.destination().id())
 *             .protocol(&#34;tcp&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Network Insights Paths using the `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:ec2/networkInsightsPath:NetworkInsightsPath test nip-00edfba169923aefd
 * ```
 * 
 */
@ResourceType(type="aws:ec2/networkInsightsPath:NetworkInsightsPath")
public class NetworkInsightsPath extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the Network Insights Path.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the Network Insights Path.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
     * 
     */
    @Export(name="destination", refs={String.class}, tree="[0]")
    private Output<String> destination;

    /**
     * @return ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }
    /**
     * ARN of the destination.
     * 
     */
    @Export(name="destinationArn", refs={String.class}, tree="[0]")
    private Output<String> destinationArn;

    /**
     * @return ARN of the destination.
     * 
     */
    public Output<String> destinationArn() {
        return this.destinationArn;
    }
    /**
     * IP address of the destination resource.
     * 
     */
    @Export(name="destinationIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationIp;

    /**
     * @return IP address of the destination resource.
     * 
     */
    public Output<Optional<String>> destinationIp() {
        return Codegen.optional(this.destinationIp);
    }
    /**
     * Destination port to analyze access to.
     * 
     */
    @Export(name="destinationPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> destinationPort;

    /**
     * @return Destination port to analyze access to.
     * 
     */
    public Output<Optional<Integer>> destinationPort() {
        return Codegen.optional(this.destinationPort);
    }
    /**
     * Protocol to use for analysis. Valid options are `tcp` or `udp`.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return Protocol to use for analysis. Valid options are `tcp` or `udp`.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * ARN of the source.
     * 
     */
    @Export(name="sourceArn", refs={String.class}, tree="[0]")
    private Output<String> sourceArn;

    /**
     * @return ARN of the source.
     * 
     */
    public Output<String> sourceArn() {
        return this.sourceArn;
    }
    /**
     * IP address of the source resource.
     * 
     */
    @Export(name="sourceIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceIp;

    /**
     * @return IP address of the source resource.
     * 
     */
    public Output<Optional<String>> sourceIp() {
        return Codegen.optional(this.sourceIp);
    }
    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
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
    public NetworkInsightsPath(String name) {
        this(name, NetworkInsightsPathArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkInsightsPath(String name, NetworkInsightsPathArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkInsightsPath(String name, NetworkInsightsPathArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/networkInsightsPath:NetworkInsightsPath", name, args == null ? NetworkInsightsPathArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkInsightsPath(String name, Output<String> id, @Nullable NetworkInsightsPathState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/networkInsightsPath:NetworkInsightsPath", name, state, makeResourceOptions(options, id));
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
    public static NetworkInsightsPath get(String name, Output<String> id, @Nullable NetworkInsightsPathState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkInsightsPath(name, id, state, options);
    }
}
