#!make
SERVICE_NAME = go-helpers

tag:
	@echo "Write the version (ex v1.0.47) : " && read VERSION && git tag $(VERSION) && git push origin $$VERSION;
	@echo "Type your commit message : " && read MESSAGE && git commit -m $$MESSAGE && git push origin $$VERSION;
