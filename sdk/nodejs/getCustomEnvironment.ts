// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getCustomEnvironment(args: GetCustomEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomEnvironmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getCustomEnvironment:getCustomEnvironment", {
        "name": args.name,
        "projectId": args.projectId,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustomEnvironment.
 */
export interface GetCustomEnvironmentArgs {
    name: string;
    projectId: string;
    teamId?: string;
}

/**
 * A collection of values returned by getCustomEnvironment.
 */
export interface GetCustomEnvironmentResult {
    readonly branchTracking: outputs.GetCustomEnvironmentBranchTracking;
    readonly description: string;
    readonly id: string;
    readonly name: string;
    readonly projectId: string;
    readonly teamId: string;
}
export function getCustomEnvironmentOutput(args: GetCustomEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomEnvironmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vercel:index/getCustomEnvironment:getCustomEnvironment", {
        "name": args.name,
        "projectId": args.projectId,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustomEnvironment.
 */
export interface GetCustomEnvironmentOutputArgs {
    name: pulumi.Input<string>;
    projectId: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
}
