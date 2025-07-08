FROM jetpackio/devbox:latest@sha256:084acab190e61dcfcc234af1a8861633ca75fe6a72394c57de488de45bba9242

# Installing your devbox project
WORKDIR /code
COPY devbox.json devbox.json
COPY devbox.lock devbox.lock
RUN sudo chown -R "${DEVBOX_USER}:${DEVBOX_USER}" /code


RUN devbox run -- echo "Installed Packages."

RUN devbox shellenv --init-hook >> ~/.profile
