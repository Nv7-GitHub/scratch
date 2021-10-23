format:
	cat Project/project.json | npx json > Project/formatted.json
	rm Project/project.json
	mv Project/formatted.json Project/project.json

build:
	zip -r Project.sb3 Project

download:
	mv ~/Downloads/Project.sb3 ./Project.zip
	rm -rf Project
	unzip Project.zip -d Project
	rm Project.zip