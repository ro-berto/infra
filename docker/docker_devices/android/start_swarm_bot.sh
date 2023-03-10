#!/bin/bash -x

trap "exit 10" SIGUSR1

SWARM_DIR=/b/swarming
SWARM_ZIP=swarming_bot.zip

DEPOT_TOOLS_DIR=/b/depot_tools
DEPOT_TOOLS_URL='https://chromium.googlesource.com/chromium/tools/depot_tools.git'
DEPOT_TOOLS_REV='da3a29e13e816459234b0b08ed1059300bae46dd'
LUCI_MACHINE_TOKEN_FILE='/var/lib/luci_machine_tokend/token.json'

# Wait until this container has access to a device before starting the bot.
START=$(/bin/date +%s)
TIMEOUT=$((60*5))
while [ ! -d /dev/bus/usb ]
do
  now=$(/bin/date +%s)
  if [[ $((now-START)) -gt $TIMEOUT ]]; then
    echo "Timed out while waiting for an available device. Quitting early." 1>&2
    exit 1
  else
    echo "Waiting for an available usb device..."
    sleep 10
  fi
done

# Some chromium tests need depot tools.
mkdir -p $DEPOT_TOOLS_DIR
chown chrome-bot:chrome-bot $DEPOT_TOOLS_DIR
su -c "cd $DEPOT_TOOLS_DIR && \
       /usr/bin/git init && \
       /usr/bin/git remote add origin $DEPOT_TOOLS_URL ; \
       /usr/bin/git fetch origin $DEPOT_TOOLS_REV && \
       /usr/bin/git reset --hard FETCH_HEAD" chrome-bot

curl_header_args=()
if [[ -r "${LUCI_MACHINE_TOKEN_FILE}" ]]; then
  luci_token="$(jq -r ".luci_machine_token" "${LUCI_MACHINE_TOKEN_FILE}")"
  curl_header_args=("--header" "X-Luci-Machine-Token: ${luci_token}")
fi

mkdir -p $SWARM_DIR
chown chrome-bot:chrome-bot $SWARM_DIR
cd $SWARM_DIR
rm -rf swarming_bot*.zip
sudo -u chrome-bot -- /usr/bin/curl -sSLOJ "${curl_header_args[@]}" "${SWARM_URL}"

echo "Starting $SWARM_ZIP"
# Run the swarming bot in the background, and immediately wait for it. This
# allows the signal trapping to actually work.
su -c "/usr/bin/python3 $SWARM_ZIP start_bot" chrome-bot &
wait %1
exit $?
