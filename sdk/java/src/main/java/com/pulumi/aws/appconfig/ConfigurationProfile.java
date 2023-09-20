// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appconfig;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.appconfig.ConfigurationProfileArgs;
import com.pulumi.aws.appconfig.inputs.ConfigurationProfileState;
import com.pulumi.aws.appconfig.outputs.ConfigurationProfileValidator;
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
 * Provides an AppConfig Configuration Profile resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.appconfig.ConfigurationProfile;
 * import com.pulumi.aws.appconfig.ConfigurationProfileArgs;
 * import com.pulumi.aws.appconfig.inputs.ConfigurationProfileValidatorArgs;
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
 *         var example = new ConfigurationProfile(&#34;example&#34;, ConfigurationProfileArgs.builder()        
 *             .applicationId(aws_appconfig_application.example().id())
 *             .description(&#34;Example Configuration Profile&#34;)
 *             .locationUri(&#34;hosted&#34;)
 *             .validators(ConfigurationProfileValidatorArgs.builder()
 *                 .content(aws_lambda_function.example().arn())
 *                 .type(&#34;LAMBDA&#34;)
 *                 .build())
 *             .tags(Map.of(&#34;Type&#34;, &#34;AppConfig Configuration Profile&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import AppConfig Configuration Profiles using the configuration profile ID and application ID separated by a colon (`:`). For example:
 * 
 * ```sh
 *  $ pulumi import aws:appconfig/configurationProfile:ConfigurationProfile example 71abcde:11xxxxx
 * ```
 * 
 */
@ResourceType(type="aws:appconfig/configurationProfile:ConfigurationProfile")
public class ConfigurationProfile extends com.pulumi.resources.CustomResource {
    /**
     * Application ID. Must be between 4 and 7 characters in length.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return Application ID. Must be between 4 and 7 characters in length.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * ARN of the AppConfig Configuration Profile.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the AppConfig Configuration Profile.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The configuration profile ID.
     * 
     */
    @Export(name="configurationProfileId", refs={String.class}, tree="[0]")
    private Output<String> configurationProfileId;

    /**
     * @return The configuration profile ID.
     * 
     */
    public Output<String> configurationProfileId() {
        return this.configurationProfileId;
    }
    /**
     * Description of the configuration profile. Can be at most 1024 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the configuration profile. Can be at most 1024 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://&lt;Document_name&gt;` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://&lt;Parameter_name&gt;` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://&lt;bucket&gt;/&lt;objectKey&gt;`.
     * 
     */
    @Export(name="locationUri", refs={String.class}, tree="[0]")
    private Output<String> locationUri;

    /**
     * @return URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://&lt;Document_name&gt;` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://&lt;Parameter_name&gt;` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://&lt;bucket&gt;/&lt;objectKey&gt;`.
     * 
     */
    public Output<String> locationUri() {
        return this.locationUri;
    }
    /**
     * Name for the configuration profile. Must be between 1 and 64 characters in length.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name for the configuration profile. Must be between 1 and 64 characters in length.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
     * 
     */
    @Export(name="retrievalRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> retrievalRoleArn;

    /**
     * @return ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
     * 
     */
    public Output<Optional<String>> retrievalRoleArn() {
        return Codegen.optional(this.retrievalRoleArn);
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
     * Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
     * 
     */
    @Export(name="validators", refs={List.class,ConfigurationProfileValidator.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ConfigurationProfileValidator>> validators;

    /**
     * @return Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
     * 
     */
    public Output<Optional<List<ConfigurationProfileValidator>>> validators() {
        return Codegen.optional(this.validators);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ConfigurationProfile(String name) {
        this(name, ConfigurationProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConfigurationProfile(String name, ConfigurationProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConfigurationProfile(String name, ConfigurationProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appconfig/configurationProfile:ConfigurationProfile", name, args == null ? ConfigurationProfileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConfigurationProfile(String name, Output<String> id, @Nullable ConfigurationProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appconfig/configurationProfile:ConfigurationProfile", name, state, makeResourceOptions(options, id));
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
    public static ConfigurationProfile get(String name, Output<String> id, @Nullable ConfigurationProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConfigurationProfile(name, id, state, options);
    }
}
