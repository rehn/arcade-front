if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.array = Object();


devdept.array.each = function(arr,fn,scope) {
  if (scope) {
    fn = fn.bind(scope)
  }
  if (!arr || arr == 'undefined') {
    return;
  }

  for (var i = 0; i < arr.length; i++) {
    fn(arr[i],i);  
  }
};

devdept.array.contains = function(arr, obj) {
  var i = arr.length;
  while (i--) {
    if(arr[i]==obj) {
      return true;
    }
  }
  return false;
};

