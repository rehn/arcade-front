if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.url=Object();


devdept.url.UrlHandler = function() {
  this.protocol = document.location.protocol
  this.update();
  devdept.events.listen(window,'hashchange',this.hashChanged,this);
};
devdept.extend(devdept.url.UrlHandler,devdept.events.EventTarget);

devdept.url.UrlHandler.prototype.update = function() {
  this.protocol = document.location.protocol
  this.url = this.protocol + "//" + document.location.hostname + document.location.pathname;
  this.path = document.location.pathname;
  this.host = document.location.hostname;
  this.hashString = document.location.hash;
  this.hash = Array();
  if (devdept.string.startsWith(this.hashString,"#")) {
    this.hash = this.hashString.substring(1).split("/"); 
  }
};

devdept.url.UrlHandler.prototype.setHash = function(str) {
  document.location.hash = str;
}

devdept.url.UrlHandler.prototype.back = function() {
  window.history.back();
}

devdept.url.UrlHandler.prototype.hashChanged = function() {
  this.update();
  this.dispatchEvent({'type':'HASH_CHANGED','value':this.hash});
};
