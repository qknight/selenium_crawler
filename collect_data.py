#! /usr/bin/env nix-shell
#! nix-shell --pure collect_data-environment.nix --command 'python3 collect_data.py'

from selenium import webdriver
from selenium import selenium
from selenium.webdriver.common.keys import Keys

from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait # available since 2.4.0
from selenium.webdriver.support import expected_conditions as EC # available since 2.26.0
from pyvirtualdisplay import Display

import sys
import os

display = Display(visible=0, size=(800, 600))
display.start()

from distutils.version import LooseVersion, StrictVersion
if LooseVersion(webdriver.__version__) < LooseVersion("2.51"):
    sys.exit("error: version of selenium ("
        + str(LooseVersion(webdriver.__version__))
        + ") is too old, needs 2.51 at least")

ff = webdriver.Firefox()

st = "https://www.swt-umweltpreis.de/profile/"
ff.get(st)

v = ff.execute_script("""
var t = document.getElementById('profile').childNodes; 

var ret = []
for (var i = 0; i < t.length; i++) {
  if ('id' in t[i]) {
    if(t[i].id.includes('profil-')) {
      var myID = t[i].id.replace("profil-","");
      var myVotes = t[i].getElementsByClassName('profile-txt-stimmen')[0].innerHTML;
      var myTitle = t[i].getElementsByClassName('archive-untertitel')[0].innerHTML;
      var myVerein = t[i].getElementsByClassName('archive-titel')[0].innerHTML;
      //console.log(myID,myVerein, myTitle, myVotes)
      var r = new Object();
      r.id = parseInt(myID);
      r.votes = parseInt(myVotes);
      r.verein = myVerein; 
      r.title = myTitle;
      ret.push(r)
    }
  }
}

var date = new Date();

var exp = {};

exp.date = date;
exp.data = ret;

var j = JSON.stringify(exp,  null, "\t");
console.log(j);
return j;
""")

print (v)

ff.quit()
display.stop()
