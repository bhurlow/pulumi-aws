// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.budgets.outputs;

import com.pulumi.aws.budgets.outputs.BudgetActionDefinitionIamActionDefinition;
import com.pulumi.aws.budgets.outputs.BudgetActionDefinitionScpActionDefinition;
import com.pulumi.aws.budgets.outputs.BudgetActionDefinitionSsmActionDefinition;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BudgetActionDefinition {
    /**
     * @return The AWS Identity and Access Management (IAM) action definition details. See IAM Action Definition.
     * 
     */
    private final @Nullable BudgetActionDefinitionIamActionDefinition iamActionDefinition;
    /**
     * @return The service control policies (SCPs) action definition details. See SCP Action Definition.
     * 
     */
    private final @Nullable BudgetActionDefinitionScpActionDefinition scpActionDefinition;
    /**
     * @return The AWS Systems Manager (SSM) action definition details. See SSM Action Definition.
     * 
     */
    private final @Nullable BudgetActionDefinitionSsmActionDefinition ssmActionDefinition;

    @CustomType.Constructor
    private BudgetActionDefinition(
        @CustomType.Parameter("iamActionDefinition") @Nullable BudgetActionDefinitionIamActionDefinition iamActionDefinition,
        @CustomType.Parameter("scpActionDefinition") @Nullable BudgetActionDefinitionScpActionDefinition scpActionDefinition,
        @CustomType.Parameter("ssmActionDefinition") @Nullable BudgetActionDefinitionSsmActionDefinition ssmActionDefinition) {
        this.iamActionDefinition = iamActionDefinition;
        this.scpActionDefinition = scpActionDefinition;
        this.ssmActionDefinition = ssmActionDefinition;
    }

    /**
     * @return The AWS Identity and Access Management (IAM) action definition details. See IAM Action Definition.
     * 
     */
    public Optional<BudgetActionDefinitionIamActionDefinition> iamActionDefinition() {
        return Optional.ofNullable(this.iamActionDefinition);
    }
    /**
     * @return The service control policies (SCPs) action definition details. See SCP Action Definition.
     * 
     */
    public Optional<BudgetActionDefinitionScpActionDefinition> scpActionDefinition() {
        return Optional.ofNullable(this.scpActionDefinition);
    }
    /**
     * @return The AWS Systems Manager (SSM) action definition details. See SSM Action Definition.
     * 
     */
    public Optional<BudgetActionDefinitionSsmActionDefinition> ssmActionDefinition() {
        return Optional.ofNullable(this.ssmActionDefinition);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BudgetActionDefinition defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable BudgetActionDefinitionIamActionDefinition iamActionDefinition;
        private @Nullable BudgetActionDefinitionScpActionDefinition scpActionDefinition;
        private @Nullable BudgetActionDefinitionSsmActionDefinition ssmActionDefinition;

        public Builder() {
    	      // Empty
        }

        public Builder(BudgetActionDefinition defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.iamActionDefinition = defaults.iamActionDefinition;
    	      this.scpActionDefinition = defaults.scpActionDefinition;
    	      this.ssmActionDefinition = defaults.ssmActionDefinition;
        }

        public Builder iamActionDefinition(@Nullable BudgetActionDefinitionIamActionDefinition iamActionDefinition) {
            this.iamActionDefinition = iamActionDefinition;
            return this;
        }
        public Builder scpActionDefinition(@Nullable BudgetActionDefinitionScpActionDefinition scpActionDefinition) {
            this.scpActionDefinition = scpActionDefinition;
            return this;
        }
        public Builder ssmActionDefinition(@Nullable BudgetActionDefinitionSsmActionDefinition ssmActionDefinition) {
            this.ssmActionDefinition = ssmActionDefinition;
            return this;
        }        public BudgetActionDefinition build() {
            return new BudgetActionDefinition(iamActionDefinition, scpActionDefinition, ssmActionDefinition);
        }
    }
}