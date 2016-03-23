if (typeof devdept === 'undefined') {
  devdept = Object();
}
devdept.date = Object();


devdept.date.Now = function() {
  var datetime = new devdept.date.Date();
  return datetime;
};
devdept.date.Date = function() {
  this.datetime = new Date();
};

devdept.date.Date.prototype.parse = function(format) {
  var month_ = (this.datetime.getMonth() < 10)?'0'+ (this.datetime.getMonth()+1):this.datetime.getMonth()+1;
  var date_ = (this.datetime.getDate() < 10)?'0'+ this.datetime.getDate():this.datetime.getDate();
  var hour_ = (this.datetime.getHours() < 10)?'0'+ this.datetime.getHours():this.datetime.getHours();
  var minute_ = (this.datetime.getMinutes() < 10)?'0'+ this.datetime.getMinutes():this.datetime.getMinutes();
  format = format.replace('YYYY',this.datetime.getFullYear());
  format = format.replace('mm',month_);
  format = format.replace('dd',date_);
  format = format.replace('HH',hour_);
  format = format.replace('ii',minute_);
  return format;
}





