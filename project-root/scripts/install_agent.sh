#!/bin/bash
# Usage: ./install_agent.sh -t <secret_token> -n <server_url>

while getopts "t:n:" opt; do
  case $opt in
    t) TOKEN="$OPTARG" ;;
    n) SERVER_URL="$OPTARG" ;;
    *) echo "Usage: $0 -t <secret_token> -n <server_url>" && exit 1 ;;
  esac
done

if [ -z "$TOKEN" ] || [ -z "$SERVER_URL" ]; then
  echo "Both secret token and server URL must be provided."
  exit 1
fi

echo "Downloading agent binary..."
# Replace the URL with the actual download location from your GitHub releases.
curl -L -o agent https://github.com/pandaaxi/homie_bot/releases/latest/download/agent
if [ $? -ne 0 ]; then
  echo "Failed to download agent binary."
  exit 1
fi

chmod +x agent

echo "Installing agent binary to /usr/local/bin..."
sudo mv agent /usr/local/bin/agent

echo "Creating systemd service file for agent..."
sudo tee /etc/systemd/system/homie-agent.service > /dev/null <<EOF
[Unit]
Description=Homie Bot Agent Service
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/agent -t ${TOKEN} -n ${SERVER_URL}
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF

echo "Reloading systemd daemon..."
sudo systemctl daemon-reload

echo "Enabling and starting the homie-agent service..."
sudo systemctl enable homie-agent.service
sudo systemctl start homie-agent.service

echo "Agent installed and service started successfully."
