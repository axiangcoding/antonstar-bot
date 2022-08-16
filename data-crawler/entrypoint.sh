#!/bin/bash
# FIXME: 依然有问题

set -m

scrapyd &
scrapyd-deploy

jobs

fg 1