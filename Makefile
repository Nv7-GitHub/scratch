format:
	cat Project/project.json | npx json > Project/formatted.json
	rm Project/project.json
	mv Project/formatted.json Project/project.json

build:
	rm -f Project.sb3
	zip -r Project.sb3 Project

download:
	mv ~/Downloads/Project.sb3 ./Project.zip
	rm -rf Project
	unzip Project.zip -d Project
	rm Project.zip

loadtest:
	rm -rf Project
	unzip testdata/Project.sb3 -d Project