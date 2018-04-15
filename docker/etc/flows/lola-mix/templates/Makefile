all: clean build

ABR_DIR := ../abr-video
ABR_LABEL := tcp-abr-hd-960x720-10s
ABR_TMPL := tcp-abr-hd-960x720-10s.yaml.in

WEB_DIR := ../web-short
WEB_LABEL := tcp-web-short
WEB_TMPL := tcp-web-short.yaml.in

clean: clean-abr clean-web
build: build-abr build-web

# Web	
build-web:
	@./builder.py \
		--base_port 5004 \
		--num_files 6 \
		--label $(WEB_LABEL) \
		--template $(WEB_TMPL) \
		--num_at_points 8 \
		--at_points_distance 8 \
		--dest_dir $(WEB_DIR)

clean-web: ; @$(RM) $(WEB_DIR)/*.yaml

# ABR
build-abr:
	@./builder.py \
		--base_port 6000 \
		--num_files 18 \
		--label $(ABR_LABEL) \
		--template $(ABR_TMPL) \
		--num_at_points 6 \
		--at_points_distance 10 \
		--dest_dir $(ABR_DIR)

clean-abr: ; @$(RM) $(ABR_DIR)/*.yaml

.PHONY: clean clean-abr clean-web build build-abr build-web
