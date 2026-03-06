ARG ALPINE_VERSION=3.23
ARG GOLANG_VERSION=1.26.0
ARG BUILD_TAGS=ui
ARG CGO_ENABLED=0
ARG VAULT_VERSION=1.21.3

# ---

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS base

RUN apk add --no-cache \
        git \
        build-base \
        bash

ARG BUILD_TAGS
ARG CGO_ENABLED
ENV BUILD_TAGS=${BUILD_TAGS}
ENV CGO_ENABLED=${CGO_ENABLED}

# ---

FROM base AS dev

ENV PATH="/home/${USERNAME}/bin:${PATH}"

RUN apk add --no-cache \
        bash-completion \
        openssh-client \
        git-bash-completion \
        sudo \
    && mkdir -p /etc/skel \
    && touch /etc/skel/.bashrc \
    && echo "source /usr/share/bash-completion/completions/git" >> /etc/skel/.bashrc

ARG USERNAME
ARG USER_ID
ARG GROUP_ID

RUN addgroup -g ${GROUP_ID} ${USERNAME} \
    && adduser -S -D -u ${USER_ID} -G ${USERNAME} ${USERNAME} \
    && echo "${USERNAME} ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME} \
    && cd /home/${USERNAME} \
    && mkdir -p -m0755 src \
    && echo "export PATH=\${HOME}/bin:\${PATH}" >> .bashrc \
    && chown -R ${USER_ID}:${GROUP_ID} src .bashrc \
    && install -m 644 /dev/null .gitconfig

RUN    go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install golang.org/x/tools/gopls@latest

# ---
    
# MUST be called from repo root (hashicorp) context
FROM base AS build

ARG VAULT_VERSION

COPY vault/v${VAULT_VERSION}/modif /root/vault/v${VAULT_VERSION}/modif
WORKDIR /root/vault/v${VAULT_VERSION}/modif
COPY .git .git

RUN make bootstrap
RUN scripts/build.sh

# ---

FROM hashicorp/vault:${VAULT_VERSION} AS target

COPY --from=build /root/vault/bin/vault bin/vault
