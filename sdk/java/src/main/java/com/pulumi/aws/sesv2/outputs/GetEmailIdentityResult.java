// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sesv2.outputs;

import com.pulumi.aws.sesv2.outputs.GetEmailIdentityDkimSigningAttribute;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetEmailIdentityResult {
    /**
     * @return ARN of the Email Identity.
     * 
     */
    private String arn;
    private String configurationSetName;
    /**
     * @return A list of objects that contains at most one element with information about the private key and selector that you want to use to configure DKIM for the identity for Bring Your Own DKIM (BYODKIM) for the identity, or, configures the key length to be used for Easy DKIM.
     * 
     */
    private List<GetEmailIdentityDkimSigningAttribute> dkimSigningAttributes;
    private String emailIdentity;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
     * 
     */
    private String identityType;
    /**
     * @return Key-value mapping of resource tags.
     * 
     */
    private Map<String,String> tags;
    /**
     * @return Specifies whether or not the identity is verified.
     * 
     */
    private Boolean verifiedForSendingStatus;

    private GetEmailIdentityResult() {}
    /**
     * @return ARN of the Email Identity.
     * 
     */
    public String arn() {
        return this.arn;
    }
    public String configurationSetName() {
        return this.configurationSetName;
    }
    /**
     * @return A list of objects that contains at most one element with information about the private key and selector that you want to use to configure DKIM for the identity for Bring Your Own DKIM (BYODKIM) for the identity, or, configures the key length to be used for Easy DKIM.
     * 
     */
    public List<GetEmailIdentityDkimSigningAttribute> dkimSigningAttributes() {
        return this.dkimSigningAttributes;
    }
    public String emailIdentity() {
        return this.emailIdentity;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
     * 
     */
    public String identityType() {
        return this.identityType;
    }
    /**
     * @return Key-value mapping of resource tags.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return Specifies whether or not the identity is verified.
     * 
     */
    public Boolean verifiedForSendingStatus() {
        return this.verifiedForSendingStatus;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEmailIdentityResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arn;
        private String configurationSetName;
        private List<GetEmailIdentityDkimSigningAttribute> dkimSigningAttributes;
        private String emailIdentity;
        private String id;
        private String identityType;
        private Map<String,String> tags;
        private Boolean verifiedForSendingStatus;
        public Builder() {}
        public Builder(GetEmailIdentityResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.configurationSetName = defaults.configurationSetName;
    	      this.dkimSigningAttributes = defaults.dkimSigningAttributes;
    	      this.emailIdentity = defaults.emailIdentity;
    	      this.id = defaults.id;
    	      this.identityType = defaults.identityType;
    	      this.tags = defaults.tags;
    	      this.verifiedForSendingStatus = defaults.verifiedForSendingStatus;
        }

        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder configurationSetName(String configurationSetName) {
            this.configurationSetName = Objects.requireNonNull(configurationSetName);
            return this;
        }
        @CustomType.Setter
        public Builder dkimSigningAttributes(List<GetEmailIdentityDkimSigningAttribute> dkimSigningAttributes) {
            this.dkimSigningAttributes = Objects.requireNonNull(dkimSigningAttributes);
            return this;
        }
        public Builder dkimSigningAttributes(GetEmailIdentityDkimSigningAttribute... dkimSigningAttributes) {
            return dkimSigningAttributes(List.of(dkimSigningAttributes));
        }
        @CustomType.Setter
        public Builder emailIdentity(String emailIdentity) {
            this.emailIdentity = Objects.requireNonNull(emailIdentity);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder identityType(String identityType) {
            this.identityType = Objects.requireNonNull(identityType);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder verifiedForSendingStatus(Boolean verifiedForSendingStatus) {
            this.verifiedForSendingStatus = Objects.requireNonNull(verifiedForSendingStatus);
            return this;
        }
        public GetEmailIdentityResult build() {
            final var o = new GetEmailIdentityResult();
            o.arn = arn;
            o.configurationSetName = configurationSetName;
            o.dkimSigningAttributes = dkimSigningAttributes;
            o.emailIdentity = emailIdentity;
            o.id = id;
            o.identityType = identityType;
            o.tags = tags;
            o.verifiedForSendingStatus = verifiedForSendingStatus;
            return o;
        }
    }
}
