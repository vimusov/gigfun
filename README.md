# What?

gigfun - Utility that takes control over your nVidia video card coolers to keep it cool and steady.

# Why?

nVidia cards can not control their own coolers! So that is a kludge to do it programmatically.

# How?

The program works as a daemon. When the GPU temperature rises it runs the coolers faster and vice versa.
Daemon should be started with Xorg because `nvidia-settings` is unable to work without it.

# Requirements

## Build

- `Go >= 1.16`;
- `just`;

## Start

- `sudo`;
- `nvidia-settings`;
- `nvidia-utils` (`nvidia-smi` utility);

# Setup

1. Build and install the binary to some place:

   ```
   just
   sudo just install /usr/local/bin
   ```

1. Create systemd unit config:

   ```
   # ~/.config/systemd/user/gigfun.service
   [Unit]
   Description=Kludge for a nVidia videocard
   After=graphical-session.target
   PartOf=graphical-session.target

   [Service]
   Type=exec
   ExecStart=/usr/bin/sudo /usr/local/bin/gigfun

   [Install]
   WantedBy=graphical-session.target
   ```

1. Reload systemd units:

   `systemctl --user daemon-reload`

1. Enable and start it:

   `systemctl --user --enable --now gigfun.service`

# License

GPL.
