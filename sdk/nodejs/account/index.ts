// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AlternativeContactArgs, AlternativeContactState } from "./alternativeContact";
export type AlternativeContact = import("./alternativeContact").AlternativeContact;
export const AlternativeContact: typeof import("./alternativeContact").AlternativeContact = null as any;
utilities.lazyLoad(exports, ["AlternativeContact"], () => require("./alternativeContact"));

export { PrimaryContactArgs, PrimaryContactState } from "./primaryContact";
export type PrimaryContact = import("./primaryContact").PrimaryContact;
export const PrimaryContact: typeof import("./primaryContact").PrimaryContact = null as any;
utilities.lazyLoad(exports, ["PrimaryContact"], () => require("./primaryContact"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:account/alternativeContact:AlternativeContact":
                return new AlternativeContact(name, <any>undefined, { urn })
            case "aws:account/primaryContact:PrimaryContact":
                return new PrimaryContact(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "account/alternativeContact", _module)
pulumi.runtime.registerResourceModule("aws", "account/primaryContact", _module)
