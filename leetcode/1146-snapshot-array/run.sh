#!/usr/bin/env bash

rustc --test main.rs && ./main --nocapture
