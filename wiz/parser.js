exports.parse = function (text) {
	return parenthesize(tokenize(text));
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
