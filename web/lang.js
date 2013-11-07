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
terminal.getSession().setMode("ace/mode/lisp");
terminal.setFontSize(16);

function output(s) {
  terminal.setValue(s);
};

;(function(exports) {
  var Env = function(scope, parent) {
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

  var global = new Env({});

  var interpret = function(str) {
    return compute(parse(str), global);
  };

  var isString = function(x) {
    return typeof x === "string"; 
  }

  var isList = function(x) {
    return x instanceof Array; // TODO: use typeof instead.
  }

  var debug = function(msg) {
    //console.log(msg);
    return;
  }

  // TODO: check list size, etc.
  var compute = function(x, env) {
    debug("compute, x = " + x.toSource() + ", env = " + env.toSource());
    if (isString(x)) {
      debug("-> Variable reference");
      return env.get(x); // variable reference
    } else if (!isList(x)) {
      debug("-> Constant literal");
      return x; // constant literal
    } else if ("quote" === x[0]) { // (quote exp)
      debug("-> quote form");
      return x[1];
    } else if ("if" === x[0]) { // (if test conseq alt)
      debug("-> if form");
      var test = x[1];
      var conseq = x[2];
      var alt = x[3];
      return compute(test, env) ? compute(conseq, env) : compute(alt, env);
    } else if ("def" === x[0]) { // (def var exp)
      debug("-> def form");
      var v = x[1];
      var e = x[2];
      env.scope[v] = compute(e, env);
      return undefined;
    } else if ("func" === x[0]) { // (func (var*) exp)
      debug("-> func form");
      var vars = x[1];
      var exp = x[2];
      var lambda = function() {
        var args = arguments;
        debug("args = " + args.toSource());
        var scope = vars.reduce(function(acc, v, i) {
          acc[v] = args[i];
          return acc;
        }, {});

        return compute(exp, new Env(scope, env));
      };
      return lambda;
    } else if ("do" === x[0]) {
      debug("-> do form");
      var exps = x.slice(1);
      var val = undefined;
      for (var i = 0; i < exps.length; i++) {
        val = compute(exps[i], env);
      };
      return val;
    } else {
      debug("-> proc");
      var list = x.map(function(exp) { return compute(exp, env); });
      if (list[0] instanceof Function) {
        var proc = list[0];
        var args = list.slice(1);
        return proc.apply(undefined, args);
      } else {
        return list;
      }
    }
  };

  var atom = function(token) {
    if (!isNaN(parseFloat(token))) {
      return parseFloat(token);
    } else if (token === "true") {
      return true;
    } else if (token === "false") {
      return false;
    } else {
      return token;
    }
  };

  // TODO: throw error on incorrect input!
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
        return parenthesize(input, list.concat(atom(token)));
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

  exports.lisp = {};
  exports.lisp.interpret = interpret;
})(typeof exports === 'undefined' ? this : exports);

function naviPlay() {
  lisp.interpret(editor.getValue());
};

document.getElementById("btn-play").onclick = naviPlay;
