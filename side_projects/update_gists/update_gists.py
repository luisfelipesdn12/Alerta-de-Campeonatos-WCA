import github, json

# Get secret information from secret.json file
# (with this dir as reference, it is in ../secret.json
# but the script will run in the `side_project dir`)
secret_json = json.load(open("./secret.json"))

# Instanciaste the GitHub object with the token
gh = github.Github(secret_json["GitHubToken"])

# Get objects of the gists to be update
resume_gist = gh.get_gist(secret_json["GitHubResumeGistURL"].split("/")[-1])
main_log_gist = gh.get_gist(secret_json["GitHubMainLogGistURL"].split("/")[-1])

# Read the files of resume json and main log
resume_content = " ".join(open("./../resume.json", encoding='utf-8').readlines())
main_log_content = open("./../main.log", encoding='utf-8').read()


resume_gist.edit(
    description="Information about Alerta-de-Campeonatos-WCA runtime",
    files={"resume.json": github.InputFileContent(content=resume_content)},
)

main_log_gist.edit(
    description="Alerta-WCA main log file",
    files={"main.log": github.InputFileContent(content=main_log_content)},
)

# Set the last commit hash in json 
secret_json["GitHubMainLogGistLastCommitHash"] = main_log_gist.history[0].version
with open("./secret.json", "w") as secret_json_file:
    json.dump(secret_json, secret_json_file, indent=4)
