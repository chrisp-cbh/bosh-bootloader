#!/bin/sh -eux

cp bosh-lite-runc.yml ${BBL_STATE_DIR}/bosh-deployment/
cp bosh-lite-vm-type.yml ${BBL_STATE_DIR}/bosh-deployment/
cp bosh-lite.yml ${BBL_STATE_DIR}/bosh-deployment/
cp external-ip-not-recommended.yml ${BBL_STATE_DIR}/bosh-deployment/

cp bosh-lite_override.tf ${BBL_STATE_DIR}/terraform/

cp cloud-config.yml ${BBL_STATE_DIR}/.bbl/cloudconfig/
cp ops.yml ${BBL_STATE_DIR}/.bbl/cloudconfig/

cp create-director.sh ${BBL_STATE_DIR}/
cp delete-director.sh ${BBL_STATE_DIR}/
