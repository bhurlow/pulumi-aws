// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApiArgs, ApiState } from "./api";
export type Api = import("./api").Api;
export const Api: typeof import("./api").Api = null as any;
utilities.lazyLoad(exports, ["Api"], () => require("./api"));

export { ApiMappingArgs, ApiMappingState } from "./apiMapping";
export type ApiMapping = import("./apiMapping").ApiMapping;
export const ApiMapping: typeof import("./apiMapping").ApiMapping = null as any;
utilities.lazyLoad(exports, ["ApiMapping"], () => require("./apiMapping"));

export { AuthorizerArgs, AuthorizerState } from "./authorizer";
export type Authorizer = import("./authorizer").Authorizer;
export const Authorizer: typeof import("./authorizer").Authorizer = null as any;
utilities.lazyLoad(exports, ["Authorizer"], () => require("./authorizer"));

export { DeploymentArgs, DeploymentState } from "./deployment";
export type Deployment = import("./deployment").Deployment;
export const Deployment: typeof import("./deployment").Deployment = null as any;
utilities.lazyLoad(exports, ["Deployment"], () => require("./deployment"));

export { DomainNameArgs, DomainNameState } from "./domainName";
export type DomainName = import("./domainName").DomainName;
export const DomainName: typeof import("./domainName").DomainName = null as any;
utilities.lazyLoad(exports, ["DomainName"], () => require("./domainName"));

export { GetApiArgs, GetApiResult, GetApiOutputArgs } from "./getApi";
export const getApi: typeof import("./getApi").getApi = null as any;
export const getApiOutput: typeof import("./getApi").getApiOutput = null as any;
utilities.lazyLoad(exports, ["getApi","getApiOutput"], () => require("./getApi"));

export { GetApisArgs, GetApisResult, GetApisOutputArgs } from "./getApis";
export const getApis: typeof import("./getApis").getApis = null as any;
export const getApisOutput: typeof import("./getApis").getApisOutput = null as any;
utilities.lazyLoad(exports, ["getApis","getApisOutput"], () => require("./getApis"));

export { GetExportArgs, GetExportResult, GetExportOutputArgs } from "./getExport";
export const getExport: typeof import("./getExport").getExport = null as any;
export const getExportOutput: typeof import("./getExport").getExportOutput = null as any;
utilities.lazyLoad(exports, ["getExport","getExportOutput"], () => require("./getExport"));

export { GetVpcLinkArgs, GetVpcLinkResult, GetVpcLinkOutputArgs } from "./getVpcLink";
export const getVpcLink: typeof import("./getVpcLink").getVpcLink = null as any;
export const getVpcLinkOutput: typeof import("./getVpcLink").getVpcLinkOutput = null as any;
utilities.lazyLoad(exports, ["getVpcLink","getVpcLinkOutput"], () => require("./getVpcLink"));

export { IntegrationArgs, IntegrationState } from "./integration";
export type Integration = import("./integration").Integration;
export const Integration: typeof import("./integration").Integration = null as any;
utilities.lazyLoad(exports, ["Integration"], () => require("./integration"));

export { IntegrationResponseArgs, IntegrationResponseState } from "./integrationResponse";
export type IntegrationResponse = import("./integrationResponse").IntegrationResponse;
export const IntegrationResponse: typeof import("./integrationResponse").IntegrationResponse = null as any;
utilities.lazyLoad(exports, ["IntegrationResponse"], () => require("./integrationResponse"));

export { ModelArgs, ModelState } from "./model";
export type Model = import("./model").Model;
export const Model: typeof import("./model").Model = null as any;
utilities.lazyLoad(exports, ["Model"], () => require("./model"));

export { RouteArgs, RouteState } from "./route";
export type Route = import("./route").Route;
export const Route: typeof import("./route").Route = null as any;
utilities.lazyLoad(exports, ["Route"], () => require("./route"));

export { RouteResponseArgs, RouteResponseState } from "./routeResponse";
export type RouteResponse = import("./routeResponse").RouteResponse;
export const RouteResponse: typeof import("./routeResponse").RouteResponse = null as any;
utilities.lazyLoad(exports, ["RouteResponse"], () => require("./routeResponse"));

export { StageArgs, StageState } from "./stage";
export type Stage = import("./stage").Stage;
export const Stage: typeof import("./stage").Stage = null as any;
utilities.lazyLoad(exports, ["Stage"], () => require("./stage"));

export { VpcLinkArgs, VpcLinkState } from "./vpcLink";
export type VpcLink = import("./vpcLink").VpcLink;
export const VpcLink: typeof import("./vpcLink").VpcLink = null as any;
utilities.lazyLoad(exports, ["VpcLink"], () => require("./vpcLink"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:apigatewayv2/api:Api":
                return new Api(name, <any>undefined, { urn })
            case "aws:apigatewayv2/apiMapping:ApiMapping":
                return new ApiMapping(name, <any>undefined, { urn })
            case "aws:apigatewayv2/authorizer:Authorizer":
                return new Authorizer(name, <any>undefined, { urn })
            case "aws:apigatewayv2/deployment:Deployment":
                return new Deployment(name, <any>undefined, { urn })
            case "aws:apigatewayv2/domainName:DomainName":
                return new DomainName(name, <any>undefined, { urn })
            case "aws:apigatewayv2/integration:Integration":
                return new Integration(name, <any>undefined, { urn })
            case "aws:apigatewayv2/integrationResponse:IntegrationResponse":
                return new IntegrationResponse(name, <any>undefined, { urn })
            case "aws:apigatewayv2/model:Model":
                return new Model(name, <any>undefined, { urn })
            case "aws:apigatewayv2/route:Route":
                return new Route(name, <any>undefined, { urn })
            case "aws:apigatewayv2/routeResponse:RouteResponse":
                return new RouteResponse(name, <any>undefined, { urn })
            case "aws:apigatewayv2/stage:Stage":
                return new Stage(name, <any>undefined, { urn })
            case "aws:apigatewayv2/vpcLink:VpcLink":
                return new VpcLink(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/api", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/apiMapping", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/authorizer", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/deployment", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/domainName", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/integration", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/integrationResponse", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/model", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/route", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/routeResponse", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/stage", _module)
pulumi.runtime.registerResourceModule("aws", "apigatewayv2/vpcLink", _module)
