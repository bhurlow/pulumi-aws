// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cloudwatch.QueryDefinitionArgs;
import com.pulumi.aws.cloudwatch.inputs.QueryDefinitionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CloudWatch Logs query definition resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudwatch.QueryDefinition;
 * import com.pulumi.aws.cloudwatch.QueryDefinitionArgs;
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
 *         var example = new QueryDefinition(&#34;example&#34;, QueryDefinitionArgs.builder()        
 *             .logGroupNames(            
 *                 &#34;/aws/logGroup1&#34;,
 *                 &#34;/aws/logGroup2&#34;)
 *             .queryString(&#34;&#34;&#34;
 * fields @timestamp, @message
 * | sort @timestamp desc
 * | limit 25
 * 
 *             &#34;&#34;&#34;)
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
 *  to = aws_cloudwatch_query_definition.example
 * 
 *  id = &#34;arn:aws:logs:us-west-2:123456789012:query-definition:269951d7-6f75-496d-9d7b-6b7a5486bdbd&#34; } Using `pulumi import`, import CloudWatch query definitions using the query definition ARN. The ARN can be found on the &#34;Edit Query&#34; page for the query in the AWS Console. For exampleconsole % pulumi import aws_cloudwatch_query_definition.example arn:aws:logs:us-west-2:123456789012:query-definition:269951d7-6f75-496d-9d7b-6b7a5486bdbd
 * 
 */
@ResourceType(type="aws:cloudwatch/queryDefinition:QueryDefinition")
public class QueryDefinition extends com.pulumi.resources.CustomResource {
    /**
     * Specific log groups to use with the query.
     * 
     */
    @Export(name="logGroupNames", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> logGroupNames;

    /**
     * @return Specific log groups to use with the query.
     * 
     */
    public Output<Optional<List<String>>> logGroupNames() {
        return Codegen.optional(this.logGroupNames);
    }
    /**
     * The name of the query.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the query.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The query definition ID.
     * 
     */
    @Export(name="queryDefinitionId", refs={String.class}, tree="[0]")
    private Output<String> queryDefinitionId;

    /**
     * @return The query definition ID.
     * 
     */
    public Output<String> queryDefinitionId() {
        return this.queryDefinitionId;
    }
    /**
     * The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
     * 
     */
    @Export(name="queryString", refs={String.class}, tree="[0]")
    private Output<String> queryString;

    /**
     * @return The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
     * 
     */
    public Output<String> queryString() {
        return this.queryString;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QueryDefinition(String name) {
        this(name, QueryDefinitionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QueryDefinition(String name, QueryDefinitionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QueryDefinition(String name, QueryDefinitionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudwatch/queryDefinition:QueryDefinition", name, args == null ? QueryDefinitionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QueryDefinition(String name, Output<String> id, @Nullable QueryDefinitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudwatch/queryDefinition:QueryDefinition", name, state, makeResourceOptions(options, id));
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
    public static QueryDefinition get(String name, Output<String> id, @Nullable QueryDefinitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QueryDefinition(name, id, state, options);
    }
}
