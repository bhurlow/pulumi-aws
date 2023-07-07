// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.globalaccelerator;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.globalaccelerator.CustomRoutingListenerArgs;
import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingListenerState;
import com.pulumi.aws.globalaccelerator.outputs.CustomRoutingListenerPortRange;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator custom routing listener.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.globalaccelerator.CustomRoutingAccelerator;
 * import com.pulumi.aws.globalaccelerator.CustomRoutingAcceleratorArgs;
 * import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingAcceleratorAttributesArgs;
 * import com.pulumi.aws.globalaccelerator.CustomRoutingListener;
 * import com.pulumi.aws.globalaccelerator.CustomRoutingListenerArgs;
 * import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingListenerPortRangeArgs;
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
 *         var exampleCustomRoutingAccelerator = new CustomRoutingAccelerator(&#34;exampleCustomRoutingAccelerator&#34;, CustomRoutingAcceleratorArgs.builder()        
 *             .name(&#34;Example&#34;)
 *             .ipAddressType(&#34;IPV4&#34;)
 *             .enabled(true)
 *             .attributes(CustomRoutingAcceleratorAttributesArgs.builder()
 *                 .flowLogsEnabled(true)
 *                 .flowLogsS3Bucket(&#34;example-bucket&#34;)
 *                 .flowLogsS3Prefix(&#34;flow-logs/&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleCustomRoutingListener = new CustomRoutingListener(&#34;exampleCustomRoutingListener&#34;, CustomRoutingListenerArgs.builder()        
 *             .acceleratorArn(exampleCustomRoutingAccelerator.id())
 *             .portRanges(CustomRoutingListenerPortRangeArgs.builder()
 *                 .fromPort(80)
 *                 .toPort(80)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Global Accelerator custom routing listeners can be imported using the `id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:globalaccelerator/customRoutingListener:CustomRoutingListener example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
 * ```
 * 
 */
@ResourceType(type="aws:globalaccelerator/customRoutingListener:CustomRoutingListener")
public class CustomRoutingListener extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of a custom routing accelerator.
     * 
     */
    @Export(name="acceleratorArn", refs={String.class}, tree="[0]")
    private Output<String> acceleratorArn;

    /**
     * @return The Amazon Resource Name (ARN) of a custom routing accelerator.
     * 
     */
    public Output<String> acceleratorArn() {
        return this.acceleratorArn;
    }
    /**
     * The list of port ranges for the connections from clients to the accelerator. Fields documented below.
     * 
     * **port_range** supports the following attributes:
     * 
     */
    @Export(name="portRanges", refs={List.class,CustomRoutingListenerPortRange.class}, tree="[0,1]")
    private Output<List<CustomRoutingListenerPortRange>> portRanges;

    /**
     * @return The list of port ranges for the connections from clients to the accelerator. Fields documented below.
     * 
     * **port_range** supports the following attributes:
     * 
     */
    public Output<List<CustomRoutingListenerPortRange>> portRanges() {
        return this.portRanges;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomRoutingListener(String name) {
        this(name, CustomRoutingListenerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomRoutingListener(String name, CustomRoutingListenerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomRoutingListener(String name, CustomRoutingListenerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:globalaccelerator/customRoutingListener:CustomRoutingListener", name, args == null ? CustomRoutingListenerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomRoutingListener(String name, Output<String> id, @Nullable CustomRoutingListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:globalaccelerator/customRoutingListener:CustomRoutingListener", name, state, makeResourceOptions(options, id));
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
    public static CustomRoutingListener get(String name, Output<String> id, @Nullable CustomRoutingListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomRoutingListener(name, id, state, options);
    }
}
