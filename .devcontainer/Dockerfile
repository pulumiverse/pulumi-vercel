FROM mcr.microsoft.com/devcontainers/base:ubuntu@sha256:00e84e24112159d45c0262f07cb013cb58cb6a415f3e1c743d4a3115ac5d76c6

# Install mise
RUN curl https://mise.run | sh
ENV PATH="/root/.local/bin:${PATH}"

# Setup mise
WORKDIR /code
RUN mise install

# Activate mise in shell
RUN echo 'eval "$(mise activate bash)"' >> ~/.bashrc
