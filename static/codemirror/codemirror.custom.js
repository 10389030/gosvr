$(function(){
  $(".code_editor").each(function(){
    // init editor
    var cm_params = {
      lineNumbers: true,
      indentUnit: 4,
      lineWrapping: true,
      cursorBlinkRate: 0,
      viewportMargin: 20,
      theme: "solarized dark",
    }


    if (this.tagName === "TEXTAREA") {
      var editor = CodeMirror.fromTextArea(this, cm_params);
    } else {
      var editor = CodeMirror(this, cm_params);
    }

    $.g_editor = editor;
  });
});
