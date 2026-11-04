FROM mcr.microsoft.com/devcontainers/base:ubuntu@sha256:b1e3b08610ae0b2599f57c3a492eebc93d1d26dbaefb4a0bad0caa7149e4ec95

# Install mise
RUN curl https://mise.run | sh
ENV PATH="/root/.local/bin:${PATH}"

# Setup mise
WORKDIR /code
RUN mise install

# Activate mise in shell
RUN echo 'eval "$(mise activate bash)"' >> ~/.bashrc
