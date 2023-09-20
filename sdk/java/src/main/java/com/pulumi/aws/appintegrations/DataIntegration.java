// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appintegrations;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.appintegrations.DataIntegrationArgs;
import com.pulumi.aws.appintegrations.inputs.DataIntegrationState;
import com.pulumi.aws.appintegrations.outputs.DataIntegrationScheduleConfig;
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
 * Provides an Amazon AppIntegrations Data Integration resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.appintegrations.DataIntegration;
 * import com.pulumi.aws.appintegrations.DataIntegrationArgs;
 * import com.pulumi.aws.appintegrations.inputs.DataIntegrationScheduleConfigArgs;
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
 *         var example = new DataIntegration(&#34;example&#34;, DataIntegrationArgs.builder()        
 *             .description(&#34;example&#34;)
 *             .kmsKey(aws_kms_key.test().arn())
 *             .sourceUri(&#34;Salesforce://AppFlow/example&#34;)
 *             .scheduleConfig(DataIntegrationScheduleConfigArgs.builder()
 *                 .firstExecutionFrom(&#34;1439788442681&#34;)
 *                 .object(&#34;Account&#34;)
 *                 .scheduleExpression(&#34;rate(1 hour)&#34;)
 *                 .build())
 *             .tags(Map.of(&#34;Key1&#34;, &#34;Value1&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Amazon AppIntegrations Data Integrations using the `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:appintegrations/dataIntegration:DataIntegration example 12345678-1234-1234-1234-123456789123
 * ```
 * 
 */
@ResourceType(type="aws:appintegrations/dataIntegration:DataIntegration")
public class DataIntegration extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the Data Integration.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the Data Integration.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Specifies the description of the Data Integration.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Specifies the description of the Data Integration.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
     * 
     */
    @Export(name="kmsKey", refs={String.class}, tree="[0]")
    private Output<String> kmsKey;

    /**
     * @return Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
     * 
     */
    public Output<String> kmsKey() {
        return this.kmsKey;
    }
    /**
     * Specifies the name of the Data Integration.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the name of the Data Integration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
     * 
     */
    @Export(name="scheduleConfig", refs={DataIntegrationScheduleConfig.class}, tree="[0]")
    private Output<DataIntegrationScheduleConfig> scheduleConfig;

    /**
     * @return A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
     * 
     */
    public Output<DataIntegrationScheduleConfig> scheduleConfig() {
        return this.scheduleConfig;
    }
    /**
     * Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
     * 
     */
    @Export(name="sourceUri", refs={String.class}, tree="[0]")
    private Output<String> sourceUri;

    /**
     * @return Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
     * 
     */
    public Output<String> sourceUri() {
        return this.sourceUri;
    }
    /**
     * Tags to apply to the Data Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Tags to apply to the Data Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataIntegration(String name) {
        this(name, DataIntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataIntegration(String name, DataIntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataIntegration(String name, DataIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appintegrations/dataIntegration:DataIntegration", name, args == null ? DataIntegrationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataIntegration(String name, Output<String> id, @Nullable DataIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appintegrations/dataIntegration:DataIntegration", name, state, makeResourceOptions(options, id));
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
    public static DataIntegration get(String name, Output<String> id, @Nullable DataIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataIntegration(name, id, state, options);
    }
}
