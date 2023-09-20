// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.quicksight.ThemeArgs;
import com.pulumi.aws.quicksight.inputs.ThemeState;
import com.pulumi.aws.quicksight.outputs.ThemeConfiguration;
import com.pulumi.aws.quicksight.outputs.ThemePermission;
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
 * Resource for managing a QuickSight Theme.
 * 
 * ## Example Usage
 * ### Basic Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.Theme;
 * import com.pulumi.aws.quicksight.ThemeArgs;
 * import com.pulumi.aws.quicksight.inputs.ThemeConfigurationArgs;
 * import com.pulumi.aws.quicksight.inputs.ThemeConfigurationDataColorPaletteArgs;
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
 *         var example = new Theme(&#34;example&#34;, ThemeArgs.builder()        
 *             .baseThemeId(&#34;MIDNIGHT&#34;)
 *             .configuration(ThemeConfigurationArgs.builder()
 *                 .dataColorPalette(ThemeConfigurationDataColorPaletteArgs.builder()
 *                     .colors(                    
 *                         &#34;#FFFFFF&#34;,
 *                         &#34;#111111&#34;,
 *                         &#34;#222222&#34;,
 *                         &#34;#333333&#34;,
 *                         &#34;#444444&#34;,
 *                         &#34;#555555&#34;,
 *                         &#34;#666666&#34;,
 *                         &#34;#777777&#34;,
 *                         &#34;#888888&#34;,
 *                         &#34;#999999&#34;)
 *                     .emptyFillColor(&#34;#FFFFFF&#34;)
 *                     .minMaxGradient(                    
 *                         &#34;#FFFFFF&#34;,
 *                         &#34;#111111&#34;)
 *                     .build())
 *                 .build())
 *             .themeId(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import a QuickSight Theme using the AWS account ID and theme ID separated by a comma (`,`). For example:
 * 
 * ```sh
 *  $ pulumi import aws:quicksight/theme:Theme example 123456789012,example-id
 * ```
 * 
 */
@ResourceType(type="aws:quicksight/theme:Theme")
public class Theme extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the theme.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the theme.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * AWS account ID.
     * 
     */
    @Export(name="awsAccountId", refs={String.class}, tree="[0]")
    private Output<String> awsAccountId;

    /**
     * @return AWS account ID.
     * 
     */
    public Output<String> awsAccountId() {
        return this.awsAccountId;
    }
    /**
     * The ID of the theme that a custom theme will inherit from. All themes inherit from one of the starting themes defined by Amazon QuickSight. For a list of the starting themes, use ListThemes or choose Themes from within an analysis.
     * 
     */
    @Export(name="baseThemeId", refs={String.class}, tree="[0]")
    private Output<String> baseThemeId;

    /**
     * @return The ID of the theme that a custom theme will inherit from. All themes inherit from one of the starting themes defined by Amazon QuickSight. For a list of the starting themes, use ListThemes or choose Themes from within an analysis.
     * 
     */
    public Output<String> baseThemeId() {
        return this.baseThemeId;
    }
    /**
     * The theme configuration, which contains the theme display properties. See configuration.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="configuration", refs={ThemeConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ ThemeConfiguration> configuration;

    /**
     * @return The theme configuration, which contains the theme display properties. See configuration.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<Optional<ThemeConfiguration>> configuration() {
        return Codegen.optional(this.configuration);
    }
    /**
     * The time that the theme was created.
     * 
     */
    @Export(name="createdTime", refs={String.class}, tree="[0]")
    private Output<String> createdTime;

    /**
     * @return The time that the theme was created.
     * 
     */
    public Output<String> createdTime() {
        return this.createdTime;
    }
    /**
     * The time that the theme was last updated.
     * 
     */
    @Export(name="lastUpdatedTime", refs={String.class}, tree="[0]")
    private Output<String> lastUpdatedTime;

    /**
     * @return The time that the theme was last updated.
     * 
     */
    public Output<String> lastUpdatedTime() {
        return this.lastUpdatedTime;
    }
    /**
     * Display name of the theme.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name of the theme.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A set of resource permissions on the theme. Maximum of 64 items. See permissions.
     * 
     */
    @Export(name="permissions", refs={List.class,ThemePermission.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ThemePermission>> permissions;

    /**
     * @return A set of resource permissions on the theme. Maximum of 64 items. See permissions.
     * 
     */
    public Output<Optional<List<ThemePermission>>> permissions() {
        return Codegen.optional(this.permissions);
    }
    /**
     * The theme creation status.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The theme creation status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Identifier of the theme.
     * 
     */
    @Export(name="themeId", refs={String.class}, tree="[0]")
    private Output<String> themeId;

    /**
     * @return Identifier of the theme.
     * 
     */
    public Output<String> themeId() {
        return this.themeId;
    }
    /**
     * A description of the current theme version being created/updated.
     * 
     */
    @Export(name="versionDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> versionDescription;

    /**
     * @return A description of the current theme version being created/updated.
     * 
     */
    public Output<Optional<String>> versionDescription() {
        return Codegen.optional(this.versionDescription);
    }
    /**
     * The version number of the theme version.
     * 
     */
    @Export(name="versionNumber", refs={Integer.class}, tree="[0]")
    private Output<Integer> versionNumber;

    /**
     * @return The version number of the theme version.
     * 
     */
    public Output<Integer> versionNumber() {
        return this.versionNumber;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Theme(String name) {
        this(name, ThemeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Theme(String name, ThemeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Theme(String name, ThemeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/theme:Theme", name, args == null ? ThemeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Theme(String name, Output<String> id, @Nullable ThemeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/theme:Theme", name, state, makeResourceOptions(options, id));
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
    public static Theme get(String name, Output<String> id, @Nullable ThemeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Theme(name, id, state, options);
    }
}
