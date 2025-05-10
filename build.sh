#!/bin/bash

set -e

export REPO_SRC=grafana/xk6
export REPO_DST=${GITHUB_REPOSITORY:-grafana/gh-xk6}

TAG_SRC=${1:-latest}

if [ "$TAG_SRC" = "latest" ]; then
    TAG_SRC=$(gh api /repos/$REPO_SRC/tags --jq '.[0].name')
fi

if gh api /repos/$REPO_DST/releases/tags/$TAG_SRC >/dev/null 2>&1; then
    cat >>${GITHUB_STEP_SUMMARY:-/dev/stderr} <<EOF
> [!WARNING]
> Tag $TAG_SRC already exists in $REPO_DST
EOF
    exit 0
fi

export DIR_WORK=build/xk6

rm -rf $DIR_WORK
gh repo clone $REPO_SRC $DIR_WORK

git -C $DIR_WORK checkout $TAG_SRC

cd $DIR_WORK
goreleaser build --id gh-xk6 --clean

echo "tag=$TAG_SRC" >>${GITHUB_OUTPUT:-/dev/stdout}
cat >>${GITHUB_STEP_SUMMARY:-/dev/stderr} <<EOF
> [!NOTE]
> The $TAG_SRC release was successful
EOF
