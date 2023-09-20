// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53domains.inputs;

import com.pulumi.aws.route53domains.inputs.RegisteredDomainAdminContactArgs;
import com.pulumi.aws.route53domains.inputs.RegisteredDomainNameServerArgs;
import com.pulumi.aws.route53domains.inputs.RegisteredDomainRegistrantContactArgs;
import com.pulumi.aws.route53domains.inputs.RegisteredDomainTechContactArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegisteredDomainState extends com.pulumi.resources.ResourceArgs {

    public static final RegisteredDomainState Empty = new RegisteredDomainState();

    /**
     * Email address to contact to report incorrect contact information for a domain, to report that the domain is being used to send spam, to report that someone is cybersquatting on a domain name, or report some other type of abuse.
     * 
     */
    @Import(name="abuseContactEmail")
    private @Nullable Output<String> abuseContactEmail;

    /**
     * @return Email address to contact to report incorrect contact information for a domain, to report that the domain is being used to send spam, to report that someone is cybersquatting on a domain name, or report some other type of abuse.
     * 
     */
    public Optional<Output<String>> abuseContactEmail() {
        return Optional.ofNullable(this.abuseContactEmail);
    }

    /**
     * Phone number for reporting abuse.
     * 
     */
    @Import(name="abuseContactPhone")
    private @Nullable Output<String> abuseContactPhone;

    /**
     * @return Phone number for reporting abuse.
     * 
     */
    public Optional<Output<String>> abuseContactPhone() {
        return Optional.ofNullable(this.abuseContactPhone);
    }

    /**
     * Details about the domain administrative contact.
     * 
     */
    @Import(name="adminContact")
    private @Nullable Output<RegisteredDomainAdminContactArgs> adminContact;

    /**
     * @return Details about the domain administrative contact.
     * 
     */
    public Optional<Output<RegisteredDomainAdminContactArgs>> adminContact() {
        return Optional.ofNullable(this.adminContact);
    }

    /**
     * Whether domain administrative contact information is concealed from WHOIS queries. Default: `true`.
     * 
     */
    @Import(name="adminPrivacy")
    private @Nullable Output<Boolean> adminPrivacy;

    /**
     * @return Whether domain administrative contact information is concealed from WHOIS queries. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> adminPrivacy() {
        return Optional.ofNullable(this.adminPrivacy);
    }

    /**
     * Whether the domain registration is set to renew automatically. Default: `true`.
     * 
     */
    @Import(name="autoRenew")
    private @Nullable Output<Boolean> autoRenew;

    /**
     * @return Whether the domain registration is set to renew automatically. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }

    /**
     * The date when the domain was created as found in the response to a WHOIS query.
     * 
     */
    @Import(name="creationDate")
    private @Nullable Output<String> creationDate;

    /**
     * @return The date when the domain was created as found in the response to a WHOIS query.
     * 
     */
    public Optional<Output<String>> creationDate() {
        return Optional.ofNullable(this.creationDate);
    }

    /**
     * The name of the registered domain.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return The name of the registered domain.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * The date when the registration for the domain is set to expire.
     * 
     */
    @Import(name="expirationDate")
    private @Nullable Output<String> expirationDate;

    /**
     * @return The date when the registration for the domain is set to expire.
     * 
     */
    public Optional<Output<String>> expirationDate() {
        return Optional.ofNullable(this.expirationDate);
    }

    /**
     * The list of nameservers for the domain.
     * 
     */
    @Import(name="nameServers")
    private @Nullable Output<List<RegisteredDomainNameServerArgs>> nameServers;

    /**
     * @return The list of nameservers for the domain.
     * 
     */
    public Optional<Output<List<RegisteredDomainNameServerArgs>>> nameServers() {
        return Optional.ofNullable(this.nameServers);
    }

    /**
     * Details about the domain registrant.
     * 
     */
    @Import(name="registrantContact")
    private @Nullable Output<RegisteredDomainRegistrantContactArgs> registrantContact;

    /**
     * @return Details about the domain registrant.
     * 
     */
    public Optional<Output<RegisteredDomainRegistrantContactArgs>> registrantContact() {
        return Optional.ofNullable(this.registrantContact);
    }

    /**
     * Whether domain registrant contact information is concealed from WHOIS queries. Default: `true`.
     * 
     */
    @Import(name="registrantPrivacy")
    private @Nullable Output<Boolean> registrantPrivacy;

    /**
     * @return Whether domain registrant contact information is concealed from WHOIS queries. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> registrantPrivacy() {
        return Optional.ofNullable(this.registrantPrivacy);
    }

    /**
     * Name of the registrar of the domain as identified in the registry.
     * 
     */
    @Import(name="registrarName")
    private @Nullable Output<String> registrarName;

    /**
     * @return Name of the registrar of the domain as identified in the registry.
     * 
     */
    public Optional<Output<String>> registrarName() {
        return Optional.ofNullable(this.registrarName);
    }

    /**
     * Web address of the registrar.
     * 
     */
    @Import(name="registrarUrl")
    private @Nullable Output<String> registrarUrl;

    /**
     * @return Web address of the registrar.
     * 
     */
    public Optional<Output<String>> registrarUrl() {
        return Optional.ofNullable(this.registrarUrl);
    }

    /**
     * Reseller of the domain.
     * 
     */
    @Import(name="reseller")
    private @Nullable Output<String> reseller;

    /**
     * @return Reseller of the domain.
     * 
     */
    public Optional<Output<String>> reseller() {
        return Optional.ofNullable(this.reseller);
    }

    /**
     * List of [domain name status codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
     * 
     */
    @Import(name="statusLists")
    private @Nullable Output<List<String>> statusLists;

    /**
     * @return List of [domain name status codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
     * 
     */
    public Optional<Output<List<String>>> statusLists() {
        return Optional.ofNullable(this.statusLists);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Details about the domain technical contact.
     * 
     */
    @Import(name="techContact")
    private @Nullable Output<RegisteredDomainTechContactArgs> techContact;

    /**
     * @return Details about the domain technical contact.
     * 
     */
    public Optional<Output<RegisteredDomainTechContactArgs>> techContact() {
        return Optional.ofNullable(this.techContact);
    }

    /**
     * Whether domain technical contact information is concealed from WHOIS queries. Default: `true`.
     * 
     */
    @Import(name="techPrivacy")
    private @Nullable Output<Boolean> techPrivacy;

    /**
     * @return Whether domain technical contact information is concealed from WHOIS queries. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> techPrivacy() {
        return Optional.ofNullable(this.techPrivacy);
    }

    /**
     * Whether the domain is locked for transfer. Default: `true`.
     * 
     */
    @Import(name="transferLock")
    private @Nullable Output<Boolean> transferLock;

    /**
     * @return Whether the domain is locked for transfer. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> transferLock() {
        return Optional.ofNullable(this.transferLock);
    }

    /**
     * The last updated date of the domain as found in the response to a WHOIS query.
     * 
     */
    @Import(name="updatedDate")
    private @Nullable Output<String> updatedDate;

    /**
     * @return The last updated date of the domain as found in the response to a WHOIS query.
     * 
     */
    public Optional<Output<String>> updatedDate() {
        return Optional.ofNullable(this.updatedDate);
    }

    /**
     * The fully qualified name of the WHOIS server that can answer the WHOIS query for the domain.
     * 
     */
    @Import(name="whoisServer")
    private @Nullable Output<String> whoisServer;

    /**
     * @return The fully qualified name of the WHOIS server that can answer the WHOIS query for the domain.
     * 
     */
    public Optional<Output<String>> whoisServer() {
        return Optional.ofNullable(this.whoisServer);
    }

    private RegisteredDomainState() {}

    private RegisteredDomainState(RegisteredDomainState $) {
        this.abuseContactEmail = $.abuseContactEmail;
        this.abuseContactPhone = $.abuseContactPhone;
        this.adminContact = $.adminContact;
        this.adminPrivacy = $.adminPrivacy;
        this.autoRenew = $.autoRenew;
        this.creationDate = $.creationDate;
        this.domainName = $.domainName;
        this.expirationDate = $.expirationDate;
        this.nameServers = $.nameServers;
        this.registrantContact = $.registrantContact;
        this.registrantPrivacy = $.registrantPrivacy;
        this.registrarName = $.registrarName;
        this.registrarUrl = $.registrarUrl;
        this.reseller = $.reseller;
        this.statusLists = $.statusLists;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.techContact = $.techContact;
        this.techPrivacy = $.techPrivacy;
        this.transferLock = $.transferLock;
        this.updatedDate = $.updatedDate;
        this.whoisServer = $.whoisServer;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegisteredDomainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegisteredDomainState $;

        public Builder() {
            $ = new RegisteredDomainState();
        }

        public Builder(RegisteredDomainState defaults) {
            $ = new RegisteredDomainState(Objects.requireNonNull(defaults));
        }

        /**
         * @param abuseContactEmail Email address to contact to report incorrect contact information for a domain, to report that the domain is being used to send spam, to report that someone is cybersquatting on a domain name, or report some other type of abuse.
         * 
         * @return builder
         * 
         */
        public Builder abuseContactEmail(@Nullable Output<String> abuseContactEmail) {
            $.abuseContactEmail = abuseContactEmail;
            return this;
        }

        /**
         * @param abuseContactEmail Email address to contact to report incorrect contact information for a domain, to report that the domain is being used to send spam, to report that someone is cybersquatting on a domain name, or report some other type of abuse.
         * 
         * @return builder
         * 
         */
        public Builder abuseContactEmail(String abuseContactEmail) {
            return abuseContactEmail(Output.of(abuseContactEmail));
        }

        /**
         * @param abuseContactPhone Phone number for reporting abuse.
         * 
         * @return builder
         * 
         */
        public Builder abuseContactPhone(@Nullable Output<String> abuseContactPhone) {
            $.abuseContactPhone = abuseContactPhone;
            return this;
        }

        /**
         * @param abuseContactPhone Phone number for reporting abuse.
         * 
         * @return builder
         * 
         */
        public Builder abuseContactPhone(String abuseContactPhone) {
            return abuseContactPhone(Output.of(abuseContactPhone));
        }

        /**
         * @param adminContact Details about the domain administrative contact.
         * 
         * @return builder
         * 
         */
        public Builder adminContact(@Nullable Output<RegisteredDomainAdminContactArgs> adminContact) {
            $.adminContact = adminContact;
            return this;
        }

        /**
         * @param adminContact Details about the domain administrative contact.
         * 
         * @return builder
         * 
         */
        public Builder adminContact(RegisteredDomainAdminContactArgs adminContact) {
            return adminContact(Output.of(adminContact));
        }

        /**
         * @param adminPrivacy Whether domain administrative contact information is concealed from WHOIS queries. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder adminPrivacy(@Nullable Output<Boolean> adminPrivacy) {
            $.adminPrivacy = adminPrivacy;
            return this;
        }

        /**
         * @param adminPrivacy Whether domain administrative contact information is concealed from WHOIS queries. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder adminPrivacy(Boolean adminPrivacy) {
            return adminPrivacy(Output.of(adminPrivacy));
        }

        /**
         * @param autoRenew Whether the domain registration is set to renew automatically. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(@Nullable Output<Boolean> autoRenew) {
            $.autoRenew = autoRenew;
            return this;
        }

        /**
         * @param autoRenew Whether the domain registration is set to renew automatically. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(Boolean autoRenew) {
            return autoRenew(Output.of(autoRenew));
        }

        /**
         * @param creationDate The date when the domain was created as found in the response to a WHOIS query.
         * 
         * @return builder
         * 
         */
        public Builder creationDate(@Nullable Output<String> creationDate) {
            $.creationDate = creationDate;
            return this;
        }

        /**
         * @param creationDate The date when the domain was created as found in the response to a WHOIS query.
         * 
         * @return builder
         * 
         */
        public Builder creationDate(String creationDate) {
            return creationDate(Output.of(creationDate));
        }

        /**
         * @param domainName The name of the registered domain.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The name of the registered domain.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param expirationDate The date when the registration for the domain is set to expire.
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(@Nullable Output<String> expirationDate) {
            $.expirationDate = expirationDate;
            return this;
        }

        /**
         * @param expirationDate The date when the registration for the domain is set to expire.
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(String expirationDate) {
            return expirationDate(Output.of(expirationDate));
        }

        /**
         * @param nameServers The list of nameservers for the domain.
         * 
         * @return builder
         * 
         */
        public Builder nameServers(@Nullable Output<List<RegisteredDomainNameServerArgs>> nameServers) {
            $.nameServers = nameServers;
            return this;
        }

        /**
         * @param nameServers The list of nameservers for the domain.
         * 
         * @return builder
         * 
         */
        public Builder nameServers(List<RegisteredDomainNameServerArgs> nameServers) {
            return nameServers(Output.of(nameServers));
        }

        /**
         * @param nameServers The list of nameservers for the domain.
         * 
         * @return builder
         * 
         */
        public Builder nameServers(RegisteredDomainNameServerArgs... nameServers) {
            return nameServers(List.of(nameServers));
        }

        /**
         * @param registrantContact Details about the domain registrant.
         * 
         * @return builder
         * 
         */
        public Builder registrantContact(@Nullable Output<RegisteredDomainRegistrantContactArgs> registrantContact) {
            $.registrantContact = registrantContact;
            return this;
        }

        /**
         * @param registrantContact Details about the domain registrant.
         * 
         * @return builder
         * 
         */
        public Builder registrantContact(RegisteredDomainRegistrantContactArgs registrantContact) {
            return registrantContact(Output.of(registrantContact));
        }

        /**
         * @param registrantPrivacy Whether domain registrant contact information is concealed from WHOIS queries. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder registrantPrivacy(@Nullable Output<Boolean> registrantPrivacy) {
            $.registrantPrivacy = registrantPrivacy;
            return this;
        }

        /**
         * @param registrantPrivacy Whether domain registrant contact information is concealed from WHOIS queries. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder registrantPrivacy(Boolean registrantPrivacy) {
            return registrantPrivacy(Output.of(registrantPrivacy));
        }

        /**
         * @param registrarName Name of the registrar of the domain as identified in the registry.
         * 
         * @return builder
         * 
         */
        public Builder registrarName(@Nullable Output<String> registrarName) {
            $.registrarName = registrarName;
            return this;
        }

        /**
         * @param registrarName Name of the registrar of the domain as identified in the registry.
         * 
         * @return builder
         * 
         */
        public Builder registrarName(String registrarName) {
            return registrarName(Output.of(registrarName));
        }

        /**
         * @param registrarUrl Web address of the registrar.
         * 
         * @return builder
         * 
         */
        public Builder registrarUrl(@Nullable Output<String> registrarUrl) {
            $.registrarUrl = registrarUrl;
            return this;
        }

        /**
         * @param registrarUrl Web address of the registrar.
         * 
         * @return builder
         * 
         */
        public Builder registrarUrl(String registrarUrl) {
            return registrarUrl(Output.of(registrarUrl));
        }

        /**
         * @param reseller Reseller of the domain.
         * 
         * @return builder
         * 
         */
        public Builder reseller(@Nullable Output<String> reseller) {
            $.reseller = reseller;
            return this;
        }

        /**
         * @param reseller Reseller of the domain.
         * 
         * @return builder
         * 
         */
        public Builder reseller(String reseller) {
            return reseller(Output.of(reseller));
        }

        /**
         * @param statusLists List of [domain name status codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
         * 
         * @return builder
         * 
         */
        public Builder statusLists(@Nullable Output<List<String>> statusLists) {
            $.statusLists = statusLists;
            return this;
        }

        /**
         * @param statusLists List of [domain name status codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
         * 
         * @return builder
         * 
         */
        public Builder statusLists(List<String> statusLists) {
            return statusLists(Output.of(statusLists));
        }

        /**
         * @param statusLists List of [domain name status codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
         * 
         * @return builder
         * 
         */
        public Builder statusLists(String... statusLists) {
            return statusLists(List.of(statusLists));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param techContact Details about the domain technical contact.
         * 
         * @return builder
         * 
         */
        public Builder techContact(@Nullable Output<RegisteredDomainTechContactArgs> techContact) {
            $.techContact = techContact;
            return this;
        }

        /**
         * @param techContact Details about the domain technical contact.
         * 
         * @return builder
         * 
         */
        public Builder techContact(RegisteredDomainTechContactArgs techContact) {
            return techContact(Output.of(techContact));
        }

        /**
         * @param techPrivacy Whether domain technical contact information is concealed from WHOIS queries. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder techPrivacy(@Nullable Output<Boolean> techPrivacy) {
            $.techPrivacy = techPrivacy;
            return this;
        }

        /**
         * @param techPrivacy Whether domain technical contact information is concealed from WHOIS queries. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder techPrivacy(Boolean techPrivacy) {
            return techPrivacy(Output.of(techPrivacy));
        }

        /**
         * @param transferLock Whether the domain is locked for transfer. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder transferLock(@Nullable Output<Boolean> transferLock) {
            $.transferLock = transferLock;
            return this;
        }

        /**
         * @param transferLock Whether the domain is locked for transfer. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder transferLock(Boolean transferLock) {
            return transferLock(Output.of(transferLock));
        }

        /**
         * @param updatedDate The last updated date of the domain as found in the response to a WHOIS query.
         * 
         * @return builder
         * 
         */
        public Builder updatedDate(@Nullable Output<String> updatedDate) {
            $.updatedDate = updatedDate;
            return this;
        }

        /**
         * @param updatedDate The last updated date of the domain as found in the response to a WHOIS query.
         * 
         * @return builder
         * 
         */
        public Builder updatedDate(String updatedDate) {
            return updatedDate(Output.of(updatedDate));
        }

        /**
         * @param whoisServer The fully qualified name of the WHOIS server that can answer the WHOIS query for the domain.
         * 
         * @return builder
         * 
         */
        public Builder whoisServer(@Nullable Output<String> whoisServer) {
            $.whoisServer = whoisServer;
            return this;
        }

        /**
         * @param whoisServer The fully qualified name of the WHOIS server that can answer the WHOIS query for the domain.
         * 
         * @return builder
         * 
         */
        public Builder whoisServer(String whoisServer) {
            return whoisServer(Output.of(whoisServer));
        }

        public RegisteredDomainState build() {
            return $;
        }
    }

}
