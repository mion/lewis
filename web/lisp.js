var editor = ace.edit("editor");
editor.setTheme("ace/theme/monokai");
editor.getSession().setMode("ace/mode/lisp");
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

// LittleLisp
;(function(exports) {
  var library = {
    first: function(x) {
      return x[0];
    },
    rest: function(x) {
      return x.slice(1);
    },
    print: function(x) {
      console.log(x);
      return x;
    }
  };

  var Context = function(scope, parent) {
    this.scope = scope;
    this.parent = parent;

    this.get = function(identifier) {
      if (identifier in this.scope) {
        return this.scope[identifier];
      } else if (this.parent !== undefined) {
        return this.parent.get(identifier);
      }
    };
  };

  var special = {
    quote: function(input, context) {
      return input[1];
    },

    '+': function(input, context) {
      console.log('form: +');
      console.log(input.slice(1));
      t = input.slice(1).reduce(function(acc, x) { return acc + interpret(x, context); });
      console.log(t);
      return interpret(t, context);
    },

    let: function(input, context) {
      var letContext = input[1].reduce(function(acc, x) {
        acc.scope[x[0].value] = interpret(x[1], context);
        return acc;
      }, new Context({}, context));

      return interpret(input[2], letContext);
    },

    lambda: function(input, context) {
      return function() {
        var lambdaArguments = arguments;
        var lambdaScope = input[1].reduce(function(acc, x, i) {
          acc[x.value] = lambdaArguments[i];
          return acc;
        }, {});

        return interpret(input[2], new Context(lambdaScope, context));
      };
    },

    if: function(input, context) {
      return interpret(input[1], context) ?
        interpret(input[2], context) :
        interpret(input[3], context);
    }
  };

  var interpretList = function(input, context) {
    if (input[0].value in special) {
      return special[input[0].value](input, context);
    } else {
      var list = input.map(function(x) { return interpret(x, context); });
      if (list[0] instanceof Function) {
        return list[0].apply(undefined, list.slice(1));
      } else {
        return list;
      }
    }
  };

  var interpret = function(input, context) {
    if (context === undefined) {
      return interpret(input, new Context(library));
    } else if (input instanceof Array) {
      return interpretList(input, context);
    } else if (input.type === "identifier") {
      return context.get(input.value);
    } else { // literal
      return input.value;
    }
  };

  var categorize = function(input) {
    if (!isNaN(parseFloat(input))) {
      return { type:'literal', value: parseFloat(input) };
    } else if (input[0] === '"' && input.slice(-1) === '"') {
      return { type:'literal', value: input.slice(1, -1) };
    } else {
      return { type:'identifier', value: input };
    }
  };

  var parenthesize = function(input, list) {
    if (list === undefined) {
      return parenthesize(input, []);
    } else {
      var token = input.shift();
      if (token === undefined) {
        return list.pop();
      } else if (token === "(") {
        list.push(parenthesize(input, []));
        return parenthesize(input, list);
      } else if (token === ")") {
        return list;
      } else {
        return parenthesize(input, list.concat(categorize(token)));
      }
    }
  };

  var tokenize = function(input) {
    return input.split('"')
                .map(function(x, i) {
                   if (i % 2 === 0) { // not in string
                     return x.replace(/\(/g, ' ( ')
                             .replace(/\)/g, ' ) ');
                   } else { // in string
                     return x.replace(/ /g, "!whitespace!");
                   }
                 })
                .join('"')
                .trim()
                .split(/\s+/)
                .map(function(x) {
                  return x.replace(/!whitespace!/g, " ");
                });
  };

  var parse = function(input) {
    return parenthesize(tokenize(input));
  };

  exports.littleLisp = {};
  exports.littleLisp.parse = parse;
  exports.littleLisp.interpret = interpret;
})(typeof exports === 'undefined' ? this : exports);

// Lewis
function lewisEvaluate() {
  input = editor.getValue()
  // parsedInput = littleLisp.parse(input)
  // console.log(parsedInput);
  // output = littleLisp.interpret(parsedInput);
  // console.log(output);

  out = littleLisp.interpret(littleLisp.parse(input));
  d = terminal.getSession().getDocument();
  if (out === undefined) {
    d.setValue('undefined');
  } else {
    d.setValue(out.toString());
  }
}

document.getElementById("btn-eval").onclick = lewisEvaluate;
