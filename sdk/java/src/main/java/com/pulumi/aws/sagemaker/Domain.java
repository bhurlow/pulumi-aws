// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.sagemaker.DomainArgs;
import com.pulumi.aws.sagemaker.inputs.DomainState;
import com.pulumi.aws.sagemaker.outputs.DomainDefaultSpaceSettings;
import com.pulumi.aws.sagemaker.outputs.DomainDefaultUserSettings;
import com.pulumi.aws.sagemaker.outputs.DomainDomainSettings;
import com.pulumi.aws.sagemaker.outputs.DomainRetentionPolicy;
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
 * Provides a SageMaker Domain resource.
 * 
 * ## Example Usage
 * ### Basic usage
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
 * import com.pulumi.aws.sagemaker.Domain;
 * import com.pulumi.aws.sagemaker.DomainArgs;
 * import com.pulumi.aws.sagemaker.inputs.DomainDefaultUserSettingsArgs;
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
 *         final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
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
 *             .path(&#34;/&#34;)
 *             .assumeRolePolicy(examplePolicyDocument.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var exampleDomain = new Domain(&#34;exampleDomain&#34;, DomainArgs.builder()        
 *             .domainName(&#34;example&#34;)
 *             .authMode(&#34;IAM&#34;)
 *             .vpcId(aws_vpc.example().id())
 *             .subnetIds(aws_subnet.example().id())
 *             .defaultUserSettings(DomainDefaultUserSettingsArgs.builder()
 *                 .executionRole(exampleRole.arn())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Using Custom Images
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sagemaker.Image;
 * import com.pulumi.aws.sagemaker.ImageArgs;
 * import com.pulumi.aws.sagemaker.AppImageConfig;
 * import com.pulumi.aws.sagemaker.AppImageConfigArgs;
 * import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigArgs;
 * import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigKernelSpecArgs;
 * import com.pulumi.aws.sagemaker.ImageVersion;
 * import com.pulumi.aws.sagemaker.ImageVersionArgs;
 * import com.pulumi.aws.sagemaker.Domain;
 * import com.pulumi.aws.sagemaker.DomainArgs;
 * import com.pulumi.aws.sagemaker.inputs.DomainDefaultUserSettingsArgs;
 * import com.pulumi.aws.sagemaker.inputs.DomainDefaultUserSettingsKernelGatewayAppSettingsArgs;
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
 *         var exampleImage = new Image(&#34;exampleImage&#34;, ImageArgs.builder()        
 *             .imageName(&#34;example&#34;)
 *             .roleArn(aws_iam_role.example().arn())
 *             .build());
 * 
 *         var exampleAppImageConfig = new AppImageConfig(&#34;exampleAppImageConfig&#34;, AppImageConfigArgs.builder()        
 *             .appImageConfigName(&#34;example&#34;)
 *             .kernelGatewayImageConfig(AppImageConfigKernelGatewayImageConfigArgs.builder()
 *                 .kernelSpec(AppImageConfigKernelGatewayImageConfigKernelSpecArgs.builder()
 *                     .name(&#34;example&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleImageVersion = new ImageVersion(&#34;exampleImageVersion&#34;, ImageVersionArgs.builder()        
 *             .imageName(exampleImage.id())
 *             .baseImage(&#34;base-image&#34;)
 *             .build());
 * 
 *         var exampleDomain = new Domain(&#34;exampleDomain&#34;, DomainArgs.builder()        
 *             .domainName(&#34;example&#34;)
 *             .authMode(&#34;IAM&#34;)
 *             .vpcId(aws_vpc.example().id())
 *             .subnetIds(aws_subnet.example().id())
 *             .defaultUserSettings(DomainDefaultUserSettingsArgs.builder()
 *                 .executionRole(aws_iam_role.example().arn())
 *                 .kernelGatewayAppSettings(DomainDefaultUserSettingsKernelGatewayAppSettingsArgs.builder()
 *                     .customImages(DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArgs.builder()
 *                         .appImageConfigName(exampleAppImageConfig.appImageConfigName())
 *                         .imageName(exampleImageVersion.imageName())
 *                         .build())
 *                     .build())
 *                 .build())
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
 *  to = aws_sagemaker_domain.test_domain
 * 
 *  id = &#34;d-8jgsjtilstu8&#34; } Using `pulumi import`, import SageMaker Domains using the `id`. For exampleconsole % pulumi import aws_sagemaker_domain.test_domain d-8jgsjtilstu8
 * 
 */
@ResourceType(type="aws:sagemaker/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
     * 
     */
    @Export(name="appNetworkAccessType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> appNetworkAccessType;

    /**
     * @return Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
     * 
     */
    public Output<Optional<String>> appNetworkAccessType() {
        return Codegen.optional(this.appNetworkAccessType);
    }
    /**
     * The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
     * 
     */
    @Export(name="appSecurityGroupManagement", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> appSecurityGroupManagement;

    /**
     * @return The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
     * 
     */
    public Output<Optional<String>> appSecurityGroupManagement() {
        return Codegen.optional(this.appSecurityGroupManagement);
    }
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Domain.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) assigned by AWS to this Domain.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
     * 
     */
    @Export(name="authMode", refs={String.class}, tree="[0]")
    private Output<String> authMode;

    /**
     * @return The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
     * 
     */
    public Output<String> authMode() {
        return this.authMode;
    }
    /**
     * The default space settings. See Default Space Settings below.
     * 
     */
    @Export(name="defaultSpaceSettings", refs={DomainDefaultSpaceSettings.class}, tree="[0]")
    private Output</* @Nullable */ DomainDefaultSpaceSettings> defaultSpaceSettings;

    /**
     * @return The default space settings. See Default Space Settings below.
     * 
     */
    public Output<Optional<DomainDefaultSpaceSettings>> defaultSpaceSettings() {
        return Codegen.optional(this.defaultSpaceSettings);
    }
    /**
     * The default user settings. See Default User Settings below.* `domain_name` - (Required) The domain name.
     * 
     */
    @Export(name="defaultUserSettings", refs={DomainDefaultUserSettings.class}, tree="[0]")
    private Output<DomainDefaultUserSettings> defaultUserSettings;

    /**
     * @return The default user settings. See Default User Settings below.* `domain_name` - (Required) The domain name.
     * 
     */
    public Output<DomainDefaultUserSettings> defaultUserSettings() {
        return this.defaultUserSettings;
    }
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * The domain&#39;s settings.
     * 
     */
    @Export(name="domainSettings", refs={DomainDomainSettings.class}, tree="[0]")
    private Output</* @Nullable */ DomainDomainSettings> domainSettings;

    /**
     * @return The domain&#39;s settings.
     * 
     */
    public Output<Optional<DomainDomainSettings>> domainSettings() {
        return Codegen.optional(this.domainSettings);
    }
    /**
     * The ID of the Amazon Elastic File System (EFS) managed by this Domain.
     * 
     */
    @Export(name="homeEfsFileSystemId", refs={String.class}, tree="[0]")
    private Output<String> homeEfsFileSystemId;

    /**
     * @return The ID of the Amazon Elastic File System (EFS) managed by this Domain.
     * 
     */
    public Output<String> homeEfsFileSystemId() {
        return this.homeEfsFileSystemId;
    }
    /**
     * The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
     * 
     */
    @Export(name="retentionPolicy", refs={DomainRetentionPolicy.class}, tree="[0]")
    private Output</* @Nullable */ DomainRetentionPolicy> retentionPolicy;

    /**
     * @return The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
     * 
     */
    public Output<Optional<DomainRetentionPolicy>> retentionPolicy() {
        return Codegen.optional(this.retentionPolicy);
    }
    /**
     * The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
     * 
     */
    @Export(name="securityGroupIdForDomainBoundary", refs={String.class}, tree="[0]")
    private Output<String> securityGroupIdForDomainBoundary;

    /**
     * @return The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
     * 
     */
    public Output<String> securityGroupIdForDomainBoundary() {
        return this.securityGroupIdForDomainBoundary;
    }
    /**
     * The SSO managed application instance ID.
     * 
     */
    @Export(name="singleSignOnManagedApplicationInstanceId", refs={String.class}, tree="[0]")
    private Output<String> singleSignOnManagedApplicationInstanceId;

    /**
     * @return The SSO managed application instance ID.
     * 
     */
    public Output<String> singleSignOnManagedApplicationInstanceId() {
        return this.singleSignOnManagedApplicationInstanceId;
    }
    /**
     * The VPC subnets that Studio uses for communication.
     * 
     */
    @Export(name="subnetIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> subnetIds;

    /**
     * @return The VPC subnets that Studio uses for communication.
     * 
     */
    public Output<List<String>> subnetIds() {
        return this.subnetIds;
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The domain&#39;s URL.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The domain&#39;s URL.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Domain(String name) {
        this(name, DomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Domain(String name, DomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Domain(String name, DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/domain:Domain", name, args == null ? DomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Domain(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/domain:Domain", name, state, makeResourceOptions(options, id));
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
    public static Domain get(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Domain(name, id, state, options);
    }
}
