// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.account;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.account.RegionArgs;
import com.pulumi.aws.account.inputs.RegionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Enable (Opt-In) or Disable (Opt-Out) a particular Region for an AWS account.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.account.Region;
 * import com.pulumi.aws.account.RegionArgs;
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
 *         var example = new Region(&#34;example&#34;, RegionArgs.builder()        
 *             .regionName(&#34;ap-southeast-3&#34;)
 *             .enabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Using `pulumi import`. For example:
 * 
 * ```sh
 * $ pulumi import aws:account/region:Region example ap-southeast-3
 * ```
 * 
 */
@ResourceType(type="aws:account/region:Region")
public class Region extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the target account when managing member accounts. Will manage current user&#39;s account by default if omitted. To use this parameter, the caller must be an identity in the organization&#39;s management account or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have all features enabled, and the organization must have trusted access enabled for the Account Management service, and optionally a delegated admin account assigned.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountId;

    /**
     * @return The ID of the target account when managing member accounts. Will manage current user&#39;s account by default if omitted. To use this parameter, the caller must be an identity in the organization&#39;s management account or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have all features enabled, and the organization must have trusted access enabled for the Account Management service, and optionally a delegated admin account assigned.
     * 
     */
    public Output<Optional<String>> accountId() {
        return Codegen.optional(this.accountId);
    }
    /**
     * Whether the region is enabled.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Whether the region is enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The region opt status.
     * 
     */
    @Export(name="optStatus", refs={String.class}, tree="[0]")
    private Output<String> optStatus;

    /**
     * @return The region opt status.
     * 
     */
    public Output<String> optStatus() {
        return this.optStatus;
    }
    /**
     * The region name to manage.
     * 
     */
    @Export(name="regionName", refs={String.class}, tree="[0]")
    private Output<String> regionName;

    /**
     * @return The region name to manage.
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Region(String name) {
        this(name, RegionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Region(String name, RegionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Region(String name, RegionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:account/region:Region", name, args == null ? RegionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Region(String name, Output<String> id, @Nullable RegionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:account/region:Region", name, state, makeResourceOptions(options, id));
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
    public static Region get(String name, Output<String> id, @Nullable RegionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Region(name, id, state, options);
    }
}
