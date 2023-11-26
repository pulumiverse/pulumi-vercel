"""A Python Pulumi program"""

import pulumi
import pulumiverse_vercel as vercel

project = vercel.Project("pulumi-vercel", root_directory="src")
deployment = vercel.Deployment("production-deployment", project_id=project.id, production=True, files=vercel.get_project_directory_output(path="./src").files)

pulumi.export("url", deployment.url)
