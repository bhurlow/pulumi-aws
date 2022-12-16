// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.route53.QueryLogArgs;
import com.pulumi.aws.route53.inputs.QueryLogState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Route53 query logging configuration resource.
 * 
 * &gt; **NOTE:** There are restrictions on the configuration of query logging. Notably,
 * the CloudWatch log group must be in the `us-east-1` region,
 * a permissive CloudWatch log resource policy must be in place, and
 * the Route53 hosted zone must be public.
 * See [Configuring Logging for DNS Queries](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html?console_help=true#query-logs-configuring) for additional details.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.Provider;
 * import com.pulumi.aws.ProviderArgs;
 * import com.pulumi.aws.cloudwatch.LogGroup;
 * import com.pulumi.aws.cloudwatch.LogGroupArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.cloudwatch.LogResourcePolicy;
 * import com.pulumi.aws.cloudwatch.LogResourcePolicyArgs;
 * import com.pulumi.aws.route53.Zone;
 * import com.pulumi.aws.route53.QueryLog;
 * import com.pulumi.aws.route53.QueryLogArgs;
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
 *         var us_east_1 = new Provider(&#34;us-east-1&#34;, ProviderArgs.builder()        
 *             .region(&#34;us-east-1&#34;)
 *             .build());
 * 
 *         var awsRoute53ExampleCom = new LogGroup(&#34;awsRoute53ExampleCom&#34;, LogGroupArgs.builder()        
 *             .retentionInDays(30)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.us-east-1())
 *                 .build());
 * 
 *         final var route53-query-logging-policyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .actions(                
 *                     &#34;logs:CreateLogStream&#34;,
 *                     &#34;logs:PutLogEvents&#34;)
 *                 .resources(&#34;arn:aws:logs:*:*:log-group:/aws/route53/*&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .identifiers(&#34;route53.amazonaws.com&#34;)
 *                     .type(&#34;Service&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var route53_query_logging_policyLogResourcePolicy = new LogResourcePolicy(&#34;route53-query-logging-policyLogResourcePolicy&#34;, LogResourcePolicyArgs.builder()        
 *             .policyDocument(route53_query_logging_policyPolicyDocument.json())
 *             .policyName(&#34;route53-query-logging-policy&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.us-east-1())
 *                 .build());
 * 
 *         var exampleComZone = new Zone(&#34;exampleComZone&#34;);
 * 
 *         var exampleComQueryLog = new QueryLog(&#34;exampleComQueryLog&#34;, QueryLogArgs.builder()        
 *             .cloudwatchLogGroupArn(awsRoute53ExampleCom.arn())
 *             .zoneId(exampleComZone.zoneId())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(route53_query_logging_policyLogResourcePolicy)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Route53 query logging configurations can be imported using their ID, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:route53/queryLog:QueryLog example_com xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
 * ```
 * 
 */
@ResourceType(type="aws:route53/queryLog:QueryLog")
public class QueryLog extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the Query Logging Config.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the Query Logging Config.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * CloudWatch log group ARN to send query logs.
     * 
     */
    @Export(name="cloudwatchLogGroupArn", refs={String.class}, tree="[0]")
    private Output<String> cloudwatchLogGroupArn;

    /**
     * @return CloudWatch log group ARN to send query logs.
     * 
     */
    public Output<String> cloudwatchLogGroupArn() {
        return this.cloudwatchLogGroupArn;
    }
    /**
     * Route53 hosted zone ID to enable query logs.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return Route53 hosted zone ID to enable query logs.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QueryLog(String name) {
        this(name, QueryLogArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QueryLog(String name, QueryLogArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QueryLog(String name, QueryLogArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/queryLog:QueryLog", name, args == null ? QueryLogArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QueryLog(String name, Output<String> id, @Nullable QueryLogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/queryLog:QueryLog", name, state, makeResourceOptions(options, id));
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
    public static QueryLog get(String name, Output<String> id, @Nullable QueryLogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QueryLog(name, id, state, options);
    }
}
