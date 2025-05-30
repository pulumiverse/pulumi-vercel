// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about an existing Microfrontend Group.
 *
 * A Microfrontend Group is a definition of a microfrontend belonging to a Vercel Team.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumi/vercel";
 *
 * const example = vercel.getMicrofrontendGroup({
 *     id: "mfe_xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 * });
 * ```
 */
export function getMicrofrontendGroup(args: GetMicrofrontendGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetMicrofrontendGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getMicrofrontendGroup:getMicrofrontendGroup", {
        "id": args.id,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMicrofrontendGroup.
 */
export interface GetMicrofrontendGroupArgs {
    /**
     * A unique identifier for the group of microfrontends. Example: mfe_12HKQaOmR5t5Uy6vdcQsNIiZgHGB
     */
    id: string;
    /**
     * The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: string;
}

/**
 * A collection of values returned by getMicrofrontendGroup.
 */
export interface GetMicrofrontendGroupResult {
    /**
     * The default app for the project. Used as the entry point for the microfrontend.
     */
    readonly defaultApp: outputs.GetMicrofrontendGroupDefaultApp;
    /**
     * A unique identifier for the group of microfrontends. Example: mfe_12HKQaOmR5t5Uy6vdcQsNIiZgHGB
     */
    readonly id: string;
    /**
     * A human readable name for the microfrontends group.
     */
    readonly name: string;
    /**
     * A slugified version of the name.
     */
    readonly slug: string;
    /**
     * The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
     */
    readonly teamId: string;
}
/**
 * Provides information about an existing Microfrontend Group.
 *
 * A Microfrontend Group is a definition of a microfrontend belonging to a Vercel Team.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumi/vercel";
 *
 * const example = vercel.getMicrofrontendGroup({
 *     id: "mfe_xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 * });
 * ```
 */
export function getMicrofrontendGroupOutput(args: GetMicrofrontendGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMicrofrontendGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vercel:index/getMicrofrontendGroup:getMicrofrontendGroup", {
        "id": args.id,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMicrofrontendGroup.
 */
export interface GetMicrofrontendGroupOutputArgs {
    /**
     * A unique identifier for the group of microfrontends. Example: mfe_12HKQaOmR5t5Uy6vdcQsNIiZgHGB
     */
    id: pulumi.Input<string>;
    /**
     * The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
}
