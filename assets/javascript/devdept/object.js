if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.object = Object();


devdept.object.each = function(obj,fn,opt_scope) {
  if (opt_scope) {
    fn = fn.bind(opt_scope)
  }
  for (var key in obj) {
    if(obj.hasOwnProperty(key)){
      fn(obj[key],key);
    }
  }
};

devdept.object.merge = function(target,src) {
  devdept.object.each(src,function(val,key) {
    target[key] = val;
  });
  return target;
};

devdept.object.count = function (obj) {
var i = 0;
  for (var key in obj) {
    if(obj.hasOwnProperty(key)){
      i++;
    }
  }
  return i;
}

devdept.object.findKey = function(obj,key) {
  var o =  obj[key];
  if (o) {
    return o;
  }
  else {
    return null;
  }
};
