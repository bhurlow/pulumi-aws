// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.servicequotas;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.servicequotas.ServiceQuotaArgs;
import com.pulumi.aws.servicequotas.inputs.ServiceQuotaState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages an individual Service Quota.
 * 
 * &gt; **NOTE:** Global quotas apply to all AWS regions, but can only be accessed in `us-east-1` in the Commercial partition or `us-gov-west-1` in the GovCloud partition. In other regions, the AWS API will return the error `The request failed because the specified service does not exist.`
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.servicequotas.ServiceQuota;
 * import com.pulumi.aws.servicequotas.ServiceQuotaArgs;
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
 *         var example = new ServiceQuota(&#34;example&#34;, ServiceQuotaArgs.builder()        
 *             .quotaCode(&#34;L-F678F1CE&#34;)
 *             .serviceCode(&#34;vpc&#34;)
 *             .value(75)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ~&gt; *NOTE* This resource does not require explicit import and will assume management of an existing service quota on resource creation. `aws_servicequotas_service_quota` can be imported by using the service code and quota code, separated by a front slash (`/`), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:servicequotas/serviceQuota:ServiceQuota example vpc/L-F678F1CE
 * ```
 * 
 */
@ResourceType(type="aws:servicequotas/serviceQuota:ServiceQuota")
public class ServiceQuota extends com.pulumi.resources.CustomResource {
    /**
     * Whether the service quota can be increased.
     * 
     */
    @Export(name="adjustable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> adjustable;

    /**
     * @return Whether the service quota can be increased.
     * 
     */
    public Output<Boolean> adjustable() {
        return this.adjustable;
    }
    /**
     * Amazon Resource Name (ARN) of the service quota.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the service quota.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Default value of the service quota.
     * 
     */
    @Export(name="defaultValue", refs={Double.class}, tree="[0]")
    private Output<Double> defaultValue;

    /**
     * @return Default value of the service quota.
     * 
     */
    public Output<Double> defaultValue() {
        return this.defaultValue;
    }
    /**
     * Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
     * 
     */
    @Export(name="quotaCode", refs={String.class}, tree="[0]")
    private Output<String> quotaCode;

    /**
     * @return Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
     * 
     */
    public Output<String> quotaCode() {
        return this.quotaCode;
    }
    /**
     * Name of the quota.
     * 
     */
    @Export(name="quotaName", refs={String.class}, tree="[0]")
    private Output<String> quotaName;

    /**
     * @return Name of the quota.
     * 
     */
    public Output<String> quotaName() {
        return this.quotaName;
    }
    @Export(name="requestId", refs={String.class}, tree="[0]")
    private Output<String> requestId;

    public Output<String> requestId() {
        return this.requestId;
    }
    @Export(name="requestStatus", refs={String.class}, tree="[0]")
    private Output<String> requestStatus;

    public Output<String> requestStatus() {
        return this.requestStatus;
    }
    /**
     * Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     * 
     */
    @Export(name="serviceCode", refs={String.class}, tree="[0]")
    private Output<String> serviceCode;

    /**
     * @return Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     * 
     */
    public Output<String> serviceCode() {
        return this.serviceCode;
    }
    /**
     * Name of the service.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return Name of the service.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
     * 
     */
    @Export(name="value", refs={Double.class}, tree="[0]")
    private Output<Double> value;

    /**
     * @return Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
     * 
     */
    public Output<Double> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceQuota(String name) {
        this(name, ServiceQuotaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceQuota(String name, ServiceQuotaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceQuota(String name, ServiceQuotaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:servicequotas/serviceQuota:ServiceQuota", name, args == null ? ServiceQuotaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceQuota(String name, Output<String> id, @Nullable ServiceQuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:servicequotas/serviceQuota:ServiceQuota", name, state, makeResourceOptions(options, id));
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
    public static ServiceQuota get(String name, Output<String> id, @Nullable ServiceQuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceQuota(name, id, state, options);
    }
}
