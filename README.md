# updatemgr

This is basically a toy for me to play with.

It's a NATS messaging-based system for managing OS package upgrades on
Debian-derived distributions.  I'm really just using it to manage my fleet of
small ARM machines (running Raspbian and Armbian).

There is a k8s deployment/Docker image that I use purely for functional
testing.

## How to run

Have NATS up and running somewhere.

Run `updatemgr agent` on the machines you want to manage.

Run `updatemgr server` on the thing you want to manage from.

It should track what updates are available on what hosts, allow you to trigger
an upgrade, track when a reboot is marked as being required (new kernel, new
glibc, etc), and allow you to reboot them.

## WARNING

There is basically no security around this thing at all.  Anyone with access to
the network could update or reboot your hosts, or probably cause a denial of
service.

It will likely burn down whatever building or hut you are in.
