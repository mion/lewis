var editor = ace.edit("editor");
editor.setTheme("ace/theme/monokai");
editor.getSession().setMode("ace/mode/javascript");
editor.setFontSize(16);
editor.getSession().setTabSize(4);
editor.getSession().setUseSoftTabs(true);

var terminal = ace.edit("terminal");
terminal.setTheme("ace/theme/ambiance");
terminal.getSession().setMode("ace/mode/javascript");
//terminal.renderer.setShowGutter(false);
terminal.setHighlightActiveLine(false);
terminal.setReadOnly(true);
terminal.setFontSize(16);

// easter egg: mion_rocks => play music
editor.commands.addCommand({
    name: 'eval',
    bindKey: {win: 'Ctrl-E',  mac: 'Command-E'},
    exec: function (editor) {
      var text = editor.session.getTextRange(editor.getSelectionRange());
      terminal.insert(">> " + text + "\n");
      try {
        var result = eval(text + ";");
        terminal.insert("=> " + eval(text) + "\n");
      } catch(err) {
        if (err instanceof SyntaxError) {
          terminal.insert("[!] The syntax is incorrect.");
        }
      }
    },
    readOnly: true // false if this command should not apply in readOnly mode
});

function output (s) {
  terminal.setValue(s);
};

function naviPlay () {
  lisp.interpret(editor.getValue());
};

