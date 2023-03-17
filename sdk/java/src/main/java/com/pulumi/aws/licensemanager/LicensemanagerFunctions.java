// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.licensemanager;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsPlainArgs;
import com.pulumi.aws.licensemanager.inputs.GetReceivedLicenseArgs;
import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensePlainArgs;
import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesPlainArgs;
import com.pulumi.aws.licensemanager.outputs.GetLicenseGrantsResult;
import com.pulumi.aws.licensemanager.outputs.GetReceivedLicenseResult;
import com.pulumi.aws.licensemanager.outputs.GetReceivedLicensesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class LicensemanagerFunctions {
    /**
     * This resource can be used to get a set of license grant ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license grant ARNs granted to your account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.AwsFunctions;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
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
     *         final var current = AwsFunctions.getCallerIdentity();
     * 
     *         final var test = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
     *             .filters(GetLicenseGrantsFilterArgs.builder()
     *                 .name(&#34;GranteePrincipalARN&#34;)
     *                 .values(String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLicenseGrantsResult> getLicenseGrants() {
        return getLicenseGrants(GetLicenseGrantsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license grant ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license grant ARNs granted to your account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.AwsFunctions;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
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
     *         final var current = AwsFunctions.getCallerIdentity();
     * 
     *         final var test = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
     *             .filters(GetLicenseGrantsFilterArgs.builder()
     *                 .name(&#34;GranteePrincipalARN&#34;)
     *                 .values(String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLicenseGrantsResult> getLicenseGrantsPlain() {
        return getLicenseGrantsPlain(GetLicenseGrantsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license grant ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license grant ARNs granted to your account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.AwsFunctions;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
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
     *         final var current = AwsFunctions.getCallerIdentity();
     * 
     *         final var test = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
     *             .filters(GetLicenseGrantsFilterArgs.builder()
     *                 .name(&#34;GranteePrincipalARN&#34;)
     *                 .values(String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLicenseGrantsResult> getLicenseGrants(GetLicenseGrantsArgs args) {
        return getLicenseGrants(args, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license grant ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license grant ARNs granted to your account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.AwsFunctions;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
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
     *         final var current = AwsFunctions.getCallerIdentity();
     * 
     *         final var test = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
     *             .filters(GetLicenseGrantsFilterArgs.builder()
     *                 .name(&#34;GranteePrincipalARN&#34;)
     *                 .values(String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLicenseGrantsResult> getLicenseGrantsPlain(GetLicenseGrantsPlainArgs args) {
        return getLicenseGrantsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license grant ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license grant ARNs granted to your account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.AwsFunctions;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
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
     *         final var current = AwsFunctions.getCallerIdentity();
     * 
     *         final var test = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
     *             .filters(GetLicenseGrantsFilterArgs.builder()
     *                 .name(&#34;GranteePrincipalARN&#34;)
     *                 .values(String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLicenseGrantsResult> getLicenseGrants(GetLicenseGrantsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("aws:licensemanager/getLicenseGrants:getLicenseGrants", TypeShape.of(GetLicenseGrantsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This resource can be used to get a set of license grant ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license grant ARNs granted to your account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.AwsFunctions;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
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
     *         final var current = AwsFunctions.getCallerIdentity();
     * 
     *         final var test = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
     *             .filters(GetLicenseGrantsFilterArgs.builder()
     *                 .name(&#34;GranteePrincipalARN&#34;)
     *                 .values(String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLicenseGrantsResult> getLicenseGrantsPlain(GetLicenseGrantsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("aws:licensemanager/getLicenseGrants:getLicenseGrants", TypeShape.of(GetLicenseGrantsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.
     * 
     * ## Example Usage
     * 
     * The following shows getting the received license data using and ARN.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicenseArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicense(GetReceivedLicenseArgs.builder()
     *             .licenseArn(&#34;arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetReceivedLicenseResult> getReceivedLicense(GetReceivedLicenseArgs args) {
        return getReceivedLicense(args, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.
     * 
     * ## Example Usage
     * 
     * The following shows getting the received license data using and ARN.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicenseArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicense(GetReceivedLicenseArgs.builder()
     *             .licenseArn(&#34;arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetReceivedLicenseResult> getReceivedLicensePlain(GetReceivedLicensePlainArgs args) {
        return getReceivedLicensePlain(args, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.
     * 
     * ## Example Usage
     * 
     * The following shows getting the received license data using and ARN.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicenseArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicense(GetReceivedLicenseArgs.builder()
     *             .licenseArn(&#34;arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetReceivedLicenseResult> getReceivedLicense(GetReceivedLicenseArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("aws:licensemanager/getReceivedLicense:getReceivedLicense", TypeShape.of(GetReceivedLicenseResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.
     * 
     * ## Example Usage
     * 
     * The following shows getting the received license data using and ARN.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicenseArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicense(GetReceivedLicenseArgs.builder()
     *             .licenseArn(&#34;arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetReceivedLicenseResult> getReceivedLicensePlain(GetReceivedLicensePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("aws:licensemanager/getReceivedLicense:getReceivedLicense", TypeShape.of(GetReceivedLicenseResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This resource can be used to get a set of license ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license ARNs issued from the AWS marketplace. Providing no filter, would provide all license ARNs for the entire account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
     *             .filters(GetReceivedLicensesFilterArgs.builder()
     *                 .name(&#34;IssuerName&#34;)
     *                 .values(&#34;AWS/Marketplace&#34;)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetReceivedLicensesResult> getReceivedLicenses() {
        return getReceivedLicenses(GetReceivedLicensesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license ARNs issued from the AWS marketplace. Providing no filter, would provide all license ARNs for the entire account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
     *             .filters(GetReceivedLicensesFilterArgs.builder()
     *                 .name(&#34;IssuerName&#34;)
     *                 .values(&#34;AWS/Marketplace&#34;)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetReceivedLicensesResult> getReceivedLicensesPlain() {
        return getReceivedLicensesPlain(GetReceivedLicensesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license ARNs issued from the AWS marketplace. Providing no filter, would provide all license ARNs for the entire account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
     *             .filters(GetReceivedLicensesFilterArgs.builder()
     *                 .name(&#34;IssuerName&#34;)
     *                 .values(&#34;AWS/Marketplace&#34;)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetReceivedLicensesResult> getReceivedLicenses(GetReceivedLicensesArgs args) {
        return getReceivedLicenses(args, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license ARNs issued from the AWS marketplace. Providing no filter, would provide all license ARNs for the entire account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
     *             .filters(GetReceivedLicensesFilterArgs.builder()
     *                 .name(&#34;IssuerName&#34;)
     *                 .values(&#34;AWS/Marketplace&#34;)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetReceivedLicensesResult> getReceivedLicensesPlain(GetReceivedLicensesPlainArgs args) {
        return getReceivedLicensesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This resource can be used to get a set of license ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license ARNs issued from the AWS marketplace. Providing no filter, would provide all license ARNs for the entire account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
     *             .filters(GetReceivedLicensesFilterArgs.builder()
     *                 .name(&#34;IssuerName&#34;)
     *                 .values(&#34;AWS/Marketplace&#34;)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetReceivedLicensesResult> getReceivedLicenses(GetReceivedLicensesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("aws:licensemanager/getReceivedLicenses:getReceivedLicenses", TypeShape.of(GetReceivedLicensesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This resource can be used to get a set of license ARNs matching a filter.
     * 
     * ## Example Usage
     * 
     * The following shows getting all license ARNs issued from the AWS marketplace. Providing no filter, would provide all license ARNs for the entire account.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
     * import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
     *         final var test = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
     *             .filters(GetReceivedLicensesFilterArgs.builder()
     *                 .name(&#34;IssuerName&#34;)
     *                 .values(&#34;AWS/Marketplace&#34;)
     *                 .build())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetReceivedLicensesResult> getReceivedLicensesPlain(GetReceivedLicensesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("aws:licensemanager/getReceivedLicenses:getReceivedLicenses", TypeShape.of(GetReceivedLicensesResult.class), args, Utilities.withVersion(options));
    }
}
