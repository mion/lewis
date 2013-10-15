var editor = ace.edit("editor");
editor.setTheme("ace/theme/monokai");
editor.getSession().setMode("ace/mode/javascript");
editor.setFontSize(16);
editor.getSession().setTabSize(2);
editor.getSession().setUseSoftTabs(true);

var terminal = ace.edit("terminal");
terminal.setTheme("ace/theme/ambiance");
terminal.renderer.setShowGutter(false);
terminal.setHighlightActiveLine(false);
terminal.setReadOnly(true);
terminal.getSession().setMode("ace/mode/plain_text");
terminal.setFontSize(16);

function print(str) {
  terminal.insert(str.toString() + "\n");
};

function naviPlay() {
  eval(editor.getValue());
};

document.getElementById("btn-play").onclick = naviPlay;
