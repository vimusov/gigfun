# What?

gigfun - Utility that takes control over your nVidia video card coolers to keep it cool and steady.

# Why?

My GeForce 3090 Ti can not control its own coolers! So that is a kludge to do it programmatically.

# How?

The program works as a daemon. When the GPU temperature rises it runs the coolers faster and vice versa.
Daemon must be started with Xorg because in requires "Coolbits" options (see below).

# Requirements

- Go >= 1.16;
- `nvidia-settings`;
- `nvidia-utils` (утилита `nvidia-smi`);

# Setup

1. Setup the option "Coolbits" how it is [described here](https://wiki.archlinux.org/title/NVIDIA/Tips_and_tricks#Enabling_overclocking):

   ```
   # /etc/X11/xorg.conf.d/nvidia.conf
   Section "Device"
       Identifier "Device0"
       Driver     "nvidia"
       Option     "Coolbits" "4"
   EndSection
   ```

2. Copy systemd unit to a local user and enable it:

   ```
   install -D -m 0644 /usr/share/gigfun/service.example ~/.config/systemd/user/gigfun.service
   systemctl --user daemon-reload
   systemctl --user enable --now gigfun.service
   ```

# License

GPL.
