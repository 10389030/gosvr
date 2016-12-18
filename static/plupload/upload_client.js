// 对Date的扩展，将 Date 转化为指定格式的String   
// 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符，   
// 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字)   
// 例子：   
// (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423   
// (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18   
Date.prototype.Format = function(fmt)   
{ //author: meizz   
  var o = {   
    "M+" : this.getMonth()+1,                 //月份   
    "d+" : this.getDate(),                    //日   
    "h+" : this.getHours(),                   //小时   
    "m+" : this.getMinutes(),                 //分   
    "s+" : this.getSeconds(),                 //秒   
    "q+" : Math.floor((this.getMonth()+3)/3), //季度   
    "S"  : this.getMilliseconds()             //毫秒   
  };   
  if(/(y+)/.test(fmt))   
    fmt=fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length));   
  for(var k in o)   
    if(new RegExp("("+ k +")").test(fmt))   
  fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));   
  return fmt;   
}  

$(function(){
  var domain = 'http://ogvugz6yc.bkt.clouddn.com';

  var uploader = Qiniu.uploader({
    runtimes: 'html5,html4',      // 上传模式，依次退化
    browse_button: 'select_file',         // 上传选择的点选按钮，必需
    uptoken_url: '/qiniu_uptoken',         // Ajax请求uptoken的Url，强烈建议设置（服务端提供）
    get_new_uptoken: true,             // 设置上传文件的时候是否每次都重新获取新的uptoken
    // downtoken_url: '/downtoken',
    // Ajax请求downToken的Url，私有空间时使用，JS-SDK将向该地址POST文件的key和domain，服务端返回的JSON必须包含url字段，url值为该文件的下载地址
    unique_names: true,              // 默认false，key为文件名。若开启该选项，JS-SDK会为每个文件自动生成key（文件名）
    save_key: true,                  // 默认false。若在服务端生成uptoken的上传策略中指定了sava_key，则开启，SDK在前端将不对key进行任何处理
    domain: domain,                  // bucket域名，下载资源时用到，必需
    max_file_size: '100mb',             // 最大文件体积限制
    max_retries: 1,                     // 上传失败最大重试次数
    dragdrop: true,                     // 开启可拖曳上传
    drop_element: 'upload',          // 拖曳上传区域元素的ID，拖曳文件或文件夹后可触发上传
    chunk_size: '4mb',                  // 分块上传时，每块的体积
    auto_start: true,                   // 选择文件后自动上传，若关闭需要自己绑定事件触发上传
    init: {
      'FilesAdded': function(up, files) {
        plupload.each(files, function(file) {
          // 文件添加进队列后，处理相关的事情
        });
      },
      'BeforeUpload': function(up, file) {
        // 每个文件上传前，处理相关的事情
      },
      'UploadProgress': function(up, file) {
        // 每个文件上传时，处理相关的事情
      },
      'FileUploaded': function(up, file, info) {
        // 每个文件上传成功后，处理相关的事情
        // 其中info是文件上传成功后，服务端返回的json，形式如：
        // {
        //    "hash": "Fh8xVqod2MQ1mocfI4S4KpRL6D98",
        //    "key": "gogopher.jpg"
        //  }
        // 查看简单反馈
        // var domain = up.getOption('domain');
        var res = JSON.parse(info);
        $.ajax({
          type: "POST",
          url: "/upload_save",
          dataType: "json",
          data: JSON.stringify(res),
        }).then(function(rsp){
          console.log(res, rsp);
          RenderAssets(res);
        });
      },
      'Error': function(up, err, errTip) {
        //上传出错时，处理相关的事情
      },
      'UploadComplete': function() {
        //队列文件处理完毕后，处理相关的事情
        RenderAssetsList();
      },
      // 'Key': function(up, file) {
      // 若想在前端对每个文件的key进行个性化处理，可以配置该函数
      // 该配置必须要在unique_names: false，save_key: false时才生效

      // var key = "";
      // do something with key here
      // return key
      // }
    }
  });

  $(".col_left").on("click", ".trash", function(){
    var id = $(this).data("id"); 
    console.log("trash id:", id);

    $.ajax({
      type: "DELETE",
      url: "/upload_delete",
      dataType: "json",
      data: JSON.stringify({"_id": id}),
    }).then(function(rsp){
      RenderAssetsList();
    });
  });

  $(".col_left").on("click", ".link", function(){
    var id = $(this).data("id"); 
    var txt = $(this).data("text");

    // should manage gloabl variables
    $.g_editor.replaceSelection("\n" + txt + "\n");
  });

  $(".col_left").on("click", ".edit", function(){
    var id = $(this).data("id"); 
    console.log("edit id:", id);

    $.ajax({
      type: "DELETE",
      url: "/update",
      dataType: "json",
      data: JSON.stringify({"_id": id}),
    }).then(function(rsp){
      RenderAssetsList();
    });
  });

  var RenderAssetsList = function() {
    $.get("/assets_list").then(function(rsp){
      $(".col_left .item:not(#upload)").remove();
      assets_list = rsp.msg.assets_list || [];
      for (var info of assets_list) {
        RenderAssets(info);
      } 
    }).fail(function(){
      alert("Fail to update assest list");
    });
  }
  RenderAssetsList();

  var RenderAssets = function(info) {
    console.log(info);

    var fname = info.fname;   
    var fsize = (info.fsize / 1024.0) + "MB";
    var mimeType = info.mimeType;
    var url = domain + "/" + info.key;
    var thumbnail = url + "-160.160"
    var createTs = (new Date(info.createTs * 1000)).Format("yyyy-MM-dd hh:mm:ss");
    console.log(createTs);

    var tmp = `
      <div class="item">
        <div class="thumbnail"><img src='${thumbnail}'/></div>
        <div class='item_info'>
          <p class='item_name'>${fname}</p>
          <p class='item_mime'>${mimeType}</p>
          <p class='item_ts'>${createTs}</p>
        </div>
        <div class='item_operation'>
          <button class="action trash" title="删除"><i class="fa fa-trash" aria-hidden="true"></i></button>
          <button class="action edit" title="编辑"><i class="fa fa-floppy-o" aria-hidden="true"></i></button>
          <button class='action link' title="链接"><i class="fa fa-files-o" aria-hidden="true"></i></button>
        </div>
      </div>
      `;
    var node = $(tmp);
    node.find(".action").data("id", info._id);
    var txt = "![" + fname + "](" + url + ")";
    console.log(txt);

    node.find(".action.link").data("text", txt);

    console.log(node.find(".action.link").data("text"));
    $("#upload").after(node);
  }
});


// domain为七牛空间对应的域名，选择某个空间后，可通过 空间设置->基本设置->域名设置 查看获取
// uploader为一个plupload对象，继承了所有plupload的方法  
