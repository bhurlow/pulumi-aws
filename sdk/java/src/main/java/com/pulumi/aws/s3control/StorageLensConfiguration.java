// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3control;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3control.StorageLensConfigurationArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationState;
import com.pulumi.aws.s3control.outputs.StorageLensConfigurationStorageLensConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage an S3 Storage Lens configuration.
 * 
 * ## Import
 * 
 * S3 Storage Lens configurations can be imported using the `account_id` and `config_id`, separated by a colon (`:`), e.g.
 * 
 * ```sh
 *  $ pulumi import aws:s3control/storageLensConfiguration:StorageLensConfiguration example 123456789012:example-1
 * ```
 * 
 */
@ResourceType(type="aws:s3control/storageLensConfiguration:StorageLensConfiguration")
public class StorageLensConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * The account ID of the owner of the S3 Storage Lens metrics export bucket.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return The account ID of the owner of the S3 Storage Lens metrics export bucket.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * The Amazon Resource Name (ARN) of the bucket.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the bucket.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ID of the S3 Storage Lens configuration.
     * 
     */
    @Export(name="configId", refs={String.class}, tree="[0]")
    private Output<String> configId;

    /**
     * @return The ID of the S3 Storage Lens configuration.
     * 
     */
    public Output<String> configId() {
        return this.configId;
    }
    /**
     * The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
     * 
     */
    @Export(name="storageLensConfiguration", refs={StorageLensConfigurationStorageLensConfiguration.class}, tree="[0]")
    private Output<StorageLensConfigurationStorageLensConfiguration> storageLensConfiguration;

    /**
     * @return The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
     * 
     */
    public Output<StorageLensConfigurationStorageLensConfiguration> storageLensConfiguration() {
        return this.storageLensConfiguration;
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
     */
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
    public StorageLensConfiguration(String name) {
        this(name, StorageLensConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StorageLensConfiguration(String name, StorageLensConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StorageLensConfiguration(String name, StorageLensConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/storageLensConfiguration:StorageLensConfiguration", name, args == null ? StorageLensConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private StorageLensConfiguration(String name, Output<String> id, @Nullable StorageLensConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/storageLensConfiguration:StorageLensConfiguration", name, state, makeResourceOptions(options, id));
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
    public static StorageLensConfiguration get(String name, Output<String> id, @Nullable StorageLensConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new StorageLensConfiguration(name, id, state, options);
    }
}
