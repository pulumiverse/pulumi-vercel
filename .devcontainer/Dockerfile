FROM jetpackio/devbox:latest@sha256:687adfc866274d710814525bdb9f40c50707a15faffb44ed7e50745d26c48777

# Installing your devbox project
WORKDIR /code
COPY devbox.json devbox.json
COPY devbox.lock devbox.lock
RUN sudo chown -R "${DEVBOX_USER}:${DEVBOX_USER}" /code


RUN devbox run -- echo "Installed Packages."

RUN devbox shellenv --init-hook >> ~/.profile
