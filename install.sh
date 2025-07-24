#!/bin/bash

# Check if the user is root
if [ "$(id -u)" != "0" ]; then
    echo "This script must be run as root" 1>&2
    exit 1
fi

# Update the package list
apt-get update

# Install dependencies
apt-get install -y wget
wget https://go.dev/dl/go1.24.5.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc

# Create the application directory
mkdir -p /opt/xray-panel

# Copy the application files
cp -r . /opt/xray-panel/

# Build the application
cd /opt/xray-panel
go build -o xray-panel .

# Create the systemd service file
cat > /etc/systemd/system/xray-panel.service << EOL
[Unit]
Description=Xray Panel
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/xray-panel
ExecStart=/opt/xray-panel/xray-panel
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOL

# Reload the systemd daemon
systemctl daemon-reload

# Enable the service
systemctl enable xray-panel

# Start the service
systemctl start xray-panel

echo "Xray Panel has been installed and started successfully."
