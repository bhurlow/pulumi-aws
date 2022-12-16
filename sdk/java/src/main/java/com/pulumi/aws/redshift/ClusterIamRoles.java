// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.redshift;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.redshift.ClusterIamRolesArgs;
import com.pulumi.aws.redshift.inputs.ClusterIamRolesState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Redshift Cluster IAM Roles resource.
 * 
 * &gt; **NOTE:** A Redshift cluster&#39;s default IAM role can be managed both by this resource&#39;s `default_iam_role_arn` argument and the `aws.redshift.Cluster` resource&#39;s `default_iam_role_arn` argument. Do not configure different values for both arguments. Doing so will cause a conflict of default IAM roles.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.redshift.ClusterIamRoles;
 * import com.pulumi.aws.redshift.ClusterIamRolesArgs;
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
 *         var example = new ClusterIamRoles(&#34;example&#34;, ClusterIamRolesArgs.builder()        
 *             .clusterIdentifier(aws_redshift_cluster.example().cluster_identifier())
 *             .iamRoleArns(aws_iam_role.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Redshift Cluster IAM Roless can be imported using the `cluster_identifier`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:redshift/clusterIamRoles:ClusterIamRoles examplegroup1 example
 * ```
 * 
 */
@ResourceType(type="aws:redshift/clusterIamRoles:ClusterIamRoles")
public class ClusterIamRoles extends com.pulumi.resources.CustomResource {
    /**
     * The name of the Redshift Cluster IAM Roles.
     * 
     */
    @Export(name="clusterIdentifier", refs={String.class}, tree="[0]")
    private Output<String> clusterIdentifier;

    /**
     * @return The name of the Redshift Cluster IAM Roles.
     * 
     */
    public Output<String> clusterIdentifier() {
        return this.clusterIdentifier;
    }
    /**
     * The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.
     * 
     */
    @Export(name="defaultIamRoleArn", refs={String.class}, tree="[0]")
    private Output<String> defaultIamRoleArn;

    /**
     * @return The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.
     * 
     */
    public Output<String> defaultIamRoleArn() {
        return this.defaultIamRoleArn;
    }
    /**
     * A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
     * 
     */
    @Export(name="iamRoleArns", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> iamRoleArns;

    /**
     * @return A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
     * 
     */
    public Output<List<String>> iamRoleArns() {
        return this.iamRoleArns;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterIamRoles(String name) {
        this(name, ClusterIamRolesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterIamRoles(String name, ClusterIamRolesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterIamRoles(String name, ClusterIamRolesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:redshift/clusterIamRoles:ClusterIamRoles", name, args == null ? ClusterIamRolesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterIamRoles(String name, Output<String> id, @Nullable ClusterIamRolesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:redshift/clusterIamRoles:ClusterIamRoles", name, state, makeResourceOptions(options, id));
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
    public static ClusterIamRoles get(String name, Output<String> id, @Nullable ClusterIamRolesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterIamRoles(name, id, state, options);
    }
}
