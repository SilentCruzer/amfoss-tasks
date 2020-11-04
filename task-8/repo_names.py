from github import Github
import re
import json
import subprocess
import time
repo_name = []
username = "amfoss"
user = Github().get_user(username)
for repo in user.get_repos():
	name = re.findall("/.*\w", str(repo))
	repo_name +=name	

print("Fetching repository names and commits\nIt may take a few seconds.....")	

for i in repo_name:
	with open('commits.json', 'a') as file:
		file.write("\nCommits of Repository " + i + " :\n" )
	subprocess.check_output("perceval git --json-line https://github.com/amfoss" + i +">>commits.json", shell = True, stderr=subprocess.STDOUT)
	time.sleep(5)



print("Finished, data can be seen in commits.json in your directory")
	