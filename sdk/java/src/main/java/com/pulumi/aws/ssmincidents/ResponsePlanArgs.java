// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ssmincidents;

import com.pulumi.aws.ssmincidents.inputs.ResponsePlanActionArgs;
import com.pulumi.aws.ssmincidents.inputs.ResponsePlanIncidentTemplateArgs;
import com.pulumi.aws.ssmincidents.inputs.ResponsePlanIntegrationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResponsePlanArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResponsePlanArgs Empty = new ResponsePlanArgs();

    /**
     * The actions that the response plan starts at the beginning of an incident.
     * 
     */
    @Import(name="action")
    private @Nullable Output<ResponsePlanActionArgs> action;

    /**
     * @return The actions that the response plan starts at the beginning of an incident.
     * 
     */
    public Optional<Output<ResponsePlanActionArgs>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * The Chatbot chat channel used for collaboration during an incident.
     * 
     */
    @Import(name="chatChannels")
    private @Nullable Output<List<String>> chatChannels;

    /**
     * @return The Chatbot chat channel used for collaboration during an incident.
     * 
     */
    public Optional<Output<List<String>>> chatChannels() {
        return Optional.ofNullable(this.chatChannels);
    }

    /**
     * The long format of the response plan name. This field can contain spaces.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The long format of the response plan name. This field can contain spaces.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
     * 
     */
    @Import(name="engagements")
    private @Nullable Output<List<String>> engagements;

    /**
     * @return The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
     * 
     */
    public Optional<Output<List<String>>> engagements() {
        return Optional.ofNullable(this.engagements);
    }

    @Import(name="incidentTemplate", required=true)
    private Output<ResponsePlanIncidentTemplateArgs> incidentTemplate;

    public Output<ResponsePlanIncidentTemplateArgs> incidentTemplate() {
        return this.incidentTemplate;
    }

    /**
     * Information about third-party services integrated into the response plan. The following values are supported:
     * 
     */
    @Import(name="integration")
    private @Nullable Output<ResponsePlanIntegrationArgs> integration;

    /**
     * @return Information about third-party services integrated into the response plan. The following values are supported:
     * 
     */
    public Optional<Output<ResponsePlanIntegrationArgs>> integration() {
        return Optional.ofNullable(this.integration);
    }

    /**
     * The name of the response plan.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the response plan.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The tags applied to the response plan.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tags applied to the response plan.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ResponsePlanArgs() {}

    private ResponsePlanArgs(ResponsePlanArgs $) {
        this.action = $.action;
        this.chatChannels = $.chatChannels;
        this.displayName = $.displayName;
        this.engagements = $.engagements;
        this.incidentTemplate = $.incidentTemplate;
        this.integration = $.integration;
        this.name = $.name;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResponsePlanArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResponsePlanArgs $;

        public Builder() {
            $ = new ResponsePlanArgs();
        }

        public Builder(ResponsePlanArgs defaults) {
            $ = new ResponsePlanArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The actions that the response plan starts at the beginning of an incident.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<ResponsePlanActionArgs> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The actions that the response plan starts at the beginning of an incident.
         * 
         * @return builder
         * 
         */
        public Builder action(ResponsePlanActionArgs action) {
            return action(Output.of(action));
        }

        /**
         * @param chatChannels The Chatbot chat channel used for collaboration during an incident.
         * 
         * @return builder
         * 
         */
        public Builder chatChannels(@Nullable Output<List<String>> chatChannels) {
            $.chatChannels = chatChannels;
            return this;
        }

        /**
         * @param chatChannels The Chatbot chat channel used for collaboration during an incident.
         * 
         * @return builder
         * 
         */
        public Builder chatChannels(List<String> chatChannels) {
            return chatChannels(Output.of(chatChannels));
        }

        /**
         * @param chatChannels The Chatbot chat channel used for collaboration during an incident.
         * 
         * @return builder
         * 
         */
        public Builder chatChannels(String... chatChannels) {
            return chatChannels(List.of(chatChannels));
        }

        /**
         * @param displayName The long format of the response plan name. This field can contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The long format of the response plan name. This field can contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param engagements The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
         * 
         * @return builder
         * 
         */
        public Builder engagements(@Nullable Output<List<String>> engagements) {
            $.engagements = engagements;
            return this;
        }

        /**
         * @param engagements The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
         * 
         * @return builder
         * 
         */
        public Builder engagements(List<String> engagements) {
            return engagements(Output.of(engagements));
        }

        /**
         * @param engagements The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
         * 
         * @return builder
         * 
         */
        public Builder engagements(String... engagements) {
            return engagements(List.of(engagements));
        }

        public Builder incidentTemplate(Output<ResponsePlanIncidentTemplateArgs> incidentTemplate) {
            $.incidentTemplate = incidentTemplate;
            return this;
        }

        public Builder incidentTemplate(ResponsePlanIncidentTemplateArgs incidentTemplate) {
            return incidentTemplate(Output.of(incidentTemplate));
        }

        /**
         * @param integration Information about third-party services integrated into the response plan. The following values are supported:
         * 
         * @return builder
         * 
         */
        public Builder integration(@Nullable Output<ResponsePlanIntegrationArgs> integration) {
            $.integration = integration;
            return this;
        }

        /**
         * @param integration Information about third-party services integrated into the response plan. The following values are supported:
         * 
         * @return builder
         * 
         */
        public Builder integration(ResponsePlanIntegrationArgs integration) {
            return integration(Output.of(integration));
        }

        /**
         * @param name The name of the response plan.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the response plan.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags The tags applied to the response plan.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags applied to the response plan.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public ResponsePlanArgs build() {
            $.incidentTemplate = Objects.requireNonNull($.incidentTemplate, "expected parameter 'incidentTemplate' to be non-null");
            return $;
        }
    }

}
