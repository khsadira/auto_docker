#!/usr/bin/env python3
import sys
import os

name = sys.argv[1]
repo_name = sys.argv[2]
#kill
#pull
#launch
os.system("docker rm -f " + name)
os.system("docker pull " + repo_name)
os.system("/devops/script/docker_run_script")
