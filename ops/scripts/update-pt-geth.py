#!/usr/bin/env python3


import json
import subprocess
import os


GETH_VERSION='v1.11.5'


def main():
	for project in ('.', 'indexer'):
		print(f'Updating {project}...')
		update_mod(project)


def update_mod(project):
	print('Replacing...')
	subprocess.run([
		'go',
		'mod',
		'edit',
		'-replace',
		f'github.com/ethereum/go-ethereum@{GETH_VERSION}=github.com/patex-ecosystem/patex-chain@main'
	], cwd=os.path.join(project), check=True)
	print('Tidying...')
	subprocess.run([
		'go',
		'mod',
		'tidy'
	], cwd=os.path.join(project), check=True)


if __name__ == '__main__':
	main()
