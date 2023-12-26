#!/usr/bin/env bash

# Kind
kind create cluster

# Flux

# Source Controller
flux install --namespace flux-system --network-policy false --components source-controller

# Azure Controller
export IMG=ghcr.io/ljtill/manager:v0.0.1
make docker-build
kind load docker-image $IMG
make install
make deploy
