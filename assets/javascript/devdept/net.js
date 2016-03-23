if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.net = Object();


devdept.net.XhrIo = function(url,method,data,scope) {
  this.url=url;
  this.method=method;
  this.data = data;
  this.scope = scope;
  this.evh = new devdept.events.EventHandler();
  this.xhr = new XMLHttpRequest();
 
  this.evh.listen(this.xhr,'load',function(e) {  
    if (this.xhr.status != 200) {
      if (typeof this.onError == 'function') {
        this.onError({ 
          status : this.xhr.status, 
          text : this.xhr.statusText
        });
      }
      return;
    }
    if (typeof this.onComplete == 'function') {
      if (this.scope){
        this.onComplete= this.onComplete.bind(this.scope);
      }
      var json={};
      try {
        var json = devdept.json.parse(this.xhr.responseText);
      } catch(w){};
      this.onComplete({ 
          status : this.xhr.status, 
          text : this.xhr.statusText,
          responseJson : json,
          responseText : this.xhr.responseText,
          responseType : this.xhr.responseType
        });
    }
  },this);

  this.evh.listen(this.xhr,'error',function(e) {
    if (typeof this.onError == 'function') {
      this.onError(this.xhr);
    }
  },this);
}
devdept.extend(devdept.net.XhrIo,devdept.events.EventTarget);

devdept.net.XhrIo.prototype = {
  send : function() {
    this.xhr.open(this.method, this.url);

    if (this.method=="POST") {
      //this.xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    }
    this.xhr.send(this.data);
  }
}