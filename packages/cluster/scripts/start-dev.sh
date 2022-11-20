#!/bin/bash
# This script is meant to be run in the User Data of each EC2 Instance while it's booting. The script uses the
# run-nomad and run-consul scripts to configure and start Nomad and Consul in client mode. Note that this script
# assumes it's running in an AMI built from the Packer template in examples/nomad-consul-ami/nomad-consul.json.

set -e

# Send the log output from this script to user-data.log, syslog, and the console
# Inspired by https://alestic.com/2010/12/ec2-user-data-output/
exec > >(tee /var/log/user-data.log|logger -t user-data -s 2>/dev/console) 2>&1

# --- Mount the persistent disk with Firecracker environments.
# See https://cloud.google.com/compute/docs/disks/add-persistent-disk#create_disk
disk_name="sdb"
mount_dir="fc-envs"

mkdir -p /mnt/disks/$mount_dir
mount /dev/$disk_name /mnt/disks/fc-envs
chmod a+w /mnt/disks/$mount_dir

# Add automatic mounting on VM restart.
cp /etc/fstab /etc/fstab.backup
UUID=$(blkid /dev/sdb | tr ' ' '\n' | grep UUID)

echo "$UUID /mnt/disks/$mount_dir xfs defaults 0 0" >> /etc/fstab

mkdir -p /var/log/nomad/
mkdir -p /etc/nomad
mkdir -p /etc/nomad/plugins
cat <<EOF >> /etc/nomad/config.hcl
log_file = "/var/log/nomad/"
log_level = "DEBUG"
plugin_dir = "/opt/nomad/plugins"
plugin "firecracker-task-driver" {
  args = ["-telemetry-api=${telemetry_api_key}"]
}
EOF