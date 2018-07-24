#!/bin/sh
# setup.sh
# uninstall script for samarkand

if [ ! -w /etc/init.d ]; then
  echo "Run this script in more privileged mode please"
else
  if [ -d /etc/samarkand ]; then
    echo "Samarkand 8080 - Simple Web Server"
    echo "Uninstalling..."

    systemctl stop samarkand
    rm /lib/systemd/system/samarkand.service
    rm -r /etc/samarkand
    
    echo "Success!"
  else
    echo "Samarkand hasn't been installed yet"
  fi
fi
