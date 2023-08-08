// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.comprehend;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.comprehend.DocumentClassifierArgs;
import com.pulumi.aws.comprehend.inputs.DocumentClassifierState;
import com.pulumi.aws.comprehend.outputs.DocumentClassifierInputDataConfig;
import com.pulumi.aws.comprehend.outputs.DocumentClassifierOutputDataConfig;
import com.pulumi.aws.comprehend.outputs.DocumentClassifierVpcConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS Comprehend Document Classifier.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketObjectv2;
 * import com.pulumi.aws.comprehend.DocumentClassifier;
 * import com.pulumi.aws.comprehend.DocumentClassifierArgs;
 * import com.pulumi.aws.comprehend.inputs.DocumentClassifierInputDataConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var documents = new BucketObjectv2(&#34;documents&#34;);
 * 
 *         var example = new DocumentClassifier(&#34;example&#34;, DocumentClassifierArgs.builder()        
 *             .dataAccessRoleArn(aws_iam_role.example().arn())
 *             .languageCode(&#34;en&#34;)
 *             .inputDataConfig(DocumentClassifierInputDataConfigArgs.builder()
 *                 .s3Uri(documents.id().applyValue(id -&gt; String.format(&#34;s3://%s/%s&#34;, aws_s3_bucket.test().bucket(),id)))
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(aws_iam_role_policy.example())
 *                 .build());
 * 
 *         var entities = new BucketObjectv2(&#34;entities&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_comprehend_document_classifier.example
 * 
 *  id = &#34;arn:aws:comprehend:us-west-2:123456789012:document_classifier/example&#34; } Using `pulumi import`, import Comprehend Document Classifier using the ARN. For exampleconsole % pulumi import aws_comprehend_document_classifier.example arn:aws:comprehend:us-west-2:123456789012:document_classifier/example
 * 
 */
@ResourceType(type="aws:comprehend/documentClassifier:DocumentClassifier")
public class DocumentClassifier extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the Document Classifier version.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the Document Classifier version.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ARN for an IAM Role which allows Comprehend to read the training and testing data.
     * 
     */
    @Export(name="dataAccessRoleArn", refs={String.class}, tree="[0]")
    private Output<String> dataAccessRoleArn;

    /**
     * @return The ARN for an IAM Role which allows Comprehend to read the training and testing data.
     * 
     */
    public Output<String> dataAccessRoleArn() {
        return this.dataAccessRoleArn;
    }
    /**
     * Configuration for the training and testing data.
     * See the `input_data_config` Configuration Block section below.
     * 
     */
    @Export(name="inputDataConfig", refs={DocumentClassifierInputDataConfig.class}, tree="[0]")
    private Output<DocumentClassifierInputDataConfig> inputDataConfig;

    /**
     * @return Configuration for the training and testing data.
     * See the `input_data_config` Configuration Block section below.
     * 
     */
    public Output<DocumentClassifierInputDataConfig> inputDataConfig() {
        return this.inputDataConfig;
    }
    /**
     * Two-letter language code for the language.
     * One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
     * 
     */
    @Export(name="languageCode", refs={String.class}, tree="[0]")
    private Output<String> languageCode;

    /**
     * @return Two-letter language code for the language.
     * One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
     * 
     */
    public Output<String> languageCode() {
        return this.languageCode;
    }
    /**
     * The document classification mode.
     * One of `MULTI_CLASS` or `MULTI_LABEL`.
     * `MULTI_CLASS` is also known as &#34;Single Label&#34; in the AWS Console.
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mode;

    /**
     * @return The document classification mode.
     * One of `MULTI_CLASS` or `MULTI_LABEL`.
     * `MULTI_CLASS` is also known as &#34;Single Label&#34; in the AWS Console.
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
    }
    /**
     * KMS Key used to encrypt trained Document Classifiers.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    @Export(name="modelKmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> modelKmsKeyId;

    /**
     * @return KMS Key used to encrypt trained Document Classifiers.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    public Output<Optional<String>> modelKmsKeyId() {
        return Codegen.optional(this.modelKmsKeyId);
    }
    /**
     * Name for the Document Classifier.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name for the Document Classifier.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Configuration for the output results of training.
     * See the `output_data_config` Configuration Block section below.
     * 
     */
    @Export(name="outputDataConfig", refs={DocumentClassifierOutputDataConfig.class}, tree="[0]")
    private Output<DocumentClassifierOutputDataConfig> outputDataConfig;

    /**
     * @return Configuration for the output results of training.
     * See the `output_data_config` Configuration Block section below.
     * 
     */
    public Output<DocumentClassifierOutputDataConfig> outputDataConfig() {
        return this.outputDataConfig;
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Name for the version of the Document Classifier.
     * Each version must have a unique name within the Document Classifier.
     * If omitted, the provider will assign a random, unique version name.
     * If explicitly set to `&#34;&#34;`, no version name will be set.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name_prefix`.
     * 
     */
    @Export(name="versionName", refs={String.class}, tree="[0]")
    private Output<String> versionName;

    /**
     * @return Name for the version of the Document Classifier.
     * Each version must have a unique name within the Document Classifier.
     * If omitted, the provider will assign a random, unique version name.
     * If explicitly set to `&#34;&#34;`, no version name will be set.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name_prefix`.
     * 
     */
    public Output<String> versionName() {
        return this.versionName;
    }
    /**
     * Creates a unique version name beginning with the specified prefix.
     * Has a maximum length of 37 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name`.
     * 
     */
    @Export(name="versionNamePrefix", refs={String.class}, tree="[0]")
    private Output<String> versionNamePrefix;

    /**
     * @return Creates a unique version name beginning with the specified prefix.
     * Has a maximum length of 37 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name`.
     * 
     */
    public Output<String> versionNamePrefix() {
        return this.versionNamePrefix;
    }
    /**
     * KMS Key used to encrypt storage volumes during job processing.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    @Export(name="volumeKmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> volumeKmsKeyId;

    /**
     * @return KMS Key used to encrypt storage volumes during job processing.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    public Output<Optional<String>> volumeKmsKeyId() {
        return Codegen.optional(this.volumeKmsKeyId);
    }
    /**
     * Configuration parameters for VPC to contain Document Classifier resources.
     * See the `vpc_config` Configuration Block section below.
     * 
     */
    @Export(name="vpcConfig", refs={DocumentClassifierVpcConfig.class}, tree="[0]")
    private Output</* @Nullable */ DocumentClassifierVpcConfig> vpcConfig;

    /**
     * @return Configuration parameters for VPC to contain Document Classifier resources.
     * See the `vpc_config` Configuration Block section below.
     * 
     */
    public Output<Optional<DocumentClassifierVpcConfig>> vpcConfig() {
        return Codegen.optional(this.vpcConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DocumentClassifier(String name) {
        this(name, DocumentClassifierArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DocumentClassifier(String name, DocumentClassifierArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DocumentClassifier(String name, DocumentClassifierArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:comprehend/documentClassifier:DocumentClassifier", name, args == null ? DocumentClassifierArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DocumentClassifier(String name, Output<String> id, @Nullable DocumentClassifierState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:comprehend/documentClassifier:DocumentClassifier", name, state, makeResourceOptions(options, id));
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
    public static DocumentClassifier get(String name, Output<String> id, @Nullable DocumentClassifierState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DocumentClassifier(name, id, state, options);
    }
}
