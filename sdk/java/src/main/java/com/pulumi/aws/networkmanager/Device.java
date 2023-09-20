// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.networkmanager;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.networkmanager.DeviceArgs;
import com.pulumi.aws.networkmanager.inputs.DeviceState;
import com.pulumi.aws.networkmanager.outputs.DeviceAwsLocation;
import com.pulumi.aws.networkmanager.outputs.DeviceLocation;
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
 * Creates a device in a global network. If you specify both a site ID and a location,
 * the location of the site is used for visualization in the Network Manager console.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.networkmanager.Device;
 * import com.pulumi.aws.networkmanager.DeviceArgs;
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
 *         var example = new Device(&#34;example&#34;, DeviceArgs.builder()        
 *             .globalNetworkId(aws_networkmanager_global_network.example().id())
 *             .siteId(aws_networkmanager_site.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import `aws_networkmanager_device` using the device ARN. For example:
 * 
 * ```sh
 *  $ pulumi import aws:networkmanager/device:Device example arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/device-07f6fd08867abc123
 * ```
 * 
 */
@ResourceType(type="aws:networkmanager/device:Device")
public class Device extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the device.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the device.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The AWS location of the device. Documented below.
     * 
     */
    @Export(name="awsLocation", refs={DeviceAwsLocation.class}, tree="[0]")
    private Output</* @Nullable */ DeviceAwsLocation> awsLocation;

    /**
     * @return The AWS location of the device. Documented below.
     * 
     */
    public Output<Optional<DeviceAwsLocation>> awsLocation() {
        return Codegen.optional(this.awsLocation);
    }
    /**
     * A description of the device.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the device.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID of the global network.
     * 
     */
    @Export(name="globalNetworkId", refs={String.class}, tree="[0]")
    private Output<String> globalNetworkId;

    /**
     * @return The ID of the global network.
     * 
     */
    public Output<String> globalNetworkId() {
        return this.globalNetworkId;
    }
    /**
     * The location of the device. Documented below.
     * 
     */
    @Export(name="location", refs={DeviceLocation.class}, tree="[0]")
    private Output</* @Nullable */ DeviceLocation> location;

    /**
     * @return The location of the device. Documented below.
     * 
     */
    public Output<Optional<DeviceLocation>> location() {
        return Codegen.optional(this.location);
    }
    /**
     * The model of device.
     * 
     */
    @Export(name="model", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> model;

    /**
     * @return The model of device.
     * 
     */
    public Output<Optional<String>> model() {
        return Codegen.optional(this.model);
    }
    /**
     * The serial number of the device.
     * 
     */
    @Export(name="serialNumber", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serialNumber;

    /**
     * @return The serial number of the device.
     * 
     */
    public Output<Optional<String>> serialNumber() {
        return Codegen.optional(this.serialNumber);
    }
    /**
     * The ID of the site.
     * 
     */
    @Export(name="siteId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> siteId;

    /**
     * @return The ID of the site.
     * 
     */
    public Output<Optional<String>> siteId() {
        return Codegen.optional(this.siteId);
    }
    /**
     * Key-value tags for the device. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value tags for the device. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The type of device.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of device.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * The vendor of the device.
     * 
     */
    @Export(name="vendor", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vendor;

    /**
     * @return The vendor of the device.
     * 
     */
    public Output<Optional<String>> vendor() {
        return Codegen.optional(this.vendor);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Device(String name) {
        this(name, DeviceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Device(String name, DeviceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Device(String name, DeviceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:networkmanager/device:Device", name, args == null ? DeviceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Device(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:networkmanager/device:Device", name, state, makeResourceOptions(options, id));
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
    public static Device get(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Device(name, id, state, options);
    }
}
