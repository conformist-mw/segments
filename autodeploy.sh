#!/bin/bash

user=$(env | grep SUDO_USER | cut -d= -f 2)
if [[ -z $user ]] && [[ $UID -ne 0 ]]; then
	echo 'The script needs to run as root' && exit 1
fi
# apt-get install python3-pip python3-virtualenv virtualenv nginx git

su - $user << commands
cd /home/$user
mkdir project
cd project
virtualenv -p python3 .
. bin/activate
git clone https://github.com/conformist-mw/segments.git
cd segments
pip install -r requirements.txt
deactivate
commands

cd /home/$user/project/segments
sed -i "s/target_user/$user/g" etc/nginx/sites-available/project
sed -i "s/target_user/$user/g" etc/systemd/system/project.service
rm /etc/nginx/sites-enabled/*
cp etc/nginx/sites-available/project /etc/nginx/sites-available/
ln -s /etc/nginx/sites-available/project /etc/nginx/sites-enabled/
cp etc/systemd/system/project.service /etc/systemd/system/
systemctl daemon-reload
systemctl start project.service
systemctl enable project.service
systemctl restart nginx.service
