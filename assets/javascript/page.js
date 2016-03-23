ToggleEnable=function(sender,name) {
  if (devdept.dom.hasClass(sender,'disabled')){
    sendEnable(name,true)
    devdept.dom.removeClass(sender,'disabled')
  }else{
    sendEnable(name,false)
    devdept.dom.addClass(sender,'disabled')
  }
}

sendEnable=function(name, value) {
  var form = new FormData();
  form.append('name', name);
  form.append('value',value);
  var req = new devdept.net.XhrIo("/enable","POST",form);
  req.onComplete = function(g) {
    g.ResponseText
  }
  req.send();
}


exe=function(name) {
  var form = new FormData();
  form.append('name', name);
  var req = new devdept.net.XhrIo("/execute","POST",form);
  req.onComplete = function(g) {
    console.log(g.ResponseText)
  }
  req.send();
}