// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.servicequotas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceQuotaPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceQuotaPlainArgs Empty = new GetServiceQuotaPlainArgs();

    /**
     * Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quota_code` or `quota_name` must be specified.
     * 
     */
    @Import(name="quotaCode")
    private @Nullable String quotaCode;

    /**
     * @return Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quota_code` or `quota_name` must be specified.
     * 
     */
    public Optional<String> quotaCode() {
        return Optional.ofNullable(this.quotaCode);
    }

    /**
     * Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quota_name` or `quota_code` must be specified.
     * 
     */
    @Import(name="quotaName")
    private @Nullable String quotaName;

    /**
     * @return Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quota_name` or `quota_code` must be specified.
     * 
     */
    public Optional<String> quotaName() {
        return Optional.ofNullable(this.quotaName);
    }

    /**
     * Service code for the quota. Available values can be found with the `aws.servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     * 
     */
    @Import(name="serviceCode", required=true)
    private String serviceCode;

    /**
     * @return Service code for the quota. Available values can be found with the `aws.servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     * 
     */
    public String serviceCode() {
        return this.serviceCode;
    }

    private GetServiceQuotaPlainArgs() {}

    private GetServiceQuotaPlainArgs(GetServiceQuotaPlainArgs $) {
        this.quotaCode = $.quotaCode;
        this.quotaName = $.quotaName;
        this.serviceCode = $.serviceCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceQuotaPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceQuotaPlainArgs $;

        public Builder() {
            $ = new GetServiceQuotaPlainArgs();
        }

        public Builder(GetServiceQuotaPlainArgs defaults) {
            $ = new GetServiceQuotaPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param quotaCode Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quota_code` or `quota_name` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder quotaCode(@Nullable String quotaCode) {
            $.quotaCode = quotaCode;
            return this;
        }

        /**
         * @param quotaName Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quota_name` or `quota_code` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder quotaName(@Nullable String quotaName) {
            $.quotaName = quotaName;
            return this;
        }

        /**
         * @param serviceCode Service code for the quota. Available values can be found with the `aws.servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
         * 
         * @return builder
         * 
         */
        public Builder serviceCode(String serviceCode) {
            $.serviceCode = serviceCode;
            return this;
        }

        public GetServiceQuotaPlainArgs build() {
            $.serviceCode = Objects.requireNonNull($.serviceCode, "expected parameter 'serviceCode' to be non-null");
            return $;
        }
    }

}
