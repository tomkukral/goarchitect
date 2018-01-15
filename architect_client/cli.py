#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
Inventory script to get the metadata from the Architect service.
"""

import click
from architect_client.lib import ArchitectClient


@click.command()
@click.argument('name')
def adapter_ansible(name):
    client = ArchitectClient()
    client.get_metadata('ansible', name)


@click.command()
@click.argument('name')
def adapter_salt_pillar(name):
    client = ArchitectClient()
    client.get_metadata('salt', name)


@click.command()
@click.argument('name')
def adapter_salt_top(name):
    client = ArchitectClient()
    client.get_inventory('salt', name)
