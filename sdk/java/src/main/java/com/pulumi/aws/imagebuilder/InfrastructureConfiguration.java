// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.imagebuilder;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.imagebuilder.InfrastructureConfigurationArgs;
import com.pulumi.aws.imagebuilder.inputs.InfrastructureConfigurationState;
import com.pulumi.aws.imagebuilder.outputs.InfrastructureConfigurationInstanceMetadataOptions;
import com.pulumi.aws.imagebuilder.outputs.InfrastructureConfigurationLogging;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Image Builder Infrastructure Configuration.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.imagebuilder.InfrastructureConfiguration;
 * import com.pulumi.aws.imagebuilder.InfrastructureConfigurationArgs;
 * import com.pulumi.aws.imagebuilder.inputs.InfrastructureConfigurationLoggingArgs;
 * import com.pulumi.aws.imagebuilder.inputs.InfrastructureConfigurationLoggingS3LogsArgs;
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
 *         var example = new InfrastructureConfiguration(&#34;example&#34;, InfrastructureConfigurationArgs.builder()        
 *             .description(&#34;example description&#34;)
 *             .instanceProfileName(aws_iam_instance_profile.example().name())
 *             .instanceTypes(            
 *                 &#34;t2.nano&#34;,
 *                 &#34;t3.micro&#34;)
 *             .keyPair(aws_key_pair.example().key_name())
 *             .securityGroupIds(aws_security_group.example().id())
 *             .snsTopicArn(aws_sns_topic.example().arn())
 *             .subnetId(aws_subnet.main().id())
 *             .terminateInstanceOnFailure(true)
 *             .logging(InfrastructureConfigurationLoggingArgs.builder()
 *                 .s3Logs(InfrastructureConfigurationLoggingS3LogsArgs.builder()
 *                     .s3BucketName(aws_s3_bucket.example().bucket())
 *                     .s3KeyPrefix(&#34;logs&#34;)
 *                     .build())
 *                 .build())
 *             .tags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_imagebuilder_infrastructure_configuration` can be imported using the Amazon Resource Name (ARN), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-configuration/example
 * ```
 * 
 */
@ResourceType(type="aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration")
public class InfrastructureConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the configuration.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the configuration.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Date when the configuration was created.
     * 
     */
    @Export(name="dateCreated", refs={String.class}, tree="[0]")
    private Output<String> dateCreated;

    /**
     * @return Date when the configuration was created.
     * 
     */
    public Output<String> dateCreated() {
        return this.dateCreated;
    }
    /**
     * Date when the configuration was updated.
     * 
     */
    @Export(name="dateUpdated", refs={String.class}, tree="[0]")
    private Output<String> dateUpdated;

    /**
     * @return Date when the configuration was updated.
     * 
     */
    public Output<String> dateUpdated() {
        return this.dateUpdated;
    }
    /**
     * Description for the configuration.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description for the configuration.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
     * 
     */
    @Export(name="instanceMetadataOptions", refs={InfrastructureConfigurationInstanceMetadataOptions.class}, tree="[0]")
    private Output</* @Nullable */ InfrastructureConfigurationInstanceMetadataOptions> instanceMetadataOptions;

    /**
     * @return Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
     * 
     */
    public Output<Optional<InfrastructureConfigurationInstanceMetadataOptions>> instanceMetadataOptions() {
        return Codegen.optional(this.instanceMetadataOptions);
    }
    /**
     * Name of IAM Instance Profile.
     * 
     */
    @Export(name="instanceProfileName", refs={String.class}, tree="[0]")
    private Output<String> instanceProfileName;

    /**
     * @return Name of IAM Instance Profile.
     * 
     */
    public Output<String> instanceProfileName() {
        return this.instanceProfileName;
    }
    /**
     * Set of EC2 Instance Types.
     * 
     */
    @Export(name="instanceTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> instanceTypes;

    /**
     * @return Set of EC2 Instance Types.
     * 
     */
    public Output<Optional<List<String>>> instanceTypes() {
        return Codegen.optional(this.instanceTypes);
    }
    /**
     * Name of EC2 Key Pair.
     * 
     */
    @Export(name="keyPair", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyPair;

    /**
     * @return Name of EC2 Key Pair.
     * 
     */
    public Output<Optional<String>> keyPair() {
        return Codegen.optional(this.keyPair);
    }
    /**
     * Configuration block with logging settings. Detailed below.
     * 
     */
    @Export(name="logging", refs={InfrastructureConfigurationLogging.class}, tree="[0]")
    private Output</* @Nullable */ InfrastructureConfigurationLogging> logging;

    /**
     * @return Configuration block with logging settings. Detailed below.
     * 
     */
    public Output<Optional<InfrastructureConfigurationLogging>> logging() {
        return Codegen.optional(this.logging);
    }
    /**
     * Name for the configuration.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name for the configuration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Key-value map of resource tags to assign to infrastructure created by the configuration.
     * 
     */
    @Export(name="resourceTags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> resourceTags;

    /**
     * @return Key-value map of resource tags to assign to infrastructure created by the configuration.
     * 
     */
    public Output<Optional<Map<String,String>>> resourceTags() {
        return Codegen.optional(this.resourceTags);
    }
    /**
     * Set of EC2 Security Group identifiers.
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> securityGroupIds;

    /**
     * @return Set of EC2 Security Group identifiers.
     * 
     */
    public Output<Optional<List<String>>> securityGroupIds() {
        return Codegen.optional(this.securityGroupIds);
    }
    /**
     * Amazon Resource Name (ARN) of SNS Topic.
     * 
     */
    @Export(name="snsTopicArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snsTopicArn;

    /**
     * @return Amazon Resource Name (ARN) of SNS Topic.
     * 
     */
    public Output<Optional<String>> snsTopicArn() {
        return Codegen.optional(this.snsTopicArn);
    }
    /**
     * EC2 Subnet identifier. Also requires `security_group_ids` argument.
     * 
     */
    @Export(name="subnetId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subnetId;

    /**
     * @return EC2 Subnet identifier. Also requires `security_group_ids` argument.
     * 
     */
    public Output<Optional<String>> subnetId() {
        return Codegen.optional(this.subnetId);
    }
    /**
     * Key-value map of resource tags to assign to the configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags to assign to the configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
     * 
     */
    @Export(name="terminateInstanceOnFailure", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> terminateInstanceOnFailure;

    /**
     * @return Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> terminateInstanceOnFailure() {
        return Codegen.optional(this.terminateInstanceOnFailure);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InfrastructureConfiguration(String name) {
        this(name, InfrastructureConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InfrastructureConfiguration(String name, InfrastructureConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InfrastructureConfiguration(String name, InfrastructureConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration", name, args == null ? InfrastructureConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InfrastructureConfiguration(String name, Output<String> id, @Nullable InfrastructureConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration", name, state, makeResourceOptions(options, id));
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
    public static InfrastructureConfiguration get(String name, Output<String> id, @Nullable InfrastructureConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InfrastructureConfiguration(name, id, state, options);
    }
}
