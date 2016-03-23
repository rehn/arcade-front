if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.dom = Object();


devdept.dom.getElement = function(id) {
  var el = document.getElementById(id);
  return el;
};

devdept.dom.createElement = function(tag,attributes,content) {
  var dom = document.createElement(tag);
  if (typeof attributes == 'object') {
    devdept.object.each(attributes,function(value,key){
      dom.setAttribute(key,value);
    });
  }
  if (content ) {
    devdept.dom.append(content,dom);
  }
  return dom;
};

devdept.dom.insertFirst = function(srcNode,targetNode) {
  if (targetNode.childNodes.length > 0 ) {
    targetNode.insertBefore(srcNode, targetNode.firstChild);
  }
  else {
    targetNode.appendChild(srcNode);
  }
}

devdept.dom.setContent = function(srcNode,targetNode) {
  devdept.dom.removeChildren(targetNode);
  if (typeof srcNode == 'string') {
    targetNode.appendChild(document.createTextNode(srcNode));
  }
  else if(Object.prototype.toString.call( srcNode ) === '[object Array]') {
    devdept.array.each(srcNode,function(value){
      devdept.dom.append(value,targetNode);
    });
  }
  else if (srcNode) {
    targetNode.appendChild(srcNode);
  }
  else{
    console.info("unknown type");
  }
}


devdept.dom.remove = function (domNode) {
  if (domNode.parentNode) {
    domNode.parentNode.removeChild( domNode );
  }
}

devdept.dom.append = function(srcNode,targetNode) {
  if (typeof srcNode == 'string') {
    targetNode.appendChild(document.createTextNode(srcNode));
  }
  else if(Object.prototype.toString.call( srcNode ) === '[object Array]') {
    devdept.array.each(srcNode,function(value){
      devdept.dom.append(value,targetNode);
    });
  }
  else if (srcNode) {
    targetNode.appendChild(srcNode);
  }
  else{
    console.info("unknown type");
  }
}

devdept.dom.removeChildren = function(element) {
  while (element.firstChild) {
    element.removeChild(element.firstChild);
  }
}
devdept.dom.getParent = function(element) {
  return element.parentNode;
}



devdept.dom.getClasses = function(element) {
  var classes = element.getAttribute('class').split(" ");
  return classes;
};


devdept.dom.swapClass = function(element,oldcls,newcls) {
 
 element.classList.add(newcls);
 element.classList.remove(oldcls);
};

devdept.dom.addClass = function(element,cls) {
 element.classList.add(cls);
};

devdept.dom.removeClass = function(element,cls) {
  element.classList.remove(cls);
};

devdept.dom.hasClass = function(element,cls) {
  return element.classList.contains(cls);
};

devdept.dom.toggleClass = function(element,cls) {
  if (element.classList.contains(cls)) {
    devdept.dom.removeClass(element,cls);
  }else{
    devdept.dom.addClass(element,cls);
  }
};


devdept.dom.handleClass = function(element,cls,add) {
  if (add) {
    element.classList.add(cls);
  }
  else {
    element.classList.remove(cls);
  }
}
devdept.dom.listen = function(obj,key,fn,scope) {
  if(scope){
    fn = fn.bind(scope);
  }
  obj.addEventListener(key,fn);
}