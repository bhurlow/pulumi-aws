// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.iam;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.iam.SecurityTokenServicePreferencesArgs;
import com.pulumi.aws.iam.inputs.SecurityTokenServicePreferencesState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an IAM Security Token Service Preferences resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.SecurityTokenServicePreferences;
 * import com.pulumi.aws.iam.SecurityTokenServicePreferencesArgs;
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
 *         var example = new SecurityTokenServicePreferences(&#34;example&#34;, SecurityTokenServicePreferencesArgs.builder()        
 *             .globalEndpointTokenVersion(&#34;v2Token&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="aws:iam/securityTokenServicePreferences:SecurityTokenServicePreferences")
public class SecurityTokenServicePreferences extends com.pulumi.resources.CustomResource {
    /**
     * The version of the STS global endpoint token. Valid values: `v1Token`, `v2Token`.
     * 
     */
    @Export(name="globalEndpointTokenVersion", refs={String.class}, tree="[0]")
    private Output<String> globalEndpointTokenVersion;

    /**
     * @return The version of the STS global endpoint token. Valid values: `v1Token`, `v2Token`.
     * 
     */
    public Output<String> globalEndpointTokenVersion() {
        return this.globalEndpointTokenVersion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityTokenServicePreferences(String name) {
        this(name, SecurityTokenServicePreferencesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityTokenServicePreferences(String name, SecurityTokenServicePreferencesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityTokenServicePreferences(String name, SecurityTokenServicePreferencesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iam/securityTokenServicePreferences:SecurityTokenServicePreferences", name, args == null ? SecurityTokenServicePreferencesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecurityTokenServicePreferences(String name, Output<String> id, @Nullable SecurityTokenServicePreferencesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iam/securityTokenServicePreferences:SecurityTokenServicePreferences", name, state, makeResourceOptions(options, id));
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
    public static SecurityTokenServicePreferences get(String name, Output<String> id, @Nullable SecurityTokenServicePreferencesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityTokenServicePreferences(name, id, state, options);
    }
}
