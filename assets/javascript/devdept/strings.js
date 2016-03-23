if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.string = Object();

devdept.string.startsWith = function(str,pattern) {
  return (str.indexOf(pattern) == 0);  
}