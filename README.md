# light-control

Control a light using the [Home Assistant](https://www.home-assistant.io/) Rest API.

I use this in concert with the waybar config in [my dotfiles](https://github.com/aloop/dotfiles) to control a light

### Configuration

By default this will attempt to load a configuration file from `light-control/config.json` under the user configuration directory specified determined by [os.UserConfigDir](https://pkg.go.dev/os#UserConfigDir).

It is not required to use this configuration file, as all the information can be alternately passed via the flags specified below.

## Usage

### Commands

-   `light-control` Outputs the current state and brightness in a json format compatible with waybar
-   `light-control toggle` Toggles the light on or off
-   `light-control brightness` Get the current brightness level as a percentage
-   `light-control brightness +10` increase brightness by 10%
-   `light-control brightness 10` sets the brightness to 10%
-   `light-control brightness -10` decrease brightness by 10%

### Flags

-   `light-control --config /path/to/config`
-   `light-control --entity light.ceiling_light`
-   `light-control --token SOME_TOKEN`
-   `light-control --host https://example.com`
