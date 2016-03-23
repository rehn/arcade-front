if (typeof devdept === 'undefined') {
  devdept = Object();
}

devdept.json = Object();

devdept.extend = function (ChildClass, ParentClass) {
  ChildClass.prototype = new ParentClass();
  ChildClass.prototype.constructor = ChildClass;
}


devdept.json.serialize = function(obj) {
  if(typeof JSON.stringify=='function') {
    return JSON.stringify(obj);
  }
};



devdept.json.parse = function(str) {
  if(typeof JSON.parse=='function') {
    return JSON.parse(str);
  }
};

