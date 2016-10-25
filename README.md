# what

this repository contains two projects for mining data online:

 * collect_data.py  
   selenium based tool which `scrapes a webpage`, produces a `javascript object` and saves it as `json`


 * inject_into_fluxdb  
   go based `influxdb` importer for the `json` content

# how

both projects come with a `nix-shell` environment definition which makes it easy to build and execute the software on any machine.

# who

* joachim schiele <js@lastlog.de>
* paul seitz <paul.m.seitz@gmail.com>
* paul rosset <paul.rosset@mp-ndt.de>

