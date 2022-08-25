#!/usr/bin/env bash

declare -r ORACLE_VERSION=21.7.0.0.0

# shellcheck disable=SC2016
# shellcheck disable=SC1004

declare -r SHORT_ORACLE_VERSION=$(echo ${ORACLE_VERSION} | sed 's/.//5g')
declare -r ORA_BASIC=instantclient-basic-linux.x64-${ORACLE_VERSION}dbru.zip
declare -r ORA_SDK=instantclient-sdk-linux.x64-${ORACLE_VERSION}dbru.zip

# bool function to test if the user is root or not (POSIX only)
is_user_root () { [ "$(id -u)" -eq 0 ]; }
error() { printf '\E[31m'; echo "$@"; printf '\E[0m' >&2; }
dontForget() {
    echo -en "\nDon't forget to add Oracle Instant Client ${ORACLE_VERSION} to your PATH variable into .bashrc or .zshrc, or maybe fish.config.\n\n"
    cat<<EOF
export ORACLE=/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')
export LD_LIBRARY_PATH=\$ORACLE:\$LD_LIBRARY_PATH
export PKG_CONFIG_PATH=\$ORACLE:\$PKG_CONFIG_PATH
EOF
}

echo $ORACLE_VERSION
echo $SHORT_ORACLE_VERSION

if !(is_user_root); then
    error 'You are just an ordinary user.'
    exit 1
fi

if [[ -d "/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')}" ]]; then
    echo "Oracle Instant Client ${ORACLE_VERSION} is installed."
    dontForget
    exit 0
fi

if [[ ! -f "/tmp/${ORA_BASIC}" ]] && [[ ! -f "/tmp/${ORA_SDK}" ]]; then
    wget -P /tmp https://download.oracle.com/otn_software/linux/instantclient/$(echo ${ORACLE_VERSION} | sed 's/\.//g')/${ORA_BASIC}
    wget -P /tmp https://download.oracle.com/otn_software/linux/instantclient/$(echo ${ORACLE_VERSION} | sed 's/\.//g')/${ORA_SDK}
fi

if [[ -f "/tmp/${ORA_BASIC}" ]]; then
    unzip /tmp/${ORA_BASIC} -d /opt/oracle
else
    error "Oracle Instant Client Basic ${ORACLE_VERSION} is not installed"
    exit 1
fi

if [[ -f "/tmp/${ORA_SDK}" ]]; then
    unzip /tmp/${ORA_SDK} -d /opt/oracle
else
    error "Oracle Instant Client SDK ${ORACLE_VERSION} is not found."
    exit 1
fi

if [[ -f "/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')/ojdbc8.jar" ]]; then
    echo "Oracle Instant Client ${ORACLE_VERSION} is installed."
else
    error "Oracle Instant Client ${ORACLE_VERSION} is not installed."
    exit 1
fi

if [[ ! -f "/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')/oci8.pc" ]]; then
    cat>/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')/oci8.pc<<EOF
prefixdir=/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')
libdir=\${prefixdir}
includedir=\${prefixdir}/sdk/include
Name: OCI
Description: Oracle database driver
Version: ${SHORT_ORACLE_VERSION}
Libs: -L\${libdir} -lclntsh
Cflags: -I\${includedir}
EOF
fi

if [[ -f "/etc/ld.so.conf.d/oracle.conf" ]]; then
    rm /etc/ld.so.conf.d/oracle.conf
fi
echo "/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')" > /etc/ld.so.conf.d/oracle.conf
ldconfig
dontForget

export ORACLE=/opt/oracle/instantclient_$(echo $ORACLE_VERSION | sed 's/.//5g;s/\./_/g')
export LD_LIBRARY_PATH=$ORACLE:$LD_LIBRARY_PATH
export PKG_CONFIG_PATH=$ORACLE:$PKG_CONFIG_PATH