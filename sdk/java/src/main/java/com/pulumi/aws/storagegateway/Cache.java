// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.storagegateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.storagegateway.CacheArgs;
import com.pulumi.aws.storagegateway.inputs.CacheState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages an AWS Storage Gateway cache.
 * 
 * &gt; **NOTE:** The Storage Gateway API provides no method to remove a cache disk. Destroying this resource does not perform any Storage Gateway actions.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.storagegateway.Cache;
 * import com.pulumi.aws.storagegateway.CacheArgs;
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
 *         var example = new Cache(&#34;example&#34;, CacheArgs.builder()        
 *             .diskId(data.aws_storagegateway_local_disk().example().id())
 *             .gatewayArn(aws_storagegateway_gateway.example().arn())
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
 *  to = aws_storagegateway_cache.example
 * 
 *  id = &#34;arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0&#34; } Using `pulumi import`, import `aws_storagegateway_cache` using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`). For exampleconsole % pulumi import aws_storagegateway_cache.example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
 * 
 */
@ResourceType(type="aws:storagegateway/cache:Cache")
public class Cache extends com.pulumi.resources.CustomResource {
    /**
     * Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     * 
     */
    @Export(name="diskId", refs={String.class}, tree="[0]")
    private Output<String> diskId;

    /**
     * @return Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     * 
     */
    public Output<String> diskId() {
        return this.diskId;
    }
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     * 
     */
    @Export(name="gatewayArn", refs={String.class}, tree="[0]")
    private Output<String> gatewayArn;

    /**
     * @return The Amazon Resource Name (ARN) of the gateway.
     * 
     */
    public Output<String> gatewayArn() {
        return this.gatewayArn;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cache(String name) {
        this(name, CacheArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cache(String name, CacheArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cache(String name, CacheArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:storagegateway/cache:Cache", name, args == null ? CacheArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cache(String name, Output<String> id, @Nullable CacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:storagegateway/cache:Cache", name, state, makeResourceOptions(options, id));
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
    public static Cache get(String name, Output<String> id, @Nullable CacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cache(name, id, state, options);
    }
}
