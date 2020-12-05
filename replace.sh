#!/bin/bash
## Author - Guilherme SETTE <developer.guilherme@gmail.com>
## Github - https://github.com/numeroSette

## This script replaces a repository name, it follows the Github name pattern
## This script accomplish a specific need in this case replace testing/sre-test-1 to the desired name
## If you need to change this default name, you can just edit this file, but feel free to contribute creating a new feature :-)

## How can I call this script
# > You need to allow it to execute depending on your OS, for example: chmod 755 or chmod +x, depending on your OS
# > After that you can just call it ./replace.sh on your terminal
# > You will need to inform a parameter, like this: ./replace.sh <Your Repository name>, for example: ./replace.sh numeroSette/RepositoryNameReplacer

## You can find above some instructions about how it works:

if [[ $1 =~ ^([a-zA-Z0-9]|\-|_)+\/([a-zA-Z0-9]|\-|_)+$ ]]; then
# Validate the argument informed if matches as username/repository, following the Github pattern for repository names

    export REPLACE_ARG_1=$(echo $1)
    # Makes the variable available to Perl in order to pass it as an argument

    find . -type f -not -path "*.git/*" ! -name "README.md" ! -name "RESULT.md" ! -name "replace.sh" -name "*" -exec perl -p -i -e 's/testing\/sre-test-1/$ENV{REPLACE_ARG_1}/g' {} \;
    # Uses find ignoring .git directory (can cause some troubles if you don't ignore it)
    # Ignore some more file as you prefer and get all the rest
    
    unset REPLACE_ARG_1
    # Unset the variable in order to prevent troubles and conflicts

    echo "Files replaced":
    grep -r --exclude-dir=.git --exclude=\*.md "$(echo $1 | sed 's:/:\\/:g')" .
    # Give some feed back about what files are replaced
else
    echo "Invalid repository name, you need to follow the Github pattern for repository names"
    # Give some feedback if the repository name doesn't match with the Github pattern for repository names
fi

## TODO ##
# Include some way to change the default repository, passing it as a parameter.
# Include some way to pass ignored directories and files as parameters

### This script works on the following OS:
## Note that it would require some dependecies depending on the version of the SO

## macOS
# $ system_profiler SPSoftwareDataType
# Software:

#     System Software Overview:

#       System Version: macOS 10.15.7 (19H15)
#       Kernel Version: Darwin 19.6.0
#       Boot Volume: Samsung SSD 860 EVO 250GB
#       Boot Mode: Normal
#       Computer Name: MacBook Pro de Guilherme
#       User Name: Guilherme Sette (stock)
#       Secure Virtual Memory: Enabled
#       System Integrity Protection: Enabled
#       Time since boot: 1:55

## Amazon Linux
# $ cat /etc/os-release
# NAME="Amazon Linux"
# VERSION="2"
# ID="amzn"
# ID_LIKE="centos rhel fedora"
# VERSION_ID="2"
# PRETTY_NAME="Amazon Linux 2"
# ANSI_COLOR="0;33"
# CPE_NAME="cpe:2.3:o:amazon:amazon_linux:2"
# HOME_URL="https://amazonlinux.com/"

## Ubuntu
# $ cat /etc/os-release
# NAME="Ubuntu"
# VERSION="18.04.5 LTS (Bionic Beaver)"
# ID=ubuntu
# ID_LIKE=debian
# PRETTY_NAME="Ubuntu 18.04.5 LTS"
# VERSION_ID="18.04"
# HOME_URL="https://www.ubuntu.com/"
# SUPPORT_URL="https://help.ubuntu.com/"
# BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
# PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
# VERSION_CODENAME=bionic
# UBUNTU_CODENAME=bionic