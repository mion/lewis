var editor = ace.edit("editor");
editor.setTheme("ace/theme/ambiance");
editor.getSession().setMode("ace/mode/javascript");
editor.setFontSize(18);
editor.getSession().setTabSize(4);
editor.getSession().setUseSoftTabs(true);

var terminal = ace.edit("terminal");
terminal.setTheme("ace/theme/terminal");
terminal.getSession().setMode("ace/mode/javascript");
terminal.renderer.setShowGutter(false);
terminal.setHighlightActiveLine(false);
terminal.setReadOnly(true);
terminal.setFontSize(16);

if(typeof(String.prototype.trim) === "undefined")
{
    String.prototype.trim = function() 
    {
        return String(this).replace(/^\s+|\s+$/g, '');
    };
}

function printInput (s) { if (s) {s.trim();} terminal.insert(s + "\n"); };
function printOutput (s) { if (s) {s.trim();} terminal.insert("=> " + s + "\n"); };

function execute () {
  "use strict";
  var text = editor.session.getTextRange(editor.getSelectionRange());
  printInput(text);
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

document.getElementById("btn-eval").onclick = execute;
