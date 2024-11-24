#!/bin/sh
bokeh serve --address 195.19.32.95 --port 33001 --allow-websocket-origin=195.19.32.95:33001 --show $1
