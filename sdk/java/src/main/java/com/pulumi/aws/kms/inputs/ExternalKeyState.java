// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExternalKeyState extends com.pulumi.resources.ResourceArgs {

    public static final ExternalKeyState Empty = new ExternalKeyState();

    /**
     * The Amazon Resource Name (ARN) of the key.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the key.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Specifies whether to disable the policy lockout check performed when creating or updating the key&#39;s policy. Setting this value to `true` increases the risk that the key becomes unmanageable. For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the AWS Key Management Service Developer Guide. Defaults to `false`.
     * 
     */
    @Import(name="bypassPolicyLockoutSafetyCheck")
    private @Nullable Output<Boolean> bypassPolicyLockoutSafetyCheck;

    /**
     * @return Specifies whether to disable the policy lockout check performed when creating or updating the key&#39;s policy. Setting this value to `true` increases the risk that the key becomes unmanageable. For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the AWS Key Management Service Developer Guide. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> bypassPolicyLockoutSafetyCheck() {
        return Optional.ofNullable(this.bypassPolicyLockoutSafetyCheck);
    }

    /**
     * Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
     * 
     */
    @Import(name="deletionWindowInDays")
    private @Nullable Output<Integer> deletionWindowInDays;

    /**
     * @return Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
     * 
     */
    public Optional<Output<Integer>> deletionWindowInDays() {
        return Optional.ofNullable(this.deletionWindowInDays);
    }

    /**
     * Description of the key.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the key.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
     * 
     */
    @Import(name="expirationModel")
    private @Nullable Output<String> expirationModel;

    /**
     * @return Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
     * 
     */
    public Optional<Output<String>> expirationModel() {
        return Optional.ofNullable(this.expirationModel);
    }

    /**
     * Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
     * 
     */
    @Import(name="keyMaterialBase64")
    private @Nullable Output<String> keyMaterialBase64;

    /**
     * @return Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
     * 
     */
    public Optional<Output<String>> keyMaterialBase64() {
        return Optional.ofNullable(this.keyMaterialBase64);
    }

    /**
     * The state of the CMK.
     * 
     */
    @Import(name="keyState")
    private @Nullable Output<String> keyState;

    /**
     * @return The state of the CMK.
     * 
     */
    public Optional<Output<String>> keyState() {
        return Optional.ofNullable(this.keyState);
    }

    /**
     * The cryptographic operations for which you can use the CMK.
     * 
     */
    @Import(name="keyUsage")
    private @Nullable Output<String> keyUsage;

    /**
     * @return The cryptographic operations for which you can use the CMK.
     * 
     */
    public Optional<Output<String>> keyUsage() {
        return Optional.ofNullable(this.keyUsage);
    }

    /**
     * Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
     * 
     */
    @Import(name="multiRegion")
    private @Nullable Output<Boolean> multiRegion;

    /**
     * @return Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> multiRegion() {
        return Optional.ofNullable(this.multiRegion);
    }

    /**
     * A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    /**
     * A key-value map of tags to assign to the key. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A key-value map of tags to assign to the key. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     * 
     */
    @Import(name="validTo")
    private @Nullable Output<String> validTo;

    /**
     * @return Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     * 
     */
    public Optional<Output<String>> validTo() {
        return Optional.ofNullable(this.validTo);
    }

    private ExternalKeyState() {}

    private ExternalKeyState(ExternalKeyState $) {
        this.arn = $.arn;
        this.bypassPolicyLockoutSafetyCheck = $.bypassPolicyLockoutSafetyCheck;
        this.deletionWindowInDays = $.deletionWindowInDays;
        this.description = $.description;
        this.enabled = $.enabled;
        this.expirationModel = $.expirationModel;
        this.keyMaterialBase64 = $.keyMaterialBase64;
        this.keyState = $.keyState;
        this.keyUsage = $.keyUsage;
        this.multiRegion = $.multiRegion;
        this.policy = $.policy;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.validTo = $.validTo;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExternalKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExternalKeyState $;

        public Builder() {
            $ = new ExternalKeyState();
        }

        public Builder(ExternalKeyState defaults) {
            $ = new ExternalKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the key.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the key.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param bypassPolicyLockoutSafetyCheck Specifies whether to disable the policy lockout check performed when creating or updating the key&#39;s policy. Setting this value to `true` increases the risk that the key becomes unmanageable. For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the AWS Key Management Service Developer Guide. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder bypassPolicyLockoutSafetyCheck(@Nullable Output<Boolean> bypassPolicyLockoutSafetyCheck) {
            $.bypassPolicyLockoutSafetyCheck = bypassPolicyLockoutSafetyCheck;
            return this;
        }

        /**
         * @param bypassPolicyLockoutSafetyCheck Specifies whether to disable the policy lockout check performed when creating or updating the key&#39;s policy. Setting this value to `true` increases the risk that the key becomes unmanageable. For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the AWS Key Management Service Developer Guide. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder bypassPolicyLockoutSafetyCheck(Boolean bypassPolicyLockoutSafetyCheck) {
            return bypassPolicyLockoutSafetyCheck(Output.of(bypassPolicyLockoutSafetyCheck));
        }

        /**
         * @param deletionWindowInDays Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
         * 
         * @return builder
         * 
         */
        public Builder deletionWindowInDays(@Nullable Output<Integer> deletionWindowInDays) {
            $.deletionWindowInDays = deletionWindowInDays;
            return this;
        }

        /**
         * @param deletionWindowInDays Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
         * 
         * @return builder
         * 
         */
        public Builder deletionWindowInDays(Integer deletionWindowInDays) {
            return deletionWindowInDays(Output.of(deletionWindowInDays));
        }

        /**
         * @param description Description of the key.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the key.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enabled Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param expirationModel Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
         * 
         * @return builder
         * 
         */
        public Builder expirationModel(@Nullable Output<String> expirationModel) {
            $.expirationModel = expirationModel;
            return this;
        }

        /**
         * @param expirationModel Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
         * 
         * @return builder
         * 
         */
        public Builder expirationModel(String expirationModel) {
            return expirationModel(Output.of(expirationModel));
        }

        /**
         * @param keyMaterialBase64 Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
         * 
         * @return builder
         * 
         */
        public Builder keyMaterialBase64(@Nullable Output<String> keyMaterialBase64) {
            $.keyMaterialBase64 = keyMaterialBase64;
            return this;
        }

        /**
         * @param keyMaterialBase64 Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
         * 
         * @return builder
         * 
         */
        public Builder keyMaterialBase64(String keyMaterialBase64) {
            return keyMaterialBase64(Output.of(keyMaterialBase64));
        }

        /**
         * @param keyState The state of the CMK.
         * 
         * @return builder
         * 
         */
        public Builder keyState(@Nullable Output<String> keyState) {
            $.keyState = keyState;
            return this;
        }

        /**
         * @param keyState The state of the CMK.
         * 
         * @return builder
         * 
         */
        public Builder keyState(String keyState) {
            return keyState(Output.of(keyState));
        }

        /**
         * @param keyUsage The cryptographic operations for which you can use the CMK.
         * 
         * @return builder
         * 
         */
        public Builder keyUsage(@Nullable Output<String> keyUsage) {
            $.keyUsage = keyUsage;
            return this;
        }

        /**
         * @param keyUsage The cryptographic operations for which you can use the CMK.
         * 
         * @return builder
         * 
         */
        public Builder keyUsage(String keyUsage) {
            return keyUsage(Output.of(keyUsage));
        }

        /**
         * @param multiRegion Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder multiRegion(@Nullable Output<Boolean> multiRegion) {
            $.multiRegion = multiRegion;
            return this;
        }

        /**
         * @param multiRegion Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder multiRegion(Boolean multiRegion) {
            return multiRegion(Output.of(multiRegion));
        }

        /**
         * @param policy A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param tags A key-value map of tags to assign to the key. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A key-value map of tags to assign to the key. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param validTo Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
         * 
         * @return builder
         * 
         */
        public Builder validTo(@Nullable Output<String> validTo) {
            $.validTo = validTo;
            return this;
        }

        /**
         * @param validTo Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
         * 
         * @return builder
         * 
         */
        public Builder validTo(String validTo) {
            return validTo(Output.of(validTo));
        }

        public ExternalKeyState build() {
            return $;
        }
    }

}
