# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectDeploymentRetentionArgs', 'ProjectDeploymentRetention']

@pulumi.input_type
class ProjectDeploymentRetentionArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 expiration_canceled: Optional[pulumi.Input[str]] = None,
                 expiration_errored: Optional[pulumi.Input[str]] = None,
                 expiration_preview: Optional[pulumi.Input[str]] = None,
                 expiration_production: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectDeploymentRetention resource.
        :param pulumi.Input[str] project_id: The ID of the Project for the retention policy
        :param pulumi.Input[str] expiration_canceled: The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_errored: The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_preview: The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_production: The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] team_id: The ID of the Vercel team.
        """
        pulumi.set(__self__, "project_id", project_id)
        if expiration_canceled is not None:
            pulumi.set(__self__, "expiration_canceled", expiration_canceled)
        if expiration_errored is not None:
            pulumi.set(__self__, "expiration_errored", expiration_errored)
        if expiration_preview is not None:
            pulumi.set(__self__, "expiration_preview", expiration_preview)
        if expiration_production is not None:
            pulumi.set(__self__, "expiration_production", expiration_production)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the Project for the retention policy
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="expirationCanceled")
    def expiration_canceled(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_canceled")

    @expiration_canceled.setter
    def expiration_canceled(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_canceled", value)

    @property
    @pulumi.getter(name="expirationErrored")
    def expiration_errored(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_errored")

    @expiration_errored.setter
    def expiration_errored(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_errored", value)

    @property
    @pulumi.getter(name="expirationPreview")
    def expiration_preview(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_preview")

    @expiration_preview.setter
    def expiration_preview(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_preview", value)

    @property
    @pulumi.getter(name="expirationProduction")
    def expiration_production(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_production")

    @expiration_production.setter
    def expiration_production(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_production", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Vercel team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


@pulumi.input_type
class _ProjectDeploymentRetentionState:
    def __init__(__self__, *,
                 expiration_canceled: Optional[pulumi.Input[str]] = None,
                 expiration_errored: Optional[pulumi.Input[str]] = None,
                 expiration_preview: Optional[pulumi.Input[str]] = None,
                 expiration_production: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectDeploymentRetention resources.
        :param pulumi.Input[str] expiration_canceled: The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_errored: The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_preview: The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_production: The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] project_id: The ID of the Project for the retention policy
        :param pulumi.Input[str] team_id: The ID of the Vercel team.
        """
        if expiration_canceled is not None:
            pulumi.set(__self__, "expiration_canceled", expiration_canceled)
        if expiration_errored is not None:
            pulumi.set(__self__, "expiration_errored", expiration_errored)
        if expiration_preview is not None:
            pulumi.set(__self__, "expiration_preview", expiration_preview)
        if expiration_production is not None:
            pulumi.set(__self__, "expiration_production", expiration_production)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter(name="expirationCanceled")
    def expiration_canceled(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_canceled")

    @expiration_canceled.setter
    def expiration_canceled(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_canceled", value)

    @property
    @pulumi.getter(name="expirationErrored")
    def expiration_errored(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_errored")

    @expiration_errored.setter
    def expiration_errored(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_errored", value)

    @property
    @pulumi.getter(name="expirationPreview")
    def expiration_preview(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_preview")

    @expiration_preview.setter
    def expiration_preview(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_preview", value)

    @property
    @pulumi.getter(name="expirationProduction")
    def expiration_production(self) -> Optional[pulumi.Input[str]]:
        """
        The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_production")

    @expiration_production.setter
    def expiration_production(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_production", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Project for the retention policy
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Vercel team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


class ProjectDeploymentRetention(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_canceled: Optional[pulumi.Input[str]] = None,
                 expiration_errored: Optional[pulumi.Input[str]] = None,
                 expiration_preview: Optional[pulumi.Input[str]] = None,
                 expiration_production: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Project Deployment Retention resource.

        A Project Deployment Retention resource defines an Deployment Retention on a Vercel Project.

        For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/security/deployment-retention).

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example", git_repository={
            "type": "github",
            "repo": "vercel/some-repo",
        })
        # An unlimited deployment retention policy that will be created
        # for this project for all deployments.
        example_unlimited = vercel.ProjectDeploymentRetention("exampleUnlimited",
            project_id=example.id,
            team_id=example.team_id,
            expiration_preview="unlimited",
            expiration_production="unlimited",
            expiration_canceled="unlimited",
            expiration_errored="unlimited")
        # A customized deployment retention policy that will be created
        # for this project for all deployments.
        example_customized = vercel.ProjectDeploymentRetention("exampleCustomized",
            project_id=example.id,
            team_id=example.team_id,
            expiration_preview="3m",
            expiration_production="1y",
            expiration_canceled="1m",
            expiration_errored="2m")
        ```

        ## Import

        You can import via the team_id and project_id.

        - team_id can be found in the team `settings` tab in the Vercel UI.

        - project_id can be found in the project `settings` tab in the Vercel UI.

        ```sh
        $ pulumi import vercel:index/projectDeploymentRetention:ProjectDeploymentRetention example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_canceled: The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_errored: The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_preview: The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_production: The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] project_id: The ID of the Project for the retention policy
        :param pulumi.Input[str] team_id: The ID of the Vercel team.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectDeploymentRetentionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Project Deployment Retention resource.

        A Project Deployment Retention resource defines an Deployment Retention on a Vercel Project.

        For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/security/deployment-retention).

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example", git_repository={
            "type": "github",
            "repo": "vercel/some-repo",
        })
        # An unlimited deployment retention policy that will be created
        # for this project for all deployments.
        example_unlimited = vercel.ProjectDeploymentRetention("exampleUnlimited",
            project_id=example.id,
            team_id=example.team_id,
            expiration_preview="unlimited",
            expiration_production="unlimited",
            expiration_canceled="unlimited",
            expiration_errored="unlimited")
        # A customized deployment retention policy that will be created
        # for this project for all deployments.
        example_customized = vercel.ProjectDeploymentRetention("exampleCustomized",
            project_id=example.id,
            team_id=example.team_id,
            expiration_preview="3m",
            expiration_production="1y",
            expiration_canceled="1m",
            expiration_errored="2m")
        ```

        ## Import

        You can import via the team_id and project_id.

        - team_id can be found in the team `settings` tab in the Vercel UI.

        - project_id can be found in the project `settings` tab in the Vercel UI.

        ```sh
        $ pulumi import vercel:index/projectDeploymentRetention:ProjectDeploymentRetention example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param ProjectDeploymentRetentionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectDeploymentRetentionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_canceled: Optional[pulumi.Input[str]] = None,
                 expiration_errored: Optional[pulumi.Input[str]] = None,
                 expiration_preview: Optional[pulumi.Input[str]] = None,
                 expiration_production: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectDeploymentRetentionArgs.__new__(ProjectDeploymentRetentionArgs)

            __props__.__dict__["expiration_canceled"] = expiration_canceled
            __props__.__dict__["expiration_errored"] = expiration_errored
            __props__.__dict__["expiration_preview"] = expiration_preview
            __props__.__dict__["expiration_production"] = expiration_production
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["team_id"] = team_id
        super(ProjectDeploymentRetention, __self__).__init__(
            'vercel:index/projectDeploymentRetention:ProjectDeploymentRetention',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expiration_canceled: Optional[pulumi.Input[str]] = None,
            expiration_errored: Optional[pulumi.Input[str]] = None,
            expiration_preview: Optional[pulumi.Input[str]] = None,
            expiration_production: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None) -> 'ProjectDeploymentRetention':
        """
        Get an existing ProjectDeploymentRetention resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_canceled: The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_errored: The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_preview: The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] expiration_production: The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        :param pulumi.Input[str] project_id: The ID of the Project for the retention policy
        :param pulumi.Input[str] team_id: The ID of the Vercel team.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectDeploymentRetentionState.__new__(_ProjectDeploymentRetentionState)

        __props__.__dict__["expiration_canceled"] = expiration_canceled
        __props__.__dict__["expiration_errored"] = expiration_errored
        __props__.__dict__["expiration_preview"] = expiration_preview
        __props__.__dict__["expiration_production"] = expiration_production
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["team_id"] = team_id
        return ProjectDeploymentRetention(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="expirationCanceled")
    def expiration_canceled(self) -> pulumi.Output[str]:
        """
        The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_canceled")

    @property
    @pulumi.getter(name="expirationErrored")
    def expiration_errored(self) -> pulumi.Output[str]:
        """
        The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_errored")

    @property
    @pulumi.getter(name="expirationPreview")
    def expiration_preview(self) -> pulumi.Output[str]:
        """
        The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_preview")

    @property
    @pulumi.getter(name="expirationProduction")
    def expiration_production(self) -> pulumi.Output[str]:
        """
        The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
        """
        return pulumi.get(self, "expiration_production")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the Project for the retention policy
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The ID of the Vercel team.
        """
        return pulumi.get(self, "team_id")
