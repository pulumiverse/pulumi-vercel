FROM mcr.microsoft.com/devcontainers/base:ubuntu@sha256:edfb983aab9c579a385dc23c57d7d3703f5ec920124d99c16204a2cac465aab4

# Install mise
RUN curl https://mise.run | sh
ENV PATH="/root/.local/bin:${PATH}"

# Setup mise
WORKDIR /code
RUN mise install

# Activate mise in shell
RUN echo 'eval "$(mise activate bash)"' >> ~/.bashrc
