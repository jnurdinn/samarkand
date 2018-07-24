#!/bin/sh
# setup.sh
# setup script for samarkand

echo "Samarkand 8080 - Simple Web Server"
if [ ! -w /etc/init.d ]; then
  echo "Run this script in more privileged mode please"
else
  echo "Installing..."
  cp samarkand /usr/bin
  cp -r config/ /etc/
  mv /etc/config /etc/samarkand

  echo "[Unit]" >> samarkand.service
  echo "Description=Samarkand 8080 - Simple Web Server" >> samarkand.service
  echo "After=network.target" >> samarkand.service
  echo "StartLimitIntervalSec=0" >> samarkand.service
  echo "[Service]" >> samarkand.service
  echo "Type=simple" >> samarkand.service
  echo "Restart=always" >> samarkand.service
  echo "RestartSec=1" >> samarkand.service
  echo "User=$USER" >> samarkand.service
  echo "ExecStart=/usr/bin/samarkand" >> samarkand.service
  echo "" >> samarkand.service
  echo "[Install]" >> samarkand.service
  echo "WantedBy=multi-user.target" >> samarkand.service

  mv samarkand.service /lib/systemd/system/

  if [ ! -d "/var/www/html" ]; then
    mkdir /var/www/html
    echo "SAMARKAND 8080 SIMPLE WEBSERVER INSTALLED SUCCESSFULLY" >> /var/www/html/index.html
  fi

  systemctl start samarkand
  echo "Success!"
  echo "Don't forget to config on /etc/samarkand/config.toml"
fi
