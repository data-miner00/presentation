ADDR ?= 127.0.0.1:3999

.PHONY: serve

# Serve every deck under talks/, with presenter notes enabled.
serve:
	present -content talks -http $(ADDR) -notes
