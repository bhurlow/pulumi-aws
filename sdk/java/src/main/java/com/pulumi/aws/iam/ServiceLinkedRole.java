// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.iam;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.iam.ServiceLinkedRoleArgs;
import com.pulumi.aws.iam.inputs.ServiceLinkedRoleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an [IAM service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.ServiceLinkedRole;
 * import com.pulumi.aws.iam.ServiceLinkedRoleArgs;
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
 *         var elasticbeanstalk = new ServiceLinkedRole(&#34;elasticbeanstalk&#34;, ServiceLinkedRoleArgs.builder()        
 *             .awsServiceName(&#34;elasticbeanstalk.amazonaws.com&#34;)
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
 *  to = aws_iam_service_linked_role.elasticbeanstalk
 * 
 *  id = &#34;arn:aws:iam::123456789012:role/aws-service-role/elasticbeanstalk.amazonaws.com/AWSServiceRoleForElasticBeanstalk&#34; } Using `pulumi import`, import IAM service-linked roles using role ARN. For exampleconsole % pulumi import aws_iam_service_linked_role.elasticbeanstalk arn:aws:iam::123456789012:role/aws-service-role/elasticbeanstalk.amazonaws.com/AWSServiceRoleForElasticBeanstalk
 * 
 */
@ResourceType(type="aws:iam/serviceLinkedRole:ServiceLinkedRole")
public class ServiceLinkedRole extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) specifying the role.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) specifying the role.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
     * 
     */
    @Export(name="awsServiceName", refs={String.class}, tree="[0]")
    private Output<String> awsServiceName;

    /**
     * @return The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
     * 
     */
    public Output<String> awsServiceName() {
        return this.awsServiceName;
    }
    /**
     * The creation date of the IAM role.
     * 
     */
    @Export(name="createDate", refs={String.class}, tree="[0]")
    private Output<String> createDate;

    /**
     * @return The creation date of the IAM role.
     * 
     */
    public Output<String> createDate() {
        return this.createDate;
    }
    /**
     * Additional string appended to the role name. Not all AWS services support custom suffixes.
     * 
     */
    @Export(name="customSuffix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customSuffix;

    /**
     * @return Additional string appended to the role name. Not all AWS services support custom suffixes.
     * 
     */
    public Output<Optional<String>> customSuffix() {
        return Codegen.optional(this.customSuffix);
    }
    /**
     * The description of the role.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the role.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the role.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the role.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The path of the role.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The path of the role.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Key-value mapping of tags for the IAM role. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of tags for the IAM role. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The stable and unique string identifying the role.
     * 
     */
    @Export(name="uniqueId", refs={String.class}, tree="[0]")
    private Output<String> uniqueId;

    /**
     * @return The stable and unique string identifying the role.
     * 
     */
    public Output<String> uniqueId() {
        return this.uniqueId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceLinkedRole(String name) {
        this(name, ServiceLinkedRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceLinkedRole(String name, ServiceLinkedRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceLinkedRole(String name, ServiceLinkedRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iam/serviceLinkedRole:ServiceLinkedRole", name, args == null ? ServiceLinkedRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceLinkedRole(String name, Output<String> id, @Nullable ServiceLinkedRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iam/serviceLinkedRole:ServiceLinkedRole", name, state, makeResourceOptions(options, id));
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
    public static ServiceLinkedRole get(String name, Output<String> id, @Nullable ServiceLinkedRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceLinkedRole(name, id, state, options);
    }
}
