#!/bin/bash

chmod +x /opt/go-working-calendar/go-working-calendar-linux
chmod +x /opt/go-working-calendar/install-service.sh
cp ./go-working-calendar.service  /etc/systemd/system
systemctl enable go-working-calendar
systemctl start go-working-calendar
