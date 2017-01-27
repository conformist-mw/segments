# Segments accounting

This project allows you doing segments (rectangles) accounting by size, type or color. Removed segments doesn't remove from database just mark them as inactive. Also removed segments have order number. 

## Installation

In order to install this project you have to download [autodeploy.sh](https://raw.githubusercontent.com/conformist-mw/segments/master/autodeploy.sh) script and run it as sudo user. 

It will create dir named `project` in your home directory and python virtualenv within it. Sudo writes needs to create systemd service file for `gunicorn` and `nginx` default server. 
So you can run in terminal:

```bash
$ sudo bash autodeploy.sh
```
After this your project avaiable [here](http://localhost).
