
====================
The Architect Client
====================

Client library and CLI for Architect API service.


Client Installation
===================

Following steps show how to deploy and configure Architect Client.

.. code-block:: bash

    pip install architect-client


SaltStack Integration
---------------------

To setup architect as Salt master Pillar source, set following configuration
to your Salt master.

.. code-block:: yaml

    ext_pillar:
      - cmd_yaml: 'architect-salt-pillar %s'

To setup architect as Salt master Tops source, set following configuration
to your Salt master.

.. code-block:: yaml

    master_tops:
       ext_nodes: architect-salt-top


Ansible Integration
-------------------

To setup architect as Ansible dynamic inventory source, set following
configuration to your Ansible control node.

.. code-block:: bash

    ansible -i architect-ansible-inventory
