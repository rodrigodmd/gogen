#!/bin/bash -l -i

CLI_BINARY="https://raw.githubusercontent.com/rodrigodmd/...."
COMMAND_NAME="gogen"

function log_op() {
  printf "\033[33m$*\033[0m\n"
}

function install_cli() {
  log_op "Installing $COMMAND_NAME"

  TMP_DIR=$(mktemp -d -t cli_dir)
  cd $TMP_DIR
  curl -o ${COMMAND_NAME} ${CLI_BINARY}
  sudo chmod 777 ${COMMAND_NAME}

  sudo rm -fr /usr/local/bin/${COMMAND_NAME}
  sudo mv ${COMMAND_NAME} /usr/local/bin/

  log_op "\n$COMMAND_NAME installed"

  rm -fr $TMP_DIR
}

install_cli

