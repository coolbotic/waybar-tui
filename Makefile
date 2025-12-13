.PHONY: build  

default: build reload


build:
	go build -o waybar-tui

reload:
	pkill -x waybar
	setsid uwsm-app -- waybar >/dev/null 2>&1 &
	hyprctl reload

# Run to configure hyprland to display test window in centre of the screen
setup-hypr:
	hyprctl keyword windowrulev2 'float,title:^PangoCairo:.*$$'
	hyprctl keyword windowrulev2 'center,title:^PangoCairo:.*$$'
	hyprctl keyword windowrulev2 'workspace current,title:^PangoCairo:.*$$'

reset-hypr:
	hyprctl reload

test-local: build
	pango-view --markup \
		--background='#1e1e2e' \
		--font="JetBrainsMono Nerd Font 9" \
		--margin=20 \
		-t "$$(./waybar-tui weather | jq -r '.tooltip')"
