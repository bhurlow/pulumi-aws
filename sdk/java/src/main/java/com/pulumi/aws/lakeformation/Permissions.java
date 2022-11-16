// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lakeformation;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.lakeformation.PermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.PermissionsState;
import com.pulumi.aws.lakeformation.outputs.PermissionsDataLocation;
import com.pulumi.aws.lakeformation.outputs.PermissionsDatabase;
import com.pulumi.aws.lakeformation.outputs.PermissionsLfTag;
import com.pulumi.aws.lakeformation.outputs.PermissionsLfTagPolicy;
import com.pulumi.aws.lakeformation.outputs.PermissionsTable;
import com.pulumi.aws.lakeformation.outputs.PermissionsTableWithColumns;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Grants permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, tables, LF-tags, and LF-tag policies. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
 * 
 * !&gt; **WARNING:** Lake Formation permissions are not in effect by default within AWS. Using this resource will not secure your data and will result in errors if you do not change the security settings for existing resources and the default security settings for new resources. See Default Behavior and `IAMAllowedPrincipals` for additional details.
 * 
 * &gt; **NOTE:** In general, the `principal` should _NOT_ be a Lake Formation administrator or the entity (e.g., IAM role) that is running the deployment. Administrators have implicit permissions. These should be managed by granting or not granting administrator rights using `aws.lakeformation.DataLakeSettings`, _not_ with this resource.
 * 
 * ## Default Behavior and `IAMAllowedPrincipals`
 * 
 * **_Lake Formation permissions are not in effect by default within AWS._** `IAMAllowedPrincipals` (i.e., `IAM_ALLOWED_PRINCIPALS`) conflicts with individual Lake Formation permissions (i.e., non-`IAMAllowedPrincipals` permissions), will cause unexpected behavior, and may result in errors.
 * 
 * When using Lake Formation, choose ONE of the following options as they are mutually exclusive:
 * 
 * 1. Use this resource (`aws.lakeformation.Permissions`), change the default security settings using `aws.lakeformation.DataLakeSettings`, and remove existing `IAMAllowedPrincipals` permissions
 * 2. Use `IAMAllowedPrincipals` without `aws.lakeformation.Permissions`
 * 
 * This example shows removing the `IAMAllowedPrincipals` default security settings and making the caller a Lake Formation admin. Since `create_database_default_permissions` and `create_table_default_permissions` are not set in the `aws.lakeformation.DataLakeSettings` resource, they are cleared.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetSessionContextArgs;
 * import com.pulumi.aws.lakeformation.DataLakeSettings;
 * import com.pulumi.aws.lakeformation.DataLakeSettingsArgs;
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
 *         final var currentCallerIdentity = AwsFunctions.getCallerIdentity();
 * 
 *         final var currentSessionContext = IamFunctions.getSessionContext(GetSessionContextArgs.builder()
 *             .arn(currentCallerIdentity.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.arn()))
 *             .build());
 * 
 *         var test = new DataLakeSettings(&#34;test&#34;, DataLakeSettingsArgs.builder()        
 *             .admins(currentSessionContext.applyValue(getSessionContextResult -&gt; getSessionContextResult.issuerArn()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * To remove existing `IAMAllowedPrincipals` permissions, use the [AWS Lake Formation Console](https://console.aws.amazon.com/lakeformation/) or [AWS CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lakeformation/batch-revoke-permissions.html).
 * 
 * `IAMAllowedPrincipals` is a hook to maintain backwards compatibility with AWS Glue. `IAMAllowedPrincipals` is a pseudo-entity group that acts like a Lake Formation principal. The group includes any IAM users and roles that are allowed access to your Data Catalog resources by your IAM policies.
 * 
 * This is Lake Formation&#39;s default behavior:
 * 
 * * Lake Formation grants `Super` permission to `IAMAllowedPrincipals` on all existing AWS Glue Data Catalog resources.
 * * Lake Formation enables &#34;Use only IAM access control&#34; for new Data Catalog resources.
 * 
 * For more details, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html).
 * 
 * ### Problem Using `IAMAllowedPrincipals`
 * 
 * AWS does not support combining `IAMAllowedPrincipals` permissions and non-`IAMAllowedPrincipals` permissions. Doing so results in unexpected permissions and behaviors. For example, this configuration grants a user `SELECT` on a column in a table.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.glue.CatalogDatabase;
 * import com.pulumi.aws.glue.CatalogDatabaseArgs;
 * import com.pulumi.aws.glue.CatalogTable;
 * import com.pulumi.aws.glue.CatalogTableArgs;
 * import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorArgs;
 * import com.pulumi.aws.lakeformation.Permissions;
 * import com.pulumi.aws.lakeformation.PermissionsArgs;
 * import com.pulumi.aws.lakeformation.inputs.PermissionsTableWithColumnsArgs;
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
 *         var exampleCatalogDatabase = new CatalogDatabase(&#34;exampleCatalogDatabase&#34;, CatalogDatabaseArgs.builder()        
 *             .name(&#34;sadabate&#34;)
 *             .build());
 * 
 *         var exampleCatalogTable = new CatalogTable(&#34;exampleCatalogTable&#34;, CatalogTableArgs.builder()        
 *             .name(&#34;abelt&#34;)
 *             .databaseName(aws_glue_catalog_database.test().name())
 *             .storageDescriptor(CatalogTableStorageDescriptorArgs.builder()
 *                 .columns(CatalogTableStorageDescriptorColumnArgs.builder()
 *                     .name(&#34;event&#34;)
 *                     .type(&#34;string&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var examplePermissions = new Permissions(&#34;examplePermissions&#34;, PermissionsArgs.builder()        
 *             .permissions(&#34;SELECT&#34;)
 *             .principal(&#34;arn:aws:iam:us-east-1:123456789012:user/SanHolo&#34;)
 *             .tableWithColumns(PermissionsTableWithColumnsArgs.builder()
 *                 .databaseName(exampleCatalogTable.databaseName())
 *                 .name(exampleCatalogTable.name())
 *                 .columnNames(&#34;event&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * The resulting permissions depend on whether the table had `IAMAllowedPrincipals` (IAP) permissions or not.
 * 
 * | Result With IAP | Result Without IAP |
 * | ---- | ---- |
 * | `SELECT` column wildcard (i.e., all columns) | `SELECT` on `&#34;event&#34;` (as expected) |
 * 
 * ## Using Lake Formation Permissions
 * 
 * Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. These implicit permissions cannot be revoked _per se_. If this resource reads implicit permissions, it will attempt to revoke them, which causes an error when the resource is destroyed.
 * 
 * There are two ways to avoid these errors. First, and the way we recommend, is to avoid using this resource with principals that have implicit permissions. A second, error-prone option, is to grant explicit permissions (and `permissions_with_grant_option`) to &#34;overwrite&#34; a principal&#39;s implicit permissions, which you can then revoke with this resource. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).
 * 
 * If the `principal` is also a data lake administrator, AWS grants implicit permissions that can cause errors using this resource. For example, AWS implicitly grants a `principal`/administrator `permissions` and `permissions_with_grant_option` of `ALL`, `ALTER`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT` on a table. If you use this resource to explicitly grant the `principal`/administrator `permissions` but _not_ `permissions_with_grant_option` of `ALL`, `ALTER`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT` on the table, this resource will read the implicit `permissions_with_grant_option` and attempt to revoke them when the resource is destroyed. Doing so will cause an `InvalidInputException: No permissions revoked` error because you cannot revoke implicit permissions _per se_. To workaround this problem, explicitly grant the `principal`/administrator `permissions` _and_ `permissions_with_grant_option`, which can then be revoked. Similarly, granting a `principal`/administrator permissions on a table with columns and providing `column_names`, will result in a `InvalidInputException: Permissions modification is invalid` error because you are narrowing the implicit permissions. Instead, set `wildcard` to `true` and remove the `column_names`.
 * 
 * ## Example Usage
 * ### Grant Permissions For A Lake Formation S3 Resource
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lakeformation.Permissions;
 * import com.pulumi.aws.lakeformation.PermissionsArgs;
 * import com.pulumi.aws.lakeformation.inputs.PermissionsDataLocationArgs;
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
 *         var example = new Permissions(&#34;example&#34;, PermissionsArgs.builder()        
 *             .principal(aws_iam_role.workflow_role().arn())
 *             .permissions(&#34;ALL&#34;)
 *             .dataLocation(PermissionsDataLocationArgs.builder()
 *                 .arn(aws_lakeformation_resource.example().arn())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Grant Permissions For A Glue Catalog Database
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lakeformation.Permissions;
 * import com.pulumi.aws.lakeformation.PermissionsArgs;
 * import com.pulumi.aws.lakeformation.inputs.PermissionsDatabaseArgs;
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
 *         var example = new Permissions(&#34;example&#34;, PermissionsArgs.builder()        
 *             .principal(aws_iam_role.workflow_role().arn())
 *             .permissions(            
 *                 &#34;CREATE_TABLE&#34;,
 *                 &#34;ALTER&#34;,
 *                 &#34;DROP&#34;)
 *             .database(PermissionsDatabaseArgs.builder()
 *                 .name(aws_glue_catalog_database.example().name())
 *                 .catalogId(&#34;110376042874&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Grant Permissions Using Tag-Based Access Control
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lakeformation.Permissions;
 * import com.pulumi.aws.lakeformation.PermissionsArgs;
 * import com.pulumi.aws.lakeformation.inputs.PermissionsLfTagPolicyArgs;
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
 *         var test = new Permissions(&#34;test&#34;, PermissionsArgs.builder()        
 *             .principal(aws_iam_role.sales_role().arn())
 *             .permissions(            
 *                 &#34;CREATE_TABLE&#34;,
 *                 &#34;ALTER&#34;,
 *                 &#34;DROP&#34;)
 *             .lfTagPolicy(PermissionsLfTagPolicyArgs.builder()
 *                 .resourceType(&#34;DATABASE&#34;)
 *                 .expressions(                
 *                     PermissionsLfTagPolicyExpressionArgs.builder()
 *                         .key(&#34;Team&#34;)
 *                         .values(&#34;Sales&#34;)
 *                         .build(),
 *                     PermissionsLfTagPolicyExpressionArgs.builder()
 *                         .key(&#34;Environment&#34;)
 *                         .values(                        
 *                             &#34;Dev&#34;,
 *                             &#34;Production&#34;)
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="aws:lakeformation/permissions:Permissions")
public class Permissions extends com.pulumi.resources.CustomResource {
    /**
     * Identifier for the Data Catalog. By default, it is the account ID of the caller.
     * 
     */
    @Export(name="catalogId", type=String.class, parameters={})
    private Output</* @Nullable */ String> catalogId;

    /**
     * @return Identifier for the Data Catalog. By default, it is the account ID of the caller.
     * 
     */
    public Output<Optional<String>> catalogId() {
        return Codegen.optional(this.catalogId);
    }
    /**
     * Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
     * 
     */
    @Export(name="catalogResource", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> catalogResource;

    /**
     * @return Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> catalogResource() {
        return Codegen.optional(this.catalogResource);
    }
    /**
     * Configuration block for a data location resource. Detailed below.
     * 
     */
    @Export(name="dataLocation", type=PermissionsDataLocation.class, parameters={})
    private Output<PermissionsDataLocation> dataLocation;

    /**
     * @return Configuration block for a data location resource. Detailed below.
     * 
     */
    public Output<PermissionsDataLocation> dataLocation() {
        return this.dataLocation;
    }
    /**
     * Configuration block for a database resource. Detailed below.
     * 
     */
    @Export(name="database", type=PermissionsDatabase.class, parameters={})
    private Output<PermissionsDatabase> database;

    /**
     * @return Configuration block for a database resource. Detailed below.
     * 
     */
    public Output<PermissionsDatabase> database() {
        return this.database;
    }
    /**
     * Configuration block for an LF-tag resource. Detailed below.
     * 
     */
    @Export(name="lfTag", type=PermissionsLfTag.class, parameters={})
    private Output<PermissionsLfTag> lfTag;

    /**
     * @return Configuration block for an LF-tag resource. Detailed below.
     * 
     */
    public Output<PermissionsLfTag> lfTag() {
        return this.lfTag;
    }
    /**
     * Configuration block for an LF-tag policy resource. Detailed below.
     * 
     */
    @Export(name="lfTagPolicy", type=PermissionsLfTagPolicy.class, parameters={})
    private Output<PermissionsLfTagPolicy> lfTagPolicy;

    /**
     * @return Configuration block for an LF-tag policy resource. Detailed below.
     * 
     */
    public Output<PermissionsLfTagPolicy> lfTagPolicy() {
        return this.lfTagPolicy;
    }
    /**
     * List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     * 
     */
    @Export(name="permissions", type=List.class, parameters={String.class})
    private Output<List<String>> permissions;

    /**
     * @return List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     * 
     */
    public Output<List<String>> permissions() {
        return this.permissions;
    }
    /**
     * Subset of `permissions` which the principal can pass.
     * 
     */
    @Export(name="permissionsWithGrantOptions", type=List.class, parameters={String.class})
    private Output<List<String>> permissionsWithGrantOptions;

    /**
     * @return Subset of `permissions` which the principal can pass.
     * 
     */
    public Output<List<String>> permissionsWithGrantOptions() {
        return this.permissionsWithGrantOptions;
    }
    /**
     * Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     * 
     */
    @Export(name="principal", type=String.class, parameters={})
    private Output<String> principal;

    /**
     * @return Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     * 
     */
    public Output<String> principal() {
        return this.principal;
    }
    /**
     * Configuration block for a table resource. Detailed below.
     * 
     */
    @Export(name="table", type=PermissionsTable.class, parameters={})
    private Output<PermissionsTable> table;

    /**
     * @return Configuration block for a table resource. Detailed below.
     * 
     */
    public Output<PermissionsTable> table() {
        return this.table;
    }
    /**
     * Configuration block for a table with columns resource. Detailed below.
     * 
     */
    @Export(name="tableWithColumns", type=PermissionsTableWithColumns.class, parameters={})
    private Output<PermissionsTableWithColumns> tableWithColumns;

    /**
     * @return Configuration block for a table with columns resource. Detailed below.
     * 
     */
    public Output<PermissionsTableWithColumns> tableWithColumns() {
        return this.tableWithColumns;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Permissions(String name) {
        this(name, PermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Permissions(String name, PermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Permissions(String name, PermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lakeformation/permissions:Permissions", name, args == null ? PermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Permissions(String name, Output<String> id, @Nullable PermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lakeformation/permissions:Permissions", name, state, makeResourceOptions(options, id));
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
    public static Permissions get(String name, Output<String> id, @Nullable PermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Permissions(name, id, state, options);
    }
}
