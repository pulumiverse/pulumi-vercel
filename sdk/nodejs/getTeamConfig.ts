// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves the configuration of an existing Vercel Team.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumi/vercel";
 *
 * const example = vercel.getTeamConfig({
 *     id: "team_xxxxxxxxxxxxxxxxxxxxxxxx",
 * });
 * ```
 */
export function getTeamConfig(args: GetTeamConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getTeamConfig:getTeamConfig", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeamConfig.
 */
export interface GetTeamConfigArgs {
    /**
     * The ID of the existing Vercel Team.
     */
    id: string;
}

/**
 * A collection of values returned by getTeamConfig.
 */
export interface GetTeamConfigResult {
    /**
     * A description of the team.
     */
    readonly description: string;
    /**
     * Hostname that'll be matched with emails on sign-up to automatically join the Team.
     */
    readonly emailDomain: string;
    /**
     * Preview feedback configuration.
     */
    readonly enablePreviewFeedback: string;
    /**
     * Production feedback configuration.
     */
    readonly enableProductionFeedback: string;
    /**
     * Indicates if ip addresses should be accessible in o11y tooling.
     */
    readonly hideIpAddresses: boolean;
    /**
     * Indicates if ip addresses should be accessible in log drains.
     */
    readonly hideIpAddressesInLogDrains: boolean;
    /**
     * The ID of the existing Vercel Team.
     */
    readonly id: string;
    /**
     * A code that can be used to join this team. Only visible to Team owners.
     */
    readonly inviteCode: string;
    /**
     * The name of the team.
     */
    readonly name: string;
    /**
     * The hostname that is used as the preview deployment suffix.
     */
    readonly previewDeploymentSuffix: string;
    /**
     * Configuration for Remote Caching.
     */
    readonly remoteCaching: outputs.GetTeamConfigRemoteCaching;
    /**
     * Configuration for SAML authentication.
     */
    readonly saml: outputs.GetTeamConfigSaml;
    /**
     * The policy for sensitive environment variables.
     */
    readonly sensitiveEnvironmentVariablePolicy: string;
    /**
     * The slug of the team. Used in the URL of the team's dashboard.
     */
    readonly slug: string;
}
/**
 * Retrieves the configuration of an existing Vercel Team.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumi/vercel";
 *
 * const example = vercel.getTeamConfig({
 *     id: "team_xxxxxxxxxxxxxxxxxxxxxxxx",
 * });
 * ```
 */
export function getTeamConfigOutput(args: GetTeamConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTeamConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vercel:index/getTeamConfig:getTeamConfig", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeamConfig.
 */
export interface GetTeamConfigOutputArgs {
    /**
     * The ID of the existing Vercel Team.
     */
    id: pulumi.Input<string>;
}
