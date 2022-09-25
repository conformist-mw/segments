deploy:
	ansible-playbook -i ansible/inventory.yml ansible/deploy.yml
