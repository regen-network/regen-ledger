#!/bin/bash

CNT=0
ITER=$1
SLEEP=$2
NUM_BLOCKS=$3
RESTART_FREQ=$4
NODE_ADDR=$5

if [ -z "$1" ]; then
  echo "Error: Need to input number of iterations to run."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Error: Need to input number of seconds to sleep between iterations."
  exit 1
fi

if [ -z "$3" ]; then
  echo "Error: Need to input block height to declare completion."
  exit 1
fi

if [ -z "$4" ]; then
  echo "Error: Need to input random container restart frequency."
  exit 1
fi

if [ -z "$5" ]; then
  echo "Error: Need to input node address to poll."
  exit 1
fi

docker_containers=( $(docker ps -q -f name=regen --format='{{.Names}}') )

while [ ${CNT} -lt "$ITER" ]; do
  curr_block=$(curl -s "$NODE_ADDR":26657/status | jq -r '.result.sync_info.latest_block_height')

  if [ ! -z "${curr_block}" ] ; then
    echo "Number of Blocks: ${curr_block}"
  fi

  if [ ! -z "${curr_block}" ] && [ "${curr_block}" -gt "${NUM_BLOCKS}" ]; then
    echo "Number of blocks reached. Success!"
    exit 0
  fi

  # Emulate network chaos:
  # - Pick a random container and restart it
  # - RESTART_FREQ is the number of blocks between restarts
  if ! ((CNT % RESTART_FREQ)); then
    rand_container=${docker_containers["$(RANDOM % ${#docker_containers[@]})"]};
    echo "Restarting random docker container ${rand_container}"
    docker restart "${rand_container}" &>/dev/null &
  fi

  (CNT++)

  sleep "$SLEEP"
done

echo "Timeout reached. Failure!"
exit 1
