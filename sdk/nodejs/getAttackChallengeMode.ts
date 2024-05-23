// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Attack Challenge Mode resource.
 *
 * Attack Challenge Mode prevent malicious traffic by showing a verification challenge for every visitor.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumi/vercel";
 *
 * const example = vercel.getAttackChallengeMode({
 *     projectId: vercel_project.example.id,
 * });
 * ```
 */
export function getAttackChallengeMode(args: GetAttackChallengeModeArgs, opts?: pulumi.InvokeOptions): Promise<GetAttackChallengeModeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vercel:index/getAttackChallengeMode:getAttackChallengeMode", {
        "projectId": args.projectId,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAttackChallengeMode.
 */
export interface GetAttackChallengeModeArgs {
    /**
     * The ID of the Project to adjust the CPU for.
     */
    projectId: string;
    /**
     * The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: string;
}

/**
 * A collection of values returned by getAttackChallengeMode.
 */
export interface GetAttackChallengeModeResult {
    /**
     * Whether Attack Challenge Mode is enabled or not.
     */
    readonly enabled: boolean;
    /**
     * The resource identifier.
     */
    readonly id: string;
    /**
     * The ID of the Project to adjust the CPU for.
     */
    readonly projectId: string;
    /**
     * The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    readonly teamId: string;
}
/**
 * Provides an Attack Challenge Mode resource.
 *
 * Attack Challenge Mode prevent malicious traffic by showing a verification challenge for every visitor.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumi/vercel";
 *
 * const example = vercel.getAttackChallengeMode({
 *     projectId: vercel_project.example.id,
 * });
 * ```
 */
export function getAttackChallengeModeOutput(args: GetAttackChallengeModeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAttackChallengeModeResult> {
    return pulumi.output(args).apply((a: any) => getAttackChallengeMode(a, opts))
}

/**
 * A collection of arguments for invoking getAttackChallengeMode.
 */
export interface GetAttackChallengeModeOutputArgs {
    /**
     * The ID of the Project to adjust the CPU for.
     */
    projectId: pulumi.Input<string>;
    /**
     * The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
}