// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.sagemaker.ModelArgs;
import com.pulumi.aws.sagemaker.inputs.ModelState;
import com.pulumi.aws.sagemaker.outputs.ModelContainer;
import com.pulumi.aws.sagemaker.outputs.ModelInferenceExecutionConfig;
import com.pulumi.aws.sagemaker.outputs.ModelPrimaryContainer;
import com.pulumi.aws.sagemaker.outputs.ModelVpcConfig;
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
 * Provides a SageMaker model resource.
 * 
 * ## Example Usage
 * 
 * Basic usage:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.sagemaker.SagemakerFunctions;
 * import com.pulumi.aws.sagemaker.inputs.GetPrebuiltEcrImageArgs;
 * import com.pulumi.aws.sagemaker.Model;
 * import com.pulumi.aws.sagemaker.ModelArgs;
 * import com.pulumi.aws.sagemaker.inputs.ModelPrimaryContainerArgs;
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
 *         final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;sagemaker.amazonaws.com&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleRole = new Role(&#34;exampleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         final var test = SagemakerFunctions.getPrebuiltEcrImage(GetPrebuiltEcrImageArgs.builder()
 *             .repositoryName(&#34;kmeans&#34;)
 *             .build());
 * 
 *         var exampleModel = new Model(&#34;exampleModel&#34;, ModelArgs.builder()        
 *             .executionRoleArn(exampleRole.arn())
 *             .primaryContainer(ModelPrimaryContainerArgs.builder()
 *                 .image(test.applyValue(getPrebuiltEcrImageResult -&gt; getPrebuiltEcrImageResult.registryPath()))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Inference Execution Config
 * 
 * * `mode` - (Required) How containers in a multi-container are run. The following values are valid `Serial` and `Direct`.
 * 
 * ## Import
 * 
 * Using `pulumi import`, import models using the `name`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:sagemaker/model:Model test_model model-foo
 * ```
 * 
 */
@ResourceType(type="aws:sagemaker/model:Model")
public class Model extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this model.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) assigned by AWS to this model.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
     * 
     */
    @Export(name="containers", refs={List.class,ModelContainer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ModelContainer>> containers;

    /**
     * @return Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
     * 
     */
    public Output<Optional<List<ModelContainer>>> containers() {
        return Codegen.optional(this.containers);
    }
    /**
     * Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     * 
     */
    @Export(name="enableNetworkIsolation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableNetworkIsolation;

    /**
     * @return Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     * 
     */
    public Output<Optional<Boolean>> enableNetworkIsolation() {
        return Codegen.optional(this.enableNetworkIsolation);
    }
    /**
     * A role that SageMaker can assume to access model artifacts and docker images for deployment.
     * 
     */
    @Export(name="executionRoleArn", refs={String.class}, tree="[0]")
    private Output<String> executionRoleArn;

    /**
     * @return A role that SageMaker can assume to access model artifacts and docker images for deployment.
     * 
     */
    public Output<String> executionRoleArn() {
        return this.executionRoleArn;
    }
    /**
     * Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
     * 
     */
    @Export(name="inferenceExecutionConfig", refs={ModelInferenceExecutionConfig.class}, tree="[0]")
    private Output<ModelInferenceExecutionConfig> inferenceExecutionConfig;

    /**
     * @return Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
     * 
     */
    public Output<ModelInferenceExecutionConfig> inferenceExecutionConfig() {
        return this.inferenceExecutionConfig;
    }
    /**
     * The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     * 
     */
    @Export(name="primaryContainer", refs={ModelPrimaryContainer.class}, tree="[0]")
    private Output</* @Nullable */ ModelPrimaryContainer> primaryContainer;

    /**
     * @return The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     * 
     */
    public Output<Optional<ModelPrimaryContainer>> primaryContainer() {
        return Codegen.optional(this.primaryContainer);
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * The `primary_container` and `container` block both support:
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * The `primary_container` and `container` block both support:
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
     * Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     * 
     */
    @Export(name="vpcConfig", refs={ModelVpcConfig.class}, tree="[0]")
    private Output</* @Nullable */ ModelVpcConfig> vpcConfig;

    /**
     * @return Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     * 
     */
    public Output<Optional<ModelVpcConfig>> vpcConfig() {
        return Codegen.optional(this.vpcConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Model(String name) {
        this(name, ModelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Model(String name, ModelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Model(String name, ModelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/model:Model", name, args == null ? ModelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Model(String name, Output<String> id, @Nullable ModelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/model:Model", name, state, makeResourceOptions(options, id));
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
    public static Model get(String name, Output<String> id, @Nullable ModelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Model(name, id, state, options);
    }
}
