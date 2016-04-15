# kube-watcher

## Description

Watches your pods for states that are anything other than Running and creates a datadog event (so that you can setup alerts on important events). Only works inside of a kubernetes cluster.

## Usage:

You must alter the deployment.yaml for your own deployment needs and add your own Datadog API key as well as Application Key. You may also want to alter the makefile 

## TODO:

* Switch from set interval to using api rate limiter
* Switch to using channels to watch pods
* Use variables in makefile to reference Vungle
* Migrate to a telegraf input plugin.
* Support Terminating
