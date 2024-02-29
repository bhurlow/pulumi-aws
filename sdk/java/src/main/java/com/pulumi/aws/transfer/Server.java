// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transfer;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.transfer.ServerArgs;
import com.pulumi.aws.transfer.inputs.ServerState;
import com.pulumi.aws.transfer.outputs.ServerEndpointDetails;
import com.pulumi.aws.transfer.outputs.ServerProtocolDetails;
import com.pulumi.aws.transfer.outputs.ServerWorkflowDetails;
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
 * Provides a AWS Transfer Server resource.
 * 
 * &gt; **NOTE on AWS IAM permissions:** If the `endpoint_type` is set to `VPC`, the `ec2:DescribeVpcEndpoints` and `ec2:ModifyVpcEndpoint` [actions](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonec2.html#amazonec2-actions-as-permissions) are used.
 * 
 * &gt; **NOTE:** Use the `aws.transfer.Tag` resource to manage the system tags used for [custom hostnames](https://docs.aws.amazon.com/transfer/latest/userguide/requirements-dns.html#tag-custom-hostname-cdk).
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.transfer.Server;
 * import com.pulumi.aws.transfer.ServerArgs;
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
 *         var example = new Server(&#34;example&#34;, ServerArgs.builder()        
 *             .tags(Map.of(&#34;Name&#34;, &#34;Example&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Security Policy Name
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.transfer.Server;
 * import com.pulumi.aws.transfer.ServerArgs;
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
 *         var example = new Server(&#34;example&#34;, ServerArgs.builder()        
 *             .securityPolicyName(&#34;TransferSecurityPolicy-2020-06&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### VPC Endpoint
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.transfer.Server;
 * import com.pulumi.aws.transfer.ServerArgs;
 * import com.pulumi.aws.transfer.inputs.ServerEndpointDetailsArgs;
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
 *         var example = new Server(&#34;example&#34;, ServerArgs.builder()        
 *             .endpointType(&#34;VPC&#34;)
 *             .endpointDetails(ServerEndpointDetailsArgs.builder()
 *                 .addressAllocationIds(aws_eip.example().id())
 *                 .subnetIds(aws_subnet.example().id())
 *                 .vpcId(aws_vpc.example().id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### AWS Directory authentication
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.transfer.Server;
 * import com.pulumi.aws.transfer.ServerArgs;
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
 *         var example = new Server(&#34;example&#34;, ServerArgs.builder()        
 *             .identityProviderType(&#34;AWS_DIRECTORY_SERVICE&#34;)
 *             .directoryId(aws_directory_service_directory.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### AWS Lambda authentication
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.transfer.Server;
 * import com.pulumi.aws.transfer.ServerArgs;
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
 *         var example = new Server(&#34;example&#34;, ServerArgs.builder()        
 *             .identityProviderType(&#34;AWS_LAMBDA&#34;)
 *             .function(aws_lambda_identity_provider.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Protocols
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.transfer.Server;
 * import com.pulumi.aws.transfer.ServerArgs;
 * import com.pulumi.aws.transfer.inputs.ServerEndpointDetailsArgs;
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
 *         var example = new Server(&#34;example&#34;, ServerArgs.builder()        
 *             .endpointType(&#34;VPC&#34;)
 *             .endpointDetails(ServerEndpointDetailsArgs.builder()
 *                 .subnetIds(aws_subnet.example().id())
 *                 .vpcId(aws_vpc.example().id())
 *                 .build())
 *             .protocols(            
 *                 &#34;FTP&#34;,
 *                 &#34;FTPS&#34;)
 *             .certificate(aws_acm_certificate.example().arn())
 *             .identityProviderType(&#34;API_GATEWAY&#34;)
 *             .url(String.format(&#34;%s%s&#34;, aws_api_gateway_deployment.example().invoke_url(),aws_api_gateway_resource.example().path()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Using Structured Logging Destinations
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudwatch.LogGroup;
 * import com.pulumi.aws.cloudwatch.LogGroupArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.transfer.Server;
 * import com.pulumi.aws.transfer.ServerArgs;
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
 *         var transferLogGroup = new LogGroup(&#34;transferLogGroup&#34;, LogGroupArgs.builder()        
 *             .namePrefix(&#34;transfer_test_&#34;)
 *             .build());
 * 
 *         final var transferAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;transfer.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var iamForTransfer = new Role(&#34;iamForTransfer&#34;, RoleArgs.builder()        
 *             .namePrefix(&#34;iam_for_transfer_&#34;)
 *             .assumeRolePolicy(transferAssumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .managedPolicyArns(&#34;arn:aws:iam::aws:policy/service-role/AWSTransferLoggingAccess&#34;)
 *             .build());
 * 
 *         var transferServer = new Server(&#34;transferServer&#34;, ServerArgs.builder()        
 *             .endpointType(&#34;PUBLIC&#34;)
 *             .loggingRole(iamForTransfer.arn())
 *             .protocols(&#34;SFTP&#34;)
 *             .structuredLogDestinations(transferLogGroup.arn().applyValue(arn -&gt; String.format(&#34;%s:*&#34;, arn)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Transfer Servers using the server `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:transfer/server:Server example s-12345678
 * ```
 *  Certain resource arguments, such as `host_key`, cannot be read via the API and imported into the provider. This provider will display a difference for these arguments the first run after import if declared in the provider configuration for an imported resource.
 * 
 */
@ResourceType(type="aws:transfer/server:Server")
public class Server extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of Transfer Server
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of Transfer Server
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when `protocols` is set to `FTPS`
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certificate;

    /**
     * @return The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when `protocols` is set to `FTPS`
     * 
     */
    public Output<Optional<String>> certificate() {
        return Codegen.optional(this.certificate);
    }
    /**
     * The directory service ID of the directory service you want to connect to with an `identity_provider_type` of `AWS_DIRECTORY_SERVICE`.
     * 
     */
    @Export(name="directoryId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> directoryId;

    /**
     * @return The directory service ID of the directory service you want to connect to with an `identity_provider_type` of `AWS_DIRECTORY_SERVICE`.
     * 
     */
    public Output<Optional<String>> directoryId() {
        return Codegen.optional(this.directoryId);
    }
    /**
     * The domain of the storage system that is used for file transfers. Valid values are: `S3` and `EFS`. The default value is `S3`.
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> domain;

    /**
     * @return The domain of the storage system that is used for file transfers. Valid values are: `S3` and `EFS`. The default value is `S3`.
     * 
     */
    public Output<Optional<String>> domain() {
        return Codegen.optional(this.domain);
    }
    /**
     * The endpoint of the Transfer Server (e.g., `s-12345678.server.transfer.REGION.amazonaws.com`)
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return The endpoint of the Transfer Server (e.g., `s-12345678.server.transfer.REGION.amazonaws.com`)
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
     * 
     */
    @Export(name="endpointDetails", refs={ServerEndpointDetails.class}, tree="[0]")
    private Output</* @Nullable */ ServerEndpointDetails> endpointDetails;

    /**
     * @return The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
     * 
     */
    public Output<Optional<ServerEndpointDetails>> endpointDetails() {
        return Codegen.optional(this.endpointDetails);
    }
    /**
     * The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn&#39;t accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
     * 
     */
    @Export(name="endpointType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endpointType;

    /**
     * @return The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn&#39;t accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
     * 
     */
    public Output<Optional<String>> endpointType() {
        return Codegen.optional(this.endpointType);
    }
    /**
     * A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`. This option only applies to servers configured with a `SERVICE_MANAGED` `identity_provider_type`.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`. This option only applies to servers configured with a `SERVICE_MANAGED` `identity_provider_type`.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * The ARN for a lambda function to use for the Identity provider.
     * 
     */
    @Export(name="function", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> function;

    /**
     * @return The ARN for a lambda function to use for the Identity provider.
     * 
     */
    public Output<Optional<String>> function() {
        return Codegen.optional(this.function);
    }
    /**
     * RSA, ECDSA, or ED25519 private key (e.g., as generated by the `ssh-keygen -t rsa -b 2048 -N &#34;&#34; -m PEM -f my-new-server-key`, `ssh-keygen -t ecdsa -b 256 -N &#34;&#34; -m PEM -f my-new-server-key` or `ssh-keygen -t ed25519 -N &#34;&#34; -f my-new-server-key` commands).
     * 
     */
    @Export(name="hostKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostKey;

    /**
     * @return RSA, ECDSA, or ED25519 private key (e.g., as generated by the `ssh-keygen -t rsa -b 2048 -N &#34;&#34; -m PEM -f my-new-server-key`, `ssh-keygen -t ecdsa -b 256 -N &#34;&#34; -m PEM -f my-new-server-key` or `ssh-keygen -t ed25519 -N &#34;&#34; -f my-new-server-key` commands).
     * 
     */
    public Output<Optional<String>> hostKey() {
        return Codegen.optional(this.hostKey);
    }
    /**
     * This value contains the message-digest algorithm (MD5) hash of the server&#39;s host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
     * 
     */
    @Export(name="hostKeyFingerprint", refs={String.class}, tree="[0]")
    private Output<String> hostKeyFingerprint;

    /**
     * @return This value contains the message-digest algorithm (MD5) hash of the server&#39;s host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
     * 
     */
    public Output<String> hostKeyFingerprint() {
        return this.hostKeyFingerprint;
    }
    /**
     * The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using `AWS_DIRECTORY_SERVICE` will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors. Use the `AWS_LAMBDA` value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the `function` argument.
     * 
     */
    @Export(name="identityProviderType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityProviderType;

    /**
     * @return The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using `AWS_DIRECTORY_SERVICE` will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors. Use the `AWS_LAMBDA` value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the `function` argument.
     * 
     */
    public Output<Optional<String>> identityProviderType() {
        return Codegen.optional(this.identityProviderType);
    }
    /**
     * Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
     * 
     */
    @Export(name="invocationRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> invocationRole;

    /**
     * @return Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
     * 
     */
    public Output<Optional<String>> invocationRole() {
        return Codegen.optional(this.invocationRole);
    }
    /**
     * Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
     * 
     */
    @Export(name="loggingRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loggingRole;

    /**
     * @return Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
     * 
     */
    public Output<Optional<String>> loggingRole() {
        return Codegen.optional(this.loggingRole);
    }
    /**
     * Specify a string to display when users connect to a server. This string is displayed after the user authenticates. The SFTP protocol does not support post-authentication display banners.
     * 
     */
    @Export(name="postAuthenticationLoginBanner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> postAuthenticationLoginBanner;

    /**
     * @return Specify a string to display when users connect to a server. This string is displayed after the user authenticates. The SFTP protocol does not support post-authentication display banners.
     * 
     */
    public Output<Optional<String>> postAuthenticationLoginBanner() {
        return Codegen.optional(this.postAuthenticationLoginBanner);
    }
    /**
     * Specify a string to display when users connect to a server. This string is displayed before the user authenticates.
     * 
     */
    @Export(name="preAuthenticationLoginBanner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> preAuthenticationLoginBanner;

    /**
     * @return Specify a string to display when users connect to a server. This string is displayed before the user authenticates.
     * 
     */
    public Output<Optional<String>> preAuthenticationLoginBanner() {
        return Codegen.optional(this.preAuthenticationLoginBanner);
    }
    /**
     * The protocol settings that are configured for your server.
     * 
     */
    @Export(name="protocolDetails", refs={ServerProtocolDetails.class}, tree="[0]")
    private Output<ServerProtocolDetails> protocolDetails;

    /**
     * @return The protocol settings that are configured for your server.
     * 
     */
    public Output<ServerProtocolDetails> protocolDetails() {
        return this.protocolDetails;
    }
    /**
     * Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server&#39;s endpoint. This defaults to `SFTP` . The available protocols are:
     * 
     */
    @Export(name="protocols", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> protocols;

    /**
     * @return Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server&#39;s endpoint. This defaults to `SFTP` . The available protocols are:
     * 
     */
    public Output<List<String>> protocols() {
        return this.protocols;
    }
    /**
     * Specifies the name of the security policy that is attached to the server. Default value is: `TransferSecurityPolicy-2018-11`. The available values are:
     * * `TransferSecurityPolicy-2024-01`
     * * `TransferSecurityPolicy-2023-05`
     * * `TransferSecurityPolicy-2022-03`
     * * `TransferSecurityPolicy-2020-06`
     * * `TransferSecurityPolicy-2018-11`
     * * `TransferSecurityPolicy-FIPS-2024-01`
     * * `TransferSecurityPolicy-FIPS-2023-05`
     * * `TransferSecurityPolicy-FIPS-2020-06`
     * * `TransferSecurityPolicy-PQ-SSH-Experimental-2023-04`
     * * `TransferSecurityPolicy-PQ-SSH-FIPS-Experimental-2023-04`
     * 
     */
    @Export(name="securityPolicyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityPolicyName;

    /**
     * @return Specifies the name of the security policy that is attached to the server. Default value is: `TransferSecurityPolicy-2018-11`. The available values are:
     * * `TransferSecurityPolicy-2024-01`
     * * `TransferSecurityPolicy-2023-05`
     * * `TransferSecurityPolicy-2022-03`
     * * `TransferSecurityPolicy-2020-06`
     * * `TransferSecurityPolicy-2018-11`
     * * `TransferSecurityPolicy-FIPS-2024-01`
     * * `TransferSecurityPolicy-FIPS-2023-05`
     * * `TransferSecurityPolicy-FIPS-2020-06`
     * * `TransferSecurityPolicy-PQ-SSH-Experimental-2023-04`
     * * `TransferSecurityPolicy-PQ-SSH-FIPS-Experimental-2023-04`
     * 
     */
    public Output<Optional<String>> securityPolicyName() {
        return Codegen.optional(this.securityPolicyName);
    }
    /**
     * A set of ARNs of destinations that will receive structured logs from the transfer server such as CloudWatch Log Group ARNs. If provided this enables the transfer server to emit structured logs to the specified locations.
     * 
     */
    @Export(name="structuredLogDestinations", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> structuredLogDestinations;

    /**
     * @return A set of ARNs of destinations that will receive structured logs from the transfer server such as CloudWatch Log Group ARNs. If provided this enables the transfer server to emit structured logs to the specified locations.
     * 
     */
    public Output<Optional<List<String>>> structuredLogDestinations() {
        return Codegen.optional(this.structuredLogDestinations);
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> url;

    /**
     * @return URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
     * 
     */
    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }
    /**
     * Specifies the workflow details. See Workflow Details below.
     * 
     */
    @Export(name="workflowDetails", refs={ServerWorkflowDetails.class}, tree="[0]")
    private Output</* @Nullable */ ServerWorkflowDetails> workflowDetails;

    /**
     * @return Specifies the workflow details. See Workflow Details below.
     * 
     */
    public Output<Optional<ServerWorkflowDetails>> workflowDetails() {
        return Codegen.optional(this.workflowDetails);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Server(String name) {
        this(name, ServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Server(String name, @Nullable ServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Server(String name, @Nullable ServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:transfer/server:Server", name, args == null ? ServerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Server(String name, Output<String> id, @Nullable ServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:transfer/server:Server", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "hostKey",
                "postAuthenticationLoginBanner",
                "preAuthenticationLoginBanner"
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
    public static Server get(String name, Output<String> id, @Nullable ServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Server(name, id, state, options);
    }
}
