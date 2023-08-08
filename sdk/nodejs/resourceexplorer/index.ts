// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { IndexArgs, IndexState } from "./index_";
export type Index = import("./index_").Index;
export const Index: typeof import("./index_").Index = null as any;
utilities.lazyLoad(exports, ["Index"], () => require("./index_"));

export { ViewArgs, ViewState } from "./view";
export type View = import("./view").View;
export const View: typeof import("./view").View = null as any;
utilities.lazyLoad(exports, ["View"], () => require("./view"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:resourceexplorer/index:Index":
                return new Index(name, <any>undefined, { urn })
            case "aws:resourceexplorer/view:View":
                return new View(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "resourceexplorer/index", _module)
pulumi.runtime.registerResourceModule("aws", "resourceexplorer/view", _module)
