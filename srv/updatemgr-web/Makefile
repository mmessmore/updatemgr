SVELTE_FILES := $(shell find src -name *.svelte)

build: public/build/bundle.js

public/build/bundle.js: $(SVELTE_FILES)
	npm run build

dev:
	npm run dev
