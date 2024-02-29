// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetIdentityPoolArgs, GetIdentityPoolResult, GetIdentityPoolOutputArgs } from "./getIdentityPool";
export const getIdentityPool: typeof import("./getIdentityPool").getIdentityPool = null as any;
export const getIdentityPoolOutput: typeof import("./getIdentityPool").getIdentityPoolOutput = null as any;
utilities.lazyLoad(exports, ["getIdentityPool","getIdentityPoolOutput"], () => require("./getIdentityPool"));

export { GetUserGroupArgs, GetUserGroupResult, GetUserGroupOutputArgs } from "./getUserGroup";
export const getUserGroup: typeof import("./getUserGroup").getUserGroup = null as any;
export const getUserGroupOutput: typeof import("./getUserGroup").getUserGroupOutput = null as any;
utilities.lazyLoad(exports, ["getUserGroup","getUserGroupOutput"], () => require("./getUserGroup"));

export { GetUserGroupsArgs, GetUserGroupsResult, GetUserGroupsOutputArgs } from "./getUserGroups";
export const getUserGroups: typeof import("./getUserGroups").getUserGroups = null as any;
export const getUserGroupsOutput: typeof import("./getUserGroups").getUserGroupsOutput = null as any;
utilities.lazyLoad(exports, ["getUserGroups","getUserGroupsOutput"], () => require("./getUserGroups"));

export { GetUserPoolClientArgs, GetUserPoolClientResult, GetUserPoolClientOutputArgs } from "./getUserPoolClient";
export const getUserPoolClient: typeof import("./getUserPoolClient").getUserPoolClient = null as any;
export const getUserPoolClientOutput: typeof import("./getUserPoolClient").getUserPoolClientOutput = null as any;
utilities.lazyLoad(exports, ["getUserPoolClient","getUserPoolClientOutput"], () => require("./getUserPoolClient"));

export { GetUserPoolClientsArgs, GetUserPoolClientsResult, GetUserPoolClientsOutputArgs } from "./getUserPoolClients";
export const getUserPoolClients: typeof import("./getUserPoolClients").getUserPoolClients = null as any;
export const getUserPoolClientsOutput: typeof import("./getUserPoolClients").getUserPoolClientsOutput = null as any;
utilities.lazyLoad(exports, ["getUserPoolClients","getUserPoolClientsOutput"], () => require("./getUserPoolClients"));

export { GetUserPoolSigningCertificateArgs, GetUserPoolSigningCertificateResult, GetUserPoolSigningCertificateOutputArgs } from "./getUserPoolSigningCertificate";
export const getUserPoolSigningCertificate: typeof import("./getUserPoolSigningCertificate").getUserPoolSigningCertificate = null as any;
export const getUserPoolSigningCertificateOutput: typeof import("./getUserPoolSigningCertificate").getUserPoolSigningCertificateOutput = null as any;
utilities.lazyLoad(exports, ["getUserPoolSigningCertificate","getUserPoolSigningCertificateOutput"], () => require("./getUserPoolSigningCertificate"));

export { GetUserPoolsArgs, GetUserPoolsResult, GetUserPoolsOutputArgs } from "./getUserPools";
export const getUserPools: typeof import("./getUserPools").getUserPools = null as any;
export const getUserPoolsOutput: typeof import("./getUserPools").getUserPoolsOutput = null as any;
utilities.lazyLoad(exports, ["getUserPools","getUserPoolsOutput"], () => require("./getUserPools"));

export { IdentityPoolArgs, IdentityPoolState } from "./identityPool";
export type IdentityPool = import("./identityPool").IdentityPool;
export const IdentityPool: typeof import("./identityPool").IdentityPool = null as any;
utilities.lazyLoad(exports, ["IdentityPool"], () => require("./identityPool"));

export { IdentityPoolProviderPrincipalTagArgs, IdentityPoolProviderPrincipalTagState } from "./identityPoolProviderPrincipalTag";
export type IdentityPoolProviderPrincipalTag = import("./identityPoolProviderPrincipalTag").IdentityPoolProviderPrincipalTag;
export const IdentityPoolProviderPrincipalTag: typeof import("./identityPoolProviderPrincipalTag").IdentityPoolProviderPrincipalTag = null as any;
utilities.lazyLoad(exports, ["IdentityPoolProviderPrincipalTag"], () => require("./identityPoolProviderPrincipalTag"));

export { IdentityPoolRoleAttachmentArgs, IdentityPoolRoleAttachmentState } from "./identityPoolRoleAttachment";
export type IdentityPoolRoleAttachment = import("./identityPoolRoleAttachment").IdentityPoolRoleAttachment;
export const IdentityPoolRoleAttachment: typeof import("./identityPoolRoleAttachment").IdentityPoolRoleAttachment = null as any;
utilities.lazyLoad(exports, ["IdentityPoolRoleAttachment"], () => require("./identityPoolRoleAttachment"));

export { IdentityProviderArgs, IdentityProviderState } from "./identityProvider";
export type IdentityProvider = import("./identityProvider").IdentityProvider;
export const IdentityProvider: typeof import("./identityProvider").IdentityProvider = null as any;
utilities.lazyLoad(exports, ["IdentityProvider"], () => require("./identityProvider"));

