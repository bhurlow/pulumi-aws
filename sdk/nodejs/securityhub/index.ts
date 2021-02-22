// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./actionTarget";
export * from "./inviteAccepter";
export * from "./member";
export * from "./organizationAdminAccount";
export * from "./productSubscription";
export * from "./standardsSubscription";

// Import resources to register:
import { Account } from "./account";
import { ActionTarget } from "./actionTarget";
import { InviteAccepter } from "./inviteAccepter";
import { Member } from "./member";
import { OrganizationAdminAccount } from "./organizationAdminAccount";
import { ProductSubscription } from "./productSubscription";
import { StandardsSubscription } from "./standardsSubscription";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:securityhub/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "aws:securityhub/actionTarget:ActionTarget":
                return new ActionTarget(name, <any>undefined, { urn })
            case "aws:securityhub/inviteAccepter:InviteAccepter":
                return new InviteAccepter(name, <any>undefined, { urn })
            case "aws:securityhub/member:Member":
                return new Member(name, <any>undefined, { urn })
            case "aws:securityhub/organizationAdminAccount:OrganizationAdminAccount":
                return new OrganizationAdminAccount(name, <any>undefined, { urn })
            case "aws:securityhub/productSubscription:ProductSubscription":
                return new ProductSubscription(name, <any>undefined, { urn })
            case "aws:securityhub/standardsSubscription:StandardsSubscription":
                return new StandardsSubscription(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "securityhub/account", _module)
pulumi.runtime.registerResourceModule("aws", "securityhub/actionTarget", _module)
pulumi.runtime.registerResourceModule("aws", "securityhub/inviteAccepter", _module)
pulumi.runtime.registerResourceModule("aws", "securityhub/member", _module)
pulumi.runtime.registerResourceModule("aws", "securityhub/organizationAdminAccount", _module)
pulumi.runtime.registerResourceModule("aws", "securityhub/productSubscription", _module)
pulumi.runtime.registerResourceModule("aws", "securityhub/standardsSubscription", _module)
