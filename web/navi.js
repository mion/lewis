var editor = ace.edit("editor");
editor.setTheme("ace/theme/monokai");
editor.getSession().setMode("ace/mode/javascript");
editor.setFontSize(18);
editor.getSession().setTabSize(4);
editor.getSession().setUseSoftTabs(true);

var terminal = ace.edit("terminal");
terminal.setTheme("ace/theme/ambiance");
terminal.getSession().setMode("ace/mode/javascript");
terminal.setFontSize(18);
terminal.renderer.setShowGutter(false);
terminal.setHighlightActiveLine(false);
terminal.setReadOnly(true);

if(typeof(String.prototype.trim) === "undefined")
{
    String.prototype.trim = function() 
    {
        return String(this).replace(/^\s+|\s+$/g, '');
    };
}

// Restore code
editor.session.setValue($.jStorage.get("source", ""));

function printInput (s) { terminal.insert(s + "\n"); };
function printOutput (s) { terminal.insert("=> " + s + "\n"); };

function execute () {
  "use strict";
  var text = editor.session.getTextRange(editor.getSelectionRange());
  printInput(text.trim());
  try {
    var result = window.eval(text);
    printOutput(result ? result.toSource() : result);
  } catch(err) {
    if (err instanceof SyntaxError) {
      printOutput("[!] The syntax is incorrect.");
    } else {
      printOutput("[!] Error: " + err.message);
    }
  }
}

//easter egg: mion_rocks => play music
editor.commands.addCommand({
    name: 'execute',
    bindKey: {win: 'Ctrl-E',  mac: 'Command-E'},
    exec: execute,
    readOnly: true // false if this command should not apply in readOnly mode
});
editor.commands.addCommand({
    name: 'save',
    bindKey: {win: 'Ctrl-S',  mac: 'Command-S'},
    exec: function () {
      var sourceCode = editor.session.getValue();
      $.jStorage.set("source", sourceCode);
      console.log("Saved.");
    },
    readOnly: true // false if this command should not apply in readOnly mode
});
