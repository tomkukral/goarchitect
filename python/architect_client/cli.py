#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
Inventory script to get the metadata from the Architect service.
"""

import os
import yaml
import click
from architect_client.libarchitect import ArchitectClient

default_inventory_api_url = os.environ.get('ARCHITECT_INVENTORY_API_URL',
                                           'http://localhost:8181')
default_inventory_name = os.environ.get('ARCHITECT_INVENTORY_NAME',
                                        'default')


@click.command()
@click.argument('resource_name')
def adapter_ansible_inventory(resource_name):
    client = ArchitectClient(default_inventory_api_url, default_inventory_name)
    data = client.get_data('ansible-inventory', resource_name)
    print(yaml.safe_dump(data))


@click.command()
@click.argument('resource_name')
def adapter_salt_pillar(resource_name):
    client = ArchitectClient(default_inventory_api_url, default_inventory_name)
    raw_data = client.get_data('salt-pillar', resource_name)
    output_data = {}
    for datum_name, datum in raw_data.items():
        output_data[datum_name] = datum['parameters']
    print(yaml.safe_dump(output_data[resource_name]))


@click.command()
@click.argument('resource_name')
def adapter_salt_top(resource_name):
    client = ArchitectClient(default_inventory_api_url, default_inventory_name)
    raw_data = client.get_data('salt-top', resource_name)
    output_data = {}
    for datum_name, datum in raw_data.items():
        output_data[datum_name] = datum['applications']
    print(yaml.safe_dump({'classes': output_data[resource_name]}))