export { ManagedUserPoolClientArgs, ManagedUserPoolClientState } from "./managedUserPoolClient";
export type ManagedUserPoolClient = import("./managedUserPoolClient").ManagedUserPoolClient;
export const ManagedUserPoolClient: typeof import("./managedUserPoolClient").ManagedUserPoolClient = null as any;
utilities.lazyLoad(exports, ["ManagedUserPoolClient"], () => require("./managedUserPoolClient"));

export { ResourceServerArgs, ResourceServerState } from "./resourceServer";
export type ResourceServer = import("./resourceServer").ResourceServer;
export const ResourceServer: typeof import("./resourceServer").ResourceServer = null as any;
utilities.lazyLoad(exports, ["ResourceServer"], () => require("./resourceServer"));

export { RiskConfigurationArgs, RiskConfigurationState } from "./riskConfiguration";
export type RiskConfiguration = import("./riskConfiguration").RiskConfiguration;
export const RiskConfiguration: typeof import("./riskConfiguration").RiskConfiguration = null as any;
utilities.lazyLoad(exports, ["RiskConfiguration"], () => require("./riskConfiguration"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));

export { UserGroupArgs, UserGroupState } from "./userGroup";
export type UserGroup = import("./userGroup").UserGroup;
export const UserGroup: typeof import("./userGroup").UserGroup = null as any;
utilities.lazyLoad(exports, ["UserGroup"], () => require("./userGroup"));

export { UserInGroupArgs, UserInGroupState } from "./userInGroup";
export type UserInGroup = import("./userInGroup").UserInGroup;
export const UserInGroup: typeof import("./userInGroup").UserInGroup = null as any;
utilities.lazyLoad(exports, ["UserInGroup"], () => require("./userInGroup"));

export { UserPoolArgs, UserPoolState } from "./userPool";
export type UserPool = import("./userPool").UserPool;
export const UserPool: typeof import("./userPool").UserPool = null as any;
utilities.lazyLoad(exports, ["UserPool"], () => require("./userPool"));

export { UserPoolClientArgs, UserPoolClientState } from "./userPoolClient";
export type UserPoolClient = import("./userPoolClient").UserPoolClient;
export const UserPoolClient: typeof import("./userPoolClient").UserPoolClient = null as any;
utilities.lazyLoad(exports, ["UserPoolClient"], () => require("./userPoolClient"));

export { UserPoolDomainArgs, UserPoolDomainState } from "./userPoolDomain";
export type UserPoolDomain = import("./userPoolDomain").UserPoolDomain;
export const UserPoolDomain: typeof import("./userPoolDomain").UserPoolDomain = null as any;
utilities.lazyLoad(exports, ["UserPoolDomain"], () => require("./userPoolDomain"));

export { UserPoolUICustomizationArgs, UserPoolUICustomizationState } from "./userPoolUICustomization";
export type UserPoolUICustomization = import("./userPoolUICustomization").UserPoolUICustomization;
export const UserPoolUICustomization: typeof import("./userPoolUICustomization").UserPoolUICustomization = null as any;
utilities.lazyLoad(exports, ["UserPoolUICustomization"], () => require("./userPoolUICustomization"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:cognito/identityPool:IdentityPool":
                return new IdentityPool(name, <any>undefined, { urn })
            case "aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag":
                return new IdentityPoolProviderPrincipalTag(name, <any>undefined, { urn })
            case "aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment":
                return new IdentityPoolRoleAttachment(name, <any>undefined, { urn })
            case "aws:cognito/identityProvider:IdentityProvider":
                return new IdentityProvider(name, <any>undefined, { urn })
            case "aws:cognito/managedUserPoolClient:ManagedUserPoolClient":
                return new ManagedUserPoolClient(name, <any>undefined, { urn })
            case "aws:cognito/resourceServer:ResourceServer":
                return new ResourceServer(name, <any>undefined, { urn })
            case "aws:cognito/riskConfiguration:RiskConfiguration":
                return new RiskConfiguration(name, <any>undefined, { urn })
            case "aws:cognito/user:User":
                return new User(name, <any>undefined, { urn })
            case "aws:cognito/userGroup:UserGroup":
                return new UserGroup(name, <any>undefined, { urn })
            case "aws:cognito/userInGroup:UserInGroup":
                return new UserInGroup(name, <any>undefined, { urn })
            case "aws:cognito/userPool:UserPool":
                return new UserPool(name, <any>undefined, { urn })
            case "aws:cognito/userPoolClient:UserPoolClient":
                return new UserPoolClient(name, <any>undefined, { urn })
            case "aws:cognito/userPoolDomain:UserPoolDomain":
                return new UserPoolDomain(name, <any>undefined, { urn })
            case "aws:cognito/userPoolUICustomization:UserPoolUICustomization":
                return new UserPoolUICustomization(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "cognito/identityPool", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/identityPoolProviderPrincipalTag", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/identityPoolRoleAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/identityProvider", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/managedUserPoolClient", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/resourceServer", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/riskConfiguration", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/user", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/userGroup", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/userInGroup", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/userPool", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/userPoolClient", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/userPoolDomain", _module)
pulumi.runtime.registerResourceModule("aws", "cognito/userPoolUICustomization", _module)
