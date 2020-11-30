from github import Github
import subprocess
import time
	
username = input("Enter username: ")
user = Github().get_user(username)
print("Fetching repository names and commits\nIt may take a few seconds.....")
for repo in user.get_repos():	
	with open('commits.json', 'a') as file:
		file.write("\nCommits of Repository " + repo.name + " :\n" )
	subprocess.check_output("perceval git --json-line https://github.com/"+username+"/" + repo.name +">>commits.json", shell = True, stderr=subprocess.STDOUT)
	time.sleep(5)

print("Finished, data can be seen in commits.json in your directory")
