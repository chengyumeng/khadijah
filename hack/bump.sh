#!/bin/sh
# Bump project version



git pull

CHANGELOG=CHANGELOG.md
GO_MAIN=main.go

VERSION=(`grep -Eo "[0-9]+\.[0-9]+\.[0-9]+[a-z0-9\-]*" $GO_MAIN`)
echo "Project current version: [$VERSION] "


BASES=(`echo $VERSION | tr '.' ' '`)
V_MAJOR=${BASES[0]}
V_MINOR=${BASES[1]}
V_PATCH=${BASES[2]}

if [ "$1" = "" ]; then
   V_PATCH=$((V_PATCH + 1))
   NEXT_VERSION="$V_MAJOR.$V_MINOR.$V_PATCH"
elif [ "$1" = "major" ]; then
    V_MAJOR=$((V_MAJOR + 1))
    V_MINOR=0
    V_PATCH=0
    NEXT_VERSION="$V_MAJOR.$V_MINOR.$V_PATCH"
elif [ "$1" = "minor" ]; then
    V_MINOR=$((V_MINOR + 1))
    V_PATCH=0
    NEXT_VERSION="$V_MAJOR.$V_MINOR.$V_PATCH"
elif [ "$1" = "patch" ]; then
    V_PATCH=$((V_PATCH + 1))
    NEXT_VERSION="$V_MAJOR.$V_MINOR.$V_PATCH"
else
    BASES=(`echo $1 | tr '.' ' '`)
    V_MAJOR=${BASES[0]}
    V_MINOR=${BASES[1]}
    V_PATCH=${BASES[2]}
    NEXT_VERSION=$1
fi
RELEASE_BRANCH="release/v$V_MAJOR.$V_MINOR"

read -p "Change branch to [$RELEASE_BRANCH]? (Y): " CONFIRM0
if [ "$CONFIRM0" = "" ]; then CONFIRM0="Y"; fi
if [ "$CONFIRM0" = "y" ]; then CONFIRM0="Y"; fi
if [ "$CONFIRM0" = "yes" ]; then CONFIRM0="Y"; fi
if [ "$CONFIRM0" = "YES" ]; then CONFIRM0="Y"; fi
if [ "$CONFIRM0" = "Y" ]; then
    exists=$(git branch | grep "$RELEASE_BRANCH")
    if [ -n "$exists" ]; then
        git checkout $RELEASE_BRANCH
        git pull origin $RELEASE_BRANCH
        git pull origin master
    else
        git checkout -b $RELEASE_BRANCH
    fi
    echo "$(tput setaf 2) Success checkout to branch: [$RELEASE_BRANCH] $(tput sgr0)"
else
    echo "Refuse to switch branches and quit!"
    exit 1
fi

read -p "Bump to [$NEXT_VERSION] and update changelog? (Y): " CONFIRM1
if [ "$CONFIRM1" = "" ]; then CONFIRM1="Y"; fi
if [ "$CONFIRM1" = "y" ]; then CONFIRM1="Y"; fi
if [ "$CONFIRM1" = "yes" ]; then CONFIRM1="Y"; fi
if [ "$CONFIRM1" = "YES" ]; then CONFIRM1="Y"; fi
if [ "$CONFIRM1" = "Y" ]; then
    sed -i "s/$VERSION/$NEXT_VERSION/" $GO_MAIN
    git changelog --no-merges --tag $NEXT_VERSION $CHANGELOG

    COMMITLOG="
Release v$NEXT_VERSION

$(git changelog -x -n -p -l)
"

    git commit -ae -m "$COMMITLOG"

    echo "$(tput setaf 2)Update changelog Done.$(tput sgr0)"
fi

read -p "Exec git tag v$NEXT_VERSION ? (Y): " CONFIRM2
if [ "$CONFIRM2" = "" ]; then CONFIRM2="Y"; fi
if [ "$CONFIRM2" = "y" ]; then CONFIRM2="Y"; fi
if [ "$CONFIRM2" = "yes" ]; then CONFIRM2="Y"; fi
if [ "$CONFIRM2" = "YES" ]; then CONFIRM2="Y"; fi
if [ "$CONFIRM2" = "Y" ]; then
    git tag v$NEXT_VERSION
    echo "$(tput setaf 2)Create tag done.$(tput sgr0)"
fi

read -p "Merge branch [$RELEASE_BRANCH] into master? (Y): " CONFIRM3
if [ "$CONFIRM3" = "" ]; then CONFIRM3="Y"; fi
if [ "$CONFIRM3" = "y" ]; then CONFIRM3="Y"; fi
if [ "$CONFIRM3" = "yes" ]; then CONFIRM3="Y"; fi
if [ "$CONFIRM3" = "YES" ]; then CONFIRM3="Y"; fi
if [ "$CONFIRM3" = "Y" ]; then
    git checkout master
    git merge $RELEASE_BRANCH
    echo "$(tput setaf 2)Merge Done.Now on branch master.$(tput sgr0)"
fi

read -p "Push master branch/release branch and tag? (Must Insert YES): " CONFIRM4
if [ "$CONFIRM4" = "YES" ]; then
    git push origin master:master
    git push origin $RELEASE_BRANCH
    git push origin v$NEXT_VERSION
    echo "$(tput setaf 2)Push Done.Now on branch master.$(tput sgr0)"
fi

