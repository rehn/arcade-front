if (!devdept) {
  devdept = Object();
}


devdept.TemplateFromFile = function(filename,data,target) {
  var xhr = new XMLHttpRequest();
  xhr.open('GET', filename, false);
  xhr.send();
  template = xhr.responseText;
  return new devdept.Template(template,data,target);
};

devdept.Template = function(template,data,target) {
    this.build(template,data,target);  
}

devdept.Template.prototype = { 
  build : function(template,data,target) {
    this.regexp = /{{([^}}]+)?}}/g;
    this.data = data;
    
    this.monitoredNodes = Object();
  

    var tmp = document.createElement('div');
    tmp.innerHTML = template;
    this.domNodes = tmp.childNodes;
    
    if (this.data !=null && this.data != false) {
      
      devdept.object.each(this.data,function(val,key){
        this.monitoredNodes[key]=[];
      },this);

      Object.defineProperty(this.data, "watch", {
        enumerable : false,
        configurable : true,
        writable : false,
        value : function (handler, scope) {
          devdept.object.each(this,function(value,key) {
            var oldval = value,
            newval = oldval,
            getter = function () {
              print
              return newval;
            },
            setter = function (val) {
              newval = val;
              handler.call(this, key, val, scope);
              return newval;
            };
            Object.defineProperty(this, key, {
               get: getter,
               set: setter,
               enumerable: false,
               configurable: true
            });
          },this);
        }
      });
      devdept.array.each(this.domNodes,this.walkthru,this);
      this.data.watch(this.update,this);
    }
    if (target) {
      this.render(target);
    }
  },
  
  update : function(key,newValue,scope) { 
    devdept.array.each(scope.monitoredNodes[key],function(item) {
      var node = item['target'];
      var orginal = item['orginal'];
      var newstr=orginal;
      while(match = scope.regexp.exec(orginal)) {
        var newData = "";
        //if (scope.data[match[1]]) {
          newData = scope.data[match[1]];
        //}
        newstr = newstr.replace(match[0],newData);
      }

      if (node.nodeType == 2) {
        node.value = newstr; 
      }
      else if (node.nodeType == 3) {                
        node.nodeValue = newstr;
      }
    });
  },

  parse : function(node) {
    if (node.nodeType == 1) {
      devdept.object.each(node.attributes,function(attribute) {
        var orginal = attribute.value;

        while (match = this.regexp.exec(attribute.value)) {
          var newData = "";
          //if (this.data[match[1]]) {
            newData = this.data[match[1]];
          //}

          attribute.value = attribute.value.replace(match[0],newData);        
          this.monitoredNodes[match[1]].push({'orginal':orginal,'target':attribute});
        }

      },this);
    }

    else if (node.nodeType == 3) {
      
      var orginal = node.nodeValue;
      
      while(match = this.regexp.exec(node.nodeValue)) {
        var newData = "";
        if (this.data[match[1]]) {
          newData = this.data[match[1]];
        }
        node.nodeValue = node.nodeValue.replace(match[0],newData);        
              
        this.monitoredNodes[match[1]].push({'orginal':orginal,'target':node});
      }
   }
  },
  
  walkthru : function(node,index) {
    this.parse(node);  
    if (node.nodeType == 1 && node.childNodes.length > 0) {
      devdept.array.each(node.childNodes,this.walkthru,this);
    }
  },

  render : function(target) {

    var a = [];
    while(this.domNodes.length) {
      a.push(target.appendChild(this.domNodes[0]));  
    } 
    this.domNodes=a;
  },

  remove : function() {
    if(!this.domNodes) return;
    for (var i = 0; i < this.domNodes.length; i++) {
      this.domNodes[i].parentElement.removeChild(this.domNodes[i]);
    }
      this.data = null;
      this.monitoredNodes = null;
      this.domNodes = null;
  },
  nodeByKey : function(key) {
    var nodes = Array();
    if (!this.monitoredNodes[key]) {
      return null;
    }
    devdept.object.each(this.monitoredNodes[key],function(node) {
      switch (node['target'].nodeType) {
        case 2:
          nodes.push(node['target'].ownerElement);
          break;
        case 3:
          nodes.push(node['target'].parentNode);
          break;
      }
    },this);
    if (nodes.length > 0 ){
      return nodes[0];
    }
    return null;
  }
};





