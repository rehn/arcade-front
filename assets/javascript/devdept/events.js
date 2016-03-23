if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.events = Object();


devdept.events.EventHandler = function() {
  this.listeners = Array();
}

devdept.events.EventHandler.prototype.listen = function(obj,key,fn,scope) {
  if(scope) {
    fn = fn.bind(scope);
  }
  if(typeof key == 'object') { 
    devdept.array.each(key,function(k) {
      obj.addEventListener(k,fn);
      this.listeners.push({'key':k,'obj':obj,'fn':fn});
    },this);
  }
  else {
    obj.addEventListener(key,fn);
    this.listeners.push({'key':key,'obj':obj,'fn':fn});
  }
}
devdept.events.EventHandler.prototype.listenOnce = function(obj,key,fn,scope) {
  if(scope) {
    fn = fn.bind(scope);
  }
  if ( typeof key == 'array') {
    devdept.array.each(key,function(k){
      obj.addEventListener(k,fn);
      this.listeners.push({'key':k,'obj':obj,'fn':fn});
    },this);
  }
  else {
    obj.addEventListener(key,fn);
    this.listeners.push({'key':key,'obj':obj,'fn':fn});
  }
}

devdept.events.EventHandler.prototype.removeAll = function() {
  devdept.array.each(this.listeners,function(value) {
    value['obj'].removeEventListener(value['key'],value['fn']);
  });
  this.listeners = Array();
}


devdept.events.listen = function(obj,key,fn,scope) {
  if(scope){
    fn = fn.bind(scope);
  }
  obj.addEventListener(key,fn);
}

devdept.events.EventTarget = function() {
  this.listeners = {};
};

devdept.events.EventTarget.prototype.listen = function(key,fn,scope) {
  if(Object.prototype.toString.call( key ) === '[object Array]'){
    devdept.array.each(key,function(value){
      if (!this.listeners[value]) {
        this.listeners[value] = [];
      }
      this.listeners[value].push({'fn':fn,'scope':scope});
    },this);
  }
  else {
    if (!this.listeners[key]) {
      this.listeners[key] = [];
    }
    this.listeners[key].push({'fn':fn,'scope':scope});  
  }
};

devdept.events.EventTarget.prototype.dispatchEvent = function(data) {
  var key = "";
  if (typeof data == 'object'){
    key = data['type'];
  }
  if (this.listeners[key]) {
    this.data = data;
    devdept.array.each(this.listeners[key],function(obj){
      var scope = obj['scope'];
      var func = obj['fn'].bind(scope);
      func(data);
    },this);
    
  }
};



