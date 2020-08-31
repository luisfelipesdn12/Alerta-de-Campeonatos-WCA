import github, json

secret_json = json.load(open("./secret.json"))

gh = github.Github(secret_json["GitHubToken"])

resume_gist = gh.get_gist(secret_json["GitHubResumeGistURL"].split("/")[-1])
main_log_gist = gh.get_gist(secret_json["GitHubMainLogGistURL"].split("/")[-1])

resume_content = " ".join(open("./../resume.json").readlines())
main_log_content = open("./../main.log").read()

resume_gist.edit(
    description="Information about Alerta-de-Campeonatos-WCA runtime",
    files={"resume.json": github.InputFileContent(content=resume_content)},
)

main_log_gist.edit(
    description="Alerta-WCA main log file",
    files={"main.log": github.InputFileContent(content=main_log_content)},
)