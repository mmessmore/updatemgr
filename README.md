# updatemgr

This is basically a toy for me to play with using tools I'm learning.

It's a NATS messaging-based system for managing OS package upgrades on
Debian-derived distributions.  I'm really just using it to manage my fleet of
small ARM machines (running Raspbian and Armbian).  But it also supports
x86_64.

I'm using it to play with NATS and bbolt from Go because I need "real" example
projects to actually learn stuff with.  I also really stink at front-ends.

This uses Svelte, and maybe one day I'll make it less ugly.  I wanted to embed
a front-end in the binary to make it self-contained.  That was fun.

Note: The server is not HA/horizontally scalable.  I'll never fix that.  It's
not necessary.

## How to run

Have NATS up and running somewhere.

Run `updatemgr agent` on the machines you want to manage.

Run `updatemgr server` on the thing you want to manage from.

You can specify the NATS server in the config file or on the command line.

It should track what updates are available on what hosts, allow you to trigger
an upgrade, track when a reboot is marked as being required (new kernel, new
glibc, etc), and allow you to reboot them.

[Releases](https://github.com/mmessmore/updatemgr/releases) have raw binaries
for x86_64, 32-bit armv7, and 64-bit arm, as well as `.deb` packages for each.

The deb packages have default configurations that are still specific to my
world, but can easily be updated for yours.  They are located in
`/etc/updatemgr/*.yml`.  They also include systemd unit files for the agent and
server, but are not enabled by default.  If already enabled, they will restart the
service on upgrade.

## CONTRIBUTING

I'll take most contributions.  The [MIT LICENSE](./LICENSE) is not negotiable,
so please don't ask.

This is a total side-project.  Issues will likely be ignored without an
accompanying PR.  I'll probably be slow to respond.

## KNOWN ISSUES

- [ ] logging config is broken
- [ ] Likely has front-end UI bugs and is ugly
- [ ] Needs authentication around the front-end
- [ ] Missing any testing
- [ ] Missing Actions for testing and maintaining library updates

## WARNING

There is no security around this thing at all.  Anyone with access to
the network could update or reboot your hosts, or probably cause a denial
of service.

It will likely burn down whatever building, yurt, or hut you are in.
