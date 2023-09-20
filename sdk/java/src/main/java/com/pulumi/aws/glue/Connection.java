// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.glue.ConnectionArgs;
import com.pulumi.aws.glue.inputs.ConnectionState;
import com.pulumi.aws.glue.outputs.ConnectionPhysicalConnectionRequirements;
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
 * Provides a Glue Connection resource.
 * 
 * ## Example Usage
 * ### Non-VPC Connection
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.glue.Connection;
 * import com.pulumi.aws.glue.ConnectionArgs;
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
 *         var example = new Connection(&#34;example&#34;, ConnectionArgs.builder()        
 *             .connectionProperties(Map.ofEntries(
 *                 Map.entry(&#34;JDBC_CONNECTION_URL&#34;, &#34;jdbc:mysql://example.com/exampledatabase&#34;),
 *                 Map.entry(&#34;PASSWORD&#34;, &#34;examplepassword&#34;),
 *                 Map.entry(&#34;USERNAME&#34;, &#34;exampleusername&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### VPC Connection
 * 
 * For more information, see the [AWS Documentation](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html#connection-JDBC-VPC).
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.glue.Connection;
 * import com.pulumi.aws.glue.ConnectionArgs;
 * import com.pulumi.aws.glue.inputs.ConnectionPhysicalConnectionRequirementsArgs;
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
 *         var example = new Connection(&#34;example&#34;, ConnectionArgs.builder()        
 *             .connectionProperties(Map.ofEntries(
 *                 Map.entry(&#34;JDBC_CONNECTION_URL&#34;, String.format(&#34;jdbc:mysql://%s/exampledatabase&#34;, aws_rds_cluster.example().endpoint())),
 *                 Map.entry(&#34;PASSWORD&#34;, &#34;examplepassword&#34;),
 *                 Map.entry(&#34;USERNAME&#34;, &#34;exampleusername&#34;)
 *             ))
 *             .physicalConnectionRequirements(ConnectionPhysicalConnectionRequirementsArgs.builder()
 *                 .availabilityZone(aws_subnet.example().availability_zone())
 *                 .securityGroupIdLists(aws_security_group.example().id())
 *                 .subnetId(aws_subnet.example().id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Glue Connections using the `CATALOG-ID` (AWS account ID if not custom) and `NAME`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:glue/connection:Connection MyConnection 123456789012:MyConnection
 * ```
 * 
 */
@ResourceType(type="aws:glue/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the Glue Connection.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the Glue Connection.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
     * 
     */
    @Export(name="catalogId", refs={String.class}, tree="[0]")
    private Output<String> catalogId;

    /**
     * @return The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
     * 
     */
    public Output<String> catalogId() {
        return this.catalogId;
    }
    /**
     * A map of key-value pairs used as parameters for this connection.
     * 
     */
    @Export(name="connectionProperties", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> connectionProperties;

    /**
     * @return A map of key-value pairs used as parameters for this connection.
     * 
     */
    public Output<Optional<Map<String,String>>> connectionProperties() {
        return Codegen.optional(this.connectionProperties);
    }
    /**
     * The type of the connection. Supported are: `CUSTOM`, `JDBC`, `KAFKA`, `MARKETPLACE`, `MONGODB`, and `NETWORK`. Defaults to `JBDC`.
     * 
     */
    @Export(name="connectionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> connectionType;

    /**
     * @return The type of the connection. Supported are: `CUSTOM`, `JDBC`, `KAFKA`, `MARKETPLACE`, `MONGODB`, and `NETWORK`. Defaults to `JBDC`.
     * 
     */
    public Output<Optional<String>> connectionType() {
        return Codegen.optional(this.connectionType);
    }
    /**
     * Description of the connection.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the connection.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A list of criteria that can be used in selecting this connection.
     * 
     */
    @Export(name="matchCriterias", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> matchCriterias;

    /**
     * @return A list of criteria that can be used in selecting this connection.
     * 
     */
    public Output<Optional<List<String>>> matchCriterias() {
        return Codegen.optional(this.matchCriterias);
    }
    /**
     * The name of the connection.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the connection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
     * 
     */
    @Export(name="physicalConnectionRequirements", refs={ConnectionPhysicalConnectionRequirements.class}, tree="[0]")
    private Output</* @Nullable */ ConnectionPhysicalConnectionRequirements> physicalConnectionRequirements;

    /**
     * @return A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
     * 
     */
    public Output<Optional<ConnectionPhysicalConnectionRequirements>> physicalConnectionRequirements() {
        return Codegen.optional(this.physicalConnectionRequirements);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, @Nullable ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, @Nullable ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:glue/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:glue/connection:Connection", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "connectionProperties",
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
    public static Connection get(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connection(name, id, state, options);
    }
}
