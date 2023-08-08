// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3control;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3control.ObjectLambdaAccessPointPolicyArgs;
import com.pulumi.aws.s3control.inputs.ObjectLambdaAccessPointPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage an S3 Object Lambda Access Point resource policy.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.AccessPoint;
 * import com.pulumi.aws.s3.AccessPointArgs;
 * import com.pulumi.aws.s3control.ObjectLambdaAccessPoint;
 * import com.pulumi.aws.s3control.ObjectLambdaAccessPointArgs;
 * import com.pulumi.aws.s3control.inputs.ObjectLambdaAccessPointConfigurationArgs;
 * import com.pulumi.aws.s3control.ObjectLambdaAccessPointPolicy;
 * import com.pulumi.aws.s3control.ObjectLambdaAccessPointPolicyArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var exampleBucketV2 = new BucketV2(&#34;exampleBucketV2&#34;);
 * 
 *         var exampleAccessPoint = new AccessPoint(&#34;exampleAccessPoint&#34;, AccessPointArgs.builder()        
 *             .bucket(exampleBucketV2.id())
 *             .build());
 * 
 *         var exampleObjectLambdaAccessPoint = new ObjectLambdaAccessPoint(&#34;exampleObjectLambdaAccessPoint&#34;, ObjectLambdaAccessPointArgs.builder()        
 *             .configuration(ObjectLambdaAccessPointConfigurationArgs.builder()
 *                 .supportingAccessPoint(exampleAccessPoint.arn())
 *                 .transformationConfigurations(ObjectLambdaAccessPointConfigurationTransformationConfigurationArgs.builder()
 *                     .actions(&#34;GetObject&#34;)
 *                     .contentTransformation(ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationArgs.builder()
 *                         .awsLambda(ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambdaArgs.builder()
 *                             .functionArn(aws_lambda_function.example().arn())
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleObjectLambdaAccessPointPolicy = new ObjectLambdaAccessPointPolicy(&#34;exampleObjectLambdaAccessPointPolicy&#34;, ObjectLambdaAccessPointPolicyArgs.builder()        
 *             .policy(exampleObjectLambdaAccessPoint.arn().applyValue(arn -&gt; serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Version&#34;, &#34;2008-10-17&#34;),
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                         jsonProperty(&#34;Action&#34;, &#34;s3-object-lambda:GetObject&#34;),
 *                         jsonProperty(&#34;Principal&#34;, jsonObject(
 *                             jsonProperty(&#34;AWS&#34;, data.aws_caller_identity().current().account_id())
 *                         )),
 *                         jsonProperty(&#34;Resource&#34;, arn)
 *                     )))
 *                 ))))
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
 *  to = aws_s3control_object_lambda_access_point_policy.example
 * 
 *  id = &#34;123456789012:example&#34; } Using `pulumi import`, import Object Lambda Access Point policies using the `account_id` and `name`, separated by a colon (`:`). For exampleconsole % pulumi import aws_s3control_object_lambda_access_point_policy.example 123456789012:example
 * 
 */
@ResourceType(type="aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy")
public class ObjectLambdaAccessPointPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The AWS account ID for the account that owns the Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return The AWS account ID for the account that owns the Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * Indicates whether this access point currently has a policy that allows public access.
     * 
     */
    @Export(name="hasPublicAccessPolicy", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hasPublicAccessPolicy;

    /**
     * @return Indicates whether this access point currently has a policy that allows public access.
     * 
     */
    public Output<Boolean> hasPublicAccessPolicy() {
        return this.hasPublicAccessPolicy;
    }
    /**
     * The name of the Object Lambda Access Point.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Object Lambda Access Point.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Object Lambda Access Point resource policy document.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return The Object Lambda Access Point resource policy document.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ObjectLambdaAccessPointPolicy(String name) {
        this(name, ObjectLambdaAccessPointPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ObjectLambdaAccessPointPolicy(String name, ObjectLambdaAccessPointPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ObjectLambdaAccessPointPolicy(String name, ObjectLambdaAccessPointPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy", name, args == null ? ObjectLambdaAccessPointPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ObjectLambdaAccessPointPolicy(String name, Output<String> id, @Nullable ObjectLambdaAccessPointPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy", name, state, makeResourceOptions(options, id));
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
    public static ObjectLambdaAccessPointPolicy get(String name, Output<String> id, @Nullable ObjectLambdaAccessPointPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ObjectLambdaAccessPointPolicy(name, id, state, options);
    }
}
